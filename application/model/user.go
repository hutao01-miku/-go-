package model

import (
	"fmt"
	"toupiao/application/tools"
)

func GetUser(name string) *User {
	//封装查询方法
	var ret User
	//ret := make(map[string]any)
	err := Conn.Table("user").Where("name = ?", name).Find(&ret).Error
	if err != nil {
		fmt.Printf("err:%s", err.Error())
	}
	return &ret
}
func GetUserV1(name string) *User { //用原声sql来写
	var ret User
	err := Conn.Raw("select * from user where name = ?", name).Find(&ret).Error
	if err != nil {
		tools.Logger.Error("[GetUserV1]err:%s", err.Error())
	}
	return &ret
}

func CreateUser(user *User) error {
	err := Conn.Create(user).Error
	if err != nil {
		fmt.Printf("err:%s", err.Error())
		return err
	}
	return nil
}
