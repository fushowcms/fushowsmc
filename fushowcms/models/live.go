package models

/*
 * Copyright (c) 2016—2017 www.fushow.cn, All rights reserved.
 */
import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"strconv"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

//查询观看记录
func (wr *WatchRecordList) GetWatchRecord(page, rows int) ([]WatchRecordList, error) {
	var (
		list []WatchRecordList
	)
	err := Engine.Table("watch_record").Alias("wr").Join("INNER", []string{"anchor_room", "ar"}, "wr.live_id = ar.id").Join("INNER", []string{"uid_info", "ui"}, "wr.uid = ui.id").Where("wr.uid=?", wr.Uid).Asc("wr.id").Limit(rows, (page-1)*rows).Find(&list)
	return list, err
}

//推流地址返回
func PlugFlow(roomid, str string) string {
	nowtime := time.Now().Unix()
	nowtime = nowtime + 31536000
	//一年后过期
	time := strconv.FormatInt(nowtime, 10)
	b := bytes.Buffer{}
	b.WriteString("/sllmzb/")
	b.WriteString(roomid)
	b.WriteString("-")
	b.WriteString(time)
	b.WriteString("-0-0")
	b.WriteString("-fushowcms")
	ss := b.String()
	h := md5.New()
	h.Write([]byte(ss))
	cipherStr := h.Sum(nil)
	hex.EncodeToString(cipherStr)
	md := hex.EncodeToString(cipherStr)
	if str == "" {
	}
	return roomid + "?vhost=" + str + "&auth_key=" + time + "-0-0-" + md
}

//推流地址返回
func InFlow(allid, intype, liveAddress string) string {
	nowtime := time.Now().Unix()
	time := strconv.FormatInt(nowtime, 10)
	b := bytes.Buffer{}
	if intype == "1" { //phone
		b.WriteString("/sllmzb/")
		b.WriteString(allid)
		b.WriteString(".m3u8-")
		b.WriteString(time)
		b.WriteString("-0-0")
		b.WriteString("-fushowcms")
		ss := b.String()
		h := md5.New()
		h.Write([]byte(ss))
		cipherStr := h.Sum(nil)
		hex.EncodeToString(cipherStr)
		md := hex.EncodeToString(cipherStr)
		return "http://" + liveAddress + "/sllmzb/" + allid + ".m3u8?auth_key=" + time + "-0-0-" + md
	}

	b.WriteString("/sllmzb/")
	b.WriteString(allid)
	b.WriteString("-")
	b.WriteString(time)
	b.WriteString("-0-0")
	b.WriteString("-fushowcms")
	ss2 := b.String()
	h2 := md5.New()
	h2.Write([]byte(ss2))
	cipherStr2 := h2.Sum(nil)
	hex.EncodeToString(cipherStr2)
	md2 := hex.EncodeToString(cipherStr2)
	return "rtmp://" + liveAddress + "/sllmzb/" + allid + "?auth_key=" + time + "-0-0-" + md2

}
