package pki

import (
	hash "pow-blockchain/hash"
	"testing"
)

func TestGenerateKeyPair(t *testing.T) {
	keyPair := GenerateKeyPair()
	if keyPair.PrivateKey == nil || keyPair.PublicKey == nil {
		t.Error("could not generate keypair")
	}
}

func TestSign(t *testing.T) {
	keyPair := GenerateKeyPair()

	message := []byte("hello world")

	hashedMessage := hash.Hash(message)

	signature := Sign(keyPair.PrivateKey, hashedMessage)
	if signature == nil {
		t.Error("failed to sign message")
	}
}

func TestVerifySignature(t *testing.T) {
	keyPair := GenerateKeyPair()

	message := []byte("hello world")

	hashedMessage := hash.Hash(message)

	signature := Sign(keyPair.PrivateKey, hashedMessage)

	verified := VerifySignature(keyPair.PublicKey, hashedMessage, signature)
	if verified == false {
		t.Errorf("expected signature to be valid, got %v", verified)
	}
}

func TestVerifySignatureError(t *testing.T) {
	keyPair := GenerateKeyPair()

	message := []byte("hello world")

	signature := Sign(keyPair.PrivateKey, message)

	verified := VerifySignature(keyPair.PublicKey, message, signature)
	if verified != false {
		t.Error("signature should not be verifiable, as the message is not hashed")
	}
}
