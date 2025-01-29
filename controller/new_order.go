package controller

import (
	"fmt"
	"net/http"
)

func (c *Controller) postNewOrder(w http.ResponseWriter, r *http.Request) {
	fmt.Println("received Post request: ", r)
	err := r.ParseForm()
	if err != nil {
		_ = c.tmpl["order/post_new"].Execute(w, struct{ Status any }{Status: err})
	}
	w.WriteHeader(http.StatusAccepted)
	fmt.Println(r.Form)
	_ = c.tmpl["order/post_new"].Execute(w, struct{ Status any }{Status: r.Form})
}

func (c *Controller) getNewOrder(w http.ResponseWriter, r *http.Request) {
	fmt.Println("received Get request: ", r)
	_ = c.tmpl["order/get_new"].Execute(w, nil)
}
