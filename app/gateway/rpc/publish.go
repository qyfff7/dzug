package rpc

import (
	"context"
	"dzug/discovery"
	pb "dzug/protos/publish"
	"go.uber.org/zap"
	"strconv"
)

type Author struct {
	ID              int64  `json:"id"`
	Name            string `json:"name"`
	FollowCount     int64  `json:"follow_count"`
	FollowerCount   int64  `json:"follower_count"`
	IsFollow        bool   `json:"is_follow"`
	Avatar          string `json:"avatar"`
	BackgroundImage string `json:"background_image"`
	Signature       string `json:"signature"`
	TotalFavorited  string `json:"total_favorited"`
	WorkCount       int64  `json:"work_count"`
	FavoriteCount   int64  `json:"favorite_count"`
}

type VideoL struct {
	ID            int64  `json:"id"`
	Author        Author `json:"author"`
	PlayURL       string `json:"play_url"`
	CoverURL      string `json:"cover_url"`
	FavoriteCount int64  `json:"favorite_count"`
	CommentCount  int64  `json:"comment_count"`
	IsFavorite    bool   `json:"is_favorite"`
	Title         string `json:"title"`
}

type Response struct {
	StatusCode int64    `json:"status_code"`
	StatusMsg  string   `json:"status_msg"`
	VideoList  []VideoL `json:"video_list"`
}

func PublishVideo(ctx context.Context, req *pb.PublishVideoReq) (resp *pb.PublishVideoResp, err error) {
	err = discovery.LoadClient("publish", &discovery.PublishClient)
	if err != nil {
		return nil, err
	}
	r, err := discovery.PublishClient.PublishVideo(ctx, req)
	if err != nil {
		zap.L().Error(err.Error())
		return
	}
	return r, nil
}

func GetPublishListByUser(ctx context.Context, req *pb.GetVideoListByUserIdReq) (resp *Response, err error) {
	err = discovery.LoadClient("publish", &discovery.PublishClient)
	if err != nil {
		zap.L().Error("加载服务发现出错")
		return nil, err
	}
	r, err := discovery.PublishClient.GetVideoListByUserId(ctx, req)
	vL := r.VideoList
	var rtnVl []VideoL
	for _, v := range vL {
		tmpU := &Author{
			ID:              v.Author.Id,
			Name:            v.Author.Name,
			FollowCount:     *v.Author.FollowCount,
			FollowerCount:   *v.Author.FollowerCount,
			Avatar:          *v.Author.Avatar,
			BackgroundImage: *v.Author.BackgroundImage,
			Signature:       *v.Author.Signature,
			TotalFavorited:  strconv.FormatInt(*v.Author.TotalFavorited, 10),
			WorkCount:       *v.Author.WorkCount,
			FavoriteCount:   *v.Author.FavoriteCount,
		}
		tmpV := VideoL{
			Author:        *tmpU,
			ID:            v.Id,
			PlayURL:       v.PlayUrl,
			CoverURL:      v.CoverUrl,
			FavoriteCount: v.FavoriteCount,
			CommentCount:  v.CommentCount,
			Title:         v.Title,
		}
		rtnVl = append(rtnVl, tmpV)
	}

	respRtn := &Response{
		VideoList:  rtnVl,
		StatusCode: 0,
		StatusMsg:  "Success",
	}
	if err != nil {
		zap.L().Error(err.Error())
		return
	}
	return respRtn, nil
}
