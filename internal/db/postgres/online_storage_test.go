package postgres

import (
	"context"
	"database/sql"
	"sync/atomic"
	"testing"
	"time"

	"github.com/readytotouch/readytotouch/internal/domain"
	"github.com/readytotouch/readytotouch/internal/env"

	"github.com/stretchr/testify/require"
)

func testOnlineStorage(
	t *testing.T,
	storage OnlineStorage,
) {
	t.Helper()

	ctx := context.Background()

	connection, err := Connection(env.Required("POSTGRES_DSN"))
	require.NoError(t, err)
	defer connection.Close()

	var (
		pair1v1 = domain.UserOnlinePair{
			UserID:    1,
			Timestamp: 1679800725,
		}
		pair2v1 = domain.UserOnlinePair{
			UserID:    2,
			Timestamp: 1679800730,
		}
		pair3v1 = domain.UserOnlinePair{
			UserID:    3,
			Timestamp: 1679800735,
		}
		pair4v1 = domain.UserOnlinePair{
			UserID:    4,
			Timestamp: 1679800740,
		}
		pair5v1 = domain.UserOnlinePair{
			UserID:    5,
			Timestamp: 1679800745,
		}
	)

	truncateOnline(t, ctx, connection)

	{
		err := storage.BatchStore(ctx, []domain.UserOnlinePair{
			pair1v1,
			pair2v1,
			pair3v1,
			pair4v1,
			pair5v1,
		})
		require.NoError(t, err)

		expectedHourlyStats(t, ctx, connection, []domain.UserOnlinePair{
			pair1v1,
			pair2v1,
			pair3v1,
			pair4v1,
			pair5v1,
		})
	}

	{
		var (
			pair1v2 = incUserOnlinePair(pair1v1, 1)
			pair2v2 = incUserOnlinePair(pair2v1, hour+1)
		)

		err := storage.BatchStore(ctx, []domain.UserOnlinePair{
			pair1v2,
			pair2v2,
		})
		require.NoError(t, err)

		expectedHourlyStats(t, ctx, connection, []domain.UserOnlinePair{
			pair1v1,
			pair2v1,
			pair2v2, // +
			pair3v1,
			pair4v1,
			pair5v1,
		})
	}
}

func benchmarkOnlineStorage(
	b *testing.B,
	storage OnlineStorage,
) {
	b.Helper()

	const (
		batch = 1000
	)

	ctx := context.Background()

	connection, err := Connection(env.Required("POSTGRES_DSN"))
	require.NoError(b, err)
	defer connection.Close()

	truncateOnline(b, ctx, connection)

	var (
		startTimestamp = time.Now().Unix()
		counter        = int64(0)
	)

	pairs := make([]domain.UserOnlinePair, batch)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		index := atomic.AddInt64(&counter, batch)

		for j := 0; j < batch; j++ {
			pairs[j] = domain.UserOnlinePair{
				UserID:    int64(i*batch + j + 1), // 0 .. 99, 100 .. 199
				Timestamp: startTimestamp + index,
			}
		}

		err := storage.BatchStore(ctx, pairs)

		require.NoError(b, err)
	}
	b.StopTimer()
}

func truncateOnline(t testing.TB, ctx context.Context, connection *sql.DB) {
	t.Helper()

	// clear
	{
		const (
			// language=SQL
			query = "TRUNCATE TABLE user_online_hourly_stats;"
		)

		_, err := connection.ExecContext(ctx, query)
		require.NoError(t, err)
	}

	{
		const (
			// language=SQL
			query = "TRUNCATE TABLE user_online_daily_stats;"
		)

		_, err := connection.ExecContext(ctx, query)
		require.NoError(t, err)
	}

	{
		const (
			// language=SQL
			query = "TRUNCATE TABLE user_online_daily_count_stats;"
		)

		_, err := connection.ExecContext(ctx, query)
		require.NoError(t, err)
	}
}

func expectedHourlyStats(t *testing.T, ctx context.Context, connection *sql.DB, expectedPairs []domain.UserOnlinePair) {
	t.Helper()

	database, err := NewDatabase(connection)
	require.NoError(t, err)
	defer database.Queries().Close()

	actualPairs, err := database.Queries().UserOnlineHourlyStats(ctx)
	require.NoError(t, err)

	hourlyActualPairs := make([]domain.UserOnlinePair, len(actualPairs))
	for i, pair := range actualPairs {
		hourlyActualPairs[i] = domain.UserOnlinePair{
			UserID:    pair.UserID,
			Timestamp: truncate(pair.CreatedAt.Unix(), hour),
		}
	}

	hourlyExpectedPairs := make([]domain.UserOnlinePair, len(expectedPairs))
	for i, pair := range expectedPairs {
		hourlyExpectedPairs[i] = domain.UserOnlinePair{
			UserID:    pair.UserID,
			Timestamp: truncate(pair.Timestamp, hour),
		}
	}

	require.Equal(t, hourlyExpectedPairs, hourlyActualPairs)
}
