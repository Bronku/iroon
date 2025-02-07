package main

import (
	"fmt"
	"net/http"
	"time"
)

func logger(in http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		log := fmt.Sprint(r.Method, " ", r.URL.String())
		in(w, r)
		log += fmt.Sprint(" ", time.Since(start))
		fmt.Println(log)
	}
}
