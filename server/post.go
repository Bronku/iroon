package server

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/Bronku/iroon/models"
)

func atoui(input string) (uint, error) {
	num, err := strconv.ParseUint(input, 10, 0)
	return uint(num), err
}

func (h *Server) postCake(r *http.Request) (any, int, error) {
	err := r.ParseForm()
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	var n models.Product
	n.ID, err = atoui(r.FormValue("id"))
	if err != nil {
		return nil, http.StatusBadRequest, err
	}
	n.Name = r.FormValue("name")
	n.Price, err = atoui(r.FormValue("price"))
	if err != nil {
		return nil, http.StatusBadRequest, err
	}
	result := h.db.Save(&n)
	return n, http.StatusAccepted, result.Error
}

func (h *Server) postOrder(r *http.Request) (any, int, error) {
	var cakes []models.Product
	result := h.db.Find(&cakes)
	if result.Error != nil {
		return nil, http.StatusInternalServerError, result.Error
	}

	err := r.ParseForm()
	log.Println("received form:", r.PostForm)
	if err != nil {
		return nil, http.StatusBadRequest, err
	}

	var n models.Order
	n.ID, err = atoui(r.FormValue("id"))
	if err != nil {
		return nil, http.StatusBadRequest, err
	}
	n.Prepaid, err = atoui(r.FormValue("paid"))
	if err != nil {
		return nil, http.StatusBadRequest, err
	}
	n.Date, err = time.Parse("2006-01-02", r.FormValue("date"))
	if err != nil {
		return nil, http.StatusBadRequest, err
	}

	n.OrderItems = make([]models.OrderItem, 0)
	for _, e := range cakes {
		count, err := atoui(r.FormValue(fmt.Sprintf("cake[%d]", e.ID)))
		if err != nil {
			continue
		}
		n.OrderItems = append(n.OrderItems, models.OrderItem{Amount: count, Product: e})
	}

	n.Name = strings.TrimSpace(r.FormValue("name"))
	n.Surname = strings.TrimSpace(r.FormValue("surname"))
	n.Phone = strings.TrimSpace(r.FormValue("phone"))
	n.Location = strings.TrimSpace(r.FormValue("location"))
	n.Status = strings.TrimSpace(r.FormValue("status"))

	log.Println("parsed order:", n)
	err = h.db.Save(&n).Error
	if err != nil {
		log.Println("save: ", err)
	}
	err = h.db.Model(&n).Association("OrderItems").Replace(n.OrderItems)
	if err != nil {
		log.Println("replace: ", err)
	}
	return n, http.StatusAccepted, nil
}
