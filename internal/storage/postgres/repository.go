package postgresql

import (
	"context"

	"github.com/readytotouch-yaaws/yaaws-go/internal/storage/postgres/dbs"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

// Repository - repository
type Repository struct {
	connection *pgxpool.Pool
	queries    *dbs.Queries
}

// NewSqlcRepository - constructor
func NewSqlcRepository(connection *pgxpool.Pool) *Repository {
	var queries = dbs.New(connection)

	return &Repository{
		connection: connection,
		queries:    queries,
	}
}

// Connection - getter
func (r *Repository) Connection() *pgxpool.Pool {
	return r.connection
}

// Queries - getter
func (r *Repository) Queries() *dbs.Queries {
	return r.queries
}

// WithTransaction - start transaction
func (r *Repository) WithTransaction(ctx context.Context, fn func(queries *dbs.Queries) error) error {
	return withTransaction(ctx, r.connection, r.queries, fn)
}

func withTransaction(ctx context.Context, db *pgxpool.Pool, queries *dbs.Queries, fn func(queries *dbs.Queries) error) (err error) {
	tx, err := db.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return
	}

	defer func() {
		if p := recover(); p != nil {
			// a panic occurred, rollback and repanic
			tx.Rollback(ctx)

			panic(p)
		} else if err != nil {
			// something went wrong, rollback
			tx.Rollback(ctx)
		} else {
			// all good, commit
			err = tx.Commit(ctx)
		}
	}()

	err = fn(queries.WithTx(tx))

	return err
}
