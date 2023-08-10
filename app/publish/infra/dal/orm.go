package dal

import (
	"context"
	"dzug/app/publish/infra/dal/model"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

func PublishVideo(ctx context.Context, token string, title string, playUrl string, coverUrl string) error {
	// TODO: 根据 token 解析userid

	var userId int64
	video := &model.Video{
		UserId:   userId,
		Title:    title,
		PlayUrl:  playUrl,
		CoverUrl: coverUrl,
	}

	err := DB.WithContext(ctx).Transaction(func(txn *gorm.DB) error {
		err := txn.Create(&video).Error
		if err != nil {
			zap.L().Error(err.Error())
			return err
		}
		err = txn.Table("user").Where("id = ?", userId).Update("work_count", gorm.Expr("work_count + ?", 1)).Error
		if err != nil {
			zap.L().Error(err.Error())
			return err
		}
		return nil
	})
	return err
}

func GetVideoListByUserId(ctx context.Context, userId int64) ([]*model.Video, error) {
	var videos []*model.Video
	if err := DB.WithContext(ctx).Where("user_id = ?", userId).Find(&videos).Error; err != nil {
		return nil, err
	}
	return videos, nil
}
