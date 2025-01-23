package v1

import (
	"fmt"
	"strings"
	"time"

	"github.com/readytotouch/readytotouch/internal/domain"
)

const (
	Go = domain.Go
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
	companiesActive string
	vacanciesActive string
}

func toFeatureNavigation(path string) featureNavigation {
	if strings.HasSuffix(path, "/companies") {
		return featureNavigation{
			companiesActive: "active",
			vacanciesActive: "",
		}
	}

	if strings.HasSuffix(path, "/vacancies") {
		return featureNavigation{
			companiesActive: "",
			vacanciesActive: "active",
		}
	}

	return featureNavigation{
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

func formatVacancyDate(s time.Time) string {
	return s.Format("02 January 2006")
}

func formatVacancyDiffDate(t time.Time) string {
	today := time.Now().UTC().Truncate(24 * time.Hour)

	if t.After(today) {
		return "Today"
	}

	if t.After(today.Add(-24 * time.Hour)) {
		return "Yesterday"
	}

	days := int(today.Sub(t).Hours() / 24)
	if days == 1 {
		return "1 day ago"
	}
	return fmt.Sprintf("%d days ago", days)
}

func isLinkedInVacancyURL(s string) bool {
	return strings.Contains(s, "https://www.linkedin.com/jobs/view/")
}

func isOttaVacancyURL(s string) bool {
	return strings.Contains(s, "https://app.welcometothejungle.com/jobs/") ||
		strings.Contains(s, "https://app.otta.com/jobs/")
}

func logo(s string) string {
	if s == "" {
		return "/assets/images/pages/common-images/unknown.svg"
	}

	return "/assets/unstable/logos/" + s
}

var (
	companyTypeName = map[domain.CompanyType]string{
		CompanyTypeProduct: "Product",
		CompanyTypeStartup: "Startup",
	}
)
