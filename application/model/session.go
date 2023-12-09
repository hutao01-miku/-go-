package model

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/sessions" //需要导入session包
)

var store = sessions.NewCookieStore([]byte("总辉编程啦啦啦"))

// session的id需要存放在cookie中
var sessionName = "session-name"

// GetSession 从session中获取一个值
func GetSession(c *gin.Context) map[interface{}]interface{} {
	//具体操作： 通过 store.Get 方法从当前请求中获取会话，然后打印会话的值并返回。
	session, _ := store.Get(c.Request, sessionName)
	fmt.Printf("session:%+v\n", session.Values)
	return session.Values
}

// SetSession 往session中设置一个值
func SetSession(c *gin.Context, name string, id int64) error {
	session, _ := store.Get(c.Request, sessionName)
	//设置会话中的 "name" 和 "id" 字段
	session.Values["name"] = name
	session.Values["id"] = id
	//并通过 session.Save 方法将更新后的会话保存到请求中，以更新客户端的Cookie。
	return session.Save(c.Request, c.Writer)
}

// FlushSession 清除session中设置的值
func FlushSession(c *gin.Context) error {
	session, _ := store.Get(c.Request, sessionName)
	//通过 store.Get 方法获取当前请求的会话,想操作肯定得先获取到
	fmt.Printf("session : %+v\n", session.Values)
	session.Values["name"] = ""
	session.Values["id"] = int64(0) //id一直会是int64的值
	//将字符段设为空
	return session.Save(c.Request, c.Writer)
	//通过 session.Save 方法将更新后的会话保存到请求中，以清空客户端的Cookie。
}
