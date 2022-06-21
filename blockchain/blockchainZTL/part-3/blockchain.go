package main

import (
	"fmt"
	"github.com/boltdb/bolt"
	"log"
)

const dbFile = "blockchain.db"
const blockBucket = "blocks"

// Blockchain tip 这个词本身有事物尖端或尾部的意思，这里指的是存储最后一个块的哈希
// 在链的末端可能出现短暂分叉的情况，所以选择 tip 其实也就是选择了哪条链
// db 存储数据库连接
type Blockchain struct {
	tip []byte
	db  *bolt.DB
}
type BlockchainIterator struct {
	currentHash []byte
	db          *bolt.DB
}

func (blockchain *Blockchain) AddBlock(data string) {
	var lastHash []byte

	// 首先获取最后一个块的哈希用于生成新块的哈希
	err := blockchain.db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(blockBucket))
		lastHash = bucket.Get([]byte("l"))

		return nil
	})
	if err != nil {
		log.Panic(err)
	}

	newBlock := NewBlock(data, lastHash)

	err = blockchain.db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(blockBucket))
		err := bucket.Put(newBlock.Hash, newBlock.Serialize())
		if err != nil {
			log.Panic(err)
		}
		err = bucket.Put([]byte("l"), newBlock.Hash)
		if err != nil {
			log.Panic(err)
		}
		blockchain.tip = newBlock.Hash
		return nil
	})

}

func NewBlockchain() *Blockchain {
	var tip []byte
	//打开一个boltDB
	db, err := bolt.Open(dbFile, 0600, nil)
	if err != nil {
		log.Panic(err)
	}

	err = db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(blockBucket))

		// 如果数据库中不存在区块链就创建一个，否则直接读取最后一个块的哈希
		if bucket == nil {
			fmt.Println("No existing blockchain found. Creating a new one...")
			genesis := NewGenesisBlock()

			bucket, err := tx.CreateBucket([]byte(blockBucket))
			if err != nil {
				log.Panic(err)
			}

			err = bucket.Put(genesis.Hash, genesis.Serialize())
			if err != nil {
				log.Panic(err)
			}

			err = bucket.Put([]byte("l"), genesis.Hash)
			if err != nil {
				log.Panic(err)
			}
			tip = genesis.Hash
		} else {
			tip = bucket.Get([]byte("l"))
		}

		return nil
	})

	if err != nil {
		log.Panic(err)
	}
	blockchain := Blockchain{tip, db}

	return &blockchain
}

func (blockchain *Blockchain) Iterator() *BlockchainIterator {
	bci := &BlockchainIterator{blockchain.tip, blockchain.db}

	return bci
}

func (i *BlockchainIterator) Next() *Block {
	var block *Block

	err := i.db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(blockBucket))
		encodedBlock := bucket.Get(i.currentHash)
		block = DeserializeBlock(encodedBlock)

		return nil
	})
	if err != nil {
		log.Panic(err)
	}
	i.currentHash = block.PrevBlockHash

	return block
}
