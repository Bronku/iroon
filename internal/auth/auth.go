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

func (a *Authenticator) ensureAuth(in http.Handler) http.Handler {
	fmt.Println("ensureAuth called")
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
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

func (a *Authenticator) Middleware(in http.Handler) http.Handler {
	handler := http.NewServeMux()
	handler.HandleFunc("GET /login", getLogin)
	handler.HandleFunc("POST /login", a.login)
	handler.Handle("/", a.ensureAuth(in))
	return handler
}
