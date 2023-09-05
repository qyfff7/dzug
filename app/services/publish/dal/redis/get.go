package redis

import (
	"context"
	"dzug/models"
	"encoding/json"
	"go.uber.org/zap"
	"math/rand"
	"strconv"
	"time"
)

func GetPublishList(userId int64) ([]*models.Video, error) {
	ctx := context.Background()
	// 根据user_id去redis中拿缓存
	key := "publish" + ":" + strconv.FormatInt(userId, 10)
	res, err := RDB.Get(ctx, key).Bytes()
	if err != nil {
		zap.L().Error("在缓存中获取用户列表失败")
		return nil, err
	}

	// 解析redis值对象
	var videoList []*models.Video
	err = json.Unmarshal(res, &videoList)
	if err != nil {
		return nil, err
	}
	// 设置过期时间
	expireTime := time.Duration(rand.Intn(24)) * time.Hour
	err = RDB.Expire(ctx, strconv.FormatInt(userId, 10), time.Duration(expireTime)).Err()
	if err != nil {
		return nil, err
	}

	return videoList, nil

}
