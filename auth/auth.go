package auth

import (
	_ "embed"
	"fmt"
	"net/http"
	"time"
)

//go:embed  login.html
var loginPage string

//go:embed wrongPassword.html
var wrongPassword string

type Authenticator struct {
	sessions map[string]token
}

func New() *Authenticator {
	return &Authenticator{sessions: make(map[string]token)}
}

func (a *Authenticator) login(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("content-type", "text/html")
		fmt.Fprint(w, loginPage)
		return
	}
	err := r.ParseForm()
	if err != nil {
		w.Header().Set("content-type", "text/html")
		fmt.Fprint(w, loginPage)
		return
	}
	login := r.PostFormValue("login")
	password := r.PostFormValue("password")
	if login != "admin" || password != "secret" {
		w.Header().Set("content-type", "text/html")
		fmt.Fprint(w, wrongPassword)
		return
	}
	key := generateKey()
	a.sessions[key] = token{userName: login, created: time.Now(), lastAccess: time.Now()}
	c := http.Cookie{
		Name:     "token",
		Value:    key,
		HttpOnly: true,
		// #todo (in prod) uncomment this line
		//Secure: true,
		SameSite: http.SameSiteStrictMode,
		Path:     "/",
	}
	http.SetCookie(w, &c)
	http.Redirect(w, r, "/", http.StatusFound)
}

func (a *Authenticator) Middleware(in http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.String() == "/login" {
			a.login(w, r)
			return
		}
		c, err := r.Cookie("token")
		if err != nil {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}
		value, ok := a.sessions[c.Value]
		if !ok || time.Since(value.lastAccess) > time.Hour*24 || time.Since(value.lastAccess) > time.Hour*240 {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}
		value.lastAccess = time.Now()
		a.sessions[c.Value] = value
		in.ServeHTTP(w, r)
	})
}
