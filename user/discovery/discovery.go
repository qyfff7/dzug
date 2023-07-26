package discovery

import (
	clientv3 "go.etcd.io/etcd/client/v3"
	"log"
	"time"
)

// ServiceRegister 创建租约注册服务
type ServiceRegister struct {
	cli     *clientv3.Client // etcd client，用于与etcd通信
	leaseID clientv3.LeaseID // 租约ID
	// 租约keepalive相应chan
	keepAliveChan <-chan *clientv3.LeaseKeepAliveResponse // 通道，只接受那一种类型，用于接收续租响应
	key           string                                  // 服务注册路径
	val           string                                  // 服务注册值
}

// putKeyWithLease 注册服务并绑定租约
func (s *ServiceRegister) putKeyWithLease(lease int64) error {
	// 设置租约时间
	resp, err := s.cli.Grant(s.cli.Ctx(), lease) // 创建一个具有指定租约时间的租约，获得租约的响应
	if err != nil {
		return err
	}
	// 注册服务并绑定租约
	_, err = s.cli.Put(s.cli.Ctx(), s.key, s.val, clientv3.WithLease(resp.ID)) // 使用cli.Put将服务注册到etcd，并用withLease绑定到注册的键上
	if err != nil {
		return err
	}
	// 设置续租，定期发送需求请求
	leaseRespChan, err := s.cli.KeepAlive(s.cli.Ctx(), resp.ID) // 设置续租，定期发送请求
	if err != nil {
		return err
	}
	log.Println(s.leaseID)
	s.keepAliveChan = leaseRespChan // 将续租响应通道赋值给keepAliveChan
	log.Printf("put key:%s val:%s success\n", s.key, s.val)
	return nil
}

func NewServiceRegister(endpoints []string, key, val string, lease int64) (*ServiceRegister, error) {
	cli, err := clientv3.New(clientv3.Config{ // 与ecd通信的 clientv3.Client对象
		Endpoints:   endpoints,
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		log.Fatal(err)
	}

	ser := &ServiceRegister{ // 创建ServiceRegister对象
		cli: cli,
		key: key,
		val: val,
	}

	// 服务注册，并申请租约设置时间keepalive
	if err := ser.putKeyWithLease(lease); err != nil {
		return nil, err
	}
	return ser, nil
}

// ListenLeaseRespChan 监听续租相应chan
func (s *ServiceRegister) ListenLeaseRespChan() {
	for leaseKeepResp := range s.keepAliveChan {
		log.Printf("续租成功: %v\n", leaseKeepResp)
	}
	log.Println("关闭续租")
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
