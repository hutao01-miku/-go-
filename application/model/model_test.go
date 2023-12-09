package model

import (
	"fmt"
	"testing"
	"time"
)

func TestGetVotes(t *testing.T) { //该方法测试vote是否生效
	NewMysql() //连接数据库
	//测试用例
	r := GetVotes() //查询方法
	fmt.Printf("ret:%+v", r)
	//%+v 是一个格式化占位符，表示以“扩展”格式输出变量的值。对于结构体（struct）类型的变量，%+v 会输出字段名和字段值。
	//和%v区别：通用的格式化占位符，根据变量的实际类型进行格式化输出。对于结构体，它会输出字段的值，但不会包括字段名。
	Close()
}
func TestGetVote(t *testing.T) {
	NewMysql() //连接数据库
	//测试用例
	r := GetVote(1) //查询方法

	fmt.Printf("ret:%+v", r)
	//%+v 是一个格式化占位符，表示以“扩展”格式输出变量的值。对于结构体（struct）类型的变量，%+v 会输出字段名和字段值。
	//和%v区别：通用的格式化占位符，根据变量的实际类型进行格式化输出。对于结构体，它会输出字段的值，但不会包括字段名。
	Close()
}
func TestDoVote(t *testing.T) {
	NewMysql() //连接数据库
	//测试用例
	r := DoVote(1, 1, []int64{1, 2}) //查询方法
	fmt.Printf("ret:%+v", r)
	Close()
}
func TestAddVote(t *testing.T) {
	NewMysql() //连接数据库
	//测试用例
	vote := Vote{
		Title:       "测试用例",
		Type:        0,
		Status:      0,
		Time:        0,
		UserId:      0,
		CreatedTime: time.Now(),
		UpdateTime:  time.Now(),
	}
	opt := make([]VoteOpt, 0)
	opt = append(opt, VoteOpt{
		Name:        "测试选项1",
		Count:       0,
		CreatedTime: time.Now(),
		UpdateTime:  time.Now(),
	})
	opt = append(opt, VoteOpt{
		Name:        "测试选项2",
		Count:       0,
		CreatedTime: time.Now(),
		UpdateTime:  time.Now(),
	})

	r := AddVote(vote, opt) //查询方法
	fmt.Printf("ret:%+v", r)
	Close()
}

func TestGetUserV1(t *testing.T) {
	NewMysql()
	a := GetUserV1("admin")
	fmt.Printf("a:%v", a)
}
