package store

import (
	"embed"
	_ "embed"
	"log"
	"strconv"
	"strings"
)

//go:embed migrations/*.sql
var migrations embed.FS

func (s *Store) loadFile(file string) {
	filename := strings.Split(file, ".")
	version, err := strconv.Atoi(filename[0])
	if err != nil {
		log.Fatal("only allowed files in migrations directory are <version>.txt")
	}
	if version <= s.version() {
		return
	}
	log.Println("migration: ", file)

	query, _ := migrations.ReadFile("migrations/" + file)
	_, err = s.db.Exec(string(query))
	if err != nil {
		log.Fatal("error executing migration: ", file, err)
	}
}

func (s *Store) loadMigrations() {
	migrationFiles, _ := migrations.ReadDir("migrations")

	for _, e := range migrationFiles {
		s.loadFile(e.Name())
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
