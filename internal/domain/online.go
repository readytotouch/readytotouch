package domain

type UserOnlineDailyCountStats struct {
	Online    int64 `json:"online"`
	UserCount int64 `json:"userCount"`
}
