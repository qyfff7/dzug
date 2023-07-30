package discovery

import (
	"fmt"
	clientv3 "go.etcd.io/etcd/client/v3"
	"log"
	"time"
)

// ServiceRegister 注册服务到etcd上
type ServiceRegister struct {
	EtcdAddrs []string // etcd集群列表
	Lease     int64    // 服务的租约时间TTL
	Key       string   // 服务名称
	Value     string   // 服务地址

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
		log.Fatal(err)
		return err
	}

	s.cli = cli

	// 启动服务
	if err := s.putKeyWithLease(); err != nil {
		log.Fatal(err)
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
		log.Fatal(err)
		return err
	}
	// 注册服务并绑定租约
	_, err = s.cli.Put(s.cli.Ctx(), s.Key, s.Value, clientv3.WithLease(resp.ID))
	if err != nil {
		log.Fatal(err)
		return err
	}
	s.keepAliveChan, err = s.cli.KeepAlive(s.cli.Ctx(), resp.ID) // 保持租约活跃
	if err != nil {
		log.Fatal(err)
		return err
	}
	log.Println("租约ID为：", s.leaseID)
	log.Printf("put key:%s val:%s success\n", s.Key, s.Value)
	return nil
}

// keepAlive 保持etcd连接
func (s *ServiceRegister) keepAlive() {
	for {
		select {
		case keepAliveResp := <-s.keepAliveChan:
			// 收到续租响应
			if keepAliveResp == nil {
				fmt.Println("续租失败")
				return
			}
			// 处理续租响应
			fmt.Println(s.Key, "：收到续租响应，续租成功")
		case <-time.After(5 * time.Second):
			// 定时执行续租操作
			if err := s.putKeyWithLease(); err != nil {
				fmt.Println("续租失败: ", err)
				return
			}
		}
	}
}

// Close 关闭租约
func (s *ServiceRegister) Close() error {
	// 撤销租约
	if _, err := s.cli.Revoke(s.cli.Ctx(), s.leaseID); err != nil { // Revoke撤销租约
		return err
	}
	log.Println("撤销租约")
	return s.cli.Close()
}
