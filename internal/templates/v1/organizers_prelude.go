package v1

import (
	"strings"

	"github.com/readytotouch/readytotouch/internal/domain"
)

type OrganizerFeature = domain.OrganizerFeature

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
