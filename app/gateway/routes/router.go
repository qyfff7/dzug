package routes

import (
	"dzug/app/gateway/http"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	ginRouter := gin.Default()
	user := ginRouter.Group("/douyin/user")
	{
		user.POST("/login", http.UserLogin)
		user.POST("/register", http.UserRegister)
	}
	relation := ginRouter.Group("/douyin/relation")
	{
		relation.POST("/action", http.RelationAction)
	}

	return ginRouter
}
