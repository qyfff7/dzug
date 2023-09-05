package dao

import (
	"context"
	"dzug/protos/comment"
	"dzug/repo"
	"fmt"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

func Comm(ctx context.Context, comment_uuid int64, user_id int64, video_id int64, contents string) (int, uint) {

	comment := &repo.Comment{
		CommentUuid: comment_uuid,
		UserId:      user_id,
		VideoId:     video_id,
		Contents:    contents,
	}
	txn := repo.DB.Begin()
	if err := txn.Create(comment).Error; err != nil {
		repo.DB.Rollback()
		zap.L().Error(err.Error())
		return 0, 0 //上传数据出错
	}
	id := comment.ID
	err := txn.Table("video").Where("id = ?", comment.VideoId).Update("comment_count", gorm.Expr("comment_count + ?", 1)).Error
	if err != nil {
		zap.L().Error(err.Error())
		return 3, id //评论数添加失败
	}

	if err := txn.Commit().Error; err != nil {
		zap.L().Error(err.Error())
		return 2, id //提交出错
	}
	return 1, id
}

// 删除评论
func Incomm(ctx context.Context, videoId int64, comment_uuid int64) int {
	comm := repo.Comment{}

	res := repo.DB.Where("id = ?", comment_uuid).Delete(&comm)

	if res.Error != nil {
		return 0
	}
	err := repo.DB.Table("video").Where("id = ?", videoId).Update("comment_count", gorm.Expr("comment_count - ?", 1)).Error
	if err != nil {
		zap.L().Error(err.Error())
		return 3 //评论数减少失败
	}
	return 1
}

func GetcommByCommentIDs(commentIDs []int64) ([]*comment.Comment, error) {
	comm := repo.Comment{}
	commentList := make([]*comment.Comment, len(commentIDs))
	for k, v := range commentIDs {
		comm.ID = uint(v)
		res := repo.DB.Where("id = ?", v).First(&comm)
		if res.Error != nil {
			zap.L().Error("读取评论失败")
			return nil, res.Error
		}
		commentList[k] = new(comment.Comment) // 初始化
		commentList[k].Id = int64(comm.ID)
		commentList[k].User = &comment.User{}
		res = repo.DB.Where("user_id = ?", comm.UserId).First(&commentList[k].User)
		if res.Error != nil {
			zap.L().Error("获取用户信息失败")
			return nil, res.Error
		}
		commentList[k].User.Id = comm.UserId // 查询评论用户信息

		commentList[k].Content = comm.Contents // 评论内容
		commentList[k].CreateDate = comm.CreatedAt.Format("2006/01/02 15:04:05.000")
		fmt.Println(commentList[k])
	}
	return commentList, nil

}

func GetcommByVideoId(videoId int64) ([]int64, error) {
	var comm []repo.Comment
	res := repo.DB.Where("video_id = ?", videoId).Find(&comm) // 查询该用户所有点赞视频数据
	if res.Error != nil {
		return nil, res.Error
	}
	ans := make([]int64, len(comm))
	for k, v := range comm { // 写入返回的数组
		ans[k] = int64(v.ID)
	}
	return ans, nil
}
