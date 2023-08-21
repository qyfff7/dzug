package dal

import (
	"context"
	repo "dzug/repo"
	"go.uber.org/zap"
)

func PublishVideo(ctx context.Context, user_id int64, title string, playUrl string, coverUrl string) error {

	video := &repo.Video{
		UserId:   user_id,
		Title:    title,
		PlayUrl:  playUrl,
		CoverUrl: coverUrl,
	}
	txn := repo.DB.Begin()

	if err := txn.Create(video).Error; err != nil {
		txn.Rollback()
		zap.L().Error(err.Error())
	}

	if err := txn.Commit().Error; err != nil {
		zap.L().Error(err.Error())
		return err
	}
	return nil
}

func GetVideoListByUserId(ctx context.Context, userId int64) ([]*repo.Video, error) {
	var videos []*repo.Video
	if err := repo.DB.Where("user_id = ?", userId).Find(&videos).Error; err != nil {
		return nil, err
	}
	return videos, nil
}

func GetUserInfoByUserId(ctx context.Context, userId int64) (*repo.User, error) {
	var userInfo *repo.User
	if err := repo.DB.Where("user_id = ?", userId).Find(&userInfo).Error; err != nil {
		return nil, nil
	}
	return userInfo, nil
}
