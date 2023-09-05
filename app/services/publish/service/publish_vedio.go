package service

import (
	"context"
	"dzug/app/services/publish/dal/dao"
	"dzug/app/services/publish/dal/redis"
	"dzug/app/services/publish/pkg/cos"
	"dzug/models"
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
	video := cos.Video{
		Title:    req.Title,
		FileName: req.FileName,
		File:     req.Data,
		UserID:   req.UserId,
	}
	cosUrl, _ := cos.UploadVideo(ctx, &video)

	// 数据库操作
	err := dao.PublishVideo(ctx, req.UserId, req.Title, cosUrl.PlayUrl, cosUrl.CoverUrl)
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
	user_id := req.UserId
	videoList, err := redis.GetPublishList(user_id)

	if err != nil {
		// 缓存未命中
		if err == r.Nil {
			// 去数据库查询
			daoVideoModel, err := dao.GetVideoListByUserId(ctx, user_id)
			if err != nil {
				return nil, err
			}

			daoUserModel, err := dao.GetUserInfoByUserId(ctx, user_id)
			if err != nil {
				return nil, err
			}
			for _, v := range daoVideoModel {
				videoListTmp := &models.Video{
					Id: int64(v.ID),
					Author: models.User{
						ID:              daoUserModel.UserId,
						Name:            daoUserModel.Name,
						FollowCount:     daoUserModel.FollowCount,
						FollowerCount:   daoUserModel.FollowerCount,
						Avatar:          daoUserModel.Avatar,
						BackgroundImage: daoUserModel.BackgroundImages,
						TotalFavorited:  daoUserModel.TotalFavorited,
						WorkCount:       daoUserModel.WorkCount,
						FavoriteCount:   daoUserModel.FavoriteCount,
					},
					PlayUrl:       v.PlayUrl,
					CoverUrl:      v.CoverUrl,
					FavoriteCount: int64(v.FavoriteCount),
					CommentCount:  int64(v.CommentCount),
					Title:         v.Title,
				}
				videoList = append(videoList, videoListTmp)
			}

			// 查询结果写入 redis
			if err := redis.PutPublishList(ctx, videoList, user_id); err != nil {
				zap.L().Error(err.Error())
			}
		}
	}

	var rtnVidelList []*pb.VideoInfo
	for _, v := range videoList {
		uInfo := pb.UserInfo{
			Id:              v.Author.ID,
			Name:            v.Author.Name,
			FollowCount:     &v.Author.FollowCount,
			FollowerCount:   &v.Author.FollowerCount,
			Avatar:          &v.Author.Avatar,
			BackgroundImage: &v.Author.BackgroundImage,
			Signature:       &v.Author.Signature,
			TotalFavorited:  &v.Author.TotalFavorited,
			WorkCount:       &v.Author.WorkCount,
			FavoriteCount:   &v.Author.FavoriteCount,
		}
		vInfo := &pb.VideoInfo{
			Author:        &uInfo,
			Id:            v.Id,
			PlayUrl:       v.PlayUrl,
			CoverUrl:      v.CoverUrl,
			FavoriteCount: v.FavoriteCount,
			CommentCount:  v.CommentCount,
			Title:         v.Title,
		}
		rtnVidelList = append(rtnVidelList, vInfo)
	}

	rtn := pb.GetVideoListByUserIdResp{
		StatusCode: 200,
		StatusMsg:  "Success",
		VideoList:  rtnVidelList,
	}

	return &rtn, nil
}
