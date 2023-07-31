package main

import (
	"dzug/user_service/Log_Conf/conf"
	"dzug/user_service/Log_Conf/logger"
	"dzug/user_service/app/gateway/routes"
	"dzug/user_service/discovery"
	"fmt"
	"go.uber.org/zap"
)

//用户服务启动文件

/*func main() {
	route := routes.NewRouter() // 启动路由
	discovery.Init()            // 启动服务发现程序
	defer discovery.Ser.Close() // 延时关闭
	_ = route.Run(":8001")
}
*/

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

	route := routes.NewRouter(conf.Config.Mode) // 启动路由

	discovery.Init()            // 启动服务发现程序
	defer discovery.Ser.Close() // 延时关闭

	//5.启动项目
	err := route.Run(fmt.Sprintf(":%d", conf.Config.Port))

	if err != nil {
		fmt.Printf("run server failed, err:%v\n", err)
		return
	}
	zap.L().Info("Server exiting")

}
