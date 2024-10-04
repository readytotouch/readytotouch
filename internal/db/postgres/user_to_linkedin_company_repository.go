package postgres

import (
	"context"
	"database/sql"
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

func (r *UserToLinkedInCompanyRepository) GetByVanityName(ctx context.Context, name string) (int64, bool, error) {
	id, err := r.db.Queries().WipLinkedInCompaniesGetByVanityName(ctx, name)
	if err == sql.ErrNoRows {
		return 0, false, nil
	}
	if err != nil {
		return 0, false, err
	}

	return id, true, nil
}

func (r *UserToLinkedInCompanyRepository) ExistsVanityName(ctx context.Context, name string) (bool, error) {
	return r.db.Queries().WipLinkedInCompanyRequestHistoryExistsVanityName(ctx, name)
}

func (r *UserToLinkedInCompanyRepository) GetRequestHistoryCount(ctx context.Context, userID int64) (int64, error) {
	return r.db.Queries().WipLinkedInCompanyRequestHistoryCount(ctx, userID)
}

func (r *UserToLinkedInCompanyRepository) CreateCompany(
	ctx context.Context,
	linkedinCompanyID int64,
	linkedinCompanyVanityName string,
	linkedinCompanyName string,
	responsePayload []byte,
	createdBy int64,
	createdAt time.Time,
) error {
	txErr := r.db.WithTransaction(ctx, func(queries *dbs.Queries) error {
		if linkedinCompanyID > 0 {
			err := queries.WipLinkedInCompaniesNew(ctx, dbs.WipLinkedInCompaniesNewParams{
				ID:         linkedinCompanyID,
				VanityName: linkedinCompanyVanityName,
				Name:       linkedinCompanyName,
				CreatedAt:  createdAt,
				CreatedBy:  createdBy,
			})
			if err != nil {
				return err
			}
		}

		err := queries.WipLinkedInCompanyRequestHistoryNew(ctx, dbs.WipLinkedInCompanyRequestHistoryNewParams{
			VanityName: linkedinCompanyVanityName,
			LinkedinCompanyID: sql.NullInt64{
				Int64: linkedinCompanyID,
				Valid: linkedinCompanyID > 0,
			},
			ResponsePayload: responsePayload,
			CreatedAt:       createdAt,
			CreatedBy:       createdBy,
		})
		if err != nil {
			return err
		}

		return nil
	})
	if txErr != nil {
		return txErr
	}

	return nil
}
