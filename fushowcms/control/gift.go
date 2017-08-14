package control

/*
 * Copyright (c) 2016—2017 www.fushow.cn, All rights reserved.
 */
import (
	"fmt"
	"fushowcms/comm"
	"fushowcms/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

//添加礼物
func GiftAdd(c *gin.Context) {
	_, filepath := Uploads(c)
	_, filepath_p := Uploads_png(c)
	var gift models.Gift
	gift.GiftName = c.PostForm("GiftName")                                //礼物名称
	gift.GiftType, _ = strconv.ParseInt(c.PostForm("GiftType"), 10, 64)   //	礼物种类
	gift.GiftAccount = c.PostForm("GiftAccount")                          //	礼物描述
	gift.BuyNumber, _ = strconv.ParseInt(c.PostForm("BuyNumber"), 10, 64) //	购买所需石榴籽数
	gift.ToNumber, _ = strconv.ParseInt(c.PostForm("ToNumber"), 10, 64)   //可兑换石榴籽数
	gift.GiftPicture = filepath                                           //	礼物图片
	gift.State, _ = strconv.ParseInt(c.PostForm("State"), 10, 64)         //	状态
	gift.GiftPicStatic = filepath_p                                       //礼物图片 静态
	if !gift.GiftAdd() {
		comm.ResponseError(c, 3130) //礼物添加失败
		return
	}
	comm.ResponseError(c, 3131) //礼物添加成功
}

//删除礼物
func GiftDel(c *gin.Context) {
	Flag = Flag + 1
	var gift models.Gift
	gift.Id, _ = strconv.ParseInt(c.PostForm("id"), 10, 64)
	if !gift.GiftDel() {
		comm.ResponseError(c, 3132) //礼物删除失败
		return
	}
	comm.ResponseError(c, 3133) //礼物删除成功
}

//修改礼物信息
func GiftUp(c *gin.Context) {
	_, filepath := Uploads(c)
	_, filepath_p := Uploads_png(c)
	var gift models.Gift
	gift.Id, _ = strconv.ParseInt(c.PostForm("myuid"), 10, 64) //根据ID修改礼物信息
	if !gift.GetGift() {
		comm.ResponseError(c, 3134) //礼物不存在
		return
	}
	gift.GiftName = c.PostForm("GiftName") //礼物名称
	fmt.Println("giftname", gift.GiftName)
	fmt.Println("giftname", c.PostForm("GiftName"))
	gift.GiftType, _ = strconv.ParseInt(c.PostForm("GiftType"), 10, 64)   //	礼物种类
	gift.GiftAccount = c.PostForm("GiftAccount")                          //	礼物描述
	gift.BuyNumber, _ = strconv.ParseInt(c.PostForm("BuyNumber"), 10, 64) //	购买所需石榴籽数
	gift.ToNumber, _ = strconv.ParseInt(c.PostForm("ToNumber"), 10, 64)   //	赠送时可兑换石榴籽数
	gift.State, _ = strconv.ParseInt(c.PostForm("State"), 10, 64)         //	状态
	gift.GiftPicture = filepath                                           //	礼物图片
	gift.GiftPicStatic = filepath_p                                       //礼物图片 静态
	if flag, _ := gift.GiftUp(); !flag {
		comm.ResponseError(c, 3135) //礼物修改失败
		return
	}
	comm.ResponseError(c, 3136) //礼物修改成功
}

//获取礼物信息
func GetGift(c *gin.Context) {
	var gift models.Gift
	gift.Id, _ = strconv.ParseInt(c.PostForm("Id"), 10, 64)
	if !gift.GetGift() {
		c.JSON(200, gin.H{"gift": "fail"})
		return
	}
	c.JSON(200, gin.H{"gift": gift})
}

//获取礼物列表
func GetGiftList(c *gin.Context) {
	page, _ := strconv.Atoi(c.PostForm("page"))
	rows, _ := strconv.Atoi(c.PostForm("rows"))
	sort := c.PostForm("sort")
	order := c.PostForm("order")
	inputid := c.PostForm("inputid")
	list, row := GetGiftList_com(page, rows, sort, order, inputid)
	c.JSON(200, gin.H{"total": row, "rows": list})
}
func GetGiftList_com(page, rows int, sort, order, inputid string) ([]models.Gift, int64) {
	var gift models.Gift
	page = page
	rows = rows
	sort = sort
	order = order
	inputid = inputid
	list, row := gift.GetGiftList(page, rows, sort, order, inputid)
	if len(list) == 0 {
		return list, row
	}
	return list, row
}
