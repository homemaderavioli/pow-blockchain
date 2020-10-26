package main

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
	messageHash := Sha256Hash(message)
	messageSignature := Sign(privateKey, messageHash)
	return Message{
		Message:          message,
		MessageHash:      messageHash,
		MessageSignature: messageSignature,
	}
}

func VerifyMessage(recievedMessage []byte, recievedMessageHash []byte) bool {
	hashOfRecievedMessage := Sha256Hash(recievedMessage)
	if bytes.Compare(hashOfRecievedMessage, recievedMessageHash) == 0 {
		return true
	}
	return false
}
