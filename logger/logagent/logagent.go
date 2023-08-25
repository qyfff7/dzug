package logagent

import (
	"dzug/conf"
	"dzug/logger/logagent/kafka"
	"dzug/logger/logagent/tailfile"
	"go.uber.org/zap"
)

// LogAgentInit 日志收集
func LogAgentInit() (err error) {

	//1. 初始化连接kafka生产者(做好准备工作)     (初始化kafka,初始化msg chan，起后台gorountine 去往kafka发msg)
	err = kafka.Init([]string{conf.LogConf.KafkaConfig.Addr}, conf.LogConf.KafkaConfig.ChanSize)
	if err != nil {
		zap.L().Error("init kafka failed, err:%v", zap.Error(err))
		return
	}
	zap.L().Info("init kafka success!")
	// 2. 根据配置中的日志路径初始化tail   （根据配置文件中指定的路径创建了一个对应的tailObj）
	err = tailfile.Init(conf.LogConf.Path)
	if err != nil {
		zap.L().Error("init tailfile failed, err:%v", zap.Error(err))
		return
	}
	zap.L().Info("init tailfile success!")
	return nil

}
