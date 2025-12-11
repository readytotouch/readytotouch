package db

import "github.com/readytotouch/readytotouch/internal/domain"

func companies17Shard() []domain.CompanyProfile {
	return []domain.CompanyProfile{
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Softdrive",
			Website: "https://www.softdrive.co/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                30148489,
				IDs:               nil,
				Alias:             "softdrivecloud",
				Name:              "Softdrive",
				Followers:         "635",
				Employees:         "11-50",
				AssociatedMembers: "14",
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
							Title:                "Go Engineer",
							ShortDescription:     "Cloud Services",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4216375542/",
							Date:                 mustDate("2025-04-24"),
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
			ShortDescription: "Platform that enables companies to deploy cloud VDI solutions, providing users with a seamless desktop experience from anywhere",
			Industries:       []domain.Industry{
				// NOP
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Command Zero",
			Website: "https://www.cmdzero.io/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                79530176,
				IDs:               nil,
				Alias:             "command-zero",
				Name:              "Command Zero",
				Followers:         "5K",
				Employees:         "11-50",
				AssociatedMembers: "28",
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
							Title:                "Go Developer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4213378861/",
							Date:                 mustDate("2025-04-23"),
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
			ShortDescription: "Autonomous and AI-assisted cyber investigations platform",
			Industries: []domain.Industry{
				domain.IndustryCyberSecurity,
			},
		},
		{
			ID:      0,                         // system
			Type:    domain.CompanyTypeStartup, // system
			Name:    "CareCar",
			Website: "https://www.carecar.co/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                18311681,
				IDs:               nil,
				Alias:             "carecarco",
				Name:              "CareCar",
				Followers:         "654",
				Employees:         "11-50",
				AssociatedMembers: "49",
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
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-CareCar-EI_IE4156942.11,18.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/CareCar-Reviews-E4156942.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/CareCar-Jobs-E4156942.htm",
				Jobs:        "128",
				Reviews:     "32",
				Salaries:    "31",
				ReviewsRate: "2.7",
				Verified:    false,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "",
			},
			Languages: domain.Languages{
				domain.Go:    {},
				domain.Rust:  {},
				domain.Zig:   {},
				domain.Scala: {},
				domain.Elixir: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Senior Software Engineer",
							ShortDescription:     "Platform & APIs",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4216110820/",
							Date:                 mustDate("2025-04-28"),
							WithSalary:           false,
							Remote:               true,
						},
					},
				},
				domain.Clojure: {},
				domain.Haskell: {},
			},
			ShortDescription: "",
			Industries:       []domain.Industry{
				// NOP
			},
		},
		{
			ID:      0,                         // system
			Type:    domain.CompanyTypeStartup, // system
			Name:    "Evident",
			Website: "https://www.evidentid.com/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                17901517,
				IDs:               nil,
				Alias:             "evident-id-inc.",
				Name:              "Evident",
				Followers:         "5K",
				Employees:         "51-200",
				AssociatedMembers: "74",
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
				Alias:     "evident",
				Employees: "90",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Evident-ID-EI_IE1554019.11,21.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Evident-ID-Reviews-E1554019.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Evident-ID-Jobs-E1554019.htm",
				Jobs:        "7",
				Reviews:     "27",
				Salaries:    "38",
				ReviewsRate: "4.4",
				Verified:    true,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "",
			},
			Languages: domain.Languages{
				domain.Go:   {},
				domain.Rust: {},
				domain.Zig:  {},
				domain.Scala: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Software Engineer (Scala)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4213956380/",
							Date:                 mustDate("2025-04-25"),
							WithSalary:           false,
							Remote:               false,
						},
					},
				},
				domain.Elixir:  {},
				domain.Clojure: {},
				domain.Haskell: {},
			},
			ShortDescription: "Third-Party Risk Management for Critical Partners",
			Industries:       []domain.Industry{
				// NOP
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "FlexTrade",
			Website: "https://flextrade.com/",
			Careers: "https://careers.flextrade.com/",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                49161,
				IDs:               nil,
				Alias:             "flextrade",
				Name:              "FlexTrade",
				Followers:         "49K",
				Employees:         "501-1K",
				AssociatedMembers: "658",
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
				Alias:     "flextrade-systems",
				Employees: "720",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-FlexTrade-Systems-Inc-EI_IE197820.11,32.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/FlexTrade-Systems-Inc-Reviews-E197820.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/FlexTrade-Systems-Inc-Jobs-E197820.htm",
				Jobs:        "23",
				Reviews:     "405",
				Salaries:    "705",
				ReviewsRate: "3.9",
				Verified:    true,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "",
			},
			Languages: domain.Languages{
				domain.Go:   {},
				domain.Rust: {},
				domain.Zig:  {},
				domain.Scala: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Senior Software Developer (Java/Scala)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4216696150/",
							Date:                 mustDate("2025-04-28"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Software Developer (Java/Scala)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4263362376/",
							Date:                 mustDate("2025-07-09"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Software Engineer (Java/Scala)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4334468004/",
							Location:             "City Of London, England, United Kingdom",
							Date:                 mustDate("2025-11-06"),
							WithSalary:           false,
							Remote:               false,
						},
					},
				},
				domain.Elixir:  {},
				domain.Clojure: {},
				domain.Haskell: {},
			},
			ShortDescription: "Broker-neutral, execution and order management trading platforms for equities, foreign exchange, options, futures and fixed income",
			Industries: []domain.Industry{
				domain.IndustryFinTech,
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Epic Games",
			Website: "https://epicgames.com/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                19327,
				IDs:               nil,
				Alias:             "epic-games",
				Name:              "Epic Games",
				Followers:         "786K",
				Employees:         "1K-5K",
				AssociatedMembers: "10,955",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:     "epicgames",
				Followers: "41.9k",
				Verified:  false,
			},
			BlindProfile: domain.BlindProfile{
				Alias:       "epic-games",
				Employees:   "1,001 to 5,000",
				Salary:      "$70K ~ $270K a year",
				Reviews:     "79",
				ReviewsRate: "4.2",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "epic-games",
				Employees: "990",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Epic-Games-EI_IE266904.11,21.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Epic-Games-Reviews-E266904.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Epic-Games-Jobs-E266904.htm",
				Jobs:        "190",
				Reviews:     "620",
				Salaries:    "1.5K",
				ReviewsRate: "3.7",
				Verified:    false,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "Epic-Games",
			},
			Languages: domain.Languages{
				domain.Go: {},
				domain.Rust: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Senior Rust Software Engineer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4181595042/",
							Location:             "Greater Montreal Metropolitan Area",
							Date:                 mustDate("2025-07-21", "2025-07-01", "2025-06-09", "2025-05-19", "2025-04-26"),
							WithSalary:           false,
							Remote:               true,
						},
						{
							Title:                "Senior Rust Software Engineer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4169030739/",
							Location:             "Cary, NC",
							Date:                 mustDate("2025-07-30", "2025-07-07", "2025-06-17", "2025-05-26", "2025-05-05"),
							WithSalary:           false,
							Remote:               false,
						},
					},
				},
				domain.Zig:     {},
				domain.Scala:   {},
				domain.Elixir:  {},
				domain.Clojure: {},
				domain.Haskell: {},
			},
			ShortDescription: "Interactive entertainment company and provider of 3D engine technology",
			Industries: []domain.Industry{
				domain.IndustryEntertainment,
				domain.IndustryGameDev,
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Notchup",
			Website: "https://www.notchup.com/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                67190926,
				IDs:               nil,
				Alias:             "notchupteam",
				Name:              "Notchup",
				Followers:         "166K",
				Employees:         "11-50",
				AssociatedMembers: "56",
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
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Notchup-EI_IE5537415.11,18.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Notchup-Reviews-E5537415.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Notchup-Jobs-E5537415.htm",
				Jobs:        "",
				Reviews:     "31",
				Salaries:    "13",
				ReviewsRate: "4.6",
				Verified:    true,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "",
			},
			Languages: domain.Languages{
				domain.Go: {},
				domain.Rust: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Rust Developer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4214284221/",
							Date:                 mustDate("2025-04-25"),
							WithSalary:           false,
							Remote:               true,
						},
						{
							Title:                "Software Engineer – Rust",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4236581310/",
							Date:                 mustDate("2025-05-28"),
							WithSalary:           false,
							Remote:               true,
						},
					},
				},
				domain.Zig:     {},
				domain.Scala:   {},
				domain.Elixir:  {},
				domain.Clojure: {},
				domain.Haskell: {},
			},
			ShortDescription: "AI Co-Pilot for Engineering Managers & CTOs",
			Industries:       []domain.Industry{
				// NOP
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Entro Security",
			Website: "https://entro.security/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                94227659,
				IDs:               nil,
				Alias:             "entro-security-non-human-identities",
				Name:              "Entro Security",
				Followers:         "5K",
				Employees:         "11-50",
				AssociatedMembers: "50",
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
							Title:                "Senior Backend Software Engineer – Golang",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4216060314/",
							Date:                 mustDate("2025-04-28"),
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
			ShortDescription: "Entro is the platform that supports a comprehensive Non-human Identity inventory that tracks and enriches exposed secrets, and the only platform that supports NHIDR (Non Human Identity Detection and Response)",
			Industries: []domain.Industry{
				domain.IndustryCyberSecurity,
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Kiddom",
			Website: "https://www.kiddom.co/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                3257139,
				IDs:               nil,
				Alias:             "kiddom",
				Name:              "Kiddom",
				Followers:         "18K",
				Employees:         "51-200",
				AssociatedMembers: "177",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:     "kiddom",
				Followers: "14",
				Verified:  false,
			},
			BlindProfile: domain.BlindProfile{
				Alias:       "Kiddom",
				Employees:   "51 to 200",
				Salary:      "",
				Reviews:     "5",
				ReviewsRate: "3.6",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "kiddom",
				Employees: "150",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Kiddom-EI_IE1575311.11,17.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Kiddom-Reviews-E1575311.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Kiddom-Jobs-E1575311.htm",
				Jobs:        "22",
				Reviews:     "64",
				Salaries:    "91",
				ReviewsRate: "3.0",
				Verified:    true,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "",
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 6,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Staff Software Engineer",
							ShortDescription:     "Platform Team (Golang)",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4185799364/",
							Date:                 mustDate("2025-04-28"),
							WithSalary:           true,
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
				domain.IndustryEdTech,
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Discern",
			Website: "https://discern.io/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                74602590,
				IDs:               nil,
				Alias:             "discern-io",
				Name:              "Discern",
				Followers:         "1K",
				Employees:         "11-50",
				AssociatedMembers: "18",
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
							Title:                "Fullstack Golang Engineer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4219007414/",
							Date:                 mustDate("2025-04-28"),
							WithSalary:           true,
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
			ShortDescription: "SaaS Performance Data & Analytics",
			Industries:       []domain.Industry{
				// NOP
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Deepcoin",
			Website: "https://www.deepcoin.com/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                67699470,
				IDs:               nil,
				Alias:             "deepcoinpro",
				Name:              "Deepcoin",
				Followers:         "7K",
				Employees:         "51-200",
				AssociatedMembers: "68",
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
							Title:                "Software Engineer – Golang Developer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4216959368/",
							Date:                 mustDate("2025-04-26"),
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
			ShortDescription: "Cryptocurrency derivatives trading platform",
			Industries:       []domain.Industry{
				// NOP
			},
			Ignore: true, // Cryptocurrency
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "CoinsPaid",
			Website: "https://coinspaid.com/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                13045933,
				IDs:               nil,
				Alias:             "cryptoprocessing-com",
				Name:              "CoinsPaid",
				Followers:         "15K",
				Employees:         "201-500",
				AssociatedMembers: "156",
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
							Title:                "Golang Engineer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4195931463/",
							Date:                 mustDate("2025-04-26"),
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
			ShortDescription: "Сrypto-financial ecosystem",
			Industries:       []domain.Industry{
				// NOP
			},
			Ignore: true, // Cryptocurrency
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Peak",
			Website: "https://www.peak.com/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                1561316,
				IDs:               nil,
				Alias:             "peak",
				Name:              "Peak",
				Followers:         "101K",
				Employees:         "51-200",
				AssociatedMembers: "193",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:     "peak",
				Followers: "26",
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
					GitHubRepositoriesCount: 27,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Big Data Platform Engineer – Java/Go",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4214247521/",
							Date:                 mustDate("2025-04-26"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Big Data Platform Engineer – Java/Go",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4253399494/",
							Location:             "Istanbul, Türkiye",
							Date:                 mustDate("2025-08-22", "2025-07-23", "2025-06-24"),
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
				domain.IndustryGameDev,
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Essity",
			Website: "https://www.essity.com/",
			Careers: "https://www.essity.com/Careers/",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                16240930,
				IDs:               []int{10676355, 16240930, 70954129},
				Alias:             "essity",
				Name:              "Essity",
				Followers:         "379K",
				Employees:         "10K+",
				AssociatedMembers: "18,334",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "",
				Verified: false,
			},
			BlindProfile: domain.BlindProfile{
				Alias:       "Essity",
				Employees:   "10,000+",
				Salary:      "",
				Reviews:     "2",
				ReviewsRate: "4",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias: "",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Essity-EI_IE1685428.11,17.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Essity-Reviews-E1685428.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Essity-Jobs-E1685428.htm",
				Jobs:        "333",
				Reviews:     "1.4K",
				Salaries:    "2.2K",
				ReviewsRate: "4.0",
				Verified:    false,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "",
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Go Engineer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4216684863/",
							Date:                 mustDate("2025-04-28"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Go Engineer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4236895751/",
							Date:                 mustDate("2025-05-28"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Graduate Go Engineer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4280052927/",
							Location:             "Rožňava, Kosice, Slovakia",
							Date:                 mustDate("2025-08-26", "2025-08-06"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Graduate Go Engineer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4309646516/",
							Location:             "Rožňava, Kosice, Slovakia",
							Date:                 mustDate("2025-10-03"),
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
			ShortDescription: "Hygiene and health company",
			Industries: []domain.Industry{
				domain.IndustryHealthTech,
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Pixellot",
			Website: "https://pixellot.tv/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                5107047,
				IDs:               nil,
				Alias:             "pixellotltd",
				Name:              "Pixellot - AI-Automated Sports Video and Analytics",
				Followers:         "22K",
				Employees:         "201-500",
				AssociatedMembers: "168",
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
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Pixellot-EI_IE1885122.11,19.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Pixellot-Reviews-E1885122.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Pixellot-Jobs-E1885122.htm",
				Jobs:        "14",
				Reviews:     "38",
				Salaries:    "47",
				ReviewsRate: "4.0",
				Verified:    false,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "",
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Backend Developer – Go",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4218604670/",
							Date:                 mustDate("2025-04-28"),
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
			ShortDescription: "AI-based automatic video and analytics solutions for the sports market",
			Industries:       []domain.Industry{
				// NOP
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Ladder",
			Website: "https://www.ladderlife.com/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                6634049,
				IDs:               nil,
				Alias:             "ladderlife",
				Name:              "Ladder",
				Followers:         "7K",
				Employees:         "51-200",
				AssociatedMembers: "120",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:     "ladderlife",
				Followers: "15",
				Verified:  true,
			},
			BlindProfile: domain.BlindProfile{
				Alias:       "Ladder-Insurance",
				Employees:   "51 to 200",
				Salary:      "",
				Reviews:     "15",
				ReviewsRate: "3.7",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "ladder",
				Employees: "150",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Ladder-EI_IE1405582.11,17.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Ladder-Reviews-E1405582.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Ladder-Jobs-E1405582.htm",
				Jobs:        "",
				Reviews:     "31",
				Salaries:    "58",
				ReviewsRate: "4.5",
				Verified:    false,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "",
			},
			Languages: domain.Languages{
				domain.Go:     {},
				domain.Rust:   {},
				domain.Zig:    {},
				domain.Scala:  {},
				domain.Elixir: {},
				domain.Clojure: {
					GitHubRepositoriesCount: 10,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Software Engineer – Full Stack (Clojure)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4209496817/",
							Date:                 mustDate("2025-05-05"),
							WithSalary:           false,
							Remote:               false,
						},
					},
				},
				domain.Haskell: {},
			},
			ShortDescription: "Ladder combines financial and insurance expertise to make life insurance easily accessible to everyone",
			Industries: []domain.Industry{
				domain.IndustryInsurTech,
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Complear",
			Website: "http://complear.com/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                75877122,
				IDs:               nil,
				Alias:             "complear",
				Name:              "Complear",
				Followers:         "3K",
				Employees:         "11-50",
				AssociatedMembers: "20",
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
				domain.Go:    {},
				domain.Rust:  {},
				domain.Zig:   {},
				domain.Scala: {},
				domain.Elixir: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Full Stack Engineer (Elixir/Liveview/React)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4219705060/",
							Date:                 mustDate("2025-05-02"),
							WithSalary:           false,
							Remote:               true,
						},
					},
				},
				domain.Clojure: {},
				domain.Haskell: {},
			},
			ShortDescription: "",
			Industries: []domain.Industry{
				domain.IndustryMedTech,
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "PDQ",
			Website: "https://www.pdq.com/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                17961724,
				IDs:               nil,
				Alias:             "pdq-com",
				Name:              "PDQ",
				Followers:         "6K",
				Employees:         "201-500",
				AssociatedMembers: "270",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:     "pdq",
				Followers: "50",
				Verified:  false,
			},
			BlindProfile: domain.BlindProfile{
				Alias: "",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "pdq",
				Employees: "100",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-PDQ-com-EI_IE2113848.11,18.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/PDQ-com-Reviews-E2113848.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/PDQ-com-Jobs-E2113848.htm",
				Jobs:        "",
				Reviews:     "58",
				Salaries:    "107",
				ReviewsRate: "4.6",
				Verified:    true,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "",
			},
			Languages: domain.Languages{
				domain.Go: {},
				domain.Rust: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Software Engineer, Rust",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4236745907/",
							Date:                 mustDate("2025-05-25"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Senior Software Engineer, Rust",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4258297981/",
							Date:                 mustDate("2025-06-28"),
							WithSalary:           false,
							Remote:               true,
						},
						{
							Title:                "Senior Software Engineer, Rust",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4262312661/",
							Date:                 mustDate("2025-07-08"),
							WithSalary:           false,
							Remote:               true,
						},
					},
				},
				domain.Zig:   {},
				domain.Scala: {},
				domain.Elixir: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Software Engineer, Elixir",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4218978070/",
							Date:                 mustDate("2025-05-01"),
							WithSalary:           false,
							Remote:               true,
						},
						{
							Title:                "Software Engineer, Elixir",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4236751515/",
							Date:                 mustDate("2025-05-25"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Senior Software Engineer, Elixir",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4258305221/",
							Date:                 mustDate("2025-07-01"),
							WithSalary:           false,
							Remote:               true,
						},
						{
							Title:                "Senior Software Engineer, Elixir",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4308900826/",
							Location:             "Romania",
							Date:                 mustDate("2025-10-01"),
							WithSalary:           false,
							Remote:               true,
						},
					},
				},
				domain.Clojure: {},
				domain.Haskell: {},
			},
			ShortDescription: "PDQ is device management for sysadmins, by sysadmins, that's simple, secure, and quick",
			Industries:       []domain.Industry{
				// NOP
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "True Anomaly",
			Website: "https://www.trueanomaly.space/",
			Careers: "https://www.trueanomaly.space/careers",
			About:   "https://www.trueanomaly.space/about-us",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                80801253,
				IDs:               nil,
				Alias:             "true-anomaly",
				Name:              "True Anomaly",
				Followers:         "41K",
				Employees:         "51-200",
				AssociatedMembers: "173",
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
				Alias:     "true-anomaly",
				Employees: "126",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-True-Anomaly-EI_IE7642233.11,23.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/True-Anomaly-Reviews-E7642233.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/True-Anomaly-Jobs-E7642233.htm",
				Jobs:        "",
				Reviews:     "17",
				Salaries:    "22",
				ReviewsRate: "3.1",
				Verified:    true,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "",
			},
			Languages: domain.Languages{
				domain.Go:    {},
				domain.Rust:  {},
				domain.Zig:   {},
				domain.Scala: {},
				domain.Elixir: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Senior Elixir Software Engineer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4219694154/",
							Location:             "Greater Colorado Springs Area",
							Date:                 mustDate("2025-09-30", "2025-09-09", "2025-08-19", "2025-07-27", "2025-07-05", "2025-06-14", "2025-05-22", "2025-04-30"),
							WithSalary:           true, // $155k/yr - $225k/yr
							Remote:               true,
						},
						{
							Title:                "Staff Elixir Software Engineer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4273981223/",
							Location:             "Greater Colorado Springs Area",
							Date:                 mustDate("2025-07-23"),
							WithSalary:           true, // $175k/yr - $255k/yr
							Remote:               true,
						},
						{
							Title:                "Staff Elixir Software Engineer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4281934098/",
							Location:             "Greater Colorado Springs Area",
							Date:                 mustDate("2025-09-20", "2025-08-08"),
							WithSalary:           true, // $175k/yr - $255k/yr
							Remote:               true,
						},
						{
							Title:                "Elixir Software Engineer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4312933268/",
							Location:             "Greater Colorado Springs Area",
							Date:                 mustDate("2025-10-11"),
							WithSalary:           true, // $110k/yr - $175k/yr
							Remote:               true,
						},
						{
							Title:                "Elixir Software Engineer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4334670217/",
							Location:             "Long Beach, CA",
							Date:                 mustDate("2025-11-29", "2025-11-06"),
							WithSalary:           true, // $90k/yr - $125k/yr
							Remote:               false,
						},
						{
							Title:                "Software Engineer II, Elixir",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4342567607/",
							Location:             "Long Beach, CA",
							Date:                 mustDate("2025-12-10"),
							WithSalary:           true, // $110k/yr - $160k/yr
							Remote:               false,
						},
					},
				},
				domain.Clojure: {},
				domain.Haskell: {},
			},
			ShortDescription: "",
			Industries:       []domain.Industry{
				// SpaceTech
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Scania Group",
			Website: "https://www.scania.com/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                3941,
				IDs:               []int{3941, 211779, 5351375, 5638270, 6040198, 6170020, 7197193, 9483639, 9630853, 10045357, 10057007, 10144562, 10304065, 10537576, 10538624, 10654301, 10853735, 10953382, 11099574, 11260008, 11535054, 11541089, 11683238, 11780674, 12615823, 14620963, 15081900, 15260433, 16241723, 18004389, 18307197, 18399387, 20487712, 24600913, 27122978, 27172150, 33184014, 36980195, 37854900, 70259502, 70411636, 71062516, 87224261, 90405534, 92822034, 92942981},
				Alias:             "scania",
				Name:              "Scania Group",
				Followers:         "627K",
				Employees:         "10K+",
				AssociatedMembers: "37,088",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:     "scania",
				Followers: "26",
				Verified:  true,
			},
			BlindProfile: domain.BlindProfile{
				Alias: "",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias: "",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Scania-EI_IE6131.11,17.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Scania-Reviews-E6131.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Scania-Jobs-E6131.htm",
				Jobs:        "366",
				Reviews:     "2.2K",
				Salaries:    "3.6K",
				ReviewsRate: "4.1",
				Verified:    true,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "",
			},
			Languages: domain.Languages{
				domain.Go:   {},
				domain.Rust: {},
				domain.Zig:  {},
				domain.Scala: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Senior Developer, Scala",
							ShortDescription:     "Autonomous Trucks",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4221388268/",
							Date:                 mustDate("2025-05-03"),
							WithSalary:           false,
							Remote:               false,
						},
					},
				},
				domain.Elixir:  {},
				domain.Clojure: {},
				domain.Haskell: {},
			},
			ShortDescription: "Scania is a provider of transport solutions",
			Industries: []domain.Industry{
				domain.IndustryAutomotive,
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "OTIV",
			Website: "https://www.otiv.ai/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                23729264,
				IDs:               nil,
				Alias:             "otiv",
				Name:              "OTIV",
				Followers:         "6K",
				Employees:         "11-50",
				AssociatedMembers: "46",
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
				domain.Go: {},
				domain.Rust: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Full Stack Developer Rust",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4223904753/",
							Date:                 mustDate("2025-05-05"),
							WithSalary:           false,
							Remote:               false,
						},
					},
				},
				domain.Zig:     {},
				domain.Scala:   {},
				domain.Elixir:  {},
				domain.Clojure: {},
				domain.Haskell: {},
			},
			ShortDescription: "Remote and autonomous driving in rail",
			Industries:       []domain.Industry{
				// NOP
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "AlertMedia",
			Website: "https://www.alertmedia.com/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                3582832,
				IDs:               nil,
				Alias:             "alertmedia",
				Name:              "AlertMedia",
				Followers:         "13K",
				Employees:         "201-500",
				AssociatedMembers: "433",
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
				Alias:     "alertmedia",
				Employees: "200",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-AlertMedia-EI_IE986631.11,21.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/AlertMedia-Reviews-E986631.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/AlertMedia-Jobs-E986631.htm",
				Jobs:        "",
				Reviews:     "164",
				Salaries:    "303",
				ReviewsRate: "4.1",
				Verified:    true,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "",
			},
			Languages: domain.Languages{
				domain.Go:    {},
				domain.Rust:  {},
				domain.Zig:   {},
				domain.Scala: {},
				domain.Elixir: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Senior Software Engineer (Elixir)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4223001348/",
							Date:                 mustDate("2025-05-07"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Senior Software Engineer (Elixir)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4241241518/",
							Date:                 mustDate("2025-06-05"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Senior Software Engineer (Elixir)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4297023255/",
							Location:             "Mexico City, Mexico",
							Date:                 mustDate("2025-09-11"),
							WithSalary:           false,
							Remote:               false,
						},
					},
				},
				domain.Clojure: {},
				domain.Haskell: {},
			},
			ShortDescription: "Helping companies protect their people during emergencies with fast, reliable communication and threat intelligence",
			Industries: []domain.Industry{
				domain.IndustryMedTech,
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Fujitsu",
			Website: "https://www.fujitsu.com/",
			Careers: "https://www.fujitsu.com/about/careers/",
			About:   "https://www.fujitsu.com/about/",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                1374,
				IDs:               []int{1374, 1377},
				Alias:             "fujitsu",
				Name:              "Fujitsu",
				Followers:         "1M",
				Employees:         "10K+",
				AssociatedMembers: "60,726",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:     "fujitsu",
				Followers: "106",
				Verified:  false,
			},
			BlindProfile: domain.BlindProfile{
				Alias:       "Fujitsu-Global",
				Employees:   "10,000+",
				Salary:      "$48K ~ $205K a year",
				Reviews:     "37",
				ReviewsRate: "3.1",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "fujitsu",
				Employees: "47,180",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Fujitsu-EI_IE3524.11,18.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Fujitsu-Reviews-E3524.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Fujitsu-Jobs-E3524.htm",
				Jobs:        "624",
				Reviews:     "7.6K",
				Salaries:    "25K",
				ReviewsRate: "3.6",
				Verified:    true,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "",
			},
			Languages: domain.Languages{
				domain.Go:   {},
				domain.Rust: {},
				domain.Zig:  {},
				domain.Scala: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Software Engineer ( Scala )",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4224532915/",
							Date:                 mustDate("2025-05-07"),
							WithSalary:           false,
							Remote:               false,
						},
					},
				},
				domain.Elixir:  {},
				domain.Clojure: {},
				domain.Haskell: {},
			},
			ShortDescription: "",
			Industries:       []domain.Industry{
				// NOP
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "OneDoc",
			Website: "https://info.onedoc.ch/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                15074341,
				IDs:               nil,
				Alias:             "onedoc.ch",
				Name:              "OneDoc",
				Followers:         "6K",
				Employees:         "11-50",
				AssociatedMembers: "54",
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
				domain.Go:   {},
				domain.Rust: {},
				domain.Zig:  {},
				domain.Scala: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Senior Software Engineer (Scala)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4225153194/",
							Date:                 mustDate("2025-05-07"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Senior Software Engineer (Scala)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4248020239/",
							Date:                 mustDate("2025-06-11"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Senior Software Engineer (Scala)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4262231132/",
							Date:                 mustDate("2025-07-04"),
							WithSalary:           false,
							Remote:               false,
						},
					},
				},
				domain.Elixir:  {},
				domain.Clojure: {},
				domain.Haskell: {},
			},
			ShortDescription: "Digital Health Platform in Switzerland",
			Industries:       []domain.Industry{
				// NOP
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Wallee Group",
			Website: "https://wallee.com/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                37869764,
				IDs:               nil,
				Alias:             "wallee-group",
				Name:              "Wallee Group",
				Followers:         "9K",
				Employees:         "51-200",
				AssociatedMembers: "99",
				Verified:          false,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:     "wallee-payment",
				Followers: "15",
				Verified:  false,
			},
			BlindProfile: domain.BlindProfile{
				Alias: "",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias: "",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Wallee-EI_IE4604933.11,17.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Wallee-Reviews-E4604933.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Wallee-Jobs-E4604933.htm",
				Jobs:        "3",
				Reviews:     "6",
				Salaries:    "20",
				ReviewsRate: "2.9",
				Verified:    false,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "",
			},
			Languages: domain.Languages{
				domain.Go:   {},
				domain.Rust: {},
				domain.Zig:  {},
				domain.Scala: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Scala Developer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4220863785/",
							Date:                 mustDate("2025-05-05"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Scala Developer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4226130637/",
							Date:                 mustDate("2025-05-09"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Scala Developer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4244541596/",
							Date:                 mustDate("2025-06-05"),
							WithSalary:           false,
							Remote:               false,
						},
					},
				},
				domain.Elixir:  {},
				domain.Clojure: {},
				domain.Haskell: {},
			},
			ShortDescription: "Omnichannel payment service provider",
			Industries: []domain.Industry{
				domain.IndustryFinTech,
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "LayerZero Labs",
			Website: "https://layerzero.network/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                75654012,
				IDs:               nil,
				Alias:             "layerzerolabs",
				Name:              "LayerZero Labs",
				Followers:         "10K",
				Employees:         "51-200",
				AssociatedMembers: "131",
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
				domain.Go: {},
				domain.Rust: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Rust Engineer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4208833223/",
							Date:                 mustDate("2025-05-07"),
							WithSalary:           false,
							Remote:               false,
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
			Industries:       []domain.Industry{
				// NOP
			},
			Ignore: true, // Blockchain
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "OpenZeppelin",
			Website: "https://www.openzeppelin.com/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                6624089,
				IDs:               nil,
				Alias:             "openzeppelin",
				Name:              "OpenZeppelin",
				Followers:         "12K",
				Employees:         "51-200",
				AssociatedMembers: "116",
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
				domain.Go: {},
				domain.Rust: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Open Source Developer – Rust",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4222275916/",
							Date:                 mustDate("2025-05-07"),
							WithSalary:           false,
							Remote:               true,
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
			Industries:       []domain.Industry{
				// NOP
			},
			Ignore: true, // Blockchain
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Caesars Entertainment",
			Website: "https://www.caesars.com/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                4210,
				IDs:               []int{4210, 11272, 114488, 676238, 803567, 883365, 946384, 2911245, 10168653, 54709651},
				Alias:             "caesars-entertainment-inc",
				Name:              "Caesars Entertainment",
				Followers:         "156K",
				Employees:         "10K+",
				AssociatedMembers: "17,856",
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
				domain.Go:   {},
				domain.Rust: {},
				domain.Zig:  {},
				domain.Scala: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Scala/Java Engineer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4226347126/",
							Date:                 mustDate("2025-05-09"),
							WithSalary:           false,
							Remote:               false,
						},
					},
				},
				domain.Elixir:  {},
				domain.Clojure: {},
				domain.Haskell: {},
			},
			ShortDescription: "",
			Industries:       []domain.Industry{
				// NOP
			},
			Ignore: true, // Casino
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "TensorZero",
			Website: "https://www.tensorzero.com/",
			Careers: "https://www.tensorzero.com/blog/",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                100773704,
				IDs:               nil,
				Alias:             "tensorzero",
				Name:              "TensorZero",
				Followers:         "493",
				Employees:         "2-10",
				AssociatedMembers: "3",
				Verified:          false,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:     "tensorzero",
				Followers: "119",
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
				domain.Go: {},
				domain.Rust: {
					GitHubRepositoriesCount: 1,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Founding Member of Technical Staff — Back-End / Systems Engineering",
							ShortDescription:     "Our model gateway, written in Rust, is the backbone of the project",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4051181311/",
							Date:                 mustDate("2024-11-11"),
							WithSalary:           true,
							Remote:               false,
						},
					},
				},
				domain.Zig:     {},
				domain.Scala:   {},
				domain.Elixir:  {},
				domain.Clojure: {},
				domain.Haskell: {},
			},
			ShortDescription: "TensorZero builds open-source infrastructure that creates a feedback loop for optimizing LLM applications — turning production data into smarter, faster, and cheaper models",
			Industries: []domain.Industry{
				domain.IndustryDevOps,
			},
		},
		{
			ID:      0,                         // system
			Type:    domain.CompanyTypeStartup, // system
			Name:    "Riza",
			Website: "https://riza.io/",
			Careers: "https://jobs.ashbyhq.com/riza",
			About:   "https://riza.io/company",
			Blog:    "https://riza.io/blog",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                96086280,
				IDs:               nil,
				Alias:             "riza-inc",
				Name:              "Riza",
				Followers:         "227",
				Employees:         "2-10",
				AssociatedMembers: "8",
				Verified:          false,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:     "riza-io",
				Followers: "26",
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
					GitHubRepositoriesCount: 5,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Backend Engineer",
							ShortDescription:     "You’ve been writing Go for at least three years",
							SwitchingOpportunity: "",
							URL:                  "https://jobs.ashbyhq.com/riza/0cd3a197-a25c-4145-b765-d6b8e518ea07",
							Date:                 mustDate("2025-04-16"),
							WithSalary:           true, // $150K – $225K • Offers Equity
							Remote:               false,
							// Google Cloud
							// Comfortable writing TypeScript / Interest in learning Rust
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
			ShortDescription: "AI writes code. Riza runs it.",
			Industries: []domain.Industry{
				domain.IndustryDevOps,
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Birdie",
			Website: "https://www.birdie.care/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                13441836,
				IDs:               nil,
				Alias:             "birdiecare",
				Name:              "Birdie",
				Followers:         "25K",
				Employees:         "51-200",
				AssociatedMembers: "160",
				Verified:          false,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:     "birdiecare",
				Followers: "19",
				Verified:  false,
			},
			BlindProfile: domain.BlindProfile{
				Alias: "",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias: "",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Birdie-United-Kingdom-EI_IE2424020.11,32.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Birdie-United-Kingdom-Reviews-E2424020.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Birdie-United-Kingdom-Jobs-E2424020.htm",
				Jobs:        "4.7K",
				Reviews:     "58",
				Salaries:    "191",
				ReviewsRate: "4.6",
				Verified:    true,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "",
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 2,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Golang Software Engineer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4224237706/",
							Date:                 mustDate("2025-05-10"),
							WithSalary:           true,
							Remote:               false,
						},
						{
							Title:                "Golang Software Engineer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4242646467/",
							Date:                 mustDate("2025-06-05"),
							WithSalary:           true,
							Remote:               false,
						},
						{
							Title:                "Senior Software Engineer (Go)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4299923099/",
							Location:             "London, England, United Kingdom",
							Date:                 mustDate("2025-10-08"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Senior Full Stack Software Engineer (TypeScript/Go)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4321464792/",
							Location:             "London, England, United Kingdom",
							Date:                 mustDate("2025-12-01"),
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
			ShortDescription: "All-in-one platform connects care management, rostering, finance, auditing, and workforce planning, surfacing insights and automating admin",
			Industries: []domain.Industry{
				domain.IndustryMedTech,
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Aikido Security",
			Website: "https://www.aikido.dev/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                88890150,
				IDs:               nil,
				Alias:             "aikido-security",
				Name:              "Aikido Security",
				Followers:         "11K",
				Employees:         "51-200",
				AssociatedMembers: "83",
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
							Title:                "Golang Engineer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4224288166/",
							Date:                 mustDate("2025-05-10"),
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
			ShortDescription: "Aikido is the get-it-done security platform for developers",
			Industries: []domain.Industry{
				domain.IndustryCyberSecurity,
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Notable Systems",
			Website: "https://notablesystems.com/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                18360848,
				IDs:               nil,
				Alias:             "notablesystems",
				Name:              "Notable Systems",
				Followers:         "1K",
				Employees:         "51-200",
				AssociatedMembers: "43",
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
							Title:                "Senior Software Engineer – Golang",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4225016951/",
							Date:                 mustDate("2025-05-10"),
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
			ShortDescription: "OCR",
			Industries:       []domain.Industry{
				// NOP
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "HAZA Foods, LLC",
			Website: "https://hazagroup.com/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                58681959,
				IDs:               nil,
				Alias:             "haza-foods-llc",
				Name:              "HAZA Foods, LLC",
				Followers:         "3K",
				Employees:         "10K+",
				AssociatedMembers: "336",
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
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-HAZA-Foods-EI_IE2359119.11,21.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/HAZA-Foods-Reviews-E2359119.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/HAZA-Foods-Jobs-E2359119.htm",
				Jobs:        "",
				Reviews:     "51",
				Salaries:    "80",
				ReviewsRate: "2.7",
				Verified:    false,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "",
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Software Engineer – Golang / Kubernetes / DevOps",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4227440633/",
							Date:                 mustDate("2025-05-10"),
							WithSalary:           true,
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
			ShortDescription: "Wendy's franchise group",
			Industries:       []domain.Industry{
				// NOP
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Proxidize",
			Website: "https://proxidize.com/",
			Careers: "https://proxidize.com/careers/",
			About:   "https://proxidize.com/about/",
			Blog:    "https://proxidize.com/blog/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                68523797,
				IDs:               nil,
				Alias:             "proxidize",
				Name:              "Proxidize",
				Followers:         "23K",
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
							Title:                "Senior Golang Developer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4223570303/",
							Date:                 mustDate("2025-05-09"),
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
			ShortDescription: "Proxy solutions",
			Industries:       []domain.Industry{
				// NOP
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Onfy",
			Website: "https://onfy.de/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                76573663,
				IDs:               nil,
				Alias:             "onfy",
				Name:              "Onfy",
				Followers:         "1K",
				Employees:         "11-50",
				AssociatedMembers: "27",
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
							Title:                "Senior Backend Developer (Go)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4227006777/",
							Date:                 mustDate("2025-05-10"),
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
			ShortDescription: "German pharmacy marketplace",
			Industries: []domain.Industry{
				domain.IndustryECommerce,
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Well",
			Website: "https://www.well.co/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                35679376,
				IDs:               nil,
				Alias:             "wellrewarded",
				Name:              "Well",
				Followers:         "7K",
				Employees:         "51-200",
				AssociatedMembers: "381",
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
				Alias:     "well",
				Employees: "540",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Well-NC-EI_IE3142517.11,18.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Well-NC-Reviews-E3142517.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Well-NC-Jobs-E3142517.htm",
				Jobs:        "9",
				Reviews:     "11",
				Salaries:    "27",
				ReviewsRate: "3.3",
				Verified:    true,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "",
			},
			Languages: domain.Languages{
				domain.Go:      {},
				domain.Rust:    {},
				domain.Zig:     {},
				domain.Scala:   {},
				domain.Elixir:  {},
				domain.Clojure: {},
				domain.Haskell: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Haskell Engineer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4214803492/",
							Date:                 mustDate("2025-04-24"),
							WithSalary:           false,
							Remote:               true,
						},
						{
							Title:                "Haskell Engineer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4248210417/",
							Date:                 mustDate("2025-06-11"),
							WithSalary:           false,
							Remote:               true,
						},
						{
							Title:                "Haskell Engineer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4279720066/",
							Location:             "Warsaw Metropolitan Area",
							Date:                 mustDate("2025-08-04"),
							WithSalary:           false,
							Remote:               true,
						},
					},
				},
			},
			ShortDescription: "A clinical-grade health partner, designed to fit any lifestyle",
			Industries: []domain.Industry{
				domain.IndustryHealthTech,
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Rewards Network",
			Website: "https://www.rewardsnetwork.com/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                10925,
				IDs:               nil,
				Alias:             "rewards-network",
				Name:              "Rewards Network",
				Followers:         "28K",
				Employees:         "201-500",
				AssociatedMembers: "710",
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
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Rewards-Network-EI_IE2216.11,26.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Rewards-Network-Reviews-E2216.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Rewards-Network-Jobs-E2216.htm",
				Jobs:        "25",
				Reviews:     "374",
				Salaries:    "438",
				ReviewsRate: "3.6",
				Verified:    true,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "",
			},
			Languages: domain.Languages{
				domain.Go:   {},
				domain.Rust: {},
				domain.Zig:  {},
				domain.Scala: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Senior Software Engineer — Scala",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4230030958/",
							Location:             "Chicago, IL",
							Date:                 mustDate("2025-10-12", "2025-07-19", "2025-06-06", "2025-05-15"),
							WithSalary:           false,
							Remote:               false,
						},
					},
				},
				domain.Elixir:  {},
				domain.Clojure: {},
				domain.Haskell: {},
			},
			ShortDescription: "Rewards Network is a fintech company providing marketing, loyalty rewards programs, and capital for the restaurant industry",
			Industries: []domain.Industry{
				domain.IndustryFinTech,
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Qdrant",
			Website: "https://qdrant.tech/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                77961893,
				IDs:               nil,
				Alias:             "qdrant",
				Name:              "Qdrant",
				Followers:         "37K",
				Employees:         "51-200",
				AssociatedMembers: "89",
				Verified:          false,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:     "qdrant",
				Followers: "1.2k",
				Verified:  true,
			},
			BlindProfile: domain.BlindProfile{
				Alias: "",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "qdrant",
				Employees: "30",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Qdrant-Solutions-EI_IE7897646.11,27.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Qdrant-Solutions-Reviews-E7897646.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Qdrant-Solutions-Jobs-E7897646.htm",
				Jobs:        "8",
				Reviews:     "5",
				Salaries:    "10",
				ReviewsRate: "3.5",
				Verified:    false,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "",
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 10,
				},
				domain.Rust: {
					GitHubRepositoriesCount: 21,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Core Rust Engineer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4229012568/",
							Date:                 mustDate("2025-05-15"),
							WithSalary:           false,
							Remote:               true,
						},
						{
							Title:                "Rust Services Developer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4278807446/",
							Location:             "Germany",
							Date:                 mustDate("2025-08-06"),
							WithSalary:           false,
							Remote:               true,
						},
						{
							Title:                "Core Rust Engineer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4294466283/",
							Location:             "Germany",
							Date:                 mustDate("2025-09-06"),
							WithSalary:           false,
							Remote:               true,
						},
					},
				},
				domain.Zig:     {},
				domain.Scala:   {},
				domain.Elixir:  {},
				domain.Clojure: {},
				domain.Haskell: {},
			},
			ShortDescription: "Massive-Scale AI Search Engine & Vector Database",
			Industries: []domain.Industry{
				domain.IndustryDevOps,
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Supra",
			Website: "http://www.supra.com/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                71652981,
				IDs:               nil,
				Alias:             "supraofficial",
				Name:              "Supra",
				Followers:         "8K",
				Employees:         "51-200",
				AssociatedMembers: "259",
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
				domain.Go: {},
				domain.Rust: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Senior Rust Compiler Engineer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4230759876/",
							Date:                 mustDate("2025-05-15"),
							WithSalary:           false,
							Remote:               true,
						},
					},
				},
				domain.Zig:     {},
				domain.Scala:   {},
				domain.Elixir:  {},
				domain.Clojure: {},
				domain.Haskell: {},
			},
			ShortDescription: "MultiVM Layer 1 built for Super dApps",
			Industries:       []domain.Industry{
				// NOP
			},
			Ignore: true, // Web3
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Swiss Life Asset Managers",
			Website: "https://www.swisslife-am.com/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                6605,
				IDs:               []int{6605, 610840, 1082449, 6842138, 10932185, 30126193, 66384503, 80740784},
				Alias:             "swiss-life-asset-management",
				Name:              "Swiss Life Asset Managers",
				Followers:         "31K",
				Employees:         "1K-5K",
				AssociatedMembers: "2,309",
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
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Swiss-Life-Asset-Management-EI_IE3015893.11,38.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Swiss-Life-Asset-Management-Reviews-E3015893.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Swiss-Life-Asset-Management-Jobs-E3015893.htm",
				Jobs:        "",
				Reviews:     "66",
				Salaries:    "142",
				ReviewsRate: "3.4",
				Verified:    false,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "",
			},
			Languages: domain.Languages{
				domain.Go: {},
				domain.Rust: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Full Stack Engineer (Rust/Svelte)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4211351867/",
							Date:                 mustDate("2025-05-15"),
							WithSalary:           false,
							Remote:               true,
						},
					},
				},
				domain.Zig:     {},
				domain.Scala:   {},
				domain.Elixir:  {},
				domain.Clojure: {},
				domain.Haskell: {},
			},
			ShortDescription: "Asset Manager offering a broad range of investment solutions in real estate, infrastructure and securities",
			Industries: []domain.Industry{
				domain.IndustryFinTech,
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Natuvion",
			Website: "https://www.natuvion.com/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                5089814,
				IDs:               nil,
				Alias:             "natuvion",
				Name:              "Natuvion",
				Followers:         "10K",
				Employees:         "201-500",
				AssociatedMembers: "384",
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
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Natuvion-EI_IE2303226.11,19.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Natuvion-Reviews-E2303226.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Natuvion-Jobs-E2303226.htm",
				Jobs:        "100",
				Reviews:     "18",
				Salaries:    "38",
				ReviewsRate: "4.0",
				Verified:    true,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "",
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Senior Go Developer",
							ShortDescription:     "Focus on gRPC & Enterprise applications",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4281091636/",
							Location:             "Berlin, Berlin, Germany",
							Date:                 mustDate("2025-09-04", "2025-08-11"),
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
				domain.Haskell: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "(Senior) Haskell Developer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4233269070/",
							Date:                 mustDate("2025-07-08", "2025-07-07", "2025-07-04", "2025-06-14", "2025-05-22"),
							WithSalary:           false,
							Remote:               false,
						},
					},
				},
			},
			ShortDescription: "Natuvion moves business-critical data and processes from one technological platform to another both smoothly and cost-efficiently",
			Industries:       []domain.Industry{
				// NOP
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "OANDA",
			Website: "https://www.oanda.com/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                28727,
				IDs:               nil,
				Alias:             "oanda",
				Name:              "OANDA",
				Followers:         "40K",
				Employees:         "201-500",
				AssociatedMembers: "759",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "",
				Verified: false,
			},
			BlindProfile: domain.BlindProfile{
				Alias:       "OANDA",
				Employees:   "201 to 500",
				Salary:      "",
				Reviews:     "12",
				ReviewsRate: "3.6",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "oanda",
				Employees: "630",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-OANDA-EI_IE100548.11,16.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/OANDA-Reviews-E100548.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/OANDA-Jobs-E100548.htm",
				Jobs:        "20",
				Reviews:     "235",
				Salaries:    "392",
				ReviewsRate: "3.6",
				Verified:    false,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "",
			},
			Languages: domain.Languages{
				domain.Go:   {},
				domain.Rust: {},
				domain.Zig:  {},
				domain.Scala: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Senior Software Engineer (Scala/Java)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4232755562/",
							Date:                 mustDate("2025-05-20"),
							WithSalary:           false,
							Remote:               false,
						},
					},
				},
				domain.Elixir:  {},
				domain.Clojure: {},
				domain.Haskell: {},
			},
			ShortDescription: "Multi-asset trading services",
			Industries: []domain.Industry{
				domain.IndustryFinTech,
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "the LEGO Group",
			Website: "https://www.lego.com/",
			Careers: "https://www.lego.com/careers",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                3724,
				IDs:               []int{3724, 683889},
				Alias:             "lego-group",
				Name:              "the LEGO Group",
				Followers:         "2M",
				Employees:         "10K+",
				AssociatedMembers: "19,541",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:     "LEGO",
				Followers: "933",
				Verified:  true,
			},
			BlindProfile: domain.BlindProfile{
				Alias:       "LEGO",
				Employees:   "10,000+",
				Salary:      "",
				Reviews:     "7",
				ReviewsRate: "4.0",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "lego",
				Employees: "19,000",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-the-LEGO-Group-EI_IE3944.11,25.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/the-LEGO-Group-Reviews-E3944.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/the-LEGO-Group-Jobs-E3944.htm",
				Jobs:        "474",
				Reviews:     "3.2K",
				Salaries:    "4.8K",
				ReviewsRate: "4.4",
				Verified:    true,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "",
			},
			Languages: domain.Languages{
				domain.Go:   {},
				domain.Rust: {},
				domain.Zig:  {},
				domain.Scala: {
					GitHubRepositoriesCount: 1,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Backend/Fullstack Engineer, Scala",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4231817942/",
							Date:                 mustDate("2025-05-21"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Backend/Fullstack Engineer, Scala",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4233764933/",
							Date:                 mustDate("2025-06-14"),
							WithSalary:           false,
							Remote:               false,
						},
					},
				},
				domain.Elixir:  {},
				domain.Clojure: {},
				domain.Haskell: {},
			},
			ShortDescription: "",
			Industries:       []domain.Industry{
				// NOP
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Deutsche Bank",
			Website: "https://db.com/",
			Careers: "https://careers.db.com/",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                1262,
				IDs:               nil,
				Alias:             "deutsche-bank",
				Name:              "Deutsche Bank",
				Followers:         "3M",
				Employees:         "10K+",
				AssociatedMembers: "75,214",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "",
				Verified: false,
			},
			BlindProfile: domain.BlindProfile{
				Alias:       "Deutsche-Bank",
				Employees:   "10,000+",
				Salary:      "$33K ~ $294K a year",
				Reviews:     "222",
				ReviewsRate: "3.3",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "deutsche-bank",
				Employees: "68,370",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Deutsche-Bank-EI_IE3150.11,24.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Deutsche-Bank-Reviews-E3150.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Deutsche-Bank-Jobs-E3150.htm",
				Jobs:        "1.9K",
				Reviews:     "15K",
				Salaries:    "29K",
				ReviewsRate: "3.8",
				Verified:    true,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "",
			},
			Languages: domain.Languages{
				domain.Go:   {},
				domain.Rust: {},
				domain.Zig:  {},
				domain.Scala: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Lead Scala Engineer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4219049156/",
							Location:             "Berlin, Berlin, Germany",
							Date:                 mustDate("2025-08-01", "2025-07-21", "2025-07-08", "2025-07-02", "2025-06-11", "2025-05-21"),
							WithSalary:           false,
							Remote:               false,
						},
					},
				},
				domain.Elixir:  {},
				domain.Clojure: {},
				domain.Haskell: {},
			},
			ShortDescription: "German bank",
			Industries: []domain.Industry{
				domain.IndustryFinTech,
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Tenstorrent",
			Website: "https://tenstorrent.com/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                10629072,
				IDs:               nil,
				Alias:             "tenstorrent-inc.",
				Name:              "Tenstorrent",
				Followers:         "57K",
				Employees:         "201-500",
				AssociatedMembers: "783",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:     "tenstorrent",
				Followers: "873",
				Verified:  false,
			},
			BlindProfile: domain.BlindProfile{
				Alias:       "Tenstorrent",
				Employees:   "201 to 500",
				Salary:      "",
				Reviews:     "21",
				ReviewsRate: "4.2",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "tenstorrent",
				Employees: "150",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Tenstorrent-EI_IE2921474.11,22.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Tenstorrent-Reviews-E2921474.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Tenstorrent-Jobs-E2921474.htm",
				Jobs:        "",
				Reviews:     "52",
				Salaries:    "112",
				ReviewsRate: "4.2",
				Verified:    false,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "",
			},
			Languages: domain.Languages{
				domain.Go: {},
				domain.Rust: {
					GitHubRepositoriesCount: 2,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Senior Rust Developer",
							ShortDescription:     "Platform Software",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4229279528/",
							Location:             "Portland, ON",
							Date:                 mustDate("2025-09-02", "2025-07-20", "2025-06-29", "2025-06-08", "2025-05-17"),
							WithSalary:           false,
							Remote:               true,
						},
					},
				},
				domain.Zig:     {},
				domain.Scala:   {},
				domain.Elixir:  {},
				domain.Clojure: {},
				domain.Haskell: {},
			},
			ShortDescription: "Tenstorrent is a company that builds computers for AI",
			Industries:       []domain.Industry{
				// NOP
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Promon",
			Website: "https://promon.io/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                679063,
				IDs:               nil,
				Alias:             "promon-as",
				Name:              "Promon",
				Followers:         "7K",
				Employees:         "51-200",
				AssociatedMembers: "129",
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
				domain.Go: {},
				domain.Rust: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Software Engineer (Java/C++/Rust)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4219347502/",
							Date:                 mustDate("2025-05-22"),
							WithSalary:           false,
							Remote:               true,
						},
					},
				},
				domain.Zig:     {},
				domain.Scala:   {},
				domain.Elixir:  {},
				domain.Clojure: {},
				domain.Haskell: {},
			},
			ShortDescription: "Safeguard app data, fight malware, and prevent app tampering",
			Industries: []domain.Industry{
				domain.IndustryCyberSecurity,
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "cside",
			Website: "https://cside.dev/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                101783504,
				IDs:               nil,
				Alias:             "csidedev",
				Name:              "cside",
				Followers:         "2K",
				Employees:         "11-50",
				AssociatedMembers: "18",
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
				domain.Go: {},
				domain.Rust: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Senior Engineer – Rust",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4233145765/",
							Date:                 mustDate("2025-05-20"),
							WithSalary:           false,
							Remote:               true,
						},
					},
				},
				domain.Zig:     {},
				domain.Scala:   {},
				domain.Elixir:  {},
				domain.Clojure: {},
				domain.Haskell: {},
			},
			ShortDescription: "Client-side security solution with a proxy solution, helping to block attacks before they reach your user",
			Industries: []domain.Industry{
				domain.IndustryCyberSecurity,
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Riot Games",
			Website: "https://www.riotgames.com/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                60870,
				IDs:               nil,
				Alias:             "riot-games",
				Name:              "Riot Games",
				Followers:         "1M",
				Employees:         "1K-5K",
				AssociatedMembers: "7,325",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:     "riotgames",
				Followers: "753",
				Verified:  false,
			},
			BlindProfile: domain.BlindProfile{
				Alias:       "Riot-Games",
				Employees:   "1,001 to 5,000",
				Salary:      "$72K ~ $299K a year",
				Reviews:     "216",
				ReviewsRate: "4.1",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "riot-games",
				Employees: "3,500",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Riot-Games-EI_IE247538.11,21.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Riot-Games-Reviews-E247538.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Riot-Games-Jobs-E247538.htm",
				Jobs:        "",
				Reviews:     "1.4K",
				Salaries:    "3.2K",
				ReviewsRate: "4.0",
				Verified:    true,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "",
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 2,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Senior Backend Software Engineer – (Golang/C++)",
							ShortDescription:     "Metagame Features",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4226950943/",
							Date:                 mustDate("2025-06-25", "2025-06-18", "2025-06-12", "2025-06-05", "2025-05-28", "2025-05-22"),
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
			ShortDescription: "Computer Games",
			Industries: []domain.Industry{
				domain.IndustryGameDev,
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Zillow",
			Website: "https://www.zillow.com/",
			Careers: "https://www.zillow.com/careers",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                13990,
				IDs:               []int{5693, 13990, 75801, 115652, 167785, 234303, 289839, 569792, 3776432, 19106063},
				Alias:             "zillow",
				Name:              "Zillow",
				Followers:         "455K",
				Employees:         "5K-10K",
				AssociatedMembers: "9,032",
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
				Alias:     "zillow",
				Employees: "8,070",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Zillow-EI_IE40802.11,17.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Zillow-Reviews-E40802.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Zillow-Jobs-E40802.htm",
				Jobs:        "88",
				Reviews:     "2.3K",
				Salaries:    "5.8K",
				ReviewsRate: "3.5",
				Verified:    true,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "",
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Senior Software Development Engineer, Platform (Golang)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4204768990/",
							Date:                 mustDate("2025-05-22"),
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
			ShortDescription: "Real Estate",
			Industries: []domain.Industry{
				domain.IndustryPropTech,
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Teamwork.com",
			Website: "https://www.teamwork.com/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                1042291,
				IDs:               nil,
				Alias:             "teamwork-com",
				Name:              "Teamwork.com",
				Followers:         "53K",
				Employees:         "201-500",
				AssociatedMembers: "535",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:     "teamwork",
				Followers: "56",
				Verified:  true,
			},
			BlindProfile: domain.BlindProfile{
				Alias: "",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "teamwork",
				Employees: "270",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Teamwork-EI_IE929006.11,19.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Teamwork-Reviews-E929006.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Teamwork-Jobs-E929006.htm",
				Jobs:        "4",
				Reviews:     "221",
				Salaries:    "591",
				ReviewsRate: "3.6",
				Verified:    true,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "",
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 35,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Software Engineer II (Golang)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4233902057/",
							Date:                 mustDate("2025-05-24"),
							WithSalary:           false,
							Remote:               true,
						},
						{
							Title:                "Software Engineer II (Golang)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4300487997/",
							Location:             "Poland",
							Date:                 mustDate("2025-09-16"),
							WithSalary:           false,
							Remote:               true,
						},
						{
							Title:                "Senior Golang Engineer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4300480968/",
							Location:             "Poland",
							Date:                 mustDate("2025-09-16"),
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
			ShortDescription: "Platform built for managing client projects, profitably",
			Industries:       []domain.Industry{
				// NOP
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "GlassFlow",
			Website: "https://www.glassflow.dev/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                100710256,
				IDs:               nil,
				Alias:             "glassflow-dev",
				Name:              "GlassFlow",
				Followers:         "1K",
				Employees:         "2-10",
				AssociatedMembers: "12",
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
							Title:                "Senior Backend Engineer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4235682115/",
							Date:                 mustDate("2025-05-28"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Senior Backend Engineer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4257648108/",
							Date:                 mustDate("2025-07-03"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Senior Backend Engineer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4274689059/",
							Location:             "Berlin, Berlin, Germany",
							Date:                 mustDate("2025-08-03"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Founding Engineer – Golang",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4291307652/",
							Location:             "Berlin, Berlin, Germany",
							Date:                 mustDate("2025-08-26"),
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
			ShortDescription: "Open-Source solution to deduplicate and join Kafka data streams for ClickHouse",
			Industries: []domain.Industry{
				domain.IndustryDevOps,
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Channable",
			Website: "https://www.channable.com/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                5396383,
				IDs:               []int{5353569, 5396383},
				Alias:             "channable",
				Name:              "Channable",
				Followers:         "22K",
				Employees:         "201-500",
				AssociatedMembers: "284",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:     "channable",
				Followers: "73",
				Verified:  true,
			},
			BlindProfile: domain.BlindProfile{
				Alias: "",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias: "",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Channable-EI_IE2933913.11,20.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Channable-Reviews-E2933913.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Channable-Jobs-E2933913.htm",
				Jobs:        "9",
				Reviews:     "71",
				Salaries:    "143",
				ReviewsRate: "4.3",
				Verified:    true,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "",
			},
			Languages: domain.Languages{
				domain.Go:      {},
				domain.Rust:    {},
				domain.Zig:     {},
				domain.Scala:   {},
				domain.Elixir:  {},
				domain.Clojure: {},
				domain.Haskell: {
					GitHubRepositoriesCount: 14,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Haskell Software Engineer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4242629004/",
							Location:             "Utrecht, Utrecht, Netherlands",
							Date:                 mustDate("2025-08-13", "2025-07-23", "2025-07-01", "2025-06-05"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Haskell Software Engineer",
							ShortDescription:     "Creatives Team",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4268705870/",
							Location:             "Utrecht, Utrecht, Netherlands",
							Date:                 mustDate("2025-08-30", "2025-08-09", "2025-07-18"),
							WithSalary:           false,
							Remote:               false,
						},
					},
				},
			},
			ShortDescription: "Channable is the all-in-one platform that provides the solutions you need for greater visibility, smarter ad campaigns, and more personalized marketing",
			Industries: []domain.Industry{
				domain.IndustryAdTech,
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Youth Inc.",
			Website: "https://www.youth.inc/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                102056923,
				IDs:               nil,
				Alias:             "youthincsports",
				Name:              "Youth Inc.",
				Followers:         "2K",
				Employees:         "11-50",
				AssociatedMembers: "27",
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
				domain.Go:    {},
				domain.Rust:  {},
				domain.Zig:   {},
				domain.Scala: {},
				domain.Elixir: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Senior Software Engineer – Elixir",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4241480550/",
							Date:                 mustDate("2025-06-01"),
							WithSalary:           false,
							Remote:               false,
						},
					},
				},
				domain.Clojure: {},
				domain.Haskell: {},
			},
			ShortDescription: "Digital media network and commerce marketplace focused exclusively on youth sports",
			Industries:       []domain.Industry{
				// NOP
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "ThreatConnect",
			Website: "https://threatconnect.com/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                1284838,
				IDs:               nil,
				Alias:             "threatconnect-inc",
				Name:              "ThreatConnect",
				Followers:         "27K",
				Employees:         "51-200",
				AssociatedMembers: "166",
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
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-ThreatConnect-EI_IE1148746.11,24.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/ThreatConnect-Reviews-E1148746.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/ThreatConnect-Jobs-E1148746.htm",
				Jobs:        "",
				Reviews:     "85",
				Salaries:    "122",
				ReviewsRate: "4.1",
				Verified:    true,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "",
			},
			Languages: domain.Languages{
				domain.Go:    {},
				domain.Rust:  {},
				domain.Zig:   {},
				domain.Scala: {},
				domain.Elixir: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Elixir Engineer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4240698572/",
							Date:                 mustDate("2025-05-31"),
							WithSalary:           false,
							Remote:               true,
						},
					},
				},
				domain.Clojure: {},
				domain.Haskell: {},
			},
			ShortDescription: "ThreatConnect enables threat intelligence operations, security operations, and cyber risk management teams to work together for more effective, efficient, and collaborative cyber defense and protection",
			Industries: []domain.Industry{
				domain.IndustryCyberSecurity,
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Mapbox",
			Website: "https://www.mapbox.com/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                3302167,
				IDs:               nil,
				Alias:             "mapbox",
				Name:              "Mapbox",
				Followers:         "57K",
				Employees:         "501-1K",
				AssociatedMembers: "833",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:     "mapbox",
				Followers: "2.1k",
				Verified:  true,
			},
			BlindProfile: domain.BlindProfile{
				Alias:       "Mapbox",
				Employees:   "201 to 500",
				Salary:      "$77K ~ $238K a year",
				Reviews:     "80",
				ReviewsRate: "2.6",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias: "",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Mapbox-EI_IE870890.11,17.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Mapbox-Reviews-E870890.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Mapbox-Jobs-E870890.htm",
				Jobs:        "",
				Reviews:     "302",
				Salaries:    "594",
				ReviewsRate: "3.4",
				Verified:    true,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "",
			},
			Languages: domain.Languages{
				domain.Go: {},
				domain.Rust: {
					GitHubRepositoriesCount: 1,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Software Development Engineer III, Rust",
							ShortDescription:     "EV Routing",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4230075544/",
							Location:             "Helsinki, Uusimaa, Finland",
							Date:                 mustDate("2025-09-22", "2025-09-01", "2025-08-10", "2025-07-20", "2025-06-28", "2025-06-05"),
							WithSalary:           false,
							Remote:               true,
						},
						{
							Title:                "Software Development Engineer III, Rust",
							ShortDescription:     "EV Routing",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4172873087/",
							Location:             "Germany",
							Date:                 mustDate("2025-07-16"),
							WithSalary:           false,
							Remote:               true,
						},
					},
				},
				domain.Zig:     {},
				domain.Scala:   {},
				domain.Elixir:  {},
				domain.Clojure: {},
				domain.Haskell: {},
			},
			ShortDescription: "Mapbox powers navigation for people, packages, and vehicles everywhere",
			Industries:       []domain.Industry{
				// NOP
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Narrative",
			Website: "https://narrative.so/",
			Careers: "https://narrative.so/jobs",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                18496828,
				IDs:               nil,
				Alias:             "narrativeapp",
				Name:              "Narrative",
				Followers:         "1K",
				Employees:         "11-50",
				AssociatedMembers: "25",
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
				domain.Go: {},
				domain.Rust: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Senior Rust Engineer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4242284379/",
							Date:                 mustDate("2025-06-05"),
							WithSalary:           false,
							Remote:               false,
						},
					},
				},
				domain.Zig:     {},
				domain.Scala:   {},
				domain.Elixir:  {},
				domain.Clojure: {},
				domain.Haskell: {},
			},
			ShortDescription: "Narrative speeds up, improves and simplifies the professional photographer’s workflow with smart and easy to use software tools",
			Industries:       []domain.Industry{
				// NOP
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Coralogix",
			Website: "https://coralogix.com/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                3763125,
				IDs:               nil,
				Alias:             "coralogix",
				Name:              "Coralogix",
				Followers:         "44K",
				Employees:         "201-500",
				AssociatedMembers: "498",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:     "coralogix",
				Followers: "96",
				Verified:  true,
			},
			BlindProfile: domain.BlindProfile{
				Alias: "",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias: "",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Coralogix-EI_IE5272217.11,20.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Coralogix-Reviews-E5272217.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Coralogix-Jobs-E5272217.htm",
				Jobs:        "61",
				Reviews:     "71",
				Salaries:    "84",
				ReviewsRate: "4.3",
				Verified:    false,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "",
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 34,
				},
				domain.Rust: {
					GitHubRepositoriesCount: 15,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Backend Engineer – Rust",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4244588178/",
							Date:                 mustDate("2025-06-05"),
							WithSalary:           false,
							Remote:               true,
						},
					},
				},
				domain.Zig: {},
				domain.Scala: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Backend Engineer (Scala)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4336403577/",
							Location:             "Berlin, Berlin, Germany",
							Date:                 mustDate("2025-12-09", "2025-11-18"),
							WithSalary:           false,
							Remote:               false,
						},
					},
				},
				domain.Elixir:  {},
				domain.Clojure: {},
				domain.Haskell: {},
			},
			ShortDescription: "Full-stack observability for logs, metrics, traces and security events",
			Industries: []domain.Industry{
				domain.IndustryDevOps,
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "ARC Analytics",
			Website: "https://arcanalytics.com/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                101811183,
				IDs:               nil,
				Alias:             "arc-risk-analytics",
				Name:              "ARC Analytics",
				Followers:         "721",
				Employees:         "2-10",
				AssociatedMembers: "2",
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
				domain.Go: {},
				domain.Rust: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Senior Backend Rust Engineer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4241100023/",
							Date:                 mustDate("2025-06-04"),
							WithSalary:           false,
							Remote:               false,
						},
					},
				},
				domain.Zig:     {},
				domain.Scala:   {},
				domain.Elixir:  {},
				domain.Clojure: {},
				domain.Haskell: {},
			},
			ShortDescription: "Providing analytical tools, data and insights to support the understanding of risk across the market",
			Industries:       []domain.Industry{
				// NOP
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "STARK",
			Website: "https://stark-defence.com/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                101537178,
				IDs:               nil,
				Alias:             "stark-defence",
				Name:              "STARK",
				Followers:         "8K",
				Employees:         "51-200",
				AssociatedMembers: "23",
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
				domain.Go: {},
				domain.Rust: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Back-End Software Engineer – C++/Rust",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4238023298/",
							Date:                 mustDate("2025-05-30"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Software Engineer – C++ or Rust",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4289968987/",
							Location:             "United States",
							Date:                 mustDate("2025-08-22"),
							WithSalary:           false,
							Remote:               true,
						},
					},
				},
				domain.Zig:     {},
				domain.Scala:   {},
				domain.Elixir:  {},
				domain.Clojure: {},
				domain.Haskell: {},
			},
			ShortDescription: "Defence technology",
			Industries:       []domain.Industry{
				// MilTech, DefTech
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "FalconX",
			Website: "https://www.falconx.io/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                18758988,
				IDs:               nil,
				Alias:             "thefalconx",
				Name:              "FalconX",
				Followers:         "60K",
				Employees:         "201-500",
				AssociatedMembers: "439",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "",
				Verified: false,
			},
			BlindProfile: domain.BlindProfile{
				Alias:       "Falconx",
				Employees:   "51-200",
				Salary:      "",
				Reviews:     "5",
				ReviewsRate: "4.6",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias: "",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-FalconX-EI_IE4615319.11,18.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/FalconX-Reviews-E4615319.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/FalconX-Jobs-E4615319.htm",
				Jobs:        "",
				Reviews:     "40",
				Salaries:    "64",
				ReviewsRate: "4.0",
				Verified:    false,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "",
			},
			Languages: domain.Languages{
				domain.Go: {},
				domain.Rust: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Senior Trading Systems Developer (Java/Rust)",
							ShortDescription:     "Electronic Trading",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4239651257/",
							Location:             "New York, NY",
							Date:                 mustDate("2025-08-06", "2025-07-16", "2025-06-03"),
							WithSalary:           true, // $179k/yr - $242k/yr
							Remote:               false,
						},
						{
							Title:                "Senior Trading Systems Developer (Java/Rust)",
							ShortDescription:     "Electronic Trading",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4299587343/",
							Location:             "New York, NY",
							Date:                 mustDate("2025-09-16"),
							WithSalary:           true, // $179k/yr - $242k/yr
							Remote:               false,
						},
					},
				},
				domain.Zig:     {},
				domain.Scala:   {},
				domain.Elixir:  {},
				domain.Clojure: {},
				domain.Haskell: {},
			},
			ShortDescription: "Digital assets prime brokerage",
			Industries: []domain.Industry{
				domain.IndustryFinTech,
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Bybit",
			Website: "https://www.bybit.com/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                12583916,
				IDs:               nil,
				Alias:             "bybitexchange",
				Name:              "Bybit",
				Followers:         "296K",
				Employees:         "1K-5K",
				AssociatedMembers: "2,312",
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
				domain.Go: {},
				domain.Rust: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Rust Developer",
							ShortDescription:     "(Wallet Backend) – Mandarin Speaker",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4238426576/",
							Date:                 mustDate("2025-05-31"),
							WithSalary:           false,
							Remote:               false,
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
			Industries:       []domain.Industry{
				// NOP
			},
			Ignore: true, // Crypto
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Mythical Games",
			Website: "https://mythicalgames.com/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                18887800,
				IDs:               nil,
				Alias:             "mythical",
				Name:              "Mythical Games",
				Followers:         "53K",
				Employees:         "201-500",
				AssociatedMembers: "244",
				Verified:          false,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:     "MythicalGames",
				Followers: "42",
				Verified:  false,
			},
			BlindProfile: domain.BlindProfile{
				Alias: "",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias: "",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Mythical-Games-EI_IE2392684.11,25.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Mythical-Games-Reviews-E2392684.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Mythical-Games-Jobs-E2392684.htm",
				Jobs:        "",
				Reviews:     "75",
				Salaries:    "110",
				ReviewsRate: "3.3",
				Verified:    true,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "",
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 2,
				},
				domain.Rust: {
					GitHubRepositoriesCount: 1,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Senior Rust Developer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4240170944/",
							Date:                 mustDate("2025-05-30"),
							WithSalary:           false,
							Remote:               true,
						},
					},
				},
				domain.Zig:     {},
				domain.Scala:   {},
				domain.Elixir:  {},
				domain.Clojure: {},
				domain.Haskell: {},
			},
			ShortDescription: "Game technology studio",
			Industries: []domain.Industry{
				domain.IndustryGameDev,
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "PayPal",
			Website: "https://www.paypal.com/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                1482,
				IDs:               []int{1482, 13308, 150981, 552278, 1310825, 2202637, 2226655},
				Alias:             "paypal",
				Name:              "PayPal",
				Followers:         "1M",
				Employees:         "10K+",
				AssociatedMembers: "34,302",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:     "paypal",
				Followers: "1.5k",
				Verified:  false,
			},
			BlindProfile: domain.BlindProfile{
				Alias:       "PayPal",
				Employees:   "10,000+",
				Salary:      "$23K ~ $282K a year",
				Reviews:     "2,170",
				ReviewsRate: "3.5",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "paypal",
				Employees: "34,620",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-PayPal-EI_IE9848.11,17.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/PayPal-Reviews-E9848.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/PayPal-Jobs-E9848.htm",
				Jobs:        "1.5K",
				Reviews:     "10K",
				Salaries:    "23K",
				ReviewsRate: "3.7",
				Verified:    true,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "",
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 7,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Senior Golang Engineer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4237498549/",
							Date:                 mustDate("2025-06-06"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Senior Software Engineer – Golang/Java",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4284788727/",
							Location:             "Bengaluru, Karnataka, India",
							Date:                 mustDate("2025-08-22"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Software Engineer – Golang",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4293689104/",
							Location:             "Bengaluru, Karnataka, India",
							Date:                 mustDate("2025-09-04"),
							WithSalary:           false,
							Remote:               false,
						},
					},
				},
				domain.Rust: {},
				domain.Zig:  {},
				domain.Scala: {
					GitHubRepositoriesCount: 5,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Senior Spark/Scala Engineer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4242758126/",
							Date:                 mustDate("2025-06-07"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Senior Software Engineer – Scala",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4301461830/",
							Location:             "New York, NY",
							Date:                 mustDate("2025-09-21"),
							WithSalary:           false,
							Remote:               false,
						},
					},
				},
				domain.Elixir:  {},
				domain.Clojure: {},
				domain.Haskell: {},
			},
			ShortDescription: "",
			Industries: []domain.Industry{
				domain.IndustryFinTech,
			},
		},
	}
}
