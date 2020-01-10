package main

import (
	"encoding/base64"
	"fmt"
	"go-bot/block"
	"go-bot/database"
	"go-bot/router"
	"go-bot/ws"
	"log"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	database.Init()
	//go sendMsg()
	block.Init()
	router.Init()
	// uncode()
}

func sendMsg() {
	p := ws.GetConnPool()
	var i int
	for {
		for _, v := range p.Pool {
			if !v.IsClosed {
				v.Send(gin.H{"code": strconv.Itoa(i)})
				i++
			}
		}
		time.Sleep(3 * time.Second)
	}
}

func uncode() {
	var enc = base64.StdEncoding
	res, err := enc.DecodeString("eyJ0IjoxLCJjIjoiOTk4IiwiZCI6IjE1Nzg2MzAxNDU0NDMifQ0=")
	if err != nil {
		log.Println(err.Error())
	}
	fmt.Println(string(res))
}
