package control

import (
	"fmt"
	m "fushowcms/models"

	"github.com/gin-gonic/gin"
)

func GetSLTVMessage(c *gin.Context) {
	var (
		sltvm m.SLTVMessage
	)
	sltvm.UrlHost = c.PostForm("url_host")
	sltvm.DbHost = c.PostForm("db_host")
	sltvm.DbPort = c.PostForm("db_port")
	sltvm.DbUser = c.PostForm("db_user")
	sltvm.DbPass = c.PostForm("db_pass")
	sltvm.DbDbname = c.PostForm("db_dbname")
	sltvm.DbSelectdb = c.PostForm("db_selectdb")
	sltvm.RedisDial = c.PostForm("redis_dial")
	sltvm.RedisDeal = c.PostForm("redis_deal")
	sltvm.RedisPass = c.PostForm("redis_pass")
	sltvm.RedisKey = c.PostForm("redis_key")
	boo := sltvm.AddSLTVMessage()
	fmt.Println(boo)
}
