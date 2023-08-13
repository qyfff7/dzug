package dao

// dao包：将所有的数据库操作封装成函数，根据业务需求进行调用

import (
	"context"
	"crypto/md5"
	"dzug/app/user/pkg/jwt"
	"dzug/app/user/pkg/snowflake"
	"dzug/models"
	"dzug/protos/user"
	"dzug/repo"
	"encoding/hex"
	"errors"
	"fmt"
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
func InsertUser(ctx context.Context, req *user.AccountReq) (*user.AccountResp, error) {

	//zap.L().Info("执行到InsertUser函数这里了")
	//1.判断用户是否存在
	exits, err := CheckUserExits(ctx, req.Username)

	if err != nil || exits {
		return nil, err
	}

	//2.生成用户ID
	userID := snowflake.GenID()

	//3.用户密码加密
	password := encryptPassword(req.Password)

	//4.构建新用户结构体
	newuser := &repo.User{
		UserId:   userID,
		Name:     req.Username,
		Password: password,
	}
	//zap.L().Info("构建新用户的结构体完毕")
	//5.保存到数据库中
	err = repo.DB.WithContext(ctx).Create(newuser).Error
	if err != nil {
		zap.L().Error("create user data fail ", zap.Error(err))
		return nil, err
	}
	zap.L().Info("用户注册成功！！！")
	//6.生成token
	token, err := jwt.GenToken(userID)
	if err != nil {
		zap.L().Error("生成tocken出错")
	}

	//7.返回相应

	resp := &user.AccountResp{
		UserId: newuser.UserId,
		Token:  token,
	}

	return resp, nil

}

func CheckAccount(ctx context.Context, req *user.AccountReq) (*user.AccountResp, error) {

	//构建登录用户
	u := repo.User{
		Name:     req.Username,
		Password: encryptPassword(req.Password),
	}

	result := repo.DB.WithContext(ctx).Where("name = ? AND password = ?", u.Name, u.Password).Limit(1).Find(&u)

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

	token, err := jwt.GenToken(u.UserId)
	if err != nil {
		zap.L().Error("token generation error")
	}

	resp := &user.AccountResp{
		UserId: u.UserId,
		Token:  token,
	}

	return resp, nil

}

func GetuserInfoByID(ctx context.Context, uid int64) (*models.User, error) {
	userInfo := new(models.User)
	//1.从user表中查找出用户的个人信息
	zap.L().Info("按照id查询用户信息")

	zap.L().Info(fmt.Sprintln(uid))
	zap.L().Info("======================")
	result := repo.DB.WithContext(ctx).Where("user_id = ? ", uid).Limit(1).Find(&userInfo)

	if result.Error != nil {
		zap.L().Info("执行获取用户信息时gorm出错")
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		zap.L().Info("当前用户不存在，请重试")
		err := errors.New("当前用户不存在，请重试")
		return nil, err
	}
	return userInfo, nil

}

// IsFollowByID 判断是否关注了该用户
func IsFollowByID(ctx context.Context, userID, autherID int64) (bool, error) {
	var rel repo.Relation
	result := repo.DB.WithContext(ctx).Table("relation").Where("user_id = ? AND to_user_id = ?", userID, autherID).Limit(1).Find(&rel)
	if result.Error != nil {
		zap.L().Info("查找关注关系时出错")
		return false, result.Error
	}
	if result.RowsAffected > 0 { //关注了该用户
		return true, nil
	}
	return false, nil //未关注
}
