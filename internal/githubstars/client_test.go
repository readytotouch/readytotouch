package githubstars

import (
	"context"
	"net/http"
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
)

func TestGetStargazersCount(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	owner := "readytotouch"
	repo := "readytotouch"
	expectedStars := 100

	httpmock.RegisterResponder("GET", "https://api.github.com/repos/readytotouch/readytotouch",
		httpmock.NewStringResponder(http.StatusOK, `{"stargazers_count": 100}`))

	client := NewClient()
	stars, err := client.GetStargazersCount(context.Background(), owner, repo)

	assert.NoError(t, err)
	assert.Equal(t, expectedStars, stars)
}
