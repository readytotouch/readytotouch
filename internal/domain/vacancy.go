package domain

type VacancySource int

const (
	UnknownVacancySource      VacancySource = 0
	ReadyToTouchVacancySource VacancySource = 1
	LinkedInVacancySource     VacancySource = 2
	OttaVacancySource         VacancySource = 3
	IndeedVacancySource       VacancySource = 4
)
