package crypto

import (
	"crypto/rand"
	"encoding/base64"
	"log"

	"golang.org/x/crypto/argon2"
)

func PasswordHash(password, salt string) string {
	return string(argon2.Key([]byte(password), []byte(salt), 3, 32*1024, 4, 32))
}

func GenerateKey() string {
	key := [32]byte{}
	if _, err := rand.Read(key[:]); err != nil {
		log.Fatal("can't generate a vaild key", err)
	}
	return base64.StdEncoding.EncodeToString(key[:])
}
