package kafka

import (
	"fmt"
	"github.com/IBM/sarama"
	"go.uber.org/zap"
	"strconv"
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
		zap.L().Debug("send msg failed, err:" + err.Error())
		return
	}
	zap.L().Debug("pid:" + string(pid) + " offset:" + strconv.FormatInt(offset, 10))
}
