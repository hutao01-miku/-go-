package model

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// 数据库操放这里
var Conn *gorm.DB

func NewMysql() {
	my := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&Local", "root", "1234", "127.0.0.1:3306", "vote")

	conn, err := gorm.Open(mysql.Open(my), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	Conn = conn
}
func Close() {
	db, _ := Conn.DB()
	err := db.Close()
	if err != nil {
		return
	}
}
