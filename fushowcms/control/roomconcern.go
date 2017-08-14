package control

/*
 * Copyright (c) 2016—2017 www.fushow.cn, All rights reserved.
 */
import (
	"fmt"
	"fushowcms/comm"
	m "fushowcms/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

//添加一条关注记录
func RoomConcernAdd(c *gin.Context) {
	var (
		ar m.AnchorRoomConcern
		ui m.UidInfo
	)
	ar.Uid, _ = strconv.ParseInt(c.PostForm("UID"), 10, 64) //关注者
	if ar.Uid == 0 {
		comm.ResponseError(c, 3175)
		return
	}
	ar.User, _ = strconv.ParseInt(c.PostForm("User"), 10, 64) //被关注者（主播）
	ui.Id = ar.Uid

	//关注者不存在
	if !ui.GetUserInfo() {
		comm.ResponseError(c, 2034)
		return
	}

	if ar.IsConcern() {
		comm.ResponseError(c, 2035)
		return
	}

	if !ar.RoomAdd() {
		comm.ResponseError(c, 2036)
		return
	}
	comm.Response(c, "关注成功")
}

//txl-del
//删除一条关注记录
func RoomConcernDel(c *gin.Context) {
	var ar m.AnchorRoomConcern
	ar.Id, _ = strconv.ParseInt(c.PostForm("UID"), 10, 64) //ID
	if ar.Id == 0 {
		comm.ResponseError(c, 3175)
		return
	}
	if !ar.RoomDel() {
		c.JSON(200, gin.H{"state": "fail"})
		return
	}
	c.JSON(200, gin.H{"state": "success"})
}

//txl-del
//获取单条关注记录
func GetRoomConcern(c *gin.Context) {
	var ar m.AnchorRoomConcern
	ar.Id, _ = strconv.ParseInt(c.PostForm("UID"), 10, 64) //ID
	if ar.Id == 0 {
		comm.ResponseError(c, 3175)
		return
	}
	if !ar.GetRoom() {
		c.JSON(200, gin.H{"state": "fail"})
		return
	}
	c.JSON(200, gin.H{"state": ar})
}

//是否已经关注
func IsConcern(c *gin.Context) {
	var ar m.AnchorRoomConcern
	fmt.Println("tjoso ")
	ar.Uid, _ = strconv.ParseInt(c.PostForm("UID"), 10, 64)
	if ar.Uid == 0 {
		comm.ResponseError(c, 3175)
		return
	}
	ar.User, _ = strconv.ParseInt(c.PostForm("User"), 10, 64)
	if ar.IsConcern() {
		comm.ResponseError(c, 2037)
		return
	}
	comm.Response(c, "未关注")
}

//txl-del
//获取所有关注记录
func GetRoomConcernList(c *gin.Context) {
	var ar m.AnchorRoomConcern
	page, _ := strconv.Atoi(c.PostForm("page"))
	rows, _ := strconv.Atoi(c.PostForm("rows"))
	list, row := ar.GetRoomList(page, rows)
	if row == 0 {
		return
	}
	c.JSON(200, gin.H{"total": row, "rows": list})
}

//txl-del
//获取某个人所有关注记录
func GetMyOrderRoomList(c *gin.Context) {
	var ar m.AnchorRoomConcern
	ar.Uid, _ = strconv.ParseInt(c.PostForm("UID"), 10, 64) //ID
	if ar.Uid == 0 {
		comm.ResponseError(c, 3175)
		return
	}
	fmt.Println("this is ar ", ar)
	list, err := ar.GetMyOrderRoomList()
	if err != nil {
		comm.ResponseError(c, 2012)
		return
	}
	comm.Response(c, list)
}

//查看我的关注列表
func GetMyAttention(c *gin.Context) {
	var (
		ar m.AnchorRoomConcern
	)

	//获取用户ID
	ar.Uid, _ = strconv.ParseInt(c.PostForm("UID"), 10, 64) //关注者ID
	if c.PostForm("UID") == "" {
		comm.ResponseError(c, 2000)
		return
	}
	//获取页数-行数
	page, _ := strconv.Atoi(c.PostForm("page"))
	rows, _ := strconv.Atoi(c.PostForm("rows"))

	attentionlist, err, total := ar.GetMyAttention(page, rows)

	if len(attentionlist) == 0 {
		comm.ResponseError(c, 2012)
		return
	}
	if err != nil {
		comm.ResponseError(c, 2039)
		return
	}

	m := make(map[string]interface{})
	m["state"] = attentionlist
	m["total"] = total
	comm.Response(c, m)
}

//取消关注
func CancelRoomCon(c *gin.Context) {
	var ar m.AnchorRoomConcern
	ar.Uid, _ = strconv.ParseInt(c.PostForm("UID"), 10, 64)
	ar.User, _ = strconv.ParseInt(c.PostForm("User"), 10, 64)
	if ar.Uid == 0 || ar.User == 0 {
		comm.ResponseError(c, 2000)
		return
	}
	if !ar.GetRoom() {
		comm.ResponseError(c, 2034)
		return
	}
	if !ar.RoomDel() {
		comm.ResponseError(c, 2040)
		return
	}
	comm.Response(c, "成功取消")

}

/*
*功能 查看我关注的正在直播的主播
*牟海龙20161013
 */
func IsOpenMyAttention(c *gin.Context) {
	var (
		ar m.AnchorRoomConcern
	)

	//获取用户ID
	ar.Uid, _ = strconv.ParseInt(c.PostForm("UID"), 10, 64) //关注者UID
	if c.PostForm("UID") == "" {
		comm.ResponseError(c, 2000)
		return
	}
	//获取页数-行数
	page, _ := strconv.Atoi(c.PostForm("page"))
	rows, _ := strconv.Atoi(c.PostForm("rows"))
	if !ar.GetRoom() {
		comm.ResponseError(c, 2028)
		return
	}
	attentionlist, err := ar.IsOpenGetMyAttention(page, rows)

	if len(attentionlist) == 0 {
		comm.ResponseError(c, 2012)
		return
	}
	if err != nil {
		comm.ResponseError(c, 2039)
		return
	}
	comm.Response(c, attentionlist)
}
