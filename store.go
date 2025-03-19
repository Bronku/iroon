// #todo: check for errors in every rows.Scan()
package main

import (
	"database/sql"
	_ "embed"
	"os"
	"strconv"
	"time"
)

//go:embed  schema.sql
var schema string

type store struct {
	db *sql.DB
}

// #todo: implement database persistance
func NewStore(filename string) (*store, error) {
	var out store
	var err error
	if filename != ":memory:" {
		os.Remove(filename)
	}
	// #todo: handle the error
	out.db, err = sql.Open("sqlite3", filename)
	if err != nil {
		return nil, err
	}

	_, err = out.db.Exec(schema)
	if err != nil {
		out.db.Close()
		return nil, err
	}
	return &out, nil
}

func (s *store) close() {
	s.db.Close()
}

func (s *store) getCakes() ([]cake, error) {
	rows, err := s.db.Query("select id, name, price from cake")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	cakes := make([]cake, 0)
	for rows.Next() {
		var c cake
		c.Amount = 0
		err = rows.Scan(&c.ID, &c.Name, &c.Price)
		if err != nil {
			return nil, err
		}
		cakes = append(cakes, c)
	}
	return cakes, err
}

func (s *store) saveCake(newCake cake) (int, error) {
	query := "insert into cake(name, price) values (?, ?) returning id;"
	if newCake.ID != -1 {
		query = "update cake set name = ? , price = ? where id = "
		query += strconv.Itoa(newCake.ID) + " returning id;"
	}

	// #todo: error handling
	row, err := s.db.Query(query, newCake.Name, newCake.Price)
	if err != nil {
		return -1, err
	}
	defer row.Close()

	// #todo: see if next is required or not to get the first element
	row.Next()
	err = row.Scan(&newCake.ID)
	return newCake.ID, err
}

// #todo: retrieve and save order contents
func (s *store) getOrder(id int) (order, error) {
	var out order
	// #todo: error handling
	row, err := s.db.Query("select id, name, surname, phone, location, order_date, delivery_date, status, paid from customer_order where id = ?;", id)
	defer row.Close()
	if err != nil {
		return out, err
	}

	var order_date, delivery_date string
	row.Next()
	row.Scan(&out.ID, &out.Name, &out.Surname, &out.Phone, &out.Location, &order_date, &delivery_date, &out.Status, &out.Paid)
	// #todo: error handling
	out.Accepted, _ = time.Parse("2006-01-02 15:04", order_date)
	out.Date, _ = time.Parse("2006-01-02 15:04", delivery_date)
	return out, nil
}

// #todo: should return an error
func (s *store) getOrders() ([]order, error) {
	var out []order
	rows, err := s.db.Query("select id, name, surname, phone, location, order_date, delivery_date, status, paid from customer_order;")
	if err != nil {
		return out, err
	}
	defer rows.Close()

	for rows.Next() {
		var o order
		var order_date, delivery_date string
		err = rows.Scan(&o.ID, &o.Name, &o.Surname, &o.Phone, &o.Location, &order_date, &delivery_date, &o.Status, &o.Paid)
		if err != nil {
			return nil, err
		}
		// #todo: error handling
		// figure out what to do with errors, as they shouldn't really crash the app, but there should probably be an indication that something went wrong
		o.Accepted, _ = time.Parse("2006-01-02 15:04", order_date)
		o.Date, _ = time.Parse("2006-01-02 15:04", delivery_date)
		out = append(out, o)
	}

	return out, nil
}

func (s *store) saveOrder(newOrder order) (int, error) {
	query := "insert into customer_order(name, surname, phone, location, order_date, delivery_date, status, paid) values (?, ?, ?, ?, ?, ?, ?, ?) returning id;"
	if newOrder.ID != -1 {
		query = "update customer_order set name = ?, surname = ?, phone = ?, location = ?, order_date = ?, delivery_date = ?, status = ?, paid = ? where id = "
		query += strconv.Itoa(newOrder.ID) + " returning id;"
	}

	// #todo: error handling
	// #todo: date format
	accepted := newOrder.Accepted.Format("2006-01-02 15:04")
	date := newOrder.Date.Format("2006-01-02 15:04")
	row, err := s.db.Query(query, newOrder.Name, newOrder.Surname, newOrder.Phone, newOrder.Location, accepted, date, newOrder.Status, newOrder.Paid)
	if err != nil {
		return -1, err
	}
	defer row.Close()

	// #todo: see if next is required or not to get the first element
	row.Next()
	err = row.Scan(&newOrder.ID)
	return newOrder.ID, err
}
