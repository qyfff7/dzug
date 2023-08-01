package registry

import (
	"context"
	"dzug/conf"
	"dzug/etcd"
	clientv3 "go.etcd.io/etcd/client/v3"
	"go.etcd.io/etcd/client/v3/concurrency"
	"go.uber.org/zap"
)

type Registry struct {
	c    *clientv3.Client
	sess *concurrency.Session
}

func NewRegistry() (*Registry, error) {
	// 新建客户端与会话
	c, err := clientv3.New(clientv3.Config{
		Endpoints: conf.Config.EtcdConfig.Addr,
	})
	if err != nil {
		return nil, err
	}

	sess, err := concurrency.NewSession(c)
	if err != nil {
		return nil, err
	}
	return &Registry{
		c:    c,
		sess: sess,
	}, nil
}

func (r *Registry) Register(ctx context.Context, si etcd.ServiceInstance) error {
	_, err := r.c.Put(ctx, si.Name, si.Address, clientv3.WithLease(r.sess.Lease()))
	if err != nil {
		zap.L().Error(si.Name + "注册失败" + err.Error())
		return err
	}
	zap.L().Info("put key: " + si.Name + " val: " + si.Address + " success")
	return nil
}

func (r *Registry) UnRegister(ctx context.Context, si etcd.ServiceInstance) error {
	_, err := r.c.Delete(ctx, si.Name)
	return err
}

func (r *Registry) Close() error {
	// Session关闭代表租约结束，会自动解约
	err := r.sess.Close()
	return err
}
