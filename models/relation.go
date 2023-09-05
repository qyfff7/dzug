package models

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type GetUserInfoListResp struct {
	Response
	UserInfos []*User `json:"user_list"`
}

type PostRelationResp struct {
	Response
}

func UserInfoListRespSuccess(c *gin.Context, list []*User) {
	c.JSON(http.StatusOK, GetUserInfoListResp{
		Response: Response{
			StatusCode: CodeSuccess,
			StatusMsg:  CodeSuccess.Msg(),
		},
		UserInfos: list,
	})
}

func PostRelationRespSuccess(c *gin.Context) {
	c.JSON(http.StatusOK, PostMessageResp{
		Response: Response{
			StatusCode: CodeSuccess,
			StatusMsg:  CodeSuccess.Msg(),
		},
	})
}

// 以下用于UserFriend类型
type UserFriend struct {
	ID              int64  `json:"id"`
	UserId          int64  `json:"user_id"`
	Name            string `json:"name"`
	FollowCount     int64  `json:"follow_count"`   // 关注总数
	FollowerCount   int64  `json:"follower_count"` // 粉丝总数
	WorkCount       int64  `json:"work_count"`
	FavoriteCount   int64  `json:"favorite_count"`
	IsFollow        bool   `json:"is_follow"`        // true-已关注，false-未关注
	Avatar          string `json:"avatar"`           //用户头像
	BackgroundImage string `json:"background_image"` //用户个人页顶部大图
	Signature       string `json:"signature"`        //个人简介
	TotalFavorited  int64  `json:"total_favorited"`  //获赞数量
	Message         string `json:"message"`
	MsgType         int64  `json:"msg_type"`
}

type GetUserFriendInfoListResp struct {
	Response
	UserInfos []*UserFriend `json:"user_list"`
}

func UserFriendInfoListRespSuccess(c *gin.Context, list []*UserFriend) {
	c.JSON(http.StatusOK, GetUserFriendInfoListResp{
		Response: Response{
			StatusCode: CodeSuccess,
			StatusMsg:  CodeSuccess.Msg(),
		},
		UserInfos: list,
	})
}
