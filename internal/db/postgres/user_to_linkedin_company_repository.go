package postgres

import (
	"context"
	"time"

	"github.com/readytotouch/readytotouch/internal/storage/postgres/dbs"
)

type UserToLinkedInCompanyRepository struct {
	db *Database
}

func NewUserToLinkedInCompanyRepository(db *Database) *UserToLinkedInCompanyRepository {
	return &UserToLinkedInCompanyRepository{db: db}
}

func (r *UserToLinkedInCompanyRepository) Upsert(ctx context.Context, userID int64, linkedinCompanyID int64, active bool, createdAt time.Time) error {
	return r.db.Queries().WipUserToLinkedInCompaniesUpsert(ctx, dbs.WipUserToLinkedInCompaniesUpsertParams{
		CreatedBy:         userID,
		LinkedinCompanyID: linkedinCompanyID,
		Active:            active,
		CreatedAt:         createdAt,
	})
}
