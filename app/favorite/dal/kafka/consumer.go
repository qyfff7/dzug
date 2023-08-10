package kafka

import (
	"fmt"
	"github.com/IBM/sarama"
	"log"
)

var messages chan *sarama.ConsumerMessage

func FavorConsumer() {
	fmt.Println("我正在执行")
	// 创建一个消息通道
	messages = make(chan *sarama.ConsumerMessage)
	//defer func() {
	//	close(messages)
	//	if err := KafkaConsumer.Close(); err != nil {
	//		log.Fatal(err)
	//	}
	//}() // 关闭消费者

	// 启动一个goroutine来消费消息
	go func() {
		for message := range messages {
			fmt.Println(string(message.Value))
		}
	}()

	// 开始消费消息
	partitionConsumer, err := KafkaConsumer.ConsumePartition("favor", 0, sarama.OffsetNewest)
	if err != nil {
		log.Fatal(err)
	}

	// 将消息发送到消息通道
	for {
		select {
		case err := <-partitionConsumer.Errors():
			log.Fatal(err)
		case msg := <-partitionConsumer.Messages():
			messages <- msg
		}
	}
}

func CloseConsumer() {
	close(messages)
	if err := KafkaConsumer.Close(); err != nil {
		log.Fatal(err)
	}
}
