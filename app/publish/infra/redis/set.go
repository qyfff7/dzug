package redis

import (
	"context"
	"dzug/app/publish/infra/dal/model"
	redisModel "dzug/app/publish/infra/redis/model"
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

func PutPublishList(ctx context.Context, videoListp []*model.Video, userId int64) error {

	if videoListp == nil || len(videoListp) == 0 {
		return nil
	}

	var videoList []redisModel.VideoCache

	for i := range videoListp {
		videoList = append(videoList, redisModel.VideoCache{
			VideoId:  videoListp[i].ID,
			UserId:   videoListp[i].UserId,
			Title:    videoListp[i].Title,
			PlayUrl:  videoListp[i].PlayUrl,
			CoverUrl: videoListp[i].CoverUrl,
		})
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
