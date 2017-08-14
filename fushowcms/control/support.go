package control

/*
 * Copyright (c) 2016—2017 www.fushow.cn, All rights reserved.
 */
import (
	"fmt"
	"fushowcms/comm"
	"fushowcms/models"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

//投注一注竞猜
//liuhan 20161130

func SupportAdd(c *gin.Context) {
	defer comm.FSLog.Unlock()
	comm.FSLog.Lock()

	var (
		smt    models.SupportManagement
		uk     models.UidInfo
		ui     models.UidInfo
		proper models.PeriodsProduct
		per    models.Periods
	)

	if c.PostForm("UID") == "" || c.PostForm("PeriodsId") == "" || c.PostForm("SupEncoding") == "" || c.PostForm("SupporNumber") == "" || c.PostForm("Odds") == "" {
		comm.ResponseError(c, 3177) // 参数错误
		return
	}
	supnum, _ := strconv.Atoi(c.PostForm("SupporNumber")) //获取投注的石榴籽数
	if supnum > 50000 {                                   // 人民币:石榴籽 = 1:100)
		comm.ResponseError(c, 3178) // 单笔最多支持500元
		return
	}

	smt.Uid, _ = strconv.ParseInt(c.PostForm("UID"), 10, 64)                   //添加用户Id
	smt.PeriodsId, _ = strconv.ParseInt(c.PostForm("PeriodsId"), 10, 64)       //添加期数Id
	smt.SupEncoding = c.PostForm("SupEncoding")                                //用户投注编码
	smt.SupporNumber, _ = strconv.ParseInt(c.PostForm("SupporNumber"), 10, 64) //支持数（投注数）
	smt.Odds = c.PostForm("Odds")                                              //添加赔率
	smt.SupporTime = time.Now().Format("2006-01-02 15:04:05")
	per.PeriodsId, _ = strconv.ParseInt(c.PostForm("PeriodsId"), 10, 64)

	if !per.GetPeriod() {
		comm.ResponseError(c, 3193) //该期数不存在
		return
	}

	if per.SupEncoding != "" {
		comm.ResponseError(c, 3194) //该期数已结束
		return
	}

	proper.PeriodsId = smt.PeriodsId
	pid := Substr(smt.SupEncoding, 1, 2)

	proper.ProductId, _ = strconv.ParseInt(pid, 10, 64)
	has := proper.GetProducts()
	if !has {
		fmt.Println("期过程表未找到")
	}

	//查询当前期 产品单个选项的支持次数，最多支持10次
	suptime := smt.GetMySupCount()
	if suptime == 10 {
		comm.ResponseError(c, 3179) // 此选项最多支持10次
		return
	}

	per.PeriodsId = smt.PeriodsId

	if !per.GetPeriod() {
		comm.ResponseError(c, 2046) // 当前期数不存在
		return
	}

	nowtime := GetDateTime()

	if nowtime > per.EndTime {
		comm.ResponseError(c, 2045) // 当前期数已结束
		return
	}

	uk.Id = smt.Uid
	ui.Id = smt.Uid
	flag, jjj := ui.UserCost(smt.SupporNumber)
	//用户扣除操作
	if !flag {
		osf := "当前余额不足"
		if strings.EqualFold(jjj, osf) {
			comm.ResponseError(c, 4040) //余额不足
		} else {
			comm.ResponseError(c, 3108) // 投注失败
		}
		return
	}
	if !uk.GetUserInfo() {
		comm.ResponseError(c, 3180) // 投注失败
		return
	}
	if !smt.SupportAdd() {
		comm.ResponseError(c, 3180) // 投注失败
		return
	}

	state := Substr(smt.SupEncoding, 4, 5)

	statenow, _ := strconv.ParseInt(state, 10, 64)

	//热度更新操作
	if statenow == 1 {
		proper.State1Hot = proper.State1Hot + smt.SupporNumber
	} else if statenow == 2 {
		proper.State2Hot = proper.State2Hot + smt.SupporNumber
	} else if statenow == 3 {
		proper.State3Hot = proper.State3Hot + smt.SupporNumber
	} else if statenow == 4 {
		proper.State4Hot = proper.State4Hot + smt.SupporNumber
	} else if statenow == 5 {
		proper.State5Hot = proper.State5Hot + smt.SupporNumber
	}

	_, err := proper.PeriodsProductUp()
	if err != nil {
		fmt.Println("过程表更新失败", err)
	}

	//投注金额supnum
	nowflag, _ := FundAddUp(float64(supnum))
	if !nowflag {
		comm.ResponseError(c, 3180) // 投注失败
		return
	}
	//更新操作
	comm.Response(c, 3181) // 投注成功
}

//删除支持
//TODO 不确定有没有此功能，应该是没有
func SupportDel(c *gin.Context) {
	var smt models.SupportManagement

	id, _ := strconv.ParseInt(c.PostForm("id"), 10, 64)
	smt.Id = id

	if !smt.SupportDel() { //调用models层删除支持方法
		c.JSON(200, gin.H{"state": "fail"})
		return
	}

	c.JSON(200, gin.H{"state": "success"})
}

//修改支持信息（通过id号）
//TODO 不确定有没有此功能
func SupportUp(c *gin.Context) {
	var smt models.SupportManagement

	smt.Id, _ = strconv.ParseInt(c.PostForm("Id"), 10, 64) //支持Id

	if !smt.GetSupport() {
		c.JSON(200, gin.H{"state": "fail"})
		return
	}
	smt.Uid, _ = strconv.ParseInt(c.PostForm("Uid"), 10, 64) //添加用户Id
	if smt.Uid == 0 {
		comm.ResponseError(c, 3175)
		return
	}
	smt.PeriodsId, _ = strconv.ParseInt(c.PostForm("PeriodsId"), 10, 64)       //添加期数Id
	smt.SupporNumber, _ = strconv.ParseInt(c.PostForm("SupporNumber"), 10, 64) //添加石榴籽	支持数（投注数）
	smt.Odds = c.PostForm("Odds")                                              //添加赔率

	smt.PrizenNmber, _ = strconv.ParseInt(c.PostForm("PrizenNmber"), 10, 64) //添加奖励支持石榴籽数	胜负状态，0：无

	if !smt.SupportUp() { //调用models层修改支持方法
		c.JSON(200, gin.H{"state": "fail"})
		return
	}
	c.JSON(200, gin.H{"state": smt})
}

//获取支持信息
//TODO 不确定有没此功能
func GetSupport(c *gin.Context) {
	var smt models.SupportManagement

	id, _ := strconv.ParseInt(c.PostForm("id"), 10, 64) //支持id
	smt.Id = id

	if !smt.GetSupport() {
		c.JSON(200, gin.H{"state": "fail"})
		return
	}
	c.JSON(200, gin.H{"state": smt})
}

//获取支持列表
//liuhan 20161130
func GetSupportList(c *gin.Context) {
	var smt models.SupportManagement

	page, _ := strconv.Atoi(c.PostForm("page"))
	rows, _ := strconv.Atoi(c.PostForm("rows"))
	inputid := c.PostForm("inputid")
	fmt.Println("inputid", inputid)
	list, row := smt.GetSupportList(page, rows, inputid)
	var arr [1]int // 声明了一个int类型的数组
	arr[0] = 99999 // 数组下标是从0开始的
	fmt.Println("row", row)
	if list == nil {
		fmt.Println("arr", arr)
		c.JSON(200, gin.H{"total": row, "rows": arr})
		return
	}
	c.JSON(200, gin.H{"total": row, "rows": list})

}

//获取支持列表
//liuhan 20161130
func GetSupportUidList(c *gin.Context) {
	var smt models.SupportManagement
	uid, _ := strconv.ParseFloat(c.PostForm("UID"), 64)
	page, _ := strconv.Atoi(c.PostForm("page"))
	rows, _ := strconv.Atoi(c.PostForm("rows"))
	smt.Uid = int64(uid)
	if smt.Uid == 0 {
		comm.ResponseError(c, 3175)
		return
	}
	list, total := smt.GetPeriodSupportUidList(page, rows)
	fmt.Println(list)
	if len(list) == 0 {
		comm.ResponseError(c, 3174) //支持数据不存在
		return
	}
	m := make(map[string]interface{})
	m["list"] = list
	m["total"] = total
	comm.Response(c, m)
}
