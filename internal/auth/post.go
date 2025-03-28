package auth

import (
	"fmt"
	"net/http"
)

func (a *Authenticator) login(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}
	login := r.PostFormValue("login")
	password := r.PostFormValue("password")
	if a.verifyCredentials(login, password) != nil {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}
	cookie, err := a.newSession(login)
	if err != nil {
		w.Header().Set("content-type", "text/html")
		fmt.Fprint(w, "internal server error")
		return
	}
	http.SetCookie(w, &cookie)
	http.Redirect(w, r, "/", http.StatusFound)
}
