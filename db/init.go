package db

import (
	"database/sql"
	"log"
	"os"
	"path/filepath"

	_ "modernc.org/sqlite"
)

func Init() *Queries {
	base, err := os.UserConfigDir()
	if err != nil {
		log.Fatal("database not found")
	}

	dir := filepath.Join(base, "vault")
	if err := os.MkdirAll(dir, 0700); err != nil {
		log.Fatal("database not found")
	}

	dbPath := filepath.Join(dir, "vault.db")
	conn, err := sql.Open("sqlite", dbPath)
	if err != nil {
		log.Fatal("couldn't connect to database")
	}

	q := New(conn)
	return q
}
