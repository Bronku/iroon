package server

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/Bronku/iroon/store"
)

func (h *Server) index(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "text/html")
	orders, err := h.s.GetOrders()
	if err != nil {
		fmt.Fprint(w, "server side error getting orders: ", err)
		return
	}
	err = h.tmpl.ExecuteTemplate(w, "index.html", orders)
	if err != nil {
		fmt.Println("error executing the template: ", err)
	}
}

func (h *Server) getOrder(w http.ResponseWriter, r *http.Request) {
	url := strings.Split(r.URL.String(), "/")
	o := store.Order{
		ID:   -1,
		Date: time.Now(),
	}
	id, err := strconv.Atoi(url[2])
	if err == nil {
		newOrder, err := h.s.GetOrder(id)
		if err == nil {
			o = newOrder
		}
	}
	type formData struct {
		Order     store.Order
		Catalogue []store.Cake
	}

	cakes, err := h.s.GetCakes()
	if err != nil {
		fmt.Fprint(w, "server side error getting available cakes: ", err)
		return
	}
	data := formData{o, cakes}
	//fmt.Println(data)

	w.Header().Set("content-type", "text/html")
	err = h.tmpl.ExecuteTemplate(w, "order.html", data)
	if err != nil {
		fmt.Println("error executing the template: ", err)
	}
}
