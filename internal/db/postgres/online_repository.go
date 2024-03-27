package postgres

import (
	"context"
	"time"

	"github.com/readytotouch/readytotouch/internal/domain"
	"github.com/readytotouch/readytotouch/internal/storage/postgres/dbs"
)

type OnlineRepository struct {
	db *Database
}

func NewOnlineRepository(db *Database) *OnlineRepository {
	return &OnlineRepository{db: db}
}

func (r *OnlineRepository) DailyCountStats(
	ctx context.Context,
	from time.Time,
	to time.Time,
) ([]domain.TimeCountStats, error) {
	rows, err := r.db.Queries().UserOnlineDailyCountStats(ctx, dbs.UserOnlineDailyCountStatsParams{
		From: from,
		To:   to,
	})
	if err != nil {
		return nil, err
	}

	result := make([]domain.TimeCountStats, len(rows))
	for i, row := range rows {
		result[i] = domain.TimeCountStats{
			Time:  row.Day,
			Count: row.UserCount,
		}
	}
	return result, nil
}
