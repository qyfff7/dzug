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

func Init() error {
	var err error
	tmp := conf.Config.DSN()
	fmt.Println(tmp)
	DB, err = gorm.Open(mysql.Open(conf.Config.DSN()), &gorm.Config{
		PrepareStmt:            true,
		SkipDefaultTransaction: true,
	})
	if err != nil {
		zap.L().Error("MySQL初始化失败：" + err.Error())
		return err
	}
	err = DB.AutoMigrate(model.Video{})
	if err != nil {
		zap.L().Error("数据表创建失败：" + err.Error())
		return err
	}
	zap.L().Info("MySQL初始化成功")
	return nil
}
