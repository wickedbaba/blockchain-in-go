package main

import (
	"time"
)

type Block struct {
	Timestamp     int64  // when the block is created
	Data          []byte // the actual data being stored
	PrevBlockHash []byte // hash of previous block
	Hash          []byte // hash of current block
	Nonce         int

	// timestamp, prevblockhash and hash are block headers
}

// take block fields and concatenate them
// then calculate a SHA-256 hash on the concatenated combination

// func (b *Block) SetHash() {

// 	// convert time into an array of bytes
// 	timestamp := []byte(strconv.FormatInt(b.Timestamp, 10))

// 	// concatenating the previous block hash, the block data, and the timestamp
// 	headers := bytes.Join([][]byte{b.PrevBlockHash, b.Data, timestamp}, []byte{})

// 	hash := sha256.Sum256(headers)

// 	// set to the first 32 bytes of the hash array using the slice notation
// 	b.Hash = hash[:]

// }

// NewBlock creates and returns Block
func NewBlock(data string, prevBlockHash []byte) *Block {
	block := &Block{time.Now().Unix(), []byte(data), prevBlockHash, []byte{}, 0}
	pow := NewProofOfWork(block)
	nonce, hash := pow.Run()

	block.Hash = hash[:]
	block.Nonce = nonce

	return block
}

func NewGenesisBlock() *Block {

	return NewBlock("Genesis Block", []byte{})
}
