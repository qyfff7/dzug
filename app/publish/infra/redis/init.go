package redis

import (
	"dzug/conf"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)

var Rdb *redis.Client

func Init() {
	addr := conf.Config.RedisConfig.Host + ":" + conf.Config.RedisConfig.Port
	Rdb = redis.NewClient(&redis.Options{
		Addr: addr,
	})
	zap.L().Info("redis 客户端初始化成功")
}
