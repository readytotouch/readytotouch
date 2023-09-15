package main

import (
	"time"

	"github.com/jackc/pgx/v5/pgtype"
)

const (
	hour = 3600
	day  = 86400
)

type UserOnlinePair struct {
	UserID int64
	Online int64
}

func toHourlyStats(pairs []UserOnlinePair) ([]int64, []pgtype.Timestamp) {
	var (
		userIDs    = make([]int64, len(pairs))
		timestamps = make([]pgtype.Timestamp, len(pairs))
	)

	for i, pair := range pairs {
		userIDs[i] = pair.UserID
		timestamps[i] = pgtype.Timestamp{
			Time:  time.Unix(truncate(pair.Online, hour), 0),
			Valid: true,
		}
	}

	return userIDs, timestamps
}

func toDailyStats(pairs []UserOnlinePair) ([]int64, []pgtype.Date) {
	var (
		userIDs    = make([]int64, len(pairs))
		timestamps = make([]pgtype.Date, len(pairs))
	)

	for i, pair := range pairs {
		userIDs[i] = pair.UserID
		timestamps[i] = pgtype.Date{
			Time:  time.Unix(truncate(pair.Online, day), 0),
			Valid: true,
		}
	}

	return userIDs, timestamps
}

func toDailyMin(pairs []UserOnlinePair) pgtype.Date {
	online := pairs[0].Online
	for _, pair := range pairs {
		online = min(online, pair.Online)
	}

	return pgtype.Date{
		Time:  time.Unix(truncate(online, day), 0),
		Valid: true,
	}
}

func truncate(source int64, d int64) int64 {
	return source - source%d
}
