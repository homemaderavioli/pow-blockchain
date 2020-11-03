package hash

import "crypto/sha256"

func Hash(plaintext []byte) []byte {
	hash := sha256.Sum256(plaintext)
	return hash[:]
}
