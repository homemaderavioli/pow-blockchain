package blockchain

import (
	"bytes"
	"crypto/rsa"

	hash "pow-blockchain/hash"
	pki "pow-blockchain/pki"
)

type Message struct {
	Message          []byte
	MessageHash      []byte
	MessageSignature []byte
}

func buildSignedMessage(privateKey *rsa.PrivateKey, message []byte) Message {
	messageHash := hash.Hash(message)
	messageSignature := pki.Sign(privateKey, messageHash)
	return Message{
		Message:          message,
		MessageHash:      messageHash,
		MessageSignature: messageSignature,
	}
}

func verifyMessage(recievedMessage []byte, recievedMessageHash []byte) bool {
	hashOfRecievedMessage := hash.Hash(recievedMessage)
	if bytes.Compare(hashOfRecievedMessage, recievedMessageHash) == 0 {
		return true
	}
	return false
}
