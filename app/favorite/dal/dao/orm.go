package dao

import (
	"context"
	"dzug/protos/favorite"
	"dzug/repo"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

// Favor 点赞数据库操作
func Favor(videoId, userId int64) error {
	favor := repo.Favorite{ // 初始化结构体
		UserId:  userId,
		VideoId: videoId,
	}
	video := repo.Video{
		Model: gorm.Model{
			ID: uint(videoId),
		},
	}
	res := repo.DB.First(&video)
	if res.Error != nil { // 查询作者id，便于后续进行更新
		zap.L().Error("查询视频作者失败")
		return res.Error
	}
	author := repo.User{
		UserId: video.UserId,
	}
	user := repo.User{
		UserId: userId,
	}

	err := repo.DB.Transaction(func(tx *gorm.DB) (err error) {
		if err = tx.Where("user_id = ?", userId).Save(&favor).Error; err != nil {
			zap.L().Error("点赞失败")
			return err
		}
		if err = tx.Model(&video).UpdateColumn("favorite_count", gorm.Expr("favorite_count + ?", 1)).Error; err != nil {
			zap.L().Error("更新视频点赞数量失败")
			return err
		}
		if err = tx.Model(&author).Where("user_id = ?", author.UserId).UpdateColumn("total_favorited", gorm.Expr("total_favorited + ?", 1)).Error; err != nil {
			zap.L().Error("更新作者获赞总数失败")
			return err
		}
		if err = tx.Model(&user).Where("user_id = ?", user.UserId).UpdateColumn("favorite_count", gorm.Expr("favorite_count + ?", 1)).Error; err != nil {
			zap.L().Error("更新用户喜欢总数失败")
			return err
		}
		return nil
	})
	return err
}

// InFavor 取消点赞数据库操作
func InFavor(videoId, userId int64) error {
	favor := repo.Favorite{ // 初始化结构体
		UserId:  userId,
		VideoId: videoId,
	}
	video := repo.Video{
		Model: gorm.Model{
			ID: uint(videoId),
		},
	}
	res := repo.DB.First(&video) // 查询作者id，便于后续事务更新数据
	if res.Error != nil {
		zap.L().Error("查询视频作者失败")
		return res.Error
	}
	author := repo.User{
		UserId: video.UserId,
	}
	user := repo.User{
		UserId: userId,
	}

	// 开启事务，写入点赞数据，插入点赞表，更新用户表，视频表
	err := repo.DB.Transaction(func(tx *gorm.DB) (err error) {
		if err = tx.Where("user_id = ?", userId).Delete(&favor).Error; err != nil {
			zap.L().Error("取消点赞失败")
			return err
		}
		if err = tx.Model(&video).UpdateColumn("favorite_count", gorm.Expr("favorite_count - ?", 1)).Error; err != nil {
			zap.L().Error("更新视频点赞数量失败")
			return err
		}
		if err = tx.Model(&author).Where("user_id = ?", author.UserId).UpdateColumn("total_favorited", gorm.Expr("total_favorited - ?", 1)).Error; err != nil {
			zap.L().Error("更新作者获赞总数失败")
			return err
		}
		if err = tx.Model(&user).Where("user_id = ?", user.UserId).UpdateColumn("favorite_count", gorm.Expr("favorite_count - ?", 1)).Error; err != nil {
			zap.L().Error("更新用户喜欢总数失败")
			return err
		}
		return nil
	})
	return err
}

// GetFavorById 获取用户的所有点赞视频
func GetFavorById(userId int64) ([]int64, error) {
	var favors []repo.Favorite
	res := repo.DB.Where("user_id = ?", userId).Find(&favors) // 查询该用户所有点赞视频数据
	if res.Error != nil {
		return nil, res.Error
	}
	ans := make([]int64, len(favors))
	for k, v := range favors { // 写入返回的数组
		ans[k] = v.VideoId
	}
	return ans, nil
}

// GetVideosByVideoIds 根据videoId返回videos
func GetVideosByVideoIds(userId int64, videoIds []int64) ([]*favorite.Video, error) {
	// 根据videoIds一个一个查，isFavorite都设置为true，作者信息也进行查询拼接
	video := repo.Video{}
	videoList := make([]*favorite.Video, len(videoIds))
	for k, v := range videoIds {
		video.ID = uint(v)
		res := repo.DB.First(&video)
		if res.Error != nil {
			zap.L().Error("读取视频信息失败")
			return nil, res.Error
		}
		videoList[k] = &favorite.Video{} // 初始化
		videoList[k].Author = &favorite.User{}
		videoList[k].Id = videoIds[k]
		res = repo.DB.Where("user_id = ?", video.UserId).First(&videoList[k].Author)
		videoList[k].Author.Id = video.UserId // 查询并设置视频作者信息

		videoList[k].Title = video.Title // 设置视频基本数据
		videoList[k].CommentCount = int64(video.CommentCount)
		videoList[k].CoverUrl = video.CoverUrl
		videoList[k].PlayUrl = video.PlayUrl
		videoList[k].FavoriteCount = int64(video.FavoriteCount)

		videoList[k].Author.IsFollow = isFollowById(userId, video.UserId) // 更新是否关注
		videoList[k].IsFavorite = true                                    // 设置已喜欢
	}
	return videoList, nil
}

func isFollowById(userId int64, authorId int64) bool {
	var rel repo.Relation
	result := repo.DB.WithContext(context.Background()).Table("relation").Where("user_id = ? AND to_user_id = ?", userId, authorId).Limit(1).Find(&rel)
	if result.Error != nil {
		zap.L().Error("查找关注关系时出错")
		return false
	}
	if result.RowsAffected > 0 { //关注了该用户
		return true
	}
	return false //未关注
}
