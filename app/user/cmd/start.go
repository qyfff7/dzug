package userservice

import (
	"dzug/app/user/service"
	"dzug/discovery"
	pb "dzug/protos/user"
)

func Start() {
	key := "user"             // 注册的名字
	value := "127.0.0.1:9000" // 注册的服务地址
	// 传入注册的服务名和注册的服务地址进行注册
	serviceRegister, grpcServer := discovery.InitRegister(key, value)
	defer serviceRegister.Close()
	defer grpcServer.Stop()
	pb.RegisterServiceServer(grpcServer, &service.Userservice{}) // 绑定grpc
	discovery.GrpcListen(grpcServer, value)                      // 开启监听
}
