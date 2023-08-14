package main

import (
	"time"

	"github.com/jackc/pgx/v5/pgtype"
)

type UserOnlinePair struct {
	UserID int64
	Online int64
}

func userOnlinePairsToPgxSlices(pairs []UserOnlinePair) ([]int64, []pgtype.Timestamp) {
	var (
		userIDs    = make([]int64, len(pairs))
		timestamps = make([]pgtype.Timestamp, len(pairs))
	)

	for i, pair := range pairs {
		userIDs[i] = pair.UserID
		timestamps[i] = pgtype.Timestamp{
			Time:  time.Unix(pair.Online, 0).UTC(),
			Valid: true,
		}
	}

	return userIDs, timestamps
}
