package blockchain

import (
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

	hashedMessage := hash(message)

	signature := sign(keyPair.PrivateKey, hashedMessage)
	if signature == nil {
		t.Error("failed to sign message")
	}
}

func TestVerifySignature(t *testing.T) {
	keyPair := GenerateKeyPair()

	message := []byte("hello world")

	hashedMessage := hash(message)

	signature := sign(keyPair.PrivateKey, hashedMessage)

	verified := verifySignature(keyPair.PublicKey, hashedMessage, signature)
	if verified == false {
		t.Errorf("expected signature to be valid, got %v", verified)
	}
}

func TestVerifySignatureError(t *testing.T) {
	keyPair := GenerateKeyPair()

	message := []byte("hello world")

	signature := sign(keyPair.PrivateKey, message)

	verified := verifySignature(keyPair.PublicKey, message, signature)
	if verified != false {
		t.Error("signature should not be verifiable, as the message is not hashed")
	}
}
