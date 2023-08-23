/*
实现方法：
GetFollowList(userId int64) ([]int64, error) {} 从Redis中获取关注列表
GetFollowerList(userId int64) ([]int64, error) {} 从 Redis 中获取粉丝列表
GetFriendList(userId int64) ([]int64, error) {} 从 Redis 中获取好友列表
func GetIsFollowById(userId int64, followId int64) (bool, error) {} Redis 中检查一个用户是否关注了另一个用户:
IsFollowKeyExist(userId int64) (bool, error) 检查关注列表在 Redis 中是否存在对应的键
IsFanKeyExist(userId int64) (bool, error){}  检查粉丝列表在 Redis 中是否存在对应的键
*/

package redis

import (
	"context"
	"dzug/conf"
	"strconv"
	"time"

	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)

func GetFollowList(ctx context.Context, userId int64) ([]int64, error) {
	zap.L().Info("start to get the follow list")
	// 1. 构建 followkey
	followKey := GetRedisKey(follow, strconv.FormatInt(userId, 10))

	// 2.从Redis中获取有序集合中指定范围的成员
	members, err := rdb.ZRange(ctx, followKey, 0, -1).Result()

	if err != nil {
		zap.L().Error("Redis中获取FollowList失败: ", zap.Error(err))
		return nil, err
	}

	// 创建切片，逐步将关注列表成员members转为int64的关注用户id
	followList := make([]int64, len(members))
	for i, member := range members {
		id, _ := strconv.ParseInt(member, 10, 64)
		followList[i] = id
	}

	// 设置key的过期时间
	rdb.Expire(ctx, followKey, time.Duration(conf.Config.RedisConfig.RedisExpire)*time.Hour)
	if err != nil {
		return nil, err
	}

	zap.L().Info("Redis缓存中获取followList成功")
	return followList, nil
}

func GetFollowerList(ctx context.Context, userId int64) ([]int64, error) {
	zap.L().Info("start to get the follower list")

	followerKey := GetRedisKey(follower, strconv.FormatInt(userId, 10))

	// 从Redis中获取有序集合中指定范围的成员
	members, err := rdb.ZRange(ctx, followerKey, 0, -1).Result()
	if err != nil {
		zap.L().Error("Redis中获取FollowerList失败: ", zap.Error(err))
		return nil, err
	}

	// 创建切片，逐步将粉丝列表成员members转为int64的粉丝用户id
	followerList := make([]int64, len(members))
	for i, member := range members {
		id, _ := strconv.ParseInt(member, 10, 64)
		followerList[i] = id
	}

	// 设置key的过期时间
	_, err = rdb.Expire(ctx, followerKey, time.Duration(conf.Config.RedisConfig.RedisExpire)*time.Hour).Result()
	if err != nil {
		return nil, err
	}

	zap.L().Info("Redis缓存中获取followerList成功")
	return followerList, nil
}

func GetFriendList(ctx context.Context, userId int64) ([]int64, error) {
	zap.L().Info("start to get the friend list")

	// 根据userID获取关注列表和粉丝列表
	followList, err := GetFollowList(ctx, userId)
	if err != nil {
		zap.L().Error("Redis中获取FriendList失败: ", zap.Error(err))
		return nil, err
	}

	followerList, err := GetFollowerList(ctx, userId)
	if err != nil {
		zap.L().Error("Redis中获取FriendList失败: ", zap.Error(err))
		return nil, err
	}

	// 创建一个映射，将关注列表和粉丝列表的用户ID添加到映射中
	// 找到关注列表和粉丝列表的交集，即好友列表
	friendList := make([]int64, 0)
	for _, id := range followList {
		for _, followerID := range followerList {
			if id == followerID {
				friendList = append(friendList, id)
				break
			}
		}
	}

	zap.L().Info("Redis缓存中获取FriendList成功")
	return friendList, nil
}

func GetIsFollow(ctx context.Context, userId int64, toUserId int64) (bool, error) {
	// 构建key
	followKey := GetRedisKey(follow, strconv.FormatInt(userId, 10))

	//获取有序集合中指定成员的分值
	score, err := rdb.ZScore(ctx, followKey, strconv.FormatInt(toUserId, 10)).Result()
	if err != nil && err != redis.Nil {
		zap.L().Error("Redis中获取对应userId是否toUserId关注失败: ", zap.Error(err))
		return false, err
	}

	// 设置key的过期时间，并返回 布尔值和错误值
	if score == 0 {
		_, err = rdb.Expire(ctx, followKey, time.Duration(conf.Config.RedisConfig.RedisExpire)*time.Hour).Result()
		if err != nil {
			return false, err
		}
		return false, nil // Redis中没有该关系 且没有错误
	}

	_, err = rdb.Expire(ctx, followKey, time.Duration(conf.Config.RedisConfig.RedisExpire)*time.Hour).Result()
	if err != nil {
		return true, err // Redis中存在该关系，但出错
	}

	zap.L().Info("Redis缓存中存在该关系")
	return true, nil // Redis中存在该关系 且没有错误
}

// 检查给定的 userId 的关注列表是否存在
func IsFollowKeyExist(ctx context.Context, userId int64) (bool, error) {
	zap.L().Info("检查给定的 userId 的关注列表是否存在")

	// 构建key
	followKey := GetRedisKey(follow, strconv.FormatInt(userId, 10))

	// Exists() 判断键是否存在于Redis, 返回存在的键的数量
	result, err := rdb.Exists(ctx, followKey).Result()
	if err != nil {
		zap.L().Error("Redis中获取followKey失败: ", zap.Error(err))
		return false, err
	}

	// 设置key的过期时间，并返回 布尔值和错误值
	if result == 0 {
		_, err = rdb.Expire(ctx, followKey, time.Duration(conf.Config.RedisConfig.RedisExpire)*time.Hour).Result()
		if err != nil {
			return false, err
		}
		return false, nil
	}

	_, err = rdb.Expire(ctx, followKey, time.Duration(conf.Config.RedisConfig.RedisExpire)*time.Hour).Result()
	if err != nil {
		return true, err
	}

	zap.L().Info("Redis缓存中存在userid的followKey")
	return true, nil
}

// 检查给定的 userId 的粉丝列表是否存在
func IsFollowerKeyExist(ctx context.Context, userId int64) (bool, error) {
	zap.L().Info("检查给定的 userId 的粉丝列表是否存在")
	// 构建key
	followerKey := GetRedisKey(follower, strconv.FormatInt(userId, 10))

	// Exists() 判断键是否存在于Redis, 返回存在的键的数量
	result, err := rdb.Exists(ctx, followerKey).Result()
	if err != nil {
		zap.L().Error("Redis中获取followerKey失败: ", zap.Error(err))
		return false, err
	}

	// 设置key的过期时间，并返回 布尔值和错误值
	if result == 0 {
		_, err = rdb.Expire(ctx, followerKey, time.Duration(conf.Config.RedisConfig.RedisExpire)*time.Hour).Result()
		if err != nil {
			return false, err
		}
		return false, nil
	}

	_, err = rdb.Expire(ctx, followerKey, time.Duration(conf.Config.RedisConfig.RedisExpire)*time.Hour).Result()
	if err != nil {
		return true, err
	}

	zap.L().Info("Redis缓存中存在userid的followerKey")
	return true, nil
}
