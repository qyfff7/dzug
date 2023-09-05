package models

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Message Gorm Data Structures

type Message struct {
	FromUserId  int64  `json:"from_user_id"`
	ToUserId    int64  `json:"to_user_id"`
	Contents    string `json:"content"`
	CreateTime  int64  `json:"create_time"`
	MessageUUID int64  `json:"id"`
}

type Thread struct {
	ThreadId     string `json:"thread_id"`
	LastPullTime int64  `json:"last_pull_time"`
}

type GetMessageListResp struct {
	Response
	Msg []*Message `json:"message_list"`
}

type PostMessageResp struct {
	Response
}

func MessageListRespSuccess(c *gin.Context, list []*Message) {
	c.JSON(http.StatusOK, GetMessageListResp{
		Response: Response{
			StatusCode: CodeSuccess,
			StatusMsg:  CodeSuccess.Msg(),
		},
		Msg: list,
	})
}

func PostMessageRespSuccess(c *gin.Context) {
	c.JSON(http.StatusOK, PostMessageResp{
		Response: Response{
			StatusCode: CodeSuccess,
			StatusMsg:  CodeSuccess.Msg(),
		},
	})
}
