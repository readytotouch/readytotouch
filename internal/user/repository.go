package user

import (
	"context"
	"time"

	"github.com/readytotouch-yaaws/yaaws-go/internal/domain"
	"github.com/readytotouch-yaaws/yaaws-go/internal/storage/postgres"
	"github.com/readytotouch-yaaws/yaaws-go/internal/storage/postgres/dbs"

	"github.com/jackc/pgx/v5/pgtype"
)

type Repository struct {
	db *postgres.Database
}

func NewRepository(db *postgres.Database) *Repository {
	return &Repository{db: db}
}

func (r *Repository) RegistrationDailyCountStats(
	ctx context.Context,
	from time.Time,
	to time.Time,
) ([]domain.TimeCountStats, error) {
	rows, err := r.db.Queries().UserRegistrationDailyCountStats(ctx, dbs.UserRegistrationDailyCountStatsParams{
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

	result := make([]domain.TimeCountStats, len(rows))
	for i, row := range rows {
		result[i] = domain.TimeCountStats{
			Time:  row.Day.Time,
			Count: row.UserCount,
		}
	}
	return result, nil
}
