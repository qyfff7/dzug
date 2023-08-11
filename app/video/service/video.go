package service

import (
	"context"
	pb "dzug/protos/video"
	"gorm.io/gorm"
)

type VideoService struct {
	db *gorm.DB
	pb.UnimplementedVideoServiceServer
}

// GetVideoFeed 通过指定latestTime和count，从DAO层获取视频基本信息，并查出当前用户是否点赞，组装后返回
func (v VideoService) GetVideoFeed(ctx context.Context, req *pb.GetVideoFeedReq) (*pb.GetVideoFeedResp, error) {

	/*var videosInfo []*pb.VideoInfo
	//1.获取分页参数
	var page int64

	//2.从数据库中获取视频封面地址、播放地址、标题、userID等信息，返回的是一个视频信息数组
	videos, nextTime, err := dao.GetVideoInfoByTime(ctx, req, page)

	for _, video := range videos {
		// 根据作者id查询作者信息

		autherID :=
		r, err := discovery.UserClient.GetUserInfo(ctx, req) // 调用注册方法

		u, err := userDao.GetuserInfoByID(ctx, auther)
		if err != nil {
			zap.L().Error("获取视频作者信息失败")
			continue
		}
		feed := &pb.VideoInfo{
			VideoId:       int64(video.ID),
			Author:        u,
			PlayUrl:       video.PlayUrl,
			CoverUrl:      video.CoverUrl,
			FavoriteCount: int64(video.FavoriteCount),
			CommentCount:  int64(video.CommentCount),
			IsFavorite:    false, //这里需要进行数据库查询
			Title:         video.Title,
		}
	}

	//3. 遍历上面得到的视频信息数组，给每一个视频补充上作者信息

	//4.返回相应

	resp, err := dao.GetVideoFeed(c, req)
	if err != nil {
		zap.L().Error("获取视频流失败", zap.Error(err))
		return nil, err
	}
	return resp, nil*/
	return nil, nil
}
