package main

import (
	"crypto/sha256"
)

type Blocks map[*Block]*Block

type LinkedBlock struct {
	block     *Block
	nextBlock *Block
	prevBlock *Block
}

type Blockchain struct {
	blocks       Blocks
	currentBlock *LinkedBlock
	length       int
}

func New(block *Block) *Blockchain {
	currectBlock := &LinkedBlock{
		block:     block,
		nextBlock: nil,
		prevBlock: nil,
	}

	blockchain := &Blockchain{
		blocks:       make(Blocks),
		currentBlock: currectBlock,
		length:       1,
	}

	blockchain.blocks[block] = block

	return blockchain
}

func (bc *Blockchain) addBlock(block *Block) {
	bc.blocks[block] = block
	topBlock := bc.currentBlock
	topBlock.nextBlock = block

	bc.currentBlock = &LinkedBlock{
		block:     block,
		nextBlock: nil,
		prevBlock: topBlock.block,
	}
	bc.length++
}

func (bc *Blockchain) getTopBlockHash() []byte {
	bSerialized := bc.currentBlock.block.blockSerialized()
	hash := sha256.Sum256(bSerialized)
	return hash[:]
}
