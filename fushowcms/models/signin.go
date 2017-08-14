//签到表

package models

/*
 * Copyright (c) 2016—2017 www.fushow.cn, All rights reserved.
 */
import (
	"fushowcms/comm"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

//添加活动
func (ss *SignIn) SignInAdd() bool {
	row, err := Engine.Insert(ss)
	if row > 0 || err == nil {
		return true
	}
	return false
}

//修改活动
func (ss *SignIn) SignInUp() bool {
	if num, err := Engine.Where("Id=?", ss.Id).Update(ss); num > 0 || err == nil {
		return true
	}
	return false
}

func (ss *SignIn) GetSigninList() (bool, []SignIn) {
	var list []SignIn
	err := Engine.Where("user_id=?", ss.UserId).Find(&list)
	if err != nil {
		return false, list
	}
	if len(list) == 0 {
		return false, list
	}
	return true, list
}

//判断今天是否签到了
func (ss *SignIn) IsSignInInfo(uid int64, nowtime string) bool {
	var s SignIn
	flag, _ := Engine.Where("user_id=? and sign_in_time >= ?", uid, nowtime).Get(s)
	return flag
}

//time:2016-11-04  txl
//参数 uid  用户id
//返回值  bool  是否成功
//       string  错误信息
//签到送石榴
func (ss *SignIn) SigninGetNumber(uid int64) (bool, string) {
	defer comm.QDLog.Unlock()
	comm.QDLog.Lock()
	var (
		user UidInfo //用户
	)
	date := time.Now().Format("2006-01-02")
	//事务开始
	session := Engine.NewSession()
	defer session.Close()
	err := session.Begin()
	user.Id = uid
	if !user.GetUserInfo() {
		session.Rollback()
		return false, "用户不存在"
	}
	//是否已经签到了？
	if ss.IsSignInInfo(uid, date) {
		session.Rollback()
		return false, "该用户今天已签到了"
	}
	ss.UserId = uid
	ss.SignInTime = time.Now().Format("2006-01-02")
	//添加签到
	if !ss.SignInAdd() {
		session.Rollback()
		return false, "签到失败"
	}
	//赠送石榴籽
	falg, mes, _ := BindGiveNumber(user.Id, 1)
	if !falg {
		return false, mes
	}
	err = session.Commit()
	if err != nil {
		return false, "错误"
	}
	return true, ""
}

func (si *SignIn) GetUserSigninList(uid int64) (bool, string, []SignIn) {
	var (
		user UidInfo
		ss   SignIn
		list []SignIn
	)
	user.Id = uid
	if !user.GetUserInfo() {
		return false, "用户不存在", list
	}
	ss.UserId = uid
	flag, list := ss.GetSigninList()
	if !flag {
		return false, "没有更多的签到记录", list
	}
	return true, "成功", list
}
