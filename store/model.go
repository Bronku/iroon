package store

import "time"

type Cake struct {
	Name   string
	ID     int
	Price  int
	Amount int
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
