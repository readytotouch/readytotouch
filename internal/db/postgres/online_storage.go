package postgres

import (
	"context"
)

type OnlineStorage interface {
	BatchStore(ctx context.Context, pairs []UserOnlinePair) error
}
