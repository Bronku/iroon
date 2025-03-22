package server

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/Bronku/iroon/store"
)

func (h *Server) postOrder(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		fmt.Println("can't parse the form")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	fmt.Println("received form: ", r.Form)

	var n store.Order
	n.ID, err = strconv.Atoi(r.FormValue("id"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Println("can't parse order id: ", err)
		return
	}

	n.Name = strings.TrimSpace(r.FormValue("name"))

	n.Surname = strings.TrimSpace(r.FormValue("surname"))

	n.Phone = strings.TrimSpace(r.FormValue("phone"))

	n.Location = strings.TrimSpace(r.FormValue("location"))

	n.Date, err = time.Parse("2006-01-02", r.FormValue("date"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Println("can't parse order date: ", err)
		return
	}

	n.Status = strings.TrimSpace(r.FormValue("status"))

	n.Paid, err = strconv.Atoi(r.FormValue("paid"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Println("can't parse order paid: ", err)
		return
	}

	n.Accepted = time.Now()

	n.Cakes = make([]store.Cake, 0)
	cakes, err := h.s.GetCakes()
	if err != nil {
		fmt.Fprint(w, "server side error getting available cakes: ", err)
		return
	}
	for _, e := range cakes {
		value := r.FormValue(fmt.Sprintf("cake[%d]", e.ID))
		if value == "" {
			continue
		}
		amount, err := strconv.Atoi(value)
		e.Amount = amount
		if err != nil {
			continue
		}
		n.Cakes = append(n.Cakes, e)
	}

	fmt.Println("parsed order: ", n)

	h.s.SaveOrder(n)

	w.Header().Set("content-type", "text/html")
	w.WriteHeader(http.StatusAccepted)
	w.Write([]byte("accepted <a href='/'>back</a>"))
}
