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

//添加产品
func ProductAdd(c *gin.Context) {
	var pro models.Product
	pro.ProductName = c.PostForm("ProductName")
	pro.State1 = c.PostForm("State1")
	pro.State2 = c.PostForm("State2")
	pro.State3 = c.PostForm("State3")
	pro.State4 = c.PostForm("State4")
	pro.State5 = c.PostForm("State5")
	pro.State6 = c.PostForm("State6")
	pro.State7 = c.PostForm("State7")
	pro.State8 = c.PostForm("State8")
	pro.State9 = c.PostForm("State9")
	pro.State10 = c.PostForm("State10")
	if !pro.ProductAdd() {
		comm.ResponseError(c, 2018)
		return
	}
	comm.Response(c, "成功")
}

//删除产品
func ProductDel(c *gin.Context) {
	var pro models.Product
	pro.Id, _ = strconv.ParseInt(c.PostForm("id"), 10, 64)
	if pro.Id == 0 {
		return
	}
	if !pro.ProductDel() {
		comm.ResponseError(c, 2018)
		return
	}
	comm.Response(c, "成功")
}

//修改产品信息
func ProductUp(c *gin.Context) {
	var pro models.Product
	pro.Id, _ = strconv.ParseInt(c.PostForm("Id"), 10, 64)

	productName := c.PostForm("ProductName")
	state1 := c.PostForm("State1")
	state2 := c.PostForm("State2")
	state3 := c.PostForm("State3")
	state4 := c.PostForm("State4")
	state5 := c.PostForm("State5")
	state6 := c.PostForm("State6")
	state7 := c.PostForm("State7")
	state8 := c.PostForm("State8")
	state9 := c.PostForm("State9")
	state10 := c.PostForm("State10")

	if !pro.GetProducts() {
		comm.ResponseError(c, 2018)
		return
	}

	pro.ProductName = productName
	pro.State1 = state1
	pro.State2 = state2
	pro.State3 = state3
	pro.State4 = state4
	pro.State5 = state5
	pro.State6 = state6
	pro.State7 = state7
	pro.State8 = state8
	pro.State9 = state9
	pro.State10 = state10

	if !pro.ProductUp() {
		comm.ResponseError(c, 2018)
		return
	}
	comm.Response(c, "成功")
}

//查询产品信息
func GetProduct(c *gin.Context) {
	var pro models.Product
	pro.Id, _ = strconv.ParseInt(c.PostForm("Id"), 10, 64)
	fmt.Println(pro.Id)
	product, err := pro.GetProduct(pro.Id)
	if err == nil {
		comm.Response(c, product)
		return
	}
	comm.ResponseError(c, 2019)
}

//查询产品信息
//后台
func GetProductList(c *gin.Context) {
	var pro models.Product
	page, _ := strconv.Atoi(c.PostForm("page"))
	rows, _ := strconv.Atoi(c.PostForm("rows"))
	inputid := c.PostForm("inputid")
	list, row := pro.GetProductLists(page, rows, inputid)
	if row == 0 {
		return
	}
	c.JSON(200, gin.H{"total": row, "rows": list})
}

//获取所有产品（后台checkbox）
func GetMyProductList(c *gin.Context) {
	var pro models.Product
	page, _ := strconv.Atoi(c.PostForm("page"))
	rows, _ := strconv.Atoi(c.PostForm("rows"))

	list, row := pro.GetProductList(page, rows)
	if row == 0 {
		comm.ResponseError(c, 2019)
		return
	}

	comm.Response(c, list)

}
