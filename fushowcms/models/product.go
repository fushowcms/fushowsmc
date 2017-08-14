package models

/*
 * Copyright (c) 2016—2017 www.fushow.cn, All rights reserved.
 */
import (
	_ "github.com/go-sql-driver/mysql"
)

//添加产品
func (pd *Product) ProductAdd() bool {
	rows, err := Engine.Insert(pd) //插入一行数据
	if rows > 0 || err != nil {
		return true
	}
	return false
}

//删除产品
func (pd *Product) ProductDel() bool {
	num, err := Engine.Delete(pd) //插入一行数据
	if num > 0 || err != nil {
		return true
	}
	return false
}

//修改产品信息
func (pd *Product) ProductUp() bool {
	if num, err := Engine.Where("Id=?", pd.Id).Update(pd); num > 0 || err == nil {
		return true
	}
	return false
}

//获取产品信息
func (pd *Product) GetProduct(uid int64) ([]Product, error) {
	var product []Product
	err := Engine.Where("id=?", uid).Find(&product)
	return product, err
}

//获取产品信息
func (pd *Product) GetProducts() bool {
	flag, _ := Engine.Get(pd)
	if flag {
		return true
	}
	return false
}

//测试
//获取产品信息列表
func (pd *Product) GetProductList(page, rows int) ([]Product, int64) {
	var list []Product
	total, _ := Engine.Where("id >?", 0).Count(pd)
	Engine.Limit(rows, (page-1)*rows).Find(&list)
	return list, total
}

//获取产品信息列表
func (pd *Product) GetProductLists(page, rows int, inputid string) ([]Product, int64) {
	var list []Product
	total, _ := Engine.Where("id >?", 0).Count(pd)
	Engine.Where("product_name like ?", "%"+inputid+"%").Limit(rows, (page-1)*rows).Find(&list)
	return list, total
}

//添加期数产品过程
func (pd *PeriodsProduct) PeriodsProductAdd() bool {
	rows, err := Engine.Insert(pd) //插入一行数据
	if rows > 0 || err != nil {
		return true
	}
	return false
}

//删除期数产品过程
func (pd *PeriodsProduct) PeriodsProductDel() error {
	_, err := Engine.Delete(pd) //插入一行数据
	return err
}

//删除期数产品过程---->通过期iD
func (pd *PeriodsProduct) PeriodsProductDelById() error {
	_, err := Engine.Where("periods_id=?", pd.PeriodsId).Delete(pd) //插入一行数据
	return err
}

//修改期数产品信息过程
func (pd *PeriodsProduct) PeriodsProductUp() (int64, error) {
	var err error
	Engine.ShowSQL(true)
	if num, err := Engine.Where("Id=?", pd.Id).Update(pd); num > 0 && err == nil {
		return 1, err
	}
	return -1, err
}

//获取期数产品信息过程
func (pd *PeriodsProduct) GetPeriodsProduct() (int64, error) {
	_, err := Engine.Get(pd)
	return 1, err
}

//获取期数产品信息过程_cnxulin
func (pd *PeriodsProduct) GetPeriodsProducts() bool {
	has, _ := Engine.Get(pd)
	return has
}

//获得过程表
func (pd *PeriodsProduct) GetProducts() bool {
	has, _ := Engine.Where("periods_id=? and product_id=?", pd.PeriodsId, pd.ProductId).Get(pd)
	return has
}

//获取某一期产品过程信息列表
func (pd *PeriodsProduct) GetPeriodsProductList() ([]PeriodsProduct, int64) {
	var list []PeriodsProduct
	total, _ := Engine.Where("periods_id=?", pd.PeriodsId).Count(pd)
	Engine.Where("periods_id=?", pd.PeriodsId).Find(&list)
	return list, total
}
