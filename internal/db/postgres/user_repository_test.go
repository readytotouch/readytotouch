package postgres

import (
	"context"
	"fmt"
	"strings"
	"sync"
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
		ctx     = context.Background()
		now     = time.Now().UTC().Truncate(time.Microsecond)
		runID   = time.Now().UnixNano()
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

	updatedAt := func(userID int64) time.Time {
		var at time.Time
		require.NoError(t, pgConnection.QueryRowContext(ctx, `SELECT updated_at FROM users WHERE id = $1`, userID).Scan(&at))

		return at
	}

	historyCount := func(userID int64) int {
		return count(`SELECT COUNT(*) FROM user_social_profile_change_history WHERE user_id = $1`, userID)
	}

	// params returns unique params for every subtest, so the subtests do not depend on each other
	params := func(tag string, provider domain.SocialProvider) *domain.SocialProviderUserCreateParams {
		return &domain.SocialProviderUserCreateParams{
			SocialProvider:       provider,
			SocialProviderUserID: fmt.Sprintf("%s-%s-%d", provider, tag, runID),
			Email:                fmt.Sprintf("user-repository-test-%s-%s-%d@example.com", tag, provider, runID),
			Username:             "test-username",
			Name:                 "Test Name",
		}
	}

	// create creates the user and registers it for the cleanup
	create := func(p *domain.SocialProviderUserCreateParams, at time.Time) int64 {
		userID, err := repository.Create(ctx, p, at)
		require.NoError(t, err)
		require.NotZero(t, userID)
		userIDs = append(userIDs, userID)

		return userID
	}

	t.Run("new user and new profile", func(t *testing.T) {
		p := params("new", domain.SocialProviderGithub)

		userID := create(p, now)

		require.Equal(t, 1, count(`SELECT COUNT(*) FROM users WHERE id = $1`, userID))
		require.Equal(t, 1, count(`SELECT COUNT(*) FROM user_social_profiles WHERE user_id = $1`, userID))
		require.Equal(t, 1, historyCount(userID))

		var gotProvider, gotProviderUserID, gotEmail, gotUsername, gotName string
		require.NoError(t, pgConnection.QueryRowContext(
			ctx,
			`SELECT social_provider::TEXT, social_provider_user_id, email, username, name FROM user_social_profiles WHERE user_id = $1`,
			userID,
		).Scan(&gotProvider, &gotProviderUserID, &gotEmail, &gotUsername, &gotName))
		require.Equal(t, string(p.SocialProvider), gotProvider)
		require.Equal(t, p.SocialProviderUserID, gotProviderUserID)
		require.Equal(t, p.Email, gotEmail)
		require.Equal(t, p.Username, gotUsername)
		require.Equal(t, p.Name, gotName)
		require.WithinDuration(t, now, updatedAt(userID), 0)
	})

	t.Run("same profile again changes nothing", func(t *testing.T) {
		p := params("same", domain.SocialProviderGithub)

		first := create(p, now)
		second, err := repository.Create(ctx, p, now.Add(time.Minute))
		require.NoError(t, err)

		require.Equal(t, first, second)
		require.Equal(t, 1, count(`SELECT COUNT(*) FROM user_social_profiles WHERE user_id = $1`, first))
		require.Equal(t, 1, historyCount(first))
		require.WithinDuration(t, now, updatedAt(first), 0)
	})

	t.Run("another provider with the same email links to the existing user", func(t *testing.T) {
		github := params("link", domain.SocialProviderGithub)
		existing := create(github, now)

		gitlab := params("link", domain.SocialProviderGitlab)
		gitlab.Email = github.Email

		userID, err := repository.Create(ctx, gitlab, now)
		require.NoError(t, err)

		require.Equal(t, existing, userID)
		require.Equal(t, 2, count(`SELECT COUNT(*) FROM user_social_profiles WHERE user_id = $1`, userID))
		require.Equal(t, 2, historyCount(userID))
	})

	t.Run("another provider with another email creates another user", func(t *testing.T) {
		github := params("other", domain.SocialProviderGithub)
		existing := create(github, now)

		google := params("other", domain.SocialProviderGoogle)

		userID := create(google, now)

		require.NotEqual(t, existing, userID)
		require.Equal(t, 1, count(`SELECT COUNT(*) FROM user_social_profiles WHERE user_id = $1`, userID))
	})

	t.Run("changed profile is updated and recorded in the history", func(t *testing.T) {
		p := params("changed", domain.SocialProviderGithub)
		existing := create(p, now)

		changed := *p
		changed.Name = "New Test Name"
		later := now.Add(time.Hour)

		userID, err := repository.Create(ctx, &changed, later)
		require.NoError(t, err)

		require.Equal(t, existing, userID)
		require.Equal(t, 2, historyCount(userID))
		require.Equal(t, 1, count(`SELECT COUNT(*) FROM user_social_profiles WHERE user_id = $1 AND name = $2`, userID, changed.Name))
		require.WithinDuration(t, later, updatedAt(userID), 0)
	})

	t.Run("changes only in the email case are ignored", func(t *testing.T) {
		p := params("email-case", domain.SocialProviderGithub)
		existing := create(p, now)

		changed := *p
		changed.Email = strings.ToUpper(p.Email)

		userID, err := repository.Create(ctx, &changed, now.Add(time.Hour))
		require.NoError(t, err)

		require.Equal(t, existing, userID)
		require.WithinDuration(t, now, updatedAt(userID), 0)
		require.Equal(t, 1, historyCount(userID))
		require.Equal(t, 1, count(`SELECT COUNT(*) FROM user_social_profiles WHERE user_id = $1 AND email = $2`, userID, p.Email))
	})

	t.Run("changes only in the username case are updated, but not recorded in the history", func(t *testing.T) {
		p := params("username-case", domain.SocialProviderGithub)
		existing := create(p, now)

		changed := *p
		changed.Username = strings.ToUpper(p.Username)
		later := now.Add(time.Hour)

		userID, err := repository.Create(ctx, &changed, later)
		require.NoError(t, err)

		require.Equal(t, existing, userID)
		require.Equal(t, 1, historyCount(userID))
		require.Equal(t, 1, count(`SELECT COUNT(*) FROM user_social_profiles WHERE user_id = $1 AND username = $2`, userID, changed.Username))
		require.WithinDuration(t, later, updatedAt(userID), 0)
	})

	t.Run("updated email is stored in lower case, the history keeps it as is", func(t *testing.T) {
		p := params("email-update", domain.SocialProviderGithub)
		existing := create(p, now)

		changed := *p
		changed.Email = fmt.Sprintf("Mixed-Case-%d@Example.com", runID)

		userID, err := repository.Create(ctx, &changed, now.Add(time.Hour))
		require.NoError(t, err)
		require.Equal(t, existing, userID)

		require.Equal(t, 1, count(`SELECT COUNT(*) FROM user_social_profiles WHERE user_id = $1 AND email = $2`, userID, strings.ToLower(changed.Email)))
		require.Equal(t, 1, count(`SELECT COUNT(*) FROM user_social_profile_change_history WHERE user_id = $1 AND email = $2`, userID, changed.Email))
	})

	t.Run("concurrent sign-ins with a changed profile are recorded once", func(t *testing.T) {
		p := params("concurrent", domain.SocialProviderGithub)
		existing := create(p, now)

		changed := *p
		changed.Name = "Concurrent Name"

		const workers = 8

		var (
			wg      sync.WaitGroup
			userIDs = make([]int64, workers)
			errs    = make([]error, workers)
		)
		for i := 0; i < workers; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				userIDs[i], errs[i] = repository.Create(ctx, &changed, now.Add(time.Hour))
			}()
		}
		wg.Wait()

		for i := 0; i < workers; i++ {
			require.NoError(t, errs[i])
			require.Equal(t, existing, userIDs[i])
		}
		require.Equal(t, 2, historyCount(existing))
	})
}
