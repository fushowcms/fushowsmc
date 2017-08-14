package models

/*
 * Copyright (c) 2016—2017 www.fushow.cn, All rights reserved.
 */
import (
	_ "github.com/go-sql-driver/mysql"
)

//添加轮播图
func (cp *CarouselPic) CarPicAdd() bool {
	row, err := Engine.Insert(cp) //插入一行数据
	if row > 0 || err == nil {
		return true
	}
	return false
}

//获取轮播图信息
func (cp *CarouselPic) GetCarPic() bool {
	_, err := Engine.Get(cp)
	if err == nil {
		return true
	}
	return false
}

//删除轮播图
func (cp *CarouselPic) CarPicDel() bool {
	row, err := Engine.Delete(cp) //删除一行数据
	if row > 0 || err == nil {
		return true
	}
	return false
}

//获取轮播图列表(总表)
func (cp *CarouselPic) GetCarPicList(page, rows int) ([]CarouselPic, int64) {
	var list []CarouselPic
	everyone := make([]CarouselPic, 0)
	Engine.Find(&everyone)
	total, _ := Engine.Where("id >?", 0).Count(cp)
	Engine.Limit(rows, (page-1)*rows).Find(&list)
	return list, total
}

//修改轮播图
func (cp *CarouselPic) CarPicUp() bool {
	Engine.ShowSQL(true)
	if num, err := Engine.Where("Id=?", cp.Id).Update(cp); num > 0 || err == nil {
		return true
	}
	return false
}

/**
 * 结构体名称：GetIndexCarPicList,查询轮播数据
 * param：page,rows(页数，行数)
 * @author 徐林->修正
 * @Time 2016-10-31
 */
func (cp *CarouselPic) GetIndexCarPicList(page, rows int) ([]CarouselPic, int64) {
	var list []CarouselPic
	total, _ := Engine.Table("carousel_pic").Where("id >?", 0).Count(cp)
	Engine.Table("carousel_pic").Alias("c").Desc("c.id").Limit(rows, (page-1)*rows).Find(&list)
	Engine.ShowSQL(true)
	return list, total
}
func (cp *CarouselPic) GetIndexCarPicLists(page, rows int) ([]CarouselPic, int64) {
	var list []CarouselPic
	Engine.ShowSQL(true)
	total, _ := Engine.Table("carousel_pic").Where("state = 1").Count(cp)
	Engine.Table("carousel_pic").Where("state = 1").Find(&list)
	Engine.ShowSQL(true)
	return list, total
}

/**
 * 结构体名称：GetDataUpdate,查询轮播数据并更新显示状态
 * @author 徐林->新增
 * @Time 2016-11-11
 */
func (cp *CarouselPic) GetDataUpdate() bool {
	sql := "update carousel_pic set state=0 where carousel_type = 1 and state = 1"
	_, err := Engine.Exec(sql)
	if err != nil {
		return false
	}
	return true
}
