package rpc

import (
	"context"
	"dzug/discovery"
	"dzug/protos/favorite"
)

// FavoriteAction 点赞操作
func FavoriteAction(ctx context.Context, req *favorite.FavoriteRequest) (resp *favorite.FavoriteResponse, err error) {
	discovery.LoadClient("favorite", &discovery.FavoriteClient)

	r, err := discovery.FavoriteClient.Favorite(ctx, req)
	if err != nil {
		return
	}
	return r, nil
}

// InFavorite 取消点赞操作
func InFavorite(ctx context.Context, req *favorite.InfavoriteRequest) (resp *favorite.InfavoriteResponse, err error) {
	discovery.LoadClient("favorite", &discovery.FavoriteClient)

	r, err := discovery.FavoriteClient.Infavorite(ctx, req)
	if err != nil {
		return
	}
	return r, nil
}

// FavoriteList 获取点赞列表
func FavoriteList(ctx context.Context, req *favorite.FavoriteListRequest) (resp *favorite.FavoriteListResponse, err error) {
	discovery.LoadClient("favorite", &discovery.FavoriteClient)

	r, err := discovery.FavoriteClient.FavoriteList(ctx, req)
	if err != nil {
		return
	}
	return r, nil
}
