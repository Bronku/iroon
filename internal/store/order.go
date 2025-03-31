package store

import (
	"database/sql"
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/Bronku/iroon/internal/models"
)

func (s *Store) parseOrderRow(row *sql.Rows) (models.Order, error) {
	var out models.Order
	var orderDate, deliveryDate string
	err := row.Scan(&out.ID, &out.Name, &out.Surname, &out.Phone, &out.Location, &orderDate, &deliveryDate, &out.Status, &out.Paid)
	if err != nil {
		return out, err
	}
	out.Accepted, _ = time.Parse("2006-01-02 15:04", orderDate)
	out.Date, _ = time.Parse("2006-01-02 15:04", deliveryDate)

	out.Cakes = make([]models.Cake, 0)
	rows, err := s.db.Query("select cake, amount from ordered_cake where customer_order = ?;", out.ID)
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

func (s *Store) getOrdersFromQuery(query string, args ...any) ([]models.Order, error) {
	var out []models.Order
	rows, err := s.db.Query(query, args...)
	if err != nil {
		return out, err
	}
	defer rows.Close()

	for rows.Next() {
		order, err := s.parseOrderRow(rows)
		if err != nil {
			fmt.Println(err)
			continue
		}
		out = append(out, order)
	}

	return out, nil
}

func (s *Store) GetOrder(id int) (models.Order, error) {
	query := "select * from customer_order where id = ?;"
	out, err := s.getOrdersFromQuery(query, id)
	if err != nil {
		return models.Order{}, err
	}
	if len(out) < 1 {
		return models.Order{}, errors.New("no output order")
	}
	return out[0], nil
}

func (s *Store) GetFilteredOrder(filter string, from, to time.Time) ([]models.Order, error) {
	if filter == "" {
		return s.GetTopOrders(from, to)
	}
	start := from.Format("2006-01-02") + " 00:00"
	end := to.Format("2006-01-02") + " 99:99"
	if to.IsZero() {
		end = "9999-99-99 99:99"
	}
	query := "select id, name, surname, phone, location, order_date, delivery_date, status, paid from order_fts(?) where delivery_date >= ? and delivery_date <= ? order by rank;"
	return s.getOrdersFromQuery(query, filter, start, end)
}

func (s *Store) GetTopOrders(from, to time.Time) ([]models.Order, error) {
	start := from.Format("2006-01-02") + " 00:00"
	end := to.Format("2006-01-02") + " 99:99"
	if to.IsZero() {
		end = "9999-99-99 99:99"
	}
	query := "select * from customer_order where status != 'done' and delivery_date >= ? and delivery_date <= ? ;"
	return s.getOrdersFromQuery(query, start, end)
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
		return 0, errors.New("the database didn't respond with an id")
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
