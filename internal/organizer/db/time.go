package db

import "time"

func mustDate(s string) time.Time {
	parsedTime, err := time.ParseInLocation("2006-01-02", s, time.UTC)
	if err != nil {
		panic(err)
	}
	return parsedTime
}
