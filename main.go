package main

import (
	client "dzug/app/gateway/cmd"
	"dzug/app/redis"
	userservice "dzug/app/user/cmd"
	"dzug/app/user/pkg/snowflake"
	"dzug/conf"
	"dzug/repo"
	"fmt"
	"time"
)

func main() {

	//1. 初始化配置文件
	if err := conf.Init([]string{"127.0.0.1:2379"}, "douyin"); err != nil {
		fmt.Printf("Config file initialization error,%#v", err)
		return
	}

	//2. 初始化日志
	//servicename := "project"
	//if err := logger.Init(conf.LogConfList, servicename); err != nil {
	//	fmt.Printf("log file initialization error,%#v", err)
	//	return
	//}
	//defer zap.L().Sync() //把缓冲区的日志，追加到文件中
	//zap.L().Info("服务启动，开始记录日志")

	//3. 初始化mysql数据库
	if err := repo.Init(); err != nil {
		fmt.Printf("mysql  init error,%#v", err)
		//zap.L().Error("初始化mysql数据库失败！！！")
		return
	}

	//defer repo.Close()

	//4.初始化redis连接
	if err := redis.Init(); err != nil {
		fmt.Printf("init redis failed, err:%v\n", err)
		return
	}
	// 程序退出关闭数据库连接
	defer redis.Close()

	//5. snowflake初始化
	if err := snowflake.Init(conf.Config.StartTime, conf.Config.MachineID); err != nil {
		//zap.L().Error("snowflake initialization error", zap.Error(err))
		return
	}
	//6.启动服务（后续可将所有的服务单独写到一个文件）

	go userservice.Start()
	time.Sleep(time.Second)
	//go videoservice.Start()
	//go favorservice.Start()
	client.Start()

}
