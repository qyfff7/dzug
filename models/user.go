package models

import (
	"dzug/protos/user"
	"github.com/gin-gonic/gin"
	"net/http"
)

type User struct {
	ID            int64  `json:"id"`
	UserId        int64  `json:"user_id"`
	Name          string `json:"name"`
	FollowCount   int64  `json:"follow_count"`   // 关注总数
	FollowerCount int64  `json:"follower_count"` // 粉丝总数
	WorkCount     int64  `json:"work_count"`
	FavoriteCount int64  `json:"favorite_count"`
	IsFollow      bool   `json:"is_follow"` // true-已关注，false-未关注

	Avatar          string `json:"avatar"`           //用户头像
	BackgroundImage string `json:"background_image"` //用户个人页顶部大图
	Signature       string `json:"signature"`        //个人简介
	TotalFavorited  int64  `json:"total_favorited"`  //获赞数量

}

type AccountResp struct {
	Response
	UserId int64  `json:"user_id,omitempty"`
	Token  string `json:"token"`
}
type GetUserInfoReq struct {
	UserID int64  `form:"user_id"`
	Token  string `form:"token"`
}

type GetUserInfoResp struct {
	Response
	User User `json:"user"`
}

func AccountRespSuccess(c *gin.Context, user *user.AccountResp) {
	c.JSON(http.StatusOK, AccountResp{
		Response: Response{
			StatusCode: CodeSuccess,
			StatusMsg:  CodeSuccess.Msg(),
		},
		UserId: user.UserId,
		Token:  user.Token,
	})
}

func UserInfoResp(u *user.GetUserInfoResp) User {
	userinfo := User{
		ID:              u.User.Id,
		Name:            u.User.Name,
		FollowCount:     u.User.FollowCount,
		FollowerCount:   u.User.FollowerCount,
		WorkCount:       u.User.WorkCount,
		FavoriteCount:   u.User.FavoriteCount,
		IsFollow:        u.User.IsFollow,
		Avatar:          u.User.Avatar,
		BackgroundImage: u.User.BackgroundImage,
		Signature:       u.User.Signature,
		TotalFavorited:  u.User.TotalFavorited,
	}
	return userinfo
}
func GetUserInfoSuccess(c *gin.Context, u User) {
	c.JSON(http.StatusOK, GetUserInfoResp{
		Response: Response{
			StatusCode: CodeSuccess,
			StatusMsg:  CodeSuccess.Msg(),
		},
		User: u,
	})
}
