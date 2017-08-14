package control

/*
 * Copyright (c) 2016—2017 www.fushow.cn, All rights reserved.
 */
import (
	"bytes"
	"crypto/sha1"
	"encoding/json"
	"fmt"
	"fushowcms/comm"
	mm "fushowcms/models"
	"io"
	"io/ioutil"
	"math/rand"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

const (
	KC_RAND_KIND_NUM   = 0                    // 纯数字
	KC_RAND_KIND_LOWER = 1                    // 小写字母
	KC_RAND_KIND_UPPER = 2                    // 大写字母
	KC_RAND_KIND_ALL   = 3                    // 数字、大小写字母
	AppId              = "XXXXXXXXXXXXXXXXXX" //微信公众平台AppId
	AppSecret          = "XXXXXXXXXXXXXXXXXX" //微信公众平台AppSecret
)

type TestWeiXin struct {
	Access_token  string `json:"access_token"`
	Expires_in    int    `json:"expires_in"`
	Refresh_token string `json:"refresh_token"`
	Openid        string `json:"openid"`
	Scope         string `json:"scope"`
	Unionid       string `json:"unionid"`
	Errcode       int    `json:"errcode"`
	Errmsg        string `json:"errmsg"`
}
type JsapiTicket struct {
	Errcode    int    `json:"errcode"`
	Errmsg     string `json:"errmsg"`
	Ticket     string `json:"ticket"`
	Expires_in int    `json:"expires_in"`
}
type Unionid struct {
	Openid     string `json:"openid"`
	Nickname   string `json:"nickname"`
	Sex        int    `json:"sex"`
	Province   string `json:"province"`
	City       string `json:"city"`
	Country    string `json:"country"`
	Headimgurl string `json:"headimgurl"`
	Privilege  string `json:"privilege"`
	Unionid    string `json:"unionid"`
}

/*微信用户信息获取*/
func WeiXinUserInfo(c *gin.Context) {
	code := c.PostForm("code")
	var str, url bytes.Buffer
	str.WriteString(`https://api.weixin.qq.com/sns/oauth2/access_token?appid=wx39232dd6a07ff4b8&secret=ad93d3b843e7b72e4b47896b925913ab&code=`)
	str.WriteString(code)
	str.WriteString(`&grant_type=authorization_code`)
	resp, _ := http.Get(str.String())
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	var test TestWeiXin
	json.Unmarshal(body, &test)
	url.WriteString(`https://api.weixin.qq.com/sns/userinfo?access_token=`)
	url.WriteString(test.Access_token)
	url.WriteString(`&openid=`)
	url.WriteString(test.Openid)
	resps, _ := http.Get(url.String())
	defer resps.Body.Close()
	bodys, _ := ioutil.ReadAll(resps.Body)
	var userinfo UserInfoModel
	json.Unmarshal(bodys, &userinfo)
	var (
		userinfoj mm.UidInfoJson
		uk        mm.UserKey
		uli       mm.UidLoginInfo
	)
	fmt.Println("userinfo", userinfo)
	userinfoj.UserName = userinfo.Unionid
	isExit := false
	//查找数据
	if !userinfoj.GetUserInfoByName() {
		userinfoj.NickName = string(Krand(10, KC_RAND_KIND_ALL))
		userinfoj.Form = "weixin"
		userinfoj.PassWord = userinfo.Unionid
		userinfoj.Favicon = userinfo.Headimgurl
		userinfoj.Type = 99
		userinfoj.RegWay = 1
		uk.UserName = userinfo.Unionid
		flag, _ := mm.UserLoginDoing(userinfoj.UidInfo, uk)
		if !flag {
			comm.ResponseError(c, 3183) //登录失败
			return
		}
		//查找数据
		if !userinfoj.GetUserInfo() {
			comm.ResponseError(c, 3184) //用户不存在
			return
		}
		isExit = true
	}
	//插入登陆信息
	uli.Uid = userinfoj.Id
	uli.LoginWay = 1
	uli.Ip = c.ClientIP()
	if !uli.UidLoginInfoAdd() {
		fmt.Println("登陆记录插入失败")
	} else {
		fmt.Println("登陆记录插入成功")
	}
	if SetCookieSession(userinfoj, c) != "" {
		comm.ResponseError(c, 3185) //session错误
		return
	}
	userinfoj.PassWord = ""
	userinfoj.UID = setNormalAesValue(strconv.FormatInt(userinfoj.Id, 10), "fushow.cms")
	if isExit {
		//微信登陆送石榴籽
		_, _, number := mm.BindGiveNumber(userinfoj.UidInfo.Id, 0)
		m := make(map[string]interface{})
		m["data"] = userinfoj
		m["Balance"] = number
		comm.Response(c, m)
		return
	}
	m := make(map[string]interface{})
	m["data"] = userinfoj
	m["Balance"] = ""
	comm.Response(c, m)
}

type UserInfoModel struct {
	Openid     string `json:"openid"`
	Nickname   string `json:"nickname"`
	Sex        int    `json:"sex"`
	Language   string `json:"language"`
	City       string `json:"city"`
	Province   string `json:"province"`
	Country    string `json:"country"`
	Headimgurl string `json:"headimgurl"`
	Privilege  string `json:"privilege"`
	Unionid    string `json:"unionid"`
	Errcode    int    `json:"errcode"`
	Errmsg     string `json:"errmsg"`
}

/*QQ登陆页面跳转*/
func QQLogin(c *gin.Context) {
	c.Writer.WriteString(`<script language='javascript'> window.location.href="https://graph.qq.com/oauth2.0/authorize?response_type=code&client_id=101359534&redirect_uri=http://www.shiliulianmeng.com&state=shiliulianmeng";</script>`)
}

type QQAccessToken struct {
	access_token  string `json:"access_token"`
	expires_in    int    `json:"expires_in"`
	refresh_token string `json:"refresh_token"`
}

type QQOpenid struct {
	Client_id string `json:"client_id"`
	Openid    string `json:"openid"`
}

type QQUserInfoModel struct {
	Ret                int    `json:"ret"`
	Msg                string `json:"msg"`
	Is_lost            int    `json:"is_lost"`
	Nickname           string `json:"nickname"`
	Gender             string `json:"gender"`
	Province           string `json:"province"`
	City               string `json:"city"`
	Year               string `json:"year"`
	Figureurl          string `json:"figureurl"`
	Figureurl_1        string `json:"figureurl_1"`
	Figureurl_2        string `json:"figureurl_2"`
	Figureurl_qq_1     string `json:"figureurl_qq_1"`
	Figureurl_qq_2     string `json:"figureurl_qq_2"`
	Is_yellow_vip      string `json:"is_yellow_vip"`
	Vip                string `json:"vip"`
	Yellow_vip_level   string `json:"yellow_vip_level"`
	Level              string `json:"level"`
	Is_yellow_year_vip string `json:"is_yellow_year_vip"`
}

/*QQ用户信息获取*/
func QQUserInfo(c *gin.Context) {
	var (
		userinfoj mm.UidInfoJson
		uk        mm.UserKey
		uli       mm.UidLoginInfo
	)
	code := c.PostForm("code")
	var str, getinfo bytes.Buffer
	str.WriteString(`https://graph.qq.com/oauth2.0/token?grant_type=authorization_code&client_id=101359534&client_secret=4ac0e832a30e73a8c0d15093b06cb846&code=`)
	str.WriteString(code)
	str.WriteString(`&redirect_uri=http://www.shiliulianmeng.com`)
	resp, err := http.Get(str.String())
	if err != nil {
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	m, _ := url.ParseQuery(string(body))
	oppenid := `https://graph.qq.com/oauth2.0/me?access_token=` + m["access_token"][0]
	oppenidResp, _ := http.Get(oppenid)
	defer oppenidResp.Body.Close()
	oppenidBody, _ := ioutil.ReadAll(oppenidResp.Body)
	qqsubstr := QQSubstr(string(oppenidBody))
	var qqopenid QQOpenid
	json.Unmarshal([]byte(qqsubstr), &qqopenid)
	getinfo.WriteString(`https://graph.qq.com/user/get_user_info?access_token=`)
	getinfo.WriteString(m["access_token"][0])
	getinfo.WriteString(`&oauth_consumer_key=101359534&openid=`)
	getinfo.WriteString(qqopenid.Openid)
	getinforesp, _ := http.Get(getinfo.String())
	defer getinforesp.Body.Close()
	getinfobody, _ := ioutil.ReadAll(getinforesp.Body)
	var userinfo QQUserInfoModel
	json.Unmarshal(getinfobody, &userinfo)
	var qqopenid1 QQOpenid
	json.Unmarshal([]byte(qqsubstr), &qqopenid1)
	getinfo.WriteString(`https://graph.qq.com/oauth2.0/me?access_token=`)
	getinfo.WriteString(m["access_token"][0])
	getinfo.WriteString(`&oauth_consumer_key=101359534&openid=`)
	getinfo.WriteString(qqopenid1.Openid)
	//查找数据
	isExit := false
	userinfoj.UserName = qqopenid1.Openid
	if !userinfoj.GetUserInfoByName() {
		userinfoj.NickName = string(Krand(10, KC_RAND_KIND_ALL))
		userinfoj.Form = "qq"
		userinfoj.PassWord = qqopenid1.Openid
		userinfoj.Favicon = userinfo.Figureurl_qq_2
		userinfoj.Type = 99
		//给字段赋值，注册设备
		userinfoj.RegWay = 1
		uk.UserName = qqopenid1.Openid
		flag, _ := mm.UserLoginDoing(userinfoj.UidInfo, uk)
		if !flag {
			comm.ResponseError(c, 3183) //登录失败
			return
		}
		if !userinfoj.GetUserInfo() {
			comm.ResponseError(c, 3184) //用户不存在
			return
		}
		isExit = true
	}
	//插入登陆信息
	uli.Uid = userinfoj.Id
	uli.LoginWay = 1
	uli.Ip = c.ClientIP()
	if !uli.UidLoginInfoAdd() {
		fmt.Println("登陆记录插入失败")
	} else {
		fmt.Println("登陆记录插入成功")
	}
	if SetCookieSession(userinfoj, c) != "" {
		comm.ResponseError(c, 3185) //session错误
		return
	}
	userinfoj.PassWord = ""
	userinfoj.UID = setNormalAesValue(strconv.FormatInt(userinfoj.Id, 10), "fushow.cms")
	//第一次第三方登录
	if isExit {
		//qq登陆送石榴籽
		_, _, number := mm.BindGiveNumber(userinfoj.UidInfo.Id, 0)
		m1 := make(map[string]interface{})
		m1["data"] = userinfoj
		m1["Balance"] = number
		comm.Response(c, m1)
		//c.JSON(200, gin.H{"state": true, "data": userinfoj, "Balance": number})
		return
	}
	m1 := make(map[string]interface{})
	m1["data"] = userinfoj
	m1["Balance"] = ""
	comm.Response(c, m1)
}

//字符串截取
func QQSubstr(str string) string {
	start := strings.Index(str, `(`) + 1
	endLen := strings.LastIndex(str, `)`)
	length := endLen - start
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

/*微博授权登录*/
func WBLogin(c *gin.Context) {
	c.Writer.WriteString(`<script language='javascript'> window.location.href="https://api.weibo.com/oauth2/authorize?client_id=1099460121&response_type=code&redirect_uri=http://tv.fushow.cn";</script>`)
}

type WBCallBack struct {
	Access_token string `json:"access_token"`
	Expires_in   int    `json:"expires_in"`
	Remind_in    string `json:"remind_in"`
	Uid          string `json:"uid"`
}

type WBUserInfo struct {
	Id                  int64  `json:"id"`
	Idstr               string `json:"idstr"`
	Class               int    `json:"class"` //
	Screen_name         string `json:"screen_name"`
	Name                string `json:"name"`
	Province            string `json:"province"`
	City                string `json:"city"`
	Location            string `json:"location"`
	Description         string `json:"description"`
	Url                 string `json:"url"`
	Profile_image_url   string `json:"profile_image_url"`
	Profile_url         string `json:"profile_url"`
	Domain              string `json:"domain"`
	Weihao              string `json:"weihao"`
	Gender              string `json:"gender"`
	Followers_count     int    `json:"followers_count"`
	Friends_count       int    `json:"friends_count"`
	Pagefriends_count   int    `json:"pagefriends_count"` //
	Statuses_count      int    `json:"statuses_count"`
	Favourites_count    int    `json:"favourites_count"`
	Created_at          string `json:"created_at"`
	Following           bool   `json:"following"`
	Allow_all_act_msg   bool   `json:"allow_all_act_msg"`
	Geo_enabled         bool   `json:"geo_enabled"`
	Verified            bool   `json:"verified"`
	Verified_type       int    `json:"verified_type"`
	Remark              string `json:"remark"`
	Ptype               int    `json:"ptype"`
	Allow_all_comment   bool   `json:"allow_all_comment"`
	Avatar_large        string `json:"avatar_large"`
	Avatar_hd           string `json:"avatar_hd"`
	Verified_reason     string `json:"verified_reason"`
	Verified_trade      string `json:"verified_trade"`      //
	Verified_reason_url string `json:"verified_reason_url"` //
	Verified_source     string `json:"verified_source"`     //
	Verified_source_url string `json:"verified_source_url"` //
	Follow_me           bool   `json:"follow_me"`
	Online_status       int    `json:"online_status"`
	Bi_followers_count  int    `json:"bi_followers_count"`
	Lang                string `json:"lang"`
	Star                int    `json:"star"`         //
	Mbtype              int    `json:"mbtype"`       //
	Mbrank              int    `json:"mbrank"`       //
	Block_word          int    `json:"block_word"`   //
	Block_app           int    `json:"block_app"`    //
	Credit_score        int    `json:"credit_score"` //
	User_ability        int    `json:"user_ability"` //
	Urank               int    `json:"urank"`        //
}

/*微博用户信息获取*/
func WBGetUserInfo(c *gin.Context) {
	//获取code
	code := c.PostForm("code")
	var url, geturl bytes.Buffer
	url.WriteString(`https://api.weibo.com/oauth2/access_token?client_id=1099460121&client_secret=761b48657a44ca3e42191c09a559ef5c&grant_type=authorization_code&code=`)
	url.WriteString(code)
	url.WriteString(`&redirect_uri=http://tv.fushow.cn`)
	resp, _ := http.Post(url.String(), "application/x-www-form-urlencoded", nil)
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	var wbcb WBCallBack
	json.Unmarshal(body, &wbcb)
	geturl.WriteString(`https://api.weibo.com/2/users/show.json?access_token=`)
	geturl.WriteString(wbcb.Access_token)
	geturl.WriteString(`&uid=`)
	geturl.WriteString(wbcb.Uid)
	wbResp, _ := http.Get(geturl.String())
	defer wbResp.Body.Close()
	oppenidBody, _ := ioutil.ReadAll(wbResp.Body)
	var wbuserinfo WBUserInfo
	json.Unmarshal(oppenidBody, &wbuserinfo)
	var (
		userinfoj mm.UidInfoJson
		uk        mm.UserKey
	)
	//登陆成功了
	userinfoj.UserName = strconv.FormatInt(wbuserinfo.Id, 10)
	isExit := false
	//查找数据
	if !userinfoj.GetUserInfoByName() {
		userinfoj.NickName = string(Krand(10, KC_RAND_KIND_ALL))
		userinfoj.Form = "weibo"
		userinfoj.PassWord = strconv.FormatInt(wbuserinfo.Id, 10)
		userinfoj.Favicon = wbuserinfo.Profile_image_url
		userinfoj.Type = 99
		uk.UserName = strconv.FormatInt(wbuserinfo.Id, 10)
		flag, _ := mm.UserLoginDoing(userinfoj.UidInfo, uk)
		if !flag {
			comm.ResponseError(c, 3183) //登录失败
			return
		}
		if !userinfoj.GetUserInfo() {
			comm.ResponseError(c, 3184) //用户不存在
			return
		}
		isExit = true
	}
	var uli mm.UidLoginInfo
	//插入登陆信息
	uli.Uid = userinfoj.Id
	uli.LoginWay = 1
	uli.Ip = c.ClientIP()
	if !uli.UidLoginInfoAdd() {
		fmt.Println("登陆记录插入失败")
	} else {
		fmt.Println("登陆记录插入成功")
	}
	if SetCookieSession(userinfoj, c) != "" {
		comm.ResponseError(c, 3185) //session错误
		return
	}
	userinfoj.PassWord = ""
	userinfoj.UID = setNormalAesValue(strconv.FormatInt(userinfoj.Id, 10), "fushow.cms")
	//是否是第一次第三方登录
	if isExit {
		//qq登陆送石榴籽
		_, _, number := mm.BindGiveNumber(userinfoj.Id, 0)
		m := make(map[string]interface{})
		m["data"] = userinfoj
		m["Balance"] = number
		comm.Response(c, m)
		return
	}
	m := make(map[string]interface{})
	m["data"] = userinfoj
	m["Balance"] = ""
	comm.Response(c, m)
}

func TestAliyun(c *gin.Context) {
	resp := c.Request
	bodys, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(bodys))
}

// 随机字符串
func Krand(size int, kind int) []byte {
	ikind, kinds, result := kind, [][]int{[]int{10, 48}, []int{26, 97}, []int{26, 65}}, make([]byte, size)
	is_all := kind > 2 || kind < 0
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < size; i++ {
		if is_all { // random ikind
			ikind = rand.Intn(3)
		}
		scope, base := kinds[ikind][0], kinds[ikind][1]
		result[i] = uint8(base + rand.Intn(scope))
	}
	return result
}

//调用微信js @liuhan
func WeChat(c *gin.Context) {
	url := c.PostForm("url")
	var strAT, strJT bytes.Buffer
	//获取access_token
	strAT.WriteString(`https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=`)
	strAT.WriteString(AppId)
	strAT.WriteString(`&secret=`)
	strAT.WriteString(AppSecret)
	accessToken, _ := http.Get(strAT.String())
	fmt.Println("accessToken", accessToken)
	defer accessToken.Body.Close()
	body, _ := ioutil.ReadAll(accessToken.Body)
	var test TestWeiXin
	json.Unmarshal(body, &test)
	fmt.Println("openid", test.Openid)
	//根据access_token获取jsapi_ticket
	strJT.WriteString(`https://api.weixin.qq.com/cgi-bin/ticket/getticket?access_token=`)
	fmt.Println("test.Access_token", test.Access_token)
	strJT.WriteString(test.Access_token)
	strJT.WriteString(`&type=jsapi`)
	jsapiTicket, _ := http.Get(strJT.String())
	fmt.Println("jsapiTicket", jsapiTicket)
	defer jsapiTicket.Body.Close()
	bodyJT, _ := ioutil.ReadAll(jsapiTicket.Body)
	var jsap JsapiTicket
	json.Unmarshal(bodyJT, &jsap)
	fmt.Println("errcode", jsap.Errcode)
	fmt.Println("errmsg", jsap.Errmsg)
	fmt.Println("ticket", jsap.Ticket)
	fmt.Println("expires_in", jsap.Expires_in)
	noncestr := strconv.Itoa(rand.Int())
	jsapi_ticket := jsap.Ticket
	timestamp := strconv.FormatInt(time.Now().Unix(), 10)
	var signature bytes.Buffer
	signature.WriteString(`jsapi_ticket=`)
	signature.WriteString(jsapi_ticket)
	signature.WriteString(`&noncestr=`)
	signature.WriteString(noncestr)
	signature.WriteString(`&timestamp=`)
	signature.WriteString(timestamp)
	signature.WriteString(`&url=`)
	signature.WriteString(url)
	fmt.Println("signature.String()", signature.String())
	h := sha1.New()
	io.WriteString(h, signature.String())
	s := fmt.Sprintf("%x", h.Sum(nil))
	fmt.Printf("% x", h.Sum(nil))
	m := make(map[string]interface{})
	m["appId"] = AppId
	m["timestamp"] = timestamp
	m["nonceStr"] = noncestr
	m["signature"] = s //h.Sum(nil)
	comm.Response(c, m)
}

// 微信公众平台授权登录
func WechatLog(c *gin.Context) {
	code := c.PostForm("code")
	var strA, strB bytes.Buffer
	var isBoo = false
	strA.WriteString(`https://api.weixin.qq.com/sns/oauth2/access_token?appid=`)
	strA.WriteString(AppId)
	strA.WriteString(`&secret=`)
	strA.WriteString(AppSecret)
	strA.WriteString(`&code=`)
	strA.WriteString(code)
	strA.WriteString(`&grant_type=authorization_code`)
	accessToken, _ := http.Get(strA.String())
	fmt.Println("accessToken", accessToken)
	defer accessToken.Body.Close()
	body, _ := ioutil.ReadAll(accessToken.Body)
	var test TestWeiXin
	json.Unmarshal(body, &test)
	strB.WriteString(`https://api.weixin.qq.com/sns/userinfo?access_token=`)
	strB.WriteString(test.Access_token)
	strB.WriteString(`&openid=`)
	strB.WriteString(test.Openid)
	strB.WriteString(`&lang=zh_CN`)
	unionid, _ := http.Get(strB.String())
	defer unionid.Body.Close()
	body, _ = ioutil.ReadAll(unionid.Body)
	var unionids Unionid
	json.Unmarshal(body, &unionids)
	var (
		userinfoj mm.UidInfoJson
		uk        mm.UserKey
		uli       mm.UidLoginInfo
	)
	if unionids.Unionid == "" {
		comm.ResponseError(c, 3183) //登录失 败
		return
	}
	userinfoj.UserName = unionids.Unionid
	isExit := false
	//查找数据
	if !userinfoj.GetUserInfoByName() {
		isBoo = true
		userinfoj.NickName = string(Krand(10, KC_RAND_KIND_ALL))
		userinfoj.Form = "weixin"
		userinfoj.PassWord = unionids.Unionid
		userinfoj.Favicon = unionids.Headimgurl
		userinfoj.Type = 99
		userinfoj.RegWay = 1
		uk.UserName = unionids.Unionid
		flag, _ := mm.UserLoginDoing(userinfoj.UidInfo, uk)
		if !flag {
			comm.ResponseError(c, 3183) //登录失败
			return
		}
		//查找数据
		if !userinfoj.GetUserInfo() {
			comm.ResponseError(c, 3184) //用户不存在
			return
		}
		isExit = true
	}
	//插入登陆信息
	uli.Uid = userinfoj.Id
	uli.LoginWay = 1
	uli.Ip = c.ClientIP()
	if !uli.UidLoginInfoAdd() {
		fmt.Println("登陆记录插入失败")
	} else {
		fmt.Println("登陆记录插入成功")
	}
	if SetCookieSession(userinfoj, c) != "" {
		comm.ResponseError(c, 3185) //session错误
		return
	}
	userinfoj.PassWord = ""
	userinfoj.UID = setNormalAesValue(strconv.FormatInt(userinfoj.Id, 10), "fushow.cms")
	if isExit {
		//微信登陆送石榴籽
		_, _, number := mm.BindGiveNumber(userinfoj.UidInfo.Id, 0)
		m := make(map[string]interface{})
		m["data"] = userinfoj
		m["Balance"] = number
		m["Openid"] = test.Openid
		comm.Response(c, m)
		return
	}
	m := make(map[string]interface{})
	m["data"] = userinfoj
	m["Balance"] = ""
	m["Openid"] = test.Openid
	m["IsBoo"] = isBoo
	comm.Response(c, m)
}

// 微信公众平台授权支付
func WechatLogPay(c *gin.Context) {
	code := c.PostForm("code")
	var strA bytes.Buffer
	strA.WriteString(`https://api.weixin.qq.com/sns/oauth2/access_token?appid=`)
	strA.WriteString(AppId)
	strA.WriteString(`&secret=`)
	strA.WriteString(AppSecret)
	strA.WriteString(`&code=`)
	strA.WriteString(code)
	strA.WriteString(`&grant_type=authorization_code`)
	accessToken, _ := http.Get(strA.String())
	fmt.Println("accessToken", accessToken)
	defer accessToken.Body.Close()
	body, _ := ioutil.ReadAll(accessToken.Body)
	var test TestWeiXin
	json.Unmarshal(body, &test)
	m := make(map[string]interface{})
	m["Openid"] = test.Openid
	comm.Response(c, m)
}
