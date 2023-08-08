package dao

// dao包：将所有的数据库操作封装成函数，根据业务需求进行调用

import (
	"context"
	"crypto/md5"
	"dzug/app/user/pkg/snowflake"
	"dzug/protos/user"
	"dzug/repo"
	"encoding/hex"
	"errors"
	"go.uber.org/zap"
)

const secret = "douyin"

var DB = repo.DB

// encryptPassword 密码加密
func encryptPassword(oPassword string) string {
	h := md5.New()
	h.Write([]byte(secret))
	return hex.EncodeToString(h.Sum([]byte(oPassword)))
}

// CheckUserExits 检查指定用户名的用户是否存在
func CheckUserExits(ctx context.Context, username string) (err error) {

	var user repo.User

	if DB.WithContext(ctx).Where("name = ?", username).First(&user).RowsAffected > 0 {
		zap.L().Info("当前用户名已存在,请更换用户名再尝试")
		err = errors.New("当前用户名已存在,请更换用户名再尝试。")
		return err
	}

	return nil

}

// InsertUser 用户注册相关数据库操作
func InsertUser(ctx context.Context, req *user.LoginAndRegisterRequest) (response *user.LoginAndRegisterResponse, err error) {

	//1.判断用户是否存在
	if err := CheckUserExits(ctx, req.Username); err != nil {
		return nil, err
	}

	//2.生成用户ID
	ID := snowflake.GenID()
	userID := uint64(ID)
	//3.用户密码加密
	password := encryptPassword(req.Password)

	//4.构建新用户结构体
	user := &repo.User{
		UserId:   userID,
		Name:     req.Username,
		Password: password,
	}
	//5.保存到数据库中
	err = DB.WithContext(ctx).Create(&user).Error
	if err != nil {
		zap.L().Error("create user data fail ", zap.Error(err))
		return nil, err
	}
	response.StatusCode = 0
	response.StatusMsg = "注册成功"
	response.UserId = userID
	response.Token = "Token - 注册成功" //这只是暂时的，后面写tocken认证

	return response, nil

}

