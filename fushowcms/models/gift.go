package models

/*
 * Copyright (c) 2016—2017 www.fushow.cn, All rights reserved.
 */
import (
	_ "github.com/go-sql-driver/mysql"
)

//增加礼物
func (gm *Gift) GiftAdd() bool {
	rows, err := Engine.Insert(gm)
	if rows > 0 || err != nil {
		return true
	}
	return false
}

//删除礼物
func (gm *Gift) GiftDel() bool {
	num, err := Engine.Delete(gm)
	if num > 0 || err != nil {
		return true
	}
	return false
}

//更新礼物信息
func (gm *Gift) GiftUp() (bool, error) {
	num, err := Engine.Where("Id=?", gm.Id).Update(gm)
	if num > 0 || err == nil {
		return true, err
	}
	return false, err
}

//获取礼物信息
func (gm *Gift) GetGift() bool {
	_, err := Engine.Get(gm)
	if err == nil {
		return true
	}
	return false
}

//根据Id查询礼物
func (gm *Gift) GetGiftCon() (bool, Gift) {
	var g Gift
	flag, _ := Engine.Where("id=?", gm.Id).Get(&g)
	return flag, g
}

//获取gift列表
func (gm *Gift) GetGiftList(page, rows int, sort, order, inputid string) ([]Gift, int64) {
	var list []Gift
	if order == "" {
		Engine.Where("gift_name like ?", "%"+inputid+"%").Desc("register_date").Limit(rows, (page-1)*rows).Find(&list)
	}
	if order == "asc" {
		Engine.Where("gift_name like ?", "%"+inputid+"%").Asc("buy_number").Limit(rows, (page-1)*rows).Find(&list)
	}
	if order == "desc" {
		Engine.Where("gift_name like ?", "%"+inputid+"%").Desc("buy_number").Limit(rows, (page-1)*rows).Find(&list)
	}
	total, _ := Engine.Where("id >?", 0).Count(gm)
	return list, total
}
