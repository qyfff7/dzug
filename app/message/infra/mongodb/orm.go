package mongodb

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

func InsertMessage(ctx context.Context, message *MgMessage) error {
	if _, err := collection.InsertOne(ctx, message); err != nil {
		fmt.Print(err)
		return err
	}
	return nil
}

func GetOldestMessage(ctx context.Context, threadId string) (*MgMessage, error) {
	cond := bson.D{
		{"thread_id", threadId},
	}
	opts := options.FindOne().SetSort(bson.D{{"create_time", 1}})
	var result MgMessage
	if err := collection.FindOne(ctx, cond, opts).Decode(&result); err != nil {
		fmt.Print(err)
		return nil, err
	}
	return &result, nil
}

func GetMessages(ctx context.Context, threadId string, preMsgTime int64) ([]*MgMessage, error) {
	cond := bson.M{
		"thread_id": threadId,
		"create_time": bson.M{
			"$gte": preMsgTime,
		},
	}
	var cursor *mongo.Cursor
	var err error
	if cursor, err = collection.Find(ctx, cond); err != nil {
		fmt.Println(err)
		return nil, err
	}
	//延迟关闭游标
	defer func() {
		if err = cursor.Close(ctx); err != nil {
			log.Fatal(err)
		}
	}()

	results := make([]*MgMessage, 0)
	//遍历游标获取结果数据
	for cursor.Next(ctx) {
		var lr MgMessage
		//反序列化Bson到对象
		if cursor.Decode(&lr) != nil {
			fmt.Print(err)
			return nil, err
		}
		results = append(results, &lr)
	}

	return results, nil
}
