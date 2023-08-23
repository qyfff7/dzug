package redis

import (
	"context"
	"dzug/conf"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
	"strconv"
)

var RDB *redis.Client

func Init() error {
	addr := conf.Config.RedisConfig.Host + ":" + strconv.Itoa(conf.Config.RedisConfig.Port)

	RDB = redis.NewClient(&redis.Options{
		Addr: addr,
	})
	ctx := context.Background()
	_, err := RDB.Ping(ctx).Result()
	if err != nil {
		zap.L().Error(err.Error())
		return err
	}
	zap.L().Info("redis初始化完成")
	return nil
}
