package main

import "testing"

func TestBuildSignedMessage(t *testing.T) {
	keyPair := GenerateKeyPair()

	messageData := []byte("hello world")

	message := buildSignedMessage(keyPair.PrivateKey, messageData)

	verified := VerifyMessage(message.Message, message.MessageHash)
	if verified == false {
		t.Error()
	}
}

func TestVerifyMessage(t *testing.T) {
	keyPair := GenerateKeyPair()

	message := []byte("hello world")

	hashedMessage := Sha256Hash(message)

	Sign(keyPair.PrivateKey, hashedMessage[:])

	verified := VerifyMessage(message, hashedMessage)
	if verified == false {
		t.Error()
	}
}
