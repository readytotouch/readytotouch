package domain

import (
	"time"
)

type SocialProvider string

const (
	SocialProviderGithub    SocialProvider = "github"
	SocialProviderGitlab    SocialProvider = "gitlab"
	SocialProviderBitbucket SocialProvider = "bitbucket"
	SocialProviderFigma     SocialProvider = "figma"
	SocialProviderDribbble  SocialProvider = "dribbble"
	SocialProviderBehance   SocialProvider = "behance"
	SocialProviderLinkedin  SocialProvider = "linkedin"
	SocialProviderXing      SocialProvider = "xing"
	SocialProviderGoogle    SocialProvider = "google"
)

type SocialUserProfile struct {
	ID                   int64
	SocialProvider       SocialProvider
	SocialProviderUserID string
	Username             string
	Name                 string
	CreatedAt            time.Time
}
