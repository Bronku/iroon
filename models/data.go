package models

import "time"

type Cake struct {
	Name         string
	ID           int
	Price        int // increments of 0.01
	Amount       int
	Category     string
	Availability string
}

type Order struct {
	ID       int
	Name     string
	Surname  string
	Phone    string
	Location string
	Date     time.Time
	Accepted time.Time
	Status   string
	Paid     int // increments of 0.01
	Cakes    []Cake
}

func (o *Order) Total() int {
	out := o.Paid * (-1)
	for _, e := range o.Cakes {
		out += e.Price * e.Amount
	}
	return out
}

func (c *Cake) Total() int {
	return c.Amount * c.Price
}
