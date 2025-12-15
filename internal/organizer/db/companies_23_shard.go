package db

import "github.com/readytotouch/readytotouch/internal/domain"

func companies23Shard() []domain.CompanyProfile {
	return []domain.CompanyProfile{
		{
			ID:         0,  // system
			Type:       "", // system
			Name:       "Alma Security",
			BaseURL:    "https://alma.security/",
			CareersURL: "https://alma.security/Career",
			AboutURL:   "https://alma.security/About-us",
			BlogURL:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                100717639,
				IDs:               nil,
				Alias:             "alma-security",
				Name:              "Alma Security",
				Followers:         "3K",
				Employees:         "11-50",
				AssociatedMembers: "21",
				Verified:          false,
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
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Backend Developer, Golang",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4350210925/",
							Location:             "Tel Aviv District, Israel",
							Date:                 mustDate("2025-12-15"),
							WithSalary:           false,
							Remote:               false,
						},
					},
				},
				domain.Rust: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Zig: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Scala: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Elixir: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Clojure: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Haskell: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Erlang: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
			},
			ShortDescription: "",
			Industries: []domain.Industry{
				domain.IndustryCyberSecurity,
			},
		},
		{
			ID:         0,  // system
			Type:       "", // system
			Name:       "ScorePlay",
			BaseURL:    "https://www.scoreplay.io/",
			CareersURL: "https://www.scoreplay.io/careers",
			AboutURL:   "https://www.scoreplay.io/company",
			BlogURL:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                13009905,
				IDs:               nil,
				Alias:             "scoreplay",
				Name:              "ScorePlay",
				Followers:         "12K",
				Employees:         "51-200",
				AssociatedMembers: "61",
				Verified:          false,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:     "ScorePlay-Inc",
				Followers: "12",
				Verified:  true,
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
					GitHubRepositoriesCount: 1,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Senior Back-end Engineer (GO)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4343636595/",
							Location:             "Paris, Île-de-France, France",
							Date:                 mustDate("2025-12-15"),
							WithSalary:           false,
							Remote:               false,
						},
					},
				},
				domain.Rust: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Zig: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Scala: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Elixir: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Clojure: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Haskell: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Erlang: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
			},
			ShortDescription: "ScorePlay builds the media backbone for sports organization",
			Industries: []domain.Industry{
				domain.IndustryEntertainment,
			},
		},
		{
			ID:         0,  // system
			Type:       "", // system
			Name:       "Open Cosmos",
			BaseURL:    "https://www.open-cosmos.com/",
			CareersURL: "https://www.open-cosmos.com/careers",
			AboutURL:   "https://www.open-cosmos.com/about",
			BlogURL:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                10352999,
				IDs:               nil,
				Alias:             "opencosmos",
				Name:              "Open Cosmos",
				Followers:         "30K",
				Employees:         "51-200",
				AssociatedMembers: "193",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:     "opencosmos",
				Followers: "12",
				Verified:  false,
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
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Junior Backend Engineer",
							ShortDescription:     "",
							SwitchingOpportunity: "Experience or strong interest in backend software development (Go, Python, or similar)",
							URL:                  "https://app.welcometothejungle.com/jobs/_EoB7Jm4",
							Date:                 mustDate("2025-12-16"),
							WithSalary:           false,
							Remote:               false,
						},
					},
				},
				domain.Rust: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Zig: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Scala: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Elixir: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Clojure: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Haskell: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Erlang: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
			},
			ShortDescription: "Open Cosmos provides simple and affordable satellite missions",
			Industries:       []domain.Industry{
				// NOP
			},
		},
		{
			ID:         0,  // system
			Type:       "", // system
			Name:       "Axiom",
			BaseURL:    "https://axiom.co/",
			CareersURL: "",
			AboutURL:   "https://axiom.co/company",
			BlogURL:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                18582643,
				IDs:               nil,
				Alias:             "axiomhq",
				Name:              "Axiom",
				Followers:         "6K",
				Employees:         "11-50",
				AssociatedMembers: "59",
				Verified:          true,
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
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Rust: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Software Engineer",
							ShortDescription:     "2+ years experience with the Rust programming language",
							SwitchingOpportunity: "",
							URL:                  "https://app.welcometothejungle.com/jobs/sLh3GI7U",
							Date:                 mustDate("2025-12-16"),
							WithSalary:           false,
							Remote:               false,
						},
					},
				},
				domain.Zig: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Scala: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Elixir: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Clojure: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Haskell: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Erlang: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
			},
			ShortDescription: "Axiom has developed its own cloud native data storage solution that its says is easier to integrate and scale than its competitors",
			Industries: []domain.Industry{
				domain.IndustryDevOps,
			},
		},

		// Template
		/*
			{
				ID:         0,  // system
				Type:       "", // system
				Name:       "",
				BaseURL:    "",
				CareersURL: "",
				AboutURL:   "",
				BlogURL:    "",
				LinkedInProfile: domain.LinkedInProfile{
					ID:                0,
					Alias:             "",
					Name:              "",
					Followers:         "",
					Employees:         "",
					AssociatedMembers: "",
					Verified:          false,
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
						GitHubRepositoriesCount: 0,
						Vacancies:               []domain.Vacancy{},
					},
					domain.Rust: {
						GitHubRepositoriesCount: 0,
						Vacancies:               []domain.Vacancy{},
					},
					domain.Zig: {
						GitHubRepositoriesCount: 0,
						Vacancies:               []domain.Vacancy{},
					},
					domain.Scala: {
						GitHubRepositoriesCount: 0,
						Vacancies:               []domain.Vacancy{},
					},
					domain.Elixir: {
						GitHubRepositoriesCount: 0,
						Vacancies:               []domain.Vacancy{},
					},
					domain.Clojure: {
						GitHubRepositoriesCount: 0,
						Vacancies:               []domain.Vacancy{},
					},
					domain.Haskell: {
						GitHubRepositoriesCount: 0,
						Vacancies:               []domain.Vacancy{},
					},
					domain.Erlang: {
						GitHubRepositoriesCount: 0,
						Vacancies:               []domain.Vacancy{},
					},
				},
				ShortDescription: "",
				Industries:       []domain.Industry{
					// NOP
				},
			},
		*/
	}
}
