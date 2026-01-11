package v1

import "github.com/readytotouch/readytotouch/internal/generated/organizers"

type CriteriaLocation struct {
	Code string
	Name string
}

func criteriaLocations() []CriteriaLocation {
	return []CriteriaLocation{}
}

func showLocation(s string) bool {
	if s == "" {
		return false
	}

	_, exists := organizers.LocationCodeMap[s]

	return exists
}
