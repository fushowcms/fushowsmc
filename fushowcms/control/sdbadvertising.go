package control

/*
 * Copyright (c) 2016—2017 www.fushow.cn, All rights reserved.
 */
import (
	"fushowcms/comm"
	m "fushowcms/models"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

//添加直播间广告
func SdbadvertisingAdd(c *gin.Context) {
	_, filepath := Uploads(c)
	sdbadURL := c.PostForm("SdbadURL")           //广告链接
	sdbadName := c.PostForm("SdbadName")         //广告名称
	sdbadBriefing := c.PostForm("SdbadBriefing") //广告简介

	var ar m.Sdbadvertising
	ar.PicURL = filepath
	ar.SdbadURL = sdbadURL
	ar.SdbadName = sdbadName
	ar.SdbadBriefing = sdbadBriefing
	ar.LiveState = 1
	ar.StartTime = time.Now().Unix() //增加时间
	if !ar.SdbadvertisingAdd() {
		comm.ResponseError(c, 3145) //增加失败
		return
	}
	comm.ResponseError(c, 3146) //增加成功
}

//删除直播间广告
func SdbadvertisingDel(c *gin.Context) {
	var ar m.Sdbadvertising
	sdbadvertisingid, _ := strconv.ParseInt(c.PostForm("id"), 10, 64) //广告编号

	ar.Id = sdbadvertisingid
	if !ar.SdbadvertisingDel() {
		comm.ResponseError(c, 3147) //广告删除失败
		return
	}
	comm.ResponseError(c, 3148) //广告删除成功
}

//修改直播间广告
func SdbadvertisingUp(c *gin.Context) {
	var ar m.Sdbadvertising
	var am m.Sdbadvertising

	_, filepath := Uploads(c)
	sdbadid, _ := strconv.ParseInt(c.PostForm("myuid"), 10, 64) //活动编号

	picURL := filepath                           //活动图片
	sdbadURL := c.PostForm("SdbadURL")           //活动链接
	sdbadName := c.PostForm("SdbadName")         //活动名称
	sdbadBriefing := c.PostForm("SdbadBriefing") //活动简介

	ar.Id = sdbadid
	am, err := ar.GetSdbad()
	if err != nil {
		comm.ResponseError(c, 3149) //官方活动不存在
		return
	}
	am.PicURL = picURL
	am.SdbadURL = sdbadURL
	am.SdbadName = sdbadName
	am.SdbadBriefing = sdbadBriefing
	if !am.SdbadvertisingUp() {
		comm.ResponseError(c, 3150) //官方活动修改失败
		return
	}
	comm.ResponseError(c, 3151) //官方活动修改成功
}

//获取所有广告列表
func GetSdadvertisinglist(c *gin.Context) {
	var ar m.Sdbadvertising
	page, _ := strconv.Atoi(c.PostForm("page"))
	rows, _ := strconv.Atoi(c.PostForm("rows"))
	list, row := ar.GetSdadvertisinglist(page, rows)
	c.JSON(200, gin.H{"total": row, "rows": list})
}

//获取直播间广告
func GetSdbadvertising(c *gin.Context) {
	var ar m.Sdbadvertising
	ar.LiveState = 1
	list, err := ar.GetSdbadvertising()
	if err != nil {
		comm.ResponseError(c, 3152) //广告不存在
		return
	}
	comm.Response(c, list)

}
