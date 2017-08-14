package control

/*
 * Copyright (c) 2016—2017 www.fushow.cn, All rights reserved.
 */
import (
	"fushowcms/comm"
	m "fushowcms/models"

	"github.com/gin-gonic/gin"
)

//网站信息设置
func SetSite(c *gin.Context) {
	var (
		ws m.Website
	)
	title := c.PostForm("title")
	keywords := c.PostForm("keywords")
	descrip := c.PostForm("descrip")
	name := c.PostForm("name")
	logo := c.PostForm("logo")
	icon := c.PostForm("icon")
	url := c.PostForm("url")

	ws.Title = title
	ws.Keywords = keywords
	ws.Descrip = descrip
	ws.Name = name
	ws.Logo = logo
	ws.Icon = icon
	ws.URL = url

	if !ws.SetSite() {
		comm.ResponseError(c, 4020)
		return
	}
	comm.Response(c, nil)
}

//获取网站信息
func GetSite(c *gin.Context) {
	var (
		ws m.Website
	)
	if !ws.GetSite() {
		comm.ResponseError(c, 4020)
		return
	}
	comm.Response(c, ws)
}

//平台收益情况
func GetWebEarnings(c *gin.Context) {
	sDate := c.PostForm("startDate")
	eDate := c.PostForm("endDate")
	site := m.GetWebEarnings(sDate, eDate)
	comm.Response(c, site)
}

//运营总数据
func GetWebInfoAll(c *gin.Context) {
	wb := m.GetWebInfoAll()
	comm.Response(c, wb)
}

//平台运营信息
func RegNumber(c *gin.Context) {
	sDate := c.PostForm("startDate")
	eDate := c.PostForm("endDate")
	site := m.RegNumber(sDate, eDate)
	comm.Response(c, site)
}
