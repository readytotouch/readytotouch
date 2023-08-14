package main

import (
	"context"
)

type OnlineStorage interface {
	BatchStore(ctx context.Context, pairs []UserOnlinePair) error
}
