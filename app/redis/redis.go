package redis

import (
	"context"
	"dzug/conf"
	"fmt"

	"github.com/redis/go-redis/v9"
)

var (
	Rdb *redis.Client
	Nil = redis.Nil
)

// Init 初始化连接
func Init(cfg *conf.RedisConfig) (err error) {
	Rdb = redis.NewClient(&redis.Options{
		Addr:         fmt.Sprintf("%s:%d", conf.Config.RedisConfig.Host, conf.Config.RedisConfig.Port),
		Password:     conf.Config.RedisConfig.Password, // no password set
		DB:           conf.Config.RedisConfig.DB,       // use default DB
		PoolSize:     conf.Config.RedisConfig.PoolSize,
		MinIdleConns: conf.Config.RedisConfig.MinIdleConns,
	})
	ctx := context.Background()
	_, err = Rdb.Ping(ctx).Result()
	if err != nil {
		return err
	}
	return nil
}

func Close() {
	_ = Rdb.Close()
}
