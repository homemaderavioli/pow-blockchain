package blockchain

import (
	"testing"
)

func buildUnsignedMessage(message []byte) Message {
	return Message{
		Message:          message,
		MessageHash:      nil,
		MessageSignature: nil,
	}
}

func getGenesisBlock() *Block {
	message := buildUnsignedMessage([]byte("James 9001"))
	block := &Block{
		BlockNumber:       0,
		Nonce:             594,
		Message:           message,
		PreviousBlockHash: nil,
	}
	return block
}

func TestGetGenesisBlockNonce(t *testing.T) {
	block := getGenesisBlock()
	block.findNonce()
	if block.Nonce != 594 {
		t.Errorf("got %d, expected nonce of 594", block.Nonce)
	}
}

func TestNewBlockchain(t *testing.T) {
	genesisBlock := getGenesisBlock()
	blockchain := initialise(genesisBlock)
	if blockchain.length() != 1 {
		t.Errorf("should be length of 1")
	}
}

func TestAddBlockToBlockchain(t *testing.T) {
	genesisBlock := getGenesisBlock()

	if genesisBlock.validBlock() == false {
		t.Errorf("invalid block")
	}

	blockchain := initialise(genesisBlock)

	message := buildUnsignedMessage([]byte("haha 1337"))
	blockNumber := blockchain.length()
	previousBlockHash := blockchain.getTopBlockHash()

	block := &Block{
		BlockNumber:       blockNumber,
		Nonce:             20333,
		Message:           message,
		PreviousBlockHash: previousBlockHash,
	}

	if block.validBlock() == false {
		t.Errorf("invalid block")
	}

	err := blockchain.addBlock(block)
	if err != nil {
		t.Errorf("%s", err)
	}

	if blockchain.length() != 2 {
		t.Errorf("should be length of 2")
	}
}

func TestAddInvalidBlockToBlockchain(t *testing.T) {
	genesisBlock := getGenesisBlock()

	if genesisBlock.validBlock() == false {
		t.Errorf("invalid block")
	}

	blockchain := initialise(genesisBlock)

	message := buildUnsignedMessage([]byte("haha 1337"))
	blockNumber := blockchain.length()
	previousBlockHash := blockchain.getTopBlockHash()

	block := &Block{
		BlockNumber:       blockNumber,
		Nonce:             20333,
		Message:           message,
		PreviousBlockHash: previousBlockHash,
	}

	if block.validBlock() == false {
		t.Errorf("invalid block")
	}

	err := blockchain.addBlock(block)
	if err != nil {
		t.Error(err)
	}

	err = blockchain.addBlock(block)
	if err == nil {
		t.Error(err)
	}
}

func TestValidateBlockChain(t *testing.T) {
	genesisBlock := getGenesisBlock()
	blockchain := initialise(genesisBlock)

	message := buildUnsignedMessage([]byte("haha 1337"))
	blockNumber := blockchain.length()
	previousBlockHash := blockchain.getTopBlockHash()

	block := &Block{
		BlockNumber:       blockNumber,
		Nonce:             20333,
		Message:           message,
		PreviousBlockHash: previousBlockHash,
	}

	if block.validBlock() == false {
		t.Errorf("invalid block")
	}

	err := blockchain.addBlock(block)
	if err != nil {
		t.Error(err)
	}

	err = blockchain.validateBlockChain()
	if err != nil {
		t.Error(err)
	}
}

func TestInvalidateBlockChain(t *testing.T) {
	genesisBlock := getGenesisBlock()
	blockchain := initialise(genesisBlock)

	message := buildUnsignedMessage([]byte("haha 1337"))
	blockNumber := blockchain.length()
	previousBlockHash := blockchain.getTopBlockHash()

	block := &Block{
		BlockNumber:       blockNumber,
		Nonce:             20333,
		Message:           message,
		PreviousBlockHash: previousBlockHash,
	}

	if block.validBlock() == false {
		t.Errorf("invalid block")
	}

	err := blockchain.addBlock(block)
	if err != nil {
		t.Errorf("%s", err)
	}

	block2 := &Block{
		BlockNumber:       blockNumber,
		Nonce:             20333,
		Message:           message,
		PreviousBlockHash: previousBlockHash,
	}

	blockchain.blocks = append(blockchain.blocks, block2)

	err = blockchain.validateBlockChain()
	if err == nil {
		t.Error(err)
	}
}
