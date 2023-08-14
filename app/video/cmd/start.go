package videoservice

import (
	"dzug/app/video/service"
	"dzug/discovery"
	"dzug/protos/video"
)

func Start() {

	key := "video"            // 注册的名字
	value := "127.0.0.1:9003" // 注册的服务地址
	// 传入注册的服务名和注册的服务地址进行注册
	serviceRegister, grpcServer := discovery.InitRegister(key, value)
	defer serviceRegister.Close()
	defer grpcServer.Stop()
	video.RegisterVideoServiceServer(grpcServer, &service.VideoService{}) // 绑定grpc
	discovery.GrpcListen(grpcServer, value)                               // 开启监听

}
