package models

/*
 * Copyright (c) 2016—2017 www.fushow.cn, All rights reserved.
 */
import (
	_ "github.com/go-sql-driver/mysql"
)

//增加活动
func (rm *OfficialFunctions) OfficialAdd() bool {
	row, err := Engine.Insert(rm) //插入一行数据
	if row > 0 || err == nil {
		return true
	}
	return false
}

//删除活动
func (rm *OfficialFunctions) OfficialDel() bool {
	row, err := Engine.Delete(rm) //插入一行数据
	if row > 0 || err == nil {
		return true
	}
	return false
}

//更新活动信息

func (rm *OfficialFunctions) OfficialUp() bool {
	row, err := Engine.Where("Id=?", rm.Id).Update(rm)
	if row > 0 && err == nil {
		return true
	}
	return false
}

/*
*作成者：刘长东；作成日：2016/08/19
*函数机能：获取活动信息
 */
func (rm *OfficialFunctions) GetOfficial() (OfficialFunctions, error) {
	var officialFunctions OfficialFunctions
	_, err := Engine.Where("Id=?", rm.Id).Get(&officialFunctions)
	return officialFunctions, err
}

//获取所有活动信息
func (rm *OfficialFunctions) GetOfficialList(page, rows int) ([]OfficialFunctions, int64) {
	var list []OfficialFunctions
	total, _ := Engine.Where("1=1", 0).Count(rm)
	Engine.Desc("id").Limit(rows, (page-1)*rows).Find(&list)
	return list, total
}

//获取显示活动信息
func (rm *OfficialFunctions) GetStartOfficialList() ([]OfficialFunctions, error) {
	var list []OfficialFunctions
	err := Engine.Where("live_state=?", rm.LiveState).Desc("id").Limit(3, 0).Find(&list)
	return list, err
}
