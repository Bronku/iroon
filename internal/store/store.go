package store

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

type Store struct {
	db *sql.DB
}

func OpenStore(filename string) (*Store, error) {
	var out Store
	var err error

	out.db, err = sql.Open("sqlite3", filename)
	if err != nil {
		return &out, err
	}

	err = out.loadMigrations()
	return &out, err
}

func (s *Store) Close() {
	if s.db != nil {
		s.db.Close()
	}
}
