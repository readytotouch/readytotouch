package v1

import (
	"fmt"

	"github.com/readytotouch/readytotouch/internal/domain"
)

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

func SocialProviderUserHeaderPhoto(v SocialProviderUser) string {
	switch v.SocialProvider {
	case github:
		return fmt.Sprintf("https://avatars.githubusercontent.com/u/%s?v=4&s=48", v.SocialProviderUserID)
	case gitlab:
		return fmt.Sprintf("https://gitlab.com/uploads/-/system/user/avatar/%s/avatar.png?width=48", v.SocialProviderUserID)
	default:
		return ""
	}
}
