package server

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/Bronku/iroon/store"
)

func (h *Server) index(r *http.Request) (any, error) {
	return h.s.GetOrders()
}

func (h *Server) order(r *http.Request) (any, error) {
	type formData struct {
		Order     store.Order
		Catalogue []store.Cake
	}
	var err error
	var data formData

	data.Catalogue, err = h.s.GetCakes()
	if err != nil {
		return nil, err
	}

	url := strings.Split(r.URL.String(), "/")
	if len(url) < 3 || url[2] == "" {
		return data, nil
	}

	id, err := strconv.Atoi(url[2])
	if err != nil {
		return nil, err
	}

	data.Order, err = h.s.GetOrder(id)
	if err != nil {
		return nil, err
	}

	return data, nil
}
