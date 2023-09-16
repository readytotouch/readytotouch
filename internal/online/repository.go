package online

import (
	"context"
	"time"

	"github.com/readytotouch-yaaws/yaaws-go/internal/domain"
	"github.com/readytotouch-yaaws/yaaws-go/internal/storage/postgres"
	"github.com/readytotouch-yaaws/yaaws-go/internal/storage/postgres/dbs"

	"github.com/jackc/pgx/v5/pgtype"
)

type Repository struct {
	db *postgres.Repository
}

func NewRepository(repository *postgres.Repository) *Repository {
	return &Repository{db: repository}
}

func (r *Repository) DailyCountStats(
	ctx context.Context,
	from time.Time,
	to time.Time,
) ([]domain.UserOnlineDailyCountStats, error) {
	rows, err := r.db.Queries().UserOnlineDailyCountStats(ctx, dbs.UserOnlineDailyCountStatsParams{
		From: pgtype.Date{
			Time:  from,
			Valid: true,
		},
		To: pgtype.Date{
			Time:  to,
			Valid: true,
		},
	})
	if err != nil {
		return nil, err
	}

	result := make([]domain.UserOnlineDailyCountStats, len(rows))
	for i, row := range rows {
		result[i] = domain.UserOnlineDailyCountStats{
			Online:    row.Online.Time.Unix(),
			UserCount: row.UserCount,
		}
	}
	return result, nil
}
