package main

import (
	"log"
	"net/http"

	"github.com/Bronku/iroon/auth"
	"github.com/Bronku/iroon/middleware"
	"github.com/Bronku/iroon/router"
)

func main() {
	var err error
	var h router.Router
	var a auth.Authenticator
	defer h.Close()

	err = h.LoadTemplates()
	if err != nil {
		log.Fatal("can't parse templates: ", err)
	}

	err = h.OenStore()
	if err != nil {
		log.Fatal("can't open the databse", err)
	}

	a = auth.New()

	// #todo: make handler functions private, and create a servemux in router
	http.HandleFunc("GET /order/", middleware.Logger(a.Authenticate(h.Form)))
	http.HandleFunc("GET /", middleware.Logger(a.Authenticate(h.Index)))
	http.HandleFunc("POST /", middleware.Logger(a.Authenticate(h.AddOrder)))
	http.ListenAndServe(":8080", nil)
}
