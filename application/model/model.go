package model

import "database/sql"

//这里存放各种gorm建表语句

type VoteOptUser struct {
	Id         sql.NullInt64  `gorm:"column:id;primary_key;AUTO_INCREMENT"`
	UserId     sql.NullInt64  `gorm:"column:user_id"`
	VoteId     sql.NullInt64  `gorm:"column:vote_id"`
	VoteOptId  sql.NullInt64  `gorm:"column:vote_opt_id"`
	CreateTime sql.NullString `gorm:"column:create_time"`
	UpdateTime sql.NullString `gorm:"column:update_time"`
}

func (v *VoteOptUser) TableName() string {
	return "vote_opt_user"
}

type VoteOpt struct {
	Id          sql.NullInt64  `gorm:"column:id;primary_key;AUTO_INCREMENT"`
	Name        sql.NullString `gorm:"column:name"`
	VoteId      sql.NullInt64  `gorm:"column:vote_id"`
	Count       sql.NullInt32  `gorm:"column:count"`
	CreatedTime sql.NullString `gorm:"column:created_time"`
	UpdateTime  sql.NullString `gorm:"column:update_time"`
}

func (v *VoteOpt) TableName() string {
	return "vote_opt"
}

type Vote struct {
	Id          sql.NullInt64  `gorm:"column:id;primary_key;AUTO_INCREMENT"`
	Title       sql.NullString `gorm:"column:title"`
	Type        sql.NullInt32  `gorm:"column:type;comment:'0是单选1是多选'"`
	Status      sql.NullInt32  `gorm:"column:status;comment:'0开放1超时'"`
	Time        sql.NullInt64  `gorm:"column:time;comment:'有效时长'"`
	UserId      sql.NullInt64  `gorm:"column:user_id;comment:'创建人是谁'"`
	CreatedTime sql.NullString `gorm:"column:created_time;comment:'创建时间'"`
	UpdatedTime sql.NullString `gorm:"column:updated_time;comment:'更新时间'"`
}

func (v *Vote) TableName() string {
	return "vote"
}

type User struct {
	Id          sql.NullInt64  `gorm:"column:id;primary_key;AUTO_INCREMENT"`
	Name        sql.NullString `gorm:"column:name"`
	Password    sql.NullString `gorm:"column:password"`
	CreatedTime sql.NullString `gorm:"column:created_time"`
	UpdateTime  sql.NullString `gorm:"column:update_time"`
}

func (u *User) TableName() string {
	return "user"
}
