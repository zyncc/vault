package db

import (
	"database/sql"
	"log"

	_ "modernc.org/sqlite"
)

func Init() *Queries {
	conn, err := sql.Open("sqlite", "./db/store.db")
	if err != nil {
		log.Fatal("couldn't connect to database")
	}

	q := New(conn)
	return q
}
