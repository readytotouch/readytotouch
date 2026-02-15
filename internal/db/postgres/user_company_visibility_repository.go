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

func (r *UserCompanyVisibilityRepository) List(
	ctx context.Context,
	userID int64,
) ([]int64, error) {
	return r.db.Queries().UserCompanyVisibilityOverrides(ctx, userID)
}
