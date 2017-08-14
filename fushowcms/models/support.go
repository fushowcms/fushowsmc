package models

/*
 * Copyright (c) 2016—2017 www.fushow.cn, All rights reserved.
 */
import (
	"bytes"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
)

/*新增支持*/
func (smt *SupportManagement) SupportAdd() bool {
	rows, err := Engine.Insert(smt)
	if rows > 0 && err == nil {
		return true
	}
	return false
}

/*删除支持*/
func (smt *SupportManagement) SupportDel() bool {
	var isok bool = false
	if num, err1 := Engine.Where("id = ?", smt.Id).Delete(smt); num > 0 && err1 == nil {
		isok = true
	}
	return isok
}

/*更新支持*/
func (smt *SupportManagement) SupportUp() bool {
	row, err := Engine.Where("Id=?", smt.Id).Update(smt)
	if row > 0 || err == nil {
		return true
	}
	return false
}

/*获取支持信息*/
func (smt *SupportManagement) GetSupport() bool {
	if _, err := Engine.Get(smt); err == nil { //注意：Get（）查询时没有结果时err不为空
		return true
	}
	return false
}

/*获取支持信息列表*/
func (smt *SupportManagement) GetSupportList(page, rows int, inputid string) ([]MySupport, int64) {
	var (
		test    []MySupport
		total   int64
		sql     string
		results []map[string][]byte
	)
	num1 := rows
	num2 := (page - 1) * rows
	if inputid == "" {
		sql = `SELECT * FROM support_management a LEFT JOIN product b on  b.id = substr(a.sup_encoding ,2,2) order by  a.suppor_time DESC limit ?,?   `
		total, _ = Engine.Count(smt)
		results, _ = Engine.Query(sql, num2, num1)
	} else {
		sql = `SELECT * FROM support_management a LEFT JOIN product b on  b.id = substr(a.sup_encoding ,2,2) WHERE a.periods_id like ?  order by  a.suppor_time DESC limit ?,?   `
		total, _ = Engine.Where("periods_id like ?", "%"+inputid+"%").Count(smt)
		results, _ = Engine.Query(sql, "%"+inputid+"%", num2, num1)
	}
	for _, value := range results {
		var smt MySupport //临时储存量
		smt.Uid, _ = strconv.ParseInt(string(value["uid"]), 10, 64)
		smt.PeriodsId, _ = strconv.ParseInt(string(value["periods_id"]), 10, 64)
		smt.SupporTime = string(value["suppor_time"])
		smt.ProductName = string(value["product_name"])
		smt.Odds = string(value["odds"])
		smt.IsWin, _ = strconv.ParseInt(string(value["is_win"]), 10, 64)
		smt.SupporNumber, _ = strconv.ParseInt(string(value["suppor_number"]), 10, 64)
		supEncoding := string(value["sup_encoding"])
		id := Substr(supEncoding, 4, 2)
		//转化int64
		b := bytes.Buffer{}                      //state拼接
		state, _ := strconv.ParseInt(id, 10, 64) //01--->1
		nowstate := strconv.FormatInt(state, 10) //to string  5-->"5"
		b.WriteString("state")                   //state小写   string(value["state"])
		b.WriteString(nowstate)                  //state1   string(value["state1"])
		ss := b.String()                         //state3
		smt.SupportState = string(value[ss])     //取用户选中值
		test = append(test, smt)                 //添加至数组
	}
	return test, total
}

/*获取期数支持信息列表*/
func (smt *SupportManagement) GetPeriodSupportList(page, rows int) []SupportManagement {
	var list []SupportManagement
	Engine.Where("periods_id = ?", smt.PeriodsId).Find(&list)
	return list

}

/*获取期数支持信息列表*/
func (smt *SupportManagement) GetPeriodSupportUidList(page, rows int) ([]MySupport, int64) {
	var (
		test    []MySupport
		sql     string
		results []map[string][]byte
	)
	num1 := rows
	num2 := (page - 1) * rows
	if num1 == 0 && num2 == 0 {
		sql = `SELECT * FROM support_management a LEFT JOIN product b on  b.id = substr(a.sup_encoding ,2,2) where a.uid = ? order by  a.suppor_time `
		results, _ = Engine.Query(sql, smt.Uid)
	} else {
		sql = `SELECT * FROM support_management a LEFT JOIN product b on  b.id = substr(a.sup_encoding ,2,2) where a.uid = ? order by  a.suppor_time DESC limit ?,?   `
		results, _ = Engine.Query(sql, smt.Uid, num2, num1)
	}
	total, _ := Engine.Where("uid =?", smt.Uid).Count(smt)
	for _, value := range results {
		var smt MySupport //临时储存量
		smt.PeriodsId, _ = strconv.ParseInt(string(value["periods_id"]), 10, 64)
		smt.SupporTime = string(value["suppor_time"])
		smt.ProductName = string(value["product_name"])
		smt.Odds = string(value["odds"])
		smt.IsWin, _ = strconv.ParseInt(string(value["is_win"]), 10, 64)
		smt.SupporNumber, _ = strconv.ParseInt(string(value["suppor_number"]), 10, 64)
		supEncoding := string(value["sup_encoding"])
		id := Substr(supEncoding, 4, 2)
		//转化int64
		b := bytes.Buffer{}                      //state拼接
		state, _ := strconv.ParseInt(id, 10, 64) //01--->1
		nowstate := strconv.FormatInt(state, 10) //to string  5-->"5"
		b.WriteString("state")                   //state小写   string(value["state"])
		b.WriteString(nowstate)                  //state1   string(value["state1"])
		ss := b.String()                         //state3
		smt.SupportState = string(value[ss])     //取用户选中值
		test = append(test, smt)                 //添加至数组
	}
	return test, total
}

func Substr(str string, start, length int) string {
	rs := []rune(str)
	rl := len(rs)
	end := 0
	if start < 0 {
		start = rl - 1 + start
	}
	end = start + length
	if start > end {
		start, end = end, start
	}
	if start < 0 {
		start = 0
	}
	if start > rl {
		start = rl
	}
	if end < 0 {
		end = 0
	}
	if end > rl {
		end = rl
	}
	return string(rs[start:end])
}

/*获取期数赢得所有信息数量*/
func (smt *SupportManagement) GetSupportWinnerNum() int {
	total, _ := Engine.Table("support_management").Alias("s").
		Join("INNER", []string{"periods", "p"}, "s.periods_id = p.periods_id").
		Join("INNER", []string{"uid_info", "u"}, "s.uid = u.id").
		Where("s.periods_id =? and p.sup_encoding like concat(\"%\",s.sup_encoding,\"%\")", smt.PeriodsId).Count(smt)
	return int(total)
}

/*获取期数支持信息列表*/
//  num显示个数    page页码
func (smt *SupportManagement) GetWinnerInfoList(page, num int) ([]Test, error) {
	var list []Test
	err := Engine.Table("support_management").Alias("s").
		Join("INNER", []string{"periods", "p"}, "s.periods_id = p.periods_id").
		Join("INNER", []string{"uid_info", "u"}, "s.uid = u.id").
		Where("s.periods_id =? and p.sup_encoding like concat(\"%\",s.sup_encoding,\"%\")", smt.PeriodsId).Limit(num, page*num).Find(&list)
	return list, err
}

//获取期数产品信息过程_cnxulin
func (sm *SupportManagement) GetSupCountByEncoding() Number {
	var list Number
	total, _ := Engine.Where("periods_id = ? and sup_encoding=?", sm.PeriodsId, sm.SupEncoding).Sum(sm, "suppor_number") //获取总数
	row, _ := Engine.Where("periods_id=? and sup_encoding=?", sm.PeriodsId, sm.SupEncoding).Count(sm)                    //获取行数
	list.SupportNumber = strconv.Itoa(int(total))
	list.RowNumber = row
	list.SupEncoding = sm.SupEncoding
	return list
}

//获取期数产品信息过程2_cnxulin
func (sm *SupportManagement) GetSupporByEncoding() ([]SupportManagement, Number) {
	var (
		list  Number
		lists []SupportManagement
	)
	Engine.Table("support_management").Alias("s").
		Where("s.periods_id =? and s.sup_encoding=?", sm.PeriodsId, sm.SupEncoding).Find(&lists)
	total, _ := Engine.Where("periods_id = ? and sup_encoding=?", sm.PeriodsId, sm.SupEncoding).Sum(sm, "suppor_number") //获取总数
	list.SupportNumber = strconv.Itoa(int(total))
	return lists, list
}

//根据我的id，期id，投注编码查询我的投注条数_cnxulin
func (sm *SupportManagement) GetMySupCount() int64 {
	var smt SupportManagement
	total, err := Engine.Where("uid=? and periods_id=? and sup_encoding=?", sm.Uid, sm.PeriodsId, sm.SupEncoding).Count(smt) //获取我的某个产品单个选项的投注次数
	if err != nil {
		return -1
	}
	return total
}
