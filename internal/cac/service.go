package cac

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/readytotouch/readytotouch/internal/db/postgres"
	"github.com/readytotouch/readytotouch/internal/linkedin"
)

type Service struct {
	mu                              sync.Mutex
	userToLinkedInCompanyRepository *postgres.UserToLinkedInCompanyRepository
	linkedinClient                  *linkedin.Client
}

func NewService(userToLinkedInCompanyRepository *postgres.UserToLinkedInCompanyRepository, linkedinClient *linkedin.Client) *Service {
	return &Service{userToLinkedInCompanyRepository: userToLinkedInCompanyRepository, linkedinClient: linkedinClient}
}

func (s *Service) Load(ctx context.Context, companyVanityName string, userID int64, now time.Time) (int64, error) {
	// Throttling to avoid LinkedIn API rate limits
	s.mu.Lock()
	defer s.mu.Unlock()

	id, found, err := s.userToLinkedInCompanyRepository.GetByVanityName(ctx, companyVanityName)
	if err != nil {
		return 0, err
	}

	if found {
		return id, nil
	}

	exists, err := s.userToLinkedInCompanyRepository.ExistsVanityName(ctx, companyVanityName)
	if err != nil {
		return 0, err
	}

	if exists {
		return 0, fmt.Errorf("cannot load company %s, it's already in the database", companyVanityName)
	}

	// Minimize the number of requests to the LinkedIn API per user
	{
		previousRequestHistoryCount, err := s.userToLinkedInCompanyRepository.GetRequestHistoryCount(ctx, userID)
		if err != nil {
			return 0, err
		}

		if previousRequestHistoryCount >= 100 {
			return 0, fmt.Errorf("cannot load company %s, the limit of 100 requests per user has been reached", companyVanityName)
		}
	}

	companies, payload, err := s.linkedinClient.CompaniesSearch(companyVanityName)
	if err != nil {
		return 0, err
	}

	// Success
	if len(companies) == 1 {
		var company = companies[0]

		err = s.userToLinkedInCompanyRepository.CreateCompany(
			ctx,
			company.ID,
			companyVanityName,
			company.Name,
			payload,
			userID,
			now,
		)
		if err != nil {
			return 0, err
		}

		return company.ID, nil
	}

	// Shame on you, username
	err = s.userToLinkedInCompanyRepository.CreateCompany(
		ctx,
		0,
		companyVanityName,
		"",
		payload,
		userID,
		now,
	)
	if err != nil {
		return 0, err
	}

	return 0, fmt.Errorf("cannot load company %s, it's already in the database", companyVanityName)
}
