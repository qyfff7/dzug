package kafka

import (
	"dzug/conf"
	"github.com/IBM/sarama"
	"go.uber.org/zap"
)

var (
	KafkaConsumer sarama.Consumer     // 消费者
	KafkaProducer sarama.SyncProducer // 生产者
)

func InitKafka() error {
	var err error
	// 成功交付的消息将在success channel返回
	KafkaConsumer, err = sarama.NewConsumer(conf.Config.KafkaConfig.Addr, nil) // 启动消费者，此时消费者开始消费
	//KafkaConsumer, err = sarama.NewConsumer(conf.Config.KafkaConfig.Addr, nil) // 启动消费者，此时消费者开始消费
	if err != nil {
		panic("kafka错误：" + err.Error())
	}

	config := sarama.NewConfig()                            // 配置生成这
	config.Producer.RequiredAcks = sarama.WaitForAll        // 发送完数据需要leader和follow都确认
	config.Producer.Partitioner = sarama.NewHashPartitioner // 生产者分区用的hash分配，我暂时只用一个分区
	config.Producer.Return.Successes = true                 //回复确认

	KafkaProducer, err = sarama.NewSyncProducer([]string{"127.0.0.1:9092"}, config) // 初始化生产者
	if err != nil {
		panic(err)
		return err
	}
	return nil
}

func InitConsumer() error {
	var err error
	// 成功交付的消息将在success channel返回
	KafkaConsumer, err = sarama.NewConsumer([]string{"127.0.0.1:9092"}, nil) // 启动消费者，此时消费者开始消费
	//KafkaConsumer, err = sarama.NewConsumer(conf.Config.KafkaConfig.Addr, nil) // 启动消费者，此时消费者开始消费
	if err != nil {
		panic("kafka错误：" + err.Error())
	}
	return nil
}

func InitProducer() error {
	var err error
	config := sarama.NewConfig()                            // 配置生成这
	config.Producer.RequiredAcks = sarama.WaitForAll        // 发送完数据需要leader和follow都确认
	config.Producer.Partitioner = sarama.NewHashPartitioner // 生产者分区用的hash分配，我暂时只用一个分区
	config.Producer.Return.Successes = true                 //回复确认

	KafkaProducer, err = sarama.NewSyncProducer([]string{"127.0.0.1:9092"}, config) // 初始化生产者
	if err != nil {
		panic(err)
		return err
	}
	return nil
}

func CloseKafka() {
	if err := KafkaProducer.Close(); err != nil {
		zap.L().Fatal("关闭生产者失败" + err.Error())
	}
	if err := KafkaConsumer.Close(); err != nil {
		zap.L().Fatal("关闭消费者失败" + err.Error())
	}
}
