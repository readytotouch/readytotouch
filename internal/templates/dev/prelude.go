package dev

import "github.com/readytotouch/readytotouch/internal/domain"

type Company = domain.Company

type CompanyCodePair struct {
	ID    int64
	Name  string
	Alias string
}
