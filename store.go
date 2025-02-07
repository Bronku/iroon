package main

import (
	"errors"
)

type store struct {
	cakes  []cake
	orders []order
}

func NewStore() *store {
	var out store
	out.cakes = make([]cake, 0)
	out.orders = make([]order, 0)
	return &out
}

func (s *store) getCakes() []cake {
	return s.cakes
}

func (s *store) saveCake(newCake cake) int {
	if newCake.ID == -1 {
		newCake.ID = len(s.cakes)
		s.cakes = append(s.cakes, newCake)
		return newCake.ID
	}
	for i := range s.cakes {
		if s.cakes[i].ID == newCake.ID {
			s.cakes[i] = newCake
			break
		}
	}
	return newCake.ID
}

func (s *store) getOrder(id int) (order, error) {
	for _, e := range s.orders {
		if e.ID == id {
			return e, nil
		}
	}
	return order{}, errors.New("order not found")
}

func (s *store) getOrders() []order {
	return s.orders
}

func (s *store) saveOrder(newOrder order) int {
	if newOrder.ID == -1 {
		newOrder.ID = len(s.orders)
		s.orders = append(s.orders, newOrder)
		return newOrder.ID
	}

	for i := range s.orders {
		if s.orders[i].ID == newOrder.ID {
			s.orders[i] = newOrder
			break
		}
	}
	return newOrder.ID
}
