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

var Flag int64 = 0

/* 添加轮播图 后台管理
 * 请求参数:uploadFile、PicName、state、CarouselType
 * @author 徐林->修正
 */
func CarPicAdd(c *gin.Context) {

	var (
		cp m.CarouselPic
		ar m.AnchorRoom
	)

	cp.CarouselType, _ = strconv.ParseInt(c.PostForm("CarouselType"), 10, 64) //轮播类型(视频：1，图片：3 2 聚合)
	cp.PicName = c.PostForm("PicName")                                        //轮播名称
	cp.State, _ = strconv.ParseInt(c.PostForm("State"), 10, 64)               //显示状态
	cp.StartTime = time.Now().Format("2006-01-02 15:04:05")                   //设置开始时间为当前时间
	cp.EndTime = time.Now().Format("2006-01-02 15:04:05")                     //设置结束时间为当前时间
	ar.Id, _ = strconv.ParseInt(c.PostForm("LiveRoom"), 10, 64)               //房间号
	cp.VideoLivePage = c.PostForm("VideoLivePage")
	if cp.CarouselType == 2 {
		cp.VideoLivePage = c.PostForm("juheLivePage")
	}

	//判断是否为视频
	if cp.CarouselType == 1 {
		if ar.Id == 0 || !ar.GetRoom() {
			comm.ResponseError(c, 3101) //房间不存在
			return
		}
	}
	if cp.PicName == "" {
		comm.ResponseError(c, 3102) //请输入轮播名称
		return
	}

	_, filepath := CUploads(c, "uploadFile") //获取上传的缩略图
	if filepath == "" {
		comm.ResponseError(c, 3104) //请上传轮播图片
		return
	}
	//图片时，需要缩略图
	if cp.CarouselType == 3 {
		_, cfilepath := CUploads(c, "limit")
		if cfilepath == "" {
			comm.ResponseError(c, 3105) //请上传缩略图图片
			return
		}
		cp.Litming = cfilepath //缩略图
	}

	cp.PicPath = filepath //正图

	if cp.CarouselType == 1 { //判断如果是视频
		cp.VideoLivePage = "/page/roomlive?roomId=" + strconv.FormatInt(ar.Id, 10) + "&" + "anchorId=" + strconv.FormatInt(ar.Uid, 10)
	}

	if !cp.CarPicAdd() {
		comm.ResponseError(c, 3106) //上传轮播图失败
		return
	}
	comm.ResponseError(c, 3107) //添加成功
}

//删除轮播图
func CarPicDel(c *gin.Context) {
	var cp m.CarouselPic
	cp.Id, _ = strconv.ParseInt(c.PostForm("id"), 10, 64)
	if !cp.CarPicDel() {
		comm.ResponseError(c, 3108) //删除失败
		return
	}
	comm.ResponseError(c, 3109) //删除成功
}

//修改轮播图
func CarPicUp(c *gin.Context) {
	var (
		cp m.CarouselPic
		ar m.AnchorRoom
	)
	cp.Id, _ = strconv.ParseInt(c.PostForm("Idss"), 10, 64) //根据ID修改礼物信息
	if !cp.GetCarPic() {
		comm.ResponseError(c, 3110) //获取轮播图信息失败
		return
	}
	cp.CarouselType, _ = strconv.ParseInt(c.PostForm("CarouselType"), 10, 64) //轮播类型
	cp.PicName = c.PostForm("PicName")                                        //轮播名称
	cp.VideoLivePage = c.PostForm("VideoLivePage")                            //轮播名称
	cp.State, _ = strconv.ParseInt(c.PostForm("State"), 10, 64)               //显示状态
	if cp.CarouselType == 3 {
		_, filepath := Uploads(c)
		_, limit := CUploads(c, "limit") //获取上传的缩略图
		cp.PicPath = filepath
		cp.Litming = limit
	} else if cp.CarouselType == 2 {
		cp.VideoLivePage = c.PostForm("juheLivePage")

	} else if cp.CarouselType == 1 {
		ar.Id, _ = strconv.ParseInt(c.PostForm("LiveRoom"), 10, 64) //房间号

		if !ar.GetRoom() {
			comm.ResponseError(c, 3111) //获取房间失败
			return
		}
		liveadd := "/page/roomlive?roomId=" + strconv.FormatInt(ar.Id, 10) + "&anchorId=" + strconv.FormatInt(ar.Uid, 10)

		cp.VideoLivePage = liveadd
	}
	if !cp.CarPicUp() {
		comm.ResponseError(c, 3112) //修改轮播图失败
		return
	}
	comm.ResponseError(c, 3113) //修改轮播图成功

}

//获取轮播图列表
func GetCarPicList(c *gin.Context) {
	var cp m.CarouselPic

	page, _ := strconv.Atoi(c.PostForm("page"))
	rows, _ := strconv.Atoi(c.PostForm("rows"))

	list, row := cp.GetCarPicList(page, rows)

	if row == 0 {
		comm.ResponseError(c, 3182) //轮播图不存在
		return
	}
	comm.Response(c, list)

}

/* 获取轮播图列表 后台管理
 * 请求参数:page、rows
 * @author 徐林->修正
 * 返回值：total(行数)、rows(列表) easyui设定必须返回total、rows
 */
func GetIndexCarPicList(c *gin.Context) {
	var cp m.CarouselPic
	page, _ := strconv.Atoi(c.PostForm("page"))
	rows, _ := strconv.Atoi(c.PostForm("rows"))

	list, row := cp.GetIndexCarPicList(page, rows)

	if row == 0 {
		c.JSON(200, gin.H{"total": 0, "rows": ""})
		return
	}
	c.JSON(200, gin.H{"total": row, "rows": list})
}
func GetIndexCarPicLists(c *gin.Context) {
	var cp m.CarouselPic
	page, _ := strconv.Atoi(c.PostForm("page"))
	rows, _ := strconv.Atoi(c.PostForm("rows"))
	list, row := cp.GetIndexCarPicLists(page, rows)
	if row == 0 {
		c.JSON(200, gin.H{"total": 0, "rows": ""})
		return
	}
	c.JSON(200, gin.H{"total": row, "rows": list})
}

//手机端获取轮播图列表
func GetCarPicLists(c *gin.Context) {
	var cp m.CarouselPic
	page, _ := strconv.Atoi(c.PostForm("page"))
	rows, _ := strconv.Atoi(c.PostForm("rows"))
	list, row := cp.GetCarPicList(page, rows)
	if row == 0 {
		return
	}
	c.JSON(200, gin.H{"total": row, "rows": list})
}
