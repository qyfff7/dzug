package models

import (
	"dzug/protos/user"
	"github.com/gin-gonic/gin"
	"net/http"
)

type User struct {
	ID            int64  `json:"id"`
	Name          string `json:"name"`
	FollowCount   int64  `json:"follow_count"`   // 关注总数
	FollowerCount int64  `json:"follower_count"` // 粉丝总数
	WorkCount     int64  `json:"work_count"`
	FavoriteCount int64  `json:"favorite_count"`
	IsFollow      bool   `json:"is_follow"` // true-已关注，false-未关注
}

type AccountResp struct {
	Response
	UserId int64  `json:"user_id,omitempty"`
	Token  string `json:"token"`
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
