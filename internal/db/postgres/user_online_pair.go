package postgres

import (
	"time"
)

const (
	hour = 3600
	day  = 86400
)

type UserOnlinePair struct {
	UserID int64
	Online int64
}

func toHourlyStats(pairs []UserOnlinePair) ([]int64, []time.Time) {
	var (
		userIDs    = make([]int64, len(pairs))
		timestamps = make([]time.Time, len(pairs))
	)

	for i, pair := range pairs {
		userIDs[i] = pair.UserID
		timestamps[i] = time.Unix(truncate(pair.Online, hour), 0)
	}

	return userIDs, timestamps
}

func toDailyStats(pairs []UserOnlinePair) ([]int64, []time.Time) {
	var (
		userIDs    = make([]int64, len(pairs))
		timestamps = make([]time.Time, len(pairs))
	)

	for i, pair := range pairs {
		userIDs[i] = pair.UserID
		timestamps[i] = time.Unix(truncate(pair.Online, day), 0)
	}

	return userIDs, timestamps
}

func toDailyMin(pairs []UserOnlinePair) time.Time {
	online := pairs[0].Online
	for _, pair := range pairs {
		online = min(online, pair.Online)
	}

	return time.Unix(truncate(online, day), 0)
}

func truncate(source int64, d int64) int64 {
	return source - source%d
}
