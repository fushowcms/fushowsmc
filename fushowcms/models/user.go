package models

/*
 * Copyright (c) 2016—2017 www.fushow.cn, All rights reserved.
 */
import (
	"fmt"
	"fushowcms/comm"
	"strconv"

	"time"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/go-xorm/xorm"
)

//查询用户是否存在
func (uk *UserKey) GetUserKey() bool {
	has, _ := Engine.Get(uk)
	return has
}

func (uk *UserKey) GetUserKeyId() bool {
	has, _ := Engine.Where("id =?", uk.Id).Get(uk)
	return has
}

//查询用户信息
func (uk *UserKey) GetUserInfo() bool {
	has, _ := Engine.Where("Id=?", uk.Id).Get(uk)
	return has
}

//删除UserKey表新增加用户
func (uk *UserKey) DelUserKey() bool {
	num, err := Engine.Delete(uk) //删除一行数据
	if num > 0 && err != nil {
		return true
	}
	return false
}

//注册用户UserKey表
func (uk *UserKey) UserKeyAdd() bool {
	rows, err := Engine.Insert(uk) //插入一行数据
	if rows > 0 && err == nil {
		return true
	}
	return false
}

//注册用户UidInfo表
func (ui *UidInfo) UserInfoAdd() bool {
	rows, err := Engine.Insert(ui) //插入一行数据
	if rows > 0 && err == nil {
		return true
	}
	return false
}

//用户登录
func (ui *UidInfo) Login() bool {
	has, _ := Engine.Where("user_name=?", ui.UserName).Get(ui)
	return has
}

//删除用户
func (ui *UidInfo) UserInfoDel() bool {
	num, err := Engine.Delete(ui) //删除一行数据
	if num > 0 || err != nil {
		return true
	}
	return false
}
func (uk *UserKey) UserKeyDel() bool {
	num, err := Engine.Delete(uk) //删除一行数据
	if num > 0 || err != nil {
		return true
	}
	return false
}

//修改用户信息
func (ui *UidInfo) UserInfoUp() bool {
	row, err := Engine.Where("Id=?", ui.Id).Update(ui)
	if row > 0 || err == nil {
		return true
	}
	return false
}

func (ui *UidInfo) UserBICost() bool {
	row, err := Engine.Where("Id=?", ui.Id).Cols("integral", "balance").Update(ui)
	if row > 0 || err == nil {
		return true
	}
	return false
}

//主播更新操作
func (ui *UidInfo) UserAnchorBalance() bool {
	sql := "UPDATE uid_info SET gift_num = ? , pomegranate_num =? WHERE id = ? "
	_, err := Engine.Exec(sql, ui.GiftNum, ui.PomegranateNum, ui.Id)
	if err != nil {
		return false
	}
	return true
}

//用户余额清零
func (ui *UidInfo) UserBalanceNull() bool {
	row, err := Engine.Where("Id=?", ui.Id).Cols("balance").Update(ui)
	if row > 0 && err == nil {
		return true
	}
	return false
}

//赠送石榴籽扣钱
func (ui *UidInfo) UserCoastBalance() bool {
	row, err := Engine.Table("uid_info").Where("id=?", ui.Id).Cols("balance").Update(ui)
	if err == nil && row > 0 {
		return true
	}
	return false
}

//判断用户名密码是否正确 （修改密码）
func (ui *UidInfo) GetUserPwd(username string) bool {
	has, err := Engine.Cols("Id", "user_name", "pass_word", "version").Where("user_name=?", username).Get(ui)
	if has && err == nil {
		return true
	}
	return false
}

//修改用户密码
func (ui *UidInfo) PassUp() bool {
	row, err := Engine.Where("Id=?", ui.Id).Update(ui)
	if row > 0 && err == nil {
		return true
	}
	return false
}

//查询用户信息
func (ui *UidInfo) GetUserInfo() bool {
	Engine.ShowSQL(true)
	flag, _ := Engine.Get(ui)
	return flag
}

func (ui *UidInfo) GetUserInfoByName() bool {
	flag, _ := Engine.Where("user_name=?", ui.UserName).Get(ui)
	return flag
}

//查询用户是否填写了昵称
func (ui *UidInfo) IsEditNick() bool {
	if ui.NickName == "" {
		return false
	}
	return true
}

//每日数据
type SiteInfo struct {
	RegNumber      int64  //注册数
	LoginNumber    int64  //登陆数---->无法判断
	RecordsNumber  string //充值数
	CoastAllNumber string //消费总额-->消费记录较多，暂时没有
	SupportNumber  string //竞猜总额
	ApplyNumber    string //体现总额-->没有体现记录
}

//平台运营信息
func RegNumber(stardate, enddates string) SiteInfo {
	var (
		uid  UidInfo      //注册数
		uidl UidLoginInfo //登陆人数
		sit  SiteInfo     //每日数据
	)

	if enddates == "" { //默认今天数据
		stardate = time.Now().Format("2006-01-02")
		//注册数
		total, err := Engine.Table("uid_info").Where("reg_time = ?", stardate).Count(uid)
		if err != nil {
			fmt.Println("今日注册数错误信息", err)
		}
		if total == 0 {
			sit.RegNumber = 0
		}
		sit.RegNumber = total

		//登陆数
		Engine.ShowSQL(true)
		uil_total, err := Engine.Table("uid_login_info").Where("login_time= ? ", stardate).Count(uidl)
		if err != nil {
			fmt.Println("今日登陆数错误信息", err)
		}
		sit.LoginNumber = uil_total

		//充值数
		rec_sql := "SELECT SUM(money) FROM recharging_records WHERE time =?"
		rec_ss, err := Engine.Query(rec_sql, stardate)
		if err != nil {
			fmt.Println("今日充值数错误信息", err)
		}
		for _, value := range rec_ss {
			mon := string(value["SUM(money)"])
			if mon == "" {
				sit.RecordsNumber = "0"
			}
			sit.RecordsNumber = mon
		}
		//竞猜数
		sup_sql := "SELECT SUM(suppor_number) FROM support_management WHERE suppor_time =?"
		sup_ss, err := Engine.Query(sup_sql, stardate)
		if err != nil {
			fmt.Println("今日充值数错误信息", err)
		}
		for _, value := range sup_ss {
			mon := string(value["SUM(suppor_number)"])
			if mon == "" {
				sit.SupportNumber = "0"
			}
			sit.SupportNumber = mon
		}

		//体现总额
		set_sql := "SELECT SUM(apply_cashing_num) FROM settlement_detail WHERE cashing_date =?"
		set_ss, err := Engine.Query(set_sql, stardate)
		if err != nil {
			fmt.Println("今日体现总额错误信息", err)
		}
		for _, value := range set_ss {
			mon := string(value["SUM(apply_cashing_num)"])
			if mon == "" {
				sit.ApplyNumber = "0"
			}
			sit.ApplyNumber = mon
		}
		//消费总额
		gift_sql := "SELECT SUM(all_number) FROM gift_give WHERE give_date = ?"
		gift_res, err := Engine.Query(gift_sql, stardate)
		if err != nil {
			fmt.Println("消费总额错误信息", err)
		}
		for _, value := range gift_res {
			mon := string(value["SUM(all_number)"])
			if mon == "" {
				sit.CoastAllNumber = "0"
			}
			sit.CoastAllNumber = mon
		}
		return sit

	} else {
		//注册数
		total, err := Engine.Table("uid_info").Where("reg_time between ? and ?", stardate, enddates).Count(uid)
		if err != nil {
			fmt.Println("这段时间内注册数错误信息", err)
		}
		sit.RegNumber = total
		//登陆数
		uil_total, err := Engine.Table("uid_login_info").Where("login_time between ? and ?", stardate, enddates).Count(uidl)
		if err != nil {
			fmt.Println("这段时间内登陆数错误信息", err)
		}
		sit.LoginNumber = uil_total
		//充值数
		sql := "SELECT SUM(money) FROM recharging_records WHERE time BETWEEN ? and ?"
		ss, err := Engine.Query(sql, stardate, enddates)
		if err != nil {
			fmt.Println("充值错误信息", err)
		}
		for _, value := range ss {
			mon := string(value["SUM(money)"])
			sit.RecordsNumber = mon
		}
		//竞猜数
		sup_sql := "SELECT SUM(suppor_number) FROM support_management WHERE suppor_time BETWEEN ? and ?"
		sup_res, err := Engine.Query(sup_sql, stardate, enddates)
		if err != nil {
			fmt.Println("支持错误信息", err)
		}
		for _, value := range sup_res {
			mon := string(value["SUM(suppor_number)"])
			sit.SupportNumber = mon
		}
		//体现总额
		set_sql := "SELECT SUM(apply_cashing_num) FROM settlement_detail WHERE cashing_date BETWEEN ? and ?"
		set_res, err := Engine.Query(set_sql, stardate, enddates)
		if err != nil {
			fmt.Println("体现总额错误信息", err)
		}
		for _, value := range set_res {
			mon := string(value["SUM(apply_cashing_num)"])
			sit.ApplyNumber = mon
		}
		//消费总额
		gift_sql := "SELECT SUM(all_number) FROM gift_give WHERE give_date BETWEEN ? and ?"
		gift_res, err := Engine.Query(gift_sql, stardate, enddates)
		if err != nil {
			fmt.Println("消费总额错误信息", err)
		}
		for _, value := range gift_res {
			mon := string(value["SUM(all_number)"])
			sit.CoastAllNumber = mon
		}
		return sit
	}
}

//消费总额 ---->搜索  两个时间段以内
// 参数 stardate 开始  stardate 结束
//返回值  消费总额
//2016-11-28  txl
func AllNumber(enddate, stardate string) float64 {
	var (
		giftgive  GiftGive          //刷礼物
		buygoods  Order             //兑换商品
		sup       SupportManagement //支持
		allnumber float64
	)
	if stardate == "" { //默认今天数据
		//刷礼物
		gift_total, err := Engine.Table("gift_give").Where("give_date = ?", enddate).Sum(giftgive, "all_number")
		if err != nil {
			fmt.Println("今日刷新礼物总价值错误信息", err)
		}
		allnumber = allnumber + gift_total
		//兑换商品
		buy_total, err := Engine.Table("order").Where("create_time = ?", enddate).Sum(buygoods, "goods_total")
		if err != nil {
			fmt.Println("今日兑换商品总价值错误信息", err)
		}
		allnumber = allnumber + buy_total
		//支持
		sup_total, err := Engine.Table("support_management").Where("suppor_time = ?", enddate).Sum(sup, "suppor_number")
		if err != nil {
			fmt.Println("今日兑换商品总价值错误信息", err)
		}
		allnumber = allnumber + sup_total
		//返回总价值
		return allnumber
	} else {
		//刷礼物
		gift_total, err := Engine.Table("gift_give").Where("give_date between ? and ?", stardate, enddate).Sum(giftgive, "all_number")
		if err != nil {
			fmt.Println("这段时间内刷新礼物总价值错误信息", err)
		}
		allnumber = allnumber + gift_total
		//兑换商品
		buy_total, err := Engine.Table("order").Where("create_time between ? and ?", stardate, enddate).Sum(buygoods, "goods_total")
		if err != nil {
			fmt.Println("这段时间内兑换商品总价值错误信息", err)
		}
		allnumber = allnumber + buy_total
		//支持
		sup_total, err := Engine.Table("support_management").Where("suppor_time between ? and ?", stardate, enddate).Sum(sup, "suppor_number")
		if err != nil {
			fmt.Println("这段时间内兑换商品总价值错误信息", err)
		}
		allnumber = allnumber + sup_total
		//返回总价值
		return allnumber
	}
}

//查询用户是否更改过昵称
func (ui *UidInfo) CheckNickName() bool {
	flag, _ := Engine.Get(ui)
	return flag
}

//修改用户昵称
func (ui *UidInfo) NickNameUp() bool {
	sql := "update uid_info set nick_name=?,nick_flag=? where id =?"
	Engine.Exec(sql, ui.NickName, ui.NickFlag, ui.Id)
	return true
}

func (ui *UidInfo) GetUserPhone() bool {
	flag, _ := Engine.Where("phone=?", ui.Phone).Get(ui)
	return flag
}

func (ui *UidInfo) GetUserNickName() bool {
	flag, _ := Engine.Where("nick_name=?", ui.NickName).Get(ui)
	return flag
}

//获取用户列表
func (uid *UidInfo) GetRootUserList(page, rows int) ([]UidInfo, int64) {
	var list []UidInfo
	total, _ := Engine.Where("id >?", 0).Count(uid)
	Engine.Limit(rows, (page-1)*rows).Find(&list)
	return list, total
}

//获取用户列表
func (uid *UidInfo) GetInfoPower(page, rows int) ([]UidInfo, int64) {
	var list []UidInfo
	total, _ := Engine.Where("id >?", 0).Count(uid)
	Engine.Limit(rows, (page-1)*rows).Desc("balance").Find(&list)
	return list, total
}

/*
*功能:查询密码是否匹配
*@muhailong
*日期:20161103
 */
func (ui *UidInfo) PassMate() bool {
	has, _ := Engine.Where("id=?", ui.Id).Get(ui)
	return has
}

/*
*功能:后台管理按类型和等级实现正倒序排序
*@muhailong
*日期:20161101
 */
func (uid *UidInfo) UserOrderByLevelADSC(page, rows int, sort string, order string, inputid string) ([]UidInfo, int64) {
	var list []UidInfo
	total, _ := Engine.Where("id >?", 0).Count(uid)
	if inputid == "" {
		if order == "asc" && sort == "Integral" {
			Engine.Table("uid_info").Alias("u").Where("u.id >?", 0).Asc("u.integral").Limit(rows, (page-1)*rows).Find(&list)
		} else if order == "desc" && sort == "Integral" {
			Engine.Table("uid_info").Alias("u").Where("u.id >?", 0).Desc("u.integral").Limit(rows, (page-1)*rows).Find(&list)
		} else if order == "asc" && sort == "Type" {
			Engine.Table("uid_info").Alias("u").Where("u.id >?", 0).Asc("u.type").Limit(rows, (page-1)*rows).Find(&list)
		} else if order == "desc" && sort == "Type" {
			Engine.Table("uid_info").Alias("u").Where("u.id >?", 0).Desc("u.type").Limit(rows, (page-1)*rows).Find(&list)
		} else {
			Engine.Table("uid_info").Alias("u").Where("u.id >?", 0).Desc("u.id").Limit(rows, (page-1)*rows).Find(&list)
		}
	} else {
		Engine.Where("id like ? or nick_name like ?", "%"+inputid+"%", "%"+inputid+"%").Desc("id").Limit(rows, (page-1)*rows).Find(&list)
	}
	return list, total
}

func (uid *UidInfo) FindUserOrderByLevelADSC(page, rows int, sort string, order string, inputid string) ([]UidInfo, int64) {
	var list []UidInfo
	total, _ := Engine.Where("id >?", 0).Count(uid)
	Engine.Table("uid_info").Alias("u").Where("u.id >?", 0).Desc("u.id").Limit(rows, (page-1)*rows).Find(&list)
	return list, total
}

/*
*功能:后台管理按类型和等级
*time :20161121  txl
 */
func (uid *UidInfo) UserOrderBySearch(page, rows int) ([]UidInfo, int64) {
	var list []UidInfo
	var uu UidInfo
	total, _ := Engine.Where("id >?", 0).Count(uu)
	err := Engine.Where("integral>? and integral<?", uid.Integral/10, uid.Integral).Limit(rows, (page-1)*rows).Find(&list)
	if err != nil {
		return list, total
	}
	return list, total
}
func (uid *UidInfo) UserOrderBySearch1(page, rows int) ([]UidInfo, int64) {
	var list []UidInfo
	var aaa UidInfo
	total, _ := Engine.Where("type = ?", uid.Type).Count(aaa)
	err := Engine.Where("type = ?", uid.Type).Limit(rows, (page-1)*rows).Find(&list)
	if err != nil {
		return list, total
	}
	return list, total
}

func (uid *UidInfo) UserOrderBySearch2(page, rows int) ([]UidInfo, int64) {
	var list []UidInfo
	var aaa UidInfo
	total, _ := Engine.Where("integral>=? and integral<?", uid.Integral/10, uid.Integral).Count(aaa)
	err := Engine.Where("integral>=? and integral<?", uid.Integral/10, uid.Integral).Limit(rows, (page-1)*rows).Find(&list)
	if err != nil {
		return list, total
	}
	return list, total
}

func (uid *UidInfo) UserOrderBySearch3(page, rows int) ([]UidInfo, int64) {
	var list []UidInfo
	var aaa UidInfo
	total, _ := Engine.Where("integral>=? and integral<? and type = ?", uid.Integral/10, uid.Integral, uid.Type).Count(aaa)
	err := Engine.Where("integral>=? and integral<? and type = ?", uid.Integral/10, uid.Integral, uid.Type).Limit(rows, (page-1)*rows).Find(&list)
	if err != nil {
		return list, total
	}
	return list, total
}
func (uid *UidInfo) UserOrderBySearch4(page, rows int) ([]UidInfo, int64) {
	var list []UidInfo
	var aaa UidInfo
	Engine.ShowSQL(true)
	total, _ := Engine.Table("uid_info").Where("integral<100").Count(aaa)
	err := Engine.Table("uid_info").Where("integral<100").Limit(rows, (page-1)*rows).Find(&list)
	if err != nil {
		return list, total
	}
	return list, total
}

/*
*功能:查询个人充值记录
*@muhailong
*日期:20161103
 */
//liuhan 修改
//Mao& Shuo 修改
func (rech *RechargingRecords) UserPayRecord(page, rows int) ([]RechargingRecords, int64) {
	var list []RechargingRecords
	total, _ := Engine.Where("uid=?", rech.Uid).Count(rech)
	Engine.Where("uid=?", rech.Uid).Desc("time").Limit(rows, (page-1)*rows).Find(&list)
	return list, total
}

/*
*功能:主播查询结算明细
*@muhailong
*日期:20161105
**/
//@liuhan 修改
func (sd *SettlementDetail) SettlementDetails(page, rows int) ([]SettlementDetail, int64) {
	var list []SettlementDetail
	total, _ := Engine.Where("Uid=?", sd.Uid).Count(sd)
	Engine.Where("Uid=?", sd.Uid).Desc("cashing_date").Limit(rows, (page-1)*rows).Find(&list)
	return list, total
}

/*
*功能:主播申请结算
*@muhailong
*日期:20161105
**/
func (sd *SettlementDetail) AnchorApplyCashing() bool {
	rows, err := Engine.Insert(sd) //插入一行数据
	if rows > 0 && err == nil {
		return true
	}
	return false
}

/*
*功能:银行卡绑定
*@muhailong
*日期:20161105
**/
func (ui *UidInfo) AnchorBindingBank() bool {
	sql := "update uid_info set bank_name=?,bank_card=?,bank_deposit=?,is_banding_bank=1 where id =?"
	Engine.Exec(sql, ui.BankName, ui.BankCard, ui.BankDeposit, ui.Id)
	return true
}

/*
*功能:判断是否绑定了银行卡
*@muhailong
*日期:20161107
**/
func (ui *UidInfo) IsBindingBank() ([]UidInfo, int64) {
	var list []UidInfo
	total, _ := Engine.Where("id=?", ui.Id).Count(ui)
	Engine.Where("id=?", ui.Id).Find(&list)
	return list, total
}

/*
*功能:判断申请结算数量是否足够
*@muhailong
*日期:20161107
**/
func (ui *UidInfo) IsEnough(pome_num int64) bool {
	if ui.PomegranateNum-pome_num >= 0 {
		return true
	} else {
		return false
	}

}

/*
*功能:判断是否已经申请结算
*@muhailong
*日期:20161124
**/
func (sd *SettlementDetail) IsMonthCashing() ([]SettlementDetail, error) {
	var list []SettlementDetail
	err := Engine.Table("settlement_detail").Where("uid = ?", sd.Uid).Desc("cashing_date").Limit(1).Find(&list)
	return list, err
}

/*
*功能:后台查询主播列表
*@muhailong
*日期:20161014
 */
func (ui *UidInfo) GetAnchorInfos(page, rows int, inputid string) ([]UidInfoSettlementDetail, int64) {
	var list []UidInfoSettlementDetail
	var sd SettlementDetail
	var total int64
	if inputid == "" {
		total, _ = Engine.Table("settlement_detail").Alias("s").Join("LEFT", []string{"uid_info", "u"}, "s.uid = u.id").Where("u.type =?", 1).Count(sd)
		sql := "SELECT * FROM `settlement_detail` AS `s` LEFT JOIN `uid_info` AS `u` ON s.uid = u.id WHERE (u.type = 1) ORDER BY cashing_date DESC"
		Engine.SQL(sql).Find(&list)
	} else {
		total, _ = Engine.Table("settlement_detail").Alias("s").Join("LEFT", []string{"uid_info", "u"}, "s.uid = u.id").Where("u.type =? and u.nick_name like ?", 1, "%"+inputid+"%").Count(sd)
		sql := "SELECT * FROM `settlement_detail` AS `s` LEFT JOIN `uid_info` AS `u` ON s.uid = u.id WHERE (u.type = 1) and u.nick_name like '%" + inputid + "%'  ORDER BY cashing_date DESC"
		Engine.SQL(sql).Find(&list)
	}

	return list, total
}

/*
*功能:后台主播结算
*@muhailong
*日期:20161014
 */
func (usd *UidInfoSettlementDetail) BalanceUp() bool {
	var (
		fund    Fund
		newfund Fund
	)
	session := Engine.NewSession()
	defer session.Close()
	session.Begin()
	sql := "update uid_info u set u.pomegranate_num=u.pomegranate_num - ? where id =?"
	sql2 := "update settlement_detail s set s.is_apply=0,s.is_cashing=1 where s.id =?"
	_, err := Engine.Exec(sql, usd.SettlementDetail.ApplyCashingNum, usd.Uid)
	if err != nil {
		session.Rollback()
		return false
	}

	_, err = Engine.Exec(sql2, usd.SettlementDetail.Id)
	if err != nil {
		session.Rollback()
		return false
	}
	//查找资金池
	if !fund.GetFundsDesc() {
		session.Rollback()
		return false
	}
	//更新资金池
	newfund.StorageFund = fund.StorageFund + float64(usd.SettlementDetail.ApplyCashingNum)
	newfund.CurrencyMoney = fund.CurrencyMoney - float64(usd.SettlementDetail.ApplyCashingNum)
	newfund.OfferAmount = newfund.StorageFund + newfund.CurrencyMoney
	//插入记录
	_, err = newfund.FundAdd()
	//错误
	if err != nil {
		fmt.Println("更新资金池错误信息", err)
		session.Rollback()
		return false
	}
	//提交
	err = session.Commit()
	if err != nil {
		return false
	}
	return true
}

/*
*功能:用户注册
*@muhailong
*日期:20161110
**/
func (ui *UidInfo) Reg() (bool, int64) {
	rows, err := Engine.Insert(ui) //插入一行数据
	if rows > 0 && err == nil {
		_, _, number := BindGiveNumber(ui.Id, 0)
		return true, number
	}
	return false, 0
}

/*
*功能:查看昵称是否存在
*@muhailong
*日期:20161111
**/
func (ui *UidInfo) CheckNick() bool {
	has, _ := Engine.Where("nick_name=?", ui.NickName).Count(ui)
	if has > 0 {
		return false
	}
	return true
}

/*
*功能:用户填写昵称
*@muhailong
*日期:20161110
**/
func (ui *UidInfo) Nick() bool {
	sql := "update uid_info u set u.nick_name=? ,u.nick_flag = 1 where u.Id =?"
	Engine.Exec(sql, ui.NickName, ui.Id)
	return true
}

/*
*功能:用户重置密码
*@muhailong
*日期:20161114
**/
func (ui *UidInfo) ResetPass(password string) bool {
	sql := "update uid_info u set u.pass_word=? where u.phone=?"
	Engine.Exec(sql, password, ui.Phone)
	return true
}

//查询用户信息列表
func (ui *UidInfo) GetUserList(page, rows int, str string) ([]UidInfo, int64) {
	var (
		list  []UidInfo
		total int64
	)
	if str == "" {
		total, _ = Engine.Where("id >?", 0).Count(ui) //查询总数
		Engine.Limit(rows, (page-1)*rows).Find(&list)
	} else {
		total, _ = Engine.Where("nick_name=?", str).Count(ui) //查询总数
		Engine.Where("nick_name=?", str).Limit(rows, (page-1)*rows).Find(&list)
	}

	return list, total
}

//参数  num2 系统回收石榴籽  num1 可兑换的石榴籽
//     uid  用户ID    anchor 主播ID   nowgift 礼物ID  number 赠送数量

//赠送礼物
func (user *UidInfo) GiftFund(num1, num2, uid, anchorid, giftid, number int64) (bool, string) {
	defer comm.FSLog.Unlock()
	comm.FSLog.Lock()
	var (
		err      string
		giftgive GiftGive
		uk       UserKey
	)
	user.Id = uid
	uk.Id = uid
	num1 = num1 * number
	num2 = num2 * number
	//n个礼物总价值
	allnum := num1 + num2
	//事务开始
	session := Engine.NewSession()
	defer session.Close()
	if num1 == 0 || uid == 0 || anchorid == 0 || giftid == 0 || number == 0 {
		err = "参数错误"
		return false, err
	}
	session.Begin()
	//用户扣除石榴籽
	if !user.GetUserInfo() {
		session.Rollback()
		err = "赠送人不存在"
		return false, err
	}
	//用户余额是否够
	if user.Balance == float64(allnum) {
		user.Balance = 0
	} else {
		user.Balance = user.Balance - float64(allnum)
	}
	if user.Balance < 0 {
		session.Rollback()
		err = "用户余额不足"
		return false, err
	}
	user.Integral = user.Integral + allnum/10
	//用户更新扣除石榴籽

	if !user.UserBICost() {
		session.Rollback()
		err = "用户扣除石榴籽失败"
		return false, err
	}
	//获取配置文件设置的更新次数
	a_times, _ := strconv.ParseInt(comm.GetConfig("GIFT", "anchor_times"), 10, 64)
	s_times, _ := strconv.ParseInt(comm.GetConfig("GIFT", "sys_times"), 10, 64)
	//主播的到的石榴籽
	if !AnchorGiftUp(allnum, a_times, num1, anchorid) {
		session.Rollback()
		err = "anchor error"
		return false, err
	}
	if !SysGiftUp(s_times, num2) {
		session.Rollback()
		err = "sys error"
		return false, err
	}
	//添加一条赠送记录
	giftgive.BenefactorId = uid     //赠送人
	giftgive.RecipientId = anchorid //接受人
	giftgive.GiftId = giftid        //礼物id
	giftgive.GiftNum = number       //礼物数量
	giftgive.AllNumber = allnum     //礼物总价值
	if !giftgive.GiftGiveAdd() {
		session.Rollback()
		err = "添加赠送记录失败"
		return false, err
	}
	session.Commit()
	return true, err
}

//用户扣除石榴籽
func (user *UidInfo) UserCost(balance int64) (bool, string) {
	var err string
	//用户是否存在
	if !user.GetUserInfo() {
		err = "该用户不存在"
		return false, err
	}
	if user.Balance == float64(balance) {
		user.Balance = 0
		if !user.UserBalanceNull() {
			return false, "用户更新失败"
		}
	} else {
		user.Balance = user.Balance - float64(balance)
		if user.Balance < 0 {
			err = "当前余额不足"
			return false, err
		}
		//用户更新扣除石榴籽
		if !user.UserInfoUp() {
			err = "竞猜失败"
			return false, err
		}
	}
	return true, err
}

func UserLoginDoing(userinfo UidInfo, uk UserKey) (bool, string) {
	session := Engine.NewSession()
	defer session.Close()
	err := session.Begin()
	if err != nil {
		return false, "系统错误"
	}
	//user_info添加数据
	_, err = session.Insert(&userinfo)
	if err != nil {
		session.Rollback()
		return false, "uid_info插入错误"
	}
	//提交数据
	errs := session.Commit()
	if errs != nil {
		return false, "提交错误"
	}
	return true, ""
}

//注册
func Register(uid UidInfoJson, uk UserKey) (bool, string, UidInfoJson) {
	var (
		ev      Event    //活动
		newev   Event    //更新活动  经费不足时自动更新state=-1
		evrec   EventRec //活动记录
		newever EventRec //活动记录
		fund    Fund     //资金池
		newfund Fund     //剩余活动资金不够一个人时，资金还回给资金池
	)
	session := Engine.NewSession()
	defer session.Close()
	err := session.Begin()
	if err != nil {
		return false, "系统错误", uid
	}
	if uk.GetUserKey() { //用户已存在
		session.Rollback()
		return false, "用户已经存在", uid
	}
	//判断该时间注册时是否赠送石榴籽
	ev.EventType = 0
	nowtime := time.Now().Format("2006-01-02 15:04:05")
	flag, number, evid := ev.GetEventInfo(nowtime)
	//活动存在时
	if flag {
		evrec.EventId = evid
		//查找活动记录
		if evrec.GetEventDesc() {
			//活动资金
			newever.BalanceNumber = evrec.BalanceNumber - number
			//活动资金是否>0
			if newever.BalanceNumber > 0 {
				uid.Balance = uid.Balance + float64(number)
				//活动资金足够时，生成活动记录
				newever.UserId = uid.Id
				newever.EventId = evrec.EventId
				newever.EventType = evrec.EventType
				newever.SponsorId = evrec.SponsorId
				//添加活动记录
				if !newever.EventAdd() {
					session.Rollback()
					return false, "活动记录更新失败", uid
				}
			}
			if !uk.UserKeyAdd() { //增加用户失败
				session.Rollback()
				return false, "增加用户失败", uid
			}
			uid.Id = uk.Id
			if !uid.UserInfoAdd() { //增加用户失败
				session.Rollback()
				return false, "增加用户失败", uid
			}
			//活动资金不足够时，结束活动
			newev.Id = evid
			newev.NowState = -1
			if !newev.StopEventUp() {
				session.Rollback()
				return false, "更新失败", uid
			}
			//活动资金不足够时，剩余资金自动还资金池
			if fund.GetFundsDesc() {
				newfund.StorageFund = fund.StorageFund + float64(evrec.BalanceNumber)     //储备金
				newfund.CurrencyMoney = fund.CurrencyMoney - float64(evrec.BalanceNumber) //流通资金
				newfund.OfferAmount = newfund.StorageFund + newfund.CurrencyMoney         //发行总额
				newfund.ModifyUserid = evrec.SponsorId
				//更新资金池
				_, err := newfund.FundAdd()
				if err != nil {
					session.Rollback()
					return false, "资金池更新失败", uid
				}
			}
			err = session.Commit()
			if err != nil {
				return false, "错误", uid
			}
			return true, "添加成功", uid
		}
	}
	if !uk.UserKeyAdd() { //增加用户失败
		session.Rollback()
		return false, "增加用户失败", uid
	}
	uid.Id = uk.Id
	if !uid.UserInfoAdd() { //增加用户失败
		session.Rollback()
		return false, "增加用户失败", uid
	}
	err = session.Commit()
	if err != nil {
		return false, "错误", uid
	}
	return true, "添加成功", uid
}

//绑定送石榴籽
// 参数  uid 参与用户
// 返回值  bool    true 存在活动 ---->>>(string  赠送成功\活动已结束\目前没有活动)     false 目前没有活动
// string  错误信息
//time  2016-11-04  txl
func BindGiveNumber(uid int64, event_type int64) (bool, string, int64) {
	var (
		ev      Event    //活动
		newev   Event    //活动
		evrec   EventRec //活动记录
		newever EventRec //新的活动记录
		fund    Fund
		newfund Fund
	)
	//事务开始
	session := Engine.NewSession()
	defer session.Close()
	session.Begin()
	ev.EventType = event_type
	nowtime := time.Now().Format("2006-01-02 15:04:05")
	flag, number, evid := ev.GetEventInfo(nowtime)
	//活动存在时，用户增加石榴
	if flag {
		evrec.EventId = evid
		//查找活动记录
		if evrec.GetEventDesc() {
			//活动资金
			newever.BalanceNumber = evrec.BalanceNumber - number
			//活动资金是否>0
			if newever.BalanceNumber >= 0 {
				var nowuid UidInfo
				nowuid.Id = uid
				if !nowuid.GetUserInfo() {
					fmt.Println("用户不存在")
				}
				nowuid.Balance = nowuid.Balance + float64(number)
				//活动资金足够时，生成活动记录
				newever.UserId = nowuid.Id
				newever.EventId = evrec.EventId
				newever.EventType = event_type //evrec.EventType
				newever.SponsorId = evrec.SponsorId
				//添加活动记录
				if !newever.EventAdd() {
					session.Rollback()
					return false, "活动记录更新失败", int64(0)
				}
				//用户增加石榴籽
				if !nowuid.UserInfoUp() {
					session.Rollback()
					return false, "用户更新失败", int64(0)
				}
				//赠送记录添加
				var gs GiveSlz
				gs.BenefactorId = 1  //系统
				gs.RecipientId = uid //接收人Id
				gs.Num = number      //数量
				if !gs.AddGiveSlz() {
					return false, "赠送记录添加失败", int64(0)
				}
				return true, "赠送成功", number
			}
			//停止活动
			newev.Id = evid
			newev.NowState = -1
			if !newev.StopEventUp() {
				session.Rollback()
				return false, "更新失败", int64(0)
			}
			//活动经费不足时，资金返回资金池
			if fund.GetFundsDesc() {
				newfund.StorageFund = fund.StorageFund + float64(evrec.BalanceNumber)     //储备金
				newfund.CurrencyMoney = fund.CurrencyMoney - float64(evrec.BalanceNumber) //流通资金
				newfund.OfferAmount = newfund.StorageFund + newfund.CurrencyMoney         //发行总额
				newfund.ModifyUserid = evrec.SponsorId
				//更新资金池
				_, err := newfund.FundAdd()
				if err != nil {
					session.Rollback()
					return false, "资金池更新失败", int64(0)
				}
			}
			return true, "活动已结束", int64(0)
		}
	}
	return true, "目前没有活动", int64(0)
}

//添加一条登陆记录
func (uil UidLoginInfo) UidLoginInfoAdd() bool {
	rows, err := Engine.Insert(uil) //插入一行数据
	if rows > 0 && err == nil {
		return true
	}
	return false
}

//用户头像修改 addby liuhan
func (ui *UidInfo) UserUpFavicon() bool {
	Engine.ShowSQL(true)
	row, err := Engine.Where("Id=?", ui.Id).Update(ui)
	Engine.ShowSQL(true)
	if row > 0 && err == nil {
		return true
	}
	return false
}
