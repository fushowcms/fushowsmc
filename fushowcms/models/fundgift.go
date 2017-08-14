package models

/*
 * Copyright (c) 2016—2017 www.fushow.cn, All rights reserved.
 */

func (fg *FundGift) FundGiftAdd() bool {
	row, err := Engine.Insert(fg) //插入一行数据
	if row > 0 || err == nil {
		return true
	}
	return false
}

//删除礼物
func (fg *FundGift) GiftDel() bool {
	num, err := Engine.Delete(fg)
	if num > 0 || err != nil {
		return true
	}
	return false
}

//更新礼物信息
func (fg *FundGift) GiftUp() (bool, error) {
	num, err := Engine.Where("Id=?", fg.Id).Update(fg)
	if num > 0 || err == nil {
		return true, err
	}
	return false, err
}

//获取礼物信息
func (fg *FundGift) GetGift() bool {
	_, err := Engine.Get(fg)
	if err == nil {
		return true
	}
	return false
}
