package dao

import (
	"dzug/protos/favorite"
	"dzug/repo"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

func Favor(videoId, userId int64) error {
	favorite := repo.Favorite{
		UserId:  userId,
		VideoId: videoId,
	}
	video := repo.Video{
		Model: gorm.Model{
			ID: uint(videoId),
		},
	}
	res := repo.DB.First(&video)
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

	err := repo.DB.Transaction(func(tx *gorm.DB) (err error) {
		if err = tx.Where("user_id = ?", userId).Save(&favorite).Error; err != nil {
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

func InFavor(videoId, userId int64) error {
	favorite := repo.Favorite{
		UserId:  userId,
		VideoId: videoId,
	}
	video := repo.Video{
		Model: gorm.Model{
			ID: uint(videoId),
		},
	}
	res := repo.DB.First(&video)
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

	err := repo.DB.Transaction(func(tx *gorm.DB) (err error) {
		if err = tx.Where("user_id = ?", userId).Delete(&favorite).Error; err != nil {
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
	res := repo.DB.Where("user_id = ?", userId).Find(&favors)
	if res.Error != nil {
		return nil, res.Error
	}
	ans := make([]int64, len(favors))
	for k, v := range favors {
		ans[k] = v.VideoId
	}
	return ans, nil
}

// GetVideosByVideoIds 根据videoId返回videos
func GetVideosByVideoIds(videoIds []int64) ([]*favorite.Video, error) {
	// 根据videoIds一个一个查，isFavorite都设置为true，作者信息根据那啥去查？
	return nil, nil
}
