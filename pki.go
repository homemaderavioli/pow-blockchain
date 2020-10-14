package main

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"fmt"
	"os"
)

type KeyPair struct {
	PrivateKey *rsa.PrivateKey
	PublicKey  *rsa.PublicKey
}

func GenerateKeyPair() KeyPair {
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		panic(err)
	}
	return KeyPair{
		PrivateKey: privateKey,
		PublicKey:  &privateKey.PublicKey,
	}
}

func Sign(privateKey *rsa.PrivateKey, message []byte) []byte {
	signature, err := rsa.SignPKCS1v15(rand.Reader, privateKey, crypto.SHA256, message)
	if err != nil {
		fmt.Fprintf(os.Stderr, "(Sign) error signing message: %s\n", err)
		return nil
	}
	return signature
}

func VerifySignature(publicKey *rsa.PublicKey, message []byte, signature []byte) bool {
	err := rsa.VerifyPKCS1v15(publicKey, crypto.SHA256, message, signature)
	if err == nil {
		return true
	}
	fmt.Fprintf(os.Stderr, "(VerifySignature) error verifying signature: %s\n", err)
	return false
}

func VerifyMessage(recievedMessage []byte, recievedMessageHash [32]byte) bool {
	hashOfRecievedMessage := sha256.Sum256(recievedMessage)
	if hashOfRecievedMessage == recievedMessageHash {
		return true
	}
	return false
}
