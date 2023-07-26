package routes

import "github.com/gin-gonic/gin"

func NewRouter() *gin.Engine {
	ginRouter := gin.Default()
	v1 := ginRouter.Group("/douyin/user")
	{
		v1.POST("/register", func(ctx *gin.Context) {

		})
	}

	return ginRouter
}
