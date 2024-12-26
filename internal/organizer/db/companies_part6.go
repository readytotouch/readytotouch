package db

import (
	"github.com/readytotouch/readytotouch/internal/domain"
)

func companiesPart6() []domain.CompanyProfile {
	return []domain.CompanyProfile{

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Hasaki.vn",
			Website: "https://hasaki.vn/",
			Careers: "",
			About:   "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                31033180,
				Alias:             "hasaki-vn/people",
				Name:              "Hasaki.vn",
				Followers:         "9K",
				Employees:         "1K-5K",
				AssociatedMembers: "243",
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
							Title:            "Senior Back-end Developer (PHP/ GO)",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4109417330/",
							Date:             mustDate("2024-12-25"),
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
			ShortDescription: "Retail Health and Personal Care Products company",
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Bedrock Streaming",
			Website: "https://bedrockstreaming.com/",
			Careers: "https://bedrockstreaming.com/career/",
			About:   "https://bedrockstreaming.com/company/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                64853943,
				Alias:             "bedrock-streaming",
				Name:              "Bedrock Streaming",
				Followers:         "30K",
				Employees:         "201-500",
				AssociatedMembers: "315",
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
							Title:            "Software Developer — Backend GO",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4079483983/",
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
			ShortDescription: "Streaming-tech venture",
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Social Discovery Group",
			Website: "https://socialdiscoverygroup.com/",
			Careers: "https://socialdiscoverygroup.com/vacancies",
			About:   "https://socialdiscoverygroup.com/about-us",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                3683382,
				Alias:             "social-discovery-group",
				Name:              "Social Discovery Group",
				Followers:         "23K",
				Employees:         "1K-5K",
				AssociatedMembers: "773",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "SDVentures",
				Verified: false,
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:            "GO — developer",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4087164770/",
							Date:             mustDate("2024-12-05"),
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
			ShortDescription: "Global technology company",
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "",
			Website: "",
			Careers: "",
			About:   "",
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
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:            "",
							ShortDescription: "",
							URL:              "",
							Date:             mustDate(""),
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
		//	ShortDescription: "",
		//},
	}
}
