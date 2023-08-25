package main

import (
	client "dzug/app/gateway/cmd"
	userservice "dzug/app/services/user/cmd"
	"dzug/conf"
	"dzug/logger"
	"fmt"
	"go.uber.org/zap"
)

func main() {

	//1.启动配置中心（将项目初始配置都存到etcd中，启动监控）
	if err := conf.Init(); err != nil {
		fmt.Printf("config center initialization error,%#v", err)
		return
	}

	//2.启动日志（获取etcd中日志的配置，进行初始化，包括kafka,es的初始化）这也是一个服务
	if err := logger.Init(); err != nil {
		fmt.Printf("log file initialization error,%#v", err)
		return
	}
	defer zap.L().Sync() //把缓冲区的日志，追加到文件中
	zap.L().Info("服务启动，开始记录日志")

	//3.各个服务启动（①获取各自的配置，进行相应的初始化，②进行业务代码操作）
	go userservice.Start()

	client.Start()

}
