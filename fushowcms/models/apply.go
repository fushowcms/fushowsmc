package models

/*
 * Copyright (c) 2016—2017 www.fushow.cn, All rights reserved.
 */
import (
	"strconv"

	_ "github.com/go-sql-driver/mysql"
)

//添加申请
func (al *Applicant) ApplyAdd() bool {
	row, err := Engine.Insert(al) //插入一行数据
	if row > 0 || err == nil {
		return true
	}

	return false
}

//删除申请(删除id)
func (al *Applicant) ApplyDel() bool {
	row, err := Engine.Delete(al) //删除一行数据
	if row > 0 || err == nil {
		return true
	}
	return false
}

func (al *Applicant) GetApplyById() bool {
	flag, _ := Engine.Where("user_id=?", al.UserId).Get(al)
	return flag
}

//判断申请记录存在
func (al *Applicant) GetApplyInfo() bool {
	flag, _ := Engine.Get(al)
	return flag
}

//判断申请记录存在
func (al *Applicant) GetApplyId() bool {
	flag, _ := Engine.Where("id =?", al.Id).Get(al)
	if !flag {
		return false
	}
	return true
}

//获取申请列表(总表)
func (al *Applicant) GetApplyList(page, rows int) ([]Applicant, int64) {
	var list []Applicant
	everyone := make([]Applicant, 0)
	Engine.Find(&everyone)
	total, _ := Engine.Where("id >? and state=0", 0).Count(al)
	Engine.Where("id >? and state=0", 0).Desc("applicant_time").Limit(rows, (page-1)*rows).Find(&list)
	return list, total
}

//获取申请列表筛选
func (al *Applicant) GetApplyLists(page, rows int, state string) ([]Applicant, int64) {
	var list []Applicant
	states, _ := strconv.ParseInt(state, 10, 64)
	if states == 3 {
		total, _ := Engine.Where("id >?", 0).Count(al)
		Engine.Where("id >? ", 0).Desc("applicant_time").Limit(rows, (page-1)*rows).Find(&list)
		return list, total
	} else {
		total, _ := Engine.Where("id >? and state=?", 0, states).Count(al)
		Engine.Where("id >? and state=?", 0, state).Desc("applicant_time").Limit(rows, (page-1)*rows).Find(&list)
		return list, total
	}
}

//修改申请
func (al *Applicant) ApplyUp() bool {
	num, err := Engine.Where("id=?", al.Id).Update(al)
	if num > 0 && err == nil {
		return true
	}
	return false
}

//重复申请时更改state状态
func (al *Applicant) ApplyStateUp() bool {
	num, err := Engine.Where("id=?", al.Id).Cols("state").Update(al)
	if num > 0 && err == nil {
		return true
	}
	return false
}

//判断是否提交过申请
func (al *Applicant) ApplicantExit() bool {
	flag, _ := Engine.Where("user_id=?", al.UserId).Get(al)
	return flag
}
