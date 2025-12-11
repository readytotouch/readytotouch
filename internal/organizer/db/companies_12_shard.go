package db

import (
	"github.com/readytotouch/readytotouch/internal/domain"
)

func companies12Shard() []domain.CompanyProfile {
	return []domain.CompanyProfile{
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Polygon Labs",
			Website: "https://polygon.technology/",
			Careers: "https://polygon.technology/careers/",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                13449964,
				IDs:               nil,
				Alias:             "polygonlabs",
				Name:              "Polygon Labs",
				Followers:         "191K",
				Employees:         "201-500",
				AssociatedMembers: "447",
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
							Title:                "Senior Rust Engineer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4102719913/",
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
			Ignore:           true, // Blockchain
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Skyworks Solutions",
			Website: "https://www.skyworksinc.com/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                165998,
				IDs:               nil,
				Alias:             "skyworks-solutions-inc",
				Name:              "Skyworks Solutions, Inc.",
				Followers:         "82K",
				Employees:         "5K-10K",
				AssociatedMembers: "5,202",
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
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Skyworks-Solutions-EI_IE762.11,29.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Skyworks-Solutions-Reviews-E762.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Skyworks-Solutions-Jobs-E762.htm",
				Jobs:        "228",
				Reviews:     "866",
				Salaries:    "1.9K",
				ReviewsRate: "3.7",
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
							Title:                "Rust Software Engineer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4027089913/",
							Date:                 mustDate("2025-03-11"),
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
			ShortDescription: "American semiconductor company",
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "F5",
			Website: "https://www.f5.com/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                4841,
				IDs:               nil,
				Alias:             "f5",
				Name:              "F5",
				Followers:         "358K",
				Employees:         "5K-10K",
				AssociatedMembers: "6,141",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:     "f5networks",
				Followers: "280",
				Verified:  false,
			},
			BlindProfile: domain.BlindProfile{
				Alias:       "F5-Networks",
				Employees:   "1,001 to 5,000",
				Salary:      "$71K ~ $265K a year",
				Reviews:     "292",
				ReviewsRate: "3.5",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "f5-networks",
				Employees: "6,500",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-F5-EI_IE9222.11,13.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/F5-Reviews-E9222.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/F5-Jobs-E9222.htm",
				Jobs:        "173",
				Reviews:     "2K",
				Salaries:    "3.9K",
				ReviewsRate: "3.8",
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
							ShortDescription:     "Proxy Solution",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4153485948/",
							Location:             "San Jose, CA",
							Date:                 mustDate("2025-10-30", "2025-10-10", "2025-09-18", "2025-08-26", "2025-08-06", "2025-07-14", "2025-06-24", "2025-06-02", "2025-04-20", "2025-03-31", "2025-03-09"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Senior Rust Developer",
							ShortDescription:     "Proxy Solution",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4190616076/",
							Date:                 mustDate("2025-03-21"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Senior Rust Developer",
							ShortDescription:     "Proxy Solution",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4198709559/",
							Date:                 mustDate("2025-05-17", "2025-04-25", "2025-04-06"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Senior Rust Developer",
							ShortDescription:     "Proxy Solution",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4250693044/",
							Date:                 mustDate("2025-06-14"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Principal Software Engineer (Rust)",
							ShortDescription:     "Distributed Cloud",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4283983721/",
							Location:             "Greater Seattle Area",
							Date:                 mustDate("2025-11-26", "2025-11-06"),
							WithSalary:           true, // $186.4k/yr - $279.6k/yr
							Remote:               false,
						},
						{
							Title:                "Principal Rust Developer",
							ShortDescription:     "Gateway Solutions",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4341814563/",
							Location:             "Greater Seattle Area",
							Date:                 mustDate("2025-11-20"),
							WithSalary:           true, // $186.4k/yr - $279.6k/yr
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
			ShortDescription: "F5 is a multi-cloud application services and security company",
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Megaport",
			Website: "https://www.megaport.com/",
			Careers: "https://www.megaport.com/careers/",
			About:   "https://www.megaport.com/about-megaport/",
			Blog:    "https://www.megaport.com/blog/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                3225864,
				IDs:               nil,
				Alias:             "megaport",
				Name:              "Megaport",
				Followers:         "32K",
				Employees:         "201-500",
				AssociatedMembers: "377",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "megaport",
				Verified: true,
			},
			BlindProfile: domain.BlindProfile{
				Alias: "",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias: "",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Megaport-EI_IE1022619.11,19.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Megaport-Reviews-E1022619.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Megaport-Jobs-E1022619.htm",
				Jobs:        "32",
				Reviews:     "124",
				Salaries:    "142",
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
							Title:                "Software Engineer (Go)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4181511391/",
							Date:                 mustDate("2025-03-13"),
							WithSalary:           false,
							Remote:               true,
						},
						{
							Title:                "Software Engineer (Go)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4324861462/",
							Location:             "Australia",
							Date:                 mustDate("2025-11-26"),
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
			ShortDescription: "Australian AWS Direct Connect eco-system",
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "FairMoney",
			Website: "https://fairmoney.io/",
			Careers: "",
			About:   "https://fairmoney.io/about/",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                11707115,
				IDs:               nil,
				Alias:             "fairmoney",
				Name:              "FairMoney",
				Followers:         "142K",
				Employees:         "501-1K",
				AssociatedMembers: "1,029",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:     "fairmoney",
				Followers: "37",
				Verified:  false,
			},
			BlindProfile: domain.BlindProfile{
				Alias: "",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "fairmoney",
				Employees: "570",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-FairMoney-EI_IE3265012.11,20.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/FairMoney-Reviews-E3265012.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/FairMoney-Jobs-E3265012.htm",
				Jobs:        "20",
				Reviews:     "78",
				Salaries:    "94",
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
							Title:                "Software Engineer — Backend (Golang)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4178759491/",
							Location:             "Istanbul, Istanbul, Türkiye",
							Date:                 mustDate("2025-07-21", "2025-06-07", "2025-05-15", "2025-04-24", "2025-04-02", "2025-03-12"),
							WithSalary:           true, // €48k/yr - €54k/yr
							Remote:               true,
						},
						{
							Title:                "Software Engineer — Backend (Golang)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4283545481/",
							Location:             "Istanbul, Istanbul, Türkiye",
							Date:                 mustDate("2025-09-04"),
							WithSalary:           true, // €48k/yr - €54k/yr
							Remote:               true,
						},
						{
							Title:                "Golang Backend Engineer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4323540300/",
							Location:             "Nigeria",
							Date:                 mustDate("2025-11-20"),
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
			ShortDescription: "Nigerian mobile banking for private and business borrowers",
			Industries: []domain.Industry{
				domain.IndustryFinTech,
			},
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Zalopay",
			Website: "https://zalopay.vn/",
			Careers: "",
			About:   "",
			Blog:    "https://engineering.zalopay.vn/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                31297705,
				IDs:               nil,
				Alias:             "zalopay",
				Name:              "Zalopay",
				Followers:         "44K",
				Employees:         "501-1K",
				AssociatedMembers: "406",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:     "zalopay-oss",
				Followers: "234",
				Verified:  false,
			},
			BlindProfile: domain.BlindProfile{
				Alias: "",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias: "",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-ZaloPay-EI_IE4832300.11,18.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/ZaloPay-Reviews-E4832300.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/ZaloPay-Jobs-E4832300.htm",
				Jobs:        "",
				Reviews:     "36",
				Salaries:    "41",
				ReviewsRate: "4.0",
				Verified:    false,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "",
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 7,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Senior Software Engineer — Golang",
							ShortDescription:     "Merchant Platform",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4180086053/",
							Date:                 mustDate("2025-03-12"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Senior Software Engineer (Golang / Java)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4214926656/",
							Date:                 mustDate("2025-04-24"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Software Engineer (Golang)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4231137930/",
							Date:                 mustDate("2025-05-21"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Middle/Senior Back-End Developer (Go)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4261894568/",
							Date:                 mustDate("2025-07-08"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Golang Developer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4280807342/",
							Location:             "Ho Chi Minh City Metropolitan Area",
							Date:                 mustDate("2025-08-06"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Senior Software Engineer (Golang)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4281582959/",
							Location:             "Ho Chi Minh City, Vietnam",
							Date:                 mustDate("2025-08-09"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Senior Golang Developer",
							ShortDescription:     "Platform",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4290559770/",
							Location:             "Ho Chi Minh City, Vietnam",
							Date:                 mustDate("2025-08-25"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Golang Developer (Junior/Middle/Senior)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4292479120/",
							Location:             "Ho Chi Minh City, Vietnam",
							Date:                 mustDate("2025-09-10"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Senior Golang Developer",
							ShortDescription:     "Platform Team",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4319286689/",
							Location:             "Ho Chi Minh City Metropolitan Area",
							Date:                 mustDate("2025-10-29"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "(Middle/Senior) Software Engineer (Golang)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4323023617/",
							Location:             "Ho Chi Minh City Metropolitan Area",
							Date:                 mustDate("2025-12-10", "2025-11-19"),
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
			ShortDescription: "Vietnamese mobile wallet and payment platform associated with the Zalo messaging app",
			Industries: []domain.Industry{
				domain.IndustryFinTech,
			},
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "TD",
			Website: "https://www.td.com/",
			Careers: "https://careers.td.com/",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                2775,
				IDs:               nil,
				Alias:             "td",
				Name:              "TD",
				Followers:         "900K",
				Employees:         "10K+",
				AssociatedMembers: "97,091",
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
				Alias:     "td-bank",
				Employees: "26.000",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-TD-EI_IE3767.11,13.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/TD-Reviews-E3767.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/TD-Jobs-E3767.htm",
				Jobs:        "1.3K",
				Reviews:     "23K",
				Salaries:    "35K",
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
							Title:                "Software Engineer II (Scala/Spark)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4182972537/",
							Date:                 mustDate("2025-03-12"),
							WithSalary:           true, // $76,800 - $115,200 CAD per year
							Remote:               false,
						},
						{
							Title:                "Senior Data Engineer (Scala, Spark, Python, Hadoop)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4266360659/",
							Location:             "Toronto, ON",
							Date:                 mustDate("2025-08-06", "2025-07-14"),
							WithSalary:           false,
							Remote:               false,
						},
					},
				},
				domain.Elixir:  {},
				domain.Clojure: {},
				domain.Haskell: {},
			},
			ShortDescription: "Canadian bank who offers a full range of financial products and services",
			Industries: []domain.Industry{
				domain.IndustryFinTech,
			},
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "AI Rudder",
			Website: "https://airudder.com/",
			Careers: "https://airudder.com/careers/",
			About:   "https://airudder.com/about-us/",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                47635489,
				IDs:               nil,
				Alias:             "airudder",
				Name:              "AI Rudder",
				Followers:         "16K",
				Employees:         "51-200",
				AssociatedMembers: "175",
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
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-AI-Rudder-EI_IE5023654.11,20.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/AI-Rudder-Reviews-E5023654.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/AI-Rudder-Jobs-E5023654.htm",
				Jobs:        "",
				Reviews:     "21",
				Salaries:    "17",
				ReviewsRate: "3.2",
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
							Title:                "Software Engineer (Go/Golang)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4184051342/",
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
			ShortDescription: "AI startup who provides AI-powered voice solutions to improve B2C communications",
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Bitpanda",
			Website: "https://www.bitpanda.com/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                18164565,
				IDs:               nil,
				Alias:             "bitpanda",
				Name:              "Bitpanda",
				Followers:         "68K",
				Employees:         "501-1K",
				AssociatedMembers: "744",
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
							Title:                "Senior Software Engineer, Golang (Digital Asset Infrastructure)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4168215806/",
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
			ShortDescription: "",
			Ignore:           true, // Crypto
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Guidewire Software",
			Website: "https://www.guidewire.com/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                9624,
				IDs:               nil,
				Alias:             "guidewire-software",
				Name:              "Guidewire Software",
				Followers:         "261K",
				Employees:         "1K-5K",
				AssociatedMembers: "3,597",
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
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Guidewire-EI_IE122537.11,20.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Guidewire-Reviews-E122537.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Guidewire-Jobs-E122537.htm",
				Jobs:        "172",
				Reviews:     "1.5K",
				Salaries:    "2.5K",
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
							Title:                "Software Engineer III (Golang)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4184033707/",
							Date:                 mustDate("2025-03-13"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "API Engineer — Micro-Services",
							ShortDescription:     "Maintain existing codebases in languages like Go",
							SwitchingOpportunity: "Knowledge of modern programming languages (specifically Go is a plus)",
							URL:                  "https://www.indeed.com/viewjob?jk=8d7ce4fc383b9eac",
							Date:                 mustDate("2025-03-25"),
							WithSalary:           true, // $134,000 - $247,000 per year
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
			Name:    "Makro PRO",
			Website: "https://www.makro.pro/",
			Careers: "https://www.makrodigitalcareer.com/",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                76238800,
				IDs:               nil,
				Alias:             "makropro",
				Name:              "Makro PRO",
				Followers:         "67K",
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
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Senior Software Engineer (Golang)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4138748551/",
							Date:                 mustDate("2025-03-13"),
							WithSalary:           false,
							Remote:               true,
						},
						{
							Title:                "Senior DevOps Engineer (Golang)",
							ShortDescription:     "Strong proficiency in Golang and experience with writing production-grade software using Golang",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4154050315/",
							Date:                 mustDate("2025-04-02"),
							WithSalary:           false,
							Remote:               true,
						},
						{
							Title:                "Senior Software Engineer (Golang)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4138747697/",
							Date:                 mustDate("2025-04-05"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Senior DevOps Engineer (Golang)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4154045928/",
							Location:             "Philippines",
							Date:                 mustDate("2025-07-19", "2025-06-05"),
							WithSalary:           false,
							Remote:               true,
						},
						{
							Title:                "Senior DevOps Engineer (Golang)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4154050314/",
							Location:             "Bangladesh",
							Date:                 mustDate("2025-07-16"),
							WithSalary:           false,
							Remote:               true,
						},
						{
							Title:                "Senior DevOps Engineer (Golang)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4279851073/",
							Location:             "Indonesia",
							Date:                 mustDate("2025-08-09"),
							WithSalary:           false,
							Remote:               true,
						},
						{
							Title:                "Senior Software Engineer (Golang)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4282248862/",
							Location:             "Bangkok, Bangkok City, Thailand",
							Date:                 mustDate("2025-10-18", "2025-09-04"),
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
			Name:    "Infoblox",
			Website: "https://www.infoblox.com/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                8697,
				IDs:               nil,
				Alias:             "infoblox",
				Name:              "Infoblox",
				Followers:         "172K",
				Employees:         "1K-5K",
				AssociatedMembers: "2,608",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:     "infobloxopen",
				Followers: "120",
				Verified:  false,
			},
			BlindProfile: domain.BlindProfile{
				Alias: "",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias: "",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Infoblox-EI_IE35108.11,19.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Infoblox-Reviews-E35108.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Infoblox-Jobs-E35108.htm",
				Jobs:        "130",
				Reviews:     "985",
				Salaries:    "1.2K",
				ReviewsRate: "4.3",
				Verified:    true,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "",
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 67,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Senior Software Engineer (Golang)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4183852152/",
							Date:                 mustDate("2025-03-13"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Staff Software Engineer Golang",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4299531568/",
							Location:             "Bengaluru, Karnataka, India",
							Date:                 mustDate("2025-10-08"),
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
			Name:    "NCR Voyix",
			Website: "https://www.ncrvoyix.com/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                96157938,
				IDs:               nil,
				Alias:             "ncrvoyix",
				Name:              "NCR Voyix",
				Followers:         "112K",
				Employees:         "10K+",
				AssociatedMembers: "6,668",
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
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-NCR-Voyix-EI_IE10016304.11,20.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/NCR-Voyix-Reviews-E10016304.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/NCR-Voyix-Jobs-E10016304.htm",
				Jobs:        "",
				Reviews:     "164",
				Salaries:    "315",
				ReviewsRate: "3.1",
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
							Title:                "Software Engineer II (Golang – Kubernetes)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4137578941/",
							Date:                 mustDate("2025-03-13"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Senior Software Engineer – Golang",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4071755601/",
							Date:                 mustDate("2025-03-19"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Software Engineer II – Edge (Golang Kubernetes)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4145763846/",
							Date:                 mustDate("2025-05-05", "2025-03-22"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Golang Software Engineer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4210732845/",
							Date:                 mustDate("2025-04-21"),
							WithSalary:           false,
							Remote:               true,
						},
						{
							Title:                "Software Engineer – Golang",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4220344158/",
							Date:                 mustDate("2025-07-06"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Lead Software Engineer Golang",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4265482835/",
							Location:             "Atlanta, GA",
							Date:                 mustDate("2025-07-15"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Software Engineer III – Golang Kubernetes",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4145767451/",
							Location:             "Hyderabad, Telangana, India",
							Date:                 mustDate("2025-09-13", "2025-08-22"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Software Engineer I – Golang Developer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4304104920/",
							Location:             "Hyderabad, Telangana, India",
							Date:                 mustDate("2025-10-15", "2025-09-22"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Software Engineer III – Golang",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4247512353/",
							Location:             "Hyderabad, Telangana, India",
							Date:                 mustDate("2025-11-08"),
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
			Name:    "Cisco",
			Website: "https://www.cisco.com/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                1063,
				IDs:               []int{1063, 3805, 11434, 92950, 687352, 2792574, 3517498, 72566046},
				Alias:             "cisco",
				Name:              "Cisco",
				Followers:         "7M",
				Employees:         "10K+",
				AssociatedMembers: "95,120",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:     "cisco",
				Followers: "986",
				Verified:  true,
			},
			BlindProfile: domain.BlindProfile{
				Alias:       "cisco",
				Employees:   "10,000+",
				Salary:      "$50K ~ $260K a year",
				Reviews:     "2,519",
				ReviewsRate: "3.6",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "cisco",
				Employees: "98,110",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Cisco-EI_IE1425.11,16.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Cisco-Reviews-E1425.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Cisco-Jobs-E1425.htm",
				Jobs:        "428",
				Reviews:     "38K",
				Salaries:    "113K",
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
							Title:                "Software Engineer – FullStack | Golang | Java",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4183841211/",
							Date:                 mustDate("2025-03-13"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Senior Software Engineer — Go",
							ShortDescription:     "3+ years of experience in Go",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4198239895/",
							Date:                 mustDate("2025-04-28", "2025-04-24", "2025-04-01"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Software Engineer – Golang Developer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4189691197/",
							Date:                 mustDate("2025-04-12"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Senior Software Engineer",
							ShortDescription:     "Developing high-quality Go and eBPF code for Cilium OSS and Enterprise",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4176497487/",
							Date:                 mustDate("2025-04-18"),
							WithSalary:           true,
							Remote:               true,
						},
						{
							Title:                "Software Engineer – Golang Developer | DevOps",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4242589805/",
							Date:                 mustDate("2025-07-02"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Software Engineer – Golang",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4310896069/",
							Location:             "Bengaluru, Karnataka, India",
							Date:                 mustDate("2025-10-07"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "AI Infrastructure Engineer – Golang",
							ShortDescription:     "Kubernetes",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4337079569/",
							Location:             "Bengaluru, Karnataka, India",
							Date:                 mustDate("2025-11-20"),
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
			Name:    "Priority Software",
			Website: "https://www.priority-software.com/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                89771,
				IDs:               nil,
				Alias:             "prioritysoftware",
				Name:              "Priority Software",
				Followers:         "19K",
				Employees:         "201-500",
				AssociatedMembers: "405",
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
							Title:                "Golang Senior Developer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4178758249/",
							Date:                 mustDate("2025-03-12"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Golang Senior Developer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4203313137/",
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
			ShortDescription: "",
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Aspire Software",
			Website: "https://www.aspiresoftware.com/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                52129049,
				IDs:               nil,
				Alias:             "aspire-software",
				Name:              "Aspire Software",
				Followers:         "36K",
				Employees:         "1K-5K",
				AssociatedMembers: "248",
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
							Title:                "Senior Software Developer – Golang",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4182723217/",
							Date:                 mustDate("2025-04-24", "2025-03-12"),
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
			Name:    "Valsoft Corporation",
			Website: "https://www.valsoftcorp.com/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                10396649,
				IDs:               nil,
				Alias:             "valsoft-corporation",
				Name:              "Valsoft Corporation",
				Followers:         "60K",
				Employees:         "1K-5K",
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
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Valsoft-EI_IE2564742.11,18.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Valsoft-Reviews-E2564742.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Valsoft-Jobs-E2564742.htm",
				Jobs:        "70",
				Reviews:     "167",
				Salaries:    "191",
				ReviewsRate: "3.4",
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
							Title:                "Senior Software Developer – Golang",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4182769280/",
							Date:                 mustDate("2025-03-12"),
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
			Name:    "Flutter Entertainment",
			Website: "https://www.flutter.com/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                14072241,
				IDs:               nil,
				Alias:             "flutter-entertainment",
				Name:              "Flutter Entertainment",
				Followers:         "64K",
				Employees:         "10K+",
				AssociatedMembers: "7,847",
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
							Title:                "Software Engineer (Scala)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4158759621/",
							Date:                 mustDate("2025-03-13"),
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
			Ignore:           true, // Gambling
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "NielsenIQ",
			Website: "https://nielseniq.com/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                69088863,
				IDs:               []int{1718, 419609, 1364767, 3126409, 18360338, 69088863},
				Alias:             "nielseniq",
				Name:              "NielsenIQ",
				Followers:         "1M",
				Employees:         "10K+",
				AssociatedMembers: "25,585",
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
				Alias:     "nielseniq",
				Employees: "14,430",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-NielsenIQ-EI_IE4796509.11,20.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/NielsenIQ-Reviews-E4796509.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/NielsenIQ-Jobs-E4796509.htm",
				Jobs:        "621",
				Reviews:     "4.4K",
				Salaries:    "7.7K",
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
							Title:                "Data Engineer (Python, Scala, Azure)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4181286159/",
							Date:                 mustDate("2025-03-14"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Data Engineer (Scala, Databricks, Azure)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4232688139/",
							Date:                 mustDate("2025-05-22"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Senior Data Engineer – Scala",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4322148530/",
							Location:             "Chennai, Tamil Nadu, India",
							Date:                 mustDate("2025-11-14"),
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
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Adform",
			Website: "https://adform.com/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                72240,
				IDs:               nil,
				Alias:             "adform",
				Name:              "Adform",
				Followers:         "42K",
				Employees:         "501-1K",
				AssociatedMembers: "858",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:     "adform",
				Followers: "9",
				Verified:  false,
			},
			BlindProfile: domain.BlindProfile{
				Alias: "",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias: "Adform",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Adform-EI_IE373601.11,17.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Adform-Reviews-E373601.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Adform-Jobs-E373601.htm",
				Jobs:        "14",
				Reviews:     "193",
				Salaries:    "280",
				ReviewsRate: "3.9",
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
					GitHubRepositoriesCount: 1,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Software Engineer, Java/Scala",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/3885026236/",
							Date:                 mustDate("2025-03-13"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Java/Scala Developer",
							ShortDescription:     "Spark & Hadoop",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4120994368/",
							Date:                 mustDate("2025-06-27", "2025-06-05", "2025-05-15", "2025-04-01"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Java/Scala Developer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4285962926/",
							Location:             "Mumbai, Maharashtra, India",
							Date:                 mustDate("2025-09-06", "2025-08-15"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Java/Scala Developer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4316070713/",
							Location:             "Mumbai, Maharashtra, India",
							Date:                 mustDate("2025-11-30"),
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
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Avathon",
			Website: "https://www.avathon.com/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                5155679,
				IDs:               nil,
				Alias:             "avathonai",
				Name:              "Avathon",
				Followers:         "45K",
				Employees:         "201-500",
				AssociatedMembers: "302",
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
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Avathon-EI_IE1343120.11,18.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Avathon-Reviews-E1343120.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Avathon-Jobs-E1343120.htm",
				Jobs:        "34",
				Reviews:     "172",
				Salaries:    "347",
				ReviewsRate: "3.1",
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
							Title:                "Software Engineer (Scala Back-End)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4053922033/",
							Date:                 mustDate("2025-04-22", "2025-03-11"),
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
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Kraken Digital Asset Exchange",
			Website: "https://kraken.com/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                3787845,
				IDs:               []int{3787845, 69222723},
				Alias:             "krakenfx",
				Name:              "Kraken Digital Asset Exchange",
				Followers:         "208K",
				Employees:         "1K-5K",
				AssociatedMembers: "851",
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
							Title:                "Senior Software Engineer – Rust",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4137684962/",
							Date:                 mustDate("2025-03-14"),
							WithSalary:           true,
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
			Ignore:           true, // Crypto
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Cryptio",
			Website: "https://cryptio.co/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                11423070,
				IDs:               nil,
				Alias:             "cryptio",
				Name:              "Cryptio",
				Followers:         "8K",
				Employees:         "11-50",
				AssociatedMembers: "97",
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
							Title:                "Senior Rust Engineer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4183189232/",
							Date:                 mustDate("2025-03-13"),
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
			Ignore:           true, // Crypto
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "RingCentral",
			Website: "https://www.ringcentral.com/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                60868,
				IDs:               nil,
				Alias:             "ringcentral",
				Name:              "RingCentral",
				Followers:         "254K",
				Employees:         "5K-10K",
				AssociatedMembers: "6,328",
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
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-RingCentral-EI_IE197577.11,22.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/RingCentral-Reviews-E197577.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/RingCentral-Jobs-E197577.htm",
				Jobs:        "184",
				Reviews:     "2K",
				Salaries:    "2.7K",
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
							Title:                "Rust developer",
							ShortDescription:     "Core Media AI",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4168118534/",
							Date:                 mustDate("2025-03-13"),
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
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "QNX",
			Website: "https://blackberry.qnx.com/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                9889,
				IDs:               nil,
				Alias:             "qnx",
				Name:              "QNX",
				Followers:         "26K",
				Employees:         "201-500",
				AssociatedMembers: "704",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:     "qnx",
				Followers: "107",
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
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Rust Systems Software Developer II",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4151541268/",
							Date:                 mustDate("2025-04-08", "2025-03-31", "2025-03-11"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Senior Systems Software Developer (Rust)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4151540281/",
							Date:                 mustDate("2025-04-11"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Rust C C++ Developer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4304996055/",
							Location:             "Greater Ottawa Metropolitan Area",
							Date:                 mustDate("2025-09-23"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Senior Rust Developer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4334667902/",
							Location:             "Greater Ottawa Metropolitan Area",
							Date:                 mustDate("2025-11-08"),
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
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Solaris",
			Website: "https://www.solarisgroup.com/",
			Careers: "https://www.solarisgroup.com/careers/",
			About:   "https://www.solarisgroup.com/about/",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                10533980,
				IDs:               nil,
				Alias:             "solariscompany",
				Name:              "Solaris SE",
				Followers:         "55K",
				Employees:         "201-500",
				AssociatedMembers: "683",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:     "solarisbank",
				Followers: "21",
				Verified:  false,
			},
			BlindProfile: domain.BlindProfile{
				Alias: "",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias: "",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Solaris-Group-EI_IE1618625.11,24.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Solaris-Group-Reviews-E1618625.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Solaris-Group-Jobs-E1618625.htm",
				Jobs:        "1",
				Reviews:     "261",
				Salaries:    "499",
				ReviewsRate: "2.9",
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
							Title:                "Senior Engineer Golang",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4181915311/",
							Date:                 mustDate("2025-03-14"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Senior Engineer Golang",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4189553947/",
							Location:             "Chennai, Tamil Nadu, India",
							Date:                 mustDate("2025-08-02", "2025-07-10", "2025-05-28", "2025-05-07", "2025-04-14", "2025-03-24"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Professional Engineer Golang",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4196521310/",
							Date:                 mustDate("2025-06-07", "2025-05-15", "2025-04-24"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Senior Engineer Golang",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4297773165/",
							Location:             "Chennai, Tamil Nadu, India",
							Date:                 mustDate("2025-10-05", "2025-09-13"),
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
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "smava",
			Website: "https://www.smava.de/",
			Careers: "https://www.smava.de/jobs",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                643119,
				IDs:               nil,
				Alias:             "smava",
				Name:              "smava",
				Followers:         "7K",
				Employees:         "501-1K",
				AssociatedMembers: "299",
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
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-smava-EI_IE870424.11,16.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/smava-Reviews-E870424.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/smava-Jobs-E870424.htm",
				Jobs:        "83",
				Reviews:     "173",
				Salaries:    "240",
				ReviewsRate: "3.1",
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
							Title:                "Senior Software Engineer – Golang & NodeJS",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4184824435/",
							Date:                 mustDate("2025-05-14", "2025-03-14"),
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
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "CoinGate",
			Website: "https://coingate.com/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                10586817,
				IDs:               nil,
				Alias:             "coingate",
				Name:              "CoinGate",
				Followers:         "9K",
				Employees:         "51-200",
				AssociatedMembers: "81",
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
							Title:                "Senior Golang Engineer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4184624827/",
							Date:                 mustDate("2025-03-14"),
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
			Ignore:           true, // Crypto
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Kiln",
			Website: "https://www.kiln.fi/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                28150174,
				IDs:               nil,
				Alias:             "kiln-fi",
				Name:              "Kiln",
				Followers:         "8K",
				Employees:         "51-200",
				AssociatedMembers: "124",
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
							Title:                "Senior Backend Engineer – Golang",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4184623805/",
							Date:                 mustDate("2025-03-14"),
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
			Ignore:           true, // Crypto
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Syniti",
			Website: "https://www.syniti.com/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                19112642,
				IDs:               []int{2724939, 10417630, 19112642, 69646738},
				Alias:             "synitidata",
				Name:              "Syniti",
				Followers:         "67K",
				Employees:         "1K-5K",
				AssociatedMembers: "1,452",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "",
				Verified: false,
			},
			BlindProfile: domain.BlindProfile{
				Alias:       "Syniti",
				Employees:   "1,001 to 5,000",
				Salary:      "",
				Reviews:     "3",
				ReviewsRate: "4.3",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "syniti",
				Employees: "1,500",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Syniti-EI_IE152790.11,17.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Syniti-Reviews-E152790.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Syniti-Jobs-E152790.htm",
				Jobs:        "24",
				Reviews:     "388",
				Salaries:    "602",
				ReviewsRate: "4.2",
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
							Title:                "Software Engineer ll (Golang)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4181840665/",
							Date:                 mustDate("2025-03-14"),
							WithSalary:           false,
							Remote:               true,
						},
						{
							Title:                "Software Engineer ll (Golang)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4223656767/",
							Date:                 mustDate("2025-05-08"),
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
			Name:    "Ringover",
			Website: "https://www.ringover.com/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                79585140,
				IDs:               []int{9180058, 79583448, 79585140},
				Alias:             "ringover",
				Name:              "Ringover North America",
				Followers:         "9K",
				Employees:         "201-500",
				AssociatedMembers: "386",
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
							Title:                "Software Engineer Golang (Senior)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4184269668/",
							Date:                 mustDate("2025-03-14"),
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
			Ignore:           true, // Waiting for the English version of the vacancy description
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Goldman Sachs",
			Website: "https://www.goldmansachs.com/",
			Careers: "https://www.goldmansachs.com/careers",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                1382,
				IDs:               nil,
				Alias:             "goldman-sachs",
				Name:              "Goldman Sachs",
				Followers:         "5M",
				Employees:         "10K+",
				AssociatedMembers: "60,944",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:     "goldmansachs",
				Followers: "1.5k",
				Verified:  false,
			},
			BlindProfile: domain.BlindProfile{
				Alias:       "goldman-sachs",
				Employees:   "10,000+",
				Salary:      "$33K ~ $295K a year",
				Reviews:     "1,576",
				ReviewsRate: "3.4",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "goldman-sachs",
				Employees: "67,370",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Goldman-Sachs-EI_IE2800.11,24.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Goldman-Sachs-Reviews-E2800.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Goldman-Sachs-Jobs-E2800.htm",
				Jobs:        "361",
				Reviews:     "23K",
				Salaries:    "54K",
				ReviewsRate: "3.8",
				Verified:    true,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "",
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 8,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Golang Software Engineer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4180220865/",
							Date:                 mustDate("2025-03-13"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Software Engineer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.indeed.com/viewjob?jk=7922a4920a8d87b5",
							Date:                 mustDate("2025-03-23"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Core Engineering, Golang Software Engineer, Analyst/ Associate",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4124674198/",
							Date:                 mustDate("2025-05-05", "2025-04-12", "2025-03-22"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Golang Software Engineer",
							ShortDescription:     "Core Engineering",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4247882590/",
							Date:                 mustDate("2025-07-01", "2025-06-11"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Go Developer (with React/Angular)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4301305644/",
							Location:             "Warsaw, Mazowieckie, Poland",
							Date:                 mustDate("2025-12-10"), // mustDate("2025-11-20", "2025-10-30", "2025-10-08", "2025-09-16"),
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
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Chợ Tốt",
			Website: "https://chotot.com/",
			Careers: "https://careers.chotot.com/",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                3780320,
				IDs:               []int{3780320, 78359837},
				Alias:             "cho-tot",
				Name:              "Chợ Tốt",
				Followers:         "28K",
				Employees:         "201-500",
				AssociatedMembers: "232",
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
							Title:                "Backend Engineer (Java, Golang, Python)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4178756795/",
							Date:                 mustDate("2025-04-24", "2025-03-12"),
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
			Name:    "Simulmedia",
			Website: "https://www.simulmedia.com/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                237095,
				IDs:               nil,
				Alias:             "simulmedia",
				Name:              "Simulmedia",
				Followers:         "13K",
				Employees:         "51-200",
				AssociatedMembers: "89",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:     "Simulmedia",
				Followers: "5",
				Verified:  false,
			},
			BlindProfile: domain.BlindProfile{
				Alias: "",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "simulmedia",
				Employees: "126",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Simulmedia-EI_IE722501.11,21.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Simulmedia-Reviews-E722501.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Simulmedia-Jobs-E722501.htm",
				Jobs:        "20",
				Reviews:     "40",
				Salaries:    "83",
				ReviewsRate: "1.9",
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
							Title:                "Senior Golang Engineer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4182751037/",
							Date:                 mustDate("2025-03-12"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Senior Golang Engineer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4314614526/",
							Location:             "Greater Drohobych Area",
							Date:                 mustDate("2025-11-26", "2025-11-05", "2025-10-15"),
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
			ShortDescription: "Simulmedia is a cross-channel TV advertising platform",
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Mindrift",
			Website: "https://mindrift.ai/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                101576014,
				IDs:               nil,
				Alias:             "mindrift-ai",
				Name:              "Mindrift",
				Followers:         "380K",
				Employees:         "501-1K",
				AssociatedMembers: "1,135",
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
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Mindrift-EI_IE9951816.11,19.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Mindrift-Reviews-E9951816.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Mindrift-Jobs-E9951816.htm",
				Jobs:        "1",
				Reviews:     "35",
				Salaries:    "25",
				ReviewsRate: "3.4",
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
							Title:                "Software Developer (Golang)",
							ShortDescription:     "AI Tutor",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4182668629/",
							Date:                 mustDate("2025-03-14"),
							WithSalary:           false,
							Remote:               true,
						},
						{
							Title:                "Software Developer (Golang)",
							ShortDescription:     "AI Tutor",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4226665425/",
							Date:                 mustDate("2025-05-10"),
							WithSalary:           false,
							Remote:               true,
						},
					},
				},
				domain.Rust: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Software Developer (Rust)",
							ShortDescription:     "AI Tutor",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4182671096/",
							Date:                 mustDate("2025-04-16", "2025-04-01"),
							WithSalary:           false,
							Remote:               true,
						},
						{
							Title:                "Software Developer (Rust)",
							ShortDescription:     "AI Tutor",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4211032583/",
							Date:                 mustDate("2025-04-17"),
							WithSalary:           false,
							Remote:               true,
						},
						{
							Title:                "Software Developer (Rust)",
							ShortDescription:     "AI Tutor",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4205583730/",
							Date:                 mustDate("2025-04-22"),
							WithSalary:           false,
							Remote:               true,
						},
						{
							Title:                "Software Developer (Rust)",
							ShortDescription:     "AI Tutor",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4226390680/",
							Date:                 mustDate("2025-05-09"),
							WithSalary:           false,
							Remote:               true,
						},
						{
							Title:                "Software Developer (Rust)",
							ShortDescription:     "AI Tutor",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4235390623/",
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
			ShortDescription: "",
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Hewlett Packard Enterprise",
			Website: "https://hpe.com/",
			Careers: "https://careers.hpe.com/",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                1025,
				IDs:               []int{1025, 162533},
				Alias:             "hewlett-packard-enterprise",
				Name:              "Hewlett Packard Enterprise",
				Followers:         "4M",
				Employees:         "10K+",
				AssociatedMembers: "74,605",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:     "HewlettPackard",
				Followers: "352",
				Verified:  false,
			},
			BlindProfile: domain.BlindProfile{
				Alias: "",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias: "",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Hewlett-Packard-Enterprise-HPE-EI_IE1093046.11,41.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Hewlett-Packard-Enterprise-HPE-Reviews-E1093046.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Hewlett-Packard-Enterprise-HPE-Jobs-E1093046.htm",
				Jobs:        "297",
				Reviews:     "20K",
				Salaries:    "32K",
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
							Title:                "Cloud Developer (Golang)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4184377636/",
							Date:                 mustDate("2025-03-14"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Multilingual Expert Software Engineer (C, Go, Rust, Python)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4249256830/",
							Date:                 mustDate("2025-06-13"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Golang Cloud developer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4258028903/",
							Date:                 mustDate("2025-07-02"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Cloud Developer (Golang)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4284414526/",
							Location:             "Bengaluru, Karnataka, India",
							Date:                 mustDate("2025-08-16"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Golang Developer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4308037975/",
							Location:             "Bengaluru, Karnataka, India",
							Date:                 mustDate("2025-10-01"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Golang Developer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4338743395/",
							Location:             "Bengaluru, Karnataka, India",
							Date:                 mustDate("2025-11-25"),
							WithSalary:           false,
							Remote:               false,
						},
					},
				},
				domain.Rust: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Multilingual Expert Software Engineer (C, Go, Rust, Python)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4249256830/",
							Date:                 mustDate("2025-06-13"),
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
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "IONOS",
			Website: "https://www.ionos.com/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                10298,
				IDs:               nil,
				Alias:             "ionos",
				Name:              "IONOS",
				Followers:         "63K",
				Employees:         "1K-5K",
				AssociatedMembers: "2,599",
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
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-IONOS-EI_IE688822.11,16.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/IONOS-Reviews-E688822.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/IONOS-Jobs-E688822.htm",
				Jobs:        "10",
				Reviews:     "530",
				Salaries:    "955",
				ReviewsRate: "3.7",
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
							Title:                "Go/Golang Developer (DNS)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4182960038/",
							Date:                 mustDate("2025-03-12"),
							WithSalary:           false,
							Remote:               true,
						},
						{
							Title:                "Senior Software Engineer Golang",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4212017777/",
							Date:                 mustDate("2025-04-24"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Software Engineer Golang",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4221103665/",
							Date:                 mustDate("2025-05-05"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Senior Software Engineer Golang",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4298230825/",
							Location:             "Berlin, Germany",
							Date:                 mustDate("2025-09-10"),
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
			Name:    "Relai",
			Website: "https://relai.app/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                19078807,
				IDs:               nil,
				Alias:             "relai-app",
				Name:              "Relai 🇨🇭",
				Followers:         "14K",
				Employees:         "11-50",
				AssociatedMembers: "64",
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
							Title:                "Senior Go Developer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4157760436/",
							Date:                 mustDate("2025-03-14"),
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
			Ignore:           true, // Crypto
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Velmie",
			Website: "https://www.velmie.com/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                18778012,
				IDs:               nil,
				Alias:             "velmie",
				Name:              "Velmie",
				Followers:         "2K",
				Employees:         "51-200",
				AssociatedMembers: "47",
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
							Title:                "Senior Go Developer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4184379493/",
							Date:                 mustDate("2025-03-14"),
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
			Industries: []domain.Industry{
				domain.IndustryFinTech,
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Clinia",
			Website: "https://clinia.com/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                11160528,
				IDs:               nil,
				Alias:             "clinia",
				Name:              "Clinia",
				Followers:         "3K",
				Employees:         "11-50",
				AssociatedMembers: "41",
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
							Title:                "Senior Go Developer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4183894600/",
							Date:                 mustDate("2025-03-14"),
							WithSalary:           false,
							Remote:               true,
						},
						{
							Title:                "Go Developer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4269501812/",
							Location:             "Montreal, QC",
							Date:                 mustDate("2025-07-19"),
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
			Name:    "Keyrock",
			Website: "https://keyrock.com/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                11259345,
				IDs:               nil,
				Alias:             "keyrock",
				Name:              "Keyrock",
				Followers:         "20K",
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
							Title:                "Senior Software Engineer – High Frequency Trading (C++ OR Rust)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4180667729/",
							Date:                 mustDate("2025-03-14"),
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
			Ignore:           true, // Crypto
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Lazer Technologies",
			Website: "https://www.lazertechnologies.com/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                37569333,
				IDs:               nil,
				Alias:             "lazer-technologies",
				Name:              "Lazer Technologies",
				Followers:         "8K",
				Employees:         "51-200",
				AssociatedMembers: "125",
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
							Title:                "Senior Full Stack Rust Engineer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4184889475/",
							Date:                 mustDate("2025-03-15"),
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
			Ignore:           true, // Crypto
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Cloudwick",
			Website: "https://cloudwick.com/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                2043823,
				IDs:               nil,
				Alias:             "cloudwick",
				Name:              "Cloudwick",
				Followers:         "13K",
				Employees:         "201-500",
				AssociatedMembers: "143",
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
				Alias:     "cloudwick",
				Employees: "180",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Cloudwick-EI_IE1181769.11,20.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Cloudwick-Reviews-E1181769.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Cloudwick-Jobs-E1181769.htm",
				Jobs:        "",
				Reviews:     "58",
				Salaries:    "179",
				ReviewsRate: "4.7",
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
							Title:                "Data Engineer (Health Domain)",
							ShortDescription:     "Design, develop, and maintain ETL processes and data pipelines with Scala/PySpark",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4182758670/",
							Date:                 mustDate("2025-03-13"),
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
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "FIS",
			Website: "https://www.fisglobal.com/",
			Careers: "https://www.fisglobal.com/careers",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                3128,
				IDs:               []int{3128, 804638, 76916563},
				Alias:             "fis",
				Name:              "FIS",
				Followers:         "756K",
				Employees:         "10K+",
				AssociatedMembers: "46,542",
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
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-FIS-EI_IE313114.11,14.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/FIS-Reviews-E313114.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/FIS-Jobs-E313114.htm",
				Jobs:        "719",
				Reviews:     "14K",
				Salaries:    "21K",
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
							Title:                "Lead Engineer – Development (BigData, Scala)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4172072369/",
							Date:                 mustDate("2025-03-12"),
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
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Nike",
			Website: "https://nike.com/",
			Careers: "https://careers.nike.com/",
			About:   "https://about.nike.com/",
			Blog:    "https://engineering.nike.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                2029,
				IDs:               []int{2029, 10818},
				Alias:             "nike",
				Name:              "Nike",
				Followers:         "6M",
				Employees:         "10K+",
				AssociatedMembers: "81,633",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:     "nike-inc",
				Followers: "229",
				Verified:  false,
			},
			BlindProfile: domain.BlindProfile{
				Alias:       "Nike",
				Employees:   "10,000+",
				Salary:      "$65K ~ $280K a year",
				Reviews:     "484",
				ReviewsRate: "3.6",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "nike",
				Employees: "99,580",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-NIKE-EI_IE1699.11,15.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/NIKE-Reviews-E1699.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/NIKE-Jobs-E1699.htm",
				Jobs:        "1.3K",
				Reviews:     "16K",
				Salaries:    "26K",
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
							Title:                "Software Engineer II — Platform Engineering",
							ShortDescription:     "Experience developing Kubernetes controllers in Golang",
							SwitchingOpportunity: "",
							URL:                  "https://www.indeed.com/viewjob?jk=6ea60544f1c6f982",
							Date:                 mustDate("2025-03-18"),
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
			ShortDescription: "American athletic footwear and apparel corporation",
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Aviatrix",
			Website: "https://aviatrix.com/",
			Careers: "https://aviatrix.com/careers/",
			About:   "https://aviatrix.com/about-us/",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                6440846,
				IDs:               nil,
				Alias:             "aviatrix-systems",
				Name:              "Aviatrix",
				Followers:         "38K",
				Employees:         "201-500",
				AssociatedMembers: "495",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:     "aviatrixsystems",
				Followers: "56",
				Verified:  false,
			},
			BlindProfile: domain.BlindProfile{
				Alias: "",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "aviatrix",
				Employees: "10",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Aviatrix-Systems-EI_IE1176649.11,27.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Aviatrix-Systems-Reviews-E1176649.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Aviatrix-Systems-Jobs-E1176649.htm",
				Jobs:        "31",
				Reviews:     "132",
				Salaries:    "266",
				ReviewsRate: "3.2",
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
							Title:                "Principal Architect/Technical Lead Manager",
							ShortDescription:     "Build and enhance automation tools and frameworks in Golang and Python to streamline operations",
							SwitchingOpportunity: "",
							URL:                  "https://www.indeed.com/viewjob?jk=7843b69900f3a0ab",
							Date:                 mustDate("2025-03-18"),
							WithSalary:           true, // $244,000 - $265,000 per year
							Remote:               true,
						},
						{
							Title:                "Staff Software Engineer – Golang Developer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4184984550/",
							Date:                 mustDate("2025-03-19"),
							WithSalary:           true, // $220,000-$226,000
							Remote:               false,
						},
						{
							Title:                "Staff Engineer – Golang Developer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4211120440/",
							Date:                 mustDate("2025-05-10", "2025-04-17"),
							WithSalary:           true, // $188,445 - $221,700 + benefits + 401(k) match + equity
							Remote:               true,
						},
						{
							Title:                "Golang Developer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4340395018/",
							Location:             "Warsaw, Mazowieckie, Poland",
							Date:                 mustDate("2025-12-02"),
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
			ShortDescription: "American multi-cloud networking software who provides connect, manage, and secure their connections between cloud providers",
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Firstbase",
			Website: "https://www.firstbase.io/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                12671153,
				IDs:               nil,
				Alias:             "firstbase-io",
				Name:              "Firstbase",
				Followers:         "14K",
				Employees:         "51-200",
				AssociatedMembers: "60",
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
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Firstbase-EI_IE4382832.11,20.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Firstbase-Reviews-E4382832.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Firstbase-Jobs-E4382832.htm",
				Jobs:        "15",
				Reviews:     "50",
				Salaries:    "49",
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
							Title:                "Senior Full-Stack Engineer – Golang",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4183968735/",
							Date:                 mustDate("2025-03-18"),
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
			ShortDescription: "",
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "NatWest Group",
			Website: "https://natwestgroup.com/",
			Careers: "https://jobs.natwestgroup.com/",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                68360623,
				IDs:               []int{4777, 163231, 12643959, 29024025, 68360623},
				Alias:             "natwest-group",
				Name:              "NatWest Group",
				Followers:         "682K",
				Employees:         "10K+",
				AssociatedMembers: "41,426",
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
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-NatWest-Group-EI_IE10222.11,24.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/NatWest-Group-Reviews-E10222.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/NatWest-Group-Jobs-E10222.htm",
				Jobs:        "207",
				Reviews:     "12K",
				Salaries:    "19K",
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
							Title:                "Back-End Engineer – Golang",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4188231036/",
							Date:                 mustDate("2025-03-18"),
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
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Commonwealth Bank",
			Website: "https://www.commbank.com.au/",
			Careers: "https://www.commbank.com.au/about-us/careers.html",
			About:   "https://www.commbank.com.au/about-us.html",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                2848,
				IDs:               []int{2848, 2850, 2851, 270126},
				Alias:             "commonwealthbank",
				Name:              "Commonwealth Bank",
				Followers:         "600K",
				Employees:         "10K+",
				AssociatedMembers: "45,071",
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
				Alias:     "commonwealth-bank-of-australia",
				Employees: "44,000",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Commonwealth-Bank-of-Australia-EI_IE7922.11,41.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Commonwealth-Bank-of-Australia-Reviews-E7922.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Commonwealth-Bank-of-Australia-Jobs-E7922.htm",
				Jobs:        "350",
				Reviews:     "7.8K",
				Salaries:    "14K",
				ReviewsRate: "3.9",
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
							Title:                "Senior Platform Engineer (AWS, Python, Golang)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4183953062/",
							Date:                 mustDate("2025-03-18"),
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
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "IKEA",
			Website: "https://ikea.com/",
			Careers: "",
			About:   "https://about.ikea.com/",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                2743,
				IDs:               []int{2743, 9854},
				Alias:             "ikea",
				Name:              "IKEA",
				Followers:         "3M",
				Employees:         "10K+",
				AssociatedMembers: "88,900",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:     "IKEA",
				Followers: "40",
				Verified:  true,
			},
			BlindProfile: domain.BlindProfile{
				Alias:       "IKEA",
				Employees:   "10,000+",
				Salary:      "$41K ~ $185K a year",
				Reviews:     "45",
				ReviewsRate: "3.8",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "ikea",
				Employees: "87,550",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-IKEA-EI_IE3957.11,15.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/IKEA-Reviews-E3957.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/IKEA-Jobs-E3957.htm",
				Jobs:        "1.4K",
				Reviews:     "18K",
				Salaries:    "25K",
				ReviewsRate: "3.8",
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
							Title:                "Junior Software Engineer — (Java/Golang) & GCP",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4185204926/",
							Date:                 mustDate("2025-03-17"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Software Engineer — (Golang/Java)",
							ShortDescription:     "Order management",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4288161120/",
							Location:             "Bengaluru, Karnataka, India",
							Date:                 mustDate("2025-08-20"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Software Engineer – Golang",
							ShortDescription:     "Fulfillment",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4311357088/",
							Location:             "Bengaluru, Karnataka, India",
							Date:                 mustDate("2025-10-08"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Senior Software Engineer – Golang/Java",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4323113146/",
							Location:             "Bengaluru, Karnataka, India",
							Date:                 mustDate("2025-11-19"),
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
			Name:    "nexos.ai",
			Website: "https://www.nexos.ai/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                105703851,
				IDs:               nil,
				Alias:             "nexos-ai",
				Name:              "nexos.ai",
				Followers:         "4K",
				Employees:         "11-50",
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
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Back-End Engineer | Middle-Senior | Go",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4188140675/",
							Date:                 mustDate("2025-03-18"),
							WithSalary:           false,
							Remote:               true,
						},
						{
							Title:                "Back-End Engineer | Middle-Senior | Go",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4264729713/",
							Date:                 mustDate("2025-07-09"),
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
			Name:    "GCash",
			Website: "https://wearegcash.com/",
			Careers: "https://wearegcash.com/careers",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                14693558,
				IDs:               nil,
				Alias:             "wearegcash",
				Name:              "GCash",
				Followers:         "213K",
				Employees:         "501-1K",
				AssociatedMembers: "3,024",
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
							URL:                  "https://www.linkedin.com/jobs/view/4183372178/",
							Location:             "Metro Manila",
							Date:                 mustDate("2025-07-27", "2025-07-05", "2025-03-18"),
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
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "mailstep",
			Website: "https://www.mailstep.cz/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                10554555,
				IDs:               nil,
				Alias:             "mail-step",
				Name:              "mailstep",
				Followers:         "2K",
				Employees:         "201-500",
				AssociatedMembers: "96",
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
							Title:                "Software Engineer PHP/Go",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4187272758/",
							Date:                 mustDate("2025-03-17"),
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
			Name:    "Lexia Learning",
			Website: "https://www.lexialearning.com/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                47092,
				IDs:               nil,
				Alias:             "lexia-learning-systems",
				Name:              "Lexia Learning",
				Followers:         "73K",
				Employees:         "501-1K",
				AssociatedMembers: "877",
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
							Title:                "Senior Fullstack Software Engineer, Rust Focused",
							ShortDescription:     "Develop modular, intuitive and scalable components, services and APIs using Rust",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4183950416/",
							Date:                 mustDate("2025-03-18"),
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
			Industries: []domain.Industry{
				domain.IndustryEdTech,
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "CertiK",
			Website: "https://www.certik.com/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                11831043,
				IDs:               nil,
				Alias:             "certik",
				Name:              "CertiK",
				Followers:         "23K",
				Employees:         "201-500",
				AssociatedMembers: "184",
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
							Title:                "Blockchain Security Engineer – (Solidity / Rust / Golang – Senior Level)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4161545800/",
							Date:                 mustDate("2025-03-16"),
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
			ShortDescription: "",
			Ignore:           true, // Web3
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "HolidayCheck",
			Website: "https://holidaycheck.com/",
			Careers: "https://careers.holidaycheck.com/",
			About:   "",
			Blog:    "https://techblog.holidaycheck.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                289231,
				IDs:               nil,
				Alias:             "holidaycheck-ag",
				Name:              "HolidayCheck AG",
				Followers:         "8K",
				Employees:         "201-500",
				AssociatedMembers: "309",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:     "holidaycheck",
				Followers: "6",
				Verified:  true,
			},
			BlindProfile: domain.BlindProfile{
				Alias: "",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias: "",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-HolidayCheck-Group-EI_IE931668.11,29.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/HolidayCheck-Group-Reviews-E931668.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/HolidayCheck-Group-Jobs-E931668.htm",
				Jobs:        "27",
				Reviews:     "47",
				Salaries:    "113",
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
					GitHubRepositoriesCount: 3,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Software Engineer – Scala",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4087274629/",
							Location:             "Munich, Bavaria, Germany",
							Date:                 mustDate("2025-07-03", "2025-06-24", "2025-06-18", "2025-06-11", "2025-05-20", "2025-04-27", "2025-03-16"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Software Engineer – Scala",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4271201559/",
							Location:             "Munich, Bavaria, Germany",
							Date:                 mustDate("2025-12-02", "2025-11-12", "2025-09-09", "2025-08-19", "2025-07-28"),
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
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Dyad",
			Website: "https://dyad.net/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                67565681,
				IDs:               nil,
				Alias:             "dyad-ai",
				Name:              "Dyad",
				Followers:         "2K",
				Employees:         "11-50",
				AssociatedMembers: "40",
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
							Title:                "Elixir Product Engineer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4188226908/",
							Date:                 mustDate("2025-03-18"),
							WithSalary:           false,
							Remote:               false,
						},
					},
				},
				domain.Clojure: {},
				domain.Haskell: {},
			},
			ShortDescription: "",
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Stylitics",
			Website: "https://www.stylitics.com/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                1616816,
				IDs:               nil,
				Alias:             "stylitics",
				Name:              "Stylitics",
				Followers:         "14K",
				Employees:         "51-200",
				AssociatedMembers: "177",
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
				Alias:     "stylitics",
				Employees: "126",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Stylitics-EI_IE1873109.11,20.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Stylitics-Reviews-E1873109.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Stylitics-Jobs-E1873109.htm",
				Jobs:        "",
				Reviews:     "71",
				Salaries:    "112",
				ReviewsRate: "2.9",
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
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Clojure Backend Engineer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4181907066/",
							Date:                 mustDate("2025-03-15"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Clojure/ClojureScript Engineer",
							ShortDescription:     "Fashion AI Automation",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4331755523/",
							Location:             "New York, NY",
							Date:                 mustDate("2025-11-02"),
							WithSalary:           true, // $135k/yr - $155k/yr
							Remote:               true,
						},
					},
				},
				domain.Haskell: {},
			},
			ShortDescription: "",
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "CERESTI HEALTH",
			Website: "https://ceresti.com/",
			Careers: "https://www.ceresti.com/careers",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                15403042,
				IDs:               nil,
				Alias:             "ceresti-health-inc.",
				Name:              "CERESTI HEALTH",
				Followers:         "1K",
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
				Alias: "Ceresti-Health",
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Senior Golang Backend Developer",
							ShortDescription:     "3+ years of experience with the Go",
							SwitchingOpportunity: "",
							URL:                  "https://www.indeed.com/viewjob?jk=ed0f373654d5bce5",
							Date:                 mustDate("2025-03-19"),
							WithSalary:           true, // From $135,000 per year
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
			ShortDescription: "A medical facility that helps clients with Alzheimer's",
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Amuse",
			Website: "https://www.amuse.io/",
			Careers: "https://careers.amuse.io/jobs",
			About:   "https://www.amuse.io/en/resources/about/",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                10607933,
				IDs:               nil,
				Alias:             "amuse.io",
				Name:              "Amuse",
				Followers:         "9K",
				Employees:         "51-200",
				AssociatedMembers: "279",
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
				Alias: "Amuse",
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Full Stack Developer",
							ShortDescription:     "You have experience using Go or a strongly typed OO language",
							SwitchingOpportunity: "",
							URL:                  "https://www.indeed.com/viewjob?jk=1aeebd1be05a4f6e",
							Date:                 mustDate("2025-03-19"),
							WithSalary:           true, // $90,000 - $115,000 per year
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
			ShortDescription: "Global distributed  music company",
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Lucid Motors",
			Website: "https://lucidmotors.com/",
			Careers: "https://lucidmotors.com/careers",
			About:   "https://lucidmotors.com/company",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                1684122,
				IDs:               nil,
				Alias:             "lucidmotors",
				Name:              "Lucid Motors",
				Followers:         "512K",
				Employees:         "1K-5K",
				AssociatedMembers: "6,366",
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
				Alias:     "lucid-motors",
				Employees: "2,250",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Lucid-Motors-EI_IE698305.11,23.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Lucid-Motors-Reviews-E698305.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Lucid-Motors-Jobs-E698305.htm",
				Jobs:        "434",
				Reviews:     "1.1K",
				Salaries:    "3.3K",
				ReviewsRate: "3.0",
				Verified:    true,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "Lucid-Motors",
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Senior Full Stack Developer",
							ShortDescription:     "Design and develop full-stack web applications, microservices, and REST APIs using languages such as Golang, Python, Java, or similar",
							SwitchingOpportunity: "",
							URL:                  "https://www.indeed.com/viewjob?jk=e35725dabddd890c",
							Date:                 mustDate("2025-03-19"),
							WithSalary:           true, // $154,000 - $211,750 per year
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
			ShortDescription: "American automotive and technology company that produces electric vehicles",
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "The ODP Corporation",
			Website: "https://www.theodpcorp.com/",
			Careers: "https://careers.theodpcorp.com/",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                75753067,
				IDs:               nil,
				Alias:             "the-odp-corporation",
				Name:              "The ODP Corporation",
				Followers:         "11K",
				Employees:         "10K+",
				AssociatedMembers: "27,156",
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
				Alias:     "the-odp-corporation",
				Employees: "120",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-The-ODP-Corporation-EI_IE9487978.11,30.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/The-ODP-Corporation-Reviews-E9487978.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/The-ODP-Corporation-Jobs-E9487978.htm",
				Jobs:        "7",
				Reviews:     "6",
				Salaries:    "17",
				ReviewsRate: "2.8",
				Verified:    true,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "The-Odp-Corporation",
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Senior Backend Developer",
							ShortDescription:     "Architect and implement solutions primarily using Golang",
							SwitchingOpportunity: "",
							URL:                  "https://www.indeed.com/viewjob?jk=164ee8448e688865",
							Date:                 mustDate("2025-03-19"),
							WithSalary:           true, // $75,500 - $117,950 per year
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
			ShortDescription: "American company who provides business services and supplies, products, and digital workplace technology solutions for small, medium, and enterprise businesses",
		},

		// BigTech
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Netflix",
			Website: "https://netflix.com/",
			Careers: "https://jobs.netflix.com/",
			About:   "https://about.netflix.com/",
			Blog:    "https://netflixtechblog.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                165158,
				IDs:               nil,
				Alias:             "netflix",
				Name:              "Netflix",
				Followers:         "11M",
				Employees:         "10K+",
				AssociatedMembers: "16,766",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:     "netflix",
				Followers: "7.9k",
				Verified:  false,
			},
			BlindProfile: domain.BlindProfile{
				Alias:       "Netflix",
				Employees:   "1,001 to 5,000",
				Salary:      "$87K ~ $827K a year",
				Reviews:     "576",
				ReviewsRate: "4.2",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "netflix",
				Employees: "14,600",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Netflix-EI_IE11891.11,18.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Netflix-Reviews-E11891.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Netflix-Jobs-E11891.htm",
				Jobs:        "466",
				Reviews:     "3.3K",
				Salaries:    "6.8K",
				ReviewsRate: "4.2",
				Verified:    true,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "Netflix",
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Software Engineer L5 — Linux Kernel Developer",
							ShortDescription:     "Proficiency in Golang, Python",
							SwitchingOpportunity: "",
							URL:                  "https://www.indeed.com/viewjob?jk=5a23d65371dc5435",
							Date:                 mustDate("2025-03-19"),
							WithSalary:           true, // $100,000 - $620,000 per year
							Remote:               true,
						},
						{
							Title:                "Software Engineer 6 — Games Experience Engineering",
							ShortDescription:     "Languages: JavaScript, Java, Python, Lua, Golang, Clojure",
							SwitchingOpportunity: "",
							URL:                  "https://www.indeed.com/viewjob?jk=ba7e24381e243687",
							Date:                 mustDate("2025-03-23"),
							WithSalary:           true, // $230,000 - $960,000 per year
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
			ShortDescription: "American online streaming service",
		},

		// BigTech
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Panasonic",
			Website: "https://holdings.panasonic/",
			Careers: "https://holdings.panasonic/global/corporate/careers.html",
			About:   "https://holdings.panasonic/global/corporate/about.html",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                3549587,
				IDs:               nil,
				Alias:             "panasonic",
				Name:              "Panasonic",
				Followers:         "302K",
				Employees:         "10K+",
				AssociatedMembers: "28,942",
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
				Alias:     "panasonic",
				Employees: "259,000",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Panasonic-EI_IE4279.11,20.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Panasonic-Reviews-E4279.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Panasonic-Jobs-E4279.htm",
				Jobs:        "1.2K",
				Reviews:     "5.1K",
				Salaries:    "8.8K",
				ReviewsRate: "3.7",
				Verified:    false,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "Panasonic-a3e6d439",
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Golang Network Developer",
							ShortDescription:     "Strong proven experience in Golang",
							SwitchingOpportunity: "",
							URL:                  "https://www.indeed.com/viewjob?jk=858278a26965db1b",
							Date:                 mustDate("2025-03-19"),
							WithSalary:           true, // $93,000 - $157,000 per year
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
			ShortDescription: "Japanese multinational company who manufactures batteries, automotive and aviation systems, industrial systems, and home repair and construction",
		},
	}
}
