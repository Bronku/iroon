package auth

import (
	"errors"
	"net/http"
	"time"

	"github.com/Bronku/iroon/crypto"
	"github.com/Bronku/iroon/models"
)

var ErrNoCookie = errors.New("session cookie not found")
var ErrSessionNotFound = errors.New("session not found")
var ErrSessionExpired = errors.New("session has expired")

func (a *Authenticator) getSession(r *http.Request) (models.Token, error) {
	cookie, err := r.Cookie("token")
	if err != nil {
		return models.Token{}, ErrNoCookie
	}
	var session models.Token
	result := a.db.First(&session, "token = ?", cookie.Value)
	if result.Error != nil {
		return models.Token{}, ErrSessionNotFound
	}
	if time.Since(session.Expiration) > 0 {
		a.db.Delete(&session)
		return models.Token{}, ErrSessionExpired
	}
	return session, nil
}

func (a *Authenticator) login(w http.ResponseWriter, r *http.Request) {
	_ = r.ParseForm()
	login := r.PostFormValue("login")
	password := r.PostFormValue("password")
	if a.verifyCredentials(login, password) != nil {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}
	cookie := a.newSession(login)
	http.SetCookie(w, &cookie)
	http.Redirect(w, r, "/", http.StatusFound)
}

func (a *Authenticator) logout(w http.ResponseWriter, r *http.Request) {
	cookie := http.Cookie{
		Name:     "Token",
		Value:    "nil",
		HttpOnly: true,
		SameSite: http.SameSiteStrictMode,
		Path:     "/",
	}
	http.SetCookie(w, &cookie)
	http.Redirect(w, r, "/", http.StatusFound)
	c, err := r.Cookie("token")
	if err != nil {
		return
	}
	a.db.Delete(&models.Token{}, c.Value)
}

func (a *Authenticator) newSession(user string) http.Cookie {
	key := crypto.GenerateKey()
	session := models.Token{
		User:       user,
		Expiration: time.Now().Add(a.config.SessionExpiration),
		Token:      key,
	}
	a.db.Create(&session)

	cookie := http.Cookie{
		Name:     "token",
		Value:    key,
		HttpOnly: true,
		SameSite: http.SameSiteStrictMode,
		Path:     "/",
		// Secure: true, #todo
	}
	return cookie
}
