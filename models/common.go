package models

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	StatusCode RespCode `json:"status_code"`
	StatusMsg  string   `json:"status_msg,omitempty"`
}

func ResponseError(c *gin.Context, code RespCode) {
	c.JSON(http.StatusInternalServerError, &Response{
		StatusCode: code,
		StatusMsg:  code.Msg(),
	})
}

func ResponseErrorWithMsg(c *gin.Context, code RespCode, msg string) {
	c.JSON(http.StatusInternalServerError, &Response{
		StatusCode: code,
		StatusMsg:  msg,
	})
}

type Comment struct {
	Id         int64  `json:"id,omitempty"`
	User       User   `json:"user"`
	Content    string `json:"content,omitempty"`
	CreateDate string `json:"create_date,omitempty"`
}

type Message struct {
	Id         int64  `json:"id,omitempty"`
	Content    string `json:"content,omitempty"`
	CreateTime string `json:"create_time,omitempty"`
}

type MessageSendEvent struct {
	UserId     int64  `json:"user_id,omitempty"`
	ToUserId   int64  `json:"to_user_id,omitempty"`
	MsgContent string `json:"msg_content,omitempty"`
}

type MessagePushEvent struct {
	FromUserId int64  `json:"user_id,omitempty"`
	MsgContent string `json:"msg_content,omitempty"`
}
