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

func GetFavorById(userId uint64) ([]uint64, error) {
	var favors []repo.Favorite
	res := repo.DB.Where("user_id = ?", userId).Find(&favors)
	if res.Error != nil {
		return nil, res.Error
	}
	ans := make([]uint64, len(favors))
	for k, v := range favors {
		ans[k] = v.VideoId
	}
	return ans, nil
}
