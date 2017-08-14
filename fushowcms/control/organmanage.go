package control

/*
 * Copyright (c) 2016—2017 www.fushow.cn, All rights reserved.
 */
import (
	"fushowcms/comm"
	m "fushowcms/models"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

//增加机构 addby liuhan
func AddOrgan(c *gin.Context) {
	var o m.Organ
	organCode := c.PostForm("OrganCode")
	organName := c.PostForm("OrganName")
	passWord := c.PostForm("NewPassWord")
	passWord = comm.SetAesValue(passWord, "fushow.cms")
	if organCode == "" {
		comm.ResponseError(c, 3153) // 联盟编码为空
		return
	}
	if organName == "" {
		comm.ResponseError(c, 3154) // 联盟名字为空
		return
	}
	o.OrganCode = organCode
	o.OrganName = organName
	o.PassWord = passWord
	//判断编码是否存在
	data1, _ := o.FindOrganByOrganCode()
	if len(data1) != 0 {
		comm.ResponseError(c, 3155) // 联盟编码已存在
		return
	}
	//判断名字是否存在
	data2, _ := o.FindOrganByOrganName()
	if len(data2) != 0 {
		comm.ResponseError(c, 3156) // 联盟名字已存在
		return
	}
	if !o.AddOrgan() {
		comm.ResponseError(c, 3157) // 联盟增加失败
		return
	} else {
		var (
			uk m.UserKey //判断用户是否存在
			ui m.UidInfo
		)
		uk.UserName = organCode
		ui.UserName = organCode
		ui.PassWord = passWord
		ui.NickName = organName
		ui.Type = 4          //4:机构管理员
		if uk.GetUserKey() { //用户已存在
			comm.ResponseError(c, 3159) //联盟编号对应的用户账号已存在，请修改
			return
		}
		if !ui.UserInfoAdd() { //增加用户失败
			comm.ResponseError(c, 3157) // 联盟增加失败
			return
		} else {
			//			c.JSON(200, gin.H{"state": true})
			comm.ResponseError(c, 3158) // 联盟增加成功
		}
	}
}

//查询机构
func FindOrgan(c *gin.Context) {
	var o m.Organ
	page, _ := strconv.Atoi(c.PostForm("page"))
	rows, _ := strconv.Atoi(c.PostForm("rows"))
	sort := c.PostForm("sort")
	order := c.PostForm("order")
	inputid := c.PostForm("inputid")
	list, row := o.FindOrgan(page, rows, sort, order, inputid)
	for i := 0; i < len(list); i++ {
		list[i].PassWord = comm.GetDecValue(list[i].PassWord, "fushow.cms")
	}
	if len(list) == 0 {
		c.JSON(200, gin.H{"total": 0, "rows": []int{}})
		return
	}
	c.JSON(200, gin.H{"total": row, "rows": list})
}

//根据机构id查询
func FindOrganById(c *gin.Context) {
	var o m.Organ
	o.OrganCode = c.PostForm("code")
	data, _ := o.FindOrganById()
	if len(data) <= 0 {
		comm.ResponseError(c, 3176) //查询联盟失败
		return
	}
	comm.Response(c, data)
}

//根据机构编码查询
func FindOrganByOrganCode(c *gin.Context) {
	var o m.Organ
	organCode := c.PostForm("OrganCode")
	if organCode == "" {
		comm.ResponseError(c, 3160) //联盟编码为空
		return
	}
	o.OrganCode = organCode
	data, _ := o.FindOrganByOrganCode()
	if len(data) == 0 {
		comm.ResponseError(c, 3161) //联盟不存在
		return
	}
	comm.Response(c, data)
}

//根据机构名称查询
func FindOrganByOrganName(c *gin.Context) {
	var o m.Organ
	organName := c.PostForm("OrganName")
	if organName == "" {
		c.JSON(200, gin.H{"state": false, "msg": "3008"}) //3008：机构编码为空
		return
	}
	o.OrganName = organName
	data, err := o.FindOrganByOrganName()
	if err != nil && data == nil {
		c.JSON(200, gin.H{"state": false, "data": nil})
		return
	}
	c.JSON(200, gin.H{"state": true, "data": data})
}

//根据用户查询
func FindOrganByUserId(c *gin.Context) {
	var o m.OrganManage
	userId, _ := strconv.ParseInt(c.PostForm("UID"), 10, 64)
	if userId == 0 {
		comm.ResponseError(c, 3175)
		return
	}
	o.UserId = userId
	data, err := o.FindOrganByUserId()
	if err != nil && data == nil {
		comm.ResponseError(c, 3161) //联盟不存在
		return
	}
	comm.Response(c, data)
}

//删除机构 addby liuhan
func DelOrgan(c *gin.Context) {
	var (
		o  m.Organ
		om m.OrganManage
	)
	o.Id, _ = strconv.ParseInt(c.PostForm("Id"), 10, 64)       //ID
	om.OrganId, _ = strconv.ParseInt(c.PostForm("Id"), 10, 64) //ID
	if o.Id == 0 {
		comm.ResponseError(c, 3162) //用户id为空
		return
	}
	if om.OrganId == 0 {
		comm.ResponseError(c, 3163) //联盟id为空
		return
	}
	//判断id是否存在
	data, _ := o.FindOrganById()
	if len(data) == 0 {
		comm.ResponseError(c, 3161) //联盟不存在
		return
	}
	//判断该机构下是否已绑定用户
	data1, _ := om.FindOrganManageByOrganId()
	if len(data1) > 0 {
		comm.ResponseError(c, 3164) //该联盟中有用户存在,不能删除
		return
	}
	//删除
	if !o.DelOrgan() {
		comm.ResponseError(c, 3165) //删除联盟失败
		return
	}
	comm.ResponseError(c, 3166) //删除联盟成功
}

//修改机构 addby liuhan
//liuhan 20161130
func UpdateOrgan(c *gin.Context) {
	var (
		o  m.Organ
		o1 m.Organ
	)
	id, _ := strconv.ParseInt(c.PostForm("Id"), 10, 64) //ID
	organCode := c.PostForm("OrganCode")
	organName := c.PostForm("OrganName")
	passWord := c.PostForm("PassWord")
	newpassWord := c.PostForm("NewPassWord")
	newpassWord = comm.SetAesValue(newpassWord, "fushow.cms")
	o.Id = id
	o1.Id = id
	o.OrganCode = organCode
	o.OrganName = organName
	o.PassWord = newpassWord
	//判断id是否存在
	data3, _ := o.FindOrganById()
	if len(data3) == 0 {
		comm.ResponseError(c, 3161) //联盟不存在
		return
	}
	if !o1.GetOrgan() {
		comm.ResponseError(c, 3161) //联盟不存在
		return
	}
	o1.OrganCode = organCode
	o1.OrganName = organName
	o1.PassWord = newpassWord
	if !o1.UpdateOrgan() {
		comm.ResponseError(c, 3167) //修改联盟失败
		return
	}
	// 如果机构修改密码，那么用户表也要对应把密码改了。
	if !strings.EqualFold(passWord, newpassWord) {
		var ui m.UidInfo
		ui.UserName = organCode
		has := ui.GetUserInfo()
		if !has {
			comm.ResponseError(c, 3167) //修改联盟失败
			return
		}
		ui.PassWord = newpassWord
		if !ui.PassUp() {
			comm.ResponseError(c, 3167) //修改联盟失败
			return
		}
	}
	comm.ResponseError(c, 3168) //修改联盟
}

//新增机构管理信息 addby liuhan
//liuhan 20161130
func AddOrganManage(c *gin.Context) {
	var (
		om  m.OrganManage
		uid m.UidInfo
	)
	userId, _ := strconv.ParseInt(c.PostForm("UID"), 10, 64)              //UserId
	organId, _ := strconv.ParseInt(c.PostForm("OrganId"), 10, 64)         //OrganId
	rechargeNum, _ := strconv.ParseInt(c.PostForm("RechargeNum"), 10, 64) //rechargeNum
	rechargeMethod := c.PostForm("RechargeMethod")                        //RechargeMethod
	if userId == 0 {
		comm.ResponseError(c, 3162) //用户id为空
		return
	}
	if organId == 0 {
		comm.ResponseError(c, 3163) //联盟id为空
		return
	}
	om.OrganId = organId
	om.UserId = userId
	om.RechargeNum = rechargeNum
	om.RechargeMethod = rechargeMethod
	if !om.AddOrganManage() {
		comm.ResponseError(c, 3157) //联盟增加失败
		return
	}
	uid.Id = userId
	if !uid.GetUserInfo() {
		comm.ResponseError(c, 3175) //用户不存在
		return
	}
	_, _, number := m.BindGiveNumber(uid.Id, 2)
	m := make(map[string]interface{})
	m["state"] = true
	m["Number"] = number
	comm.Response(c, m)
}

//根据机构id查询机构管理表 //查询机构下所有用户 addby liuhan
func FindOrganManageByOrganId(c *gin.Context) {
	var om m.OrganManage
	page, _ := strconv.Atoi(c.PostForm("page"))
	rows, _ := strconv.Atoi(c.PostForm("rows"))
	organId, _ := strconv.ParseInt(c.PostForm("OrganId"), 10, 64) //OrganId
	sort := c.PostForm("sort")
	order := c.PostForm("order")
	inputid := c.PostForm("inputid")
	om.OrganId = organId
	list, row := om.FindOrganManageByOrganIds(page, rows, sort, order, inputid)
	if len(list) == 0 {
		c.JSON(200, gin.H{"total": row, "rows": []int{}})
		return
	}
	c.JSON(200, gin.H{"total": row, "rows": list})
}

//查询机构下充值的总额人数 addby liuhan
func SumOrganManageByOrganId(c *gin.Context) {
	var om m.OrganManage
	organId, _ := strconv.ParseInt(c.PostForm("OrganId"), 10, 64) //OrganId
	if organId == 0 {
		comm.ResponseError(c, 3163) //联盟id为空
		return
	}
	om.OrganId = organId
	total, err, count := om.SumOrganManageByOrganId()
	if err != nil {
		comm.ResponseError(c, 3169) //失败
		return
	}
	m := make(map[string]interface{})
	m["state"] = true
	m["total"] = total
	m["count"] = count
	comm.Response(c, m)
}
