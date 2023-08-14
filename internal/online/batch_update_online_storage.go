package main

import (
	"context"

	postgresql "github.com/readytotouch-yaaws/yaaws-go/internal/storage/postgres"
	"github.com/readytotouch-yaaws/yaaws-go/internal/storage/postgres/dbs"
)

type BatchUpdateOnlineStorage struct {
	repository *postgresql.Repository
}

func NewBatchUpdateOnlineStorage(repository *postgresql.Repository) *BatchUpdateOnlineStorage {
	return &BatchUpdateOnlineStorage{repository: repository}
}

func (s *BatchUpdateOnlineStorage) BatchStore(ctx context.Context, pairs []UserOnlinePair) error {
	userIDs, timestamps := userOnlinePairsToPgxSlices(pairs)

	return s.repository.Queries().UserOnlineBatchUpdate(ctx, dbs.UserOnlineBatchUpdateParams{
		UserIds: userIDs,
		Onlines: timestamps,
	})
}
