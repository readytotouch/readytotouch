package domain

import "time"

type TimeCountStats struct {
	Time  time.Time `json:"time"`
	Count int64     `json:"count"`
}
