package auth

import (
	"errors"
	"net/http"
	"time"

	"github.com/Bronku/iroon/crypto"
	"github.com/Bronku/iroon/models"
)

func (a *Authenticator) getSession(r *http.Request) (models.Token, error) {
	c, err := r.Cookie("token")
	if err != nil {
		return models.Token{}, err
	}
	var session models.Token
	result := a.db.First(&session, c.Value)
	if result.Error != nil {
		return models.Token{}, errors.New("session not found")
	}
	if time.Since(session.Expiration) > 0 {
		a.db.Delete(&session)
		return models.Token{}, errors.New("session expired")
	}
	return session, nil
}

func (a *Authenticator) login(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
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
	var session models.Token
	session.User = user
	session.Expiration = time.Now().Add(time.Hour * 24)
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
