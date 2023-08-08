package redis

import (
	"dzug/conf"
	"github.com/redis/go-redis/v9"
	"strconv"
)

var Rdb *redis.Client

func initRedis() {
	addr := conf.Config.RedisConfig.Host + ":" + strconv.Itoa(conf.Config.RedisConfig.Port)
	Rdb = redis.NewClient(&redis.Options{
		Addr: addr,
	})
}
