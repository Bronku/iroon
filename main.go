package main

import (
	"html/template"
	"log"
	"net/http"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	var h handler
	templates, err := template.ParseFiles("index.html", "order.html")
	if err != nil {
		log.Fatal("can't parse templates: ", err)
	}
	h.tmpl = templates
	// #todo: error handling
	h.s, _ = NewStore("./foo.db")

	h.s.saveCake(cake{"Sernik ulubiony", -1, 65, 0})
	h.s.saveCake(cake{"Malinowa chmurka", -1, 150, 0})
	h.s.saveCake(cake{"Mako sernik", -1, 150, 0})
	h.s.saveCake(cake{"Rolada makowa", -1, 80, 0})
	h.s.saveCake(cake{"Wieniec bezowy", -1, 120, 0})

	h.s.saveOrder(order{-1, "Albert", "Camus", "123456789", "Kartuzy", time.Now().AddDate(0, 0, 7), time.Now(), "accepted", 0, nil})
	h.s.saveOrder(order{-1, "George", "Orwell", "", "Kartuzy", time.Now().AddDate(0, 1, 0), time.Now(), "accepted", 0, nil})
	h.s.saveOrder(order{-1, "Karl", "Marx", "0700", "Somonino", time.Now(), time.Now(), "accepted", 0, nil})

	http.HandleFunc("GET /order/", logger(h.form))
	http.HandleFunc("GET /", logger(h.index))
	http.HandleFunc("POST /", logger(h.addOrder))
	http.ListenAndServe(":8080", nil)
}
