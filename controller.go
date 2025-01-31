package main

import (
	"fmt"
	"html/template"
	"io/fs"
	"net/http"
	"strings"
)

type controller struct {
	tmpl map[string]*template.Template
	http.Handler
}

func LoadTemplates(fs fs.ReadDirFS) map[string]*template.Template {
	files, err := fs.ReadDir("templates")
	if err != nil {
		fmt.Println("can't read the templates directory")
		return nil
	}
	tmpl := make(map[string]*template.Template)
	for _, e := range files {
		if e.IsDir() {
			continue
		}
		name := e.Name()
		name = strings.Split(name, ".")[0]
		tmpl[name], err = template.ParseFS(fs, "templates/"+name+".html")
		if err != nil {
			fmt.Println("can't open template: " + name)
		}
	}
	return tmpl
}

func (c *controller) LoadRouter(pub fs.FS) {
	serveMux := http.NewServeMux()

	serveMux.Handle("GET /", http.FileServerFS(unwrap(fs.Sub(pub, "public"))))
	serveMux.HandleFunc("POST /order/new", c.postNewOrder)
	serveMux.HandleFunc("GET /cake/options", c.cakeOptions)

	c.Handler = serveMux
}
