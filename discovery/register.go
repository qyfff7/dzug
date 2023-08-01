package discovery

import (
	clientv3 "go.etcd.io/etcd/client/v3"
	"go.uber.org/zap"
	"time"
)

// ServiceRegister 注册服务到etcd上
type ServiceRegister struct {
	EtcdAddrs []string // etcd集群列表
	Lease     int64    // 服务的租约时间TTL
	Key       string   // 服务名称
	Value     string   // 服务地址，这个东西更改后，应该服务地址列表和代表的权值

	cli     *clientv3.Client // etcd client，用于与etcd通信
	leaseID clientv3.LeaseID // 租约ID
	// 租约keepalive相应chan
	keepAliveChan <-chan *clientv3.LeaseKeepAliveResponse // 通道，只接受那一种类型，用于接收续租响应
	// todo 暂时没有日志记录
}

// NewServiceRegister 通过传入的ServiceRegister 启动一个服务注册项目，并返回错误
func (s *ServiceRegister) NewServiceRegister() error {
	// 配置clientv3 服务器
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   s.EtcdAddrs,
		DialTimeout: 5 * time.Second, // 5s 连接超时
	})
	if err != nil {
		zap.L().Error("新建clientv3失败" + err.Error())
		return err
	}
	s.cli = cli

	// 启动服务
	if err = s.putKeyWithLease(); err != nil {
		zap.L().Info("启动服务失败")
		return err
	}
	go s.keepAlive() // 保持连接
	return nil
}

// putKeyWithLease 通过租约来启动一个etcd连接
func (s *ServiceRegister) putKeyWithLease() error {
	// 设置租约时间
	resp, err := s.cli.Grant(s.cli.Ctx(), s.Lease) // 得到租约ID
	if err != nil {
		zap.L().Error("设置租约失败" + err.Error())
		return err
	}
	// 注册服务并绑定租约
	_, err = s.cli.Put(s.cli.Ctx(), s.Value, s.Key, clientv3.WithLease(resp.ID)) // !!!key，value，故意填反的
	if err != nil {
		zap.L().Error("注册服务绑定租约失败" + err.Error())
		return err
	}
	// 保持租约活跃
	s.keepAliveChan, err = s.cli.KeepAlive(s.cli.Ctx(), resp.ID)
	if err != nil {
		zap.L().Error("租约保活失败" + err.Error())
		return err
	}
	zap.L().Debug("put key: " + s.Key + " val: " + s.Value + " success")
	return nil
}

// keepAlive 保持etcd连接
func (s *ServiceRegister) keepAlive() {
	for {
		select {
		case keepAliveResp := <-s.keepAliveChan:
			// 收到续租响应
			if keepAliveResp == nil {
				zap.L().Debug("续租失败")
				return
			}
			// 处理续租响应
			zap.L().Debug(s.Key + "收到续租响应，续租成功 " + time.Now().Format("2006-01-02 15:04:05"))
		case <-time.After(5 * time.Second):
			// 定时执行续租操作
			if err := s.putKeyWithLease(); err != nil {
				zap.L().Error("定时续租失败" + err.Error())
				return
			}
		}
	}
}

// Close 关闭租约
func (s *ServiceRegister) Close() error {
	// 撤销租约
	// todo 优雅关闭？
	if _, err := s.cli.Revoke(s.cli.Ctx(), s.leaseID); err != nil { // Revoke撤销租约
		zap.L().Error("撤销租约失败：" + err.Error())
		return err
	}
	zap.L().Debug(s.Key + "撤销租约")
	return s.cli.Close()
}
