package blockchain

import (
	"bytes"
	"errors"
)

type Blockchain struct {
	blocks []*Block
}

func New(keyPair KeyPair) *Blockchain {
	messageData := []byte("hello world")
	message := buildSignedMessage(keyPair.PrivateKey, messageData)
	block := genesisBlock(message)

	blockchain := initialise(block)

	return blockchain
}

func initialise(block *Block) *Blockchain {
	blocks := make([]*Block, 0)
	blocks = append(blocks, block)
	blockchain := &Blockchain{
		blocks: blocks,
	}

	return blockchain
}

func (bc *Blockchain) length() int {
	return len(bc.blocks)
}

func (bc *Blockchain) addBlock(block *Block) error {
	currentBlock := bc.getTopBlock()

	if validLink(block, currentBlock) == false {
		return errors.New("cannot add block as the link is invalid")
	}

	bc.blocks = append(bc.blocks, block)
	return nil
}

func (bc *Blockchain) getTopBlock() *Block {
	return bc.blocks[len(bc.blocks)-1]
}

func (bc *Blockchain) getTopBlockHash() []byte {
	block := bc.getTopBlock()
	return block.blockHash()
}

func (bc *Blockchain) validateBlockChain() error {
	for i := 1; i < len(bc.blocks); i++ {
		currentBlock := bc.blocks[i]
		previousBlock := bc.blocks[i-1]

		if validLink(currentBlock, previousBlock) == false {
			return errors.New("blockchain has an invalid link")
		}
	}
	return nil
}

func validLink(a *Block, b *Block) bool {
	aHash := a.PreviousBlockHash
	bHash := b.blockHash()
	return bytes.Compare(aHash, bHash) == 0
}
