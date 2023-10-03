package domain

import "time"

type TimeCountStats struct {
	Time  time.Time `json:"time"`
	Count int64     `json:"count"`
}

type UserOnlinePair struct {
	UserID    int64
	Timestamp int64
}

type UserOnlinePairs []UserOnlinePair

func (ps UserOnlinePairs) Len() int {
	return len(ps)
}

func (ps UserOnlinePairs) Less(i, j int) bool {
	return ps.less(ps[i], ps[j])
}

func (ps UserOnlinePairs) Swap(i, j int) {
	ps[i], ps[j] = ps[j], ps[i]
}

func (ps UserOnlinePairs) less(l, r UserOnlinePair) bool {
	if l.UserID == r.UserID {
		return l.Timestamp < r.Timestamp
	}

	return l.UserID < r.UserID
}
