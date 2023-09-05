package redis

import (
	"context"
	"dzug/app/services/comment/dal/dao"

	"strconv"
	"time"

	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)

var coPrefix = "comm:" // redis中comm前缀
var key string         // redis的key
func AddComm(ctx context.Context, videoId int64, comment_uuid int64) int {
	InitRedis()
	cmd, err := exist(videoId)
	if err != nil {
		return 0
	}
	cmd = Rdb.SAdd(ctx, key, strconv.FormatInt(comment_uuid, 10))
	if cmd.Val() == 0 { // 已经存在这个value了
		return 2
	}

	return 1
}

func DelComm(ctx context.Context, videoId int64, comment_uuid int64) int {
	InitRedis()
	cmd, err := exist(videoId)
	if err != nil {
		return 0
	}
	cmd = Rdb.SRem(context.Background(), key, comment_uuid) // 这个key现在已经存在了，去删除这个videoId
	if cmd.Val() == 0 {                                     // 已经存在这个value了
		return 2
	}
	return 1
}

func GetComm(ctx context.Context, videoId int64) ([]int64, error) {
	InitRedis()
	_, err := exist(videoId)
	if err != nil {
		return nil, err
	}

	cmd := Rdb.SMembers(context.Background(), key)
	commentIDs := make([]int64, len(cmd.Val()))
	for k, v := range cmd.Val() {
		value, _ := strconv.Atoi(v)
		commentIDs[k] = int64(value)
	}

	return commentIDs, cmd.Err()
}

func exist(videoId int64) (*redis.IntCmd, error) {
	key = coPrefix + strconv.FormatInt(videoId, 10)
	cmd := Rdb.Exists(context.Background(), key)
	if cmd.Err() != nil {
		zap.L().Error("查询redis失败")
		return nil, cmd.Err()
	}
	if cmd.Val() == 0 {
		zap.L().Debug("当前user不存在redis，正在查找数据库")
		err := getSet(videoId)
		if err != nil {
			return nil, err
		}
	}
	return cmd, nil
}

// getSet 从数据库中得到该video的所有评论数据，写入redis中
func getSet(videoId int64) error {
	videoIds, err := dao.GetcommByVideoId(videoId)
	if err != nil {
		zap.L().Error("获取用户" + strconv.FormatInt(videoId, 10) + "点赞数据失败")
		return err
	}
	key = coPrefix + strconv.FormatInt(videoId, 10)
	ctx := context.Background()
	for _, v := range videoIds {
		Rdb.SAdd(ctx, key, v)
	}
	Rdb.Expire(ctx, key, time.Hour*3) // 设置过期时间
	return nil
}
