package main

import (
	"context"
	"testing"

	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/stretchr/testify/require"
)

const (
	dataSourceName = "postgresql://yaaws_user:yaaws_password@postgres:5432/yaaws?sslmode=disable&timezone=UTC"
)

func TestPing(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}

	ctx := context.Background()

	connection, err := pgxpool.New(ctx, dataSourceName)
	require.NoError(t, err)
	defer connection.Close()

	require.NoError(t, connection.Ping(ctx))
}
