package postgres

import "github.com/readytotouch/readytotouch/internal/domain"

func incUserOnlinePair(source domain.UserOnlinePair, shift int64) domain.UserOnlinePair {
	return domain.UserOnlinePair{
		UserID:    source.UserID,
		Timestamp: source.Timestamp + shift,
	}
}
