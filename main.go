package main

import (
	"go-bot/conn"
	"go-bot/router"
)

func main() {
	conn.Init()
	router.Init()
}
