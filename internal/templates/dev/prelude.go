package dev

import "github.com/readytotouch/readytotouch/internal/organizer/domain"

type Company = domain.Company

type CompanyCodePair struct {
	ID    int64
	Alias string
}
