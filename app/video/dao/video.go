package dao

import (
	"context"
	"dzug/protos/video"
	"dzug/repo"
	"fmt"
	"go.uber.org/zap"
	"time"
)

/*func GetVideoFeed(ctx context.Context, req *pb.GetVideoFeedReq) (*pb.GetVideoFeedResp, error) {

	return nil, nil

}*/

// GetVideoInfoByTime 根据时间戳返回最近count个视频,还需要返回next time
func GetVideoInfoByTime(ctx context.Context, req *video.GetVideoListByTimeReq, feedcount int64) ([]*repo.Video, int64, error) {

	//1.按照时间倒序，查询所有的视频
	videos := make([]*repo.Video, 0, feedcount)
	//var videos []*repo.Video
	if req.LatestTime == 0 {
		req.LatestTime = time.Now().Unix()
	}
	//2.将时间戳转换成时间
	t := time.Unix(req.LatestTime, 0)
	//3.从video表中查询出前feedcount个video信息
	if err := repo.DB.WithContext(ctx).Where("created_at < ?", t).Limit(int(feedcount)).Order("created_at DESC").Find(&videos).Error; err != nil {
		zap.L().Info("获取所有视频orm语句出错")
		return nil, 0, err
	}
	zap.L().Info("查到的视频数量" + fmt.Sprintln(len(videos)))
	for i, _ := range videos {
		zap.L().Info("视频id" + fmt.Sprintln(videos[i].ID))
	}
	zap.L().Info("从video表中查询最新的视频，orm操作完成")
	//4.如果查询到了新视频，更新nextTime
	var nextTime int64
	// 查到了新视频
	if len(videos) != 0 {
		nextTime = videos[len(videos)-1].CreatedAt.Unix()
	}
	return videos, nextTime, nil

	/*
		videosInfos := make([]*video.Video, 0, len(videos))



		var videosInfo *video.Video
		//var videosInfos []*video.Video
		videosInfos := make([]*video.Video, 0, len(videos))
		for i, _ := range videos {
			if req.Token != "" {
				u, _ := jwt.ParseToken(req.Token)
				IsFavorite, _ := IsFavoriteByID(ctx, u.UserID, videos[i].UserId)
				videosInfo.IsFavorite = IsFavorite
			} else {
				videosInfo.IsFavorite = false
			}
			videosInfo.VideoId = int64(videos[i].ID)
			videosInfo.AutherId = videos[i].UserId
			videosInfo.PlayUrl = videos[i].PlayUrl
			videosInfo.CoverUrl = videos[i].CoverUrl
			videosInfo.Title = videos[i].Title
			videosInfo.FavoriteCount = int64(videos[i].FavoriteCount)
			videosInfo.CommentCount = int64(videos[i].CommentCount)
			zap.L().Info(fmt.Sprintln(videos[i].ID))
			videosInfos = append(videosInfos, videosInfo)

		}
		return videosInfos, nextTime, nil
	*/
}

// IsFavoriteByID 判断是否点赞了该视频
func IsFavoriteByID(ctx context.Context, userID int64, videoID uint) (bool, error) {
	var rel repo.Favorite
	result := repo.DB.WithContext(ctx).Where("user_id = ? AND video_id = ?", userID, videoID).Limit(1).Find(&rel)
	if result.Error != nil {
		zap.L().Info("查找点赞关系时出错")
		return false, result.Error
	}
	if result.RowsAffected > 0 { //点赞
		return true, nil
	}
	return false, nil //未点赞
}
