package redis

import (
	"context"
	"dzug/app/publish/infra/dal/model"
	redisVideoModle "dzug/app/publish/infra/redis/model"
	"encoding/json"
	"math/rand"
	"strconv"
	"time"
)

func GetPublishList(userId int64) ([]*model.Video, error) {
	ctx := context.Background()
	// 根据user_id去redis中拿缓存
	key := "publish" + ":" + strconv.FormatInt(userId, 10)
	res, err := RDB.Get(ctx, key).Bytes()
	if err != nil {
		return nil, err
	}

	// 解析redis值对象
	videoRedisListp := new([]redisVideoModle.VideoCache)
	err = json.Unmarshal(res, videoRedisListp)
	if err != nil {
		return nil, err
	}
	// 设置过期时间
	expireTime := time.Duration(rand.Intn(24)) * time.Hour
	err = RDB.Expire(ctx, strconv.FormatInt(userId, 10), time.Duration(expireTime)).Err()
	if err != nil {
		return nil, err
	}

	videoRedisList := *videoRedisListp
	videoList := make([]*model.Video, len(videoRedisList))

	for it := range videoRedisList {
		videoList[it] = &model.Video{
			UserId:   videoRedisList[it].UserId,
			Title:    videoRedisList[it].Title,
			PlayUrl:  videoRedisList[it].PlayUrl,
			CoverUrl: videoRedisList[it].CoverUrl,
		}
		videoList[it].ID = videoRedisList[it].VideoId
	}
	return videoList, nil

}
