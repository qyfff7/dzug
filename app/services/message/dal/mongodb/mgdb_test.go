package mongodb

// import (
// 	"context"
// 	"encoding/json"
// 	"fmt"
// 	"testing"
// )

// func TestMongoDB(t *testing.T) {
// MdbInit()
//ts := time.Now().Unix() - 100000
//InsertMessage(&MgMessage{
//	ThreadId:    "1_2",
//	FromUserId:  1,
//	ToUserId:    2,
//	Contents:    "test1",
//	MessageUUID: 1,
//	CreateTime:  ts,
//})
//messages, err := GetMessages("1_2", ts)
//fmt.Printf("message len: %d", len(messages))
//assert.Equal(t, err, nil)
//for _, msg := range messages {
//	res, _ := json.Marshal(msg)
//	fmt.Printf(string(res))
//}
// 	message, _ := GetOldestMessage(context.TODO(), "1_2")
// 	res, _ := json.Marshal(message)
// 	fmt.Printf(string(res))
// }
