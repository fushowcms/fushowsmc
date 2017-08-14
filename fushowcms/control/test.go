package control

/*
 * Copyright (c) 2016—2017 www.fushow.cn, All rights reserved.
 */
import (
	"bytes"
	"fmt"
	"fushowcms/models"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func OverWeb(c *gin.Context) {
	if !models.OverWeb() {
		c.JSON(200, "fail")
		return
	}
	c.JSON(200, "ok")
}

type TestMM struct {
	Uid      string
	GiftId   string
	AnchorId string
	Number   string
}

//测试
func TestGiftNumber(c *gin.Context) {
	client := &http.Client{}
	uid := c.PostForm("id")
	var nowm TestMM
	nowm.GiftId = "1"
	nowm.Uid = uid
	nowm.Number = "1"
	nowm.AnchorId = "1"
	b := bytes.Buffer{}
	b.WriteString("GiftId=1&Number=1&AnchorId=1&UID=")
	b.WriteString(uid)
	body := ioutil.NopCloser(strings.NewReader(b.String())) //把form数据编下码
	for i := 0; i < 100; i++ {
		req, _ := http.NewRequest("POST", "http://192.168.1.200/page/givegiftnumadd", body)
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value") //这个一定要加，不加form的值post不过去，被坑了两小时
		fmt.Printf("%+v\n", req)                                                         //看下发送的结构
		resp, err := client.Do(req)                                                      //发送
		defer resp.Body.Close()                                                          //一定要关闭resp.Body
		data, _ := ioutil.ReadAll(resp.Body)
		fmt.Println(string(data), err)
	}
}
