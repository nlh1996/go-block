package main

import (
	"fmt"
	"go-bot/block"
	"go-bot/database"
	"go-bot/router"
	"go-bot/ws"
	"strconv"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	database.Init()
	go sendMsg()
	block.Init()
	router.Init()
	// testMap()
}

func sendMsg() {
	p := ws.GetConnPool()
	var i int
	for {
		p.Lock()
		for _, v := range p.Pool {
			if !v.IsClosed {
				v.Send(gin.H{"code": strconv.Itoa(i)})
				i++
			}
		}
		p.Unlock()
		time.Sleep(3 * time.Second)
	}
}

type box struct {
	m map[int]int
	sync.Mutex
}

func testMap() {
	b := box{}
	b.m = make(map[int]int)
	for i := 0; i < 100000; i++ {
		go func(i int) {
			b.Lock()
			defer b.Unlock()
			b.m[i] = i
		}(i)
	}
	b.Lock()
	for _, v := range b.m {
		fmt.Println(v)
	}
	b.Unlock()
}
