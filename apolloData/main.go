// ApolloDemo project main.go
package main

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"

	"github.com/garyburd/redigo/redis"
)

var Mapollo []Apollo

// 生成连接池
var pool = newPoolUser()

func main() {

	WriteApollo() //写入apollo配置文件

}

// 重写生成连接池方法
func newPoolUser() *redis.Pool {
	a := 80
	b := 1200
	return &redis.Pool{
		MaxIdle:   a,
		MaxActive: b,
		Dial: func() (redis.Conn, error) {

			c, err := redis.Dial("tcp", "127.0.0.1:6379",redis.DialPassword("Fushow666@"))
			if err != nil {
				panic(err.Error())
			}
			return c, err
		},
	}
}

//写入apollo配置文件
func WriteApollo() {
	fmt.Println("strat")
	DoApolloMain()
	ticker := time.NewTicker(10*60 * time.Second)
	for {
		<-ticker.C
		DoApolloMain()
		fmt.Println("strat2")
	}
}

type Apollo struct {
	UserName string
	PassWord []byte
}

//写入apollo配置文件操作
func DoApolloMain() {
	//获取数据
	list := GetApolloData()
	for m := 0; m < len(list); m++ {
		fmt.Println("m", list[m])
		isflag := SetApolloIM(&list[m], int64(m))
		if !isflag {
			fmt.Println("********************cuowu")
		}
	}

	fname := "/mnt/mybroker/etc/users.properties"
	f, err := os.OpenFile(fname, os.O_CREATE|os.O_RDWR|os.O_APPEND, os.ModeAppend|os.ModePerm)

	defer f.Close()
	if err != nil {
		fmt.Println("文件未找到", f, err)
		return
	}

	//Redis 获取数据
	var addData []Apollo //数组  配置数据
	for s := 0; s < len(list); s++ {
		mData := GetApolloIM(int64(s))
		addData = append(addData, mData)
	}

	//写入配置文件中
	var myData string
	for i := 0; i < len(addData); i++ {
		fmt.Println("i", i)
		if len(addData) == 0 {
			fmt.Println("Mapollo无数据")
			break
		}

		//先剪裁配置数据 1206字节
		err = os.Truncate(fname, 1206)
		if err != nil {
			fmt.Println("裁剪失败")
			continue
		}

		username := addData[i].UserName
		//[]byte转换统一格式
		pwd := ByteToHex(addData[i].PassWord)

		//md5加密
		md5Ctx := md5.New()
		md5Ctx.Write([]byte(pwd))
		cipherStr := md5Ctx.Sum(nil)

		md5pwd := hex.EncodeToString(cipherStr)
		//第一条数据换行
		if i == 0 {
			myData = myData + "\n" + username + "=" + md5pwd + "\n"
		} else {
			myData = myData + username + "=" + md5pwd + "\n"
		}

	}

	//写入操作
	f.WriteString(myData)

}

//apollo数据
func GetApolloData() []Apollo {
	var list []Apollo

	if len(Mapollo) > 0 {
		Mapollo = nil
	}

	//循环次数   apollo用户组
	//username  :

	for i := 0; i < 10; i++ {

		var data Apollo
		//username = "admin"

		username := "admin" + strconv.FormatInt(int64(i), 10)

		number := randInt(6, 12)

		pwd := GetRandomString(number)
		//加密AES   秘钥
		aesPwd, err := AesEncrypt([]byte(pwd), []byte("apollofushowcms1"))
		fmt.Println("pwd", aesPwd)
		if err != nil {
			panic(err)
		}

		data.UserName = username

		data.PassWord = aesPwd

		list = append(list, data)

	}
	return list

}

//最大密码值与最小密码值
func randInt(min int, max int) int {
	return min + rand.Intn(max-min)
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

//将IM信息存入Radis
func SetApolloIM(b *Apollo, id int64) bool {
	rs := pool.Get()
	defer rs.Close()
	var key string
	key = "apollo" + strconv.FormatInt(id, 10) //key是固定值
	if err := rs.Send("HMSET", redis.Args{}.Add(key).AddFlat(b)...); err != nil {
		panic(err)
		fmt.Println("err", err)
	}
	//rs.Send("Expire", key, comm.GetConfig("SESSION", "apollo_time"))
	return true
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
