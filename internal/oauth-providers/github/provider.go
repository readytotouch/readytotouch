package github

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
	"sync"

	"golang.org/x/oauth2"

	"github.com/readytotouch/readytotouch/internal/domain"
	pkgOAuthProviders "github.com/readytotouch/readytotouch/internal/oauth-providers"

	"github.com/valyala/fasthttp"
)

type GithubOAuthProvider struct {
	config *oauth2.Config
}

func NewGithubOAuthProvider(config *oauth2.Config) *GithubOAuthProvider {
	return &GithubOAuthProvider{config: config}
}

func (p *GithubOAuthProvider) SocialProvider() domain.SocialProvider {
	return domain.SocialProviderGithub
}

func (p *GithubOAuthProvider) GetConfig() *oauth2.Config {
	return p.config
}

func (p *GithubOAuthProvider) GetUser(accessToken string) (*domain.OAuthProviderUserResponse, error) {
	var (
		wg                       sync.WaitGroup
		userResponse             *UserResponse
		userResponseErr          error
		userEmailListResponse    []UserEmail
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

	verifyEmail, verifyEmailErr := getVerifyEmail(userEmailListResponse)
	if verifyEmailErr != nil {
		return nil, verifyEmailErr
	}

	wg.Wait()

	if userResponseErr != nil {
		return nil, userResponseErr
	}

	if userResponse.ID == 0 {
		return nil, pkgOAuthProviders.ErrAbsentID
	}
	if userResponse.Login == "" {
		return nil, pkgOAuthProviders.ErrAbsentUsername
	}
	return &domain.OAuthProviderUserResponse{
		ID:       strconv.FormatUint(userResponse.ID, 10),
		Email:    verifyEmail,
		Username: userResponse.Login,
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
	request.Header.Set("Authorization", "token "+accessToken)

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

func getUserEmail(accessToken string) ([]UserEmail, error) {
	var (
		request  = fasthttp.AcquireRequest()
		response = fasthttp.AcquireResponse()
	)

	defer func() {
		fasthttp.ReleaseRequest(request)
		fasthttp.ReleaseResponse(response)
	}()

	request.SetRequestURI(apiUserEmailListURI)
	request.Header.Set("Authorization", "token "+accessToken)

	httpErr := fasthttp.Do(request, response)
	if httpErr != nil {
		return nil, httpErr
	}
	if response.StatusCode() != http.StatusOK {
		return nil, pkgOAuthProviders.ErrUnexpectedHttpStatusCode
	}

	// main email and @users.noreply.github.com
	var result []UserEmail
	unmarshalErr := json.Unmarshal(response.Body(), &result)
	if unmarshalErr != nil {
		return nil, unmarshalErr
	}
	return result, nil
}

func getVerifyEmail(emails []UserEmail) (string, error) {
	for _, email := range emails {
		// just in case
		if email.Email == "" {
			continue
		}

		if email.Primary && email.Verified {
			return email.Email, nil
		}
	}

	return "", pkgOAuthProviders.ErrAbsentEmail
}
