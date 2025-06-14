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
		part01             = companies01Shard()
		part02             = companies02Shard()
		part03             = companies03Shard()
		part04             = companies04Shard()
		part05             = companies05Shard()
		part06             = companies06Shard()
		part07             = companies07Shard()
		part08             = companies08Shard()
		part09             = companies09Shard()
		part10             = companies10Shard()
		part11             = companies11Shard()
		part12             = companies12Shard()
		part13             = companies13Shard()
		part14             = companies14Shard()
		part15             = companies15Shard()
		part16             = companies16Shard()
		part17             = companies17Shard()
		part18             = companies18Shard()
		part19             = companies19Shard()
		part20             = companies20Shard()
		outsourceCompanies = companiesOutsourceCompanies()
		recruitingAgencies = companiesRecruitingAgencies()
		result             = make([]domain.CompanyProfile, 0, len(part01)+
			len(part02)+
			len(part03)+
			len(part04)+
			len(part05)+
			len(part06)+
			len(part07)+
			len(part08)+
			len(part09)+
			len(part10)+
			len(part11)+
			len(part12)+
			len(part13)+
			len(part14)+
			len(part15)+
			len(part16)+
			len(part17)+
			len(part18)+
			len(part19)+
			len(part20)+
			len(outsourceCompanies)+
			len(recruitingAgencies),
		)
	)

	const debug = false

	if debug {
		fmt.Println(
			"part01", len(part01), "\n",
			"part02", len(part02), "\n",
			"part03", len(part03), "\n",
			"part04", len(part04), "\n",
			"part05", len(part05), "\n",
			"part06", len(part06), "\n",
			"part07", len(part07), "\n",
			"part08", len(part08), "\n",
			"part09", len(part09), "\n",
			"part10", len(part10), "\n",
			"part11", len(part11), "\n",
			"part12", len(part12), "\n",
			"part13", len(part13), "\n",
			"part14", len(part14), "\n",
			"part15", len(part15), "\n",
			"part16", len(part16), "\n",
			"part17", len(part17), "\n",
			"part18", len(part18), "\n",
			"part19", len(part19), "\n",
			"part20", len(part20), "\n",
			len(part01)+
				len(part02)+
				len(part03)+
				len(part04)+
				len(part05)+
				len(part06)+
				len(part07)+
				len(part08)+
				len(part09)+
				len(part10)+
				len(part11)+
				len(part12)+
				len(part13)+
				len(part14)+
				len(part15)+
				len(part16)+
				len(part17)+
				len(part18)+
				len(part19)+
				len(part20),
			(len(part01)+
				len(part02)+
				len(part03)+
				len(part04)+
				len(part05)+
				len(part06)+
				len(part07)+
				len(part08)+
				len(part09)+
				len(part10)+
				len(part11)+
				len(part12)+
				len(part13)+
				len(part14)+
				len(part15)+
				len(part16)+
				len(part17)+
				len(part18)+
				len(part19)+
				len(part20)+
				0)/64,
		)
	}

	result = append(result, part01...)
	result = append(result, part02...)
	result = append(result, part03...)
	result = append(result, part04...)
	result = append(result, part05...)
	result = append(result, part06...)
	result = append(result, part07...)
	result = append(result, part08...)
	result = append(result, part09...)
	result = append(result, part10...)
	result = append(result, part11...)
	result = append(result, part12...)
	result = append(result, part13...)
	result = append(result, part14...)
	result = append(result, part15...)
	result = append(result, part16...)
	result = append(result, part17...)
	result = append(result, part18...)
	result = append(result, part19...)
	result = append(result, part20...)
	result = append(result, outsourceCompanies...)
	result = append(result, recruitingAgencies...)

	return result
}

func CloneCompanies() []domain.CompanyProfile {
	return slices.Clone(companies)
}
