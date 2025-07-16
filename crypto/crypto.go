package crypto

import (
	"crypto/rand"
	"encoding/base64"

	"golang.org/x/crypto/argon2"
)

func PasswordHash(password, salt string) string {
	var time uint32 = 3
	var memory uint32 = 32 * 1024
	var threads uint8 = 4
	var length uint32 = 32

	return string(argon2.Key([]byte(password), []byte(salt), time, memory, threads, length))
}

func GenerateKey() string {
	key := [32]byte{}
	// rand.Read always panics when encountering an error, so checking it is pointless, as the program has already panicked
	_, _ = rand.Read(key[:])

	return base64.StdEncoding.EncodeToString(key[:])
}
