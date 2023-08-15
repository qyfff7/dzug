package redis

import (
	"context"
	"dzug/conf"
	"dzug/models"
	"dzug/repo"
	"fmt"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
	"strconv"
	"time"
)

// redis key 注意使用命名空间的方式,方便查询和拆分
const (
	Prefix      = "douyin:"   // 项目key前缀
	KeyUserInfo = "userinfo:" //保存用户信息
	KeyUserId   = "userId:"   //保存用户信息
)

// 给redis key加上前缀
func GetRedisKey(servicename, key string) string {
	return Prefix + servicename + key
}

var (
	Rdb *redis.Client
	Nil = redis.Nil
)

// Init 初始化连接
func Init(cfg *conf.RedisConfig) (err error) {
	Rdb = redis.NewClient(&redis.Options{
		Addr:         fmt.Sprintf("%s:%d", cfg.Host, cfg.Port),
		Password:     cfg.Password, // no password set
		DB:           cfg.DB,       // use default DB
		PoolSize:     cfg.PoolSize,
		MinIdleConns: cfg.MinIdleConns,
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

func AddUser(ctx context.Context, newuser *repo.User) error {
	//以userid为key,用户的其他所有信息为value,存到redis中
	userinfo := make(map[string]interface{})
	userinfo["user_id"] = newuser.UserId
	userinfo["name"] = newuser.Name
	userinfo["password"] = newuser.Password
	userinfo["background_images"] = newuser.BackgroundImages
	userinfo["follower_count"] = newuser.FollowerCount
	userinfo["follow_count"] = newuser.FollowCount
	userinfo["favorite_count"] = newuser.FavoriteCount
	userinfo["total_favorited"] = newuser.TotalFavorited
	userinfo["signature"] = newuser.Signature
	userinfo["avatar"] = newuser.Avatar
	userinfo["work_count"] = newuser.WorkCount

	ok, _ := Rdb.SIsMember(ctx, GetRedisKey(KeyUserId, ""), newuser.UserId).Result()
	if !ok {
		err := Rdb.SAdd(ctx, GetRedisKey(KeyUserId, ""), newuser.UserId).Err()
		if err != nil {
			zap.L().Error("用户id存到redis出错", zap.Error(err))
			return err
		}
		err = Rdb.HMSet(ctx, GetRedisKey(KeyUserInfo, strconv.Itoa(int(newuser.UserId))), userinfo).Err()
		if err != nil {
			zap.L().Error("用户信息存到redis出错", zap.Error(err))
			return err
		}
		Rdb.Expire(ctx, GetRedisKey(KeyUserId, ""), time.Duration(conf.Config.RedisConfig.RedisExpire)*time.Hour)
		//Rdb.Expire(ctx, GetRedisKey(KeyUserInfo, strconv.Itoa(int(newuser.UserId))), time.Duration(conf.Config.RedisConfig.RedisExpire)*time.Second)
	}

	return nil
}

func GetUserInfoByID(ctx context.Context, userid int64) (*models.User, error) {

	userinfo := []string{"user_id", "name", "signature", "work_count",
		"total_favorited", "background_images", "follower_count", "follow_count",
		"avatar", "favorite_count"}
	vals, err := Rdb.HMGet(ctx, GetRedisKey(KeyUserInfo, strconv.Itoa(int(userid))), userinfo...).Result()
	if err != nil {
		zap.L().Error("", zap.Error(err))
		return nil, err
	}
	uInfo := new(models.User)

	uInfo.ID, _ = strconv.ParseInt(fmt.Sprintf("%s", vals[0]), 10, 64)
	uInfo.Name = vals[1].(string)
	uInfo.Signature = vals[2].(string)
	uInfo.WorkCount, _ = strconv.ParseInt(fmt.Sprintf("%s", vals[3]), 10, 64)
	uInfo.TotalFavorited, _ = strconv.ParseInt(fmt.Sprintf("%s", vals[4]), 10, 64)
	uInfo.BackgroundImage = vals[5].(string)
	uInfo.FollowerCount, _ = strconv.ParseInt(fmt.Sprintf("%s", vals[6]), 10, 64)
	uInfo.FollowCount, _ = strconv.ParseInt(fmt.Sprintf("%s", vals[7]), 10, 64)
	uInfo.Avatar = vals[8].(string)
	uInfo.FavoriteCount, _ = strconv.ParseInt(fmt.Sprintf("%s", vals[9]), 10, 64)
	uInfo.IsFollow = false
	return uInfo, nil
}
