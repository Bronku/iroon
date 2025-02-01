package main

import (
	"embed"
	"fmt"
	"os"

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
	publicfs := os.DirFS("public")
	c.LoadRouter(publicfs)
	c.cakes = make(map[int]Cake)
	c.cakes[0] = Cake{"Sernik", 120, 0}
	c.cakes[1] = Cake{"Malinowa chmórka", 120, 1}
	c.cakes[2] = Cake{"Beza Pavlova", 8, 2}

	err := http.ListenAndServe(":8080", c)
	if err != nil {
		fmt.Println("Can't start the server: ", err)
	}
}
