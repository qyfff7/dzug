package redis

import (
	"context"
	"dzug/app/publish/infra/dal/model"
)

func DelPublishList(ctx context.Context, token string) error {
	err := RDB.Del(ctx, token).Err()
	if err != nil {
		return err
	}
	return nil
}

func PutPublishList(videoList []*model.Video, userId int64) error {
	return nil
}
