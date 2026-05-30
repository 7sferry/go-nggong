package db

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func Open(url string) (*sql.DB, error) {
	conn, err := sql.Open("postgres", url)
	if err != nil {
		return nil, err
	}
	if err := conn.Ping(); err != nil {
		_ = conn.Close()
		return nil, err
	}
	return conn, nil
}
