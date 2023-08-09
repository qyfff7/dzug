package handlers

import (
	"dzug/app/user/pkg/jwt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func List(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{
		"msg": "cishi",
	})
}

// Feed 视频流
func Feed(c *gin.Context) {

	//1.获取当前请求中的token
	authHeader := c.Request.Header.Get("Authorization")
	parts := strings.SplitN(authHeader, " ", 2)
	token := parts[1]
	//2.从token中解析处userID
	u, err := jwt.ParseToken(token)
	if err != nil {
		//错误处理
		return
	}
	userID := u.UserID
	//此处userID就是当前用户的ID

	c.JSON(http.StatusOK, gin.H{
		"msg": userID,
	})

}
