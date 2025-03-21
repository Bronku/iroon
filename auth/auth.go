package auth

import (
	"net/http"
	"strconv"
)

type Authenticator struct {
}

func (a *Authenticator) Authenticate(in http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		c := http.Cookie{
			Name: "RequestCounter",
		}
		cookie, err := r.Cookie("RequestCounter")
		if err != nil {
			cookie = &http.Cookie{
				Value: "0",
			}
		}
		count, err := strconv.Atoi(cookie.Value)
		if err != nil {
			c.Value = "1"
		} else {
			c.Value = strconv.Itoa(count + 1)
		}
		http.SetCookie(w, &c)
		in(w, r)
	}
}
