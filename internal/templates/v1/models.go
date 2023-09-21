package v1

import "github.com/readytotouch-yaaws/yaaws-go/internal/domain"

const (
	appVersion = 1
)

type SocialUserProfile = domain.SocialUserProfile

func SocialUserProfileName(v SocialUserProfile) string {
	if v.Name == "" {
		return v.Username
	}

	return v.Name
}
