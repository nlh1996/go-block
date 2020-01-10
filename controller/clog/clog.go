package clog

import (
	"encoding/base64"
	"fmt"
	"log"
	"strings"

	"github.com/gin-gonic/gin"
)

type logData struct {
	Client      string
	WebSource   string
	DeviceId    string
	Environment string
	CustomInfo  string
	LogPageNo   int
	FileDate    string
	LogArray    string
}

// LogFromClient 收集客户端异常日志.
func LogFromClient(c *gin.Context) {
	data := &logData{}
	if err := c.Bind(data); err != nil {
		log.Println(err)
	}

	var enc = base64.StdEncoding
	arr := strings.Split(data.LogArray, ",")
	for _, v := range arr {
		d := strings.Split(v[16:], "%")
		var decodeData string
		if d[0][len(d[0])-1] == 48 {
			decodeData = d[0] + "="
		}
		if d[0][len(d[0])-1] == 81 {
			decodeData = d[0] + "=="
		}
		res, err := enc.DecodeString(decodeData)
		if err != nil {
			log.Println(err.Error(), string(res), decodeData)
		}
		fmt.Println(string(res))
	}
	c.String(200, "ok")
}
