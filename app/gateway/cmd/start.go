package client

import (
	"dzug/app/gateway/routes"
	"dzug/conf"
	"dzug/discovery"
	"fmt"
	"go.uber.org/zap"
)

//这里相当于是客户端，去访问各个服务

func Start() {

	if err := conf.Init(); err != nil {
		fmt.Printf("Config file initialization error,%#v", err)
		return
	}

	go func() {
		err := conf.Collectlog()
		if err != nil {
			zap.L().Error("log collect error ,", zap.Error(err))
		}
	}()

	//1.注册路由，并启动服务发现程序
	r := routes.NewRouter(conf.Config.Mode) // 启动路由
	discovery.InitDiscovery()               // 启动服务发现程序
	defer discovery.SerDiscovery.Close()    // 延时关闭服务发现程序

	//2.启动项目
	err := r.Run(fmt.Sprintf(":%d", conf.Config.Port))
	// err := r.Run()

	if err != nil {
		fmt.Printf("run server failed, err:%v\n", err)
		return
	}
	zap.L().Info("Server exiting")
}

func main() {
	Start()
}
