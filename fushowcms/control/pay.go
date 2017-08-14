package control

/*
 * Copyright (c) 2016—2017 www.fushow.cn, All rights reserved.
 */
import (
	"fushowcms/models"

	"strconv"

	"github.com/gin-gonic/gin"
)

func GetPayList(c *gin.Context) {
	var rec models.RechargingRecords
	page, _ := strconv.Atoi(c.PostForm("page"))
	rows, _ := strconv.Atoi(c.PostForm("rows"))
	uid := c.PostForm("inputid")
	if uid == "" {
		list, total := rec.GetPayList(rows, page)
		if len(list) == 0 {
			c.JSON(200, gin.H{"total": "", "order": "fail", "rows": []int{}})
			return
		}

		c.JSON(200, gin.H{"total": total, "rows": list})
		return
	}

	slist, stotal := rec.GetPaySearchList(rows, page, uid)
	if stotal == 0 {
		c.JSON(200, gin.H{"total": "", "order": "fail", "rows": []int{}})
		return
	}
	c.JSON(200, gin.H{"total": stotal, "rows": slist})
}
