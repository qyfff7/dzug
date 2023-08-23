package kafka

import (
	"context"
	dao2 "dzug/app/services/message/dal/dao"
	mongodb2 "dzug/app/services/message/dal/mongodb"
	"encoding/json"
	"fmt"
	"github.com/IBM/sarama"
	"time"
)

func MessageHandler(kfMessage *sarama.ConsumerMessage) error {

	var msg dao2.Message
	if err := json.Unmarshal(kfMessage.Value, &msg); err != nil {
		fmt.Printf("Unmarshal message fail " + err.Error())
		return err
	}

	parent := context.Background()
	ctx, cancel := context.WithTimeout(parent, 5*time.Second)
	defer cancel()
	if err := dao2.CreateMessage(ctx, &msg); err != nil {
		fmt.Printf("Insert message in db fail " + err.Error())
		return err
	}

	if err := mongodb2.InsertMessage(ctx, &mongodb2.MgMessage{
		ThreadId:    msg.ThreadId,
		FromUserId:  msg.FromUserId,
		ToUserId:    msg.ToUserId,
		Contents:    msg.Contents,
		MessageUUID: msg.MessageUUID,
		CreateTime:  msg.CreateTime,
	}); err != nil {
		fmt.Printf("Insert message in mongo db fail " + err.Error())
		return err
	}

	return nil
}
