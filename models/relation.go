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
