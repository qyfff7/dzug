package kafka

import (
	"github.com/IBM/sarama"
	"go.uber.org/zap"
	"strconv"
)

func SendMsg(topic string, key string, message string) error {
	msg := &sarama.ProducerMessage{
		Topic: topic,
		Key:   sarama.StringEncoder(key),
		Value: sarama.StringEncoder(message),
	}
	pid, offset, err := KafkaProducer.SendMessage(msg) // 发送消息进消息队列
	if err != nil {
		zap.L().Debug("send msg failed, err:" + err.Error())
		return err
	}
	zap.L().Debug("pid:" + string(pid) + " offset:" + strconv.FormatInt(offset, 10))
	return nil
}
