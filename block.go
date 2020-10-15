package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strings"
)

type Message struct {
	Message          []byte
	MessageHash      []byte
	MessageSignature []byte
}

type Block struct {
	BlockNumber       int
	Nonce             int
	Message           Message
	PreviousBlockHash []byte
}

const numberOfLeadingZeros int = 4
const maxNonceLimit int = 9999999

func GenesisBlock() *Block {
	message := buildMessage("James 9001")
	block := &Block{
		BlockNumber:       0,
		Nonce:             0,
		Message:           message,
		PreviousBlockHash: nil,
	}
	block.findNonce()
	return block
}

func buildMessage(message string) Message {
	messagePlainText := []byte(message)
	messageHash := sha256.Sum256(messagePlainText)
	messageSignature := SignMessage(messageHash[:])
	return Message{
		Message:          messagePlainText,
		MessageHash:      messageHash[:],
		MessageSignature: messageSignature,
	}
}

func (b *Block) findNonce() bool {
	for nonce := 0; nonce < maxNonceLimit; nonce++ {
		if b.tryNonce(nonce) {
			fmt.Printf("%d\n", nonce)
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
	bSerialized := b.blockSerialized()
	if blockHashHasLeadingZeros(bSerialized) {
		return true
	}
	return false
}

func (b *Block) blockSerialized() []byte {
	return []byte(fmt.Sprintf("%v", b))
}

func blockHashHasLeadingZeros(block []byte) bool {
	blockHash := sha256.Sum256(block)
	hexBlockHash := hex.EncodeToString(blockHash[:])
	//fmt.Printf("%v\n", hexBlockHash)
	leadingNCharacters := hexBlockHash[:numberOfLeadingZeros]
	if leadingNCharacters == strings.Repeat("0", numberOfLeadingZeros) {
		return true
	}
	return false
}
