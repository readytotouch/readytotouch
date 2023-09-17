package postgres

import (
	"context"

	"github.com/readytotouch-yaaws/yaaws-go/internal/storage/postgres"
	"github.com/readytotouch-yaaws/yaaws-go/internal/storage/postgres/dbs"
)

type BatchUpdateOnlineStorage struct {
	repository *postgres.Database
}

func NewBatchUpdateOnlineStorage(repository *postgres.Database) *BatchUpdateOnlineStorage {
	return &BatchUpdateOnlineStorage{repository: repository}
}

func (s *BatchUpdateOnlineStorage) BatchStore(ctx context.Context, pairs []UserOnlinePair) error {
	// user_online_hourly_stats
	{
		hourUserIDs, hourOnlineTimestamps := toHourlyStats(pairs)
		rowsAffected, err := s.repository.Queries().UserOnlineHourlyStatsUpsert(ctx, dbs.UserOnlineHourlyStatsUpsertParams{
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
	{
		dayUserIDs, dayOnlineTimestamps := toDailyStats(pairs)
		rowsAffected, err := s.repository.Queries().UserOnlineDailyStatsUpsert(ctx, dbs.UserOnlineDailyStatsUpsertParams{
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
		err := s.repository.Queries().UserOnlineDailyCountStatsUpsert(ctx, toDailyMin(pairs))
		if err != nil {
			return err
		}
	}

	return nil
}
