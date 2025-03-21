package main

import (
	"html/template"
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	var err error
	var h handler
	defer h.close()

	h.tmpl, err = template.ParseFiles("index.html", "order.html")
	if err != nil {
		log.Fatal("can't parse templates: ", err)
	}

	h.s, err = openStore("./foo.db")
	if err != nil {
		log.Fatal("can't open the databse", err)
	}

	http.HandleFunc("GET /order/", logger(h.form))
	http.HandleFunc("GET /", logger(h.index))
	http.HandleFunc("POST /", logger(h.addOrder))
	go http.ListenAndServe(":8080", nil)
	adminConsole()
}
