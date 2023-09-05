package handlers

import (
	"dzug/app/gateway/rpc"
	"dzug/app/services/user/pkg/jwt"
	"dzug/models"
	pb "dzug/protos/message"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"strconv"
)

func MessageChatList(ctx *gin.Context) {
	var msgListReq pb.GetMessageListReq

	zap.L().Info("Getting Message List!")

	//if err := ctx.Bind(&msgListReq); err != nil {
	//	zap.L().Error("Get Message List with invalid param", zap.Error(err))
	//	models.ResponseError(ctx, models.CodeInvalidParam)
	//	return
	//}
	msgListReq.Token = ctx.Query("token")
	toUserId, _ := strconv.ParseInt(ctx.Query("to_user_id"), 10, 64)
	msgListReq.ToUserId = toUserId
	preMsgTime, _ := strconv.ParseInt(ctx.Query("pre_msg_time"), 10, 64)
	msgListReq.PreMsgTime = preMsgTime
	//fmt.Printf("Got message list request: %++v \n", msgListReq)

	u, err := jwt.ParseToken(msgListReq.Token)
	if err != nil {
		zap.L().Error("解析Token出错", zap.Error(err))
		return
	}
	fmt.Println(u.UserID)
	msgListReq.Token = strconv.FormatInt(u.UserID, 10)

	msgListResp, err := rpc.MessageChatList(ctx, &msgListReq)
	if err != nil {
		zap.L().Error("Get Message List rpc error", zap.Error(err))
		models.ResponseErrorWithMsg(ctx, models.CodeServerBusy, err.Error())
		return
	}
	msgList := make([]*models.Message, 0)
	for _, m := range msgListResp.MessageInfos {
		msgList = append(msgList, &models.Message{
			FromUserId:  m.FromUserId,
			ToUserId:    m.ToUserId,
			Contents:    m.Content,
			CreateTime:  m.CreateTime,
			MessageUUID: m.MessageId,
		})
	}
	models.MessageListRespSuccess(ctx, msgList)
}

func MessagePostAction(ctx *gin.Context) {
	msgPostReq := new(pb.CreateMessageReq)

	zap.L().Info("Posting a Message!")

	msgPostReq.Token = ctx.Query("token")
	toUserId, _ := strconv.ParseInt(ctx.Query("to_user_id"), 10, 64)
	msgPostReq.ToUserId = toUserId
	n, _ := strconv.Atoi(ctx.Query("action_type"))
	msgPostReq.ActionType = int32(n)
	msgPostReq.Content = ctx.Query("content")

	fmt.Printf("Got Chat request: %++v\n", msgPostReq)

	u, err := jwt.ParseToken(msgPostReq.Token)
	if err != nil {
		zap.L().Error("解析Token出错", zap.Error(err))
		return
	}
	fmt.Println(u.UserID)
	msgPostReq.Token = strconv.FormatInt(u.UserID, 10)

	tmp, _ := json.Marshal(msgPostReq)
	fmt.Println(string(tmp))

	_, err = rpc.MessagePostAction(ctx, msgPostReq)
	if err != nil {
		zap.L().Error("Post message rpc error", zap.Error(err))
		models.ResponseErrorWithMsg(ctx, models.CodeServerBusy, err.Error())
		return
	}
	models.PostMessageRespSuccess(ctx)
}
