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

//查询广播信息 addby liuhan
func FindBroadcast(c *gin.Context) {
	var b m.Broadcast
	data, err := b.FindBroadcast()
	if err != nil && data == nil {
		c.JSON(200, gin.H{"state": false, "data": nil})
		return
	}
	c.JSON(200, gin.H{"state": true, "data": data})
}

//增加广播信息 addby liuhan
func AddBroadcast(c *gin.Context) {
	var b m.Broadcast
	userId, _ := strconv.ParseInt(c.PostForm("UID"), 10, 64)
	content := c.PostForm("Content")
	if userId == 0 {
		c.JSON(200, gin.H{"state": false, "msg": "3002"}) //3002：用户id为空
		return
	}
	if content == "" {
		c.JSON(200, gin.H{"state": false, "msg": "3020"}) //3020：内容为空
		return
	}
	b.UserId = userId
	b.Content = content
	if !b.AddBroadcast() {
		c.JSON(200, gin.H{"state": "fail"})
		return
	}
	c.JSON(200, gin.H{"state": true})
}

//增加广播信息redis addby liuhan
func AddBroadcastRedis(c *gin.Context) {
	var b m.Broadcast
	userId, _ := strconv.ParseInt(c.PostForm("UID"), 10, 64)
	if userId == 0 {
		comm.ResponseError(c, 3175)
		return
	}
	content := c.PostForm("Content")
	b.Id = time.Now().Unix()
	b.UserId = userId
	b.Content = content
	if SetBroadcastNew(b, c) != "" {
		return
	}
	comm.ResponseError(c, 3190) //全频道广播发送成功
}
