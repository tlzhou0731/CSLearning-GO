package main

import (
	"fmt"
	"strconv"
)

func (cli *CLI) printChain() {
	bc := NewBlockchain("")
	defer bc.db.Close()

	bci := bc.Iterator()
	for {
		block := bci.Next()

		fmt.Printf("Prev hash: %x\n", block.PrevBlockHash)
		fmt.Printf("Hash: %x\n", block.Hash)
		pow := NewProofOfWork(block)
		fmt.Printf("PoW: %s\n", strconv.FormatBool(pow.Validate()))
		fmt.Println()

		if len(block.PrevBlockHash) == 0 {
			break
		}
	}
}

func (cli *CLI) send(from, to string, amount int) {
	bc := NewBlockchain(from)
	defer bc.db.Close()

	tx := NewUTXOTransaction(from, to, amount, bc)
	bc.MineBlock([]*Transaction{tx})
	fmt.Printf("Send Success!")
}

func (cli *CLI) getBalance(address string) {
	blockchain := NewBlockchain(address)
	defer blockchain.db.Close()

	balance := 0
	UTXOs := blockchain.FindUTXO(address)

	for _, out := range UTXOs {
		balance += out.Value
	}
	fmt.Printf("Balance of %s : %d\n", address, balance)
}

func (cli *CLI) createBlockchain(address string) {
	blockchain := CreateBlockchain(address)
	blockchain.db.Close()
	fmt.Println("Create Blockchain Done!")
}
