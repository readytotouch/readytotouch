package githubstars

import (
	"context"
	"net/http"
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/require"
)

func TestGetStargazersCount(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	const (
		owner         = "readytotouch"
		repo          = "readytotouch"
		expectedStars = 100
	)

	httpmock.RegisterResponder(
		"GET",
		"https://api.github.com/repos/readytotouch/readytotouch",
		httpmock.NewStringResponder(http.StatusOK, `{"stargazers_count":100}`),
	)

	stars, err := GetStargazersCount(context.Background(), owner, repo)

	require.NoError(t, err)
	require.Equal(t, int32(expectedStars), stars)
}
