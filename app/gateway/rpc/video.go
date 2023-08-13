package rpc

import (
	"context"
	"dzug/discovery"
	pb "dzug/protos/publish"
	"go.uber.org/zap"
)

func PublishVideo(ctx context.Context, req *pb.PublishVideoReq) (resp *pb.PublishVideoResp, err error) {
	discovery.LoadClient("publish", &discovery.PublishClient)
	r, err := discovery.PublishClient.PublishVideo(ctx, req)
	if err != nil {
		zap.L().Error(err.Error())
		return
	}
	return r, nil
}

func GetPublishListByUser(ctx context.Context, req *pb.GetVideoListByUserIdReq) (resp *pb.GetVideoListByUserIdResp, err error) {
	discovery.LoadClient("publish", &discovery.PublishClient)
	r, err := discovery.PublishClient.GetVideoListByUserId(ctx, req)
	if err != nil {
		zap.L().Error(err.Error())
		return
	}
	return r, nil
}
