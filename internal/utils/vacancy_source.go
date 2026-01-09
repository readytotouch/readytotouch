package utils

import (
	"strings"

	"github.com/readytotouch/readytotouch/internal/domain"
)

func IsLinkedInVacancyURL(s string) bool {
	return strings.Contains(s, "https://www.linkedin.com/jobs/view/")
}

func IsOttaVacancyURL(s string) bool {
	return strings.Contains(s, "https://app.welcometothejungle.com/jobs/") ||
		strings.Contains(s, "https://app.otta.com/jobs/")
}

func IsIndeedVacancyURL(s string) bool {
	return strings.Contains(s, "https://www.indeed.com/viewjob")
}

func IsXingVacancyURL(s string) bool {
	return strings.Contains(s, "https://www.xing.com/jobs/")
}

func DetectVacancySource(s string) domain.VacancySource {
	if IsLinkedInVacancyURL(s) {
		return domain.LinkedInVacancySource
	}
	if IsOttaVacancyURL(s) {
		return domain.OttaVacancySource
	}
	if IsIndeedVacancyURL(s) {
		return domain.IndeedVacancySource
	}

	return domain.UnknownVacancySource
}
