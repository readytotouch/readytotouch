package domain

import "golang.org/x/oauth2"

type OAuthProvider interface {
	SocialProvider() SocialProvider
	GetConfig() *oauth2.Config
	GetUser(accessToken string) (*OAuthProviderUserResponse, error)
}

type OAuthProviderUserResponse struct {
	ID       string
	Email    string
	Username string
	Name     string
}
