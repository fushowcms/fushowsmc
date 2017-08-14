package models

/*
 * Copyright (c) 2016—2017 www.fushow.cn, All rights reserved.
 */
import (
	_ "github.com/go-sql-driver/mysql"
)

//增加直播间广告
func (rm *Dbadvertising) DbadvertisingAdd() bool {
	row, err := Engine.Insert(rm) //插入一行数据
	if row > 0 || err == nil {
		return true
	}
	return false
}

//删除直播间广告
func (rm *Dbadvertising) DbadvertisingDel() bool {
	row, err := Engine.Delete(rm) //插入一行数据
	//	fmt.Printf("rrrr %d\n", rm)
	if row > 0 || err == nil {
		return true
	}
	return false
}

//修改直播间广告信息

func (rm *Dbadvertising) DbadvertisingUp() bool {
	row, err := Engine.Where("Id=?", rm.Id).Update(rm)
	if row > 0 && err == nil {
		return true
	}

	return false
}

//获取广告信息
func (rm *Dbadvertising) GetDbad() (Dbadvertising, error) {
	var Dbadvertising Dbadvertising
	_, err := Engine.Where("Id=?", rm.Id).Get(&Dbadvertising)
	return Dbadvertising, err
}

//获取所有广告信息
func (rm *Dbadvertising) GetDbadvertisinglist(page, rows int) ([]Dbadvertising, int64) {
	var list []Dbadvertising
	total, _ := Engine.Where("1=1", 0).Count(rm)
	Engine.Desc("id").Limit(rows, (page-1)*rows).Find(&list)
	//	fmt.Println(len(list))
	return list, total
}

//获取直播间广告
func (rm *Dbadvertising) GetDbadvertising() ([]Dbadvertising, error) {
	var list []Dbadvertising
	err := Engine.Where("live_state=?", rm.LiveState).Desc("id").Limit(3, 0).Find(&list)
	return list, err
}
