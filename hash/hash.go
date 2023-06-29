package hash

import (
	"crypto/rand"
	"encoding/base64"
	"golang.org/x/crypto/argon2"
	"strings"
)

func HashPassword(password string) (string, error) {
	salt := make([]byte, 16)
	_, err := rand.Read(salt)
	if err != nil {
		return "", err
	}

	hash := argon2.IDKey([]byte(password), salt, 1, 64*1024, 4, 32)

	encodedHash := base64.StdEncoding.EncodeToString(hash)
	encodedSalt := base64.StdEncoding.EncodeToString(salt)

	return encodedSalt + ":" + encodedHash, nil
}

func ComparePasswords(hashedPassword, providedPassword string) bool {
	parts := strings.Split(hashedPassword, ":")
	salt, _ := base64.StdEncoding.DecodeString(parts[0])
	hashed := argon2.IDKey([]byte(providedPassword), salt, 1, 64*1024, 4, 32)
	encodedHash := base64.StdEncoding.EncodeToString(hashed)

	return parts[1] == encodedHash
}
