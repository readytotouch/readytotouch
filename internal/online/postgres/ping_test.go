package postgres

import (
	"context"
	"testing"

	"github.com/readytotouch-yaaws/yaaws-go/internal/env"

	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/stretchr/testify/require"
)

func TestPing(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}

	ctx := context.Background()

	connection, err := pgxpool.New(ctx, env.Must("POSTGRES_DSN"))
	require.NoError(t, err)
	defer connection.Close()

	require.NoError(t, connection.Ping(ctx))
}
