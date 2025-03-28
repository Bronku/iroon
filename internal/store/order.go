package store

import (
	"errors"
	"strconv"
	"time"

	"github.com/Bronku/iroon/internal/models"
)

func (s *Store) GetOrder(id int) (models.Order, error) {
	var out models.Order
	row, err := s.db.Query("select id, name, surname, phone, location, order_date, delivery_date, status, paid from customer_order where id = ?;", id)
	if err != nil {
		return out, err
	}
	defer row.Close()

	var order_date, delivery_date string
	row.Next()
	err = row.Scan(&out.ID, &out.Name, &out.Surname, &out.Phone, &out.Location, &order_date, &delivery_date, &out.Status, &out.Paid)
	if err != nil {
		return out, err
	}
	out.Accepted, _ = time.Parse("2006-01-02 15:04", order_date)
	out.Date, _ = time.Parse("2006-01-02 15:04", delivery_date)

	out.Cakes = make([]models.Cake, 0)
	rows, err := s.db.Query("select cake, amount from ordered_cake where customer_order = ?;", id)
	if err != nil {
		return out, err
	}
	defer rows.Close()
	for rows.Next() {
		var newCake models.Cake
		err = rows.Scan(&newCake.ID, &newCake.Amount)
		if err != nil {
			return out, err
		}
		cakeData, err := s.GetCake(newCake.ID)
		if err != nil {
			continue
		}
		newCake.Name = cakeData.Name
		newCake.Price = cakeData.Price
		out.Cakes = append(out.Cakes, newCake)
	}

	return out, nil
}

func (s *Store) GetOrders() ([]models.Order, error) {
	var out []models.Order
	rows, err := s.db.Query("select id, name, surname, phone, location, order_date, delivery_date, status, paid from customer_order;")
	if err != nil {
		return out, err
	}
	defer rows.Close()

	for rows.Next() {
		var o models.Order
		var order_date, delivery_date string
		err = rows.Scan(&o.ID, &o.Name, &o.Surname, &o.Phone, &o.Location, &order_date, &delivery_date, &o.Status, &o.Paid)
		if err != nil {
			return nil, err
		}
		o.Accepted, _ = time.Parse("2006-01-02 15:04", order_date)
		o.Date, _ = time.Parse("2006-01-02 15:04", delivery_date)

		o.Cakes = make([]models.Cake, 0)
		rows, err := s.db.Query("select cake, amount from ordered_cake where customer_order = ?;", o.ID)

		if err != nil {
			return out, err
		}
		defer rows.Close()
		for rows.Next() {
			var newCake models.Cake
			err = rows.Scan(&newCake.ID, &newCake.Amount)
			if err != nil {
				return out, err
			}
			cakeData, err := s.GetCake(newCake.ID)
			if err != nil {
				continue
			}
			newCake.Name = cakeData.Name
			newCake.Price = cakeData.Price
			o.Cakes = append(o.Cakes, newCake)
		}

		out = append(out, o)
	}

	return out, nil
}

func (s *Store) SaveOrder(newOrder models.Order) (int, error) {
	query := "insert into customer_order(name, surname, phone, location, order_date, delivery_date, status, paid) values (?, ?, ?, ?, ?, ?, ?, ?) returning id;"
	if newOrder.ID != 0 {
		query = "update customer_order set name = ?, surname = ?, phone = ?, location = ?, order_date = ?, delivery_date = ?, status = ?, paid = ? where id = "
		query += strconv.Itoa(newOrder.ID) + " returning id;"
	}
	tx, err := s.db.Begin()
	if err != nil {
		return 0, err
	}

	accepted := newOrder.Accepted.Format("2006-01-02 15:04")
	date := newOrder.Date.Format("2006-01-02 15:04")
	row, err := tx.Query(query, newOrder.Name, newOrder.Surname, newOrder.Phone, newOrder.Location, accepted, date, newOrder.Status, newOrder.Paid)
	if err != nil {
		_ = tx.Rollback()
		return 0, err
	}
	defer row.Close()

	if !row.Next() {
		_ = tx.Rollback()
		return 0, errors.New("The database didn't respond with an id")
	}
	err = row.Scan(&newOrder.ID)
	if err != nil {
		_ = tx.Rollback()
		return newOrder.ID, err
	}

	// remove all ordered_cakes associated with this order before inserting
	query = "delete from ordered_cake where customer_order = ?;"
	_, err = tx.Exec(query, newOrder.ID)
	if err != nil {
		_ = tx.Rollback()
		return newOrder.ID, err
	}

	// add all ordered_cakes for this order
	query = "insert into ordered_cake(customer_order, cake, amount) values (?,?,?);"
	for _, e := range newOrder.Cakes {
		_, err := tx.Exec(query, newOrder.ID, e.ID, e.Amount)
		if err != nil {
			_ = tx.Rollback()
			return newOrder.ID, err
		}
	}

	err = tx.Commit()
	return newOrder.ID, err
}
