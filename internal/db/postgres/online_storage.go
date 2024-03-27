package postgres

import (
	"context"

	"github.com/readytotouch/readytotouch/internal/domain"
)

type OnlineStorage interface {
	BatchStore(ctx context.Context, pairs []domain.UserOnlinePair) error
}
