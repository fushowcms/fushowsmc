package control

/*
 * Copyright (c) 2016—2017 www.fushow.cn, All rights reserved.
 */
import (
	"fushowcms/comm"
	"fushowcms/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

//txl-del
func GiftGiveAdd(c *gin.Context) {
	var gift models.GiftGive
	gift.RecipientId, _ = strconv.ParseInt(c.PostForm("RecipientId"), 10, 64)
	gift.BenefactorId, _ = strconv.ParseInt(c.PostForm("BenefactorId"), 10, 64)
	gift.GiftId, _ = strconv.ParseInt(c.PostForm("GiftId"), 10, 64)
	gift.GiftNum, _ = strconv.ParseInt(c.PostForm("GiftNum"), 10, 64)
	if !gift.GiftGiveAdd() {

		c.JSON(200, gin.H{"state": "fail"})
		return
	}
	c.JSON(200, gin.H{"state": "success"})
}

//txl-del
func GetGiftGiveList(c *gin.Context) {
	var gift models.GiftGive
	page, _ := strconv.Atoi(c.PostForm("page"))
	rows, _ := strconv.Atoi(c.PostForm("rows"))
	list, row := gift.GetGiftGiveList(page, rows)
	if row == 0 {
		return
	}
	c.JSON(200, gin.H{"total": row, "rows": list})
}

//刷礼物  周榜
//参数
func GetGiftGiveWeeks(c *gin.Context) {
	anchorid, _ := strconv.ParseInt(c.PostForm("AnchorId"), 10, 64)
	_, list := GetGiftGiveWeeks_com(anchorid)
	comm.Response(c, list)
}

func GetGiftGiveWeeks_com(AnchorId int64) (bool, []models.WeeksGift) {
	var gw models.GiftGive
	anchorid := AnchorId
	if anchorid == 0 {
	}
	flag, list := gw.GetGiftGiveWeeks(anchorid)
	if !flag {
	}
	return flag, list
}

//刷礼物  总榜
//参数
func GetGiftGiveMonths(c *gin.Context) {
	anchorid, _ := strconv.ParseInt(c.PostForm("AnchorId"), 10, 64)
	_, list := GetGiftGiveMonths_com(anchorid)
	comm.Response(c, list)
}

func GetGiftGiveMonths_com(AnchorId int64) (bool, []models.WeeksGift) {
	var gw models.GiftGive
	anchorid := AnchorId
	if anchorid == 0 {
	}
	flag, list := gw.GetGiftGiveMonth(anchorid)
	if !flag {
	}
	return flag, list
}

//根据userId查询刷礼物记录 addby liuhan
func FindByUserIdAll(c *gin.Context) {
	var gg models.GiftGive
	benefactorId, _ := strconv.ParseInt(c.PostForm("UID"), 10, 64)
	page, _ := strconv.Atoi(c.PostForm("page"))
	rows, _ := strconv.Atoi(c.PostForm("rows"))
	if benefactorId == 0 {
		comm.ResponseError(c, 2000) //参数错误
		return
	}
	gg.BenefactorId = benefactorId
	data, err, total := gg.FindByUserIdAll(page, rows)
	if err != nil && data == nil {
		comm.ResponseError(c, 2012) //没有更多数据
		return
	}
	m := make(map[string]interface{})
	m["data"] = data
	m["total"] = total
	m["id"] = benefactorId
	comm.Response(c, m)
}

//根据userId查询赠送石榴籽记录 addby liuhan
func FindByUserIdGiveSlz(c *gin.Context) {
	var (
		gs models.GiveSlz
	)
	userId, _ := strconv.ParseInt(c.PostForm("UID"), 10, 64)
	//获取页数-行数
	page, _ := strconv.Atoi(c.PostForm("page"))
	rows, _ := strconv.Atoi(c.PostForm("rows"))
	if userId == 0 {
		comm.ResponseError(c, 2000) //参数错误
		return
	}
	data, _, total := gs.FindByUserIdGiveSlz(page, rows, userId)
	if len(data) <= 0 {
		comm.ResponseError(c, 2012) //没有更多数据
		return
	}
	m := make(map[string]interface{})
	if total == 0 {
		m["data"] = []int{}
	} else {
		m["data"] = data
	}
	m["total"] = total
	comm.Response(c, m)
}

//后台赠送石榴籽记录
//time  2016-11-11 txl
func GetNumberList(c *gin.Context) {
	var gs models.GiveSlz
	page, _ := strconv.Atoi(c.PostForm("page"))
	rows, _ := strconv.Atoi(c.PostForm("rows"))
	nickname := c.PostForm("NickName")
	if nickname == "" {
		lista, totala := gs.GetNumberList(page, rows)
		if len(lista) == 0 {
			c.JSON(200, gin.H{"total": totala, "rows": []int{}})
			return
		}
		c.JSON(200, gin.H{"total": totala, "rows": lista})
		return
	}
	listb, totalb := gs.GetNumberLikeList(page, rows, nickname)
	if len(listb) == 0 {
		c.JSON(200, gin.H{"total": listb, "rows": []int{}})
		return
	}
	c.JSON(200, gin.H{"total": totalb, "rows": listb})
}

//后台查询刷礼物记录 addby liuhan
func FindGiveAll(c *gin.Context) {
	var gg models.GiftGive
	page, _ := strconv.Atoi(c.PostForm("page"))
	rows, _ := strconv.Atoi(c.PostForm("rows"))
	inputid := c.PostForm("inputid")
	data, _, total := gg.FindGiveAll(page, rows, inputid)
	if len(data) == 0 {
		c.JSON(200, gin.H{"state": false, "rows": []int{}})
		return
	}
	c.JSON(200, gin.H{"total": total, "rows": data})
}

//根据Recipient查询刷礼物记录 addby liuhan
func FindByRecipient(c *gin.Context) {
	var gg models.GiftGive
	recipientId, _ := strconv.ParseInt(c.PostForm("UID"), 10, 64)
	page, _ := strconv.Atoi(c.PostForm("page"))
	rows, _ := strconv.Atoi(c.PostForm("rows"))
	if recipientId == 0 {
		comm.ResponseError(c, 2000) //参数错误
		return
	}
	gg.RecipientId = recipientId
	data, err, total := gg.FindByRecipient(page, rows)
	if err != nil && data == nil {
		comm.ResponseError(c, 2012) //没有更多数据
		return
	}
	m := make(map[string]interface{})
	m["data"] = data
	m["total"] = total
	m["id"] = recipientId
	comm.Response(c, m)
}

//根据benefactorId查询刷礼物记录 addby liuhan
func FindByBenefactor(c *gin.Context) {
	var gg models.GiftGive
	benefactorId, _ := strconv.ParseInt(c.PostForm("UID"), 10, 64)
	page, _ := strconv.Atoi(c.PostForm("page"))
	rows, _ := strconv.Atoi(c.PostForm("rows"))
	if benefactorId == 0 {
		c.JSON(200, gin.H{"state": false, "msg": "3006"}) //3006：用户id为空
		return
	}
	gg.BenefactorId = benefactorId
	data, err, total := gg.FindByBenefactor(page, rows)
	if err != nil && data == nil {
		c.JSON(200, gin.H{"state": false, "data": nil})
		return
	}
	c.JSON(200, gin.H{"state": true, "data": data, "total": total, "id": benefactorId})
}
