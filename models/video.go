package models

import (
	"dzug/protos/video"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Video struct {
	Id            int64  `json:"id,omitempty"`
	Author        User   `json:"user_id,omitempty"`
	PlayUrl       string `json:"play_url,omitempty"`
	CoverUrl      string `json:"cover_url,omitempty"`
	FavoriteCount int64  `json:"favorite_count,omitempty"`
	CommentCount  int64  `json:"comment_count,omitempty"`
	IsFavorite    bool   `json:"is_favorite,omitempty"`
	Title         string `json:"title,omitempty"`
}

// Feed
type Feed struct {
	Response
	NextTime  int64    `json:"nextTime,omitempty"`
	VideoList []*Video `json:"video_list,omitempty"`
}

type VideoList struct {
}

func GetFeedResp(v *video.Video, author User) *Video {
	vv := &Video{
		Id:            v.VideoId,
		Author:        author,
		PlayUrl:       v.PlayUrl,
		CoverUrl:      v.CoverUrl,
		FavoriteCount: v.FavoriteCount,
		CommentCount:  v.CommentCount,
		IsFavorite:    v.IsFavorite,
		Title:         v.Title,
	}
	return vv
}
func GetFeedSuccess(c *gin.Context, v []*Video, next int64) {
	c.JSON(http.StatusOK, Feed{
		Response: Response{
			StatusCode: CodeSuccess,
			StatusMsg:  CodeSuccess.Msg(),
		},
		NextTime:  next,
		VideoList: v,
	})
}
