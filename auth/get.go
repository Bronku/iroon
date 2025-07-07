package auth

import (
	"fmt"
	"net/http"

	_ "embed"
)

//go:embed  templates/login.html
var loginHTML string

// returns a simple login page
func getLogin(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("content-type", "text/html")
	fmt.Fprint(w, loginHTML)
	return
}
