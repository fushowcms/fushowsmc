package control

/*
 * Copyright (c) 2016—2017 www.fushow.cn, All rights reserved.
 */
import (
	"fushowcms/comm"
	"github.com/gin-gonic/gin"
	m "fushowcms/models"
	"strconv"
)

/*
 * 功能:获取网站后台侧边栏 总管不返回联盟单位，联盟管理只返回联盟单位
 * 请求参数1个: UID
 * 返回值2个:
 * @徐林 20161115
 * @牟海龙 20161130修改
 */
func GetSidebar(c *gin.Context) {
	var (
		userinfo m.UidInfo
		al       m.AuthorityList //网站后台侧边栏
	)
	userinfo.Id, _ = strconv.ParseInt(c.PostForm("UID"), 10, 64)
	if c.PostForm("UID") != "" {
		if !userinfo.GetUserInfo() {
			comm.ResponseError(c, 4010)
			return
		}
	} else {
		comm.ResponseError(c, 4017)
		return
	}
	list := al.GetSidebar(userinfo.Type)
	comm.Response(c, gin.H{"state": "true", "list": list})
}
