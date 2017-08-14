package models

/*
 * Copyright (c) 2016—2017 www.fushow.cn, All rights reserved.
 */

func OverWeb() bool {
	var is bool = true
	var n Fund
	//资金池清空
	_, err := Engine.Table("fund").Where("id>0").Delete(n)
	if err != nil {
		is = false
	}
	var ui UidInfo
	ui.Balance = 0
	ui.GiftNum = 0
	ui.PomegranateNum = 0
	//用户清空
	sql := "update uid_info set balance=0,gift_num=0,integral=0,pomegranate_num=0 "
	_, err = Engine.Exec(sql)
	if err != nil {
		is = false
	}
	var fund Fund
	fund.CurrencyMoney = 1000000                             //流通金
	fund.StorageFund = 1000000                               //储备金
	fund.OfferAmount = fund.StorageFund + fund.CurrencyMoney //发行额
	//添加资金池
	_, err = Engine.Table("fund").Insert(fund)
	if err != nil {
		is = false
	}
	sql2 := "delete from gift_give"
	_, err = Engine.Exec(sql2)
	//赠送记录清空
	if err != nil {
		is = false
	}
	//admin加钱
	sqluid := "update uid_info set balance = 1000000 where id = 1"
	_, err = Engine.Exec(sqluid)
	if err != nil {
		is = false
	}
	return is
}
