package controller

import (
	"html/template"
	"io/fs"
	"net/http"
)

type Controller struct {
	tmpl map[string]*template.Template
	http.Handler
}

// #todo: load all automatically
func (c *Controller) LoadTemplates(fs fs.FS) {
	c.tmpl = make(map[string]*template.Template)
	c.tmpl["order/confirmation"], _ = template.ParseFS(fs, "templates/order/confirmation.html")
	c.tmpl["cake/options"], _ = template.ParseFS(fs, "templates/cake/options.html")
}

func (c *Controller) LoadRouter(pub fs.FS) {
	serveMux := http.NewServeMux()

	serveMux.Handle("GET /", http.FileServerFS(unwrap(fs.Sub(pub, "public"))))
	serveMux.HandleFunc("POST /order/new", c.postNewOrder)
	serveMux.HandleFunc("GET /cake/options", c.cakeOptions)

	c.Handler = serveMux
}
