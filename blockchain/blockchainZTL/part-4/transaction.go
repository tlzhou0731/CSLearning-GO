package main

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"encoding/hex"
	"fmt"
	"log"
)

const subsidy = 10

// Transaction 一笔交易由一些输入（input）和输出（output）组合而来：
// 交易仅仅是通过一个脚本（script）来锁定（lock）一些值（value），而这些值只可以被锁定它们的人解锁（unlock）。
type Transaction struct {
	ID   []byte
	VIn  []TXInput
	VOut []TXOutput
}

// TXInput 包含 3 部分
// TxID: 一个交易输入引用了之前一笔交易的一个输出, ID 表明是之前哪笔交易
// VOut: 一笔交易可能有多个输出，VOut 为输出的索引
// ScriptSig: 提供解锁输出 TxID:VOut 的数据
type TXInput struct {
	TxID      []byte
	VOut      int
	ScriptSig string
}

// TXOutput 包含两部分
// Value: 有多少币，就是存储在 Value 里面
// ScriptPubKey: 对输出进行锁定
// 在当前实现中，ScriptPubKey 将仅用一个字符串来代替
type TXOutput struct {
	Value        int
	ScriptPubKey string
}

func (tx *Transaction) SetID() {
	var encoded bytes.Buffer
	var hash [32]byte

	enc := gob.NewEncoder(&encoded)
	err := enc.Encode(tx)
	if err != nil {
		log.Panic(err)
	}
	hash = sha256.Sum256(encoded.Bytes())
	tx.ID = hash[:]
}

// IsCoinbase 判断是否是 coinbase 交易
// coinbase交易：每个挖出这个区块的交易
func (tx Transaction) IsCoinbase() bool {
	return len(tx.VIn) == 1 && len(tx.VIn[0].TxID) == 0 && tx.VIn[0].VOut == -1
}

func NewCoinbaseTX(to, data string) *Transaction {
	if data == "" {
		data = fmt.Sprintf("Reward to %s", to)
	}

	txIn := TXInput{[]byte{}, -1, data}
	txOut := TXOutput{subsidy, to}
	tx := Transaction{nil, []TXInput{txIn}, []TXOutput{txOut}}
	tx.SetID()

	return &tx
}

func (in *TXInput) CanUnlockOutputWith(unlockingData string) bool {
	return in.ScriptSig == unlockingData
}

func (out *TXOutput) CanBeUnlockedWith(unlockingData string) bool {
	return out.ScriptPubKey == unlockingData
}

func NewUTXOTransaction(from, to string, amount int, blockchain *Blockchain) *Transaction {
	var inputs []TXInput
	var outputs []TXOutput

	acc, validOutputs := blockchain.FindSpendableOutputs(from, amount)

	if acc < amount {
		log.Panic("ERROR: Not enough funds")
	}

	//Build a list of inputs
	for txid, outs := range validOutputs {
		txID, err := hex.DecodeString(txid)
		if err != nil {
			log.Panic(err)
		}
		for _, out := range outs {
			input := TXInput{txID, out, from}
			inputs = append(inputs, input)
		}
	}

	//Build a list of outputs
	outputs = append(outputs, TXOutput{amount, to})
	if acc > amount {
		//找零
		outputs = append(outputs, TXOutput{acc - amount, from})
	}
	tx := Transaction{nil, inputs, outputs}
	tx.SetID()

	return &tx
}
