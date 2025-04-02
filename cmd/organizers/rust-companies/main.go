package main

import (
	"fmt"
	"os"
	"slices"
	"sort"
	"strings"

	"github.com/readytotouch/readytotouch/internal/domain"
	"github.com/readytotouch/readytotouch/internal/organizer/db"
)

func main() {
	var (
		companies = db.CloneCompanies()
		verified  = os.Getenv("VERIFIED") == "true"
		github    = os.Getenv("GITHUB") == "true"
	)

	// filter with Rust vacancies and not ignore

	var ss = make([]string, 0, len(companies))
	for _, company := range companies {
		if verified {
			if company.Ignore {
				continue
			}

			if !company.GitHubProfile.Verified {
				continue
			}
		}

		if len(company.Languages[domain.Rust].Vacancies) == 0 && !slices.Contains(company.SyncSources, domain.RustCompanies) {
			continue
		}

		if github {
			ss = append(ss, fmt.Sprintf("%s %s https://github.com/%s", company.Name, company.Website, company.GitHubProfile.Login))
		} else {
			ss = append(ss, fmt.Sprintf("%s %s", company.Name, company.Website))
		}

	}

	// Used for development purposes
	sort.Slice(ss, func(i, j int) bool {
		return strings.ToLower(ss[i]) < strings.ToLower(ss[j])
	})

	for _, s := range ss {
		fmt.Println(s)
	}
}
