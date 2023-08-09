package dal

import (
	"dzug/app/publish/infra/dal/model"
	"dzug/conf"
	"fmt"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Init() {
	var err error
	tmp := conf.Config.DSN()
	fmt.Println(tmp)
	DB, err = gorm.Open(mysql.Open(conf.Config.DSN()), &gorm.Config{
		PrepareStmt:            true,
		SkipDefaultTransaction: true,
	})
	if err != nil {
		zap.L().Error("初始化数据库失败" + err.Error())
	}
	err = DB.AutoMigrate(model.Video{})
	if err != nil {
		zap.L().Error("初始化数据表失败" + err.Error())
	}
	zap.L().Info("数据库初始化成功")

}
