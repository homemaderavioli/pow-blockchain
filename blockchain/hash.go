package blockchain

import "crypto/sha256"

func hash(plaintext []byte) []byte {
	hash := sha256.Sum256(plaintext)
	return hash[:]
}
