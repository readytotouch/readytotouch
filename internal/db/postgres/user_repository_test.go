package postgres

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/readytotouch/readytotouch/internal/domain"
	"github.com/readytotouch/readytotouch/internal/env"

	"github.com/stretchr/testify/require"
)

func TestUserRepositoryCreate(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}

	pgConnection, err := Connection(env.Required("POSTGRES_DSN"))
	require.NoError(t, err)
	defer pgConnection.Close()

	database, err := NewDatabase(pgConnection)
	require.NoError(t, err)
	defer database.Queries().Close()

	repository := NewUserRepository(database)

	var (
		ctx    = context.Background()
		now    = time.Now().UTC()
		suffix = time.Now().UnixNano()
		email  = fmt.Sprintf("user-repository-test-%d@example.com", suffix)

		userIDs []int64
	)

	defer func() {
		for _, id := range userIDs {
			_, err := pgConnection.ExecContext(ctx, `DELETE FROM user_social_profile_change_history WHERE user_id = $1`, id)
			require.NoError(t, err)
			_, err = pgConnection.ExecContext(ctx, `DELETE FROM user_social_profiles WHERE user_id = $1`, id)
			require.NoError(t, err)
			_, err = pgConnection.ExecContext(ctx, `DELETE FROM users WHERE id = $1`, id)
			require.NoError(t, err)
		}
	}()

	count := func(query string, args ...any) int {
		var n int
		require.NoError(t, pgConnection.QueryRowContext(ctx, query, args...).Scan(&n))

		return n
	}

	github := &domain.SocialProviderUserCreateParams{
		SocialProvider:       domain.SocialProviderGithub,
		SocialProviderUserID: fmt.Sprintf("github-%d", suffix),
		Email:                email,
		Username:             "test-username",
		Name:                 "Test Name",
	}

	t.Run("new user and new profile", func(t *testing.T) {
		userID, err := repository.Create(ctx, github, now)
		require.NoError(t, err)
		require.NotZero(t, userID)
		userIDs = append(userIDs, userID)

		require.Equal(t, 1, count(`SELECT COUNT(*) FROM users WHERE id = $1`, userID))
		require.Equal(t, 1, count(`SELECT COUNT(*) FROM user_social_profiles WHERE user_id = $1`, userID))
		require.Equal(t, 1, count(`SELECT COUNT(*) FROM user_social_profile_change_history WHERE user_id = $1`, userID))

		var (
			gotProvider, gotProviderUserID, gotEmail, gotUsername, gotName string
		)
		require.NoError(t, pgConnection.QueryRowContext(
			ctx,
			`SELECT social_provider::TEXT, social_provider_user_id, email, username, name FROM user_social_profiles WHERE user_id = $1`,
			userID,
		).Scan(&gotProvider, &gotProviderUserID, &gotEmail, &gotUsername, &gotName))
		require.Equal(t, string(github.SocialProvider), gotProvider)
		require.Equal(t, github.SocialProviderUserID, gotProviderUserID)
		require.Equal(t, github.Email, gotEmail)
		require.Equal(t, github.Username, gotUsername)
		require.Equal(t, github.Name, gotName)
	})

	t.Run("same profile again is idempotent", func(t *testing.T) {
		first, err := repository.Create(ctx, github, now)
		require.NoError(t, err)
		second, err := repository.Create(ctx, github, now.Add(time.Minute))
		require.NoError(t, err)

		require.Equal(t, first, second)
		require.Equal(t, 1, count(`SELECT COUNT(*) FROM user_social_profiles WHERE user_id = $1`, first))
		require.Equal(t, 1, count(`SELECT COUNT(*) FROM user_social_profile_change_history WHERE user_id = $1`, first))
	})

	t.Run("another provider with the same email links to the existing user", func(t *testing.T) {
		existing, err := repository.Create(ctx, github, now)
		require.NoError(t, err)

		gitlab := &domain.SocialProviderUserCreateParams{
			SocialProvider:       domain.SocialProviderGitlab,
			SocialProviderUserID: fmt.Sprintf("gitlab-%d", suffix),
			Email:                email,
			Username:             "test-username-gitlab",
			Name:                 "Test Name GitLab",
		}

		userID, err := repository.Create(ctx, gitlab, now)
		require.NoError(t, err)

		require.Equal(t, existing, userID)
		require.Equal(t, 2, count(`SELECT COUNT(*) FROM user_social_profiles WHERE user_id = $1`, userID))
		require.Equal(t, 2, count(`SELECT COUNT(*) FROM user_social_profile_change_history WHERE user_id = $1`, userID))
	})

	t.Run("another provider with another email creates another user", func(t *testing.T) {
		existing, err := repository.Create(ctx, github, now)
		require.NoError(t, err)

		google := &domain.SocialProviderUserCreateParams{
			SocialProvider:       domain.SocialProviderGoogle,
			SocialProviderUserID: fmt.Sprintf("google-%d", suffix),
			Email:                fmt.Sprintf("other-%d@example.com", suffix),
			Username:             "test-username-google",
			Name:                 "Test Name Google",
		}

		userID, err := repository.Create(ctx, google, now)
		require.NoError(t, err)
		userIDs = append(userIDs, userID)

		require.NotEqual(t, existing, userID)
		require.Equal(t, 1, count(`SELECT COUNT(*) FROM user_social_profiles WHERE user_id = $1`, userID))
	})

	t.Run("changed profile is updated and recorded in the history", func(t *testing.T) {
		existing, err := repository.Create(ctx, github, now)
		require.NoError(t, err)

		changed := *github
		changed.Name = "New Test Name"

		userID, err := repository.Create(ctx, &changed, now.Add(time.Hour))
		require.NoError(t, err)

		require.Equal(t, existing, userID)
		require.Equal(t, 3, count(`SELECT COUNT(*) FROM user_social_profile_change_history WHERE user_id = $1`, userID))
	})
}
