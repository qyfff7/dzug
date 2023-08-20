package handlers

import (
	"dzug/app/gateway/rpc"
	"dzug/app/user/pkg/jwt"
	"dzug/models"
	pb "dzug/protos/relation"
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type relationAction struct {
	Token      string `json:"token"`
	toUserId   int64  `json:"to_user_id"`
	ActionType int32  `json:"action_type"`
}

type RelationListReq struct {
	UserId int64  `json:"user_id"`
	Token  string `json:"token"`
}

// RelationAction 关注/取关操作
func RelationAction(ctx *gin.Context) {

	var rReq relationAction
	rReq.Token = ctx.Query("token")
	actype, _ := strconv.Atoi(ctx.Query("action_type"))
	rReq.ActionType = int32(actype)
	tUserId, _ := strconv.ParseInt(ctx.Query("to_user_id"), 10, 64)
	rReq.toUserId = tUserId

	friendListReq := pb.DouyinRelationActionRequest{
		Token:      rReq.Token,
		ToUserId:   rReq.toUserId,
		ActionType: rReq.ActionType,
	}

	_, err := rpc.RelationAction(ctx, &friendListReq)
	zap.L().Info("Start To Call RPC RelationAction Service!")
	if err != nil {
		zap.L().Error("Call Relation Action Service failed: ", zap.Error(err))
		models.ResponseErrorWithMsg(ctx, models.CodeServerBusy, err.Error())
		return
	}

	models.PostRelationRespSuccess(ctx)

}

func RelationFollowList(ctx *gin.Context) {
	var followReq RelationListReq
	userId, _ := jwt.GetUserID(ctx)
	zap.L().Info(fmt.Sprintf("userId:", userId))
	followReq.UserId = userId
	followReq.Token = ctx.Query("token")

	followListReq := pb.DouyinRelationFollowListRequest{
		UserId: followReq.UserId,
		Token:  followReq.Token,
	}

	relationResp, err := rpc.RelationFollowList(ctx, &followListReq)
	if err != nil {
		zap.L().Error("Post relation action rpc error", zap.Error(err))
		models.ResponseErrorWithMsg(ctx, models.CodeServerBusy, err.Error())
		return
	}

	userInfoList := make([]*models.User, 0)
	for _, m := range relationResp.UserList {
		resp, _ := json.Marshal(m)
		zap.L().Info(string(resp))

		userInfoList = append(userInfoList, &models.User{
			ID:            m.Id,
			Name:          m.Name,
			FollowCount:   m.FollowCount,
			FollowerCount: m.FollowerCount,
			WorkCount:     m.WorkCount,
			FavoriteCount: m.WorkCount,
			IsFollow:      m.IsFollow,

			Avatar:          m.Avatar,
			BackgroundImage: m.BackgroundImage,
			Signature:       m.Signature,
			TotalFavorited:  m.TotalFavorited,
		})
	}

	models.UserInfoListRespSuccess(ctx, userInfoList)
}

func RelationFanList(ctx *gin.Context) {
	var fansReq RelationListReq
	userId, _ := jwt.GetUserID(ctx)
	zap.L().Info(fmt.Sprintf("userId:", userId))
	fansReq.UserId = userId
	fansReq.Token = ctx.Query("token")

	followerListReq := pb.DouyinRelationFollowerListRequest{
		UserId: fansReq.UserId,
		Token:  fansReq.Token,
	}

	relationResp, err := rpc.RelationFollowerList(ctx, &followerListReq)
	if err != nil {
		zap.L().Error("Get fan List rpc error", zap.Error(err))
		models.ResponseErrorWithMsg(ctx, models.CodeServerBusy, err.Error())
		return
	}

	userInfoList := make([]*models.User, 0)
	for _, m := range relationResp.UserList {
		userInfoList = append(userInfoList, &models.User{
			ID:            m.Id,
			Name:          m.Name,
			FollowCount:   m.FollowCount,
			FollowerCount: m.FollowerCount,
			WorkCount:     m.WorkCount,
			FavoriteCount: m.WorkCount,
			IsFollow:      m.IsFollow,

			Avatar:          m.Avatar,
			BackgroundImage: m.BackgroundImage,
			Signature:       m.Signature,
			TotalFavorited:  m.TotalFavorited,
		})
	}

	models.UserInfoListRespSuccess(ctx, userInfoList)
}

func RelationFriendList(ctx *gin.Context) {
	var friendsReq RelationListReq
	userId, _ := jwt.GetUserID(ctx)
	zap.L().Info(fmt.Sprintf("userId:", userId))
	friendsReq.UserId = userId
	friendsReq.Token = ctx.Query("token")

	friendListReq := pb.DouyinRelationFriendListRequest{
		UserId: friendsReq.UserId,
		Token:  friendsReq.Token,
	}

	relationResp, err := rpc.RelationFriendList(ctx, &friendListReq)
	if err != nil {
		zap.L().Error("Get Friend List rpc error", zap.Error(err))
		models.ResponseErrorWithMsg(ctx, models.CodeServerBusy, err.Error())
		return
	}

	friendInfoList := make([]*models.FriendUser, 0)
	for _, m := range relationResp.UserList {
		friendInfoList = append(friendInfoList, &models.FriendUser{
			Msg:     m.Message,
			MsgType: m.MsgType,
		})
	}

	models.FriendInfoListRespSuccess(ctx, friendInfoList)
}
