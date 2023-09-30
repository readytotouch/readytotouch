package postgres

import (
	"context"
	"testing"

	"github.com/readytotouch-yaaws/yaaws-go/internal/env"

	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/stretchr/testify/require"
)

func TestBatchUpdateOnlineStorage(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}

	ctx := context.Background()

	connection, err := pgxpool.New(ctx, env.Required("POSTGRES_DSN"))
	require.NoError(t, err)
	defer connection.Close()

	storage := NewBatchUpdateOnlineStorage(NewDatabase(connection))

	testOnlineStorage(t, storage)
}

func BenchmarkBatchUpdateOnlineStorage(b *testing.B) {
	if testing.Short() {
		b.Skip()
	}

	ctx := context.Background()

	connection, err := pgxpool.New(ctx, env.Required("POSTGRES_DSN"))
	require.NoError(b, err)
	defer connection.Close()

	storage := NewBatchUpdateOnlineStorage(NewDatabase(connection))

	benchmarkOnlineStorage(b, storage)
}
