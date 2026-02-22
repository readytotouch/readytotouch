package postgres

import (
	"context"
	"time"

	"github.com/readytotouch/readytotouch/internal/storage/postgres/dbs"
)

type UserCompanyVisibilityRepository struct {
	db *Database
}

func NewUserCompanyVisibilityRepository(db *Database) *UserCompanyVisibilityRepository {
	return &UserCompanyVisibilityRepository{db: db}
}

func (r *UserCompanyVisibilityRepository) Upsert(
	ctx context.Context,
	userID int64,
	companyID int64,
	visibility bool,
	createdAt time.Time,
) error {
	return r.db.WithTransaction(ctx, func(queries *dbs.Queries) error {
		affected, err := queries.UserCompanyVisibilityOverrideUpsert(ctx, dbs.UserCompanyVisibilityOverrideUpsertParams{
			UserID:     userID,
			CompanyID:  companyID,
			Visibility: visibility,
			CreatedAt:  createdAt,
		})
		if err != nil {
			return err
		}

		if affected == 0 {
			return nil
		}

		return queries.UserCompanyVisibilityOverrideHistoryNew(ctx, dbs.UserCompanyVisibilityOverrideHistoryNewParams{
			UserID:     userID,
			CompanyID:  companyID,
			Visibility: visibility,
			CreatedAt:  createdAt,
		})
	})
}

func (r *UserCompanyVisibilityRepository) GetMap(
	ctx context.Context,
	userID int64,
) (map[int64]bool, error) {
	if userID == 0 {
		return nil, nil
	}

	rows, err := r.db.Queries().UserCompanyVisibilityOverrides(ctx, userID)
	if err != nil {
		return nil, err
	}

	result := make(map[int64]bool, len(rows))
	for _, row := range rows {
		result[row] = true
	}

	return result, nil
}
