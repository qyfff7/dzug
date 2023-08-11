package models

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type ResponseData struct {
	Code RespCode    `json:"status_code"`
	Msg  string      `json:"status_msg"`
	Data interface{} `json:"data,omitempty"`
}

func ResponseError(c *gin.Context, code RespCode, data interface{}) {
	c.JSON(http.StatusInternalServerError, &ResponseData{
		Code: code,
		Msg:  code.Msg(),
		Data: data,
	})
}

func ResponseErrorWithMsg(c *gin.Context, code RespCode, msg string, data interface{}) {
	c.JSON(http.StatusInternalServerError, &ResponseData{
		Code: code,
		Msg:  msg,
		Data: data,
	})
}
func ResponseSuccess(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, &ResponseData{
		Code: CodeSuccess,
		Msg:  CodeSuccess.Msg(),
		Data: data,
	})
}
