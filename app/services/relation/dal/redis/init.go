package redis

import (
	"context"
	"dzug/conf"
	"fmt"

	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)

// redis key 注意使用命名空间的方式,方便查询和拆分
const (
	Prefix   = "douyin:"   // 项目key前缀
	follow   = "follow:"   // 关注关系
	follower = "follower:" // 粉丝关系

	KeyUserInfo = "userinfo:" //保存用户信息
	KeyUserId   = "userId:"   //保存用户信息
)

// 给redis key加上前缀
func GetRedisKey(servicename, key string) string {
	return Prefix + servicename + key
}

var (
	rdb *redis.Client
	Nil = redis.Nil
)

func Close() {
	_ = rdb.Close()
}

// Init 初始化连接
func Init() (err error) {
	rdb = redis.NewClient(&redis.Options{
		Addr:         fmt.Sprintf("%s:%d", conf.Config.RedisConfig.Host, conf.Config.RedisConfig.Port),
		Password:     conf.Config.RedisConfig.Password, // no password set
		DB:           conf.Config.RedisConfig.DB,       // use default DB
		PoolSize:     conf.Config.RedisConfig.PoolSize,
		MinIdleConns: conf.Config.RedisConfig.MinIdleConns,
	})
	ctx := context.Background()
	_, err = rdb.Ping(ctx).Result()
	if err != nil {
		zap.L().Error("redis初始化失败", zap.Error(err))
		return err
	}
	zap.L().Info("redis初始化完成")
	return nil
}
