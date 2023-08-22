package messageservice

import (
	kf2 "dzug/app/message/infra/kafka"
	"dzug/app/message/infra/mongodb"
	"dzug/app/message/service"
	"dzug/app/user/pkg/snowflake"
	"dzug/conf"
	"dzug/discovery"
	"dzug/kafka"
	pb "dzug/protos/message"
	"dzug/repo"
	"go.uber.org/zap"
)

func Start() {
	/*//1. 初始化配置文件
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
	zap.L().Info("服务启动，开始记录日志")*/

	//3. 初始化数据库
	repo.Init()

	//4. 初始化Kafka
	kafka.InitKafka()
	go kafka.ConsumeMsg("message", kf2.MessageHandler)

	//5. snowflake初始化
	if err := snowflake.Init(conf.Config.StartTime, conf.Config.MachineID); err != nil {
		zap.L().Error("snowflake initialization error", zap.Error(err))
		return
	}

	//6. mongodb初始化
	if err := mongodb.Init(); err != nil {
		zap.L().Error("Mongodb initialization error", zap.Error(err))
		return
	}

	key := "message"          // 注册的名字
	value := "127.0.0.1:9006" // 注册的服务地址
	// 传入注册的服务名和注册的服务地址进行注册
	serviceRegister, grpcServer := discovery.InitRegister(key, value)
	defer serviceRegister.Close()
	defer grpcServer.Stop()
	pb.RegisterDouyinMessageServiceServer(grpcServer, service.MsgSvr{}) // 绑定grpc
	discovery.GrpcListen(grpcServer, value)                             // 开启监听
}
