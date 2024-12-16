package v1

import "time"

const (
	overOneMonth     = "Over one month"
	overTwoMonths    = "Over two months"
	overThreeMonths  = "Over three months"
	overFourMonths   = "Over four months"
	overFiveMonths   = "Over five months"
	overSixMonths    = "Over six months"
	overSevenMonths  = "Over seven months"
	overEightMonths  = "Over eight months"
	overNineMonths   = "Over nine months"
	overTenMonths    = "Over ten months"
	overElevenMonths = "Over eleven months"
	overAYear        = "Over a year"
)

type vacancyPeriod struct {
	name   string
	before time.Time
	used   bool
}

type vacancyPeriods []*vacancyPeriod

func newVacancyPeriods() vacancyPeriods {
	today := time.Now().UTC().Truncate(24 * time.Hour)

	return vacancyPeriods{
		{name: overOneMonth, before: today.AddDate(0, -1, 0), used: false},
		{name: overTwoMonths, before: today.AddDate(0, -2, 0), used: false},
		{name: overThreeMonths, before: today.AddDate(0, -3, 0), used: false},
		{name: overFourMonths, before: today.AddDate(0, -4, 0), used: false},
		{name: overFiveMonths, before: today.AddDate(0, -5, 0), used: false},
		{name: overSixMonths, before: today.AddDate(0, -6, 0), used: false},
		{name: overSevenMonths, before: today.AddDate(0, -7, 0), used: false},
		{name: overEightMonths, before: today.AddDate(0, -8, 0), used: false},
		{name: overNineMonths, before: today.AddDate(0, -9, 0), used: false},
		{name: overTenMonths, before: today.AddDate(0, -10, 0), used: false},
		{name: overElevenMonths, before: today.AddDate(0, -11, 0), used: false},
		{name: overAYear, before: today.AddDate(0, -12, 0), used: false},
	}
}

func (v vacancyPeriods) over(vacancyDate time.Time) []string {
	var names []string

	for _, p := range v {
		if !p.used && vacancyDate.Before(p.before) {
			names = append(names, p.name)
			p.used = true
		}
	}

	return names
}
