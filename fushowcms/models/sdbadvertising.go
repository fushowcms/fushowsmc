package models

/*
 * Copyright (c) 2016—2017 www.fushow.cn, All rights reserved.
 */
import (
	_ "github.com/go-sql-driver/mysql"
)

//增加直播间广告
func (rm *Sdbadvertising) SdbadvertisingAdd() bool {
	row, err := Engine.Insert(rm) //插入一行数据
	if row > 0 || err == nil {
		return true
	}
	return false
}

//删除直播间广告
func (rm *Sdbadvertising) SdbadvertisingDel() bool {
	row, err := Engine.Delete(rm) //插入一行数据
	if row > 0 || err == nil {
		return true
	}
	return false
}

//修改直播间广告信息

func (rm *Sdbadvertising) SdbadvertisingUp() bool {
	row, err := Engine.Where("Id=?", rm.Id).Update(rm)
	if row > 0 && err == nil {
		return true
	}
	return false
}

//获取广告信息
func (rm *Sdbadvertising) GetSdbad() (Sdbadvertising, error) {
	var Sdbadvertising Sdbadvertising
	_, err := Engine.Where("Id=?", rm.Id).Get(&Sdbadvertising)
	return Sdbadvertising, err
}

//获取所有广告信息
func (rm *Sdbadvertising) GetSdadvertisinglist(page, rows int) ([]Sdbadvertising, int64) {
	var list []Sdbadvertising
	total, _ := Engine.Where("1=1", 0).Count(rm)
	Engine.Desc("id").Limit(rows, (page-1)*rows).Find(&list)
	return list, total
}

//获取直播间广告
func (rm *Sdbadvertising) GetSdbadvertising() ([]Sdbadvertising, error) {
	var list []Sdbadvertising
	err := Engine.Where("live_state=?", rm.LiveState).Desc("id").Limit(3, 0).Find(&list)
	return list, err
}
