package linkedin

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

type Client struct {
	apiKey string
}

func NewClient(apiKey string) *Client {
	return &Client{apiKey: apiKey}
}

func (c *Client) CompaniesSearch(vanityName string) ([]Company, []byte, error) {
	if c.apiKey == "" {
		return nil, nil, fmt.Errorf("api key is required")
	}

	if vanityName == "" {
		return nil, nil, fmt.Errorf("vanity name is required")
	}

	var (
		path   = "https://api.linkedin.com/rest/organizations"
		values = url.Values{
			"q":          {"vanityName"},
			"vanityName": {vanityName},
		}
	)

	request, err := http.NewRequest("GET", path+"?"+values.Encode(), nil)
	if err != nil {
		// should be unreachable
		return nil, nil, err
	}

	request.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.apiKey))
	request.Header.Add("X-Restli-Protocol-Version", "2.0.0")
	request.Header.Add("Linkedin-Version", "202404")

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return nil, nil, err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return nil, nil, fmt.Errorf("unexpected status code: %d", response.StatusCode)
	}

	content, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to read response body: %w", err)
	}

	var parsedResponse CompaniesResponse
	if err := json.Unmarshal(content, &parsedResponse); err != nil {
		return nil, nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	return c.prepareCompaniesResponse(parsedResponse.Elements), content, nil
}

func (c *Client) prepareCompaniesResponse(source []CompanyResponse) []Company {
	companies := make([]Company, len(source))
	for i, element := range source {
		companies[i] = Company{
			ID:         element.ID,
			VanityName: element.VanityName,
			Name:       element.Name.Localized["en_US"],
		}
	}

	return companies
}
