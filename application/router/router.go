package router

//路由，各中响应和请求路径放这里
import (
	"github.com/gin-gonic/gin"
	"toupiao/application/logic"
)

func New() {
	r := gin.Default()
	r.LoadHTMLGlob("application/view/*")
	//var user model.User"/login", logic.GetLogin)
	r.GET("/login", logic.GetLogin)
	r.POST("/login", logic.DoLogin)
	//相关路径放这里
	if err := r.Run(":8080"); err != nil {
		panic("启动失败")
	}
}
