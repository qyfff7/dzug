package handlers

import (
	"dzug/app/gateway/rpc"
	"dzug/app/services/user/pkg/jwt"
	"dzug/models"
	pb "dzug/protos/message"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
	"strconv"
)

func MessageChatList(ctx *gin.Context) {
	var msgListReq pb.GetMessageListReq

	zap.L().Info("Getting Message List!")

	if err := ctx.Bind(&msgListReq); err != nil {
		zap.L().Error("Get Message List with invalid param", zap.Error(err))
		models.ResponseError(ctx, models.CodeInvalidParam)
		return
	}

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

	if err := ctx.ShouldBind(msgPostReq); err != nil {
		zap.L().Error("Post message with invalid param", zap.Error(err))
		// 判断err是不是validator.ValidationErrors 类型
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			// 非validator.ValidationErrors类型错误直接返回
			models.ResponseError(ctx, models.CodeInvalidParam)
			return
		}
		err, _ := json.Marshal(removeTopStruct(errs.Translate(trans)))
		models.ResponseErrorWithMsg(ctx, models.CodeInvalidParam, string(err))
		return
	}

	fmt.Println(msgPostReq.Token)
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
