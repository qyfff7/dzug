package routes

import (
	"dzug/app/gateway/handlers"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	ginRouter := gin.Default()
	user := ginRouter.Group("/douyin/user")
	{
		user.POST("/login", handlers.UserLogin)
		user.POST("/register", handlers.UserRegister)
	}
	relation := ginRouter.Group("/douyin/relation")
	{
		relation.POST("/action", handlers.RelationAction)
	}

	return ginRouter
}
