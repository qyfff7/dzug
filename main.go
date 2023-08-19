package main

import (
	commentservice "dzug/app/comment/cmd"
	client "dzug/app/gateway/cmd"
	"dzug/app/redis"
	"dzug/conf"
	"dzug/logger"
	"dzug/repo"

	"fmt"

	"go.uber.org/zap"
)

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
	//3. 初始化mysql数据库
	if err := repo.Init(); err != nil {
		fmt.Printf("mysql  init error,%#v", err)
		zap.L().Error("初始化mysql数据库失败！！！")
		return
	}

	//defer repo.Close()

	//4.初始化redis连接
	if err := redis.Init(conf.Config.RedisConfig); err != nil {
		fmt.Printf("init redis failed, err:%v\n", err)
		return
	}
	// 程序退出关闭数据库连接
	defer redis.Close()
	//defer repo.Close()
	go commentservice.Start()

	client.Start()

}
