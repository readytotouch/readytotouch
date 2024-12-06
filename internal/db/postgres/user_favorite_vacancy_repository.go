package postgres

import (
	"context"
	"time"

	"github.com/readytotouch/readytotouch/internal/storage/postgres/dbs"
)

type UserFavoriteVacancyRepository struct {
	db *Database
}

func NewUserFavoriteVacancyRepository(db *Database) *UserFavoriteVacancyRepository {
	return &UserFavoriteVacancyRepository{db: db}
}

func (r *UserFavoriteVacancyRepository) Upsert(
	ctx context.Context,
	userID int64,
	vacancyID int64,
	favorite bool,
	createdAt time.Time,
) error {
	return r.db.Queries().UserFavoriteVacanciesUpsert(ctx, dbs.UserFavoriteVacanciesUpsertParams{
		UserID:    userID,
		VacancyID: vacancyID,
		Favorite:  favorite,
		CreatedAt: createdAt,
	})
}

func (r *UserFavoriteVacancyRepository) GetMap(
	ctx context.Context,
	userID int64,
	vacancyIDs []int64,
) (map[int64]bool, error) {
	rows, err := r.db.Queries().UserFavoriteVacancies(ctx, dbs.UserFavoriteVacanciesParams{
		UserID:                 userID,
		VacancyIdsFilterExists: len(vacancyIDs) > 0,
		VacancyIds:             vacancyIDs,
	})
	if err != nil {
		return nil, err
	}

	result := make(map[int64]bool, len(rows))
	for _, row := range rows {
		result[row] = true
	}

	return result, nil
}
