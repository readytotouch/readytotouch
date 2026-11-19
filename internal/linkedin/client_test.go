package linkedin

import (
	"net/http"
	"testing"
	"time"

	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/require"
)

const (
	fixtureSuccess = `{"paging":{"start":0,"count":10,"links":[],"total":1},"elements":[{"vanityName":"readytotouch","localizedName":"ReadyToTouch","name":{"localized":{"en_US":"ReadyToTouch","uk_UA":"ReadyToTouch"},"preferredLocale":{"country":"US","language":"en"}},"primaryOrganizationType":"NONE","locations":[],"id":97909464,"localizedWebsite":"https://readytotouch.com/","logoV2":{"cropped":"urn:li:digitalmediaAsset:D4D0BAQFYYEr94PL59Q","original":"urn:li:digitalmediaAsset:D4D0BAQFYYEr94PL59Q","cropInfo":{"x":0,"width":0,"y":0,"height":0}}}]}`

	organizationsURL = "https://api.linkedin.com/v2/organizations?q=vanityName&vanityName=readytotouch"
)

var expectedCompanies = []Company{
	{
		ID:         97909464,
		VanityName: "readytotouch",
		Name:       "ReadyToTouch",
	},
}

func tokenResponder(accessToken string, calls *int) httpmock.Responder {
	return func(request *http.Request) (*http.Response, error) {
		*calls++

		if err := request.ParseForm(); err != nil {
			return nil, err
		}

		if request.PostForm.Get("grant_type") != "refresh_token" ||
			request.PostForm.Get("refresh_token") != "refresh" ||
			request.PostForm.Get("client_id") != "id" ||
			request.PostForm.Get("client_secret") != "secret" {
			return httpmock.NewStringResponse(400, `{"error":"invalid_request"}`), nil
		}

		return httpmock.NewStringResponse(200, `{"access_token":"`+accessToken+`","expires_in":5183999,"refresh_token":"refresh","refresh_token_expires_in":31534600}`), nil
	}
}

func authorizedResponder(token string) httpmock.Responder {
	return func(request *http.Request) (*http.Response, error) {
		if request.Header.Get("Authorization") != "Bearer "+token {
			return httpmock.NewStringResponse(401, `{}`), nil
		}

		return httpmock.NewStringResponse(200, fixtureSuccess), nil
	}
}

func refreshConfig() Config {
	return Config{ClientID: "id", ClientSecret: "secret", RefreshToken: "refresh"}
}

func TestClientNoCredentials(t *testing.T) {
	_, _, err := NewClient(Config{}).CompaniesSearch("readytotouch")
	require.Error(t, err)
}

func TestClientRefreshesTokenOnceAndCachesIt(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	var refreshCalls int
	httpmock.RegisterResponder(http.MethodPost, tokenURL, tokenResponder("fresh", &refreshCalls))
	httpmock.RegisterResponder(http.MethodGet, organizationsURL, authorizedResponder("fresh"))

	client := NewClient(refreshConfig())

	for i := 0; i < 3; i++ {
		response, _, err := client.CompaniesSearch("readytotouch")
		require.NoError(t, err)
		require.Equal(t, expectedCompanies, response)
	}

	require.Equal(t, 1, refreshCalls)
}

func TestClientRefreshesTokenCloseToExpiration(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	var refreshCalls int
	httpmock.RegisterResponder(http.MethodPost, tokenURL, tokenResponder("fresh", &refreshCalls))
	httpmock.RegisterResponder(http.MethodGet, organizationsURL, authorizedResponder("fresh"))

	now := time.Date(2026, 11, 18, 0, 0, 0, 0, time.UTC)
	client := NewClient(refreshConfig())
	client.now = func() time.Time { return now }

	_, _, err := client.CompaniesSearch("readytotouch")
	require.NoError(t, err)
	require.Equal(t, 1, refreshCalls)

	// ~59 days later: still valid for more than refreshBefore
	now = now.Add(54 * 24 * time.Hour)
	_, _, err = client.CompaniesSearch("readytotouch")
	require.NoError(t, err)
	require.Equal(t, 1, refreshCalls)

	// less than refreshBefore left
	now = now.Add(2 * 24 * time.Hour)
	_, _, err = client.CompaniesSearch("readytotouch")
	require.NoError(t, err)
	require.Equal(t, 2, refreshCalls)
}

func TestClientRetriesOnceOnUnauthorized(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	var refreshCalls int
	httpmock.RegisterResponder(http.MethodPost, tokenURL, tokenResponder("fresh", &refreshCalls))
	httpmock.RegisterResponder(http.MethodGet, organizationsURL, authorizedResponder("fresh"))

	client := NewClient(refreshConfig())
	// a token which looks valid, but was revoked
	client.token = "revoked"
	client.tokenExpiresAt = time.Now().Add(30 * 24 * time.Hour)

	response, _, err := client.CompaniesSearch("readytotouch")
	require.NoError(t, err)
	require.Equal(t, expectedCompanies, response)
	require.Equal(t, 1, refreshCalls)
}

func TestClientUnauthorizedAfterRefresh(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	var refreshCalls int
	httpmock.RegisterResponder(http.MethodPost, tokenURL, tokenResponder("fresh", &refreshCalls))
	httpmock.RegisterResponder(http.MethodGet, organizationsURL, authorizedResponder("other"))

	_, _, err := NewClient(refreshConfig()).CompaniesSearch("readytotouch")
	require.EqualError(t, err, "unexpected status code: 401")
	require.Equal(t, 2, refreshCalls)
}

func TestClientRefreshFailure(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder(http.MethodPost, tokenURL, httpmock.NewStringResponder(400, `{"error":"invalid_grant"}`))

	_, _, err := NewClient(refreshConfig()).CompaniesSearch("readytotouch")
	require.EqualError(t, err, "failed to refresh token: unexpected status code: 400")
}

func TestClientRefreshRequiresClientCredentials(t *testing.T) {
	_, _, err := NewClient(Config{RefreshToken: "refresh"}).CompaniesSearch("readytotouch")
	require.Error(t, err)
}
