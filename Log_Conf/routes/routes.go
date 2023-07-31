package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"zap-viper/Log_Conf/logger"
)

func SetupRouter(mode string) *gin.Engine {

	//如果配置文件中的mode设置为release模式，则gin框架也设置为发布模式(终端不输出任何信息)
	if mode == gin.ReleaseMode {
		gin.SetMode(gin.ReleaseMode) // gin设置成发布模式
	}

	//创建web服务时，不使用默认的gin.Default()
	r := gin.New()
	//使用自己写的两个中间件，
	r.Use(logger.GinLogger(), logger.GinRecovery(true))

	r.GET("/login", func(c *gin.Context) {
		c.String(http.StatusOK, "测试日志与项目配置")
	})
	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "404",
		})
	})

	return r

}
