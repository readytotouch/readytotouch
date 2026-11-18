package postgres

import (
	"context"
	"database/sql"
	"strings"
	"time"

	"github.com/readytotouch/readytotouch/internal/domain"
	"github.com/readytotouch/readytotouch/internal/storage/postgres/dbs"
)

type UserRepository struct {
	db *Database
}

func NewUserRepository(db *Database) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) SocialUserProfilesByUser(ctx context.Context, id int64) ([]domain.SocialProviderUser, error) {
	rows, err := r.db.Queries().SocialUserProfilesByUser(ctx, id)
	if err != nil {
		return nil, err
	}

	result := make([]domain.SocialProviderUser, len(rows))
	for i, row := range rows {
		result[i] = domain.SocialProviderUser{
			SocialProvider:       domain.SocialProvider(row.SocialProvider),
			SocialProviderUserID: row.SocialProviderUserID,
			Username:             row.Username,
			Name:                 row.Name,
			CreatedAt:            row.CreatedAt,
		}
	}
	return result, nil
}

func (r *UserRepository) Create(ctx context.Context, params *domain.SocialProviderUserCreateParams, createdAt time.Time) (int64, error) {
	var (
		userID int64
	)

	txErr := r.db.WithTransaction(ctx, func(queries *dbs.Queries) error {
		prev, err := queries.UserSocialProfileGetByID(ctx, dbs.UserSocialProfileGetByIDParams{
			SocialProvider:       dbs.SocialProvider(params.SocialProvider),
			SocialProviderUserID: params.SocialProviderUserID,
		})
		switch err {
		case nil:
			userID = prev.UserID

			return r.update(ctx, queries, params, prev, createdAt)
		case sql.ErrNoRows:
			id, err := queries.UserSocialProfileCreate(ctx, dbs.UserSocialProfileCreateParams{
				Email:                params.Email,
				CreatedAt:            createdAt,
				SocialProvider:       dbs.SocialProvider(params.SocialProvider),
				SocialProviderUserID: params.SocialProviderUserID,
				Username:             params.Username,
				Name:                 params.Name,
			})
			if err != nil {
				return err
			}

			userID = id

			return nil
		default:
			return err
		}
	})
	if txErr != nil {
		return 0, txErr
	}

	return userID, nil
}

func (r *UserRepository) RegistrationDailyCountStats(
	ctx context.Context,
	from time.Time,
	to time.Time,
) ([]domain.TimeCountStats, error) {
	rows, err := r.db.Queries().UserRegistrationDailyCountStats(ctx, dbs.UserRegistrationDailyCountStatsParams{
		From: from,
		To:   to,
	})
	if err != nil {
		return nil, err
	}

	result := make([]domain.TimeCountStats, len(rows))
	for i, row := range rows {
		result[i] = domain.TimeCountStats{
			Time:  row.Day,
			Count: row.UserCount,
		}
	}
	return result, nil
}

func (r *UserRepository) SocialUserProfiles(ctx context.Context, socialProviders []dbs.SocialProvider, limit int32) ([]domain.SocialProviderUser, error) {
	rows, err := r.db.Queries().SocialUserProfiles(ctx, dbs.SocialUserProfilesParams{
		SocialProviderFilterExists: len(socialProviders) > 0,
		SocialProviders:            socialProviders,
		Limit:                      limit,
	})
	if err != nil {
		return nil, err
	}

	result := make([]domain.SocialProviderUser, len(rows))
	for i, row := range rows {
		result[i] = domain.SocialProviderUser{
			SocialProvider:       domain.SocialProvider(row.SocialProvider),
			SocialProviderUserID: row.SocialProviderUserID,
			Username:             row.Username,
			Name:                 row.Name,
			CreatedAt:            row.CreatedAt,
		}
	}
	return result, nil
}

func (r *UserRepository) update(ctx context.Context, queries *dbs.Queries, params *domain.SocialProviderUserCreateParams, prev dbs.UserSocialProfileGetByIDRow, now time.Time) error {
	sensitivitySame := strings.EqualFold(params.Email, prev.Email) &&
		params.Username == prev.Username &&
		params.Name == prev.Name
	if sensitivitySame {
		return nil
	}

	err := queries.UserSocialProfileUpdate(ctx, dbs.UserSocialProfileUpdateParams{
		Email:     params.Email,
		Username:  params.Username,
		Name:      params.Name,
		UpdatedAt: now,
		ID:        prev.ID,
	})
	if err != nil {
		return err
	}

	err = queries.UsersUpdate(ctx, dbs.UsersUpdateParams{
		UpdatedAt: now,
		ID:        prev.UserID,
	})
	if err != nil {
		return err
	}

	insensitivitySame := strings.EqualFold(params.Email, prev.Email) &&
		strings.EqualFold(params.Username, prev.Username) &&
		strings.EqualFold(params.Name, prev.Name)
	if insensitivitySame {
		return nil
	}

	err = queries.UserSocialProfileChangeHistoryNew(ctx, dbs.UserSocialProfileChangeHistoryNewParams{
		UserID:              prev.UserID,
		UserSocialProfileID: prev.ID,
		Email:               params.Email,
		Username:            params.Username,
		Name:                params.Name,
		CreatedAt:           now,
	})
	if err != nil {
		return err
	}

	return nil
}
