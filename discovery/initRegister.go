package discovery

import (
	"dzug/app/services/config_center"
	"fmt"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"net"
)

// InitRegister 放入服务名称和服务的链接，返回etcdServiceRegister和grpc服务器
func InitRegister(key, value string) (*ServiceRegister, *grpc.Server) {
	endpoints := config_center.ProjBaseConf.EtcdAddr
	sRegister := &ServiceRegister{
		EtcdAddrs: endpoints,
		Lease:     5,
		Key:       key,
		Value:     value,
	}
	err := sRegister.newServiceRegister()
	if err != nil {
		zap.L().Fatal("初始化服务注册失败" + err.Error())
	}

	server := grpc.NewServer()
	return sRegister, server
}

// GrpcListen 监听grpc链接
func GrpcListen(server *grpc.Server, value string) {
	lis, err := net.Listen("tcp", value)
	if err != nil {
		zap.L().Fatal("启动监听失败" + err.Error())
	}
	fmt.Println("listening...")
	if err = server.Serve(lis); err != nil {
		zap.L().Fatal("连接grpc服务失败" + err.Error())
		return
	}
}
