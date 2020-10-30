package main

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

func GenesisBlock(message Message) *Block {
	block := &Block{
		BlockNumber:       0,
		Nonce:             0,
		Message:           message,
		PreviousBlockHash: nil,
	}
	block.findNonce()
	return block
}

func (b *Block) blockHash() []byte {
	bSerialized := blockSerialized(b)
	hash := Sha256Hash(bSerialized)
	return hash
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
	blockHash := Sha256Hash(block)
	hexBlockHash := hex.EncodeToString(blockHash)
	leadingNCharacters := hexBlockHash[:numberOfLeadingZeros]
	if leadingNCharacters == strings.Repeat("0", numberOfLeadingZeros) {
		return true
	}
	return false
}
