package main

import (
	"fmt"
	"strconv"
)

func main() {
	//part-2
	bc := NewBlockChain()

	//bc.AddBlock("Send 1 BTC to Ivan")
	//bc.AddBlock("Send 2 more BTC to Ivan")

	for i := 1; i <= 10; i++ {
		si := strconv.Itoa(i)
		data := "send " + si + " BTC to ZT"
		bc.AddBlock(data)
	}

	for _, block := range bc.blocks {
		fmt.Printf("Prev hash: %x\n", block.PrevBlockHash)
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)
		pow := NewProofOfWork(block)
		fmt.Printf("PoW: %s\n", strconv.FormatBool(pow.Validate()))
		fmt.Println()
	}

}
