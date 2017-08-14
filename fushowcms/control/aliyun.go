package control

/*
 * Copyright (c) 2016—2017 www.fushow.cn, All rights reserved.
 */
import (
	"fmt"
	"time"

	"fushowcms/comm"
	m "fushowcms/models"
	"strconv"
	"strings"

	"github.com/BPing/aliyun-live-go-sdk/client"
	"github.com/BPing/aliyun-live-go-sdk/device/live"
	"github.com/gin-gonic/gin"
)

const (
	//访问地址：https://ak-console.aliyun.com/#/accesskey
	AccessKeyId     = "" //您的阿里云Id
	AccessKeySecret = "" //您的阿里云Secret
)

// 协程
func GetDoMainName() {
	// 定时为60秒
	ticker := time.NewTicker(1 * time.Minute)
	for {
		<-ticker.C
		cert := client.NewCredentials(AccessKeyId, AccessKeySecret)
		var ss int64 = 1
		for i := 0; i < 10; i++ {
			name := "live" + strconv.FormatInt(ss, 10)
			nowlive := comm.GetConfig("FS", name)
			liveM := live.NewLive(cert, nowlive, `sllmzb`, nil).SetDebug(true)
			resp := live.OnlineInfoResponse{}
			//循环liveList
			liveM.StreamOnlineUserNum("", &resp)
			for i := 0; i < len(resp.OnlineUserInfo.LiveStreamOnlineUserNumInfo); i++ {
				roomid := resp.OnlineUserInfo.LiveStreamOnlineUserNumInfo[i].StreamUrl[strings.LastIndex(resp.OnlineUserInfo.LiveStreamOnlineUserNumInfo[i].StreamUrl, "/")+1 : len(resp.OnlineUserInfo.LiveStreamOnlineUserNumInfo[i].StreamUrl)]
				arr := strings.Split(roomid, "-")
				var anchorLiveNumber m.AnchorLiveNumber
				anchorLiveNumber.RoomId, _ = strconv.ParseInt(arr[1], 10, 64)
				anchorLiveNumber.LiveNumber = SumNumber(resp.OnlineUserInfo.LiveStreamOnlineUserNumInfo[i].UserNumber) //虚拟人数
				anchorLiveNumber.LiveNumberUp()
			}
			ss++
		}
	}
}

//获取所有流
func GetStreams() {
	cert := client.NewCredentials(AccessKeyId, AccessKeySecret)
	liveM := live.NewLive(cert, `live.shiliu88.com`, `sllmzb`, nil).SetDebug(true)

	resp := live.OnlineRoomInfo{}
	liveM.StreamsOnlineListWithApp(`sllmzb`, &resp)
	list := make(map[int]m.AnchorInfo)
	for i := 0; i < len(resp.OnlineInfo.LiveStreamOnlineInfo); i++ {
		var (
			anchorroom m.AnchorRoom
			uid        m.UidInfo
			nowroom    m.AnchorInfo
		)
		roomid := resp.OnlineInfo.LiveStreamOnlineInfo[i].StreamName
		arr := strings.Split(roomid, "-")
		anchorroom.Id, _ = strconv.ParseInt(arr[1], 10, 64)
		uid.Id, _ = strconv.ParseInt(arr[0], 10, 64)
		anchorroom.GetRoom()
		uid.GetUserInfo()
		nowroom.Id = anchorroom.Id
		nowroom.Uid, _ = strconv.ParseInt(arr[0], 10, 64)
		nowroom.RoomAlias = anchorroom.RoomAlias
		nowroom.LiveCover = "http://fushow.oss-cn-shanghai.aliyuncs.com/sllmzb/" + arr[0] + "-" + arr[1] + ".jpg"
		nowroom.LiveAddress = anchorroom.LiveAddress
		nowroom.NickName = uid.NickName
		list[i] = nowroom
	}

	if len(list) == 0 {

	}

}

//在线人数获取
func GetOnlineUser(StreamName string) {
	cert := client.NewCredentials(AccessKeyId, AccessKeySecret)
	liveM := live.NewLive(cert, `live.shiliu88.com`, `sllmzb`, nil).SetDebug(true)
	resp := live.OnlineInfoResponse{}
	liveM.StreamOnlineUserNum(StreamName, &resp)
}

//禁止推流列表
func GetBlockList(c *gin.Context) {
	cert := client.NewCredentials(AccessKeyId, AccessKeySecret)
	liveM := live.NewLive(cert, `live.shiliu88.com`, `sllmzb`, nil).SetDebug(true)
	resp := live.StreamListResponse{}
	liveM.StreamsBlockList(&resp)
	list := make(map[string]m.AnchorRoom)
	if len(resp.StreamUrls.StreamUrl) > 0 {
		for i := 0; i < len(resp.StreamUrls.StreamUrl); i++ {
			var ar m.AnchorRoom
			streamUrl := resp.StreamUrls.StreamUrl[i]
			roomid := Substr(streamUrl, 28, 31)
			arr := strings.Split(roomid, "-")
			ar.Id, _ = strconv.ParseInt(arr[1], 10, 64)
			ar.GetRoom()
			ar.LiveState = 3 //禁止直播
			ar.RoomUp()

			list[string(i)] = ar
		}
	}
	m := make(map[string]interface{})
	m["rows"] = list
	m["total"] = len(list)
	comm.Response(c, m)
}

//首页房间信息

func GetALiRoomInfo(c *gin.Context) {
	page, _ := strconv.Atoi(c.PostForm("page"))
	rows, _ := strconv.Atoi(c.PostForm("rows"))
	index := c.PostForm("index")
	if index != "" {
		rows = 8
	}
	ing_list, ing_list_normal, stick, s_list_normal, total := GetAlllist(page, rows)
	comm.Response(c, gin.H{"data_normal": ing_list, "data_normal_state": ing_list_normal, "data_stick": stick, "data_stick_state": s_list_normal, "total": total})
}

func GetAlllist(page, rows int) ([]m.AnchorRoomss, []m.AnchorRoomss, []m.AnchorRoomss, []m.AnchorInfo, int64) {
	//正在直播的
	var (
		ar              m.AnchorRoom
		ing_list_normal []m.AnchorRoomss
		s_list_normal   []m.AnchorInfo
	)
	//正在直播   page  页数  rows 个数
	ing_list, _, total := ar.GetRoomListIng(page, rows)
	if len(ing_list) < 10 {
		//不足10个房间时

		roomnum := 10 - len(ing_list)

		_, ing_list_normal = ar.GetRoomNoIng(roomnum)
	}
	//精彩推荐
	stick, errs, s_total := ar.GetRoomListStick()
	var stickId []string
	//去重
	for i := 0; i < len(stick); i++ {
		stickId = append(stickId, strconv.FormatInt(stick[i].Id, 10))
	}
	fmt.Println("-------------------", errs, s_total, stick)
	if len(stick) < 5 {
		//不足4个房间时
		roomnum := 5 - len(stick)
		_, s_list_normal = ar.GetRoomStickIng(roomnum, stickId)
	}
	return ing_list, ing_list_normal, stick, s_list_normal, total
}

//回调方法
func PublishCallBack(c *gin.Context) {
	action := c.Request.URL.Query()["action"][0]
	str := c.Request.URL.Query()["id"][0] //房间号
	arr := strings.Split(str, "-")
	var (
		livenumber m.AnchorLiveNumber
		anchorroom m.AnchorRoom
		anchor_con m.AnchorRoomTime
	)
	anchorroom.Id, _ = strconv.ParseInt(arr[1], 10, 64)
	anchorroom.GetRoom()
	//直播记录
	anchor_con.RoomId = anchorroom.Id
	anchor_con.Uid, _ = strconv.ParseInt(arr[0], 10, 64)
	nowtime := time.Now().Format("2006-01-02 15:04:05")
	switch action {
	//主播状态更新
	case `publish`:
		anchorroom.LiveState = 1
		anchorroom.AddTime = time.Now().Unix()
		fmt.Println("time", time.Now().Unix())
		anchor_con.StartTime = nowtime
		if !anchor_con.AnchorRoomTimeAdd() {
			fmt.Println("直播记录添加失败")
		}
	case `publish_done`:
		//更新主播人数及状态
		anchorroom.LiveState = 0
		//		anchorroom.UserNumber = 0
		//		anchorroom.LiveNumber = 0
		livenumber.LiveNumber = 0
		livenumber.RoomId = anchorroom.Id
		if !livenumber.LiveNumberUp() {
			fmt.Println("人数更新失败")
		}
		anchorroom.AddTime = 0
		if !anchor_con.GetRoomTime() {
			fmt.Println("结束时，未找到直播记录")
		}
		anchor_con.EndTime = nowtime
		the_time, _ := time.Parse("2006-01-02 15:04:05", anchor_con.StartTime)
		endtime, _ := time.Parse("2006-01-02 15:04:05", anchor_con.EndTime)
		u_time := the_time.Unix()
		u_endtime := endtime.Unix()
		anchor_con.AnchormTime = (u_endtime - u_time) / 60 //转换分钟
		if !anchor_con.AnchorRoomTimeUpdate() {
			fmt.Println("添加结束时间失败")
		}
	}
	fmt.Println("addtime", anchorroom.AddTime)
	anchorroom.LiveCover = "http://fushow.oss-cn-shanghai.aliyuncs.com/sllmzb/" + strconv.FormatInt(anchorroom.Uid, 10) + "-" + strconv.FormatInt(anchorroom.Id, 10) + ".jpg"
	anchorroom.RoomLiveState()
}

//虚拟人数
func SumNumber(num int64) int64 {
	return num*63 + 6
}
