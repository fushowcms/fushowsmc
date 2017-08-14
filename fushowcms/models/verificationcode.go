package models

/*
 * Copyright (c) 2016—2017 www.fushow.cn, All rights reserved.
 */
import (
	_ "github.com/go-xorm/xorm"
)

//添加短信验证码表
func (pd *VerificationCode) VerificationCodeAdd() bool {
	rows, err := Engine.Insert(pd) //插入一行数据
	if rows > 0 || err != nil {
		return true
	}
	return false
}

//查找验证码手机号是否存在
func (pd *VerificationCode) GetVerificationCode() bool {
	flag, _ := Engine.Where("phone=?", pd.Phone).Get(pd)
	return flag
}

//更新短信验证码表
func (pd *VerificationCode) VerificationCodeUpdate() bool {
	num, err := Engine.Where("phone=?", pd.Phone).Update(pd)
	if num > 0 || err != nil {
		return true
	}
	return false
}
