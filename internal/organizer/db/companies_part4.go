package db

import (
	"github.com/readytotouch/readytotouch/internal/domain"
)

func companiesPart4() []domain.CompanyProfile {
	return []domain.CompanyProfile{
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "limango GmbH - A member of the otto group",
			Website: "https://limango.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:       9381585,
				Alias:    "limango-gmbh",
				Name:     "limango GmbH - A member of the otto group",
				Verified: true,
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Senior GO developer (m/w/d) — 100% remote possible",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4096830999/",
							Date:             mustDate("2024-12-13"),
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
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Cycloid - Sustainable Platform Engineering",
			Website: "https://www.cycloid.io/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:       10071522,
				Alias:    "cycloid",
				Name:     "Cycloid - Sustainable Platform Engineering",
				Verified: true,
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 40,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Senior Backend GO Developer - Remote",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4075626268/",
							Date:             mustDate("2024-12-12"),
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
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Varonis",
			Website: "https://www.varonis.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:       25403,
				Alias:    "varonis",
				Name:     "Varonis",
				Verified: true,
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Backend Developer (Python, GO)",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4093559031/",
							Date:             mustDate("2024-12-10"),
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
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "WorldTech IT, LLC",
			Website: "https://wtit.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:       1811897,
				Alias:    "worldtechit",
				Name:     "WorldTech IT, LLC",
				Verified: true,
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Golang Software Engineer",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4101282630/",
							Date:             mustDate("2024-12-17"),
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
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Applied Research Solutions",
			Website: "https://www.appliedres.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:       6017665,
				Alias:    "applied-research-solutions",
				Name:     "Applied Research Solutions",
				Verified: true,
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 11,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Senior GO Developer",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4101198898/",
							Date:             mustDate("2024-12-16"),
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
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "WEX",
			Website: "https://www.wexinc.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:       11637,
				Alias:    "wexinc",
				Name:     "WEX",
				Verified: true,
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Specialist Java/Golang Developer",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4048143482/",
							Date:             mustDate("2024-12-14"),
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
		},
		// Template short
		//{
		//	ID:      0,  // system
		//	Type:    "", // system
		//	Name:    "",
		//	Website: "",
		//	LinkedInProfile: domain.LinkedInProfile{
		//		ID:       0,
		//		Alias:    "",
		//		Name:     "",
		//		Verified: false,
		//	},
		//	Languages: domain.Languages{
		//		domain.Go: {
		//			GitHubRepositoriesCount: 0,
		//			Vacancies: []domain.Vacancy{
		//				{
		//					Title:            "",
		//					ShortDescription: "",
		//					URL:              "",
		//					Date:             mustDate(""),
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
		//},
	}
}
