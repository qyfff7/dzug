package redis

// import (
// 	"errors"
// 	"fmt"
// 	"strconv"
// 	"testing"
// 	"time"

// 	"github.com/magiconair/properties/assert"
// 	"github.com/redis/go-redis/v9"
// 	"go.uber.org/zap"
// 	"golang.org/x/net/context"
// )

// var (
// 	rrdb *redis.Client
// )

// const customExpiration = 30 * time.Minute // 设置为 24 小时
// func InitRedisConnection() {
// 	var err error
// 	ctx := context.Background()
// 	// 设置 Redis 连接参数
// 	options := &redis.Options{
// 		Addr:         "localhost:6379", // Redis 服务器地址
// 		Password:     "",               // 密码
// 		DB:           0,                // 数据库编号
// 		PoolSize:     10,               // 连接池大小
// 		MinIdleConns: 5,                // 最小空闲连接数
// 	}
// 	// 创建 Redis 客户端
// 	client := redis.NewClient(options)
// 	if err = client.Ping(ctx).Err(); err != nil {
// 		fmt.Println("failed to connect to Redis: %w", err)
// 		panic(err)
// 	}
// 	// 将客户端赋值给全局变量
// 	rrdb = client
// 	fmt.Println("Redis connection initialized")
// }

// // 关闭 Redis 连接
// func CloseRedis() {
// 	if rrdb != nil {
// 		_ = rrdb.Close()
// 		fmt.Println("Redis connection closed")
// 	}
// }

// func TestRedis(t *testing.T) {
// 	InitRedisConnection()
// 	ctx := context.Background()

// fmt.Println("关注")
// err := AddRelation_test(ctx, 7093779806818304, 7093627553583104)
// assert.Equal(t, err, nil)
// err = AddRelation_test(ctx, 7093779806818304, 5191404976345088)
// assert.Equal(t, err, nil)
// err = AddRelation_test(ctx, 7093779806818304, 5477693852225536)
// assert.Equal(t, err, nil)
// err = AddRelation_test(ctx, 5897088445452288, 7093779806818304)
// assert.Equal(t, err, nil)
// err = AddRelation_test(ctx, 7093779806818304, 5897088445452288)
// assert.Equal(t, err, nil)

// followIds, err := GetFollowList_test(ctx, 7093627553583104)
// assert.Equal(t, err, nil)
// for _, followid := range followIds {
// 	fmt.Printf("followid: %v\n", followid)
// }

// followIds, err = GetFollowerList_test(ctx, 7093779806818304)
// assert.Equal(t, err, nil)
// for _, followid := range followIds {
// 	fmt.Printf("followid: %v\n", followid)
// }

// followIds, err = GetFriendList_test(ctx, 7093779806818304)
// assert.Equal(t, err, nil)
// for _, followid := range followIds {
// 	fmt.Printf("followid: %v\n", followid)
// }

// fmt.Println("取消关注")
// err := DeleteRelation_test(ctx, 7093779806818304, 7093627553583104)
// assert.Equal(t, err, nil)
// err = DeleteRelation_test(ctx, 7093779806818304, 5191404976345088)
// assert.Equal(t, err, nil)
// err = DeleteRelation_test(ctx, 7093779806818304, 5477693852225536)
// assert.Equal(t, err, nil)
// err = DeleteRelation_test(ctx, 5897088445452288, 7093779806818304)
// assert.Equal(t, err, nil)

// fmt.Println("显示关注列表")
// followIds, err := GetFollowList_test(ctx, 7093779806818304)
// assert.Equal(t, err, nil)
// for _, followid := range followIds {
// 	fmt.Printf("followid: %v\n", followid)
// }

// fmt.Println("取消关注后 再显示关注列表")
// err := DeleteRelation_test(ctx, 7093779806818304, 7093627553583104)
// assert.Equal(t, err, nil)
// followIds, err := GetFollowList_test(ctx, 7093779806818304)
// assert.Equal(t, err, nil)
// for _, followid := range followIds {
// 	fmt.Printf("followid: %v\n", followid)
// }

// 重复添加关系：
// err := AddRelation_test(ctx, 7093779806818304, 7093627553583104)
// if err != nil {
// 	fmt.Printf("err: %v\n", err)
// }
// assert.Equal(t, err, nil)

// // 删除不存在的关系：
// err := DeleteRelation_test(ctx, 00, 123465)
// if err != nil {
// 	fmt.Printf("err: %v\n", err)
// }
// assert.Equal(t, err, nil)

// 测试PutFollowList 给定List添加到redis中
// 	err := PutFollowerList_test(ctx, 7093627553583104, []int64{7093779806818304, 5897088445452288, 5191404976345088, 5477693852225536})
// 	assert.Equal(t, err, nil)
// 	// 读取 Redis 中的关注列表，验证是否正确写入
// 	followIds, err := GetFollowerList_test(ctx, 7093627553583104)
// 	assert.Equal(t, err, nil)
// 	for _, followid := range followIds {
// 		fmt.Printf("followerid: %v\n", followid)
// 	}

// 	// 测试PutFollowerList 给定List添加到redis中 是set 和 score_set的问题

// 	// flag, err := GetIsFollow_test(ctx, 7093779806818304, 7093627553583104)
// 	// assert.Equal(t, err, nil)
// 	// if flag == true {
// 	// 	fmt.Println("关注了")
// 	// } else {
// 	// 	fmt.Println("没关注")
// 	// }

// 	// flag, err = GetIsFollow_test(ctx, 5477693852225536, 7093627553583104)
// 	// assert.Equal(t, err, nil)
// 	// if flag == true {
// 	// 	fmt.Println("关注了")
// 	// } else {
// 	// 	fmt.Println("没关注")
// 	// }

// 	// // 有
// 	// flag, err := IsFollowKeyExist_test(ctx, 5897088445452288)
// 	// assert.Equal(t, err, nil)
// 	// if flag == true {
// 	// 	fmt.Println("有对应id的followkey在redis中")
// 	// } else {
// 	// 	fmt.Println("没有有对应id的followkey在redis中")
// 	// }

// 	// // 无
// 	// flag, err = IsFollowKeyExist_test(ctx, 5477693852225536)
// 	// assert.Equal(t, err, nil)
// 	// if flag == true {
// 	// 	fmt.Println("有对应id的followkey在redis中")
// 	// } else {
// 	// 	fmt.Println("没有有对应id的followkey在redis中")
// 	// }

// 	// // 有
// 	// flag, err = IsFollowerKeyExist_test(ctx, 5477693852225536)
// 	// assert.Equal(t, err, nil)
// 	// if flag == true {
// 	// 	fmt.Println("有对应id的followerkey在redis中")
// 	// } else {
// 	// 	fmt.Println("没有有对应id的followerkey在redis中")
// 	// }

// 	// // 无
// 	// flag, err = IsFollowerKeyExist_test(ctx, 5897088445452288)
// 	// assert.Equal(t, err, nil)
// 	// if flag == true {
// 	// 	fmt.Println("有对应id的followerkey在redis中")
// 	// } else {
// 	// 	fmt.Println("没有有对应id的followerkey在redis中")
// 	// }

// }

// // 在Redis中添加一个关系
// func AddRelation_test(ctx context.Context, userId int64, toUserId int64) error {
// 	fmt.Println("开始添加关系")
// 	// 1. 构建关注关系followKey 和 被关注followerKey，获取当前时间戳
// 	if userId == toUserId {
// 		return errors.New("you can't unfollow yourself")
// 	}
// 	followKey := GetRedisKey(follow, strconv.FormatInt(userId, 10))
// 	followerKey := GetRedisKey(follower, strconv.FormatInt(toUserId, 10))

// 	// 2. 判断redis中 该关注关系是否已存在
// 	_, err := rrdb.ZScore(ctx, followKey, strconv.FormatInt(toUserId, 10)).Result()
// 	if err != nil {
// 		if err == redis.Nil {
// 			fmt.Println("关注关系不存在，可以添加")
// 			// 关注关系不存在，可以添加
// 			// 3. 指定有序集合成员及其分数的结构体
// 			now := time.Now().Unix()
// 			followMember := redis.Z{
// 				Score:  float64(now),
// 				Member: strconv.FormatInt(toUserId, 10),
// 			}
// 			followerMember := redis.Z{
// 				Score:  float64(now),
// 				Member: strconv.FormatInt(userId, 10),
// 			}
// 			// 4. 添加关注关系和粉丝关系
// 			if err := rrdb.ZAdd(ctx, followKey, followMember).Err(); err != nil {
// 				fmt.Println("redis添加关注关系失败", zap.Error(err))
// 				return err
// 			}
// 			if err := rrdb.ZAdd(ctx, followerKey, followerMember).Err(); err != nil {
// 				fmt.Println("redis添加粉丝关系失败", zap.Error(err))
// 				rrdb.ZRem(ctx, followKey, toUserId) // 如果粉丝关系添加失败，则移除关注关系中的对应成员
// 				return err
// 			}
// 			// 5.设置followkey和followerkey的过期时间
// 			_, err = rrdb.Expire(ctx, followKey, customExpiration).Result()
// 			if err != nil {
// 				fmt.Println("redis设置followKey过期时间失败", zap.Error(err))
// 				return err
// 			}
// 			_, err = rrdb.Expire(ctx, followerKey, customExpiration).Result()
// 			if err != nil {
// 				fmt.Println("redis设置followerKey过期时间失败", zap.Error(err))
// 				return err
// 			}

// 			// ***************************** 6. 增加用户的两个计数器：关注数 和 粉丝数 *****************************
// 			// 使用 HINCRBY 命令增加关注者的 follow_count
// 			_, err := rrdb.HIncrBy(ctx, GetRedisKey(KeyUserInfo, strconv.FormatInt(userId, 10)), "follow_count", 1).Result()
// 			if err != nil {
// 				fmt.Println("redis增加关注者的follow_count失败", zap.Error(err))
// 				return err
// 			}

// 			// 使用 HINCRBY 命令增加粉丝的 follower_count
// 			_, err = rrdb.HIncrBy(ctx, GetRedisKey(KeyUserInfo, strconv.FormatInt(toUserId, 10)), "follower_count", 1).Result()
// 			if err != nil {
// 				fmt.Println("redis增加粉丝的follower_count失败", zap.Error(err))
// 				// 将之前增加的关注者的 follow_count 恢复
// 				_, _ = rrdb.HIncrBy(ctx, GetRedisKey(KeyUserInfo, strconv.FormatInt(userId, 10)), "follow_count", -1).Result()
// 				return err
// 			}
// 		} else {
// 			// 发生了其他错误
// 			fmt.Println("发生了错误", zap.Error(err))
// 			return err
// 		}
// 	} else {
// 		// 关注关系已存在
// 		fmt.Println("关注关系已存在，不能添加")
// 		return errors.New("关注关系已存在，不能添加")
// 	}

// 	fmt.Println("添加关系结束，添加成功")
// 	return nil
// }

// // 在Redis中删除一段关系
// func DeleteRelation_test(ctx context.Context, userId int64, toUserId int64) error {
// 	// 1. 构建关注关系followKey 和 被关注followerKey，获取当前时间戳
// 	if userId == toUserId {
// 		return errors.New("you can't unfollow yourself")
// 	}
// 	followKey := GetRedisKey(follow, strconv.FormatInt(userId, 10))
// 	followerKey := GetRedisKey(follower, strconv.FormatInt(toUserId, 10))

// 	// 2. 判断 redis 中关系是否存在
// 	exists, err := rrdb.ZScore(ctx, followKey, strconv.FormatInt(toUserId, 10)).Result()
// 	if err != nil {
// 		if err == redis.Nil {
// 			// 关注关系不存在，不可以删除
// 			fmt.Println("关注关系不存在，不能删除")
// 			return errors.New("关注关系不存在，不能删除")
// 		} else {
// 			// 发生了其他错误
// 			fmt.Println("发生了错误", zap.Error(err))
// 			return err
// 		}
// 	} else {
// 		// 关注关系已存在，可以进行删除操作
// 		fmt.Println("关注关系已存在，可以进行删除操作")
// 		// 3. 从有序集合中删除关注关系和被关注关系
// 		_, err := rrdb.ZRem(ctx, followKey, strconv.FormatInt(toUserId, 10)).Result()
// 		if err != nil {
// 			fmt.Println("redis删除关注关系失败", zap.Error(err))
// 			return err
// 		}

// 		_, err = rrdb.ZRem(ctx, followerKey, strconv.FormatInt(userId, 10)).Result()
// 		if err != nil {
// 			rrdb.ZAdd(ctx, followKey, redis.Z{Score: exists, Member: strconv.FormatInt(toUserId, 10)}) // 如果删除粉丝关系失败，将关注关系还原
// 			fmt.Println("redis删除粉丝关系失败", zap.Error(err))
// 			return err
// 		}

// 		// 4.设置followkey和followerkey的过期时间
// 		_, err = rrdb.Expire(ctx, followKey, customExpiration).Result()
// 		if err != nil {
// 			fmt.Println("redis设置followKey过期时间失败", zap.Error(err))
// 			return err
// 		}
// 		_, err = rrdb.Expire(ctx, followerKey, customExpiration).Result()
// 		if err != nil {
// 			fmt.Println("redis设置followerKey过期时间失败", zap.Error(err))
// 			return err
// 		}

// 		// 5. ***************************** 减少用户的两个计数器：关注数和粉丝数  *****************************
// 		// 使用 HINCRBY 命令减少关注者的 follow_count
// 		_, err = rrdb.HIncrBy(ctx, GetRedisKey(KeyUserInfo, strconv.FormatInt(userId, 10)), "follow_count", -1).Result()
// 		if err != nil {
// 			fmt.Println("redis减少关注者的follow_count失败", zap.Error(err))
// 			return err
// 		}

// 		// 使用 HINCRBY 命令减少粉丝的 follower_count
// 		_, err = rrdb.HIncrBy(ctx, GetRedisKey(KeyUserInfo, strconv.FormatInt(toUserId, 10)), "follower_count", -1).Result()
// 		if err != nil {
// 			fmt.Println("redis减少粉丝的follower_count失败", zap.Error(err))
// 			// 将之前减少的关注者的 follow_count 恢复
// 			_, _ = rrdb.HIncrBy(ctx, GetRedisKey(KeyUserInfo, strconv.FormatInt(userId, 10)), "follow_count", 1).Result()
// 			return err
// 		}
// 	}

// 	return nil
// }

// func GetFollowList_test(ctx context.Context, userId int64) ([]int64, error) {
// 	// 1. 构建 followkey
// 	followKey := GetRedisKey(follow, strconv.FormatInt(userId, 10))

// 	// 2.从Redis中获取有序集合中指定范围的成员
// 	members, err := rrdb.ZRange(ctx, followKey, 0, -1).Result()

// 	if err != nil {
// 		fmt.Println("Redis中获取FollowList失败: ", zap.Error(err))
// 		return nil, err
// 	}

// 	// 创建切片，逐步将关注列表成员members转为int64的关注用户id
// 	followList := make([]int64, len(members))
// 	for i, member := range members {
// 		id, _ := strconv.ParseInt(member, 10, 64)
// 		followList[i] = id
// 	}

// 	// 设置key的过期时间
// 	rrdb.Expire(ctx, followKey, customExpiration)
// 	if err != nil {
// 		return nil, err
// 	}

// 	fmt.Println("Redis缓存中获取followList成功")
// 	return followList, nil
// }

// func GetFollowerList_test(ctx context.Context, userId int64) ([]int64, error) {
// 	fmt.Println("获取粉丝列表：")
// 	followerKey := GetRedisKey(follower, strconv.FormatInt(userId, 10))

// 	// 从Redis中获取有序集合中指定范围的成员
// 	members, err := rrdb.ZRange(ctx, followerKey, 0, -1).Result()
// 	if err != nil {
// 		fmt.Println("Redis中获取FollowerList失败: ", zap.Error(err))
// 		return nil, err
// 	}

// 	// 创建切片，逐步将粉丝列表成员members转为int64的粉丝用户id
// 	followerList := make([]int64, len(members))
// 	for i, member := range members {
// 		id, _ := strconv.ParseInt(member, 10, 64)
// 		followerList[i] = id
// 	}

// 	// 设置key的过期时间
// 	_, err = rrdb.Expire(ctx, followerKey, customExpiration).Result()
// 	if err != nil {
// 		return nil, err
// 	}

// 	fmt.Println("Redis缓存中获取followerList成功")
// 	return followerList, nil
// }

// func GetFriendList_test(ctx context.Context, userId int64) ([]int64, error) {
// 	fmt.Println("获取好友列表")

// 	// 根据userID获取关注列表和粉丝列表
// 	followList, err := GetFollowList_test(ctx, userId)
// 	if err != nil {
// 		fmt.Println("Redis中获取FriendList失败: ", zap.Error(err))
// 		return nil, err
// 	}

// 	followerList, err := GetFollowerList_test(ctx, userId)
// 	if err != nil {
// 		fmt.Println("Redis中获取FriendList失败: ", zap.Error(err))
// 		return nil, err
// 	}

// 	// 找到关注列表和粉丝列表的交集，即好友列表
// 	friendList := make([]int64, 0)
// 	for _, id := range followList {
// 		for _, followerID := range followerList {
// 			if id == followerID {
// 				friendList = append(friendList, id)
// 				break
// 			}
// 		}
// 	}

// 	fmt.Println("Redis缓存中获取FriendList成功")
// 	return friendList, nil
// }

// // 把FollowList 关注用户列表数据写回redis
// func PutFollowList_test(ctx context.Context, userId int64, FollowIdList []int64) error {

// 	// 特判：如果FollowIdList为空，直接返回nil
// 	if FollowIdList == nil || len(FollowIdList) == 0 {
// 		return nil
// 	}

// 	fmt.Println("开始将followlist添加入redis")

// 	// 1. 构建关注列表的键
// 	followKey := GetRedisKey(follow, strconv.FormatInt(userId, 10))

// 	// 2. 使用管道一次性执行多个命令，设置过期时间并将关注列表中的用户ID添加到集合中
// 	pipe := rrdb.Pipeline()

// 	pipe.Expire(ctx, followKey, customExpiration) //设置过期时间

// 	// 3.将关注列表中的用户ID添加到 Redis 集合中
// 	now := float64(time.Now().Unix())
// 	for _, uid := range FollowIdList {
// 		followMember := redis.Z{
// 			Score:  float64(now),
// 			Member: strconv.FormatInt(uid, 10),
// 		}
// 		pipe.ZAdd(ctx, followKey, followMember)
// 	}

// 	// 4. 执行 Redis 管道中的所有命令
// 	_, err := pipe.Exec(ctx)
// 	if err != nil {
// 		fmt.Println("将follorList写回Redis, 管道执行失败", zap.Error(err))
// 		return err
// 	}

// 	fmt.Println("成功将followlist添加入redis")
// 	return nil
// }

// // 把FollowerList 粉丝列表数据写回redis
// func PutFollowerList_test(ctx context.Context, userId int64, FollowerList []int64) error {
// 	// 特判：如果 FollowerList 为空，直接返回 nil
// 	if FollowerList == nil || len(FollowerList) == 0 {
// 		return nil
// 	}

// 	fmt.Println("开始将followerlist添加入redis")
// 	// 1. 构建粉丝列表的键
// 	followerKey := GetRedisKey(follower, strconv.FormatInt(userId, 10))

// 	// 2. 使用管道一次性执行多个命令，设置过期时间并将粉丝列表中的用户ID添加到集合中
// 	pipe := rrdb.Pipeline()

// 	// 设置过期时间
// 	pipe.Expire(ctx, followerKey, customExpiration)

// 	// 将粉丝列表中的用户ID添加到 Redis 集合中
// 	now := float64(time.Now().Unix())
// 	for _, uid := range FollowerList {
// 		followerMember := redis.Z{
// 			Score:  float64(now),
// 			Member: strconv.FormatInt(uid, 10),
// 		}
// 		pipe.ZAdd(ctx, followerKey, followerMember)
// 	}

// 	// 执行 Redis 管道中的所有命令
// 	_, err := pipe.Exec(ctx)
// 	if err != nil {
// 		return err
// 	}

// 	fmt.Println("成功将followerlist添加入redis")

// 	return nil
// }

// // 检查给定的 userId 的关注列表是否存在
// func IsFollowKeyExist_test(ctx context.Context, userId int64) (bool, error) {
// 	fmt.Println("检查给定的 userId 的关注列表是否存在")
// 	// 构建key
// 	followKey := GetRedisKey(follow, strconv.FormatInt(userId, 10))

// 	// Exists() 判断键是否存在于Redis, 返回存在的键的数量
// 	result, err := rrdb.Exists(ctx, followKey).Result()
// 	if err != nil {
// 		fmt.Println("Redis中获取followKey失败: ", zap.Error(err))
// 		return false, err
// 	}

// 	// 设置key的过期时间，并返回 布尔值和错误值
// 	if result == 0 {
// 		_, err = rrdb.Expire(ctx, followKey, customExpiration).Result()
// 		if err != nil {
// 			return false, err
// 		}
// 		return false, nil
// 	}

// 	_, err = rrdb.Expire(ctx, followKey, customExpiration).Result()
// 	if err != nil {
// 		return true, err
// 	}

// 	fmt.Println("Redis缓存中存在userid的followKey")
// 	return true, nil
// }

// // 检查给定的 userId 的粉丝列表是否存在
// func IsFollowerKeyExist_test(ctx context.Context, userId int64) (bool, error) {
// 	fmt.Println("检查给定的 userId 的粉丝列表是否存在")
// 	// 构建key
// 	followerKey := GetRedisKey(follower, strconv.FormatInt(userId, 10))

// 	// Exists() 判断键是否存在于Redis, 返回存在的键的数量
// 	result, err := rrdb.Exists(ctx, followerKey).Result()
// 	if err != nil {
// 		fmt.Println("Redis中获取followerKey失败: ", zap.Error(err))
// 		return false, err
// 	}

// 	// 设置key的过期时间，并返回 布尔值和错误值
// 	if result == 0 {
// 		_, err = rrdb.Expire(ctx, followerKey, customExpiration).Result()
// 		if err != nil {
// 			return false, err
// 		}
// 		return false, nil
// 	}

// 	_, err = rrdb.Expire(ctx, followerKey, customExpiration).Result()
// 	if err != nil {
// 		return true, err
// 	}

// 	fmt.Println("Redis缓存中存在userid的followerKey")
// 	return true, nil
// }

// func GetIsFollow_test(ctx context.Context, userId int64, toUserId int64) (bool, error) {
// 	// 构建key
// 	followKey := GetRedisKey(follow, strconv.FormatInt(userId, 10))

// 	//获取有序集合中指定成员的分值
// 	score, err := rrdb.ZScore(ctx, followKey, strconv.FormatInt(toUserId, 10)).Result()
// 	if err != nil && err != redis.Nil {
// 		fmt.Println("Redis中获取对应userId是否toUserId关注失败: ", zap.Error(err))
// 		return false, err
// 	}

// 	// 设置key的过期时间，并返回 布尔值和错误值
// 	if score == 0 {
// 		_, err = rrdb.Expire(ctx, followKey, customExpiration).Result()
// 		if err != nil {
// 			return false, err
// 		}
// 		return false, nil // Redis中没有该关系 且没有错误
// 	}

// 	_, err = rrdb.Expire(ctx, followKey, customExpiration).Result()
// 	if err != nil {
// 		return true, err // Redis中存在该关系，但出错
// 	}

// 	return true, nil // Redis中存在该关系 且没有错误
// }
