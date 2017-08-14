package control

/*
 * Copyright (c) 2016—2017 www.fushow.cn, All rights reserved.
 */
import (
	"encoding/base64"
	"encoding/xml"
	"errors"
	"fmt"
	"fushowcms/comm"
	m "fushowcms/models"
	"image"
	"image/draw"
	"image/jpeg"
	"image/png"
	"io"
	"io/ioutil"
	"math/rand"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"

	alidayu "github.com/alidayu-master"
	"github.com/dchest/captcha"
	"github.com/garyburd/redigo/redis"
)

// 重写生成连接池方法
func newPoolUser() *redis.Pool {
	a, _ := strconv.Atoi(comm.GetConfig("POOL", "max_idle"))
	b, _ := strconv.Atoi(comm.GetConfig("POOL", "max_active"))
	return &redis.Pool{
		MaxIdle:   a,
		MaxActive: b,
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

type SessionUidInfo struct {
	Id            int64
	UserName      string
	NickName      string
	Type          uint64 //普通用户:0,主播:1,房管:2,总管:3,管理员:255
	Banned        bool   //默认flash，禁言时：true
	BannedEndTime int64  //禁言结束时间，默认空
	Level         uint64
	Favicon       string
	CheckToken    string
	Balance       float64 //余额
}

//生成随机字符串
func GetRandomString(leng int) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano() * rand.Int63()))
	for i := 0; i < leng; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}

func GetRandomBytes(leng int) []byte {
	str := "0123456789"
	bytes := []byte(str)
	result := make([]byte, leng) //[]byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano() * rand.Int63()))
	for i := 0; i < leng; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return result
}

//随机字符加密
func setNormalAesValue(v, atime string) string {
	v = GetRandomString(2) + v + GetRandomString(1)
	key := []byte("fushow" + atime)
	aec, err := comm.AesEncrypt([]byte(v), []byte(key))
	if err != nil {
		panic(err)
	}
	return base64.StdEncoding.EncodeToString(aec)
}

//解密
func getNormalDecValue(v, atime string) string {
	key := []byte("fushow" + atime)
	dec, err := comm.AesDecrypt(v, key)
	if err != nil {
		panic(err)
	}
	data := SubEndstr(string(dec), 2, len(string(dec))-1)
	return data
}

//将session从Redis移除
func DelSess(ui *SessionUidInfo) bool {
	rs := pool.Get()
	defer rs.Close()
	var key string
	key = "sess:" + strconv.FormatInt(ui.Id, 10)
	if _, err := rs.Do("DEL", redis.Args{}.Add(key).AddFlat(ui)...); err != nil {
		panic(err)
	}
	return true
}

//通过redis获取session ui的具体内容
func GetSess(ui *SessionUidInfo, isPC bool) bool {
	//更新
	var key string = "sessm"
	rs := pool.Get()
	defer rs.Close()
	if isPC {
		key = "sess"
	}
	key = key + ":" + strconv.FormatInt(ui.Id, 10)
	v, err := redis.Values(rs.Do("HGETALL", key))
	if err != nil {
		return false
	}
	if err := redis.ScanStruct(v, ui); err != nil {
		return false
	}
	if len(v) == 0 {
		return false
	}
	return true
}

//通过redis获取广播内容 addby liuhan
func GetApolloIM(id int64) Apollo {
	var b Apollo
	rs := pool.Get()
	defer rs.Close()
	key := "apollo" + strconv.FormatInt(id, 10)
	v, err := redis.Values(rs.Do("HGETALL", key))
	if err != nil {
		fmt.Println("err", err)
	}
	if err := redis.ScanStruct(v, &b); err != nil {
		fmt.Println("err:", err)
	}
	return b
}

type Apollo struct {
	UserName string
	PassWord []byte
}

//根据session判断用户操作权限
func IsAccess() gin.HandlerFunc { //c *gin.Context
	return func(c *gin.Context) {
		// 判断是否是合法主机
		var isPC bool = true

		//判断访问来源
		if c.Request.UserAgent() == "fushowphone:ios" || c.Request.UserAgent() == "fushowphone:android" {
			isPC = false
			//return
		}
		cookie, err := c.Request.Cookie(comm.GetConfig("SESSION", "cookie_name"))
		backURL := c.Request.Referer()
		if backURL == "" {
			backURL = c.Request.RequestURI
		}
		if err == nil {
			cookievalue := cookie.Value
			var ui SessionUidInfo
			ui.Id, _ = strconv.ParseInt(comm.GetDecValue(cookievalue, "fushow.cms"), 10, 64)
			if !GetSess(&ui, isPC) {
				c.Writer.Write([]byte("<b>会话过期，请重新登录</b>\n <script language='javascript'> window.location.href='/page/login_index?backurl=" + backURL + "'; </script>"))
				return
			}
		} else {
			c.Writer.Write([]byte("<b>您未登录，请重新登录</b>\n <script language='javascript'> window.location.href='/page/login_index?backurl=" + backURL + "'; </script>"))
			return
		}

	}
}

//根据session判断用户操作权限
func IsPageAccess() gin.HandlerFunc { //c *gin.Context
	return func(c *gin.Context) {
		//判断访问来源
		if c.Request.UserAgent() == "fushowphone:ios" || c.Request.UserAgent() == "fushowphone:android" {
			fmt.Println(c.Request.UserAgent(), "******************************访问")
			return
		}
		cookie, err := c.Request.Cookie(comm.GetConfig("SESSION", "cookie_name"))

		if err == nil {
			//更新redis 、cookie
			rs := pool.Get()
			defer rs.Close()
			cookievalue := cookie.Value
			id, _ := strconv.ParseInt(comm.GetDecValue(cookievalue, "fushow.cms"), 10, 64)
			// TODO: 试试不用改cookie
			SetCookie(c, true, cookievalue)
			//读取配置文件  更新时间
			dataTime := comm.GetConfig("SESSION", "expiration_time")
			rs.Send("Expire", "sess:"+strconv.FormatInt(id, 10), dataTime) //设置SESSION键的过期时间,单位:秒
		} else {
			return
		}

	}
}

//退出登录操作Seesion Cookie
func DelCookieSession(ui m.UidInfoJson, c *gin.Context) string {
	var su SessionUidInfo
	su.Id = ui.Id
	if !DelSess(&su) {
		return "session set fail"
	}
	cookie := http.Cookie{Name: comm.GetConfig("SESSION", "cookie_name"), Value: "", Path: "/", MaxAge: -1, HttpOnly: true}
	http.SetCookie(c.Writer, &cookie)
	return ""
}

//根据session判断用户操作权限
func IsRootAccess() gin.HandlerFunc { //c *gin.Context
	return func(c *gin.Context) {
		cookie, err := c.Request.Cookie(comm.GetConfig("SESSION", "cookie_name"))
		if err == nil {
			cookievalue := cookie.Value
			var ui SessionUidInfo
			ui.Id, _ = strconv.ParseInt(comm.GetDecValue(cookievalue, "fushow.cms"), 10, 64)
			if !GetSess(&ui, true) {
				c.Writer.Write([]byte("<b>会话过期，请重新登录</b>\n <script language='javascript'> window.location.href='/page/loginrootmyadmin'; </script>"))
			} else {
				if ui.Type <= 1 {
					c.Writer.Write([]byte("<b>请联系管理员</b>\n <script language='javascript'> window.location.href='/page/index'; </script>"))
				}
			}
		} else {
			c.Writer.Write([]byte("<b>您未登录，请重新登录</b>\n <script language='javascript'> window.location.href='/page/loginrootmyadmin'; </script>"))
		}
	}
}

func defaultCheckRedirect(req *http.Request, via []*http.Request) error {
	if len(via) >= 10 {
		return errors.New("stopped after 10 redirects")
	}
	return nil
}

//手机端的搜索
func GetSearchList(c *gin.Context) {
	var (
		ar m.AnchorRoom //roomResult -> rr
	)
	value := c.PostForm("value")
	if value == "" {
		fmt.Println("没输入东西")
	}
	livelist, userlist := ar.SearchRoomList(value)
	c.JSON(200, gin.H{"room": livelist, "user": userlist})
}

//将session存入Radis
func SetSess(ui *SessionUidInfo, isPC bool) bool {
	var str string
	rs := pool.Get()
	defer rs.Close()
	var key string
	if isPC {
		key = "sess:" + strconv.FormatInt(ui.Id, 10)
		str = comm.GetConfig("SESSION", "expiration_time")
	} else {
		key = "sessm:" + strconv.FormatInt(ui.Id, 10)
		str = comm.GetConfig("SESSION", "expiration_time_phone")
	}
	if err := rs.Send("HMSET", redis.Args{}.Add(key).AddFlat(ui)...); err != nil {
		panic(err)
		fmt.Println("out")
	}
	err := rs.Send("Expire", key, str) //设置SESSION键的过期时间,单位:秒
	fmt.Println("err", err)
	return true
}

//给session结构体内容赋值
func SetCookieSession(ui m.UidInfoJson, c *gin.Context) string {
	by, _ := ui.LoginTime.GobEncode()
	token := base64.StdEncoding.EncodeToString(by)
	ui.CheckToken = token
	var su SessionUidInfo
	var isPC bool = true
	su.Id = ui.Id
	su.UserName = ui.UserName
	su.NickName = ui.NickName
	su.Type = ui.Type
	su.Level = ui.Level
	su.Favicon = ui.Favicon
	su.CheckToken = ui.CheckToken
	av := comm.SetAesValue(strconv.FormatInt(ui.Id, 10), "fushow.cms")
	//判断登陆来源
	if c.Request.UserAgent() == "fushowphone:ios" || c.Request.UserAgent() == "fushowphone:android" {
		isPC = false
	}
	//redis一个月过期
	SetCookie(c, isPC, av)
	if !SetSess(&su, isPC) {
		//过期时间不同Redis
		return "session set fail"
	}
	//cookie相同
	return ""
}

//判断设置cookie
func SetCookie(c *gin.Context, isPc bool, uid string) {
	var strDate string
	if isPc {
		strDate = comm.GetConfig("SESSION", "expiration_time")
	} else {
		strDate = comm.GetConfig("SESSION", "expiration_time_phone")

	}
	cookietime, _ := strconv.ParseInt(strDate, 10, 64)
	cookie := http.Cookie{Name: comm.GetConfig("SESSION", "cookie_name"), Value: uid, Path: "/", MaxAge: /*86400*/ int(cookietime), HttpOnly: true}
	http.SetCookie(c.Writer, &cookie)
}

//文件上传
type Size interface {
	Size() int64
}

// 获取文件信息的接口
type Stat interface {
	Stat() (os.FileInfo, error)
}

func Upload(c *gin.Context) {
	fileName := time.Now().Unix()
	//获取文件内容 要这样获取
	file, _, err := c.Request.FormFile("uploadFile")
	fmt.Println("file", file)
	fmt.Println("file1", c.PostForm("uploadFile"))
	// 获取文件大小的接口
	if statInterface, ok := file.(Stat); ok {
		fileInfo, _ := statInterface.Stat()
		fmt.Println("上传文件的大小为: %d", fileInfo.Size())
		c.JSON(200, gin.H{"state": "失败"})
	}
	if sizeInterface, ok := file.(Size); ok {
		if sizeInterface.Size() > 204800 {
			c.JSON(200, gin.H{"state": "上传文件过大"})
			return
		}
		fmt.Println("上传文件的大小为: %d", sizeInterface.Size())
	}
	if err != nil {
		fmt.Println("error :", err)
		return
	}
	defer file.Close()
	//创建文件
	fW, err := os.Create("./static/upload/" + strconv.FormatInt(fileName, 10) + ".jpg")

	if err != nil {
		c.JSON(200, gin.H{"state": "文件创建失败"})
		return
	}
	defer fW.Close()
	_, err = io.Copy(fW, file)
	if err != nil {
		c.JSON(200, gin.H{"state": "文件保存失败"})
		return
	}
	c.JSON(200, gin.H{"state": "success", "imgNmae": strconv.FormatInt(fileName, 10) + ".jpg"})
}

func Uploads(c *gin.Context) (bool, string) {
	//获取文件内容 要这样获取
	fileName1 := time.Now().Unix()
	file, head, err := c.Request.FormFile("uploadFile")
	if err != nil {
		return false, ""
	}
	defer file.Close()
	//创建文件
	fW, err := os.Create("./static/upload/" + strconv.FormatInt(fileName1, 10) + "999")

	if err != nil {
		c.JSON(200, gin.H{"state": "文件创建失败"})
		return false, ""
	}
	defer fW.Close()
	_, err = io.Copy(fW, file)
	if err != nil {
		c.JSON(200, gin.H{"state": "文件保存失败"})
		fmt.Println("chuangjian shi bai 2")
		return false, ""
	}
	fmt.Println("file", head.Filename)
	return true, "/static/upload/" + strconv.FormatInt(fileName1, 10) + "999"
}

func Uploads_png(c *gin.Context) (bool, string) {
	//获取文件内容 要这样获取
	file, head, err := c.Request.FormFile("uploadFile_png")
	if err != nil {
		return false, ""
	}
	defer file.Close()
	//创建文件
	fW, err := os.Create("./static/upload/" + head.Filename)
	if err != nil {
		c.JSON(200, gin.H{"state": "文件创建失败"})
		return false, ""
	}
	defer fW.Close()
	_, err = io.Copy(fW, file)
	if err != nil {
		c.JSON(200, gin.H{"state": "文件保存失败"})
		fmt.Println("chuangjian shi bai 2")
		return false, ""
	}
	return true, "/static/upload/" + head.Filename
}

//TODO 页面没有用到  路由没有
func CUploads(c *gin.Context, name string) (bool, string) {
	fileName := time.Now().UnixNano()
	fileNames := strconv.FormatInt(fileName, 10) + string(Krand(4, KC_RAND_KIND_ALL))
	//获取文件内容 要这样获取
	file, _, err := c.Request.FormFile(name)
	if err != nil {
		return false, ""
	}
	defer file.Close()
	//创建文件
	fW, err := os.Create("./static/upload/" + fileNames)
	if err != nil {
		c.JSON(200, gin.H{"state": "文件创建失败"})
		return false, ""
	}
	defer fW.Close()
	_, err = io.Copy(fW, file)
	if err != nil {
		c.JSON(200, gin.H{"state": "文件保存失败"})
		return false, ""
	}
	return true, "/static/upload/" + fileNames
}

//处理get请求
func GetActionRoot(c *gin.Context) {
	var tplname string
	action := c.Param("action")
	if action != "" {
		tplname = "root" + action + ".html"
	}
	if !IsFileExist("views/" + tplname) {
		c.HTML(http.StatusFound, "user/404.html", nil)
		return
	}
	c.HTML(http.StatusOK, tplname, nil)

}

//处理get请求
func GetActionUser(c *gin.Context) {
	var tplname string
	action := c.Param("action")
	if action != "" {
		tplname = "user" + action + ".html"
	} else {
		tplname = "user/index.html"
	}
	if !IsFileExist("views/" + tplname) {
		c.HTML(http.StatusFound, "user/404.html", nil)

		return
	}
	c.HTML(http.StatusOK, tplname, nil)
}

func Getimagecode(c *gin.Context) {
	// 生成图片,随机数
	code := captcha.New()
	//生成图形验证码图片
	f, err := os.Create("./static/upload/captcha/" + code + ".png")
	if err != nil {
		comm.ResponseError(c, 4061)
	}
	defer f.Close()
	captcha.WriteImage(f, code, 240, 80)
	//返回图形验证码图片地址和id
	dd := struct {
		CaptchaId string
		ImageURL  string
	}{
		CaptchaId: code,
		ImageURL:  "/static/upload/captcha/" + code + ".png",
	}
	comm.Response(c, dd)
}

//处理get请求
func GetActionHelp(c *gin.Context) {
	var tplname string
	action := c.Param("action")
	if action != "" {
		tplname = "help" + action + ".html"
	} else {
		tplname = "help/index.html"
	}
	if !IsFileExist("views/" + tplname) {
		c.HTML(http.StatusFound, "user/404.html", nil)

		return
	}
	c.HTML(http.StatusOK, tplname, nil)
}

//处理get请求
func GetActionNews(c *gin.Context) {
	var tplname string
	action := c.Param("action")
	if action != "" {
		tplname = "news" + action + ".html"
	} else {
		tplname = "news/index.html"
	}
	if !IsFileExist("views/" + tplname) {
		c.HTML(http.StatusFound, "user/404.html", nil)
		return
	}
	c.HTML(http.StatusOK, tplname, nil)

}

//处理get请求
func GetActionRoomLive(c *gin.Context) {
	var tplname string
	action := c.Param("action")
	if action != "" {
		tplname = "page/roomlive.html"
		GetRoomLivePage(c)
		return
	} else {
		tplname = "page/index.html"
		return
	}
	if !IsFileExist("views/" + tplname) {
		c.HTML(http.StatusFound, "user/404.html", nil)
		return
	}
	c.HTML(http.StatusOK, tplname, nil)
}

func GetActionOutLive(c *gin.Context) {
	var tplname string
	action := c.Param("action")
	if action != "" {
		tplname = "page/outlive.html"
		GetOutLivePage(c)
		return
	} else {
		tplname = "page/index.html"
		return
	}
	if !IsFileExist("views/" + tplname) {
		c.HTML(http.StatusFound, "user/404.html", nil)
		return
	}
	c.HTML(http.StatusOK, tplname, nil)

}

//处理get请求
func GetActionCate(c *gin.Context) {
	var tplname string
	action := c.Param("action")
	if action != "" {
		GetCatePage(c)
		tplname = "page/cate.html"
		return
	} else {
		tplname = "page/index.html"
		return
	}
	if !IsFileExist("views/" + tplname) {
		c.HTML(http.StatusFound, "user/404.html", nil)
		return
	}
	c.HTML(http.StatusOK, tplname, nil)

}

//WEB头像上传
func UserUpload(c *gin.Context) {
	var user m.UidInfo
	user.Id, _ = strconv.ParseInt(c.PostForm("UID"), 10, 64)
	if c.PostForm("UID") == "" {
		c.JSON(200, gin.H{"state": "参数错误"})
		return
	}
	if !user.GetUserInfo() {
		c.JSON(200, gin.H{"state": "用户不存在"})
		return
	}
	file, head, err := c.Request.FormFile("uploadFile")
	if _, ok := file.(Stat); ok {
		c.JSON(200, gin.H{"state": "失败"})
	}
	if sizeInterface, ok := file.(Size); ok {
		fmt.Println("tsize", sizeInterface.Size())
		if sizeInterface.Size() < 2048 {
			c.JSON(200, gin.H{"state": "上传文件过大"})
			return
		}
	}
	if err != nil {
		c.JSON(200, gin.H{"state": "出错"})
		fmt.Println("error :", err)
		return
	}
	defer file.Close()
	//创建文件
	fW, err := os.Create("./static/upload/" + head.Filename)
	if err != nil {
		c.JSON(200, gin.H{"state": "文件创建失败"})
		return
	}
	defer fW.Close()
	_, err = io.Copy(fW, file)
	if err != nil {
		fmt.Println("131313")
		c.JSON(200, gin.H{"state": "文件保存失败"})
		return
	}
	user.Favicon = "/static/upload/" + head.Filename
	if !user.UserInfoUp() {
		c.JSON(200, gin.H{"state": "用户信息更新失败"})
		return
	}
	c.JSON(200, gin.H{"state": "success"})

}

//处理get请求
func GetActionPage(c *gin.Context) {
	var tplname string
	action := c.Param("action")
	fmt.Println("action", action)
	if strings.EqualFold(action, "/listsClass") {
		GetListsClassPage(c)
		return
	}
	if strings.EqualFold(action, "/lists") {
		GetListsPage(c)
		return
	}
	if strings.EqualFold(action, "/alltvlive") {
		GetAlltvLivePage(c)
		return
	}
	if strings.EqualFold(action, "/roomlive/*") {
		GetRoomLivePage(c)
		return
	}
	if action != "" {
		tplname = "page" + action + ".html"
	} else {
		GetIndexPage(c)
		return
	}
	if !IsFileExist("views/" + tplname) {
		c.HTML(http.StatusFound, "user/404.html", nil)
		return
	}
	c.HTML(http.StatusOK, tplname, nil)
	return
}

func GetActionTemp(c *gin.Context) {
	var tplname string
	action := c.Param("action")
	if action == "" {
		tplname = "temp/index.html"
	} else {
		tplname = "page" + action + ".html"
	}
	c.HTML(200, tplname, nil)
}

func IsFileExist(fil string) bool {
	f, err := os.Open(fil)
	if err != nil && os.IsNotExist(err) {
		return false
	}
	defer f.Close()
	return true
}

func GoodsUpload(c *gin.Context) {
	var user m.UidInfo
	user.Id, _ = strconv.ParseInt(c.PostForm("UID"), 10, 64)
	filename := c.PostForm("Filename")
	if c.PostForm("UID") == "" {
		c.JSON(200, gin.H{"state": "参数错误"})
		return
	}
	if !user.GetUserInfo() {
		c.JSON(200, gin.H{"state": "用户不存在"})
		return
	}
	file, head, err := c.Request.FormFile("uploadFile")
	if _, ok := file.(Stat); ok {
		c.JSON(200, gin.H{"state": "失败"})
	}
	if sizeInterface, ok := file.(Size); ok {
		fmt.Println("tsize", sizeInterface.Size())
		if sizeInterface.Size() < 2048 {
			c.JSON(200, gin.H{"state": "上传文件过大"})
			return
		}
	}
	if err != nil {
		c.JSON(200, gin.H{"state": "出错"})
		return
	}
	defer file.Close()
	//创建文件
	fmt.Println(",.,.,.,.,.,.:", head.Filename) //不可删
	fW, err := os.Create("./static/upload/goods/" + filename)
	if err != nil {
		c.JSON(200, gin.H{"state": "文件创建失败"})
		return
	}
	defer fW.Close()
	_, err = io.Copy(fW, file)
	if err != nil {
		c.JSON(200, gin.H{"state": "文件保存失败"})
		return
	}
	c.JSON(200, gin.H{"state": "success"})

}

func CategoryUpload(c *gin.Context) {
	var user m.UidInfo
	user.Id, _ = strconv.ParseInt(c.PostForm("UID"), 10, 64)
	filename := c.PostForm("Filename")
	if c.PostForm("UID") == "" {
		c.JSON(200, gin.H{"state": "参数错误"})
		return
	}
	if !user.GetUserInfo() {
		c.JSON(200, gin.H{"state": "用户不存在"})
		return

	}

	file, head, err := c.Request.FormFile("uploadFile")
	if _, ok := file.(Stat); ok {
		c.JSON(200, gin.H{"state": "失败"})
	}
	if sizeInterface, ok := file.(Size); ok {
		fmt.Println("tsize", sizeInterface.Size())
		if sizeInterface.Size() < 2048 {
			c.JSON(200, gin.H{"state": "上传文件过大"})
			return
		}
	}
	if err != nil {
		c.JSON(200, gin.H{"state": "出错"})
		return
	}
	defer file.Close()
	//创建文件
	fmt.Println(",.,.,.,.,.,.:", head.Filename) //不可删
	fmt.Println("filename:", filename)          //不可删
	fW, err := os.Create("./static/upload/category/" + filename)
	fmt.Println("err:", err)
	if err != nil {
		c.JSON(200, gin.H{"state": "文件创建失败"})
		return
	}
	defer fW.Close()
	_, err = io.Copy(fW, file)
	if err != nil {
		c.JSON(200, gin.H{"state": "文件保存失败"})
		return
	}
	c.JSON(200, gin.H{"state": "success"})
}

//获取期ID
type PIDX struct {
	Dateday string //当前时间
	Id      int64  //当前id
	Pid     string
}

var Pid PIDX

func InitPid() {
	Pid.Dateday = strings.Replace(time.Now().Format("2006-01-02"), "-", "", -1)
	Pid.Id = 0
	Pid.Pid = Pid.Dateday + strconv.FormatInt(Pid.Id, 10)
}

type Application struct {
	Name string `xml:"app_name,attr"`
	Data string `xml:"app_data,attr"`
}
type AppLog struct {
	Applications []Application `xml:"application"`
}
type CDR struct {
	XMLName xml.Name `xml:"cdr"`
	AppLog  AppLog   `xml:"app_log"`
}
type Result struct {
	Code  string `xml:"code"`
	Msg   string `xml:"msg"`
	Smsid string `xml:"smsid"`
}

// 手机端注册连接SMS短信服务器接口
func MobileSms(c *gin.Context) {
	account := comm.GetConfig("SMS", "account")
	password := comm.GetConfig("SMS", "password")
	urls := comm.GetConfig("SMS", "url")
	mobile := c.PostForm("mobile")
	//随机生成6位验证码
	code := rand.Intn(1000000)
	content := "您的验证码是：" + strconv.Itoa(code) + "。请不要把验证码泄露给其他人。如非本人操作，可不用理会！"
	//向短信服务器发送请求
	reqest, err := http.PostForm(urls, url.Values{"account": {account}, "password": {password}, "mobile": {mobile}, "content": {content}})
	if err != nil {
		fmt.Println(err)
		return
	}
	defer reqest.Body.Close()
	body, err := ioutil.ReadAll(reqest.Body)
	var result Result
	if err != nil {
		panic(err)
		return
	}
	errs := xml.Unmarshal(body, &result)
	fmt.Println("res", result)
	fmt.Println("err", errs)
	fmt.Println(string(body))
	if errs != nil {
		c.JSON(200, gin.H{"flag": true, "data": result, "error": errs})
		return
	}
	c.JSON(200, gin.H{"flag": false, "data": result, "error": errs})
}

/*
 * 功能:手机端忘记密码，验证手机短信
 * 请求参数2个: username、mobile、type(0:忘记密码，1:修改手机号)
 * 返回值2个: flag (判断短信服务器请求是否成功) 、 data (错误信息)
 * @徐林 20161012
 */
//TODO 老版获取手机验证码
func MobileSmS(c *gin.Context) {
	var uidinfo m.UidInfo
	//请求的参数 (手机号)
	uidinfo.UserName = c.PostForm("username")
	mobile := c.PostForm("mobile")
	types := c.PostForm("type")
	//查询用户名是否存在
	if !uidinfo.GetUserInfo() {
		c.JSON(200, gin.H{"flag": false, "data": "用户不存在"})
		return
	}
	//如果忘记密码修改密码
	if strings.EqualFold(types, "0") {
		//判断手机号
		if !strings.EqualFold(uidinfo.Phone, mobile) {
			c.JSON(200, gin.H{"flag": false, "data": "手机号错误"})
			return
		}
	}
	//随机生成6位验证码
	code := rand.Intn(1000000)
	str := strconv.Itoa(code)
	//向短信服务器发送请求
	alidayu.AppKey = "23477883"
	alidayu.AppSecret = "10a97d80159cfd05439d7f912de24e8a"
	success, resp := alidayu.SendSMS(mobile, "石榴联盟直播", "SMS_25815187", `{"code":"`+str+`","time":"5"}`)
	fmt.Println("Success:", success)
	fmt.Println(resp)
	if !success {

		return
	}
	//把短信验证码存到数据库中 更新用户表
	uidinfo.VerifyCode = strconv.Itoa(code)
	if !uidinfo.UserInfoUp() {
		//数据库更新失败
		c.JSON(200, gin.H{"flag": false})
		return
	}
	c.JSON(200, gin.H{"flag": true})

}

/**
 *fushou短信验证 addby liuhan 170122
 */
func VerificationCode(c *gin.Context) {
	keycode := c.PostForm("keycode")
	capt := c.PostForm("captcha")
	types, _ := strconv.ParseInt(c.PostForm("type"), 10, 64)
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
	if mobile == "" {
		comm.ResponseError(c, 4001) //手机号为空
		return
	}
	uidinfo.Phone = mobile
	if types == 1 {
		if !uidinfo.GetUserPhone() {
			comm.ResponseError(c, 4006) //不存在该手机号！
			return
		}
	} else {
		if uidinfo.GetUserPhone() {
			comm.ResponseError(c, 4002) //手机号已注册
			return
		}
	}

	//随机生成6位验证码
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	code := Substr(strconv.Itoa(r.Int()), 0, 6)
	account := comm.GetConfig("SMS", "account")
	password := comm.GetConfig("SMS", "password")
	urls := comm.GetConfig("SMS", "url")
	//随机生成6位验证码
	content := "您的验证码是：" + code + "。请不要把验证码泄露给其他人。如非本人操作，可不用理会！"
	//向短信服务器发送请求
	reqest, err := http.PostForm(urls, url.Values{"account": {account}, "password": {password}, "mobile": {mobile}, "content": {content}})
	if err != nil {
		fmt.Println(err)
		return
	}
	defer reqest.Body.Close()
	body, err := ioutil.ReadAll(reqest.Body)
	var result Result
	if err != nil {
		panic(err)
		return
	}
	errs := xml.Unmarshal(body, &result)
	fmt.Println("errs", errs)
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
	comm.Response(c, nil) //返回成功
}

/*
 * 功能:手机端忘记密码，判断输入验证码与数据库是否相同
 * 请求参数2个: username、mobile、code
 * 返回值2个: flag (判断短信服务器请求是否成功) 、 data (错误信息)
 * @徐林 20161012
 */
func CheckCode(c *gin.Context) {
	var uidinfo m.UidInfo
	//请求的参数 (手机号)
	uidinfo.UserName = c.PostForm("username")
	mobile := c.PostForm("mobile")
	code := c.PostForm("code")
	types := c.PostForm("type")

	fmt.Println("--------------------------", types)
	//查询手机号是否存在
	if !uidinfo.GetUserInfo() {
		c.JSON(200, gin.H{"flag": false, "data": "用户不存在"})
		return
	}
	if strings.EqualFold(types, "0") {
		//判断手机号
		if !strings.EqualFold(uidinfo.Phone, mobile) {
			c.JSON(200, gin.H{"flag": false, "data": "手机号错误"})
			return
		}
	}
	//判断验证码
	if !strings.EqualFold(uidinfo.VerifyCode, code) {
		c.JSON(200, gin.H{"flag": false, "data": "验证码错误"})
		return
	}
	c.JSON(200, gin.H{"flag": true})
}
func InitDatabase(c *gin.Context) {
	m.InitDataBase("admin", comm.SetAesValue("lol666@", "fushow.cms"))
	comm.Response(c, nil)
}

// 增加广播信息  addby liuhan
func SetBroadcastNew(b m.Broadcast, c *gin.Context) string {
	var bs Broadcasts
	bs.Id = b.Id
	bs.UserId = b.UserId
	bs.Content = b.Content
	if !SetBroadcast(&bs) {
		return "err"
	} else {
		return ""
	}
}

//将广播信息存入Radis addby liuhan
func SetBroadcast(b *Broadcasts) bool {
	rs := pool.Get()
	defer rs.Close()
	var key string
	key = "broadcast" //key是固定值
	if err := rs.Send("HMSET", redis.Args{}.Add(key).AddFlat(b)...); err != nil {
		panic(err)
		fmt.Println("err", err)
	}
	rs.Send("Expire", key, comm.GetConfig("SESSION", "broadcast_time"))
	return true
}

//广播管理  addby liuhan
type Broadcasts struct {
	Id      int64
	UserId  int64  //用户名称
	Content string //内容
}

//从资金从划拨
func FundAddUp(num float64) (bool, string) {
	var (
		fun     m.Fund //获取资金池
		newfund m.Fund //更新资金池
	)
	_, err := fun.GetFundDesc()
	if err != nil {
		return false, "资金池获取失败"
	}
	newfund.StorageFund = fun.StorageFund + num     // 储备金添加
	newfund.CurrencyMoney = fun.CurrencyMoney - num //流动资金减少
	newfund.OfferAmount = newfund.StorageFund + newfund.CurrencyMoney
	_, err = newfund.FundAdd()
	if err != nil {
		return false, "资金池更新失败"
	}
	return true, ""
}

//截取图片  addby liuhan
func ImageresizerUpload(c *gin.Context) {
	file, _, err := c.Request.FormFile("uploadFile")
	fileName := c.PostForm("FileName")
	arr := strings.Split(fileName, ".")
	i, _ := strconv.Atoi(c.PostForm("X"))
	j, _ := strconv.Atoi(c.PostForm("Y"))
	x1, _ := strconv.Atoi(c.PostForm("X1"))
	y1, _ := strconv.Atoi(c.PostForm("Y1"))
	if _, ok := file.(Stat); ok {
		comm.ResponseError(c, 3169) //失败
		return
	}
	if sizeInterface, ok := file.(Size); ok {
		fmt.Println("tsize", sizeInterface.Size())
		if sizeInterface.Size() > 2048000 {
			comm.ResponseError(c, 3170) //上传文件过大
			return
		}
	}
	if err != nil {
		comm.ResponseError(c, 3171) //出错
		return
	}
	// 上传原图存本地路径
	file1, err := os.Create("./static/upload/abcdefg.jpg")
	_, err = io.Copy(file1, file)
	file.Close()
	file1.Close()
	var format = arr[len(arr)-1]
	if strings.EqualFold(format, "jpeg") || strings.EqualFold(format, "jpg") {
		// 重新读回进行操作
		file2, _ := os.Open("./static/upload/abcdefg.jpg")
		file3, _ := os.Create("./static/upload/" + arr[0] + ".jpg")
		img, _ := jpeg.Decode(file2)                                               // jpg
		jpg := image.NewRGBA(image.Rect(0, 0, x1, y1))                             //要截图的大小
		draw.Draw(jpg, image.Rect(0, 0, x1, y1), img, image.Point{i, j}, draw.Src) //截取图片的一部分
		jpeg.Encode(file3, jpg, nil)
		file2.Close()
		file3.Close()
		err := os.Remove("./static/upload/abcdefg.jpg") //删除原文件
		fmt.Println("err", err)
	} else if strings.EqualFold(format, "png") {
		// 重新读回进行操作
		file2, _ := os.Open("./static/upload/abcdefg.jpg")
		file3, _ := os.Create("./static/upload/" + arr[0] + ".jpg")
		img, _ := png.Decode(file2)                                                // jpg
		jpg := image.NewRGBA(image.Rect(0, 0, x1, y1))                             //要截图的大小
		draw.Draw(jpg, image.Rect(0, 0, x1, y1), img, image.Point{i, j}, draw.Src) //截取图片的一部分
		jpeg.Encode(file3, jpg, nil)
		file2.Close()
		file3.Close()
		err := os.Remove("./static/upload/abcdefg.jpg") //删除原文件
		fmt.Println("err", err)
	} else {
		comm.ResponseError(c, 3172) //图片格式不正确
		return
	}
	comm.ResponseError(c, 3173) //图片格式不正确
}

//截取字符串 start 起点下标 end 终点下标(不包括)
func SubEndstr(str string, start int, end int) string {
	rs := []rune(str)
	length := len(rs)
	if len(str) < 0 {
		fmt.Println("字符串截取失败")
		return ""
	}
	if start < 0 || start > length {
		//panic("start is wrong")
		fmt.Println("截取字符串参数错误,起点下标错误")
	}
	if end < 0 || end > length {
		//panic("end is wrong")
		fmt.Println("截取字符串参数错误,终点下标错误")
	}
	return string(rs[start:end])
}

//获得当前日期
func GetDate() string {
	return time.Now().Format("2006-01-02")
}

//获得当前时间
func GetDateTime() string {
	return time.Now().Format("2006-01-02 15:04:05")
}
