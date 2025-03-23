package store

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

type Store struct {
	db    *sql.DB
	cakes []Cake
}

func OpenStore(filename string) *Store {
	var out Store
	var err error

	out.db, err = sql.Open("sqlite3", filename)
	if err != nil {
		log.Fatal("Can't open the database", filename, err)
	}

	out.loadMigrations()
	return &out
}

func (s *Store) Close() {
	if s.db != nil {
		s.db.Close()
	}
}
