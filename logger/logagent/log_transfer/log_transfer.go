package log_transfer

import (
	"dzug/conf"
	"dzug/logger/logagent/es"
	"dzug/logger/logagent/kafka"
	"go.uber.org/zap"
)

func Init() {
	//  连接ES
	err := es.Init(conf.LogConf.Address, conf.LogConf.Topic, conf.LogConf.GoNum, conf.LogConf.MaxSize)
	if err != nil {
		zap.L().Error("Init es failed,err: ", zap.Error(err))
		panic(err)
	}
	zap.L().Info("Init ES success")
	// 初始化kafka 消费者
	err = kafka.ConsumerInit([]string{conf.LogConf.Addr}, conf.LogConf.Topic)
	if err != nil {
		zap.L().Error("connect to kafka consumer failed,err: ", zap.Error(err))
		panic(err)
	}
	zap.L().Info("Init kafka success")
	zap.L().Info("log transfer start ...")
	// 在这儿停顿!
	select {}
}
