package control

/*
 * Copyright (c) 2016—2017 www.fushow.cn, All rights reserved.
 */
import (
	"fmt"
	"fushowcms/comm"
	m "fushowcms/models"
	"sort"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/BPing/aliyun-live-go-sdk/client"
	"github.com/BPing/aliyun-live-go-sdk/device/live"
)

//查询观看直播记录
func WatchRecord(c *gin.Context) {
	var (
		wr m.WatchRecordList
	)
	//获取用户ID
	wr.Uid, _ = strconv.ParseInt(c.PostForm("UID"), 10, 64)
	if wr.Uid == 0 {
		comm.ResponseError(c, 3175)
		return
	}
	//获取页数-行数
	page, _ := strconv.Atoi(c.PostForm("page"))
	rows, _ := strconv.Atoi(c.PostForm("rows"))
	watchlist, err := wr.GetWatchRecord(page, rows)
	if len(watchlist) == 0 {
		c.JSON(200, gin.H{"state": "null"}) //没有观看记录
		return
	}
	if err != nil {
		c.JSON(200, gin.H{"state": "error"}) //查询失败
		return
	}
	c.JSON(200, gin.H{"message": true, "list": watchlist})
}

//获取直播串流码
//time  2016-11-17  txl
//参数  AnchorId  主播Id
func GetPlugFlow(c *gin.Context) {
	fmt.Println("start", GetDateTime())
	var (
		wr     m.AnchorRoom
		isflag bool = false
		str    string
		plug   string
	)
	//获取用户ID
	wr.Uid, _ = strconv.ParseInt(c.PostForm("UID"), 10, 64)
	regcharge := c.PostForm("Type")
	if wr.Uid == 0 {
		comm.ResponseError(c, 3175)
		return
	}
	//获取直播房间
	if !wr.GetRoomUid() {
		comm.ResponseError(c, 3186) //您不是主播
		return
	}
	//重置  1
	if regcharge != "" {
		// 判断如果房间为禁播房间，不允许重置推流码
		if wr.LiveState == 2 {
			comm.ResponseError(c, 3196) //该房间为禁播房间，禁止重置流密钥
			return
		}
		//重新分配域名  第二次
		allid := strconv.FormatInt(wr.Uid, 10) + "-" + strconv.FormatInt(wr.Id, 10)
		str = GetRealmName()
		plug = m.PlugFlow(allid, str)
		wr.LiveAddress = plug
	} else {
		if wr.LiveAddress == "" {
			isflag = true
		}
		allid := strconv.FormatInt(wr.Uid, 10) + "-" + strconv.FormatInt(wr.Id, 10)
		// 随机分配阿里服务器  第一次执行分配
		if isflag {
			str = GetRealmName()
			plug = m.PlugFlow(allid, str)
			wr.LiveAddress = plug
		}
	}
	plug = wr.LiveAddress
	if plug == "" {
		comm.ResponseError(c, 3187) //返回错误
		return
	}
	//更新DB中LiveAddress的值
	if !wr.RoomUp() {
		comm.ResponseError(c, 3187) //返回错误
		return
	}
	mm := make(map[string]interface{})
	mm["errMsg"] = plug
	fmt.Println("end", GetDateTime())
	comm.Response(c, mm)
}

type RealmNameNum struct {
	Key   string
	Value int
}

// GetRealmName 查所有域名选出最少的地址
func GetRealmName() string {
	var str = ""
	var people []RealmNameNum
	cert := client.NewCredentials(AccessKeyId, AccessKeySecret)
	var ss int64 = 1
	// 把10个域名 查一遍  取到所有线上在线人数
	for i := 0; i < 10; i++ {
		var nows RealmNameNum
		name := "live" + strconv.FormatInt(ss, 10)
		nowlive := comm.GetConfig("FS", name)
		// 取配置文件流域名
		liveM := live.NewLive(cert, nowlive, `sllmzb`, nil).SetDebug(true)
		resp := live.OnlineInfoResponse{}
		// 获取在线人
		liveM.StreamOnlineUserNum("", &resp)
		nows.Key = nowlive
		nows.Value = len(resp.OnlineUserInfo.LiveStreamOnlineUserNumInfo)
		people = append(people, nows)
		ss++
	}
	sort.Sort(realmNameNum(people))
	// 找到一个最少的流
	for j := 9; j >= 0; j-- {
		if people[j].Value != 10 {
			str = people[j].Key
			break
		}
	}
	// 最少的可用的流
	return str
}

type realmNameNum []RealmNameNum

func (a realmNameNum) Len() int { // 重写 Len() 方法
	return len(a)
}
func (a realmNameNum) Swap(i, j int) { // 重写 Swap() 方法
	a[i], a[j] = a[j], a[i]
}
func (a realmNameNum) Less(i, j int) bool { // 重写 Less() 方法， 从大到小排序
	return a[j].Value < a[i].Value
}

//获取直播地址
//time  2016-11-17  txl
//参数  AnchorId  主播Id   Type  1---->手机端  m3u8   默认pc端
//liuhan 20161202
func GetInFlow(c *gin.Context) {
	//	//获取用户ID
	Uid, _ := strconv.ParseInt(c.PostForm("AnchorId"), 10, 64)
	intype := c.PostForm("Type") //1-->//phone   m3u8
	mm := make(map[string]interface{})
	plug := GetInFlow_com(Uid, intype)
	mm["errMsg"] = plug
	comm.Response(c, mm)
}
func GetInFlow_com(AnchorId int64, Type string) string {
	var wr m.AnchorRoom
	//获取用户ID
	wr.Uid = AnchorId
	intype := Type //1-->//phone   m3u8
	//获取直播房间
	if !wr.GetRoomUid() {
	}
	if wr.LiveAddress == "" {
		return ""
	}
	allid := strconv.FormatInt(wr.Uid, 10) + "-" + strconv.FormatInt(wr.Id, 10)
	start := strings.Index(wr.LiveAddress, "=")
	end := strings.Index(wr.LiveAddress, "&")
	data := SubEndstr(wr.LiveAddress, start+1, end)
	plug := m.InFlow(allid, intype, data)
	if plug == "" {
	}
	return plug
}
