package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Bronku/iroon/auth"
	"github.com/Bronku/iroon/logging"
	"github.com/Bronku/iroon/server"
	"github.com/Bronku/iroon/store"
)

func main() {
	s := store.OpenStore("./foo.db")
	s.AddUser("admin", "secret")
	defer s.Close()
	h := server.New(s)
	defer h.Close()

	var handler http.Handler = h
	handler = logging.Middleware(handler)
	handler = auth.New(s).Middleware(handler)
	fmt.Println("starting server")
	log.Fatal(http.ListenAndServe(":8080", handler))
}
