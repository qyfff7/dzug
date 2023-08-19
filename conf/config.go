package conf

import (
	"dzug/conf/etcd"
	"dzug/conf/kafka"
	"dzug/conf/tailfile"
	"dzug/logger"
	"dzug/models"
	"fmt"
	"go.uber.org/zap"
)

// Config 全局变量，用来保存项目所有的配置信息
var Config = new(models.ProjectConfig)
var LogConfList []*models.LogConfig

// Init
func Init(etcdaddr []string, projName string) (err error) {
	//1. 初始化etcd连接
	err = etcd.Init(etcdaddr)
	if err != nil {
		//zap.L().Error("init etcd failed, err:", zap.Error(err))
		fmt.Println("init etcd failed, err:" + err.Error())
		return
	}
	// 2.从etcd中拉取项目所有的配置项
	Config, err = etcd.GetProjectConf(projName)
	if err != nil {
		//zap.L().Error("get conf from etcd failed, err:", zap.Error(err))
		fmt.Printf("get conf from etcd failed, err:%s", err)
		return
	}
	fmt.Printf("%s", Config)
	//初始化日志
	if err := logger.Init(Config.LogConfig); err != nil {
		fmt.Printf("log file initialization error,%#v", err)
		return
	}
	defer zap.L().Sync() //把缓冲区的日志，追加到文件中
	zap.L().Info("服务启动，开始记录日志")

	//3. 初始化连接kafka(做好准备工作)     (初始化kafka,初始化msg chan，起后台gorountine 去往kafka发msg)
	err = kafka.Init([]string{Config.KafkaConfig.Addr}, Config.KafkaConfig.ChanSize)
	if err != nil {
		zap.L().Error("init kafka failed, err:%v", zap.Error(err))
		return
	}
	zap.L().Info("init kafka success!")

	// 4.派一个小弟去监控etcd中 日志配置的变化
	//go etcd.WatchConf(Config.LogConfig)

	// 5. 根据配置中的日志路径初始化tail   （根据配置文件中指定的路径创建了一个对应的tailObj）
	err = tailfile.Init(LogConfList)
	if err != nil {
		//zap.L().Error("init tailfile failed, err:%v", zap.Error(err))
		fmt.Printf("init tailfile failed, err:%v", zap.Error(err))
		return
	}
	//zap.L().Info("init tailfile success!")
	fmt.Printf("init tailfile success!")
	//6. run
	confrun()
	return
}
func confrun() {
	select {} //这里死循环，让程序不停的运行
}
