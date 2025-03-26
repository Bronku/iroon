package server

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/Bronku/iroon/internal/store"
)

func (h *Server) index(r *http.Request) (any, int, error) {
	data, err := h.s.GetOrders()
	return data, http.StatusOK, err
}

func (h *Server) cakes(r *http.Request) (any, int, error) {
	data, err := h.s.GetCakes()
	return data, http.StatusOK, err
}

func (h *Server) cake(r *http.Request) (any, int, error) {
	url := strings.Split(r.URL.String(), "/")
	if len(url) < 3 || url[2] == "" {
		return store.Cake{}, http.StatusOK, nil
	}

	id, err := strconv.Atoi(url[2])
	if err != nil {
		return nil, http.StatusBadRequest, err
	}

	data, err := h.s.GetCake(id)
	if err != nil {
		return nil, http.StatusNotFound, err
	}

	return data, http.StatusOK, nil
}

func (h *Server) order(r *http.Request) (any, int, error) {
	type formData struct {
		Order     store.Order
		Catalogue []store.Cake
	}
	var err error
	var data formData

	data.Catalogue, err = h.s.GetCakes()
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	url := strings.Split(r.URL.String(), "/")
	if len(url) < 3 || url[2] == "" {
		return data, http.StatusOK, nil
	}

	id, err := strconv.Atoi(url[2])
	if err != nil {
		return nil, http.StatusBadRequest, err
	}

	data.Order, err = h.s.GetOrder(id)
	if err != nil {
		return nil, http.StatusNotFound, err
	}

	return data, http.StatusOK, nil
}
