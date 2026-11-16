package linkedin

import (
	"net/http"
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/require"
)

const (
	fixtureSuccess = `{"paging":{"start":0,"count":10,"links":[],"total":1},"elements":[{"vanityName":"readytotouch","localizedName":"ReadyToTouch","name":{"localized":{"en_US":"ReadyToTouch","uk_UA":"ReadyToTouch"},"preferredLocale":{"country":"US","language":"en"}},"primaryOrganizationType":"NONE","locations":[],"id":97909464,"localizedWebsite":"https://readytotouch.com/","logoV2":{"cropped":"urn:li:digitalmediaAsset:D4D0BAQFYYEr94PL59Q","original":"urn:li:digitalmediaAsset:D4D0BAQFYYEr94PL59Q","cropInfo":{"x":0,"width":0,"y":0,"height":0}}}]}`
)

func TestClientSuccess(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder(
		http.MethodGet,
		"https://api.linkedin.com/v2/organizations?q=vanityName&vanityName=readytotouch",
		httpmock.NewStringResponder(
			200,
			fixtureSuccess,
		),
	)

	client := NewClient("secret_dev_team")

	response, payload, err := client.CompaniesSearch("readytotouch")
	require.NoError(t, err)
	require.Equal(t, []Company{
		{
			ID:         97909464,
			VanityName: "readytotouch",
			Name:       "ReadyToTouch",
		},
	}, response)
	require.Equal(t, []byte(fixtureSuccess), payload)
}
