package main

import (
	client "dzug/app/gateway/cmd"
	"dzug/app/services/config_center"
	userservice "dzug/app/services/user/cmd"
	"dzug/logger"
	"fmt"
	"go.uber.org/zap"
)

func main() {

	//1.启动配置中心（将项目初始配置都存到etcd中，启动监控）
	config_center.Init()

	//2.初始化日志（获取etcd中日志的配置，进行初始化，包括kafka,es的初始化）这也是一个服务
	if err := logger.Init(); err != nil {
		fmt.Printf("log file initialization error,%#v", err)
		return
	}
	defer zap.L().Sync() //把缓冲区的日志，追加到文件中
	zap.L().Info("服务启动，开始记录日志")

	//3.各个服务启动（①获取各自的配置，进行相应的初始化，②进行业务代码操作）

	go userservice.Start()
	client.Start()

	/*//1. 初始化配置中心
	if err := conf.Init(); err != nil {
		fmt.Printf("Config file initialization error,%#v", err)
		return
	}
	//2.初始化kafka消费者和ES
	go transfer.Init()

	//3. 初始化mysql数据库
	if err := repo.Init(); err != nil {
		fmt.Printf("mysql  init error,%#v", err)
		zap.L().Error("初始化mysql数据库失败！！！")
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
		zap.L().Error("snowflake initialization error", zap.Error(err))
		return
	}
	//6.启动日志收集
	go func() {
		err := conf.Collectlog()
		if err != nil {
			zap.L().Error("log collect error ,", zap.Error(err))
		}
	}()

	//6.启动服务（后续可将所有的服务单独写到一个文件）
	go userservice.Start()
	time.Sleep(time.Second)
	go videoservice.Start()
	go favorservice.Start()
	go messageservice.Start()
	go commentservice.Start()
	go relationservice.Start()
	go publishservice.Start()
	client.Start()*/

	//1.初始化配置中心（将项目初始配置都存到etcd中，启动监控）
	//2.初始化日志（获取etcd中日志的配置，进行初始化，包括kafka,es的初始化）这也是一个服务
	//3.各个服务启动（①获取各自的配置，进行相应的初始化，②进行业务代码操作）

	/*time.Sleep(time.Second)
	fmt.Println(config_center.ProjConf.Name)
	fmt.Println("2222" + config_center.ProjConf.Name)
	time.Sleep(time.Second * 2)
	go func() {
		for {
			fmt.Println("2222" + config_center.ProjConf.Name + "hahah")
			fmt.Println("2222" + config_center.ProjConf.Name)
		}

	}()

	select {}*/
}
