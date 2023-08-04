package main

import (
	"dzug/app/favorite/service"
	"dzug/conf"
	"dzug/discovery"
	"dzug/logger"
	pb "dzug/protos/favorite"
	"dzug/repo"
	"fmt"
	"go.uber.org/zap"
)

func main() {
	//1. 初始化配置文件
	if err := conf.Init(); err != nil {
		fmt.Printf("Config file initialization error,%#v", err)
		return
	}

	//2. 初始化日志
	if err := logger.Init(conf.Config.LogConfig, conf.Config.Mode); err != nil {
		fmt.Printf("log file initialization error,%#v", err)
		return
	}
	defer zap.L().Sync() //把缓冲区的日志，追加到文件中
	zap.L().Info("服务启动，开始记录日志")

	repo.Init() // 初始化数据库

	key := "favorite"         // 注册的名字
	value := "127.0.0.1:9003" // 注册的服务地址
	// 传入注册的服务名和注册的服务地址进行注册
	serviceRegister, grpcServer := discovery.InitRegister(key, value)
	defer serviceRegister.Close()
	defer grpcServer.Stop()
	pb.RegisterDouyinFavoriteActionServiceServer(grpcServer, &service.FavoriteSrv{}) // 绑定grpc
	discovery.GrpcListen(grpcServer, value)                                          // 开启监听
}
