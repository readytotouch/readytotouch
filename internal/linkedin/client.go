package linkedin

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"sync"
	"time"
)

type Client struct {
	httpClient *http.Client
	now        func() time.Time

	mu             sync.Mutex
	config         Config
	token          string
	tokenExpiresAt time.Time
}

func NewClient(config Config) *Client {
	return &Client{
		httpClient: http.DefaultClient,
		now:        time.Now,
		config:     config,
	}
}

func (c *Client) CompaniesSearch(vanityName string) ([]Company, []byte, error) {
	if vanityName == "" {
		return nil, nil, fmt.Errorf("vanity name is required")
	}

	var (
		path   = "https://api.linkedin.com/v2/organizations"
		values = url.Values{
			"q":          {"vanityName"},
			"vanityName": {vanityName},
		}
	)

	response, err := c.get(path + "?" + values.Encode())
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

// get performs an authorized GET request. On 401 it refreshes the token once and retries.
func (c *Client) get(rawURL string) (*http.Response, error) {
	c.mu.Lock()
	defer c.mu.Unlock()

	for attempt := 0; ; attempt++ {
		token, err := c.accessToken(c.now())
		if err != nil {
			return nil, err
		}

		request, err := http.NewRequest("GET", rawURL, nil)
		if err != nil {
			// should be unreachable
			return nil, err
		}

		request.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))
		request.Header.Add("X-Restli-Protocol-Version", "2.0.0")
		request.Header.Add("Linkedin-Version", "202404")

		response, err := c.httpClient.Do(request)
		if err != nil {
			return nil, err
		}

		if response.StatusCode == http.StatusUnauthorized && attempt == 0 {
			response.Body.Close()
			c.invalidate()

			continue
		}

		return response, nil
	}
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
