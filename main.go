package main

import (
	"html/template"
	"log"
	"net/http"

	"github.com/Bronku/iroon/auth"
	_ "github.com/mattn/go-sqlite3"
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

	h.s, err = openStore("./foo.db")
	if err != nil {
		log.Fatal("can't open the databse", err)
	}

	a = auth.New()

	http.HandleFunc("GET /order/", logger(a.Authenticate(h.form)))
	http.HandleFunc("GET /", logger(a.Authenticate(h.index)))
	http.HandleFunc("POST /", logger(a.Authenticate(h.addOrder)))
	http.ListenAndServe(":8080", nil)
}
