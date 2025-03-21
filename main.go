package main

import (
	"html/template"
	"log"
	"net/http"

	"github.com/Bronku/iroon/auth"
	"github.com/Bronku/iroon/middleware"
	"github.com/Bronku/iroon/store"
)

func main() {
	var err error
	var h handler
	var a auth.Authenticator
	defer h.close()

	h.tmpl, err = template.ParseFiles("index.html", "order.html")
	if err != nil {
		log.Fatal("can't parse templates: ", err)
	}

	h.s, err = store.OpenStore("./foo.db")
	if err != nil {
		log.Fatal("can't open the databse", err)
	}

	a = auth.New()

	http.HandleFunc("GET /order/", middleware.Logger(a.Authenticate(h.form)))
	http.HandleFunc("GET /", middleware.Logger(a.Authenticate(h.index)))
	http.HandleFunc("POST /", middleware.Logger(a.Authenticate(h.addOrder)))
	http.ListenAndServe(":8080", nil)
}
