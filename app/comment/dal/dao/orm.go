package dao

import (
	"context"
	"dzug/protos/comment"
	"dzug/repo"
	"fmt"

	"strconv"

	"go.uber.org/zap"
)

func Comm(ctx context.Context, comment_uuid int64, user_id int64, video_id int64, contents string) int {
	comment := &repo.Comment{
		CommentUuid: comment_uuid,
		UserId:      user_id,
		VideoId:     video_id,
		Contents:    contents,
	}

	if err := repo.DB.Create(comment).Error; err != nil {
		repo.DB.Rollback()
		zap.L().Error(err.Error())
		return 0 //上传数据出错
	}

	if err := repo.DB.Commit().Error; err != nil {
		zap.L().Error(err.Error())
		return 2 //提交出错
	}
	return 1
}

// 删除评论
func Incomm(comment_uuid int64) int {
	var comm []repo.Comment

	res := repo.DB.Where("comment_uuid = ?", comment_uuid).Delete(&comm)
	if res.Error != nil {
		return 0
	}
	return 1
}

func GetcommByVideoId(commentIDs []int64) ([]*comment.Comment, error) {
	comm := repo.Comment{}
	commentList := make([]*comment.Comment, len(commentIDs))
	for k, v := range commentIDs {
		comm.ID = uint(v)
		res := repo.DB.First(&comm)
		if res.Error != nil {
			zap.L().Error("读取评论失败")
			return nil, res.Error
		}
		commentList[k] = new(comment.Comment) // 初始化
		commentList[k].Id = commentIDs[k]
		commentList[k].User = &comment.User{}
		res = repo.DB.Where("user_id = ?", comm.UserId).First(&commentList[k].User)
		commentList[k].User.Id = comm.UserId // 查询并设置视频作者信息

		commentList[k].Content = comm.Contents // 设置视频基本数据
		commentList[k].CreateDate = strconv.FormatInt(comm.CreateTime, 10)
		fmt.Println(commentList[k])
	}
	return commentList, nil

}
