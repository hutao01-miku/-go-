package model

import (
	"fmt"
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
