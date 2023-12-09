package model

import (
	"fmt"
	"gorm.io/gorm"
	"time"
)

func GetVotes() []Vote {
	//输出Vote类型的值
	//封装查询方法,查询投票项目的详情
	ret := make([]Vote, 0) //切片类型，0表示长度为0
	//ret := make(map[string]any)
	err := Conn.Table("vote").Find(&ret).Error
	if err != nil {
		fmt.Printf("err:%s", err.Error())
	}
	return ret
}

func GetVote(id int64) VoteWithOpt {
	//封装查询方法，查询指定ID的投票项目及其选项和属性
	//同时查两个表vote和vote_opt的部分
	var ret Vote //vote查的是单个表名及其属性
	err := Conn.Table("vote").Where("id=?", id).First(&ret).Error
	if err != nil {
		fmt.Printf("err:%s", err.Error())
	}
	opt := make([]VoteOpt, 0) //opt查的是对应vote_id的多条数据。所以用切片
	err1 := Conn.Table("vote_opt").Where("vote_id=?", id).Find(&opt).Error
	if err1 != nil {
		fmt.Printf("err:%s", err.Error())
	}
	return VoteWithOpt{ //返回一个包含查询结果的结构体
		Vote: ret,
		Opt:  opt,
	}
}

// DoVote 函数用于处理用户的投票操作，接收用户ID、投票ID和选项ID列表作为参数，把数据存入数据库中
// ，返回一个布尔值表示操作是否成功
func DoVote(userId, voteId int64, optIds []int64) bool { //太复杂，//gorm中最常用的事务处理方法
	tx := Conn.Begin() //创建一个数据库事务，Begin() 方法开始一个新的事务并返回对应的 *gorm.DB 对象 tx。

	var ret Vote
	err := tx.Table("vote").Where("id=?", voteId).First(&ret).Error
	//在事务中执行一个查询操作，根据给定的 voteId 查询 vote 表中的数据，并将结果存储在 ret 变量中。
	if err != nil { //出现错误，打印出来并回滚；
		fmt.Printf("err:%s", err.Error())
		tx.Rollback()
	}
	//var oldVoteUser VoteOptUser
	//err = tx.Table("vote_opt_user").Where("vote_id=? ans", voteId).First(&oldVoteUser).Error
	//if err != nil {
	//	fmt.Printf("err:%s")
	//	tx.Rollback()
	//}
	//if oldVoteUser.Id > 0 {
	//	fmt.Printf("用户已投票")
	//	tx.Rollback()
	//}

	for _, value := range optIds { // 遍历选项ID列表
		err := tx.Table("vote_opt").Where("id=?", value).Update("count", gorm.Expr("count+?", 1)).Error
		if err != nil {
			//更新数据库中的选项计数
			fmt.Printf("err:%s", err.Error())
			tx.Rollback()
		}
		user := VoteOptUser{ //创建一个新的VoteOptUser结构体实例对应数据库的vote_opt_user表
			VoteId:     voteId,
			UserId:     userId,
			VoteOptId:  value,
			CreateTime: time.Now(),
			UpdateTime: time.Now(),
		}
		err = tx.Create(&user).Error
		if err != nil {
			fmt.Printf("err:%s", err.Error())
			tx.Rollback()
		}
		//使用Create方法将新的VoteOptUser记录创建到数据库中
	}
	tx.Commit()
	return true
}

// 用原生sql优化
func DoVoteV3(userId, voteId int64, optIds []int64) bool {
	tx := Conn.Begin() //创建一个数据库事务，Begin() 方法开始一个新的事务并返回对应的 *gorm.DB 对象 tx。
	var ret Vote
	err := tx.Raw("select * from vote where id = ?", voteId).Scan(&ret).Error
	//在事务中执行一个查询操作，根据给定的 voteId 查询 vote 表中的数据，并将结果存储在 ret 变量中。
	if err != nil { //出现错误，打印出来并回滚；
		fmt.Printf("err:%s", err.Error())
		tx.Rollback()
	}

	var oldVoteUser VoteOptUser
	err = tx.Raw("select * from vote_opt_user where vote_id=? and user_id=?", voteId, voteId).Scan(&oldVoteUser).Error
	if err != nil {
		fmt.Printf("err:%s")
		tx.Rollback()
	}

	if oldVoteUser.Id > 0 {
		fmt.Printf("用户已投票")
		tx.Rollback()
	}
	for _, value := range optIds { // 遍历选项ID列表
		tx.Exec("update vote_opt set count=count+1 where id=? limit 1", value)
		err := tx.Exec("update vote_opt set count=count+1 where id=? limit 1", value).Error
		if err != nil {
			//更新数据库中的选项计数
			fmt.Printf("err:%s", err.Error())
			tx.Rollback()
		}
		user := VoteOptUser{ //创建一个新的VoteOptUser结构体实例对应数据库的vote_opt_user表
			VoteId:     voteId,
			UserId:     userId,
			VoteOptId:  value,
			CreateTime: time.Now(),
			UpdateTime: time.Now(),
		}
		err = tx.Create(&user).Error
		if err != nil {
			fmt.Printf("err:%s", err.Error())
			tx.Rollback()
		}
		//使用Create方法将新的VoteOptUser记录创建到数据库中
	}
	tx.Commit()
	return true
}

func AddVote(vote Vote, opt []VoteOpt) error {
	err := Conn.Transaction(func(tx *gorm.DB) error {
		// 创建投票记录
		if err := tx.Create(&vote).Error; err != nil {
			return err
		}
		// 遍历选项并关联到投票记录
		for _, voteOpt := range opt {
			voteOpt.VoteId = vote.Id
			if err := tx.Create(&voteOpt).Error; err != nil {
				return err
			}
		}
		return nil
	})
	return err
}
func DeleteVote(id int64) bool {
	if err := Conn.Transaction(func(tx *gorm.DB) error {

		if err := tx.Exec("delete form vote where id=? limit 1", id).Error; err != nil {
			fmt.Printf("err:%s", err.Error())
			return err
		}

		if err := tx.Exec("delete form vote_opt where vote_id=? limit 1", id).Error; err != nil {
			fmt.Printf("err:%s", err.Error())
			return err
		}

		if err := tx.Exec("delete form vote_opt_user where vote_id=?", id).Error; err != nil {
			fmt.Printf("err:%s", err.Error())
			return err
		}
		return nil
	}); err != nil {
		fmt.Printf("err:%s", err.Error())
		return false
	}

	return true
}
func UpdateVote(vote Vote, opt []VoteOpt) error {

	err := Conn.Transaction(func(tx *gorm.DB) error {
		// 创建投票记录
		if err := tx.Save(&vote).Error; err != nil {
			return err
		}
		// 遍历选项并关联到投票记录
		for _, voteOpt := range opt {
			voteOpt.VoteId = vote.Id
			if err := tx.Create(&voteOpt).Error; err != nil {
				return err
			}
		}
		return nil
	})
	return err
}

func GetVoteHistory(userId, voteId int64) []VoteOptUser {
	ret := make([]VoteOptUser, 0)

	err := Conn.Table("vote_opt_user").Where("vote_id=? ans", voteId).First(&ret).Error
	if err != nil {
		fmt.Printf("err:%s")
	}
	return ret
}
func EndVote() {
	votes := make([]Vote, 0)
	err := Conn.Table("vote").Where("status=?", 1).Find(&votes).Error
	if err != nil {
		return
	}
	now := time.Now().Unix() // 获取当前时间戳
	for _, vote := range votes {
		if vote.Time+vote.CreatedTime.Unix() <= now {
			// 执行更新操作
			Conn.Table("vote").Where("id=?", vote.Id).Update("status", 0)
		}
	}
	return
}

// 原声sql写法
func EneVoteV1() {
	votes := make([]Vote, 0)
	err := Conn.Raw("select * from vote where status=?", 1).Find(&votes).Error
	if err != nil {
	}
	err = Conn.Table("vote").Where("status=?", 1).Find(&votes).Error
	if err != nil {
		return
	}
	now := time.Now().Unix() // 获取当前时间戳
	for _, vote := range votes {
		if vote.Time+vote.CreatedTime.Unix() <= now {
			// 执行更新操作
			Conn.Table("vote").Where("id=?", vote.Id).Update("status", 0)
		}
	}
	return
}
