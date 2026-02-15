package postgres

import (
	"context"
	"time"

	"github.com/readytotouch/readytotouch/internal/storage/postgres/dbs"
)

type UserIndustryVisibilityRepository struct {
	db *Database
}

func NewUserIndustryVisibilityRepository(db *Database) *UserIndustryVisibilityRepository {
	return &UserIndustryVisibilityRepository{db: db}
}

func (r *UserIndustryVisibilityRepository) Upsert(
	ctx context.Context,
	userID int64,
	industryID int64,
	visibility bool,
	createdAt time.Time,
) error {
	return r.db.WithTransaction(ctx, func(queries *dbs.Queries) error {
		affected, err := queries.UserIndustryVisibilityOverrideUpsert(ctx, dbs.UserIndustryVisibilityOverrideUpsertParams{
			UserID:     userID,
			IndustryID: industryID,
			Visibility: visibility,
			CreatedAt:  createdAt,
		})
		if err != nil {
			return err
		}

		if affected == 0 {
			return nil
		}

		return queries.UserIndustryVisibilityOverrideHistoryNew(ctx, dbs.UserIndustryVisibilityOverrideHistoryNewParams{
			UserID:     userID,
			IndustryID: industryID,
			Visibility: visibility,
			CreatedAt:  createdAt,
		})
	})
}
