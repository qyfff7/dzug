package repo

import (
	"dzug/conf"
	"fmt"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

// DB 提供给对外操作数据库
var DB *gorm.DB

func Init() (err error) {

	//conf.Init() // 初始化配置文件

	mysqlConfig := conf.Config.MySQLConfig
	link := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		mysqlConfig.User, mysqlConfig.Password, mysqlConfig.Host, mysqlConfig.Port, mysqlConfig.DB)
	DB, err = gorm.Open(mysql.Open(link), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
		NamingStrategy: &schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		zap.L().Error(link + "连接数据库失败" + err.Error())
		panic(err)
	}
	err = DB.AutoMigrate(User{}, Video{}, Comment{}, Message{}, Favorite{}, Relation{}) // 迁移数据表
	if err != nil {
		zap.L().Error("数据表初始化失败")
		panic(err)
	}
	return err
}
