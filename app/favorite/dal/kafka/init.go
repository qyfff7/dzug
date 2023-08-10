package kafka

import (
	"github.com/IBM/sarama"
)

var (
	KafkaConsumer sarama.Consumer
	KafkaProducer sarama.SyncProducer
	KafkaAddr     = []string{"127.0.0.1:9092"}
)

func InitKafka() {
	var err error // 成功交付的消息将在success channel返回
	KafkaConsumer, err = sarama.NewConsumer(KafkaAddr, nil)
	if err != nil {
		panic("kafka错误：" + err.Error())
	}

	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll // 发送完数据需要leader和follow都确认
	config.Producer.Partitioner = sarama.NewHashPartitioner

	config.Producer.Return.Successes = true
	KafkaProducer, err = sarama.NewSyncProducer(KafkaAddr, config)
	if err != nil {
		panic(err)
		return
	}

}
