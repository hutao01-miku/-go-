package application

// 这个里边放启动器的代码
import (
	"toupiao/application/model"
	"toupiao/application/router"
)

func Start() {

	model.NewMysql()
	defer func() { //最后运行，结束数据库
		model.Close()
	}()
	router.New()
}
