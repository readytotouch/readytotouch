package v1

import "github.com/readytotouch-yaaws/yaaws-go/internal/domain"

const (
	appVersion = 1
)

const (
	github = domain.SocialProviderGithub
	gitlab = domain.SocialProviderGitlab
)

// SocialProviderUser reference
type SocialProviderUser = domain.SocialProviderUser

func SocialProviderUserName(v SocialProviderUser) string {
	if v.Name == "" {
		return v.Username
	}

	return v.Name
}
