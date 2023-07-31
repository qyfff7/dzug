package discovery

import (
	"context"
	"fmt"
	"go.etcd.io/etcd/api/v3/mvccpb"
	clientv3 "go.etcd.io/etcd/client/v3"
	"log"
	"math/rand"
	"sync"
	"time"
)

// ServiceDiscovery 用于服务发现
type ServiceDiscovery struct {
	EtcdAddrs []string

	cli        *clientv3.Client
	serverList map[string][]string // 存储解析后的地址
	lock       sync.Mutex
}

// NewServiceDiscovery 新建发现结构体
func (s *ServiceDiscovery) NewServiceDiscovery() (err error) {
	s.cli, err = clientv3.New(clientv3.Config{
		Endpoints:   s.EtcdAddrs,
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		log.Println(err)
	}
	s.serverList = make(map[string][]string)
	return
}

func (s *ServiceDiscovery) watchService(target string) error {
	// 获取target的所有键值对
	resp, err := s.cli.Get(context.Background(), target, clientv3.WithPrefix())
	if err != nil {
		log.Println(err)
		return err
	}

	// 将键值对放入新增服务地址中
	for _, ev := range resp.Kvs {
		s.setServiceList(string(ev.Key), string(ev.Value))
	}

	// 启动协程持续性监听
	go s.watcher(target)
	return nil
}

// setServiceList 设置地址列表
func (s *ServiceDiscovery) setServiceList(key, value string) {
	s.lock.Lock()
	defer s.lock.Unlock()
	if !existValue(s.serverList[value], key) { // 如果没有了这个地址 ！！！反着放的
		s.serverList[value] = append(s.serverList[value], key)
		log.Println("put key :", key, " val:", value)
	}
}

func existValue(strs []string, str string) bool {
	for _, s := range strs {
		if s == str {
			return true
		}
	}
	return false
}

// delServiceList 从列表中删除服务
func (s *ServiceDiscovery) delServiceList(key string) {
	s.lock.Lock()
	defer s.lock.Unlock()
	delete(s.serverList, key)
	log.Println("delete key: ", key)
}

// watcher 监视服务列表
func (s *ServiceDiscovery) watcher(target string) {
	watchChan := s.cli.Watch(context.Background(), target, clientv3.WithPrefix())
	log.Println("Watching target: ", target, "...")
	for w := range watchChan {
		for _, ev := range w.Events {
			switch ev.Type {
			case mvccpb.PUT: //修改或者新增
				s.setServiceList(string(ev.Kv.Key), string(ev.Kv.Value))
			case mvccpb.DELETE: //删除
				s.delServiceList(string(ev.Kv.Key))
			}
		}
	}
}

// GetServices 获取服务中所有的服务
func (s *ServiceDiscovery) GetServices() map[string][]string {
	s.lock.Lock()
	defer s.lock.Unlock()
	return s.serverList
}

// GetServiceByKey 通过key 获取服务链接
func (s *ServiceDiscovery) GetServiceByKey(target string) (value string) {
	rand.Seed(time.Now().UnixNano())
	// 生成随机整数
	randomNum := rand.Intn(len(s.serverList[target])) // target 下随机选一个链接进行调用，负载均衡 /:fade
	fmt.Println(target, " 调用的链接为：", s.serverList[target][randomNum])
	return s.serverList[target][randomNum]
}

func (s *ServiceDiscovery) Close() error {
	return s.cli.Close()
}
