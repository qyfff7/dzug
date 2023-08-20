package db

// import (
// 	"context"
// 	"encoding/json"
// 	"fmt"
// 	"testing"
// 	"time"

// 	"github.com/magiconair/properties/assert"
// )

// func TestMysql(t *testing.T) {
// 	Init()
// 	ts := time.Now().Unix()
// 	err := CreateMessage(context.TODO(), &Message{
// 		ThreadId:    "1",
// 		ToUserId:    5174765849415680,
// 		FromUserId:  5191404976345088,
// 		Contents:    "test",
// 		MessageUUID: 1,
// 		CreateTime:  ts,
// 	})
// 	assert.Equal(t, err, nil)
// 	messages, err := GetMessageList(context.TODO(), 5174765849415680, 5191404976345088, ts-5)
// 	assert.Equal(t, err, nil)
// 	for _, msg := range messages {
// 		res, _ := json.Marshal(msg)
// 		fmt.Println(string(res))
// 	}
// }
