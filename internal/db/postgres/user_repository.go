package postgres

import (
	"context"
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

// Create signs in the social profile: creates it (and the user if needed) or updates it if it has changed.
// It is a single statement, so it is atomic and does not need an explicit transaction.
//
// Concurrent first sign-ins of the same profile are not handled: the unique index makes the second one fail.
func (r *UserRepository) Create(ctx context.Context, params *domain.SocialProviderUserCreateParams, createdAt time.Time) (int64, error) {
	return r.db.Queries().UserSocialProfileUpsert(ctx, dbs.UserSocialProfileUpsertParams{
		SocialProvider:       dbs.SocialProvider(params.SocialProvider),
		SocialProviderUserID: params.SocialProviderUserID,
		Email:                params.Email,
		Username:             params.Username,
		Name:                 params.Name,
		CreatedAt:            createdAt,
	})
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
