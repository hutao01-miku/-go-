package router

//路由，各中响应和请求路径放这里
import (
	"github.com/gin-gonic/gin"
	"net/http"
	"toupiao/application/logic"
)

func New() {
	r := gin.Default()
	r.LoadHTMLGlob("application/view/*")
	// 创建一个路由组（Group），用于添加中间件 checkUser
	index := r.Group("")
	index.Use(checkUser)
	// 定义路径和对应的处理函数
	index.GET("/index", logic.Index)

	r.GET("/", logic.Index) //主页

	r.GET("/login", logic.GetLogin)
	r.POST("/login", logic.DoLogin)
	//相关路径放这里
	if err := r.Run(":8080"); err != nil {
		panic("启动失败")
	}
}
func checkUser(context *gin.Context) {
	// 从请求中获取 Cookie 中的 "name" 值
	name, err := context.Cookie("name")
	if (err != nil) || name == "" {
		// 如果获取 Cookie 出现错误或者 Cookie 中的 "name" 值为空，则重定向到 "/login" 路径
		context.Redirect(http.StatusFound, "/login")

	}
	context.Next()
}
