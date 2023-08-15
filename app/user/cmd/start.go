package userservice

import (
	"dzug/app/user/service"
	"dzug/conf"
	"dzug/discovery"
	pb "dzug/protos/user"
)

func Start() {

	// 传入注册的服务名和注册的服务地址进行注册
	serviceRegister, grpcServer := discovery.InitRegister(conf.Config.UserServiceName, conf.Config.UserServiceUrl)
	defer serviceRegister.Close()
	defer grpcServer.Stop()
	pb.RegisterServiceServer(grpcServer, &service.Userservice{}) // 绑定grpc
	discovery.GrpcListen(grpcServer, conf.Config.UserServiceUrl) // 开启监听
}
