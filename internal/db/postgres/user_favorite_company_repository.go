package postgres

import (
	"context"
	"database/sql"
	"time"

	"github.com/readytotouch/readytotouch/internal/storage/postgres/dbs"
)

type UserFavoriteCompanyRepository struct {
	db *Database
}

func NewUserFavoriteCompanyRepository(db *Database) *UserFavoriteCompanyRepository {
	return &UserFavoriteCompanyRepository{db: db}
}

func (r *UserFavoriteCompanyRepository) Upsert(
	ctx context.Context,
	userID int64,
	companyID int64,
	favorite bool,
	createdAt time.Time,
) error {
	return r.db.Queries().UserFavoriteCompaniesUpsert(ctx, dbs.UserFavoriteCompaniesUpsertParams{
		UserID:    userID,
		CompanyID: companyID,
		Favorite:  favorite,
		CreatedAt: createdAt,
	})
}

func (r *UserFavoriteCompanyRepository) GetMap(
	ctx context.Context,
	userID int64,
	companyIDs []int64,
) (map[int64]bool, error) {
	rows, err := r.db.Queries().UserFavoriteCompanies(ctx, dbs.UserFavoriteCompaniesParams{
		UserID:                 userID,
		CompanyIdsFilterExists: len(companyIDs) > 0,
		CompanyIds:             companyIDs,
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

func (r *UserFavoriteCompanyRepository) Stats(ctx context.Context, companyID int64, from time.Time) (int64, int64, error) {
	row, err := r.db.Queries().UserFavoriteCompaniesStats(ctx, dbs.UserFavoriteCompaniesStatsParams{
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
