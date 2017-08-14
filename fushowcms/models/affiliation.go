package models

/*
 * Copyright (c) 2016—2017 www.fushow.cn, All rights reserved.
 */
import (
	_ "github.com/go-sql-driver/mysql"
	//"github.com/go-xorm/xorm"
	"fmt"
)

//创建时间 2016-10-13 Txl
//添加申请
func (al *Affiliation) AffiliationAdd() bool {
	row, err := Engine.Insert(al) //插入一行数据
	if row > 0 || err == nil {
		return true
	}
	return false

}

//创建时间 2016-10-13 Txl
//删除申请————>通过机构编号删除
func (al *Affiliation) AffiliationDel() bool {
	fmt.Println(al.Id)
	row, err := Engine.Where("Id=?", al.Id).Delete(al) //删除一行数据
	if row > 0 || err == nil {
		return true
	}
	return false
}

//创建时间 2016-10-13 Txl
//修改推荐机构信息  --->通过机构编号修改
func (al *Affiliation) AffiliationUpdate() bool {
	if num, err := Engine.Where("AffId=?", al.AffId).Update(al); num > 0 || err == nil {
		return true
	}
	return false
}

//创建时间 2016-10-13 Txl
//获取推荐机构列表(总表)
func (al *Affiliation) GetAffiliationList(page, rows int) ([]Affiliation, int64) {
	var list []Affiliation
	total, _ := Engine.Where("id >?", 0).Count(al)
	Engine.Limit(rows, (page-1)*rows).Find(&list)
	return list, total
}

//创建时间 2016-10-13 Txl
//获取推荐表
func (al *Affiliation) GetAffiliation() bool {
	flag, _ := Engine.Get(al)
	if flag {
		return true
	}
	return false
}

//创建时间 2016-10-13 Txl
//WEB端获取所有的机构信息表
//无分页功能
func (al *Affiliation) GetAllAffiliation() ([]Affiliation, int64) {
	var list []Affiliation
	total, _ := Engine.Where("id >?", 0).Count(al)
	Engine.Find(&list)
	return list, total
}
