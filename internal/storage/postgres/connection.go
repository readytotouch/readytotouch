package postgres

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

func Connection(ctx context.Context, dsn string) (*pgxpool.Pool, error) {
	return pgxpool.New(ctx, dsn)
}

func MustConnection(ctx context.Context, dsn string) *pgxpool.Pool {
	var connection, err = Connection(ctx, dsn)
	if err != nil {
		panic(err)
	}
	return connection
}
