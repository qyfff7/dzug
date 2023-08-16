package redis

import (
	"context"
	"strconv"
)

func AddComm(ctx context.Context, videoId int64, comment_uuid int64) int {
	err := Rdb.SAdd(ctx, strconv.FormatInt(videoId, 10), strconv.FormatInt(comment_uuid, 10)).Err()
	if err != nil {
		return 0
	}

	return 1
}

func DelComm(ctx context.Context, videoId int64, comment_uuid int64) int {
	_, err := Rdb.Exists(context.Background(), strconv.FormatInt(comment_uuid, 10)).Result()
	if err != nil {
		return 0
	}

	ans := Rdb.SRem(context.Background(), strconv.FormatInt(videoId, 10), comment_uuid) // 这个key现在已经存在了，去删除这个videoId
	if ans.Val() == 0 {                                                                 // 已经存在这个value了
		return 2
	}
	return 1
}

func GetComm(ctx context.Context, videoId int64) ([]int64, error) {

	_, err := Rdb.Exists(context.Background(), strconv.FormatInt(videoId, 10)).Result()
	if err != nil {
		return nil, err
	}
	cmd := Rdb.SMembers(context.Background(), strconv.FormatInt(videoId, 10))
	commentIDs := make([]int64, len(cmd.Val()))
	for k, v := range cmd.Val() {
		value, _ := strconv.Atoi(v)
		commentIDs[k] = int64(value)
	}

	return commentIDs, err
}
