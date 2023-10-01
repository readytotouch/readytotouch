package github

type UserResponse struct {
	ID    uint64 `json:"id"`
	Login string `json:"login"`
	Name  string `json:"name"`
	// NodeID            string      `json:"node_id"`
	// AvatarURL         string      `json:"avatar_url"`
	// GravatarID        string      `json:"gravatar_id"`
	// URL               string      `json:"url"`
	// HTMLURL           string      `json:"html_url"`
	// FollowersURL      string      `json:"followers_url"`
	// FollowingURL      string      `json:"following_url"`
	// GistsURL          string      `json:"gists_url"`
	// StarredURL        string      `json:"starred_url"`
	// SubscriptionsURL  string      `json:"subscriptions_url"`
	// OrganizationsURL  string      `json:"organizations_url"`
	// ReposURL          string      `json:"repos_url"`
	// EventsURL         string      `json:"events_url"`
	// ReceivedEventsURL string      `json:"received_events_url"`
	// Type              string      `json:"type"`
	// SiteAdmin         bool        `json:"site_admin"`
	// Company           interface{} `json:"company"`
	// Blog              string      `json:"blog"`
	// Location          string      `json:"location"`
	// Email             interface{} `json:"email"`
	// Hireable          interface{} `json:"hireable"`
	// Bio               interface{} `json:"bio"`
	// TwitterUsername   interface{} `json:"twitter_username"`
	// PublicRepos       int         `json:"public_repos"`
	// PublicGists       int         `json:"public_gists"`
	// Followers         int         `json:"followers"`
	// Following         int         `json:"following"`
	// CreatedAt         time.Time   `json:"created_at"`
	// UpdatedAt         time.Time   `json:"updated_at"`
}

type UserEmail struct {
	Email    string `json:"email"`
	Primary  bool   `json:"primary"`
	Verified bool   `json:"verified"`
	// Visibility string `json:"visibility"`
}
