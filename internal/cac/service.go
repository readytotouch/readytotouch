package cac

import (
	"time"

	"github.com/readytotouch/readytotouch/internal/db/postgres"
	"github.com/readytotouch/readytotouch/internal/linkedin"
)

type Service struct {
	userToLinkedInCompanyRepository *postgres.UserToLinkedInCompanyRepository
	linkedinClient                  *linkedin.Client
}

func NewService(userToLinkedInCompanyRepository *postgres.UserToLinkedInCompanyRepository, linkedinClient *linkedin.Client) *Service {
	return &Service{userToLinkedInCompanyRepository: userToLinkedInCompanyRepository, linkedinClient: linkedinClient}
}

func (s *Service) Add(companyVanityName string, userID int64, now time.Time) (int64, error) {
	return 0, nil
}
