package models

/*
 * Copyright (c) 2016—2017 www.fushow.cn, All rights reserved.
 */
import (
	_ "github.com/go-sql-driver/mysql"
)

//查询广播信息 addby liuhan
func (b *Broadcast) FindBroadcast() ([]Broadcast, error) {
	var list []Broadcast
	err := Engine.Table("broadcast").Find(&list)
	Engine.ShowSQL(true)
	return list, err
}

//增加广播信息 addby liuhan
func (b *Broadcast) AddBroadcast() bool {
	row, err := Engine.Insert(b) //插入一行数据
	if row > 0 || err == nil {
		return true
	}
	return false
}
