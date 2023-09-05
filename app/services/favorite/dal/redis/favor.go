package redis

import (
	"context"
	"dzug/app/services/favorite/dal/dao"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
	"strconv"
	"time"
)

var faPrefix = "favor:" // redis中favor前缀
var key string          // redis的key

// AddFavor 点赞操作
// 1. 查询redis中是否有userid这个set，没有则从mysql中读取加入缓存，这部分出错返回，则返回代码0
// 2. 把videoId加入set，如果加入成功，返回nil，并在service的函数中加入消息队列，等待写入mysql中 返回1
// 3. 如果加入失败，redis返回0，即这个video id 已经存在则返 2
func AddFavor(userId, videoId int64) int {
	cmd, err := exist(userId)
	if err != nil {
		return 0
	}
	cmd = Rdb.SAdd(context.Background(), key, videoId) // 这个key现在已经存在了，去添加这个videoId
	if cmd.Val() == 0 {                                // 已经存在这个value了
		return 2
	}
	return 1 // 理论上为1就是操作正确，应该不会有其他意外情况
}

// DelFavor 取消点赞操作
func DelFavor(userId, videoId int64) int {
	cmd, err := exist(userId)
	if err != nil {
		return 0
	}
	cmd = Rdb.SRem(context.Background(), key, videoId) // 这个key现在已经存在了，去删除这个videoId
	if cmd.Val() == 0 {                                // 已经存在这个value了
		return 2
	}
	return 1
}

// GetVideosByUserId 根据用户id获取用户的喜欢列表
func GetVideosByUserId(userId int64) ([]int64, error) {
	_, err := exist(userId)
	if err != nil {
		return nil, err
	}
	cmd := Rdb.SMembers(context.Background(), key)
	videos := make([]int64, len(cmd.Val()))
	for k, v := range cmd.Val() {
		value, _ := strconv.Atoi(v)
		videos[k] = int64(value)
	}
	return videos, err
}

// 确保当前key存在redis中
func exist(userId int64) (*redis.IntCmd, error) {
	key = faPrefix + strconv.FormatInt(userId, 10)
	cmd := Rdb.Exists(context.Background(), key)
	if cmd.Err() != nil {
		zap.L().Error("查询redis失败")
		return nil, cmd.Err()
	}
	if cmd.Val() == 0 {
		zap.L().Debug("当前user不存在redis，正在查找数据库")
		err := getSet(userId)
		if err != nil {
			return nil, err
		}
	}
	return cmd, nil
}

// getSet 从数据库中得到该user的所有点赞视频数据，写入redis中
func getSet(userId int64) error {
	videoIds, err := dao.GetFavorById(userId)
	if err != nil {
		zap.L().Error("获取用户" + strconv.FormatInt(userId, 10) + "点赞数据失败")
		return err
	}
	//key = faPrefix + strconv.FormatInt(userId, 10)
	ctx := context.Background()
	for _, v := range videoIds {
		Rdb.SAdd(ctx, key, v)
	}
	Rdb.Expire(ctx, key, time.Hour*3) // 设置过期时间
	return nil
}
