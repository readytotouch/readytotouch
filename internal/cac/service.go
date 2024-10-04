package cac

import (
	"time"

	"github.com/readytotouch/readytotouch/internal/linkedin"
)

type Service struct {
	linkedinClient *linkedin.Client
}

func NewService(linkedinClient *linkedin.Client) *Service {
	return &Service{linkedinClient: linkedinClient}
}

func (s *Service) Add(companyVanityName string, userID int64, now time.Time) (int64, error) {
	return 0, nil
}
