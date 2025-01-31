package main

import (
	"embed"
	"fmt"

	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

//go:embed public
var public embed.FS

//go:embed templates
var templates embed.FS

func main() {
	var c controller
	c.tmpl = LoadTemplates(templates)
	c.LoadRouter(public)
	c.cakes = make(map[int]Cake)
	c.cakes[0] = Cake{"Sernik", 120}
	c.cakes[1] = Cake{"Malinowa chmórka", 120}
	c.cakes[2] = Cake{"Beza Pavlova", 8}

	err := http.ListenAndServe(":8080", c)
	if err != nil {
		fmt.Println("Can't start the server: ", err)
	}
}
