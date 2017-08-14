//活动记录表
package models

/*
 * Copyright (c) 2016—2017 www.fushow.cn, All rights reserved.
 */
import (
	_ "github.com/go-sql-driver/mysql"
)

//添加活动
func (evr *EventRec) EventAdd() bool {
	row, err := Engine.Insert(evr)
	if row > 0 || err == nil {
		return true
	}
	return false
}

//获取活动记录信息
func (evr *EventRec) GetEventDesc() bool {
	flag, _ := Engine.Where("event_id =?", evr.EventId).Desc("date_time").Get(evr)
	return flag
}

func (evr *EventRec) GetEventDet() (bool, []EventRec) {
	var list []EventRec
	err := Engine.Where("event_id = ? and event_type=?", evr.EventId, evr.EventType).Find(&list)
	if err != nil {
		return false, list
	}
	if len(list) == 0 {
		return false, list
	}
	return true, list
}
