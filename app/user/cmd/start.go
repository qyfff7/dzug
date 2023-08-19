package userservice

import (
	"dzug/app/user/service"
	"dzug/conf"
	"dzug/discovery"
	"dzug/logger"
	pb "dzug/protos/user"
	"fmt"
	"go.uber.org/zap"
)

func Start() {
	//1.初始化日志

	servicename := "user"
	if err := logger.Init(conf.LogConfList, servicename); err != nil {
		fmt.Printf("log file initialization error,%#v", err)
		return
	}
	defer zap.L().Sync() //把缓冲区的日志，追加到文件中
	zap.L().Info("user服务启动，开始记录日志")

	// 传入注册的服务名和注册的服务地址进行注册
	serviceRegister, grpcServer := discovery.InitRegister(conf.Config.UserServiceName, conf.Config.UserServiceUrl)
	defer serviceRegister.Close()
	defer grpcServer.Stop()
	pb.RegisterServiceServer(grpcServer, &service.Userservice{}) // 绑定grpc
	discovery.GrpcListen(grpcServer, conf.Config.UserServiceUrl) // 开启监听
}
