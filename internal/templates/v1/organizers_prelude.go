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
		domain.IndustryAutomotive,
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
		domain.IndustryAgroTech,
		domain.IndustryInsurTech,
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

	if strings.HasSuffix(path, "/jobs") {
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

func vacancyDescription(vacancy domain.PreparedVacancy) string {
	if vacancy.ShortDescription == "" {
		return vacancy.SwitchingOpportunity
	}

	if vacancy.SwitchingOpportunity == "" {
		return vacancy.ShortDescription
	}

	return vacancy.ShortDescription + ". " + vacancy.SwitchingOpportunity
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

func isIndeedVacancyURL(s string) bool {
	return strings.Contains(s, "https://www.indeed.com/viewjob")
}

func logo(l domain.CompanyLogo) string {
	if l.V1 != "" {
		return "/assets/unstable/logos-v1/" + l.V1
	}

	if l.V0 != "" {
		return "/assets/unstable/logos-v0/" + l.V0
	}

	return "/assets/images/pages/common-images/unknown.svg"
}

func logoV1(l domain.CompanyLogo) string {
	if l.V1 != "" {
		return "/assets/unstable/logos-v1/" + l.V1
	}

	return "/assets/images/pages/common-images/unknown.svg"
}

func logoV0(l domain.CompanyLogo) string {
	if l.V0 != "" {
		return "/assets/unstable/logos-v1/" + l.V0
	}

	return "/assets/images/pages/common-images/unknown.svg"
}

var (
	companyTypeName = map[domain.CompanyType]string{
		CompanyTypeProduct: "Product",
		CompanyTypeStartup: "Startup",
	}
)
