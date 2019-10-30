package main

import (
	"go-bot/block"
	"go-bot/database"
	"go-bot/router"
	"go-bot/ws"
	"strconv"
	"time"
)

type test struct {
	id int
}

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
				v.WriteMessage([]byte(strconv.Itoa(i)))
				i++
			}
		}
		time.Sleep(3 * time.Second)
	}
}
