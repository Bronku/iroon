package main

import (
	"fmt"
	"net/http"
)

func (c *controller) cakeOptions(w http.ResponseWriter, _ *http.Request) {
	type cake struct {
		Name string
		ID   int
	}
	c.tmpl["available_cakes"].Execute(w, struct{ Cakes []cake }{Cakes: []cake{{"Hello", 1}, {"World", 2}}})
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
