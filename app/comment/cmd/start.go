package commentservice

import (
	"dzug/app/comment/service"
	"dzug/conf"
	"dzug/discovery"

	pb "dzug/protos/comment"

	"go.uber.org/zap"
)

func Start() {

	zap.L().Info("服务启动，开始记录日志")

	serviceRegister, grpcServer := discovery.InitRegister(conf.Config.CommentServiceName, conf.Config.CommentServiceUrl)
	defer serviceRegister.Close()
	defer grpcServer.Stop()
	pb.RegisterDouyinCommentServiceServer(grpcServer, &service.CommentSrv{}) // 绑定grpc
	discovery.GrpcListen(grpcServer, conf.Config.CommentServiceUrl)          // 开启监听

}
