package db

import (
	"context"
	"dzug/repo"
)

func CreateMessage(ctx context.Context, message *Message) error {

	err := repo.DB.WithContext(ctx).Create(&message).Error

	return err
}

func GetMessageList(ctx context.Context, userId int64, toUserId int64, latestTime int64) ([]*Message, error) {
	var messages []*Message
	err := repo.DB.WithContext(ctx).Where("from_user_id = ? AND to_user_id = ? AND create_time >= ?",
		userId, toUserId, latestTime).Or("from_user_id = ? AND to_user_id = ? AND create_time >= ?",
		toUserId, userId, latestTime).Find(&messages).Error
	if err != nil {
		return nil, err
	}
	return messages, nil
}
