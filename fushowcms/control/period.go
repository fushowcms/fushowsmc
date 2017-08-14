package control

/*
 * Copyright (c) 2016—2017 www.fushow.cn, All rights reserved.
 */
import (
	"fmt"
	"fushowcms/comm"
	m "fushowcms/models"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

//var Engine *xorm.Engine

//添加期数
func PeriodAdd(c *gin.Context) {
	var (
		per m.Periods
	)
	//事务开始

	per.PeriodsId, _ = strconv.ParseInt(m.GetPerId(), 10, 64)
	if per.PeriodsId == 0 {
		comm.ResponseError(c, 5033)
		return
	}
	per.StartTime = c.PostForm("StartTime")
	per.EndTime = c.PostForm("EndTime")
	per.ATeam = c.PostForm("ATeam")
	per.BTeam = c.PostForm("BTeam")
	per.State, _ = strconv.ParseInt(c.PostForm("State"), 10, 64)

	proEncoding := c.PostForm("ProEncoding") //产品选择编码

	if len(proEncoding) <= 0 {
		comm.ResponseError(c, 5039)
		return
	}
	per.ProEncoding = proEncoding

	//判断ar.Id是否已#号开头
	if !strings.HasPrefix(c.PostForm("RoomId"), "#") {
		comm.ResponseError(c, 5001)
		return
	}

	//字符串替换
	str := strings.Replace(c.PostForm("RoomId"), "#_#", "#", -1)
	str = strings.Replace(str, "_", "", -1)
	//字符串分割
	canSplit := func(c rune) bool { return c == '#' }
	roomid := strings.FieldsFunc(str, canSplit)

	for _, num := range roomid {
		var ar m.AnchorRoom
		ar.Id, _ = strconv.ParseInt(num, 10, 64)
		//判断房间是否存在
		if !ar.IsRoomExist() {
			comm.ResponseError(c, 5002)
			return
		}
	}

	// 判断结束时间是否大于开始时间
	t1, err := time.Parse("2006-01-02 15:04:05", per.StartTime)
	t2, err := time.Parse("2006-01-02 15:04:05", per.EndTime)
	if err == nil && t2.Before(t1) {
		comm.ResponseError(c, 5003)
		return
	}
	for _, nums := range roomid {
		per.RoomId += "#" + nums
	}

	proen := strings.Split(proEncoding, ",")
	for _, value := range proen {
		var proper m.PeriodsProduct
		proper.PeriodsId = per.PeriodsId
		proper.ProductId, _ = strconv.ParseInt(value, 10, 64)
		//热度
		proper.State1Hot, _ = strconv.ParseInt(c.PostForm(value+"State1Hot"), 10, 64)
		proper.State2Hot, _ = strconv.ParseInt(c.PostForm(value+"State2Hot"), 10, 64)
		proper.State3Hot, _ = strconv.ParseInt(c.PostForm(value+"State3Hot"), 10, 64)
		proper.State4Hot, _ = strconv.ParseInt(c.PostForm(value+"State4Hot"), 10, 64)
		proper.State5Hot, _ = strconv.ParseInt(c.PostForm(value+"State5Hot"), 10, 64)
		proper.State5Hot, _ = strconv.ParseInt(c.PostForm(value+"State5Hot"), 10, 64)
		proper.State6Hot, _ = strconv.ParseInt(c.PostForm(value+"State6Hot"), 10, 64)
		proper.State7Hot, _ = strconv.ParseInt(c.PostForm(value+"State7Hot"), 10, 64)
		proper.State8Hot, _ = strconv.ParseInt(c.PostForm(value+"State8Hot"), 10, 64)
		proper.State9Hot, _ = strconv.ParseInt(c.PostForm(value+"State9Hot"), 10, 64)
		proper.State10Hot, _ = strconv.ParseInt(c.PostForm(value+"State10Hot"), 10, 64)
		//赔率
		proper.State1Odds = c.PostForm(value + "State1Odds")
		proper.State2Odds = c.PostForm(value + "State2Odds")
		proper.State3Odds = c.PostForm(value + "State3Odds")
		proper.State4Odds = c.PostForm(value + "State4Odds")
		proper.State5Odds = c.PostForm(value + "State5Odds")
		proper.State6Odds = c.PostForm(value + "State6Odds")
		proper.State7Odds = c.PostForm(value + "State7Odds")
		proper.State8Odds = c.PostForm(value + "State8Odds")
		proper.State9Odds = c.PostForm(value + "State9Odds")
		proper.State10Odds = c.PostForm(value + "State10Odds")
		if proper.State1Odds != "" {
			match1, _ := regexp.MatchString("((^\\d+):([1-9]\\d*\\.\\d*|0\\.\\d*[1-9]\\d|\\d*)$)", proper.State1Odds)
			fmt.Println("match1", match1)
			if !match1 {
				comm.ResponseError(c, 5038)
				return
			}
		}
		if proper.State2Odds != "" {
			match2, _ := regexp.MatchString("((^\\d+):([1-9]\\d*\\.\\d*|0\\.\\d*[1-9]\\d|\\d*)$)", proper.State2Odds)
			fmt.Println("match2", match2)
			if !match2 {
				comm.ResponseError(c, 5038)
				return
			}
		}
		if proper.State3Odds != "" {
			match3, _ := regexp.MatchString("((^\\d+):([1-9]\\d*\\.\\d*|0\\.\\d*[1-9]\\d|\\d*)$)", proper.State3Odds)
			fmt.Println("match3", match3)
			if !match3 {
				comm.ResponseError(c, 5038)
				return
			}
		}
		if proper.State4Odds != "" {
			match4, _ := regexp.MatchString("((^\\d+):([1-9]\\d*\\.\\d*|0\\.\\d*[1-9]\\d|\\d*)$)", proper.State4Odds)
			fmt.Println("match4", match4)
			if !match4 {
				comm.ResponseError(c, 5038)
				return
			}
		}
		if proper.State5Odds != "" {
			match5, _ := regexp.MatchString("((^\\d+):([1-9]\\d*\\.\\d*|0\\.\\d*[1-9]\\d|\\d*)$)", proper.State5Odds)
			fmt.Println("match5", match5)
			if !match5 {
				comm.ResponseError(c, 5038)
				return
			}
		}
		if proper.State6Odds != "" {
			match6, _ := regexp.MatchString("((^\\d+):([1-9]\\d*\\.\\d*|0\\.\\d*[1-9]\\d|\\d*)$)", proper.State6Odds)
			fmt.Println("match6", match6)
			if !match6 {
				comm.ResponseError(c, 5038)
				return
			}
		}
		if proper.State7Odds != "" {
			match7, _ := regexp.MatchString("((^\\d+):([1-9]\\d*\\.\\d*|0\\.\\d*[1-9]\\d|\\d*)$)", proper.State7Odds)
			fmt.Println("match7", match7)
			if !match7 {
				comm.ResponseError(c, 5038)
				return
			}
		}
		if proper.State8Odds != "" {
			match8, _ := regexp.MatchString("((^\\d+):([1-9]\\d*\\.\\d*|0\\.\\d*[1-9]\\d|\\d*)$)", proper.State8Odds)
			fmt.Println("match8", match8)
			if !match8 {
				comm.ResponseError(c, 5038)
				return
			}
		}
		if proper.State9Odds != "" {
			match9, _ := regexp.MatchString("((^\\d+):([1-9]\\d*\\.\\d*|0\\.\\d*[1-9]\\d|\\d*)$)", proper.State9Odds)
			fmt.Println("match9", match9)
			if !match9 {
				comm.ResponseError(c, 5038)
				return
			}
		}
		if proper.State10Odds != "" {
			match10, _ := regexp.MatchString("((^\\d+):([1-9]\\d*\\.\\d*|0\\.\\d*[1-9]\\d|\\d*)$)", proper.State10Odds)
			fmt.Println("match10", match10)
			if !match10 {
				comm.ResponseError(c, 5038)
				return
			}
		}

		if !proper.PeriodsProductAdd() {
			comm.ResponseError(c, 5005)
			return
		}
	}

	//期Id 不可重复
	if !per.PeriodAdd() {
		comm.ResponseError(c, 5004)
		return
	}

	comm.Response(c, "成功添加")
}

//期过程属性表
func GetPeriodIdProduct(c *gin.Context) {
	var (
		sup_encoding1  string
		sup_encoding2  string
		sup_encoding3  string
		sup_encoding4  string
		sup_encoding5  string
		sup_encoding6  string
		sup_encoding7  string
		sup_encoding8  string
		sup_encoding9  string
		sup_encoding10 string
		sm             m.SupportManagement //查询支持的数量
		per            m.PeriodsProduct
	)
	number := make(map[string]m.Number)

	per.PeriodsId, _ = strconv.ParseInt(c.PostForm("PeriodsId"), 10, 64) //期Id
	product := c.PostForm("ProEncoding")                                 //多个产品编码 1,3,7

	list, rows := per.GetPeriodsProductList()
	if rows <= 0 {
		comm.ResponseError(c, 5011)
		return
	}

	var i int = 0
	if product != "" && len(product) > 0 {
		product += ","
	}

	//判断字符串是否含有逗号
	if strings.Contains(product, ",") {
		//字符串分割
		productids := strings.Split(product, ",")
		for _, num := range productids {
			var nowper m.PeriodsProduct
			nowper.PeriodsId = per.PeriodsId
			nowper.ProductId, _ = strconv.ParseInt(num, 10, 64)
			nowper.GetPeriodsProducts()
			sm.PeriodsId = nowper.PeriodsId

			//判断产品状态值是否为空
			if nowper.State1Odds != "" {
				//判断产品ID是否小于10
				if nowper.ProductId < 10 {
					sup_encoding1 = "#0" + strconv.FormatInt(int64(nowper.ProductId), 10) + ">01"
				} else {
					sup_encoding1 = "#" + strconv.FormatInt(int64(nowper.ProductId), 10) + ">01"
				}
				sm.SupEncoding = sup_encoding1
				count := sm.GetSupCountByEncoding()
				number[strconv.Itoa(i)] = count
				i++
			}
			if nowper.State2Odds != "" {
				//判断产品ID是否小于10
				if nowper.ProductId < 10 {
					sup_encoding2 = "#0" + strconv.FormatInt(int64(nowper.ProductId), 10) + ">02"
				} else {
					sup_encoding2 = "#" + strconv.FormatInt(int64(nowper.ProductId), 10) + ">02"
				}
				sm.SupEncoding = sup_encoding2
				count := sm.GetSupCountByEncoding()
				number[strconv.Itoa(i)] = count
				i++
			}

			if nowper.State3Odds != "" {
				//判断产品ID是否小于10
				if nowper.ProductId < 10 {
					sup_encoding3 = "#0" + strconv.FormatInt(int64(nowper.ProductId), 10) + ">03"
				} else {
					sup_encoding3 = "#" + strconv.FormatInt(int64(nowper.ProductId), 10) + ">03"
				}
				sm.SupEncoding = sup_encoding3
				count := sm.GetSupCountByEncoding()
				number[strconv.Itoa(i)] = count
				i++
			}

			if nowper.State4Odds != "" {
				//判断产品ID是否小于10
				if nowper.ProductId < 10 {
					sup_encoding4 = "#0" + strconv.FormatInt(int64(nowper.ProductId), 10) + ">04"
				} else {
					sup_encoding4 = "#" + strconv.FormatInt(int64(nowper.ProductId), 10) + ">04"
				}
				sm.SupEncoding = sup_encoding4
				count := sm.GetSupCountByEncoding()
				number[strconv.Itoa(i)] = count
				i++
			}
			if nowper.State5Odds != "" {
				//判断产品ID是否小于10
				if nowper.ProductId < 10 {
					sup_encoding5 = "#0" + strconv.FormatInt(int64(nowper.ProductId), 10) + ">05"
				} else {
					sup_encoding5 = "#" + strconv.FormatInt(int64(nowper.ProductId), 10) + ">05"
				}
				sm.SupEncoding = sup_encoding5
				count := sm.GetSupCountByEncoding()
				number[strconv.Itoa(i)] = count
				i++
			}
			if nowper.State6Odds != "" {
				//判断产品ID是否小于10
				if nowper.ProductId < 10 {
					sup_encoding6 = "#0" + strconv.FormatInt(int64(nowper.ProductId), 10) + ">06"
				} else {
					sup_encoding6 = "#" + strconv.FormatInt(int64(nowper.ProductId), 10) + ">06"
				}
				sm.SupEncoding = sup_encoding6
				count := sm.GetSupCountByEncoding()
				number[strconv.Itoa(i)] = count
				i++
			}
			if nowper.State7Odds != "" {
				//判断产品ID是否小于10
				if nowper.ProductId < 10 {
					sup_encoding7 = "#0" + strconv.FormatInt(int64(nowper.ProductId), 10) + ">07"
				} else {
					sup_encoding7 = "#" + strconv.FormatInt(int64(nowper.ProductId), 10) + ">07"
				}
				sm.SupEncoding = sup_encoding7
				count := sm.GetSupCountByEncoding()
				number[strconv.Itoa(i)] = count
				i++
			}
			if nowper.State8Odds != "" {
				//判断产品ID是否小于10
				if nowper.ProductId < 10 {
					sup_encoding8 = "#0" + strconv.FormatInt(int64(nowper.ProductId), 10) + ">08"
				} else {
					sup_encoding8 = "#" + strconv.FormatInt(int64(nowper.ProductId), 10) + ">08"
				}
				sm.SupEncoding = sup_encoding8
				count := sm.GetSupCountByEncoding()
				number[strconv.Itoa(i)] = count
				i++
			}
			if nowper.State9Odds != "" {
				//判断产品ID是否小于10
				if nowper.ProductId < 10 {
					sup_encoding9 = "#0" + strconv.FormatInt(int64(nowper.ProductId), 10) + ">09"
				} else {
					sup_encoding9 = "#" + strconv.FormatInt(int64(nowper.ProductId), 10) + ">09"
				}
				sm.SupEncoding = sup_encoding9
				count := sm.GetSupCountByEncoding()
				number[strconv.Itoa(i)] = count
				i++
			}
			if nowper.State10Odds != "" {
				//判断产品ID是否小于10
				if nowper.ProductId < 10 {
					sup_encoding10 = "#0" + strconv.FormatInt(int64(nowper.ProductId), 10) + ">10"
				} else {
					sup_encoding10 = "#" + strconv.FormatInt(int64(nowper.ProductId), 10) + ">10"
				}
				sm.SupEncoding = sup_encoding10
				count := sm.GetSupCountByEncoding()
				number[strconv.Itoa(i)] = count
				i++
			}

			fmt.Println("number", number)
		}
	}

	m := make(map[string]interface{})
	m["list"] = list
	m["total"] = rows
	m["hotnum"] = number
	comm.Response(c, m)
}

//删除期数
func PeriodDel(c *gin.Context) {
	var per m.Periods
	per.Id, _ = strconv.ParseInt(c.PostForm("id"), 10, 64)
	fmt.Println("id", c.PostForm("id"))
	fmt.Println("aaad", c.PostForm("Id"))
	if c.PostForm("id") == "" {
		comm.ResponseError(c, 5021)
		return
	}

	if !per.GetPeriod() {
		comm.ResponseError(c, 5022)
		return
	}

	// 判断开始、结束、现在时间
	nowtime := time.Now()
	starttime, _ := time.Parse("2006-01-02 15:04:05", per.StartTime)
	endtime, _ := time.Parse("2006-01-02 15:04:05", per.EndTime)

	if starttime.Before(nowtime) && nowtime.Before(endtime) {
		comm.ResponseError(c, 5023)
		return
	}

	if endtime.Before(nowtime) {
		comm.ResponseError(c, 5024)
		return
	}

	if !per.PeriodDel() {
		comm.ResponseError(c, 5025)
		return
	}
	comm.Response(c, "删除成功")
}

//修改期数
func PeriodEndCodingUP(c *gin.Context) {
	var per m.Periods
	per.PeriodsId, _ = strconv.ParseInt(c.PostForm("PeriodsId"), 10, 64)
	if !per.GetPeriod() {
		comm.ResponseError(c, 5026)
		return
	}
	per.State = 1
	per.SupEncoding = c.PostForm("SupEncoding")
	if !per.PeriodUp() {
		comm.ResponseError(c, 5027)
	} else {
		//插入记录表
		uid, _ := strconv.ParseInt(c.PostForm("UID"), 10, 64)
		if uid == 0 {
			comm.ResponseError(c, 3162)
		}
		var supportRecord m.SupportRecord
		supportRecord.PerId = per.PeriodsId
		supportRecord.SubmitterId = uid
		supportRecord.Type = 1
		result := strings.Split(per.SupEncoding, "#")
		str := ""
		for mm, res := range result {
			if mm == 0 {
				continue
			}
			secResult := strings.Split(res, ">")
			fmt.Println("secResult", secResult)
			var (
				pro  []m.Product
				pro1 m.Product
			)
			pro1.Id, _ = strconv.ParseInt(secResult[0], 10, 64)
			pro, _ = pro1.GetProduct(pro1.Id)
			fmt.Println(pro, "++++", len(pro))
			if len(pro) > 0 {
				str += "#产品名称："
				str += pro[0].ProductName
				str += ",产品状态："
				if strings.EqualFold(secResult[1], "01") {
					str += pro[0].State1
				} else if strings.EqualFold(secResult[1], "02") {
					str += pro[0].State2
				} else if strings.EqualFold(secResult[1], "03") {
					str += pro[0].State3
				} else if strings.EqualFold(secResult[1], "04") {
					str += pro[0].State4
				} else if strings.EqualFold(secResult[1], "05") {
					str += pro[0].State5
				} else if strings.EqualFold(secResult[1], "06") {
					str += pro[0].State6
				} else if strings.EqualFold(secResult[1], "07") {
					str += pro[0].State7
				} else if strings.EqualFold(secResult[1], "08") {
					str += pro[0].State8
				}
			}
		}
		supportRecord.Result = str
		boo := supportRecord.SupportRecordAdd()
		if !boo {
			comm.ResponseError(c, 4049)
		} else {
			comm.Response(c, per)
		}
	}
}

//time 2016-11-25 txl
//修改期数
func PeriodUP(c *gin.Context) {
	var (
		per m.Periods
	)
	//获取客户端POST请求参数
	per.Id, _ = strconv.ParseInt(c.PostForm("myuid"), 10, 64)
	if !per.GetPeriod() {
		comm.ResponseError(c, 5026)
		return
	}

	// 判断开始、结束、现在时间
	nowtime := time.Now()

	updatestarttime, _ := time.Parse("2006-01-02 15:04:05", c.PostForm("StartTime")) //获取修改的开始时间
	updateendtime, _ := time.Parse("2006-01-02 15:04:05", c.PostForm("EndTime"))     //获取修改的结束时间

	if c.PostForm("StartTime") != "" {
		per.StartTime = c.PostForm("StartTime")
		if nowtime.Before(updatestarttime) {
			comm.ResponseError(c, 5028)
			return
		}
	}
	if c.PostForm("EndTime") != "" {
		per.EndTime = c.PostForm("EndTime")
		if updateendtime.Before(updatestarttime) {
			comm.ResponseError(c, 5028)
			return
		}
	}

	per.ATeam = c.PostForm("ATeam")
	per.BTeam = c.PostForm("BTeam")
	per.ProEncoding = c.PostForm("ProEncoding")

	if !per.PeriodUp() {
		comm.ResponseError(c, 5027)
		return
	}
	proen := strings.Split(per.ProEncoding, ",")
	for _, value := range proen {
		var (
			proper m.PeriodsProduct
		)

		proper.PeriodsId = per.PeriodsId
		proper.ProductId, _ = strconv.ParseInt(value, 10, 64)

		//判断是否存在过程表
		if !proper.GetPeriodsProducts() {
			fmt.Println("期数该产品不存在")
		}
		//热度
		proper.State1Hot, _ = strconv.ParseInt(c.PostForm(value+"State1Hot"), 10, 64)
		proper.State2Hot, _ = strconv.ParseInt(c.PostForm(value+"State2Hot"), 10, 64)
		proper.State3Hot, _ = strconv.ParseInt(c.PostForm(value+"State3Hot"), 10, 64)
		proper.State4Hot, _ = strconv.ParseInt(c.PostForm(value+"State4Hot"), 10, 64)
		proper.State5Hot, _ = strconv.ParseInt(c.PostForm(value+"State5Hot"), 10, 64)
		proper.State6Hot, _ = strconv.ParseInt(c.PostForm(value+"State6Hot"), 10, 64)
		proper.State7Hot, _ = strconv.ParseInt(c.PostForm(value+"State7Hot"), 10, 64)
		proper.State8Hot, _ = strconv.ParseInt(c.PostForm(value+"State8Hot"), 10, 64)
		proper.State9Hot, _ = strconv.ParseInt(c.PostForm(value+"State9Hot"), 10, 64)
		proper.State10Hot, _ = strconv.ParseInt(c.PostForm(value+"State10Hot"), 10, 64)
		//赔率
		proper.State1Odds = c.PostForm(value + "State1Odds")
		proper.State2Odds = c.PostForm(value + "State2Odds")
		proper.State3Odds = c.PostForm(value + "State3Odds")
		proper.State4Odds = c.PostForm(value + "State4Odds")
		proper.State5Odds = c.PostForm(value + "State5Odds")
		proper.State6Odds = c.PostForm(value + "State6Odds")
		proper.State7Odds = c.PostForm(value + "State7Odds")
		proper.State8Odds = c.PostForm(value + "State8Odds")
		proper.State9Odds = c.PostForm(value + "State9Odds")
		proper.State10Odds = c.PostForm(value + "State10Odds")
		_, err := proper.PeriodsProductUp()

		if err != nil {
			fmt.Println("热度修改失败", err)
		}
	}
	comm.Response(c, "修改成功")

}

//txl-del
//获取期数信息
func GetPeriod(c *gin.Context) {
	var per m.Periods
	per.PeriodsId, _ = strconv.ParseInt(c.PostForm("PeriodsId"), 10, 64)
	if !per.GetPeriod() {
		comm.ResponseError(c, 5026)
		return
	}
	comm.Response(c, per)
}

//获取期数列表
func GetPeriodList(c *gin.Context) {
	//	var per m.Periods
	var per m.PeriodsRoom
	page, _ := strconv.Atoi(c.PostForm("page"))
	rows, _ := strconv.Atoi(c.PostForm("rows"))
	state := c.PostForm("state")
	inputid := c.PostForm("inputid")
	nowdata := c.PostForm("nowdata")
	list, row := per.GetPeriodList(page, rows, state, inputid, nowdata)
	fmt.Println("row", row)
	if row == 0 {
		c.JSON(200, gin.H{"total": row, "rows": []int{}})
		return
	}
	c.JSON(200, gin.H{"total": row, "rows": list})
}

//txl-del
//获取当前期数所有的产品
func GetPerProList(c *gin.Context) {
	var perpro m.PeriodsProduct
	page, _ := strconv.Atoi(c.PostForm("page"))
	rows, _ := strconv.Atoi(c.PostForm("rows"))
	perpro.PeriodsId, _ = strconv.ParseInt(c.PostForm("PerProid"), 10, 64)
	list, row := perpro.GetNowPerProList(page, rows)
	if row == 0 {
		c.JSON(200, gin.H{"message": true, "date": ""})
	}
	c.JSON(200, gin.H{"message": true, "date": list})
}

//txl-del
func GetPeriodIdProductName(c *gin.Context) {
	PeriodsId, _ := strconv.ParseInt(c.PostForm("PeriodId"), 10, 64)
	list, row := m.GetPeriodIdProductName(PeriodsId)
	if row <= 0 {
		c.JSON(200, gin.H{"state": false, "date": "", "row": row})
		return
	}
	c.JSON(200, gin.H{"state": true, "date": list, "row": row})
}

//获取当前期的产品ID和对应名称 (即将开始的一期)
func CurrentPeriodBase(c *gin.Context) {
	roomId := c.PostForm("roomId")
	res, indexs, ATeam, BTeam := CurrentPeriodBase_com(roomId)
	time := GetDateTime()
	m := make(map[string]interface{})
	m["ATeam"] = ATeam
	m["BTeam"] = BTeam
	m["rows"] = indexs
	m["data"] = res
	m["nowtime"] = time
	comm.Response(c, m)
}

func CurrentPeriodBase_com(roomId string) ([]m.WebProductShow, int, string, string) {
	res, indexs, ATeam, BTeam := m.GetTimePerProList(roomId)
	if indexs <= 0 {
	}
	start_time := GetDateTime()
	st, _ := time.Parse("2006-01-02 15:04:05", start_time)
	et, _ := time.Parse("2006-01-02 15:04:05", res[0].EndTime)
	u_st := st.Unix()
	u_et := et.Unix()
	timess := u_et - u_st
	fmt.Println("time", timess)
	m := make(map[string]interface{})
	m["ATeam"] = ATeam
	m["BTeam"] = BTeam
	m["rows"] = indexs
	m["data"] = res
	m["time"] = timess
	return res, indexs, ATeam, BTeam
}

//获取当前期产品明细
func CurrentPeriodDetails(c *gin.Context) {
	perid := c.PostForm("PeriodId")
	proid := c.PostForm("ProductId")
	res, err := m.CurrentPeriodDetails(perid, proid)
	if err != nil {
		comm.ResponseError(c, 5029)
		return
	}
	comm.Response(c, res)
}

//获取更多期
func PeriodMore(c *gin.Context) {
	var all m.PeriodsAllContent
	nowtime := time.Now().Local()
	list, rows := all.GetMorePeriodList(nowtime.String())
	if rows <= 0 {
		c.JSON(200, gin.H{"state": false, "rows": rows, "date": ""})
		return
	}
	c.JSON(200, gin.H{"state": false, "rows": rows, "date": list})
}

//Phone端更多期  A  VS  B
func PeriodPhoneMore(c *gin.Context) {
	var per m.Periods
	nowtime := time.Now().Local()
	list, rows := per.GetPeriodsPhone(nowtime.String())
	if rows <= 0 {
		c.JSON(200, gin.H{"state": false, "rows": rows, "date": "No More Periods"})
		return
	}
	c.JSON(200, gin.H{"state": true, "rows": rows, "date": list})

}

func PerProName(c *gin.Context) {
	var per m.Periods
	per.PeriodsId, _ = strconv.ParseInt(c.PostForm("PeriodsId"), 10, 64)
	perid := c.PostForm("PeriodsId")
	if !per.GetPeriod() {
		c.JSON(200, gin.H{"state": false, "data": "", "message": "期数不存在"})
	}
	s := strings.Split(per.ProEncoding, ",")
	list := m.GetPerProName(perid, s)
	if len(list) <= 0 {
		c.JSON(200, gin.H{"state": false, "data": "", "message": "查询错误"})
	}
	c.JSON(200, gin.H{"state": true, "data": list, "message": "查询成功"})
}

//期数核算
func PerAccounting(c *gin.Context) {
	var (
		per m.Periods
		smt m.SupportManagementResult
	)
	per.PeriodsId, _ = strconv.ParseInt(c.PostForm("PeriodsId"), 10, 64)
	if per.PeriodsId == 0 {
		comm.ResponseError(c, 2000)
		return
	}
	smt.PeriodsId = per.PeriodsId
	if smt.IsHere() {
		comm.ResponseError(c, 5030)
		return
	}

	if !per.GetPeriod() {
		comm.ResponseError(c, 5036)
		return
	}

	if per.SupEncoding == "" {
		comm.ResponseError(c, 5037)
		return
	}

	flag, _ := per.PerAccounting(per.StartTime)
	if !flag {
		comm.ResponseError(c, 5031)
		return
	}
	flag, _ = m.UpdateInfo(per.PeriodsId)
	if !flag {
		comm.ResponseError(c, 5032)
		return
	}
	// 添加核算记录
	uid, _ := strconv.ParseInt(c.PostForm("UID"), 10, 64)
	if uid == 0 {
		comm.ResponseError(c, 3162)
	}
	var supportRecord m.SupportRecord
	supportRecord.PerId = per.PeriodsId
	supportRecord.SubmitterId = uid
	supportRecord.Type = 2
	result := strings.Split(per.SupEncoding, "#")
	str := ""
	for mm, res := range result {
		if mm == 0 {
			continue
		}
		secResult := strings.Split(res, ">")
		var (
			pro  []m.Product
			pro1 m.Product
		)
		pro1.Id, _ = strconv.ParseInt(secResult[0], 10, 64)
		pro, _ = pro1.GetProduct(pro1.Id)
		if len(pro) > 0 {
			str += "#产品名称："
			str += pro[0].ProductName
			str += ",产品状态："
			if strings.EqualFold(secResult[1], "01") {
				str += pro[0].State1
			} else if strings.EqualFold(secResult[1], "02") {
				str += pro[0].State2
			} else if strings.EqualFold(secResult[1], "03") {
				str += pro[0].State3
			} else if strings.EqualFold(secResult[1], "04") {
				str += pro[0].State4
			} else if strings.EqualFold(secResult[1], "05") {
				str += pro[0].State5
			} else if strings.EqualFold(secResult[1], "06") {
				str += pro[0].State6
			} else if strings.EqualFold(secResult[1], "07") {
				str += pro[0].State7
			} else if strings.EqualFold(secResult[1], "08") {
				str += pro[0].State8
			}
		}
	}
	supportRecord.Result = str
	boo := supportRecord.SupportRecordAdd()
	if !boo {
		comm.ResponseError(c, 4049)
	} else {
		comm.Response(c, "核算完成")
	}
}

//返回当前期收益详情
type EarningDetails struct {
	//PeriodId            int64  //期Id
	ProductId            string  //产品Id
	ProductName          string  //产品名
	State1               string  //状态1
	State1SupmanCount    int64   //状态1的支持人数
	State1SuptotalCount  float64 // 状态1的支持总石榴数
	State2               string  //状态2
	State2SupmanCount    int64   //状态2的支持人数
	State2SuptotalCount  float64 // 状态2的支持总石榴数
	State3               string  //状态3
	State3SupmanCount    int64   //状态3的支持人数
	State3SuptotalCount  float64 // 状态3的支持总石榴数
	State4               string  //状态4
	State4SupmanCount    int64   //状态4的支持人数
	State4SuptotalCount  float64 // 状态4的支持总石榴数
	State5               string  //状态5
	State5SupmanCount    int64   //状态5的支持人数
	State5SuptotalCount  float64 // 状态5的支持总石榴数
	State6               string  //状态6
	State6SupmanCount    int64   //状态6的支持人数
	State6SuptotalCount  float64 // 状态6的支持总石榴数
	State7               string  //状态7
	State7SupmanCount    int64   //状态7的支持人数
	State7SuptotalCount  float64 // 状态7的支持总石榴数
	State8               string  //状态8
	State8SupmanCount    int64   //状态8的支持人数
	State8SuptotalCount  float64 // 状态8的支持总石榴数
	State9               string  //状态9
	State9SupmanCount    int64   //状态9的支持人数
	State9SuptotalCount  float64 // 状态9的支持总石榴数
	State10              string  //状态10
	State10SupmanCount   int64   //状态10的支持人数
	State10SuptotalCount float64 // 状态10的支持总石榴数
	SupManCount          int64   //支持总人数
	SupTotalCount        float64 //支持总石榴数
	EarningsCount        float64 //当前产品收益
}

//期数管理 期数收益详情 cnxulin
func EarningsDetails(c *gin.Context) {

	var (
		per                 m.Periods
		ed                  EarningDetails
		peroidtotalearnings float64 //当前期的总盈利
	)

	periodsid, _ := strconv.ParseInt(c.PostForm("PeriodsId"), 10, 64)
	productid := c.PostForm("ProductId")
	fmt.Println("periodsid", periodsid)
	fmt.Println("productid", productid)

	if c.PostForm("PeriodsId") == "" {
		comm.ResponseError(c, 5033)
		return
	}
	per.PeriodsId = periodsid
	if !per.GetPeriod() {
		comm.ResponseError(c, 5034)
		return
	}

	detaillist := make(map[string]EarningDetails)

	//字符串分割
	productids := strings.Split(productid, ",")
	for i, num := range productids {
		var (
			nowper     m.PeriodsProduct
			sm         m.SupportManagement     //支持管理
			product    m.Product               //产品
			iswinCount float64             = 0 //所有用户赢得石榴数
		)
		nowper.PeriodsId = periodsid
		nowper.ProductId, _ = strconv.ParseInt(num, 10, 64)

		nowper.GetPeriodsProducts()

		sm.PeriodsId = nowper.PeriodsId
		ed.ProductId = num //产品ID

		//根据产品id查询产品表
		product.Id, _ = strconv.ParseInt(num, 10, 64)
		if !product.GetProducts() {
			fmt.Println("产品不存在")
		}
		ed.ProductName = product.ProductName //产品名

		number := make(map[string]m.Number)

		//判断产品状态值是否为空
		if nowper.State1Odds != " " {
			//判断产品ID是否小于10
			if nowper.ProductId < 10 {
				sm.SupEncoding = "#0" + strconv.FormatInt(int64(nowper.ProductId), 10) + ">01"
			} else {
				sm.SupEncoding = "#" + strconv.FormatInt(int64(nowper.ProductId), 10) + ">01"
			}

			list, count := sm.GetSupporByEncoding() // 获取支持数与赔率记录

			number["0"] = count

			ed.State1 = product.State1                                                    //状态1名称
			ed.State1SupmanCount = int64(len(list))                                       //状态1支持人数
			ed.State1SuptotalCount, _ = strconv.ParseFloat(number["0"].SupportNumber, 10) //状态1的支持总石榴数

			for i := 0; i < len(list); i++ {
				if strings.Contains(per.SupEncoding, list[i].SupEncoding) { //赢 按赔率算出赢得数量
					arr := strings.Split(list[i].Odds, ":")
					k, _ := strconv.ParseFloat(arr[0], 64)
					v, _ := strconv.ParseFloat(arr[1], 64)
					modds := v / k

					iswinCount = iswinCount + float64(list[i].SupporNumber)*modds

				}
			}

		}

		if nowper.State2Odds != " " {
			//判断产品ID是否小于10
			if nowper.ProductId < 10 {
				sm.SupEncoding = "#0" + strconv.FormatInt(int64(nowper.ProductId), 10) + ">02"
			} else {
				sm.SupEncoding = "#" + strconv.FormatInt(int64(nowper.ProductId), 10) + ">02"
			}

			list, count := sm.GetSupporByEncoding() // 获取支持数与赔率记录

			number["1"] = count

			ed.State2 = product.State2                                                    //状态1名称
			ed.State2SupmanCount = int64(len(list))                                       //状态1支持人数
			ed.State2SuptotalCount, _ = strconv.ParseFloat(number["1"].SupportNumber, 10) //状态1的支持总石榴数

			for i := 0; i < len(list); i++ {
				if strings.Contains(per.SupEncoding, list[i].SupEncoding) { //赢 按赔率算出赢得数量
					arr := strings.Split(list[i].Odds, ":")
					k, _ := strconv.ParseFloat(arr[0], 64)
					v, _ := strconv.ParseFloat(arr[1], 64)
					modds := v / k

					iswinCount = iswinCount + float64(list[i].SupporNumber)*modds

				}
			}

		}

		if nowper.State3Odds != " " {
			//判断产品ID是否小于10
			if nowper.ProductId < 10 {
				sm.SupEncoding = "#0" + strconv.FormatInt(int64(nowper.ProductId), 10) + ">03"
			} else {
				sm.SupEncoding = "#" + strconv.FormatInt(int64(nowper.ProductId), 10) + ">03"
			}

			list, count := sm.GetSupporByEncoding() // 获取支持数与赔率记录

			number["2"] = count

			ed.State3 = product.State3                                                    //状态1名称
			ed.State3SupmanCount = int64(len(list))                                       //状态1支持人数
			ed.State3SuptotalCount, _ = strconv.ParseFloat(number["2"].SupportNumber, 10) //状态1的支持总石榴数

			for i := 0; i < len(list); i++ {
				if strings.Contains(per.SupEncoding, list[i].SupEncoding) { //赢 按赔率算出赢得数量
					arr := strings.Split(list[i].Odds, ":")
					k, _ := strconv.ParseFloat(arr[0], 64)
					v, _ := strconv.ParseFloat(arr[1], 64)
					modds := v / k

					iswinCount = iswinCount + float64(list[i].SupporNumber)*modds

				}
			}
		}

		if nowper.State4Odds != " " {
			//判断产品ID是否小于10
			if nowper.ProductId < 10 {
				sm.SupEncoding = "#0" + strconv.FormatInt(int64(nowper.ProductId), 10) + ">04"
			} else {
				sm.SupEncoding = "#" + strconv.FormatInt(int64(nowper.ProductId), 10) + ">04"
			}

			list, count := sm.GetSupporByEncoding() // 获取支持数与赔率记录

			number["3"] = count

			ed.State4 = product.State4                                                    //状态1名称
			ed.State4SupmanCount = int64(len(list))                                       //状态1支持人数
			ed.State4SuptotalCount, _ = strconv.ParseFloat(number["3"].SupportNumber, 10) //状态1的支持总石榴数

			for i := 0; i < len(list); i++ {
				if strings.Contains(per.SupEncoding, list[i].SupEncoding) { //赢 按赔率算出赢得数量
					arr := strings.Split(list[i].Odds, ":")
					k, _ := strconv.ParseFloat(arr[0], 64)
					v, _ := strconv.ParseFloat(arr[1], 64)
					modds := v / k

					iswinCount = iswinCount + float64(list[i].SupporNumber)*modds

				}
			}
		}

		if nowper.State5Odds != " " {
			//判断产品ID是否小于10
			if nowper.ProductId < 10 {
				sm.SupEncoding = "#0" + strconv.FormatInt(int64(nowper.ProductId), 10) + ">05"
			} else {
				sm.SupEncoding = "#" + strconv.FormatInt(int64(nowper.ProductId), 10) + ">05"
			}

			list, count := sm.GetSupporByEncoding() // 获取支持数与赔率记录

			number["4"] = count

			ed.State5 = product.State5                                                    //状态1名称
			ed.State5SupmanCount = int64(len(list))                                       //状态1支持人数
			ed.State5SuptotalCount, _ = strconv.ParseFloat(number["4"].SupportNumber, 10) //状态1的支持总石榴数

			for i := 0; i < len(list); i++ {
				if strings.Contains(per.SupEncoding, list[i].SupEncoding) { //赢 按赔率算出赢得数量
					arr := strings.Split(list[i].Odds, ":")
					k, _ := strconv.ParseFloat(arr[0], 64)
					v, _ := strconv.ParseFloat(arr[1], 64)
					modds := v / k

					iswinCount = iswinCount + float64(list[i].SupporNumber)*modds

				}
			}
		}
		if nowper.State6Odds != " " {
			//判断产品ID是否小于10
			if nowper.ProductId < 10 {
				sm.SupEncoding = "#0" + strconv.FormatInt(int64(nowper.ProductId), 10) + ">06"
			} else {
				sm.SupEncoding = "#" + strconv.FormatInt(int64(nowper.ProductId), 10) + ">06"
			}

			list, count := sm.GetSupporByEncoding() // 获取支持数与赔率记录

			number["5"] = count

			ed.State6 = product.State6
			ed.State6SupmanCount = int64(len(list))
			ed.State6SuptotalCount, _ = strconv.ParseFloat(number["5"].SupportNumber, 10)

			for i := 0; i < len(list); i++ {
				if strings.Contains(per.SupEncoding, list[i].SupEncoding) { //赢 按赔率算出赢得数量
					arr := strings.Split(list[i].Odds, ":")
					k, _ := strconv.ParseFloat(arr[0], 64)
					v, _ := strconv.ParseFloat(arr[1], 64)
					modds := v / k

					iswinCount = iswinCount + float64(list[i].SupporNumber)*modds

				}
			}
		}
		if nowper.State7Odds != " " {
			//判断产品ID是否小于10
			if nowper.ProductId < 10 {
				sm.SupEncoding = "#0" + strconv.FormatInt(int64(nowper.ProductId), 10) + ">07"
			} else {
				sm.SupEncoding = "#" + strconv.FormatInt(int64(nowper.ProductId), 10) + ">07"
			}

			list, count := sm.GetSupporByEncoding() // 获取支持数与赔率记录

			number["6"] = count

			ed.State7 = product.State7
			ed.State7SupmanCount = int64(len(list))
			ed.State7SuptotalCount, _ = strconv.ParseFloat(number["6"].SupportNumber, 10)

			for i := 0; i < len(list); i++ {
				if strings.Contains(per.SupEncoding, list[i].SupEncoding) { //赢 按赔率算出赢得数量
					arr := strings.Split(list[i].Odds, ":")
					k, _ := strconv.ParseFloat(arr[0], 64)
					v, _ := strconv.ParseFloat(arr[1], 64)
					modds := v / k

					iswinCount = iswinCount + float64(list[i].SupporNumber)*modds

				}
			}
		}
		if nowper.State8Odds != " " {
			//判断产品ID是否小于10
			if nowper.ProductId < 10 {
				sm.SupEncoding = "#0" + strconv.FormatInt(int64(nowper.ProductId), 10) + ">08"
			} else {
				sm.SupEncoding = "#" + strconv.FormatInt(int64(nowper.ProductId), 10) + ">08"
			}

			list, count := sm.GetSupporByEncoding() // 获取支持数与赔率记录

			number["7"] = count

			ed.State8 = product.State8
			ed.State8SupmanCount = int64(len(list))
			ed.State8SuptotalCount, _ = strconv.ParseFloat(number["7"].SupportNumber, 10)

			for i := 0; i < len(list); i++ {
				if strings.Contains(per.SupEncoding, list[i].SupEncoding) { //赢 按赔率算出赢得数量
					arr := strings.Split(list[i].Odds, ":")
					k, _ := strconv.ParseFloat(arr[0], 64)
					v, _ := strconv.ParseFloat(arr[1], 64)
					modds := v / k

					iswinCount = iswinCount + float64(list[i].SupporNumber)*modds

				}
			}
		}
		if nowper.State9Odds != " " {
			//判断产品ID是否小于10
			if nowper.ProductId < 10 {
				sm.SupEncoding = "#0" + strconv.FormatInt(int64(nowper.ProductId), 10) + ">09"
			} else {
				sm.SupEncoding = "#" + strconv.FormatInt(int64(nowper.ProductId), 10) + ">09"
			}

			list, count := sm.GetSupporByEncoding() // 获取支持数与赔率记录

			number["8"] = count

			ed.State9 = product.State9
			ed.State9SupmanCount = int64(len(list))
			ed.State9SuptotalCount, _ = strconv.ParseFloat(number["8"].SupportNumber, 10)

			for i := 0; i < len(list); i++ {
				if strings.Contains(per.SupEncoding, list[i].SupEncoding) { //赢 按赔率算出赢得数量
					arr := strings.Split(list[i].Odds, ":")
					k, _ := strconv.ParseFloat(arr[0], 64)
					v, _ := strconv.ParseFloat(arr[1], 64)
					modds := v / k

					iswinCount = iswinCount + float64(list[i].SupporNumber)*modds

				}
			}
		}
		if nowper.State10Odds != " " {
			//判断产品ID是否小于10
			if nowper.ProductId < 10 {
				sm.SupEncoding = "#0" + strconv.FormatInt(int64(nowper.ProductId), 10) + ">10"
			} else {
				sm.SupEncoding = "#" + strconv.FormatInt(int64(nowper.ProductId), 10) + ">10"
			}

			list, count := sm.GetSupporByEncoding() // 获取支持数与赔率记录

			number["9"] = count

			ed.State10 = product.State10
			ed.State10SupmanCount = int64(len(list))
			ed.State10SuptotalCount, _ = strconv.ParseFloat(number["9"].SupportNumber, 10)

			for i := 0; i < len(list); i++ {
				if strings.Contains(per.SupEncoding, list[i].SupEncoding) { //赢 按赔率算出赢得数量
					arr := strings.Split(list[i].Odds, ":")
					k, _ := strconv.ParseFloat(arr[0], 64)
					v, _ := strconv.ParseFloat(arr[1], 64)
					modds := v / k

					iswinCount = iswinCount + float64(list[i].SupporNumber)*modds

				}
			}
		}
		//		}
		//当前产品的支持人数
		ed.SupManCount = ed.State1SupmanCount + ed.State2SupmanCount + ed.State3SupmanCount + ed.State4SupmanCount + ed.State5SupmanCount
		//当前产品的支持总石榴数
		ed.SupTotalCount = ed.State1SuptotalCount + ed.State2SuptotalCount + ed.State3SuptotalCount + ed.State4SuptotalCount + ed.State5SuptotalCount
		//当前产品收益 = 当前产品的支持总石榴数 - 用户支持获胜数
		ed.EarningsCount = ed.SupTotalCount - iswinCount
		detaillist[strconv.Itoa(i)] = ed

		peroidtotalearnings += ed.EarningsCount
		fmt.Println("数据", ed)
	}
	m := make(map[string]interface{})
	m["detaillist"] = detaillist
	m["periodId"] = periodsid
	m["peroidtotalearnings"] = peroidtotalearnings
	comm.Response(c, m)

}

//查询资金流向列表
func GetFundAccounting(c *gin.Context) {
	var cp m.FundAccounting
	page, _ := strconv.Atoi(c.PostForm("page"))
	rows, _ := strconv.Atoi(c.PostForm("rows"))
	inputid := c.PostForm("inputid")
	if inputid != "" {
		list, row, _ := cp.GetFunAccountingByIdOrUserId(inputid, page, rows)
		if row == 0 {
			c.JSON(200, gin.H{"total": 0, "rows": ""})
			return
		}
		c.JSON(200, gin.H{"total": row, "rows": list})
		return
	}
	list, row := cp.GetFundAccounting(page, rows)
	if row == 0 {
		c.JSON(200, gin.H{"total": 0, "rows": ""})
		return
	}
	c.JSON(200, gin.H{"total": row, "rows": list})
}

//后台提交核算记录@liuhan 170314
func GetSupportRecord(c *gin.Context) {
	var sr m.SupportRecord
	page, _ := strconv.Atoi(c.PostForm("page"))
	rows, _ := strconv.Atoi(c.PostForm("rows"))
	inputid := c.PostForm("inputid")
	list, row := sr.GetSupportRecord(page, rows, inputid)
	if row == 0 {
		c.JSON(200, gin.H{"total": 0, "rows": ""})
		return
	}
	c.JSON(200, gin.H{"total": row, "rows": list})
}
