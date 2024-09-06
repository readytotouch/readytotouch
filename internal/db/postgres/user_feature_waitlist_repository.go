package postgres

import (
	"context"
	"database/sql"
	"time"

	"github.com/readytotouch/readytotouch/internal/domain"
	"github.com/readytotouch/readytotouch/internal/storage/postgres/dbs"
)

type UserFeatureWaitlistRepository struct {
	db *Database
}

func NewUserFeatureWaitlistRepository(db *Database) *UserFeatureWaitlistRepository {
	return &UserFeatureWaitlistRepository{db: db}
}

func (r *UserFeatureWaitlistRepository) Upsert(
	ctx context.Context,
	userID int64,
	feature dbs.FeatureWait,
	active bool,
	createdAt time.Time,
) error {
	txErr := r.db.WithTransaction(ctx, func(queries *dbs.Queries) error {
		affected, err := queries.UserFeatureWaitlistUpsert(ctx, dbs.UserFeatureWaitlistUpsertParams{
			UserID:    userID,
			Feature:   feature,
			Active:    active,
			CreatedAt: createdAt,
		})
		if err != nil {
			return err
		}
		if affected == 0 {
			return nil
		}

		var diff int64
		if active {
			diff = 1 // +1
		} else {
			diff = -1
		}

		err = queries.UserFeatureWaitlistStatsUpsert(ctx, dbs.UserFeatureWaitlistStatsUpsertParams{
			Feature:   feature,
			UserCount: diff,
		})
		if err != nil {
			return err
		}

		return nil
	})
	if txErr != nil {
		return txErr
	}

	return nil
}

func (r *UserFeatureWaitlistRepository) SubscribedState(
	ctx context.Context,
	userID int64,
	feature dbs.FeatureWait,
) (bool, error) {
	active, err := r.db.Queries().UserFeatureWaitlist(ctx, dbs.UserFeatureWaitlistParams{
		UserID:  userID,
		Feature: feature,
	})
	if err == sql.ErrNoRows {
		return false, nil
	}
	if err != nil {
		return false, err
	}

	return active, nil
}

func (r *UserFeatureWaitlistRepository) DailyCountStats(
	ctx context.Context,
	feature dbs.FeatureWait,
	from time.Time,
	to time.Time,
) ([]domain.TimeCountStats, error) {
	rows, err := r.db.Queries().UserFeatureWaitlistDailyStats(ctx, dbs.UserFeatureWaitlistDailyStatsParams{
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
			Count: row.SubscriberCount,
		}
	}
	return result, nil
}

func (r *UserFeatureWaitlistRepository) TotalCount(ctx context.Context, feature dbs.FeatureWait) (int64, error) {
	count, err := r.db.Queries().UserFeatureWaitlistStats(ctx, feature)
	if err == sql.ErrNoRows {
		return 0, nil
	}
	if err != nil {
		return 0, err
	}

	return count, nil
}
