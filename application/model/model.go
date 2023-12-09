package model

import (
	"time"
)

//这里存放各种gorm建表语句

type VoteOptUser struct {
	Id         int64     `gorm:"column:id;primary_key;AUTO_INCREMENT"`
	UserId     int64     `gorm:"column:user_id;NOT NULL"`
	VoteId     int64     `gorm:"column:vote_id;NOT NULL"`
	VoteOptId  int64     `gorm:"column:vote_opt_id;NOT NULL"`
	CreateTime time.Time `gorm:"column:create_time;NOT NULL"`
	UpdateTime time.Time `gorm:"column:update_time;NOT NULL"`
}

func (v *VoteOptUser) TableName() string {
	return "vote_opt_user"
}

type VoteOpt struct {
	Id          int64     `gorm:"column:id;primary_key;AUTO_INCREMENT"`
	Name        string    `gorm:"column:name;NOT NULL"`
	VoteId      int64     `gorm:"column:vote_id;NOT NULL"`
	Count       int64     `gorm:"column:count;NOT NULL"`
	CreatedTime time.Time `gorm:"column:created_time;NOT NULL"`
	UpdateTime  time.Time `gorm:"column:update_time;NOT NULL"`
}

func (v *VoteOpt) TableName() string {
	return "vote_opt"
}

type Vote struct {
	Id          int64     `gorm:"column:id;primary_key;AUTO_INCREMENT"`
	Title       string    `gorm:"column:title;NOT NULL"`
	Type        int32     `gorm:"column:type;NOT NULL;comment:'0是单选1是多选'"`
	Status      int32     `gorm:"column:status;NOT NULL;comment:'0开放1超时'"`
	Time        int64     `gorm:"column:time;NOT NULL;comment:'有效时长'"`
	UserId      int64     `gorm:"column:user_id;NOT NULL;comment:'创建人是谁'"`
	CreatedTime time.Time `gorm:"column:created_time;NOT NULL;comment:'创建时间'"`
	UpdateTime  time.Time `gorm:"column:updated_time;NOT NULL;comment:'更新时间'"`
}

func (v *Vote) TableName() string {
	return "vote"
}

type User struct {
	Id          int64     `gorm:"column:id;primary_key;AUTO_INCREMENT"`
	Name        string    `gorm:"column:name;NOT NULL"`
	Password    string    `gorm:"column:password;NOT NULL"`
	CreatedTime time.Time `gorm:"column:created_time;NOT NULL"`
	UpdateTime  time.Time `gorm:"column:update_time;NOT NULL"`
	Uuid        string    `gorm:"column:uuid;NOT NULL"`
}

func (u *User) TableName() string {
	return "user"
}

type VoteWithOpt struct {
	Vote Vote
	Opt  []VoteOpt
}
