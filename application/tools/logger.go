package tools

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"io"
	"os"
)

var Logger *logrus.Entry

func NewLogger() {
	logStore := logrus.New()
	logStore.SetLevel(logrus.DebugLevel)

	// 同时写到多个输出
	w1 := os.Stdout // 写到控制台
	w2, err := os.OpenFile("./vote.log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	//写到文件中；
	if err != nil {
		// 错误处理，例如打印错误并退出程序
		fmt.Println("Error opening log file:", err)
		os.Exit(1)
	}

	logStore.SetOutput(io.MultiWriter(w1, w2))     // io.MultiWriter 返回一个 io.Writer 对象
	logStore.SetFormatter(&logrus.JSONFormatter{}) //生成json格式的日志
	Logger = logStore.WithFields(logrus.Fields{
		"name": "香香编程喵喵喵",
		"app":  "voteV2",
	})
}

//logStore.AddHook()//钩子函数：如果出现非常严重的错误，直接微信或者邮箱报警，日志分割，能将日志中塞入一些字段；
