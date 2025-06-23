package logging

import (
	"log"
	"net/http"
	"time"
)

func Middleware(in http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		in.ServeHTTP(w, r)
		log.Println(r.Method, r.URL.String(), time.Since(start))
	})
}
