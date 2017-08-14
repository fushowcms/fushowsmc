package models

/*
 * Copyright (c) 2016—2017 www.fushow.cn, All rights reserved.
 */

//增加资金池
func (fun *Fund) FundAdd() (int64, error) {
	rows, err := Engine.Insert(fun)
	return rows, err
}

//获取资金池信息
func (fun *Fund) GetFund() (int64, error) {
	_, err := Engine.Get(fun)
	return 1, err
}

//更新资金池信息
func (fun *Fund) FundUp() (int64, error) {
	var err error
	if num, err := Engine.Where("Id=?", fun.Id).Update(fun); num > 0 && err == nil {
		return 1, err
	}
	return -1, err
}

// 获取资金池信息
func (fun *Fund) GetFundDesc() (int64, error) {
	_, err := Engine.Table("fund").Alias("t").Desc("t.id").Get(fun)
	return 1, err
}

//获取资金池信息
func (fun *Fund) GetFundsDesc() bool {
	flag, _ := Engine.Table("fund").Alias("t").Desc("t.id").Get(fun)
	return flag
}

//获取资金池信息
func (fun *Fund) GetAllFundDesc(page, rows int) ([]Fund, int64) {
	var list []Fund
	total, _ := Engine.Where("id >?", 0).Count(fun)
	Engine.Desc("add_time").Limit(rows, (page-1)*rows).Find(&list)
	return list, total
}

func (gr *RechargingRecords) IsExist() bool {
	falg, _ := Engine.Table("recharging_records").Where("trade_no = ?", gr.TradeNo).Get(gr)
	return falg
}
