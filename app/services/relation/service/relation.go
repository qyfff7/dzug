package service

import (
	"context"
	"dzug/app/services/relation/dal/dao"
	"dzug/app/services/user/pkg/jwt"
	"dzug/protos/relation"
	"errors"
	"fmt"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

type RelationSrv struct {
	db *gorm.DB
	relation.UnimplementedDouyinRelationActionServiceServer
}

// 在此处执行关注或取消关注的逻辑
// 获取参数：token 用户token，对方用户to_user_id，action操作：1-关注,2-取消关注
// 返回参数：status_code 状态码 0-成功/其他值-失败，status_msg 返回状态描述
func (r *RelationSrv) DouyinRelationAction(ctx context.Context, req *relation.DouyinRelationActionRequest) (*relation.DouyinRelationActionResponse, error) {

	zap.L().Info("开始进行关注或取关操作")

	if req.Token == "" {
		return nil, errors.New("token is null")
	}

	u, _ := jwt.ParseToken(req.Token)

	userID := u.UserID // 从 token 中获取用户 ID
	toUserID := req.ToUserId
	actionType := req.ActionType

	if actionType == 1 {
		err := dao.FollowUser(ctx, userID, toUserID)
		if err != nil {
			return &relation.DouyinRelationActionResponse{
				StatusCode: 500,
				StatusMsg:  "关注用户失败：" + err.Error(),
			}, nil
		}
	} else if actionType == 2 {
		err := dao.UnFollowUser(ctx, userID, toUserID)
		if err != nil {
			return &relation.DouyinRelationActionResponse{
				StatusCode: 500,
				StatusMsg:  "取消关注用户失败：" + err.Error(),
			}, nil
		}
	} else {
		return &relation.DouyinRelationActionResponse{
			StatusCode: 400,
			StatusMsg:  "无效的关系操作类型",
		}, nil
	}

	return &relation.DouyinRelationActionResponse{
		StatusCode: 0,
		StatusMsg:  "关系操作成功",
	}, nil
}

func (r *RelationSrv) DouyinRelationFollowList(ctx context.Context, req *relation.DouyinRelationFollowListRequest) (*relation.DouyinRelationFollowListResponse, error) {
	// 在此处获取用户的关注列表
	// 获取参数：user_id用户id，token用户的token
	// 返回参数：status_code状态码 0:成功 其他值:失败，status_msg返回状态描述，User user_list用户信息列表
	zap.L().Info("开始获取用户的关注列表, userid: " + string(req.UserId))
	userID := req.UserId
	followIdList, err := dao.GetFollowList(ctx, userID)
	userInfoList, err := GetUserInfoList(ctx, userID, followIdList)
	if err != nil {
		return &relation.DouyinRelationFollowListResponse{
			StatusCode: 500,
			StatusMsg:  "获取关注列表失败：" + err.Error(),
		}, nil
	}
	return &relation.DouyinRelationFollowListResponse{
		StatusCode: 0,
		StatusMsg:  "获取关注列表成功",
		UserList:   userInfoList,
	}, nil
}

func (r *RelationSrv) DouyinRelationFollowerList(ctx context.Context, req *relation.DouyinRelationFollowerListRequest) (*relation.DouyinRelationFollowerListResponse, error) {
	// 在此处获取用户的粉丝列表
	// 获取参数：user_id用户id，token用户的token
	// 返回参数：status_code状态码 0:成功 其他值:失败，status_msg返回状态描述，User user_list用户信息列表

	userID := req.UserId
	fmt.Printf("userID: %v\n", userID)
	followerIdList, err := dao.GetFollowerList(ctx, userID)
	userInfoList, err := GetUserInfoList(ctx, userID, followerIdList)
	if err != nil {
		return &relation.DouyinRelationFollowerListResponse{
			StatusCode: 500,
			StatusMsg:  "获取粉丝列表失败：" + err.Error(),
		}, nil
	}
	return &relation.DouyinRelationFollowerListResponse{
		StatusCode: 0,
		StatusMsg:  "获取粉丝列表成功",
		UserList:   userInfoList,
	}, nil
}

func (r *RelationSrv) DouyinRelationFriendList(ctx context.Context, req *relation.DouyinRelationFriendListRequest) (*relation.DouyinRelationFriendListResponse, error) {
	// 在此处获取用户的好友列表
	// 获取参数：user_id用户id，token用户的token
	// 返回参数：status_code状态码 0:成功 其他值:失败，status_msg返回状态描述，返回表 FriendUser user_list
	// FriendUser: User + (message = 1; 和该好友的最新聊天消息,message消息的类型，0 => 当前用户接收的消息， 1 => 当前用户发送的消息)

	userID := req.UserId
	zap.L().Info("开始获取用户的好友列表, userid: " + string(userID))
	fmt.Printf("userID: %v\n", userID)
	friendIdList, err := dao.GetFriendList(ctx, userID)

	friendInfoList, err := GetUserInfoList(ctx, userID, friendIdList)

	if err != nil {
		return &relation.DouyinRelationFriendListResponse{
			StatusCode: 500,
			StatusMsg:  "获取好友列表失败：" + err.Error(),
		}, nil
	}
	return &relation.DouyinRelationFriendListResponse{
		StatusCode: 0,
		StatusMsg:  "获取好友列表成功",
		UserList:   friendInfoList,
	}, nil
}

func GetUserInfoList(ctx context.Context, UserID int64, followIdList []int64) ([]*relation.User, error) {
	zap.L().Info("已经获取关注或者粉丝的id列表，开始获取用户的具体信息")

	// 根据idList获取所有的信息list
	users, err := dao.GetUsersByIDList(ctx, UserID, followIdList)
	if err != nil {
		return nil, err
	}

	var userProtos []*relation.User
	for _, user := range users {
		userProto := &relation.User{
			Id:              user.UserId,
			Name:            user.Name,
			FollowCount:     user.FollowCount,
			FollowerCount:   user.FollowerCount,
			IsFollow:        user.IsFollow,
			Avatar:          user.Avatar,
			BackgroundImage: user.BackgroundImage,
			Signature:       user.Signature,
			TotalFavorited:  user.TotalFavorited,
			WorkCount:       user.WorkCount,
			FavoriteCount:   user.FavoriteCount,
		}
		userProtos = append(userProtos, userProto)
	}
	return userProtos, nil
}
