package application

import (
	"awesomeProject1/application/model"
	"awesomeProject1/application/router"
)

// 这个里边放启动器的代码
func Start() {
	model.NewMysql()
	defer func() {
		model.Close()
	}()

	router.New()
}

//func main() {
//	Start()
//	fmt.Println("Hello World")
//}
