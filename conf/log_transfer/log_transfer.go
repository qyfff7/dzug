package log_transfer

import (
	"dzug/conf"
	"dzug/conf/es"
	"dzug/conf/kafka"
	"go.uber.org/zap"
)

func Init() {
	//  连接ES
	err := es.Init(conf.Config.ESConf.Address, conf.Config.LogConfig.Topic, conf.Config.ESConf.GoNum, conf.Config.ESConf.MaxSize)
	if err != nil {
		//fmt.Printf("Init es failed,err:%v\n", err)
		zap.L().Error("Init es failed,err: ", zap.Error(err))
		panic(err)
	}
	zap.L().Info("Init ES success")
	// 初始化kafka 消费者
	err = kafka.ConsumerInit([]string{conf.Config.KafkaConfig.Addr}, conf.Config.LogConfig.Topic)
	if err != nil {
		zap.L().Error("connect to kafka consumer failed,err: ", zap.Error(err))
		panic(err)
	}
	zap.L().Info("Init kafka success")
	zap.L().Info("log transfer start ...")
	// 在这儿停顿!
	select {}
}
