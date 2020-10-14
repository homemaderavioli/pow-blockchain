package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

func main() {
	keyPair := GenerateKeyPair()
	fmt.Println(keyPair)

	message := []byte("hello world")
	fmt.Printf("Messaged: %s\n", message)

	hashedMessage := sha256.Sum256(message)
	fmt.Printf("Hashed Messaged: %s\n", hex.EncodeToString(hashedMessage[:]))

	signature := Sign(keyPair.PrivateKey, hashedMessage[:])
	fmt.Printf("Signature: %s\n", hex.EncodeToString(signature))

	fmt.Printf("Message Verfied: %v\n", VerifyMessage(message, hashedMessage))

	verified := VerifySignature(keyPair.PublicKey, hashedMessage[:], signature)
	fmt.Printf("Signature Verified: %v\n", verified)

}
