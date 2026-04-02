package db

import "github.com/readytotouch/readytotouch/internal/domain"

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
							Title:                "Senior Backend (Go) Engineer",
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
	}
}
