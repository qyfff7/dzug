package registry

import (
	"context"
	"dzug/etcd"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"net"
	"time"
)

type Server struct {
	registry        *Registry
	registerTimeout time.Duration
}

func NewServer() (*Server, error) {
	return &Server{
		registerTimeout: 10 * time.Second,
	}, nil
}

func (s *Server) StartRegisterAndListen(key, val string) error {
	si := etcd.ServiceInstance{
		Name:    key,
		Address: val,
	}
	registry, err := NewRegistry()
	s.registry = registry
	if err != nil {
		zap.L().Error("启动注册中心失败" + err.Error())
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), s.registerTimeout)
	defer cancel()
	err = s.registry.Register(ctx, si)
	if err != nil {
		return err
	}
	_ = s.registry.Close()
	return nil
}

func (s *Server) StartGRPCListen(server *grpc.Server, addr string) {
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		zap.L().Error("启动监听失败")
	}
	if err = server.Serve(lis); err != nil {
		zap.L().Error("连接grpc服务失败" + err.Error())
		return
	}
}
