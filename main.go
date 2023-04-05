package main

import (
	"fmt"
)

var pl = fmt.Println
var pf = fmt.Printf

func main() {

	newChain := NewBlockchain()

	newChain.AddBlock("Send 1 Bruh Points to henry")
	newChain.AddBlock("Send 2 more Bruh points to henry")

	for _, block := range newChain.blocks {

		pf("Prev hash: %x\n", block.PrevBlockHash)
		pf("Data: %s\n", block.Data)
		pf("Current Hash: %x\n", block.Hash)
		pl()
	}
}
