package blockchain

import "crypto/sha256"

func Sha256Hash(plaintext []byte) []byte {
	hash := sha256.Sum256(plaintext)
	return hash[:]
}
