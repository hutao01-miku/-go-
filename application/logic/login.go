package logic

import (
	"awesomeProject1/application/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetLogin(context *gin.Context) {
	context.HTML(http.StatusOK, "login.tmpl", nil)
}
func DoLogin(context *gin.Context) {
	var user model.User
	ret := make(map[string]any)
	_ = context.ShouldBind(&user)

	//ret = model.GetUser(&user)
	//if err := model.Conn.Table("user").Where("name = ?", user.Name).Find(&ret).Error; err != nil {
	//	fmt.Printf("err:%s", err.Error())
	//	context.JSON(http.StatusBadGateway, map[string]string{
	//		"msg": "用户名或密码错误",
	//	})
	//}
	context.JSON(http.StatusOK, ret)
}
