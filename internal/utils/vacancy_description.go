package utils

func VacancyDescription(shortDescription string, switchingOpportunity string) string {
	if shortDescription == "" {
		return switchingOpportunity
	}

	if switchingOpportunity == "" {
		return shortDescription
	}

	return shortDescription + ". " + switchingOpportunity
}
