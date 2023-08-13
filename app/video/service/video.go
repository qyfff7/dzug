package service

import (
	"context"
	"dzug/app/video/dao"
	"dzug/protos/video"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type VideoService struct {
	db *gorm.DB
	video.UnimplementedVideoServiceServer
}

// GetVideoFeed 通过指定latestTime和count，从DAO层获取视频基本信息，并查出当前用户是否点赞，组装后返回
func (v VideoService) GetVideoFeed(ctx context.Context, req *video.GetVideoListByTimeReq) (*video.GetVideoListByTimeResp, error) {

	//1.获取分页和大小参数
	var page, size int64

	//2.从数据库中获取视频封面地址、播放地址、标题、userID等信息，返回的是一个视频信息数组
	videos, nextTime, err := dao.GetVideoInfoByTime(ctx, req, page, size)

	if err != nil {
		zap.L().Error("获取视频信息失败", zap.Error(err))
		return nil, err
	}

	//4.返回相应
	resp := &video.GetVideoListByTimeResp{
		VideoList: videos,
		NextTime:  nextTime,
	}

	return resp, nil

}
