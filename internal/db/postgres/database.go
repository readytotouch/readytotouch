package postgres

import (
	"context"
	"database/sql"

	"github.com/readytotouch/readytotouch/internal/storage/postgres/dbs"
)

// Database - repository
type Database struct {
	connection *sql.DB
	queries    *dbs.Queries
}

// NewDatabase - constructor
func NewDatabase(connection *sql.DB) (*Database, error) {
	var queries, err = dbs.Prepare(context.Background(), connection)
	if err != nil {
		return nil, err
	}

	return &Database{
		connection: connection,
		queries:    queries,
	}, err
}

func MustDatabase(connection *sql.DB) *Database {
	var repository, err = NewDatabase(connection)
	if err != nil {
		panic(err)
	}
	return repository
}

// Connection - getter
func (r *Database) Connection() *sql.DB {
	return r.connection
}

// Queries - getter
func (r *Database) Queries() *dbs.Queries {
	return r.queries
}

// WithTransaction - start transaction
func (r *Database) WithTransaction(ctx context.Context, fn func(queries *dbs.Queries) error) error {
	return withTransaction(ctx, r.connection, r.queries, fn)
}

func withTransaction(ctx context.Context, db *sql.DB, queries *dbs.Queries, fn func(queries *dbs.Queries) error) (err error) {
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return
	}

	defer func() {
		if p := recover(); p != nil {
			// a panic occurred, rollback and repanic
			tx.Rollback()

			panic(p)
		} else if err != nil {
			// something went wrong, rollback
			tx.Rollback()
		} else {
			// all good, commit
			err = tx.Commit()
		}
	}()

	err = fn(queries.WithTx(tx))

	return err
}
