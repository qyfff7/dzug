package routes

import (
	"dzug/app/gateway/handlers"
	"dzug/app/gateway/middlewares"
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

	ginRouter.LoadHTMLFiles("./templates/index.html")
	ginRouter.Static("/static", "./static")

	ginRouter.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": "Main website",
		})
	})

	tourist := ginRouter.Group("/douyin")
	{
		tourist.GET("/feed/", handlers.Feed)                   //视频流
		tourist.POST("/user/login/", handlers.UserLogin)       //用户登录路由
		tourist.POST("/user/register/", handlers.UserRegister) //用户注册路由
	}

	user := ginRouter.Group("/douyin")
	user.Use(middlewares.JWTAuthMiddleware())
	{
		user.GET("/user/", handlers.UserInfo) //用户信息路由
	}

	favorite := ginRouter.Group("/douyin/favorite")
	favorite.Use(middlewares.JWTAuthMiddleware())
	{
		favorite.POST("/action/", handlers.FavoriteAction)
		favorite.GET("/list/", handlers.FavoriteList)
	}

	ginRouter.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "404 not found",
		})
	})
	return ginRouter
}
