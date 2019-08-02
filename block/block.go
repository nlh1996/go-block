package block

import (
	"go-bot/dao"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

// Block 区块
type Block struct {
	Timestamp     int64
	Data          []byte
	PrevBlockHash []byte
	Hash          []byte
	Nonce         int
}

// Blockchain 区块链
type Blockchain struct {
	blocks []*Block
}

// BlockchainIterator 区块链迭代器
type BlockchainIterator struct {
	currentHash []byte
	collection  *mongo.Collection
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
	block := &Block{time.Now().Unix(), []byte(data), prevBlockHash, []byte{}, 0}
	pow := NewProofOfWork(block)
	nonce, hash := pow.Run()

	block.Hash = hash[:]
	block.Nonce = nonce

	if nonce < 10000000 {
		// 新块存储到数据库
		dao.InsertOne(block)
	}
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

// func (b *Block) setHash() {
// 	timestamp := []byte(strconv.FormatInt(b.Timestamp, 10))
// 	headers := bytes.Join([][]byte{b.PrevBlockHash, b.Data, timestamp}, []byte{})
// 	hash := sha256.Sum256(headers)
// 	b.Hash = hash[:]
// }

// AddBlock .
func (bc *Blockchain) AddBlock(data string) int {
	preBlock := bc.blocks[len(bc.blocks)-1]
	newBlock := newBlock(data, preBlock.Hash)
	bc.blocks = append(bc.blocks, newBlock)

	return len(bc.blocks) - 1
}

// GetBlockByIndex .
func (bc *Blockchain) GetBlockByIndex(index int) *Block {
	if index < len(bc.blocks) && index >= 0 {
		return bc.blocks[index]
	}
	return nil
}

// GetBlockchain .
func (bc *Blockchain) GetBlockchain() []*Block {
	return bc.blocks
}
