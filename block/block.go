package block

import (
	"bytes"
	"crypto/sha256"
	"strconv"
	"time"
)

// Block 区块
type Block struct {
	Timestamp     int64
	Data          []byte
	PrevBlockHash []byte
	Hash          []byte
}

// Blockchain 区块链
type Blockchain struct {
	blocks []*Block
}

// 单例 。
var instance *Blockchain

// GetInstance .
func GetInstance() *Blockchain {
	if instance == nil {
		instance = newBlockChain()
	}
	return instance
}

// NewBlock .
func newBlock(data string, prevBlockHash []byte) *Block {
	block := &Block{time.Now().Unix(), []byte(data), prevBlockHash, []byte{}}
	block.setHash()
	return block
}

// 创世块新建
func newGenesisBlock() *Block {
	return newBlock("Genesis Block", []byte{})
}

// 创建带有创世块的区块链
func newBlockChain() *Blockchain {
	return &Blockchain{[]*Block{newGenesisBlock()}}
}

func (b *Block) setHash() {
	timestamp := []byte(strconv.FormatInt(b.Timestamp, 10))
	headers := bytes.Join([][]byte{b.PrevBlockHash, b.Data, timestamp}, []byte{})
	hash := sha256.Sum256(headers)
	b.Hash = hash[:]
}

// AddBlock .
func (bc *Blockchain) AddBlock(data string) {
	preBlock := bc.blocks[len(bc.blocks)-1]
	newBlock := newBlock(data, preBlock.Hash)
	bc.blocks = append(bc.blocks, newBlock)
}

// GetBlockByIndex .
func (bc *Blockchain) GetBlockByIndex(index int) *Block {
	return bc.blocks[index-1]
}
