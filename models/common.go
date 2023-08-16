package models

import (
	"net/http"

	"github.com/gin-gonic/gin"
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
