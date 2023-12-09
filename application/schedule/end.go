package schedule

//增加定时器功能
import (
	"fmt"
	"time"
	"toupiao/application/model"
)

func Start() {
	//使用协程（goroutine）的目的是让 EndVote 函数在一个独立的并发执行线程中运行，而不会阻塞 Start 函数的执行。不然会阻塞start函数执行，意味着EndVote函数会一直执行下去；
	//Start 函数启动一个 goroutine,在这个 goroutine中调用了 voteEnd 函数。
	go func() { //使用 go 关键字创建 goroutine 表示这个函数是异步执行的，不会阻塞当前程序的执行。
		EndVote()
	}()
	return
}

func EndVote() {
	t := time.NewTicker(86400 * time.Second)
	//每秒触发一次
	defer t.Stop() //最后运行关闭定时器

	for {
		select { //监听定时器的触发事件，
		case <-t.C:
			fmt.Printf("定时器voteEnd启动")
			//执行函数
			model.EndVote()
			fmt.Println("EndVote 运行完毕")
		}
	}
}
