package main

import (
	"crypto/sha256"
	"testing"
)

func TestGenerateKeyPair(t *testing.T) {
	keyPair := GenerateKeyPair()
	if keyPair.PrivateKey == nil || keyPair.PublicKey == nil {
		t.Error()
	}
}

func TestSign(t *testing.T) {
	keyPair := GenerateKeyPair()

	message := []byte("hello world")

	hashedMessage := sha256.Sum256(message)

	signature := Sign(keyPair.PrivateKey, hashedMessage[:])
	if signature == nil {
		t.Error("Failed to sign message")
	}
}

func TestVerifySignature(t *testing.T) {
	keyPair := GenerateKeyPair()

	message := []byte("hello world")

	hashedMessage := sha256.Sum256(message)

	signature := Sign(keyPair.PrivateKey, hashedMessage[:])

	verified := VerifySignature(keyPair.PublicKey, hashedMessage[:], signature)
	if verified == false {
		t.Error()
	}
}

func TestVerifyMessage(t *testing.T) {
	keyPair := GenerateKeyPair()

	message := []byte("hello world")

	hashedMessage := sha256.Sum256(message)

	Sign(keyPair.PrivateKey, hashedMessage[:])

	verified := VerifyMessage(message, hashedMessage)
	if verified == false {
		t.Error()
	}
}
