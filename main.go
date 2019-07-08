package main

import (
	"fmt"
	"go-bot/block"
)

func main() {
	bc := block.GetInstance()
	bc.AddBlock("Send 1 BTC to Ivan")
	bc.AddBlock("Send 2 more BTC to Ivan")

	fmt.Printf("Prev.hash: %x\n", bc.GetBlockByIndex(2).PrevBlockHash)
	fmt.Printf("Data: %s\n", bc.GetBlockByIndex(2).Data)
	fmt.Printf("Hash: %x\n", bc.GetBlockByIndex(2).Hash)
}
