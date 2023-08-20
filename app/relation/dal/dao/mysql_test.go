package dao

import (
	"context"
	"dzug/repo"
	"errors"
	"fmt"
	"testing"

	"github.com/magiconair/properties/assert"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func Init() {
	var err error
	dsn := "root:root@tcp(127.0.0.1:3306)/douyin?charset=utf8mb4&parseTime=True&loc=Local"

	db, err = gorm.Open(mysql.Open(dsn),
		&gorm.Config{
			PrepareStmt:            true,
			SkipDefaultTransaction: true,
		},
	)
	if err != nil {
		panic(err)
	}
}

func TestMysql(t *testing.T) {
	Init()
	// // 测试关注别人 已经关注过 -> 会出错
	// fmt.Println("测试关注别人")
	// err := FollowUser_test(context.TODO(), 7093779806818304, 7093627553583104)
	// if err != nil {
	// 	fmt.Printf("err: %v\n", err)
	// }
	// assert.Equal(t, err, nil)

	// // 测试取关别人
	// fmt.Println("测试取关别人")
	// err = UnFollowUser_test(context.TODO(), 7093779806818304, 7093627553583104)
	// if err != nil {
	// 	fmt.Printf("err: %v\n", err)
	// }
	// assert.Equal(t, err, nil)

	// err := FollowUser_test(context.TODO(), 7093779806818304, 7093627553583104)
	// err = FollowUser_test(context.TODO(), 7093779806818304, 5191404976345088)
	// err = FollowUser_test(context.TODO(), 7093779806818304, 5477693852225536)
	// err = FollowUser_test(context.TODO(), 7093779806818304, 5897115414827008)

	// err = FollowUser_test(context.TODO(), 5897088445452288, 7093779806818304)
	// err = FollowUser_test(context.TODO(), 5897088445452288, 5477693852225536)
	// err = FollowUser_test(context.TODO(), 5897088445452288, 7093627553583104)

	// err = FollowUser_test(context.TODO(), 7093627553583104, 7093779806818304)
	// err = FollowUser_test(context.TODO(), 7093627553583104, 5897088445452288)
	// err = FollowUser_test(context.TODO(), 7093627553583104, 5191404976345088)
	// err = FollowUser_test(context.TODO(), 7093627553583104, 5477693852225536)

	// 测试关注列表
	fmt.Println("测试关注列表")
	followIds, err := GetFollowList_test(context.TODO(), 7093779806818304)
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
	assert.Equal(t, err, nil)

	for _, followid := range followIds {
		fmt.Printf("followid: %v\n", followid)
	}

	// 测试粉丝列表      粉丝列表个数好像不对
	fmt.Println("测试粉丝列表")
	followerIds, err := GetFollowerList_test(context.TODO(), 5477693852225536)
	assert.Equal(t, err, nil)

	for _, followerid := range followerIds {
		fmt.Printf("followerid: %v\n", followerid)
	}

	// 测试好友列表
	fmt.Println("测试好友列表")
	friendids, err := GetFriendList_test(context.TODO(), 7093779806818304)
	assert.Equal(t, err, nil)

	for _, friendid := range friendids {
		fmt.Printf("friendid: %v\n", friendid)
	}

}

// 关注
func FollowUser_test(ctx context.Context, userID, toUserID int64) error {
	if userID == toUserID {
		return errors.New("不能关注自己")
	}

	// 检查是否已经关注过用户
	var tmp repo.Relation
	err := db.WithContext(ctx).Table("relation").Where("user_id = ? AND to_user_id = ?", userID, toUserID).First(&tmp).Error
	if err == nil {
		return errors.New("已经关注过该用户")
	} else if !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}

	relation := repo.Relation{
		UserId:   userID,
		ToUserId: toUserID,
	}

	// 开启事务：
	err = db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		// 创建关注关系
		if err := tx.Table("relation").Create(&relation).Error; err != nil {
			zap.L().Error("关注别人，MySQL数据库添加关注关系失败: ", zap.Error(err))
			return err
		}

		// 更新关注者的关注数
		if err := tx.Table("user").Where("user_id = ?", userID).Update("follow_count", gorm.Expr("follow_count + ?", 1)).Error; err != nil {
			zap.L().Error("关注别人，MySQL数据库更新关注者的关注数失败: ", zap.Error(err))
			return err
		}

		// 更新被关注者的粉丝数
		if err := tx.Table("user").Where("user_id = ?", toUserID).Update("follower_count", gorm.Expr("follower_count + ?", 1)).Error; err != nil {
			zap.L().Error("关注别人，MySQL数据库更新被关注者的粉丝数失败: ", zap.Error(err))
			return err
		}

		return nil
	})

	if err != nil {
		zap.L().Error("MySQL关注别人失败: ", zap.Error(err))
		return err
	}

	zap.L().Info("MySQL成功关注别人")
	return nil
}

// 获取粉丝列表
func GetFollowerList_test(ctx context.Context, userID int64) ([]int64, error) {
	var followerList []repo.Relation

	// 在 relation 表中查找所有 to_user_id 字段等于给定 userID 的数据
	err := db.WithContext(ctx).Where("to_user_id = ?", userID).Find(&followerList).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		fmt.Printf("MySQL数据库中查询到该用户没有粉丝, err: %v\n", err)
		return make([]int64, 0), err
	}
	if err != nil {
		fmt.Printf("MySQL数据库中查找该用户的粉丝失败, err: %v\n", err)
		return nil, err
	}

	// 提取UserID的所有被关注关系中的 粉丝ID
	userIDs := make([]int64, len(followerList))
	for i, relation := range followerList {
		userIDs[i] = int64(relation.UserId)
	}

	fmt.Printf("MySQL数据库成功获取粉丝列表, err: %v\n", err)
	return userIDs, nil
}

// 获取关注列表
func GetFollowList_test(ctx context.Context, userID int64) ([]int64, error) {
	var followList []repo.Relation
	// 在 relation 表中查找所有 user_id 字段等于给定 userID 的数据
	err := db.WithContext(ctx).Where("user_id = ?", userID).Find(&followList).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		fmt.Printf("MySQL数据库中查询到该用户没有关注者, err: %v\n", err)
		return make([]int64, 0), err
	}
	if err != nil {
		fmt.Printf("MySQL数据库中查询该用户的关注者失败, err: %v\n", err)
		return nil, err
	}
	// 提取UserID的所有关注关系中的 关注ID
	toUserIDs := make([]int64, len(followList))
	for i, relation := range followList {
		toUserIDs[i] = int64(relation.ToUserId)
	}

	fmt.Printf("MySQL数据库成功获取关注列表, err: %v\n", err)
	return toUserIDs, nil
}

// 取关
func UnFollowUser_test(ctx context.Context, userID, toUserID int64) error {
	if userID == toUserID {
		fmt.Println("不能取关自己")
		return errors.New("不能取关自己")
	}

	// 检查是否关注过用户 前提：数据库中必须有该条数据
	var temp repo.Relation
	err := db.WithContext(ctx).Table("relation").Where("user_id = ? AND to_user_id = ?", userID, toUserID).First(&temp).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		fmt.Printf("MySQL数据库中查询到 用户没有关注过该用户，不能取关, err: %v\n", err)
		return errors.New("用户没有关注过该用户，不能取关")
	} else if err != nil {
		fmt.Printf("MySQL数据库查询 关注关系失败, err: %v\n", err)
		return err
	}

	relation := repo.Relation{
		UserId:   userID,
		ToUserId: toUserID,
	}

	// 开启事务：
	err = db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		// 更新关注者的关注数
		if err := tx.Table("user").Where("user_id = ?", userID).Update("follow_count", gorm.Expr("follow_count - ?", 1)).Error; err != nil {
			zap.L().Error("取关别人，MySQL数据库更新关注者的关注数失败: ", zap.Error(err))
			return err
		}

		// 更新被关注者的粉丝数
		if err := tx.Table("user").Where("user_id = ?", toUserID).Update("follower_count", gorm.Expr("follower_count - ?", 1)).Error; err != nil {
			zap.L().Error("取关别人，MySQL数据库更新被关注者的粉丝数失败: ", zap.Error(err))
			return err
		}

		// 删除关注关系
		if err := tx.Table("relation").Where("user_id = ? AND to_user_id = ?", userID, toUserID).Delete(&relation).Error; err != nil {
			zap.L().Error("取关别人，MySQL数据库软删除关注关系失败: ", zap.Error(err))
			return err
		}
		return nil
	})

	if err != nil {
		fmt.Printf("MySQL数据库取关别人失败, err: %v\n", err)
		return err
	}

	fmt.Printf("MySQL数据库成功取关别人\n")

	return nil
}

// 获取好友列表(互关)
func GetFriendList_test(ctx context.Context, userID int64) ([]int64, error) {
	var friendList []repo.Relation

	// 在 relation 表中查找当前UserID互关的数据
	err := db.WithContext(ctx).Where("user_id = ? AND to_user_id IN (?)", userID, db.WithContext(ctx).Table("relation").Select("user_id").Where("to_user_id = ?", userID)).Find(&friendList).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		fmt.Printf("MySQL数据库中查询到该用户没有好友, err: %v\n", err)
		return make([]int64, 0), err
	}
	if err != nil {
		fmt.Printf("MySQL数据库中查找该用户的好友失败, err: %v\n", err)
		return nil, err
	}
	userIDS := make([]int64, len(friendList))
	for i, fan := range friendList {
		userIDS[i] = int64(fan.UserId)
	}

	fmt.Printf("MySQL数据库成功获取好友列表\n")
	return userIDS, nil
}
