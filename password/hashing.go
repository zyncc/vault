package password

import (
	"crypto/rand"
	"crypto/subtle"
	"encoding/base64"
	"fmt"

	"golang.org/x/crypto/argon2"
)

func HashPassword(password string) (string, string, error) {
	saltBytes := make([]byte, 16)
	_, err := rand.Read(saltBytes)
	if err != nil {
		return "", "", err
	}

	hashBytes := argon2.IDKey([]byte(password), saltBytes, 3, 64*1024, 2, 32)

	return base64.RawStdEncoding.EncodeToString(hashBytes), base64.RawStdEncoding.EncodeToString(saltBytes), nil
}

func CompareHash(password string, hash string, salt string) (bool, error) {
	hashBytes, err := base64.RawStdEncoding.DecodeString(hash)
	if err != nil {
		return false, fmt.Errorf("failed to compare password")
	}
	saltBytes, err := base64.RawStdEncoding.DecodeString(salt)
	if err != nil {
		return false, fmt.Errorf("failed to compare password")
	}

	newHash := argon2.IDKey([]byte(password), saltBytes, 3, 64*1024, 2, 32)

	return subtle.ConstantTimeCompare(hashBytes, newHash) == 1, nil
}
