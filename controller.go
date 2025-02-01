package main

import (
	"html/template"
	"io/fs"
	"net/http"
)

type Cake struct {
	Name  string
	Price int
	ID    int
}

type controller struct {
	tmpl  map[string]*template.Template
	cakes map[int]Cake
	http.Handler
}

func (c *controller) LoadRouter(pub fs.FS) {
	serveMux := http.NewServeMux()

	//temporary, change in prod
	//serveMux.Handle("GET /", http.FileServerFS(unwrap(fs.Sub(pub, "public"))))
	serveMux.Handle("GET /", http.FileServerFS(pub))
	serveMux.HandleFunc("POST /order/new", c.postNewOrder)
	serveMux.HandleFunc("GET /available_cakes", c.cakeOptions)
	serveMux.HandleFunc("GET /edit_cakes", c.cakeOptions)
	serveMux.HandleFunc("GET /cake_editor", c.cakeEditor)
	serveMux.HandleFunc("GET /cake_editor/{ID}", c.cakeEditor)
	serveMux.HandleFunc("POST /new_cake", c.newCake)

	c.Handler = serveMux
}
