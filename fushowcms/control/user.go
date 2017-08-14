package control

/*
 * Copyright (c) 2016—2017 www.fushow.cn, All rights reserved.
 */
import (
	"fmt"
	"fushowcms/comm"
	m "fushowcms/models"
	"io"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"

	"time"

	"github.com/gin-gonic/gin"

	alidayu "github.com/alidayu-master"
	"github.com/dchest/captcha"
)

//后台用户登录,返回格式已修改
func LoginRootMyAdmin(c *gin.Context) {
	var (
		uk  m.UserKey //判断用户是否存在
		ui  m.UidInfoJson
		uil m.UidLoginInfo
	)
	keycode := c.PostForm("keycode")
	capt := c.PostForm("captcha")
	fmt.Println("capt", capt)
	fmt.Println("keycode", keycode)
	if keycode == "" {
		comm.ResponseError(c, 4066) //请输入图形验证码
		return
	}
	if !captcha.VerifyString(capt, keycode) {
		comm.ResponseError(c, 4065) //图形验证码输入错误
		return
	}
	uk.UserName = c.PostForm("username")
	ui.UserName = c.PostForm("username")
	password := c.PostForm("password")
	password1 := comm.SetAesValue(password, "fushow.cms")

	//获取数据库密码
	if !ui.Login() {
		comm.ResponseError(c, 4028)
		return
	}

	//密码长度相等
	if !strings.EqualFold(ui.PassWord, password1) {
		comm.ResponseError(c, 4011)
		return
	}

	if ui.Type <= 2 {
		comm.ResponseError(c, 4029)
		return
	}
	//登录成功时生成cookie、token,并复制到session的结构体,存入redis
	if SetCookieSession(ui, c) != "" {
		comm.ResponseError(c, 4030)
		return
	} //ssion set
	ui.UID = setNormalAesValue(strconv.FormatInt(ui.Id, 10), "fushow.cms")
	ui.PassWord = ""

	//生成登陆记录
	uil.Uid = ui.Id
	uil.LoginWay = 2
	uil.Ip = c.ClientIP()
	if !uil.UidLoginInfoAdd() {
		fmt.Println("登陆记录插入失败")
	}
	comm.Response(c, ui)
}

/*用户登录
*@muhailong已修改后台返回格式
*20161130
 */
func UnLogin(c *gin.Context) {
	var ui m.UidInfoJson
	ui.Id, _ = strconv.ParseInt(c.PostForm("UID"), 10, 64)
	if ui.Id == 0 {
		comm.ResponseError(c, 3175)
		return
	}
	//登录成功时生成cookie、token,并复制到session的结构体,存入redis
	if DelCookieSession(ui, c) != "" {
		comm.ResponseError(c, 4018)
		return
	} //ssion set
	ui.UID = comm.SetAesValue(strconv.FormatInt(ui.Id, 10), "fushow.cms")
	comm.Response(c, ui)
}

//删除用户,返回已修改
func UserDel(c *gin.Context) {
	var (
		ui m.UidInfo
		uk m.UserKey
	)

	//获取客户端参数id
	ui.Id, _ = strconv.ParseInt(c.PostForm("id"), 10, 64)
	uk.Id = ui.Id

	//删除UidInfo表
	if !ui.UserInfoDel() {
		comm.ResponseError(c, 4031)
		return
	}

	//删除UserKey表
	if !uk.UserKeyDel() {
		comm.ResponseError(c, 4031)
		return
	}
	comm.Response(c, nil)
}

//用户信息修改,后台返回数据格式已修改
func UserUp(c *gin.Context) {
	var ui m.UidInfo

	//获取客户端参数
	ui.Id, _ = strconv.ParseInt(c.PostForm("UID"), 10, 64)
	if ui.Id == 0 {
		comm.ResponseError(c, 3175)
		return
	}
	//后台添加
	nickname := c.PostForm("NickName")
	phone := c.PostForm("Phone")
	types, _ := strconv.ParseUint(c.PostForm("Type"), 0, 64)
	balance, _ := strconv.ParseInt(c.PostForm("Balance"), 10, 64)
	password := c.PostForm("PassWord")
	password1 := comm.SetAesValue(password, "fushow.cms")
	remark := c.PostForm("Remark")

	if ui.Type != 255 && ui.Type != 3 && ui.Type != 2 && ui.Type != 1 {
		ui.Type = 0
	}

	if !ui.GetUserInfo() {
		c.JSON(200, gin.H{"state": "fail"})
		return
	}
	if nickname != "" {
		ui.NickName = nickname
	}

	if phone != "" {
		ui.Phone = phone
	}

	if c.PostForm("Type") != "" {
		ui.Type = types
	}

	if c.PostForm("Balance") != "" {
		ui.Balance = float64(balance)
	}
	if c.PostForm("PassWord") != "" {
		ui.PassWord = password1
	}

	if remark != "" {
		ui.Remark = remark
	}

	if nickname == "" && phone == "" && c.PostForm("Type") == "" && c.PostForm("Balance") == "" && remark == "" {
		file, header, err := c.Request.FormFile("upload")
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(header.Filename)
		if header.Filename != "" {

			nowtime := time.Now().Format("2006-01-02 15:04:05")
			stringtime := strings.Replace(nowtime, " ", "", -1)
			stringtime = strings.Replace(stringtime, ":", "", -1)
			stringtime = strings.Replace(stringtime, "-", "", -1)

			out, err := os.Create("./static/upload/" + stringtime + "image" + c.PostForm("Id") + ".jpg")
			if err != nil {
				log.Fatal(err)
			}
			defer out.Close()
			_, err = io.Copy(out, file)

			if err != nil {
				log.Fatal(err)
			}

			ui.Favicon = "/static/upload/" + stringtime + "image" + c.PostForm("Id") + ".jpg"

		}
	}
	//修改用户信息
	if !ui.UserInfoUp() { //失败返回
		comm.ResponseError(c, 4023)
		return
	} else { //成功返回
		comm.Response(c, ui.Favicon)
	}
}

//用户密码修改,感觉无用,先保留,不确定,未修改
func PassUp(c *gin.Context) {
	var (
		ui m.UidInfo
	)

	//获取客户端参数
	username := c.PostForm("username")       //用户名
	password := c.PostForm("password")       //密码
	passwordNew := c.PostForm("newpassword") //新密码

	//获取用户名密码
	has := ui.GetUserPwd(username)

	if !has {
		c.JSON(200, gin.H{"state": "user not exist"})
		return
	}
	if password != "" { //找回密码通过验证手机 不需验证密码

		//密码长度相等
		if len(password) == len(ui.PassWord) {
			for i := 0; i < len(password); i++ {
				if password[i] != ui.PassWord[i] {
					c.JSON(200, gin.H{"state": "password error"})
					return
				}
			}
		}
	}
	ui.PassWord = passwordNew

	//赋予新密码
	if !ui.PassUp() {
		c.JSON(200, gin.H{"state": "fail"}) //失败返回
		return
	}
	c.JSON(200, gin.H{"state": "success"}) //成功返回

}

//用户密码修改,后台返回格式已修改
func PcPassUp(c *gin.Context) {
	var (
		ui m.UidInfo
	)
	//获取客户端参数
	userId, _ := strconv.ParseInt(c.PostForm("UID"), 10, 64) //用户id
	if userId == 0 {
		comm.ResponseError(c, 3175)
		return
	}
	password := c.PostForm("password")                          //密码
	password1 := comm.SetAesValue(password, "fushow.cms")       //原密码加密
	passwordNew := c.PostForm("newpassword")                    //新密码
	passwordNew1 := comm.SetAesValue(passwordNew, "fushow.cms") //新密码加密
	ui.Id = userId
	ui.PassWord = password1
	//获取用户名密码
	has := ui.GetUserInfo()
	if !has {
		comm.ResponseError(c, 4011)
		return
	}
	if !strings.EqualFold(ui.PassWord, password1) {
		comm.ResponseError(c, 4033)
		return
	}
	ui.PassWord = passwordNew1
	//赋予新密码
	if !ui.PassUp() {
		comm.ResponseError(c, 4013)
		return
	}
	comm.Response(c, 4034)
}

/*
 *功能：手机用户忘记密码，验证后修改密码
 *参数：用户名、新密码（username、newpassword）
 *返回：flag (修改状态)
 *@徐林 20161012
 */
//感觉无用,先保留,不确定,未修改
func MobilePassUp(c *gin.Context) {
	var (
		ui m.UidInfo
	)

	//获取客户端参数
	username := c.PostForm("username")       //用户名
	passwordNew := c.PostForm("newpassword") //新密码
	ui.UserName = username

	if !ui.GetUserInfo() {
		c.JSON(200, gin.H{"flag": false, "data": "用户不存在"})
		return
	}

	ui.PassWord = passwordNew
	//赋予新密码
	if !ui.PassUp() {
		c.JSON(200, gin.H{"flag": "修改失败"}) //失败返回
		return
	}
	c.JSON(200, gin.H{"flag": "修改成功"}) //成功返回

}

//获取用户信息,返回格式已修改
func GetUser(c *gin.Context) {
	var ui m.UidInfoJson

	//获取客户端参数
	id, _ := strconv.ParseInt(c.PostForm("UID"), 10, 64)
	if id == 0 {
		comm.ResponseError(c, 4010)
		return
	}
	ui.Id = id
	// 根据id 获取用户信息
	if !ui.GetUserInfo() {
		comm.ResponseError(c, 4020)
		return
	} else {
		ui.PassWord = ""
		ui.BankCard = 0
		ui.BankName = ""
		ui.BankDeposit = ""
		ui.UID = comm.SetAesValue(strconv.FormatInt(ui.Id, 10), "fushow.cms")
		comm.Response(c, ui)
	}
}

//获取用户信息，感觉无用，没改
func GetUserName(c *gin.Context) {
	var ui m.UidInfo

	//获取客户端参数
	ui.UserName = c.PostForm("username")
	if c.PostForm("username") == "" {
		c.JSON(200, gin.H{"flag": false, "message": ""})
		return
	}
	// 根据id 获取用户信息
	if !ui.GetUserInfo() {
		c.JSON(200, gin.H{"state": "fail"})
	} else {
		ui.PassWord = ""
		c.JSON(200, gin.H{"flag": true, "message": ui})
	}

}

//获取用户列表,感觉无用，没改
func GetUserList(c *gin.Context) {
	var ui m.UidInfo

	check := c.PostForm("inputid") // 昵称

	//获取页数-行数
	page, _ := strconv.Atoi(c.PostForm("page"))
	rows, _ := strconv.Atoi(c.PostForm("rows"))

	userlist, total := ui.GetUserList(page, rows, check)

	if total == 0 { //行数为0
		return
	}
	c.JSON(200, gin.H{"total": total, "rows": userlist})
}

//后台用户列表,感觉无用，没改
func GetRootUserList(c *gin.Context) {
	var uid m.UidInfo
	page, _ := strconv.Atoi(c.PostForm("page"))
	rows, _ := strconv.Atoi(c.PostForm("rows"))
	list, row := uid.GetRootUserList(page, rows)

	if row == 0 {
		return
	}
	c.JSON(200, gin.H{"total": row, "rows": list})
}

/*
*功能:验证密码是否匹配
*@muhailong
*时间:20161103 后台返回已修改
 */
func PassMate(c *gin.Context) {
	var ui m.UidInfo
	ui.Id, _ = strconv.ParseInt(c.PostForm("UID"), 10, 64)
	if ui.Id == 0 {
		comm.ResponseError(c, 3175)
		return
	}
	pass := c.PostForm("pass")
	pass1 := comm.SetAesValue(pass, "fushow.cms")
	has := ui.PassMate()
	if has == false {
		comm.ResponseError(c, 4010)
		return
	}
	if pass1 != ui.PassWord {
		comm.ResponseError(c, 4035)
		return
	}
	comm.Response(c, nil)
}

/*
*功能:后台管理按用户等级正序排序
*@muhailong
*时间:20161101
 */
//默认查询无法修改
func UserOrderByLevelADSC(c *gin.Context) {
	var uid m.UidInfo
	page, _ := strconv.Atoi(c.PostForm("page"))
	rows, _ := strconv.Atoi(c.PostForm("rows"))
	sort := c.PostForm("sort")
	order := c.PostForm("order")
	inputid := c.PostForm("inputid")
	list, row := uid.UserOrderByLevelADSC(page, rows, sort, order, inputid)
	if row == 0 {
		return
	}
	c.JSON(200, gin.H{"total": row, "rows": list})
}

//等级排序，后台返回已修改
func UserOrderBySearch(c *gin.Context) {
	page, _ := strconv.Atoi(c.PostForm("page"))
	rows, _ := strconv.Atoi(c.PostForm("rows"))
	var list []m.UidInfo
	var row int64
	var uid m.UidInfo
	utype := c.PostForm("Type") //4---->权限
	leve := c.PostForm("Leve")
	if leve == "-1" {
		leve = ""
	}
	if utype == "-1" {
		utype = ""
	}
	if leve == "" && utype == "" {
		list, row = uid.FindUserOrderByLevelADSC(page, rows, "", "desc", "")
	}
	if leve == "1" {
		leve = "1000"
	} else if leve == "2" {
		leve = "10000"
	} else if leve == "3" {
		leve = "100000"
	} else if leve == "4" {
		leve = "1000000"
	} else if leve == "5" {
		leve = "10000000"
	} else if leve == "6" {
		leve = "20000000"
	} else if leve == "7" {
		leve = "30000000"
	}

	var arr [1]int // 声明了一个int类型的数组
	arr[0] = 99999 // 数组下标是从0开始的

	if !strings.EqualFold(utype, "") && strings.EqualFold(leve, "") {
		uid.Type, _ = strconv.ParseUint(utype, 10, 64)
		list, row = uid.UserOrderBySearch1(page, rows)
	}
	if strings.EqualFold(utype, "") && !strings.EqualFold(leve, "") {
		uid.Integral, _ = strconv.ParseInt(leve, 10, 64)
		list, row = uid.UserOrderBySearch2(page, rows)
	}
	if !strings.EqualFold(utype, "") && !strings.EqualFold(leve, "") {
		uid.Type, _ = strconv.ParseUint(utype, 10, 64)
		uid.Integral, _ = strconv.ParseInt(leve, 10, 64)
		list, row = uid.UserOrderBySearch3(page, rows)
	}

	if len(list) <= 0 {
		c.JSON(200, gin.H{"total": row, "rows": arr})
		return
	}

	c.JSON(200, gin.H{"total": row, "rows": list})
}

/*
*功能:后台主播信息
*@muhailong
*时间:20161014
 */
//后台管理默认调用，无法修改
func GetAnchorInfos(c *gin.Context) {
	var uid m.UidInfo
	page, _ := strconv.Atoi(c.PostForm("page"))
	rows, _ := strconv.Atoi(c.PostForm("rows"))
	inputid := c.PostForm("inputid")
	list, row := uid.GetAnchorInfos(page, rows, inputid)
	if len(list) == 0 {
		c.JSON(200, gin.H{"total": row, "rows": []int{}})
		return
	}
	c.JSON(200, gin.H{"total": row, "rows": list})
}

/*
*功能:后台主播结算
*@muhailong
*时间:20161014
 */
//后台返回已修改
func AnchorBalance(c *gin.Context) {
	var (
		ui  m.UidInfo
		usd m.UidInfoSettlementDetail
	)
	ui.Id, _ = strconv.ParseInt(c.PostForm("Suid"), 10, 64)
	fmt.Println("id11", ui.Id)
	usd.Uid, _ = strconv.ParseInt(c.PostForm("Suid"), 10, 64)
	usd.SettlementDetail.Id, _ = strconv.ParseInt(c.PostForm("Id"), 10, 64)
	usd.SettlementDetail.ApplyCashingNum, _ = strconv.ParseInt(c.PostForm("ApplyCashingNum"), 10, 64)
	if ui.Id == 0 {
		comm.ResponseError(c, 4017)
		return
	}
	if !ui.GetUserInfo() {
		comm.ResponseError(c, 4010)
		return
	}
	if !usd.BalanceUp() {
		comm.ResponseError(c, 4036)
		return
	}
	comm.Response(c, "结算成功")
}

/*
*功能:查询是否修改过昵称
*@muhailong
*时间:20161031
 */
//返回结果已修改
func CheckNickNameChange(c *gin.Context) {
	var ui m.UidInfo
	//获取客户端参数
	id, _ := strconv.ParseInt(c.PostForm("UID"), 10, 64)
	if id == 0 {
		comm.ResponseError(c, 3175)
		return
	}
	ui.Id = id
	// 根据id 获取用户信息
	if !ui.CheckNickName() {
		comm.ResponseError(c, 4010)
		return
	} else {
		comm.Response(c, ui)
	}
}

/*
*功能:修改昵称
*@muhailong
*时间:20161031
 */
//返回格式已修改
func ChangeNickname(c *gin.Context) {
	var ui m.UidInfo
	id, _ := strconv.ParseInt(c.PostForm("UID"), 10, 64) //用户id
	ui.Id = id
	nickname := c.PostForm("nickname") //新修改昵称
	nickflag := c.PostForm("nickflag") //昵称标识

	if !ui.GetUserInfo() {
		comm.ResponseError(c, 4010)
		return
	}

	if KeyWord(nickname) {
		comm.ResponseError(c, 4062)
		return
	}
	ui.NickName = nickname
	ui.NickFlag, _ = strconv.ParseBool(nickflag)

	//修改昵称
	if !ui.NickNameUp() {
		comm.ResponseError(c, 4032)
		return
	}
	comm.Response(c, nil)
}

/*
*功能:过滤昵称关键词
*@Mao Yanlei & Zhang Shuo
*日期:20170303
 */
func KeyWord(word string) bool {
	var wordMap = map[int]string{0: "赌博", 1: "军火", 2: "毒品", 3: "冰毒", 4: "习近平", 5: "江泽民", 6: "胡锦涛", 7: "温家宝", 8: "习仲勋", 9: "共产党", 10: "反共", 11: "反党", 12: "FUCK", 13: "真善忍", 14: "李洪志", 15: "肉棍", 16: "淫靡", 17: "迷昏药", 18: "窃听器", 19: "买卖枪支", 20: "麻醉药", 21: "短信群发器", 22: "摇头丸", 23: "黑社会", 24: "枪决女犯", 25: "出售假币", 26: "共产党", 27: "手枪", 28: "色情", 29: "手淫", 30: "强奸", 31: "干你", 32: "操你妈", 33: "你妈逼", 34: "cao你妈", 35: "淫荡", 36: "阴道", 37: "女优", 38: "性虐待", 39: "强奸犯", 40: "日你妈", 41: "操你", 42: "你真丑", 43: "丑b", 44: "丑鬼", 45: "性交", 46: "性生活", 47: "想日你", 48: "生孩子", 49: "骚b", 50: "骚货", 51: "穷b", 52: "淫秽", 53: "性器官", 54: "鸡巴", 55: "狗b", 56: "诈骗", 57: "摇头丸", 58: "K粉", 59: "嫖娼", 60: "叫床", 61: "叫春", 62: "发情", 63: "性病", 64: "梅毒", 65: "阴道炎", 66: "溃烂", 67: "乳房", 68: "小姐", 69: "情妇", 70: "卖淫", 71: "法轮功", 72: "法轮大法"}
	for _, val := range wordMap {
		if strings.Contains(word, val) {
			return true
		}
	}
	return false

}

/*
*功能:个人充值记录
*@muhailong
*日期:20161103
 */
//liuhan 修改,后台返回已修改
func UserPayRecord(c *gin.Context) {
	var rech m.RechargingRecords
	rech.Uid, _ = strconv.ParseInt(c.PostForm("UID"), 10, 64)
	if rech.Uid == 0 {
		comm.ResponseError(c, 3175)
		return
	}
	page, _ := strconv.Atoi(c.PostForm("page"))
	rows, _ := strconv.Atoi(c.PostForm("rows"))
	list, row := rech.UserPayRecord(page, rows)
	if row == 0 {
		comm.ResponseError(c, 4028)
		return
	}
	m := make(map[string]interface{})
	m["rows"] = list
	m["total"] = row
	comm.Response(c, m)
}

/*
*功能:主播申请结算
*@muhailong
*日期:20161105
**/
//后台返回已修改
func AnchorApplyCashing(c *gin.Context) {
	var sd m.SettlementDetail
	var ui m.UidInfo
	sd.Uid, _ = strconv.ParseInt(c.PostForm("UID"), 10, 64)
	if sd.Uid == 0 {
		comm.ResponseError(c, 3175)
		return
	}
	sd.ApplyCashingNum, _ = strconv.ParseInt(c.PostForm("applycashingnum"), 10, 64)
	sd.Cashing, _ = strconv.ParseFloat(c.PostForm("cashing"), 64)
	sd.CashingDate = c.PostForm("cashingdate")
	sd.IsApply = true
	ui.PomegranateNum, _ = strconv.ParseInt(c.PostForm("applycashingnum"), 10, 64)
	if !sd.AnchorApplyCashing() {
		comm.ResponseError(c, 4037)
		return
	}
	comm.Response(c, nil)
}

/*
*功能:主播银行卡信息绑定
*@muhailong
*日期:20161105
**/
//后台返回已修改
func AnchorBindingBank(c *gin.Context) {
	var ui m.UidInfo
	ui.Id, _ = strconv.ParseInt(c.PostForm("UID"), 10, 64)
	if ui.Id == 0 {
		comm.ResponseError(c, 3175)
		return
	}
	ui.BankName = c.PostForm("bankname")
	ui.BankCard, _ = strconv.ParseInt(c.PostForm("bankcard"), 10, 64)
	ui.BankDeposit = c.PostForm("bankdeposit")
	if !ui.AnchorBindingBank() {
		comm.ResponseError(c, 4039)
		return
	}
	comm.Response(c, nil)
}

/*
*功能:判断主播是否绑定银行卡
*@muhailong
*日期:20161107
**/
//返回已修改
func IsBindingBank(c *gin.Context) {
	var ui m.UidInfo
	ui.Id, _ = strconv.ParseInt(c.PostForm("UID"), 10, 64)
	if ui.Id == 0 {
		comm.ResponseError(c, 3175)
		return
	}
	list, row := ui.IsBindingBank()
	if row == 0 {
		comm.ResponseError(c, 4028)
		return
	}
	m := make(map[string]interface{})
	m["rows"] = list
	m["total"] = row
	comm.Response(c, m)
}

/*
*功能:判断申请结算数量是否足够
*@muhailong
*日期:20161107
**/
//返回已修改
func IsEnough(c *gin.Context) {
	var ui m.UidInfo
	ui.Id, _ = strconv.ParseInt(c.PostForm("UID"), 10, 64)
	if ui.Id == 0 {
		comm.ResponseError(c, 3175)
		return
	}
	pome_num, _ := strconv.ParseInt(c.PostForm("applycashingnum"), 10, 64)
	if !ui.GetUserInfo() {
		comm.ResponseError(c, 4010)
		return
	}
	flag := ui.IsEnough(pome_num)
	if !flag {
		comm.ResponseError(c, 4061)
		return
	}
	comm.Response(c, nil)
}

/*
*功能:主播查询结算明细
*@muhailong
*日期:20161105
**/
//后台返回已修改
func SettlementDetails(c *gin.Context) {
	var sd m.SettlementDetail
	sd.Uid, _ = strconv.ParseInt(c.PostForm("UID"), 10, 64)
	if sd.Uid == 0 {
		comm.ResponseError(c, 3175)
		return
	}
	page, _ := strconv.Atoi(c.PostForm("page"))
	rows, _ := strconv.Atoi(c.PostForm("rows"))
	list, row := sd.SettlementDetails(page, rows)
	if row == 0 {
		comm.ResponseError(c, 4028)
		return
	}
	m := make(map[string]interface{})
	m["rows"] = list
	m["total"] = row
	comm.Response(c, m)
}

/*
*功能:判断主播本月是否已经申请过结算
*@muhailong
*日期:20161124
**/
func IsMonthCashing(c *gin.Context) {
	var sd m.SettlementDetail
	sd.Uid, _ = strconv.ParseInt(c.PostForm("UID"), 10, 64)
	data, _ := sd.IsMonthCashing()
	if len(data) != 0 {
		comm.Response(c, data)
		return
	}
	comm.Response(c, nil)
}

//礼物赠送
//参数  num2 系统回收石榴籽
//     uid  用户ID    anchor 主播ID   nowgift 礼物ID   赠送数量
//后台返回已修改
func GiveGiftNumAdd(c *gin.Context) {

	var (
		gift m.Gift
		user m.UidInfo
	)
	gift.Id, _ = strconv.ParseInt(c.PostForm("GiftId"), 10, 64)
	uid, _ := strconv.ParseInt(c.PostForm("UID"), 10, 64)
	if uid == 0 {
		comm.ResponseError(c, 3175)
		return
	}
	anchor, _ := strconv.ParseInt(c.PostForm("AnchorId"), 10, 64)
	number, _ := strconv.ParseInt(c.PostForm("Number"), 10, 64)

	flag, nowgift := gift.GetGiftCon()
	if !flag {
		comm.ResponseError(c, 4041)
		return
	}

	if number < 1 {
		comm.ResponseError(c, 2047)
		return
	}

	//系统回收石榴籽=购买时所花石榴籽  - 可兑换石榴籽
	num2 := nowgift.BuyNumber - nowgift.ToNumber

	_, errMsg := user.GiftFund(nowgift.ToNumber, num2, uid, anchor, nowgift.Id, number)
	if errMsg == "赠送人不存在" {
		comm.ResponseError(c, 2049)
		return
	}
	if errMsg == "用户余额不足" {
		comm.ResponseError(c, 2050)
		return
	}
	if errMsg == "用户扣除石榴籽失败" {
		comm.ResponseError(c, 2051)
		return
	}

	comm.Response(c, nil)
}

//赠送石榴籽,后台返回已修改
func GiveNumber(c *gin.Context) {
	var (
		ui   m.UidInfo
		toui m.UidInfo
		gs   m.GiveSlz
	)
	mid := c.PostForm("UID")
	mnum := c.PostForm("Number")
	mtoId := c.PostForm("ToId")
	if mid == mtoId {
		comm.ResponseError(c, 4042)
		return
	}
	if mid == "" {
		comm.ResponseError(c, 4043)
		return
	} else if mnum == "" {
		comm.ResponseError(c, 4044)
		return
	} else if mtoId == "" {
		comm.ResponseError(c, 4045)
		return
	}
	id, _ := strconv.ParseInt(mid, 10, 64)
	toId, _ := strconv.ParseInt(mtoId, 10, 64)
	num, _ := strconv.ParseInt(mnum, 10, 64)
	if num < 5000 {
		comm.ResponseError(c, 4064)
		return
	}
	ui.Id = id
	toui.Id = toId
	gs.BenefactorId, _ = strconv.ParseInt(mid, 10, 64)
	gs.RecipientId, _ = strconv.ParseInt(mtoId, 10, 64)
	gs.Num, _ = strconv.ParseInt(mnum, 10, 64)
	if !ui.GetUserInfo() {
		comm.ResponseError(c, 4046)
		return
	}
	if !toui.GetUserInfo() {
		comm.ResponseError(c, 4047)
		return
	}

	if ui.Balance == float64(num) {
		ui.Balance = 0
	} else {
		ui.Balance = ui.Balance - float64(num)
	}
	if ui.Balance < 0 {
		comm.ResponseError(c, 4048)
		return
	}

	if ui.UserCoastBalance() {
		toui.Balance = toui.Balance + float64(num)
		if toui.UserCoastBalance() {
			// addby liuhan
			if !gs.AddGiveSlz() {
				comm.ResponseError(c, 4049)
				return
			} else {
				number := strconv.FormatInt(num, 10)
				comm.Response(c, gin.H{"state": true, "message": "您成功赠送给" + toui.NickName + "用户," + number + "个石榴籽"})
				return
			}
		}
	}
	comm.ResponseError(c, 4050)
	return
}

//后台添加用户,后台管理默认调用，无法修改
func RootAddUser(c *gin.Context) {
	var (
		uk m.UserKey //判断用户是否存在
		ui m.UidInfo
	)

	//获取客户端参数
	uk.UserName = c.PostForm("UserName")
	ui.UserName = c.PostForm("UserName")
	ui.RegTime = time.Unix(1389058332, 0).Format("2006-01-02 15:04:05")
	ui.PassWord = comm.GetConfig("Root", "pass")

	//后台添加
	ui.NickName = c.PostForm("NickName")
	ui.Phone = c.PostForm("Phone")
	ui.Type, _ = strconv.ParseUint(c.PostForm("Type"), 0, 64)
	ui.Balance, _ = strconv.ParseFloat(c.PostForm("Balance"), 10)
	ui.Level = 0
	ui.Remark = c.PostForm("Remark")

	if ui.Type != 255 && ui.Type != 3 && ui.Type != 2 && ui.Type != 1 {
		ui.Type = 99
	}

	if uk.GetUserKey() { //用户已存在
		c.JSON(200, gin.H{"state": "exist"})
		return
	}

	if !uk.UserKeyAdd() { //增加用户失败
		c.JSON(200, gin.H{"state": "failuk"})
		return
	}

	ui.Id = uk.Id
	if !ui.UserInfoAdd() { //增加用户失败
		uk.DelUserKey() // 删除UserKey表新增加数据
		c.JSON(200, gin.H{"state": "failui"})
		return
	}

	ui.PassWord = ""
	c.JSON(200, ui) //成功返回

}

//后台用户列表--->按余额,路由中没查到，无用，暂留
func GetInfoPower(c *gin.Context) {
	var uid m.UidInfo
	page, _ := strconv.Atoi(c.PostForm("page"))
	rows, _ := strconv.Atoi(c.PostForm("rows"))
	list, row := uid.GetInfoPower(page, rows)
	if row == 0 {
		return
	}
	c.JSON(200, gin.H{"total": row, "rows": list})

}

//WEB成为主播,后台返回已修改
func UserIdNumber(c *gin.Context) {

	var (
		user m.UidInfo
		al   m.Applicant
	)
	user.Id, _ = strconv.ParseInt(c.PostForm("UID"), 10, 64)
	IdNumber := c.PostForm("IdNumber") //身份证号码
	fielname := c.PostForm("uploadFile")
	phone := c.PostForm("Phone")
	name := c.PostForm("Name")
	if c.PostForm("UID") == "" {
		comm.ResponseError(c, 4017)
		return
	}
	if !user.GetUserInfo() {
		comm.ResponseError(c, 4010)
		return
	}
	if user.Type != 99 {
		comm.ResponseError(c, 4051)
		return
	}
	user.IdentityPic = "/static/upload/" + fielname
	user.RealName = name
	user.Phone = phone
	user.IdNumber = IdNumber
	if !user.UserInfoUp() {
		comm.ResponseError(c, 4052)
		return
	}
	al.UserId = user.Id

	if !al.GetApplyById() { //判断是否已经申请
		//不存在时
		al.State = 0
		if !al.ApplyAdd() {
			comm.ResponseError(c, 4053)
			return
		}
	} else {
		//存在时
		al.State = 0
		if !al.ApplyStateUp() {
			comm.ResponseError(c, 4054)
			return
		}
	}
	comm.Response(c, nil)
}

//注册账号  -->通过手机号
//参数  mobile  注册手机号
//错误信息   message  1.参数错误 2.该手机号已注册 3.短信服务请求异常 4.数据库更新失败  5.发送失败
//2016-10-16   txl
//后台返回已修改
func ByPhoneRegPhone(c *gin.Context) {
	keycode := c.PostForm("keycode")
	capt := c.PostForm("captcha")
	if keycode == "" {
		comm.ResponseError(c, 4066) //请输入图形验证码
		return
	}
	if !captcha.VerifyString(capt, keycode) {
		comm.ResponseError(c, 4065) //图形验证码输入错误
		return
	}
	var (
		uidinfo m.UidInfo
		vcode   m.VerificationCode
	)
	mobile := c.PostForm("mobile")
	if c.PostForm("mobile") == "" {
		comm.ResponseError(c, 4001)
		return
	}
	uidinfo.Phone = mobile
	if uidinfo.GetUserPhone() {
		comm.ResponseError(c, 4002)
		return
	}
	//随机生成6位验证码
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	code := Substr(strconv.Itoa(r.Int()), 0, 6)
	//	向短信服务器发送请求
	alidayu.AppKey = "23477883"
	alidayu.AppSecret = "10a97d80159cfd05439d7f912de24e8a"
	success, resp := alidayu.SendSMS(mobile, "石榴联盟直播", "SMS_25635195", `{"code":"`+code+`","time":"5"}`)
	fmt.Println("Success:", success)
	fmt.Println(resp)
	if !success {
		comm.ResponseError(c, 4004)
		return
	}

	//添加短信验证码表
	vcode.Phone = mobile
	if vcode.GetVerificationCode() {
		vcode.Code = code
		if !vcode.VerificationCodeUpdate() {
			comm.ResponseError(c, 4003)
			return
		}
	}
	vcode.Code = code
	if !vcode.VerificationCodeAdd() {
		comm.ResponseError(c, 4004)
		return
	}
	comm.Response(c, nil)
}

//验证验证码是否正确
//参数 mobile-->手机号   code-->验证码
//错误信息 message 1.参数错误 2.不存在该手机号 3.验证码错误
//后台返回已修改
func IsVerification(c *gin.Context) {
	var vcode m.VerificationCode
	mobile := c.PostForm("mobile")
	code := c.PostForm("code")

	if mobile == "" || code == "" {
		comm.ResponseError(c, 4005)
		return
	}
	vcode.Phone = mobile
	if !vcode.GetVerificationCode() {
		comm.ResponseError(c, 4006)
		return
	}
	if !strings.EqualFold(vcode.Code, code) {
		comm.ResponseError(c, 4007)
		return
	}
	comm.Response(c, nil)
}

//截取字符串
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

type ThirdPartyLanding struct {
	Appid string
	Token string
	Type  string
	Code  string
	//	url      string
}

//向手机发送验证码 --> 通过手机号
//参数 mobile 注册手机号
//错误信息   message  1.参数错误 2.该手机号已注册  3.短信服务请求异常  4.数据库更新失败  5.发送失败

func getVercodeFromServer(keycode string, capt string, mobile string, functionType string) int {
	if keycode == "" {
		return 4063
	}
	if !captcha.VerifyString(capt, keycode) {
		return 4062
	}
	var (
		uidinfo m.UidInfo
		vcode   m.VerificationCode
	)

	if mobile == "" {
		return 4001
	}
	if functionType == "" {
		return 2000
	} else if functionType == "regval" || functionType == "bindphone" {
		uidinfo.Phone = mobile
		if uidinfo.GetUserPhone() {
			return 4002
		}
	} else if functionType == "loseregval" {
		if !uidinfo.GetUserPhone() {
			return 4010
		}
	} else {
		return 2000
	}

	//随机生成6位验证码
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	code := Substr(strconv.Itoa(r.Int()), 0, 6)

	//	向短信服务器发送请求
	alidayu.AppKey = "23477883"
	alidayu.AppSecret = "10a97d80159cfd05439d7f912de24e8a"
	success, resp := alidayu.SendSMS(mobile, "石榴联盟直播", "SMS_25635195", `{"code":"`+code+`","time":"5"}`)
	fmt.Println("Success:", success)
	fmt.Println(resp)
	if !success {
		return 4016
	}
	//添加短信验证码表
	vcode.Phone = mobile
	if vcode.GetVerificationCode() {
		vcode.Code = code
		if !vcode.VerificationCodeUpdate() {
			return 4003
		}
	}
	vcode.Code = code
	if !vcode.VerificationCodeAdd() {
		return 4004
	}
	return 0
}

//注册账号  -->通过手机号(新)
//参数  mobile  注册手机号
//错误信息   message  1.参数错误 2.该手机号已注册 3.短信服务请求异常 4.数据库更新失败  5.发送失败
//2016-11-10  txl
func RegVal(c *gin.Context) {
	keycode := c.PostForm("keycode")
	captcha := c.PostForm("captcha")
	mobile := c.PostForm("mobile")
	functionType := "regval"
	errorCode := getVercodeFromServer(keycode, captcha, mobile, functionType)
	if errorCode != 0 {
		comm.ResponseError(c, errorCode)
	} else {
		comm.Response(c, nil)
	}
}

//验证验证码是否正确(新)
//参数 mobile-->手机号   code-->验证码
//错误信息 message 1.参数错误 2.不存在该手机号 3.验证码错误
func IsVerCode(c *gin.Context) {
	var (
		uidinfo m.UidInfo
		vcode   m.VerificationCode
	)
	mobile := c.PostForm("mobile")
	code := c.PostForm("code")

	if mobile == "" || code == "" {
		comm.ResponseError(c, 4005)
		return
	}

	uidinfo.Phone = mobile
	if uidinfo.GetUserPhone() {
		comm.ResponseError(c, 4002) //手机号已注册
		return
	}
	vcode.Phone = mobile
	if !vcode.GetVerificationCode() {
		comm.ResponseError(c, 4006)
		return
	}
	if !strings.EqualFold(vcode.Code, code) {
		comm.ResponseError(c, 4007)
		return
	}
	comm.Response(c, nil)
}

//手机注册
func Reg(c *gin.Context) {
	var (
		ui               m.UidInfo
		uidinfo          m.UidInfo
		verificationcode m.VerificationCode
		//		regRecord        m.RegRecord
	)

	ui.Phone = c.PostForm("mobile")
	passWord := c.PostForm("password")
	code := c.PostForm("code")
	way, errway := strconv.ParseInt(c.PostForm("way"), 10, 64)

	uidinfo.Phone = ui.Phone
	if uidinfo.GetUserPhone() {
		comm.ResponseError(c, 4002) //手机号已注册
		return
	}
	//判断验证码表中是否存在该记录
	verificationcode.Phone = ui.Phone
	verificationcode.Code = code
	if !verificationcode.GetVerificationCode() {
		comm.ResponseError(c, 4063)
		return
	}

	ui.PassWord = comm.SetAesValue(passWord, "fushow.cms")
	ui.NickName = string(Krand(10, KC_RAND_KIND_ALL))
	ui.Type = 99
	ui.Ip = c.ClientIP()
	if errway == nil {
		ui.RegWay = way
	} else {
		ui.RegWay = 0
	}
	flag, _ := ui.Reg()
	if !flag {
		comm.ResponseError(c, 4008)
		return
	}

	var (
		uij m.UidInfoJson
		uil m.UidLoginInfo
	)

	uij.Phone = c.PostForm("mobile")
	if !uij.GetUserInfo() {
		comm.ResponseError(c, 4010)
		return
	}

	//登录成功时生成cookie、token,并复制到session的结构体,存入redis
	if SetCookieSession(uij, c) != "" {
		comm.Response(c, gin.H{"session": "session error"})
		return
	} //ssion set
	uij.UID = setNormalAesValue(strconv.FormatInt(uij.Id, 10), "fushow.cms")
	uij.PassWord = ui.PassWord

	//生成登陆记录
	uil.Uid = uij.Id
	uil.Ip = c.ClientIP()

	uil.LoginWay = 1
	if errway == nil {
		uil.LoginWay = way
	} else {
		uil.LoginWay = 1
	}
	if !uil.UidLoginInfoAdd() {
		fmt.Println("登陆记录插入失败")
	} else {
		fmt.Println("登陆记录插入成功")
	}

	comm.Response(c, gin.H{"list": uij})
}

//判断昵称是否存在
func CheckNick(c *gin.Context) {
	var ui m.UidInfo
	ui.NickName = c.PostForm("nickname")
	ishas := ui.CheckNick()
	comm.Response(c, gin.H{"state": ishas})
}

//填写昵称
func Nick(c *gin.Context) {
	var ui m.UidInfo
	ui.NickName = c.PostForm("nickname")
	ui.Id, _ = strconv.ParseInt(c.PostForm("UID"), 10, 64)
	if !ui.Nick() {
		comm.ResponseError(c, 4009)
		return
	}
	comm.Response(c, nil)
}

//登录
func Log(c *gin.Context) {
	keycode := c.PostForm("keycode")
	capt := c.PostForm("captcha")
	if keycode == "" {
		comm.ResponseError(c, 4063) //请输入图形验证码
		return
	}
	if !captcha.VerifyString(capt, keycode) {
		comm.ResponseError(c, 4062) //图形验证码输入错误
		return
	}
	var (
		ui  m.UidInfoJson
		uil m.UidLoginInfo
	)
	if keycode == "" {
		comm.ResponseError(c, 4066) //请输入图形验证码
		return
	}
	if !captcha.VerifyString(capt, keycode) {
		comm.ResponseError(c, 4065) //图形验证码输入错误
		return
	}
	ui.Phone = c.PostForm("mobile")
	password := c.PostForm("password")
	way, errway := strconv.ParseInt(c.PostForm("way"), 10, 64)
	if !ui.GetUserInfo() {
		comm.ResponseError(c, 4010)
		return
	}

	password1 := comm.SetAesValue(password, "fushow.cms")
	if !strings.EqualFold(ui.PassWord, password1) {
		comm.ResponseError(c, 4011)
		return
	}
	//登录成功时生成cookie、token,并复制到session的结构体,存入redis
	if SetCookieSession(ui, c) != "" {
		comm.Response(c, gin.H{"session": "session error"})
		return
	} //ssion set
	ui.UID = setNormalAesValue(strconv.FormatInt(ui.Id, 10), "fushow.cms")
	ui.PassWord = comm.SetAesValue(password, "fushow.cms")

	//生成登陆记录
	uil.Uid = ui.Id
	uil.Ip = c.ClientIP()

	uil.LoginWay = 1
	if errway == nil {
		uil.LoginWay = way
	} else {
		uil.LoginWay = 1
	}
	if !uil.UidLoginInfoAdd() {
		fmt.Println("登陆记录插入失败")
	} else {
		fmt.Println("登陆记录插入成功")
	}

	comm.Response(c, gin.H{"list": ui})
}

//忘记密码手机验证码
func LoseRegval(c *gin.Context) {
	keycode := c.PostForm("keycode")
	captcha := c.PostForm("captcha")
	mobile := c.PostForm("mobile")
	functionType := "loseregval"
	errorCode := getVercodeFromServer(keycode, captcha, mobile, functionType)
	if errorCode != 0 {
		comm.ResponseError(c, errorCode)
	} else {
		comm.Response(c, nil)
	}
}

//重置密码
func ResetPass(c *gin.Context) {
	var (
		ui    m.UidInfo
		vcode m.VerificationCode
	)
	code := c.PostForm("code")
	vcode.Phone = c.PostForm("mobile")
	ui.Phone = vcode.Phone
	password := c.PostForm("password")
	password1 := comm.SetAesValue(password, "fushow.cms") //密码加密
	if !vcode.GetVerificationCode() {
		comm.ResponseError(c, 4006)
		return
	}
	if !strings.EqualFold(vcode.Code, code) {
		comm.ResponseError(c, 4007)
		return
	}
	if !ui.ResetPass(password1) {
		comm.ResponseError(c, 4013)
		return
	}
	comm.Response(c, nil)
}

//第三方登陆
func DiSanFangDengLu(c *gin.Context) {
	/*	var (
			test ThirdPartyLanding
			strs bytes.Buffer
		)
		test.Appid = "157ca89018b36a"
		test.Token = "96e26c0180780fec8d7ab49b18edd1d1"
		test.Type = "get_user_info"
		test.Code = c.PostForm("code")
		url := "http://open.51094.com/user/auth.html"
		strs.WriteString("type=")
		strs.WriteString(test.Type)
		strs.WriteString("&code=")
		strs.WriteString(test.Code)
		strs.WriteString("&appid=")
		strs.WriteString(test.Appid)
		strs.WriteString("&token=")
		strs.WriteString(test.Token)
		//发送请求
		resp, err := http.Post(url,
			"application/x-www-form-urlencoded",
			strings.NewReader(strs.String()))
		if err != nil {
			fmt.Println(err)
		}
		defer resp.Body.Close()
		//读取返回数据
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			// handle error
		}
		var config ConfigStruct
		//转换为对象
		json.Unmarshal(body, &config)
		if config.From == `` || config.Uniq == `` {
			c.JSON(200, gin.H{"state": false, "data": nil})
			return
		}
		//登陆成功了
		var (
			userinfo m.UidInfoJson
			uk       m.UserKey
		)
		str := strings.Replace(config.Name, " ", "", -1)
		userinfo.NickName = str
		userinfo.Form = config.From
		userinfo.PassWord = config.Uniq + config.From
		userinfo.Favicon = config.Img
		userinfo.UserName = test.Code
		uk.UserName = test.Code
		flag, message := m.UserLoginDoing(userinfo, uk)
		if !flag {
			c.JSON(200, gin.H{"state": false, "data": message})
			return
		}

		//查找数据
		if !userinfo.GetUserInfo() {
			c.JSON(200, gin.H{"state": false, "data": "用户不存在"})
			return
		}

		if SetCookieSession(userinfo, c) != "" {
			c.JSON(200, gin.H{"state": false, "data": "后台cookie"})
			return
		}
		userinfo.PassWord = ""
		userinfo.UID = comm.SetAesValue(strconv.FormatInt(userinfo.Id, 10), "fushow.cms")
		c.JSON(200, gin.H{"state": true, "data": userinfo})
	*/
}

type ConfigStruct struct {
	Name string `json:"name"`
	Img  string `json:"img"`
	Sex  int    `json:"sex"`
	Uniq string `json:"uniq"`
	From string `json:"from"`
}

//后台返回已修改
func PhoneLogin(c *gin.Context) {
	var (
		userinfoj m.UidInfoJson
		uk        m.UserKey
		uli       m.UidLoginInfo
	)
	//登陆成功了

	nickname := c.PostForm("nickname")
	openid := c.PostForm("openid")
	src := c.PostForm("src")
	form := c.PostForm("form")
	fmt.Println("nickname", nickname)
	fmt.Println("openid", openid)
	fmt.Println("src", src)
	fmt.Println("form", form)
	userinfoj.UserName = openid
	fmt.Println("openid", openid)
	if openid == "" {
		comm.ResponseError(c, 4010)
		return
	}
	//注册来源，1、PC  3、IOS  4、H5   5、Android原生  0、注册时未传信息
	way, errWay := strconv.ParseInt(c.PostForm("way"), 10, 64)
	//查找数据
	if !userinfoj.GetUserInfo() {
		//未找到数据，插入数据
		uk.UserName = openid
		userinfoj.NickName = string(Krand(10, KC_RAND_KIND_ALL))
		userinfoj.Form = form
		userinfoj.PassWord = openid
		userinfoj.Favicon = src
		userinfoj.Type = 99
		if errWay == nil {
			userinfoj.RegWay = way
		} else {
			userinfoj.RegWay = 0
		}
		flag, _ := m.UserLoginDoing(userinfoj.UidInfo, uk)
		if !flag {
			comm.ResponseError(c, 4055)
			return
		}

	}

	//找到数据或者未找到 但是已经成功插入数据后
	////////////////////////////
	//插入登陆信息
	uli.Uid = userinfoj.Id
	uli.Ip = c.ClientIP()
	if errWay == nil {
		uli.LoginWay = way
	} else {
		uli.LoginWay = 1 //因为该字段规范是后来定义的，所以，不会出现零
	}
	if !uli.UidLoginInfoAdd() {
		fmt.Println("登陆记录插入失败")
	} else {
		fmt.Println("登陆记录插入成功")
	}
	////////////////////////////

	var userinfok m.UidInfoJson
	userinfok.UserName = openid
	userinfok.UidInfo.GetUserInfo()
	if SetCookieSession(userinfok, c) != "" {
		comm.ResponseError(c, 4056)
		return
	}
	userinfok.PassWord = ""
	userinfok.UID = setNormalAesValue(strconv.FormatInt(userinfok.Id, 10), "fushow.cms")
	fmt.Println("userinfok", userinfok)
	comm.Response(c, userinfok)
}

//未找到使用的地方，暂留
func SearchOfType(c *gin.Context) {
	ptype, _ := strconv.ParseInt(c.PostForm("Type"), 10, 64)
	pleve, _ := strconv.ParseInt(c.PostForm("Leve"), 10, 64)
	fmt.Println(ptype, pleve)
	if pleve == 0 {
		c.JSON(200, gin.H{"total": 0, "rows": ""})
		return
	}
}

//用户头像修改 addby liuhan
//后台返回已修改
func UserUpFavicon(c *gin.Context) {
	var ui m.UidInfo
	//获取客户端参数
	ui.Id, _ = strconv.ParseInt(c.PostForm("UID"), 10, 64)
	if ui.Id == 0 {
		comm.ResponseError(c, 3175)
		return
	}
	if !ui.GetUserInfo() {
		comm.ResponseError(c, 4010)
		return
	}
	ui.Favicon = c.PostForm("Favicon")
	//修改用户信息
	if !ui.UserUpFavicon() { //失败返回
		comm.ResponseError(c, 4057)
		return
	} else { //成功返回
		comm.Response(c, nil)
	}

}

//修改绑定手机--发送验证码  -->通过手机号
func ByPhoneBindEditPhone(c *gin.Context) {
	keycode := c.PostForm("keycode")
	captcha := c.PostForm("captcha")
	mobile := c.PostForm("mobile")
	functionType := "bindphone"
	errorCode := getVercodeFromServer(keycode, captcha, mobile, functionType)
	if errorCode != 0 {
		comm.ResponseError(c, errorCode)
	} else {
		comm.Response(c, nil)
	}
}

func SlzPassUp(c *gin.Context) {
	var (
		ui m.UidInfo
	)
	//获取客户端参数
	userId, _ := strconv.ParseInt(c.PostForm("UID"), 10, 64) //用户id
	if userId == 0 {
		comm.ResponseError(c, 3175)
		return
	}
	password := c.PostForm("password")                    //密码
	password1 := comm.SetAesValue(password, "fushow.cms") //原密码加密
	ui.Id = userId
	ui.PassWord = password1
	//获取用户名密码
	has := ui.GetUserInfo()
	if !has {
		comm.ResponseError(c, 4011)
		return
	}
	if !strings.EqualFold(ui.PassWord, password1) {
		comm.ResponseError(c, 4033)
		return
	}
	comm.Response(c, 4034)
}

func GetNowDate(c *gin.Context) {
	//获取时间戳
	timestamp := time.Now().Unix()
	fmt.Println(timestamp)
	//格式化为字符串,tm为Time类型
	tm := time.Unix(timestamp, 0)
	comm.Response(c, tm.Format("2006-01-02"))
}
