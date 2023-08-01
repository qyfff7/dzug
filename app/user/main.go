package main

import (
	"dzug/app/user/models"
	"dzug/app/user/pkg/snowflake"
	"dzug/user_service/Log_Conf/conf"
	"dzug/user_service/Log_Conf/logger"
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

	//3. 初始化mysql数据库

	models.InitDB()

	//...

	models.InsertData()

	//4. snowflake初始化
	if err := snowflake.Init(conf.Config.StartTime, conf.Config.MachineID); err != nil {
		zap.L().Error("snowflake initialization error", zap.Error(err))
		return
	}

	////4.注册路由
	//r := routes.SetupRouter(conf.Config.Mode)
	//
	////5.启动项目
	//if err := r.Run(fmt.Sprintf(":%d", conf.Config.Port));err != nil {
	//	fmt.Printf("run server failed, err:%v\n", err)
	//	return
	//}
	//zap.L().Info("Server exiting")

}
