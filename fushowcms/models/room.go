package models

/*
 * Copyright (c) 2016—2017 www.fushow.cn, All rights reserved.
 */
import (
	"bytes"
	"fmt"
	"fushowcms/comm"
	"strconv"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

//增加房间
func (am *AnchorRoom) RoomAdd() bool {
	row, err := Engine.Insert(am) //插入一行数据
	if row > 0 || err == nil {
		return true
	}
	return false
}

//增加房间
func (aln *AnchorLiveNumber) LiveNumberAdd() bool {
	row, err := Engine.Insert(aln) //插入一行数据
	if row > 0 || err == nil {
		return true
	}
	return false
}

//删除房间id
func (am *AnchorRoom) RoomDel() bool {
	row, err := Engine.Delete(am) //插入一行数据
	if row > 0 || err == nil {
		return true
	}
	return false
}

//更新房间信息
func (am *AnchorRoom) RoomUp() bool {
	comm.ShowStack()
	row, err := Engine.Where("Id=?", am.Id).Update(am)
	if row > 0 && err == nil {
		return true
	}
	return false
}

//更新房间在线人数
func (aln *AnchorLiveNumber) LiveNumberUp() bool {
	sql := "update `anchor_live_number` set `live_number`=? where `room_id`=?;"
	_, err := Engine.Exec(sql, aln.LiveNumber, aln.RoomId)
	if err != nil {
		return false
	}
	return true
}

//更新房间信息
func (am *AnchorRoom) RoomUpStaick() bool {
	Engine.ShowSQL(true)
	sql := "UPDATE anchor_room set room_stick = '1' WHERE id = ?"
	_, err := Engine.Exec(sql, am.Id)
	if err != nil {
		return false
	}
	return true
}

//恢复直播更新
func (am *AnchorRoom) RoomLiveState() bool {
	Engine.ShowSQL(true)
	num, err := Engine.Table("anchor_room").Where("id=?", am.Id).Cols("live_state", "user_number", "live_number", "add_time", "live_cover").Update(am)
	if num > 0 && err == nil {
		return true
	}
	return false
}

//获取phone单个房间信息
func (ar *AnchorRoomInfo) GetPhoneRoom(uid int64) []AnchorRoomInfo {
	var list []AnchorRoomInfo
	Engine.ShowSQL(true)
	Engine.Table("anchor_room").Alias("r").Join("LEFT", []string{"uid_info", "u"}, "r.uid = u.id").Join("LEFT", []string{"anchor_live_number", "aln"}, "r.id = aln.room_id").Join("LEFT", []string{"category_two", "ct"}, "r.room_type = ct.id").Where("r.uid=?", uid).Find(&list)
	return list
}

func (ar *AnchorRoomInfo) GetAnchorRoomById(roomId int64) []AnchorRoomInfo {
	var list []AnchorRoomInfo
	Engine.ShowSQL(true)
	Engine.Table("anchor_room").Alias("r").Join("LEFT", []string{"uid_info", "u"}, "r.uid = u.id").Where("r.id=?", roomId).Find(&list)

	return list
}

//获取单个房间信息
func (am *AnchorRoom) GetRoom() bool {
	flag, _ := Engine.Get(am)
	return flag
}

//通过主播Id获得房间是否存在
func (am *AnchorRoom) GetRoomUid() bool {
	flag, _ := Engine.Table("anchor_room").Where("uid=?", am.Uid).Get(am)
	return flag
}

//取消置顶sql
func (ar *AnchorRoom) CalcelStaick(id, uid int64) bool {
	sql := "update `anchor_room` set room_stick=0 ,modify_user_id=? where id=?"
	_, err := Engine.Exec(sql, uid, id)
	if err != nil {
		return false
	}
	return true
}

//直播房间表
type AnchorRooms struct {
	NickName     string //真实姓名
	Id           int64  //房间Id	房间Id：主播Id = 1:1
	Uid          int64  `xorm:"index"` //主播Id
	RoomType     uint64 `xorm:"index"` //房间分类	0:普通房间，1:竞猜房间
	RoomAlias    string //房间别名
	RoomNotice   string //房间公告
	LiveAddress  string //直播地址
	LiveCover    string //直播封面
	LiveState    uint64 //直播状态	0:未直播，1:直播中 2 禁止推流
	RoomStick    int64  //是否置顶   0:不置顶 1:置顶
	PeriodsId    int64  //期Id	房管设定，包括房间分类和期
	UserNumber   int64  //在线人数
	LiveNumber   int64  //虚拟人数
	AddTime      int64  //增加时间	添加时间
	ModifyTime   string `xorm:"datetime updated"` //修改时间	修改时间
	DelTime      string `xorm:"datetime deleted"` //记录删除时间
	ModifyUserId int64  `xorm:"index"`            //用户Id	记录修改者账号
	Version      int64  `xorm:"version"`
}

//获取所有房间列表
func (ar *AnchorRoom) GetRoomList(page, rows int, inputid string) ([]AnchorRooms, int64) {
	var list []AnchorRooms
	total, _ := Engine.Where("id >?", 0).Count(ar)
	Engine.Table("anchor_room").Alias("r").Join("LEFT", []string{"uid_info", "u"}, "r.uid = u.id").Where("room_alias like ?", "%"+inputid+"%").Desc("r.id").Limit(rows, (page-1)*rows).Find(&list)
	return list, total
}

//直播房间表
type AnchorRoomss struct {
	Id              int64  //房间Id	房间Id：主播Id = 1:1
	Uid             int64  `xorm:"index"` //主播Id
	RoomType        uint64 `xorm:"index"` //房间分类	0:普通房间，1:竞猜房间
	NickName        string
	RoomAlias       string //房间别名
	RoomNotice      string //房间公告
	LiveAddress     string //直播地址
	LiveCover       string //直播封面
	LiveState       uint64 //直播状态	0:未直播，1:直播中 2 禁止推流
	RoomStick       int64  //是否置顶   0:不置顶 1:置顶
	PeriodsId       int64  //期Id	房管设定，包括房间分类和期
	UserNumber      int64  //在线人数
	LiveNumber      int64  //虚拟人数
	AddTime         int64  //增加时间	添加时间
	ModifyTime      string `xorm:"datetime updated"` //修改时间	修改时间
	DelTime         string `xorm:"datetime deleted"` //记录删除时间
	ModifyUserId    int64  `xorm:"index"`            //用户Id	记录修改者账号
	Version         int64  `xorm:"version"`
	TwoCategoryName string
}

//获取正在直播房间列表
func (ar *AnchorRoom) GetRoomListIng(page, rows int) ([]AnchorRoomss, error, int64) {
	var list []AnchorRoomss
	Engine.ShowSQL(true)
	total, _ := Engine.Table("anchor_room").Alias("r").Join("LEFT", []string{"uid_info", "u"}, "r.uid = u.id").Join("LEFT", []string{"anchor_live_number", "aln"}, "aln.room_id = r.id").Where("r.live_state=1").Count(ar)

	err := Engine.Table("anchor_room").Alias("r").Join("LEFT", []string{"uid_info", "u"}, "r.uid = u.id").Join("LEFT", []string{"anchor_live_number", "aln"}, "aln.room_id = r.id").Where("r.live_state=1").Desc("aln.live_number").Limit(rows, (page-1)*rows).Find(&list)
	return list, err, total
}

//获取精彩推荐的房间
func (ar *AnchorRoom) GetRoomListStick() ([]AnchorRoomss, error, int64) {
	var list []AnchorRoomss
	Engine.ShowSQL(true)
	total, _ := Engine.Table("anchor_room").Alias("r").Join("LEFT", []string{"uid_info", "u"}, "r.uid = u.id").Join("LEFT", []string{"anchor_live_number", "aln"}, "aln.room_id = r.id").Join("LEFT", []string{"category_two", "ct"}, "r.room_type = ct.id").Where("1=1").Count(ar)
	err := Engine.Table("anchor_room").Alias("r").Join("LEFT", []string{"uid_info", "u"}, "r.uid = u.id").Join("LEFT", []string{"anchor_live_number", "aln"}, "aln.room_id = r.id").Join("LEFT", []string{"category_two", "ct"}, "r.room_type = ct.id").Where("1=1").Desc("aln.live_number").Limit(5, 0).Find(&list)
	return list, err, total
}

//不足8个房间时补充
func (ar *AnchorRoom) GetRoomNoIng(rows int) (bool, []AnchorRoomss) {
	var list []AnchorRoomss
	Engine.ShowSQL(true)
	err := Engine.Table("anchor_room").Alias("r").Join("LEFT", []string{"uid_info", "u"}, "r.uid = u.id").Join("LEFT", []string{"anchor_live_number", "aln"}, "aln.room_id = r.id").Join("LEFT", []string{"anchor_room_time", "art"}, "art.id = u.id").Where("r.live_state=0").Desc("art.end_time").Limit(rows, 0).Find(&list)
	if err != nil {
		return false, list
	}
	return true, list
}

//不足4个房间时补充   竞猜直播
func (ar *AnchorRoom) GetRoomStickIng(rows int, stickId []string) (bool, []AnchorInfo) {
	var list []AnchorInfo
	Engine.ShowSQL(true)
	sql := bytes.Buffer{}
	if len(stickId) == 0 {
		sql.WriteString("SELECT * FROM `anchor_room` AS `r` LEFT JOIN `uid_info` AS `u` ON r.uid = u.id  LEFT JOIN `anchor_live_number` AS `aln` ON r.id = aln.room_id LEFT JOIN `category_two` as `ct` ON r.room_type = ct.id  WHERE r.live_state=1 ")
	} else {
		sql.WriteString("SELECT * FROM `anchor_room` AS `r` LEFT JOIN `uid_info` AS `u` ON r.uid = u.id LEFT JOIN `anchor_live_number` AS `aln` ON r.id = aln.room_id LEFT JOIN `category_two` as `ct` ON r.room_type = ct.id WHERE r.live_state=1 and r.Id NOT IN (")
		mm := 1
		for i := 0; i < len(stickId); i++ {
			if mm == len(stickId) {
				sql.WriteString(stickId[i] + ",-9999")
			} else {
				sql.WriteString(stickId[i] + ",")
			}
			mm++
		}
		sql.WriteString(")")
	}
	sql.WriteString(" ORDER BY `aln`.`live_number` DESC LIMIT ")
	sql.WriteString(strconv.FormatInt(int64(rows), 10) + "")
	ss := sql.String()
	err := Engine.SQL(ss).Find(&list)
	if err != nil {
		return false, list
	}
	return true, list
}

//获取正在直播列表   page  分页  num 条数
func (ar *AnchorInfo) GetRoomTypeListIng(page, num int) ([]AnchorInfo, error) {
	var list []AnchorInfo
	err := Engine.Table("anchor_room").Alias("r").Join("INNER", []string{"uid_info", "u"}, "r.uid = u.id").Where("r.live_state=1 ").Desc("r.user_number").Limit(num, (page-1)*num).Find(&list)
	return list, err
}

//获取前八个房间列表
func (rm *AnchorInfo) GetIndexRoomList(page int, num int) ([]AnchorInfo, error) {
	var list []AnchorInfo
	err := Engine.Table("anchor_room").Alias("r").Join("INNER", []string{"uid_info", "u"}, "r.uid  = u.id").Asc("r.id").Limit(num, (page-1)*num).Find(&list)

	return list, err
}

//获取所有房间列表
func (rm *AnchorInfo) GetListRoomList(page int, num int) ([]AnchorInfo, error) {
	var list []AnchorInfo
	err := Engine.Table("anchor_room").Alias("r").Join("INNER", []string{"uid_info", "u"}, "r.uid = u.id").Asc("r.id").Limit(num, (page-1)*num).Find(&list)

	return list, err
}

//获取直播记录表
func (rm *AnchorRoomTime) GetAnchorRoomTimeList(page, rows int) (bool, []AnchorRoomTime, int64) {
	var list []AnchorRoomTime
	total, _ := Engine.Table("anchor_room_time").Alias("t").Where("t.uid =?", rm.Uid).Count(rm)
	err := Engine.Table("anchor_room_time").Alias("t").Where("t.uid =?", rm.Uid).Desc("t.id").Limit(rows, (page-1)*rows).Find(&list)
	if err != nil {
		return false, list, total
	}
	return true, list, total
}

//获取直播记录表
func (rm *AnchorRoomTime) GetRootPlayList(page, rows int) (bool, []AnchorRoomTime, int64) {
	var list []AnchorRoomTime
	if rm.Uid == 0 {
		total, _ := Engine.Table("anchor_room_time").Alias("t").Count(rm)
		err := Engine.Table("anchor_room_time").Alias("t").Desc("start_time").Limit(rows, (page-1)*rows).Find(&list)
		if err != nil {
			return false, list, total
		}
		Engine.ShowSQL(true)
		return true, list, total
	} else {
		total, _ := Engine.Table("anchor_room_time").Alias("t").Where("t.uid =?", rm.Uid).Count(rm)
		err := Engine.Table("anchor_room_time").Alias("t").Where("t.uid =?", rm.Uid).Desc("start_time").Limit(rows, (page-1)*rows).Find(&list)
		if err != nil {
			return false, list, total
		}
		Engine.ShowSQL(true)
		return true, list, total
	}

}

//获取当前房间主播所有信息
func (rm *AnchorRoom) GetAnchorInfo() ([]AnchorRoomAllInfo, error) {
	var list []AnchorRoomAllInfo
	err := Engine.Table("anchor_room").Alias("r").Join("LEFT", []string{"uid_info", "u"}, "r.uid = u.id").Where("r.id =?", rm.Id).Find(&list)
	Engine.ShowSQL(true)
	return list, err
}

//通过联合查询，返回【用户】和【房间】列表
func (ui *AnchorRoom) SearchRoomList(str string) ([]AnchorRoom, []UidInfo) {
	var (
		listRoom []AnchorRoom
		listUser []UidInfo
	)

	if str == "" {
	} else {
		err1 := Engine.Table("uid_info").Alias("u").Join("LEFT", []string{"anchor_room", "r"}, "u.id = r.uid").Where("r.id like ? or r.room_alias LIKE ? OR u.nick_name LIKE ?", "%"+str+"%", "%"+str+"%", "%"+str+"%").Find(&listRoom)
		err2 := Engine.Table("uid_info").Alias("u").Join("LEFT", []string{"anchor_room", "r"}, "u.id = r.uid").Where("r.id like ? or r.room_alias LIKE ? OR u.nick_name LIKE ?", "%"+str+"%", "%"+str+"%", "%"+str+"%").Find(&listUser)
		fmt.Println("err1", err1)
		fmt.Println("err2", err2)
	}
	return listRoom, listUser
}

//根据房间别名查询
func (rm *AnchorInfo) GetListRoomLikeList(str string, page int, num int) ([]AnchorInfo, error, int64) {
	var list []AnchorInfo
	total, _ := Engine.Table("anchor_room").Alias("r").Join("INNER", []string{"uid_info", "u"}, "r.uid = u.id").Where("r.id like ? or r.room_alias like ? or u.nick_name like ?", "%"+str+"%", "%"+str+"%", "%"+str+"%").Count(rm)
	err := Engine.Table("anchor_room").Alias("r").Join("INNER", []string{"uid_info", "u"}, "r.uid = u.id").Where("r.id like ? or r.room_alias like ? or u.nick_name like ?", "%"+str+"%", "%"+str+"%", "%"+str+"%").Asc("r.id").Limit(num, (page-1)*num).Find(&list)
	return list, err, total
}

/* 功能：判断房间是否存在
 * 添加一期产品时 首先判断此房间是否存在
 * @author 徐林->新增
 * @Time 20161103
 */
func (ar *AnchorRoom) IsRoomExist() bool {
	has, err := Engine.Table("anchor_room").Where("id=?", ar.Id).Get(ar)
	if err != nil {
		return false
	}
	return has
}

//根据id查询 addby liuhan
func (rum *RoomUserManage) FindByIdAll() ([]RoomUserManage, error) {
	var list []RoomUserManage
	err := Engine.Table("room_user_manage").Where("id =?", rum.Id).Find(&list)
	return list, err
}

//根据房间id查询所有房管 addby liuhan
func (rum *RoomUserManage) FindByRoomIdAll(page, rows int) ([]RoomUserManage, error, int64) {
	var list []RoomUserManage
	err := Engine.Table("room_user_manage").Where("room_id =?", rum.RoomId).Limit(rows, (page-1)*rows).Find(&list)
	total, _ := Engine.Table("room_user_manage").Where("room_id =?", rum.RoomId).Count(rum)
	return list, err, total
}

//根据房间id和用户id查询所有房管 addby liuhan
func (rum *RoomUserManage) FindByRoomIdUserIdAll() ([]RoomUserManage, error) {
	var list []RoomUserManage
	err := Engine.Table("room_user_manage").Where("room_id =? and user_id=?", rum.RoomId, rum.UserId).Find(&list)
	return list, err
}

//增加房管 addby liuhan
func (rum *RoomUserManage) AddRoomUserManage() bool {
	row, err := Engine.Insert(rum) //插入一行数据
	if row > 0 || err == nil {
		return true
	}
	return false
}

//删除房管 addby liuhan
func (rum *RoomUserManage) DelRoomUserManage() bool {
	row, err := Engine.Delete(rum)
	if row > 0 || err == nil {
		return true
	}
	return false
}

//查看置顶房间数量
func (ar *AnchorRoom) GetStaick() int64 {
	total, _ := Engine.Where("room_stick=1").Count(ar)
	return total
}

//添加禁言  addby liuhan
func (num *NotUserSpeak) AddNotUserSpeak() bool {
	row, err := Engine.Insert(num) //插入一行数据
	if row > 0 || err == nil {
		return true
	}
	return false
}

//查询房间内禁言人员 addby liuhan
func (rum *NotUserSpeak) FindNotUserSpeakAll() ([]NotUserSpeak, error) {
	var list []NotUserSpeak
	k := time.Now()
	d, _ := time.ParseDuration("-0.25h")
	nowtime := k.Add(d).Format("2006-01-02 15:04:05")
	err := Engine.Table("not_user_speak").Where("room_id =? and user_id =? and modify_time >=?", rum.RoomId, rum.UserId, nowtime).Find(&list)
	Engine.ShowSQL(true)
	return list, err
}

//房间筛选  txl  2014-11-15
func (ar *AnchorRoom) ChooseRoom(page, num int) (bool, []AnchorRoom, int64) {
	var list []AnchorRoom
	tp, _ := Engine.Table("anchor_room").Where("id>0").Count(ar)
	if ar.RoomStick == 3 {
		err := Engine.Where("live_state=?", ar.LiveState).Limit(num, (page-1)*num).Find(&list)
		if err != nil {
			return false, list, tp
		}
		return true, list, tp
	}
	err := Engine.Where("live_state=? and room_stick=?", ar.LiveState, ar.RoomStick).Limit(num, (page-1)*num).Find(&list)
	if err != nil {
		return false, list, tp
	}
	return true, list, tp
}

//根据分类查询 addby liuhan
func (rm *AnchorInfo) GetRoomAliasByRoomType(roomType string, page int, num int) ([]AnchorInfo, error, int64) {
	var list []AnchorInfo
	Engine.ShowSQL(true)
	total, _ := Engine.Table("anchor_room").Alias("r").Join("LEFT", []string{"uid_info", "u"}, "r.uid = u.id").Join("LEFT", []string{"category_two", "ct"}, "r.room_type = ct.id").Join("LEFT", []string{"anchor_live_number", "aln"}, "aln.room_id = r.id").Where("r.room_type = ? ", roomType).Count(rm)
	err := Engine.Table("anchor_room").Alias("r").Join("LEFT", []string{"uid_info", "u"}, "r.uid = u.id").Join("LEFT", []string{"category_two", "ct"}, "r.room_type = ct.id").Join("LEFT", []string{"anchor_live_number", "aln"}, "aln.room_id = r.id").Where("r.room_type = ? ", roomType).Asc("r.id").Limit(num, (page-1)*num).Find(&list)
	return list, err, total
}

//根据分类查询 addby liuhan
func (rm *AnchorInfo) GetRoomAliasByRoomTypes(roomType string) ([]AnchorInfo, error, int64) {
	var list []AnchorInfo
	Engine.ShowSQL(true)
	total, _ := Engine.Table("anchor_room").Alias("r").Join("LEFT", []string{"uid_info", "u"}, "r.uid = u.id").Join("LEFT", []string{"category_two", "ct"}, "r.room_type = ct.id").Join("LEFT", []string{"anchor_live_number", "aln"}, "aln.room_id = r.id").Where("r.room_type = ? ", roomType).Count(rm)
	err := Engine.Table("anchor_room").Alias("r").Join("LEFT", []string{"uid_info", "u"}, "r.uid = u.id").Join("LEFT", []string{"category_two", "ct"}, "r.room_type = ct.id").Join("LEFT", []string{"anchor_live_number", "aln"}, "aln.room_id = r.id").Where("r.room_type = ? ", roomType).Asc("r.id").Limit(4).Find(&list)
	return list, err, total
}

// 热门直播8个 addby liuhan 161227
func (ar *AnchorRoom) GetHotLiveList() ([]AnchorInfo, error, int64) {
	var list []AnchorInfo
	Engine.ShowSQL(true)
	total, _ := Engine.Table("anchor_room").Alias("r").Join("LEFT", []string{"uid_info", "u"}, "r.uid = u.id").Join("LEFT", []string{"anchor_live_number", "aln"}, "r.id = aln.room_id").Join("LEFT", []string{"category_two", "ct"}, "r.room_type = ct.id").Where("1 =1").Count(ar)
	err := Engine.Table("anchor_room").Alias("r").Join("LEFT", []string{"uid_info", "u"}, "r.uid = u.id").Join("LEFT", []string{"anchor_live_number", "aln"}, "r.id = aln.room_id").Join("LEFT", []string{"category_two", "ct"}, "r.room_type = ct.id").Where("1 =1").Desc("aln.live_number").Limit(8).Find(&list)
	return list, err, total
}

//获取phone单个房间信息
func (ar *AnchorRoomInfo) GetPhoneRoomByRoomId(roomId int64) []AnchorRoomInfo {
	var list []AnchorRoomInfo
	Engine.Table("anchor_room").Alias("r").Join("LEFT", []string{"uid_info", "u"}, "r.uid = u.id").Join("LEFT", []string{"anchor_live_number", "aln"}, "r.id = aln.room_id").Where("r.id=?", roomId).Find(&list)
	return list
}
func FindByRoomIdUserIdAll_com(roomId, userId int64) []RoomUserManage {
	var rum RoomUserManage
	if roomId == 0 {
	}
	if userId == 0 {
	}
	rum.RoomId = roomId
	rum.UserId = userId
	data, _ := rum.FindByRoomIdUserIdAll()
	if data == nil {
	}
	return data
}
