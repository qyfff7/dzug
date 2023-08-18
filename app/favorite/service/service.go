package service

import (
	"context"
	"dzug/app/favorite/dal/dao"
	"dzug/app/favorite/dal/kafka"
	"dzug/app/favorite/dal/redis"
	"dzug/protos/favorite"
	"errors"
	"go.uber.org/zap"
)

type FavoriteSrv struct {
	favorite.UnimplementedDouyinFavoriteActionServiceServer
}

// Favorite
// 1. 用户点击点赞按钮
// 2. 获取到用户id和视频id
// 3. 先去redis里寻找是否有集合 key favor:userId value video id
// 4. 有则再把value存入set中
// 5. 无则从mysql中查询点赞列表进行缓存（读取出这个user的所有点赞数据），然后redis再获取一次，有没有这个key和value，有这个key value就返回，没有就添加key value
// 6. 根据写入redis的情况，放入消息队列中，定时写入数据库中
// redis key：userId value：videoId
func (f *FavoriteSrv) Favorite(ctx context.Context, in *favorite.FavoriteRequest) (*favorite.FavoriteResponse, error) {
	userId := in.UserId
	videoId := in.VideoId
	ans := redis.AddFavor(userId, videoId)
	if ans == 0 {
		return &favorite.FavoriteResponse{
			StatusCode: 500,
			StatusMsg:  "服务器错误",
		}, errors.New("redis或mysql数据库错误")
	} else if ans == 2 {
		return &favorite.FavoriteResponse{ // 不需要写入消息队列
<<<<<<< HEAD
			StatusCode: 0,
=======
			StatusCode: 200,
>>>>>>> feat(-): message module
			StatusMsg:  "重复点赞操作",
		}, nil
	}
	err := kafka.Sender(userId, videoId, 1)
	if err != nil { // 放入消息队列失败，要删除redis中的value数据，保持和数据库的一致性
		if redis.DelFavor(userId, videoId) == 0 {
			zap.L().Error("redis或mysql数据库错误")
		}
		return nil, err
	}
	return &favorite.FavoriteResponse{
<<<<<<< HEAD
		StatusCode: 0,
=======
		StatusCode: 200,
>>>>>>> feat(-): message module
		StatusMsg:  "点赞成功",
	}, nil
}

// Infavorite 取消点赞
func (f *FavoriteSrv) Infavorite(ctx context.Context, in *favorite.InfavoriteRequest) (*favorite.InfavoriteResponse, error) {
	userId := in.UserId
	videoId := in.VideoId
	ans := redis.DelFavor(userId, videoId)
	if ans == 0 {
		return &favorite.InfavoriteResponse{
			StatusCode: 500,
			StatusMsg:  "服务器错误",
		}, errors.New("redis或mysql数据库错误")
	} else if ans == 2 {
		return &favorite.InfavoriteResponse{
<<<<<<< HEAD
			StatusCode: 0,
=======
			StatusCode: 200,
>>>>>>> feat(-): message module
			StatusMsg:  "重复取消点赞操作",
		}, nil
	}
	err := kafka.Sender(userId, videoId, 2)
	if err != nil { // 放入消息队列失败，把key value放回redis中，保持和数据库中的一致性
		if redis.AddFavor(userId, videoId) == 0 { // 不管是1还是2，都达成目标
			zap.L().Error("redis或mysql数据库错误")
		}
		return nil, err
	}
	return &favorite.InfavoriteResponse{
<<<<<<< HEAD
		StatusCode: 0,
=======
		StatusCode: 200,
>>>>>>> feat(-): message module
		StatusMsg:  "取消点赞成功",
	}, nil
}

// FavoriteList 获取点赞列表
func (f *FavoriteSrv) FavoriteList(ctx context.Context, in *favorite.FavoriteListRequest) (*favorite.FavoriteListResponse, error) {
	videoIds, err := redis.GetVideosByUserId(in.UserId) // 根据用户id获取点赞视频id
	if err != nil {
		zap.L().Error("redis获取点赞视频id列表失败")
		return nil, err
	}
	videos, err := dao.GetVideosByVideoIds(in.UserId, videoIds) // 根据点赞视频id获取视频信息，并且介为返回的视频信息
	if err != nil {
		zap.L().Error("获取点赞视频列表失败")
		return &favorite.FavoriteListResponse{
			StatusCode: 500,
			StatusMsg:  "获取点赞视频列表失败",
		}, nil
	}
	return &favorite.FavoriteListResponse{
		StatusCode: 200,
		StatusMsg:  "获取点赞视频列表成功",
		VideoList:  videos,
	}, nil
}
