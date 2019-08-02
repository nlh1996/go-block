package block

import (
	"bytes"
	"crypto/sha256"
	"math/big"
	"strconv"
)

const targetBits = 24

// ProofOfWork .
type ProofOfWork struct {
	block  *Block
	target *big.Int
}

// NewProofOfWork .
func NewProofOfWork(b *Block) *ProofOfWork {
	target := big.NewInt(1)
	target.Lsh(target, uint(256-targetBits))
	pow := &ProofOfWork{b, target}
	return pow
}

func (pow *ProofOfWork) prepareData(nonce int) []byte {
	data := bytes.Join(
		[][]byte{
			pow.block.PrevBlockHash,
			pow.block.Data,
			IntToHex(pow.block.Timestamp),
			IntToHex(int64(targetBits)),
			IntToHex(int64(nonce)),
		},
		[]byte{},
	)

	return data
}

// IntToHex .
func IntToHex(data int64) (arr []byte) {
	return []byte(strconv.FormatInt(data, 10))
}

// Run .
func (pow *ProofOfWork) Run() (int, []byte) {
	var hashInt big.Int
	var hash [32]byte
	nonce := 0

	for nonce < 10000000 {
		data := pow.prepareData(nonce)
		hash = sha256.Sum256(data)
		hashInt.SetBytes(hash[:])

		if hashInt.Cmp(pow.target) == -1 {
			break
		}
		nonce++
	}
	return nonce, hash[:]
}

// Validate .
// func (pow *ProofOfWork) Validate() bool {
// 	var hashInt big.Int

// 	data := pow.prepareData(pow.block.Nonce)
// 	hash := sha256.Sum256(data)
// 	hashInt.SetBytes(hash[:])

// 	isValid := hashInt.Cmp(pow.target) == -1

// 	return isValid
// }
