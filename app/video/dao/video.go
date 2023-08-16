package dao

import (
	"context"
	"dzug/protos/video"
	"dzug/repo"
	"go.uber.org/zap"
	"time"
)

// GetVideoInfoByTime 根据时间戳返回最近count个视频,还需要返回next time
func GetVideoInfoByTime(ctx context.Context, req *video.GetVideoListByTimeReq, feedcount int64) ([]*repo.Video, int64, error) {

	//1.按照时间倒序，查询所有的视频
	videos := make([]*repo.Video, 0, feedcount)
	if req.LatestTime <= 0 || req.LatestTime > time.Now().Unix() {
		req.LatestTime = time.Now().Unix()
	}
	//2.将时间戳转换成时间
	t := time.Unix(req.LatestTime, 0)
	//3.从video表中查询出前feedcount个video信息
	if err := repo.DB.WithContext(ctx).Where("created_at < ?", t).Limit(int(feedcount)).Order("created_at DESC").Find(&videos).Error; err != nil {
		zap.L().Info("获取所有视频orm语句出错")
		return nil, 0, err
	}
	//4.如果查询到了新视频，更新nextTime
	var nextTime int64
	// 查到了新视频
	if len(videos) != 0 {
		nextTime = videos[len(videos)-1].CreatedAt.Unix()
	}
	return videos, nextTime, nil
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
