package controller

import (
	"fmt"
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
	c.tmpl["order/confirmation.html"], _ = template.ParseFS(fs, "templates/order/confirmation.html")
	c.tmpl["new_order.html"], _ = template.ParseFS(fs, "templates/new_order.html")
}

func (c *Controller) LoadRouter(pub fs.FS) {
	serveMux := http.NewServeMux()

	serveMux.Handle("GET /", http.FileServerFS(unwrap(fs.Sub(pub, "public"))))
	serveMux.HandleFunc("POST /", c.HandlePost)
	serveMux.HandleFunc("GET /new_order.html", c.HandleGet)

	c.Handler = serveMux
}

func (c *Controller) HandlePost(w http.ResponseWriter, r *http.Request) {
	fmt.Println("received Post request: ", r)
	err := r.ParseForm()
	if err != nil {
		_ = c.tmpl["order/confirmation.html"].Execute(w, struct{ Status any }{Status: err})
	}
	w.WriteHeader(http.StatusAccepted)
	fmt.Println(r.Form)
	_ = c.tmpl["order/confirmation.html"].Execute(w, struct{ Status any }{Status: r.Form})
}

func (c *Controller) HandleGet(w http.ResponseWriter, r *http.Request) {
	fmt.Println("received Get request: ", r)
	_ = c.tmpl["new_order.html"].Execute(w, nil)
}
