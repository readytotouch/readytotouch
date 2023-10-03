package postgres

import (
	"context"

	"github.com/readytotouch-yaaws/yaaws-go/internal/domain"
)

type OnlineStorage interface {
	BatchStore(ctx context.Context, pairs []domain.UserOnlinePair) error
}
