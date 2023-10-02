package postgres

import (
	"database/sql"
)

func Connection(dsn string) (*sql.DB, error) {
	return sql.Open("postgres", dsn)
}

func MustConnection(dsn string) *sql.DB {
	connection, err := Connection(dsn)
	if err != nil {
		panic(err)
	}
	return connection
}
