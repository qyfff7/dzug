package service

import (
	"context"
	"dzug/app/favorite/dal/dao"
	"dzug/protos/favorite"
	"fmt"
)

type FavoriteSrv struct {
	favorite.UnimplementedDouyinFavoriteActionServiceServer
}

// Favorite 点赞操作（测试用）
// 先从redis里获取点赞记录，如果点赞记录key存在在redis中，直接添加value
// 如果没有这个key，就从数据库里读取，确保有这个key，添加到redis中后，再进行更改
// redis key：userId value：videoId
func (f *FavoriteSrv) Favorite(ctx context.Context, in *favorite.FavoriteRequest) (*favorite.FavoriteResponse, error) {
	userId := uint64(in.UserId)
	videoId := uint64(in.VideoId)
	//keys := redis.Rdb.Keys(context.Background(), strconv.FormatUint(userId, 10))

	err := dao.Favorite(videoId, userId)
	if err != nil {
		fmt.Println("错误为：", err.Error())
		return &favorite.FavoriteResponse{
			StatusCode: 400,
			StatusMsg:  "点赞失败",
		}, err
	}
	return &favorite.FavoriteResponse{
		StatusCode: 200,
		StatusMsg:  "点赞成功",
	}, nil
}

// Infavorite 取消点赞
func (f *FavoriteSrv) Infavorite(ctx context.Context, in *favorite.InfavoriteRequest) (*favorite.InfavoriteResponse, error) {
	return &favorite.InfavoriteResponse{
		StatusCode: 200,
		StatusMsg:  "调用成功，你成功进行了一次取消收藏操作",
	}, nil
}

// FavoriteList 获取点赞列表
func (f *FavoriteSrv) FavoriteList(ctx context.Context, in *favorite.FavoriteListRequest) (*favorite.FavoriteListResponse, error) {
	return &favorite.FavoriteListResponse{
		StatusCode: 200,
		StatusMsg:  "调用成功，你成功进行了一次收藏列表操作",
	}, nil
}
