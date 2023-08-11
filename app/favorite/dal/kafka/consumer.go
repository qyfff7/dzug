package kafka

import (
	"dzug/repo"
	"errors"
	"fmt"
	"github.com/IBM/sarama"
	"log"
	"strconv"
	"strings"
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
			update(string(message.Value))
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

func update(message string) error {
	arr := strings.Split(message, ":")
	favor := repo.Favorite{}
	favor.UserId, _ = strconv.ParseInt(arr[0], 10, 64)
	favor.VideoId, _ = strconv.ParseInt(arr[1], 10, 64)
	action := arr[2]
	if action == "1" {
		// 进行点赞
		fmt.Println("我们进行了一次点赞操作")
		return nil
	} else if action == "2" {
		// 进行取消点赞操作
		fmt.Println("我们进行了一次取消点赞操作")
		return nil
	}
	return errors.New("非法操作")
}
