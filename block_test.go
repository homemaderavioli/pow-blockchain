package main

import (
	"testing"
)

func TestTryNotFindNonce(t *testing.T) {
	block := &Block{
		BlockNumber:       0,
		Nonce:             0,
		Message:           Message{},
		PreviousBlockHash: nil,
	}
	nonce := 9001
	notFound := block.tryNonce(nonce)
	if notFound {
		t.Errorf("what")
	}
}

func TestHasValidNonce(t *testing.T) {
	block := &Block{
		BlockNumber:       0,
		Nonce:             46802,
		Message:           Message{},
		PreviousBlockHash: nil,
	}
	bSerialized := block.blockSerialized()
	hasLeadingZeros := blockHashHasLeadingZeros(bSerialized)
	if hasLeadingZeros == false {
		t.Errorf("not found")
	}
}
func TestHasInvalidNonce(t *testing.T) {
	block := &Block{
		BlockNumber:       0,
		Nonce:             9001,
		Message:           Message{},
		PreviousBlockHash: nil,
	}
	bSerialized := block.blockSerialized()
	hasLeadingZeros := blockHashHasLeadingZeros(bSerialized)
	if hasLeadingZeros {
		t.Errorf("not found")
	}
}
