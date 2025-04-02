package v1

import (
	"fmt"

	"github.com/readytotouch/readytotouch/internal/domain"
)

const (
	appVersion = 17
)

const (
	github    = domain.SocialProviderGithub
	gitlab    = domain.SocialProviderGitlab
	bitbucket = domain.SocialProviderBitbucket
)

// SocialProviderUser reference
type SocialProviderUser = domain.SocialProviderUser

func SocialProviderUserName(v SocialProviderUser) string {
	if v.Name == "" {
		return v.Username
	}

	return v.Name
}

func SocialProviderUserHeaderPhoto(vs []SocialProviderUser) string {
	for _, v := range vs {
		switch v.SocialProvider {
		case github:
			return fmt.Sprintf("https://avatars.githubusercontent.com/u/%s?v=4&s=48", v.SocialProviderUserID)
		case gitlab:

			// GitLab outdated API
			// return fmt.Sprintf("https://gitlab.com/uploads/-/system/user/avatar/%s/avatar.png?width=48", v.SocialProviderUserID)

		case bitbucket:
			// Unknown
		}
	}

	return ""
}
