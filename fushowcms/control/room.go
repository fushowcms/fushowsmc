package control

/*
 * Copyright (c) 2016—2017 www.fushow.cn, All rights reserved.
 */
import (
	"bytes"
	"encoding/base64"
	"fmt"
	"fushowcms/comm"
	m "fushowcms/models"
	"io/ioutil"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

//txl -del
//添加房间
func RoomAdd(c *gin.Context) {
	var ar m.AnchorRoom

	ar.Uid, _ = strconv.ParseInt(c.PostForm("AnchorId"), 10, 64)              //主播Id
	ar.RoomType, _ = strconv.ParseUint(c.PostForm("RoomType"), 10, 64)        //房间分类 0：普通房间，1：竞猜房间
	ar.RoomAlias = c.PostForm("RoomAlias")                                    //房间别名
	ar.LiveCover = c.PostForm("LiveCover")                                    //直播封面
	ar.LiveState, _ = strconv.ParseUint(c.PostForm("LiveState"), 10, 64)      //直播状态	0：未直播，1：直播中
	ar.PeriodsId, _ = strconv.ParseInt(c.PostForm("PeriodsId"), 10, 64)       //期数Id
	ar.ModifyUserId, _ = strconv.ParseInt(c.PostForm("ModifyUserId"), 10, 64) //修改者ID

	if !ar.RoomAdd() {
		c.JSON(200, gin.H{"state": "fail"})
		return
	}
	s := strconv.FormatInt(ar.Id, 64)

	ar.LiveAddress = ":8080/user/roomlive?roomId=" + s
	c.JSON(200, gin.H{"state": "success"})
}

//删除房间
func RoomDel(c *gin.Context) {
	var ar m.AnchorRoom
	if c.PostForm("id") == "" {
		comm.ResponseError(c, 2000)
		return
	}
	ar.Id, _ = strconv.ParseInt(c.PostForm("id"), 10, 64) //ID
	if !ar.RoomDel() {
		comm.ResponseError(c, 2018)
		return
	}
	comm.Response(c, "成功")
}

//修改房间信息
func RoomUp(c *gin.Context) {
	var ar m.AnchorRoom
	ar.Id, _ = strconv.ParseInt(c.PostForm("RoomId"), 10, 64) //ID

	roomtype, _ := strconv.ParseUint(c.PostForm("RoomType"), 10, 64) //房间分类 0：普通房间，1：竞猜房间
	roomAlias := c.PostForm("RoomAlias")                             //房间别名
	liveAddress := c.PostForm("LiveAddress")                         //直播地址
	roomnotice := c.PostForm("RoomNotice")
	liveCover := c.PostForm("LiveCover")                                    //直播封面
	liveState, _ := strconv.ParseUint(c.PostForm("LiveState"), 10, 64)      //直播状态	0：未直播，1：直播中
	modifyUserId, _ := strconv.ParseInt(c.PostForm("ModifyUserId"), 10, 64) //修改者ID

	if !ar.GetRoom() {
		comm.ResponseError(c, 2020)
		return
	}
	ar.RoomNotice = roomnotice
	ar.RoomType = roomtype
	ar.RoomAlias = roomAlias
	ar.LiveAddress = liveAddress
	ar.LiveCover = liveCover
	ar.LiveState = liveState
	ar.ModifyUserId = modifyUserId

	if !ar.RoomUp() {
		comm.ResponseError(c, 2021)
		return
	}

	comm.Response(c, ar)
}

//获取单个房间信息
func GetRoom(c *gin.Context) {
	var (
		ar  m.AnchorRoomInfo
		arc m.AnchorRoomConcern
	)
	uid, _ := strconv.ParseInt(c.PostForm("AnchorId"), 10, 64) //ID

	list := ar.GetPhoneRoom(uid)
	// 获取主播关注数
	count, err := arc.GetMyAttentionCount(uid)
	if err != nil {
		comm.ResponseError(c, 2019)
		return
	}

	m := make(map[string]interface{})
	m["state"] = list
	m["attention"] = GumNumber(count)
	comm.Response(c, m)
}

//获取单个房间信息
func GetRoomByRoomId(c *gin.Context) {
	var (
		ar  m.AnchorRoomInfo
		arc m.AnchorRoomConcern
	)
	roomId, _ := strconv.ParseInt(c.PostForm("RoomId"), 10, 64) //roomid
	if roomId == 0 {
		comm.ResponseError(c, 2028)
		return
	}
	list := ar.GetPhoneRoomByRoomId(roomId)
	if len(list) == 0 {
		comm.ResponseError(c, 2028)
		return
	}
	uid := list[0].AnchorRoom.Uid
	list[0].UidInfo.PassWord = ""
	list[0].UidInfo.Phone = ""
	list[0].UidInfo.IdNumber = ""
	list[0].UidInfo.IdentityPic = ""
	list[0].UidInfo.BankName = ""
	list[0].UidInfo.BankDeposit = ""
	// 获取主播关注数
	count, err := arc.GetMyAttentionCount(uid)
	if err != nil {
		comm.ResponseError(c, 2019)
		return
	}
	m := make(map[string]interface{})
	m["state"] = list
	m["attention"] = GumNumber(count)
	comm.Response(c, m)
}

func GetRoomInfo(c *gin.Context) {
	var ar m.AnchorRoom

	ar.Uid, _ = strconv.ParseInt(c.PostForm("UID"), 10, 64)
	if ar.Uid == 0 {
		comm.ResponseError(c, 2000)
		return
	}
	if !ar.GetRoom() {
		comm.ResponseError(c, 2020)
		return
	}

	comm.Response(c, ar)
}

//txl-root
//获取所有房间列表
func GetRoomList(c *gin.Context) {
	var ar m.AnchorRoom
	page, _ := strconv.Atoi(c.PostForm("page"))
	rows, _ := strconv.Atoi(c.PostForm("rows"))
	inputid := c.PostForm("inputid")
	list, row := ar.GetRoomList(page, rows, inputid)
	if row == 0 {
		c.JSON(200, gin.H{"total": row, "rows": []int{}})
		return
	}
	c.JSON(200, gin.H{"total": row, "rows": list})

}

//获取直播记录表
func GetAnchorRoomTimeList(c *gin.Context) {
	var ar m.AnchorRoomTime
	uid, _ := strconv.ParseInt(c.PostForm("UID"), 10, 64) //ID
	if uid == 0 {
		comm.ResponseError(c, 3175)
		return
	}
	page, _ := strconv.Atoi(c.PostForm("page"))
	rows, _ := strconv.Atoi(c.PostForm("rows"))
	ar.Uid = uid
	flag, list, total := ar.GetAnchorRoomTimeList(page, rows)
	if !flag {
		comm.ResponseError(c, 2019)
		return
	}
	m := make(map[string]interface{})
	m["data"] = list
	m["total"] = total
	comm.Response(c, m)
}

//txl-del
//获取正在直播的房间  分页  page    条数 rows
func GetRoomTypeListIng(c *gin.Context) {
	var ar m.AnchorInfo
	page, _ := strconv.Atoi(c.PostForm("page"))
	rows, _ := strconv.Atoi(c.PostForm("rows"))
	list, err := ar.GetRoomTypeListIng(page, rows)
	if err != nil {
		c.JSON(200, gin.H{"state": false, "errCode ": 2021, "errMsg": "错误信息，请刷新重试"})
		return
	}
	c.JSON(200, gin.H{"state": true, "errCode ": 2000, "errMsg": list})
}

//txl-del
//获取首页推荐房间列表
func GetIndexRoomList(c *gin.Context) {

	var ar m.AnchorInfo
	page, _ := strconv.Atoi(c.PostForm("page"))
	num, _ := strconv.Atoi(c.PostForm("num"))
	if page == 0 && num == 0 {
		page = 1
		num = 8
	}
	list, err := ar.GetIndexRoomList(page, num)
	if err != nil || list == nil {
		c.JSON(200, gin.H{"message": true, "state": "", "page": 1})
		return
	}

	c.JSON(200, gin.H{"message": false, "state": list, "page": page})
}

//txl-del
//获取List推荐房间列表
func GetListRoomList(c *gin.Context) {
	var ar m.AnchorInfo
	page, _ := strconv.Atoi(c.PostForm("page"))
	num, _ := strconv.Atoi(c.PostForm("num"))
	if page == 0 && num == 0 {
		page = 1
		num = 8
	}
	list, err := ar.GetListRoomList(page, num)
	if err != nil || list == nil {
		c.JSON(200, gin.H{"message": true, "state": "", "page": 1})
		return
	}
	c.JSON(200, gin.H{"message": false, "state": list, "page": page})

}

func GetAnchorInfo(c *gin.Context) {
	var ar m.AnchorRoom
	ar.Id, _ = strconv.ParseInt(c.PostForm("RoomId"), 10, 64) //房间Id
	anchorRoomAllInfo, err := ar.GetAnchorInfo()
	if err != nil && anchorRoomAllInfo == nil {
		c.JSON(200, gin.H{"state": "fail"})
		return
	}
	c.JSON(200, gin.H{"state": anchorRoomAllInfo})
}

//取消关注
func CancelOrderRoom(c *gin.Context) {
	var order m.AnchorRoomConcern
	order.Uid, _ = strconv.ParseInt(c.PostForm("UID"), 10, 64)
	order.User, _ = strconv.ParseInt(c.PostForm("User"), 10, 64)
	if c.PostForm("UID") == "" && c.PostForm("User") == "" {
		//参数错误
		comm.ResponseError(c, 2000)
		return
	}

	if !order.GetRoom() {
		//关注不存在
		comm.ResponseError(c, 2022)
		return
	}
	if !order.RoomDel() {
		//取消失败
		comm.ResponseError(c, 2023)
		return
	}
	//取消成功

	comm.Response(c, "成功取消")
}

//根据别名查询(模糊)
func SelRoomAlias(c *gin.Context) {
	var ar m.AnchorInfo
	str := c.PostForm("roomAlias")
	page, _ := strconv.Atoi(c.PostForm("page"))
	rows, _ := strconv.Atoi(c.PostForm("rows"))
	data, err, total := ar.GetListRoomLikeList(str, page, rows)
	if err != nil {
		comm.ResponseError(c, 2019)
		return
	}
	m := make(map[string]interface{})
	m["data"] = data
	m["total"] = total
	comm.Response(c, m)
}

//根据房间id查询所有房管 addby liuhan
func FindByRoomIdAll(c *gin.Context) {
	var rum m.RoomUserManage
	roomId, _ := strconv.ParseInt(c.PostForm("RoomId"), 10, 64)
	page, _ := strconv.Atoi(c.PostForm("page"))
	rows, _ := strconv.Atoi(c.PostForm("rows"))
	if roomId == 0 {
		comm.ResponseError(c, 2000)
		return
	}
	rum.RoomId = roomId
	data, err, total := rum.FindByRoomIdAll(page, rows)
	if err != nil && data == nil {
		comm.ResponseError(c, 2019)
		return
	}

	m := make(map[string]interface{})
	m["data"] = data
	m["total"] = total
	comm.Response(c, m)

}

//增加房管 addby liuhan
func AddRoomUserManage(c *gin.Context) {
	var rum m.RoomUserManage
	roomId, _ := strconv.ParseInt(c.PostForm("RoomId"), 10, 64)
	userId, _ := strconv.ParseInt(c.PostForm("UID"), 10, 64)
	if roomId == 0 {
		comm.ResponseError(c, 2000)
		return
	}
	if userId == 0 {
		comm.ResponseError(c, 2000)
		return
	}

	var user m.UidInfo
	user.Id = userId

	if !user.GetUserInfo() {
		comm.ResponseError(c, 2004)
		return
	}

	rum.RoomId = roomId
	rum.UserId = userId
	rum.ModifyBy = userId
	rum.NickName = user.NickName

	//根据房间id和用户id判断该用户是否是房管
	data, _ := rum.FindByRoomIdUserIdAll()
	if len(data) != 0 {
		comm.ResponseError(c, 2024)
		return
	}
	if !rum.AddRoomUserManage() {
		comm.ResponseError(c, 2025)
		return
	}
	comm.Response(c, "更新成功")
}

//删除房管 addby liuhan
func DelRoomUserManage(c *gin.Context) {
	var rum m.RoomUserManage
	rum.Id, _ = strconv.ParseInt(c.PostForm("Id"), 10, 64) //ID
	if rum.Id == 0 {
		comm.ResponseError(c, 2000)
		return
	}
	//判断id是否存在
	data, _ := rum.FindByIdAll()
	if len(data) == 0 {
		comm.ResponseError(c, 2004)
		return
	}
	if !rum.DelRoomUserManage() {
		comm.ResponseError(c, 2026)
		return
	}
	comm.Response(c, "您成功取消了")
}

//直播间中删除房管 addby wenhan
func DelFrontUserManage(c *gin.Context) {
	var rum m.RoomUserManage

	roomId, _ := strconv.ParseInt(c.PostForm("RoomId"), 10, 64)
	userId, _ := strconv.ParseInt(c.PostForm("UID"), 10, 64)

	rum.RoomId = roomId
	rum.UserId = userId

	//判断id是否存在
	data, _ := rum.FindByRoomIdUserIdAll()
	if len(data) == 0 {
		comm.ResponseError(c, 2027)
		return
	}
	if !rum.DelRoomUserManage() {
		comm.ResponseError(c, 2026)
		return
	}
	comm.Response(c, "您成功取消该用户的房管")
}

//添加禁言  addby liuhan
func AddNotUserSpeak(c *gin.Context) {
	var nus m.NotUserSpeak

	nus.UserId, _ = strconv.ParseInt(c.PostForm("UID"), 10, 64) //用户Id
	nus.RoomId, _ = strconv.ParseInt(c.PostForm("RoomId"), 10, 64)
	if nus.UserId == 0 {
		comm.ResponseError(c, 3175)
		return
	}

	if !nus.AddNotUserSpeak() {
		comm.ResponseError(c, 2018)
		return
	}
	comm.Response(c, "成功")
}

/*筛选房间  txl  2016-11-15*/
func ChooseRoom(c *gin.Context) {
	choose := c.PostForm("Choose")
	if choose == "" {
		c.JSON(200, gin.H{"total": 0, "rows": ""})
		return
	}
	stick, _ := strconv.ParseInt(c.PostForm("Staick"), 10, 64)
	page, _ := strconv.Atoi(c.PostForm("page"))
	rows, _ := strconv.Atoi(c.PostForm("rows"))

	var ar m.AnchorRoom
	chooses, _ := strconv.ParseInt(choose, 10, 64)
	if chooses == 1 {
		//禁封 	LiveState-->2
		ar.LiveState = 2
	} else if chooses == 2 {
		//直播中 LiveState-->1
		ar.LiveState = 1
	} else if chooses == 3 {
		//休息中 LiveState-->0
		ar.LiveState = 0
	}
	//1-->置顶中

	if stick == 1 { //1-->置顶中
		//1-->置顶中
		ar.RoomStick = 1
	} else {
		ar.RoomStick = 0
	}

	flag, list, tp := ar.ChooseRoom(page, rows)
	if !flag {
		c.JSON(200, gin.H{"total": 0, "rows": ""})
		return
	}
	c.JSON(200, gin.H{"total": tp, "rows": list})
}

//后台获取直播记录表 addby liuhan
func GetRootPlayList(c *gin.Context) {
	var ar m.AnchorRoomTime
	fmt.Println("uid1", c.PostForm("userId"))
	uid, _ := strconv.ParseInt(c.PostForm("userId"), 10, 64) //ID
	page, _ := strconv.Atoi(c.PostForm("page"))
	rows, _ := strconv.Atoi(c.PostForm("rows"))
	ar.Uid = uid

	_, list, total := ar.GetRootPlayList(page, rows)

	if total == 0 {
		c.JSON(200, gin.H{"total": total, "rows": []int{}})
		return
	}
	c.JSON(200, gin.H{"total": total, "rows": list})
}

//通过房间Id查找房间信息
func GetRoomByRId(c *gin.Context) {
	roomid := c.PostForm("RoomId")
	if roomid == "" {
		comm.ResponseError(c, 2000)
		return
	}

	var rm m.AnchorRoom
	rm.Id, _ = strconv.ParseInt(roomid, 10, 64)
	if !rm.GetRoom() {
		comm.ResponseError(c, 2048)
		return
	}

	comm.Response(c, rm.Uid)
}

func PandaInfo(c *gin.Context) {
	roomid := c.Request.URL.Query()["roomid"][0]
	if roomid == "" {
		comm.ResponseError(c, 2000)
	}
	response, _ := http.Get("http://www.panda.tv/api_room_v2?roomid=" + roomid)
	defer response.Body.Close()
	body, _ := ioutil.ReadAll(response.Body)
	comm.Response(c, string(body))
}

//关注数虚拟人数
func GumNumber(num int64) int64 {
	return num*441 + 6
}

//连接房间Im
func GetRoomIm(c *gin.Context) {
	m := make(map[string]interface{})
	if c.PostForm("UID") == "" {
		m["username"] = "shiliuAide"
		m["password"] = "shiliutv"
		comm.Response(c, m)
		return
	}

	var ui SessionUidInfo
	var isPC bool = true
	ui.Id, _ = strconv.ParseInt(c.PostForm("UID"), 10, 64)

	//判断访问来源
	if c.Request.UserAgent() == "fushowphone:ios" || c.Request.UserAgent() == "fushowphone:android" {
		isPC = false
	}

	if !GetSess(&ui, isPC) {
		m["username"] = "shiliuAide"
		m["password"] = "shiliutv"
		comm.Response(c, m)
		return
	}

	coordinate, _ := strconv.ParseInt(getNumberIndex(1), 10, 64)
	mData := GetApolloIM(coordinate)
	m["username"] = mData.UserName
	m["password"] = ByteToHex(mData.PassWord)

	comm.Response(c, m)
}

//byte转16进制字符串
func ByteToHex(data []byte) string {
	buffer := new(bytes.Buffer)
	for _, b := range data {

		s := strconv.FormatInt(int64(b&0xff), 16)
		if len(s) == 1 {
			buffer.WriteString("0")
		}
		buffer.WriteString(s)
	}

	return buffer.String()
}

func base64Decode(src string) ([]byte, error) {
	return base64.StdEncoding.DecodeString(src)
}

//生成随机字符串
func getNumberIndex(leng int) string {
	str := "0123456789"
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano() * rand.Int63()))

	for i := 0; i < leng; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}

type UserLevel struct {
	Level int
	Name  string
}

//获取用户等级列表与等级名称的对应
func GetLevelMap(c *gin.Context) {

	comm.Response(c, []UserLevel{{0, "白板"}, {1, "青铜5"}, {2, "青铜4"}, {3, "青铜3"}, {4, "青铜2"}, {5, "青铜1"}, {6, "白银5"}, {7, "白银4"}, {8, "白银3"}, {9, "白银2"}, {10, "白银1"}, {11, "黄金5"}, {12, "黄金4"}, {13, "黄金3"}, {14, "黄金2"}, {15, "黄金1"}, {16, "铂金5"}, {17, "铂金4"}, {18, "铂金3"}, {19, "铂金2"}, {20, "铂金1"}, {21, "钻石5"}, {22, "钻石4"}, {23, "钻石3"}, {24, "钻石2"}, {25, "钻石1"}, {26, "大师"}, {27, "王者"}})
}

func FindByRoomIdUserIdAll_com(roomId, userId int64) []m.RoomUserManage {
	var rum m.RoomUserManage

	if roomId == 0 {
		fmt.Println("房间ID为空")
	}
	if userId == 0 {
		fmt.Println("用户ID为空")
	}
	rum.RoomId = roomId
	rum.UserId = userId
	data, _ := rum.FindByRoomIdUserIdAll()
	if data == nil {
		fmt.Println("暂无房间")
	}
	return data
}
func FindNotUserSpeakAll_com(UID, RoomId int64) []m.NotUserSpeak {
	var nus m.NotUserSpeak

	nus.UserId = UID
	nus.RoomId = RoomId

	data, err := nus.FindNotUserSpeakAll()
	if err != nil && data == nil {
	}
	return data
}
func GetRoom_com(AnchorId int64) ([]m.AnchorRoomInfo, int64) {
	var (
		ar  m.AnchorRoomInfo
		arc m.AnchorRoomConcern
	)
	uid := AnchorId

	list := ar.GetPhoneRoom(uid)
	// 获取主播关注数
	count, err := arc.GetMyAttentionCount(uid)
	if err != nil {

	}

	m := make(map[string]interface{})
	m["state"] = list
	m["attention"] = GumNumber(count)
	return list, count
	//	comm.Response(c, m)
}

//根据分类查询 addby liuhan
func GetRoomAliasByRoomType(c *gin.Context) {
	var ar m.AnchorInfo
	roomType := c.PostForm("RoomType")
	fmt.Println(roomType)
	page, _ := strconv.Atoi(c.PostForm("page"))
	rows, _ := strconv.Atoi(c.PostForm("rows"))
	data, err, total := ar.GetRoomAliasByRoomType(roomType, page, rows)
	if err != nil {
		comm.ResponseError(c, 2019)
		return
	}
	m := make(map[string]interface{})
	m["data"] = data
	m["total"] = total
	comm.Response(c, m)
}

// 热门直播8个 addby liuhan 161227
func GetHotLiveList(c *gin.Context) {
	var (
		ar            m.AnchorRoom
		s_list_normal []m.AnchorInfo
	)
	//精彩推荐8个
	stick, _, _ := ar.GetHotLiveList()
	var stickId []string
	//去重
	for i := 0; i < len(stick); i++ {
		stickId = append(stickId, strconv.FormatInt(stick[i].Id, 10))
	}
	if len(stick) < 8 {
		//不足4个房间时
		roomnum := 8 - len(stick)
		_, s_list_normal = ar.GetRoomStickIng(roomnum, stickId)
	}
	comm.Response(c, gin.H{"data_stick": stick, "data_stick_state": s_list_normal})

}
