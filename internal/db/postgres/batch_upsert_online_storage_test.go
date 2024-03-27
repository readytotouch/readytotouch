package postgres

import (
	"testing"

	"github.com/readytotouch/readytotouch/internal/env"

	"github.com/stretchr/testify/require"
)

func TestBatchUpdateOnlineStorage(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}

	pgConnection, err := Connection(env.Required("POSTGRES_DSN"))
	require.NoError(t, err)
	defer pgConnection.Close()

	database, err := NewDatabase(pgConnection)
	require.NoError(t, err)
	defer database.Queries().Close()

	storage := NewBatchUpsertOnlineStorage(database)

	testOnlineStorage(t, storage)
}

func BenchmarkBatchUpdateOnlineStorage(b *testing.B) {
	if testing.Short() {
		b.Skip()
	}

	pgConnection, err := Connection(env.Required("POSTGRES_DSN"))
	require.NoError(b, err)
	defer pgConnection.Close()

	database, err := NewDatabase(pgConnection)
	require.NoError(b, err)
	defer database.Queries().Close()

	storage := NewBatchUpsertOnlineStorage(database)

	benchmarkOnlineStorage(b, storage)
}
