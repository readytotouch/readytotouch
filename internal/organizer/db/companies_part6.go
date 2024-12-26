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
				Alias:             "hasaki-vn",
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
			Name:    "Nutanix",
			Website: "https://www.nutanix.com/",
			Careers: "https://www.nutanix.com/company/careers",
			About:   "https://www.nutanix.com/company",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                735085,
				Alias:             "nutanix",
				Name:              "Nutanix",
				Followers:         "521K",
				Employees:         "5K-10K",
				AssociatedMembers: "8,633",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "nutanix",
				Verified: true,
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Software Engineer (C++/Golang/Systems Development)",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4110235695/",
							Date:             mustDate("2024-12-26"),
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
			ShortDescription: "Private and hybrid cloud software provider",
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Umbra",
			Website: "https://umbra.space/",
			Careers: "https://umbra.space/careers/",
			About:   "https://umbra.space/about/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                18111757,
				Alias:             "umbraspace",
				Name:              "Umbra",
				Followers:         "16K",
				Employees:         "51-200",
				AssociatedMembers: "156",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "",
				Verified: false,
			},
			Languages: domain.Languages{
				domain.Go: {},
				domain.Rust: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Senior Spacecraft Flight/Software Engineer (Rust)",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4105560497/",
							Date:             mustDate("2024-12-24"),
						},
					},
				},
				domain.Zig:     {},
				domain.Scala:   {},
				domain.Elixir:  {},
				domain.Clojure: {},
				domain.Haskell: {},
			},
			ShortDescription: "Creator of space technology",
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Cognite",
			Website: "https://www.cognite.com",
			Careers: "https://www.cognite.com/company/careers",
			About:   "https://www.cognite.com/en/company/about-us-cognite",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                18045548,
				Alias:             "cognitedata",
				Name:              "Cognite",
				Followers:         "58K",
				Employees:         "501-1K",
				AssociatedMembers: "657",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "cognitedata",
				Verified: true,
			},
			Languages: domain.Languages{
				domain.Go: {},
				domain.Rust: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Senior Software Engineer (Rust/Kotlin)",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4106227833/",
							Date:             mustDate("2024-12-20"),
						},
					},
				},
				domain.Zig:     {},
				domain.Scala:   {},
				domain.Elixir:  {},
				domain.Clojure: {},
				domain.Haskell: {},
			},
			ShortDescription: "Automates and scales industrial data contextualization ",
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
