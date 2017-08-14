package control

/*
 * Copyright (c) 2016—2017 www.fushow.cn, All rights reserved.
 */
import (
	"fmt"
	"fushowcms/comm"
	m "fushowcms/models"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
)

func SetUp(c *gin.Context) {
	ip := c.ClientIP()
	phone := c.PostForm("Phone")
	db_host := c.PostForm("DbHost")
	db_port := c.PostForm("DbPort")
	db_user := c.PostForm("DbUser")
	db_pass := c.PostForm("DbPass")
	db_dbname := c.PostForm("DbDbname")
	db_selectdb := c.PostForm("DbSelectdb")
	redis_dial := c.PostForm("RedisDial")
	redis_deal := c.PostForm("RedisDeal")
	redis_pass := c.PostForm("RedisPass")
	redis_key := c.PostForm("RedisKey")
	db_user = comm.SetAesValue(db_user, "fushow.cms")
	db_pass = comm.SetAesValue(db_pass, "fushow.cms")
	redis_pass = comm.SetAesValue(redis_pass, "fushow.cms")
	userFile := "./conf/app.conf"
	fout, err := os.Create(userFile)
	defer fout.Close()
	if err != nil {
		fmt.Println(userFile, err)
		return
	}
	fout.WriteString("[DB]\n")
	fout.WriteString("host = " + db_host + "\n")
	fout.WriteString("port = " + db_port + "\n")
	fout.WriteString("user = " + db_user + "\n")
	fout.WriteString("pass = " + db_pass + "\n")
	fout.WriteString("dbname = " + db_dbname + "\n")
	fout.WriteString("selectdb = " + db_selectdb + "\n")
	fout.WriteString("\n\n")
	fout.WriteString("[REDIS]\n")
	fout.WriteString("dial = " + redis_dial + "\n")
	fout.WriteString("deal = " + redis_deal + "\n")
	fout.WriteString("pass = " + redis_pass + "\n")
	fout.WriteString("key = " + redis_key + "\n")
	fout.WriteString("\n\n")
	fout.WriteString("[Fund]\n")
	fout.WriteString("smoney = 5000000\n")
	fout.WriteString("cmoney = 0\n")
	fout.WriteString("user_money= 1000000\n")
	fout.WriteString("\n\n")
	fout.WriteString("[GIFT]\n")
	fout.WriteString("anchor_times = 2\n")
	fout.WriteString("sys_times = 4\n")
	fout.WriteString("\n\n")
	fout.WriteString("[POOL]\n")
	fout.WriteString("max_idle = 80\n")
	fout.WriteString("max_active = 1200\n")
	fout.WriteString("\n\n")
	fout.WriteString("[SESSION]\n")
	fout.WriteString("expiration_time = 3600\n")
	fout.WriteString("broadcast_time = 120\n")
	fout.WriteString("cookie_name =	www.fushow.cn\n")
	fout.WriteString("\n\n")
	fout.WriteString("[SMS]\n")
	fout.WriteString("account = \n")
	fout.WriteString("password = \n")
	fout.WriteString("url = https://106.ihuyi.com/webservice/sms.php?method=Submit\n")
	fout.WriteString("\n\n")
	fout.WriteString("[Root]\n")
	fout.WriteString("pass = fushow\n")
	fout.WriteString("\n\n")
	fout.WriteString("[FS]\n")
	fout.WriteString("live1 = live0.fushow.cn\n")
	fout.WriteString("live2 = live1.fushow.cn\n")
	fout.WriteString("live3 = live2.fushow.cn\n")
	fout.WriteString("live4 = live3.fushow.cn\n")
	fout.WriteString("live5 = live4.fushow.cn\n")
	fout.WriteString("live6 = live5.fushow.cn\n")
	fout.WriteString("live7 = live6.fushow.cn\n")
	fout.WriteString("live8 = live7.fushow.cn\n")
	fout.WriteString("live9 = live8.fushow.cn\n")
	fout.WriteString("live10 = live9.fushow.cn\n")
	fout.WriteString("time = 1\n")
	v := url.Values{}
	v.Add("ip", ip)
	v.Add("phone", phone)
	respet, _ := http.NewRequest("POST", "http://tv.fushow.cn/page/addsetUp", strings.NewReader(v.Encode()))
	respet.Header.Set("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")
	client := &http.Client{}
	fmt.Println("respet", respet)
	response, _ := client.Do(respet)
	fmt.Println("response", response)
	defer response.Body.Close()
	comm.Response(c, nil)
}

func AddSetUp(c *gin.Context) {
	var (
		su m.SetUp
	)
	su.Ip = c.PostForm("ip")
	su.Phone = c.PostForm("phone")
	boo := su.AddSetUp()
	fmt.Println(boo)
}
