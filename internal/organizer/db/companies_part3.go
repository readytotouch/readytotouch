package db

import (
	"github.com/readytotouch/readytotouch/internal/domain"
)

func companiesPart3() []domain.CompanyProfile {
	return []domain.CompanyProfile{
		{
			ID:   0,  // system
			Type: "", // system
			Name: "Capital One",
			URL:  "https://www.capitalonecareers.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:       1419,
				Alias:    "capital-one",
				Name:     "Capital One",
				Verified: true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "capitalone",
				GoRepositoryCount: 5,
				Verified:          false,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Capital-One-EI_IE3736.11,22.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Capital-One-Reviews-E3736.htm",
				Verified:    true,
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go: []string{
					"https://www.linkedin.com/jobs/view/4058793283/",
				},
				domain.Rust:    nil,
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
			},
			ShortDescription: "American bank holding company",
			Industries: []domain.Industry{
				domain.IndustryFinTech,
			},
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
				domain.Czechia,
			},
		},
		{
			ID:   0,  // system
			Type: "", // system
			Name: "Dynatrace",
			URL:  "https://www.dynatrace.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:       125999,
				Alias:    "dynatrace",
				Name:     "Dynatrace",
				Verified: true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "dynatrace",
				GoRepositoryCount: 11,
				Verified:          true,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Dynatrace-EI_IE309684.11,20.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Dynatrace-Reviews-E309684.htm",
				Verified:    true,
			},
			OttaProfileSlug:   "Dynatrace",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go: []string{
					"https://www.linkedin.com/jobs/view/3951768378/",
				},
				domain.Rust:    nil,
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
			},
			ShortDescription: "Digital and application performance monitoring",
			Industries:       []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
				domain.Czechia,
			},
		},
		{
			ID:   0,  // system
			Type: "", // system
			Name: "emnify",
			URL:  "https://www.emnify.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:       5131340,
				Alias:    "emnify",
				Name:     "emnify",
				Verified: true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "emnify",
				GoRepositoryCount: 7,
				Verified:          true,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-EMnify-EI_IE1356231.11,17.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/EMnify-Reviews-E1356231.htm",
				Verified:    true,
			},
			OttaProfileSlug:   "emnify-1",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go: []string{
					"https://www.linkedin.com/jobs/view/4058102205/",
				},
				domain.Rust:    nil,
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
			},
			ShortDescription: "Cloud-based services for IoT communications",
			Industries:       []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
			},
		},
		{
			ID:   0,  // system
			Type: "", // system
			Name: "Moody's Corporation",
			URL:  "https://www.moodys.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:       165033,
				Alias:    "moodys-corporation",
				Name:     "Moody's Corporation",
				Verified: true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "",
				GoRepositoryCount: 0,
				Verified:          false,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Moody-s-EI_IE11303.11,18.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Moody-s-Reviews-E11303.htm",
				Verified:    true,
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go: nil,
				domain.Rust: []string{
					"https://www.linkedin.com/jobs/view/4060347505/",
				},
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
			},
			ShortDescription: "American business and financial services company",
			Industries: []domain.Industry{
				domain.IndustryFinTech,
			},
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
			},
		},
		{
			ID:   0,  // system
			Type: "", // system
			Name: "Lytx, Inc.",
			URL:  "https://www.lytx.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:       26167,
				Alias:    "lytxinc",
				Name:     "Lytx, Inc.",
				Verified: true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "lytx",
				GoRepositoryCount: 0,
				Verified:          false,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Lytx-EI_IE813859.11,15.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Lytx-Reviews-E813859.htm",
				Verified:    true,
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go: nil,
				domain.Rust: []string{
					"https://www.linkedin.com/jobs/view/3951803632/",
				},
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
			},
			ShortDescription: "Fleet and compliance management solutions",
			Industries:       []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Applied Systems",
			Website: "https://www1.appliedsystems.com/en-us",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:       162482,
				Alias:    "applied-systems",
				Name:     "Applied Systems",
				Verified: true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "",
				Verified: false,
			},
			BlindProfile: domain.BlindProfile{
				Alias:       "",
				Employees:   "",
				Salary:      "",
				Reviews:     "",
				ReviewsRate: "",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "",
				Employees: "",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "",
				ReviewsURL:  "",
				Verified:    false,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "",
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Software Engineer (Golang/React)",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4068444754/",
							Date:             mustDate("2024-11-04"),
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
			ShortDescription:          "",
			DealroomURL:               "",
			CrunchbaseURL:             "",
			PitchbookURL:              "",
			YahooFinanceURL:           "",
			GoogleFinanceURL:          "",
			YCombinatorURL:            "",
			Industries:                []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Yum! Brands",
			Website: "https://www.yum.com/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:       3409,
				Alias:    "yum-brands",
				Name:     "Yum! Brands",
				Verified: true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "",
				Verified: false,
			},
			BlindProfile: domain.BlindProfile{
				Alias:       "",
				Employees:   "",
				Salary:      "",
				Reviews:     "",
				ReviewsRate: "",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "",
				Employees: "",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "",
				ReviewsURL:  "",
				Verified:    false,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "",
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Sr. Software Engineer (GoLang)",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4027419578/",
							Date:             mustDate("2024-11-01"),
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
			ShortDescription:          "",
			DealroomURL:               "",
			CrunchbaseURL:             "",
			PitchbookURL:              "",
			YahooFinanceURL:           "",
			GoogleFinanceURL:          "",
			YCombinatorURL:            "",
			Industries:                []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Nokia",
			Website: "https://www.nokia.com/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:       1070,
				Alias:    "nokia",
				Name:     "Nokia",
				Verified: true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "",
				Verified: false,
			},
			BlindProfile: domain.BlindProfile{
				Alias:       "",
				Employees:   "",
				Salary:      "",
				Reviews:     "",
				ReviewsRate: "",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "",
				Employees: "",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "",
				ReviewsURL:  "",
				Verified:    false,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "",
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Senior Software Engineer — Golang developer",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4062804124/",
							Date:             mustDate("2024-10-31"),
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
			ShortDescription:          "",
			DealroomURL:               "",
			CrunchbaseURL:             "",
			PitchbookURL:              "",
			YahooFinanceURL:           "",
			GoogleFinanceURL:          "",
			YCombinatorURL:            "",
			Industries:                []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{},
		},
		{
			ID:      0,                         // system
			Type:    domain.CompanyTypeStartup, // system
			Name:    "FLYR",
			Website: "https://flyr.com/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:       2906214,
				Alias:    "flyrinc",
				Name:     "FLYR",
				Verified: true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "",
				Verified: false,
			},
			BlindProfile: domain.BlindProfile{
				Alias:       "",
				Employees:   "",
				Salary:      "",
				Reviews:     "",
				ReviewsRate: "",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "",
				Employees: "",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "",
				ReviewsURL:  "",
				Verified:    false,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "",
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Staff Golang Engineer",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/3992302026/",
							Date:             mustDate("2024-10-31"),
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
			ShortDescription:          "",
			DealroomURL:               "",
			CrunchbaseURL:             "",
			PitchbookURL:              "",
			YahooFinanceURL:           "",
			GoogleFinanceURL:          "",
			YCombinatorURL:            "",
			Industries:                []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Venatus",
			Website: "https://www.venatus.com/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:       1116413,
				Alias:    "we-are-venatus",
				Name:     "Venatus",
				Verified: true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "",
				Verified: false,
			},
			BlindProfile: domain.BlindProfile{
				Alias:       "",
				Employees:   "",
				Salary:      "",
				Reviews:     "",
				ReviewsRate: "",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "",
				Employees: "",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "",
				ReviewsURL:  "",
				Verified:    false,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "",
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Senior Golang Engineer",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/3993864259/",
							Date:             mustDate("2024-10-30"),
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
			ShortDescription:          "",
			DealroomURL:               "",
			CrunchbaseURL:             "",
			PitchbookURL:              "",
			YahooFinanceURL:           "",
			GoogleFinanceURL:          "",
			YCombinatorURL:            "",
			Industries:                []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Rocket Lab",
			Website: "https://www.rocketlabusa.com/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:       3308308,
				Alias:    "rocket-lab-limited",
				Name:     "Rocket Lab",
				Verified: true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "",
				Verified: false,
			},
			BlindProfile: domain.BlindProfile{
				Alias:       "",
				Employees:   "",
				Salary:      "",
				Reviews:     "",
				ReviewsRate: "",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "",
				Employees: "",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "",
				ReviewsURL:  "",
				Verified:    false,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "",
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Senior Software Engineer — Operations Software — Golang",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4025814866/",
							Date:             mustDate("2024-10-30"),
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
			ShortDescription:          "",
			DealroomURL:               "",
			CrunchbaseURL:             "",
			PitchbookURL:              "",
			YahooFinanceURL:           "",
			GoogleFinanceURL:          "",
			YCombinatorURL:            "",
			Industries:                []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{},
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
		//		ID:       0,
		//		Alias:    "",
		//		Name:     "",
		//		Verified: false,
		//	},
		//	GitHubProfile: domain.GitHubProfile{
		//		Login:             "",
		//		Verified:          false,
		//	},
		//	BlindProfile: domain.BlindProfile{
		//		Alias:       "",
		//		Employees:   "",
		//		Salary:      "",
		//		Reviews:     "",
		//		ReviewsRate: "",
		//	},
		//	LevelsFyiProfile: domain.LevelsFyiProfile{
		//		Alias:     "",
		//		Employees: "",
		//	},
		//	GlassdoorProfile: domain.GlassdoorProfile{
		//		OverviewURL: "",
		//		ReviewsURL:  "",
		//		Verified:    false,
		//	},
		//	IndeedProfile: domain.IndeedProfile{
		//		Alias: "",
		//	},
		//	OttaProfileSlug:   "",
		//	YouTubeChannelURL: "",
		//	GoMainLanguage:    false,
		//	Languages: domain.Languages{
		//		domain.Go: {
		//			GitHubRepositoriesCount: 0,
		//			Vacancies:               []domain.Vacancy{
		//				{
		//					Title:            "",
		//					ShortDescription: "",
		//					URL:              "",
		//					Date:             time.Time{},
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
		//	ShortDescription:          "",
		//	DealroomURL:               "",
		//	CrunchbaseURL:             "",
		//	PitchbookURL:              "",
		//	YahooFinanceURL:           "",
		//	GoogleFinanceURL:          "",
		//	YCombinatorURL:            "",
		//	Industries:                []domain.Industry{},
		//	HasEmployeesFromCountries: []domain.Country{},
		//},
	}
}
