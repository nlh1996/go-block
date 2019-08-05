package block

import (
	"errors"
	"go-bot/conn"
	"go-bot/dao"
	"go-bot/utils"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/mgo.v2/bson"
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

// 单例 。
var instance *Blockchain

// GetInstance .
func GetInstance() *Blockchain {
	if instance == nil {
		instance = newBlockChain()
	}
	return instance
}

// Init 初始化区块链.
func Init() {
	block := &Block{}
	opts := options.FindOne()
	opts.SetSort(bson.M{"timestamp": -1})
	conn.GetCollection().FindOne(utils.GetCtx(), bson.M{}, opts).Decode(block)
	if block == nil {
		instance = nil
		return
	}
	instance = &Blockchain{}
	instance.blocks = append(instance.blocks, block)
}

// NewBlock .
func newBlock(data string, prevBlockHash []byte) *Block {
	block := &Block{time.Now().Unix(), []byte(data), prevBlockHash, []byte{}, 0}
	// 工作量证明
	pow := NewProofOfWork(block)
	nonce, hash := pow.Run()
	block.Hash = hash[:]
	block.Nonce = nonce
	// 是否证明通过
	if pow.Validate() {
		// 新块存储到数据库
		dao.InsertOne(block)
		return block
	}
	return nil
}

// 创世块新建
func newGenesisBlock() *Block {
	bk := newBlock("Genesis Block", []byte{})
	if bk == nil {
		log.Fatal("创世块创建失败！")
	}
	return bk
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
func (bc *Blockchain) AddBlock(data string) error {
	preBlock := bc.blocks[len(bc.blocks)-1]
	newBlock := newBlock(data, preBlock.Hash)
	if newBlock != nil {
		bc.blocks = append(bc.blocks, newBlock)
		return nil
	}
	return errors.New("加入失败")
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

// BlockchainIterator 区块链迭代器
type BlockchainIterator struct {
	currentHash []byte
	collection  *mongo.Collection
}

// NewIterator .
func (bc *Blockchain) NewIterator() *BlockchainIterator {
	bci := &BlockchainIterator{}
	bci.collection = conn.GetCollection()

	return bci
}

// Next .
func (bci *BlockchainIterator) Next() *Block {
	block := &Block{}
	ctx := utils.GetCtx()
	if bci.currentHash == nil {
		opts := options.FindOne()
		opts.SetSort(bson.M{"timestamp": -1})
		bci.collection.FindOne(ctx, bson.M{}, opts).Decode(block)
	} else {
		bci.collection.FindOne(ctx, bson.M{"hash": bci.currentHash}).Decode(block)
	}
	bci.currentHash = block.PrevBlockHash
	return block
}
