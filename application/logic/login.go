package logic

//逻辑层，相当于controller
import (
	"github.com/gin-gonic/gin"
	"net/http"
	"toupiao/application/model"
	"toupiao/application/tools"
)

type User struct {
	Name     string `json:"name" form:"name"`
	Password string `json:"password" form:"password"`
}

func GetLogin(context *gin.Context) {
	context.HTML(http.StatusOK, "login.tmpl", nil)
}
func DoLogin(context *gin.Context) {
	var user User
	//ret := make(map[string]any)
	err := context.ShouldBind(&user)
	if err != nil {
		context.JSON(http.StatusOK, tools.Ecode{
			//Code:   0,
			Message: "账号密码错误！",
		})
	}
	//作用是了解析请求中的数据将其绑定到user结构体上
	ret := model.GetUser(user.Name)
	if ret.ID < 1 || ret.Password != user.Password {
		context.JSON(http.StatusOK, tools.Ecode{
			//Code:   0,
			Message: err.Error(), //这里有风险，

		})
	}
	context.SetCookie("name", user.Name, 3600, "/", "", true, false)
	context.JSON(http.StatusOK, tools.Ecode{
		Message: "登陆成功",
		Data:    nil,
	})
}
