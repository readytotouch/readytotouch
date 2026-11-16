package google

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/readytotouch/readytotouch/internal/domain"
	pkgOAuthProviders "github.com/readytotouch/readytotouch/internal/oauth-providers"

	"github.com/valyala/fasthttp"
	"golang.org/x/oauth2"
)

var (
	defaultClient = &fasthttp.Client{
		ReadBufferSize: 16 * 1024,
	}
)

type GoogleOAuthProvider struct {
	config *oauth2.Config
}

func NewGoogleOAuthProvider(config *oauth2.Config) *GoogleOAuthProvider {
	return &GoogleOAuthProvider{config: config}
}

func (p *GoogleOAuthProvider) SocialProvider() domain.SocialProvider {
	return domain.SocialProviderGoogle
}

func (p *GoogleOAuthProvider) GetConfig() *oauth2.Config {
	return p.config
}

func (p *GoogleOAuthProvider) GetUser(accessToken string) (*domain.OAuthProviderUserResponse, error) {
	userResponse, userResponseErr := getUser(accessToken)
	if userResponseErr != nil {
		return nil, userResponseErr
	}

	if userResponse.Sub == "" {
		return nil, pkgOAuthProviders.ErrAbsentID
	}

	if userResponse.Name == "" {
		return nil, pkgOAuthProviders.ErrAbsentName
	}

	if userResponse.Email == "" {
		return nil, pkgOAuthProviders.ErrAbsentEmail
	}

	if !userResponse.EmailVerified {
		return nil, pkgOAuthProviders.ErrUnverifiedEmail
	}

	return &domain.OAuthProviderUserResponse{
		ID:       userResponse.Sub,
		Email:    userResponse.Email,
		Username: "",
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
	request.Header.Set("Authorization", "Bearer "+accessToken)

	httpErr := defaultClient.Do(request, response)
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
