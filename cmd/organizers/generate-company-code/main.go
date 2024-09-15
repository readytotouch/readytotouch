package main

import (
	"os"

	"github.com/readytotouch/readytotouch/internal/organizer/db"
	"github.com/readytotouch/readytotouch/internal/protos/organizers"
	"github.com/readytotouch/readytotouch/internal/templates/dev"
)

func main() {
	var maxCompanyCode int32 = 0

	for code := range organizers.CompanyCode_name {
		maxCompanyCode = max(maxCompanyCode, code)
	}

	err := os.WriteFile(
		"./protos/organizers/company_code.proto",
		[]byte(dev.CompanyCode(organizers.CompanyCode_name, maxCompanyCode, db.Companies())),
		0644,
	)
	if err != nil {
		panic(err)
	}
}
