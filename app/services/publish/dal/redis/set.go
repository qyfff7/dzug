package redis

import (
	"context"
	"dzug/models"
	"encoding/json"
	"math/rand"
	"strconv"
	"time"
)

func DelPublishList(ctx context.Context, user_id int64) error {
	key := "publish" + ":" + strconv.FormatInt(user_id, 10)
	err := RDB.Del(ctx, key).Err()
	if err != nil {
		return err
	}
	return nil
}

func PutPublishList(ctx context.Context, videoList []*models.Video, userId int64) error {

	if videoList == nil || len(videoList) == 0 {
		return nil
	}
	ub, err := json.Marshal(videoList)
	if err != nil {
		return err
	}

	key := "publish" + ":" + strconv.FormatInt(userId, 10)
	expireTime := time.Duration(rand.Intn(24)) * time.Hour
	_, err = RDB.Set(ctx, key, ub, time.Duration(expireTime)).Result()
	if err != nil {
		_ = RDB.Del(ctx, key).Err()
		return err
	}
	return nil
}
