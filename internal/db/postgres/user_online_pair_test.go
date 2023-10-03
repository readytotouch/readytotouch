package postgres

import "github.com/readytotouch-yaaws/yaaws-go/internal/domain"

func incUserOnlinePair(source domain.UserOnlinePair, shift int64) domain.UserOnlinePair {
	return domain.UserOnlinePair{
		UserID:    source.UserID,
		Timestamp: source.Timestamp + shift,
	}
}
