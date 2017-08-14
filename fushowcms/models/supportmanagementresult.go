package models

/*
 * Copyright (c) 2016—2017 www.fushow.cn, All rights reserved.
 */
import (
	_ "github.com/go-sql-driver/mysql"
)

/*iscunzai*/
func (smt *SupportManagementResult) IsHere() bool {
	flag, _ := Engine.Get(smt)
	return flag
}
