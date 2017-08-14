package control

/*
 * Copyright (c) 2016—2017 www.fushow.cn, All rights reserved.
 */
import (
	"github.com/gin-gonic/gin"
	"fushowcms/models"
	"strconv"
)

//获取资金池全部信息
func GetAllFundDesc(c *gin.Context) {
	var fund models.Fund
	page, _ := strconv.Atoi(c.PostForm("page"))
	rows, _ := strconv.Atoi(c.PostForm("rows"))
	list, row := fund.GetAllFundDesc(page, rows)
	if row == 0 {
		return
	}
	c.JSON(200, gin.H{"total": row, "rows": list})
}
