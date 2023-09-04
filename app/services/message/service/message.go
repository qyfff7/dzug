package service

import (
	"context"
	dao2 "dzug/app/services/message/dal/dao"
	"dzug/app/services/message/dal/kafka"
	mongodb2 "dzug/app/services/message/dal/mongodb"
	"dzug/app/services/user/pkg/snowflake"
	pb "dzug/protos/message"
	"encoding/json"
	"fmt"
	"go.uber.org/zap"
	"strconv"
	"time"
)

type MsgSvr struct {
	pb.UnimplementedDouyinMessageServiceServer
}

func getThreadId(fromUserId int64, toUserId int64) string {
	if fromUserId < toUserId {
		return strconv.FormatInt(fromUserId, 10) + "_" + strconv.FormatInt(toUserId, 10)
	} else {
		return strconv.FormatInt(toUserId, 10) + "_" + strconv.FormatInt(fromUserId, 10)
	}
}

func (MsgSvr) CreateMessage(ctx context.Context, request *pb.CreateMessageReq) (*pb.CreateMessageResp, error) {

	uuid := snowflake.GenID()
	userId, _ := strconv.ParseInt(request.Token, 10, 64)

	message := &dao2.Message{
		ThreadId:    getThreadId(userId, request.ToUserId),
		FromUserId:  userId,
		ToUserId:    request.ToUserId,
		Contents:    request.Content,
		MessageUUID: uuid,
		CreateTime:  time.Now().Unix(),
	}
	res, err := json.Marshal(message)
	if err != nil {
		fmt.Println(err)
	}

	// 把Message发送到Kafka
	if err := kafka.SendMsg("message", "", string(res)); err != nil {
		return &pb.CreateMessageResp{
			BaseResp: &pb.BaseResp{
				StatusCode: 500,
				StatusMsg:  "Kafka服务错误",
			},
		}, nil
	}

	return &pb.CreateMessageResp{
		BaseResp: &pb.BaseResp{
			StatusCode: 200,
			StatusMsg:  "调用成功，你成功发送了一条消息",
		},
	}, nil

}

func (MsgSvr) GetMessageList(ctx context.Context, request *pb.GetMessageListReq) (*pb.GetMessageListResp, error) {

	//userId := tokenToUserId(request.Token)
	userId, _ := strconv.ParseInt(request.Token, 10, 64)

	threadId := getThreadId(request.ToUserId, userId)
	fmt.Printf("Message between: %d, %d. Thread: %s", request.ToUserId, userId, threadId)

	oldestCache, err := mongodb2.GetOldestMessage(ctx, threadId)
	if err != nil {
		zap.L().Error("获取缓存记录失败", zap.Error(err))
		return &pb.GetMessageListResp{
			BaseResp: &pb.BaseResp{
				StatusCode: 500,
				StatusMsg:  "Kafka服务错误",
			},
		}, nil
	}

	var infos []*pb.MessageInfo
	if request.PreMsgTime < oldestCache.CreateTime {
		//fmt.Printf("Pull messages from db ")
		msgs, err := dao2.GetMessageList(ctx, userId, request.ToUserId, request.PreMsgTime)
		if err != nil {
			fmt.Printf("Get messages from db fail: " + err.Error())
			return nil, err
		}
		infos = messagesToInfo(msgs)
	} else {
		//fmt.Printf("Pull messages from mongodb ")
		mgMessages, err := mongodb2.GetMessages(ctx, threadId, request.PreMsgTime)
		if err != nil {
			fmt.Printf("Get messages from cache fail: " + err.Error())
			return nil, err
		}
		infos = mgMessagesToInfo(mgMessages)
	}
	//for _, info := range infos {
	//	fmt.Printf("Got msg from DB: %++v \n", info)
	//}

	//err = dao2.UpdateThreadPullInfo(ctx, threadId)
	//if err != nil {
	//	return nil, err
	//}

	return &pb.GetMessageListResp{
		BaseResp: &pb.BaseResp{
			StatusCode: 200,
			StatusMsg:  "调用成功，你成功查询了消息记录",
		},
		MessageInfos: infos,
	}, nil
}

func messagesToInfo(messages []*dao2.Message) []*pb.MessageInfo {
	infos := make([]*pb.MessageInfo, 0)
	for _, msg := range messages {
		infos = append(infos, &pb.MessageInfo{
			MessageId:  msg.MessageUUID,
			FromUserId: msg.FromUserId,
			ToUserId:   msg.ToUserId,
			Content:    msg.Contents,
			CreateTime: msg.CreateTime,
		})
	}
	return infos
}

func mgMessagesToInfo(messages []*mongodb2.MgMessage) []*pb.MessageInfo {
	infos := make([]*pb.MessageInfo, 0)
	for _, msg := range messages {
		infos = append(infos, &pb.MessageInfo{
			MessageId:  msg.MessageUUID,
			FromUserId: msg.FromUserId,
			ToUserId:   msg.ToUserId,
			Content:    msg.Contents,
			CreateTime: msg.CreateTime,
		})
	}
	return infos
}
