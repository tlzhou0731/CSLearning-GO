package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"math"
	"math/big"
)

// 难度值，这里表示哈希的前 24 位必须是 0
const targetBits = 24

const maxNonce = math.MaxInt64

// ProofOfWork 每个块的工作量都必须要证明，所有有个指向 Block 的指针
// target 是目标，我们最终要找的哈希必须要小于目标
type ProofOfWork struct {
	block  *Block
	target *big.Int
}

// NewProofOfWork target 等于 1 左移 256 - targetBits 位
func NewProofOfWork(b *Block) *ProofOfWork {
	target := big.NewInt(1)
	target.Lsh(target, uint(256-targetBits))

	pow := &ProofOfWork{b, target}

	return pow
}

// 工作量证明用到的数据有：PrevBlockHash, Data, Timestamp, targetBits, nonce
// ProofOfWork根据nonce, 即count来准备数据
func (pow *ProofOfWork) prepareData(nonce int) []byte {
	data := bytes.Join(
		[][]byte{
			pow.block.PrevBlockHash,
			pow.block.Data,
			IntToHex(pow.block.Timestamp),
			IntToHex(int64(targetBits)),
			IntToHex(int64(nonce)),
		},
		[]byte{},
	)

	return data
}

// Run 工作量证明的核心就是寻找有效哈希
func (pow *ProofOfWork) Run() (int, []byte) {
	var hashInt big.Int
	var hash [32]byte
	nonce := 0

	fmt.Printf("Mining the block containing \"%s\"\n", pow.block.Data)
	for nonce < maxNonce {
		//pow准备根据nonce准备数据
		data := pow.prepareData(nonce)
		//将数据进行hash
		hash = sha256.Sum256(data)
		//将hash转换为整数
		hashInt.SetBytes(hash[:])
		//将hash 转换的整数与这个pow的target比较
		if hashInt.Cmp(pow.target) == -1 {
			fmt.Printf("\r%x", hash)
			break
		} else {
			nonce++
		}
	}
	fmt.Print("\n\n")
	//返回这次挖到的nonce和hash值
	return nonce, hash[:]
}

// Validate 验证工作量，只要哈希小于目标就是有效工作量
func (pow *ProofOfWork) Validate() bool {
	var hashInt big.Int

	data := pow.prepareData(pow.block.Nonce)
	hash := sha256.Sum256(data)
	hashInt.SetBytes(hash[:])

	isValid := hashInt.Cmp(pow.target) == -1

	return isValid
}
