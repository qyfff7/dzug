package dao

import (
	"dzug/repo"
)

func Favorite(videoId, userId uint64) error {
	favorite := repo.Favorite{
		UserId:  videoId,
		VideoId: userId,
	}
	res := repo.DB.Create(&favorite) // 记得加上地址
	return res.Error
}
