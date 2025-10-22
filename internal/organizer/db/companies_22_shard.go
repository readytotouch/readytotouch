package db

import "github.com/readytotouch/readytotouch/internal/domain"

func companies22Shard() []domain.CompanyProfile {
	return []domain.CompanyProfile{
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Centra",
			Website: "https://centra.com/",
			Careers: "https://centra.teamtailor.com/",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                10813841,
				IDs:               nil,
				Alias:             "centra-commerce",
				Name:              "Centra",
				Followers:         "11K",
				Employees:         "51-200",
				AssociatedMembers: "313",
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
					Vacancies: []domain.Vacancy{
						{
							Title:                "Staff Engineer Go",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4310350066/",
							Location:             "Wrocław, Dolnośląskie, Poland",
							Date:                 mustDate("2025-10-08"),
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
			ShortDescription: "The fashion brand commerce platform",
			Industries: []domain.Industry{
				domain.IndustryECommerce,
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Lifemote",
			Website: "https://www.lifemote.com/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                15143129,
				IDs:               nil,
				Alias:             "lifemote-networks",
				Name:              "Lifemote",
				Followers:         "5K",
				Employees:         "11-50",
				AssociatedMembers: "61",
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
							Title:                "Backend Engineer – Go Developer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4311254582/",
							Location:             "Türkiye",
							Date:                 mustDate("2025-10-08"),
							WithSalary:           false,
							Remote:               true,
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
			ShortDescription: "WiFi Analytics and Proactive Maintenance for Internet Service Providers to find and fix home network issues",
			Industries:       []domain.Industry{
				// NOP
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Chili Piper",
			Website: "https://www.chilipiper.com/",
			Careers: "https://www.chilipiper.com/careers",
			About:   "https://www.chilipiper.com/about",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                2687110,
				IDs:               nil,
				Alias:             "chilipiper",
				Name:              "Chili Piper",
				Followers:         "81K",
				Employees:         "51-200",
				AssociatedMembers: "256",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:     "Chili-Piper",
				Followers: "21",
				Verified:  true,
			},
			BlindProfile: domain.BlindProfile{
				Alias: "",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "chili-piper",
				Employees: "180",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Chili-Piper-EI_IE1835737.11,22.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Chili-Piper-Reviews-E1835737.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Chili-Piper-Jobs-E1835737.htm",
				Jobs:        "",
				Reviews:     "100",
				Salaries:    "151",
				ReviewsRate: "4.5",
				Verified:    true,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "Chili-Piper",
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 1,
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
					GitHubRepositoriesCount: 5,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Scala Engineer",
							SubTitle:             "",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							Location:             "",
							URL:                  "https://www.chilipiper.com/careers?ashby_jid=ab556557-83cf-467d-90fb-5119dabf146c#work",
							Date:                 mustDate("2025-10-07"),
							WithSalary:           true, // $87K – $138K
							Remote:               true,
						},
					},
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
			ShortDescription: "Demand Conversion Platform",
			Industries:       []domain.Industry{
				// NOP
			},
			PinnedUntil: mustDate("2025-10-31"),
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "OpenVPN Inc.",
			Website: "https://openvpn.net/",
			Careers: "https://careers.openvpn.net/",
			About:   "https://openvpn.net/about/",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                15353748,
				IDs:               nil,
				Alias:             "openvpn",
				Name:              "OpenVPN Inc.",
				Followers:         "7K",
				Employees:         "51-200",
				AssociatedMembers: "183",
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
					Vacancies: []domain.Vacancy{
						{
							Title:                "Senior Software Engineer (Golang)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4315588096/",
							Location:             "European Union",
							Date:                 mustDate("2025-10-22"),
							WithSalary:           false,
							Remote:               true,
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

		// Template
		//{
		//	ID:      0,  // system
		//	Type:    "", // system
		//	Name:    "",
		//	Website: "",
		//	Careers: "",
		//	About:   "",
		//	Blog:    "",
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
		//	BlindProfile: domain.BlindProfile{
		//		Alias: "",
		//	},
		//	LevelsFyiProfile: domain.LevelsFyiProfile{
		//		Alias: "",
		//	},
		//	GlassdoorProfile: domain.GlassdoorProfile{
		//		OverviewURL: "",
		//	},
		//	IndeedProfile: domain.IndeedProfile{
		//		Alias: "",
		//	},
		//	Languages: domain.Languages{
		//		domain.Go: {
		//			GitHubRepositoriesCount: 0,
		//			Vacancies:               []domain.Vacancy{},
		//		},
		//		domain.Rust: {
		//			GitHubRepositoriesCount: 0,
		//			Vacancies:               []domain.Vacancy{},
		//		},
		//		domain.Zig: {
		//			GitHubRepositoriesCount: 0,
		//			Vacancies:               []domain.Vacancy{},
		//		},
		//		domain.Scala: {
		//			GitHubRepositoriesCount: 0,
		//			Vacancies:               []domain.Vacancy{},
		//		},
		//		domain.Elixir: {
		//			GitHubRepositoriesCount: 0,
		//			Vacancies:               []domain.Vacancy{},
		//		},
		//		domain.Clojure: {
		//			GitHubRepositoriesCount: 0,
		//			Vacancies:               []domain.Vacancy{},
		//		},
		//		domain.Haskell: {
		//			GitHubRepositoriesCount: 0,
		//			Vacancies:               []domain.Vacancy{},
		//		},
		//		domain.Erlang: {
		//			GitHubRepositoriesCount: 0,
		//			Vacancies:               []domain.Vacancy{},
		//		},
		//	},
		//	ShortDescription: "",
		//	Industries:       []domain.Industry{
		//		// NOP
		//	},
		//},
	}
}
