package postgres

import (
	"context"
	"sync/atomic"
	"testing"
	"time"

	postgresql "github.com/readytotouch-yaaws/yaaws-go/internal/storage/postgres"

	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/stretchr/testify/require"
)

func testOnlineStorage(
	t *testing.T,
	storage OnlineStorage,
) {
	t.Helper()

	ctx := context.Background()

	connection, err := pgxpool.New(ctx, dataSourceName)
	require.NoError(t, err)
	defer connection.Close()

	var (
		pair1v1 = UserOnlinePair{
			UserID: 1,
			Online: 1679800725,
		}
		pair2v1 = UserOnlinePair{
			UserID: 2,
			Online: 1679800730,
		}
		pair3v1 = UserOnlinePair{
			UserID: 3,
			Online: 1679800735,
		}
		pair4v1 = UserOnlinePair{
			UserID: 4,
			Online: 1679800740,
		}
		pair5v1 = UserOnlinePair{
			UserID: 5,
			Online: 1679800745,
		}
	)

	truncateOnline(t, ctx, connection)

	{
		err := storage.BatchStore(ctx, []UserOnlinePair{
			pair1v1,
			pair2v1,
			pair3v1,
			pair4v1,
			pair5v1,
		})
		require.NoError(t, err)

		expectedHourlyStats(t, ctx, connection, []UserOnlinePair{
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

		err := storage.BatchStore(ctx, []UserOnlinePair{
			pair1v2,
			pair2v2,
		})
		require.NoError(t, err)

		expectedHourlyStats(t, ctx, connection, []UserOnlinePair{
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

	connection, err := pgxpool.New(ctx, dataSourceName)
	require.NoError(b, err)
	defer connection.Close()

	truncateOnline(b, ctx, connection)

	var (
		startTimestamp = time.Now().Unix()
		counter        = int64(0)
	)

	pairs := make([]UserOnlinePair, batch)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		index := atomic.AddInt64(&counter, batch)

		for j := 0; j < batch; j++ {
			pairs[j] = UserOnlinePair{
				UserID: int64(i*batch + j + 1), // 0 .. 99, 100 .. 199
				Online: startTimestamp + index,
			}
		}

		err := storage.BatchStore(ctx, pairs)

		require.NoError(b, err)
	}
	b.StopTimer()
}

func truncateOnline(t testing.TB, ctx context.Context, connection *pgxpool.Pool) {
	t.Helper()

	// clear
	{
		const (
			// language=SQL
			query = "TRUNCATE TABLE user_online_hourly_stats;"
		)

		_, err := connection.Exec(ctx, query)
		require.NoError(t, err)
	}

	{
		const (
			// language=SQL
			query = "TRUNCATE TABLE user_online_daily_stats;"
		)

		_, err := connection.Exec(ctx, query)
		require.NoError(t, err)
	}

	{
		const (
			// language=SQL
			query = "TRUNCATE TABLE user_online_daily_count_stats;"
		)

		_, err := connection.Exec(ctx, query)
		require.NoError(t, err)
	}
}

func expectedHourlyStats(t *testing.T, ctx context.Context, connection *pgxpool.Pool, expectedPairs []UserOnlinePair) {
	t.Helper()

	repository := postgresql.NewRepository(connection)
	actualPairs, err := repository.Queries().UserOnlineHourlyStats(ctx)
	require.NoError(t, err)

	hourlyActualPairs := make([]UserOnlinePair, len(actualPairs))
	for i, pair := range actualPairs {
		hourlyActualPairs[i] = UserOnlinePair{
			UserID: pair.UserID,
			Online: truncate(pair.Online.Time.Unix(), hour),
		}
	}

	hourlyExpectedPairs := make([]UserOnlinePair, len(expectedPairs))
	for i, pair := range expectedPairs {
		hourlyExpectedPairs[i] = UserOnlinePair{
			UserID: pair.UserID,
			Online: truncate(pair.Online, hour),
		}
	}

	require.Equal(t, hourlyExpectedPairs, hourlyActualPairs)
}
