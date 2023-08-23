package publishservice

import (
	"dzug/app/services/publish/dal/redis"
	"dzug/app/services/publish/pkg/oss"
	"dzug/app/services/publish/service"
	"dzug/conf"
	"dzug/discovery"
	pb "dzug/protos/publish"
	"dzug/repo"
)

func Start() {
	/*if err := conf.Init(); err != nil {
		fmt.Printf("Config file initialization error,%#v", err)
		return
	}

	if err := logger.Init(conf.Config.LogConfig); err != nil {
		fmt.Printf("log file initialization error,%#v", err)
		return
	}

	defer zap.L().Sync()
	zap.L().Info("服务启动，开始记录日志")*/

	// 初始化redis
	err := redis.Init()
	if err != nil {
		return
	}
	oss.Init()
	_ = repo.Init()
	serviceRegister, grpcServer := discovery.InitRegister(conf.Config.PublishServiceName, conf.Config.PublishServiceUrl)
	defer serviceRegister.Close()
	defer grpcServer.Stop()
	pb.RegisterPublishServiceServer(grpcServer, &service.VideoServer{})
	discovery.GrpcListen(grpcServer, conf.Config.PublishServiceUrl)
}
