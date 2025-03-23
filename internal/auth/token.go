package auth

import (
	"crypto/rand"
	"encoding/base64"
	"time"
)

type token struct {
	userName   string
	created    time.Time
	lastAccess time.Time
}

func generateKey() string {
	key := [32]byte{}
	rand.Read(key[:])
	return base64.StdEncoding.EncodeToString(key[:])
}
