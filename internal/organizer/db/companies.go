package db

import "github.com/readytotouch/readytotouch/internal/domain"

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
		result = make([]domain.CompanyProfile, 0, len(part1)+
			len(part2)+
			len(part3)+
			len(part4)+
			len(part5)+
			len(part6)+
			len(part7)+
			len(part8)+
			len(part9))
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

	return result
}
