package control

/*
 * Copyright (c) 2016—2017 www.fushow.cn, All rights reserved.
 */
import (
	"fmt"
	"fushowcms/comm"
	"github.com/gin-gonic/gin"
	m "fushowcms/models"
	"strconv"
	"time"
)

//添加官方活动
func OfficialAdd(c *gin.Context) {
	_, filepath := Uploads(c)
	officialURL := c.PostForm("OfficialURL")           //活动链接
	officialName := c.PostForm("OfficialName")         //活动名称
	officialBriefing := c.PostForm("OfficialBriefing") //活动简介
	var ar m.OfficialFunctions
	ar.PicURL = filepath
	ar.OfficialURL = officialURL
	ar.OfficialName = officialName
	ar.OfficialBriefing = officialBriefing
	ar.LiveState = 1
	ar.StartTime = time.Now().Unix() //增加时间
	if !ar.OfficialAdd() {
		comm.ResponseError(c, 3145) //官方活动增加失败
		return
	}
	comm.ResponseError(c, 3146) //官方活动增加成功
}

//删除官方活动
func OfficialDel(c *gin.Context) {
	var ar m.OfficialFunctions
	officialid, _ := strconv.ParseInt(c.PostForm("id"), 10, 64) //活动编号

	ar.Id = officialid
	if !ar.OfficialDel() {
		comm.ResponseError(c, 3147) //官方活动删除失败
		return
	}
	comm.ResponseError(c, 3148) //官方活动删除成功
}

//修改官方活动
func OfficialUp(c *gin.Context) {
	var ar m.OfficialFunctions
	var am m.OfficialFunctions
	fmt.Println(c.PostForm("myuid"))
	_, filepath := Uploads(c)
	officialid, _ := strconv.ParseInt(c.PostForm("myuid"), 10, 64) //活动编号
	picURL := filepath                                             //活动图片
	officialURL := c.PostForm("OfficialURL")                       //活动链接
	officialName := c.PostForm("OfficialName")                     //活动名称
	officialBriefing := c.PostForm("OfficialBriefing")             //活动简介
	ar.Id = officialid
	am, err := ar.GetOfficial()
	if err != nil {
		comm.ResponseError(c, 3149) //官方活动不存在
		return
	}
	am.PicURL = picURL
	am.OfficialURL = officialURL
	am.OfficialName = officialName
	am.OfficialBriefing = officialBriefing
	if !am.OfficialUp() {
		comm.ResponseError(c, 3150) //官方活动修改失败
		return
	}
	comm.ResponseError(c, 3151) //官方活动修改成功
}

//获取单个官方活动
func GetOfficial(c *gin.Context) {
	var ar m.OfficialFunctions
	officialid, _ := strconv.ParseInt(c.PostForm("officialid"), 10, 64)
	ar.Id = officialid
	list, err := ar.GetOfficial()
	if err != nil {
		c.JSON(200, gin.H{"state": "未查询到"})
		return
	}
	c.JSON(200, gin.H{"state": list})
}

//获取所有活动列表
func GetOfficialList(c *gin.Context) {
	var ar m.OfficialFunctions
	page, _ := strconv.Atoi(c.PostForm("page"))
	rows, _ := strconv.Atoi(c.PostForm("rows"))
	list, row := ar.GetOfficialList(page, rows)
	c.JSON(200, gin.H{"total": row, "rows": list})
}

//获取显示活动信息
func GetStartOfficialList(c *gin.Context) {
	var ar m.OfficialFunctions
	ar.LiveState = 1
	list, err := ar.GetStartOfficialList()
	if err != nil {
		comm.ResponseError(c, 3152) //官方活动不存在
		return
	}
	comm.Response(c, list)
}
