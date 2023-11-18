package logic

//逻辑层，相当于controller
import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"toupiao/application/model"
)

func GetLogin(context *gin.Context) {
	context.HTML(http.StatusOK, "login.tmpl", nil)
}
func DoLogin(context *gin.Context) {
	var user model.User
	ret := make(map[string]any)
	_ = context.ShouldBind(&user)
	ret = model.GetUser(&user)
	err := model.Conn.Table("user").Where("name = ?", user.Name).Find(&ret).Error
	if err != nil {
		fmt.Printf("err:%s", err.Error())
		context.JSON(http.StatusBadGateway, map[string]string{
			"msg": "用户名或密码错误",
		})
	}
	context.JSON(http.StatusOK, ret)
}
