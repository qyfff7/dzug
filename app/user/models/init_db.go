package models

import (
	"dzug/conf"
	"fmt"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() (err error) {
	// 链接数据库 --格式: 用户名:密码@协议(IP:port)/数据库名？xxx&yyy&
	//conn, err := gorm.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/test")
	DBConnectString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=%v&loc=%s",
		conf.Config.MySQLConfig.User, conf.Config.MySQLConfig.Password, conf.Config.MySQLConfig.Host,
		conf.Config.MySQLConfig.Port, conf.Config.MySQLConfig.DB, conf.Config.MySQLConfig.Charset,
		conf.Config.MySQLConfig.ParseTime, conf.Config.MySQLConfig.Loc,
	)
	DB, err = gorm.Open(mysql.Open(DBConnectString), &gorm.Config{
		PrepareStmt:            true, //缓存预编译命令
		SkipDefaultTransaction: true, //禁用默认事务操作
		//Logger:                 logger.Default.LogMode(logger.Global), //打印sql语句
	})
	if err != nil {
		zap.L().Error("Database connection failure，", zap.Error(err))
		return
	}
	if err = DB.AutoMigrate(&User{}); err != nil {
		zap.L().Error("Failed to create database,", zap.Error(err))
		return
	}
	return
}
func InsertData() (err error) {
	// 先创建数据 --- 创建对象
	var u User
	u.Name = "zhangsan3"
	u.Password = "aaa"
	u.FollowCount = 1
	u.FollowerCount = 2
	u.WorkCount = 3
	u.FavoriteCount = 4

	// 插入(创建)数据
	if err = DB.Create(&u).Error; err != nil {
		zap.L().Error("插入数据库失败,", zap.Error(err))
		return
	}
	return
}
