package main

import "time"

type cake struct {
	Name   string
	ID     int
	Price  int
	Amount int
}

type order struct {
	ID       int
	Name     string
	Surname  string
	Phone    string
	Location string
	Date     time.Time
	Accepted time.Time
	Paid     int // increments of 0.01
	Cakes    []cake
}
