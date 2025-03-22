package main

import (
	"log"
	"net/http"

	"github.com/Bronku/iroon/auth"
	"github.com/Bronku/iroon/logging"
	"github.com/Bronku/iroon/router"
)

func main() {
	h, err := router.New()
	if err != nil {
		log.Fatal(err)
	}
	defer h.Close()

	var handler http.Handler = h
	handler = logging.Middleware(handler)
	handler = auth.New().Middleware(handler)
	log.Fatal(http.ListenAndServe(":8080", handler))
}
