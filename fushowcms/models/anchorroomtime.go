package models

/*
 * Copyright (c) 2016—2017 www.fushow.cn, All rights reserved.
 */
import (
	_ "github.com/go-sql-driver/mysql"
)

//添加一条直播记录
func (at *AnchorRoomTime) AnchorRoomTimeAdd() bool {
	row, err := Engine.Insert(at) //插入一行数据
	if row > 0 || err == nil {
		return true
	}

	return false
}

//更新结束时间
func (at *AnchorRoomTime) AnchorRoomTimeUpdate() bool {
	num, err := Engine.Table("anchor_room_time").Where("id=?", at.Id).Cols("end_time", "anchorm_time").Update(at)
	if num > 0 && err == nil {
		return true
	}
	return false
}

//获取单条直播记录
func (am *AnchorRoomTime) GetRoomTime() bool {
	flag, _ := Engine.Table("anchor_room_time").Alias("t").Desc("t.start_time").Get(am)
	return flag
}
