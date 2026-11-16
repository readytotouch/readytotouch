package postgres

import (
	"context"
	"slices"
	"testing"
	"time"

	"github.com/readytotouch/readytotouch/internal/domain"
	"github.com/readytotouch/readytotouch/internal/env"

	"github.com/stretchr/testify/require"
)

func TestGithubRepositoryStarsRepository(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}

	pgConnection, err := Connection(env.Required("POSTGRES_DSN"))
	require.NoError(t, err)
	defer pgConnection.Close()

	database, err := NewDatabase(pgConnection)
	require.NoError(t, err)
	defer database.Queries().Close()

	repository := NewGithubRepositoryStarsRepository(database)

	const (
		owner    = "test_owner"
		repo     = "test_repo"
		stars    = 8096
		newStars = stars + 100
	)

	var (
		ctx = context.Background()
		now = time.Now().UTC()
	)

	// First upsert
	{
		err = repository.Upsert(ctx, owner, repo, stars, now)
		require.NoError(t, err)

		rows, err := repository.All(ctx)
		require.NoError(t, err)

		require.True(
			t,
			slices.Contains(rows, domain.GithubRepositoryStar{
				Owner:           owner,
				Repo:            repo,
				StargazersCount: stars,
			}),
		)
	}

	// Second upsert with updated stars
	{
		err = repository.Upsert(ctx, owner, repo, newStars, now)
		require.NoError(t, err)

		rows, err := repository.All(ctx)
		require.NoError(t, err)

		require.True(
			t,
			slices.Contains(rows, domain.GithubRepositoryStar{
				Owner:           owner,
				Repo:            repo,
				StargazersCount: newStars,
			}),
		)

		actualStars, err := repository.Get(ctx, owner, repo)
		require.NoError(t, err)
		require.Equal(t, newStars, actualStars)
	}
}
