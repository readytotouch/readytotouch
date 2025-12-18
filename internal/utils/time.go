package utils

import "time"

func TimePointerOrNil(t time.Time) *time.Time {
	if t.IsZero() {
		return nil
	}

	return &t
}
