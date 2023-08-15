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

	resp, err = discovery.FavoriteClient.FavoriteList(ctx, req)
	return
}

//type Videos struct {
//	Id            int64  `json:"id"`             // 视频唯一标识
//	Author        *Users  `json:"author"`         // 视频作者信息
//	PlayUrl       string `json:"play_url"`       // 视频播放地址
//	CoverUrl      string `json:"cover_url"`      // 视频封面地址
//	FavoriteCount int64  `json:"favorite_count"` // 视频的点赞总数
//	CommentCount  int64  `json:"comment_count"`  // 视频的评论总数
//	IsFavorite    bool   `json:"is_favorite"`    // true-已点赞，false-未点赞
//	Title         string `json:"title"`
//}
//
//type Users struct {
//	Id              int64  `json:"id"`               // 用户id
//	Name            string `json:"name"`             // 用户名称
//	FollowCount     int64  `json:"follow_count"`     // 关注总数
//	FollowerCount   int64  `json:"follower_count"`   // 粉丝总数
//	IsFollow        bool   `json:"is_follow"`        // true-已关注，false-未关注
//	Avatar          string `json:"avatar"`           //用户头像
//	BackgroundImage string `json:"background_image"` //用户个人页顶部大图
//	Signature       string `json:"signature"`        //个人简介
//	TotalFavorited  int64  `json:"total_favorited"`  //获赞数量
//	WorkCount       int64  `json:"work_count"`       //作品数量
//	FavoriteCount   int64  `json:"favorite_count"`   //点赞数量
//}
