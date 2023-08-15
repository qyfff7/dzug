package videoservice

import (
	"dzug/app/video/service"
	"dzug/conf"
	"dzug/discovery"
	"dzug/protos/video"
)

func Start() {

	// 传入注册的服务名和注册的服务地址进行注册
	serviceRegister, grpcServer := discovery.InitRegister(conf.Config.VideoServiceName, conf.Config.VideoServiceUrl)
	defer serviceRegister.Close()
	defer grpcServer.Stop()
	video.RegisterVideoServiceServer(grpcServer, &service.VideoService{}) // 绑定grpc
	discovery.GrpcListen(grpcServer, conf.Config.VideoServiceUrl)         // 开启监听

}