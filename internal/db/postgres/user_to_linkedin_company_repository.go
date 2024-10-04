package postgres

import (
	"context"
	"time"

	"github.com/readytotouch/readytotouch/internal/domain"
	"github.com/readytotouch/readytotouch/internal/storage/postgres/dbs"
)

type UserToLinkedInCompanyRepository struct {
	db *Database
}

func NewUserToLinkedInCompanyRepository(db *Database) *UserToLinkedInCompanyRepository {
	return &UserToLinkedInCompanyRepository{db: db}
}

func (r *UserToLinkedInCompanyRepository) List(ctx context.Context, userID int64) ([]domain.LinkedInProfileResponse, error) {
	rows, err := r.db.Queries().WipUserLinkedInCompanies(ctx, userID)
	if err != nil {
		return nil, err
	}

	result := make([]domain.LinkedInProfileResponse, len(rows))
	for i, row := range rows {
		result[i] = domain.LinkedInProfileResponse{
			ID:    row.ID,
			Alias: row.VanityName,
			Name:  row.Name,
		}
	}
	return result, nil
}

func (r *UserToLinkedInCompanyRepository) Add(ctx context.Context, userID int64, linkedinCompanyID int64, createdAt time.Time) error {
	return r.db.Queries().WipUserToLinkedInCompaniesAdd(ctx, dbs.WipUserToLinkedInCompaniesAddParams{
		CreatedBy:         userID,
		LinkedinCompanyID: linkedinCompanyID,
		CreatedAt:         createdAt,
	})
}

func (r *UserToLinkedInCompanyRepository) Delete(ctx context.Context, userID int64, linkedinCompanyID int64, updatedAt time.Time) error {
	return r.db.Queries().WipUserToLinkedInCompaniesDelete(ctx, dbs.WipUserToLinkedInCompaniesDeleteParams{
		UpdatedAt:         updatedAt,
		UpdatedBy:         userID,
		UserID:            userID,
		LinkedinCompanyID: linkedinCompanyID,
	})
}
