package control

/*
 * Copyright (c) 2016—2017 www.fushow.cn, All rights reserved.
 */
import (
	m "fushowcms/models"
	"strconv"
	"time"

	comm "fushowcms/comm"

	"github.com/gin-gonic/gin"
)

//签到
//liuhan 20161202
func SigninAdd(c *gin.Context) {
	var (
		uid m.UidInfo
		ss  m.SignIn
		ev  m.Event
	)
	uid.Id, _ = strconv.ParseInt(c.PostForm("UID"), 10, 64)
	if c.PostForm("UID") == "" {
		comm.ResponseError(c, 3187) //参数错误
		return
	}
	flag, _ := ss.SigninGetNumber(uid.Id)
	if !flag {
		comm.ResponseError(c, 3189) //活动已结束
		return
	}

	ev.EventType = 1
	nowtime := time.Now().Format("2006-01-02 15:04:05")
	_, number, _ := ev.GetEventInfo(nowtime)
	mm := make(map[string]interface{})
	mm["errMsg"] = strconv.FormatInt(number, 10)
	comm.Response(c, mm)
}

//签到记录
//TODO 页面没用到
func GetSignInList(c *gin.Context) {
	var si m.SignIn
	uid, _ := strconv.ParseInt(c.PostForm("UID"), 10, 64)
	if c.PostForm("UID") == "" {
		c.JSON(200, gin.H{"state": false, "errCode ": 2014, "errMsg": "参数错误"})
		return
	}

	flag, mes, list := si.GetUserSigninList(uid)
	if !flag {
		c.JSON(200, gin.H{"state": false, "errCode ": 2015, "errMsg": mes})
		return
	}
	c.JSON(200, gin.H{"state": true, "errCode ": 2000, "errMsg": list})
}

//判断是否签到过
//liuhan 20161202
func IsSigned(c *gin.Context) {
	var si m.SignIn
	uid, _ := strconv.ParseInt(c.PostForm("UID"), 10, 64)
	if uid <= 0 {
		comm.ResponseError(c, 9001) // 参数错误
		return
	}
	flag := si.IsSignInInfo(uid, time.Now().Format("2006-01-02 00:00:00"))
	mm := make(map[string]interface{})
	mm["flag"] = flag
	comm.Response(c, mm)
}
