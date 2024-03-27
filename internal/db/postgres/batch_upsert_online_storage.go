package postgres

import (
	"context"

	"github.com/readytotouch/readytotouch/internal/domain"
	"github.com/readytotouch/readytotouch/internal/storage/postgres/dbs"
)

type BatchUpsertOnlineStorage struct {
	db *Database
}

func NewBatchUpsertOnlineStorage(db *Database) *BatchUpsertOnlineStorage {
	return &BatchUpsertOnlineStorage{db: db}
}

func (s *BatchUpsertOnlineStorage) BatchStore(ctx context.Context, pairs []domain.UserOnlinePair) error {
	// user_online_hourly_stats
	{
		hourUserIDs, hourOnlineTimestamps := toHourlyStats(pairs)
		rowsAffected, err := s.db.Queries().UserOnlineHourlyStatsUpsert(ctx, dbs.UserOnlineHourlyStatsUpsertParams{
			UserIds:    hourUserIDs,
			CreatedAts: hourOnlineTimestamps,
		})
		if err != nil {
			return err
		}
		if rowsAffected == 0 {
			return nil
		}
	}

	// user_online_daily_stats
	dayUserIDs, dayOnlineTimestamps := toDailyStats(pairs)
	{
		rowsAffected, err := s.db.Queries().UserOnlineDailyStatsUpsert(ctx, dbs.UserOnlineDailyStatsUpsertParams{
			UserIds:    dayUserIDs,
			CreatedAts: dayOnlineTimestamps,
		})
		if err != nil {
			return err
		}
		if rowsAffected == 0 {
			return nil
		}
	}

	// user_online_daily_count_stats
	{
		err := s.db.Queries().UserOnlineDailyCountStatsUpsert(ctx, dayOnlineTimestamps)
		if err != nil {
			return err
		}
	}

	return nil
}
