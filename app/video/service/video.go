package service

import (
	"context"
	"dzug/app/user/pkg/jwt"
	"dzug/app/video/dao"
	"dzug/conf"
	"dzug/protos/video"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type VideoService struct {
	db *gorm.DB
	video.UnimplementedVideoServiceServer
}

// GetVideoListByTime 通过指定latestTime和count，从DAO层获取视频基本信息，并查出当前用户是否点赞，组装后返回
func (v VideoService) GetVideoListByTime(ctx context.Context, req *video.GetVideoListByTimeReq) (*video.GetVideoListByTimeResp, error) {

	//1.从数据库中获取视频封面地址、播放地址、标题、userID等信息，返回的是一个视频信息数组
	videos, nextTime, err := dao.GetVideoInfoByTime(ctx, req, conf.Config.FeedCount)
	if err != nil {
		zap.L().Error("获取视频信息失败", zap.Error(err))
		return nil, err
	}

	videosInfos := make([]*video.Video, 0, len(videos))
	//2.如果用户已登录,查询是否点赞
	for _, v := range videos {
		var videosInfo video.Video
		if req.Token != "" {
			u, _ := jwt.ParseToken(req.Token)
			IsFavorite, _ := dao.IsFavoriteByID(ctx, u.UserID, v.ID)
			videosInfo.IsFavorite = IsFavorite
		}
		videosInfo.VideoId = int64(v.ID)
		videosInfo.AutherId = v.UserId
		videosInfo.PlayUrl = v.PlayUrl
		videosInfo.CoverUrl = v.CoverUrl
		videosInfo.Title = v.Title
		videosInfo.FavoriteCount = int64(v.FavoriteCount)
		videosInfo.CommentCount = int64(v.CommentCount)
		videosInfos = append(videosInfos, &videosInfo)
	}

	//3.组装,返回相应
	resp := &video.GetVideoListByTimeResp{
		VideoList: videosInfos,
		NextTime:  nextTime,
	}
	return resp, nil

}
