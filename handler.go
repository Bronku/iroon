package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type handler struct {
	tmpl *template.Template
	s    *store
}

func (h *handler) form(w http.ResponseWriter, r *http.Request) {
	url := strings.Split(r.URL.String(), "/")
	o := order{
		ID:   -1,
		Date: time.Now(),
	}
	id, err := strconv.Atoi(url[2])
	if err == nil {
		newOrder, err := h.s.getOrder(id)
		if err == nil {
			o = newOrder
		}
	}
	type formData struct {
		Order     order
		Catalogue []cake
	}
	data := formData{o, h.s.getCakes()}
	fmt.Println(data)

	w.Header().Set("content-type", "text/html")
	err = h.tmpl.ExecuteTemplate(w, "order.html", data)
	if err != nil {
		fmt.Println("error executing the template: ", err)
	}
}

func (h *handler) addOrder(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		fmt.Println("can't parse the form")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	fmt.Println("received form: ", r.Form)

	var n order
	n.ID, err = strconv.Atoi(r.FormValue("id"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Println("can't parse order id: ", err)
		return
	}
	n.Name = r.FormValue("name")
	n.Surname = r.FormValue("surname")
	n.Phone = r.FormValue("phone")
	n.Location = r.FormValue("location")
	n.Date, err = time.Parse("2006-01-02", r.FormValue("date"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Println("can't parse order date: ", err)
		return
	}
	n.Paid, err = strconv.Atoi(r.FormValue("paid"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Println("can't parse order paid: ", err)
		return
	}
	n.Cakes = make([]cake, 0)

	h.s.saveOrder(n)

	fmt.Println("parsed order: ", n)

	w.Header().Set("content-type", "text/html")
	w.WriteHeader(http.StatusAccepted)
	w.Write([]byte("accepted <a href='/'>back</a>"))
}

func (h *handler) index(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "text/html")
	err := h.tmpl.ExecuteTemplate(w, "index.html", h.s.getOrders())
	if err != nil {
		fmt.Println("error executing the template: ", err)
	}
}
