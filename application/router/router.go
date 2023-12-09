package router

//路由，各中响应和请求路径放这里
import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"toupiao/application/logic"
	"toupiao/application/model"
	"toupiao/application/tools"
)

func New() {
	r := gin.Default()
	r.LoadHTMLGlob("application/view/*")
	r.Static("css", "./application/css")
	r.Static("/static", "./application/img")
	// 创建一个路由组（Group），用于添加中间件 checkUser
	index := r.Group("")
	index.Use(checkUser)
	{ //vote相关

		index.GET("/index", logic.Index) //静态页面
		// 定义路径和对应的处理函数
		index.GET("/votes", logic.GetVotes)
		index.GET("/vote", logic.GetVoteInfo)
		index.POST("/vote", logic.DoVote)

		index.POST("/vote/add", logic.AddVote)
		index.POST("/vote/update", logic.UpdateVote)
		index.POST("/vote/delete", logic.DeleteVote)

		index.GET("result", logic.ResultInfo)
		index.GET("result/info", logic.ResultVote)
	}
	//restFul风格接口
	//{
	//	//读
	//	index.GET("/votes", logic.GetVotes)
	//	index.GET("/vote", logic.GetVoteInfo)
	//	//
	//	index.POST("/vote", logic.AddVote)
	//	index.PUT("/vote", logic.UpdateVote)
	//	index.DELETE("/vote", logic.DeleteVote)
	//	index.GET("/vote/result", logic.ResultVote)
	//	index.POST("/do_vote", logic.DoVote)
	//}

	r.GET("/", logic.Index) //主页
	{
		r.GET("/login", logic.GetLogin)
		r.POST("/login", logic.DoLogin)
		r.GET("/logout", logic.Logout)
		//相关路径放这里

		//user
		r.POST("/user/create", logic.CreateUser)
		r.GET("/register", logic.RegisterUser)
	}
	r.GET("/captcha", func(context *gin.Context) {
		captcha, err := tools.CaptchaGenerate()
		if err != nil {
			context.JSON(http.StatusOK, tools.ECode{
				Code:    10005,
				Message: err.Error(),
			})
			return
		}
		context.JSON(http.StatusOK, tools.ECode{
			Data: captcha,
		})
	})

	r.POST("/captcha/verify", func(context *gin.Context) {
		var param tools.CaptchaData
		if err := context.ShouldBindJSON(&param); err != nil {
			context.JSON(http.StatusOK, tools.ParamErr)
			return
		}
		fmt.Printf("参数为: %+v", param)
		if !tools.CaptchaVerify(param) {
			context.JSON(http.StatusOK, tools.ECode{
				Code:    10008,
				Message: "验证失败",
			})
			return
		}
		context.JSON(http.StatusOK, tools.OK)
	})

	if err := r.Run(":8081"); err != nil {
		panic("启动失败")
	}
}
func checkUser(context *gin.Context) {
	var name string
	var id int64
	values := model.GetSession(context)
	if v, ok := values["name"]; ok {
		name = v.(string)
	}
	if v, ok := values["id"]; ok {
		id = v.(int64) //todo 存在bug
	}
	if name == "" || id <= 0 {
		context.JSON(http.StatusOK, tools.ECode{
			Code:    0,
			Message: "用户未登录",
		})
		context.Abort()
	}

	//从请求中获取 Cookie 中的 "name" 值
	name, err := context.Cookie("name")
	if (err != nil) || name == "" {
		// 如果获取 Cookie 出现错误或者 Cookie 中的 "name" 值为空，则重定向到 "/login" 路径
		context.Redirect(http.StatusFound, "/login")

	}
	context.Next()
}
