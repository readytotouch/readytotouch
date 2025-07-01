package db

import (
	"github.com/readytotouch/readytotouch/internal/domain"
)

func companies11Shard() []domain.CompanyProfile {
	return []domain.CompanyProfile{
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "NBCUniversal",
			Website: "https://www.nbcuniversal.com/",
			Careers: "https://www.nbcunicareers.com/",
			About:   "https://www.nbcuniversal.com/about",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                1828,
				Alias:             "nbcuniversal-inc-",
				Name:              "NBCUniversal",
				Followers:         "1M",
				Employees:         "10K+",
				AssociatedMembers: "58,328",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "nbcu-ds",
				Verified: false,
			},
			BlindProfile: domain.BlindProfile{
				Alias:       "NBCUniversal",
				Employees:   "10,000+",
				Salary:      "$67K ~ $276K a year",
				Reviews:     "226",
				ReviewsRate: "3.6",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "nbc",
				Employees: "52,370",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.ca/Overview/Working-at-NBCUniversal-EI_IE32038.11,23.htm",
				ReviewsURL:  "https://www.glassdoor.ca/Reviews/NBCUniversal-Reviews-E32038.htm",
				JobsURL:     "https://www.glassdoor.ca/Jobs/NBCUniversal-Jobs-E32038.htm",
				Jobs:        "464",
				Reviews:     "6.9k",
				Salaries:    "13K",
				ReviewsRate: "4.0",
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
							Title:                "Senior Data Engineer (Scala/Spark)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4136839875/",
							Date:                 mustDate("2025-02-20"),
							WithSalary:           true, // $115.000 - $145.000 per year
							Remote:               true,
						},
						{
							Title:                "Senior Data Engineer (Scala/Spark)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4169604784/",
							Date:                 mustDate("2025-03-01"),
							WithSalary:           true, // $115.000 - $145.000 per year
							Remote:               true,
						},
					},
				},
				domain.Elixir:  {},
				domain.Clojure: {},
				domain.Haskell: {},
			},
			ShortDescription: "American multinational mass media and entertainment conglomerate",
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Sendle",
			Website: "https://www.sendle.com/",
			Careers: "https://try.sendle.com/careers",
			About:   "https://try.sendle.com/en-au/about",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                4846518,
				Alias:             "sendle",
				Name:              "Sendle",
				Followers:         "9K",
				Employees:         "51-200",
				AssociatedMembers: "130",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "sendle",
				Verified: true,
			},
			BlindProfile: domain.BlindProfile{
				Alias: "",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "sendle",
				Employees: "180",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.ca/Overview/Working-at-Sendle-EI_IE2860030.11,17.htm",
				ReviewsURL:  "https://www.glassdoor.ca/Reviews/Sendle-Reviews-E2860030.htm",
				JobsURL:     "https://www.glassdoor.ca/Jobs/Sendle-Jobs-E2860030.htm",
				Jobs:        "0",
				Reviews:     "71",
				Salaries:    "82",
				ReviewsRate: "3.8",
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
							Title:                "Senior Product Engineer (React/Elixir)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4157295270/",
							Date:                 mustDate("2025-03-05"), // mustDate("2025-02-20"),
							WithSalary:           false,
							Remote:               true,
						},
					},
				},
				domain.Clojure: {},
				domain.Haskell: {},
			},
			ShortDescription: "Australian virtual courier company which provides courier services",
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Aaronson Rappaport Feinstein & Deutsch, LLP",
			Website: "https://www.arfdlaw.com/",
			Careers: "",
			About:   "https://www.arfdlaw.com/our-firm/",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                2436230,
				Alias:             "arfdlaw",
				Name:              "Aaronson Rappaport Feinstein & Deutsch, LLP",
				Followers:         "3K",
				Employees:         "51-200",
				AssociatedMembers: "120",
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
							Title:                "Software Developer (Rust)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4156586021/",
							Date:                 mustDate("2025-02-20"),
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
			ShortDescription: "It is an American law firm",
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Canva",
			Website: "https://www.canva.com/",
			Careers: "https://www.lifeatcanva.com/en/jobs/",
			About:   "https://www.canva.com/about/",
			Blog:    "https://www.canva.dev/blog/engineering/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                2850862,
				Alias:             "canva",
				Name:              "Canva",
				Followers:         "2M",
				Employees:         "1K-5K",
				AssociatedMembers: "7,103",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "canva",
				Verified: false,
			},
			BlindProfile: domain.BlindProfile{
				Alias:       "canva",
				Employees:   "501 to 1,000",
				Salary:      "",
				Reviews:     "382",
				ReviewsRate: "3.8",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "canva",
				Employees: "1,500",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Canva-EI_IE1013251.11,16.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Canva-Reviews-E1013251.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Canva-Jobs-E1013251.htm",
				Jobs:        "80",
				Reviews:     "880",
				Salaries:    "1.8K",
				ReviewsRate: "4.1",
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
							Title:                "Senior Platform Engineer (Golang, Kubernetes)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4238207309/",
							Date:                 mustDate("2025-05-28"),
							WithSalary:           false,
							Remote:               false, // open to remote across ANZ
						},
					},
				},
				domain.Rust: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Staff Frontend Engineer (Rust)",
							ShortDescription:     "Editing Foundations",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4160846209/",
							Date:                 mustDate("2025-04-09"), // mustDate("2025-04-05"), // mustDate("2025-03-14"), // mustDate("2025-02-21"),
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
			ShortDescription: "Australian multinational software company that provides a graphic design platform",
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Adtran",
			Website: "https://www.adtran.com/",
			Careers: "https://www.adtran.com/en/about-us/careers",
			About:   "https://www.adtran.com/en/about-us",
			Blog:    "https://www.blog.adtran.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                8049,
				Alias:             "adtran",
				Name:              "Adtran",
				Followers:         "49K",
				Employees:         "1K-5K",
				AssociatedMembers: "2,408",
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
				Alias:     "adtran",
				Employees: "2,000",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-ADTRAN-EI_IE3056.11,17.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/ADTRAN-Reviews-E3056.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/ADTRAN-Jobs-E3056.htm",
				Jobs:        "95",
				Reviews:     "441",
				Salaries:    "816",
				ReviewsRate: "3.6",
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
							Title:                "Senior Software Engineer (Rust and C++)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4160542551/",
							Date:                 mustDate("2025-04-27"), // mustDate("2025-04-05"), // mustDate("2025-03-14"), // mustDate("2025-02-20"),
							WithSalary:           false,
							Remote:               false,
						},
					},
				},
				domain.Zig: {},
				domain.Scala: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Senior Software Engineer (Backend/Scala)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4188798669/",
							Date:                 mustDate("2025-06-14"), // mustDate("2025-05-22"), // mustDate("2025-04-30"), // mustDate("2025-04-10"),
							WithSalary:           false,
							Remote:               false,
						},
					},
				},
				domain.Elixir:  {},
				domain.Clojure: {},
				domain.Haskell: {},
			},
			ShortDescription: "American telecommunications networking equipment vendor",
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Photoroom",
			Website: "https://www.photoroom.com/",
			Careers: "https://www.photoroom.com/company#careers",
			About:   "https://www.photoroom.com/company",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                26243215,
				Alias:             "photoroom",
				Name:              "Photoroom",
				Followers:         "14K",
				Employees:         "11-50",
				AssociatedMembers: "471",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "photoroom",
				Verified: true,
			},
			BlindProfile: domain.BlindProfile{
				Alias: "",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "photoroom",
				Employees: "60",
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
							URL:                  "https://www.linkedin.com/jobs/view/4159248269/",
							Date:                 mustDate("2025-02-20"),
							WithSalary:           true, // $90.000 - $110.000 per year + Stock-Options/BSPCE
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
			ShortDescription: "AI-Photo editing software",
			YahooFinanceURL:  "https://www.ycombinator.com/companies/photoroom",
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Etsy",
			Website: "https://www.etsy.com/",
			Careers: "https://careers.etsy.com/global/en",
			About:   "https://www.etsy.com/about",
			Blog:    "https://www.etsy.com/codeascraft",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                67849,
				Alias:             "etsy",
				Name:              "Etsy",
				Followers:         "300K",
				Employees:         "1K-5K",
				AssociatedMembers: "8,730",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "etsy",
				Verified: true,
			},
			BlindProfile: domain.BlindProfile{
				Alias: "",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "etsy",
				Employees: "6,120",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Etsy-EI_IE42751.11,15.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Etsy-Reviews-E42751.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Etsy-Jobs-E42751.htm",
				Jobs:        "67",
				Reviews:     "640",
				Salaries:    "1.7K",
				ReviewsRate: "4.0",
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
							Title:                "Senior Software Engineer II, Scala",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4139270225/",
							Date:                 mustDate("2025-02-21"),
							WithSalary:           true, // $205.000 - $237.000 per year
							Remote:               false,
						},
					},
				},
				domain.Elixir:  {},
				domain.Clojure: {},
				domain.Haskell: {},
			},
			ShortDescription: "American online e-commerce marketplace",
			Industries: []domain.Industry{
				domain.IndustryECommerce,
			},
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "AppFlowy",
			Website: "https://appflowy.com/",
			Careers: "https://appflowy.com/join",
			About:   "https://appflowy.com/about",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                81474610,
				IDs:               nil,
				Alias:             "appflowy",
				Name:              "AppFlowy",
				Followers:         "3K",
				Employees:         "2-10",
				AssociatedMembers: "9",
				Verified:          false,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "AppFlowy-IO",
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
			OttaProfileSlug:   "",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go: {},
				domain.Rust: {
					GitHubRepositoriesCount: 7,
					Vacancies:               nil,
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
			Industries:                nil,
			HasEmployeesFromCountries: nil,
			RustFoundationMember:      true,
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Arm",
			Website: "https://www.arm.com/",
			Careers: "https://careers.arm.com/",
			About:   "https://www.arm.com/company",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                4472,
				IDs:               []int{4472, 102650, 569760, 1584887, 1968088, 2855159, 5103215, 18122054},
				Alias:             "arm",
				Name:              "Arm",
				Followers:         "541K",
				Employees:         "5K-10K",
				AssociatedMembers: "9,416",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "",
				Verified: false,
			},
			BlindProfile: domain.BlindProfile{
				Alias:       "arm",
				Employees:   "5,001 to 10,000",
				Salary:      "$50K ~ $307K a year",
				Reviews:     "218",
				ReviewsRate: "3.8",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias: "arm",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Arm-EI_IE7834.11,14.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Arm-Reviews-E7834.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Arm-Jobs-E7834.htm",
				Jobs:        "501",
				Reviews:     "2.8K",
				Salaries:    "15K",
				ReviewsRate: "4.4",
				Verified:    true,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "Arm",
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Fullstack Web Developer – Golang",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4236800842/",
							Date:                 mustDate("2025-05-28"),
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
			ShortDescription:          "",
			DealroomURL:               "",
			CrunchbaseURL:             "",
			PitchbookURL:              "",
			YahooFinanceURL:           "",
			GoogleFinanceURL:          "",
			YCombinatorURL:            "",
			Industries:                nil,
			HasEmployeesFromCountries: nil,
			RustFoundationMember:      true,
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Devolutions",
			Website: "https://devolutions.net/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                1533251,
				IDs:               nil,
				Alias:             "devolutions-inc-",
				Name:              "Devolutions",
				Followers:         "5K",
				Employees:         "51-200",
				AssociatedMembers: "195",
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
			OttaProfileSlug:   "",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go:      {},
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
			Industries:                nil,
			HasEmployeesFromCountries: nil,
			RustFoundationMember:      true,
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Embecosm",
			Website: "https://embecosm.com/",
			Careers: "",
			About:   "https://embecosm.com/about-us/",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                2933938,
				IDs:               nil,
				Alias:             "embecosm",
				Name:              "Embecosm",
				Followers:         "451",
				Employees:         "11-50",
				AssociatedMembers: "17",
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
			OttaProfileSlug:   "",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go:      {},
				domain.Rust:    {},
				domain.Zig:     {},
				domain.Scala:   {},
				domain.Elixir:  {},
				domain.Clojure: {},
				domain.Haskell: {},
			},
			ShortDescription:          "Delivering open source software tool chain, embedded operating system and AI services",
			DealroomURL:               "",
			CrunchbaseURL:             "",
			PitchbookURL:              "",
			YahooFinanceURL:           "",
			GoogleFinanceURL:          "",
			YCombinatorURL:            "",
			Industries:                nil,
			HasEmployeesFromCountries: nil,
			RustFoundationMember:      true,
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Helsing",
			Website: "https://helsing.ai/",
			Careers: "https://helsing.ai/jobs",
			About:   "https://helsing.ai/company",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                71719093,
				IDs:               nil,
				Alias:             "helsing",
				Name:              "Helsing",
				Followers:         "42K",
				Employees:         "201-500",
				AssociatedMembers: "371",
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
				Alias: "helsing",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Helsing-EI_IE6763957.11,18.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Helsing-Reviews-E6763957.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Helsing-Jobs-E6763957.htm",
				Jobs:        "",
				Reviews:     "38",
				Salaries:    "61",
				ReviewsRate: "4.4",
				Verified:    true,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "",
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go:      {},
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
			Industries:                nil,
			HasEmployeesFromCountries: nil,
			RustFoundationMember:      true,
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "JetBrains",
			Website: "https://www.jetbrains.com/",
			Careers: "https://www.jetbrains.com/careers/jobs/",
			About:   "https://www.jetbrains.com/company/",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                12515,
				IDs:               nil,
				Alias:             "jetbrains",
				Name:              "JetBrains",
				Followers:         "354K",
				Employees:         "1K-5K",
				AssociatedMembers: "2,292",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "jetbrains",
				Verified: true,
			},
			BlindProfile: domain.BlindProfile{
				Alias:       "jetbrains",
				Employees:   "1,001 to 5,000",
				Salary:      "",
				Reviews:     "33",
				ReviewsRate: "4.2",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias: "jetbrains",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-JetBrains-EI_IE222299.11,20.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/JetBrains-Reviews-E222299.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/JetBrains-Jobs-E222299.htm",
				Jobs:        "30",
				Reviews:     "288",
				Salaries:    "737",
				ReviewsRate: "4.3",
				Verified:    true,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "Jetbrains",
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go: {},
				domain.Rust: {
					GitHubRepositoriesCount: 2,
					Vacancies:               nil,
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
			Industries:       nil,
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
				domain.Czechia,
			},
			RustFoundationMember: true,
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Garena",
			Website: "https://www.garena.sg/",
			Careers: "https://careers.garena.com/global/careers",
			About:   "https://careers.garena.com/global/about",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                13678254,
				IDs:               nil,
				Alias:             "garena",
				Name:              "Garena",
				Followers:         "304K",
				Employees:         "1K-5K",
				AssociatedMembers: "5,066",
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
							Title:                "Full Stack Developer (Golang, ReactJS)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4163890859/",
							Date:                 mustDate("2025-02-24"),
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
			ShortDescription: "Singapur online games developer and publisher",
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "MoneyLion",
			Website: "https://www.moneylion.com/",
			Careers: "https://www.moneylion.com/careers/",
			About:   "https://www.moneylion.com/about/",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                3489789,
				IDs:               nil,
				Alias:             "moneylion",
				Name:              "MoneyLion",
				Followers:         "33K",
				Employees:         "501-1K",
				AssociatedMembers: "588",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "MoneyLion",
				Verified: true,
			},
			BlindProfile: domain.BlindProfile{
				Alias: "",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "moneylion",
				Employees: "450",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-MoneyLion-EI_IE1053194.11,20.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/MoneyLion-Reviews-E1053194.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/MoneyLion-Jobs-E1053194.htm",
				Jobs:        "37",
				Reviews:     "261",
				Salaries:    "432",
				ReviewsRate: "3.7",
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
							Title:                "Staff Backend Engineer (Scala)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4160597734/",
							Date:                 mustDate("2025-03-05"), // mustDate("2025-02-21"),
							WithSalary:           false,
							Remote:               true,
						},
					},
				},
				domain.Elixir:  {},
				domain.Clojure: {},
				domain.Haskell: {},
			},
			ShortDescription: "American fintech platform that specializes in providing fast short-term loans to people with limited credit",
			Industries: []domain.Industry{
				domain.IndustryFinTech,
			},
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Loblaw Companies Limited",
			Website: "https://www.loblaw.ca/",
			Careers: "https://careers.loblaw.ca/",
			About:   "https://www.loblaw.ca/en/who-we-are/",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                164756,
				IDs:               []int{164756, 9450129, 79658411},
				Alias:             "loblaw-companies-limited",
				Name:              "Loblaw Companies Limited",
				Followers:         "243K",
				Employees:         "10K+",
				AssociatedMembers: "23,193",
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
				Alias:     "loblaw-companies",
				Employees: "200,000",
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
							Title:                "Senior Developer (Elixir)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4139223224/",
							Date:                 mustDate("2025-02-21"),
							WithSalary:           false,
							Remote:               false,
						},
					},
				},
				domain.Clojure: {},
				domain.Haskell: {},
			},
			ShortDescription: "Canadian retailer and distributor of food and pharmaceutical products",
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Qonto",
			Website: "https://qonto.com/",
			Careers: "https://qonto.com/en/careers",
			About:   "https://qonto.com/en/about",
			Blog:    "https://medium.com/qonto-way",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                13205888,
				IDs:               nil,
				Alias:             "qonto",
				Name:              "Qonto",
				Followers:         "135K",
				Employees:         "1K-5K",
				AssociatedMembers: "1,936",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "qonto",
				Verified: false,
			},
			BlindProfile: domain.BlindProfile{
				Alias: "",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "qonto",
				Employees: "660",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Qonto-EI_IE1879266.11,16.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Qonto-Reviews-E1879266.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Qonto-Jobs-E1879266.htm",
				Jobs:        "32",
				Reviews:     "341",
				Salaries:    "727",
				ReviewsRate: "4.3",
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
							Title:                "Senior Go Backend Engineer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4144346084/",
							Date:                 mustDate("2025-02-26"),
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
			ShortDescription: "French FinTech platform for businesses and freelancers",
			Industries: []domain.Industry{
				domain.IndustryFinTech,
			},
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "iCapital",
			Website: "https://icapital.com/",
			Careers: "https://icapital.com/careers/",
			About:   "https://icapital.com/about-us/",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                3080201,
				IDs:               nil,
				Alias:             "icapital-network-inc",
				Name:              "iCapital",
				Followers:         "67K",
				Employees:         "1K-5K",
				AssociatedMembers: "1,671",
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
				Alias:     "icapital",
				Employees: "690",
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
							Title:                "Back-End Engineer, Scala",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4168470393/",
							Date:                 mustDate("2025-02-26"),
							WithSalary:           true, // $120.000 - $160.000 per year
							Remote:               false,
						},
					},
				},
				domain.Elixir:  {},
				domain.Clojure: {},
				domain.Haskell: {},
			},
			ShortDescription: "American FinTech company that provides platform solutions related to alternative investments in the private markets",
			Industries: []domain.Industry{
				domain.IndustryFinTech,
			},
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Expedia Group",
			Website: "https://expediagroup.com/",
			Careers: "https://careers.expediagroup.com/",
			About:   "https://www.expediagroup.com/who-we-are/our-story/default.aspx",
			Blog:    "https://medium.com/expedia-group-tech",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                2751,
				IDs:               []int{2751, 34274, 86849, 150884, 404429, 1965675},
				Alias:             "expedia",
				Name:              "Expedia Group",
				Followers:         "807K",
				Employees:         "10K+",
				AssociatedMembers: "22,684",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:     "expediagroup",
				Followers: "159",
				Verified:  false,
			},
			BlindProfile: domain.BlindProfile{
				Alias: "",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias: "",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Expedia-Group-EI_IE9876.11,24.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Expedia-Group-Reviews-E9876.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Expedia-Group-Jobs-E9876.htm",
				Jobs:        "172",
				Reviews:     "8.2K",
				Salaries:    "20K",
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
							Title:                "Data Engineer III — Scala/Spark",
							ShortDescription:     "Media Solutions",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4141235282/",
							Date:                 mustDate("2025-02-25"),
							WithSalary:           true, // $128.000 - $179.500 per year
							Remote:               false,
						},
					},
				},
				domain.Elixir:  {},
				domain.Clojure: {},
				domain.Haskell: {},
			},
			ShortDescription: "American travel technology company that owns and operates travel fare aggregators and travel metasearch engines",
			YahooFinanceURL:  "https://finance.yahoo.com/quote/EXPE/",
			GoogleFinanceURL: "https://www.google.com/finance/quote/EXPE:NASDAQ",
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Esri",
			Website: "https://www.esri.com/",
			Careers: "https://www.esri.com/en-us/about/careers/overview",
			About:   "https://www.esri.com/en-us/about/about-esri/overview",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                5311,
				IDs:               []int{5311, 2905030},
				Alias:             "esri",
				Name:              "Esri",
				Followers:         "430K",
				Employees:         "5K-10K",
				AssociatedMembers: "6,805",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:     "esri",
				Followers: "2k",
				Verified:  true,
			},
			BlindProfile: domain.BlindProfile{
				Alias: "",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "esri",
				Employees: "5,750",
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
							Title:                "Software Development Engineer II – Rust",
							ShortDescription:     "Design and develop high-performance components for the ArcGIS API for Python using Rust and PyO3",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4190174086/",
							Date:                 mustDate("2025-04-15"),
							WithSalary:           false,
							Remote:               false,
						},
					},
				},
				domain.Zig: {},
				domain.Scala: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Java/Scala Software Development Engineer II",
							ShortDescription:     "Raster and Imagery",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4075632689/",
							Date:                 mustDate("2025-02-25"),
							WithSalary:           true, // $97.344 - $162.240 per year
							Remote:               false,
						},
					},
				},
				domain.Elixir:  {},
				domain.Clojure: {},
				domain.Haskell: {},
			},
			ShortDescription: "American geographic information system (GIS) software",
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "CloudLinux",
			Website: "https://www.cloudlinux.com/",
			Careers: "https://www.cloudlinux.com/about-us-company-jobs/",
			About:   "https://www.cloudlinux.com/about-us/",
			Blog:    "https://blog.cloudlinux.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                632769,
				IDs:               nil,
				Alias:             "cloudlinux",
				Name:              "CloudLinux",
				Followers:         "33K",
				Employees:         "51-200",
				AssociatedMembers: "245",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "cloudlinux",
				Verified: false,
			},
			BlindProfile: domain.BlindProfile{
				Alias: "",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "cloudlinux",
				Employees: "150",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-CloudLinux-EI_IE1448501.11,21.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/CloudLinux-Reviews-E1448501.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/CloudLinux-Jobs-E1448501.htm",
				Jobs:        "24",
				Reviews:     "67",
				Salaries:    "46",
				ReviewsRate: "4.2",
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
							Title:                "Senior Python/Go Developer",
							ShortDescription:     "Imunify360 team",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4170791654/",
							Date:                 mustDate("2025-04-02"), // mustDate("2025-03-13"), // mustDate("2025-03-05"), // mustDate("2025-02-28"),
							WithSalary:           false,
							Remote:               true,
						},
						{
							Title:                "Senior Python/Go Developer",
							ShortDescription:     "Imunify360 team",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4197920577/",
							Date:                 mustDate("2025-04-18"), // mustDate("2025-04-14"),
							WithSalary:           false,
							Remote:               true,
						},
						{
							Title:                "Senior Go Developer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4208666795/",
							Date:                 mustDate("2025-04-17"),
							WithSalary:           false,
							Remote:               true,
						},
						{
							Title:                "Senior Python/Go Developer",
							ShortDescription:     "Imunify360 team",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4213174081/",
							Date:                 mustDate("2025-05-15"), // mustDate("2025-05-07"),
							WithSalary:           false,
							Remote:               true,
						},
						{
							Title:                "Senior Go Developer",
							ShortDescription:     "Imunify360 team",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4227285156/",
							Date:                 mustDate("2025-06-04"), // mustDate("2025-05-28"),
							WithSalary:           false,
							Remote:               true,
						},
						{
							Title:                "Senior Python/Go Developer",
							ShortDescription:     "Imunify360 team",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4245418480/",
							Date:                 mustDate("2025-06-11"),
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
			ShortDescription: "Is a Linux based operating system designed to give shared hosting providers a more stable and secure OS",
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Green Packet",
			Website: "https://www.greenpacket.com/",
			Careers: "https://www.greenpacket.com/careers",
			About:   "https://www.greenpacket.com/company/about-us",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                81702,
				IDs:               nil,
				Alias:             "greenpacket",
				Name:              "Green Packet",
				Followers:         "13K",
				Employees:         "51-200",
				AssociatedMembers: "172",
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
							Title:                "Full Stack Developer (PHP,Go)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4168191417/",
							Date:                 mustDate("2025-02-28"),
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
			ShortDescription: "Malaysian Telecommunications, Internet Service Providers, Website Hosting & Internet-Related Services, and Mobile company",
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Bloomreach",
			Website: "https://www.bloomreach.com/",
			Careers: "https://www.bloomreach.com/en/careers",
			About:   "https://www.bloomreach.com/en/about-us",
			Blog:    "https://dev.to/bloomreach",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                36076,
				IDs:               nil,
				Alias:             "bloomreach",
				Name:              "Bloomreach",
				Followers:         "70K",
				Employees:         "501-1K",
				AssociatedMembers: "1,072",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "bloomreach",
				Verified: false,
			},
			BlindProfile: domain.BlindProfile{
				Alias: "",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "bloomreach",
				Employees: "850",
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
							Title:                "Senior Go Platform Developer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4168467917/",
							Date:                 mustDate("2025-04-12"), // mustDate("2025-03-27"), // mustDate("2025-02-26"),
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
			ShortDescription: "Is a digital marketing and website optimization company",
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Fidelity Investments",
			Website: "https://www.fidelity.com/",
			Careers: "https://www.fidelity.com/about-fidelity/careers",
			About:   "https://www.fidelity.com/about-fidelity/our-company",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                1307,
				IDs:               nil,
				Alias:             "fidelity-investments",
				Name:              "Fidelity Investments",
				Followers:         "1M",
				Employees:         "10K+",
				AssociatedMembers: "80,034",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:     "fidelity",
				Followers: "244",
				Verified:  true,
			},
			BlindProfile: domain.BlindProfile{
				Alias: "",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "fidelity-investments",
				Employees: "60,570",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Fidelity-Investments-EI_IE2786.11,31.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Fidelity-Investments-Reviews-E2786.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Fidelity-Investments-Jobs-E2786.htm",
				Jobs:        "1.1K",
				Reviews:     "19K",
				Salaries:    "37K",
				ReviewsRate: "4.1",
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
							Title:                "Principal Full Stack Engineer (Golang)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4168166029/",
							Date:                 mustDate("2025-03-21"), // mustDate("2025-02-28"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Principal Full Stack Engineer (Golang)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4214349134/",
							Date:                 mustDate("2025-06-05"), // mustDate("2025-04-22"),
							WithSalary:           false,
							Remote:               false,
						},
					},
				},
				domain.Rust: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Senior Software Engineer (RUST/Python)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4231780612/",
							Date:                 mustDate("2025-06-08"),
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
			ShortDescription: "American multinational financial services corporation",
			Industries: []domain.Industry{
				domain.IndustryFinTech,
			},
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Bread Financial",
			Website: "https://www.breadfinancial.com/",
			Careers: "https://www.breadfinancial.com/en/who-we-are/careers.html",
			About:   "https://www.breadfinancial.com/en/who-we-are.html",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                84108929,
				IDs:               nil,
				Alias:             "bread-financial",
				Name:              "Bread Financial",
				Followers:         "118K",
				Employees:         "5K-10K",
				AssociatedMembers: "4,001",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "getbread", // redirect
				Verified: false,
			},
			BlindProfile: domain.BlindProfile{
				Alias: "",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "bread-financial",
				Employees: "1,190",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Bread-Financial-EI_IE11421.11,26.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Bread-Financial-Reviews-E11421.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Bread-Financial-Jobs-E11421.htm",
				Jobs:        "24",
				Reviews:     "1.8K",
				Salaries:    "2.4K",
				ReviewsRate: "4.5",
				Verified:    true,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "",
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 17,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Golang Engineer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4169628634/",
							Date:                 mustDate("2025-03-06"), // mustDate("2025-03-01"),
							WithSalary:           true,                   // $87.900 - $198.900 per year
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
			ShortDescription: "FinTech financial services company providing simple, personalized payment, lending and saving solutions",
			Industries: []domain.Industry{
				domain.IndustryFinTech,
			},
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Prosus Group",
			Website: "https://www.prosus.com/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                40677740,
				IDs:               nil,
				Alias:             "prosusgroup",
				Name:              "Prosus Group",
				Followers:         "88K",
				Employees:         "10K+",
				AssociatedMembers: "296",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:     "ProsusAI",
				Followers: "38",
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
							Title:                "Senior Software Engineer (Go)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4146942038/",
							Date:                 mustDate("2025-03-01"),
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
			ShortDescription: "Netherlands consumer internet group and technology investor company",
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Appear",
			Website: "https://www.appear.net/",
			Careers: "https://careers.appear.net/jobs",
			About:   "https://www.appear.net/company-overview/",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                99029,
				IDs:               nil,
				Alias:             "appear-net",
				Name:              "Appear",
				Followers:         "8K",
				Employees:         "51-200",
				AssociatedMembers: "228",
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
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Appear-EI_IE5458916.11,17.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Appear-Reviews-E5458916.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Appear-Jobs-E5458916.htm",
				Jobs:        "1",
				Reviews:     "31",
				Salaries:    "79",
				ReviewsRate: "3.4",
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
							Title:                "Senior Rust Developer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4174424707/",
							Date:                 mustDate("2025-03-03"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Senior Rust Developer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4224562038/",
							Date:                 mustDate("2025-06-17"), // mustDate("2025-05-28"), // mustDate("2025-05-07"),
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
			ShortDescription: "Software company, who specialize in high-capacity solutions for media processing and content delivery, catering to the ever-changing needs of the media, entertainment, and sports industries",
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Worldline",
			Website: "https://worldline.com/",
			Careers: "https://jobs.worldline.com/",
			About:   "https://worldline.com/en/home/top-navigation/about-worldline/who-we-are",
			Blog:    "https://blog.worldline.tech/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                2611817,
				IDs:               []int{95402, 2611817},
				Alias:             "worldlineglobal",
				Name:              "Worldline",
				Followers:         "349K",
				Employees:         "10K+",
				AssociatedMembers: "15,596",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:     "worldline",
				Followers: "71",
				Verified:  false,
			},
			BlindProfile: domain.BlindProfile{
				Alias: "",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "worldline",
				Employees: "20,000",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Worldline-EI_IE2068593.11,20.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Worldline-Reviews-E2068593.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Worldline-Jobs-E2068593.htm",
				Jobs:        "296",
				Reviews:     "1.7K",
				Salaries:    "3K",
				ReviewsRate: "3.6",
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
							Title:                "Go Developer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4172652078/",
							Date:                 mustDate("2025-03-31"), // mustDate("2025-03-04"),
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
			ShortDescription: "French electronic payment services and banking",
			Industries: []domain.Industry{
				domain.IndustryFinTech,
			},
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Whatnot",
			Website: "https://www.whatnot.com/",
			Careers: "https://careers.whatnot.com/",
			About:   "https://www.whatnot.com/about-us",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                52160608,
				IDs:               nil,
				Alias:             "whatnot-inc",
				Name:              "Whatnot",
				Followers:         "58K",
				Employees:         "501-1K",
				AssociatedMembers: "1,058",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "Whatnot-Inc",
				Verified: true,
			},
			BlindProfile: domain.BlindProfile{
				Alias: "",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "whatnot",
				Employees: "270",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Whatnot-EI_IE5065998.11,18.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Whatnot-Reviews-E5065998.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Whatnot-Jobs-E5065998.htm",
				Jobs:        "30",
				Reviews:     "182",
				Salaries:    "225",
				ReviewsRate: "4.2",
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
					GitHubRepositoriesCount: 9,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Elixir Engineer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4169557069/",
							Date:                 mustDate("2025-04-12"), // mustDate("2025-03-20"), // mustDate("2025-02-22"),
							WithSalary:           false,
							Remote:               false,
						},
					},
				},
				domain.Clojure: {},
				domain.Haskell: {},
			},
			ShortDescription: "British community marketplace for buyers and sellers",
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "DNSFilter",
			Website: "https://www.dnsfilter.com/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                10288375,
				IDs:               nil,
				Alias:             "dnsfilter",
				Name:              "DNSFilter",
				Followers:         "23K",
				Employees:         "51-200",
				AssociatedMembers: "171",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:     "dnsfilter",
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
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-DNSFilter-EI_IE2146068.11,20.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/DNSFilter-Reviews-E2146068.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/DNSFilter-Jobs-E2146068.htm",
				Jobs:        "",
				Reviews:     "59",
				Salaries:    "82",
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
							Title:                "Senior Software Engineer – Golang",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4174530724/",
							Date:                 mustDate("2025-03-06"),
							WithSalary:           true,
							Remote:               true,
						},
						{
							Title:                "Senior Software Engineer – Golang",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4214973364/",
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
			ShortDescription: "DNS security that identifies threats",
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Forcepoint",
			Website: "https://www.forcepoint.com/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                7584467,
				IDs:               nil,
				Alias:             "forcepoint",
				Name:              "Forcepoint",
				Followers:         "166K",
				Employees:         "1K-5K",
				AssociatedMembers: "1,813",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "",
				Verified: false,
			},
			BlindProfile: domain.BlindProfile{
				Alias:       "forcepoint",
				Employees:   "1,001 to 5,000 Employees",
				Salary:      "$75K ~ $230K a year",
				Reviews:     "25",
				ReviewsRate: "3.0",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "forcepoint",
				Employees: "2,250",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Forcepoint-EI_IE1262390.11,21.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Forcepoint-Reviews-E1262390.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Forcepoint-Jobs-E1262390.htm",
				Jobs:        "27",
				Reviews:     "1.3K",
				Salaries:    "1.8K",
				ReviewsRate: "3.7",
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
							Title:                "Principal Engineer (Golang, Java, AWS & Data Engineering)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4157762453/",
							Date:                 mustDate("2025-03-06"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Senior Software Engineer – Golang",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4169998476/",
							Date:                 mustDate("2025-03-24"), // mustDate("2025-03-18"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Software Engineer II – Golang",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4214032831/",
							Date:                 mustDate("2025-04-21"),
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
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "NU",
			Website: "https://www.thenu.com/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                76545362,
				IDs:               nil,
				Alias:             "thenu",
				Name:              "NU",
				Followers:         "5K",
				Employees:         "11-50",
				AssociatedMembers: "98",
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
							Title:                "Golang Back-End Engineer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4174841291/",
							Date:                 mustDate("2025-03-05"),
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
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Barracuda",
			Website: "https://www.barracuda.com/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                13957,
				IDs:               nil,
				Alias:             "barracuda-networks",
				Name:              "Barracuda",
				Followers:         "69K",
				Employees:         "1K-5K",
				AssociatedMembers: "2,148",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:     "barracudanetworks",
				Followers: "21",
				Verified:  false,
			},
			BlindProfile: domain.BlindProfile{
				Alias:       "Barracuda",
				Employees:   "1,001 to 5,000",
				Salary:      "$53K ~ $335K a year",
				Reviews:     "20",
				ReviewsRate: "3.6",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "barracuda-networks",
				Employees: "1,500",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "",
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "",
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 8,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Software Engineer (React, Java Spring, Golang)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4173786381/",
							Date:                 mustDate("2025-03-05"),
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
				domain.IndustryCyberSecurity,
			},
		},
		{
			ID:      0,                         // system
			Type:    domain.CompanyTypeStartup, // system
			Name:    "FlatPeak",
			Website: "https://flatpeak.com/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                82675445,
				IDs:               nil,
				Alias:             "flatpeak",
				Name:              "FlatPeak",
				Followers:         "1K",
				Employees:         "11-50",
				AssociatedMembers: "6",
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
							Title:                "Backend Engineer (Golang)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4172858065/",
							Date:                 mustDate("2025-03-04"),
							WithSalary:           true, // £65,000-£80,000 + bonus and stock options
							Remote:               false,
						},
						{
							Title:                "Software Engineer – Golang",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4192226484/",
							Date:                 mustDate("2025-03-24"),
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
			ShortDescription: "Energy tariff data and optimisation platform",
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Teradata",
			Website: "https://www.teradata.com/",
			Careers: "https://www.teradata.com/about-us/careers",
			About:   "https://www.teradata.com/about-us",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                1466,
				IDs:               []int{1466, 14996},
				Alias:             "teradata",
				Name:              "Teradata",
				Followers:         "404K",
				Employees:         "10K+",
				AssociatedMembers: "10,263",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:     "teradata",
				Followers: "106",
				Verified:  false,
			},
			BlindProfile: domain.BlindProfile{
				Alias:       "teradata",
				Employees:   "10,000+",
				Salary:      "$50K ~ $252K a year",
				Reviews:     "122",
				ReviewsRate: "3.3",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "teradata",
				Employees: "10,760",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Teradata-EI_IE14638.11,19.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Teradata-Reviews-E14638.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Teradata-Jobs-E14638.htm",
				Jobs:        "47",
				Reviews:     "3.7K",
				Salaries:    "5.4K",
				ReviewsRate: "3.7",
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
							Title:                "Cloud Engineer (Golang/Java/Python – RestAPI)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4149196862/",
							Date:                 mustDate("2025-03-04"),
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
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Teramind",
			Website: "https://www.teramind.co/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                5090184,
				IDs:               nil,
				Alias:             "teramindco",
				Name:              "Teramind",
				Followers:         "60K",
				Employees:         "51-200",
				AssociatedMembers: "161",
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
							Title:                "Senior Golang Software Engineer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4172639728/",
							Date:                 mustDate("2025-03-04"),
							WithSalary:           false,
							Remote:               true,
						},
						{
							Title:                "Senior Golang Software Engineer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4172643376/",
							Date:                 mustDate("2025-03-04"),
							WithSalary:           false,
							Remote:               true,
						},
						{
							Title:                "Senior Golang Software Engineer",
							ShortDescription:     "2+ years experience in Golang",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4198065578/",
							Date:                 mustDate("2025-04-01"),
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
			ShortDescription: "Teramind Tracker is an employee monitoring and insider threat detection software that records screenshots, audio/video, keystrokes, and other activities",
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Lotus's",
			Website: "https://www.lotuss.com/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                3830488,
				IDs:               nil,
				Alias:             "lotusretails",
				Name:              "Lotus's",
				Followers:         "70K",
				Employees:         "10K+",
				AssociatedMembers: "2,932",
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
							Title:                "Senior Software Engineer – Backend (Golang)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4117492485/",
							Date:                 mustDate("2025-03-04"),
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
			ShortDescription: "Lotus’s is an omni-channel retailer in Thailand",
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Aescape",
			Website: "https://aescape.com/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                11234163,
				IDs:               nil,
				Alias:             "aescape",
				Name:              "Aescape",
				Followers:         "13K",
				Employees:         "51-200",
				AssociatedMembers: "159",
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
							Title:                "Senior/Staff Backend Engineer, Golang",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4172541606/",
							Date:                 mustDate("2025-03-04"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Staff/Principal Backend Engineer, Golang",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4240542373/",
							Date:                 mustDate("2025-06-09"),
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
			ShortDescription: "Building intelligent massage therapy to help people live better, longer",
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "SMC",
			Website: "https://www.smcusa.com/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                6122,
				IDs:               nil,
				Alias:             "smc",
				Name:              "SMC",
				Followers:         "156K",
				Employees:         "10K+",
				AssociatedMembers: "8,749",
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
							Title:                "Lead Software Engineer – Golang",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4172507663/",
							Date:                 mustDate("2025-03-04"),
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
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Crypto",
			Website: "https://www.crypto.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                13322447,
				IDs:               nil,
				Alias:             "cryptocom",
				Name:              "Crypto.com",
				Followers:         "590K",
				Employees:         "1K-5K",
				AssociatedMembers: "6,502",
				Verified:          true,
			},
			Ignore: true, // Crypto
		},
		{
			ID:      0,                         // system
			Type:    domain.CompanyTypeStartup, // system
			Name:    "Depixen",
			Website: "https://www.depixen.com/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                19062020,
				IDs:               nil,
				Alias:             "depixen-sw",
				Name:              "Depixen",
				Followers:         "8K",
				Employees:         "11-50",
				AssociatedMembers: "33",
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
							Title:                "Senior Backend Developer | Senior Go Developer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4174831245/",
							Date:                 mustDate("2025-03-05"),
							WithSalary:           false,
							Remote:               false,
						},
					},
				},
				domain.Rust:  {},
				domain.Zig:   {},
				domain.Scala: {},
				domain.Elixir: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Elixir Developer",
							ShortDescription:     "Business Automation",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4153262881/",
							Date:                 mustDate("2025-04-01"),
							WithSalary:           false,
							Remote:               false,
						},
					},
				},
				domain.Clojure: {},
				domain.Haskell: {},
			},
			ShortDescription: "AI-Powered Autonomous Product Management Systems, Ontology-Based Digital Transformation, RDF-Based Semantic Applications",
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Climatiq",
			Website: "https://www.climatiq.io/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                72159281,
				IDs:               nil,
				Alias:             "climatiq",
				Name:              "Climatiq",
				Followers:         "8K",
				Employees:         "11-50",
				AssociatedMembers: "30",
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
							Title:                "Senior Full Stack Engineer – Climate Tech – Rust & TypeScript",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4175571523/",
							Date:                 mustDate("2025-03-07"),
							WithSalary:           false,
							Remote:               true,
						},
						{
							Title:                "Senior Full Stack Engineer – Climate Tech – Rust & TypeScript",
							ShortDescription:     "Experience using Rust in production or side projects with significant scope",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4200870310/",
							Date:                 mustDate("2025-04-04"),
							WithSalary:           false,
							Remote:               true,
						},
						{
							Title:                "Senior Full Stack Engineer – Climate Tech – Rust & TypeScript",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4226687011/",
							Date:                 mustDate("2025-05-10"),
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
			ShortDescription: "The scope 1, 2, and 3 carbon calculation engine",
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "commercetools",
			Website: "https://www.commercetools.com/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                126050,
				IDs:               nil,
				Alias:             "commercetools",
				Name:              "commercetools",
				Followers:         "41K",
				Employees:         "501-1K",
				AssociatedMembers: "696",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:     "commercetools",
				Followers: "269",
				Verified:  true,
			},
			BlindProfile: domain.BlindProfile{
				Alias: "",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias: "",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-commercetools-EI_IE1304424.11,24.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/commercetools-Reviews-E1304424.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/commercetools-Jobs-E1304424.htm",
				Jobs:        "20",
				Reviews:     "185",
				Salaries:    "320",
				ReviewsRate: "3.5",
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
							Title:                "Back-End Engineer – Scala",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4175367766/",
							Date:                 mustDate("2025-03-07"),
							WithSalary:           false,
							Remote:               true,
						},
						{
							Title:                "Senior Scala Engineer",
							ShortDescription:     "Typelevel & Elasticsearch",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4172530640/",
							Date:                 mustDate("2025-03-11"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Senior Backend Engineer (Scala & Elasticsearch)",
							ShortDescription:     "5+ years experience as Software Engineer working with Scala",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4200235969/",
							Date:                 mustDate("2025-04-03"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Senior Scala Engineer – Distributed Systems",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4229801981/",
							Date:                 mustDate("2025-05-15"),
							WithSalary:           false,
							Remote:               true,
						},
						{
							Title:                "Senior Scala Engineer – Distributed Systems",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4256296165/",
							Date:                 mustDate("2025-06-30"),
							WithSalary:           false,
							Remote:               false,
						},
					},
				},
				domain.Elixir:  {},
				domain.Clojure: {},
				domain.Haskell: {},
			},
			ShortDescription: "Composable commerce",
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Net Health",
			Website: "https://www.nethealth.com/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                120271,
				IDs:               []int{90636, 120271},
				Alias:             "net-health-systems-inc-",
				Name:              "Net Health",
				Followers:         "24K",
				Employees:         "501-1K",
				AssociatedMembers: "591",
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
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Net-Health-EI_IE813969.11,21.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Net-Health-Reviews-E813969.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Net-Health-Jobs-E813969.htm",
				Jobs:        "11",
				Reviews:     "188",
				Salaries:    "352",
				ReviewsRate: "3.4",
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
							Title:                "Senior Elixir Developer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4148903233/",
							Date:                 mustDate("2025-03-12"), // mustDate("2025-03-06"),
							WithSalary:           true,                   // $104.000 - $130.000 per year
							Remote:               false,
						},
					},
				},
				domain.Clojure: {},
				domain.Haskell: {},
			},
			ShortDescription: "Specialized software and analytics that serves the continuum of restorative care, from hospital to home",
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Dscout",
			Website: "https://dscout.com/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                2438748,
				IDs:               nil,
				Alias:             "dscout",
				Name:              "Dscout",
				Followers:         "25K",
				Employees:         "201-500",
				AssociatedMembers: "255",
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
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-dscout-EI_IE1311226.11,17.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/dscout-Reviews-E1311226.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/dscout-Jobs-E1311226.htm",
				Jobs:        "7",
				Reviews:     "39",
				Salaries:    "124",
				ReviewsRate: "4.2",
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
							Title:                "Lead Software Engineer – Elixir",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4149944830/",
							Date:                 mustDate("2025-02-14"),
							WithSalary:           false,
							Remote:               true,
						},
					},
				},
				domain.Clojure: {},
				domain.Haskell: {},
			},
			ShortDescription: "Experience Research Platform",
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Viasat",
			Website: "https://www.viasat.com/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                5770,
				IDs:               nil,
				Alias:             "viasat",
				Name:              "Viasat",
				Followers:         "166K",
				Employees:         "5K-10K",
				AssociatedMembers: "6,382",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:     "Viasat",
				Followers: "30",
				Verified:  true,
			},
			BlindProfile: domain.BlindProfile{
				Alias:       "Viasat",
				Employees:   "5,001 to 10,000",
				Salary:      "$43K ~ $874K a year",
				Reviews:     "140",
				ReviewsRate: "3.6",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "viasat",
				Employees: "6,000",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Viasat-EI_IE5500.11,17.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Viasat-Reviews-E5500.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Viasat-Jobs-E5500.htm",
				Jobs:        "419",
				Reviews:     "1.9K",
				Salaries:    "3.6K",
				ReviewsRate: "3.6",
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
							Title:                "Embedded Software Engineer – Rust",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4207056896/",
							Date:                 mustDate("2025-05-07"), // mustDate("2025-05-03"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Embedded Software Engineer – Rust",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4207060632/",
							Date:                 mustDate("2025-05-07"),
							WithSalary:           false,
							Remote:               false,
						},
					},
				},
				domain.Zig:    {},
				domain.Scala:  {},
				domain.Elixir: {},
				domain.Clojure: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Senior Software Engineer (Java, Clojure)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4166326665/",
							Date:                 mustDate("2025-02-28"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Senior Software Engineer (Python, Clojure)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4213829549/",
							Date:                 mustDate("2025-05-10"), // mustDate("2025-04-21"),
							WithSalary:           false,
							Remote:               false,
						},
					},
				},
				domain.Haskell: {},
			},
			ShortDescription: "Telecommunications",
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Pitch",
			Website: "https://pitch.com/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                27222819,
				IDs:               nil,
				Alias:             "pitchhq",
				Name:              "Pitch",
				Followers:         "30K",
				Employees:         "51-200",
				AssociatedMembers: "64",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:     "pitch-io",
				Followers: "38",
				Verified:  false,
			},
			BlindProfile: domain.BlindProfile{
				Alias: "",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias: "",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Pitch-EI_IE2235698.11,16.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Pitch-Reviews-E2235698.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Pitch-Jobs-E2235698.htm",
				Jobs:        "1",
				Reviews:     "70",
				Salaries:    "79",
				ReviewsRate: "4.6",
				Verified:    true,
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
					GitHubRepositoriesCount: 5,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Senior Backend Clojure Engineer (m/f/d)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4128083196/",
							Date:                 mustDate("2025-02-07"),
							WithSalary:           false,
							Remote:               true,
						},
					},
				},
				domain.Haskell: {},
			},
			ShortDescription: "The complete pitching platform",
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Doccla",
			Website: "https://www.doccla.com/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                33432702,
				IDs:               nil,
				Alias:             "doccla",
				Name:              "Doccla",
				Followers:         "11K",
				Employees:         "51-200",
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
				domain.Go:    {},
				domain.Rust:  {},
				domain.Zig:   {},
				domain.Scala: {},
				domain.Elixir: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Senior Software Engineer – Elixir",
							ShortDescription:     "Design and implementation of RESTful microservices in Elixir",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4008290372/",
							Date:                 mustDate("2024-09-17"),
							WithSalary:           false,
							Remote:               true,
						},
					},
				},
				domain.Clojure: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Senior Software Engineer – Clojure",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4068248164/",
							Date:                 mustDate("2024-11-07"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Clojure Engineer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4228305138/",
							Date:                 mustDate("2025-05-13"),
							WithSalary:           false,
							Remote:               true,
						},
					},
				},
				domain.Haskell: {},
			},
			ShortDescription: "Supporting healthcare organisations deliver virtual ward services proven to reduce costs while improving outcomes",
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "University of Virginia",
			Website: "https://www.virginia.edu/",
			Careers: "https://jobs.virginia.edu/",
			About:   "https://www.virginia.edu/aboutuva/",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                4298,
				IDs:               nil,
				Alias:             "university-of-virginia",
				Name:              "University of Virginia",
				Followers:         "312K",
				Employees:         "10K+", // https://www.linkedin.com/search/results/people/?currentCompany=%5B%224298%22%5D
				AssociatedMembers: "15,000",
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
				Alias:     "university-of-virginia",
				Employees: "8,397",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-University-of-Virginia-EI_IE8336.11,33.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/University-of-Virginia-Reviews-E8336.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/University-of-Virginia-Jobs-E8336.htm",
				Jobs:        "579",
				Reviews:     "2.5K",
				Salaries:    "5.9K",
				ReviewsRate: "4.3",
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
							Title:                "Senior Golang Developer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4177332766/",
							Date:                 mustDate("2025-03-13"), // mustDate("2025-03-07"),
							WithSalary:           true,                   // $110,000 - $130,000 per year
							Remote:               false,
						},
						{
							Title:                "Senior Golang Developer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4205310300/",
							Date:                 mustDate("2025-04-10"),
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
			ShortDescription: "American public institution of higher learning",
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Phrase",
			Website: "https://phrase.com/",
			Careers: "https://phrase.com/careers/",
			About:   "https://phrase.com/about/",
			Blog:    "https://phrase.com/blog/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                1344091,
				IDs:               nil,
				Alias:             "phraseplatform",
				Name:              "Phrase",
				Followers:         "36K",
				Employees:         "201-500",
				AssociatedMembers: "374",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "phrase",
				Verified: true,
			},
			BlindProfile: domain.BlindProfile{
				Alias: "",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias: "",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Phrase-EI_IE1535968.11,17.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Phrase-Reviews-E1535968.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Phrase-Jobs-E1535968.htm",
				Jobs:        "7",
				Reviews:     "75",
				Salaries:    "126",
				ReviewsRate: "3.9",
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
							Title:                "Senior Full Stack (Elixir) Engineer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4176875904/",
							Date:                 mustDate("2025-03-09"),
							WithSalary:           false,
							Remote:               true,
						},
						{
							Title:                "Senior Full Stack (Elixir) Engineer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4214281866/",
							Date:                 mustDate("2025-04-25"),
							WithSalary:           false,
							Remote:               true,
						},
					},
				},
				domain.Clojure: {},
				domain.Haskell: {},
			},
			ShortDescription: "Czech AI-led translation technology",
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Mr Price Group",
			Website: "https://mrpricegroup.com/",
			Careers: "https://mrpcareers.com/our-careers/",
			About:   "https://mrpricegroup.com/about-us/",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                47016,
				IDs:               []int{47016, 3110091},
				Alias:             "mr-price-group",
				Name:              "Mr Price Group",
				Followers:         "622K",
				Employees:         "10K+",
				AssociatedMembers: "16,870",
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
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Mr-Price-EI_IE224003.11,19.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Mr-Price-Reviews-E224003.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Mr-Price-Jobs-E224003.htm",
				Jobs:        "118",
				Reviews:     "382",
				Salaries:    "438",
				ReviewsRate: "3.6",
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
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Lead Clojure Developer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4173567999/",
							Date:                 mustDate("2025-03-05"),
							WithSalary:           false,
							Remote:               true,
						},
					},
				},
				domain.Haskell: {},
			},
			ShortDescription: "SA omni-channel, fashion value retailer",
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "PlentyONE",
			Website: "https://www.plentyone.com/",
			Careers: "https://www.plentyone.com/careers",
			About:   "https://www.plentyone.com/about",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                5079248,
				IDs:               nil,
				Alias:             "plentyone-ecommerce",
				Name:              "PlentyONE",
				Followers:         "5K",
				Employees:         "201-500",
				AssociatedMembers: "237",
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
							Title:                "Golang Developer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4177522085/",
							Date:                 mustDate("2025-03-09"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Golang Developer",
							ShortDescription:     "Plentychannel",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4178761414/",
							Date:                 mustDate("2025-03-13"),
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
			ShortDescription: "E-Commerce ERP system for online retailers, brands and manufacturers",
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Future",
			Website: "https://www.future.co/",
			Careers: "https://www.future.co/careers",
			About:   "https://www.future.co/about",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                33275761,
				IDs:               nil,
				Alias:             "future-research-inc",
				Name:              "Future",
				Followers:         "13K",
				Employees:         "201-500",
				AssociatedMembers: "654",
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
							Title:                "Backend Engineer (Golang)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4179354991/",
							Date:                 mustDate("2025-03-10"),
							WithSalary:           true, // Salary range: $160,000 - 215,000 / year + equity
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
			ShortDescription: "Online fitness training platform",
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Los Alamos National Laboratory",
			Website: "https://www.lanl.gov/",
			Careers: "https://www.lanl.gov/careers",
			About:   "https://www.lanl.gov/about",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                5327,
				IDs:               nil,
				Alias:             "los-alamos-national-laboratory",
				Name:              "Los Alamos National Laboratory",
				Followers:         "157K",
				Employees:         "10K+",
				AssociatedMembers: "12,716",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:     "LANL",
				Followers: "487",
				Verified:  true,
			},
			BlindProfile: domain.BlindProfile{
				Alias: "",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "los-alamos-national-laboratory",
				Employees: "12,752",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Los-Alamos-National-Laboratory-EI_IE35234.11,41.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Los-Alamos-National-Laboratory-Reviews-E35234.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Los-Alamos-National-Laboratory-Jobs-E35234.htm",
				Jobs:        "345",
				Reviews:     "1.2K",
				Salaries:    "2.9K",
				ReviewsRate: "4.1",
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
							Title:                "DevOps Python/Rust Developer (Scientist 1/2)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4181750667/",
							Date:                 mustDate("2025-05-14"), // mustDate("2025-04-22"), // mustDate("2025-04-01"), // mustDate("2025-03-11"),
							WithSalary:           true,                   // Scientist 1 ($94,5-$154,600), Scientist 2 ($104,100-$172,200)
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
			ShortDescription: "American laboratory's who provide scientific and engineering support to national security programs",
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Volkswagen Digital Solutions",
			Website: "https://www.vwds.pt/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                19035234,
				IDs:               nil,
				Alias:             "volkswagen-digital-solutions",
				Name:              "Volkswagen Digital Solutions",
				Followers:         "47K",
				Employees:         "501-1K",
				AssociatedMembers: "694",
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
							Title:                "Go Developer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4157350713/",
							Date:                 mustDate("2025-03-11"),
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
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Binance",
			Website: "https://www.binance.com/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                13336409,
				IDs:               nil,
				Alias:             "binance",
				Name:              "Binance",
				Followers:         "824K",
				Employees:         "5K-10K",
				AssociatedMembers: "10,428",
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
							Title:                "Golang Developer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4179397668/",
							Date:                 mustDate("2025-03-11"),
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
			ShortDescription: "Digital Asset Exchange",
			Ignore:           true, // Cryptocurrency
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Crossover",
			Website: "https://www.crossover.com/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                9308035,
				IDs:               nil,
				Alias:             "crossover",
				Name:              "Crossover",
				Followers:         "8M",
				Employees:         "5K-10K",
				AssociatedMembers: "2,742",
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
				domain.Go:      {},
				domain.Rust:    {},
				domain.Zig:     {},
				domain.Scala:   {},
				domain.Elixir:  {},
				domain.Clojure: {},
				domain.Haskell: {},
			},
			ShortDescription: "Recruitment company",
			Ignore:           true,
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Optable",
			Website: "https://optable.co/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                64779622,
				IDs:               nil,
				Alias:             "optableco",
				Name:              "Optable",
				Followers:         "3K",
				Employees:         "11-50",
				AssociatedMembers: "57",
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
							Title:                "Experienced Backend Developer (Golang)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4178647999/",
							Date:                 mustDate("2025-03-11"),
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
			ShortDescription: "Optable is an identity management and data collaboration platform",
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "ScaleOps",
			Website: "https://scaleops.com/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                79471518,
				IDs:               nil,
				Alias:             "scaleops-sh",
				Name:              "ScaleOps - Cloud-Native Optimization",
				Followers:         "9K",
				Employees:         "51-200",
				AssociatedMembers: "62",
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
							Title:                "Senior Backend Engineer Go",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4181796179/",
							Date:                 mustDate("2025-03-11"),
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
			ShortDescription: "Effortlessly reduce cloud costs by up to 80% by automatically optimizing and scaling Kubernetes resources in runtime",
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "bet365",
			Website: "https://www.bet365careers.com/en",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                40163,
				IDs:               nil,
				Alias:             "bet365",
				Name:              "bet365",
				Followers:         "237K",
				Employees:         "5K-10K",
				AssociatedMembers: "6,619",
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
							Title:                "Software Engineer (Go), Sports Platform",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4179415502/",
							Date:                 mustDate("2025-03-11"),
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
			Ignore:           true, // Gambling
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Zenport",
			Website: "https://zenport.io/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                13725003,
				IDs:               nil,
				Alias:             "zenport",
				Name:              "Zenport Inc.",
				Followers:         "2K",
				Employees:         "11-50",
				AssociatedMembers: "10",
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
							Title:                "Senior Backend Developer (Go)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4181727305/",
							Date:                 mustDate("2025-03-12"),
							WithSalary:           false,
							Remote:               true,
						},
						{
							Title:                "Senior Backend Developer (Go)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4206443530/",
							Date:                 mustDate("2025-04-12"),
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
			ShortDescription: "",
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "FirstBatch",
			Website: "https://www.firstbatch.xyz/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                80364972,
				IDs:               nil,
				Alias:             "firstbatchai",
				Name:              "FirstBatch",
				Followers:         "6K",
				Employees:         "11-50",
				AssociatedMembers: "19",
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
							Title:                "Rust Backend Engineer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4178518636/",
							Date:                 mustDate("2025-03-11"),
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
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Syndica",
			Website: "https://www.syndica.io/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                76168224,
				IDs:               nil,
				Alias:             "syndica-io",
				Name:              "Syndica",
				Followers:         "2K",
				Employees:         "11-50",
				AssociatedMembers: "24",
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
							Title:                "Senior Rust Software Engineer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4179432496/",
							Date:                 mustDate("2025-03-11"),
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
			Ignore:           true, // Web3
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Nelly Solutions",
			Website: "https://www.getnelly.de/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                74522753,
				IDs:               nil,
				Alias:             "nelly-solutions",
				Name:              "Nelly Solutions",
				Followers:         "6K",
				Employees:         "51-200",
				AssociatedMembers: "132",
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
							Title:                "(Senior) Software Engineer, Rust",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4125742838/",
							Date:                 mustDate("2025-06-24"), // mustDate("2025-05-28"), // mustDate("2025-05-13"), // mustDate("2025-04-14"), // mustDate("2025-04-04"), // mustDate("2025-03-21"), // mustDate("2025-03-11"),
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
			ShortDescription: "Nelly is a healthcare payments company",
		},
	}
}
