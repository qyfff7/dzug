package kafka

import (
	"dzug/app/favorite/dal/dao"
	"dzug/repo"
	"errors"
	"github.com/IBM/sarama"
	"go.uber.org/zap"
	"strconv"
	"strings"
)

var messages chan *sarama.ConsumerMessage

func FavorConsumer() {
	zap.L().Debug("消息队列消费者开始消费")
	// 创建一个消息通道
	messages = make(chan *sarama.ConsumerMessage)

	// 启动一个goroutine来消费消息
	go func() {
		for message := range messages {
			err := update(string(message.Value)) // 我的数据库更新的地方
			if err != nil {
				zap.L().Fatal("消息队列更新点赞失败")
			}
		}
	}()

	// 开始消费消息，消费favor主题，partition 0
	partitionConsumer, err := KafkaConsumer.ConsumePartition("favor", 0, sarama.OffsetNewest)
	if err != nil {
		zap.L().Fatal("消费消息队列失败" + err.Error())
	}

	// 将消息发送到消息通道
	for {
		select {
		case err := <-partitionConsumer.Errors():
			zap.L().Fatal(err.Error())
		case msg := <-partitionConsumer.Messages():
			messages <- msg
		}
	}
}

// CloseConsumer 关闭消费者，在启动消息的地方进行defer close，close时关闭消息通道
func CloseConsumer() {
	close(messages)
	if err := KafkaConsumer.Close(); err != nil {
		zap.L().Fatal("关闭消费者失败" + err.Error())
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
		err := dao.Favor(favor.VideoId, favor.UserId)
		if err != nil {
			zap.L().Error("用户 " + strconv.FormatInt(favor.UserId, 10) + " 点赞 " + strconv.FormatInt(favor.VideoId, 10) + " 失败")
			return err
		}
		return nil
	} else if action == "2" {
		// 进行取消点赞操作
		err := dao.InFavor(favor.VideoId, favor.UserId)
		if err != nil {
			zap.L().Error("用户 " + strconv.FormatInt(favor.UserId, 10) + " 取消点赞 " + strconv.FormatInt(favor.VideoId, 10) + " 失败")
			return err
		}
		return nil
	}
	return errors.New("非法操作")
}
