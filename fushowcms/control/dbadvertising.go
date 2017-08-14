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
func DbadvertisingAdd(c *gin.Context) {
	_, filepath := Uploads(c)
	dbadURL := c.PostForm("DbadURL")           //广告链接
	dbadName := c.PostForm("DbadName")         //广告名称
	dbadBriefing := c.PostForm("DbadBriefing") //广告简介
	var ar m.Dbadvertising
	ar.PicURL = filepath
	ar.DbadURL = dbadURL
	ar.DbadName = dbadName
	ar.DbadBriefing = dbadBriefing
	ar.LiveState = 1
	ar.StartTime = time.Now().Unix() //增加时间
	if !ar.DbadvertisingAdd() {
		comm.ResponseError(c, 3145) //官方活动增加失败
		return
	}
	comm.ResponseError(c, 3146) //官方活动增加成功
}

//删除直播间广告
func DbadvertisingDel(c *gin.Context) {
	var ar m.Dbadvertising
	dbadvertisingid, _ := strconv.ParseInt(c.PostForm("id"), 10, 64) //广告编号
	ar.Id = dbadvertisingid
	if !ar.DbadvertisingDel() {
		comm.ResponseError(c, 3147) //广告删除失败
		return
	}
	comm.ResponseError(c, 3148) //广告删除成功
}

//修改直播间广告
func DbadvertisingUp(c *gin.Context) {
	var ar m.Dbadvertising
	var am m.Dbadvertising
	_, filepath := Uploads(c)
	dbadid, _ := strconv.ParseInt(c.PostForm("myuid"), 10, 64) //活动编号
	picURL := filepath                                         //活动图片
	dbadURL := c.PostForm("DbadURL")                           //活动链接
	dbadName := c.PostForm("DbadName")                         //活动名称
	dbadBriefing := c.PostForm("DbadBriefing")                 //活动简介
	ar.Id = dbadid
	am, err := ar.GetDbad()
	if err != nil {
		comm.ResponseError(c, 3149) //官方活动不存在
		return
	}
	am.PicURL = picURL
	am.DbadURL = dbadURL
	am.DbadName = dbadName
	am.DbadBriefing = dbadBriefing
	if !am.DbadvertisingUp() {
		comm.ResponseError(c, 3150) //官方活动修改失败
		return
	}
	comm.ResponseError(c, 3151) //官方活动修改成功
}

//获取所有广告列表
func GetDbadvertisinglist(c *gin.Context) {
	var ar m.Dbadvertising
	page, _ := strconv.Atoi(c.PostForm("page"))
	rows, _ := strconv.Atoi(c.PostForm("rows"))
	list, row := ar.GetDbadvertisinglist(page, rows)
	c.JSON(200, gin.H{"total": row, "rows": list})
}

//获取直播间广告
func GetDbadvertising(c *gin.Context) {
	var ar m.Dbadvertising
	ar.LiveState = 1
	list, err := ar.GetDbadvertising()
	if err != nil {
		comm.ResponseError(c, 3152) //广告不存在
		return
	}
	comm.Response(c, list)

}
