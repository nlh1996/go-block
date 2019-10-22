package main

import (
	"go-bot/block"
	"go-bot/database"
	"go-bot/router"
	"go-bot/ws"
	"strconv"
	"time"
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
		if len(p.Pool) != 0 {
			p.Pool[0].WriteMessage([]byte(strconv.Itoa(i)))
			i++
		}
		time.Sleep(3 * time.Second)
	}
}
