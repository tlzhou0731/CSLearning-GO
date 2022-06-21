package main

import "fmt"

func main() {
	blockchain := NewBlockchain()

	blockchain.AddBlock("send -from Ivan -to Pedro -amount 6")
	blockchain.AddBlock("send -from Ivan -to TTL -amount 6")

	for _, block := range blockchain.blocks {
		fmt.Printf("prevHash %x\n", block.prevBlockHash)
		fmt.Printf("Data %x\n", block.Data)
		fmt.Printf("Hash %x\n", block.Hash)
		fmt.Println()
	}

}
