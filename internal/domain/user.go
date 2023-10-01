package domain

import (
	"time"
)

const (
	RegistrationHistoryLimit = 32
)

type SocialProviderUser struct {
	SocialProvider       SocialProvider
	SocialProviderUserID string
	Username             string
	Name                 string
	CreatedAt            time.Time
}

type SocialProviderUserCreateParams struct {
	SocialProvider       SocialProvider
	SocialProviderUserID string
	Email                string
	Username             string
	Name                 string
}
