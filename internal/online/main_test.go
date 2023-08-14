package main

import (
	"context"
	"testing"

	"github.com/jackc/pgx/v5"

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

	connection, err := pgx.Connect(ctx, dataSourceName)
	require.NoError(t, err)
	defer connection.Close(ctx)

	require.NoError(t, connection.Ping(ctx))
}
