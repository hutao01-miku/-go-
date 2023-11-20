package model

import "fmt"

func GetVotes() []Vote { //输出Vote类型的值
	//封装查询方法,查询投票项目的详情
	ret := make([]Vote, 0) //切片类型，0表示长度为0
	//ret := make(map[string]any)
	err := Conn.Table("vote").Find(&ret).Error
	if err != nil {
		fmt.Printf("err:%s", err.Error())
	}
	return ret
}
