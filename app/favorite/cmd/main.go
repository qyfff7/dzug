package main

import (
	"dzug/app/favorite/dal/kafka"
	"dzug/app/favorite/dal/redis"
	"dzug/app/favorite/service"
	"dzug/conf"
	"dzug/discovery"
	"dzug/logger"
	pb "dzug/protos/favorite"
	"dzug/repo"
	"fmt"
	"go.uber.org/zap"
	"time"
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

	repo.Init()       // 初始化数据库
	redis.InitRedis() // 初始化redis
	kafka.InitKafka() // 初始化kafka
	//
	go kafka.FavorConsumer()    // 启动消费者开始监听
	time.Sleep(1 * time.Second) // 因为下面的内容执行太快了，所以上面没有读取出来，但其实可以读取出来的
	defer kafka.KafkaProducer.Close()
	defer kafka.CloseConsumer()

	key := "favorite"         // 注册的名字
	value := "127.0.0.1:9003" // 注册的服务地址
	// 传入注册的服务名和注册的服务地址进行注册
	serviceRegister, grpcServer := discovery.InitRegister(key, value)
	defer serviceRegister.Close()
	defer grpcServer.Stop()
	pb.RegisterDouyinFavoriteActionServiceServer(grpcServer, &service.FavoriteSrv{}) // 绑定grpc
	discovery.GrpcListen(grpcServer, value)                                          // 开启监听
}
