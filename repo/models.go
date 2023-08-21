package repo

import (
	"gorm.io/gorm"
	"time"
)

// 数据库模型

// Comment 评论表
type Comment struct {
	ID          uint   `gorm:"primarykey"`
	CommentUuid int64  `gorm:"column:comment_uuid;type:bigint(20) unsigned;default:0;comment:评论uuid;NOT NULL;unique" json:"comment_uuid"`
	UserId      int64  `gorm:"column:user_id;type:bigint(20) unsigned;default:0;comment:评论作者id;NOT NULL" json:"user_id"`
	VideoId     int64  `gorm:"column:video_id;type:bigint(20) unsigned;default:0;comment:评论视频id;NOT NULL" json:"video_id"`
	Contents    string `gorm:"column:contents;type:varchar(255);comment:评论内容;NOT NULL" json:"contents"`
	CreateTime  int64  `gorm:"column:create_time;type:bigint(20) unsigned;default:0;comment:自设创建时间(unix);NOT NULL" json:"create_time"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"uniqueIndex:f"`
}

func (m *Comment) TableName() string {
	return "comment"
}

// Favorite 点赞表
type Favorite struct {
	ID        uint  `gorm:"primarykey"`
	UserId    int64 `gorm:"column:user_id;type:bigint(20) unsigned;default:0;comment:用户id;NOT NULL;uniqueIndex:f" json:"user_id"`
	VideoId   int64 `gorm:"column:video_id;type:bigint(20) unsigned;default:0;comment:视频id;NOT NULL;uniqueIndex:f" json:"video_id"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"uniqueIndex:f"`
}

func (m *Favorite) TableName() string {
	return "favorite"
}

// Message 消息表
type Message struct {
	ID          uint `gorm:"primarykey"`
	CreatedAt   time.Time
	UpdatedAt   time.Time      `gorm:"uniqueIndex:m"`
	DeletedAt   gorm.DeletedAt `gorm:"index"`
	MessageUuid int64          `gorm:"column:message_uuid;type:bigint(20) unsigned;default:0;comment:消息uuid;NOT NULL;" json:"message_uuid"`
	ToUserId    int64          `gorm:"column:to_user_id;type:bigint(20) unsigned;default:0;comment:该消息接收者的id;NOT NULL;uniqueIndex:m" json:"to_user_id"`
	FromUserId  int64          `gorm:"column:from_user_id;type:bigint(20) unsigned;default:0;comment:该消息发送者的id;NOT NULL;uniqueIndex:m" json:"from_user_id"`
	Contents    string         `gorm:"column:contents;type:varchar(255);comment:消息内容;NOT NULL" json:"contents"`
	CreateTime  int64          `gorm:"column:create_time;type:bigint(20) unsigned;default:0;comment:自设创建时间(unix);NOT NULL" json:"create_time"`
}

func (m *Message) TableName() string {
	return "message"
}

type Relation struct {
	ID        uint           `gorm:"primarykey"`
	UserId    int64          `gorm:"foreignKey:UserId;references:UserId;comment:用户的UserId;type:bigint(20) unsigned;default:0;NOT NULL;uniqueIndex:r" json:"user_id"`
	ToUserId  int64          `gorm:"foreignKey:ToUserId;references:UserId;comment:关注目标用户的UserId;type:bigint(20) unsigned;default:0;NOT NULL;uniqueIndex:r" json:"to_user_id"`
	CreatedAt time.Time      `gorm:"null" json:"created_at"`
	UpdatedAt time.Time      `gorm:"null" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"uniqueIndex:r"`
}

func (m *Relation) TableName() string {
	return "relation"
}

// User 用户表
type User struct {
	gorm.Model
	UserId           int64      `gorm:"column:user_id;type:bigint(20);comment:用户真正的id;NOT NULL;unique" json:"user_id"` // UserId 是Relation表中，UserId和ToUserId的外键
	Comment          []Comment  `gorm:"foreignKey:UserId;references:UserId"`
	Favorite         []Favorite `gorm:"foreignKey:UserId;references:UserId"`
	Video            []Video    `gorm:"foreignKey:UserId;references:UserId"`
	Relation         []Relation `gorm:"foreignKey:ToUserId;references:UserId"`
	Relation1        []Relation `gorm:"foreignKey:UserId;references:UserId"`
	Message          []Message  `gorm:"foreignKey:ToUserId;references:UserId"`
	Message1         []Message  `gorm:"foreignKey:FromUserId;references:UserId"`
	Name             string     `gorm:"column:name;type:varchar(32);comment:用户名称;NOT NULL;unique" json:"name"`
	BackgroundImages string     `gorm:"column:background_images;type:varchar(255);comment:主页背景图;default:" json:"background_images"`
	Avatar           string     `gorm:"column:avatar;type:varchar(255);comment:用户头像;default:" json:"avatar"`
	Signature        string     `gorm:"column:signature;type:varchar(255);comment:个人简介;default:" json:"signature"`
	Password         string     `gorm:"column:password;type:varchar(255);comment:密码，已加密;NOT NULL" json:"password"`
	FollowCount      int64      `gorm:"column:follow_count;type:bigint(20) unsigned;default:0;comment:关注人数;NOT NULL" json:"follow_count"`
	FollowerCount    int64      `gorm:"column:follower_count;type:bigint(20) unsigned;default:0;comment:粉丝人数;NOT NULL" json:"follower_count"`
	WorkCount        int64      `gorm:"column:work_count;type:bigint(20) unsigned;default:0;comment:作品数;NOT NULL" json:"work_count"`
	FavoriteCount    int64      `gorm:"column:favorite_count;type:bigint(20) unsigned;default:0;comment:点赞视频数;NOT NULL" json:"favorite_count"`
	TotalFavorited   int64      `gorm:"column:total_favorited;type:bigint(20);default:0;comment:获赞数"`
}

func (m *User) TableName() string {
	return "user"
}

// Video 视频表
type Video struct {
	gorm.Model
	UserId        int64  `gorm:"column:user_id;type:bigint(20) unsigned;default:0;comment:user表主键;NOT NULL;" json:"user_id"`
	Title         string `gorm:"column:title;type:varchar(128);comment:视频标题;NOT NULL" json:"title"`
	PlayUrl       string `gorm:"column:play_url;type:varchar(225);comment:视频地址;NOT NULL" json:"play_url"`
	CoverUrl      string `gorm:"column:cover_url;type:varchar(225);comment:封面地址;NOT NULL" json:"cover_url"`
	FavoriteCount uint   `gorm:"column:favorite_count;type:int(15) unsigned;default:0;comment:获赞数量;NOT NULL" json:"favorite_count"`
	CommentCount  uint   `gorm:"column:comment_count;type:int(15) unsigned;default:0;comment:评论数量;NOT NULL" json:"comment_count"`
	Comment       []Comment
	Favorite      []Favorite
}

func (m *Video) TableName() string {
	return "video"
}
