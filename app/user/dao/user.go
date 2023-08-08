package dao

// dao包：将所有的数据库操作封装成函数，根据业务需求进行调用

import (
	"dzug/app/user/pkg/snowflake"
	"errors"
	db "gorm.io/gorm"
)

// CheckUserExits 检查指定用户名的用户是否存在
func CheckUserExits(username string) (err error) {

	sqlStr := `select count(user_id) from user where username = ?`
	var count int64
	if err := db.Get(&count, sqlStr, username); err != nil {
		return err
	}
	if count > 0 {
		return errors.New("用户已存在")
	}

	return
}

// InsertUser 用户注册相关数据库操作
func InsertUser() {
	//1.判断用户是否存在
	if err := CheckUserExits(); err != nil {

	}
	//2.生成用户ID
	snowflake.GenID()
	//3.用户密码加密
	//4.保存到数据库中
}
