package main

type LinkedBlock struct {
	block     *Block
	nextBlock *Block
	prevBlock *Block
}

type Blockchain struct {
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
		currentBlock: currectBlock,
		length:       1,
	}

	return blockchain
}

func (bc *Blockchain) addBlock(block *Block) {
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
	bSerialized := blockSerialized(bc.currentBlock.block)
	hash := Sha256Hash(bSerialized)
	return hash
}

func (bc *Blockchain) validateBlockChain() bool {

	return false
}
