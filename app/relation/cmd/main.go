package main

import (
	"dzug/app/relation/service"
	"dzug/discovery"
	pb "dzug/idl/relation"
)

func main() {
	key := "relation"         // 注册的名字
	value := "127.0.0.1:9001" // 注册的地址
	// 传入注册的服务名和注册的服务地址进行注册
	serviceRegister, grpcServer := discovery.InitRegister(key, value)
	defer serviceRegister.Close()
	defer grpcServer.Stop()
	pb.RegisterDouyinRelationActionServiceServer(grpcServer, &service.RelationSrv{}) // 绑定grpc
	discovery.GrpcListen(grpcServer, value)                                          // 开启监听
}
