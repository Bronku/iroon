package auth

import (
	_ "embed"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/Bronku/iroon/internal/models"
	"github.com/Bronku/iroon/internal/store"
)

//go:embed  templates/login.html
var loginPage string

//go:embed templates/wrongPassword.html
var wrongPassword string

type Authenticator struct {
	sessions map[string]models.Token
	s        *store.Store
}

func New(s *store.Store) *Authenticator {
	var out Authenticator
	var err error
	out.s = s
	out.sessions, err = s.GetSessions()
	fmt.Println(out.sessions)
	if err != nil {
		log.Fatal(err)
	}
	return &out
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
	cookie, err := a.newSession(login)
	// #todo: change to a some sort of internal server error
	if err != nil {
		w.Header().Set("content-type", "text/html")
		fmt.Fprint(w, "internal server error")
		return
	}
	http.SetCookie(w, &cookie)
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
		if !ok {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}
		if time.Since(value.Expiration) > 0 {
			delete(a.sessions, c.Value)
			err := a.s.CleanSessions()
			if err != nil {
				fmt.Println("error cleaning the sessions", err)
			}
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}
		in.ServeHTTP(w, r)
	})
}
