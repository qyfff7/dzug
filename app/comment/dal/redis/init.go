package redis

import (
	"context"
	"dzug/conf"
	"fmt"
	"strconv"
	"time"

	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)

var Rdb *redis.Client

func InitRedis() {
	addr := conf.Config.RedisConfig.Host + ":" + strconv.Itoa(conf.Config.RedisConfig.Port)
	Rdb = redis.NewClient(&redis.Options{
		Addr: addr,
	})
	ctx := context.Background()

	_, err := Rdb.Ping(ctx).Result()
	if err != nil {
		zap.L().Fatal(fmt.Sprintf("连接redis失败" + err.Error()))
	}

	Rdb.Options().DialTimeout = time.Second * 3
	Rdb.Options().ReadTimeout = 0 // 默认读写时间 3s
	Rdb.Options().WriteTimeout = 0
	//Rdb.Options().ContextTimeoutEnabled = true
	//Rdb.Options().PoolSize
}
