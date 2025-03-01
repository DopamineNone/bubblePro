package crypto

import (
	"bytes"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"golang.org/x/crypto/pbkdf2"
)

func generateSalt(length int) (salt []byte, err error) {
	salt = make([]byte, length)
	_, err = rand.Read(salt)
	if err != nil {
		return nil, err
	}
	return
}

func GenerateHashedString(input []byte, hashLength, saltLength, iteration int) (hash string, err error) {
	var salt []byte
	if salt, err = generateSalt(saltLength); err != nil {
		return "", err
	}
	hashBytes := pbkdf2.Key(input, salt, iteration, hashLength, sha256.New)
	hash = hex.EncodeToString(append(hashBytes, salt...))
	return
}

func VerifyIfPlainAndHashMatched(plain []byte, hash string, hashLength, iteration int) bool {
	hashBytes, err := hex.DecodeString(hash)
	if err != nil || len(hashBytes) < hashLength {
		return false
	}
	newHash := pbkdf2.Key(plain, hashBytes[hashLength:], iteration, hashLength, sha256.New)
	return bytes.Equal(newHash, hashBytes[:hashLength])
}
