package linkedin

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
)

const (
	tokenURL = "https://www.linkedin.com/oauth/v2/accessToken"

	// Refresh the access token in advance, so a request never goes out with a token about to expire.
	refreshBefore = 5 * 24 * time.Hour
)

// Config describes how the client authenticates.
//
// The access token is obtained from RefreshToken and renewed automatically.
//
// LinkedIn does not extend the refresh token lifetime (~365 days) on refresh,
// so it must be re-issued manually once a year.
type Config struct {
	ClientID     string
	ClientSecret string
	RefreshToken string
}

type tokenResponse struct {
	AccessToken  string `json:"access_token"`
	ExpiresIn    int64  `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
}

// accessToken returns a valid access token, refreshing it if needed. Must be called with c.mu held.
func (c *Client) accessToken(now time.Time) (string, error) {
	if c.token != "" && now.Add(refreshBefore).Before(c.tokenExpiresAt) {
		return c.token, nil
	}

	return c.refresh(now)
}

func (c *Client) refresh(now time.Time) (string, error) {
	if c.config.ClientID == "" || c.config.ClientSecret == "" || c.config.RefreshToken == "" {
		return "", fmt.Errorf("client id, client secret and refresh token are required")
	}

	form := url.Values{
		"grant_type":    {"refresh_token"},
		"refresh_token": {c.config.RefreshToken},
		"client_id":     {c.config.ClientID},
		"client_secret": {c.config.ClientSecret},
	}

	request, err := http.NewRequest(http.MethodPost, tokenURL, strings.NewReader(form.Encode()))
	if err != nil {
		// should be unreachable
		return "", err
	}
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	response, err := c.httpClient.Do(request)
	if err != nil {
		return "", fmt.Errorf("failed to refresh token: %w", err)
	}
	defer response.Body.Close()

	content, err := io.ReadAll(response.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read refresh response body: %w", err)
	}

	if response.StatusCode != http.StatusOK {
		// the body is intentionally not included: it is not needed to diagnose and may be sensitive
		return "", fmt.Errorf("failed to refresh token: unexpected status code: %d", response.StatusCode)
	}

	var parsed tokenResponse
	if err := json.Unmarshal(content, &parsed); err != nil {
		return "", fmt.Errorf("failed to unmarshal refresh response: %w", err)
	}

	if parsed.AccessToken == "" {
		return "", fmt.Errorf("refresh response has no access token")
	}

	c.token = parsed.AccessToken
	c.tokenExpiresAt = now.Add(time.Duration(parsed.ExpiresIn) * time.Second)
	if parsed.RefreshToken != "" {
		c.config.RefreshToken = parsed.RefreshToken
	}

	return c.token, nil
}

// invalidate drops the cached access token, so the next call refreshes it. Must be called with c.mu held.
func (c *Client) invalidate() {
	c.token = ""
	c.tokenExpiresAt = time.Time{}
}
