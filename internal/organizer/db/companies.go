package db

import (
	"fmt"
	"slices"

	"github.com/readytotouch/readytotouch/internal/domain"
)

var (
	companies = Companies()
)

func Companies() []domain.CompanyProfile {
	var (
		part1              = companies01Shard()
		part2              = companies02Shard()
		part3              = companies03Shard()
		part4              = companies04Shard()
		part5              = companies05Shard()
		part6              = companies06Shard()
		part7              = companies07Shard()
		part8              = companies08Shard()
		part9              = companies09Shard()
		part10             = companies10Shard()
		part11             = companies11Shard()
		part12             = companies12Shard()
		putsourceCompanies = companiesOutsourceCompanies()
		recruitingAgencies = companiesRecruitingAgencies()
		result             = make([]domain.CompanyProfile, 0, len(part1)+
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
			len(part12)+
			len(putsourceCompanies)+
			len(recruitingAgencies) +
			+0, // to speed up adding a new part
		)
	)

	fmt.Println(
		"part1", len(part1), "\n",
		"part2", len(part2), "\n",
		"part3", len(part3), "\n",
		"part4", len(part4), "\n",
		"part5", len(part5), "\n",
		"part6", len(part6), "\n",
		"part7", len(part7), "\n",
		"part8", len(part8), "\n",
		"part9", len(part9), "\n",
		"part10", len(part10), "\n",
		"part11", len(part11), "\n",
		"part12", len(part12), "\n",
		len(part1)+
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
			len(part12),
		(len(part1)+
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
			len(part12))/64,
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
	result = append(result, putsourceCompanies...)
	result = append(result, recruitingAgencies...)

	return result
}

func CloneCompanies() []domain.CompanyProfile {
	return slices.Clone(companies)
}
