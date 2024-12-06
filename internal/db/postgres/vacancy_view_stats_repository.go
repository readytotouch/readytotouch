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

func (r *VacancyViewStatsRepository) Upsert(ctx context.Context, vacancyID int64, userID int64, createdAt time.Time) error {
	var (
		hour = createdAt.Truncate(time.Hour)
		day  = createdAt.Truncate(24 * time.Hour)
	)

	return r.db.WithTransaction(ctx, func(queries *dbs.Queries) error {
		{
			affected, err := queries.UserVacancyViewHourlyStatsUpsert(ctx, dbs.UserVacancyViewHourlyStatsUpsertParams{
				VacancyID: vacancyID,
				CreatedAt: hour,
				UserID:    userID,
			})
			if err != nil {
				return err
			}

			if affected == 0 {
				return nil
			}
		}

		{
			affected, err := queries.UserVacancyViewDailyStatsUpsert(ctx, dbs.UserVacancyViewDailyStatsUpsertParams{
				VacancyID: vacancyID,
				CreatedAt: day,
				UserID:    userID,
			})
			if err != nil {
				return err
			}

			if affected == 0 {
				return nil
			}
		}

		err := queries.VacancyViewDailyStatsUpsert(ctx, dbs.VacancyViewDailyStatsUpsertParams{
			VacancyID: vacancyID,
			CreatedAt: day,
		})
		if err != nil {
			return err
		}

		return nil
	})
}
