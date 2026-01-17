package githubstars

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type Client struct {
}

func NewClient() *Client {
	return &Client{}
}

type repositoryResponse struct {
	StargazersCount int `json:"stargazers_count"`
}

func (c *Client) GetStargazersCount(ctx context.Context, owner, repo string) (int, error) {
	url := fmt.Sprintf("https://api.github.com/repos/%s/%s", owner, repo)

	request, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return 0, err
	}

	request.Header.Set("User-Agent", "readytotouch-app")

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return 0, err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return 0, fmt.Errorf("unexpected status code: %d", response.StatusCode)
	}

	var parsedResponse repositoryResponse
	if err := json.NewDecoder(response.Body).Decode(&parsedResponse); err != nil {
		return 0, fmt.Errorf("failed to decode response: %w", err)
	}

	return parsedResponse.StargazersCount, nil
}
