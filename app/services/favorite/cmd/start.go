package favorservice

import (
	kafka2 "dzug/app/services/favorite/dal/kafka"
	"dzug/app/services/favorite/dal/redis"
	"dzug/app/services/favorite/service"
	"dzug/conf"
	"dzug/discovery"
	pb "dzug/protos/favorite"
	"time"
)

func Start() {
	////1. 初始化配置文件
	//if err := conf.Init(); err != nil {
	//	fmt.Printf("Config file initialization error,%#v", err)
	//	return
	//}
	//
	////2. 初始化日志
	//if err := logger.Init(conf.Config.LogConfig, conf.Config.Mode); err != nil {
	//	fmt.Printf("log file initialization error,%#v", err)
	//	return
	//}
	//defer zap.L().Sync() //把缓冲区的日志，追加到文件中
	//zap.L().Info("服务启动，开始记录日志")

	//repo.Init()       // 初始化数据库
	redis.InitRedis()  // 初始化redis
	kafka2.InitKafka() // 初始化kafka
	//
	go kafka2.FavorConsumer()   // 启动消费者开始监听
	time.Sleep(1 * time.Second) // 因为下面的内容执行太快了，所以上面没有读取出来，但其实可以读取出来的
	defer kafka2.KafkaProducer.Close()
	defer kafka2.CloseConsumer()

	// 传入注册的服务名和注册的服务地址进行注册
	serviceRegister, grpcServer := discovery.InitRegister(conf.Config.FavoriteServiceName, conf.Config.FavoriteServiceUrl)
	defer serviceRegister.Close()
	defer grpcServer.Stop()
	pb.RegisterDouyinFavoriteActionServiceServer(grpcServer, &service.FavoriteSrv{}) // 绑定grpc
	discovery.GrpcListen(grpcServer, conf.Config.FavoriteServiceUrl)                 // 开启监听
}
