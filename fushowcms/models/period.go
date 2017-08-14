package models

/*
 * Copyright (c) 2016—2017 www.fushow.cn, All rights reserved.
 */
import (
	"bytes"
	"strconv"
	"strings"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

//增加期数
func (per *Periods) PeriodAdd() bool {
	rows, err := Engine.Insert(per)
	if rows > 0 && err == nil {
		return true
	}
	return false
}

//删除期数
func (per *Periods) PeriodDel() bool {
	rows, err := Engine.Delete(per)
	if rows > 0 && err == nil {
		return true
	}
	return false
}

//通过期Id删除
func (per *Periods) PeriodDelForPid() bool {
	rows, err := Engine.Where("periods_id=?", per.PeriodsId).Delete(per)
	if rows > 0 && err == nil {
		return true
	}
	return false
}

//修改期数信息
func (per *Periods) PeriodUp() bool {
	rows, err := Engine.Where("Id=?", per.Id).Update(per)
	if rows > 0 && err == nil {
		return true
	}
	return false
}

//获取期数信息
func (per *Periods) GetPeriod() bool {
	flag, _ := Engine.Get(per)
	if flag {
		return true
	}
	return false
}

//获取当前期数所有产品---竞猜产品
func (perpro *PeriodsProduct) GetNowPerProList(page, rows int) ([]PeriodsProduct, int64) {
	var list []PeriodsProduct
	total, _ := Engine.Where("periods_id =?", perpro.PeriodsId).Count(perpro)
	Engine.Where("periods_id =?", perpro.PeriodsId).Limit(rows, (page-1)*rows).Find(&list)
	return list, total
}

//web端产品展示
type WebProductShow struct {
	PeriodsId   string
	ProductId   string
	ProductName string
	StartTime   string
	EndTime     string
}

func GetPeriodIdProductName(PeriodsId int64) ([]WebProductShow, int) {
	var (
		wps          WebProductShow
		wproductshow []WebProductShow
	)

	sql := "SELECT x.periods_id,x.product_id,y.product_name FROM (SELECT b.* FROM (SELECT periods_id FROM periods where periods_id=? ORDER BY id ASC LIMIT 1) as a left join periods_product as b on a.periods_id = b.periods_id) as x INNER JOIN product as y on x.product_id = y.id"
	results, _ := Engine.Query(sql, PeriodsId)
	for _, value := range results {
		wps.PeriodsId = string(value["periods_id"])
		wps.ProductId = string(value["product_id"])
		wps.ProductName = string(value["product_name"])
		wproductshow = append(wproductshow, wps)
	}
	return wproductshow, len(results)
}

//WEB获取当前期数（即将开始的第一场比赛）
func GetTimePerProList(roomId string) ([]WebProductShow, int, string, string) {
	var (
		wproductshow []WebProductShow
		wps          WebProductShow
		ATeam        string
		BTeam        string
	)
	Engine.ShowSQL(true)
	b := bytes.Buffer{}
	b.WriteString("SELECT x.periods_id,x.product_id,y.product_name,a_team,b_team,x.start_time,x.end_time FROM (SELECT b.periods_id,b.product_id,a_team,b_team,a.start_time,a.end_time FROM (SELECT periods_id,a_team,b_team,start_time,end_time FROM periods where room_id REGEXP ")
	b.WriteString("? AND state=0 AND end_time > DATE_FORMAT(NOW() AND del_time IS NOT NULL, '%Y-%c-%d %H:%i:%s') ")
	b.WriteString(" ORDER BY start_time DESC limit 1) as a left join periods_product as b on a.periods_id = b.periods_id) as x INNER JOIN product as y on x.product_id = y.id")
	ss := b.String()
	results, _ := Engine.Query(ss, "#"+roomId+"#|#"+roomId+"$")

	for _, value := range results {
		wps.PeriodsId = string(value["periods_id"])
		wps.ProductId = string(value["product_id"])
		wps.ProductName = string(value["product_name"])
		wps.StartTime = string(value["start_time"])
		wps.EndTime = string(value["end_time"])
		ATeam = string(value["a_team"])
		BTeam = string(value["b_team"])
		wproductshow = append(wproductshow, wps)
	}
	return wproductshow, len(results), ATeam, BTeam
}

//2 WEB端详细产品热度赔率投注数
type WebProductShowContent struct {
	PeriodsId   string
	ProductId   string
	ProductName string
	State1      string
	State2      string
	State3      string
	State4      string
	State5      string
	State6      string
	State7      string
	State8      string
	State9      string
	State10     string
	State1Hot   string
	State2Hot   string
	State3Hot   string
	State4Hot   string
	State5Hot   string
	State6Hot   string
	State7Hot   string
	State8Hot   string
	State9Hot   string
	State10Hot  string
	State1Odds  string
	State2Odds  string
	State3Odds  string
	State4Odds  string
	State5Odds  string
	State6Odds  string
	State7Odds  string
	State8Odds  string
	State9Odds  string
	State10Odds string
}

func GetPerProName(pid string, proid []string) []WebProductShowContent {
	var wps []WebProductShowContent
	sql := "SELECT * FROM (SELECT * FROM periods_product WHERE periods_id = ? AND product_id = ?) AS A INNER JOIN product AS B ON A.product_id = B.id"
	for i := 0; i < len(proid); i++ {
		var web WebProductShowContent
		results, _ := Engine.Query(sql, pid, proid[i])
		for _, value := range results {
			web.PeriodsId = string(value["periods_id"])
			web.ProductId = string(value["product_id"])
			web.ProductName = string(value["product_name"])
			web.State1 = string(value["state1"])
			web.State2 = string(value["state2"])
			web.State3 = string(value["state3"])
			web.State4 = string(value["state4"])
			web.State5 = string(value["state5"])
			web.State6 = string(value["state6"])
			web.State7 = string(value["state7"])
			web.State8 = string(value["state8"])
			web.State1Hot = string(value["state1_hot"])
			web.State2Hot = string(value["state2_hot"])
			web.State3Hot = string(value["state3_hot"])
			web.State4Hot = string(value["state4_hot"])
			web.State5Hot = string(value["state5_hot"])
			web.State6Hot = string(value["state6_hot"])
			web.State7Hot = string(value["state7_hot"])
			web.State8Hot = string(value["state8_hot"])
			web.State1Odds = string(value["state1_odds"])
			web.State2Odds = string(value["state2_odds"])
			web.State3Odds = string(value["state3_odds"])
			web.State4Odds = string(value["state4_odds"])
			web.State5Odds = string(value["state5_odds"])
			web.State6Odds = string(value["state6_odds"])
			web.State7Odds = string(value["state7_odds"])
			web.State8Odds = string(value["state8_odds"])
		}
		wps = append(wps, web)
	}
	return wps
}

//获取当前最近一起的比赛
func CurrentPeriodDetails(pid, proid string) (WebProductShowContent, error) {
	var wps WebProductShowContent
	sql := "SELECT * FROM (SELECT * FROM periods_product WHERE periods_id = ? AND product_id = ?) AS A INNER JOIN product AS B ON A.product_id = B.id"
	results, err := Engine.Query(sql, pid, proid)
	for _, value := range results {
		wps.PeriodsId = string(value["periods_id"])
		wps.ProductId = string(value["product_id"])
		wps.ProductName = string(value["product_name"])
		wps.State1 = string(value["state1"])
		wps.State2 = string(value["state2"])
		wps.State3 = string(value["state3"])
		wps.State4 = string(value["state4"])
		wps.State5 = string(value["state5"])
		wps.State6 = string(value["state6"])
		wps.State7 = string(value["state7"])
		wps.State8 = string(value["state8"])
		wps.State9 = string(value["state9"])
		wps.State10 = string(value["state10"])
		wps.State1Hot = string(value["state1_hot"])
		wps.State2Hot = string(value["state2_hot"])
		wps.State3Hot = string(value["state3_hot"])
		wps.State4Hot = string(value["state4_hot"])
		wps.State5Hot = string(value["state5_hot"])
		wps.State6Hot = string(value["state6_hot"])
		wps.State7Hot = string(value["state7_hot"])
		wps.State8Hot = string(value["state8_hot"])
		wps.State9Hot = string(value["state9_hot"])
		wps.State10Hot = string(value["state10_hot"])
		wps.State1Odds = string(value["state1_odds"])
		wps.State2Odds = string(value["state2_odds"])
		wps.State3Odds = string(value["state3_odds"])
		wps.State4Odds = string(value["state4_odds"])
		wps.State5Odds = string(value["state5_odds"])
		wps.State6Odds = string(value["state6_odds"])
		wps.State7Odds = string(value["state7_odds"])
		wps.State8Odds = string(value["state8_odds"])
		wps.State9Odds = string(value["state9_odds"])
		wps.State10Odds = string(value["state10_odds"])
	}
	return wps, err
}

//获得更多期数列表WEB
//2016-9-6
func (per *PeriodsAllContent) GetMorePeriodList(nowtime string) ([]PeriodsAllContent, int64) {
	var list []PeriodsAllContent
	total, _ := Engine.Where("id>?", 0).Count(per)
	Engine.Where("id>?", 0).Find(&list)
	return list, total

}

/*
*功能:查询期数信息列表---后台 （修改）
*修改原因：后台期数管理增加房间别名
*@cnxulin
*日期:20161114
 */
func (per *Periods) GetPeriodList(page, rows int, state, input, nowdata string) ([]PeriodsRoom, int64) {
	var (
		list    PeriodsRoom
		now     Periods
		lists   []PeriodsRoom
		sql     string
		results []map[string][]byte
	)
	total, _ := Engine.Table("periods").Where("id>0").Count(&now)
	if state != "" {
		sql = "SELECT p.*,r.room_alias FROM periods p  INNER JOIN anchor_room r WHERE substring_index(p.room_id,'#',-1) = r.id AND p.del_time is null AND p.state = 0 order by p.start_time desc Limit ?,?"
		results, _ = Engine.Query(sql, (page-1)*rows, rows)
	} else {
		sql = "SELECT p.*,r.room_alias FROM periods p  INNER JOIN anchor_room r WHERE substring_index(p.room_id,'#',-1) = r.id AND p.del_time is null order by p.start_time desc Limit ?,?"
		results, _ = Engine.Query(sql, (page-1)*rows, rows)
	}
	if input != "" {
		sql = "SELECT p.*,r.room_alias FROM periods p  INNER JOIN anchor_room r WHERE substring_index(p.room_id,'#',-1) = r.id AND p.del_time is null AND r.id like ? order by p.start_time desc Limit ?,?"
		results, _ = Engine.Query(sql, "%"+input+"%", (page-1)*rows, rows)
	}
	if nowdata != "" {
		sql = "SELECT p.*,r.room_alias FROM periods p  INNER JOIN anchor_room r WHERE substring_index(p.room_id,'#',-1) = r.id AND p.del_time is null AND p.state=1 order by p.start_time desc Limit ?,?"
		results, _ = Engine.Query(sql, (page-1)*rows, rows)
	}
	//截取#号后面的数字（房间ID）根据房间ID与期数管理的房间ID来查询期数信息与房间别名(rows, (page-1)*rows)
	for _, value := range results {
		id, _ := strconv.ParseInt(string(value["id"]), 10, 64)
		periods, _ := strconv.ParseInt(string(value["periods_id"]), 10, 64)
		state, _ := strconv.ParseInt(string(value["state"]), 10, 64)
		list.Id = id
		list.PeriodsId = periods
		list.StartTime = string(value["start_time"])
		list.EndTime = string(value["end_time"])
		list.RoomId = string(value["room_id"])
		list.ATeam = string(value["a_team"])
		list.BTeam = string(value["b_team"])
		list.State = state
		list.SubmitTime = string(value["submit_time"])
		list.VerifyTime = string(value["verify_time"])
		list.ProEncoding = string(value["pro_encoding"])
		list.SupEncoding = string(value["sup_encoding"])
		list.RoomAlias = string(value["room_alias"])
		lists = append(lists, list)
	}
	return lists, total
}

//手机端更多期   A VS B
func (per *Periods) GetPeriodsPhone(nowtime string) ([]Periods, int64) {
	var list []Periods
	total, _ := Engine.Where("start_time >?", nowtime).Count(per)
	Engine.Where("start_time >?", nowtime).Find(&list)
	return list, total
}

//期数核算
func (per *Periods) SetPeriodComputation() {
	var sub SupportManagement
	sub.PeriodsId = per.PeriodsId
	list := sub.GetPeriodSupportList(1, 1000)
	session := Engine.NewSession()
	defer session.Close()
	err := session.Begin()
	//期数id核算字符串//$:#01>03#02>08#03->05#
	for _, sm := range list {
		var tmpStat int64
		if strings.Contains(per.SupEncoding, sm.SupEncoding) {
			tmpStat = 1
			//从系统划拨石榴籽  ---》
			//xxx()
			//给赢家划拨石榴籽
			//yy()
			//用户投注结果更新
		} else {
			tmpStat = 0
		}
		suport := SupportManagement{
			Id:               sm.Id,
			Uid:              sm.Uid,
			PeriodsId:        sm.PeriodsId,
			SupEncoding:      sm.SupEncoding,
			SupporNumber:     sm.SupporNumber,
			Odds:             sm.Odds,
			ComputationState: tmpStat,
			SupporTime:       sm.SupporTime,
			PrizenNmber:      sm.PrizenNmber,
		}
		_, err = session.Where("id=?", suport.Id).Update(&suport)
		Engine.ShowSQL(true)
		if err != nil {
			session.Rollback()
			return
		}
	}
	err = session.Commit()
	if err != nil {
		return
	}
}

func RoundNum(a, b int) int {
	if a%b == 0 {
		return a / b
	} else {
		return a/b + 1
	}
}

func PerNumAll(perid int64) int64 {
	ss := new(SupportManagement)
	total, _ := Engine.Where("periods_id =?", perid).Sum(ss, "suppor_number")
	return int64(total)
}

//期数核算
func (per *Periods) PerAccounting(StartTime string) (bool, string) {
	//取支持表带核算数据数量 count，例如 带核算总数1230条
	var (
		sub       SupportManagement
		moneyPond float64
		smrs      []SupportManagementResult
	)
	sub.PeriodsId = per.PeriodsId
	//获得所有投注人的石榴籽
	allMoney := PerNumAll(sub.PeriodsId)
	//前台来的 期数
	num := sub.GetSupportWinnerNum()
	session := Engine.NewSession()
	defer session.Close()
	//取整 (入)
	roundNum := RoundNum(num, 100)
	for i := 1; i <= roundNum; i++ {
		session.Begin()
		moneyPond = 0
		//事物开始
		//获取赢得数据
		list, _ := sub.GetWinnerInfoList(i-1, 100)
		for _, value := range list {
			smr := SupportManagementResult{
				Uid:              value.Uid,
				PeriodsId:        value.PeriodsId,
				ComputationState: 1,
				IsWin:            true,
			}
			smrs = append(smrs, smr)
			arr := strings.Split(value.Odds, ":")
			k, _ := strconv.ParseFloat(arr[0], 64)
			v, _ := strconv.ParseFloat(arr[1], 64)
			odds := v / k
			peilv := odds * value.SupporNumber
			moneyPond += peilv
			if !updataUser(value.Uid, peilv) {
				session.Rollback()
				return false, "用户余额未添加"
			}
			//  添加核算资金流向表
			var user FundAccounting
			user.Uid = value.Uid
			user.PeriodsId = value.PeriodsId
			user.Money = peilv
			if !user.FundAccountingAdd() {
				session.Rollback()
				return false, "添加记录失败"
			}
		}
		_, err := session.Insert(&smrs)
		if err != nil {
			moneyPond = 0
			session.Rollback()
			return false, "操作数据错误"
		}
		var fund, fundTwo Fund
		_, err = fund.GetFundDesc()
		if err != nil {
			moneyPond = 0
			session.Rollback()
			return false, "资金池未找到"
		}
		//核算期的利润 = 所有人投注的钱 - 赢的钱  利润可为负
		//添加结算记录
		var per_era PeriodsEarnings
		per_era.PeriodsId = per.PeriodsId
		per_era.AllNumber = allMoney
		per_era.WinNumber = int64(moneyPond)
		per_era.EarningsNumber = allMoney - int64(moneyPond)
		if !per_era.PeriodsEarningsAdd() {
			moneyPond = 0
			session.Rollback()
			return false, "添加记录失败"
		}
		fundTwo.StorageFund = fund.StorageFund - float64(moneyPond)     //储备金chubeimoney  zijinmoney
		fundTwo.CurrencyMoney = fund.CurrencyMoney + float64(moneyPond) //流动金liudongmoney
		fundTwo.OfferAmount = fundTwo.StorageFund + fundTwo.CurrencyMoney
		//插入资金池流动动向
		_, err = session.Insert(&fundTwo)
		if err != nil {
			moneyPond = 0
			session.Rollback()
			return false, "资金次更新失败"
		}
		session.Commit()
	}
	return true, ""

}

//插入期数核算记录
func (per *PeriodsEarnings) PeriodsEarningsAdd() bool {
	rows, err := Engine.Insert(per)
	if rows > 0 || err != nil {
		return true
	}
	return false

}

//更新用户石榴籽
func updataUser(uid int64, money float64) bool {
	var user UidInfo
	user.Id = uid
	if !user.GetUserInfo() {
		return false
	}
	user.Balance = user.Balance + money
	if !user.UserInfoUp() {
		return false
	}
	return true
}

//更新这一期所有人的支持状态、核算状态
func UpdateInfo(perid int64) (bool, string) {
	//更新这一期投注的人  状态  ComputationState int64     //核算状态，0：未核算，1：核算结束
	//						 IsWin            int64     //支持状态 胜负状态，0：竞猜中，1：胜 2：负
	user := new(SupportManagement)
	user.ComputationState = 1
	user.IsWin = 2
	//更新这一期所有人的支持状态、核算状态
	_, err := Engine.Cols("computation_state", "is_win").Where("periods_id=?", perid).Update(user)
	if err != nil {
		return false, "核算状态更新失败"
	}
	var (
		per  Periods
		list SupportManagement
	)
	per.PeriodsId = perid
	if !per.GetPeriod() {
		return false, "该期数不存在"
	}
	per.State = 2
	if !per.PeriodUp() {
		return false, "期数修改失败"
	}

	list.IsWin = 1
	_, err = Engine.Table("support_management").Where("periods_id =? and ? like concat(\"%\",sup_encoding,\"%\")", per.PeriodsId, per.SupEncoding).Update(list)
	if err != nil {
		return false, "赢的用户更新失败"
	}
	return true, ""

}

//获取期ID
type PIDX struct {
	Dateday string //当前时间
	Id      int64  //当前id
	Pid     string
}

func (pd *PerId) GetDbPid() bool {
	flag, _ := Engine.Get(pd)
	return flag
}

//tongyitian
func (pd *PerId) GetDbPiduP() bool {
	total, err := Engine.Table("per_id").Where("id =?", pd.Id).Cols("pid", "pers_id", "date").Update(pd)
	if err != nil || total < 0 {
		return false
	}
	return true
}

var Pid PIDX

//获得期Id
func GetPerId() string {
	var pd PerId
	if !pd.GetDbPid() {
		date := time.Now().Format("2006-01-02")
		pd.Date = date
		date = strings.Replace(date, "-", "", -1)
		//add数据库
		pd.Pid = 1
		pd.PersId = date + "1"
		_, err := Engine.Insert(pd)
		if err != nil {
		}
		return pd.PersId
	} else {
		//存在时
		if pd.Date == time.Now().Format("2006-01-02") {
			pd.Pid++
			date := strings.Replace(pd.Date, "-", "", -1)
			pd.PersId = date + strconv.FormatInt(pd.Pid, 10)

			//同一天第二期
			if !pd.GetDbPiduP() {

			}
			return pd.PersId
		} else {
			//第二天时
			date := time.Now().Format("2006-01-02")
			pd.Date = date
			date = strings.Replace(date, "-", "", -1)
			pd.Pid = 1
			pd.PersId = date + strconv.FormatInt(pd.Pid, 10)
			//修改数据
			if !pd.GetDbPiduP() {

			}
			return pd.PersId
		}
	}

}

//资金流向列表
func (cp *FundAccounting) GetFundAccounting(page, rows int) ([]FundAccounting, int64) {
	var list []FundAccounting
	total, _ := Engine.Table("fund_accounting").Where("id >?", 0).Count(cp)
	Engine.Table("fund_accounting").Alias("c").Desc("c.id").Limit(rows, (page-1)*rows).Find(&list)
	Engine.ShowSQL(true)
	return list, total
}

//模糊查询资金流向表
func (cp *FundAccounting) GetFunAccountingByIdOrUserId(str string, page, rows int) ([]FundAccounting, int64, error) {
	var list []FundAccounting
	total, _ := Engine.Table("fund_accounting").Alias("r").Where("r.uid like ? or r.periods_id like ?", "%"+str+"%", "%"+str+"%").Count(cp)
	err := Engine.Table("fund_accounting").Alias("r").Where("r.uid like ? or r.periods_id like ?", "%"+str+"%", "%"+str+"%").Asc("r.id").Limit(rows, (page-1)*rows).Find(&list)
	return list, total, err
}

func (per *FundAccounting) FundAccountingAdd() bool {
	rows, err := Engine.Insert(per)
	if rows > 0 || err != nil {
		return true
	}
	return false
}

//增加后台提交支持结果记录表
func (sr *SupportRecord) SupportRecordAdd() bool {
	rows, err := Engine.Insert(sr)
	if rows > 0 && err == nil {
		return true
	}
	return false
}

//后台提交核算记录@liuhan 170314
func (sr *SupportRecord) GetSupportRecord(page, rows int, inputid string) ([]SupportRecord, int64) {
	var list []SupportRecord
	var total int64
	if inputid == "" {
		total, _ = Engine.Table("support_record").Where("id >?", 0).Count(sr)
		Engine.Table("support_record").Alias("c").Desc("c.id").Limit(rows, (page-1)*rows).Find(&list)
		return list, total
	} else {
		total, _ = Engine.Table("support_record").Where("id >? and per_id like ? or submitter_id like ?", 0, "%"+inputid+"%", "%"+inputid+"%").Count(sr)
		Engine.Table("support_record").Alias("c").Where("id >? and per_id like ? or submitter_id like ?", 0, "%"+inputid+"%", "%"+inputid+"%").Desc("c.id").Limit(rows, (page-1)*rows).Find(&list)
		return list, total
	}
}
