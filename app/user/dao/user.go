package dao

// dao包：将所有的数据库操作封装成函数，根据业务需求进行调用

import (
	"context"
	"crypto/md5"
	"dzug/app/user/pkg/jwt"
	"dzug/app/user/pkg/snowflake"
	"dzug/protos/user"
	"dzug/repo"
	"encoding/hex"
	"errors"
	"go.uber.org/zap"
)

const secret = "douyin"

// encryptPassword 密码加密
func encryptPassword(oPassword string) string {
	h := md5.New()
	h.Write([]byte(secret))
	return hex.EncodeToString(h.Sum([]byte(oPassword)))
}

// CheckUserExits 检查指定用户名的用户是否存在
func CheckUserExits(ctx context.Context, username string) (check bool, err error) {

	var user repo.User
	//zap.L().Info("开始执行CheckUserExits")
	result := repo.DB.WithContext(ctx).Where("name = ?", username).Limit(1).Find(&user)

	if result.Error != nil {
		zap.L().Info("查询是否存在当前用户时出错")
		return false, result.Error
	}

	if result.RowsAffected > 0 {
		zap.L().Info("当前用户名已存在,请更换用户名再尝试")
		err = errors.New("当前用户名已存在,请更换用户名再尝试。")
		return true, err
	}

	//zap.L().Info("执行到CheckUserExits函数这里了，执行完毕")
	return false, nil

}

// InsertUser 用户注册相关数据库操作
func InsertUser(ctx context.Context, req *user.LoginAndRegisterRequest) (*user.LoginAndRegisterResponse, error) {

	//zap.L().Info("执行到InsertUser函数这里了")
	//1.判断用户是否存在
	exits, err := CheckUserExits(ctx, req.Username)

	if err != nil || exits {
		return nil, err
	}

	//2.生成用户ID
	ID := snowflake.GenID()
	userID := uint64(ID)
	//3.用户密码加密
	password := encryptPassword(req.Password)

	//4.构建新用户结构体
	newuser := &repo.User{
		UserId:   int64(userID),
		Name:     req.Username,
		Password: password,
	}
	//zap.L().Info("构建新用户的结构体完毕")
	//5.保存到数据库中
	err = repo.DB.WithContext(ctx).Create(&newuser).Error
	if err != nil {
		zap.L().Error("create user data fail ", zap.Error(err))
		return nil, err
	}
	zap.L().Info("用户注册成功！！！")
	//6.生成token
	token, err := jwt.GenToken(uint64(newuser.UserId))
	if err != nil {
		zap.L().Error("生成tocken出错")
	}

	//7.返回相应
	resp := &user.LoginAndRegisterResponse{
		StatusCode: 0,
		StatusMsg:  "注册成功",
		UserId:     uint64(newuser.UserId),
		Token:      token,
	}

	return resp, nil

}

func Login(ctx context.Context, req *user.LoginAndRegisterRequest) (*user.LoginAndRegisterResponse, error) {

	//构建登录用户
	u := repo.User{
		Name:     req.Username,
		Password: encryptPassword(req.Password),
	}
	result := repo.DB.WithContext(ctx).Where("name = ? AND password >= ?", u.Name, u.Password).Limit(1).Find(&u)

	if result.Error != nil {
		zap.L().Info("执行用户登录sql查询时出错")
		return nil, result.Error
	}

	if result.RowsAffected == 0 {
		zap.L().Info("用户名或密码错误")
		err := errors.New("用户名或密码错误")
		return nil, err
	}
	zap.L().Info("User login successful！！！")

	token, err := jwt.GenToken(uint64(u.UserId))
	if err != nil {
		zap.L().Error("token generation error")
	}

	resp := &user.LoginAndRegisterResponse{
		StatusCode: 0,
		StatusMsg:  "用户登录成功",
		UserId:     uint64(u.UserId),
		Token:      token,
	}

	return resp, nil

}

func GetuserInfo(ctx context.Context, req *user.UserInfoRequest) (*user.UserInfoResponse, error) {
	//1.从user表中查找出用户的个人信息
	uInfo := new(repo.User)
	result := repo.DB.WithContext(ctx).Where("user_id = ? ", req.UserId).Limit(1).Find(&uInfo)

	if result.Error != nil {
		zap.L().Info("执行获取用户信息时出错")
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		zap.L().Info("未找到当前用户")
		err := errors.New("未找到当前用户")
		return nil, err
	}
	//2.获取当前视频的用户ID

	//3.从relation表中,查找出是否关注
	var to_user_id uint64
	to_user_id = 211

	isfollow, err := IsFollowByID(ctx, req.UserId, to_user_id)
	if err != nil {
		return nil, err
	}
	//3.构建返回结构
	userInfo := &user.User{
		Id:              uint64(uInfo.UserId),
		Name:            uInfo.Name,
		FollowCount:     &uInfo.FollowCount,
		FollowerCount:   &uInfo.FollowerCount,
		Avatar:          &uInfo.Avatar,
		BackgroundImage: &uInfo.BackgroundImages,
		Signature:       &uInfo.Signature,
		TotalFavorited:  &uInfo.TotalFavorited,
		WorkCount:       &uInfo.WorkCount,
		FavoriteCount:   &uInfo.FavoriteCount,
		IsFollow:        isfollow,
	}

	//zap.L().Info("执行到dao/user/GetuserInfo函数这里了")

	resp := &user.UserInfoResponse{
		StatusCode: 0,
		StatusMsg:  "获取用户信息成功",
		User:       userInfo,
	}
	return resp, nil

}

// IsFollowByID 判断是否关注了该用户
func IsFollowByID(ctx context.Context, userID, touserId uint64) (bool, error) {
	var rel repo.Relation
	result := repo.DB.WithContext(ctx).Table("relation").Where("user_id = ? AND to_user_id = ?", userID, touserId).Limit(1).Find(&rel)
	if result.Error != nil {
		zap.L().Info("查找关注关系时出错")
		return false, result.Error
	}
	if result.RowsAffected > 0 { //关注了该用户
		return true, nil
	}
	return false, nil //未关注

}
