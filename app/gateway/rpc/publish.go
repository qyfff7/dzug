package rpc

import (
	"context"
	"dzug/discovery"
	pb "dzug/protos/publish"
	"go.uber.org/zap"
)

func PublishVideo(ctx context.Context, req *pb.PublishVideoReq) (resp *pb.PublishVideoResp, err error) {
	err = discovery.LoadClient("publish", &discovery.PublishClient)
	if err != nil {
		return nil, err
	}
	r, err := discovery.PublishClient.PublishVideo(ctx, req)
	if err != nil {
		zap.L().Error(err.Error())
		return
	}
	return r, nil
}

func GetPublishListByUser(ctx context.Context, req *pb.GetVideoListByUserIdReq) (resp *pb.GetVideoListByUserIdResp, err error) {
	err = discovery.LoadClient("publish", &discovery.PublishClient)
	if err != nil {
		zap.L().Error("加载服务发现出错")
		return nil, err
	}
	r, err := discovery.PublishClient.GetVideoListByUserId(ctx, req)
	if err != nil {
		zap.L().Error(err.Error())
		return
	}
	return r, nil
}
