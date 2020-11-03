package blockchain

import (
	"bytes"
	"crypto/rsa"
)

type Message struct {
	Message          []byte
	MessageHash      []byte
	MessageSignature []byte
}

func buildSignedMessage(privateKey *rsa.PrivateKey, message []byte) Message {
	messageHash := hash(message)
	messageSignature := sign(privateKey, messageHash)
	return Message{
		Message:          message,
		MessageHash:      messageHash,
		MessageSignature: messageSignature,
	}
}

func verifyMessage(recievedMessage []byte, recievedMessageHash []byte) bool {
	hashOfRecievedMessage := hash(recievedMessage)
	if bytes.Compare(hashOfRecievedMessage, recievedMessageHash) == 0 {
		return true
	}
	return false
}
