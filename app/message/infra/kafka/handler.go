package kafka

import (
	"context"
	"dzug/app/message/infra/db"
	"dzug/app/message/infra/mongodb"
	"encoding/json"
	"fmt"
	"github.com/IBM/sarama"
	"time"
)

func MessageHandler(kfMessage *sarama.ConsumerMessage) error {

	var msg db.Message
	if err := json.Unmarshal(kfMessage.Value, &msg); err != nil {
		fmt.Printf("Unmarshal message fail " + err.Error())
		return err
	}

	parent := context.Background()
	ctx, cancel := context.WithTimeout(parent, 5*time.Second)
	defer cancel()
	if err := db.CreateMessage(ctx, &msg); err != nil {
		fmt.Printf("Insert message in db fail " + err.Error())
		return err
	}

	if err := mongodb.InsertMessage(ctx, &mongodb.MgMessage{
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
