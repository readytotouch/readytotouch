package db

import (
	"slices"

	"github.com/readytotouch/readytotouch/internal/domain"
)

var (
	companies = Companies()
)

func Companies() []domain.CompanyProfile {
	var (
		part1  = companiesPart1()
		part2  = companiesPart2()
		part3  = companiesPart3()
		part4  = companiesPart4()
		part5  = companiesPart5()
		part6  = companiesPart6()
		part7  = companiesPart7()
		part8  = companiesPart8()
		part9  = companiesPart9()
		part10 = companiesPart10()
		part11 = companiesPart11()
		part12 = companiesPart12()
		result = make([]domain.CompanyProfile, 0, len(part1)+
			len(part2)+
			len(part3)+
			len(part4)+
			len(part5)+
			len(part6)+
			len(part7)+
			len(part8)+
			len(part9)+
			len(part10)+
			len(part11)+
			len(part12))
	)

	result = append(result, part1...)
	result = append(result, part2...)
	result = append(result, part3...)
	result = append(result, part4...)
	result = append(result, part5...)
	result = append(result, part6...)
	result = append(result, part7...)
	result = append(result, part8...)
	result = append(result, part9...)
	result = append(result, part10...)
	result = append(result, part11...)
	result = append(result, part12...)

	return result
}

func CloneCompanies() []domain.CompanyProfile {
	return slices.Clone(companies)
}
