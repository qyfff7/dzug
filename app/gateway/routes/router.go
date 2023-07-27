package routes

import (
	"dzug/app/gateway/http"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	ginRouter := gin.Default()
	v1 := ginRouter.Group("/douyin/user")
	{
		v1.POST("/login", http.UserLogin)
		v1.POST("/register", http.UserRegister)
	}

	return ginRouter
}
