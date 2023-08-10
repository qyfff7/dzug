package service

import (
	"context"
	"dzug/app/publish/infra/dal"
	"dzug/app/publish/infra/redis"
	"dzug/app/publish/pkg/oss"
	pb "dzug/protos/publish"
	"go.uber.org/zap"
)

type VideoServer struct {
	pb.UnimplementedPublishServiceServer
}

// PublishVideo 视频投稿服务
func (p *VideoServer) PublishVideo(ctx context.Context, req *pb.PublishVideoReq) (*pb.PublishVideoResp, error) {
	resp := new(pb.PublishVideoResp)
	// 删除旧的视频列表
	if err := redis.DelPublishList(ctx, req.Token); err != nil {
		zap.L().Error(err.Error())
	}

	// TODO: 根据token获取userid
	var userID int64

	// 对象存储操作
	video := oss.Video{
		Title:    req.Title,
		FileName: req.FileName,
		File:     req.Data,
		UserID:   userID,
	}
	ossUrl, _ := oss.UploadVideoToOss(ctx, &video)

	// 数据库操作
	err := dal.PublishVideo(ctx, req.Token, req.Title, ossUrl.PlayUrl, ossUrl.CoverUrl)
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
