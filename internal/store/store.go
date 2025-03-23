package store

import (
	"database/sql"
	"embed"
	_ "embed"
	"errors"
	"strconv"
	"strings"

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

	err = out.loadSchema()
	return &out, err
}

func (s *Store) Close() {
	if s.db != nil {
		s.db.Close()
	}
}

func (s *Store) version() int {
	out := -1
	row, err := s.db.Query("PRAGMA user_version;")
	if err != nil {
		return out
	}
	defer row.Close()
	row.Next()
	_ = row.Scan(&out)
	return out
}

//go:embed migrations/*.sql
var migrations embed.FS

func (s *Store) loadSchema() error {
	if s.db == nil {
		return errors.New("database doesn't exist")
	}

	migration_files, err := migrations.ReadDir("migrations")
	if err != nil {
		return err
	}
	for _, e := range migration_files {
		filename := strings.Split(e.Name(), ".")
		version, err := strconv.Atoi(filename[0])
		if err != nil {
			continue
		}
		if version <= s.version() {
			continue
		}
		query, err := migrations.ReadFile("migrations/" + e.Name())
		if err != nil {
			continue
		}
		_, err = s.db.Exec(string(query))
		if err != nil {
			return err
		}
	}

	return nil
}
