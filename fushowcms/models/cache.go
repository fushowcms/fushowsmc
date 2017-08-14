//作成者：曹，作成日：2016/09/05
package models

/*
 * Copyright (c) 2016—2017 www.fushow.cn, All rights reserved.
 */
import (
	"fmt"
	"fushowcms/comm"
	"strconv"

	"github.com/garyburd/redigo/redis"
)

// 重写生成连接池方法
func newPoolUser() *redis.Pool {
	return &redis.Pool{
		MaxIdle:   80,
		MaxActive: 12000,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", comm.GetConfig("REDIS", "dial"), redis.DialPassword(comm.GetConfig("REDIS", "pass")))
			if err != nil {
				panic(err.Error())
			}
			return c, err
		},
	}
}

// 生成连接池
var pool = newPoolUser()

type AnchorGift struct {
	Time      int64
	Number    int64 //主播获得
	AllNumber int64 //主播热度值
}

//定期刷新主播礼物
//作成者：曹，作成日：2016/09/05
func AnchorGiftUp(allnum, time, num, id int64) bool {
	rs := pool.Get()
	defer rs.Close()
	var agu AnchorGift
	var ui UidInfo
	key := "AnchorGiftUp:" + strconv.FormatInt(id, 10)
	v, err := redis.Values(rs.Do("HGETALL", key))
	if err != nil {
		return false
	}
	if err := redis.ScanStruct(v, &agu); err != nil {
		return false
	}

	agu.Time++
	agu.Number = agu.Number + num
	agu.AllNumber = agu.AllNumber + allnum
	if agu.Time >= time {
		agu.Time = 0
		ui.Id = id
		if !ui.GetUserInfo() {
			return false
		}
		//主播石榴币更 新
		ui.PomegranateNum = ui.PomegranateNum + agu.Number
		//礼物总价值  不清零
		ui.GiftNum = ui.GiftNum + agu.AllNumber
		//主播更新
		if !ui.UserAnchorBalance() { //更新主播数据库的余额
			return false
		}
		agu.Number = 0
		agu.AllNumber = 0
	}
	fmt.Println("keyz:", key)
	if _, err := rs.Do("HMSET", redis.Args{}.Add(key).AddFlat(&agu)...); err != nil {
		panic(err)
	}
	return true

}
func SysGiftUp(time, num int64) bool {
	rs := pool.Get()
	defer rs.Close()
	var agu AnchorGift
	key := "SysGiftUp:"
	v, err := redis.Values(rs.Do("HGETALL", key))
	if err != nil {
		return false
	}
	if err := redis.ScanStruct(v, &agu); err != nil {
		return false
	}
	fmt.Println("key2:", key)
	agu.Time++
	agu.Number = agu.Number + num
	if agu.Time >= time {
		var fund, fundTwo Fund
		_, err := fund.GetFundDesc()
		if err != nil {
			return false
		}
		fundTwo.StorageFund = fund.StorageFund + float64(agu.Number)     //储备金  zijinmoney
		fundTwo.CurrencyMoney = fund.CurrencyMoney - float64(agu.Number) //流通金
		fundTwo.OfferAmount = fundTwo.StorageFund + fundTwo.CurrencyMoney
		_, err = fundTwo.FundAdd()
		if err != nil {
			return false
		}
		agu.Time = 0
		agu.Number = 0
	}
	if _, err := rs.Do("HMSET", redis.Args{}.Add(key).AddFlat(&agu)...); err != nil {
		panic(err)
	}
	return true

}
