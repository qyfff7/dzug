package kafka

import (
	"fmt"
	"github.com/IBM/sarama"
	"go.uber.org/zap"
	"strconv"
)

func Sender(userId, videoId int64, ActionType int) error {
	msg := &sarama.ProducerMessage{} // 发送的消息
	msg.Topic = "favor"              // 消息的主题
	//key := fmt.Sprintf("%d", userId)
	key := "onlyKey"                                              // 暂时锁定这一个分区
	value := fmt.Sprintf("%d:%d:%d", userId, videoId, ActionType) // 写入我要写的信息，这个格式我感觉string比较方便，有直接的方法就用的string作为消息的类型了
	msg.Key = sarama.StringEncoder(key)                           // 我希望限定只用一个分区，保证消息被顺序消费，我不确定多个分区会对我的顺序消费是否有影响，感觉有，所以进行了限定
	msg.Value = sarama.StringEncoder(value)                       // 把value放入msg里
	pid, offset, err := KafkaProducer.SendMessage(msg)            // 发送消息进消息队列
	if err != nil {
		zap.L().Debug("send msg failed, err:" + err.Error())
		return err
	}
	zap.L().Debug("pid:" + string(pid) + " offset:" + strconv.FormatInt(offset, 10))
	return nil
}
