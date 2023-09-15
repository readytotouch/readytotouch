package main

import (
	"context"

	postgresql "github.com/readytotouch-yaaws/yaaws-go/internal/storage/postgres"
	"github.com/readytotouch-yaaws/yaaws-go/internal/storage/postgres/dbs"
)

type BatchUpdateOnlineStorage struct {
	repository *postgresql.Repository
}

func NewBatchUpdateOnlineStorage(repository *postgresql.Repository) *BatchUpdateOnlineStorage {
	return &BatchUpdateOnlineStorage{repository: repository}
}

func (s *BatchUpdateOnlineStorage) BatchStore(ctx context.Context, pairs []UserOnlinePair) error {
	// user_online_hourly_stats
	{
		hourUserIDs, hourOnlineTimestamps := toHourlyStats(pairs)
		rowsAffected, err := s.repository.Queries().UserOnlineHourlyStatsUpsert(ctx, dbs.UserOnlineHourlyStatsUpsertParams{
			UserIds: hourUserIDs,
			Onlines: hourOnlineTimestamps,
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
			UserIds: dayUserIDs,
			Onlines: dayOnlineTimestamps,
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
