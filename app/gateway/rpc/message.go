package rpc

import (
	"context"
	"dzug/discovery"
	"dzug/protos/message"
)

func MessageChatList(ctx context.Context, req *message.GetMessageListReq) (resp *message.GetMessageListResp, err error) {
	discovery.LoadClient("message", &discovery.MessageClient)
	r, err := discovery.MessageClient.GetMessageList(ctx, req) // 调用注册方法
	if err != nil {
		return
	}
	return r, nil
}

func MessagePostAction(ctx context.Context, req *message.CreateMessageReq) (resp *message.CreateMessageResp, err error) {
	discovery.LoadClient("message", &discovery.MessageClient)
	r, err := discovery.MessageClient.CreateMessage(ctx, req) // 调用注册方法
	if err != nil {
		return
	}
	return r, nil
}
