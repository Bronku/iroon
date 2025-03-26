package store

import (
	"errors"
)

type Cake struct {
	Name         string
	ID           int
	Price        int
	Amount       int
	Category     string
	Availability string
}

func (s *Store) cakeCount() int {
	out := 0
	rows, err := s.db.Query("select count(*) from cake;")
	if err != nil {
		return out
	}
	defer rows.Close()
	_ = rows.Next()
	_ = rows.Scan(&out)
	return out
}

func (s *Store) loadCakes() ([]Cake, error) {
	out := make([]Cake, 0, s.cakeCount())

	rows, err := s.db.Query("select id, name, price, category, availability from cake order by id asc;")
	if err != nil {
		return out, err
	}
	defer rows.Close()

	for rows.Next() {
		var cake Cake
		err = rows.Scan(&cake.ID, &cake.Name, &cake.Price, &cake.Category, &cake.Availability)
		if err != nil {
			continue
		}
		out = append(out, cake)
	}

	return out, nil
}

func (s *Store) searchCakes(id int) (int, error) {
	for i, e := range s.cakes {
		if e.ID != id {
			continue
		}
		return i, nil
	}
	return 0, errors.New("cake not found")
}

func (s *Store) GetCake(id int) (Cake, error) {
	if id <= 0 {
		return Cake{}, errors.New("Invalid cake id")
	}

	i, err := s.searchCakes(id)
	if err != nil {
		return Cake{}, err
	}
	return s.cakes[i], err
}

func (s *Store) GetCakes() ([]Cake, error) {
	result := make([]Cake, len(s.cakes))
	copy(result, s.cakes)
	return result, nil
}

func (s *Store) updateCake(newCake Cake) error {
	i, err := s.searchCakes(newCake.ID)
	if err != nil {
		return err
	}
	s.cakes[i] = newCake
	return nil
}

// #todo: implement
func (s *Store) SyncCakes() error {
	return nil
}

func (s *Store) SaveCake(newCake Cake) (int, error) {
	if newCake.ID != 0 {
		return newCake.ID, s.updateCake(newCake)
	}
	newCake.ID = len(s.cakes) + 1
	s.cakes = append(s.cakes, newCake)
	return newCake.ID, nil
}
