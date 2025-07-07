package models

import (
	"time"

	"gorm.io/gorm"
)

// money amounts are in incremetns of 0.01

type OrderItem struct {
	Product   Product
	Amount    uint
	ProductID uint `gorm:"primaryKey;autoIncrement:false"`
	OrderID   uint `gorm:"primaryKey;autoIncrement:false"`
}

type Order struct {
	gorm.Model
	Name       string
	Surname    string
	Phone      string
	Location   string
	Status     string
	Prepaid    uint
	Date       time.Time
	OrderItems []OrderItem
}

func (c *OrderItem) Total() uint {
	return c.Amount * c.Product.Price
}

func (o *Order) Total() uint {
	var out uint
	for _, e := range o.OrderItems {
		out += e.Total()
	}
	out -= o.Prepaid
	return out
}
