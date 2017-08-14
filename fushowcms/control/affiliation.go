package control

/*
 * Copyright (c) 2016—2017 www.fushow.cn, All rights reserved.
 */
import (
	"fmt"
	m "fushowcms/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

//添加一个机构
//创建时间 2016-10-13 Txl
//AffId				机构编码-->8位由前台输入
//InstitutionName   机构名称
func AffiliationAdd(c *gin.Context) {
	var aff m.Affiliation
	aff.AffId = c.PostForm("AffId")
	aff.InstitutionName = c.PostForm("InstitutionName")
	if aff.AffId == "" || aff.InstitutionName == "" {
		c.JSON(200, gin.H{"state": false, "message": "参数错误"})
		return
	}
	if len(aff.AffId) != 8 {
		c.JSON(200, gin.H{"state": false, "message": "机构编码位数错误,请重新输入"})
		return
	}
	if !aff.AffiliationAdd() {
		c.JSON(200, gin.H{"state": false, "message": "添加错误,机构编码不可重复"})
		return
	}
	c.JSON(200, gin.H{"state": true, "message": "恭喜您成功添加" + aff.InstitutionName})
}

//删除一个机构----->通过AffId删除
//创建时间 2016-10-13 Txl
//AffId				机构编码-->8位由前台输入
func AffiliationDel(c *gin.Context) {
	var aff m.Affiliation
	aff.Id, _ = strconv.ParseInt(c.PostForm("id"), 10, 64)
	if c.PostForm("id") == "" {
		c.JSON(200, gin.H{"state": false, "message": "参数错误"})
		return
	}
	if !aff.AffiliationDel() {
		c.JSON(200, gin.H{"state": false, "message": "机构编码错误"})
		return
	}
	c.JSON(200, gin.H{"state": true, "message": "成功删除"})
}

//修改机构信息
//创建时间 2016-10-13 Txl
//AffId				机构编码-->8位由前台输入
//InstitutionName   机构名称
func AffiliationUp(c *gin.Context) {
	var aff m.Affiliation
	aff.AffId = c.PostForm("AffId")
	institutionName := c.PostForm("InstitutionName")
	fmt.Println("ssss", aff.AffId, institutionName)
	if aff.AffId == "" || institutionName == "" {
		c.JSON(200, gin.H{"state": false, "message": "参数错误"})
		return
	}
	if !aff.GetAffiliation() {
		c.JSON(200, gin.H{"state": false, "message": "机构不存在"})
		return
	}
	aff.InstitutionName = institutionName
	if !aff.AffiliationUpdate() {
		c.JSON(200, gin.H{"state": false, "message": "更新失败"})
		return
	}
	c.JSON(200, gin.H{"state": true, "message": "更新成功"})
}

//获取机构信息列表
//创建时间 2016-10-13 Txl
//无参
//返回数据   total:列表条数    rows:机构信息列表数据
func GetAffiliationList(c *gin.Context) {
	var aff m.Affiliation
	page, _ := strconv.Atoi(c.PostForm("page"))
	rows, _ := strconv.Atoi(c.PostForm("rows"))
	list, row := aff.GetAffiliationList(page, rows)
	fmt.Println(row)
	if row == 0 {
		return
	}
	c.JSON(200, gin.H{"total": row, "rows": list})
}

func GetAllAffiliation(c *gin.Context) {
	var aff m.Affiliation
	list, row := aff.GetAllAffiliation()
	if row == 0 {
		c.JSON(200, gin.H{"state": false, "data": ""})
		return
	}
	c.JSON(200, gin.H{"state": true, "data": list})
}
