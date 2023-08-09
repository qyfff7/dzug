package service

import (
	"context"
	"dzug/app/publish/infra/dal"
	"dzug/app/publish/infra/redis"
	pb "dzug/protos/publish"
	"go.uber.org/zap"
)

type VideoServer struct {
	pb.UnimplementedPublishServiceServer
}

func (p *VideoServer) PublishVideo(ctx context.Context, req *pb.PublishVideoReq) (*pb.PublishVideoResp, error) {
	resp := new(pb.PublishVideoResp)
	if err := redis.DelPublishList(ctx, req.VideoBaseInfo.UserId); err != nil {
		zap.L().Error(err.Error())
	}
	err := dal.PublishVideo(ctx, req.VideoBaseInfo.UserId, req.VideoBaseInfo.Title, req.VideoBaseInfo.PlayUrl, req.VideoBaseInfo.CoverUrl)
	if err != nil {
		resp.BaseResp.StatusCode = 400
		resp.BaseResp.StatusMsg = "发布失败"
		zap.L().Error(err.Error())
		return resp, err
	}
	resp.BaseResp.StatusCode = 200
	resp.BaseResp.StatusMsg = "发布成功"
	return resp, nil
}
func (p *VideoServer) GetVideoListByUserId(ctx context.Context, req *pb.GetVideoListByUserIdReq) (*pb.GetVideoListByUserIdResp, error) {
	resp := pb.GetVideoListByUserIdResp{}
	return &resp, nil
}
