package store

import (
	"database/sql"
	"log"

	"github.com/Bronku/iroon/internal/models"
	_ "github.com/mattn/go-sqlite3"
)

type Store struct {
	db    *sql.DB
	cakes []models.Cake
}

func OpenStore(filename string) *Store {
	var out Store
	var err error

	out.db, err = sql.Open("sqlite3", filename)
	if err != nil {
		log.Fatal("Can't open the database", filename, err)
	}

	out.loadMigrations()
	out.cakes, err = out.loadCakes()
	if err != nil {
		log.Fatal(err)
	}
	return &out
}

func (s *Store) Close() {
	if s.db != nil {
		s.db.Close()
	}
}
