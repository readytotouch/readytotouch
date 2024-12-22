package main

import (
	"encoding/json"
	"fmt"
	"go/format"
	"os"

	"github.com/readytotouch/readytotouch/internal/domain"
	"github.com/readytotouch/readytotouch/internal/organizer/db"
	"github.com/readytotouch/readytotouch/internal/protos/organizers"
	"github.com/readytotouch/readytotouch/internal/templates/dev"
)

func main() {
	var (
		companies = db.Companies()
	)

	generateCompanies(companies)
	generateVacancies(companies)
}

func generateCompanies(companies []domain.CompanyProfile) {
	var (
		maxID = int64(0)
		pairs = make([]*dev.CompanyCodePair, 0, len(companies))
		names = make([]string, 0, len(companies))
	)

	// Assert that all company aliases are present in the map
	for alias := range organizers.CompanyStartupMap {
		if _, ok := organizers.CompanyAliasMap[alias]; !ok {
			panic(fmt.Sprintf("Company alias not found: %s", alias))
		}
	}

	for _, company := range companies {
		id := organizers.CompanyAliasMap[company.LinkedInProfile.Alias]

		// If we added the ignored company before, then keep it
		if id == 0 && company.Ignore {
			continue
		}

		pair := &dev.CompanyCodePair{
			ID:    id,
			Name:  company.LinkedInProfile.Name,
			Alias: company.LinkedInProfile.Alias,
		}

		pairs = append(pairs, pair)
		names = append(names, company.LinkedInProfile.Name)
		maxID = max(maxID, id)
	}

	for _, pair := range pairs {
		if pair.ID == 0 {
			maxID++
			pair.ID = maxID
		}
	}

	output, err := format.Source([]byte(dev.CompanyCode(pairs)))
	if err != nil {
		panic(err)
	}

	err = os.WriteFile("./internal/protos/organizers/company_code.go", output, 0644)
	if err != nil {
		panic(err)
	}

	fmt.Println("Company code generated successfully")
	fmt.Printf("Max ID: %d\n", maxID)

	// json.NewEncoder(os.Stdout).Encode(pairs)
	json.NewEncoder(os.Stdout).Encode(names) // hide
}

func generateVacancies(companies []domain.CompanyProfile) {
	const (
		count = 1024
	)

	var (
		maxID        = int64(0)
		pairs        = make([]*dev.VacancyCodePair, 0, count)
		urlExistsMap = make(map[string]bool, count)
	)

	for _, company := range companies {
		for _, language := range company.Languages {
			for _, vacancy := range language.Vacancies {
				if urlExistsMap[vacancy.URL] {
					continue
				}

				pair := &dev.VacancyCodePair{
					ID:           organizers.VacancyUrlMap[vacancy.URL],
					URL:          vacancy.URL,
					CompanyAlias: company.LinkedInProfile.Alias,
				}

				pairs = append(pairs, pair)
				maxID = max(maxID, pair.ID)

				urlExistsMap[vacancy.URL] = true
			}
		}
	}

	for _, pair := range pairs {
		if pair.ID == 0 {
			maxID++
			pair.ID = maxID
		}
	}

	output, err := format.Source([]byte(dev.VacancyCode(pairs)))
	if err != nil {
		panic(err)
	}

	err = os.WriteFile("./internal/protos/organizers/vacancy_code.go", output, 0644)
	if err != nil {
		panic(err)
	}

	fmt.Println("Vacancy code generated successfully")
	fmt.Printf("Max ID: %d\n", maxID)
}
