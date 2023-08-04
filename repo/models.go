package repo

import (
	"gorm.io/gorm"
	"time"
)

// 数据库模型

// Comment 评论表
type Comment struct {
	gorm.Model
	CommentUuid uint64 `gorm:"column:comment_uuid;type:bigint(20) unsigned;default:0;comment:评论uuid;NOT NULL;unique" json:"comment_uuid"`
	UserId      uint64 `gorm:"column:user_id;type:bigint(20) unsigned;default:0;comment:评论作者id;NOT NULL" json:"user_id"`
	VideoId     uint64 `gorm:"column:video_id;type:bigint(20) unsigned;default:0;comment:评论视频id;NOT NULL" json:"video_id"`
	Contents    string `gorm:"column:contents;type:varchar(255);comment:评论内容;NOT NULL" json:"contents"`
	CreateTime  uint64 `gorm:"column:create_time;type:bigint(20) unsigned;default:0;comment:自设创建时间(unix);NOT NULL" json:"create_time"`
}

func (m *Comment) TableName() string {
	return "comment"
}

// Favorite 点赞表
type Favorite struct {
	ID        uint   `gorm:"primarykey"`
	UserId    uint64 `gorm:"column:user_id;type:bigint(20) unsigned;default:0;comment:用户id;NOT NULL;uniqueIndex:f" json:"user_id"`
	VideoId   uint64 `gorm:"column:video_id;type:bigint(20) unsigned;default:0;comment:视频id;NOT NULL;uniqueIndex:f" json:"video_id"`
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
	MessageUuid uint64         `gorm:"column:message_uuid;type:bigint(20) unsigned;default:0;comment:消息uuid;NOT NULL" json:"message_uuid"`
	ToUserId    uint64         `gorm:"column:to_user_id;type:bigint(20) unsigned;default:0;comment:该消息接收者的id;NOT NULL;uniqueIndex:m" json:"to_user_id"`
	FromUserId  uint64         `gorm:"column:from_user_id;type:bigint(20) unsigned;default:0;comment:该消息发送者的id;NOT NULL;uniqueIndex:m" json:"from_user_id"`
	Contents    string         `gorm:"column:contents;type:varchar(255);comment:消息内容;NOT NULL" json:"contents"`
	CreateTime  uint64         `gorm:"column:create_time;type:bigint(20) unsigned;default:0;comment:自设创建时间(unix);NOT NULL" json:"create_time"`
}

func (m *Message) TableName() string {
	return "message"
}

func (m *Relation) TableName() string {
	return "relation"
}

type Relation struct {
	ID        uint   `gorm:"primarykey"`
	UserId    uint64 `gorm:"column:user_id;type:bigint(20) unsigned;default:0;comment:用户id;NOT NULL;uniqueIndex:r" json:"user_id"`
	ToUserId  uint64 `gorm:"column:to_user_id;type:bigint(20) unsigned;default:0;comment:关注目标的用户id;NOT NULL;uniqueIndex:r" json:"to_user_id"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"uniqueIndex:r"`
}

// User 用户表
type User struct {
	gorm.Model
	UserId        uint64 `gorm:"column:userid;type:bigint(20);comment:用户id;NOT NULL;unique" json:"user_id"` // UserId 是Relation表中，UserId和ToUserId的外键
	Comment       []Comment
	Favorite      []Favorite
	Video         []Video
	Relation      []Relation `gorm:"foreignKey:ToUserId"`
	Relation1     []Relation `gorm:"foreignKey:UserId"`
	Message       []Message  `gorm:"foreignKey:ToUserId"`
	Message1      []Message  `gorm:"foreignKey:FromUserId"`
	Name          string     `gorm:"column:name;type:varchar(32);comment:用户名称;NOT NULL;unique" json:"name"`
	Password      string     `gorm:"column:password;type:varchar(255);comment:密码，已加密;NOT NULL" json:"password"`
	FollowCount   uint64     `gorm:"column:follow_count;type:bigint(20) unsigned;default:0;comment:关注人数;NOT NULL" json:"follow_count"`
	FollowerCount uint64     `gorm:"column:follower_count;type:bigint(20) unsigned;default:0;comment:粉丝人数;NOT NULL" json:"follower_count"`
	WorkCount     uint64     `gorm:"column:work_count;type:bigint(20) unsigned;default:0;comment:作品数;NOT NULL" json:"work_count"`
	FavoriteCount uint64     `gorm:"column:favorite_count;type:bigint(20) unsigned;default:0;comment:点赞视频数;NOT NULL" json:"favorite_count"`
}

func (m *User) TableName() string {
	return "user"
}

// Video 视频表
type Video struct {
	gorm.Model
	UserId        uint64 `gorm:"column:user_id;type:bigint(20) unsigned;default:0;comment:user表主键;NOT NULL" json:"user_id"`
	Title         string `gorm:"column:title;type:varchar(128);comment:视频标题;NOT NULL" json:"title"`
	PlayUrl       string `gorm:"column:play_url;type:varchar(128);comment:视频地址;NOT NULL" json:"play_url"`
	CoverUrl      string `gorm:"column:cover_url;type:varchar(128);comment:封面地址;NOT NULL" json:"cover_url"`
	FavoriteCount uint   `gorm:"column:favorite_count;type:int(15) unsigned;default:0;comment:获赞数量;NOT NULL" json:"favorite_count"`
	CommentCount  uint   `gorm:"column:comment_count;type:int(15) unsigned;default:0;comment:评论数量;NOT NULL" json:"comment_count"`
	Comment       []Comment
	Favorite      []Favorite
}

func (m *Video) TableName() string {
	return "video"
}
