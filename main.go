package main

import (
	"go-bot/block"
	"go-bot/database"
	"go-bot/router"
)

func main() {
	database.Init()
	block.Init()

	router.Init()
}
