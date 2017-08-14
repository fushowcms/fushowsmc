package control

/*
 * Copyright (c) 2016—2017 www.fushow.cn, All rights reserved.
 */
import (
	m "fushowcms/models"
	"time"

	"fushowcms/comm"
	"strconv"

	"github.com/gin-gonic/gin"
)

//获取活动列表
func GetEventList(c *gin.Context) {
	var ev m.Event
	page, _ := strconv.Atoi(c.PostForm("page"))
	rows, _ := strconv.Atoi(c.PostForm("rows"))
	list, row := ev.GetApplyList(page, rows)
	if row == 0 {
		return
	}
	c.JSON(200, gin.H{"total": row, "rows": list})
}

//添加活动
func EventAdd(c *gin.Context) {
	var (
		ev m.Event
	)
	name := c.PostForm("EventName")
	stime := c.PostForm("StartTime")
	etime := c.PostForm("EndTime")
	allnumber, _ := strconv.ParseInt(c.PostForm("AllNumber"), 10, 64)
	userid, _ := strconv.ParseInt(c.PostForm("UID"), 10, 64)
	number, _ := strconv.ParseInt(c.PostForm("Number"), 10, 64)
	eventtype, _ := strconv.ParseInt(c.PostForm("EventType"), 10, 64)
	if userid == 0 {
		comm.ResponseError(c, 3175)
		return
	}
	if name == "" {
		comm.ResponseError(c, 3114) //活动名称不能为空
		return
	}
	if number < 0 {
		comm.ResponseError(c, 3115) //赠送石榴籽不能为负
		return
	}
	ev.EventName = name
	ev.StartTime = stime
	ev.EndTime = etime
	ev.AllNumber = allnumber
	ev.Number = number
	ev.EventType = eventtype
	ev.NowState = 1
	flag, _ := ev.EventAddInfo(ev, userid)
	if !flag {
		comm.ResponseError(c, 3117) //活动添加失败
		return
	}
	comm.ResponseError(c, 3118) //活动添加成功
}

//删除活动
func EventDel(c *gin.Context) {
	var ev m.Event
	if c.PostForm("id") == "" {
		comm.ResponseError(c, 3119) //活动删除失败
		return
	}
	ev.Id, _ = strconv.ParseInt(c.PostForm("id"), 10, 64)
	if !ev.EventDel() {
		comm.ResponseError(c, 3119) //活动删除失败
		return
	}
	comm.ResponseError(c, 3120) //活动删除成功
}

//修改活动
func EventUp(c *gin.Context) {
	var ev m.Event
	name := c.PostForm("EventName")
	stime := c.PostForm("StartTime")
	etime := c.PostForm("EndTime")
	number, _ := strconv.ParseInt(c.PostForm("Number"), 10, 64)
	eventtype, _ := strconv.ParseInt(c.PostForm("EventType"), 10, 64)
	allNumber, _ := strconv.ParseInt(c.PostForm("AllNumber"), 10, 64)
	nowtime := time.Now().Format("2006-01-02 15:04:05")
	t1, err := time.Parse("2006-01-02 15:04:05", nowtime)
	t2, err := time.Parse("2006-01-02 15:04:05", etime)
	//判断活动时间已经过期
	if err == nil && t2.Before(t1) {
		comm.ResponseError(c, 3121) //活动时间已经结束
		return
	}
	if c.PostForm("Id") == "" {
		comm.ResponseError(c, 3122) //活动不存在
		return
	}
	if name == "" {
		comm.ResponseError(c, 3123) //活动名称不能为空
		return
	}
	if number < 0 {
		comm.ResponseError(c, 3124) //赠送石榴籽不能为负
		return
	}
	ev.EventName = name
	ev.StartTime = stime
	ev.EndTime = etime
	ev.Number = number
	ev.AllNumber = allNumber
	ev.EventType = eventtype
	ev.Id, _ = strconv.ParseInt(c.PostForm("Id"), 10, 64)
	if !ev.EventUp() {
		comm.ResponseError(c, 3125) //活动修改失败
		return
	}
	comm.ResponseError(c, 3126) //活动修改成功
}

//活动详情
func GetEventRecInfo(c *gin.Context) {
	etype, _ := strconv.ParseInt(c.PostForm("Type"), 10, 64)
	id, _ := strconv.ParseInt(c.PostForm("Id"), 10, 64)
	page, _ := strconv.Atoi(c.PostForm("page"))
	rows, _ := strconv.Atoi(c.PostForm("rows"))
	var eve_det m.EventRec
	eve_det.EventId = id
	eve_det.EventType = etype
	list, total := eve_det.GetEventRecList(page, rows)
	if total == 0 {
		return
	}
	c.JSON(200, gin.H{"total": total, "rows": list})
}

//活动截止
func EventOver(c *gin.Context) {
	var e m.Event
	evid, _ := strconv.ParseInt(c.PostForm("Id"), 10, 64)
	uid, _ := strconv.ParseInt(c.PostForm("UID"), 10, 64)
	if uid == 0 {
		comm.ResponseError(c, 3175)
		return
	}
	if c.PostForm("Id") == "" {
		comm.ResponseError(c, 3127) //参数错误
		return
	}
	flag, mes := e.EventOver(evid, uid)
	if !flag {
		comm.ResponseError(c, 3128) //截止活动失败
		c.JSON(200, gin.H{"state": false, "errCode ": 2002, "errMsg": mes})
		return
	}
	comm.ResponseError(c, 3129) //活动已结束
}

//判断绑定联盟活动是否存在
func GetEventInfo(c *gin.Context) {
	var e m.Event
	nowtime := time.Now().Format("2006-01-02 15:04:05")
	_, number, _ := e.GetEventInfos(nowtime)
	m := make(map[string]interface{})
	m["state"] = true
	m["Number"] = number
	comm.Response(c, m)
}
