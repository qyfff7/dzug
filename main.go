package main

import (
	favorservice "dzug/app/favorite/cmd"
	client "dzug/app/gateway/cmd"
	"dzug/app/redis"
	userservice "dzug/app/user/cmd"
	"dzug/app/user/pkg/snowflake"
	videoservice "dzug/app/video/cmd"
	"dzug/conf"
	transfer "dzug/conf/log_transfer"
	"dzug/repo"
	"fmt"
	"go.uber.org/zap"
	"time"
)

func main() {

	//1. 初始化配置文件
	if err := conf.Init(); err != nil {
		fmt.Printf("Config file initialization error,%#v", err)
		return
	}
	//2.初始化kafka消费者和ES
	go transfer.Init()

	//3. 初始化mysql数据库
	if err := repo.Init(); err != nil {
		//fmt.Printf("mysql  init error,%#v", err)
		zap.L().Error("初始化mysql数据库失败！！！")
		return
	}

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

	go func() {
		err := conf.Collectlog()
		if err != nil {
			zap.L().Error("log collect error ,", zap.Error(err))
		}
	}()

	time.Sleep(time.Second)
	go userservice.Start()
	time.Sleep(time.Second)
	go videoservice.Start()
	time.Sleep(time.Second)
	go favorservice.Start()
	client.Start()

}
