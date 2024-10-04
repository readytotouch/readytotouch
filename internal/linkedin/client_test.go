package linkedin

import (
	"net/http"
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/require"
)

const (
	fixtureSuccess = `{"paging":{"start":0,"count":10,"links":[]},"elements":[{"vanityName":"readytotouch","website":{"localized":{"en_US":"https://readytotouch.com/"},"preferredLocale":{"country":"US","language":"en"}},"localizedName":"ReadyToTouch","foundedOn":{"year":2020},"created":{"actor":"urn:li:person:qwmBoAitPD","time":1691066065447},"groups":[],"description":{"localized":{"en_US":"Yet another anonymous work search","uk_UA":"Ще один анонімний пошук роботи"},"preferredLocale":{"country":"US","language":"en"}},"versionTag":"1564428138","defaultLocale":{"country":"US","language":"en"},"organizationType":"NON_PROFIT","alternativeNames":[],"specialties":[],"localizedSpecialties":[],"name":{"localized":{"en_US":"ReadyToTouch","uk_UA":"ReadyToTouch"},"preferredLocale":{"country":"US","language":"en"}},"primaryOrganizationType":"NONE","locations":[],"lastModified":{"actor":"urn:li:person:qwmBoAitPD","time":1719792410009},"id":97909464,"localizedDescription":"Yet another anonymous work search","autoCreated":false,"localizedWebsite":"https://readytotouch.com/","logoV2":{"cropped":"urn:li:digitalmediaAsset:D4D0BAQFYYEr94PL59Q","original":"urn:li:digitalmediaAsset:D4D0BAQFYYEr94PL59Q","cropInfo":{"x":0,"width":0,"y":0,"height":0}}}]}`
)

func TestClientSuccess(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder(
		http.MethodGet,
		"https://api.linkedin.com/rest/organizations?q=vanityName&vanityName=readytotouch",
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
