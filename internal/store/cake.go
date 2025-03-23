package store

import (
	"errors"
	"strconv"
)

type Cake struct {
	Name   string
	ID     int
	Price  int
	Amount int
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
		return 0, err
	}
	defer row.Close()

	if !row.Next() {
		return 0, errors.New("The database didn't respond with an id")
	}
	err = row.Scan(&newCake.ID)
	return newCake.ID, err
}
