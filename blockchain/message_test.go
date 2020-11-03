package blockchain

import "testing"

func TestVerifySignedMessage(t *testing.T) {
	keyPair := GenerateKeyPair()

	messageData := []byte("hello world")

	message := buildSignedMessage(keyPair.PrivateKey, messageData)

	verified := verifyMessage(message.Message, message.MessageHash)
	if verified == false {
		t.Errorf("expected the message to be valid, got %t", verified)
	}
}

func TestVerifyMessage(t *testing.T) {
	message := []byte("hello world")

	hashedMessage := hash(message)

	verified := verifyMessage(message, hashedMessage)
	if verified == false {
		t.Errorf("expected the message to be valid, got %t", verified)
	}
}
