package main

import (
	"go-bot/block"
	"go-bot/database"
	"go-bot/router"

	"go-bot/ws"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	database.Init()
	go sendMsg()
	block.Init()
	router.Init()
}

func sendMsg() {
	p := ws.GetInstance()
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
