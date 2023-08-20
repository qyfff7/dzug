package models

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type FriendUser struct {
	Msg     string `json:"message"`
	MsgType int64  `json:"msgType"`
}

type GetUserInfoListResp struct {
	Response
	UserInfos []*User `json:"user_list"`
}

type GetFriendInfoListResp struct {
	Response
	FriendInfos []*FriendUser `json:"user_list"`
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

func FriendInfoListRespSuccess(c *gin.Context, list []*FriendUser) {
	c.JSON(http.StatusOK, GetFriendInfoListResp{
		Response: Response{
			StatusCode: CodeSuccess,
			StatusMsg:  CodeSuccess.Msg(),
		},
		FriendInfos: list,
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
