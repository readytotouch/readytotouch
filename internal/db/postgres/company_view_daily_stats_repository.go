package postgres

import (
	"context"
	"database/sql"
	"time"

	"github.com/readytotouch/readytotouch/internal/domain"
	"github.com/readytotouch/readytotouch/internal/storage/postgres/dbs"
)

type CompanyViewDailyStatsRepository struct {
	db *Database
}

func NewCompanyViewDailyStatsRepository(db *Database) *CompanyViewDailyStatsRepository {
	return &CompanyViewDailyStatsRepository{db: db}
}

func (r *CompanyViewDailyStatsRepository) Upsert(ctx context.Context, companyID int64, createdAt time.Time) error {
	err := r.db.Queries().CompanyViewDailyStatsUpsert(ctx, dbs.CompanyViewDailyStatsUpsertParams{
		CompanyID: companyID,
		CreatedAt: createdAt.Truncate(24 * time.Hour),
	})
	if err != nil {
		return err
	}

	return nil
}

func (r *CompanyViewDailyStatsRepository) DailyCountStats(
	ctx context.Context,
	companyID int64,
	from time.Time,
	to time.Time,
) ([]domain.TimeCountStats, error) {
	rows, err := r.db.Queries().CompanyViewDailyStats(ctx, dbs.CompanyViewDailyStatsParams{
		From:      from,
		To:        to,
		CompanyID: companyID,
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

func (r *CompanyViewDailyStatsRepository) Stats(ctx context.Context, companyID int64, from time.Time) (int64, int64, error) {
	row, err := r.db.Queries().CompanyTotalViews(ctx, dbs.CompanyTotalViewsParams{
		From:      from,
		CompanyID: companyID,
	})
	if err == sql.ErrNoRows {
		return 0, 0, nil
	}
	if err != nil {
		return 0, 0, err
	}

	return row.Count, row.CountSince, nil
}
