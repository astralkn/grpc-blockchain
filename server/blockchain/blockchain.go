package blockchain

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"time"
)

type Block struct {
	Hash          string
	PrevBlockHash string
	Data          string
}

type BlockChain struct {
	Blocks []*Block
}

func (b *Block) setHash() {
	hash := sha256.Sum256([]byte(b.PrevBlockHash + b.Data))
	b.Hash = hex.EncodeToString(hash[:])
}

func NewBlock(data, prevBlockHash string) *Block {
	block := &Block{Data: data, PrevBlockHash: prevBlockHash}
	block.setHash()
	return block
}

func (b *BlockChain) AddBlock(data string) *Block {
	prevBlock := b.Blocks[len(b.Blocks)-1]
	newBlock := NewBlock(data, prevBlock.Hash)
	b.Blocks = append(b.Blocks, newBlock)
	return newBlock
}

func NewBlockChain() *BlockChain {
	return &BlockChain{[]*Block{NewGenesisBlock()}}
}

func NewGenesisBlock() *Block {
	block := NewBlock(fmt.Sprint("Genesis_Block_%s", time.Now().String()), "")
	return block
}
