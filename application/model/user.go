package model

import (
	"fmt"
)

type User struct {
	Name     string `json:"name" form:"name"`
	Password string `json:"password" form:"password"`
}

func GetUser(user *User) map[string]any {

	ret := make(map[string]any)
	if err := Conn.Table("user").Where("name = ?", user.Name).Find(&ret).Error; err != nil {
		fmt.Printf("err:%s", err.Error())
	}
	return ret
}
