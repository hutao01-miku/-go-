package logic

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"toupiao/application/model"
	"toupiao/application/tools"
)

func Index(context *gin.Context) { //渲染前端：把投票表都传给前端
	ret := model.GetVotes()
	context.HTML(http.StatusOK, "index.tmpl", gin.H{"vote": ret})
}

func GetVotes(context *gin.Context) {
	ret := model.GetVotes()
	context.JSON(http.StatusOK, tools.ECode{
		Data: ret,
	})
}
func GetVoteInfo(context *gin.Context) {
	//处理HTTP请求，从查询参数中获取一个ID，然后使用该ID调用model.GetVote(id)
	//来获取投票信息，最后将结果渲染到一个HTML模板中。
	var id int64
	idStr := context.Query("id")
	id, _ = strconv.ParseInt(idStr, 10, 64)
	// 将idStr转换为int64类型的整数，并将结果存储在id变量中
	ret := model.GetVote(id) //向ret中传入id中的值

	//logrus.Errorf("[error]ret:%+v", ret)
	//tools.Logger.Errorf("[error]ret:%+v", ret)

	//log.Printf("[print]ret:%+v", ret)
	//log.Panicf("[fatal]ret:%+v", ret)
	context.JSON(http.StatusOK, tools.ECode{
		Data: ret,
	})
	context.HTML(http.StatusOK, "vote.tmpl", gin.H{"vote": ret})
	//把ret的值传给模版
}

func DoVote(context *gin.Context) {
	userIdStr, _ := context.Cookie("Id")
	voteIdStr, _ := context.GetPostForm("vote_id")
	optStr, _ := context.GetPostForm("opt[]")

	userId, _ := strconv.ParseInt(userIdStr, 10, 64)
	voteId, _ := strconv.ParseInt(voteIdStr, 10, 64)
	old := model.GetVoteHistory(userId, voteId)
	if len(old) >= 1 {
		context.JSON(200, tools.ECode{
			Code:    10010,
			Message: "您已投过票",
		})
	}
	//model.GetVoteHistory()
	opt := make([]int64, 0)    //存储投票选项
	for _, v := range optStr { //遍历投票选项字符串数组
		optId, _ := strconv.ParseInt(string(v), 10, 64)
		opt = append(opt, optId) //将转换后的整数添加到opt切片中；
	}
	model.DoVote(userId, voteId, opt) //调用DoVote（）方法，把接收的数据传给数据库
	context.JSON(200, tools.ECode{
		Message: "投票完成",
		Data:    nil,
	})

}
