package models

/*
 * Copyright (c) 2016—2017 www.fushow.cn, All rights reserved.
 */
import (
	"strconv"

	_ "github.com/go-sql-driver/mysql"
)

//增加关注记录
func (am *AnchorRoomConcern) RoomAdd() bool {
	row, err := Engine.Insert(am) //插入一行数据
	if row > 0 || err == nil {
		return true
	}
	return false
}

//判断是否已经关注
func (am *AnchorRoomConcern) IsConcern() bool {
	has, _ := Engine.Where("Uid =? and User = ?", am.Uid, am.User).Get(am)
	return has
}

//删除关注记录
func (am *AnchorRoomConcern) RoomDel() bool {
	row, err := Engine.Where("Uid =? and User = ?", am.Uid, am.User).Delete(am) //删
	if row > 0 || err == nil {
		return true
	}
	return false
}

//更新关注记录
func (am *AnchorRoomConcern) RoomUp() bool {
	row, err := Engine.Where("Id=?", am.Id).Update(am)
	if row > 0 && err == nil {
		return true
	}
	return false
}

//获取一条关注记录
func (am *AnchorRoomConcern) GetRoom() bool {
	_, err := Engine.Get(am)
	if err == nil {
		return true
	}
	return false
}

//获取所有关注表
func (ar *AnchorRoomConcern) GetRoomList(page, rows int) ([]AnchorRoomConcern, int64) {
	var list []AnchorRoomConcern
	total, _ := Engine.Where("id >?", 0).Count(ar)
	Engine.Limit(rows, (page-1)*rows).Find(&list)
	return list, total
}

//获取所有关注表
func (ar *AnchorRoomConcern) GetRoomLists(page, rows int) ([]AnchorRoomConcern, int64) {
	var list []AnchorRoomConcern
	total, _ := Engine.Where("uid =?", ar.Uid).Count(ar)
	Engine.Limit(rows, (page-1)*rows).Find(&list)
	return list, total
}

//获取某个人关注记录
func (ar *AnchorRoomConcern) GetMyOrderRoomList() ([]MyOrderRoom, error) {
	var list []MyOrderRoom
	err := Engine.Table("anchor_room_concern").Alias("a").Join("INNER", []string{"anchor_room", "r"}, "a.user = r.uid").Join("INNER", []string{"anchor_room_time", "t"}, "a.user = t.uid").Where("a.uid=?", ar.Uid).Find(&list)
	return list, err
}

//查看我的关注列表
func (ar *AnchorRoomConcern) GetMyAttention(page, rows int) ([]MyOrderRoom, error, int64) {
	var (
		list []MyOrderRoom
	)
	sql := "SELECT r.id,r.uid,r.live_address,r.live_cover,r.live_state,t.nick_name,r.room_alias,r.room_type FROM `anchor_room_concern` AS `a` LEFT JOIN `anchor_room` AS `r` ON a.user = r.uid LEFT JOIN `uid_info` AS `t` ON a.user = t.id WHERE a.uid=? ORDER BY r.live_state Desc LIMIT ?,?"
	results, err := Engine.Query(sql, ar.Uid, (page-1)*rows, rows)
	for _, value := range results {
		var (
			now MyOrderRoom
			ar  AnchorRoomConcern
		)
		now.LiveAddress = string(value["live_address"])
		now.User, _ = strconv.ParseInt(string(value["uid"]), 10, 64) //主播id
		now.LiveState, _ = strconv.ParseInt(string(value["live_state"]), 10, 64)
		now.RoomAlias = string(value["room_alias"])
		now.NickName = string(value["nick_name"])
		now.LiveCover = string(value["live_cover"])
		now.Id, _ = strconv.ParseInt(string(value["id"]), 10, 64) //房间id
		ar.User = now.User
		total, _ := Engine.Where("user=?", ar.User).Count(ar)
		now.LiveNumber = total
		list = append(list, now)
	}
	total, _ := Engine.Table("anchor_room_concern").Alias("a").Where("a.uid=?", ar.Uid).Count(ar)
	return list, err, total
}

/*
*功能 查看我关注的主播是否开播
*牟海龙20161012
 */
func (ar *AnchorRoomConcern) IsOpenGetMyAttention(page, rows int) ([]MyOrderRoom, error) {
	var (
		list []MyOrderRoom
	)
	err := Engine.Table("anchor_room_concern").Alias("a").Join("INNER", []string{"anchor_room", "r"}, "a.user = r.uid").Join("INNER", []string{"uid_info", "t"}, "a.user = t.id").Where("a.uid=? and r.live_state=?", ar.Uid, 1).Limit(rows, (page-1)*rows).Find(&list)
	return list, err
}

/* 功能 查询我的关注数
 * @author 徐林 20161102
 */
func (ar *AnchorRoomConcern) GetMyAttentionCount(uid int64) (int64, error) {
	total, err := Engine.Where("user =?", uid).Count(ar)
	return total, err
}
