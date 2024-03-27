package postgres

import (
	"time"

	"github.com/readytotouch/readytotouch/internal/domain"
)

const (
	hour = 3600
	day  = 86400
)

func toHourlyStats(pairs []domain.UserOnlinePair) ([]int64, []time.Time) {
	var (
		userIDs    = make([]int64, len(pairs))
		timestamps = make([]time.Time, len(pairs))
	)

	for i, pair := range pairs {
		userIDs[i] = pair.UserID
		timestamps[i] = time.Unix(truncate(pair.Timestamp, hour), 0)
	}

	return userIDs, timestamps
}

func toDailyStats(pairs []domain.UserOnlinePair) ([]int64, []time.Time) {
	var (
		userIDs    = make([]int64, len(pairs))
		timestamps = make([]time.Time, len(pairs))
	)

	for i, pair := range pairs {
		userIDs[i] = pair.UserID
		timestamps[i] = time.Unix(truncate(pair.Timestamp, day), 0)
	}

	return userIDs, timestamps
}

func truncate(source int64, d int64) int64 {
	return source - source%d
}
