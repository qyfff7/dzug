package service

import (
	"context"
	"dzug/app/video/dao"
	pb "dzug/protos/video"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type VideoService struct {
	db *gorm.DB
	pb.UnimplementedVideoServiceServer
}

// GetVideoFeed 通过指定latestTime和count，从DAO层获取视频基本信息，并查出当前用户是否点赞，组装后返回
func (v VideoService) GetVideoFeed(c context.Context, req *pb.GetVideoFeedReq) (*pb.GetVideoFeedResp, error) {

	//1. 从redis中获取视频封面地址、播放地址、标题、userID等信息，返回的是一个视频信息数组

	//2. 遍历上面得到的视频信息数组，给每一个视频补充上作者信息

	//3.返回相应

	resp, err := dao.GetVideoFeed(c, req)
	if err != nil {
		zap.L().Error("获取视频流失败", zap.Error(err))
		return nil, err
	}
	return resp, nil
}
