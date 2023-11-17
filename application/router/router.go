package router

import (
	"awesomeProject1/application/logic"
	"github.com/gin-gonic/gin"
)

func New() {
	r := gin.Default()
	r.LoadHTMLGlob("application/view/*")
	//var user model.User"/login", logic.GetLogin)
	r.GET("/login", logic.GetLogin)
	r.POST("/login", logic.DoLogin)

	//	g.POST("/login", func(context *gin.Context) {
	//		var user User
	//		ret := make(map[string]string)
	//		_ = context.ShouldBind(&user)
	//		err := conn.Table("user").Where("name = ? and password = ?", user.Name, user.Password).Find(&user)
	//		if err != nil {
	//			context.JSON(http.StatusBadGateway, map[string]string{
	//				"msg": "用户名或密码错误",
	//			})
	//相关路径放这里
	if err := r.Run(":8080"); err != nil {
		panic("启动失败")
	}
}
