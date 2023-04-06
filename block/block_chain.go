package block

import (
	"bitcoin/constant"
	"fmt"
	"github.com/boltdb/bolt"
	"log"
)

type BlockChain struct {
	tip []byte
	db  *bolt.DB
}

type BlockChainIterator struct {
	currentHash []byte
	db          *bolt.DB
}

type Chain struct {
	Blocks []*Block
}

// AddBlock
// @Description: 添加区块
// @receiver chain
// @param data
func (chain *BlockChain) AddBlock(data string) {
	var lastHash []byte
	// 获取最后一个块的hash
	err := chain.db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(constant.BLOCK_BUCKET))
		lastHash = bucket.Get([]byte("1"))
		return nil
	})
	if err != nil {
		log.Panic(err)
	}

	newBlock := NewBlock(data, lastHash)

	chain.db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(constant.BLOCK_BUCKET))
		err = bucket.Put(newBlock.Hash, newBlock.Serialize())
		if err != nil {
			log.Panic(err)
		}
		err = bucket.Put([]byte("1"), newBlock.Hash)
		if err != nil {
			log.Panic(err)
		}
		chain.tip = newBlock.Hash
		return nil
	})
}

// NewBlockChain
// @Description: 新建区块链
// @return *BlockChain
func NewBlockChain() *BlockChain {
	var tip []byte
	db, err := bolt.Open(constant.DB_File, 0600, nil)
	if err != nil {
		log.Panic(err)
	}
	db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(constant.BLOCK_BUCKET))

		if bucket == nil {
			fmt.Printf("no existing blockchain found creating a new one...\n")
			genesis := NewGenesisBlock()

			bucket, err = tx.CreateBucket([]byte(constant.BLOCK_BUCKET))
			if err != nil {
				log.Panic(err)
			}
			err = bucket.Put(genesis.Hash, genesis.Serialize())
			if err != nil {
				log.Panic(err)
			}
			err = bucket.Put([]byte("1"), genesis.Hash)
			if err != nil {
				log.Panic(err)
			}
			tip = genesis.Hash
		} else {
			tip = bucket.Get([]byte("1"))
		}
		return nil
	})
	if err != nil {
		log.Panic(err)
	}
	bc := BlockChain{tip: tip, db: db}
	return &bc
}

func (chain *BlockChain) Iterator() *BlockChainIterator {
	chainIterator := BlockChainIterator{chain.tip, chain.db}
	return &chainIterator
}

func (i *BlockChainIterator) Next() *Block {
	var block *Block
	err := i.db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(constant.BLOCK_BUCKET))
		encoderBlock := bucket.Get(i.currentHash)
		block = Deserialize(encoderBlock)
		return nil
	})
	if err != nil {
		log.Panic(err)
	}
	i.currentHash = block.PrevBlockHash
	return block
}

func (chain *BlockChain) CloseDB() error {
	err := chain.db.Close()
	return err
}
