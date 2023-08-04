package service

import (
	"context"
	"dzug/protos/favorite"
	"dzug/repo"
	"fmt"
)

type FavoriteSrv struct {
	favorite.UnimplementedDouyinFavoriteActionServiceServer
}

// Favorite 点赞操作（测试用）
func (f *FavoriteSrv) Favorite(ctx context.Context, in *favorite.FavoriteRequest) (*favorite.FavoriteResponse, error) {
	if repo.DB == nil {
		panic("fuck")
	}
	fa := &repo.Favorite{
		UserId:  uint64(in.UserId),
		VideoId: uint64(in.VideoId),
	}
	fmt.Println(fa)
	fmt.Println(in.UserId)
	res := repo.DB.Create(fa)
	if res.Error != nil {
		fmt.Println("错误为：", res.Error.Error())
		return &favorite.FavoriteResponse{
			StatusCode: 400,
			StatusMsg:  "点赞失败",
		}, res.Error
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
