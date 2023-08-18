package discovery

import (
	"context"
	"errors"
	"go.etcd.io/etcd/api/v3/mvccpb"
	clientv3 "go.etcd.io/etcd/client/v3"
	"go.uber.org/zap"
	"math/rand"
	"sync"
	"time"
)

// serviceDiscovery 用于服务发现
type serviceDiscovery struct {
	EtcdAddrs []string

	cli        *clientv3.Client
	serverList map[string][]string // 存储解析后的地址
	lock       sync.Mutex
}

// NewServiceDiscovery 新建发现结构体
func (s *serviceDiscovery) newServiceDiscovery() (err error) {
	s.cli, err = clientv3.New(clientv3.Config{
		Endpoints:   s.EtcdAddrs,
		DialTimeout: 3 * time.Second,
	})
	if err != nil {
		zap.L().Error("建立服务发现失败: " + err.Error() + "（请检查etcd是否启动，并且已注册有服务）")
		return
	}
	s.serverList = make(map[string][]string)
	return
}

func (s *serviceDiscovery) watchService(target string) error {
	// 建立连接超时，3秒未连接上，超时
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(3)*time.Second)
	defer cancel()
	// 获取target的所有键值对，即所有服务地址
	resp, err := s.cli.Get(ctx, target, clientv3.WithPrefix())
	if err != nil {
		zap.L().Error("获取服务列表失败：", zap.Error(err))
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
func (s *serviceDiscovery) setServiceList(key, value string) {
	s.lock.Lock()
	defer s.lock.Unlock()
	if !existValue(s.serverList[value], key) { // 如果没有了这个地址 ！！！反着放的
		s.serverList[value] = append(s.serverList[value], key)
		zap.L().Debug("put key :" + key + " val:" + value)
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
func (s *serviceDiscovery) delServiceList(key string) {
	s.lock.Lock()
	defer s.lock.Unlock()
	delete(s.serverList, key)
	zap.L().Debug("delete key: " + key)
}

// watcher 监视服务列表
func (s *serviceDiscovery) watcher(target string) {
	watchChan := s.cli.Watch(context.Background(), target, clientv3.WithPrefix())
	zap.L().Debug("Watching target: " + target + "...")
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

// getServices 获取服务中所有的服务
func (s *serviceDiscovery) getServices() map[string][]string {
	s.lock.Lock()
	defer s.lock.Unlock()
	return s.serverList
}

// getServiceByKey 通过key 获取服务链接
func (s *serviceDiscovery) getServiceByKey(target string) (value string, err error) {
	rand.Seed(time.Now().UnixNano())
	if len(s.serverList[target]) == 0 {
		zap.L().Error("服务列表中无该服务")
		return "", errors.New("服务列表中无该服务")
	}
	// 生成随机整数
	randomNum := rand.Intn(len(s.serverList[target])) // target 下随机选一个链接进行调用，负载均衡 /:fade
	zap.L().Debug(target + " 调用的链接为：" + s.serverList[target][randomNum])
	return s.serverList[target][randomNum], nil
}

func (s *serviceDiscovery) Close() error {
	return s.cli.Close()
}
