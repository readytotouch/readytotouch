package db

import "github.com/readytotouch/readytotouch/internal/domain"

func Companies() []domain.CompanyProfile {
	var (
		part1  = companiesPart1()
		part2  = companiesPart2()
		part3  = companiesPart3()
		part4  = companiesPart4()
		result = make([]domain.CompanyProfile, 0, len(part1)+len(part2)+len(part3)+len(part4))
	)

	result = append(result, part1...)
	result = append(result, part2...)
	result = append(result, part3...)
	result = append(result, part4...)

	return result
}
