package models

/*
 * Copyright (c) 2016—2017 www.fushow.cn, All rights reserved.
 */
import (
	"fmt"
	"strconv"
	"strings"
)

func (al *AuthorityList) GetNowSidebar() bool {
	Engine.ShowSQL(true)
	flag, _ := Engine.Where("id=?", al.Id).Get(al)
	return flag
}

//根据type查询后台侧边栏列表信息
//time 2016-11-19 txl
func (al *AuthorityList) GetSidebar(types uint64) []AuthorityList {
	var (
		alist []AuthorityList
		ut    TypeProcess
	)

	ut.TypeId = types

	//该权限没有可访问的sietbar
	if !ut.GetTypeList() {
		return alist
	}
	sid := strings.Split(ut.AuthorityListId, ",")
	for i := 0; i < len(sid); i++ {
		//没有更多的权限时跳出
		if sid[i] == "" {
			fmt.Println("kong")
			return alist
		}
		var now AuthorityList
		now.Id, _ = strconv.ParseInt(sid[i], 10, 64)
		//得到当前的sietbar
		fmt.Println("now.Id", now.Id)
		if !now.GetNowSidebar() {
			fmt.Println("bucun")
			//			return alist
		} else {
			//加入数组中
			alist = append(alist, now)
		}
	}
	//返回该权限所有的sitebar
	return alist
}
