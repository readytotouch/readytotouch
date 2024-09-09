package postgres

import (
	"context"
	"database/sql"
	"time"

	"github.com/readytotouch/readytotouch/internal/domain"
	"github.com/readytotouch/readytotouch/internal/storage/postgres/dbs"
)

type FeatureViewStatsRepository struct {
	db *Database
}

func NewFeatureViewStatsRepository(db *Database) *FeatureViewStatsRepository {
	return &FeatureViewStatsRepository{db: db}
}

func (r *FeatureViewStatsRepository) Upsert(ctx context.Context, feature dbs.FeatureWait, createdAt time.Time) error {
	err := r.db.Queries().FeatureViewDailyStatsUpsert(ctx, dbs.FeatureViewDailyStatsUpsertParams{
		Feature:   feature,
		CreatedAt: createdAt.Truncate(24 * time.Hour),
	})
	if err != nil {
		return err
	}

	err = r.db.Queries().FeatureViewStatsUpsert(ctx, feature)
	if err != nil {
		return err
	}

	return nil
}

func (r *FeatureViewStatsRepository) DailyCountStats(
	ctx context.Context,
	feature dbs.FeatureWait,
	from time.Time,
	to time.Time,
) ([]domain.TimeCountStats, error) {
	rows, err := r.db.Queries().FeatureViewDailyStats(ctx, dbs.FeatureViewDailyStatsParams{
		From:    from,
		To:      to,
		Feature: feature,
	})
	if err != nil {
		return nil, err
	}

	result := make([]domain.TimeCountStats, len(rows))
	for i, row := range rows {
		result[i] = domain.TimeCountStats{
			Time:  row.Day,
			Count: row.ViewCount,
		}
	}
	return result, nil
}

func (r *FeatureViewStatsRepository) TotalCount(ctx context.Context, feature dbs.FeatureWait) (int64, error) {
	count, err := r.db.Queries().FeatureViewStats(ctx, feature)
	if err == sql.ErrNoRows {
		return 0, nil
	}
	if err != nil {
		return 0, err
	}

	return count, nil
}
