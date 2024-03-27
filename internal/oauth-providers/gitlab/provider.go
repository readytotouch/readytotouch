package gitlab

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"golang.org/x/oauth2"

	"github.com/readytotouch/readytotouch/internal/domain"
	pkgOAuthProviders "github.com/readytotouch/readytotouch/internal/oauth-providers"

	"github.com/valyala/fasthttp"
)

type GitlabOAuthProvider struct {
	config *oauth2.Config
}

func NewGitlabOAuthProvider(config *oauth2.Config) *GitlabOAuthProvider {
	return &GitlabOAuthProvider{config: config}
}

func (p *GitlabOAuthProvider) SocialProvider() domain.SocialProvider {
	return domain.SocialProviderGitlab
}

func (p *GitlabOAuthProvider) GetConfig() *oauth2.Config {
	return p.config
}

func (p *GitlabOAuthProvider) GetUser(accessToken string) (*domain.OAuthProviderUserResponse, error) {
	userResponse, userResponseErr := getUser(accessToken)
	if userResponseErr != nil {
		return nil, userResponseErr
	}

	if userResponse.ID == 0 {
		return nil, pkgOAuthProviders.ErrAbsentID
	}
	if userResponse.Username == "" {
		return nil, pkgOAuthProviders.ErrAbsentUsername
	}
	if userResponse.Email == "" {
		return nil, pkgOAuthProviders.ErrAbsentEmail
	}
	return &domain.OAuthProviderUserResponse{
		ID:       strconv.FormatUint(userResponse.ID, 10),
		Email:    userResponse.Email,
		Username: userResponse.Username,
		Name:     strings.TrimSpace(userResponse.Name),
	}, nil
}

func getUser(accessToken string) (*UserResponse, error) {
	var (
		request  = fasthttp.AcquireRequest()
		response = fasthttp.AcquireResponse()
	)

	defer func() {
		fasthttp.ReleaseRequest(request)
		fasthttp.ReleaseResponse(response)
	}()

	request.SetRequestURI(apiUserURI)
	request.URI().QueryArgs().Add("access_token", accessToken)

	httpErr := fasthttp.Do(request, response)
	if httpErr != nil {
		return nil, httpErr
	}
	if response.StatusCode() != http.StatusOK {
		return nil, pkgOAuthProviders.ErrUnexpectedHttpStatusCode
	}

	var result UserResponse
	unmarshalErr := json.Unmarshal(response.Body(), &result)
	if unmarshalErr != nil {
		return nil, unmarshalErr
	}
	return &result, nil
}
