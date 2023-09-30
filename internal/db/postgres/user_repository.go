package postgres

import (
	"context"
	"strings"
	"time"

	"github.com/readytotouch-yaaws/yaaws-go/internal/domain"
	"github.com/readytotouch-yaaws/yaaws-go/internal/storage/postgres/dbs"

	"github.com/jackc/pgx/v5/pgtype"
)

type UserRepository struct {
	db *Database
}

func NewUserRepository(db *Database) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) Upsert(ctx context.Context, params *domain.SocialProviderUserCreateParams) (int64, error) {
	var (
		userID int64
		now    = time.Now().Truncate(time.Second).UTC()
	)

	txErr := r.db.WithTransaction(ctx, func(queries *dbs.Queries) error {
		row, err := queries.UserSocialProfileGet(ctx, dbs.UserSocialProfileGetParams{
			SocialProvider:       dbs.SocialProvider(params.SocialProvider),
			SocialProviderUserID: params.SocialProviderUserID,
		})
		switch err {
		case nil:
			userID = row.UserID

			return r.update(ctx, queries, params, row, now)
		}

		return nil
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
		From: pgtype.Date{
			Time:  from,
			Valid: true,
		},
		To: pgtype.Date{
			Time:  to,
			Valid: true,
		},
	})
	if err != nil {
		return nil, err
	}

	result := make([]domain.TimeCountStats, len(rows))
	for i, row := range rows {
		result[i] = domain.TimeCountStats{
			Time:  row.Day.Time,
			Count: row.UserCount,
		}
	}
	return result, nil
}

func (r *UserRepository) SocialUserProfiles(ctx context.Context, limit int32) ([]domain.SocialProviderUser, error) {
	rows, err := r.db.Queries().SocialUserProfiles(ctx, limit)
	if err != nil {
		return nil, err
	}

	result := make([]domain.SocialProviderUser, len(rows))
	for i, row := range rows {
		result[i] = domain.SocialProviderUser{
			ID:                   row.ID,
			SocialProvider:       domain.SocialProvider(row.SocialProvider),
			SocialProviderUserID: row.SocialProviderUserID,
			Username:             row.Username,
			Name:                 row.Name,
			CreatedAt:            row.CreatedAt.Time,
		}
	}
	return result, nil
}

func (r *UserRepository) update(ctx context.Context, queries *dbs.Queries, params *domain.SocialProviderUserCreateParams, row dbs.UserSocialProfileGetRow, now time.Time) error {
	sensitivitySame := params.Email == row.Email &&
		params.Username == row.Username &&
		params.Name == row.Name
	if sensitivitySame {
		return nil
	}

	err := queries.UserSocialProfileUpdate(ctx, dbs.UserSocialProfileUpdateParams{
		Email:    row.Email,
		Username: row.Username,
		Name:     row.Name,
		UpdatedAt: pgtype.Timestamp{
			Time:             now,
			InfinityModifier: 0,
			Valid:            true,
		},
		ID: row.ID,
	})
	if err != nil {
		return err
	}

	insensitivitySame := strings.EqualFold(params.Email, row.Email) &&
		strings.EqualFold(params.Username, row.Username) &&
		strings.EqualFold(params.Name, row.Name)
	if insensitivitySame {
		return nil
	}

	err = queries.UserSocialProfileChangeHistoryNew(ctx, dbs.UserSocialProfileChangeHistoryNewParams{
		UserID:              row.UserID,
		UserSocialProfileID: row.ID,
		Email:               params.Email,
		Username:            params.Username,
		Name:                params.Name,
		CreatedAt: pgtype.Timestamp{
			Time:             now,
			InfinityModifier: 0,
			Valid:            true,
		},
	})
	if err != nil {
		return err
	}

	return nil
}
