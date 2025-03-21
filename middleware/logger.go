package middleware

import (
	"fmt"
	"net/http"
	"time"
)

func Logger(in http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		//log := fmt.Sprint(r.Method, " ", r.URL.String())
		fmt.Println(r.Method, r.URL.String())
		in(w, r)
		fmt.Println(r.Method, "finished in", time.Since(start))
	}
}
