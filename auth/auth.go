package auth

import (
	_ "embed"
	"fmt"
	"net/http"
)

//go:embed  login.html
var loginPage string

//go:embed wrongPassword.html
var wrongPassword string

type Authenticator struct {
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
	c := http.Cookie{
		Name:  "token",
		Value: "good",
	}
	http.SetCookie(w, &c)
	http.Redirect(w, r, "/", http.StatusFound)
}

func (a *Authenticator) Authenticate(in http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.URL.String() == "/login" {
			a.login(w, r)
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
