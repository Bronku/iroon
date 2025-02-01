package main

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

func (c *controller) cakeOptions(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	template := strings.Split(r.URL.String(), "/")[1]
	c.tmpl[template].Execute(w, c.cakes)
}

func (c *controller) cakeEditor(w http.ResponseWriter, r *http.Request) {
	url := strings.Split(r.URL.String(), "/")
	if len(url) < 3 {
		c.tmpl["cake_editor"].Execute(w, Cake{"", 0, -1})
		return
	}

	id, err := strconv.Atoi(url[2])
	if err != nil {
		c.tmpl["cake_editor"].Execute(w, Cake{"", 0, -1})
		return
	}

	cake, exists := c.cakes[id]
	if !exists {
		c.tmpl["cake_editor"].Execute(w, Cake{"", 0, -1})
		return
	}

	c.tmpl["cake_editor"].Execute(w, cake)
}

func (c *controller) newCake(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		c.tmpl["alert"].Execute(w, err)
		return
	}

	var in Cake
	in.Name = r.FormValue("name")
	in.Price, err = strconv.Atoi(r.FormValue("price"))
	if err != nil {
		c.tmpl["alert"].Execute(w, err)
		return
	}
	in.ID, err = strconv.Atoi(r.FormValue("id"))
	if err != nil {
		c.tmpl["alert"].Execute(w, err)
		return
	}

	if in.ID == -1 {
		in.ID = len(c.cakes)
	}

	c.cakes[in.ID] = in
	_ = c.tmpl["cake"].Execute(w, in)
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
