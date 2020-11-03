package pki

import (
	"bytes"
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
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
		return nil
	}
	return signature
}

func VerifySignature(publicKey *rsa.PublicKey, message []byte, signature []byte) bool {
	err := rsa.VerifyPKCS1v15(publicKey, crypto.SHA256, message, signature)
	if err == nil {
		return true
	}
	return false
}

func GetBase64PublicKey(publicKey *rsa.PublicKey) (string, error) {
	publicKeyData, _ := x509.MarshalPKIXPublicKey(publicKey)
	publicKeyBlock := &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: publicKeyData,
	}

	var keyBuffer bytes.Buffer

	err := pem.Encode(&keyBuffer, publicKeyBlock)
	if err != nil {
		return "", err
	}

	base64PublicKey := base64.StdEncoding.EncodeToString(keyBuffer.Bytes())
	return base64PublicKey, nil
}
