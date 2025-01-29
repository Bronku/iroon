package controller

import (
	"net/http"
)

func (c *Controller) cakeOptions(w http.ResponseWriter, _ *http.Request) {
	type cake struct {
		Name string
		ID   int
	}
	//w.Header().Set("Content-Type", "text/html")
	c.tmpl["cake/options"].Execute(w, struct{ Cakes []cake }{Cakes: []cake{{"Hello", 1}, {"World", 2}}})
}
