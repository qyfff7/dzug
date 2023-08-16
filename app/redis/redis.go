package redis

import (
	"context"
	"dzug/conf"
	"fmt"

	"github.com/redis/go-redis/v9"
)

var (
	client *redis.Client
	Nil    = redis.Nil
)

// Init 初始化连接
func Init(cfg *conf.RedisConfig) (err error) {
	client = redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%d", cfg.Host, cfg.Port),
	})
	ctx := context.Background()
	_, err = client.Ping(ctx).Result()
	if err != nil {
		return err
	}
	return nil
}

func Close() {
	_ = client.Close()
}
