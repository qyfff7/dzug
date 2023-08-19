package logagent

import (
	"dzug/conf"
	"dzug/logger/logagent/etcd"
	"dzug/logger/logagent/kafka"
	"fmt"
	"go.uber.org/zap"
)

func run() {
	select {} //这里死循环，让程序不停的运行
}
func Start() {
	//1. 初始化etcd连接
	err := etcd.Init([]string{"127.0.0.1:2379"})
	if err != nil {
		//zap.L().Error("init etcd failed, err:", zap.Error(err))
		fmt.Println("init etcd failed, err:" + err.Error())
		return
	}
	// 2.从etcd中拉取项目所有的配置项
	Config, err := etcd.GetConf("ProjectConfig")
	if err != nil {
		//zap.L().Error("get conf from etcd failed, err:", zap.Error(err))
		fmt.Printf("get conf from etcd failed, err:%s", err)
		return
	}
	fmt.Printf("%#v", Config)

	//3. 初始化连接kafka(做好准备工作)     (初始化kafka,初始化msg chan，起后台gorountine 去往kafka发msg)
	err = kafka.Init(Config.KafkaConfig.Addr, Config.KafkaConfig.ChanSize)
	if err != nil {
		zap.L().Error("init kafka failed, err:%v", zap.Error(err))
		return
	}
	zap.L().Info("init kafka success!")

	// 4.派一个小弟去监控etcd中 conf.Config.EtcdConfig.LogCollectKey 对应值的变化
	go etcd.WatchConf(conf.Config.EtcdConfig.LogCollectKey)

	/*	// 5. 根据配置中的日志路径初始化tail   （根据配置文件中指定的路径创建了一个对应的tailObj）
		err = tailfile.Init(Config)
		if err != nil {
			zap.L().Error("init tailfile failed, err:%v", zap.Error(err))
			return
		}
		zap.L().Info("init tailfile success!")
		//6. run
		run()*/
}
