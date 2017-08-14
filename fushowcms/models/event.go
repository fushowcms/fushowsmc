package models

/*
 * Copyright (c) 2016—2017 www.fushow.cn, All rights reserved.
 */
import (
	"strconv"

	_ "github.com/go-sql-driver/mysql"
)

//添加活动
func (ev *Event) EventAdd() bool {
	row, err := Engine.Insert(ev)
	if row > 0 || err == nil {
		return true
	}
	return false
}

//修改活动
func (ev *Event) EventUp() bool {
	if num, err := Engine.Where("Id=?", ev.Id).Update(ev); num > 0 || err == nil {
		return true
	}
	return false
}

//截止活动
func (ev *Event) StopEventUp() bool {
	num, err := Engine.Where("id=?", ev.Id).Cols("now_state").Update(ev)
	if num > 0 && err == nil {
		return true
	}
	return false
}

//删除活动
func (ev *Event) EventDel() bool {
	row, err := Engine.Delete(ev) //删除一行数据
	if row > 0 || err == nil {
		return true
	}
	return false
}

//删除活动
func (ev *Event) GetEvent() bool {
	row, _ := Engine.Get(ev) //删除一行数据

	return row
}

//判断活动记录存在
//bool  是否真   int64 石榴籽   int64   活动id
func (ev *Event) GetEventInfo(nowtime string) (bool, int64, int64) {
	var list []Event
	err := Engine.Where("? between start_time and end_time and event_type =? and now_state = 1", nowtime, ev.EventType).Find(&list)
	if err != nil {
		return false, 0, 0
	}
	if len(list) == 0 {
		return false, 0, 0
	}

	number := list[0].Number
	id := list[0].Id
	return true, number, id
}

//判断活动记录存在
//bool  是否真   int64 石榴籽   int64   活动id
func (ev *Event) GetEventInfos(nowtime string) (bool, int64, int64) {
	var list []Event
	err := Engine.Where("? between start_time and end_time and event_type =2 and now_state = 1", nowtime).Find(&list)
	if err != nil {
		return false, 0, 0
	}
	if len(list) == 0 {
		return false, 0, 0
	}

	number := list[0].Number
	id := list[0].Id
	return true, number, id
}

//获取活动列表(总表)
func (ev *Event) GetApplyList(page, rows int) ([]Event, int64) {
	var list []Event
	everyone := make([]Event, 0)
	Engine.Find(&everyone)
	total, _ := Engine.Where("id >? ", 0).Count(ev)
	Engine.Where("id >?", 0).Desc("id").Limit(rows, (page-1)*rows).Find(&list)
	return list, total
}

//活动详情
func (evr *EventRec) GetEventRecList(page, rows int) ([]EventRec, int64) {
	var list []EventRec
	total, _ := Engine.Where("event_id=? and event_type=?", evr.EventId, evr.EventType).Count(evr)
	Engine.Where("event_id=? and event_type=?", evr.EventId, evr.EventType).Limit(rows, (page-1)*rows).Desc("date_time").Find(&list)
	return list, total
}

func (ev *Event) EventAddInfo(em Event, userid int64) (bool, string) {
	var (
		fund    Fund
		newfund Fund
		uid     UidInfo
		evrec   EventRec
	)
	session := Engine.NewSession()
	defer session.Close()
	err := session.Begin()
	if err != nil {
		return false, "系统错误"
	}
	uid.Id = userid
	if !uid.GetUserInfo() {
		session.Rollback()
		return false, "该负责人不存在"
	}
	if uid.Type <= 3 {
		session.Rollback()
		return false, "该负责人没有权限添加活动"
	}
	//get资金池
	if !fund.GetFundsDesc() {
		session.Rollback()
		return false, "未找到资金池"
	}
	//update资金池---取钱
	newfund.StorageFund = fund.StorageFund - float64(em.AllNumber)     //储备金)
	newfund.CurrencyMoney = fund.CurrencyMoney + float64(em.AllNumber) //流通金
	newfund.OfferAmount = newfund.StorageFund + newfund.CurrencyMoney
	newfund.ModifyUserid = userid
	_, err = newfund.FundAdd()
	if err != nil {
		session.Rollback()
		return false, "资金池划拨错误"
	}
	em.NickName = uid.NickName
	//添加活动
	if !em.EventAdd() {
		session.Rollback()
		return false, "活动添加失败"
	}
	evrec.BalanceNumber = em.AllNumber
	evrec.EventId = em.Id
	evrec.EventType = em.EventType
	evrec.SponsorId = userid
	if !evrec.EventAdd() {
		session.Rollback()
		return false, "活动记录添加失败"
	}
	return true, "添加成功"
}

//截止活动
func (e *Event) EventOver(id, uid int64) (bool, string) {
	var (
		ev      Event
		newev   Event
		evre    EventRec
		fund    Fund
		newfund Fund
	)
	session := Engine.NewSession()
	defer session.Close()
	err := session.Begin()
	if err != nil {
		return false, "系统错误"
	}
	ev.Id = id
	if !ev.GetEvent() {
		session.Rollback()
		return false, "活动不存在"
	}
	if ev.NowState == 0 {
		session.Rollback()
		return false, "活动已经结束"
	}
	newev.Id = ev.Id
	newev.NowState = 0
	if !newev.StopEventUp() {
		session.Rollback()
		return false, "更新失败"
	}
	evre.EventId = ev.Id
	if !evre.GetEventDesc() {
		session.Rollback()
		return false, "活动资金错误"
	}
	if !fund.GetFundsDesc() {
		session.Rollback()
		return false, "资金池错误"
	}
	newfund.OfferAmount = fund.OfferAmount                                   //发行总额
	newfund.StorageFund = fund.StorageFund + float64(evre.BalanceNumber)     //储备金
	newfund.CurrencyMoney = fund.CurrencyMoney - float64(evre.BalanceNumber) //流通金
	newfund.ModifyUserid = uid
	_, err = newfund.FundAdd()
	if err != nil {
		session.Rollback()
		return false, "资金池划拨错误"
	}
	return true, strconv.FormatInt(evre.BalanceNumber, 10)
}
