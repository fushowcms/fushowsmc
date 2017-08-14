package control

/*
 * Copyright (c) 2016—2017 www.fushow.cn, All rights reserved.
 */
import (
	"fushowcms/comm"
	m "fushowcms/models"
	"io"
	"log"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

//txl-del
//添加申请
func ApplyAdd(c *gin.Context) {
	var (
		al m.Applicant
		ui m.UidInfo
	)
	al.UserId, _ = strconv.ParseInt(c.PostForm("UID"), 10, 64) //添加用户Id
	if al.UserId == 0 {
		comm.ResponseError(c, 3175)
		return
	}
	al.State = 0 //添加审核状态	0：未审核，1：审核通过，2：审核未通过
	ui.Id = al.UserId
	realname := c.PostForm("RealName") //真实姓名
	idnumber := c.PostForm("IdNumber") //身份证号
	phone := c.PostForm("Phone")
	if !ui.GetUserInfo() {
		comm.ResponseError(c, 2004)
		return
	}
	ui.RealName = realname
	ui.IdNumber = idnumber
	ui.Phone = phone
	file, header, err := c.Request.FormFile("upload")
	if header.Filename == "" {
		comm.ResponseError(c, 2005)
		return
	}
	out, err := os.Create("./static/upload/" + "photo" + c.PostForm("UID") + ".png")
	if err != nil {
		log.Fatal(err)
	}
	defer out.Close()
	_, err = io.Copy(out, file)

	if err != nil {
		log.Fatal(err)
	}
	ui.IdentityPic = "/static/upload/" + "photo" + c.PostForm("UID") + ".png"
	//修改用户信息
	if !ui.UserInfoUp() {
		comm.ResponseError(c, 2006)
		return
	}
	if !al.ApplyAdd() {
		comm.ResponseError(c, 2007)
		return
	}
	comm.Response(c, "success")
}

//获取某一条申请详情
func GetApplyInfo(c *gin.Context) {
	var (
		ui  m.UidInfo
		al  m.Applicant
		now m.ApplicantInfo
	)
	ui.Id, _ = strconv.ParseInt(c.PostForm("id"), 10, 64) //删除用户Id
	flag := ui.GetUserInfo()
	if !flag {
		comm.ResponseError(c, 2004)
		return
	}
	al.Id, _ = strconv.ParseInt(c.PostForm("applyId"), 10, 64)
	isExist := al.GetApplyInfo()
	if !isExist {
		comm.ResponseError(c, 2007)
		return
	}
	now.Id = al.Id
	now.Uid = ui.Id
	now.NickName = ui.UserName
	now.Phone = ui.Phone
	now.Type = ui.Type
	now.Level = ui.Level
	now.RealName = ui.RealName
	now.IdNumber = ui.IdNumber
	now.IdentityPic = ui.IdentityPic
	now.ApplicantTime = al.ApplicantTime
	now.ApplyId = al.Id
	now.State = al.State
	comm.Response(c, now)
}

//同意申请
func ApplyArg(c *gin.Context) {
	var (
		ui      m.UidInfo
		cheeckr m.AnchorRoom
		al      m.Applicant
	)
	ui.Id, _ = strconv.ParseInt(c.PostForm("Id"), 10, 64) //删除用户Id
	cheeckr.Uid = ui.Id

	if !ui.GetUserInfo() {
		comm.ResponseError(c, 2004)
		return
	}
	if cheeckr.GetRoom() {
		comm.ResponseError(c, 2008)
		return
	}
	ui.Type = 1 //主播

	if !ui.UserInfoUp() {
		comm.ResponseError(c, 2006)
		return
	}
	al.Id, _ = strconv.ParseInt(c.PostForm("ApplyId"), 10, 64)
	if !al.GetApplyInfo() {
		comm.ResponseError(c, 2007)
		return
	}
	al.State = 1
	if !al.ApplyStateUp() {
		comm.ResponseError(c, 2009)
		return
	}
	var ar m.AnchorRoom
	ar.Uid = ui.Id                               //主播Id
	ar.RoomType = 9999999999                     //房间分类 0：普通房间，1：竞猜房间
	ar.RoomAlias = "欢迎来到" + ui.NickName + "的直播间" //房间别名
	ar.LiveState = 0                             //直播状态	0：未直播，1：直播中
	if !ar.RoomAdd() {
		comm.ResponseError(c, 2010)
		return
	}
	var aln m.AnchorLiveNumber
	aln.RoomId = ar.Id
	aln.LiveNumber = 0
	if !aln.LiveNumberAdd() {
		comm.ResponseError(c, 2010)
		return
	}
	comm.Response(c, "您同意该申请")
}

//拒绝申请
func ApplyRefused(c *gin.Context) {
	var (
		ui m.UidInfo
		al m.Applicant
	)
	ui.Id, _ = strconv.ParseInt(c.PostForm("Id"), 10, 64)

	if !ui.GetUserInfo() {
		comm.ResponseError(c, 2004)
		return
	}
	ui.Type = 0
	if !ui.UserInfoUp() {
		comm.ResponseError(c, 2006)
		return
	}
	al.Id, _ = strconv.ParseInt(c.PostForm("ApplyId"), 10, 64)
	if !al.GetApplyId() {
		comm.ResponseError(c, 2007)
		return
	}
	al.State = 2
	if !al.ApplyStateUp() {
		comm.ResponseError(c, 2009)
		return
	}
	comm.Response(c, "您拒绝了该申请")
}

//获取申请列表筛选
func GetApplyLists(c *gin.Context) {
	var al m.Applicant
	page, _ := strconv.Atoi(c.PostForm("page"))
	rows, _ := strconv.Atoi(c.PostForm("rows"))
	state := c.PostForm("state")
	list, row := al.GetApplyLists(page, rows, state)
	var arr [1]int // 声明了一个int类型的数组
	arr[0] = 99999 // 数组下标是从0开始的
	if row == 0 {
		c.JSON(200, gin.H{"total": row, "rows": arr})
		return
	}
	c.JSON(200, gin.H{"total": row, "rows": list})
}

//获取申请列表
func GetApplyList(c *gin.Context) {
	var al m.Applicant
	page, _ := strconv.Atoi(c.PostForm("page"))
	rows, _ := strconv.Atoi(c.PostForm("rows"))

	list, row := al.GetApplyList(page, rows)
	var arr [1]int // 声明了一个int类型的数组
	arr[0] = 99999 // 数组下标是从0开始的
	if row == 0 {
		c.JSON(200, gin.H{"total": row, "rows": arr})
		return
	}
	c.JSON(200, gin.H{"total": row, "rows": list})
}

//判断是否已经提交
func IsApplyExit(c *gin.Context) {
	var al m.Applicant
	al.UserId, _ = strconv.ParseInt(c.PostForm("UID"), 10, 64)
	//参数错误
	if c.PostForm("UID") == "" {
		comm.ResponseError(c, 2000)
		return
	}
	//不存在
	if !al.ApplicantExit() {
		comm.ResponseError(c, 2001)
		return
	}
	//同意
	if al.State == 1 {
		comm.ResponseError(c, 2002)
		return
	} else if al.State == 2 { //拒绝
		comm.ResponseError(c, 2003)
		return
	}
	comm.Response(c, "正在审核中")
}
