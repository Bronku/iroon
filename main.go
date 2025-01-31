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

	err := http.ListenAndServe(":8080", c)
	if err != nil {
		fmt.Println("Can't start the server: ", err)
	}
}
