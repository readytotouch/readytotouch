package domain

import "time"

type GithubRepositoryStar struct {
	Owner           string
	Repo            string
	StargazersCount int
	UpdatedAt       time.Time
}
