package store

import (
	"database/sql"
	_ "embed"
	"errors"
	"os"
	"strconv"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

//go:embed  schema.sql
var schema string

type Store struct {
	db *sql.DB
}

func OpenStore(filename string) (*Store, error) {
	var out Store
	var err error
	_, err = os.Stat(filename)
	if errors.Is(err, os.ErrNotExist) {
		defer out.loadSchema()
	}
	out.db, err = sql.Open("sqlite3", filename)
	return &out, err
}

func (s *Store) Close() {
	if s.db != nil {
		s.db.Close()
	}
}

func (s *Store) loadSchema() error {
	if s.db == nil {
		return errors.New("database doesn't exist")
	}
	_, err := s.db.Exec(schema)
	return err
}

func (s *Store) GetCake(id int) (Cake, error) {
	out := Cake{ID: id}
	row, err := s.db.Query("select name, price from cake where id = ?;", id)
	if err != nil {
		return out, err
	}
	defer row.Close()
	row.Next()
	err = row.Scan(&out.Name, &out.Price)
	return out, err
}

func (s *Store) GetCakes() ([]Cake, error) {
	rows, err := s.db.Query("select id, name, price from cake")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	cakes := make([]Cake, 0)
	for rows.Next() {
		var c Cake
		c.Amount = 0
		err = rows.Scan(&c.ID, &c.Name, &c.Price)
		if err != nil {
			return nil, err
		}
		cakes = append(cakes, c)
	}
	return cakes, err
}

func (s *Store) SaveCake(newCake Cake) (int, error) {
	query := "insert into cake(name, price) values (?, ?) returning id;"
	if newCake.ID != 0 {
		query = "update cake set name = ? , price = ? where id = "
		query += strconv.Itoa(newCake.ID) + " returning id;"
	}

	row, err := s.db.Query(query, newCake.Name, newCake.Price)
	if err != nil {
		return -1, err
	}
	defer row.Close()

	row.Next()
	err = row.Scan(&newCake.ID)
	return newCake.ID, err
}

func (s *Store) GetOrder(id int) (Order, error) {
	var out Order
	row, err := s.db.Query("select id, name, surname, phone, location, order_date, delivery_date, status, paid from customer_order where id = ?;", id)
	defer row.Close()
	if err != nil {
		return out, err
	}

	var order_date, delivery_date string
	row.Next()
	err = row.Scan(&out.ID, &out.Name, &out.Surname, &out.Phone, &out.Location, &order_date, &delivery_date, &out.Status, &out.Paid)
	if err != nil {
		return out, err
	}
	out.Accepted, _ = time.Parse("2006-01-02 15:04", order_date)
	out.Date, _ = time.Parse("2006-01-02 15:04", delivery_date)

	out.Cakes = make([]Cake, 0)
	rows, err := s.db.Query("select cake, amount from ordered_cake where customer_order = ?;", id)
	if err != nil {
		return out, err
	}
	defer rows.Close()
	for rows.Next() {
		var newCake Cake
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

func (s *Store) GetOrders() ([]Order, error) {
	var out []Order
	rows, err := s.db.Query("select id, name, surname, phone, location, order_date, delivery_date, status, paid from customer_order;")
	if err != nil {
		return out, err
	}
	defer rows.Close()

	for rows.Next() {
		var o Order
		var order_date, delivery_date string
		err = rows.Scan(&o.ID, &o.Name, &o.Surname, &o.Phone, &o.Location, &order_date, &delivery_date, &o.Status, &o.Paid)
		if err != nil {
			return nil, err
		}
		o.Accepted, _ = time.Parse("2006-01-02 15:04", order_date)
		o.Date, _ = time.Parse("2006-01-02 15:04", delivery_date)

		o.Cakes = make([]Cake, 0)
		rows, err := s.db.Query("select cake, amount from ordered_cake where customer_order = ?;", o.ID)

		if err != nil {
			return out, err
		}
		defer rows.Close()
		for rows.Next() {
			var newCake Cake
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

func (s *Store) SaveOrder(newOrder Order) (int, error) {
	query := "insert into customer_order(name, surname, phone, location, order_date, delivery_date, status, paid) values (?, ?, ?, ?, ?, ?, ?, ?) returning id;"
	if newOrder.ID != 0 {
		query = "update customer_order set name = ?, surname = ?, phone = ?, location = ?, order_date = ?, delivery_date = ?, status = ?, paid = ? where id = "
		query += strconv.Itoa(newOrder.ID) + " returning id;"
	}
	tx, err := s.db.Begin()
	if err != nil {
		return -1, nil
	}

	accepted := newOrder.Accepted.Format("2006-01-02 15:04")
	date := newOrder.Date.Format("2006-01-02 15:04")
	row, err := tx.Query(query, newOrder.Name, newOrder.Surname, newOrder.Phone, newOrder.Location, accepted, date, newOrder.Status, newOrder.Paid)
	if err != nil {
		tx.Rollback()
		return -1, err
	}
	defer row.Close()

	row.Next()
	err = row.Scan(&newOrder.ID)
	if err != nil {
		tx.Rollback()
		return newOrder.ID, err
	}

	// remove all ordered_cakes associated with this order before inserting
	query = "delete from ordered_cake where customer_order = ?;"
	_, err = tx.Exec(query, newOrder.ID)
	if err != nil {
		tx.Rollback()
		return newOrder.ID, err
	}

	// add all ordered_cakes for this order
	query = "insert into ordered_cake(customer_order, cake, amount) values (?,?,?);"
	for _, e := range newOrder.Cakes {
		_, err := tx.Exec(query, newOrder.ID, e.ID, e.Amount)
		if err != nil {
			tx.Rollback()
			return newOrder.ID, err
		}
	}

	err = tx.Commit()
	return newOrder.ID, err
}
