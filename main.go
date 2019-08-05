package main

import (
	"go-bot/block"
	"go-bot/conn"
	"go-bot/router"
)

func main() {
	conn.Init()
	block.Init()
	
	router.Init()
}
