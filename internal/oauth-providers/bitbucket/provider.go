package bitbucket

import (
	"encoding/json"
	"net/http"
	"strings"
	"sync"

	"golang.org/x/oauth2"

	"github.com/readytotouch/readytotouch/internal/domain"
	pkgOAuthProviders "github.com/readytotouch/readytotouch/internal/oauth-providers"

	"github.com/valyala/fasthttp"
)

type BitbucketOAuthProvider struct {
	config *oauth2.Config
}

func NewBitbucketOAuthProvider(config *oauth2.Config) *BitbucketOAuthProvider {
	return &BitbucketOAuthProvider{config: config}
}

func (p *BitbucketOAuthProvider) SocialProvider() domain.SocialProvider {
	return domain.SocialProviderBitbucket
}

func (p *BitbucketOAuthProvider) GetConfig() *oauth2.Config {
	return p.config
}

func (p *BitbucketOAuthProvider) GetUser(accessToken string) (*domain.OAuthProviderUserResponse, error) {
	var (
		wg                       sync.WaitGroup
		userResponse             *UserResponse
		userResponseErr          error
		userEmailListResponse    *UserEmailsResponse
		userEmailListResponseErr error
	)

	{
		wg.Add(1)
		go func() {
			defer wg.Done()

			userResponse, userResponseErr = getUser(accessToken)
		}()
	}

	userEmailListResponse, userEmailListResponseErr = getUserEmail(accessToken)
	if userEmailListResponseErr != nil {
		return nil, userEmailListResponseErr
	}

	verifyEmail, verifyEmailErr := getVerifyEmail(userEmailListResponse.Values)
	if verifyEmailErr != nil {
		return nil, verifyEmailErr
	}

	wg.Wait()

	if userResponseErr != nil {
		return nil, userResponseErr
	}

	if userResponse.UUID == "" {
		return nil, pkgOAuthProviders.ErrAbsentID
	}
	if userResponse.Username == "" {
		return nil, pkgOAuthProviders.ErrAbsentUsername
	}
	return &domain.OAuthProviderUserResponse{
		ID:       userResponse.UUID,
		Email:    verifyEmail,
		Username: userResponse.Username,
		Name:     strings.TrimSpace(userResponse.DisplayName),
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

func getUserEmail(accessToken string) (*UserEmailsResponse, error) {
	var (
		request  = fasthttp.AcquireRequest()
		response = fasthttp.AcquireResponse()
	)

	defer func() {
		fasthttp.ReleaseRequest(request)
		fasthttp.ReleaseResponse(response)
	}()

	request.SetRequestURI(apiUserEmailListURI)
	request.Header.Set("Authorization", "Bearer "+accessToken)

	httpErr := fasthttp.Do(request, response)
	if httpErr != nil {
		return nil, httpErr
	}
	if response.StatusCode() != http.StatusOK {
		return nil, pkgOAuthProviders.ErrUnexpectedHttpStatusCode
	}

	var result UserEmailsResponse
	unmarshalErr := json.Unmarshal(response.Body(), &result)
	if unmarshalErr != nil {
		return nil, unmarshalErr
	}
	return &result, nil
}

func getVerifyEmail(emails []UserEmail) (string, error) {
	for _, email := range emails {
		// just in case
		if email.Email == "" {
			continue
		}

		if email.IsPrimary && email.IsConfirmed {
			return email.Email, nil
		}
	}

	return "", pkgOAuthProviders.ErrAbsentEmail
}
