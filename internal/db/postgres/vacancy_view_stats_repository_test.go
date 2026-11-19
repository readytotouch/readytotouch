package postgres

import (
	"context"
	"sync"
	"testing"
	"time"

	"github.com/readytotouch/readytotouch/internal/env"

	"github.com/stretchr/testify/require"
)

func TestVacancyViewStatsRepositoryUpsert(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}

	pgConnection, err := Connection(env.Required("POSTGRES_DSN"))
	require.NoError(t, err)
	defer pgConnection.Close()

	database, err := NewDatabase(pgConnection)
	require.NoError(t, err)
	defer database.Queries().Close()

	repository := NewVacancyViewStatsRepository(database)

	var (
		ctx = context.Background()
		day = time.Date(2030, time.January, 10, 0, 0, 0, 0, time.UTC)

		userIDs   []int64
		companyID int64
		vacancyID int64
	)

	// the fixtures: the vacancy and the users, who view it
	newUser := func() int64 {
		var id int64
		require.NoError(t, pgConnection.QueryRowContext(ctx, `INSERT INTO users (created_at, updated_at) VALUES (NOW(), NOW()) RETURNING id`).Scan(&id))
		userIDs = append(userIDs, id)

		return id
	}

	owner := newUser()
	require.NoError(t, pgConnection.QueryRowContext(ctx, `INSERT INTO companies (created_at, created_by) VALUES (NOW(), $1) RETURNING id`, owner).Scan(&companyID))
	require.NoError(t, pgConnection.QueryRowContext(ctx, `INSERT INTO vacancies (company_id, created_at, created_by) VALUES ($1, NOW(), $2) RETURNING id`, companyID, owner).Scan(&vacancyID))

	defer func() {
		for _, query := range []string{
			`DELETE FROM user_vacancy_view_hourly_stats WHERE vacancy_id = $1`,
			`DELETE FROM user_vacancy_view_daily_stats WHERE vacancy_id = $1`,
			`DELETE FROM vacancy_view_daily_stats WHERE vacancy_id = $1`,
			`DELETE FROM vacancies WHERE id = $1`,
		} {
			_, err := pgConnection.ExecContext(ctx, query, vacancyID)
			require.NoError(t, err)
		}

		_, err := pgConnection.ExecContext(ctx, `DELETE FROM companies WHERE id = $1`, companyID)
		require.NoError(t, err)

		for _, id := range userIDs {
			_, err := pgConnection.ExecContext(ctx, `DELETE FROM users WHERE id = $1`, id)
			require.NoError(t, err)
		}
	}()

	count := func(query string) int64 {
		var n int64
		require.NoError(t, pgConnection.QueryRowContext(ctx, query, vacancyID).Scan(&n))

		return n
	}

	hourlyRows := func() int64 {
		return count(`SELECT COUNT(*) FROM user_vacancy_view_hourly_stats WHERE vacancy_id = $1`)
	}
	dailyRows := func() int64 {
		return count(`SELECT COUNT(*) FROM user_vacancy_view_daily_stats WHERE vacancy_id = $1`)
	}
	totalUsers := func() int64 {
		return count(`SELECT COALESCE(SUM(user_count), 0) FROM vacancy_view_daily_stats WHERE vacancy_id = $1`)
	}

	userA, userB := newUser(), newUser()

	// the steps depend on each other: they describe the life of one vacancy
	t.Run("the first view of the user counts", func(t *testing.T) {
		require.NoError(t, repository.Upsert(ctx, vacancyID, userA, day.Add(10*time.Hour+15*time.Minute)))

		require.Equal(t, int64(1), hourlyRows())
		require.Equal(t, int64(1), dailyRows())
		require.Equal(t, int64(1), totalUsers())
	})

	t.Run("the same user in the same hour changes nothing", func(t *testing.T) {
		require.NoError(t, repository.Upsert(ctx, vacancyID, userA, day.Add(10*time.Hour+45*time.Minute)))

		require.Equal(t, int64(1), hourlyRows())
		require.Equal(t, int64(1), dailyRows())
		require.Equal(t, int64(1), totalUsers())
	})

	t.Run("the same user in another hour of the same day is a new hour, but not a new daily user", func(t *testing.T) {
		require.NoError(t, repository.Upsert(ctx, vacancyID, userA, day.Add(11*time.Hour+5*time.Minute)))

		require.Equal(t, int64(2), hourlyRows())
		require.Equal(t, int64(1), dailyRows())
		require.Equal(t, int64(1), totalUsers())
	})

	t.Run("another user counts", func(t *testing.T) {
		require.NoError(t, repository.Upsert(ctx, vacancyID, userB, day.Add(11*time.Hour+10*time.Minute)))

		require.Equal(t, int64(3), hourlyRows())
		require.Equal(t, int64(2), dailyRows())
		require.Equal(t, int64(2), totalUsers())
	})

	t.Run("the same user the next day counts again", func(t *testing.T) {
		require.NoError(t, repository.Upsert(ctx, vacancyID, userA, day.Add(24*time.Hour+time.Hour)))

		require.Equal(t, int64(4), hourlyRows())
		require.Equal(t, int64(3), dailyRows())
		require.Equal(t, int64(3), totalUsers())
		require.Equal(t, int64(2), count(`SELECT COUNT(*) FROM vacancy_view_daily_stats WHERE vacancy_id = $1`))
	})

	t.Run("the stats are visible through the repository", func(t *testing.T) {
		stats, err := repository.Stats(ctx, []int64{vacancyID}, day)
		require.NoError(t, err)
		require.Equal(t, map[int64]int64{vacancyID: 3}, stats)

		stats, err = repository.Stats(ctx, []int64{vacancyID}, day.Add(24*time.Hour))
		require.NoError(t, err)
		require.Equal(t, map[int64]int64{vacancyID: 1}, stats)
	})

	t.Run("concurrent views of the same user are counted once", func(t *testing.T) {
		var (
			user = newUser()
			at   = day.Add(48*time.Hour + 3*time.Hour)
		)

		const workers = 16

		var (
			wg   sync.WaitGroup
			errs = make([]error, workers)
		)
		for i := 0; i < workers; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				errs[i] = repository.Upsert(ctx, vacancyID, user, at)
			}()
		}
		wg.Wait()

		for _, err := range errs {
			require.NoError(t, err)
		}

		require.Equal(t, int64(1), count(`SELECT COUNT(*) FROM user_vacancy_view_daily_stats WHERE vacancy_id = $1 AND created_at = '2030-01-12'`))
		require.Equal(t, int64(1), count(`SELECT COALESCE(SUM(user_count), 0) FROM vacancy_view_daily_stats WHERE vacancy_id = $1 AND created_at = '2030-01-12'`))
	})

	t.Run("concurrent views of different users are all counted", func(t *testing.T) {
		const workers = 16

		var (
			wg   sync.WaitGroup
			errs = make([]error, workers)
			at   = day.Add(72*time.Hour + 3*time.Hour)
			ids  = make([]int64, workers)
		)
		for i := 0; i < workers; i++ {
			ids[i] = newUser()
		}
		for i := 0; i < workers; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				errs[i] = repository.Upsert(ctx, vacancyID, ids[i], at)
			}()
		}
		wg.Wait()

		for _, err := range errs {
			require.NoError(t, err)
		}

		require.Equal(t, int64(workers), count(`SELECT COALESCE(SUM(user_count), 0) FROM vacancy_view_daily_stats WHERE vacancy_id = $1 AND created_at = '2030-01-13'`))
	})
}
