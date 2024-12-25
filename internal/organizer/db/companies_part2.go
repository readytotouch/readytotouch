package db

import (
	"github.com/readytotouch/readytotouch/internal/domain"
)

func companiesPart2() []domain.CompanyProfile {
	return []domain.CompanyProfile{
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Trusting Social",
			Website: "https://www.trustingsocial.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                3506386,
				Alias:             "trustingsocial",
				Name:              "Trusting Social",
				Followers:         "13K",
				Employees:         "200+",
				AssociatedMembers: "321",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "",
				GoRepositoryCount: 0,
				Verified:          false,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-TrustingSocial-EI_IE741535.11,25.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/TrustingSocial-Reviews-E741535.htm",
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go: []string{
					"https://www.linkedin.com/jobs/view/4042510486/",
				},
				domain.Rust:    nil,
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
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
			ShortDescription: "AI fintech company revolutionizing credit scoring using big data technology",
			Industries: []domain.Industry{
				domain.IndustryFinTech,
			},
			HasEmployeesFromCountries: []domain.Country{},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "McAfee",
			Website: "https://www.mcafee.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                2336,
				Alias:             "mcafee",
				Name:              "McAfee",
				Followers:         "320K",
				Employees:         "1K+",
				AssociatedMembers: "5,531",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "mcafee",
				GoRepositoryCount: 0,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-McAfee-EI_IE2244.11,17.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/McAfee-Reviews-E2244.htm",
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go: []string{
					"https://www.linkedin.com/jobs/view/4041686516/",
				},
				domain.Rust:    nil,
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
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
			ShortDescription: "Antivirus software company",
			Industries: []domain.Industry{
				domain.IndustryCyberSecurity,
			},
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
				domain.Czechia,
			},
		},
		{
			ID:      0,                         // system
			Type:    domain.CompanyTypeStartup, // system
			Name:    "Hadrian",
			Website: "https://www.hadrian.co/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                71668100,
				Alias:             "hadrianspace",
				Name:              "Hadrian",
				Followers:         "13K",
				Employees:         "50+",
				AssociatedMembers: "182",
				Verified:          false,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "",
				GoRepositoryCount: 0,
				Verified:          false,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Hadrian-EI_IE5364733.11,18.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Hadrian-Reviews-E5364733.htm",
			},
			OttaProfileSlug:   "Hadrian-2",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go: []string{
					"https://app.welcometothejungle.com/jobs/63inUEkF",
				},
				domain.Rust:    nil,
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
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
			ShortDescription:          "Building autonomous factories for aerospace components",
			Industries:                []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Scaleway",
			Website: "https://www.scaleway.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                9427185,
				Alias:             "scaleway",
				Name:              "Scaleway",
				Followers:         "33K",
				Employees:         "500+",
				AssociatedMembers: "561",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "scaleway",
				GoRepositoryCount: 23,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Scaleway-EI_IE1861876.11,19.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Scaleway-Reviews-E1861876.htm",
			},
			OttaProfileSlug:   "Scaleway",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go: []string{
					"https://app.welcometothejungle.com/jobs/dFlIR1ow",
				},
				domain.Rust:    nil,
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
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
			ShortDescription:          "Cloud provider",
			Industries:                []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Demandbase",
			Website: "https://www.demandbase.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                89759,
				Alias:             "demandbase",
				Name:              "Demandbase",
				Followers:         "68K",
				Employees:         "500+",
				AssociatedMembers: "981",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "Demandbase",
				GoRepositoryCount: 1,
				Verified:          true,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Demandbase-EI_IE271272.11,21.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Demandbase-Reviews-E271272.htm",
				Verified:    true,
			},
			OttaProfileSlug:   "Demandbase",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go:   nil,
				domain.Rust: nil,
				domain.Zig:  nil,
				domain.Scala: []string{
					"https://www.linkedin.com/jobs/view/4056224841/", // Help develop and maintain the backend that powers our underlying SaaS platform, primarily with Scala microservices
					"https://app.welcometothejungle.com/jobs/sKrHcCx0",
				},
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
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
			ShortDescription:          "Go-To-Market engagement & marketing software",
			Industries:                []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Overwolf",
			Website: "https://www.overwolf.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                519584,
				Alias:             "overwolf.com",
				Name:              "Overwolf",
				Followers:         "21K",
				Employees:         "50+",
				AssociatedMembers: "193",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "overwolf",
				GoRepositoryCount: 0,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Overwolf-EI_IE1963582.11,19.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Overwolf-Reviews-E1963582.htm",
			},
			OttaProfileSlug:   "Overwolf",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go: []string{
					"https://app.welcometothejungle.com/jobs/duQ9FWwK",
				},
				domain.Rust:    nil,
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
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
			ShortDescription: "Gaming app & mod development framework",
			Industries: []domain.Industry{
				domain.IndustryGameDev,
			},
			HasEmployeesFromCountries: []domain.Country{},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Bitfount",
			Website: "https://www.bitfount.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                69183486,
				Alias:             "bitfount",
				Name:              "Bitfount",
				Followers:         "1K",
				Employees:         "10+",
				AssociatedMembers: "25",
				Verified:          false,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "bitfount",
				GoRepositoryCount: 0, // Rust 1
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "",
				ReviewsURL:  "",
				Verified:    false,
			},
			OttaProfileSlug:   "Bitfount",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go: nil,
				domain.Rust: []string{
					"https://www.linkedin.com/jobs/view/4044732400/",
					"https://app.welcometothejungle.com/jobs/AGxChIIS",
				},
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
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
			ShortDescription:          "Federated AI & distributed data science platform",
			Industries:                []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{},
		},
		{
			ID:      0,                         // system
			Type:    domain.CompanyTypeStartup, // system
			Name:    "Hinge",
			Website: "https://hinge.co/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                2603651,
				Alias:             "hinge-app",
				Name:              "Hinge",
				Followers:         "46K",
				Employees:         "50+",
				AssociatedMembers: "431",
				Verified:          false,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "Hinge",
				GoRepositoryCount: 4,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Hinge-EI_IE871901.11,16.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Hinge-Reviews-E871901.htm",
			},
			OttaProfileSlug:   "Hinge",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go: []string{
					"https://www.linkedin.com/jobs/view/4043430809/",
					"https://app.welcometothejungle.com/jobs/V73cck97",
				},
				domain.Rust:    nil,
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
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
			ShortDescription:          "Dating app for meaningful relationships",
			Industries:                []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{},
		},
		{
			ID:      0,                         // system
			Type:    domain.CompanyTypeStartup, // system
			Name:    "Firework",
			Website: "https://firework.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                18700473,
				Alias:             "fireworkhq",
				Name:              "Firework",
				Followers:         "26K",
				Employees:         "200+",
				AssociatedMembers: "371",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "",
				GoRepositoryCount: 0,
				Verified:          false,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Firework-EI_IE2300706.11,19.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Firework-Reviews-E2300706.htm",
			},
			OttaProfileSlug:   "Firework",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go:    nil,
				domain.Rust:  nil,
				domain.Zig:   nil,
				domain.Scala: nil,
				domain.Elixir: []string{
					"https://app.welcometothejungle.com/jobs/srNX_4cO",
				},
				domain.Clojure: nil,
				domain.Haskell: nil,
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
			ShortDescription: "Video-based eCommerce platform",
			Industries: []domain.Industry{
				domain.IndustryECommerce,
			},
			HasEmployeesFromCountries: []domain.Country{},
		},
		{
			ID:      0,                         // system
			Type:    domain.CompanyTypeStartup, // system
			Name:    "Breakroom",
			Website: "https://www.breakroom.cc/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                18458137,
				Alias:             "breakroom",
				Name:              "Breakroom",
				Followers:         "2K",
				Employees:         "10+",
				AssociatedMembers: "20",
				Verified:          false,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "breakroom",
				GoRepositoryCount: 0, // Elixir 13
			},
			BlindProfile:     domain.BlindProfile{},
			LevelsFyiProfile: domain.LevelsFyiProfile{},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Breakroom-EI_IE8002334.11,20.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Breakroom-Reviews-E8002334.htm",
			},
			IndeedProfile:     domain.IndeedProfile{},
			OttaProfileSlug:   "Breakroom",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go:    {},
				domain.Rust:  {},
				domain.Zig:   {},
				domain.Scala: {},
				domain.Elixir: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Software Engineer",
							ShortDescription: "Our platform is built on Elixir, using the Phoenix web framework",
							URL:              "https://www.linkedin.com/jobs/view/4052987616/",
							Date:             mustDate("2024-10-21"),
						},
					},
				},
				domain.Clojure: {},
				domain.Haskell: {},
			},
			ShortDescription:          "The people-powered job comparison site",
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
			Name:    "Ansa",
			Website: "https://www.ansa.dev/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                91158109,
				Alias:             "getansa",
				Name:              "Ansa",
				Followers:         "2K",
				Employees:         "2+",
				AssociatedMembers: "27",
				Verified:          false,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "",
				GoRepositoryCount: 0,
				Verified:          false,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "",
				ReviewsURL:  "",
				Verified:    false,
			},
			OttaProfileSlug:   "Ansa",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go: []string{
					"https://app.welcometothejungle.com/jobs/EjVpa6gl",
				},
				domain.Rust:    nil,
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
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
			ShortDescription:          "Branded customer wallet infrastructure",
			Industries:                []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "1GLOBAL",
			Website: "https://www.1global.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                39711,
				Alias:             "1global",
				Name:              "1GLOBAL",
				Followers:         "61K",
				Employees:         "500+",
				AssociatedMembers: "463",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "",
				GoRepositoryCount: 0,
				Verified:          false,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-1GLOBAL-EI_IE9805631.11,18.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/1GLOBAL-Reviews-E9805631.htm",
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go: []string{
					"https://www.linkedin.com/jobs/view/4048972208/",
				},
				domain.Rust:    nil,
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
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
			ShortDescription: "eSIM",
			Industries:       []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
				domain.Czechia,
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Eco-Movement",
			Website: "https://eco-movement.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                15223273,
				Alias:             "eco-movement-charge-point-data",
				Name:              "Eco-Movement",
				Followers:         "3K",
				Employees:         "10+",
				AssociatedMembers: "50",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "",
				GoRepositoryCount: 0,
				Verified:          false,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Eco-Movement-EI_IE5353422.11,23.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Eco-Movement-Reviews-E5353422.htm",
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go: []string{
					"https://www.linkedin.com/jobs/view/4046511967/",
				},
				domain.Rust:    nil,
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
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
			ShortDescription:          "Charging station data platform",
			Industries:                []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Network Optix",
			Website: "https://www.networkoptix.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                1896041,
				Alias:             "network-optix",
				Name:              "Network Optix",
				Followers:         "22K",
				Employees:         "200+",
				AssociatedMembers: "183",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "networkoptix",
				GoRepositoryCount: 0,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Network-Optix-EI_IE1103945.11,24.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Network-Optix-Reviews-E1103945.htm",
				Verified:    true,
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go: []string{
					"https://www.linkedin.com/jobs/view/4028418131/",
				},
				domain.Rust:    nil,
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
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
			ShortDescription: "Convert video data into practical insights for businesses and industries of all kinds",
			Industries:       []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Cinemo",
			Website: "https://www.cinemo.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                564661,
				Alias:             "cinemo",
				Name:              "Cinemo",
				Followers:         "6K",
				Employees:         "200+",
				AssociatedMembers: "262",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "",
				GoRepositoryCount: 0,
				Verified:          false,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Cinemo-EI_IE2189711.11,17.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Cinemo-Reviews-E2189711.htm",
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Software Engineer / Golang Developer",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4050001754/",
							Date:             mustDate("2024-10-25"),
						},
						{
							Title:            "(Senior) Software Engineer – React & Golang",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/3921611198/",
							Date:             mustDate("2024-12-21"),
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
			ShortDescription: "Make every screen an opportunity",
			Industries:       []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
				domain.Czechia,
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Readdle",
			Website: "https://readdle.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                629551,
				Alias:             "readdle",
				Name:              "Readdle",
				Followers:         "7K",
				Employees:         "200+",
				AssociatedMembers: "280",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "readdle",
				GoRepositoryCount: 5,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Readdle-EI_IE613771.11,18.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Readdle-Reviews-E613771.htm",
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go: []string{
					"https://www.linkedin.com/jobs/view/3971228819/",
				},
				domain.Rust:    nil,
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
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
			ShortDescription: "Creating apps",
			Industries:       []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Ipsos",
			Website: "https://www.ipsos.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                4318,
				Alias:             "ipsos",
				Name:              "Ipsos",
				Followers:         "398K",
				Employees:         "10K+",
				AssociatedMembers: "17,733",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "",
				GoRepositoryCount: 0,
				Verified:          false,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Ipsos-EI_IE13063.11,16.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Ipsos-Reviews-E13063.htm",
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go: []string{
					"https://www.linkedin.com/jobs/view/4053349624/",
				},
				domain.Rust:    nil,
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
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
			ShortDescription: "Market research company",
			Industries:       []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
				domain.Czechia,
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Bedrock Streaming",
			Website: "https://bedrockstreaming.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                64853943,
				Alias:             "bedrock-streaming",
				Name:              "Bedrock Streaming",
				Followers:         "30K",
				Employees:         "200+",
				AssociatedMembers: "315",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "bedrockstreaming",
				Verified: true,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Bedrock-Streaming-EI_IE5955934.11,28.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Bedrock-Streaming-Reviews-E5955934.htm",
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 7,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Software Developer — Backend GO / PHP",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4049210489/",
							Date:             mustDate("2024-10-25"),
						},
						{
							Title:            "Software Developer —  Backend GO",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4079483983/",
							Date:             mustDate("2024-12-11"),
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
			ShortDescription:          "Bedrock creates and operates full-scope streaming platforms",
			Industries:                []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Whalebone",
			Website: "https://www.whalebone.io/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                10961729,
				Alias:             "whaleboneio",
				Name:              "Whalebone",
				Followers:         "9K",
				Employees:         "50+",
				AssociatedMembers: "118",
				Verified:          false,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "whalebone",
				GoRepositoryCount: 5,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Whalebone-EI_IE4889910.11,20.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Whalebone-Reviews-E4889910.htm",
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go: []string{
					"https://www.linkedin.com/jobs/view/4007694502/",
				},
				domain.Rust:    nil,
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
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
			Industries: []domain.Industry{
				domain.IndustryCyberSecurity,
			},
			HasEmployeesFromCountries: []domain.Country{
				domain.Czechia,
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "OneSignal",
			Website: "https://onesignal.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                6424376,
				Alias:             "onesignal",
				Name:              "OneSignal",
				Followers:         "9K",
				Employees:         "50+",
				AssociatedMembers: "164",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "OneSignal",
				GoRepositoryCount: 5, // Rust 32
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-OneSignal-EI_IE1952551.11,20.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/OneSignal-Reviews-E1952551.htm",
			},
			OttaProfileSlug:   "OneSignal",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go: []string{
					"https://app.welcometothejungle.com/jobs/wxzE8h5B",
				},
				domain.Rust: []string{
					"https://app.welcometothejungle.com/jobs/wxzE8h5B",
				},
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
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
			ShortDescription:          "Customer messaging and engagement platform",
			Industries:                []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{},
		},
		{
			ID:      0,                         // system
			Type:    domain.CompanyTypeStartup, // system
			Name:    "Metabase",
			Website: "https://www.metabase.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                6460313,
				Alias:             "metabase",
				Name:              "Metabase",
				Followers:         "15K",
				Employees:         "50+",
				AssociatedMembers: "94",
				Verified:          false,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "metabase",
				GoRepositoryCount: 0, // Clojure 30
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Metabase-CA-EI_IE6095700.11,22.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Metabase-CA-Reviews-E6095700.htm",
			},
			OttaProfileSlug:   "Metabase",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go:     nil,
				domain.Rust:   nil,
				domain.Zig:    nil,
				domain.Scala:  nil,
				domain.Elixir: nil,
				domain.Clojure: []string{
					"https://app.welcometothejungle.com/jobs/K21IaTJ2",
				},
				domain.Haskell: nil,
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
			ShortDescription: "Open-source Business Intelligence tool",
			Industries:       []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Anjuna Security",
			Website: "https://www.anjuna.io/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                18438300,
				Alias:             "anjuna-security",
				Name:              "Anjuna Security",
				Followers:         "4K",
				Employees:         "10+",
				AssociatedMembers: "47",
				Verified:          false,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "anjuna-security",
				GoRepositoryCount: 2,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Anjuna-Security-EI_IE3150191.11,26.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Anjuna-Security-Reviews-E3150191.htm",
			},
			OttaProfileSlug:   "Anjuna",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go: []string{
					"https://www.linkedin.com/jobs/view/4052028718/",
					"https://app.welcometothejungle.com/jobs/yxGdxMhb",
				},
				domain.Rust: []string{
					"https://www.linkedin.com/jobs/view/4052028718/",
					"https://app.welcometothejungle.com/jobs/yxGdxMhb",
				},
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
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
			ShortDescription: "Public-cloud network security",
			Industries: []domain.Industry{
				domain.IndustryCyberSecurity,
			},
			HasEmployeesFromCountries: []domain.Country{},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Aqua Security",
			Website: "https://www.aquasec.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                10034420,
				Alias:             "aquasecteam",
				Name:              "Aqua Security",
				Followers:         "63K",
				Employees:         "500+",
				AssociatedMembers: "639",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "aquasecurity",
				GoRepositoryCount: 101,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Aqua-Security-Software-EI_IE1785939.11,33.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Aqua-Security-Software-Reviews-E1785939.htm",
			},
			OttaProfileSlug:   "Aqua-Security",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go: []string{
					"https://www.linkedin.com/jobs/view/4045761623/",
				},
				domain.Rust:    nil,
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
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
			ShortDescription: "Cloud native security platform",
			Industries: []domain.Industry{
				domain.IndustryCyberSecurity,
			},
			HasEmployeesFromCountries: []domain.Country{},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Verve",
			Website: "https://verve.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                108605,
				Alias:             "verve-ad-solutions",
				Name:              "Verve",
				Followers:         "23K",
				Employees:         "200+",
				AssociatedMembers: "470",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "",
				GoRepositoryCount: 0,
				Verified:          false,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Verve-EI_IE4432646.11,16.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Verve-Reviews-E4432646.htm",
			},
			OttaProfileSlug:   "Verve",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go: []string{
					"https://www.linkedin.com/jobs/view/4015809270/",
				},
				domain.Rust:    nil,
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
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
			ShortDescription: "Real time omnichannel ad platform",
			Industries: []domain.Industry{
				domain.IndustryAdTech,
			},
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
				domain.Czechia,
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Wargaming",
			Website: "https://wargaming.com/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                127309,
				Alias:             "wargaming-net",
				Name:              "Wargaming",
				Followers:         "131K",
				Employees:         "1K+",
				AssociatedMembers: "3,921",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "wgnet",
				GoRepositoryCount: 3,
			},
			BlindProfile:     domain.BlindProfile{},
			LevelsFyiProfile: domain.LevelsFyiProfile{},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Wargaming-EI_IE381713.11,20.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Wargaming-Reviews-E381713.htm",
			},
			IndeedProfile:     domain.IndeedProfile{},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Software Engineer (Python/Go)",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4077489876/",
							Date:             mustDate("2024-11-17"),
						},
					},
				},
				domain.Rust: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Rust Software Engineer (Platform)",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4077490300/",
							Date:             mustDate("2024-11-17"),
						},
					},
				},
				domain.Zig:     {},
				domain.Scala:   {},
				domain.Elixir:  {},
				domain.Clojure: {},
				domain.Haskell: {},
			},
			ShortDescription: "Game developer and publisher",
			DealroomURL:      "",
			CrunchbaseURL:    "",
			PitchbookURL:     "",
			YahooFinanceURL:  "",
			GoogleFinanceURL: "",
			YCombinatorURL:   "",
			Industries: []domain.Industry{
				domain.IndustryGameDev,
			},
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
				domain.Czechia,
			},
		},
		{
			ID:      0,                         // system
			Type:    domain.CompanyTypeStartup, // system
			Name:    "Napier AI",
			Website: "https://www.napier.ai/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                15197985,
				Alias:             "napier",
				Name:              "Napier AI",
				Followers:         "17K",
				Employees:         "200+",
				AssociatedMembers: "229",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "napier-ai",
				Verified: true,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Napier-Technologies-EI_IE4370088.11,30.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Napier-Technologies-Reviews-E4370088.htm",
			},
			OttaProfileSlug:   "Napier",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Back-end Engineer (Go)",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4044948110/",
							Date:             mustDate("2024-12-18"),
						},
						{
							Title:            "Back-end Engineer (Go)",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4106841758/",
							Date:             mustDate("2024-12-24"),
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
			ShortDescription: "AI-driven AML solutions transform financial crime compliance from legal obligation to competitive edge",
			Industries:       []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Kaseya",
			Website: "https://www.kaseya.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                21377,
				Alias:             "kaseya",
				Name:              "Kaseya",
				Followers:         "157K",
				Employees:         "5K+",
				AssociatedMembers: "4,836",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "kaseya",
				GoRepositoryCount: 0,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Kaseya-EI_IE262966.11,17.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Kaseya-Reviews-E262966.htm",
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go: []string{
					"https://www.linkedin.com/jobs/view/3908613207/",
				},
				domain.Rust:    nil,
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
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
			ShortDescription: "Provider of AI-powered cybersecurity and IT management software",
			Industries: []domain.Industry{
				domain.IndustryCyberSecurity,
			},
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
			},
		},
		{
			ID:      0,                         // system
			Type:    domain.CompanyTypeStartup, // system
			Name:    "Omio",
			Website: "https://www.omio.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                2754440,
				Alias:             "omio",
				Name:              "Omio",
				Followers:         "50K",
				Employees:         "200+",
				AssociatedMembers: "361",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "omio-labs",
				GoRepositoryCount: 2,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Omio-EI_IE2855962.11,15.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Omio-Reviews-E2855962.htm",
			},
			OttaProfileSlug:   "Omio",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go: []string{
					"https://www.linkedin.com/jobs/view/4025460474/",
				},
				domain.Rust:    nil,
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
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
			ShortDescription: "Comparison & booking website for public transport",
			Industries:       []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
				domain.Czechia,
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Liebherr",
			Website: "https://www.liebherr.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                11057,
				Alias:             "liebherr",
				Name:              "Liebherr Group",
				Followers:         "398K",
				Employees:         "10K+",
				AssociatedMembers: "10,005",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "",
				GoRepositoryCount: 0,
				Verified:          false,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Liebherr-Group-EI_IE787369.11,25.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Liebherr-Group-Reviews-E787369.htm",
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go: []string{
					"https://www.linkedin.com/jobs/view/3821037210/",
				},
				domain.Rust:    nil,
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
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
			ShortDescription: "Construction equipment manufacturer",
			Industries:       []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
				domain.Czechia,
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Fiserv",
			Website: "https://www.fiserv.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                3364,
				Alias:             "fiserv",
				Name:              "Fiserv",
				Followers:         "755K",
				Employees:         "10K+",
				AssociatedMembers: "35,296",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "fiserv",
				GoRepositoryCount: 0,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Fiserv-EI_IE1384.11,17.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Fiserv-Reviews-E1384.htm",
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go: []string{
					"https://www.linkedin.com/jobs/view/4026685486/",
				},
				domain.Rust:    nil,
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
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
			Industries:       []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
				domain.Czechia,
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Gett",
			Website: "https://www.gett.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                1514929,
				Alias:             "gettaxi",
				Name:              "Gett",
				Followers:         "49K",
				Employees:         "500+",
				AssociatedMembers: "862",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "GettEngineering",
				GoRepositoryCount: 8,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Gett-EI_IE810731.11,15.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Gett-Reviews-E810731.htm",
			},
			OttaProfileSlug:   "Gett",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go: []string{
					"https://www.linkedin.com/jobs/view/3819936157/",
				},
				domain.Rust:    nil,
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
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
			ShortDescription: "SaaS platform for mobility providers",
			Industries:       []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Stairwell",
			Website: "https://stairwell.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                68606343,
				Alias:             "stairwell-inc",
				Name:              "Stairwell",
				Followers:         "4K",
				Employees:         "10+",
				AssociatedMembers: "65",
				Verified:          false,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "stairwell-inc",
				GoRepositoryCount: 0,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "",
				ReviewsURL:  "",
				Verified:    false,
			},
			OttaProfileSlug:   "Stairwell",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go: []string{
					"https://app.welcometothejungle.com/jobs/hnl2J9ws",
				},
				domain.Rust:    nil,
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
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
			ShortDescription: "Threat detection & cybersecurity platform",
			Industries: []domain.Industry{
				domain.IndustryCyberSecurity,
			},
			HasEmployeesFromCountries: []domain.Country{},
		},
		{
			ID:      0,                         // system
			Type:    domain.CompanyTypeStartup, // system
			Name:    "Solo.io",
			Website: "https://www.solo.io/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                11074869,
				Alias:             "solo.io",
				Name:              "Solo.io",
				Followers:         "14K",
				Employees:         "50+",
				AssociatedMembers: "248",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "solo-io",
				GoRepositoryCount: 82,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-solo-io-EI_IE5382785.11,18.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/solo-io-Reviews-E5382785.htm",
			},
			OttaProfileSlug:   "Solo-io",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go: []string{
					"https://app.welcometothejungle.com/jobs/vuvd0QAu",
				},
				domain.Rust:    nil,
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
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
			ShortDescription: "Delivering API infrastructure software",
			Industries: []domain.Industry{
				domain.IndustryCyberSecurity,
				domain.IndustryDevOps,
			},
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Splash Damage",
			Website: "https://www.splashdamage.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                47756,
				Alias:             "splash-damage",
				Name:              "Splash Damage",
				Followers:         "47K",
				Employees:         "200+",
				AssociatedMembers: "329",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "splash-damage",
				GoRepositoryCount: 0,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Splash-Damage-EI_IE470252.11,24.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Splash-Damage-Reviews-E470252.htm",
			},
			OttaProfileSlug:   "SplashDamage",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go: []string{
					"https://app.welcometothejungle.com/jobs/A7daqq7I",
				},
				domain.Rust:    nil,
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
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
			ShortDescription: "Game development studio",
			Industries: []domain.Industry{
				domain.IndustryGameDev,
			},
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
			},
		},
		{
			ID:      0,                         // system
			Type:    domain.CompanyTypeStartup, // system
			Name:    "Ditto",
			Website: "https://ditto.live/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                18702497,
				Alias:             "dittolive",
				Name:              "Ditto",
				Followers:         "4K",
				Employees:         "50+",
				AssociatedMembers: "143",
				Verified:          false,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "getditto",
				GoRepositoryCount: 6, // Rust 30
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "",
				ReviewsURL:  "",
				Verified:    false,
			},
			OttaProfileSlug:   "ditto",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go: nil,
				domain.Rust: []string{
					"https://www.linkedin.com/jobs/view/4050904655/",
					"https://app.welcometothejungle.com/jobs/U3QtZINa",
				},
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
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
			ShortDescription:          "Data sync platform",
			Industries:                []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Pinecone",
			Website: "https://www.pinecone.io/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                20299330,
				Alias:             "pinecone-io",
				Name:              "Pinecone",
				Followers:         "63K",
				Employees:         "50+",
				AssociatedMembers: "178",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "pinecone-io",
				GoRepositoryCount: 6, // Rust 9
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Pinecone-Systems-EI_IE6085804.11,27.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Pinecone-Systems-Reviews-E6085804.htm",
			},
			OttaProfileSlug:   "Pinecone",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go: []string{
					"https://app.welcometothejungle.com/jobs/ZudU5aZg",
					"https://www.linkedin.com/jobs/view/4002028518/",
				},
				domain.Rust: []string{
					"https://app.welcometothejungle.com/jobs/ZudU5aZg",
					"https://www.linkedin.com/jobs/view/4002028518/",
				},
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
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
			ShortDescription: "Vector database",
			Industries: []domain.Industry{
				domain.IndustryDevOps,
			},
			HasEmployeesFromCountries: []domain.Country{
				domain.Czechia,
			},
		},
		{
			ID:      0,                         // system
			Type:    domain.CompanyTypeStartup, // system
			Name:    "Beyond Identity",
			Website: "https://www.beyondidentity.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                64665737,
				Alias:             "beyond-identity-inc",
				Name:              "Beyond Identity",
				Followers:         "14K",
				Employees:         "50+",
				AssociatedMembers: "140",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "gobeyondidentity",
				Verified: true,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Beyond-Identity-EI_IE3403008.11,26.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Beyond-Identity-Reviews-E3403008.htm",
			},
			OttaProfileSlug:   "Beyond-Identity",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go: {},
				domain.Rust: {
					GitHubRepositoriesCount: 2,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Senior Software Engineer",
							ShortDescription: "We develop most of our server-side code in Rust and it's a plus if you already love Rust, or want a chance to become a Rust expert",
							URL:              "https://www.linkedin.com/jobs/view/4068292201/",
							Date:             mustDate("2024-11-25"),
						},
					},
				},
				domain.Zig:     {},
				domain.Scala:   {},
				domain.Elixir:  {},
				domain.Clojure: {},
				domain.Haskell: {},
			},
			ShortDescription: "Passwordless identity management solutions",
			Industries: []domain.Industry{
				domain.IndustryCyberSecurity,
			},
			HasEmployeesFromCountries: []domain.Country{},
		},
		{
			ID:      0,                         // system
			Type:    domain.CompanyTypeStartup, // system
			Name:    "Nightfall AI",
			Website: "https://www.nightfall.ai/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                19031754,
				Alias:             "nightfall-ai",
				Name:              "Nightfall AI",
				Followers:         "16K",
				Employees:         "50+",
				AssociatedMembers: "110",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "nightfallai",
				GoRepositoryCount: 5,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Nightfall-AI-EI_IE3097406.11,23.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Nightfall-AI-Reviews-E3097406.htm",
			},
			OttaProfileSlug:   "Nightfall-AI",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Staff Software Engineer",
							ShortDescription: "Experience in programming with Go, C++, Java or a related language",
							URL:              "https://www.linkedin.com/jobs/view/4068295195/",
							Date:             mustDate("2024-11-25"),
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
			ShortDescription: "Cloud-native data protection platform",
			Industries: []domain.Industry{
				domain.IndustryCyberSecurity,
			},
			HasEmployeesFromCountries: []domain.Country{},
		},
		{
			ID:      0,                         // system
			Type:    domain.CompanyTypeStartup, // system
			Name:    "Cyberhaven",
			Website: "https://www.cyberhaven.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                10639445,
				Alias:             "cyberhaven",
				Name:              "Cyberhaven",
				Followers:         "10K",
				Employees:         "50+",
				AssociatedMembers: "181",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "CyberhavenInc",
				Verified: false,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Cyberhaven-EI_IE2985068.11,21.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Cyberhaven-Reviews-E2985068.htm",
			},
			OttaProfileSlug:   "Cyberhaven",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 13,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Backend Developer",
							ShortDescription: "You get to work with a modern and constantly evolving microservices-based software stack which includes Go, Kubernetes, BigTable, BigQuery, Spanner, Redis, Docker, etcd",
							URL:              "https://www.linkedin.com/jobs/view/4056670384/",
							Date:             mustDate("2024-10-25"),
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
			ShortDescription: "Data detection & response platform",
			Industries: []domain.Industry{
				domain.IndustryCyberSecurity,
			},
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
			},
		},
		{
			ID:      0,                         // system
			Type:    domain.CompanyTypeStartup, // system
			Name:    "Fonoa",
			Website: "https://www.fonoa.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                30129363,
				Alias:             "fonoa",
				Name:              "Fonoa",
				Followers:         "19K",
				Employees:         "50+",
				AssociatedMembers: "140",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "Fonoa-Tech",
				GoRepositoryCount: 0,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Fonoa-EI_IE4125781.11,16.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Fonoa-Reviews-E4125781.htm",
			},
			OttaProfileSlug:   "Fonoa",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Backend Engineer",
							ShortDescription: "For the backend, we use Node and Go",
							URL:              "https://www.linkedin.com/jobs/view/3960607071/",
							Date:             mustDate("2024-11-25"),
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
			ShortDescription: "Tax automation for the internet economy",
			Industries: []domain.Industry{
				domain.IndustryFinTech,
			},
			HasEmployeesFromCountries: []domain.Country{},
		},
		{
			ID:      0,                         // system
			Type:    domain.CompanyTypeStartup, // system
			Name:    "Payrails",
			Website: "https://www.payrails.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                75580859,
				Alias:             "payrails",
				Name:              "Payrails",
				Followers:         "7K",
				Employees:         "50+",
				AssociatedMembers: "68",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "",
				GoRepositoryCount: 0,
				Verified:          false,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "",
				ReviewsURL:  "",
				Verified:    false,
			},
			OttaProfileSlug:   "Payrails",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Senior Software Engineer",
							ShortDescription: "You show vast experience in developing software in large-scale production systems, ideally using Golang",
							URL:              "https://www.linkedin.com/jobs/view/4026258568/",
							Date:             mustDate("2024-09-25"),
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
			ShortDescription: "Operating system for payments and financial services",
			Industries: []domain.Industry{
				domain.IndustryFinTech,
			},
			HasEmployeesFromCountries: []domain.Country{
				domain.Czechia,
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Teleport",
			Website: "https://goteleport.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                7941233,
				Alias:             "go-teleport",
				Name:              "Teleport",
				Followers:         "20K",
				Employees:         "50+",
				AssociatedMembers: "193",
				Verified:          false,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "gravitational",
				GoRepositoryCount: 116,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Teleport-EI_IE1967263.11,19.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Teleport-Reviews-E1967263.htm",
			},
			OttaProfileSlug:   "Teleport",
			YouTubeChannelURL: "",
			GoMainLanguage:    true,
			Vacancies: domain.Vacancies{
				domain.Go: []string{
					"https://www.linkedin.com/jobs/view/4000249330/",
					"https://app.welcometothejungle.com/jobs/vhRdlAui",
				},
				domain.Rust:    nil,
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
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
			ShortDescription: "Identity-native infrastructure access",
			Industries: []domain.Industry{
				domain.IndustryCyberSecurity,
				domain.IndustryDevOps,
			},
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
			},
		},
		{
			ID:      0,                         // system
			Type:    domain.CompanyTypeStartup, // system
			Name:    "Bigblue",
			Website: "https://www.bigblue.co/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                18776044,
				Alias:             "bigblue-co",
				Name:              "Bigblue",
				Followers:         "11K",
				Employees:         "50+",
				AssociatedMembers: "198",
				Verified:          false,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "",
				GoRepositoryCount: 0,
				Verified:          false,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Bigblue-EI_IE3379658.11,18.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Bigblue-Reviews-E3379658.htm",
			},
			OttaProfileSlug:   "Bigblue",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go: []string{
					"https://app.welcometothejungle.com/jobs/UU1oL2xl",
				},
				domain.Rust:    nil,
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
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
			ShortDescription: "Developer of eCommerce logistics tools",
			Industries: []domain.Industry{
				domain.IndustryECommerce,
			},
			HasEmployeesFromCountries: []domain.Country{},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Taxfix",
			Website: "https://taxfix.de/en/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                10791106,
				Alias:             "taxfix",
				Name:              "Taxfix",
				Followers:         "25K",
				Employees:         "200+",
				AssociatedMembers: "339",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "taxfix",
				GoRepositoryCount: 0,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Taxfix-EI_IE1301207.11,17.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Taxfix-Reviews-E1301207.htm",
			},
			OttaProfileSlug:   "Taxfix",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go: []string{
					"https://app.welcometothejungle.com/jobs/whJEvleb",
				},
				domain.Rust:    nil,
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
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
			ShortDescription: "Tax filing app",
			Industries: []domain.Industry{
				domain.IndustryFinTech,
			},
			HasEmployeesFromCountries: []domain.Country{},
		},
		{
			ID:      0,                         // system
			Type:    domain.CompanyTypeStartup, // system
			Name:    "Clerk",
			Website: "https://clerk.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                69336355,
				Alias:             "clerkinc",
				Name:              "Clerk.com",
				Followers:         "5K",
				Employees:         "50+",
				AssociatedMembers: "147",
				Verified:          false,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "clerk",
				GoRepositoryCount: 8,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "",
				ReviewsURL:  "",
				Verified:    false,
			},
			OttaProfileSlug:   "Clerk",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go: []string{
					"",
				},
				domain.Rust:    nil,
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
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
			ShortDescription:          "Identity tool for React applications",
			Industries:                []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "ClearScore",
			Website: "https://clearscore.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                9320086,
				Alias:             "clearscore",
				Name:              "ClearScore",
				Followers:         "46K",
				Employees:         "200+",
				AssociatedMembers: "380",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "ClearScore",
				GoRepositoryCount: 2, // Scala 2
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-ClearScore-EI_IE1046600.11,21.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/ClearScore-Reviews-E1046600.htm",
			},
			OttaProfileSlug:   "ClearScore",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go:   nil,
				domain.Rust: nil,
				domain.Zig:  nil,
				domain.Scala: []string{
					"https://app.welcometothejungle.com/jobs/RmSslgyY",
				},
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
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
			ShortDescription: "Financial wellbeing marketplace",
			Industries: []domain.Industry{
				domain.IndustryFinTech,
			},
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
			},
		},
		{
			ID:      0,                         // system
			Type:    domain.CompanyTypeStartup, // system
			Name:    "OneSchema",
			Website: "https://www.oneschema.co/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                74704820,
				Alias:             "oneschema",
				Name:              "OneSchema",
				Followers:         "3K",
				Employees:         "10+",
				AssociatedMembers: "21",
				Verified:          false,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "oneschema",
				GoRepositoryCount: 0,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-OneSchema-EI_IE7681431.11,20.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/OneSchema-Reviews-E7681431.htm",
			},
			OttaProfileSlug:   "OneSchema",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go: nil,
				domain.Rust: []string{
					"https://app.welcometothejungle.com/jobs/YvpdZmyz",
				},
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
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
			ShortDescription:          "Cloud-based data importing and validation platform",
			Industries:                []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{},
		},
		{
			ID:      0,                         // system
			Type:    domain.CompanyTypeStartup, // system
			Name:    "StrongDM",
			Website: "https://www.strongdm.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                9446266,
				Alias:             "strongdm",
				Name:              "StrongDM",
				Followers:         "10K",
				Employees:         "50+",
				AssociatedMembers: "145",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "strongdm",
				GoRepositoryCount: 15,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-strongDM-EI_IE4514684.11,19.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/strongDM-Reviews-E4514684.htm",
			},
			OttaProfileSlug:   "strongDM",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go: []string{
					"https://www.linkedin.com/jobs/view/3978025966/",
					"https://app.welcometothejungle.com/jobs/lys1kK5P",
				},
				domain.Rust:    nil,
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
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
			ShortDescription: "Manages auditable database access",
			Industries: []domain.Industry{
				domain.IndustryCyberSecurity,
			},
			HasEmployeesFromCountries: []domain.Country{},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "JMA Wireless",
			Website: "https://jmawireless.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                3105986,
				Alias:             "jmawireless",
				Name:              "JMA Wireless",
				Followers:         "28K",
				Employees:         "500+",
				AssociatedMembers: "730",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "",
				GoRepositoryCount: 0,
				Verified:          false,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-JMA-Wireless-EI_IE875877.11,23.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/JMA-Wireless-Reviews-E875877.htm",
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go: []string{
					"https://www.linkedin.com/jobs/view/4048251309/",
				},
				domain.Rust:    nil,
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
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
			ShortDescription:          "JMA Wireless designs and delivers cutting-edge wireless technology",
			Industries:                []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "ON2IT Cybersecurity",
			Website: "https://on2it.net/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                489607,
				Alias:             "on2it-cybersecurity",
				Name:              "ON2IT Cybersecurity",
				Followers:         "5K",
				Employees:         "50+",
				AssociatedMembers: "96",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "on2itsecurity",
				GoRepositoryCount: 6,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-ON2IT-EI_IE1075463.11,16.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/ON2IT-Reviews-E1075463.htm",
				Verified:    true,
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go: []string{
					"https://www.linkedin.com/jobs/view/3513136315/", // "We mainly use PHP and Go"
				},
				domain.Rust:    nil,
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
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
			ShortDescription: "CyberSecurity service provider",
			Industries: []domain.Industry{
				domain.IndustryCyberSecurity,
			},
			HasEmployeesFromCountries: []domain.Country{},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Cycloid",
			Website: "https://www.cycloid.io/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                10071522,
				Alias:             "cycloid",
				Name:              "Cycloid - Sustainable Platform Engineering",
				Followers:         "11K",
				Employees:         "10+",
				AssociatedMembers: "25",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "cycloidio",
				GoRepositoryCount: 40,
				Verified:          false,
			},
			BlindProfile:     domain.BlindProfile{},
			LevelsFyiProfile: domain.LevelsFyiProfile{},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Cycloid-io-EI_IE1562422.11,21.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Cycloid-io-Reviews-E1562422.htm",
				Verified:    false,
			},
			IndeedProfile:     domain.IndeedProfile{},
			OttaProfileSlug:   "Cycloid",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 40,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Senior Backend Go Developer",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4075626268/",
							Date:             mustDate("2024-12-12"), // mustDate("2024-11-17"),
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
			ShortDescription: "Hybrid cloud DevOps platform",
			DealroomURL:      "",
			CrunchbaseURL:    "",
			PitchbookURL:     "",
			YahooFinanceURL:  "",
			GoogleFinanceURL: "",
			YCombinatorURL:   "",
			Industries: []domain.Industry{
				domain.IndustryDevOps,
			},
			HasEmployeesFromCountries: []domain.Country{},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Circuit",
			Website: "https://circuit.ai/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                36014538,
				Alias:             "circuitknowledge",
				Name:              "Circuit",
				Followers:         "5K",
				Employees:         "10+",
				AssociatedMembers: "57",
				Verified:          false,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "",
				GoRepositoryCount: 0,
				Verified:          false,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Circuit-TX-EI_IE10004565.11,21.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Circuit-TX-Reviews-E10004565.htm",
				Verified:    true,
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go: []string{
					"https://www.linkedin.com/jobs/view/3984171089/", // 5+ years of experience with Go
				},
				domain.Rust:    nil,
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
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
			ShortDescription: "AI-powered platform transforms information into actionable knowledge",
			Industries: []domain.Industry{
				domain.IndustryEdTech,
			},
			HasEmployeesFromCountries: []domain.Country{},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Aylo",
			Website: "https://www.aylo.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                98389423,
				Alias:             "ayloservices",
				Name:              "Aylo",
				Followers:         "22K",
				Employees:         "1K+",
				AssociatedMembers: "1,520",
				Verified:          false,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "MindGeekOSS",
				GoRepositoryCount: 1,
				Verified:          false,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Aylo-EI_IE9049055.11,15.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Aylo-Reviews-E9049055.htm",
				Verified:    true,
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go: []string{
					"https://www.linkedin.com/jobs/view/3968183921/", // Familiarity of Golang
				},
				domain.Rust:    nil,
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
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
			ShortDescription: "Adult content platforms",
			Industries: []domain.Industry{
				domain.IndustryEntertainment,
			},
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
				domain.Czechia,
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "FactSet",
			Website: "https://www.factset.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                163755,
				Alias:             "factset",
				Name:              "FactSet",
				Followers:         "243K",
				Employees:         "10K+",
				AssociatedMembers: "13,540",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "factset",
				GoRepositoryCount: 1,
				Verified:          true,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-FactSet-EI_IE6066.11,18.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/FactSet-Reviews-E6066.htm",
				Verified:    true,
			},
			OttaProfileSlug:   "Factset",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go: []string{
					"https://www.linkedin.com/jobs/view/4018413250/", // Proficiency with Go and protobuf for minimum of 1+ years
				},
				domain.Rust:    nil,
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
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
			ShortDescription: "Providing instant access to financial data and analytics that investors use to make decisions",
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
			Name:    "Voodoo",
			Website: "https://voodoo.io/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                5353630,
				Alias:             "voodoo.io",
				Name:              "Voodoo",
				Followers:         "77K",
				Employees:         "500+",
				AssociatedMembers: "694",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "VoodooTeam",
				Verified: false,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Voodoo-EI_IE1889746.11,17.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Voodoo-Reviews-E1889746.htm",
				Verified:    true,
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 3,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Senior Backend Engineer — Golang",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/3966903286/",
							Date:             mustDate("2024-12-04"),
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
			ShortDescription: "Company that creates mobile games and apps",
			Industries: []domain.Industry{
				domain.IndustryGameDev,
			},
			HasEmployeesFromCountries: []domain.Country{},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "home24",
			Website: "https://www.home24.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                2653180,
				Alias:             "home24",
				Name:              "home24",
				Followers:         "11K",
				Employees:         "1K+",
				AssociatedMembers: "447",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "home24",
				GoRepositoryCount: 1,
				Verified:          false,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-home24-SE-EI_IE3166011.11,20.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/home24-SE-Reviews-E3166011.htm",
				Verified:    false,
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go: []string{
					"https://www.linkedin.com/jobs/view/3872092645/", // Our tech stack includes Golang
				},
				domain.Rust:    nil,
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
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
			ShortDescription: "Home & living e-commerce platform",
			Industries: []domain.Industry{
				domain.IndustryECommerce,
			},
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Australian Broadcasting Corporation",
			Website: "https://www.abc.net.au/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                2242,
				Alias:             "australian-broadcasting-corporation",
				Name:              "Australian Broadcasting Corporation",
				Followers:         "342K",
				Employees:         "1K+",
				AssociatedMembers: "8,087",
				Verified:          false,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "abcnews",
				GoRepositoryCount: 0,
				Verified:          false,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Australian-Broadcasting-EI_IE416033.11,34.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Australian-Broadcasting-Reviews-E416033.htm",
				Verified:    false,
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go: []string{
					"https://www.linkedin.com/jobs/view/4050919710/", // Outstanding enterprise-level experience with Golang and extensive background in server-side development
				},
				domain.Rust:    nil,
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
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
			ShortDescription: "Independent public broadcaster",
			Industries:       []domain.Industry{
				// Media
			},
			HasEmployeesFromCountries: []domain.Country{},
		},
		{
			ID:      0,                         // system
			Type:    domain.CompanyTypeStartup, // system
			Name:    "Nitro",
			Website: "https://www.gonitro.com/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                76281,
				Alias:             "go-nitro",
				Name:              "Nitro Software",
				Followers:         "16K",
				Employees:         "200+",
				AssociatedMembers: "325",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "nitro",
				GoRepositoryCount: 21, // Scala 5
				Verified:          true,
			},
			BlindProfile:     domain.BlindProfile{},
			LevelsFyiProfile: domain.LevelsFyiProfile{},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Nitro-EI_IE402100.11,16.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Nitro-Reviews-E402100.htm",
				Verified:    true,
			},
			IndeedProfile:     domain.IndeedProfile{},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go:   {},
				domain.Rust: {},
				domain.Zig:  {},
				domain.Scala: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Scala Developer",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4042171556/",
							Date:             mustDate("2024-11-08"),
						},
					},
				},
				domain.Elixir:  {},
				domain.Clojure: {},
				domain.Haskell: {},
			},
			ShortDescription: "Document productivity SaaS company",
			DealroomURL:      "",
			CrunchbaseURL:    "",
			PitchbookURL:     "",
			YahooFinanceURL:  "",
			GoogleFinanceURL: "",
			YCombinatorURL:   "",
			Industries:       []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{
				domain.Czechia,
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Appodeal",
			Website: "https://appodeal.com/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                19139082,
				Alias:             "appodeal",
				Name:              "Appodeal",
				Followers:         "7K",
				Employees:         "200+",
				AssociatedMembers: "199",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "appodeal",
				GoRepositoryCount: 2, // Scala 2
				Verified:          true,
			},
			BlindProfile:     domain.BlindProfile{},
			LevelsFyiProfile: domain.LevelsFyiProfile{},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Appodeal-EI_IE4606703.11,19.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Appodeal-Reviews-E4606703.htm",
				Verified:    false,
			},
			IndeedProfile:     domain.IndeedProfile{},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go: {},
				domain.Rust: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Rust Engineer / Real-Time Bidding",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4068403021/",
							Date:             mustDate("2024-11-04"),
						},
					},
				},
				domain.Zig: {},
				domain.Scala: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Scala Developer",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4068296805/",
							Date:             mustDate("2024-11-04"),
						},
					},
				},
				domain.Elixir:  {},
				domain.Clojure: {},
				domain.Haskell: {},
			},
			ShortDescription: "",
			DealroomURL:      "",
			CrunchbaseURL:    "",
			PitchbookURL:     "",
			YahooFinanceURL:  "",
			GoogleFinanceURL: "",
			YCombinatorURL:   "",
			Industries: []domain.Industry{
				domain.IndustryAdTech,
			},
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Tala",
			Website: "https://tala.co/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                2647513,
				Alias:             "talamobile",
				Name:              "Tala",
				Followers:         "52K",
				Employees:         "500+",
				AssociatedMembers: "792",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "inventure",
				GoRepositoryCount: 2, // Scala 2
				Verified:          true,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Tala-EI_IE1264165.11,15.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Tala-Reviews-E1264165.htm",
				Verified:    true,
			},
			OttaProfileSlug:   "Tala",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go:   nil,
				domain.Rust: nil,
				domain.Zig:  nil,
				domain.Scala: []string{
					"https://www.linkedin.com/jobs/view/4048262224/", // Senior Scala Developer
				},
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
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
			ShortDescription: "Financial platform",
			Industries: []domain.Industry{
				domain.IndustryFinTech,
			},
			HasEmployeesFromCountries: []domain.Country{},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "MacroHealth",
			Website: "https://macrohealth.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                19012089,
				Alias:             "macrohealth",
				Name:              "MacroHealth",
				Followers:         "13K",
				Employees:         "50+",
				AssociatedMembers: "120",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "",
				GoRepositoryCount: 0,
				Verified:          false,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-MacroHealth-EI_IE3277977.11,22.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/MacroHealth-Reviews-E3277977.htm",
				Verified:    false,
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go:   nil,
				domain.Rust: nil,
				domain.Zig:  nil,
				domain.Scala: []string{
					"https://www.linkedin.com/jobs/view/4051275869/", // Maintain and extend applications using Scala and Java
				},
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
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
			ShortDescription: "Health Market-as-a-Service platform that enables Payers, Providers, and Health Market Partners",
			Industries:       []domain.Industry{
				// HealthTech, eCommerce?
			},
			HasEmployeesFromCountries: []domain.Country{},
		},
		{
			// Part of DoorDash
			ID:      0,                         // system
			Type:    domain.CompanyTypeStartup, // system
			Name:    "Wolt",
			Website: "https://careers.wolt.com/en",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                3995846,
				Alias:             "wolt-oy",
				Name:              "Wolt",
				Followers:         "177K",
				Employees:         "10K+",
				AssociatedMembers: "9,722",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "woltapp",
				GoRepositoryCount: 0, // Scala 2
				Verified:          true,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Wolt-EI_IE1800227.11,15.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Wolt-Reviews-E1800227.htm",
				Verified:    true,
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go:   nil,
				domain.Rust: nil,
				domain.Zig:  nil,
				domain.Scala: []string{
					"https://www.linkedin.com/jobs/view/4008937766/", // You have solid experience designing, building, and integrating quality software (microservices) in Scala
				},
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
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
			ShortDescription: "Company making it easy to discover and get restaurants and shops delivered home & to the office",
			Industries:       []domain.Industry{
				// Food delivery
			},
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
				domain.Czechia,
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Bynder",
			Website: "https://www.bynder.com/en/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                2427738,
				Alias:             "bynder",
				Name:              "Bynder",
				Followers:         "38K",
				Employees:         "500+",
				AssociatedMembers: "603",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "bynder",
				GoRepositoryCount: 3, // Scala 0
				Verified:          true,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Bynder-EI_IE1003066.11,17.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Bynder-Reviews-E1003066.htm",
				Verified:    false,
			},
			OttaProfileSlug:   "Bynder",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go:   nil,
				domain.Rust: nil,
				domain.Zig:  nil,
				domain.Scala: []string{
					"https://www.linkedin.com/jobs/view/4011968287/", // Experience building microservices with containers, Scala and Reactive
				},
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
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
			ShortDescription:          "Digital asset management platform",
			Industries:                []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{},
		},
		{
			ID:      0,                         // system
			Type:    domain.CompanyTypeStartup, // system
			Name:    "OVO",
			Website: "https://company.ovo.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                643101,
				Alias:             "ovoenergy",
				Name:              "OVO",
				Followers:         "70K",
				Employees:         "5K+",
				AssociatedMembers: "3,258",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "ovotech",
				GoRepositoryCount: 10, // Scala 29
				Verified:          true,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-OVO-EI_IE767881.11,14.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/OVO-Reviews-E767881.htm",
				Verified:    true,
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go:   nil,
				domain.Rust: nil,
				domain.Zig:  nil,
				domain.Scala: []string{
					"https://www.linkedin.com/jobs/view/4051591843/", // Scala Software Engineer
				},
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
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
			ShortDescription: "Clean affordable energy",
			Industries:       []domain.Industry{
				// Energy
			},
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
			},
		},
		{
			ID:      0,                         // system
			Type:    domain.CompanyTypeStartup, // system
			Name:    "Avetta",
			Website: "https://www.avetta.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                7584447,
				Alias:             "avetta",
				Name:              "Avetta",
				Followers:         "60K",
				Employees:         "500+",
				AssociatedMembers: "886",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "avetta",
				GoRepositoryCount: 0,
				Verified:          false,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Avetta-EI_IE1182913.11,17.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Avetta-Reviews-E1182913.htm",
				Verified:    true,
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go:   nil,
				domain.Rust: nil,
				domain.Zig:  nil,
				domain.Scala: []string{
					"https://www.linkedin.com/jobs/view/4049869381/", // Scala Software Engineer — Backend
				},
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
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
			ShortDescription:          "Contractor risk management platform",
			Industries:                []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Adverty",
			Website: "https://adverty.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                15233801,
				Alias:             "adverty",
				Name:              "Adverty",
				Followers:         "4K",
				Employees:         "10+",
				AssociatedMembers: "40",
				Verified:          false,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "",
				GoRepositoryCount: 0,
				Verified:          false,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Adverty-EI_IE7208426.11,18.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Adverty-Reviews-E7208426.htm",
				Verified:    false,
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go:   nil,
				domain.Rust: nil,
				domain.Zig:  nil,
				domain.Scala: []string{
					"https://www.linkedin.com/jobs/view/4053026269/", // Backend Scala Engineer
				},
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
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
			ShortDescription: "In-game ad platform",
			Industries: []domain.Industry{
				domain.IndustryAdTech,
			},
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Unzer",
			Website: "https://www.unzer.com/en/", // https://www.youpaylater.com/en/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                67571244,
				Alias:             "unzer",
				Name:              "Unzer",
				Followers:         "11K",
				Employees:         "500+",
				AssociatedMembers: "605",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "unzerdev",
				GoRepositoryCount: 0,
				Verified:          false,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Unzer-EI_IE2871480.11,16.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Unzer-Reviews-E2871480.htm",
				Verified:    false,
			},
			OttaProfileSlug:   "Unzer",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go:   nil,
				domain.Rust: nil,
				domain.Zig:  nil,
				domain.Scala: []string{
					"https://www.linkedin.com/jobs/view/4016971905/", // Scala Software Engineer
				},
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
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
			ShortDescription: "All-in-one payment solutions",
			Industries: []domain.Industry{
				domain.IndustryFinTech,
			},
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "HERE Technologies",
			Website: "https://www.here.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                3237134,
				Alias:             "here",
				Name:              "HERE Technologies",
				Followers:         "238K",
				Employees:         "5K+",
				AssociatedMembers: "8,155",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "heremaps",
				GoRepositoryCount: 0, // Scala 2
				Verified:          false,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-HERE-Technologies-EI_IE753677.11,28.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/HERE-Technologies-Reviews-E753677.htm",
				Verified:    true,
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go:   nil,
				domain.Rust: nil,
				domain.Zig:  nil,
				domain.Scala: []string{
					"https://www.linkedin.com/jobs/view/4053090497/", // Software Engineer (Scala/Java)
				},
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
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
			ShortDescription: "Location platform company",
			Industries:       []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
				domain.Czechia,
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "UBS",
			Website: "https://www.ubs.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                1214,
				Alias:             "ubs",
				Name:              "UBS",
				Followers:         "2M",
				Employees:         "10K+",
				AssociatedMembers: "115,948",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "",
				GoRepositoryCount: 0,
				Verified:          false,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-UBS-EI_IE3419.11,14.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/UBS-Reviews-E3419.htm",
				Verified:    false,
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go:   nil,
				domain.Rust: nil,
				domain.Zig:  nil,
				domain.Scala: []string{
					"https://www.linkedin.com/jobs/view/4034024007/", // Data Engineer — Scala
				},
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
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
			Industries: []domain.Industry{
				domain.IndustryFinTech,
			},
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
				domain.Czechia,
			},
		},
		{
			ID:      0,                         // system
			Type:    domain.CompanyTypeStartup, // system
			Name:    "DroneShield",
			Website: "https://www.droneshield.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                7575557,
				Alias:             "droneshield",
				Name:              "DroneShield",
				Followers:         "18K",
				Employees:         "50+",
				AssociatedMembers: "201",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "",
				GoRepositoryCount: 0,
				Verified:          false,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-DroneShield-EI_IE5175947.11,22.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/DroneShield-Reviews-E5175947.htm",
				Verified:    false,
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go: []string{
					"https://www.linkedin.com/jobs/view/4032752026/", // Software Engineer — Go
				},
				domain.Rust:    nil,
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
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
			ShortDescription: "Counter-UAS solutions",
			Industries:       []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
			},
		},
		{
			ID:      0,                         // system
			Type:    domain.CompanyTypeStartup, // system
			Name:    "Automox",
			Website: "https://www.automox.com/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                7941257,
				Alias:             "automox",
				Name:              "Automox",
				Followers:         "11K",
				Employees:         "200+",
				AssociatedMembers: "252",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "automox",
				GoRepositoryCount: 0,
				Verified:          false,
			},
			BlindProfile:     domain.BlindProfile{},
			LevelsFyiProfile: domain.LevelsFyiProfile{},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Automox-EI_IE2236326.11,18.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Automox-Reviews-E2236326.htm",
				Verified:    true,
			},
			IndeedProfile:     domain.IndeedProfile{},
			OttaProfileSlug:   "Automox",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Staff Software Engineer (Golang)",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4076698129/",
							Date:             mustDate("2024-11-14"),
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
			ShortDescription: "Cloud-native automated patch management platform",
			DealroomURL:      "",
			CrunchbaseURL:    "",
			PitchbookURL:     "",
			YahooFinanceURL:  "",
			GoogleFinanceURL: "",
			YCombinatorURL:   "",
			Industries: []domain.Industry{
				domain.IndustryDevOps, // Programming experience in Golang is recommended
			},
			HasEmployeesFromCountries: []domain.Country{},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Darktrace",
			Website: "https://darktrace.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                5013440,
				Alias:             "darktrace",
				Name:              "Darktrace",
				Followers:         "217K",
				Employees:         "1K+",
				AssociatedMembers: "2,741",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "darktrace",
				GoRepositoryCount: 0,
				Verified:          false,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Darktrace-EI_IE1059406.11,20.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Darktrace-Reviews-E1059406.htm",
				Verified:    true,
			},
			OttaProfileSlug:   "Darktrace",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go: nil,
				domain.Rust: []string{
					"https://www.linkedin.com/jobs/view/4056914662/", // Predominately working with core software modules written in Python and Rust,
				},
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
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
			Industries: []domain.Industry{
				domain.IndustryCyberSecurity,
			},
			HasEmployeesFromCountries: []domain.Country{
				domain.Czechia,
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Poppi Technologies",
			Website: "https://www.poppitechnologies.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                93580213,
				Alias:             "poppi-technologies",
				Name:              "Poppi Technologies",
				Followers:         "300",
				Employees:         "10+",
				AssociatedMembers: "7",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "",
				GoRepositoryCount: 0,
				Verified:          false,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "",
				ReviewsURL:  "",
				Verified:    false,
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go: nil,
				domain.Rust: []string{
					"https://www.linkedin.com/jobs/view/4056511585/",
				},
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
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
			Industries: []domain.Industry{
				domain.IndustryFinTech,
			},
			HasEmployeesFromCountries: []domain.Country{},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Booz Allen Hamilton",
			Website: "https://www.boozallen.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                1403,
				Alias:             "booz-allen-hamilton",
				Name:              "Booz Allen Hamilton",
				Followers:         "716K",
				Employees:         "10K+",
				AssociatedMembers: "39,622",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "boozallen",
				GoRepositoryCount: 0,
				Verified:          false,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Booz-Allen-Hamilton-EI_IE2735.11,30.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Booz-Allen-Hamilton-Reviews-E2735.htm",
				Verified:    true,
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go: nil,
				domain.Rust: []string{
					"https://www.linkedin.com/jobs/view/4056310581/",
				},
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
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
			Industries:       []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
				domain.Czechia,
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Flexera",
			Website: "https://www.flexera.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                574962,
				Alias:             "flexera",
				Name:              "Flexera",
				Followers:         "111K",
				Employees:         "1K+",
				AssociatedMembers: "1,669",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "flexera-public",
				GoRepositoryCount: 32,
				Verified:          false,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Flexera-EI_IE304392.11,18.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Flexera-Reviews-E304392.htm",
				Verified:    true, // openCompany ???
			},
			OttaProfileSlug:   "Flexera",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go: []string{
					"https://www.linkedin.com/jobs/view/4033986551/", // Design, implement and run microservices written in Go and C#
				},
				domain.Rust:    nil,
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
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
			ShortDescription: "IT management solutions",
			Industries:       []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
			},
		},
		{
			ID:      0,                         // system
			Type:    domain.CompanyTypeStartup, // system
			Name:    "Flowhub",
			Website: "https://flowhub.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                4812785,
				Alias:             "flowhub",
				Name:              "Flowhub",
				Followers:         "36K",
				Employees:         "50+",
				AssociatedMembers: "81",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "",
				GoRepositoryCount: 0,
				Verified:          false,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Flowhub-EI_IE1000410.11,18.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Flowhub-Reviews-E1000410.htm",
				Verified:    true,
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go: []string{
					"https://www.linkedin.com/jobs/view/4053457618/",
				},
				domain.Rust:    nil,
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
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
			ShortDescription: "Cannabis retail platform",
			Industries: []domain.Industry{
				domain.IndustryECommerce,
			},
			HasEmployeesFromCountries: []domain.Country{},
		},
		{
			ID:      0,                         // system
			Type:    domain.CompanyTypeStartup, // system
			Name:    "Teachable",
			Website: "https://teachable.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                4815896,
				Alias:             "teachable",
				Name:              "Teachable",
				Followers:         "37K",
				Employees:         "50+",
				AssociatedMembers: "294",
				Verified:          false,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "usefedora",
				GoRepositoryCount: 0,
				Verified:          false,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Teachable-EI_IE1614390.11,20.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Teachable-Reviews-E1614390.htm",
				Verified:    true,
			},
			OttaProfileSlug:   "Teachable",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go: []string{
					"https://www.linkedin.com/jobs/view/4055915782/", // Proficiency using Go in a production environment
				},
				domain.Rust:    nil,
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
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
			ShortDescription: "Online teaching platform",
			Industries: []domain.Industry{
				domain.IndustryEdTech,
			},
			HasEmployeesFromCountries: []domain.Country{},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Lloyds Banking Group",
			Website: "https://www.lloydsbankinggroup.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                417361,
				Alias:             "lloyds-banking-group",
				Name:              "Lloyds Banking Group",
				Followers:         "485K",
				Employees:         "10K+",
				AssociatedMembers: "60,748",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "",
				GoRepositoryCount: 0,
				Verified:          false,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "",
				ReviewsURL:  "",
				Verified:    false,
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go: []string{
					"https://www.linkedin.com/jobs/view/4056705575/", // Strong proficiency in Golang
				},
				domain.Rust:    nil,
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
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
			Name:    "CyberArk",
			Website: "https://www.cyberark.com/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                26630,
				Alias:             "cyber-ark-software",
				Name:              "CyberArk",
				Followers:         "224K",
				Employees:         "1K+",
				AssociatedMembers: "3,825",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "cyberark",
				Verified: false,
			},
			BlindProfile:     domain.BlindProfile{},
			LevelsFyiProfile: domain.LevelsFyiProfile{},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-CyberArk-EI_IE30042.11,19.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/CyberArk-Reviews-E30042.htm",
				Verified:    false,
			},
			IndeedProfile:     domain.IndeedProfile{},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 21,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Senior Staff Software Engineer, Golang",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4077219379/",
							Date:             mustDate("2024-11-19"),
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
			ShortDescription: "Machine Identity Security",
			DealroomURL:      "",
			CrunchbaseURL:    "",
			PitchbookURL:     "",
			YahooFinanceURL:  "",
			GoogleFinanceURL: "",
			YCombinatorURL:   "",
			Industries: []domain.Industry{
				domain.IndustryCyberSecurity,
			},
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
				domain.Czechia,
			},
		},
		{
			ID:      0,                         // system
			Type:    domain.CompanyTypeStartup, // system
			Name:    "PayNearMe",
			Website: "https://home.paynearme.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                1137260,
				Alias:             "paynearme",
				Name:              "PayNearMe",
				Followers:         "38K",
				Employees:         "200+",
				AssociatedMembers: "227",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "paynearme",
				GoRepositoryCount: 0,
				Verified:          false,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-PayNearMe-EI_IE954212.11,20.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/PayNearMe-Reviews-E954212.htm",
				Verified:    true,
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go: []string{
					"https://www.linkedin.com/jobs/view/4058466151/", // Producing high-performance, reliable, and secure service implementations in Go
					"https://www.linkedin.com/jobs/view/4058460891/",
				},
				domain.Rust:    nil,
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
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
			Industries: []domain.Industry{
				domain.IndustryFinTech,
			},
			HasEmployeesFromCountries: []domain.Country{},
		},
		{
			ID:      0,                         // system
			Type:    domain.CompanyTypeStartup, // system
			Name:    "Creative Fabrica",
			Website: "https://www.creativefabrica.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                10854260,
				Alias:             "creative-fabrica",
				Name:              "Creative Fabrica",
				Followers:         "9K",
				Employees:         "50+",
				AssociatedMembers: "380",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "creativefabrica",
				GoRepositoryCount: 0,
				Verified:          true,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Creative-Fabrica-EI_IE3892703.11,27.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Creative-Fabrica-Reviews-E3892703.htm",
				Verified:    true,
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go: []string{
					"https://www.linkedin.com/jobs/view/4001594627/", // Experience with Go
				},
				domain.Rust:    nil,
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
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
			ShortDescription:          "",
			Industries:                []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{},
		},
		{
			ID:      0,                         // system
			Type:    domain.CompanyTypeStartup, // system
			Name:    "CAST AI",
			Website: "https://cast.ai/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                30767670,
				Alias:             "cast-ai",
				Name:              "CAST AI",
				Followers:         "13K",
				Employees:         "50+",
				AssociatedMembers: "218",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "castai",
				GoRepositoryCount: 25,
				Verified:          true,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Cast-AI-EI_IE3133117.11,18.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Cast-AI-Reviews-E3133117.htm",
				Verified:    true,
			},
			OttaProfileSlug:   "CAST-AI",
			YouTubeChannelURL: "",
			GoMainLanguage:    true,
			Vacancies: domain.Vacancies{
				domain.Go: []string{
					"https://www.linkedin.com/jobs/view/3982723810/", // GoLang is our language
				},
				domain.Rust:    nil,
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
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
			ShortDescription: "Cloud cost optimization & automation",
			Industries: []domain.Industry{
				domain.IndustryDevOps,
			},
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Grasshopper",
			Website: "https://grasshopperasia.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                14373895,
				Alias:             "grasshopperasia",
				Name:              "Grasshopper",
				Followers:         "4K",
				Employees:         "50+",
				AssociatedMembers: "110",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "ghpr-asia",
				GoRepositoryCount: 13, // Rust 5
				Verified:          false,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Grasshopper-EI_IE1145365.11,22.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Grasshopper-Reviews-E1145365.htm",
				Verified:    true,
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go: nil,
				domain.Rust: []string{
					"https://www.linkedin.com/jobs/view/3848418692/", // Utilise your expertise in Rust to write robust and performant code
				},
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
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
			ShortDescription: "Quantitative trading technology provide",
			Industries: []domain.Industry{
				domain.IndustryFinTech,
			},
			HasEmployeesFromCountries: []domain.Country{},
		},
		{
			ID:      0,                         // system
			Type:    domain.CompanyTypeStartup, // system
			Name:    "Swish Analytics",
			Website: "https://swishanalytics.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                3886282,
				Alias:             "swish-analytics",
				Name:              "Swish Analytics",
				Followers:         "9K",
				Employees:         "10+",
				AssociatedMembers: "122",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "swishanalytics",
				GoRepositoryCount: 0,
				Verified:          false,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Swish-Analytics-EI_IE2104030.11,26.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Swish-Analytics-Reviews-E2104030.htm",
				Verified:    false,
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go: nil,
				domain.Rust: []string{
					"https://www.linkedin.com/jobs/view/4054776468/", // Proactively improve our Rust and Python codebase
				},
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
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
			ShortDescription:          "Sports Betting Solutions",
			Industries:                []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Abbott",
			Website: "https://www.abbott.com/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                1612,
				Alias:             "abbott-",
				Name:              "Abbott",
				Followers:         "4M",
				Employees:         "10K+",
				AssociatedMembers: "164,356",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "",
				GoRepositoryCount: 0,
				Verified:          false,
			},
			BlindProfile:     domain.BlindProfile{},
			LevelsFyiProfile: domain.LevelsFyiProfile{},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Abbott-EI_IE12.11,17.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Abbott-Reviews-E12.htm",
				Verified:    true,
			},
			IndeedProfile:     domain.IndeedProfile{},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Senior Software Engineer, Cloud (Go/Golang)",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4076610998/",
							Date:             mustDate("2024-11-14"),
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
			DealroomURL:      "",
			CrunchbaseURL:    "",
			PitchbookURL:     "",
			YahooFinanceURL:  "",
			GoogleFinanceURL: "",
			YCombinatorURL:   "",
			Industries: []domain.Industry{
				domain.IndustryHealthTech,
			},
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
				domain.Czechia,
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Dealer Tire",
			Website: "https://www.dealertire.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                21738,
				Alias:             "dealer-tire",
				Name:              "Dealer Tire",
				Followers:         "53K",
				Employees:         "1K+",
				AssociatedMembers: "1,434",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "",
				GoRepositoryCount: 0,
				Verified:          false,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Dealer-Tire-EI_IE39809.11,22.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Dealer-Tire-Reviews-E39809.htm",
				Verified:    true,
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go: []string{
					"https://www.linkedin.com/jobs/view/4053146259/",
				},
				domain.Rust:    nil,
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
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
			ShortDescription:          "Distributor of tires and parts",
			Industries:                []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "JPMorganChase",
			Website: "https://www.jpmorganchase.com/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                1068,
				Alias:             "jpmorganchase",
				Name:              "JPMorganChase",
				Followers:         "5M",
				Employees:         "10K+",
				AssociatedMembers: "301,391",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "jpmorganchase",
				GoRepositoryCount: 3,
				Verified:          false,
			},
			BlindProfile:     domain.BlindProfile{},
			LevelsFyiProfile: domain.LevelsFyiProfile{},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-JPMorganChase-EI_IE5224839.11,24.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/JPMorganChase-Reviews-E5224839.htm",
				Verified:    true,
			},
			IndeedProfile:     domain.IndeedProfile{},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Software Engineer III — Golang, AWS, Terraform",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4050118492/",
							Date:             mustDate("2024-11-03"),
						},
						{
							Title:            "Software Engineer II —  Core Engineering —  GO / Python",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4105782609/",
							Date:             mustDate("2024-12-19"),
						},
					},
				},
				domain.Rust: {},
				domain.Zig:  {},
				domain.Scala: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Lead Software Engineer — Scala",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4055220537/",
							Date:             mustDate("2024-10-25"),
						},
					},
				},
				domain.Elixir:  {},
				domain.Clojure: {},
				domain.Haskell: {},
			},
			ShortDescription: "",
			DealroomURL:      "",
			CrunchbaseURL:    "",
			PitchbookURL:     "",
			YahooFinanceURL:  "",
			GoogleFinanceURL: "",
			YCombinatorURL:   "",
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
			Name:    "Outreach",
			Website: "https://www.outreach.io/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                3559595,
				Alias:             "outreach-saas",
				Name:              "Outreach",
				Followers:         "220K",
				Employees:         "500+",
				AssociatedMembers: "1,317",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "getoutreach",
				GoRepositoryCount: 63,
				Verified:          true,
			},
			BlindProfile:     domain.BlindProfile{},
			LevelsFyiProfile: domain.LevelsFyiProfile{},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Outreach-EI_IE1276320.11,19.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Outreach-Reviews-E1276320.htm",
				Verified:    true,
			},
			IndeedProfile:     domain.IndeedProfile{},
			OttaProfileSlug:   "Outreach",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 63,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Senior Staff Software Engineer (Backend — Golang)",
							ShortDescription: "We primarily use Go language for our next gen applications",
							URL:              "https://www.linkedin.com/jobs/view/4074568575/",
							Date:             mustDate("2024-11-14"),
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
			ShortDescription: "Sales execution platform",
			DealroomURL:      "",
			CrunchbaseURL:    "",
			PitchbookURL:     "",
			YahooFinanceURL:  "",
			GoogleFinanceURL: "",
			YCombinatorURL:   "",
			Industries:       []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
				domain.Czechia,
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Tucows",
			Website: "https://www.tucows.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                166530,
				Alias:             "tucows",
				Name:              "Tucows",
				Followers:         "36K",
				Employees:         "500+",
				AssociatedMembers: "841",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "tucows",
				GoRepositoryCount: 4,
				Verified:          false,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Tucows-EI_IE6018.11,17.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Tucows-Reviews-E6018.htm",
				Verified:    true,
			},
			OttaProfileSlug:   "Tucows",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go: []string{
					"https://www.linkedin.com/jobs/view/4004609791/", // Proficiency in Python or Golang programming languages
				},
				domain.Rust:    nil,
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
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
			ShortDescription: "Domain registrar",
			Industries:       []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "ING",
			Website: "https://www.ing.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                2594164,
				IDs:               []int{2594164, 299201, 1523089, 387202, 71625717, 3702090, 3622172, 40711, 516340, 18071794, 2038590, 698107, 2105579, 2924419, 2722195, 215713},
				Alias:             "ing",
				Name:              "ING",
				Followers:         "559K",
				Employees:         "10K+",
				AssociatedMembers: "66,484",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "ing-bank",
				GoRepositoryCount: 9, // Scala 4
				Verified:          false,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-ING-EI_IE4264.11,14.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/ING-Reviews-E4264.htm",
				Verified:    true,
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go:   nil,
				domain.Rust: nil,
				domain.Zig:  nil,
				domain.Scala: []string{
					"https://www.linkedin.com/jobs/view/4053346980/", // 5+ years of experience in Scala
				},
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
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
			ShortDescription: "Bank",
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
			Name:    "Remitly",
			Website: "https://www.remitly.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                2261199,
				Alias:             "remitly",
				Name:              "Remitly",
				Followers:         "59K",
				Employees:         "1K+",
				AssociatedMembers: "2,947",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "remitly",
				GoRepositoryCount: 8,
				Verified:          true,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Remitly-EI_IE1044836.11,18.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Remitly-Reviews-E1044836.htm",
				Verified:    true,
			},
			OttaProfileSlug:   "Remitly",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go: []string{
					"https://www.linkedin.com/jobs/view/4050285116/", // The ideal candidate is a backend engineer with Golang experience
				},
				domain.Rust:    nil,
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
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
			ShortDescription: "Payments company",
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
			Name:    "Level All",
			Website: "https://www.levelall.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                79136550,
				Alias:             "level-all",
				Name:              "Level All",
				Followers:         "17K",
				Employees:         "50+",
				AssociatedMembers: "88",
				Verified:          false,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "Level-All",
				Verified: false,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "",
				ReviewsURL:  "",
				Verified:    false,
			},
			OttaProfileSlug:   "Level-All",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go:    {},
				domain.Rust:  {},
				domain.Zig:   {},
				domain.Scala: {},
				domain.Elixir: {
					GitHubRepositoriesCount: 6,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Senior Software Engineer",
							ShortDescription: "We’re looking for a Senior Elixir Software Engineer to contribute to our core web platform and supporting applications built in Elixir, Phoenix, LiveView, TypeScript, and Tailwind, AlpineJS backed with a Postgres database",
							URL:              "https://www.linkedin.com/jobs/view/4092211710/",
							Date:             mustDate("2024-12-06"),
						},
					},
				},
				domain.Clojure: {},
				domain.Haskell: {},
			},
			ShortDescription: "Personalized college & career guidance",
			Industries: []domain.Industry{
				domain.IndustryEdTech,
			},
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "LittleLives",
			Website: "https://www.littlelives.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                6577681,
				Alias:             "littlelives",
				Name:              "LittleLives",
				Followers:         "4K",
				Employees:         "10+",
				AssociatedMembers: "76",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "teamlittlelives",
				GoRepositoryCount: 0,
				Verified:          false,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-LittleLives-EI_IE959760.11,22.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/LittleLives-Reviews-E959760.htm",
				Verified:    false,
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go:    nil,
				domain.Rust:  nil,
				domain.Zig:   nil,
				domain.Scala: nil,
				domain.Elixir: []string{
					"https://www.linkedin.com/jobs/view/4046403243/", // strong foundation in Elixir
				},
				domain.Clojure: nil,
				domain.Haskell: nil,
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
			ShortDescription: "School Management System",
			Industries: []domain.Industry{
				domain.IndustryEdTech,
			},
			HasEmployeesFromCountries: []domain.Country{},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "BILL",
			Website: "https://www.bill.com/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                113254,
				Alias:             "bill",
				Name:              "BILL",
				Followers:         "60K",
				Employees:         "1K+",
				AssociatedMembers: "3,137",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "",
				GoRepositoryCount: 0,
				Verified:          false,
			},
			BlindProfile:     domain.BlindProfile{},
			LevelsFyiProfile: domain.LevelsFyiProfile{},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-BILL-EI_IE801594.11,15.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/BILL-Reviews-E801594.htm",
				Verified:    true,
			},
			IndeedProfile:     domain.IndeedProfile{},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "",
			Languages: domain.Languages{
				domain.Go:    {},
				domain.Rust:  {},
				domain.Zig:   {},
				domain.Scala: {},
				domain.Elixir: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Software Engineer (Elixir)",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/3906024674/",
							Date:             mustDate("2024-11-04"),
						},
					},
				},
				domain.Clojure: {},
				domain.Haskell: {},
			},
			ShortDescription: "Financial automation software",
			DealroomURL:      "",
			CrunchbaseURL:    "",
			PitchbookURL:     "",
			YahooFinanceURL:  "",
			GoogleFinanceURL: "",
			YCombinatorURL:   "",
			Industries: []domain.Industry{
				domain.IndustryFinTech,
			},
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
			},
		},
		{
			ID:      0,                         // system
			Type:    domain.CompanyTypeStartup, // system
			Name:    "Community",
			Website: "https://www.community.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                18945668,
				Alias:             "communitydotcom",
				Name:              "Community",
				Followers:         "17K",
				Employees:         "50+",
				AssociatedMembers: "81",
				Verified:          false,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "",
				GoRepositoryCount: 0,
				Verified:          false,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-community-com-EI_IE3005175.11,24.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/community-com-Reviews-E3005175.htm",
				Verified:    true,
			},
			OttaProfileSlug:   "Community",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go:    nil,
				domain.Rust:  nil,
				domain.Zig:   nil,
				domain.Scala: nil,
				domain.Elixir: []string{
					"https://www.linkedin.com/jobs/view/4041216496/", // Design, develop, and maintain robust, scalable backend services primarily using Elixir
				},
				domain.Clojure: nil,
				domain.Haskell: nil,
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
			ShortDescription:          "SMS marketing platform",
			Industries:                []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{},
		},
		{
			ID:      0,                         // system
			Type:    domain.CompanyTypeStartup, // system
			Name:    "BigPay",
			Website: "https://bigpayme.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                1067442,
				Alias:             "bigpayme",
				Name:              "BigPay",
				Followers:         "25K",
				Employees:         "200+",
				AssociatedMembers: "190",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "bigpay",
				GoRepositoryCount: 0, // Rust 2
				Verified:          false,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-BigPay-EI_IE2497022.11,17.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/BigPay-Reviews-E2497022.htm",
				Verified:    true,
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go: nil,
				domain.Rust: []string{
					"https://www.linkedin.com/jobs/view/4052948882/", // Senior Backend Engineer (Rust)
				},
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
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
			ShortDescription: "Fintech company",
			Industries: []domain.Industry{
				domain.IndustryFinTech,
			},
			HasEmployeesFromCountries: []domain.Country{},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Rabbet",
			Website: "https://rabbet.com/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                18937694,
				Alias:             "rabbet",
				Name:              "Rabbet",
				Followers:         "3K",
				Employees:         "10+",
				AssociatedMembers: "43",
				Verified:          false,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "Rabbet",
				GoRepositoryCount: 0, // Elixir 3
				Verified:          false,
			},
			BlindProfile:     domain.BlindProfile{},
			LevelsFyiProfile: domain.LevelsFyiProfile{},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Rabbet-EI_IE1805037.11,17.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Rabbet-Reviews-E1805037.htm",
				Verified:    false,
			},
			IndeedProfile:     domain.IndeedProfile{},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go:    {},
				domain.Rust:  {},
				domain.Zig:   {},
				domain.Scala: {},
				domain.Elixir: {
					GitHubRepositoriesCount: 3,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Senior Full Stack Engineer (Elixir / React)",
							ShortDescription: "Back-end is primarily Elixir/Phoenix/Absinthe and hosted in GCP",
							URL:              "https://www.linkedin.com/jobs/view/4041341979/",
							Date:             mustDate("2024-10-14"),
						},
					},
				},
				domain.Clojure: {},
				domain.Haskell: {},
			},
			ShortDescription: "",
			DealroomURL:      "",
			CrunchbaseURL:    "",
			PitchbookURL:     "",
			YahooFinanceURL:  "",
			GoogleFinanceURL: "",
			YCombinatorURL:   "",
			Industries: []domain.Industry{
				domain.IndustryFinTech,
			},
			HasEmployeesFromCountries: []domain.Country{},
		},
		{
			ID:      0,                         // system
			Type:    domain.CompanyTypeStartup, // system
			Name:    "Agentero",
			Website: "https://agentero.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                17946876,
				Alias:             "agentero",
				Name:              "Agentero",
				Followers:         "2K",
				Employees:         "10+",
				AssociatedMembers: "42",
				Verified:          false,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "",
				GoRepositoryCount: 0,
				Verified:          false,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Agentero-EI_IE2215731.11,19.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Agentero-Reviews-E2215731.htm",
				Verified:    false,
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go: []string{
					"https://www.linkedin.com/jobs/view/4050442624/", // You will write clean, robust, and performant software using Go
				},
				domain.Rust:    nil,
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
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
			ShortDescription:          "Digital insurance network",
			Industries:                []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Marbill Technologies",
			Website: "https://marbill.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                18914675,
				Alias:             "marbill-technologies",
				Name:              "Marbill Technologies",
				Followers:         "10K",
				Employees:         "200+",
				AssociatedMembers: "95",
				Verified:          false,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "",
				GoRepositoryCount: 0,
				Verified:          false,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "",
				ReviewsURL:  "",
				Verified:    false,
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go: []string{
					"https://www.linkedin.com/jobs/view/4057281091/", // 5+ years experience with Go
				},
				domain.Rust:    nil,
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
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
			ShortDescription: "Platform for the e-commerce industry",
			Industries: []domain.Industry{
				domain.IndustryECommerce,
			},
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "BiTaksi",
			Website: "https://bitaksi.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                2995287,
				Alias:             "bitaksi",
				Name:              "BiTaksi",
				Followers:         "8K",
				Employees:         "50+",
				AssociatedMembers: "118",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "BiTaksi",
				GoRepositoryCount: 0,
				Verified:          true,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-BiTaksi-EI_IE2994794.11,18.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/BiTaksi-Reviews-E2994794.htm",
				Verified:    false,
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go: []string{
					"https://www.linkedin.com/jobs/view/4053855198/", // Experience with Golang
				},
				domain.Rust:    nil,
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
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
			ShortDescription:          "Taxi-hailing app",
			Industries:                []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{},
		},
		{
			ID:      0,                         // system
			Type:    domain.CompanyTypeStartup, // system
			Name:    "Oddin.gg",
			Website: "https://oddin.gg/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                19137419,
				Alias:             "oddingg",
				Name:              "Oddin.gg",
				Followers:         "6K",
				Employees:         "200+",
				AssociatedMembers: "73",
				Verified:          false,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "oddin-gg",
				GoRepositoryCount: 1,
				Verified:          false,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Oddin-EI_IE5084636.11,16.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Oddin-Reviews-E5084636.htm",
				Verified:    false,
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go: []string{
					"https://www.linkedin.com/jobs/view/4052970416/", // 5 years of experience with Golang
				},
				domain.Rust:    nil,
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
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
			ShortDescription: "Esports betting ecosystem",
			Industries:       []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{
				domain.Czechia,
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Mimecast",
			Website: "https://www.mimecast.com/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                55895,
				Alias:             "mimecast",
				Name:              "Mimecast",
				Followers:         "102K",
				Employees:         "1K+",
				AssociatedMembers: "2,608",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "mimecast",
				GoRepositoryCount: 1,
				Verified:          false,
			},
			BlindProfile:     domain.BlindProfile{},
			LevelsFyiProfile: domain.LevelsFyiProfile{},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Mimecast-EI_IE221309.11,19.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Mimecast-Reviews-E221309.htm",
				Verified:    true,
			},
			IndeedProfile:     domain.IndeedProfile{},
			OttaProfileSlug:   "Mimecast",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Senior Software Engineer — GO",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/3987632672/",
							Date:             mustDate("2024-11-16"),
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
			ShortDescription: "Cloud cybersecurity services",
			DealroomURL:      "",
			CrunchbaseURL:    "",
			PitchbookURL:     "",
			YahooFinanceURL:  "",
			GoogleFinanceURL: "",
			YCombinatorURL:   "",
			Industries: []domain.Industry{
				domain.IndustryCyberSecurity,
			},
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
				domain.Czechia,
			},
		},
		{
			ID:      0,                         // system
			Type:    domain.CompanyTypeStartup, // system
			Name:    "Sentry",
			Website: "https://sentry.io/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                6424460,
				Alias:             "getsentry",
				Name:              "Sentry",
				Followers:         "20K",
				Employees:         "200+",
				AssociatedMembers: "392",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "getsentry",
				GoRepositoryCount: 22, // Rust 67
				Verified:          true,
			},
			BlindProfile:     domain.BlindProfile{},
			LevelsFyiProfile: domain.LevelsFyiProfile{},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Sentry-CA-EI_IE1622271.11,20.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Sentry-CA-Reviews-E1622271.htm",
				Verified:    true,
			},
			IndeedProfile:     domain.IndeedProfile{},
			OttaProfileSlug:   "Sentry",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 22,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Senior Software Engineer (Go), SDK",
							ShortDescription: "Work in the client infrastructure team to improve and evolve our https://github.com/getsentry/sentry-go",
							URL:              "https://www.linkedin.com/jobs/view/4042362901/",
							Date:             mustDate("2024-11-16"),
						},
					},
				},
				domain.Rust: {
					GitHubRepositoriesCount: 67,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Senior Software Engineer, Core Product",
							ShortDescription: "Familiarity with Rust",
							URL:              "https://www.linkedin.com/jobs/view/3997074001/",
							Date:             mustDate("2024-11-08"),
						},
					},
				},
				domain.Zig:     {},
				domain.Scala:   {},
				domain.Elixir:  {},
				domain.Clojure: {},
				domain.Haskell: {},
			},
			ShortDescription: "Error tracking and app monitoring software",
			DealroomURL:      "",
			CrunchbaseURL:    "",
			PitchbookURL:     "",
			YahooFinanceURL:  "",
			GoogleFinanceURL: "",
			YCombinatorURL:   "",
			Industries: []domain.Industry{
				domain.IndustryDevOps,
			},
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
				domain.Czechia,
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Track24",
			Website: "https://www.track24.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                315653,
				Alias:             "track24",
				Name:              "Track24",
				Followers:         "9K",
				Employees:         "10+",
				AssociatedMembers: "33",
				Verified:          false,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "",
				GoRepositoryCount: 0,
				Verified:          false,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Track24-EI_IE1463252.11,18.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Track24-Reviews-E1463252.htm",
				Verified:    true,
			},
			OttaProfileSlug:   "Track24",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go:    nil,
				domain.Rust:  nil,
				domain.Zig:   nil,
				domain.Scala: nil,
				domain.Elixir: []string{
					"https://www.linkedin.com/jobs/view/4055781039/", // Elixir Engineer
				},
				domain.Clojure: nil,
				domain.Haskell: nil,
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
			ShortDescription:          "Operational risk management and communications solutions",
			Industries:                []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{},
		},
		{
			ID:      0,                         // system
			Type:    domain.CompanyTypeStartup, // system
			Name:    "Hiive",
			Website: "https://www.hiive.com/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                76581616,
				Alias:             "hiive",
				Name:              "Hiive",
				Followers:         "10K",
				Employees:         "50+",
				AssociatedMembers: "105",
				Verified:          false,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "",
				GoRepositoryCount: 0,
				Verified:          false,
			},
			BlindProfile:     domain.BlindProfile{},
			LevelsFyiProfile: domain.LevelsFyiProfile{},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Hiive-Canada-EI_IE6926718.11,23.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Hiive-Canada-Reviews-E6926718.htm",
				Verified:    true,
			},
			IndeedProfile:     domain.IndeedProfile{},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go:    {},
				domain.Rust:  {},
				domain.Zig:   {},
				domain.Scala: {},
				domain.Elixir: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Senior Software Engineer (Elixir)",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4060816724/",
							Date:             mustDate("2024-10-29"),
						},
					},
				},
				domain.Clojure: {},
				domain.Haskell: {},
			},
			ShortDescription: "All-in-one liquidity platform for private companies and their shareholders",
			DealroomURL:      "",
			CrunchbaseURL:    "",
			PitchbookURL:     "",
			YahooFinanceURL:  "",
			GoogleFinanceURL: "",
			YCombinatorURL:   "",
			Industries: []domain.Industry{
				domain.IndustryFinTech,
			},
			HasEmployeesFromCountries: []domain.Country{},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "ExpressVPN",
			Website: "https://www.expressvpn.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                14488856,
				Alias:             "expressvpn",
				Name:              "ExpressVPN",
				Followers:         "13K",
				Employees:         "500+",
				AssociatedMembers: "237",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "expressvpn",
				GoRepositoryCount: 0, // Rust 3
				Verified:          true,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-ExpressVPN-EI_IE3355140.11,21.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/ExpressVPN-Reviews-E3355140.htm",
				Verified:    true,
			},
			OttaProfileSlug:   "ExpressVPN",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go: []string{
					"https://www.linkedin.com/jobs/view/4058715430/", // Backend Golang Engineer
				},
				domain.Rust:    nil,
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
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
			ShortDescription: "VPN",
			Industries: []domain.Industry{
				domain.IndustryCyberSecurity,
			},
			HasEmployeesFromCountries: []domain.Country{},
		},
		{
			ID:      0,                         // system
			Type:    domain.CompanyTypeStartup, // system
			Name:    "Shuttle",
			Website: "https://www.shuttle.dev/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                67174759,
				Alias:             "shuttle-yc",
				Name:              "Shuttle",
				Followers:         "3K",
				Employees:         "10+",
				AssociatedMembers: "20",
				Verified:          false,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "shuttle-hq",
				GoRepositoryCount: 1, // Rust 26
				Verified:          false,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "",
				ReviewsURL:  "",
				Verified:    false,
			},
			OttaProfileSlug:   "shuttle",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go: nil,
				domain.Rust: []string{
					"https://www.ycombinator.com/companies/shuttle/jobs/pEYmbIL-senior-software-engineer",
				},
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
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
			ShortDescription:          "Serverless Rust platform",
			Industries:                []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Second Spectrum",
			Website: "https://www.secondspectrum.com/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                3562222,
				Alias:             "second-spectrum",
				Name:              "Second Spectrum",
				Followers:         "9K",
				Employees:         "50+",
				AssociatedMembers: "181",
				Verified:          true,
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
				domain.Go: {},
				domain.Rust: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Staff Software Engineer (Realtime Systems, Rust)",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/3930229478/",
							Date:             mustDate("2024-11-08"),
						},
					},
				},
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
			Name:    "Worldpay",
			Website: "https://www.worldpay.com/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                260387,
				Alias:             "worldpay",
				Name:              "Worldpay",
				Followers:         "221K",
				Employees:         "5K+",
				AssociatedMembers: "8,605",
				Verified:          true,
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
				domain.Go:   {},
				domain.Rust: {},
				domain.Zig:  {},
				domain.Scala: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Senior Software Engineer — Scala & AWS",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/3982972845/",
							Date:             mustDate("2024-11-06"),
						},
					},
				},
				domain.Elixir:  {},
				domain.Clojure: {},
				domain.Haskell: {},
			},
			ShortDescription: "",
			DealroomURL:      "",
			CrunchbaseURL:    "",
			PitchbookURL:     "",
			YahooFinanceURL:  "",
			GoogleFinanceURL: "",
			YCombinatorURL:   "",
			Industries:       []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Exabeam",
			Website: "https://www.exabeam.com/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                3181474,
				Alias:             "exabeam",
				Name:              "Exabeam",
				Followers:         "42K",
				Employees:         "500+",
				AssociatedMembers: "880",
				Verified:          true,
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
				domain.Go:   {},
				domain.Rust: {},
				domain.Zig:  {},
				domain.Scala: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Senior Software Engineer — Java/Scala",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4048214298/",
							Date:             mustDate("2024-11-06"),
						},
					},
				},
				domain.Elixir:  {},
				domain.Clojure: {},
				domain.Haskell: {},
			},
			ShortDescription: "",
			DealroomURL:      "",
			CrunchbaseURL:    "",
			PitchbookURL:     "",
			YahooFinanceURL:  "",
			GoogleFinanceURL: "",
			YCombinatorURL:   "",
			Industries: []domain.Industry{
				domain.IndustryCyberSecurity,
			},
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Selby Jennings",
			Website: "https://www.selbyjennings.co.uk/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                106584,
				Alias:             "selby-jennings",
				Name:              "Selby Jennings",
				Followers:         "789K",
				Employees:         "1K+",
				AssociatedMembers: "833",
				Verified:          true,
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
				domain.Go: {},
				domain.Rust: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Rust Developer | Digital Asset Trading",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4059312901/",
							Date:             mustDate("2024-10-25"),
						},
					},
				},
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
			Name:    "DataBeacon",
			Website: "https://databeacon.aero/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                37898796,
				Alias:             "databeaconaero",
				Name:              "DataBeacon",
				Followers:         "3K",
				Employees:         "10+",
				AssociatedMembers: "8",
				Verified:          false,
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
				domain.Go: {},
				domain.Rust: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Senior Backend Software Engineer, Rust & Python",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4060807594/",
							Date:             mustDate("2024-11-01"),
						},
					},
				},
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
			Name:    "MNTN",
			Website: "https://mountain.com/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                480382,
				Alias:             "wearemntn",
				Name:              "MNTN",
				Followers:         "112K",
				Employees:         "200+",
				AssociatedMembers: "663",
				Verified:          true,
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
				domain.Go: {},
				domain.Rust: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Senior Rust Engineer",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4064837533/",
							Date:             mustDate("2024-11-04"),
						},
					},
				},
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
			Name:    "Writer",
			Website: "https://writer.com/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                67088679,
				Alias:             "getwriter",
				Name:              "Writer",
				Followers:         "49K",
				Employees:         "500+",
				AssociatedMembers: "1,351",
				Verified:          false,
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
				domain.Go:   {},
				domain.Rust: {},
				domain.Zig:  {},
				domain.Scala: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Senior Software Engineer, Backend (Scala)",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/3972187342/",
							Date:             mustDate("2024-11-04"),
						},
					},
				},
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
			Name:    "dLocal",
			Website: "https://www.dlocal.com/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                15156062,
				Alias:             "dlocal",
				Name:              "dLocal",
				Followers:         "140K",
				Employees:         "500+",
				AssociatedMembers: "1,129",
				Verified:          true,
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
							Title:            "Senior Full Stack Engineer (Golang and Node)",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4070887973/",
							Date:             mustDate("2024-11-07"),
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
			DealroomURL:      "",
			CrunchbaseURL:    "",
			PitchbookURL:     "",
			YahooFinanceURL:  "",
			GoogleFinanceURL: "",
			YCombinatorURL:   "",
			Industries:       []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
			},
		},
		{
			ID:      0,                         // system
			Type:    domain.CompanyTypeStartup, // system
			Name:    "Clearstory.build",
			Website: "https://www.clearstory.build/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                18591580,
				Alias:             "clearstory-build",
				Name:              "Clearstory.build",
				Followers:         "7K",
				Employees:         "50+",
				AssociatedMembers: "61",
				Verified:          true,
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
							Title:            "Senior Backend Golang Developer — Philippines",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4034614195/",
							Date:             mustDate("2024-11-02"),
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
			Name:    "Millennium",
			Website: "https://www.mlp.com/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                164987,
				Alias:             "millennium-partners",
				Name:              "Millennium",
				Followers:         "258K",
				Employees:         "5K+",
				AssociatedMembers: "7,236",
				Verified:          true,
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
							Title:            "Backend Engineer — Go",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/3943478838/",
							Date:             mustDate("2024-11-03"),
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
			DealroomURL:      "",
			CrunchbaseURL:    "",
			PitchbookURL:     "",
			YahooFinanceURL:  "",
			GoogleFinanceURL: "",
			YCombinatorURL:   "",
			Industries:       []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
				domain.Czechia,
			},
		},
		{
			ID:      0,                         // system
			Type:    domain.CompanyTypeStartup, // system
			Name:    "Mondoo",
			Website: "https://mondoo.com/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                75622407,
				Alias:             "mondoo",
				Name:              "Mondoo",
				Followers:         "3K",
				Employees:         "10+",
				AssociatedMembers: "53",
				Verified:          false,
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
							Title:            "Go Engineer",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4065503885/",
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
			ShortDescription: "",
			DealroomURL:      "",
			CrunchbaseURL:    "",
			PitchbookURL:     "",
			YahooFinanceURL:  "",
			GoogleFinanceURL: "",
			YCombinatorURL:   "",
			Industries: []domain.Industry{
				domain.IndustryCyberSecurity,
			},
			HasEmployeesFromCountries: []domain.Country{},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "fullinfo",
			Website: "https://www.fullinfo.com/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                71773445,
				Alias:             "fullinfo",
				Name:              "fullinfo",
				Followers:         "8K",
				Employees:         "10+",
				AssociatedMembers: "13",
				Verified:          false,
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
							Title:            "Senior Backend Engineer (Go)",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4065504018/",
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
			Name:    "Bloomberg",
			Website: "https://www.bloomberg.com/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                2494,
				Alias:             "bloomberg",
				Name:              "Bloomberg",
				Followers:         "2M",
				Employees:         "10K+",
				AssociatedMembers: "24,725",
				Verified:          true,
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
							URL:              "https://www.linkedin.com/jobs/view/4062165305/",
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
			ShortDescription: "",
			DealroomURL:      "",
			CrunchbaseURL:    "",
			PitchbookURL:     "",
			YahooFinanceURL:  "",
			GoogleFinanceURL: "",
			YCombinatorURL:   "",
			Industries:       []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
				domain.Czechia,
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Sainsbury's",
			Website: "https://www.sainsburys.co.uk/",
			Careers: "https://sainsburys.jobs/",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                5015,
				Alias:             "sainsburys",
				Name:              "Sainsbury's",
				Followers:         "388K",
				Employees:         "10K+",
				AssociatedMembers: "61,975",
				Verified:          true,
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
							Title:            "Fullstack Engineer — React & Golang",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4062980858/",
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
			Name:    "Stratascale — An SHI Company",
			Website: "https://www.stratascale.com/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                4791832,
				Alias:             "stratascale-llc",
				Name:              "Stratascale – An SHI Company",
				Followers:         "13K",
				Employees:         "200+",
				AssociatedMembers: "233",
				Verified:          true,
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
							Title:            "Senior Software Engineer — Golang",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4005692040/",
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
			ShortDescription: "",
			DealroomURL:      "",
			CrunchbaseURL:    "",
			PitchbookURL:     "",
			YahooFinanceURL:  "",
			GoogleFinanceURL: "",
			YCombinatorURL:   "",
			Industries: []domain.Industry{
				domain.IndustryCyberSecurity,
			},
			HasEmployeesFromCountries: []domain.Country{},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Gartner",
			Website: "https://www.gartner.com/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                2078,
				Alias:             "gartner",
				Name:              "Gartner",
				Followers:         "2M",
				Employees:         "10K+",
				AssociatedMembers: "24,462",
				Verified:          true,
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
							Title:            "Lead Backend Software Engineer (Golang) — Peer Experiences",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4037373961/",
							Date:             mustDate("2024-10-25"),
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
			DealroomURL:      "",
			CrunchbaseURL:    "",
			PitchbookURL:     "",
			YahooFinanceURL:  "",
			GoogleFinanceURL: "",
			YCombinatorURL:   "",
			Industries:       []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
				domain.Czechia,
			},
		},
		{
			ID:      0,                         // system
			Type:    domain.CompanyTypeStartup, // system
			Name:    "Extreme Networks",
			Website: "https://www.extremenetworks.com/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                4761,
				Alias:             "extreme-networks",
				Name:              "Extreme Networks",
				Followers:         "165K",
				Employees:         "1K+",
				AssociatedMembers: "3,875",
				Verified:          true,
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
							Title:            "Senior Software Engineer — Golang — Cloud Networking Developer",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4055078837/",
							Date:             mustDate("2024-10-25"),
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
			Name:    "Yalo",
			Website: "https://www.yalo.ai/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                10251920,
				Alias:             "yalo",
				Name:              "Yalo",
				Followers:         "106K",
				Employees:         "200+",
				AssociatedMembers: "310",
				Verified:          true,
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
							Title:            "Golang Engineer",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4037222694/",
							Date:             mustDate("2024-10-25"),
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
			Name:    "reMarkable",
			Website: "https://remarkable.com/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                10785500,
				Alias:             "remarkable-as",
				Name:              "reMarkable",
				Followers:         "44K",
				Employees:         "200+",
				AssociatedMembers: "565",
				Verified:          true,
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
							Title:            "Senior Golang Developer",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4061088133/",
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
			ShortDescription: "",
			DealroomURL:      "",
			CrunchbaseURL:    "",
			PitchbookURL:     "",
			YahooFinanceURL:  "",
			GoogleFinanceURL: "",
			YCombinatorURL:   "",
			Industries:       []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{
				domain.Czechia,
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Synopsys",
			Website: "https://synopsys.com/",
			Careers: "https://careers.synopsys.com/",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                2457,
				Alias:             "synopsys",
				Name:              "Synopsys",
				Followers:         "737K",
				Employees:         "10K+",
				AssociatedMembers: "19,480",
				Verified:          true,
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
							Title:            "Senior AI Platform Engineer — Golang and Kubernetes",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4027571222/",
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
			ShortDescription: "",
			DealroomURL:      "",
			CrunchbaseURL:    "",
			PitchbookURL:     "",
			YahooFinanceURL:  "",
			GoogleFinanceURL: "",
			YCombinatorURL:   "",
			Industries:       []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
				domain.Czechia,
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "PubMatic",
			Website: "https://pubmatic.com/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                167624,
				Alias:             "pubmatic",
				Name:              "PubMatic",
				Followers:         "335K",
				Employees:         "500+",
				AssociatedMembers: "1,229",
				Verified:          true,
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
							Title:            "Principal Software Engineer (Golang)",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4044830862/",
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
			ShortDescription: "",
			DealroomURL:      "",
			CrunchbaseURL:    "",
			PitchbookURL:     "",
			YahooFinanceURL:  "",
			GoogleFinanceURL: "",
			YCombinatorURL:   "",
			Industries: []domain.Industry{
				domain.IndustryAdTech,
			},
			HasEmployeesFromCountries: []domain.Country{},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Fanatics",
			Website: "https://www.fanaticsinc.com/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                68543,
				Alias:             "fanatics-inc-",
				Name:              "Fanatics",
				Followers:         "198K",
				Employees:         "10K+",
				AssociatedMembers: "9,825",
				Verified:          true,
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
				domain.Go:    {},
				domain.Rust:  {},
				domain.Zig:   {},
				domain.Scala: {},
				domain.Elixir: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Software Engineer III",
							ShortDescription: "Our stack primarily uses Elixir with the Phoenix Framework",
							URL:              "https://www.linkedin.com/jobs/view/4011218533/",
							Date:             mustDate("2024-11-04"),
						},
					},
				},
				domain.Clojure: {},
				domain.Haskell: {},
			},
			ShortDescription: "",
			DealroomURL:      "",
			CrunchbaseURL:    "",
			PitchbookURL:     "",
			YahooFinanceURL:  "",
			GoogleFinanceURL: "",
			YCombinatorURL:   "",
			Industries:       []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Kantox",
			Website: "https://www.kantox.com/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                1871617,
				Alias:             "kantox",
				Name:              "Kantox",
				Followers:         "39K",
				Employees:         "200+",
				AssociatedMembers: "291",
				Verified:          true,
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
				domain.Go:    {},
				domain.Rust:  {},
				domain.Zig:   {},
				domain.Scala: {},
				domain.Elixir: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Elixir Software Engineer",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4072221760/",
							Date:             mustDate("2024-10-14"),
						},
					},
				},
				domain.Clojure: {},
				domain.Haskell: {},
			},
			ShortDescription: "",
			DealroomURL:      "",
			CrunchbaseURL:    "",
			PitchbookURL:     "",
			YahooFinanceURL:  "",
			GoogleFinanceURL: "",
			YCombinatorURL:   "",
			Industries: []domain.Industry{
				domain.IndustryFinTech,
			},
			HasEmployeesFromCountries: []domain.Country{},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "SmartRent",
			Website: "https://smartrent.com/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                16165131,
				Alias:             "smartrentdotcom",
				Name:              "SmartRent",
				Followers:         "16K",
				Employees:         "500+",
				AssociatedMembers: "434",
				Verified:          true,
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
				domain.Go:    {},
				domain.Rust:  {},
				domain.Zig:   {},
				domain.Scala: {},
				domain.Elixir: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Senior Software Engineer (Elixir)",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4071960242/",
							Date:             mustDate("2024-11-08"),
						},
					},
				},
				domain.Clojure: {},
				domain.Haskell: {},
			},
			ShortDescription:          "Home automation solutions for property managers and renters",
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
			Name:    "Denim",
			Website: "https://www.denim.com/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                11026766,
				Alias:             "join-denim",
				Name:              "Denim",
				Followers:         "7K",
				Employees:         "50+",
				AssociatedMembers: "203",
				Verified:          true,
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
				domain.Go:    {},
				domain.Rust:  {},
				domain.Zig:   {},
				domain.Scala: {},
				domain.Elixir: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Senior Software Engineer",
							ShortDescription: "Deep experience building large applications with Elixir",
							URL:              "https://www.linkedin.com/jobs/view/4064330407/",
							Date:             mustDate("2024-11-04"),
						},
					},
				},
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
			Name:    "Thryve.Earth",
			Website: "https://www.thryve.earth/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                82124834,
				Alias:             "thryve-earth",
				Name:              "Thryve.Earth",
				Followers:         "11K",
				Employees:         "2+",
				AssociatedMembers: "20",
				Verified:          false,
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
				domain.Go:    {},
				domain.Rust:  {},
				domain.Zig:   {},
				domain.Scala: {},
				domain.Elixir: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Full-Stack Developer (Elixir Phoenix)",
							ShortDescription: "Develop and maintain web applications using Elixir",
							URL:              "https://www.linkedin.com/jobs/view/4058317241/",
							Date:             mustDate("2024-10-28"),
						},
					},
				},
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
			Name:    "Zubale",
			Website: "https://www.zubale.com/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                33183626,
				Alias:             "zubale",
				Name:              "Zubale",
				Followers:         "104K",
				Employees:         "500+",
				AssociatedMembers: "1,152",
				Verified:          true,
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
				domain.Go:    {},
				domain.Rust:  {},
				domain.Zig:   {},
				domain.Scala: {},
				domain.Elixir: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Elixir/Scala/Rust Backend Engineer",
							ShortDescription: "More than 3 years of experience in Elixir",
							URL:              "https://www.linkedin.com/jobs/view/4076675943/",
							Date:             mustDate("2024-11-14"),
						},
					},
				},
				domain.Clojure: {},
				domain.Haskell: {},
			},
			ShortDescription: "",
			DealroomURL:      "",
			CrunchbaseURL:    "",
			PitchbookURL:     "",
			YahooFinanceURL:  "",
			GoogleFinanceURL: "",
			YCombinatorURL:   "",
			Industries: []domain.Industry{
				domain.IndustryECommerce,
			},
			HasEmployeesFromCountries: []domain.Country{},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Priceline",
			Website: "https://www.priceline.com/",
			Careers: "https://careers.priceline.com/",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                7451,
				Alias:             "priceline-com",
				Name:              "Priceline",
				Followers:         "72K",
				Employees:         "1K+",
				AssociatedMembers: "1,859",
				Verified:          true,
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
				domain.Go: {},
				domain.Rust: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Senior Software Engineer (Rust)",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4072949018/",
							Date:             mustDate("2024-11-14"),
						},
					},
				},
				domain.Zig:     {},
				domain.Scala:   {},
				domain.Elixir:  {},
				domain.Clojure: {},
				domain.Haskell: {},
			},
			ShortDescription: "",
			DealroomURL:      "",
			CrunchbaseURL:    "",
			PitchbookURL:     "",
			YahooFinanceURL:  "",
			GoogleFinanceURL: "",
			YCombinatorURL:   "",
			Industries:       []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "New Relic",
			Website: "https://newrelic.com/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                426253,
				Alias:             "new-relic-inc-",
				Name:              "New Relic",
				Followers:         "148K",
				Employees:         "1K+",
				AssociatedMembers: "2,926",
				Verified:          true,
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
							Title:            "Lead Software Engineer — Golang",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4058040204/",
							Date:             mustDate("2024-11-14"),
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
			ShortDescription: "Data for Engineers",
			DealroomURL:      "",
			CrunchbaseURL:    "",
			PitchbookURL:     "",
			YahooFinanceURL:  "",
			GoogleFinanceURL: "",
			YCombinatorURL:   "",
			Industries:       []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
				domain.Czechia,
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Riskified",
			Website: "https://www.riskified.com/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                2821604,
				Alias:             "riskified",
				Name:              "Riskified",
				Followers:         "53K",
				Employees:         "500+",
				AssociatedMembers: "720",
				Verified:          true,
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
				domain.Go:   {},
				domain.Rust: {},
				domain.Zig:  {},
				domain.Scala: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Scala Developer — App Infra Team",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4047670627/",
							Date:             mustDate("2024-11-14"),
						},
					},
				},
				domain.Elixir:  {},
				domain.Clojure: {},
				domain.Haskell: {},
			},
			ShortDescription: "",
			DealroomURL:      "",
			CrunchbaseURL:    "",
			PitchbookURL:     "",
			YahooFinanceURL:  "",
			GoogleFinanceURL: "",
			YCombinatorURL:   "",
			Industries:       []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
			},
		},
		{
			ID:      0,                         // system
			Type:    domain.CompanyTypeStartup, // system
			Name:    "Omilia",
			Website: "https://omilia.com/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                1057955,
				Alias:             "omilia-ltd",
				Name:              "Omilia - Conversational Intelligence",
				Followers:         "21K",
				Employees:         "200+",
				AssociatedMembers: "384",
				Verified:          true,
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
							Title:            "Golang Engineer",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4071728951/",
							Date:             mustDate("2024-11-14"),
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
			DealroomURL:      "",
			CrunchbaseURL:    "",
			PitchbookURL:     "",
			YahooFinanceURL:  "",
			GoogleFinanceURL: "",
			YCombinatorURL:   "",
			Industries:       []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
				domain.Czechia,
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Rakuten Viber",
			Website: "https://www.viber.com/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                1710293,
				Alias:             "rakuten-viber",
				Name:              "Rakuten Viber",
				Followers:         "41K",
				Employees:         "200+",
				AssociatedMembers: "710",
				Verified:          true,
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
							Title:            "Golang Developer",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4076696523/",
							Date:             mustDate("2024-11-14"),
						},
						{
							Title:            "Golang Developer",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4106769274/",
							Date:             mustDate("2024-12-20"),
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
			DealroomURL:      "",
			CrunchbaseURL:    "",
			PitchbookURL:     "",
			YahooFinanceURL:  "",
			GoogleFinanceURL: "",
			YCombinatorURL:   "",
			Industries:       []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Nebius",
			Website: "https://nebius.com/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                89802307,
				Alias:             "nebius",
				Name:              "Nebius",
				Followers:         "21K",
				Employees:         "200",
				AssociatedMembers: "499",
				Verified:          false,
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
							Title:            "Senior Backend Engineer (Go)",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4074905433/",
							Date:             mustDate("2024-11-14"),
						},
						{
							Title:            "Senior Backend Engineer (Go)",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4102045400/",
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
			ShortDescription: "Cloud platform specifically designed to train AI models",
			DealroomURL:      "",
			CrunchbaseURL:    "",
			PitchbookURL:     "",
			YahooFinanceURL:  "",
			GoogleFinanceURL: "",
			YCombinatorURL:   "",
			Industries:       []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Zoom",
			Website: "https://www.zoom.com/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                2532259,
				Alias:             "zoom",
				Name:              "Zoom",
				Followers:         "588K",
				Employees:         "5K+",
				AssociatedMembers: "11,159",
				Verified:          true,
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
							Title:            "Security Engineer (PHP Laravel & Go) — Workvivo",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4075517057/",
							Date:             mustDate("2024-11-14"),
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
			Name:    "LexisNexis Risk Solutions",
			Website: "https://risk.lexisnexis.com/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                10371941,
				Alias:             "lexisnexis-risk-solutions",
				Name:              "LexisNexis Risk Solutions",
				Followers:         "245K",
				Employees:         "10K+",
				AssociatedMembers: "8,927",
				Verified:          true,
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
							Title:            "Go Full Stack Developer",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4077459247/",
							Date:             mustDate("2024-11-17"),
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
			DealroomURL:      "",
			CrunchbaseURL:    "",
			PitchbookURL:     "",
			YahooFinanceURL:  "",
			GoogleFinanceURL: "",
			YCombinatorURL:   "",
			Industries:       []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
				domain.Czechia,
			},
		},
		{
			ID:      0,                         // system
			Type:    domain.CompanyTypeStartup, // system
			Name:    "SKIDOS",
			Website: "https://www.skidos.com/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                9246535,
				Alias:             "skidos-games",
				Name:              "SKIDOS",
				Followers:         "6K",
				Employees:         "10+",
				AssociatedMembers: "33",
				Verified:          false,
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
							Title:            "Golang Developer",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4075801293/",
							Date:             mustDate("2024-11-16"),
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
			DealroomURL:      "",
			CrunchbaseURL:    "",
			PitchbookURL:     "",
			YahooFinanceURL:  "",
			GoogleFinanceURL: "",
			YCombinatorURL:   "",
			Industries: []domain.Industry{
				domain.IndustryEdTech,
			},
			HasEmployeesFromCountries: []domain.Country{},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "SingleStore",
			Website: "https://www.singlestore.com/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                71851558,
				Alias:             "singlestore",
				Name:              "SingleStore",
				Followers:         "40K",
				Employees:         "200+",
				AssociatedMembers: "520",
				Verified:          false,
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
							Title:            "Software Engineer, Nova Platform (Go & K8s)",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4039359993/",
							Date:             mustDate("2024-11-16"),
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
			ShortDescription: "Cloud-native database",
			DealroomURL:      "",
			CrunchbaseURL:    "",
			PitchbookURL:     "",
			YahooFinanceURL:  "",
			GoogleFinanceURL: "",
			YCombinatorURL:   "",
			Industries:       []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
				domain.Czechia,
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Diabolocom",
			Website: "https://www.diabolocom.com/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                902343,
				Alias:             "diabolocom",
				Name:              "Diabolocom",
				Followers:         "24K",
				Employees:         "50+",
				AssociatedMembers: "119",
				Verified:          true,
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
							Title:            "Software Engineer (Golang)",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4059794979/",
							Date:             mustDate("2024-11-17"),
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
			ShortDescription:          "Contact Center",
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
			Name:    "FlexAI",
			Website: "https://www.flex.ai/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                100770131,
				Alias:             "flexaihq",
				Name:              "FlexAI",
				Followers:         "4K",
				Employees:         "10+",
				AssociatedMembers: "63",
				Verified:          false,
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
							Title:            "Backend Engineer — Golang (Senior)",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4059611387/",
							Date:             mustDate("2024-10-25"),
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
			DealroomURL:      "",
			CrunchbaseURL:    "",
			PitchbookURL:     "",
			YahooFinanceURL:  "",
			GoogleFinanceURL: "",
			YCombinatorURL:   "",
			Industries:       []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
			},
		},
		{
			ID:      0,                         // system
			Type:    domain.CompanyTypeStartup, // system
			Name:    "Stockbit",
			Website: "https://stockbit.com/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                6398473,
				Alias:             "stockbit",
				Name:              "Stockbit",
				Followers:         "31K",
				Employees:         "500+",
				AssociatedMembers: "454",
				Verified:          true,
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
							Title:            "Lead Backend Engineer (Golang)",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4058347229/",
							Date:             mustDate("2024-10-25"),
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
			Name:    "Teaching Strategies",
			Website: "https://teachingstrategies.com/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                56801,
				Alias:             "teaching-strategies-llc",
				Name:              "Teaching Strategies, LLC",
				Followers:         "98K",
				Employees:         "200+",
				AssociatedMembers: "1,061",
				Verified:          true,
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
							Title:            "Senior Software Engineer (Golang)",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4023232816/",
							Date:             mustDate("2024-10-25"),
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
			DealroomURL:      "",
			CrunchbaseURL:    "",
			PitchbookURL:     "",
			YahooFinanceURL:  "",
			GoogleFinanceURL: "",
			YCombinatorURL:   "",
			Industries: []domain.Industry{
				domain.IndustryEdTech,
			},
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Sophos",
			Website: "https://www.sophos.com/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                5053,
				Alias:             "sophos",
				Name:              "Sophos",
				Followers:         "481K",
				Employees:         "1K+",
				AssociatedMembers: "4,769",
				Verified:          true,
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
							Title:            "Senior Software Engineer — Golang Developer (Network security domain)",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4042181078/",
							Date:             mustDate("2024-11-15"),
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
			DealroomURL:      "",
			CrunchbaseURL:    "",
			PitchbookURL:     "",
			YahooFinanceURL:  "",
			GoogleFinanceURL: "",
			YCombinatorURL:   "",
			Industries: []domain.Industry{
				domain.IndustryCyberSecurity,
			},
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
				domain.Czechia,
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Macquarie Group",
			Website: "https://www.macquarie.com/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                3537,
				Alias:             "macquariegroup",
				Name:              "Macquarie Group",
				Followers:         "649K",
				Employees:         "10K+",
				AssociatedMembers: "20,814",
				Verified:          true,
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
							Title:            "Software Engineer — Python/Golang",
							ShortDescription: "We use Golang and Python",
							URL:              "https://www.linkedin.com/jobs/view/4022294274/",
							Date:             mustDate("2024-10-25"),
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
			DealroomURL:      "",
			CrunchbaseURL:    "",
			PitchbookURL:     "",
			YahooFinanceURL:  "",
			GoogleFinanceURL: "",
			YCombinatorURL:   "",
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
			Name:    "Ayoconnect",
			Website: "https://www.ayoconnect.com/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                7598856,
				Alias:             "ayoconnect",
				Name:              "Ayoconnect",
				Followers:         "71K",
				Employees:         "200+",
				AssociatedMembers: "167",
				Verified:          false,
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
							Title:            "Junior Software Engineer (Golang)",
							ShortDescription: "Design, develop, and maintain software applications using Golang",
							URL:              "https://www.linkedin.com/jobs/view/4057794419/",
							Date:             mustDate("2024-10-25"),
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
			ShortDescription: "Payments and Banking Enablement for Indonesia",
			DealroomURL:      "",
			CrunchbaseURL:    "",
			PitchbookURL:     "",
			YahooFinanceURL:  "",
			GoogleFinanceURL: "",
			YCombinatorURL:   "",
			Industries: []domain.Industry{
				domain.IndustryFinTech,
			},
			HasEmployeesFromCountries: []domain.Country{},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Camping World",
			Website: "https://www.campingworld.com/",
			Careers: "https://www.campingworldcareers.com/",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                23441,
				Alias:             "campingworld",
				Name:              "Camping World",
				Followers:         "85K",
				Employees:         "10K+",
				AssociatedMembers: "6,875",
				Verified:          true,
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
							Title:            "Senior Back End Engineer (TypeScript/Golang)",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4056215700/",
							Date:             mustDate("2024-10-25"),
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
			Name:    "Shopee",
			Website: "https://shopee.com/",
			Careers: "https://careers.shopee.sg/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                6451760,
				Alias:             "shopee",
				Name:              "Shopee",
				Followers:         "2M",
				Employees:         "5K+",
				AssociatedMembers: "47,126",
				Verified:          true,
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Engineer, Backend (Python/ Golang)",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4045101256/",
							Date:             mustDate("2024-11-16"),
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
			ShortDescription: "E-commerce platform",
			Industries: []domain.Industry{
				domain.IndustryECommerce,
			},
			HasEmployeesFromCountries: []domain.Country{},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Guardsquare",
			Website: "https://www.guardsquare.com/",
			Careers: "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                5012731,
				Alias:             "guardsquare",
				Name:              "Guardsquare",
				Followers:         "7K",
				Employees:         "50+",
				AssociatedMembers: "163",
				Verified:          true,
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Backend Engineer (Go)",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/3777868160/",
							Date:             mustDate("2024-11-14"),
						},
						{
							Title:            "Web Engineer (GO/React)",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4020630528/",
							Date:             mustDate("2024-12-18"),
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
			ShortDescription: "Mobile application protection",
			Industries: []domain.Industry{
				domain.IndustryCyberSecurity,
			},
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
				domain.Czechia,
			},
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "JFrog",
			Website: "https://jfrog.com/",
			Careers: "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                455737,
				Alias:             "jfrog-ltd",
				Name:              "JFrog",
				Followers:         "66K",
				Employees:         "1K+",
				AssociatedMembers: "1,958",
				Verified:          true,
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Senior Golang Developer",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4077432516/",
							Date:             mustDate("2024-11-16"),
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
			ShortDescription: "Software supply chain platform",
			Industries: []domain.Industry{
				domain.IndustryDevOps,
			},
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "HP",
			Website: "https://www.hp.com/",
			Careers: "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                5390798,
				Alias:             "hp",
				Name:              "HP",
				Followers:         "5M",
				Employees:         "10K+",
				AssociatedMembers: "122,744",
				Verified:          true,
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Golang Developer",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4078297154/",
							Date:             mustDate("2024-11-17"),
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
				domain.Czechia,
			},
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Mirantis",
			Website: "https://www.mirantis.com/",
			Careers: "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                7793,
				Alias:             "mirantis",
				Name:              "Mirantis",
				Followers:         "70K",
				Employees:         "500+",
				AssociatedMembers: "589",
				Verified:          true,
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Senior Software Engineer (Golang/Kubernetes)",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4075754584/",
							Date:             mustDate("2024-11-14"),
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
			ShortDescription: "Open source infrastructure for containers and virtual machines",
			Industries: []domain.Industry{
				domain.IndustryDevOps,
			},
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
				domain.Czechia,
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Zscaler",
			Website: "https://www.zscaler.com/",
			Careers: "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                234625,
				Alias:             "zscaler",
				Name:              "Zscaler",
				Followers:         "387K",
				Employees:         "5K+",
				AssociatedMembers: "8,436",
				Verified:          true,
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Staff Software Engineer, Data Forwarding (Golang)",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4077097120/",
							Date:             mustDate("2024-11-15"),
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
				domain.IndustryCyberSecurity,
				domain.IndustryDevOps,
			},
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Pure Storage",
			Website: "https://www.purestorage.com/",
			Careers: "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                1632202,
				Alias:             "purestorage",
				Name:              "Pure Storage",
				Followers:         "493K",
				Employees:         "5K",
				AssociatedMembers: "6,341",
				Verified:          true,
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Golang Engineer, FlashArray",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4072597478/",
							Date:             mustDate("2024-11-13"),
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
			ShortDescription: "Data storage",
			Industries:       []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
				domain.Czechia,
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Teya",
			Website: "https://www.teya.com/",
			Careers: "https://www.teya.com/careers",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                93621491,
				Alias:             "teya-global",
				Name:              "Teya",
				Followers:         "65K",
				Employees:         "1K+",
				AssociatedMembers: "735",
				Verified:          true,
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Senior Software Engineer (Golang)",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4073753061/",
							Date:             mustDate("2024-11-14"),
						},
						{
							Title:            "Senior Software Engineer (Golang)",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4107304437/",
							Date:             mustDate("2024-12-24"),
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
			Name:    "TomTom",
			Website: "https://www.tomtom.com/",
			Careers: "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                166474,
				Alias:             "tomtom",
				Name:              "TomTom",
				Followers:         "177K",
				Employees:         "1K+",
				AssociatedMembers: "5,196",
				Verified:          true,
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Full-Stack Developer (GO)",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4075571548/",
							Date:             mustDate("2024-11-14"),
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
			ShortDescription: "Mapping and location technology",
			Industries:       []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
				domain.Czechia,
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Postscript",
			Website: "https://postscript.io/",
			Careers: "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                18771172,
				Alias:             "postscriptio",
				Name:              "Postscript",
				Followers:         "15K",
				Employees:         "200+",
				AssociatedMembers: "295",
				Verified:          true,
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Staff Golang Software Engineer",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4071761682/",
							Date:             mustDate("2024-11-19"),
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
			ShortDescription:          "SMS marketing platform for ecommerce companies",
			Industries:                []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Impossible Cloud",
			Website: "https://www.impossiblecloud.com/",
			Careers: "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                76137251,
				Alias:             "impossiblecloud",
				Name:              "Impossible Cloud",
				Followers:         "3K",
				Employees:         "50+",
				AssociatedMembers: "51",
				Verified:          false,
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Junior Engineer — Golang & Distributed Systems",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4079088056/",
							Date:             mustDate("2024-11-19"),
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
			ShortDescription: "Cloud storage",
			Industries:       []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Arista Networks",
			Website: "https://www.arista.com/",
			Careers: "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                80069,
				Alias:             "arista-networks-inc",
				Name:              "Arista Networks",
				Followers:         "414K",
				Employees:         "1K+",
				AssociatedMembers: "4,697",
				Verified:          true,
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Software Engineer (Java/Golang)",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4077123852/",
							Date:             mustDate("2024-11-19"),
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
			ShortDescription: "Data-Driven Networking",
			Industries:       []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
				domain.Czechia,
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "PaveNow",
			Website: "https://www.pavenow.io/",
			Careers: "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                96200958,
				Alias:             "pavenow",
				Name:              "PaveNow",
				Followers:         "1K",
				Employees:         "10+",
				AssociatedMembers: "15",
				Verified:          false,
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Senior Golang Developer",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4077178228/",
							Date:             mustDate("2024-11-19"),
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
			ShortDescription: "Financial management for small and medium-sized businesses",
			Industries: []domain.Industry{
				domain.IndustryFinTech,
			},
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "LemFi",
			Website: "https://lemfi.com/",
			Careers: "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                68838211,
				Alias:             "lemfi",
				Name:              "LemFi",
				Followers:         "12K",
				Employees:         "10+",
				AssociatedMembers: "239",
				Verified:          false,
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Senior Backend Engineer — Golang",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4080247545/",
							Date:             mustDate("2024-11-19"),
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
