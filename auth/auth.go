package auth

import (
	"fmt"
	"net/http"
)

type Authenticator struct {
}

func (a *Authenticator) Authenticate(in http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.URL.String() == "/loginok" {
			c := http.Cookie{
				Name:  "token",
				Value: "good",
			}
			http.SetCookie(w, &c)
			http.Redirect(w, r, "/", http.StatusFound)
			return
		}
		if r.URL.String() == "/login" {
			w.Header().Set("content-type", "text/html")
			fmt.Fprint(w, "<a href='/loginok'>Login</a>")
			return
		}
		if _, err := r.Cookie("token"); err != nil {
			fmt.Println("user not logged in")
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}
		in(w, r)
	}
}
