package models

type User struct {
	ID            int64  `json:"id"`
	Name          string `json:"name"`
	FollowCount   int64  `json:"follow_count"`   // 关注总数
	FollowerCount int64  `json:"follower_count"` // 粉丝总数
	WorkCount     int64  `json:"work_count"`
	FavoriteCount int64  `json:"favorite_count"`
	IsFollow      bool   `json:"is_follow"` // true-已关注，false-未关注
}

// 获取用户信息
type GetUserInfoReq struct {
	UserId int64 `form:"user_id" json:"user_id" binding:"required"`
}

type GetUserInfoResp struct {
	User *User `json:"user,omitempty"`
}

// 用户注册
type AccountReq struct {
	Username string `form:"username" json:"username" binding:"required,max=32" msg:"最长32个字符串"`
	Password string `form:"password" json:"password" binding:"required,max=32" msg:"最长32个字符串"`
}

type AccountResp struct {
	UserID int64  `json:"user_id"`
	Token  string `json:"token"` // 用户鉴权token
}
