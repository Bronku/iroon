package store

import (
	"embed"
	_ "embed"
	"errors"
	"strconv"
	"strings"
)

//go:embed migrations/*.sql
var migrations embed.FS

func (s *Store) loadFile(file string) error {
	filename := strings.Split(file, ".")
	version, err := strconv.Atoi(filename[0])
	if err != nil {
		return err
	}
	if version <= s.version() {
		return nil
	}
	query, err := migrations.ReadFile("migrations/" + file)
	if err != nil {
		return err
	}
	_, err = s.db.Exec(string(query))
	return err
}

func (s *Store) loadMigrations() error {
	if s.db == nil {
		return errors.New("database doesn't exist")
	}

	migration_files, err := migrations.ReadDir("migrations")
	if err != nil {
		return err
	}

	for _, e := range migration_files {
		_ = s.loadFile(e.Name())
	}

	return nil
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
