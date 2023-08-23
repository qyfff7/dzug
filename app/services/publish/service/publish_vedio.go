package service

import (
	"context"
	"dzug/app/services/publish/dal/dao"
	"dzug/app/services/publish/dal/redis"
	"dzug/app/services/publish/pkg/oss"
	pb "dzug/protos/publish"
	r "github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)

type VideoServer struct {
	pb.UnimplementedPublishServiceServer
}

// PublishVideo 视频投稿服务
func (p *VideoServer) PublishVideo(ctx context.Context, req *pb.PublishVideoReq) (*pb.PublishVideoResp, error) {
	resp := new(pb.PublishVideoResp)
	// 删除旧的视频列表
	if err := redis.DelPublishList(ctx, req.UserId); err != nil {
		zap.L().Error(err.Error())
	}

	// 对象存储操作
	video := oss.Video{
		Title:    req.Title,
		FileName: req.FileName,
		File:     req.Data,
		UserID:   req.UserId,
	}
	ossUrl, _ := oss.UploadVideoToOss(ctx, &video)

	// 数据库操作
	err := dao.PublishVideo(ctx, req.UserId, req.Title, ossUrl.PlayUrl, ossUrl.CoverUrl)
	if err != nil {
		resp.StatusCode = 400
		resp.StatusMsg = "发布失败"
		zap.L().Error(err.Error())
		return resp, err
	}
	resp.StatusCode = 200
	resp.StatusMsg = "发布成功"
	return resp, nil
}

func (p *VideoServer) GetVideoListByUserId(ctx context.Context, req *pb.GetVideoListByUserIdReq) (*pb.GetVideoListByUserIdResp, error) {
	resp := pb.GetVideoListByUserIdResp{}
	user_id := req.UserId

	videoModels, err := redis.GetPublishList(user_id)

	if err != nil {
		// 缓存未命中
		if err == r.Nil {
			// 去数据库查询
			videoModels, err = dao.GetVideoListByUserId(ctx, user_id)
			if err != nil {
				return nil, err
			}

			// 查询结果写入 redis
			if err := redis.PutPublishList(ctx, videoModels, user_id); err != nil {
				zap.L().Error(err.Error())
			}
		}
	}

	// 获取userInfo
	userInfo, err := dao.GetUserInfoByUserId(ctx, user_id)
	rspUserInfo := &pb.UserInfo{
		Id:              userInfo.UserId,
		Name:            userInfo.Name,
		FollowCount:     &userInfo.FollowCount,
		FollowerCount:   &userInfo.FollowerCount,
		Avatar:          &userInfo.Avatar,
		BackgroundImage: &userInfo.BackgroundImages,
		Signature:       &userInfo.Signature,
		TotalFavorited:  &userInfo.TotalFavorited,
		WorkCount:       &userInfo.WorkCount,
		FavoriteCount:   &userInfo.FavoriteCount,
	}

	var videoInfoList []*pb.VideoInfo
	for i := 0; i < len(videoModels); i++ {
		tmp := &pb.VideoInfo{
			PlayUrl:       videoModels[i].PlayUrl,
			CoverUrl:      videoModels[i].CoverUrl,
			FavoriteCount: int64(videoModels[i].FavoriteCount),
			CommentCount:  int64(videoModels[i].CommentCount),
			Id:            int64(videoModels[i].ID),
			Title:         videoModels[i].Title,
			Author:        rspUserInfo,
		}
		videoInfoList = append(videoInfoList, tmp)
	}

	resp.StatusCode = 200
	resp.StatusMsg = "success"
	resp.VideoList = videoInfoList

	return &resp, nil
}
