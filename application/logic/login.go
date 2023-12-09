package logic

//逻辑层，相当于controller
import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"regexp"
	"time"
	"toupiao/application/model"
	"toupiao/application/tools"
)

type User struct { //登录页面的记录的账号和密码
	Name string `json:"name" form:"name"`
	//标签指定了在JSON和form格式的数据中如何读取该字段
	Password     string `json:"password" form:"password"`
	CaptchaId    string `json:"captcha_id" form:"captcha_id"`
	CaptchaValue string `json:"captcha_value" form:"captcha_value"`
}

// GetLogin 处理用户访问登录页面的GET请求
func GetLogin(context *gin.Context) {
	context.HTML(http.StatusOK, "login.tmpl", nil)
}
func RegisterUser(context *gin.Context) {
	context.HTML(http.StatusOK, "register.html", nil)
}

// DoLogin 处理用户提交登录表单的POST请求
func DoLogin(context *gin.Context) {
	var user User
	err := context.ShouldBind(&user) //该方法解析能JSON或form数据并将结果存入user变量
	if err != nil {
		context.JSON(http.StatusOK, tools.ECode{
			Code:    0, //自定义码
			Message: "账号密码错误！",
		})
		return
	}
	if !tools.CaptchaVerify(tools.CaptchaData{
		CaptchaId: user.CaptchaId,
		Data:      user.CaptchaValue,
	}) {
		context.JSON(http.StatusOK, tools.ECode{
			Code:    10010,
			Message: "验证码校验失败！",
		})
		return
	}
	//作用是把数据库中的数据绑定到user结构体上
	ret := model.GetUser(user.Name)
	// 通过调用model模块的GetUser函数获取数据库中与传入用户名对应的用户信息
	if ret.Id < 1 || ret.Password != user.Password {
		context.JSON(http.StatusOK, tools.ECode{
			//Code:   0,
			Message: err.Error(), //这里有风险，

		})
		return
	}
	context.SetCookie("name", user.Name, 3600, "/", "", true, false)
	//设置两个名为"name"的cookie，第一个存储用户名，第二个存储用户ID，两者都将在一小时后过期
	context.SetCookie("userId", fmt.Sprint(ret.Id), 3600, "/", "", true, false)
	_ = model.SetSession(context, user.Name, ret.Id)
	context.JSON(http.StatusOK, tools.ECode{
		Message: "登陆成功",
		Data:    nil,
	})
	return
}

// Logout 处理用户注销操作的GET或POST请求
func Logout(context *gin.Context) {
	context.SetCookie("name", "", 3600, "/", "", true, false)
	// 清除所有名为"name"的cookie,重定向到登陆页面
	//context.SetCookie("name", "", 3600, "/", "", true, false)
	_ = model.FlushSession(context)
	context.Redirect(http.StatusFound, "/login")
}

// 新创建一个结构体
type CUser struct {
	Name      string `json:"name"`
	Password  string `json:"password"`
	Password2 string `json:"password_2"`
}

func CreateUser(context *gin.Context) {
	var user CUser
	if err := context.ShouldBind(&user); err != nil {
		context.JSON(200, tools.ECode{
			Code:    10001,
			Message: err.Error(),
		})
		return
	}
	fmt.Println("user:%+v", user)
	encryptV1(user.Password)
	//return

	if user.Name == "" || user.Password == "" || user.Password2 == "" {
		context.JSON(200, tools.ParamErr)
		return
	}
	//校验密码
	if user.Password != user.Password2 {
		context.JSON(200, tools.ECode{
			Code:    10003,
			Message: "两次密码不同",
		})
		return
	}

	nameLen := len(user.Name)
	passwordLen := len(user.Password)
	if nameLen > 16 || nameLen < 8 || passwordLen > 16 || passwordLen < 8 {
		context.JSON(200, tools.ECode{
			Code:    10005,
			Message: "账号或密码要大于8位小于16位！",
		})
		return
	}
	//密码不能是纯数字，应该是数字+小写字母+大写字母
	regex := regexp.MustCompile(`^[0-9]+$`)
	if regex.MatchString(user.Password) {
		context.JSON(http.StatusOK, tools.ECode{
			Code:    10006,
			Message: "密码不能为纯数字", // 这里有风险
		})
		return
	}
	//会有风险，并发安全,当同名用户同时注册时，会出现重名
	if oldUser := model.GetUser(user.Name); oldUser.Id > 0 {
		context.JSON(200, tools.ECode{
			Code:    10004,
			Message: "用户名已存在！",
		})
		return
	}
	//用户
	newUser := model.User{
		Name:        user.Name,
		Password:    encryptV1(user.Password),
		CreatedTime: time.Now(),
		UpdateTime:  time.Now(),
		Uuid:        tools.GetUUID(),
	}
	err := model.CreateUser(&newUser)
	if err != nil {
		context.JSON(200, tools.ECode{
			Code:    10007,
			Message: "新用户创建失败",
		})
		return
	}
	context.JSON(200, tools.ECode{
		Message: "注册成功",
	})
	return
}

// 数据加密：第二种方式，在第一种基础上进行操作，具体来说是把传过来的密码使用
func encryptV1(pwd string) string {
	secretString := "香香编程喵喵喵"
	newPwd := pwd + secretString // Concatenate the secret string to the password
	hash := md5.New()
	hash.Write([]byte(newPwd))
	hashBytes := hash.Sum(nil)
	hashString := hex.EncodeToString(hashBytes)
	fmt.Printf("加密后的密码: %s\n", hashString)
	return hashString
}

// 第三种方式加密。使用不同的加密方式：
func encryptV2(pwd string) string {
	// 基于Blowfish实现加密。简单快速，但有安全风险
	// golang.org/x/crypto/中有大量的加密算法
	newPwd, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println("密码加密失败:", err)
		return ""
	}
	newPwdStr := string(newPwd)
	fmt.Printf("加密后的密码: %s\n", newPwdStr)
	return newPwdStr
}

// 数据加密:第一种方式，直接用md5加密，很容易被装库试出来
func encrypt(pwd string) string {
	hash := md5.New()
	hash.Write([]byte(pwd))
	hashBytes := hash.Sum(nil)

	hashString := hex.EncodeToString(hashBytes)
	fmt.Printf("加密后的密码: %s\n", hashString)
	return hashString
}
