package service

import (
	"context"
	"dzug/app/user/dao"
	"dzug/protos/user"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type Userservice struct {
	db *gorm.DB
	user.UnimplementedServiceServer
}

func (s *Userservice) Register(c context.Context, req *user.LoginAndRegisterRequest) (*user.LoginAndRegisterResponse, error) {
	resp, err := dao.InsertUser(c, req)
	if err != nil {
		zap.L().Error("用户注册失败", zap.Error(err))
		return nil, err
	}
	return resp, nil
}
func (s *Userservice) Login(ctx context.Context, req *user.LoginAndRegisterRequest) (*user.LoginAndRegisterResponse, error) {

	//1.dao层进行数据库查询操作
	resp, err := dao.Login(ctx, req)
	if err != nil {
		zap.L().Error("用户登录失败", zap.Error(err))
		return nil, err
	}
	return resp, nil

}

// Register 用户注册
/*func (s *user_service) Register(ctx context.Context, req *user.LoginAndRegisterRequest) (*user.TokenResponse, error) {

	// 1、请求参数校验
	if err := req.Validate(); err != nil {
		return nil, status.Error(codes.InvalidArgument,
			constant.Code2Msg(constant.ERROR_ARGS_VALIDATE))
	}

	// 2、根据 Username 查询此用户是否已经注册
	userReq := newGetUserReq()
	userReq.Username = req.Username
	po, err := s.getUser(ctx, userReq)
	if err != nil {
		return nil, status.Error(codes.Unavailable,
			constant.Code2Msg(constant.ERROR_ACQUIRE))
	}

	if po.Id > 0 {
		// 用户已存在
		return nil, status.Error(codes.AlreadyExists,
			constant.Code2Msg(constant.WRONG_EXIST_USERS))
	}

	// 3、未注册-创建用户，注册-返回提示
	po = user.NewUserPo(req.Hash())
	insertRes, err := s.insert(ctx, po)

	if err != nil {
		return nil, status.Error(codes.Unknown,
			constant.Code2Msg(constant.ERROR_SAVE))
	}

	// 4、颁发Token并返回
	response := user.NewTokenResponse(insertRes.Id, s.token(ctx, insertRes))

	return response, nil
}
*/

/*
func (s *user_service) UserInfo(ctx context.Context, req *user.UserInfoRequest) (*user.UserInfoResponse, error) {

	// 请求参数校验
	if err := req.Validate(); err != nil {
		s.l.Errorf("user UserInfo：参数校验失败，%s", err.Error())
		return nil, status.Error(codes.InvalidArgument,
			constant.Code2Msg(constant.ERROR_ARGS_VALIDATE))
	}

	response := user.NewUserInfoResponse()
	response.User = user.NewDefaultUser()
	// get user info, user += userInfo

	userReq := newGetUserReq()
	userReq.UserId = req.UserId
	po, err := s.getUser(ctx, userReq)
	if err != nil {
		return nil, status.Error(codes.Unavailable,
			constant.Code2Msg(constant.ERROR_ACQUIRE))
	}
	response.User = po.Po2vo()

	// 将Token放入Ctx
	tkCtx := context.WithValue(ctx, constant.REQUEST_TOKEN, req.Token)

	return response, s.composeInfo(tkCtx, response.User)
}

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
