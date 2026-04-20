package db

import (
	"github.com/readytotouch/readytotouch/internal/domain"
)

func companies25Shard() []domain.CompanyProfile {
	return []domain.CompanyProfile{
		{
			ID:         0,  // system
			Type:       "", // system
			Name:       "Trendyol Group",
			BaseURL:    "https://www.trendyol.com/",
			CareersURL: "https://www.trendyol.com/careers",
			AboutURL:   "",
			BlogURL:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                1127847,
				IDs:               nil,
				Alias:             "trendyolgroup",
				Name:              "Trendyol Group",
				Followers:         "555K",
				Employees:         "10K+",
				AssociatedMembers: "11,321",
				Verified:          true,
				Date:              mustDate("2026-04-01"),
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "",
				Verified: false,
			},
			BlindProfile: domain.BlindProfile{
				Alias: "",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias: "",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Trendyol-EI_IE722443.11,19.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Trendyol-Reviews-E722443.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Trendyol-Jobs-E722443.htm",
				Jobs:        "", // Becomes empty in the new design version
				Reviews:     "", // Becomes empty in the new design version
				Salaries:    "", // Becomes empty in the new design version
				ReviewsRate: "",
				Verified:    false,
				Date:        mustDate("2026-04-03"),
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "",
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoryCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Go Developer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4390917180/",
							Location:             "Istanbul, Türkiye",
							CloudProviders:       []domain.CloudProvider{},
							Date:                 mustDate("2026-03-28"),
							WithSalary:           false,
							Remote:               false,
						},
					},
				},
				domain.Rust: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Zig: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Scala: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Elixir: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Clojure: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Haskell: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Erlang: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
			},
			ShortDescription: "e-commerce platform in Türkiye",
			Industries: []domain.Industry{
				domain.IndustryECommerce,
			},
		},
		{
			ID:         0,  // system
			Type:       "", // system
			Name:       "Vivino",
			BaseURL:    "https://www.vivino.com/",
			CareersURL: "https://careers.vivino.com/",
			AboutURL:   "https://www.vivino.com/about",
			BlogURL:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                1148120,
				IDs:               nil,
				Alias:             "vivino",
				Name:              "Vivino",
				Followers:         "56K",
				Employees:         "51-200",
				AssociatedMembers: "291",
				Verified:          true,
				Date:              mustDate("2026-04-03"),
			},
			GitHubProfile: domain.GitHubProfile{
				Login:     "Vivino",
				Followers: "73",
				Verified:  true,
				Date:      mustDate("2026-04-03"),
			},
			BlindProfile: domain.BlindProfile{
				Alias: "",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias: "",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Vivino-EI_IE1350937.11,17.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Vivino-Reviews-E1350937.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Vivino-Jobs-E1350937.htm",
				Jobs:        "", // Becomes empty in the new design version
				Reviews:     "", // Becomes empty in the new design version
				Salaries:    "", // Becomes empty in the new design version
				ReviewsRate: "",
				Verified:    false,
				Date:        mustDate("2026-04-03"),
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "",
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoryCount: 27,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Senior Backend Engineer (Go)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4396662660/",
							Location:             "Copenhagen, Capital Region of Denmark, Denmark",
							CloudProviders:       []domain.CloudProvider{},
							Date:                 mustDate("2026-04-02"),
							WithSalary:           false,
							Remote:               false,
						},
					},
				},
				domain.Rust: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Zig: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Scala: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Elixir: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Clojure: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Haskell: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Erlang: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
			},
			ShortDescription: "",
			Industries:       []domain.Industry{
				// NOP
			},
		},
		{
			ID:         0,  // system
			Type:       "", // system
			Name:       "Chromatic",
			BaseURL:    "https://www.chromatic.com/",
			CareersURL: "https://www.chromatic.com/company/careers",
			AboutURL:   "https://www.chromatic.com/company/about",
			BlogURL:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                18289391,
				IDs:               nil,
				Alias:             "chromaticcom",
				Name:              "Chromatic",
				Followers:         "4K",
				Employees:         "11-50",
				AssociatedMembers: "67",
				Verified:          true,
				Date:              mustDate("2026-04-11"),
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "",
				Verified: false,
			},
			BlindProfile: domain.BlindProfile{
				Alias: "",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias: "",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "",
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "",
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Rust: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Zig: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Scala: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Elixir: {
					GitHubRepositoryCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Backend Engineer (Elixir)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4399868227/",
							Location:             "United States",
							CloudProviders:       []domain.CloudProvider{},
							Date:                 mustDate("2026-04-09"),
							WithSalary:           false,
							Remote:               true,
						},
					},
				},
				domain.Clojure: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Haskell: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Erlang: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
			},
			ShortDescription: "Frontend UI Testing and Review Platform for Teams",
			Industries:       []domain.Industry{
				// NOP
			},
		},
		{
			ID:         0,  // system
			Type:       "", // system
			Name:       "OpenYield",
			BaseURL:    "http://www.openyld.com/",
			CareersURL: "https://openyld.com/careers/",
			AboutURL:   "https://openyld.com/about-us/",
			BlogURL:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                92428976,
				IDs:               nil,
				Alias:             "openyield",
				Name:              "OpenYield",
				Followers:         "3K",
				Employees:         "2-10",
				AssociatedMembers: "8",
				Verified:          false,
				Date:              mustDate("2026-04-11"),
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "",
				Verified: false,
			},
			BlindProfile: domain.BlindProfile{
				Alias: "",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias: "",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "",
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "",
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Rust: {
					GitHubRepositoryCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Senior C++/Rust Developer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4395949066/",
							Location:             "New York, NY",
							CloudProviders:       []domain.CloudProvider{domain.AWS},
							Date:                 mustDate("2026-04-08"),
							WithSalary:           false,
							Remote:               false,
						},
					},
				},
				domain.Zig: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Scala: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Elixir: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Clojure: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Haskell: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Erlang: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
			},
			ShortDescription: "Bond marketplace",
			Industries: []domain.Industry{
				domain.IndustryFinTech,
			},
		},
		{
			ID:         0,  // system
			Type:       "", // system
			Name:       "Atomic Semi",
			BaseURL:    "https://atomicsemi.com/",
			CareersURL: "https://atomicsemi.com/careers/",
			AboutURL:   "https://atomicsemi.com/about/",
			BlogURL:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                89976423,
				IDs:               nil,
				Alias:             "atomic-semi",
				Name:              "Atomic Semi",
				Followers:         "8K",
				Employees:         "51-200",
				AssociatedMembers: "78",
				Verified:          false,
				Date:              mustDate("2026-04-11"),
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "",
				Verified: false,
			},
			BlindProfile: domain.BlindProfile{
				Alias: "",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias: "",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Atomic-Semi-EI_IE9926998.11,22.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Atomic-Semi-Reviews-E9926998.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Atomic-Semi-Jobs-E9926998.htm",
				Jobs:        "", // Becomes empty in the new design version
				Reviews:     "", // Becomes empty in the new design version
				Salaries:    "", // Becomes empty in the new design version
				ReviewsRate: "",
				Verified:    true,
				Date:        mustDate("2026-04-12"),
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "",
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Rust: {
					GitHubRepositoryCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Rust Software Engineer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4400690851/",
							Location:             "San Francisco, CA",
							CloudProviders:       []domain.CloudProvider{},
							Date:                 mustDate("2026-04-11"),
							WithSalary:           true, // $125k/yr - $195k/yr
							Remote:               false,
						},
						{
							Title:                "Software Engineer, Rust",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4401485694/",
							Location:             "San Francisco, CA",
							CloudProviders:       []domain.CloudProvider{},
							Date:                 mustDate("2026-04-15"),
							WithSalary:           false,
							Remote:               false,
						},
					},
				},
				domain.Zig: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Scala: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Elixir: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Clojure: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Haskell: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Erlang: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
			},
			ShortDescription: "Atomic Semi is building a small, fast semiconductor fab",
			Industries:       []domain.Industry{
				// NOP
			},
		},
		{
			ID:         0,  // system
			Type:       "", // system
			Name:       "Termius",
			BaseURL:    "https://termius.com/",
			CareersURL: "https://termius.com/careers",
			AboutURL:   "https://termius.com/about",
			BlogURL:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                14507352,
				IDs:               nil,
				Alias:             "termius",
				Name:              "Termius",
				Followers:         "3K",
				Employees:         "11-50",
				AssociatedMembers: "30",
				Verified:          true,
				Date:              mustDate("2026-04-11"),
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "",
				Verified: false,
			},
			BlindProfile: domain.BlindProfile{
				Alias: "",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias: "",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "",
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "",
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Rust: {
					GitHubRepositoryCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Senior Product Engineer (Rust)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4398683341/",
							Location:             "Spain",
							CloudProviders:       []domain.CloudProvider{},
							Date:                 mustDate("2026-04-09"),
							WithSalary:           false,
							Remote:               true,
						},
						{
							Title:                "Senior Product Engineer (Rust)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4402021996/",
							Location:             "Auckland, Auckland, New Zealand",
							CloudProviders:       []domain.CloudProvider{},
							Date:                 mustDate("2026-04-14"),
							WithSalary:           false,
							Remote:               false,
						},
					},
				},
				domain.Zig: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Scala: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Elixir: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Clojure: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Haskell: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Erlang: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
			},
			ShortDescription: "Termius is an SSH client",
			Industries: []domain.Industry{
				domain.IndustryDevOps,
			},
		},
		{
			ID:         0,  // system
			Type:       "", // system
			Name:       "Tentris",
			BaseURL:    "https://tentris.io/",
			CareersURL: "",
			AboutURL:   "",
			BlogURL:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                83062420,
				IDs:               nil,
				Alias:             "tentris",
				Name:              "Tentris",
				Followers:         "526",
				Employees:         "2-10",
				AssociatedMembers: "4",
				Verified:          false,
				Date:              mustDate("2026-04-16"),
			},
			GitHubProfile: domain.GitHubProfile{
				Login:     "tentris",
				Followers: "9",
				Verified:  false,
				Date:      mustDate("2026-04-16"),
			},
			BlindProfile: domain.BlindProfile{
				Alias: "",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias: "",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "",
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "",
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Rust: {
					GitHubRepositoryCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Junior System Engineer",
							ShortDescription:     "Developing graph database internals in C++/Rust",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4399937525/",
							Location:             "Cologne, North Rhine-Westphalia, Germany",
							CloudProviders:       []domain.CloudProvider{},
							Date:                 mustDate("2026-04-15"),
							WithSalary:           false,
							Remote:               false,
						},
					},
				},
				domain.Zig: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Scala: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Elixir: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Clojure: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Haskell: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Erlang: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
			},
			ShortDescription: "RDF graph database",
			Industries: []domain.Industry{
				domain.IndustryDevOps,
			},
		},
		{
			ID:         0,  // system
			Type:       "", // system
			Name:       "Cato Networks",
			BaseURL:    "https://www.catonetworks.com/",
			CareersURL: "https://www.catonetworks.com/careers/",
			AboutURL:   "https://www.catonetworks.com/company/",
			BlogURL:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                3533853,
				IDs:               nil,
				Alias:             "cato-networks",
				Name:              "Cato Networks",
				Followers:         "191K",
				Employees:         "1K-5K",
				AssociatedMembers: "1,618",
				Verified:          true,
				Date:              mustDate("2026-04-16"),
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "",
				Verified: false,
			},
			BlindProfile: domain.BlindProfile{
				Alias: "",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias: "",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "",
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "",
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Rust: {
					GitHubRepositoryCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "C/Rust Engineer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4353962622/",
							Location:             "Tel Aviv District, Israel",
							CloudProviders:       []domain.CloudProvider{},
							Date:                 mustDate("2026-04-15"),
							WithSalary:           false,
							Remote:               false,
						},
					},
				},
				domain.Zig: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Scala: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Elixir: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Clojure: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Haskell: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Erlang: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
			},
			ShortDescription: "SASE platform",
			Industries: []domain.Industry{
				domain.IndustryCyberSecurity,
			},
		},
		{
			ID:         0,  // system
			Type:       "", // system
			Name:       "Consumer Edge",
			BaseURL:    "https://www.consumeredge.com/",
			CareersURL: "",
			AboutURL:   "https://www.consumeredge.com/about-us/",
			BlogURL:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                236535,
				IDs:               nil,
				Alias:             "consumeredge",
				Name:              "Consumer Edge",
				Followers:         "11K",
				Employees:         "51-200",
				AssociatedMembers: "153",
				Verified:          true,
				Date:              mustDate("2026-04-20"),
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "",
				Verified: false,
			},
			BlindProfile: domain.BlindProfile{
				Alias: "",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias: "",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "",
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "",
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Rust: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Zig: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Scala: {
					GitHubRepositoryCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Back-End Software Engineer (Scala)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4403871153/",
							Location:             "Dublin, County Dublin, Ireland",
							CloudProviders:       []domain.CloudProvider{domain.AWS, domain.GCP},
							Date:                 mustDate("2026-04-19"),
							WithSalary:           false,
							Remote:               true,
						},
					},
				},
				domain.Elixir: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Clojure: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Haskell: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Erlang: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
			},
			ShortDescription: "Consumer Edge (CE) is a consumer alternative data firm providing insights as a service solutions",
			Industries: []domain.Industry{
				domain.IndustryMarTech,
			},
		},
		{
			ID:         0,  // system
			Type:       "", // system
			Name:       "Prodege, LLC",
			BaseURL:    "https://www.prodege.com/",
			CareersURL: "https://www.prodege.com/careers/",
			AboutURL:   "https://www.prodege.com/about/",
			BlogURL:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                1015739,
				IDs:               []int{1015739, 3012494, 15863446},
				Alias:             "prod-g-llc",
				Name:              "Prodege, LLC",
				Followers:         "28K",
				Employees:         "501-1K",
				AssociatedMembers: "629",
				Verified:          true,
				Date:              mustDate("2026-04-20"),
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "",
				Verified: false,
			},
			BlindProfile: domain.BlindProfile{
				Alias: "",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias: "",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "",
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "",
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Rust: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Zig: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Scala: {
					GitHubRepositoryCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Principal Engineer (Scala)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4392022466/",
							Location:             "United States",
							CloudProviders:       []domain.CloudProvider{},
							Date:                 mustDate("2026-04-18"),
							WithSalary:           false,
							Remote:               true,
						},
					},
				},
				domain.Elixir: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Clojure: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Haskell: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Erlang: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
			},
			ShortDescription: "Consumer insights platform",
			Industries: []domain.Industry{
				domain.IndustryMarTech,
			},
		},
		{
			ID:         0,  // system
			Type:       "", // system
			Name:       "Astronomer",
			BaseURL:    "https://www.astronomer.io/",
			CareersURL: "https://www.astronomer.io/careers/",
			AboutURL:   "https://www.astronomer.io/about-us/",
			BlogURL:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                10019299,
				IDs:               nil,
				Alias:             "astronomer",
				Name:              "Astronomer",
				Followers:         "142K",
				Employees:         "201-500",
				AssociatedMembers: "4,619",
				Verified:          true,
				Date:              mustDate("2026-04-20"),
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "",
				Verified: false,
			},
			BlindProfile: domain.BlindProfile{
				Alias: "",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias: "",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "",
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "",
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoryCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Senior Software Engineer (Node.js, Golang)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4404321800/",
							Location:             "New York, NY",
							CloudProviders:       []domain.CloudProvider{},
							Date:                 mustDate("2026-04-20"),
							WithSalary:           false,
							Remote:               false,
						},
					},
				},
				domain.Rust: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Zig: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Scala: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Elixir: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Clojure: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Haskell: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Erlang: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
			},
			ShortDescription: "",
			Industries:       []domain.Industry{
				// NOP
			},
		},
		{
			ID:         0,  // system
			Type:       "", // system
			Name:       "Trustly",
			BaseURL:    "https://www.trustly.com/",
			CareersURL: "https://www.trustly.com/careers",
			AboutURL:   "https://www.trustly.com/what-is-trustly",
			BlogURL:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                2372346,
				IDs:               nil,
				Alias:             "trustly",
				Name:              "Trustly",
				Followers:         "78K",
				Employees:         "1K-5K",
				AssociatedMembers: "889",
				Verified:          true,
				Date:              mustDate("2026-04-21"),
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "",
				Verified: false,
			},
			BlindProfile: domain.BlindProfile{
				Alias: "",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias: "",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "",
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "",
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoryCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Senior Software Engineer (Golang)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4343925593/",
							Location:             "Stockholm, Stockholm County, Sweden",
							CloudProviders:       []domain.CloudProvider{},
							Date:                 mustDate("2026-04-18"),
							WithSalary:           false,
							Remote:               false,
						},
					},
				},
				domain.Rust: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Zig: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Scala: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Elixir: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Clojure: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Haskell: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Erlang: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
			},
			ShortDescription: "Online banking payments",
			Industries: []domain.Industry{
				domain.IndustryFinTech,
			},
		},
		{
			ID:         0,  // system
			Type:       "", // system
			Name:       "Prelude",
			BaseURL:    "https://prelude.so/",
			CareersURL: "",
			AboutURL:   "",
			BlogURL:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                92698744,
				IDs:               nil,
				Alias:             "prelude-so",
				Name:              "Prelude",
				Followers:         "3K",
				Employees:         "51-200",
				AssociatedMembers: "63",
				Verified:          false,
				Date:              mustDate("2026-04-21"),
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "",
				Verified: false,
			},
			BlindProfile: domain.BlindProfile{
				Alias: "",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias: "",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "",
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "",
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoryCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Backend Engineer (Go)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4403822707/",
							Location:             "Paris, Île-de-France, France",
							CloudProviders:       []domain.CloudProvider{},
							Date:                 mustDate("2026-04-19"),
							WithSalary:           false,
							Remote:               false,
						},
					},
				},
				domain.Rust: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Zig: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Scala: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Elixir: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Clojure: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Haskell: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Erlang: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
			},
			ShortDescription: "SMS & OTP verification infrastructure built for developers",
			Industries: []domain.Industry{
				domain.IndustryMarTech,
			},
		},
		{
			ID:         0,  // system
			Type:       "", // system
			Name:       "Barter",
			BaseURL:    "https://getbarter.com/",
			CareersURL: "https://getbarter.com/careers/",
			AboutURL:   "",
			BlogURL:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                93634266,
				IDs:               nil,
				Alias:             "getbarter-app",
				Name:              "Barter",
				Followers:         "2K",
				Employees:         "11-50",
				AssociatedMembers: "37",
				Verified:          false,
				Date:              mustDate("2026-04-21"),
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "",
				Verified: false,
			},
			BlindProfile: domain.BlindProfile{
				Alias: "",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias: "",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "",
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "",
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoryCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Back-end Developer (Go)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4403423575/",
							Location:             "Portugal",
							CloudProviders:       []domain.CloudProvider{},
							Date:                 mustDate("2026-04-18"),
							WithSalary:           false,
							Remote:               true,
						},
					},
				},
				domain.Rust: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Zig: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Scala: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Elixir: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Clojure: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Haskell: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Erlang: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
			},
			ShortDescription: "Exchange your product or service for content and online reach",
			Industries: []domain.Industry{
				domain.IndustryMarTech,
			},
		},
	}
}
