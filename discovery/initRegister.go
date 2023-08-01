package discovery

import (
	"dzug/conf"
	"fmt"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"net"
)

// InitRegister 放入服务名称和服务的链接，返回etcdServiceRegister和grpc服务器
func InitRegister(key, value string) (*ServiceRegister, *grpc.Server) {
	fmt.Println(conf.Config.EtcdConfig.Addr)
	endpoints := conf.Config.EtcdConfig.Addr
	serviceRegister := &ServiceRegister{
		EtcdAddrs: endpoints,
		Lease:     5,
		Key:       key,
		Value:     value,
	}
	err := serviceRegister.NewServiceRegister()
	if err != nil {
		zap.L().Error("初始化服务注册失败" + err.Error())
	}

	server := grpc.NewServer()
	return serviceRegister, server
}

// GrpcListen 监听grpc链接
func GrpcListen(server *grpc.Server, value string) {
	lis, err := net.Listen("tcp", value)
	if err != nil {
		zap.L().Error("启动监听失败")
	}
	if err = server.Serve(lis); err != nil {
		zap.L().Error("连接grpc服务失败" + err.Error())
		return
	}
}
