package main

import (
	"dzug/app/publish/infra/redis"
	"dzug/app/publish/pkg/oss"
	"dzug/app/publish/service"
	"dzug/conf"
	"dzug/discovery"
	"dzug/logger"
	pb "dzug/protos/publish"
	"dzug/repo"
	"fmt"
	"go.uber.org/zap"
)

func main() {
	if err := conf.Init(); err != nil {
		fmt.Printf("Config file initialization error,%#v", err)
		return
	}

	if err := logger.Init(conf.Config.LogConfig, conf.Config.Mode); err != nil {
		fmt.Printf("log file initialization error,%#v", err)
		return
	}

	defer zap.L().Sync()
	zap.L().Info("服务启动，开始记录日志")

	err := redis.Init()
	if err != nil {
		return
	}
	oss.Init()
	_ = repo.Init()
	key := "publish"
	value := "127.0.0.1:9003"
	serviceRegister, grpcServer := discovery.InitRegister(key, value)
	defer serviceRegister.Close()
	defer grpcServer.Stop()
	pb.RegisterPublishServiceServer(grpcServer, &service.VideoServer{})
	discovery.GrpcListen(grpcServer, value)
}
