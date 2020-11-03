package blockchain

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
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

func sign(privateKey *rsa.PrivateKey, message []byte) []byte {
	signature, err := rsa.SignPKCS1v15(rand.Reader, privateKey, crypto.SHA256, message)
	if err != nil {
		return nil
	}
	return signature
}

func verifySignature(publicKey *rsa.PublicKey, message []byte, signature []byte) bool {
	err := rsa.VerifyPKCS1v15(publicKey, crypto.SHA256, message, signature)
	if err == nil {
		return true
	}
	return false
}
