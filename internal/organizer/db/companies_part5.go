package db

import (
	"github.com/readytotouch/readytotouch/internal/domain"
)

func companiesPart5() []domain.CompanyProfile {
	return []domain.CompanyProfile{

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Atomic",
			Website: "https://www.atomic.com/",
			Careers: "https://www.atomic.com/en/karriere",
			About:   "https://www.atomic.com/en/brand",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                6542309,
				Alias:             "atomicskiing",
				Name:              "Atomic",
				Followers:         "18K",
				Employees:         "501-1K", // 501-1K
				AssociatedMembers: "346",
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
							Title:            "Backend Golang Software Engineer",
							ShortDescription: "Build and scale backend systems for a FinTech SaaS platform, focusing on low-latency, high-reliability, and scalability",
							URL:              "https://www.linkedin.com/jobs/view/4106745259/",
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
			ShortDescription: "EU sports equipment",
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Outpost24",
			Website: "https://outpost24.com/",
			Careers: "https://careers.outpost24.com/",
			About:   "https://outpost24.com/company/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                37910,
				Alias:             "outpost24",
				Name:              "Outpost24",
				Followers:         "8K",
				Employees:         "201-500", // 201-500
				AssociatedMembers: "245",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "outpost24",
				Verified: false,
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Senior Golang Developer",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4106466480/",
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
			ShortDescription: "Cyber risk management",
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Lumen Technologies",
			Website: "https://www.lumen.com",
			Careers: "https://jobs.lumen.com/global/en",
			About:   "https://www.lumen.com/en-us/about.html",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                47664328,
				Alias:             "lumentechnologies",
				Name:              "Lumen Technologies",
				Followers:         "540K",
				Employees:         "10K+",
				AssociatedMembers: "49,722",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "CenturyLink",
				Verified: false,
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Principal Golang Developer",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4106292200/",
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
			ShortDescription: "American telecommunications company",
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "SAS",
			Website: "https://www.sas.com/",
			Careers: "https://www.sas.com/en_us/careers.html",
			About:   "https://www.sas.com/en_us/company-information.html",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                1491,
				Alias:             "sas",
				Name:              "SAS",
				Followers:         "916K",
				Employees:         "10K+",
				AssociatedMembers: "17,267",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "sassoftware",
				Verified: true,
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Software Developer (Golang / Kubernetes)",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4105023413/",
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
			ShortDescription: "American multinational developer of analytics and artificial intelligence software",
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Miquido",
			Website: "https://www.miquido.com",
			Careers: "https://careers.miquido.com/",
			About:   "https://www.miquido.com/about-us/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                2307119,
				Alias:             "miquido",
				Name:              "Miquido",
				Followers:         "7K",
				Employees:         "201-500",
				AssociatedMembers: "196",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "miquido",
				Verified: true,
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Senior Golang Developer",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4104781741/",
							Date:             mustDate("2024-12-19"),
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
			ShortDescription: "Software development company",
			Ignore:           true, // Miquido is an outsourcing company, according to the description
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "CellphoneS",
			Website: "https://cellphones.com.vn/",
			Careers: "",
			About:   "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                14590467,
				Alias:             "cellphones",
				Name:              "CellphoneS",
				Followers:         "2K",
				Employees:         "501-1K",
				AssociatedMembers: "327",
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
							Title:            "Senior Golang Developer",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4097283037/",
							Date:             mustDate("2024-12-19"),
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
			ShortDescription: "Mobile phone retail system",
			Ignore:           true, // For now, we are adding companies that have an English vacancy description and an English company description, so CellphoneS has been removed.
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Koch",
			Website: "https://www.kochinc.com/",
			Careers: "https://jobs.kochcareers.com/",
			About:   "https://www.kochinc.com/about",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                164621,
				Alias:             "kochinc",
				Name:              "Koch",
				Followers:         "254K",
				Employees:         "10K+",
				AssociatedMembers: "5,931",
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
							Title:            "Golang Developer",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4103430030/",
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
			ShortDescription: "American multinational conglomerate corporation",
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "TradingView",
			Website: "https://www.tradingview.com/",
			Careers: "https://www.tradingview.com/careers/",
			About:   "https://www.tradingview.com/about/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                3179069,
				Alias:             "tradingview",
				Name:              "TradingView",
				Followers:         "78K",
				Employees:         "201-500",
				AssociatedMembers: "1,562",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "tradingview",
				Verified: true,
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Senior Golang Developer — Alerts",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4102796032/",
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
			ShortDescription: "Charting platform and social network",
			Ignore:           true, // Currently ignored
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Groq",
			Website: "https://groq.com/",
			Careers: "https://groq.com/careers/",
			About:   "https://groq.com/about-us/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                12994410,
				Alias:             "groq",
				Name:              "Groq",
				Followers:         "119K",
				Employees:         "51-200",
				AssociatedMembers: "340",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "groq",
				Verified: true,
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Staff Software Engineer, Golang",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4108714658/",
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
			ShortDescription: "American artificial intelligence (AI) company",
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "PayRetailers",
			Website: "https://www.payretailers.com/",
			Careers: "https://www.payretailers.com/en/work-with-us/",
			About:   "https://www.payretailers.com/en/about-us/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                11060015,
				Alias:             "pay-retailers",
				Name:              "PayRetailers",
				Followers:         "34K",
				Employees:         "201-500",
				AssociatedMembers: "370",
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
							Title:            "Senior Software Engineer, Golang",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4105750575/",
							Date:             mustDate("2024-12-20"),
						},
						{
							Title:                "Senior Software Engineer, Golang",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4114715170/",
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
			ShortDescription: "Payment service provider",
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Zendesk",
			Website: "https://www.zendesk.com/",
			Careers: "https://jobs.zendesk.com/us/en",
			About:   "https://www.zendesk.com/about/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                418095,
				Alias:             "zendesk",
				Name:              "Zendesk",
				Followers:         "560K",
				Employees:         "5K-10K",
				AssociatedMembers: "6,655",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "zendesk",
				Verified: true,
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Software Engineer II, Golang",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4105423482/",
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
			ShortDescription: "Provides software-as-a-service products related to customer support",
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Mobica",
			Website: "https://mobica.com/",
			Careers: "https://mobica.com/careers",
			About:   "https://mobica.com/about_us",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                29470,
				Alias:             "mobica",
				Name:              "Mobica",
				Followers:         "57K",
				Employees:         "501-1K",
				AssociatedMembers: "728",
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
							Title:            "Senior Software Engineer — Golang",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4105552556/",
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
			ShortDescription: "Provides integrated interior and exterior solutions",
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "EDB",
			Website: "https://www.enterprisedb.com/",
			Careers: "https://www.enterprisedb.com/careers",
			About:   "https://www.enterprisedb.com/company",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                14958,
				Alias:             "edbpostgres",
				Name:              "EDB",
				Followers:         "47K",
				Employees:         "501-1K",
				AssociatedMembers: "1,335",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "enterprisedb",
				Verified: false,
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Platform Engineer (Golang)",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4083004984/",
							Date:             mustDate("2024-12-19"),
						},
						{
							Title:                "Platform Engineer (Golang)",
							ShortDescription:     "3-5 years of experience with Golang, Python, Kubernetes, GitHub Actions, OpenShift",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4115835388/",
							Date:                 mustDate("2025-01-03"),
							WithSalary:           true, // $125-150k
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
			ShortDescription: "Open source, pre-configured, certified binary PostgreSQL distribution that simplifies enterprise deployment",
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Altair",
			Website: "https://altair.com/",
			Careers: "https://altair.com/careers",
			About:   "https://altair.com/company",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                8323,
				Alias:             "altair-engineering",
				Name:              "Altair",
				Followers:         "207K",
				Employees:         "1K-5K",
				AssociatedMembers: "4,098",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "altairengineering",
				Verified: false,
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Senior Software Engineer — Golang",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4103107908/",
							Date:             mustDate("2024-12-19"),
						},
						{
							Title:                "Senior Software Engineer — Golang",
							ShortDescription:     "A minimum of 4 years of experience in Golang or other high-level language",
							SwitchingOpportunity: "Some of this experience should be in Golang; however, an interest in Golang and experience with other programming languages such as C++, Python, and Java are suitable",
							URL:                  "https://www.linkedin.com/jobs/view/4116215019/",
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
			ShortDescription: "American multinational information technology company",
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Reward Gateway",
			Website: "https://www.rewardgateway.com/",
			Careers: "https://careers.rewardgateway.com/",
			About:   "https://www.rewardgateway.com/company/overview",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                200327,
				Alias:             "reward-gateway",
				Name:              "Reward Gateway",
				Followers:         "34K",
				Employees:         "501-1K",
				AssociatedMembers: "798",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "rewardgateway",
				Verified: false,
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Backend Engineer (Golang)",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4103640949/",
							Date:             mustDate("2024-12-19"),
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
			ShortDescription: "Online platform that allows employers to provide rewards and recognition",
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "NN Group",
			Website: "https://www.nn-group.com",
			Careers: "https://www.nn-group.com/careers/engaging-employees.htm",
			About:   "https://www.nn-group.com/our-company/who-we-are.htm",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                23110483,
				Alias:             "nn-group",
				Name:              "NN Group",
				Followers:         "31K",
				Employees:         "10K+",
				AssociatedMembers: "21,871",
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
							Title:            "Senior Rust Developer",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4106226520/",
							Date:             mustDate("2024-12-21"),
						},
					},
				},
				domain.Zig:     {},
				domain.Scala:   {},
				domain.Elixir:  {},
				domain.Clojure: {},
				domain.Haskell: {},
			},
			ShortDescription: "International financial services company",
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "American Airlines",
			Website: "https://www.americanairlines.com/",
			Careers: "https://jobs.aa.com/",
			About:   "https://www.aa.com/i18n/customer-service/about-us/history-of-american-airlines.jsp",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                2640,
				Alias:             "american-airlines",
				Name:              "American Airlines",
				Followers:         "1M",
				Employees:         "10K+",
				AssociatedMembers: "53,291",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "americanairlines",
				Verified: false,
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Engineer/Sr Engineer, Kubernetes GO — Developer",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4092856552/",
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
			ShortDescription: "Major airline in the United States",
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Cake by VPBank — Digital Bank",
			Website: "https://cake.vn/",
			Careers: "https://cake.vn/recruitment",
			About:   "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                76580235,
				Alias:             "cake-by-vpbank",
				Name:              "Cake by VPBank — Digital Bank",
				Followers:         "12K",
				Employees:         "501-1K",
				AssociatedMembers: "175",
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
							Title:            "Senior/Staff/Principal Software Engineer (GO, GCP)",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4108779705/",
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
			ShortDescription: "Digital bank in Vietnam",
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Cheesecake Labs",
			Website: "https://cheesecakelabs.com/",
			Careers: "https://cheesecakelabs.com/careers/",
			About:   "https://cheesecakelabs.com/about/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                5251039,
				Alias:             "cheesecake-labs",
				Name:              "Cheesecake Labs",
				Followers:         "16K",
				Employees:         "51-200",
				AssociatedMembers: "139",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "CheesecakeLabs",
				Verified: true,
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Software Engineer Fullstack — GO",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4109304866/",
							Date:                 mustDate("2024-12-24"),
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
			Name:    "Maya",
			Website: "https://www.maya.ph/",
			Careers: "https://www.maya.ph/careers",
			About:   "https://www.maya.ph/about-us",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                80861757,
				Alias:             "mayaph",
				Name:              "Maya",
				Followers:         "205K",
				Employees:         "501-1K",
				AssociatedMembers: "1,556",
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
							Title:            "Software Engineer (Java/GO)",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4103938575/",
							Date:             mustDate("2025-01-08"), // mustDate("2024-12-19"),
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
			ShortDescription: "Digital Bank in the Philippines",
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Ashley Furniture Industries",
			Website: "https://www.ashleyfurnitureindustriesllc.com",
			Careers: "https://careers.ashleyfurniture.com",
			About:   "https://www.ashleyfurnitureindustriesllc.com/en/company/company-overview",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                162547,
				Alias:             "ashley-furniture-industries",
				Name:              "Ashley Furniture Industries",
				Followers:         "102K",
				Employees:         "10K+",
				AssociatedMembers: "9,746",
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
							Title:            "Sr Rust Developer",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4080571465/",
							Date:             mustDate("2024-12-18"),
						},
					},
				},
				domain.Zig:     {},
				domain.Scala:   {},
				domain.Elixir:  {},
				domain.Clojure: {},
				domain.Haskell: {},
			},
			ShortDescription: "American home furnishings manufacturer and retailer",
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
		//},
	}
}
