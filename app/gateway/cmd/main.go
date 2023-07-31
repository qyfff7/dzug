package main

import (
	"dzug/app/gateway/routes"
	"dzug/discovery"
)

func main() {
	route := routes.NewRouter()          // 启动路由
	discovery.InitDiscovery()            // 启动服务发现程序
	defer discovery.SerDiscovery.Close() // 延时关闭
	_ = route.Run(":8001")
}
