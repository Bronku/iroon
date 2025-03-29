package server

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/Bronku/iroon/internal/models"
)

func (h *Server) orders(r *http.Request) (any, int, error) {
	//r.URL.Query().Get()
	data, err := h.s.GetTopOrders()
	return data, http.StatusOK, err
}

func (h *Server) ordersSearch(r *http.Request) (any, int, error) {
	q := r.URL.Query().Get("q")
	fmt.Println(q)
	data, err := h.s.GetTopOrders()
	data = data[:3]
	return data, http.StatusOK, err
}

func (h *Server) cakes(r *http.Request) (any, int, error) {
	data, err := h.s.GetCakes()
	return data, http.StatusOK, err
}

func (h *Server) cake(r *http.Request) (any, int, error) {
	url := strings.Split(r.URL.String(), "/")
	if len(url) < 3 || url[2] == "" {
		return models.Cake{}, http.StatusOK, nil
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
		Order     models.Order
		Catalogue []models.Cake
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
