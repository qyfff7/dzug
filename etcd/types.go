package etcd

import (
	"context"
	"io"
)

type Registry interface {
	Register(ctx context.Context, si ServiceInstance) error
	UnRegister(ctx context.Context, instance ServiceInstance) error

	io.Closer
}

type ServiceInstance struct {
	Name    string
	Address string
}

type Event struct {
}
