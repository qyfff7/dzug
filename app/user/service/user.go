package service

import (
	"context"
	"dzug/app/user/dao"
	"dzug/app/user/pkg/jwt"
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
		zap.L().Error("用户登录失败", zap.Error(err))
		return nil, err
	}
	return resp, nil

}

func (s *Userservice) GetUserInfo(ctx context.Context, req *pb.GetUserInfoReq) (*pb.GetUserInfoResp, error) {

	//1.获取当前已经登录用户的id（未登录的话，提示需要登录）
	u, err := jwt.ParseToken(req.Token)
	if err != nil {
		//错误处理
		zap.L().Error("解析Token出错")
		return nil, err
	}
	userID := u.UserID

	//2.根据请求中视频作者的id，获取相应的作者信息
	uInfo, err := dao.GetuserInfoByID(ctx, req.UserId)
	if err != nil {
		zap.L().Error("获取视频用户信息失败", zap.Error(err))
		return nil, err
	}
	//3.从relation表中,查找出是否关注
	isfollow, err := dao.IsFollowByID(ctx, userID, req.UserId)
	if err != nil {
		zap.L().Error("查询是否关注信息出错！")
		return nil, err
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

	resp := &pb.GetUserInfoResp{
		/*StatusCode: 0,
		StatusMsg:  "获取用户信息成功",*/
		User: userInfo,
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
