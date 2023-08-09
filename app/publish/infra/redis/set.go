package redis

import (
	"context"
	"dzug/app/publish/infra/dal/model"
	"strconv"
)

func DelPublishList(ctx context.Context, userId int64) error {
	err := RDB.Del(ctx, strconv.FormatInt(userId, 10)).Err()
	if err != nil {
		return err
	}
	return nil
}

func PutPublishList(videoList []*model.Video, userId int64) error {
	return nil
}
