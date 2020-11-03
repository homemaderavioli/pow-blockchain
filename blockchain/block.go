package blockchain

import (
	"encoding/hex"
	"fmt"
	"strings"
)

type Block struct {
	BlockNumber       int
	Nonce             int
	Message           Message
	PreviousBlockHash []byte
}

const numberOfLeadingZeros int = 4
const maxNonceLimit int = 9999999

func genesisBlock(message Message) *Block {
	block := newBlock(0, message, nil)
	return block
}

func newBlock(number int, message Message, hash []byte) *Block {
	block := &Block{
		BlockNumber:       number,
		Nonce:             0,
		Message:           message,
		PreviousBlockHash: hash,
	}
	block.findNonce()
	return block
}

func (b *Block) findNonce() bool {
	for nonce := 0; nonce < maxNonceLimit; nonce++ {
		if b.tryNonce(nonce) {
			return true
		}
	}
	return false
}

func (b *Block) tryNonce(nonce int) bool {
	b.Nonce = nonce
	return b.validBlock()
}

func (b *Block) validBlock() bool {
	bSerialized := blockSerialized(b)
	if blockHashHasLeadingZeros(bSerialized) {
		return true
	}
	return false
}

func blockSerialized(b *Block) []byte {
	return []byte(fmt.Sprintf("%v", b))
}

func blockHashHasLeadingZeros(block []byte) bool {
	blockHash := hash(block)
	hexBlockHash := hex.EncodeToString(blockHash)
	leadingNCharacters := hexBlockHash[:numberOfLeadingZeros]
	if leadingNCharacters == strings.Repeat("0", numberOfLeadingZeros) {
		return true
	}
	return false
}

func (b *Block) blockHash() []byte {
	bSerialized := blockSerialized(b)
	hash := hash(bSerialized)
	return hash
}

func (b *Block) prettyPrint() {
	fmt.Printf("==== BLOCK %d ====\n", b.BlockNumber)
	fmt.Printf("Nonce: %d\n", b.Nonce)
	fmt.Printf("Data: %s\n", b.Message.Message)
	fmt.Printf("Previous block hash: %s\n", hex.EncodeToString(b.PreviousBlockHash))
	fmt.Println("=================")
}
