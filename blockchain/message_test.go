package blockchain

import (
	"testing"

	hash "pow-blockchain/hash"
	pki "pow-blockchain/pki"
)

func TestVerifySignedMessage(t *testing.T) {
	keyPair := pki.GenerateKeyPair()

	messageData := []byte("hello world")

	message := buildSignedMessage(keyPair.PrivateKey, messageData)

	verified := verifyMessage(message.Message, message.MessageHash)
	if verified == false {
		t.Errorf("expected the message to be valid, got %t", verified)
	}
}

func TestVerifyMessage(t *testing.T) {
	message := []byte("hello world")

	hashedMessage := hash.Hash(message)

	verified := verifyMessage(message, hashedMessage)
	if verified == false {
		t.Errorf("expected the message to be valid, got %t", verified)
	}
}
