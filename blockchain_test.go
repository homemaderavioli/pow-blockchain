package main

import (
	"encoding/hex"
	"fmt"
	"testing"
)

func getGenesisBlock() *Block {
	message := buildMessage("James 9001")
	block := &Block{
		BlockNumber:       0,
		Nonce:             53263,
		Message:           message,
		PreviousBlockHash: nil,
	}
	return block
}

func TestGetGenesisBlockNonce(t *testing.T) {
	t.Skip()
	block := getGenesisBlock()
	block.findNonce()
	fmt.Printf("%d\n", block.Nonce)
}

func TestNewBlockchain(t *testing.T) {
	genesisBlock := getGenesisBlock()
	blockchain := New(genesisBlock)
	if blockchain.length != 1 {
		t.Errorf("should be length of 1")
	}
}

func TestAddBlockToBlockchain(t *testing.T) {
	genesisBlock := getGenesisBlock()

	if genesisBlock.validBlock() == false {
		t.Errorf("invalid block")
	}

	blockchain := New(genesisBlock)

	if genesisBlock.validBlock() == false {
		t.Errorf("invalid block")
	}

	message := buildMessage("haha 1337")
	blockNumber := blockchain.length
	previousBlockHash := blockchain.getTopBlockHash()

	fmt.Printf("%s\n", hex.EncodeToString(previousBlockHash))

	block := &Block{
		BlockNumber:       blockNumber,
		Nonce:             555,
		Message:           message,
		PreviousBlockHash: previousBlockHash,
	}

	if block.validBlock() == false {
		t.Errorf("invalid block")
	}

	blockchain.addBlock(block)

	if blockchain.length != 2 {
		t.Errorf("should be length of 2")
	}

	fmt.Printf("%v\n", blockchain)
}
