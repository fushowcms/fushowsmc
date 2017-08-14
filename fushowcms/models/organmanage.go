package models

/*
 * Copyright (c) 2016—2017 www.fushow.cn, All rights reserved.
 */
import (
	_ "github.com/go-sql-driver/mysql"
)

//增加机构 addby liuhan
func (o *Organ) AddOrgan() bool {
	row, err := Engine.Insert(o) //插入一行数据
	if row > 0 || err == nil {
		return true
	}
	return false
}

//查询机构 addby liuhan
func (o *Organ) FindOrgan(page, rows int, sort, order, inputid string) ([]Organ, int64) {
	var list []Organ
	if order == "" {
		Engine.Where("organ_name like ?", "%"+inputid+"%").Desc("modify_time").Limit(rows, (page-1)*rows).Find(&list)
	}
	if order == "asc" {
		Engine.Where("organ_name like ?", "%"+inputid+"%").Asc("organ_code").Limit(rows, (page-1)*rows).Find(&list)
	}
	if order == "desc" {
		Engine.Where("organ_name like ?", "%"+inputid+"%").Desc("organ_code").Limit(rows, (page-1)*rows).Find(&list)
	}
	total, _ := Engine.Where("id >?", 0).Count(o)
	return list, total

}

//根据机构id查询机构 addby liuhan
func (o *Organ) FindOrganById() ([]Organ, error) {
	var list []Organ
	err := Engine.Table("organ").Where("id =?", o.Id).Find(&list)
	return list, err
}

//根据机构code查询机构 addby liuhan
func (o *Organ) FindOrganByOrganCode() ([]Organ, error) {
	var list []Organ
	err := Engine.Table("organ").Where("organ_code =?", o.OrganCode).Find(&list)
	return list, err
}

//根据机构名称查询机构 addby liuhan
func (o *Organ) FindOrganByOrganName() ([]Organ, error) {
	var list []Organ
	err := Engine.Table("organ").Where("organ_name =?", o.OrganName).Find(&list)
	Engine.ShowSQL(true)
	return list, err
}

//机构表管理Dto
type OrganManages struct {
	Id             int64
	UserId         int64  //用户名称
	OrganName      string //联盟名称
	OrganId        int64  //机构id
	RechargeNum    int64  //充值数量
	RechargeMethod string //充值方式
	ModifyTime     string `xorm:"datetime created"`
	DelTime        string `xorm:"datetime deleted"` //记录删除时间
	Version        int64  `xorm:"version"`
}

//根据用户id查询机构人员 addby liuhan
func (o *OrganManage) FindOrganByUserId() ([]OrganManages, error) {
	var list []OrganManages
	err := Engine.Table("organ_manage").Alias("om").Join("LEFT", []string{"organ", "o"}, "om.organ_id = o.organ_code").Where("om.user_id =?", o.UserId).Find(&list)
	Engine.ShowSQL(true)
	return list, err
}

//删除机构 addby liuhan
func (o *Organ) DelOrgan() bool {
	row, err := Engine.Delete(o)
	if row > 0 || err == nil {
		return true
	}
	return false
}

//修改机构 addby liuhan
func (o *Organ) UpdateOrgan() bool {
	row, err := Engine.Where("id=?", o.Id).Update(o)
	if row > 0 || err == nil {
		return true
	}
	return false
}

//查询机构 addby liuhan
func (o *Organ) GetOrgan() bool {
	_, err := Engine.Get(o)
	if err == nil {
		return true
	}
	return false
}

//增加机构管理信息 addby liuhan
func (om *OrganManage) AddOrganManage() bool {
	row, err := Engine.Insert(om) //插入一行数据
	if row > 0 || err == nil {
		return true
	}
	return false
}

//根据机构id查询机构管理表 addby liuhan
func (om *OrganManage) FindOrganManageByOrganId() ([]OrganManage, error) {
	var list []OrganManage
	err := Engine.Table("organ_manage").Where("organ_id =?", om.OrganId).Find(&list)
	Engine.ShowSQL(true)
	return list, err

}

//机构表管理Dto
type OrganManageUser struct {
	Id             int64
	UserId         int64 //用户名称
	NickName       string
	OrganId        int64  //机构id
	RechargeNum    int64  //充值数量
	RechargeMethod string //充值方式
	ModifyTime     string `xorm:"datetime created"`
	DelTime        string `xorm:"datetime deleted"` //记录删除时间
	Version        int64  `xorm:"version"`
}

//根据机构id查询机构管理表 addby liuhan
func (om *OrganManage) FindOrganManageByOrganIds(page, rows int, sort, order, inputid string) ([]OrganManageUser, int64) {
	var list []OrganManageUser
	Engine.Table("organ_manage").Alias("om").Join("LEFT", []string{"uid_info", "ui"}, "om.user_id = ui.id").Where("organ_id= ?", om.OrganId).And("recharge_num != 0").Desc("om.modify_time").Limit(rows, (page-1)*rows).Find(&list)
	total, _ := Engine.Where("id >?", 0).Count(om)
	return list, total
}

//查询机构下充值的总额 addby liuhan
func (om *OrganManage) SumOrganManageByOrganId() (float64, error, string) {
	ss := new(OrganManage)
	total, err := Engine.Table("organ_manage").Where("organ_id =?", om.OrganId).Sum(ss, "recharge_num")
	sql := "SELECT COUNT(DISTINCT user_id) FROM `organ_manage` WHERE organ_id = ?"
	count, err := Engine.Query(sql, om.OrganId)
	var mm string
	for _, value := range count {
		mm = string(value["COUNT(DISTINCT user_id)"])
	}
	return total, err, mm
}
