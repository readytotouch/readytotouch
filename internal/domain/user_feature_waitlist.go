package domain

type ChartStats struct {
	DailyStats []TimeCountStats `json:"daily_stats"`
	TotalCount int64            `json:"total_count"`
}

type UserFeatureWaitlistPageStats struct {
	Subscribers ChartStats `json:"subscribers"`
	Views       ChartStats `json:"views"`
}
