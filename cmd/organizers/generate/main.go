package main

import (
	"encoding/json"
	"fmt"
	"go/format"
	"os"
	"sort"

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
	generateLogosSearch(companies)
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
		if company.Name == "" {
			panic("Company name is empty")
		}
		if company.LinkedInProfile.Alias == "" {
			panic(fmt.Sprintf("Company LinkedIn alias is empty for company: %s", company.Name))
		}
		if company.LinkedInProfile.Name == "" {
			panic(fmt.Sprintf("Company LinkedIn name is empty for company: %s", company.Name))
		}

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
		maxID                       = int64(0)
		pairs                       = make([]*dev.VacancyCodePair, 0, count)
		urlCompanyLanguageExistsMap = make(map[string]map[string]map[int]bool, count)
	)

	for _, company := range companies {
		for language, languageProfile := range company.Languages {
			for _, vacancy := range languageProfile.Vacancies {
				if vacancy.Title == "" {
					panic(fmt.Sprintf("Vacancy title is empty for company: %s", company.Name))
				}
				if vacancy.URL == "" {
					panic(fmt.Sprintf("Vacancy URL is empty for company: %s", company.Name))
				}

				companiesCount := len(urlCompanyLanguageExistsMap[vacancy.URL])

				// Assert
				{
					if companiesCount > 1 {
						panic(fmt.Sprintf("Vacancy URL: %s is used in multiple companies", vacancy.URL))
					}

					if urlCompanyLanguageExistsMap[vacancy.URL][company.LinkedInProfile.Alias][language] {
						panic(fmt.Sprintf("Vacancy URL: %s is duplicated", vacancy.URL))
					}

					if urlCompanyLanguageExistsMap[vacancy.URL] == nil {
						urlCompanyLanguageExistsMap[vacancy.URL] = make(map[string]map[int]bool, 1)
					}

					if urlCompanyLanguageExistsMap[vacancy.URL][company.LinkedInProfile.Alias] == nil {
						urlCompanyLanguageExistsMap[vacancy.URL][company.LinkedInProfile.Alias] = make(map[int]bool, 1)
					}

					urlCompanyLanguageExistsMap[vacancy.URL][company.LinkedInProfile.Alias][language] = true
				}

				if companiesCount > 0 {
					continue
				}

				vacancyID := organizers.VacancyUrlMap[vacancy.URL]

				// If we added the ignored company before, then keep it
				if vacancyID == 0 && company.Ignore {
					continue
				}

				pair := &dev.VacancyCodePair{
					ID:           vacancyID,
					URL:          vacancy.URL,
					CompanyAlias: company.LinkedInProfile.Alias,
				}

				pairs = append(pairs, pair)
				maxID = max(maxID, pair.ID)
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

func generateLogosSearch(companies []domain.CompanyProfile) {
	var (
		pairs = make([]*dev.CompanyCodePair, 0, len(companies))
	)

	for _, company := range companies {
		id := organizers.CompanyAliasMap[company.LinkedInProfile.Alias]

		if id == 0 || !company.LinkedInProfile.Verified || company.Ignore {
			continue
		}

		pairs = append(pairs, &dev.CompanyCodePair{
			ID:    id,
			Name:  company.LinkedInProfile.Name,
			Alias: company.LinkedInProfile.Alias,
		})
	}

	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i].Name < pairs[j].Name
	})

	err := os.WriteFile("./private/logos_search.md", []byte(dev.LogosSearch(pairs)), 0644)
	if err != nil {
		panic(err)
	}
}
