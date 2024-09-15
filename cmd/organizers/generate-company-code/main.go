package main

import (
	"net/url"
	"os"
	"regexp"
	"strings"

	"github.com/readytotouch/readytotouch/internal/organizer/db"
	"github.com/readytotouch/readytotouch/internal/organizer/domain"
	"github.com/readytotouch/readytotouch/internal/protos/organizers"
	"github.com/readytotouch/readytotouch/internal/templates/dev"
)

func main() {
	var (
		maxCompanyCode  int32 = 0
		letterOnlyRegex       = regexp.MustCompile(`^[a-zA-Z]+$`)
		companies             = make([]domain.Company, 0, len(db.Companies()))
	)

	for code := range organizers.CompanyCode_name {
		maxCompanyCode = max(maxCompanyCode, code)
	}

	for _, company := range db.Companies() {
		if _, exists := organizers.CompanyCode_value[company.LinkedInProfile.Alias]; exists {
			continue
		}

		if !letterOnlyRegex.MatchString(company.LinkedInProfile.Alias) {
			continue
		}

		site, err := site(company.URL)
		if err != nil {
			panic(err)
		}

		if company.LinkedInProfile.Alias != site {
			continue
		}

		companies = append(companies, company)
	}

	err := os.WriteFile(
		"./protos/organizers/company_code.proto",
		[]byte(dev.CompanyCode(organizers.CompanyCode_name, maxCompanyCode, companies)),
		0644,
	)
	if err != nil {
		panic(err)
	}
}

func site(s string) (string, error) {
	parsedURL, err := url.Parse(s)
	if err != nil {
		return "", err
	}

	result, _, _ := strings.Cut(strings.TrimPrefix(parsedURL.Host, "www."), ".")

	return result, nil
}
