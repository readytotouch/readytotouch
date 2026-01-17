package postgres

import (
	"context"
	"testing"
	"time"

	"github.com/readytotouch/readytotouch/internal/env"
	"github.com/stretchr/testify/assert"
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

	ctx := context.Background()
	owner := "test_owner"
	repo := "test_repo"
	stars := 42
	now := time.Now().Truncate(time.Microsecond).UTC()

	// Test Upsert
	err = repository.Upsert(ctx, owner, repo, stars, now)
	require.NoError(t, err)

	// Test ListAll
	repos, err := repository.ListAll(ctx)
	require.NoError(t, err)

	found := false
	for _, r := range repos {
		if r.Owner == owner && r.Repo == repo {
			assert.Equal(t, stars, r.StargazersCount)
			// Postgres might have slightly different time precision, but Truncate should help
			assert.WithinDuration(t, now, r.UpdatedAt.UTC(), time.Second)
			found = true
			break
		}
	}
	assert.True(t, found, "Repository not found in ListAll")

	// Test Update via Upsert
	newStars := 100
	newNow := time.Now().Truncate(time.Microsecond).UTC()
	err = repository.Upsert(ctx, owner, repo, newStars, newNow)
	require.NoError(t, err)

	repos, err = repository.ListAll(ctx)
	require.NoError(t, err)

	found = false
	for _, r := range repos {
		if r.Owner == owner && r.Repo == repo {
			assert.Equal(t, newStars, r.StargazersCount)
			assert.WithinDuration(t, newNow, r.UpdatedAt.UTC(), time.Second)
			found = true
			break
		}
	}
	assert.True(t, found, "Updated repository not found in ListAll")
}
