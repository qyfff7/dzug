package dao

import (
	"context"
	"dzug/app/user/pkg/jwt"
	"dzug/protos/video"
	"dzug/repo"
	"go.uber.org/zap"
	"math"
)

/*func GetVideoFeed(ctx context.Context, req *pb.GetVideoFeedReq) (*pb.GetVideoFeedResp, error) {

	return nil, nil

}*/

// GetVideoInfoByTime 根据时间戳返回最近count个视频,还需要返回next time
func GetVideoInfoByTime(ctx context.Context, req *video.GetVideoListByTimeReq, page int64, size int64) ([]*video.Video, int64, error) {

	//1.按照时间倒序，查询所有的视频
	videos := make([]*repo.Video, 0, size)

	if err := repo.DB.WithContext(ctx).Where("created_at < ?", req.LatestTime).Limit(int(size)).Order("created_at DESC").Find(&videos).Error; err != nil {
		zap.L().Info("获取所有视频orm语句出错")
		return nil, 0, err
	}
	var nextTime int64
	nextTime = math.MaxInt64

	if len(videos) != 0 { // 查到了新视频
		nextTime = videos[0].CreatedAt.Unix()
	}
	var videosInfo *video.Video
	var videosInfos []*video.Video

	for i, _ := range videos {

		if req.Token != "" {
			u, _ := jwt.ParseToken(req.Token)
			IsFavorite, _ := IsFavoriteByID(ctx, u.UserID, videos[i].UserId)
			videosInfo.IsFavorite = IsFavorite
		}

		videosInfo.VideoId = int64(videos[i].ID)
		videosInfo.AutherId = videos[i].UserId
		videosInfo.PlayUrl = videos[i].PlayUrl
		videosInfo.CoverUrl = videos[i].CoverUrl
		videosInfo.Title = videos[i].Title
		videosInfo.FavoriteCount = int64(videos[i].FavoriteCount)
		videosInfo.CommentCount = int64(videos[i].CommentCount)
		videosInfos = append(videosInfos, videosInfo)
	}
	return videosInfos, nextTime, nil
}

// IsFavoriteByID 判断是否点赞了该视频
func IsFavoriteByID(ctx context.Context, userID, videoID int64) (bool, error) {
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
