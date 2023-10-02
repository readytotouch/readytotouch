package bitbucket

type UserResponse struct {
	UUID        string `json:"uuid"`
	Username    string `json:"username"`
	DisplayName string `json:"display_name"`
}

type UserEmailsResponse struct {
	Values []UserEmail `json:"values"`
}

type UserEmail struct {
	Email       string `json:"email"`
	IsPrimary   bool   `json:"is_primary"`
	IsConfirmed bool   `json:"is_confirmed"`
}
