package models

/*
 * Copyright (c) 2016—2017 www.fushow.cn, All rights reserved.
 */
import (
	"time"

	_ "github.com/go-sql-driver/mysql"
)

//增加礼物
func (gm *GiftGive) GiftGiveAdd() bool {
	rows, err := Engine.Insert(gm)
	if rows > 0 || err != nil {
		return true
	}
	return false
}

//删除礼物
func (gm *GiftGive) GiftGiveDel() bool {
	num, err := Engine.Delete(gm)
	if num > 0 || err != nil {
		return true
	}
	return false
}

//更新礼物信息
func (gm *GiftGive) GiftGiveUp() bool {
	if num, err := Engine.Where("Id=?", gm.Id).Update(gm); num > 0 || err == nil {
		return true
	}
	return false
}

//获取礼物信息
func (gm *GiftGive) GetGiftGive() bool {
	_, err := Engine.Get(gm)
	if err == nil {
		return true
	}
	return false
}

//获取gift列表
func (gm *GiftGive) GetGiftGiveList(page, rows int) ([]GiftGive, int64) {
	var list []GiftGive
	total, _ := Engine.Where("id >?", 0).Count(gm)
	Engine.Limit(rows, (page-1)*rows).Find(&list)
	return list, total
}

type WeeksGift struct {
	Favicon  string
	NickName string
	Integral int64
}

//获取周榜
//time 2016-11-03  txl
func (gm *GiftGive) GetGiftGiveWeeks(anchorid int64) (bool, []WeeksGift) {
	var list []WeeksGift
	k := time.Now()
	d, _ := time.ParseDuration("-24h")
	//一周以前的时间
	nowtime := k.Add(d * 7).Format("2006-01-02 15:04:05")
	sql := "SELECT recipient_id, benefactor_id, SUM(all_number) AS Number, ui.nick_name,ui.favicon,ui.integral FROM gift_give AS gg LEFT JOIN uid_info AS ui ON gg.benefactor_id = ui.id WHERE gg.recipient_id = ? AND gg.give_date > ? GROUP BY benefactor_id ORDER BY SUM(all_number) DESC LIMIT 0,10"
	err := Engine.SQL(sql, anchorid, nowtime).Find(&list)
	if err != nil {
		return false, list
	}
	if len(list) == 0 {
		return false, list
	}
	return true, list
}

//获取总榜
//time 2016-11-03  txl
func (gm *GiftGive) GetGiftGiveMonth(anchorid int64) (bool, []WeeksGift) {
	var list []WeeksGift
	k := time.Now()
	d, _ := time.ParseDuration("-24h")
	//31天前以前的时间
	nowtime := k.Add(d * 31).Format("2006-01-02 15:04:05")
	sql := "SELECT recipient_id, benefactor_id, SUM(all_number) AS Number, ui.nick_name,ui.favicon,ui.integral FROM gift_give AS gg LEFT JOIN uid_info AS ui ON gg.benefactor_id = ui.id WHERE gg.recipient_id = ? AND gg.give_date > ? GROUP BY benefactor_id ORDER BY SUM(all_number) DESC LIMIT 0,10"
	err := Engine.SQL(sql, anchorid, nowtime).Find(&list)
	if err != nil {
		return false, list
	}
	if len(list) == 0 {
		return false, list
	}
	return true, list
}

type GiftGiveInfo struct {
	Id             string //记录Id
	BenefactorId   string `xorm:"index"`            //赠送人Id`
	RecipientId    string `xorm:"index"`            //接收人Id
	GiveDate       string `xorm:"datetime created"` //赠送时间
	GiftId         string //Id
	GiftName       string //礼物名称
	GiftNum        string //数量
	AllNumber      string //总价值
	BenefactorName string //赠送人昵称
	RecipientName  string //接收人昵称
}

//根据userId查询刷礼物记录 addby liuhan
func (gg *GiftGive) FindByUserIdAll(page, rows int) ([]GiftGiveInfo, error, int64) {
	var (
		list []GiftGiveInfo
		wps  GiftGiveInfo
	)
	page1 := rows
	rows1 := (page - 1) * rows
	total, _ := Engine.Table("gift_give").Alias("gg").Join("LEFT", []string{"gift", "g"}, "gg.gift_id = g.id").Where("benefactor_id =? or recipient_id=?", gg.BenefactorId, gg.BenefactorId).Count(gg)
	sql := "SELECT gg.benefactor_id,ui.nick_name AS BenefactorName,gg.recipient_id,uii.nick_name AS RecipientName,gg.give_date,gg.gift_id,g.gift_name,gg.gift_num,gg.all_number FROM gift_give AS gg LEFT JOIN gift AS g ON gg.gift_id = g.id LEFT JOIN uid_info AS ui ON ui.id = gg.benefactor_id LEFT JOIN uid_info AS uii ON uii.id = gg.recipient_id WHERE gg.benefactor_id =? or gg.recipient_id=? ORDER BY give_date DESC Limit ?,?"
	results, err := Engine.Query(sql, gg.BenefactorId, gg.BenefactorId, rows1, page1)
	for _, value := range results {
		wps.Id = string(value["id"])
		wps.BenefactorId = string(value["benefactor_id"])
		wps.RecipientId = string(value["recipient_id"])
		wps.GiftNum = string(value["gift_num"])
		wps.GiveDate = string(value["give_date"])
		wps.BenefactorName = string(value["BenefactorName"])
		wps.RecipientName = string(value["RecipientName"])
		wps.GiftId = string(value["gift_id"])
		wps.GiftName = string(value["gift_name"])
		list = append(list, wps)
	}
	return list, err, total
}

//后台查询刷礼物记录 addby liuhan
func (gg *GiftGive) FindGiveAll(page, rows int, inputid string) ([]GiftGiveInfo, error, int64) {
	var (
		list []GiftGiveInfo
		wps  GiftGiveInfo
	)
	page1 := rows
	rows1 := (page - 1) * rows
	total, _ := Engine.Table("gift_give").Alias("gg").Join("LEFT", []string{"gift", "g"}, "gg.gift_id = g.id").Join("LEFT", []string{"uid_info", "ui"}, "ui.id = gg.benefactor_id").Where("nick_name like ? or ui.id like ?", "%"+inputid+"%", "%"+inputid+"%").Count(gg)
	sql := "SELECT gg.benefactor_id,ui.nick_name AS BenefactorName,gg.recipient_id,uii.nick_name AS RecipientName,gg.give_date,gg.gift_id,g.gift_name,gg.gift_num,gg.all_number FROM gift_give AS gg LEFT JOIN gift AS g ON gg.gift_id = g.id LEFT JOIN uid_info AS ui ON ui.id = gg.benefactor_id LEFT JOIN uid_info AS uii ON uii.id = gg.recipient_id WHERE ui.nick_name like ? or ui.id like ?   ORDER BY give_date DESC Limit ?,?"
	results, err := Engine.Query(sql, "%"+inputid+"%", "%"+inputid+"%", rows1, page1)
	for _, value := range results {
		wps.Id = string(value["id"])
		wps.BenefactorId = string(value["benefactor_id"])
		wps.RecipientId = string(value["recipient_id"])
		wps.GiftNum = string(value["gift_num"])
		wps.GiveDate = string(value["give_date"])
		wps.BenefactorName = string(value["BenefactorName"])
		wps.RecipientName = string(value["RecipientName"])
		wps.GiftId = string(value["gift_id"])
		wps.GiftName = string(value["gift_name"])
		list = append(list, wps)
	}
	return list, err, total
}

//根据userId查询赠送石榴籽记录 addby liuhan
func (gs *GiveSlz) FindByUserIdGiveSlz(page, rows int, id int64) ([]GiveNumberInfo, error, int64) {
	var (
		list []GiveNumberInfo
		wps  GiveNumberInfo
	)
	Engine.ShowSQL(true)
	total, _ := Engine.Table("give_slz").Where("benefactor_id =? or recipient_id =?", id, id).Count(gs)
	page1 := rows
	rows1 := (page - 1) * rows
	sql := "SELECT a.id,a.benefactor_id,a.recipient_id,a.num,a.give_date,a.nick_name AS BenefactorName,uir.nick_name AS RecipientName FROM(SELECT gslz.id,gslz.benefactor_id,gslz.recipient_id,gslz.num,gslz.give_date,ui.nick_name		FROM	give_slz AS gslz	LEFT JOIN uid_info AS ui ON gslz.benefactor_id = ui.id		WHERE	(benefactor_id = ? OR recipient_id = ?	)	) AS a LEFT JOIN uid_info uir ON uir.id = a.recipient_id ORDER BY a.`id` DESC  Limit ?,?"
	results, err := Engine.Query(sql, id, id, rows1, page1)
	for _, value := range results {
		wps.Id = string(value["id"])
		wps.BenefactorId = string(value["benefactor_id"])
		wps.RecipientId = string(value["recipient_id"])
		wps.Num = string(value["num"])
		wps.GiveDate = string(value["give_date"])
		wps.BenefactorName = string(value["BenefactorName"])
		wps.RecipientName = string(value["RecipientName"])
		list = append(list, wps)
	}
	return list, err, total
}

//赠送石榴籽记录--->全纪录  txl
func (gs *GiveSlz) GetNumberList(page, rows int) ([]GiveNumberInfo, int64) {
	var (
		list []GiveNumberInfo
		wps  GiveNumberInfo
	)
	total, _ := Engine.Where("id >?", 0).Count(gs)
	sql := "SELECT a.num,a.give_date,a.benefactor_id,a.recipient_id,a.nick_name RecipientName,uir.nick_name BenefactorName FROM( SELECT gs.give_date,gs.num,gs.benefactor_id,gs.recipient_id,ui.nick_name FROM `give_slz` gs LEFT JOIN uid_info ui on gs.recipient_id = ui.id) as a LEFT JOIN uid_info uir ON uir.id = a.benefactor_id ORDER BY a.give_date	DESC limit ?,? "
	results, _ := Engine.Query(sql, (page-1)*rows, rows)
	for _, value := range results {
		wps.Id = string(value["id"])
		wps.BenefactorId = string(value["benefactor_id"])
		wps.RecipientId = string(value["recipient_id"])
		wps.Num = string(value["num"])
		wps.GiveDate = string(value["give_date"])
		wps.BenefactorName = string(value["BenefactorName"])
		wps.RecipientName = string(value["RecipientName"])

		list = append(list, wps)
	}
	return list, total
}

type GiveNumberInfo struct {
	Id             string //记录Id
	BenefactorId   string //赠送人Id
	RecipientId    string //接收人
	Num            string //数量
	GiveDate       string //赠送时间
	BenefactorName string //赠送人昵称
	RecipientName  string //接收人昵称

}

//赠送石榴籽记录-->模糊查询   txl
func (gs *GiveSlz) GetNumberLikeList(page, rows int, uid string) ([]GiveNumberInfo, int) {
	var (
		wps  GiveNumberInfo
		list []GiveNumberInfo
	)
	sql2 := "SELECT a.id,a.benefactor_id,a.recipient_id,a.num,a.give_date,a.nick_name AS BenefactorName,uir.nick_name AS RecipientName FROM (SELECT gslz.id,gslz.benefactor_id,gslz.recipient_id,gslz.num,gslz.give_date, ui.nick_name FROM give_slz AS gslz LEFT JOIN uid_info AS ui ON gslz.benefactor_id = ui.id WHERE ui.id IN ( SELECT ui.id FROM uid_info AS ui WHERE ui.nick_name LIKE ?)) AS a LEFT JOIN uid_info uir ON uir.id = a.recipient_id ORDER BY a.give_Date	DESC"
	res_total, _ := Engine.Query(sql2, "%"+uid+"%")
	sql := "SELECT a.id,a.benefactor_id,a.recipient_id,a.num,a.give_date,a.nick_name AS BenefactorName,uir.nick_name AS RecipientName FROM (SELECT gslz.id,gslz.benefactor_id,gslz.recipient_id,gslz.num,gslz.give_date, ui.nick_name FROM give_slz AS gslz LEFT JOIN uid_info AS ui ON gslz.benefactor_id = ui.id WHERE ui.id IN ( SELECT ui.id FROM uid_info AS ui WHERE ui.nick_name LIKE ?)) AS a LEFT JOIN uid_info uir ON uir.id = a.recipient_id ORDER BY a.give_Date	DESC limit ?,? "
	results, _ := Engine.Query(sql, "%"+uid+"%", (page-1)*rows, rows)
	for _, value := range results {
		wps.Id = string(value["id"])
		wps.BenefactorId = string(value["benefactor_id"])
		wps.RecipientId = string(value["recipient_id"])
		wps.Num = string(value["num"])
		wps.GiveDate = string(value["give_date"])
		wps.BenefactorName = string(value["BenefactorName"])
		wps.RecipientName = string(value["RecipientName"])
		list = append(list, wps)
	}
	return list, len(res_total)
}

//增加增送石榴籽记录表 addby liuhan
func (gs *GiveSlz) AddGiveSlz() bool {
	row, err := Engine.Insert(gs) //插入一行数据
	if row > 0 || err == nil {
		return true
	}
	return false
}

//根据Recipient查询刷礼物记录 addby liuhan
func (gg *GiftGive) FindByRecipient(page, rows int) ([]GiftGiveInfo, error, int64) {
	var (
		list []GiftGiveInfo
		wps  GiftGiveInfo
	)
	page1 := rows
	rows1 := (page - 1) * rows
	Engine.ShowSQL(true)
	total, _ := Engine.Where("recipient_id=?", gg.RecipientId).Count(gg)
	sql := "SELECT gg.benefactor_id,ui.nick_name AS BenefactorName,gg.recipient_id,uii.nick_name AS RecipientName,gg.give_date,gg.gift_id,g.gift_name,gg.gift_num,gg.all_number FROM gift_give AS gg LEFT JOIN gift AS g ON gg.gift_id = g.id LEFT JOIN uid_info AS ui ON ui.id = gg.benefactor_id LEFT JOIN uid_info AS uii ON uii.id = gg.recipient_id WHERE gg.recipient_id=? ORDER BY give_date DESC Limit ?,?"
	results, err := Engine.Query(sql, gg.RecipientId, rows1, page1)
	for _, value := range results {
		wps.Id = string(value["id"])
		wps.BenefactorId = string(value["benefactor_id"])
		wps.RecipientId = string(value["recipient_id"])
		wps.GiftNum = string(value["gift_num"])
		wps.GiveDate = string(value["give_date"])
		wps.BenefactorName = string(value["BenefactorName"])
		wps.RecipientName = string(value["RecipientName"])
		wps.GiftId = string(value["gift_id"])
		wps.GiftName = string(value["gift_name"])
		list = append(list, wps)
	}
	return list, err, total
}

//根据benefactorId查询刷礼物记录 addby liuhan
func (gg *GiftGive) FindByBenefactor(page, rows int) ([]GiftGiveInfo, error, int64) {
	var (
		list []GiftGiveInfo
		wps  GiftGiveInfo
	)
	page1 := rows
	rows1 := (page - 1) * rows
	Engine.ShowSQL(true)
	total, _ := Engine.Where("benefactor_id=?", gg.BenefactorId).Count(gg)
	sql := "SELECT gg.all_number,ui.nick_name AS BenefactorName,gg.recipient_id,uii.nick_name AS RecipientName,gg.give_date,gg.gift_id,g.gift_name,gg.gift_num,gg.all_number FROM gift_give AS gg LEFT JOIN gift AS g ON gg.gift_id = g.id LEFT JOIN uid_info AS ui ON ui.id = gg.benefactor_id LEFT JOIN uid_info AS uii ON uii.id = gg.recipient_id WHERE gg.Benefactor_id=? ORDER BY give_date DESC Limit ?,?"
	results, err := Engine.Query(sql, gg.BenefactorId, rows1, page1)
	for _, value := range results {
		wps.AllNumber = string(value["all_number"])
		wps.Id = string(value["id"])
		wps.BenefactorId = string(value["benefactor_id"])
		wps.RecipientId = string(value["recipient_id"])
		wps.GiftNum = string(value["gift_num"])
		wps.GiveDate = string(value["give_date"])
		wps.BenefactorName = string(value["BenefactorName"])
		wps.RecipientName = string(value["RecipientName"])
		wps.GiftId = string(value["gift_id"])
		wps.GiftName = string(value["gift_name"])
		list = append(list, wps)
	}
	return list, err, total
}
