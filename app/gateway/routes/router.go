package routes

import (
	"dzug/app/gateway/handlers"
	"dzug/logger"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func NewRouter(mode string) *gin.Engine {

	// 初始化gin框架内置的校验器使用的翻译器
	if err := handlers.InitTrans("zh"); err != nil {
		fmt.Printf("init validator trans failed, err:%v\n", err)
		return nil
	}

	//如果配置文件中的mode设置为release模式，则gin框架也设置为发布模式(终端不输出任何信息)
	if mode == gin.ReleaseMode {
		gin.SetMode(gin.ReleaseMode) // gin设置成发布模式
	}

	ginRouter := gin.New()
	ginRouter.Use(logger.GinLogger(), logger.GinRecovery(true)) // 使用自己的两个中间件

	user := ginRouter.Group("/douyin/user")
	{
		user.GET("/", handlers.UserInfo) //用户信息路由
		//user.POST("/login", handlers.UserLogin)       //用户登录路由
		user.POST("/register", handlers.UserRegister) //用户注册路由
	}

	ginRouter.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "404",
		})
	})
	return ginRouter
}
