package main

type Message struct {
	Message          []byte
	MessageHash      []byte
	MessageSignature []byte
}

type Block struct {
	BlockNumber       int
	Nonce             int
	Message           Message
	PreviousBlockHash []byte
}
