package models

type Video struct {
	Id            int64  `json:"id,omitempty"`
	Auther        User   `json:"user_id,omitempty"`
	PlayUrl       string `json:"play_url,omitempty"`
	CoverUrl      string `json:"cover_url,omitempty"`
	FavoriteCount int64  `json:"favorite_count,omitempty"`
	CommentCount  int64  `json:"comment_count,omitempty"`
	IsFavorite    bool   `json:"is_favorite,omitempty"`
	Title         string `json:"title,omitempty"`
}

type GetVideoListByTimeResp struct {
	Response
	NextTime  int64 `json:"nextTime,omitempty"`
	videoList Video
}

// Feed
type Feed struct {
	Response
	NextTime int64 `json:"nextTime,omitempty"`
	*Video
	*User `json:"user"`
}