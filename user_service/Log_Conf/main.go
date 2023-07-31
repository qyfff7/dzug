package main

import (
	"dzug/user_service/Log_Conf/conf"
	"dzug/user_service/Log_Conf/logger"
	"dzug/user_service/Log_Conf/routes"
	"fmt"
	"go.uber.org/zap"
)

// main 项目主程序
func main() {

	//1. 初始化配置文件
	if err := conf.Init(); err != nil {
		fmt.Printf("Config file initialization error,%#v", err)
		return
	}

	//2. 初始化日志
	if err := logger.Init(conf.Config.LogConfig, conf.Config.Mode); err != nil {
		fmt.Printf("log file initialization error,%#v", err)
		return
	}
	defer zap.L().Sync() //把缓冲区的日志，追加到文件中
	zap.L().Info("服务启动，开始记录日志")

	//3. 初始化mysql、Redis、etcd等相关组件
	//...

	//4.注册路由
	r := routes.SetupRouter(conf.Config.Mode)

	//5.启动项目
	err := r.Run(fmt.Sprintf(":%d", conf.Config.Port))

	if err != nil {
		fmt.Printf("run server failed, err:%v\n", err)
		return
	}
	zap.L().Info("Server exiting")

}
