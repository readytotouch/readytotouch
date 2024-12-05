package v1

import (
	"strings"

	"github.com/readytotouch/readytotouch/internal/domain"
)

const (
	Go = 0
)

const (
	CompanyTypeProduct = domain.CompanyTypeProduct
	CompanyTypeStartup = domain.CompanyTypeStartup
)

var (
	industries = []Industry{
		domain.IndustryCyberSecurity,
		domain.IndustryHealthTech,
		domain.IndustryFinTech,
		domain.IndustryAdTech,
		domain.IndustryMarTech,
		domain.IndustryDevOps,
		domain.IndustryBigData,
		domain.IndustrySocialMedia,
		domain.IndustryEntertainment,
		domain.IndustryPropTech,
	}
)

type (
	Organizer        = domain.Organizer
	OrganizerFeature = domain.OrganizerFeature
	CompanyStats     struct {
		TotalViews         int64
		LastMonthViews     int64
		TotalFavorites     int64
		LastMonthFavorites int64
	}
	Industry = domain.Industry
)

type featureNavigation struct {
	vacanciesURL    string
	companiesURL    string
	companiesActive string
	vacanciesActive string
}

func toFeatureNavigation(path string) featureNavigation {
	if strings.HasSuffix(path, "/companies") {
		return featureNavigation{
			vacanciesURL:    strings.Replace(path, "/companies", "/vacancies", 1),
			companiesURL:    path,
			companiesActive: "active",
			vacanciesActive: "",
		}
	}

	if strings.HasSuffix(path, "/vacancies") {
		return featureNavigation{
			vacanciesURL:    path,
			companiesURL:    strings.Replace(path, "/vacancies", "/companies", 1),
			companiesActive: "",
			vacanciesActive: "active",
		}
	}

	return featureNavigation{
		vacanciesURL:    "javascript:void(0)",
		companiesURL:    "javascript:void(0)",
		companiesActive: "",
		vacanciesActive: "",
	}
}

func aliases(source []Industry) string {
	result := make([]string, len(source))
	for i, v := range source {
		result[i] = v.Alias
	}
	return strings.Join(result, ",")
}

func industryNames(source []Industry) string {
	result := make([]string, len(source))
	for i, v := range source {
		result[i] = v.Name
	}
	return strings.Join(result, ", ")
}

func formatLinkedInFollowers(s string) string {
	if s == "" {
		return "???"
	}

	return s
}

func formatLinkedInEmployees(s string) string {
	if s == "" {
		return "???"
	}

	return s
}

func formatLinkedInAssociatedMembers(s string) string {
	if s == "" {
		return "???"
	}

	return s
}

func fetchGitHubRepositoriesCount(company domain.CompanyProfile, language domain.Language) int {
	return company.Languages[language].GitHubRepositoriesCount
}

func formatBlindEmployees(s string) string {
	if s == "" {
		return "???"
	}

	return s
}

func formatBlindSalary(s string) string {
	if s == "" {
		return "$??K ~ $???K a year"
	}

	if strings.HasSuffix(s, "a year") {
		return s
	}

	return s + " " + "a year"
}

func formatBlindReviews(s string) string {
	if s == "" {
		return "???"
	}

	return s
}

func formatLevelsFyiEmployees(s string) string {
	if s == "" {
		return "???"
	}

	return s
}

func formatGlassdoorJobs(s string) string {
	if s == "" {
		return "???"
	}

	return s
}

func formatGlassdoorReviews(s string) string {
	if s == "" {
		return "???"
	}

	return s
}

func formatGlassdoorSalaries(s string) string {
	if s == "" {
		return "???"
	}

	return s
}

func formatGlassdoorReviewsRate(s string) string {
	if s == "" {
		return "?.?"
	}

	return s
}

var (
	companyTypeName = map[domain.CompanyType]string{
		CompanyTypeProduct: "Product",
		CompanyTypeStartup: "Startup",
	}
)
