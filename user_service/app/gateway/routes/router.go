package routes

import (
	"dzug/user_service/Log_Conf/logger"
	handlers2 "dzug/user_service/app/gateway/handlers"
	"github.com/gin-gonic/gin"
)

func NewRouter(mode string) *gin.Engine {

	if mode == gin.ReleaseMode {
		gin.SetMode(gin.ReleaseMode) // gin设置成发布模式
	}
	//创建web服务时，不使用默认的gin.Default()
	ginRouter := gin.New()
	//使用自己写的两个中间件，
	ginRouter.Use(logger.GinLogger(), logger.GinRecovery(true))

	user := ginRouter.Group("/douyin/user")
	{
		user.POST("/login", handlers2.UserLogin)
		user.POST("/register", handlers2.UserRegister)
	}

	return ginRouter
}
