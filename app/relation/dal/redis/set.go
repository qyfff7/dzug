/*
实现方法：
func AddRelation(userId int64, toUserId int64) 增加关系
func DeleteRelation(userId int64, toUserId int64) 删除关系
func PutFollowList(userId int64, FollowList []int64) 将数据库数据写入到redis中
func PutFollowerList(userId int64, FollowerList []int64) 将数据库文件写入到redis中
*/
package redis

import (
	"context"
	"dzug/conf"
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/redis/go-redis/v9"

	"go.uber.org/zap"
)

// 先写redis 再写mysql的整体框架

// 在Redis中添加一个关系
func AddRelation(ctx context.Context, userId int64, toUserId int64) error {
	// 1. 构建关注关系followKey 和 被关注followerKey，获取当前时间戳
	if userId == toUserId {
		return errors.New("you can't follow yourself")
	}
	followKey := GetRedisKey(follow, strconv.FormatInt(userId, 10))
	followerKey := GetRedisKey(follower, strconv.FormatInt(toUserId, 10))

	// 2. 判断redis中 该关注关系是否已存在
	_, err := rdb.ZScore(ctx, followKey, strconv.FormatInt(toUserId, 10)).Result()
	if err != nil {
		if err == redis.Nil {
			zap.L().Info("关系不存在，可以添加")
			// 3. 指定有序集合成员及其分数的结构体
			now := time.Now().Unix()
			followMember := redis.Z{
				Score:  float64(now),
				Member: strconv.FormatInt(toUserId, 10),
			}
			followerMember := redis.Z{
				Score:  float64(now),
				Member: strconv.FormatInt(userId, 10),
			}
			// 4. 添加关注关系和粉丝关系
			if err := rdb.ZAdd(ctx, followKey, followMember).Err(); err != nil {
				zap.L().Error("redis添加关注关系失败", zap.Error(err))
				return err
			}
			if err := rdb.ZAdd(ctx, followerKey, followerMember).Err(); err != nil {
				zap.L().Error("redis添加粉丝关系失败", zap.Error(err))
				rdb.ZRem(ctx, followKey, toUserId) // 如果粉丝关系添加失败，则移除关注关系中的对应成员
				return err
			}
			// 5.设置followkey和followerkey的过期时间
			_, err = rdb.Expire(ctx, followKey, time.Duration(conf.Config.RedisConfig.RedisExpire)*time.Hour).Result()
			if err != nil {
				zap.L().Error("redis设置followKey过期时间失败", zap.Error(err))
				return err
			}
			_, err = rdb.Expire(ctx, followerKey, time.Duration(conf.Config.RedisConfig.RedisExpire)*time.Hour).Result()
			if err != nil {
				zap.L().Error("redis设置followerKey过期时间失败", zap.Error(err))
				return err
			}

			// ***************************** 6. 增加用户的两个计数器：关注数 和 粉丝数 *****************************
			// 使用 HINCRBY 命令增加关注者的 follow_count
			_, err := rdb.HIncrBy(ctx, GetRedisKey(KeyUserInfo, strconv.FormatInt(userId, 10)), "follow_count", 1).Result()
			if err != nil {
				fmt.Println("redis增加关注者的follow_count失败", zap.Error(err))
				return err
			}

			// 使用 HINCRBY 命令增加粉丝的 follower_count
			_, err = rdb.HIncrBy(ctx, GetRedisKey(KeyUserInfo, strconv.FormatInt(toUserId, 10)), "follower_count", 1).Result()
			if err != nil {
				fmt.Println("redis增加粉丝的follower_count失败", zap.Error(err))
				// 将之前增加的关注者的 follow_count 恢复
				_, _ = rdb.HIncrBy(ctx, GetRedisKey(KeyUserInfo, strconv.FormatInt(userId, 10)), "follow_count", -1).Result()
				return err
			}
		} else {
			// 发生了其他错误
			fmt.Println("发生了错误", zap.Error(err))
			return err
		}
	} else {
		// 关注关系已存在
		zap.L().Error("redis中关系已经存在，添加失败", zap.Error(err))
		return errors.New("redis中关系已经存在，添加失败")
	}

	return nil
}

// 在Redis中删除一段关系
func DeleteRelation(ctx context.Context, userId int64, toUserId int64) error {
	// 1. 构建关注关系followKey 和 被关注followerKey，获取当前时间戳
	if userId == toUserId {
		return errors.New("you can't unfollow yourself")
	}
	followKey := GetRedisKey(follow, strconv.FormatInt(userId, 10))
	followerKey := GetRedisKey(follower, strconv.FormatInt(toUserId, 10))

	// 2. 判断 redis 中关系是否存在
	exists, err := rdb.ZScore(ctx, followKey, strconv.FormatInt(toUserId, 10)).Result()
	if err != nil {
		if err == redis.Nil {
			zap.L().Error("redis中关系不存在，删除失败", zap.Error(err))
			return errors.New("redis中关系不存在，删除失败")
		} else {
			// 发生了其他错误
			zap.L().Error("发生了错误", zap.Error(err))
			return err
		}
	} else {
		zap.L().Info("关注关系已存在，可以进行删除操作")
		// 3. 从有序集合中删除关注关系和被关注关系
		_, err := rdb.ZRem(ctx, followKey, strconv.FormatInt(toUserId, 10)).Result()
		if err != nil {
			zap.L().Error("redis删除关注关系失败", zap.Error(err))
			return err
		}

		_, err = rdb.ZRem(ctx, followerKey, strconv.FormatInt(userId, 10)).Result()
		if err != nil {
			rdb.ZAdd(ctx, followKey, redis.Z{Score: exists, Member: strconv.FormatInt(toUserId, 10)}) // 如果删除粉丝关系失败，将关注关系还原
			zap.L().Error("redis删除粉丝关系失败", zap.Error(err))
			return err
		}

		// 4.设置followkey和followerkey的过期时间
		_, err = rdb.Expire(ctx, followKey, time.Duration(conf.Config.RedisConfig.RedisExpire)*time.Hour).Result()
		if err != nil {
			zap.L().Error("redis设置followKey过期时间失败", zap.Error(err))
			return err
		}
		_, err = rdb.Expire(ctx, followerKey, time.Duration(conf.Config.RedisConfig.RedisExpire)*time.Hour).Result()
		if err != nil {
			zap.L().Error("redis设置followerKey过期时间失败", zap.Error(err))
			return err
		}

		// 5. ***************************** 减少用户的两个计数器：关注数和粉丝数  *****************************
		// 使用 HINCRBY 命令减少关注者的 follow_count
		_, err = rdb.HIncrBy(ctx, GetRedisKey(KeyUserInfo, strconv.FormatInt(userId, 10)), "follow_count", -1).Result()
		if err != nil {
			fmt.Println("redis减少关注者的follow_count失败", zap.Error(err))
			return err
		}

		// 使用 HINCRBY 命令减少粉丝的 follower_count
		_, err = rdb.HIncrBy(ctx, GetRedisKey(KeyUserInfo, strconv.FormatInt(toUserId, 10)), "follower_count", -1).Result()
		if err != nil {
			fmt.Println("redis减少粉丝的follower_count失败", zap.Error(err))
			// 将之前减少的关注者的 follow_count 恢复
			_, _ = rdb.HIncrBy(ctx, GetRedisKey(KeyUserInfo, strconv.FormatInt(userId, 10)), "follow_count", 1).Result()
			return err
		}
	}

	return nil
}

// 把FollowList 关注用户列表数据写回redis
func PutFollowList(ctx context.Context, userId int64, FollowIdList []int64) error {

	// 特判：如果FollowIdList为空，直接返回nil
	if FollowIdList == nil || len(FollowIdList) == 0 {
		return nil
	}

	// 1. 构建关注列表的键
	followKey := GetRedisKey(follow, strconv.FormatInt(userId, 10))

	// 2. 使用管道一次性执行多个命令，设置过期时间并将关注列表中的用户ID添加到集合中
	pipe := rdb.Pipeline()

	pipe.Expire(ctx, followKey, time.Duration(conf.Config.RedisConfig.RedisExpire)*time.Hour) //设置过期时间

	// 3. 将关注列表中的用户ID添加到 Redis 集合中
	now := float64(time.Now().Unix())
	for _, uid := range FollowIdList {
		followMember := redis.Z{
			Score:  float64(now),
			Member: strconv.FormatInt(uid, 10),
		}
		pipe.ZAdd(ctx, followKey, followMember)
	}

	// 5. 执行 Redis 管道中的所有命令
	_, err := pipe.Exec(ctx)
	if err != nil {
		zap.L().Error("将followerList写回Redis, 管道执行失败", zap.Error(err))
		return err
	}

	return nil
}

// 把FollowerList 粉丝列表数据写回redis
func PutFollowerList(ctx context.Context, userId int64, FollowerIdList []int64) error {
	// 特判：如果 FollowerList 为空，直接返回 nil
	if FollowerIdList == nil || len(FollowerIdList) == 0 {
		return nil
	}

	// 1. 构建粉丝列表的键
	followerKey := GetRedisKey(follower, strconv.FormatInt(userId, 10))

	// 2. 使用管道一次性执行多个命令，设置过期时间并将粉丝列表中的用户ID添加到集合中
	pipe := rdb.Pipeline()

	// 设置过期时间
	pipe.Expire(ctx, followerKey, time.Duration(conf.Config.RedisConfig.RedisExpire)*time.Hour)

	// 使用 SAdd 命令将粉丝列表中的用户ID添加到 Redis 集合中
	now := float64(time.Now().Unix())
	for _, uid := range FollowerIdList {
		followerMember := redis.Z{
			Score:  float64(now),
			Member: strconv.FormatInt(uid, 10),
		}
		pipe.ZAdd(ctx, followerKey, followerMember)
	}

	// 执行 Redis 管道中的所有命令
	_, err := pipe.Exec(ctx)
	if err != nil {
		return err
	}

	return nil
}
