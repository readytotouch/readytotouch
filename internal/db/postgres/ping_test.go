package postgres

import (
	"context"
	"testing"

	"github.com/readytotouch/readytotouch/internal/env"

	"github.com/stretchr/testify/require"
)

func TestPing(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}

	ctx := context.Background()

	connection, err := Connection(env.Required("POSTGRES_DSN"))
	require.NoError(t, err)
	defer connection.Close()

	require.NoError(t, connection.PingContext(ctx))
}
