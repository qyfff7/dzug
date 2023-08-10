package kafka

import (
	"fmt"
	"github.com/IBM/sarama"
)

func Sender(userId, videoId int64, ActionType int) {
	msg := &sarama.ProducerMessage{}
	msg.Topic = "favor"
	//key := fmt.Sprintf("%d", userId)
	key := "onlyKey" // 暂时锁定这一个分区
	value := fmt.Sprintf("%d:%d:%d", userId, videoId, ActionType)
	msg.Key = sarama.StringEncoder(key)
	msg.Value = sarama.StringEncoder(value)
	pid, offset, err := KafkaProducer.SendMessage(msg)
	if err != nil {
		fmt.Println("send msg failed, err:", err)
		return
	}
	fmt.Printf("pid:%v offset:%v\n", pid, offset)
}
