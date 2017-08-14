package control

/*
 * Copyright (c) 2016—2017 www.fushow.cn, All rights reserved.
 */
import (
	m "fushowcms/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

//权限列表
func GetUidTypeList(c *gin.Context) {
	var (
		ut m.UidType
	)
	page, _ := strconv.Atoi(c.PostForm("page"))
	rows, _ := strconv.Atoi(c.PostForm("rows"))
	total, list := ut.GetUidTypeList(page, rows)
	if len(list) <= 0 {
		return
	}
	c.JSON(200, gin.H{"total": total, "rows": list})
}

//人物权限添加
func UidTypeAdd(c *gin.Context) {
	var (
		ut m.UidType
	)
	ut.Id, _ = strconv.ParseUint(c.PostForm("Id"), 10, 64)

	if c.PostForm("TypeName") == "" {
		c.JSON(200, gin.H{"state": false, "message": "权限名称不能为空"})
		return
	}

	if ut.IsUidType() {
		c.JSON(200, gin.H{"state": false, "message": "该权限Id已存在"})
		return
	}
	ut.TypeName = c.PostForm("TypeName")

	if !ut.UidTypeAdd() {
		c.JSON(200, gin.H{"state": false, "message": "添加失败"})
		return
	}
	c.JSON(200, gin.H{"state": true, "message": "添加成功"})
}

//人物权限修改
func UidTypeUp(c *gin.Context) {
	var (
		ut m.UidType
	)
	ut.Id, _ = strconv.ParseUint(c.PostForm("Id"), 10, 64)

	if !ut.IsUidType() {
		c.JSON(200, gin.H{"state": false, "message": "该权限不存在"})
		return
	}

	ut.TypeName = c.PostForm("TypeName")
	if !ut.UidTypeUp() {
		c.JSON(200, gin.H{"state": false, "message": "添加失败"})
		return
	}
	c.JSON(200, gin.H{"state": true, "message": "修改成功"})
}

//sitebar 数据列表
func GetAuthorityList(c *gin.Context) {
	var al m.AuthorityList
	list := al.GetAuthorityList()
	c.JSON(200, gin.H{"data": list})
}

//权限过程表
func TypeProcessAdd(c *gin.Context) {
	var (
		tp  m.TypeProcess
		uid m.UidInfo
	)

	//操作人
	tp.Uid, _ = strconv.ParseInt(c.PostForm("UID"), 10, 64)

	if tp.Uid == 0 {
		c.JSON(200, gin.H{"state": false, "errMes": "操作人不存在"})
		return
	}
	uid.Id = tp.Uid

	if !uid.GetUserInfo() {
		c.JSON(200, gin.H{"state": false, "errMes": "操作人不存在"})
		return
	}

	if uid.Type < 254 {
		c.JSON(200, gin.H{"state": false, "errMes": "您没有该权限"})
		return
	}

	//TypeId
	tp.TypeId, _ = strconv.ParseUint(c.PostForm("TypeId"), 10, 64)
	//不存在时  add
	if !tp.GetTypeList() {
		//没有记录
		tp.AuthorityListId = c.PostForm("list")

		if !tp.TypeProcessAdd() {
			c.JSON(200, gin.H{"state": false, "errMes": "添加失败"})
			return
		}

		c.JSON(200, gin.H{"state": true, "errMes": "添加成功"})
		return
	}

	tp.AuthorityListId = c.PostForm("list")

	if !tp.TypeProcessUpAu() {
		c.JSON(200, gin.H{"state": false, "errMes": "错误修改"})
		return
	}

	c.JSON(200, gin.H{"state": true, "errMes": "修改成功"})

}

func GetMyAuList(c *gin.Context) {
	var ut m.TypeProcess

	ut.TypeId, _ = strconv.ParseUint(c.PostForm("TypeId"), 10, 64)

	if !ut.GetTypeList() {
		c.JSON(200, gin.H{"state": false, "data": ""})
		return
	}

	c.JSON(200, gin.H{"state": true, "data": ut})
}
