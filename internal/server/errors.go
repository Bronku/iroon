package server

import (
	"fmt"
	"net/http"
)

func errorPage(err error, status int) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(status)
		fmt.Fprint(w, err.Error())
	}
}
