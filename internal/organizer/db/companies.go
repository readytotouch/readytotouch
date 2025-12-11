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
		shard01            = companies01Shard()
		shard02            = companies02Shard()
		shard03            = companies03Shard()
		shard04            = companies04Shard()
		shard05            = companies05Shard()
		shard06            = companies06Shard()
		shard07            = companies07Shard()
		shard08            = companies08Shard()
		shard09            = companies09Shard()
		shard10            = companies10Shard()
		shard11            = companies11Shard()
		shard12            = companies12Shard()
		shard13            = companies13Shard()
		shard14            = companies14Shard()
		shard15            = companies15Shard()
		shard16            = companies16Shard()
		shard17            = companies17Shard()
		shard18            = companies18Shard()
		shard19            = companies19Shard()
		shard20            = companies20Shard()
		shard21            = companies21Shard()
		shard22            = companies22Shard()
		shard23            = companies23Shard()
		outsourceCompanies = companiesOutsourceCompanies()
		recruitingAgencies = companiesRecruitingAgencies()
		result             = make([]domain.CompanyProfile, 0, len(shard01)+
			len(shard02)+
			len(shard03)+
			len(shard04)+
			len(shard05)+
			len(shard06)+
			len(shard07)+
			len(shard08)+
			len(shard09)+
			len(shard10)+
			len(shard11)+
			len(shard12)+
			len(shard13)+
			len(shard14)+
			len(shard15)+
			len(shard16)+
			len(shard17)+
			len(shard18)+
			len(shard19)+
			len(shard20)+
			len(shard21)+
			len(shard22)+
			len(shard23)+
			len(outsourceCompanies)+
			len(recruitingAgencies),
		)
	)

	result = append(result, shard01...)
	result = append(result, shard02...)
	result = append(result, shard03...)
	result = append(result, shard04...)
	result = append(result, shard05...)
	result = append(result, shard06...)
	result = append(result, shard07...)
	result = append(result, shard08...)
	result = append(result, shard09...)
	result = append(result, shard10...)
	result = append(result, shard11...)
	result = append(result, shard12...)
	result = append(result, shard13...)
	result = append(result, shard14...)
	result = append(result, shard15...)
	result = append(result, shard16...)
	result = append(result, shard17...)
	result = append(result, shard18...)
	result = append(result, shard19...)
	result = append(result, shard20...)
	result = append(result, shard21...)
	result = append(result, shard22...)
	result = append(result, shard23...)
	result = append(result, outsourceCompanies...)
	result = append(result, recruitingAgencies...)

	return result
}

func CloneCompanies() []domain.CompanyProfile {
	return slices.Clone(companies)
}
