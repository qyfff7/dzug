package service

import (
	"context"
	"dzug/app/user/dao"
	pb "dzug/protos/user"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type Userservice struct {
	db *gorm.DB
	pb.UnimplementedServiceServer
}

func (s *Userservice) Register(c context.Context, req *pb.LoginAndRegisterRequest) (*pb.LoginAndRegisterResponse, error) {
	resp, err := dao.InsertUser(c, req)
	if err != nil {
		zap.L().Error("用户注册失败", zap.Error(err))
		return nil, err
	}
	return resp, nil
}
func (s *Userservice) Login(ctx context.Context, req *pb.LoginAndRegisterRequest) (*pb.LoginAndRegisterResponse, error) {

	//dao层进行数据库查询操作
	resp, err := dao.Login(ctx, req)
	if err != nil {
		zap.L().Error("用户登录失败", zap.Error(err))
		return nil, err
	}
	return resp, nil

}

func (s *Userservice) UserInfo(ctx context.Context, req *pb.UserInfoRequest) (*pb.UserInfoResponse, error) {

	resp, err := dao.GetuserInfo(ctx, req)
	if err != nil {
		zap.L().Error("获取用户信息失败", zap.Error(err))
		return nil, err
	}
	return resp, nil

}

/*
func (s *user_service) UserMap(ctx context.Context, req *user.UserMapRequest) (*user.UserMapResponse, error) {

	// 1、获取用户列表 []User
	userPoRes, err := s.userList(ctx, req.UserIds)

	// 这里为什么不把错误合并在一起返回，因为有可能这里已经报错了。就没必要往后面操作了
	if err != nil {
		switch e := err.(type) {
		case *custom.Exception:
			return nil, status.Error(codes.NotFound, e.Error())
		default:
			return nil, status.Error(codes.Unknown, e.Error())
		}
	}

	// 将Token放入Ctx
	tkCtx := context.WithValue(ctx, constant.REQUEST_TOKEN, req.Token)

	// 2、转换为 Map[UserId] = User
	UserMap := make(map[int64]*user.User)
	for _, po := range userPoRes {
		vo := po.Po2vo()
		err = s.composeInfo(tkCtx, vo)
		if err != nil {
			return nil, err
		}
		UserMap[vo.Id] = vo
	}

	return &user.UserMapResponse{UserMap: UserMap}, nil
}

func (s *user_service) composeInfo(ctx context.Context, uResp *user.User) error {

	var (
		wait = sync.WaitGroup{}
		errs = make([]error, 0)
	)

	wait.Add(3)

	// 组合 followListCount、followerListCount、isFollow
	go func() {
		defer wait.Done()

		errs = append(errs, s.composeRelation(ctx, uResp)...)
	}()

	// 组合 publishCount
	go func() {
		defer wait.Done()

		errs = append(errs, s.composeVideo(ctx, uResp)...)
	}()

	// 组合 favoriteCount
	go func() {
		defer wait.Done()

		errs = append(errs, s.composeFavorite(ctx, uResp)...)
	}()

	wait.Wait()

	// 查看后台调用时，是否有错误产生
	for _, err := range errs {
		if err != nil {
			switch e := err.(type) {
			case *custom.Exception:
				return status.Error(codes.NotFound, e.Error())
			default:
				return status.Error(codes.Unknown, e.Error())
			}
		}
	}

	return nil
}*/
