package main

import (
	"log"
	"net/http"

	"github.com/Bronku/iroon/auth"
	"github.com/Bronku/iroon/middleware"
	"github.com/Bronku/iroon/router"
)

func main() {
	h, err := router.New()
	if err != nil {
		log.Fatal(err)
	}
	defer h.Close()

	var a auth.Authenticator
	a = auth.New()

	// #todo: make handler functions private, and create a servemux in router
	http.HandleFunc("GET /order/", middleware.Logger(a.Authenticate(h.Form)))
	http.HandleFunc("GET /", middleware.Logger(a.Authenticate(h.Index)))
	http.HandleFunc("POST /", middleware.Logger(a.Authenticate(h.AddOrder)))
	http.ListenAndServe(":8080", nil)
}
