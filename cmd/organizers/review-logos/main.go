package main

import (
	"encoding/json"
	"fmt"
	"go/format"
	"os"
	"path/filepath"
	"slices"
	"strings"

	"github.com/readytotouch/readytotouch/internal/domain"
	resize "github.com/readytotouch/readytotouch/internal/logo-resize"
	"github.com/readytotouch/readytotouch/internal/organizer/db"
	"github.com/readytotouch/readytotouch/internal/templates/dev"
)

const (
	mapping = "./public/logos/mapping.json"
)

type Logo struct {
	Alias            string `json:"alias"`
	Source           string `json:"source"`
	Destination      string `json:"destination"`
	Destination72x72 string `json:"destination_72x72"`
}

func main() {
	syncLogos(db.CloneCompanies())
}

func syncLogos(companies []domain.CompanyProfile) {
	logos := assertFetchLogos()

	logoMap := make(map[string]Logo, len(logos))
	for _, logo := range logos {
		if logo.Alias == "" {
			panic("logo alias cannot be empty")
		}

		if _, ok := logoMap[logo.Alias]; ok {
			panic(fmt.Sprintf("duplicate logo alias: %s", logo.Alias))
		}
		logoMap[logo.Alias] = logo
	}

	for _, company := range companies {
		if company.Ignore {
			continue
		}

		logo := logoMap[company.LinkedInProfile.Alias]
		logo.Alias = company.LinkedInProfile.Alias
		logoMap[company.LinkedInProfile.Alias] = logo

		if logo.Source == "" {
			// Nothing to resize

			continue
		}

		if logo.Destination != "" && logo.Destination72x72 != "" {
			// Logo already resized

			continue
		}

		ext := filepath.Ext(logo.Source)

		switch ext {
		case ".png", ".jpg", ".jpeg":
			// NOP
		case ".webp":
			ext = ".png" // Convert webp to png
		default:
			panic(fmt.Sprintf("unexpected ext: %s for logo: %s", ext, logo.Source))
		}

		logo.Destination = logo.Alias + ext
		logo.Destination72x72 = logo.Destination

		{
			err := resize.Resize112x56("./public/logos/original/"+logo.Source, "./public/logos/112x56/"+logo.Destination)
			if err != nil {
				panic(fmt.Sprintf("failed to resize logo %s: %v", logo.Source, err))
			}
		}

		{
			err := resize.Resize72x72("./public/logos/original/"+logo.Source, "./public/logos/72x72/"+logo.Destination)
			if err != nil {
				panic(fmt.Sprintf("failed to resize logo %s: %v", logo.Source, err))
			}
		}

		logoMap[company.LinkedInProfile.Alias] = logo
	}

	logos = make([]Logo, 0, len(logoMap))
	for _, logo := range logoMap {
		logos = append(logos, logo)
	}
	slices.SortFunc(logos, func(a, b Logo) int {
		return strings.Compare(a.Alias, b.Alias)
	})

	assertStoreLogos(logos)

	generateLogos(logos)
}

func generateLogos(logos []Logo) {
	aliasImagePairs := make([]*dev.CompanyLogoPair, len(logos))
	for i, logo := range logos {
		aliasImagePairs[i] = &dev.CompanyLogoPair{
			Alias: logo.Alias,
			Logo:  logo.Destination,
		}
	}

	output, err := format.Source([]byte(dev.CompanyLogo(aliasImagePairs, 2)))
	if err != nil {
		panic(err)
	}

	err = os.WriteFile("./internal/generated/organizers/company_logo_v2.go", output, 0644)
	if err != nil {
		panic(err)
	}
}

func assertFetchLogos() []Logo {
	var (
		logos []Logo
	)

	file, err := os.Open(mapping)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	err = json.NewDecoder(file).Decode(&logos)
	if err != nil {
		panic(err)
	}

	return logos
}

func assertStoreLogos(logos []Logo) {
	file, err := os.Create(mapping)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	data, err := json.MarshalIndent(logos, "", "  ")
	if err != nil {
		panic(err)
	}

	_, err = file.Write(data)
	if err != nil {
		panic(err)
	}
}
