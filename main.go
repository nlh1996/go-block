package main

import (
	"fmt"
	"go-bot/block"
	"go-bot/conn"
	"go-bot/dao"
	"go-bot/router"
	"strconv"
)

func main() {
	conn.Init()
	router.Init()
	
	bc := block.GetInstance()
	bc.AddBlock("Send 1 BTC to Ivan")
	bc.AddBlock("Send 2 BTC to Ivan")
	bc.AddBlock("Send 3 BTC to Ivan")
	for _, bk := range bc.GetBlockchain() {
		fmt.Printf("Prev.hash: %x\n", bk.PrevBlockHash)
		fmt.Printf("Data: \"%s\"\n", bk.Data)
		fmt.Printf("Hash: %x\n", bk.Hash)
		pow := block.NewProofOfWork(bk)
		fmt.Printf("PoW: %s\n\n", strconv.FormatBool(pow.Validate()))
		dao.InsertOne(bk)
	}
}
