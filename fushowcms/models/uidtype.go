package models

/*
 * Copyright (c) 2016—2017 www.fushow.cn, All rights reserved.
 */
//添加人物权限
func (ut *UidType) UidTypeAdd() bool {
	row, err := Engine.Insert(ut) //插入一行数据
	if row > 0 || err == nil {
		return true
	}
	return false
}

//删除人物权限
func (ut *UidType) UidTypeDel() bool {
	row, err := Engine.Delete(ut) //插入一行数据
	if row > 0 || err == nil {
		return true
	}
	return false
}

//更新人物权限
func (ut *UidType) UidTypeUp() bool {
	row, err := Engine.Where("Id=?", ut.Id).Update(ut)
	if row > 0 && err == nil {
		return true
	}
	return false
}

//获取人物权限
func (ut *UidType) GetUidType() bool {
	flag, _ := Engine.Get(ut)
	return flag
}

//判断该权限是否存在
func (ut *UidType) IsUidType() bool {
	flag, _ := Engine.Where("id=?", ut.Id).Get(ut)
	return flag
}

//人物权限列表
func (ut *UidType) GetUidTypeList(page, rows int) (int64, []UidType) {
	var list []UidType
	total, _ := Engine.Where("id>?", 0).Count(ut)
	Engine.Where("id >?", 0).Limit(rows, (page-1)*rows).Find(&list)
	return total, list
}

//添加人物权限
func (ut *TypeProcess) TypeProcessAdd() bool {
	row, err := Engine.Insert(ut) //插入一行数据
	if row > 0 || err == nil {
		return true
	}
	return false
}

//删除人物权限
func (ut *TypeProcess) TypeProcessDel() bool {
	row, err := Engine.Delete(ut) //插入一行数据
	if row > 0 || err == nil {
		return true
	}
	return false
}

//更新人物权限
func (ut *TypeProcess) TypeProcessUp() bool {
	row, err := Engine.Where("Id=?", ut.Id).Update(ut)
	if row > 0 && err == nil {
		return true
	}
	return false
}

//重复申请时更改state状态
func (ut *TypeProcess) TypeProcessUpAu() bool {
	num, err := Engine.Where("id=?", ut.Id).Cols("authority_list_id").Update(ut)
	if num > 0 && err == nil {
		return true
	}
	return false
}

//获取单个人物权限
func (ut *TypeProcess) GetTypeProcess() bool {
	flag, _ := Engine.Get(ut)
	return flag
}

//获得typeid的值
func (ut *TypeProcess) GetTypeList() bool {
	flag, _ := Engine.Where("type_id=?", ut.TypeId).Get(ut)
	return flag
}

//获取sitebar 列表数据
func (ut *AuthorityList) GetAuthorityList() []AuthorityList {
	var list []AuthorityList
	Engine.Where("id>?", 0).Find(&list)
	return list
}
