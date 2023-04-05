package main

import (
	"bytes"
	"crypto/sha256"
	"strconv"
	"time"
)

type Block struct {
	Timestamp     int64  // when the block is created
	Data          []byte // the actual data being stored
	PrevBlockHash []byte // hash of previous block
	Hash          []byte // hash of current block

	// timestamp, prevblockhash and hash are block headers
}

// take block fields and concatenate them
// then calculate a SHA-256 hash on the concatenated combination

func (b *Block) SetHash() {

	// convert time into an array of bytes
	timestamp := []byte(strconv.FormatInt(b.Timestamp, 10))

	// concatenating the previous block hash, the block data, and the timestamp
	headers := bytes.Join([][]byte{b.PrevBlockHash, b.Data, timestamp}, []byte{})

	hash := sha256.Sum256(headers)

	// set to the first 32 bytes of the hash array using the slice notation
	b.Hash = hash[:]

}

func NewBlock(data string, prevBlockHash []byte) *Block {
	block := &Block{time.Now().Unix(), []byte(data), prevBlockHash, []byte{}}
	block.SetHash()
	return block

}

func NewGenesisBlock() *Block {

	return NewBlock("Genesis Block", []byte{})
}
