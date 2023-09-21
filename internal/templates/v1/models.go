package v1

import "github.com/readytotouch-yaaws/yaaws-go/internal/domain"

const (
	appVersion = 1
)

const (
	github = domain.SocialProviderGithub
	gitlab = domain.SocialProviderGitlab
)

type SocialUserProfile = domain.SocialUserProfile

func SocialUserProfileName(v SocialUserProfile) string {
	if v.Name == "" {
		return v.Username
	}

	return v.Name
}
