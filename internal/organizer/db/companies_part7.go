package db

import (
	"github.com/readytotouch/readytotouch/internal/domain"
)

func companiesPart7() []domain.CompanyProfile {
	return []domain.CompanyProfile{

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Checkmarx",
			Website: "https://checkmarx.com/",
			Careers: "https://checkmarx.com/company/careers/",
			About:   "https://checkmarx.com/company/about-checkmarx/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                223472,
				Alias:             "checkmarx",
				Name:              "Checkmarx",
				Followers:         "109K",
				Employees:         "501-1K",
				AssociatedMembers: "902",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "checkmarx",
				Verified: true,
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 12,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Senior Go Developer",
							ShortDescription:     "With Golang as the primary technology, you will cross skill yourself in other programming languages like Java and TypeScript/JavaScript and technologies required in the project",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4114108035/",
							Date:                 mustDate("2025-01-02"),
							WithSalary:           false,
						},
					},
				},
				domain.Rust:    {},
				domain.Zig:     {},
				domain.Scala:   {},
				domain.Elixir:  {},
				domain.Clojure: {},
				domain.Haskell: {},
			},
			ShortDescription: "",
			Industries:       []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
			},
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Visa",
			Website: "https://visa.com/",
			Careers: "https://corporate.visa.com/en/careers.html",
			About:   "https://corporate.visa.com/en/about-visa.html",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                2190,
				Alias:             "visa",
				Name:              "Visa",
				Followers:         "2M",
				Employees:         "10K+",
				AssociatedMembers: "28,261",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "visa",
				Verified: false,
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 1,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Staff Systems Engineer — Go",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4056869486/",
							Date:                 mustDate("2025-01-03"),
							WithSalary:           false,
							Remote:               false,
						},
					},
				},
				domain.Rust:    {},
				domain.Zig:     {},
				domain.Scala:   {},
				domain.Elixir:  {},
				domain.Clojure: {},
				domain.Haskell: {},
			},
			ShortDescription: "",
			Industries: []domain.Industry{
				domain.IndustryFinTech,
			},
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
				domain.Czechia,
			},
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Datazip",
			Website: "https://datazip.io/",
			Careers: "",
			About:   "",
			Blog:    "https://datazip.io/blog",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                80069364,
				Alias:             "datazipio",
				Name:              "Datazip",
				Followers:         "8K",
				Employees:         "11-50",
				AssociatedMembers: "19",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "datazip-inc",
				Verified: false,
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 3,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Founding Backend Developer (Golang)",
							ShortDescription:     "Write clean, efficient, and well-documented code in Golang (must) and Python",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4114492171/",
							Date:                 mustDate("2025-01-03"),
							WithSalary:           false,
							Remote:               false,
						},
					},
				},
				domain.Rust:    {},
				domain.Zig:     {},
				domain.Scala:   {},
				domain.Elixir:  {},
				domain.Clojure: {},
				domain.Haskell: {},
			},
			ShortDescription:          "Composable Lakehouse Platform",
			Industries:                []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{},
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Baker Hughes",
			Website: "https://www.bakerhughes.com/",
			Careers: "https://careers.bakerhughes.com/",
			About:   "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                4334,
				Alias:             "bakerhughes",
				Name:              "Baker Hughes",
				Followers:         "3M",
				Employees:         "10K+",
				AssociatedMembers: "64,211",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "BakerHughes",
				Verified: false,
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Senior Golang Developer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4106957208/",
							Date:                 mustDate("2025-01-03"),
							WithSalary:           false,
							Remote:               false,
						},
					},
				},
				domain.Rust:    {},
				domain.Zig:     {},
				domain.Scala:   {},
				domain.Elixir:  {},
				domain.Clojure: {},
				domain.Haskell: {},
			},
			ShortDescription: "Baker Hughes is an energy technology company that provides solutions for energy and industrial customers worldwide",
			YahooFinanceURL:  "https://finance.yahoo.com/quote/BKR/",
			GoogleFinanceURL: "https://www.google.com/finance/quote/BKR:NASDAQ",
			Industries:       []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
				domain.Czechia,
			},
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Centrica",
			Website: "https://www.centrica.com/",
			Careers: "https://www.lifeatcentrica.com/jobs/",
			About:   "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                2766,
				Alias:             "centrica",
				Name:              "Centrica",
				Followers:         "143K",
				Employees:         "10K+",
				AssociatedMembers: "19,604",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "",
				Verified: false,
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Software Engineer (Golang)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4076005752/",
							Date:                 mustDate("2025-01-03"),
							WithSalary:           true, // Enjoy a generous market salary of up to £60,000
							Remote:               false,
						},
					},
				},
				domain.Rust:    {},
				domain.Zig:     {},
				domain.Scala:   {},
				domain.Elixir:  {},
				domain.Clojure: {},
				domain.Haskell: {},
			},
			ShortDescription:          "Centrica is a leading energy services and solutions provider founded on a 200-year heritage of serving people",
			Industries:                []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{},
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Mural Health",
			Website: "https://www.muralhealth.com/",
			Careers: "https://www.muralhealth.com/careers",
			About:   "https://www.muralhealth.com/about-us",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                89290648,
				Alias:             "muralhealth",
				Name:              "Mural Health",
				Followers:         "23K",
				Employees:         "11-50",
				AssociatedMembers: "36",
				Verified:          false,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "MuralHealth",
				Verified: false,
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Golang Engineer",
							ShortDescription:     "Assist in designing, developing, and maintaining efficient and scalable backend systems and APIs using Golang",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4115784363/",
							Date:                 mustDate("2025-01-03"),
							WithSalary:           false,
							Remote:               true,
						},
					},
				},
				domain.Rust:    {},
				domain.Zig:     {},
				domain.Scala:   {},
				domain.Elixir:  {},
				domain.Clojure: {},
				domain.Haskell: {},
			},
			ShortDescription: "Mural Health is a clinical trials technology company that is dedicated to making it easy to be a participant in clinical trials",
			Industries: []domain.Industry{
				domain.IndustryHealthTech,
			},
			HasEmployeesFromCountries: []domain.Country{},
		},

		// Template short
		//{
		//	ID:      0,  // system
		//	Type:    "", // system
		//	Name:    "",
		//	Website: "",
		//	Careers: "",
		//	About:   "",
		//	LinkedInProfile: domain.LinkedInProfile{
		//		ID:                0,
		//		Alias:             "",
		//		Name:              "",
		//		Followers:         "",
		//		Employees:         "",
		//		AssociatedMembers: "",
		//		Verified:          false,
		//	},
		//	GitHubProfile: domain.GitHubProfile{
		//		Login:    "",
		//		Verified: false,
		//	},
		//	Languages: domain.Languages{
		//		domain.Go: {
		//			GitHubRepositoriesCount: 0,
		//			Vacancies: []domain.Vacancy{
		//				{
		//					Title:                "",
		//					ShortDescription:     "",
		//					SwitchingOpportunity: "",
		//					URL:                  "",
		//					Date:                 mustDate(""),
		//					WithSalary:           false,
		//				},
		//			},
		//		},
		//		domain.Rust:    {},
		//		domain.Zig:     {},
		//		domain.Scala:   {},
		//		domain.Elixir:  {},
		//		domain.Clojure: {},
		//		domain.Haskell: {},
		//	},
		//	ShortDescription: "",
		//	Industries:       []domain.Industry{},
		//	HasEmployeesFromCountries: []domain.Country{
		//		domain.Ukraine,
		//		domain.Czechia,
		//	},
		//},
	}
}
