package main

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

func (c *controller) cakeOptions(w http.ResponseWriter, r *http.Request) {
	type cake struct {
		Name string
		ID   int
	}
	w.Header().Set("Content-Type", "text/html")
	template := strings.Split(r.URL.String(), "/")[1]
	c.tmpl[template].Execute(w, struct{ Cakes []cake }{Cakes: []cake{{"Hello", 1}, {"World", 2}}})
}

func (c *controller) editCakes(w http.ResponseWriter, r *http.Request) {
	type cake struct {
		Name  string
		ID    int
		Price int
	}
	w.Header().Set("Content-Type", "text/html")
	c.tmpl["cake_editor"].Execute(w, cake{"Hello", 1, 120})
}

func (c *controller) newCake(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		c.tmpl["alert"].Execute(w, err)
		return
	}

	type cake struct {
		Name  string
		ID    int
		Price int
	}
	var out cake
	out.Name = r.FormValue("name")
	out.Price, err = strconv.Atoi(r.FormValue("price"))
	if err != nil {
		c.tmpl["alert"].Execute(w, err)
		return
	}
	out.ID, err = strconv.Atoi(r.FormValue("id"))
	if err != nil {
		c.tmpl["alert"].Execute(w, err)
		return
	}
	_ = c.tmpl["cake"].Execute(w, out)
}

func (c *controller) postNewOrder(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		_ = c.tmpl["confirm_order"].Execute(w, struct{ Status any }{Status: err})
	}
	w.WriteHeader(http.StatusAccepted)
	fmt.Println(r.Form)
	_ = c.tmpl["confirm_order"].Execute(w, struct{ Status any }{Status: r.Form})
}
