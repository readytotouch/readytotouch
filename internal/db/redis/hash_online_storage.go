package redis

import (
	"context"
	"strconv"

	"github.com/readytotouch/readytotouch/internal/domain"

	"github.com/redis/go-redis/v9"
)

type HashOnlineStorage struct {
	client *redis.Client
}

func NewHashOnlineStorage(client *redis.Client) *HashOnlineStorage {
	return &HashOnlineStorage{client: client}
}

func (s *HashOnlineStorage) Store(ctx context.Context, pair domain.UserOnlinePair) error {
	return s.client.HSet(
		ctx,
		"h:online:main",
		strconv.FormatInt(pair.UserID, 10),
		pair.Timestamp,
	).Err()
}

func (s *HashOnlineStorage) Count(ctx context.Context) (int64, error) {
	return s.client.HLen(ctx, "h:online:main").Result()
}

func (s *HashOnlineStorage) GetAndClear(ctx context.Context) ([]domain.UserOnlinePair, error) {
	var (
		oldKey = "h:online:main"
		newKey = "h:online:tmp"
	)

	err := s.client.Rename(ctx, oldKey, newKey).Err()
	if err != nil {
		return nil, err
	}

	userOnlineMap, err := s.client.HGetAll(ctx, newKey).Result()
	if err != nil {
		return nil, err
	}

	result := make([]domain.UserOnlinePair, 0, len(userOnlineMap))
	for stringUserID, stringTimestamp := range userOnlineMap {
		userID, err := strconv.ParseInt(stringUserID, 10, 64)
		if err != nil {
			// unreachable, ignore for article

			// logging or use https://github.com/hashicorp/go-multierror

			// just in case
			continue
		}

		timestamp, err := strconv.ParseInt(stringTimestamp, 10, 64)
		if err != nil {
			// unreachable, ignore for article

			// logging or use https://github.com/hashicorp/go-multierror

			// just in case
			continue
		}

		result = append(result, domain.UserOnlinePair{
			UserID:    userID,
			Timestamp: timestamp,
		})
	}

	return result, nil
}
