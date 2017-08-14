package models

/*
 * Copyright (c) 2016—2017 www.fushow.cn, All rights reserved.
 */
import (
	"fmt"
	"strconv"
	"time"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/go-xorm/xorm"
)

//网站信息设置
func (ws *Website) SetSite() bool {
	has, errs := Engine.IsTableEmpty(ws)
	if !has && errs == nil {
		Engine.Update(ws) // 有数据则更新
		return true
	}
	// 没有数据则添加
	rows, err := Engine.Insert(ws)
	if rows > 0 && err == nil {
		fmt.Println(rows)
		return true
	}
	return false
}

//获取网站信息
func (ws *Website) GetSite() bool {
	has, err := Engine.Desc("id").Limit(1).Get(ws)
	if has && err == nil {
		return true
	}
	return false
}

//运营总数据
func GetWebInfoAll() WebInfo {
	var (
		ui UidInfo
		wb WebInfo
	)
	//用户总人数
	total, err := Engine.Table("uid_info").Where("id>?", 0).Count(ui)
	if err != nil {
		fmt.Println("用户总人数错误信息", err)
	}
	wb.UidNumber = total
	//总充值金额  求和
	rec_sql := "SELECT SUM(money) FROM recharging_records WHERE state = 1"
	sup_res, err := Engine.Query(rec_sql)
	if err != nil {
		fmt.Println("总充值金额错误信息", err)
	}
	for _, value := range sup_res {
		mon := string(value["SUM(money)"])
		wb.RecNumber = mon
	}
	//总体现金额
	set_sql := "SELECT SUM(cashing) AS allnumber FROM settlement_detail WHERE is_cashing = true"
	set_res, err := Engine.Query(set_sql)
	if err != nil {
		fmt.Println("总体现金额额错误信息", err)
	}
	for _, value := range set_res {
		mon := string(value["allnumber"])
		wb.SetAllNumber = mon
	}
	//待审核提现数---->
	var set SettlementDetail
	set_total, err := Engine.Table("settlement_detail").Where("is_cashing =false").Count(set)
	if err != nil {
		fmt.Println("申请未处理错误信息", err)
	}
	wb.WithNumber = set_total
	//待申请认证数
	var al Applicant
	al_total, err := Engine.Table("applicant").Where("state=0").Count(al)
	if err != nil {
		fmt.Println("申请未处理错误信息", err)
	}
	wb.AuditNumber = al_total
	//待结算竞猜数 ---->指石榴籽
	var (
		list       []Periods
		per_number int64
	)
	Engine.Table("periods").Where("state=1").Find(&list)
	for i := 0; i < len(list); i++ {
		//循环读取没有审核的期数
		per_sql := "SELECT SUM(suppor_number)AS AllNumber FROM support_management WHERE periods_id = ?"
		per_res, err := Engine.Query(per_sql, list[i].PeriodsId)
		if err != nil {
			fmt.Println("待结算竞猜数错误信息", err)
		}
		for _, value := range per_res {
			mon, _ := strconv.ParseInt(string(value["AllNumber"]), 10, 64)
			per_number += mon
		}
	}

	wb.ClreanNumer = per_number

	return wb
}

//运营总数据
type WebInfo struct {
	UidNumber    int64  //用户总数
	RecNumber    string //总充值金额
	SetAllNumber string //总体现金额
	WithNumber   int64  //待审核提现数---->没有体现记录
	AuditNumber  int64  //待申请认证数---->申请主播?
	ClreanNumer  int64  //待结算竞猜数
}

//平台收益情况
func GetWebEarnings(startDate, endDate string) SiteEarnings {
	var site_earnings SiteEarnings
	if endDate == "" {
		//礼物收益
		startDate = time.Now().Format("2006-01-02")
		var (
			list          []Gift
			giftAllNumber int64
		)
		err := Engine.Table("gift").Where("id >?", 0).Find(&list)
		if err != nil {
			fmt.Println("礼物错误信息", err)
		}
		for i := 0; i < len(list); i++ {
			per_sql := "SELECT SUM(gift_num)AS allnumber FROM gift_give WHERE gift_id = ?"
			per_res, err := Engine.Query(per_sql, list[i].Id)
			if err != nil {
				fmt.Println("待结算竞猜数错误信息", err)
			}
			for _, value := range per_res {
				mon, _ := strconv.ParseInt(string(value["allnumber"]), 10, 64)
				thisgift := mon * (list[i].BuyNumber - list[i].ToNumber)
				giftAllNumber += thisgift
			}
		}
		site_earnings.GiftNumber = giftAllNumber
		//竞猜收益
		per_ear := "SELECT SUM(earnings_number) AS allnumber FROM periods_earnings "
		perear_res, err := Engine.Query(per_ear)
		if err != nil {
			fmt.Println("竞猜收益错误信息", err)
		}
		for _, value := range perear_res {
			mon, _ := strconv.ParseInt(string(value["allnumber"]), 10, 64)
			site_earnings.PeriodsNumber = mon
		}

		//商品收益
		order_sql := "SELECT SUM(goods_total) AS allnumber FROM `order` "
		order_res, err := Engine.Query(order_sql)
		if err != nil {
			fmt.Println("商品收益错误信息", err)
		}
		for _, value := range order_res {
			mon, _ := strconv.ParseInt(string(value["allnumber"]), 10, 64)
			site_earnings.GoodsNumber = mon
		}
		site_earnings.AllBalance = site_earnings.GiftNumber + site_earnings.PeriodsNumber + site_earnings.GoodsNumber + site_earnings.OtherNumber
		return site_earnings
	} else {
		//礼物收益
		var (
			list          []Gift
			giftAllNumber int64
		)
		err := Engine.Table("gift").Where("id >?", 0).Find(&list)
		if err != nil {
			fmt.Println("礼物错误信息", err)
		}
		for i := 0; i < len(list); i++ {
			per_sql := "SELECT SUM(gift_num)AS allnumber FROM gift_give WHERE gift_id = ? AND give_date BETWEEN ? AND ?"
			per_res, err := Engine.Query(per_sql, list[i].Id, startDate, endDate)
			if err != nil {
				fmt.Println("待结算竞猜数错误信息", err)
			}
			for _, value := range per_res {
				mon, _ := strconv.ParseInt(string(value["allnumber"]), 10, 64)
				thisgift := mon * (list[i].BuyNumber - list[i].ToNumber)
				giftAllNumber += thisgift
			}
		}
		site_earnings.GiftNumber = giftAllNumber
		//竞猜收益
		per_ear := "SELECT SUM(earnings_number) AS allnumber FROM periods_earnings WHERE creat_time BETWEEN ? AND ?"
		perear_res, err := Engine.Query(per_ear, startDate, endDate)
		if err != nil {
			fmt.Println("竞猜收益错误信息", err)
		}
		for _, value := range perear_res {
			mon, _ := strconv.ParseInt(string(value["allnumber"]), 10, 64)
			site_earnings.PeriodsNumber = mon
		}
		//商品收益
		order_sql := "SELECT SUM(goods_total) AS allnumber FROM `order`  WHERE create_time BETWEEN ? AND ? "
		order_res, err := Engine.Query(order_sql, startDate, endDate)
		if err != nil {
			fmt.Println("商品收益错误信息", err)
		}
		for _, value := range order_res {
			mon, _ := strconv.ParseInt(string(value["allnumber"]), 10, 64)
			site_earnings.GoodsNumber = mon
		}
		site_earnings.AllBalance = site_earnings.GiftNumber + site_earnings.PeriodsNumber + site_earnings.GoodsNumber + site_earnings.OtherNumber
		return site_earnings
	}
}

//平台收益情况
type SiteEarnings struct {
	AllBalance    int64 //平台收益
	GiftNumber    int64 //礼物收益---->所有赠送人礼物总价值？
	PeriodsNumber int64 //竞猜收益---->所有投注数量
	GoodsNumber   int64 //商品收益
	OtherNumber   int64 //其他收益
}

//采集房间信息
type RoomInfoSpider struct {
	Id        int64
	SId       int64 `xorm:"index"`
	RId       string
	Title     string //房间标题
	Thumbimg  string //缩略图地址
	Number    string //房间人数
	NickName  string //主播名称
	RoomType  string //房间分类
	Form      string //采集来源
	CreatTime string //采集时间
}
