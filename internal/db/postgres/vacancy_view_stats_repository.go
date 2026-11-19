package postgres

import (
	"context"
	"time"

	"github.com/readytotouch/readytotouch/internal/storage/postgres/dbs"
)

type VacancyViewStatsRepository struct {
	db *Database
}

func NewVacancyViewStatsRepository(db *Database) *VacancyViewStatsRepository {
	return &VacancyViewStatsRepository{db: db}
}

func (r *VacancyViewStatsRepository) Stats(ctx context.Context, vacancyIDs []int64, from time.Time) (map[int64]int64, error) {
	rows, err := r.db.Queries().VacancyViewStats(ctx, dbs.VacancyViewStatsParams{
		VacancyIds: vacancyIDs,
		From:       from,
	})
	if err != nil {
		return nil, err
	}

	result := make(map[int64]int64, len(rows))
	for _, row := range rows {
		result[row.VacancyID] = row.Count
	}
	return result, nil
}

// Upsert counts a view of the vacancy by the user, once per hour and once per day.
// It is a single statement, so it is atomic and does not need an explicit transaction.
func (r *VacancyViewStatsRepository) Upsert(ctx context.Context, vacancyID int64, userID int64, createdAt time.Time) error {
	return r.db.Queries().VacancyViewUpsert(ctx, dbs.VacancyViewUpsertParams{
		VacancyID: vacancyID,
		Hour:      createdAt.Truncate(time.Hour),
		UserID:    userID,
		Day:       createdAt.Truncate(24 * time.Hour),
	})
}
