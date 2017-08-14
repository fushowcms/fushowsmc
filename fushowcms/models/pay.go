package models

/*
 * Copyright (c) 2016—2017 www.fushow.cn, All rights reserved.
 */
import (
	_ "github.com/go-sql-driver/mysql"
)

//不带搜索
func (od *RechargingRecords) GetPayList(rows, page int) ([]RechargingRecords, int64) {
	var list []RechargingRecords
	total, _ := Engine.Count(od)
	Engine.Desc("time").Limit(rows, (page-1)*rows).Find(&list)
	return list, total
}

//搜索
func (od *RechargingRecords) GetPaySearchList(rows, page int, uid string) ([]RechargingRecords, int64) {
	var list []RechargingRecords
	Engine.ShowSQL(true)
	total, _ := Engine.Where("uid like ?", "%"+uid+"%").Count(od)
	Engine.Where("uid like ?", "%"+uid+"%").Desc("time").Limit(rows, (page-1)*rows).Find(&list)
	return list, total
}
