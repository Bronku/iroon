package auth

import (
	"crypto/rand"
	"encoding/base64"
	"log"
	"time"
)

type token struct {
	userName   string
	created    time.Time
	lastAccess time.Time
}

func generateKey() string {
	key := [32]byte{}
	if _, err := rand.Read(key[:]); err != nil {
		log.Fatal("can't generate a vaild key", err)
	}
	return base64.StdEncoding.EncodeToString(key[:])
}
