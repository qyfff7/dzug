package routes

import (
	"dzug/app/gateway/handlers"
	"dzug/logger"
	"github.com/gin-gonic/gin"
)

func NewRouter(mode string) *gin.Engine {
	//如果配置文件中的mode设置为release模式，则gin框架也设置为发布模式(终端不输出任何信息)
	if mode == gin.ReleaseMode {
		gin.SetMode(gin.ReleaseMode) // gin设置成发布模式
	}

	ginRouter := gin.New()
	ginRouter.Use(logger.GinLogger(), logger.GinRecovery(true)) // 使用自己的两个中间件

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
