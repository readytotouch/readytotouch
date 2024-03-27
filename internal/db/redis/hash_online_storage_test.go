package redis

import (
	"context"
	"sort"
	"testing"

	"github.com/readytotouch/readytotouch/internal/domain"

	"github.com/stretchr/testify/require"
)

func TestRedisHashOnlineStorage(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}

	testHashOnlineStorage(t, "redis:6379")
}

func testHashOnlineStorage(t *testing.T, addr string) {
	t.Helper()

	ctx := context.Background()

	client, err := Client(ctx, addr)
	require.NoError(t, err)

	require.NoError(t, client.FlushAll(ctx).Err())

	storage := NewHashOnlineStorage(client)

	expected := []domain.UserOnlinePair{
		{
			UserID:    10000001,
			Timestamp: 1679800725,
		},
		{
			UserID:    10000002,
			Timestamp: 1679800730,
		},
		{
			UserID:    10000003,
			Timestamp: 1679800735,
		},
	}

	for _, pair := range expected {
		err := storage.Store(ctx, pair)

		require.NoError(t, err)
	}

	actualCount, err := storage.Count(ctx)
	require.NoError(t, err)
	require.Equal(t, int64(len(expected)), int64(actualCount))

	actual, err := storage.GetAndClear(ctx)
	require.NoError(t, err)

	requireUserOnlinePairsEqual(t, expected, actual)
}

func requireUserOnlinePairsEqual(t testing.TB, expected, actual []domain.UserOnlinePair) {
	t.Helper()

	require.Equal(t, len(expected), len(actual))

	sort.Sort(domain.UserOnlinePairs(expected))
	sort.Sort(domain.UserOnlinePairs(actual))

	require.Equal(t, expected, actual)
}
