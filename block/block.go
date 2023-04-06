package block

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"fmt"
	"strconv"
	"time"
)

// Block
// @Description: 区块
type Block struct {
	// Timestamp
	// @Description: 时间戳
	//
	Timestamp int64
	// Data
	// @Description: 区块数据
	//
	Data []byte
	// PrevBlockHash
	// @Description: 前一个区块hash
	//
	PrevBlockHash []byte
	// Hash
	// @Description: 该区块hash
	//
	Hash []byte
	// Nonce
	// @Description: 遍历次数
	//
	Nonce int
}

func (b *Block) SetHash() {
	timestamp := []byte(strconv.FormatInt(b.Timestamp, 10))
	headers := bytes.Join([][]byte{b.PrevBlockHash, b.Data, timestamp}, []byte{})
	hash := sha256.Sum256(headers)
	b.Hash = hash[:]
}

func NewBlock(data string, prevBlockHash []byte) *Block {
	block := &Block{time.Now().Unix(), []byte(data), prevBlockHash, []byte{}, 0}
	block.SetHash()
	pow := NewProofOfWork(block)
	nounce, hash := pow.Run()

	block.Hash = hash[:]
	block.Nonce = nounce
	return block
}

func NewGenesisBlock() *Block {
	return NewBlock("Genesis Block", []byte{})
}

// Serialize
// @Description: block序列化
// @receiver b
// @return []byte
func (b *Block) Serialize() []byte {
	var result bytes.Buffer
	encoder := gob.NewEncoder(&result)
	err := encoder.Encode(b)
	if err != nil {
		fmt.Errorf("block serialize failed err:%s", err)
	}
	return result.Bytes()
}

func Deserialize(d []byte) *Block {
	var block Block
	decoder := gob.NewDecoder(bytes.NewReader(d))
	err := decoder.Decode(&block)
	if err != nil {
		fmt.Errorf("block deserialize failed err:%s", err)
	}
	return &block
}
