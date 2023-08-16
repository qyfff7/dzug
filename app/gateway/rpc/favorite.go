package rpc

import (
	"context"
	"dzug/discovery"
	"dzug/protos/favorite"
	"go.uber.org/zap"
)

// FavoriteAction 点赞操作
func FavoriteAction(ctx context.Context, req *favorite.FavoriteRequest) (resp *favorite.FavoriteResponse, err error) {
	resp = &favorite.FavoriteResponse{}
	err = discovery.LoadClient("favorite", &discovery.FavoriteClient)
	if err != nil {
		zap.L().Error("调用点赞服务失败")
		resp.StatusCode = 500
		resp.StatusMsg = "点赞失败"
		return
	}
	resp, err = discovery.FavoriteClient.Favorite(ctx, req)
	return
}

// InFavorite 取消点赞操作
func InFavorite(ctx context.Context, req *favorite.InfavoriteRequest) (resp *favorite.InfavoriteResponse, err error) {
	resp = &favorite.InfavoriteResponse{}
	err = discovery.LoadClient("favorite", &discovery.FavoriteClient)
	if err != nil {
		zap.L().Error("调用取消点赞服务失败")
		resp.StatusCode = 500
		resp.StatusMsg = "取消点赞失败"
		return
	}

	resp, err = discovery.FavoriteClient.Infavorite(ctx, req)
	return
}

// FavoriteList 获取点赞列表
func FavoriteList(ctx context.Context, req *favorite.FavoriteListRequest) (resp *favorite.FavoriteListResponse, err error) {
	resp = &favorite.FavoriteListResponse{}
	err = discovery.LoadClient("favorite", &discovery.FavoriteClient)
	if err != nil {
		zap.L().Error("获取点赞列表服务失败")
		resp.StatusCode = 500
		resp.StatusMsg = "获取点赞列表失败"
		return
	}
	// 将resp转为response类型

	resp, err = discovery.FavoriteClient.FavoriteList(ctx, req)
	return
}
