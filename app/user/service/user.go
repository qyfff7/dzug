package service

import (
	"context"
	"dzug/app/user/dao"
	"dzug/app/user/pkg/jwt"
	"dzug/models"
	pb "dzug/protos/user"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type Userservice struct {
	db *gorm.DB
	pb.UnimplementedServiceServer
}

func (s *Userservice) Register(c context.Context, req *pb.AccountReq) (*pb.AccountResp, error) {
	resp, err := dao.InsertUser(c, req)
	if err != nil {
		zap.L().Error("用户注册失败", zap.Error(err))
		return nil, err
	}
	return resp, nil
}
func (s *Userservice) Login(ctx context.Context, req *pb.AccountReq) (*pb.AccountResp, error) {

	//dao层进行数据库查询操作
	resp, err := dao.CheckAccount(ctx, req)
	if err != nil {
		zap.L().Error("用户登录失败：", zap.Error(err))
		return nil, err
	}
	return resp, nil

}

func (s *Userservice) GetUserInfo(ctx context.Context, req *pb.GetUserInfoReq) (resp *pb.GetUserInfoResp, err error) {
	var uInfo *models.User
	isfollow := false

	if req.Token != "" {
		//1.获取当前已经登录用户的id（未登录的话，提示需要登录）
		u, err := jwt.ParseToken(req.Token)
		if err != nil {
			zap.L().Error("解析Token出错", zap.Error(err))
			return nil, err
		}

		//查询本人的信息  或未登录查视频作者
		if u.UserID == req.UserId {
			uInfo, err = dao.GetuserInfoByID(ctx, req.UserId)
			if err != nil {
				zap.L().Error("获取用户个人信息失败", zap.Error(err))
				return nil, err
			}
		} else {
			//2.根据请求中视频作者的id，获取相应的作者信息
			uInfo, err = dao.GetuserInfoByID(ctx, req.UserId)
			if err != nil {
				zap.L().Error("获取视频用户信息失败", zap.Error(err))
				return nil, err
			}
			//3.从relation表中,查找出是否关注
			isfollow, err = dao.IsFollowByID(ctx, u.UserID, req.UserId)
			if err != nil {
				zap.L().Error("查询是否关注信息出错！")
				return nil, err
			}
		}

	} else {
		uInfo, err = dao.GetuserInfoByID(ctx, req.UserId)
		if err != nil {
			zap.L().Error("获取用户个人信息失败", zap.Error(err))
			return nil, err
		}
	}

	//3.构建返回结构
	userInfo := &pb.User{
		Id:              uInfo.ID,
		Name:            uInfo.Name,
		FollowCount:     uInfo.FollowCount,
		FollowerCount:   uInfo.FollowerCount,
		Avatar:          uInfo.Avatar,
		BackgroundImage: uInfo.BackgroundImage,
		Signature:       uInfo.Signature,
		TotalFavorited:  uInfo.TotalFavorited,
		WorkCount:       uInfo.WorkCount,
		FavoriteCount:   uInfo.FavoriteCount,
		IsFollow:        isfollow,
	}
	resp = &pb.GetUserInfoResp{
		User: userInfo,
	}
	return resp, nil
}
