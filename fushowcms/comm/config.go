package comm

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"fmt"
	"os"
	"runtime"
	"strconv"
	"sync"

	"github.com/gin-gonic/gin"

	"github.com/widuu/goini"
)

//conf := goini.SetConfig("./conf/conf.ini") //goini.SetConfig(filepath) 其中filepath是你ini 配置文件的所在位置
var InitConf = goini.SetConfig("conf/app.conf") //goini.SetConfig(filepath) 其中filepath是你conf配置文件的所在位置
var MsgConf = goini.SetConfig("conf/msg.conf")

// GetConfig 统一读配置
func GetConfig(section, name string) string {
	v := InitConf.GetValue(section, name)
	if v != "no value" {
		if section == "DB" && name == "user" || name == "pass" {
			return GetDecValue(v, "fushow.cms")
		}
		if section == "REDIS" && name == "pass" {
			return GetDecValue(v, "fushow.cms")
		}
	}
	return v
}

// SetConfig 统一写配置
func SetConfig(section, name, data string) {
	InitConf.SetValue(section, name, data)
}

//数据库
func AesEncrypt(origData, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockSize := block.BlockSize()
	origData = PKCS5Padding(origData, blockSize)
	blockMode := cipher.NewCBCEncrypter(block, key[:blockSize])
	crypted := make([]byte, len(origData))
	// 根据CryptBlocks方法的说明，如下方式初始化crypted也可以
	// crypted := origData
	blockMode.CryptBlocks(crypted, origData)
	return crypted, nil
}

func AesDecrypt(crypted string, key []byte) ([]byte, error) {
	bt, _ := base64.StdEncoding.DecodeString(crypted)
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockSize := block.BlockSize()
	blockMode := cipher.NewCBCDecrypter(block, key[:blockSize])
	origData := make([]byte, len(bt))
	blockMode.CryptBlocks(origData, bt)
	origData = PKCS5UnPadding(origData)
	return origData, nil
}
func PKCS5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func PKCS5UnPadding(origData []byte) []byte {
	length := len(origData)
	// 去掉最后一个字节 unpadding 次
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}

//不添加随机字符   加密
func SetAesValue(v, atime string) string {
	key := []byte("fushow" + atime)
	aec, err := AesEncrypt([]byte(v), []byte(key))
	if err != nil {
		panic(err)
	}
	return base64.StdEncoding.EncodeToString(aec)
}

//解密
func GetDecValue(v, atime string) string {
	key := []byte("fushow" + atime)
	dec, err := AesDecrypt(v, key)
	if err != nil {
		panic(err)
	}
	return string(dec)
}

var FSLog = &sync.Mutex{}
var QDLog = &sync.Mutex{}

// 常量
const (
	ERROR_0       = "ok"
	ERROR_DEFAULT = "操作异常，请检查！"
)

// 每个人错误提示信息写conf/app.conf文件,所有错误都在写[ERROR]字段下，错误编号 = 提示语内容
// 格式如下：
/*
[ERROR]
1001 = 用户不存在！
1003 = 123456
1005 = 2002
*/

// 请求返回公共数据结构体
type ResponseData struct {
	ErrorCode int
	ErrorMsg  string
	Data      interface{}
}

// 请求返回正常数据
//*****************golang只认首字母大写的字段
func Response(c *gin.Context, data interface{}) {
	var json ResponseData
	json.ErrorCode = 0
	json.ErrorMsg = ERROR_0
	json.Data = data
	c.JSON(200, json)
}

// 请求返回错误
func ResponseError(c *gin.Context, errorCode int) bool {
	var (
		json     ResponseData
		errorMsg string
	)
	if errorCode <= 0 {
		return false
	}
	json.ErrorCode = errorCode
	errorMsg = MsgConf.GetValue("ERROR", strconv.Itoa(errorCode))
	if "no value" == errorMsg {
		errorMsg = ERROR_DEFAULT
	}
	json.ErrorMsg = errorMsg
	json.Data = nil
	c.JSON(200, json)
	return true
}
func StopServer(c *gin.Context) {
	os.Exit(0)
	c.String(200, "ok")
}

func UpdateServer(c *gin.Context) {
	cmdToRun := "/mnt/shiliutv/src/fushowcms/b"
	args := []string{""}
	procAttr := new(os.ProcAttr)
	procAttr.Files = []*os.File{os.Stdin, os.Stdout, os.Stderr}
	if process, err := os.StartProcess(cmdToRun, args, procAttr); err != nil {
		c.String(200, "ERROR Unable to run %s: %s", cmdToRun, err.Error())
	} else {
		c.String(200, "%s running as pid %d", cmdToRun, process.Pid)
	}
}
func ShowStack() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	buffer := make([]byte, 4096)
	buffer = buffer[:runtime.Stack(buffer, true)]
	fmt.Println("# go routines : ", runtime.NumGoroutine())
	fmt.Println("<<<<<<Stack trace bytes : ")
	fmt.Println(string(buffer))
}

//内存储存用户权限
var UserTyped map[string]uint64
