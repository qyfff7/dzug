package relationservice

import (
	"dzug/app/services/relation/service"
	"dzug/conf"
	"dzug/discovery"
	pb "dzug/protos/relation"
)

func Start() {

	// 传入注册的服务名和注册的服务地址进行注册
	serviceRegister, grpcServer := discovery.InitRegister(conf.Config.RelationServiceName, conf.Config.RelationServiceUrl)
	defer serviceRegister.Close()
	defer grpcServer.Stop()
	pb.RegisterDouyinRelationActionServiceServer(grpcServer, &service.RelationSrv{}) // 绑定grpc
	discovery.GrpcListen(grpcServer, conf.Config.RelationServiceUrl)                 // 开启监听
}
