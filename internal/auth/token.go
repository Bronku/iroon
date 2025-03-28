package auth

import (
	"crypto/rand"
	"encoding/base64"
	"log"
	"net/http"
	"time"

	"github.com/Bronku/iroon/internal/models"
)

func (a *Authenticator) newSession(user string) (http.Cookie, error) {
	key := generateKey()
	var session models.Token
	var cookie http.Cookie
	session.User = user
	session.Expiration = time.Now().Add(time.Hour * 24)
	a.sessions[key] = session
	err := a.s.AddSession(key, user, session.Expiration)
	if err != nil {
		return cookie, err
	}

	cookie.Name = "token"
	cookie.Value = key
	cookie.HttpOnly = true
	cookie.SameSite = http.SameSiteStrictMode
	cookie.Path = "/"
	//cookie.Secure = true
	return cookie, nil
}

func generateKey() string {
	key := [32]byte{}
	if _, err := rand.Read(key[:]); err != nil {
		log.Fatal("can't generate a vaild key", err)
	}
	return base64.StdEncoding.EncodeToString(key[:])
}
