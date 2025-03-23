package logging

import (
	"fmt"
	"net/http"
	"time"
)

func Middleware(in http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		fmt.Println(r.Method, r.URL.String())
		in.ServeHTTP(w, r)
		fmt.Println(r.Method, "finished in", time.Since(start))
	})
}
