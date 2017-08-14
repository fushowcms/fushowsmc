package models

/*
 * Copyright (c) 2016—2017 www.fushow.cn, All rights reserved.
 */
import (
	_ "github.com/go-sql-driver/mysql"
)

//增加充值记录
func (records *RechargingRecords) RecordsAdd() bool {
	row, err := Engine.Insert(records) //插入一行数据
	if row > 0 || err == nil {
		return true
	}
	return false
}

//更新充值记录
func (records *RechargingRecords) RecordsUp() bool {
	row, err := Engine.Where("trade_no=?", records.TradeNo).Update(records)
	if row > 0 && err == nil {
		return true
	}
	return false
}
