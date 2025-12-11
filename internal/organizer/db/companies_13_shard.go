package db

import "github.com/readytotouch/readytotouch/internal/domain"

func companies13Shard() []domain.CompanyProfile {
	return []domain.CompanyProfile{
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "TRON DAO",
			Website: "https://tron.network/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                18289948,
				IDs:               []int{29722, 18289948},
				Alias:             "trondao",
				Name:              "TRON DAO",
				Followers:         "54K",
				Employees:         "201-500",
				AssociatedMembers: "375",
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
							Title:                "Blockchain Systems Engineer Golang",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4112783642/",
							Date:                 mustDate("2025-03-19"),
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
			Ignore:           true, // Blockchain | Crypto | Web3
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Gnosis",
			Website: "https://gnosis.io/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                18098361,
				IDs:               nil,
				Alias:             "gnosis-limited",
				Name:              "Gnosis",
				Followers:         "8K",
				Employees:         "51-200",
				AssociatedMembers: "162",
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
							Title:                "Senior Golang Engineer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4186509026/",
							Date:                 mustDate("2025-03-19"),
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
			Ignore:           true, // Blockchain
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Netscribes",
			Website: "https://www.netscribes.com/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                25974,
				IDs:               nil,
				Alias:             "netscribes",
				Name:              "Netscribes",
				Followers:         "103K",
				Employees:         "1K-5K",
				AssociatedMembers: "1,401",
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
				Alias:     "netscribes",
				Employees: "2,000",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Netscribes-EI_IE218712.11,21.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Netscribes-Reviews-E218712.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Netscribes-Jobs-E218712.htm",
				Jobs:        "14",
				Reviews:     "861",
				Salaries:    "787",
				ReviewsRate: "3.6",
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
							Title:                "Senior Software Engineer – Golang",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4184795681/",
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
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Warner Bros. Discovery",
			Website: "https://wbd.com/",
			Careers: "https://careers.wbd.com/",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                73234923,
				IDs:               []int{2046, 2470, 4749, 9620, 167119, 200300, 1393227, 10538105, 26998697, 73234923},
				Alias:             "warner-bros-discovery",
				Name:              "Warner Bros. Discovery",
				Followers:         "811K",
				Employees:         "10K+",
				AssociatedMembers: "39,992",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "",
				Verified: false,
			},
			BlindProfile: domain.BlindProfile{
				Alias:       "Warner-Bros-Discovery",
				Employees:   "10,000+",
				Salary:      "",
				Reviews:     "61",
				ReviewsRate: "2.9",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "warner-bros-discovery",
				Employees: "24,250",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Warner-Bros-Discovery-EI_IE12777.11,32.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Warner-Bros-Discovery-Reviews-E12777.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Warner-Bros-Discovery-Jobs-E12777.htm",
				Jobs:        "3.4K",
				Reviews:     "4.8K",
				Salaries:    "9.3K",
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
							Title:                "Staff Software Engineer – Golang",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4187284222/",
							Location:             "Bengaluru, Karnataka, India",
							Date:                 mustDate("2025-07-25", "2025-07-03", "2025-06-12", "2025-05-22", "2025-03-17"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Senior Software Engineer – Golang",
							ShortDescription:     "MSC Team",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4269395988/",
							Location:             "Bengaluru, Karnataka, India",
							Date:                 mustDate("2025-09-20", "2025-08-09"),
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
			Name:    "Gamingtec",
			Website: "https://gamingtec.com/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                9276193,
				IDs:               nil,
				Alias:             "gamingtec",
				Name:              "Gamingtec",
				Followers:         "6K",
				Employees:         "51-200",
				AssociatedMembers: "145",
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
							URL:                  "https://www.linkedin.com/jobs/view/4187105723/",
							Date:                 mustDate("2025-03-19"),
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
			Ignore:           true, // Gambling
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "VOIS",
			Website: "https://www.vodafone.com/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                167822,
				IDs:               nil,
				Alias:             "vois",
				Name:              "VOIS",
				Followers:         "345K",
				Employees:         "10K+",
				AssociatedMembers: "18,326",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:     "vodafone",
				Followers: "75",
				Verified:  false,
			},
			BlindProfile: domain.BlindProfile{
				Alias:       "Vodafone",
				Employees:   "10,000+",
				Salary:      "$20K ~ $450K a year",
				Reviews:     "56",
				ReviewsRate: "3.8",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "vodafone",
				Employees: "165,610",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Vodafone-EI_IE5775.11,19.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Vodafone-Reviews-E5775.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Vodafone-Jobs-E5775.htm",
				Jobs:        "2.4K",
				Reviews:     "19K",
				Salaries:    "27K",
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
							Title:                "Manager Golang Developer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4180356164/",
							Date:                 mustDate("2025-03-19"),
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
			Name:    "Spydra",
			Website: "https://www.spydra.app/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                84968130,
				IDs:               nil,
				Alias:             "spydra",
				Name:              "Spydra",
				Followers:         "50K",
				Employees:         "11-50",
				AssociatedMembers: "36",
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
							Title:                "Golang Developer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4184559478/",
							Date:                 mustDate("2025-03-19"),
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
			Ignore:           true, // Blockchain
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Veact GmbH",
			Website: "https://veact.net/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                1412127,
				IDs:               nil,
				Alias:             "veact",
				Name:              "Veact GmbH",
				Followers:         "2K",
				Employees:         "51-200",
				AssociatedMembers: "111",
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
							Title:                "Scala Developer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4182628362/",
							Date:                 mustDate("2025-03-15"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Scala Developer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4221250883/",
							Date:                 mustDate("2025-05-02"),
							WithSalary:           false,
							Remote:               false,
						},
					},
				},
				domain.Elixir:  {},
				domain.Clojure: {},
				domain.Haskell: {},
			},
			ShortDescription: "Veact is developing applications, processes and products to improve sales success in the car trade market using data-based marketing",
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Walmart Data Ventures",
			Website: "https://www.walmartdataventures.com/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                80427781,
				IDs:               nil,
				Alias:             "walmartdataventures",
				Name:              "Walmart Data Ventures",
				Followers:         "15K",
				Employees:         "201-500",
				AssociatedMembers: "244",
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
							Title:                "Senior, Data Engineer (Scala/Spark)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4189070192/",
							Date:                 mustDate("2025-03-19"),
							WithSalary:           true,
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
			Name:    "Funnel",
			Website: "https://funnel.io/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                9397630,
				IDs:               nil,
				Alias:             "funnel-io",
				Name:              "Funnel",
				Followers:         "54K",
				Employees:         "201-500",
				AssociatedMembers: "437",
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
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Funnel-EI_IE1559225.11,17.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Funnel-Reviews-E1559225.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Funnel-Jobs-E1559225.htm",
				Jobs:        "11",
				Reviews:     "96",
				Salaries:    "171",
				ReviewsRate: "3.9",
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
							Title:                "Senior Data Platform Engineer – Rust",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4142024603/",
							Date:                 mustDate("2025-06-14", "2025-05-03", "2025-04-11", "2025-03-20"),
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
			ShortDescription: "Marketing intelligence platform",
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "OKX",
			Website: "https://www.okx.com/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                13323120,
				IDs:               nil,
				Alias:             "okxofficial",
				Name:              "OKX",
				Followers:         "462K",
				Employees:         "1K-5K",
				AssociatedMembers: "5,933",
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
							Title:                "Blockchain Engineer (Rust/Go)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4187130229/",
							Date:                 mustDate("2025-03-20"),
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
			Ignore:           true, // Crypto
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Lingopal",
			Website: "https://lingopal.com/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                98079583,
				IDs:               nil,
				Alias:             "lingopal-ai",
				Name:              "Lingopal",
				Followers:         "50K",
				Employees:         "51-200",
				AssociatedMembers: "31",
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
							URL:                  "https://www.linkedin.com/jobs/view/4188980146/",
							Date:                 mustDate("2025-03-20"),
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
			ShortDescription: "Translate LIVE speech into any language in real time",
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Scanbo",
			Website: "https://www.scanbo.com/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                13271827,
				IDs:               nil,
				Alias:             "scanbo",
				Name:              "Scanbo",
				Followers:         "2K",
				Employees:         "2-10",
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
							Title:                "Rust Blockchain Developer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4185486374/",
							Date:                 mustDate("2025-03-22"),
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
			ShortDescription: "SCANBO is a HealthCare platform dedicated to making medical diagnosis available at the point of care and at home",
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "SurrealDB",
			Website: "https://surrealdb.com/",
			Careers: "https://surrealdb.com/careers",
			About:   "https://medium.com/surrealdb/about",
			Blog:    "https://surrealdb.com/blog",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                71534355,
				IDs:               nil,
				Alias:             "surrealdb",
				Name:              "SurrealDB",
				Followers:         "7K",
				Employees:         "11-50",
				AssociatedMembers: "43",
				Verified:          false,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "surrealdb",
				Verified: true,
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
			OttaProfileSlug: "SurrealDB",
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 11,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Rust: {
					GitHubRepositoriesCount: 25,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Senior Software Engineer",
							ShortDescription:     "Proficient in programming languages including Rust, C, or Go",
							SwitchingOpportunity: "",
							URL:                  "https://rustjobs.dev/featured-jobs/SurrealDB-Senior-Software-Engineer-Query-Language-QL-TcCVMEInkwcjBpupa3Ys",
							Date:                 mustDate("2024-02-03"),
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
			ShortDescription: "Serverless cloud database",
			Industries: []domain.Industry{
				domain.IndustryDevOps,
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Clockwork Labs",
			Website: "https://clockworklabs.io/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                19088131,
				IDs:               nil,
				Alias:             "clockwork-labs",
				Name:              "Clockwork Labs",
				Followers:         "3K",
				Employees:         "2-10",
				AssociatedMembers: "31",
				Verified:          false,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "clockworklabs",
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
					GitHubRepositoriesCount: 9,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Zig:     {},
				domain.Scala:   {},
				domain.Elixir:  {},
				domain.Clojure: {},
				domain.Haskell: {},
			},
			ShortDescription: "Computer Games",
		},
		{
			ID:      0,        // system
			Type:    "",       // system
			Name:    "Polars", // Polars BV
			Website: "https://pola.rs/",
			Careers: "",
			About:   "https://pola.rs/about-us/",
			Blog:    "https://pola.rs/posts/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                98439348,
				IDs:               nil,
				Alias:             "pola-rs",
				Name:              "Polars",
				Followers:         "20K",
				Employees:         "2-10",
				AssociatedMembers: "14",
				Verified:          false,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "pola-rs",
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
					GitHubRepositoriesCount: 8,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Rust Backend Engineer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://hiring.pola.rs/o/rust-backend-engineer",
							Date:                 mustDate("2025-03-21"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Database Engineer",
							ShortDescription:     "Experience in Rust",
							SwitchingOpportunity: "",
							URL:                  "https://hiring.pola.rs/o/database-engineer",
							Date:                 mustDate("2025-03-21"),
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
			ShortDescription: "Lightning-fast DataFrame library for Rust and Python",
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Tsunami Tsolutions",
			Website: "https://tsunamitsolutions.com/",
			Careers: "https://tsunamitsolutions.com/index.php/careers/",
			About:   "https://tsunamitsolutions.com/index.php/who-we-are/",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                983003,
				IDs:               nil,
				Alias:             "tsunami-tsolutions-llc.",
				Name:              "Tsunami Tsolutions",
				Followers:         "3K",
				Employees:         "51-200",
				AssociatedMembers: "127",
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
					Vacancies:               []domain.Vacancy{
						// Unclear if this is a Go position https://www.indeed.com/viewjob?jk=e5b01bd47b175b64
					},
				},
				domain.Rust:    {},
				domain.Zig:     {},
				domain.Scala:   {},
				domain.Elixir:  {},
				domain.Clojure: {},
				domain.Haskell: {},
			},
			ShortDescription: "American aviation & aerospace company that offers Engineering Support and Information Solutions",
			Ignore:           true, // Unclear about the tech stack
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Draper",
			Website: "https://www.draper.com/",
			Careers: "https://www.draper.com/careers",
			About:   "https://www.draper.com/about",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                163500,
				IDs:               nil,
				Alias:             "draper",
				Name:              "Draper",
				Followers:         "46K",
				Employees:         "1K-5K",
				AssociatedMembers: "2,475",
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
				Alias:     "draper",
				Employees: "1,750",
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
					Vacancies:               []domain.Vacancy{
						// Unclear if this is a Go position https://www.indeed.com/viewjob?jk=f77b58384382b852
					},
				},
				domain.Rust:    {},
				domain.Zig:     {},
				domain.Scala:   {},
				domain.Elixir:  {},
				domain.Clojure: {},
				domain.Haskell: {},
			},
			ShortDescription: "American nonprofit engineering innovation company",
			Ignore:           true, // Unclear about the tech stack
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Mayo Clinic",
			Website: "https://www.mayoclinic.org/",
			Careers: "https://jobs.mayoclinic.org/",
			About:   "https://www.mayoclinic.org/about-mayo-clinic",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                4725,
				IDs:               nil,
				Alias:             "mayo-clinic",
				Name:              "Mayo Clinic",
				Followers:         "1M",
				Employees:         "10K+",
				AssociatedMembers: "48,372",
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
				Alias:     "mayo-clinic",
				Employees: "63,000",
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
					Vacancies:               []domain.Vacancy{
						// Unclear if this is a Go position https://www.indeed.com/viewjob?jk=4c3c7f291e1c965f
					},
				},
				domain.Rust:    {},
				domain.Zig:     {},
				domain.Scala:   {},
				domain.Elixir:  {},
				domain.Clojure: {},
				domain.Haskell: {},
			},
			ShortDescription: "American nonprofit academic medical center focused on integrated health care, education, and research",
			Ignore:           true, // Unclear about the tech stack
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "University of Florida",
			Website: "https://ufl.edu/",
			Careers: "https://jobs.ufl.edu/",
			About:   "https://www.ufl.edu/about/",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                4282,
				IDs:               []int{4282, 7974417, 15101148, 15102044, 18869387},
				Alias:             "uflorida",
				Name:              "University of Florida",
				Followers:         "606K",
				Employees:         "",
				AssociatedMembers: "",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "smartsystemslab-uf",
				Verified: false,
			},
			BlindProfile: domain.BlindProfile{
				Alias: "",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "university-of-florida",
				Employees: "19,453",
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
					Vacancies:               []domain.Vacancy{
						// Unclear if this is a Go position https://www.indeed.com/viewjob?jk=630a513e4b7e6d8e
					},
				},
				domain.Rust:    {},
				domain.Zig:     {},
				domain.Scala:   {},
				domain.Elixir:  {},
				domain.Clojure: {},
				domain.Haskell: {},
			},
			ShortDescription: "American public land-grant research university",
			Ignore:           true, // Unclear about the tech stack
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Storage Rentals of America",
			Website: "https://www.sroa.com/",
			Careers: "https://www.sroa.com/careers",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                33228180,
				IDs:               nil,
				Alias:             "storage-rentals-america",
				Name:              "Storage Rentals of America",
				Followers:         "5K",
				Employees:         "501-1K",
				AssociatedMembers: "200",
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
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Storage-Rentals-of-America-EI_IE1406238.11,37.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Storage-Rentals-of-America-Reviews-E1406238.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Storage-Rentals-of-America-Jobs-E1406238.htm",
				Jobs:        "26",
				Reviews:     "83",
				Salaries:    "88",
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
							Title:                "DevOps Engineer",
							ShortDescription:     "Experience in programming with TypeScript and Golang, with the ability to write SQL scripts for database management and operations",
							SwitchingOpportunity: "",
							URL:                  "https://www.indeed.com/viewjob?jk=7195850460dbc38c",
							Date:                 mustDate("2025-03-23"),
							WithSalary:           true, // $75,000 - $95,000 per year
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
			ShortDescription: "American self-storage company who provides wide range of rentable storage properties",
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Red Hat",
			Website: "https://www.redhat.com/",
			Careers: "https://www.redhat.com/en/jobs",
			About:   "https://www.redhat.com/en/about",
			Blog:    "https://www.redhat.com/en/blog",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                3545,
				IDs:               nil,
				Alias:             "red-hat",
				Name:              "Red Hat",
				Followers:         "1M",
				Employees:         "10K+",
				AssociatedMembers: "19,585",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:     "redhatofficial",
				Followers: "2.2k",
				Verified:  false,
			},
			BlindProfile: domain.BlindProfile{
				Alias:       "Red-Hat",
				Employees:   "10,000+",
				Salary:      "$37K ~ $250K a year",
				Reviews:     "443",
				ReviewsRate: "4.0",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "red-hat",
				Employees: "18,600",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Red-Hat-EI_IE8868.11,18.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Red-Hat-Reviews-E8868.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Red-Hat-Jobs-E8868.htm",
				Jobs:        "294",
				Reviews:     "5.3K",
				Salaries:    "10K",
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
							Title:                "Software Engineer",
							ShortDescription:     "2+ years of experience working in a Linux environment with Golang or Python",
							SwitchingOpportunity: "",
							URL:                  "https://www.indeed.com/viewjob?jk=974e99549d320db1",
							Date:                 mustDate("2025-03-23"),
							WithSalary:           true, // $94,550 - $151,170 per year
							Remote:               false,
						},
						{
							Title:                "Senior Software Engineer",
							ShortDescription:     "RHEL Engineering/Go/Python/Linux",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4198719078/",
							Date:                 mustDate("2025-04-26"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Principal Software Engineer — Openshift/Golang",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4222831992/",
							Date:                 mustDate("2025-05-07"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Principal Software Engineer — Openshift/Golang",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4232684429/",
							Date:                 mustDate("2025-05-22"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Principal Software Engineer — OpenShift Serverless (Golang/Kubernetes)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4255916516/",
							Date:                 mustDate("2025-06-25"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Software Engineer — Golang",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4324902117/",
							Location:             "Bengaluru, Karnataka, India",
							Date:                 mustDate("2025-11-27"),
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
			ShortDescription: "American software company that provides open source software products to enterprises and is a subsidiary of IBM",
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "OCC",
			Website: "https://www.theocc.com/",
			Careers: "https://www.theocc.com/careers",
			About:   "https://www.theocc.com/company-information/what-is-occ",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                47420,
				IDs:               nil,
				Alias:             "occ",
				Name:              "OCC",
				Followers:         "31K",
				Employees:         "1K-5K",
				AssociatedMembers: "2,568",
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
				Alias:     "occ",
				Employees: "2,250",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-The-Options-Clearing-Corporation-EI_IE24085.11,43.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/The-Options-Clearing-Corporation-Reviews-E24085.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/The-Options-Clearing-Corporation-Jobs-E24085.htm",
				Jobs:        "51",
				Reviews:     "324",
				Salaries:    "678",
				ReviewsRate: "3.2",
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
							Title:                "Software Engineering: Java — Associate Principal",
							ShortDescription:     "Experience with app development in Golang",
							SwitchingOpportunity: "",
							URL:                  "https://www.indeed.com/viewjob?jk=0ce9ca11c038853d",
							Date:                 mustDate("2025-03-23"),
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
			ShortDescription: "American clearing house who specializes in equity derivatives clearing",
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Motive",
			Website: "https://gomotive.com/",
			Careers: "https://gomotive.com/company/careers/",
			About:   "",
			Blog:    "https://medium.com/motive-eng",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                3271606,
				IDs:               nil,
				Alias:             "motive-inc",
				Name:              "Motive",
				Followers:         "540K",
				Employees:         "1K-5K",
				AssociatedMembers: "4,492",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "GoMotive",
				Verified: false,
			},
			BlindProfile: domain.BlindProfile{
				Alias: "",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "motive",
				Employees: "4,024",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Motive-EI_IE1018424.11,17.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Motive-Reviews-E1018424.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Motive-Jobs-E1018424.htm",
				Jobs:        "102",
				Reviews:     "1.5K",
				Salaries:    "2K",
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
							Title:                "Staff System Software Engineer",
							ShortDescription:     "Proficiency in Golang, Kotlin, and C++",
							SwitchingOpportunity: "",
							URL:                  "https://www.indeed.com/viewjob?jk=8206a2447ac1e213",
							Date:                 mustDate("2025-03-23"),
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
			ShortDescription: "American technology company that creates software used by truck",
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Freddie Mac",
			Website: "https://www.freddiemac.com/",
			Careers: "https://careers.freddiemac.com/",
			About:   "https://www.freddiemac.com/about",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                3140,
				IDs:               nil,
				Alias:             "freddie-mac",
				Name:              "Freddie Mac",
				Followers:         "184K",
				Employees:         "5K-10K",
				AssociatedMembers: "11,147",
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
				Alias:     "freddie-mac",
				Employees: "9,250",
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
					Vacancies:               []domain.Vacancy{
						// Unclear if this is a Rust position https://www.indeed.com/viewjob?jk=1ddd9df71b91cb71
					},
				},
				domain.Zig:     {},
				domain.Scala:   {},
				domain.Elixir:  {},
				domain.Clojure: {},
				domain.Haskell: {},
			},
			ShortDescription: "American government-sponsored agency charged with keeping mortgage markets liquid",
			Ignore:           true, // Unclear about the tech stack
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Stoke Space",
			Website: "https://www.stokespace.com/",
			Careers: "https://www.stokespace.com/careers/",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                53254054,
				IDs:               nil,
				Alias:             "stoke-space",
				Name:              "Stoke Space",
				Followers:         "45K",
				Employees:         "201-500",
				AssociatedMembers: "215",
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
				Alias:     "stoke-space",
				Employees: "85",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-STOKE-Space-Technologies-EI_IE6521079.11,35.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/STOKE-Space-Technologies-Reviews-E6521079.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/STOKE-Space-Technologies-Jobs-E6521079.htm",
				Jobs:        "25",
				Reviews:     "11",
				Salaries:    "19",
				ReviewsRate: "4.5",
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
							Title:                "Vehicle Software Rust Developer",
							ShortDescription:     "Experience in Rust",
							SwitchingOpportunity: "",
							URL:                  "https://www.indeed.com/viewjob?jk=e6442d48c3e89a90",
							Date:                 mustDate("2025-03-23"),
							WithSalary:           true, // $94,100 - $275,600 per year
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
			ShortDescription: "American space launch company who developing a fully and rapidly reusable space launch vehicle",
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "GreyNoise Intelligence",
			Website: "https://greynoise.io/",
			Careers: "https://www.greynoise.io/careers",
			About:   "https://www.greynoise.io/about",
			Blog:    "https://www.greynoise.io/blog",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                11204989,
				IDs:               nil,
				Alias:             "greynoise",
				Name:              "GreyNoise Intelligence",
				Followers:         "11K",
				Employees:         "11-50",
				AssociatedMembers: "62",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "greynoise-intelligence",
				Verified: false,
			},
			BlindProfile: domain.BlindProfile{
				Alias: "",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "greynoise-intelligence",
				Employees: "59",
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
							Title:                "Senior Software Engineer (Platform)",
							ShortDescription:     "5+ years of production software engineering experience",
							SwitchingOpportunity: "Working with Go (Golang) for backend development",
							URL:                  "https://www.indeed.com/viewjob?jk=40120128f063af25",
							Date:                 mustDate("2025-03-24"),
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
			ShortDescription: "American cybersecurity company",
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Universal Music Group",
			Website: "https://www.universalmusic.com/",
			Careers: "https://www.umusiccareers.com/",
			About:   "https://www.universalmusic.com/company/",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                3007,
				IDs:               nil,
				Alias:             "universalmusicgroup",
				Name:              "Universal Music Group",
				Followers:         "951K",
				Employees:         "5K-10K",
				AssociatedMembers: "12,018",
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
				Alias:     "universal-music-group",
				Employees: "8,250",
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
					Vacancies:               []domain.Vacancy{
						// Unclear if this is a Go position https://www.indeed.com/viewjob?jk=b9f6faf52fb64156
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
			Ignore:           true, // Unclear about the tech stack
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Circles",
			Website: "https://circles.co/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                6449768,
				IDs:               nil,
				Alias:             "circles1",
				Name:              "Circles",
				Followers:         "69K",
				Employees:         "501-1K",
				AssociatedMembers: "1,324",
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
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Circles-EI_IE1739088.11,18.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Circles-Reviews-E1739088.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Circles-Jobs-E1739088.htm",
				Jobs:        "43",
				Reviews:     "624",
				Salaries:    "718",
				ReviewsRate: "2.7",
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
							Title:                "Staff Software Engineer (Golang / Node.js)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4187988746/",
							Date:                 mustDate("2025-03-21"),
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
			ShortDescription: "Circles is a technology company reimagining the telco industry with its SaaS platform, helping telco operators launch and operate successful digital brands",
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "PulsePoint",
			Website: "https://www.pulsepoint.com/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                2349382,
				IDs:               nil,
				Alias:             "pulsepoint",
				Name:              "PulsePoint",
				Followers:         "52K",
				Employees:         "51-200",
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
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-PulsePoint-EI_IE483561.11,21.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/PulsePoint-Reviews-E483561.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/PulsePoint-Jobs-E483561.htm",
				Jobs:        "12",
				Reviews:     "142",
				Salaries:    "230",
				ReviewsRate: "4.4",
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
							Title:                "Senior K8s/Golang Engineer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4188677602/",
							Date:                 mustDate("2025-03-22"),
							WithSalary:           false,
							Remote:               true,
						},
						{
							Title:                "Senior Platform Engineer, Golang/K8s",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4245747257/",
							Date:                 mustDate("2025-06-06"),
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
			ShortDescription: "PulsePoint is a technology company using data to accelerate healthcare marketing",
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Armada",
			Website: "https://www.armada.ai/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                97187862,
				IDs:               nil,
				Alias:             "armadaai",
				Name:              "Armada",
				Followers:         "13K",
				Employees:         "201-500",
				AssociatedMembers: "245",
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
				Alias:     "armadaai",
				Employees: "",
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
							Title:                "Software Engineer (Golang)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4190620206/",
							Date:                 mustDate("2025-05-10", "2025-04-18", "2025-03-21"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Software Engineer (Golang)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4233832043/",
							Date:                 mustDate("2025-06-12"),
							WithSalary:           true,
							Remote:               false,
						},
						{
							Title:                "Lead Engineer (Golang)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4203555647/",
							Location:             "Thiruvananthapuram, Kerala, India",
							Date:                 mustDate("2025-10-03", "2025-09-10", "2025-07-30", "2025-07-07"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Engineer (Golang)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4203558222/",
							Location:             "Thiruvananthapuram, Kerala, India",
							Date:                 mustDate("2025-07-30"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Engineer (Golang)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4289286338/",
							Location:             "Thiruvananthapuram, Kerala, India",
							Date:                 mustDate("2025-09-13", "2025-08-21"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Senior Engineer (Golang)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4203551864/",
							Location:             "Thiruvananthapuram, Kerala, India",
							Date:                 mustDate("2025-09-10"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Senior Golang Engineer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4301532995/",
							Location:             "Bellevue, WA",
							Date:                 mustDate("2025-09-16"),
							WithSalary:           true, // $144k/yr - $180k/yr
							Remote:               false,
						},
						{
							Title:                "Senior Golang Engineer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4308497799/",
							Location:             "Bellevue, WA",
							Date:                 mustDate("2025-11-12", "2025-10-22", "2025-10-01"),
							WithSalary:           true, // $144k/yr - $180k/yr
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
			ShortDescription: "Computing platform",
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Hubble.Build",
			Website: "https://hubble.build/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                13184745,
				IDs:               nil,
				Alias:             "hubble-pte.-ltd.",
				Name:              "Hubble.Build",
				Followers:         "34K",
				Employees:         "51-200",
				AssociatedMembers: "76",
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
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Hubble-Build-EI_IE2211505.11,23.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Hubble-Build-Reviews-E2211505.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Hubble-Build-Jobs-E2211505.htm",
				Jobs:        "4",
				Reviews:     "40",
				Salaries:    "54",
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
							Title:                "Software Engineer (Golang / Ruby on Rails)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4190270866/",
							Date:                 mustDate("2025-03-21"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Software Engineer (Golang / Ruby on Rails)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4216283297/",
							Date:                 mustDate("2025-05-21", "2025-04-24"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Software Engineer (Golang / Ruby on Rails)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4304601709/",
							Location:             "Ho Chi Minh City, Vietnam",
							Date:                 mustDate("2025-11-28", "2025-11-08", "2025-09-27", "2025-09-23"),
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
			ShortDescription: "Hubble.Build is a construction management platform that seamlessly connects stakeholders across the entire value chain to build better, faster, safer, and more cost-effective projects",
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Skillz",
			Website: "https://www.skillz.com/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                3104354,
				IDs:               nil,
				Alias:             "skillz",
				Name:              "Skillz",
				Followers:         "20K",
				Employees:         "51-200",
				AssociatedMembers: "215",
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
							Title:                "Backend Engineer – Java / Golang",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4192054490/",
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
			ShortDescription: "",
			Ignore:           true, // Gambling
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Unity",
			Website: "https://unity.com/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                212669,
				IDs:               []int{5676, 212669, 1388753, 15091249, 18301696},
				Alias:             "unity",
				Name:              "Unity",
				Followers:         "745K",
				Employees:         "5K-10K",
				AssociatedMembers: "6,927",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:     "unity-technologies",
				Followers: "10.8k",
				Verified:  false,
			},
			BlindProfile: domain.BlindProfile{
				Alias:       "unity%20technologies",
				Employees:   "1,001 to 5,000",
				Salary:      "$48K ~ $290K a year",
				Reviews:     "659",
				ReviewsRate: "3.2",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "unity",
				Employees: "5,440",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Unity-EI_IE455854.11,16.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Unity-Reviews-E455854.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Unity-Jobs-E455854.htm",
				Jobs:        "79",
				Reviews:     "1.8K",
				Salaries:    "4.3K",
				ReviewsRate: "3.3",
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
							Title:                "Senior Golang Backend Engineer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4191748355/",
							Date:                 mustDate("2025-05-05", "2025-04-14", "2025-03-23"),
							WithSalary:           false,
							Remote:               false,
						},
					},
				},
				domain.Rust: {},
				domain.Zig:  {},
				domain.Scala: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Senior Backend Scala Developer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4276985492/",
							Location:             "Tel Aviv-Yafo, Tel Aviv District, Israel",
							Date:                 mustDate("2025-10-02", "2025-09-10", "2025-08-20", "2025-07-28"),
							WithSalary:           false,
							Remote:               false,
						},
					},
				},
				domain.Elixir:  {},
				domain.Clojure: {},
				domain.Haskell: {},
			},
			ShortDescription: "Platform for creating and operating real-time 3D content",
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "adjoe",
			Website: "https://adjoe.io/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                12601466,
				IDs:               nil,
				Alias:             "adjoe",
				Name:              "adjoe",
				Followers:         "11K",
				Employees:         "51-200",
				AssociatedMembers: "149",
				Verified:          false,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:     "adjoeio",
				Followers: "29",
				Verified:  false,
			},
			BlindProfile: domain.BlindProfile{
				Alias: "",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias: "",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Adjoe-EI_IE2297473.11,16.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Adjoe-Reviews-E2297473.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Adjoe-Jobs-E2297473.htm",
				Jobs:        "21",
				Reviews:     "40",
				Salaries:    "61",
				ReviewsRate: "4.4",
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
							Title:                "Senior Go Developer",
							ShortDescription:     "Advertising Dashboard Team",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4190144087/",
							Date:                 mustDate("2025-03-24"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Senior Go Developer",
							ShortDescription:     "Advertising Dashboard Team",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4198154575/",
							Date:                 mustDate("2025-04-03"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Senior Go Backend Developer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4207721435/",
							Date:                 mustDate("2025-04-12"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Senior Go Backend Developer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4215663737/",
							Date:                 mustDate("2025-04-27"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Senior Go Developer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4220387316/",
							Date:                 mustDate("2025-05-04"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Go Backend Developer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4233718886/",
							Date:                 mustDate("2025-05-24"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Senior Go Developer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4238684540/",
							Date:                 mustDate("2025-06-03"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Senior Go Backend Developer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4248017872/",
							Date:                 mustDate("2025-06-11"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Senior Go Backend Developer",
							ShortDescription:     "Advertising Dashboard Team",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4256997924/",
							Date:                 mustDate("2025-07-02"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Go Backend Developer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4271281097/",
							Location:             "Hamburg, Hamburg, Germany",
							Date:                 mustDate("2025-07-28"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Senior Go Backend Developer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4291790681/",
							Location:             "Hamburg, Hamburg, Germany",
							Date:                 mustDate("2025-09-02"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Senior Go Backend Developer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4306523613/",
							Location:             "Hamburg, Hamburg, Germany",
							Date:                 mustDate("2025-10-01"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Senior Go Backend Developer",
							ShortDescription:     "Publisher Experience Team",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4347936773/",
							Location:             "Hamburg, Hamburg, Germany",
							Date:                 mustDate("2025-11-29"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Senior Go Backend Developer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4348521903/",
							Location:             "Hamburg, Hamburg, Germany",
							Date:                 mustDate("2025-12-05"),
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
				domain.IndustryAdTech,
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Docebo",
			Website: "https://www.docebo.com/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                210447,
				IDs:               nil,
				Alias:             "docebo",
				Name:              "Docebo",
				Followers:         "89K",
				Employees:         "501-1K",
				AssociatedMembers: "1,023",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:     "docebo",
				Followers: "18",
				Verified:  false,
			},
			BlindProfile: domain.BlindProfile{
				Alias: "",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "docebo",
				Employees: "700",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Docebo-EI_IE1293093.11,17.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Docebo-Reviews-E1293093.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Docebo-Jobs-E1293093.htm",
				Jobs:        "",
				Reviews:     "313",
				Salaries:    "695",
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
							Title:                "Backend Developer – Go",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4192297683/",
							Date:                 mustDate("2025-03-24"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Software Engineer (Golang)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4205969730/",
							Date:                 mustDate("2025-05-22", "2025-05-07", "2025-04-24"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Software Engineer (Golang)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4205975145/",
							Date:                 mustDate("2025-05-15", "2025-05-07"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Software Engineer Go",
							ShortDescription:     "Platform Team",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4255704335/",
							Location:             "Italy",
							Date:                 mustDate("2025-07-15"),
							WithSalary:           false,
							Remote:               true,
						},
						{
							Title:                "Software Engineer – Golang",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4283458837/",
							Location:             "Gibraltar",
							Date:                 mustDate("2025-09-02", "2025-08-12"),
							WithSalary:           false,
							Remote:               true,
						},
						{
							Title:                "Software Engineer Go",
							ShortDescription:     "Platform Team",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4292890797/",
							Location:             "Italy",
							Date:                 mustDate("2025-09-16", "2025-08-30"),
							WithSalary:           false,
							Remote:               true,
						},
						{
							Title:                "Software Engineer (Go)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4296187995/",
							Location:             "Birmingham, England, United Kingdom",
							Date:                 mustDate("2025-09-06"),
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
			ShortDescription: "Learning platform",
			Industries: []domain.Industry{
				domain.IndustryEdTech,
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Sona",
			Website: "https://www.getsona.com/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                79810814,
				IDs:               nil,
				Alias:             "getsona",
				Name:              "Sona (getsona.com)",
				Followers:         "10K",
				Employees:         "51-200",
				AssociatedMembers: "137",
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
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Sona-UK-EI_IE5901353.11,18.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Sona-UK-Reviews-E5901353.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Sona-UK-Jobs-E5901353.htm",
				Jobs:        "",
				Reviews:     "14",
				Salaries:    "20",
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
							Title:                "AI Software Engineer (Elixir, LLMs)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4175895182/",
							Date:                 mustDate("2025-03-22"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Full Stack Engineer (Elixir + LiveView)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4175892408/",
							Date:                 mustDate("2025-04-12", "2025-03-30"),
							WithSalary:           false,
							Remote:               true,
						},
						{
							Title:                "Full Stack Engineer (Elixir, LiveView, LLMs)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4209148989/",
							Date:                 mustDate("2025-05-06", "2025-04-21"),
							WithSalary:           false,
							Remote:               true,
						},
					},
				},
				domain.Clojure: {},
				domain.Haskell: {},
			},
			ShortDescription: "Workforce management solution",
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "ActiveProspect",
			Website: "https://activeprospect.com/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                523272,
				IDs:               []int{523272, 10653877},
				Alias:             "activeprospect",
				Name:              "ActiveProspect",
				Followers:         "21K",
				Employees:         "51-200",
				AssociatedMembers: "150",
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
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-ActiveProspect-EI_IE977924.11,25.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/ActiveProspect-Reviews-E977924.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/ActiveProspect-Jobs-E977924.htm",
				Jobs:        "6",
				Reviews:     "74",
				Salaries:    "79",
				ReviewsRate: "4.5",
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
							URL:                  "https://www.linkedin.com/jobs/view/4190011323/",
							Date:                 mustDate("2025-03-20"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Senior Software Engineer (Elixir)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4201324957/",
							Date:                 mustDate("2025-04-04"),
							WithSalary:           false,
							Remote:               false,
						},
					},
				},
				domain.Clojure: {},
				domain.Haskell: {},
			},
			ShortDescription: "Consent-based marketing platform",
			Industries: []domain.Industry{
				domain.IndustryMarTech,
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Ledger",
			Website: "https://www.ledger.com/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                5389345,
				IDs:               nil,
				Alias:             "ledgerhq",
				Name:              "Ledger",
				Followers:         "79K",
				Employees:         "501-1K",
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
							Title:                "Senior Scala Backend Engineer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4192446500/",
							Date:                 mustDate("2025-03-24"),
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
			Ignore:           true, // Crypto
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "TileDB",
			Website: "https://tiledb.com/",
			Careers: "https://ats.rippling.com/tiledb-careers/jobs",
			About:   "https://tiledb.com/about/",
			Blog:    "https://tiledb.com/blog",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                11301003,
				IDs:               nil,
				Alias:             "tiledb-inc",
				Name:              "TileDB",
				Followers:         "13K",
				Employees:         "51-200",
				AssociatedMembers: "75",
				Verified:          false,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:     "TileDB-Inc",
				Followers: "112",
				Verified:  false,
			},
			BlindProfile: domain.BlindProfile{
				Alias: "",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias: "",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-TileDB-EI_IE7553494.11,17.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/TileDB-Reviews-E7553494.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/TileDB-Jobs-E7553494.htm",
				Jobs:        "5",
				Reviews:     "3",
				Salaries:    "9",
				ReviewsRate: "5.0",
				Verified:    false,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "",
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 4,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Staff Software Engineer – Golang (required)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4201033819/",
							Date:                 mustDate("2025-04-04"),
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
			ShortDescription: "Embedded storage engine designed to support the storage and access of both dense and sparse multi-dimensional arrays",
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Redpanda Data",
			Website: "https://redpanda.com/",
			Careers: "https://www.redpanda.com/careers",
			About:   "https://www.redpanda.com/about-us",
			Blog:    "https://www.redpanda.com/blog",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                35528665,
				IDs:               nil,
				Alias:             "redpanda-data",
				Name:              "Redpanda Data",
				Followers:         "19K",
				Employees:         "51-200",
				AssociatedMembers: "170",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "redpanda-data",
				Verified: true,
			},
			BlindProfile: domain.BlindProfile{
				Alias: "",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "redpanda-data",
				Employees: "183",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Redpanda-Data-EI_IE8570301.11,24.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Redpanda-Data-Reviews-E8570301.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Redpanda-Data-Jobs-E8570301.htm",
				Jobs:        "5",
				Reviews:     "18",
				Salaries:    "30",
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
							Title:                "Software Engineer",
							ShortDescription:     "Strong understanding and experience of Go",
							SwitchingOpportunity: "",
							URL:                  "https://www.indeed.com/viewjob?jk=5c8079196dde8330",
							Date:                 mustDate("2025-03-25"),
							WithSalary:           true, // $185,000 - $225,000 per year
							Remote:               true,
						},
					},
				},
				domain.Rust: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{
						// Waiting for Rust position
					},
				},
				domain.Zig:     {},
				domain.Scala:   {},
				domain.Elixir:  {},
				domain.Clojure: {},
				domain.Haskell: {},
			},
			ShortDescription: "Modern streaming data platform for developers",
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Quizlet",
			Website: "https://quizlet.com/",
			Careers: "https://quizlet.com/careers",
			About:   "https://quizlet.com/mission",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                807920,
				IDs:               nil,
				Alias:             "quizlet",
				Name:              "Quizlet",
				Followers:         "47K",
				Employees:         "201-500",
				AssociatedMembers: "572",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "quizlet",
				Verified: true,
			},
			BlindProfile: domain.BlindProfile{
				Alias: "",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "quizlet",
				Employees: "100",
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
					Vacancies:               []domain.Vacancy{
						// Waiting for Go position
					},
				},
				domain.Rust:    {},
				domain.Zig:     {},
				domain.Scala:   {},
				domain.Elixir:  {},
				domain.Clojure: {},
				domain.Haskell: {},
			},
			ShortDescription: "American company that provides tools for studying and learning",
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Svix",
			Website: "https://www.svix.com/",
			Careers: "https://www.svix.com/careers/",
			About:   "https://www.svix.com/about/",
			Blog:    "https://www.svix.com/blog/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                71559487,
				IDs:               nil,
				Alias:             "svix",
				Name:              "Svix",
				Followers:         "2K",
				Employees:         "11-50",
				AssociatedMembers: "10",
				Verified:          false,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "svix",
				Verified: true,
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
							ShortDescription:     "Experience building production services with Rust",
							SwitchingOpportunity: "",
							URL:                  "https://www.indeed.com/viewjob?jk=c5a456ba1962bc7e",
							Date:                 mustDate("2025-03-25"),
							WithSalary:           true, // $130,000 - $160,000 per year
							Remote:               true,
						},
						{
							Title:                "Tech lead",
							ShortDescription:     "Experience with Rust, Postgres, and Redis",
							SwitchingOpportunity: "",
							URL:                  "https://www.indeed.com/viewjob?jk=1f4022f006ac8456",
							Date:                 mustDate("2025-03-25"),
							WithSalary:           true, // $190,000 - $210,000 per year
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
			ShortDescription: "Webhooks-as-a-service platform that empowers companies of all sizes to send webhooks easily and reliably",
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Anduril Industries",
			Website: "https://www.anduril.com/",
			Careers: "https://www.anduril.com/careers/",
			About:   "https://www.anduril.com/mission/",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                18293159,
				IDs:               nil,
				Alias:             "andurilindustries",
				Name:              "Anduril Industries",
				Followers:         "256K",
				Employees:         "1K-5K",
				AssociatedMembers: "4,143",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:     "anduril",
				Followers: "333",
				Verified:  true,
			},
			BlindProfile: domain.BlindProfile{
				Alias: "",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "anduril-industries",
				Employees: "530",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Anduril-EI_IE3546800.11,18.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Anduril-Reviews-E3546800.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Anduril-Jobs-E3546800.htm",
				Jobs:        "707",
				Reviews:     "251",
				Salaries:    "546",
				ReviewsRate: "3.9",
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
							Title:                "Game Developer — Sensor Simulation",
							ShortDescription:     "Rust programming experience",
							SwitchingOpportunity: "",
							URL:                  "https://www.indeed.com/viewjob?jk=1645f9d3eb114feb",
							Date:                 mustDate("2025-03-25"),
							WithSalary:           true, // $138,000 - $207,000 per year
							Remote:               false,
						},
						{
							Title:                "Senior Software Engineer, Rust",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4275736021/",
							Location:             "Broomfield, CO",
							Date:                 mustDate("2025-09-06", "2025-08-17", "2025-07-26"),
							WithSalary:           true, // $142.8k/yr - $214.2k/yr
							Remote:               false,
						},
						{
							Title:                "Senior Software Engineer, Rust",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4301541466/",
							Location:             "Fort Collins, CO",
							Date:                 mustDate("2025-10-30", "2025-10-08", "2025-09-16"),
							WithSalary:           true, // $142.8k/yr - $214.2k/yr
							Remote:               false,
						},
					},
				},
				domain.Zig:     {},
				domain.Scala:   {},
				domain.Elixir:  {},
				domain.Clojure: {},
				domain.Haskell: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Embedded Haskell Developer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/3988413256/",
							Location:             "Costa Mesa, CA",
							Date:                 mustDate("2025-09-02", "2025-08-11", "2025-07-20", "2025-06-28", "2025-06-08", "2025-05-17", "2025-04-24"),
							WithSalary:           true, // $140k/yr - $252k/yr
							Remote:               false,
						},
						{
							Title:                "Full Stack Haskell Engineer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4261799654/",
							Date:                 mustDate("2025-07-04"),
							WithSalary:           true,
							Remote:               false,
						},
						{
							Title:                "Senior Embedded Haskell Engineer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4297611580/",
							Location:             "Costa Mesa, CA",
							Date:                 mustDate("2025-09-08"),
							WithSalary:           true, // $168k/yr - $275k/yr
							Remote:               false,
						},
						{
							Title:                "Intern Embedded Haskell Developer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4311329988/",
							Location:             "Costa Mesa, CA",
							Date:                 mustDate("2025-12-10", "2025-11-19"),
							WithSalary:           false,
							Remote:               false,
						},
					},
				},
			},
			ShortDescription: "American defense technology company that specializes in autonomous systems",
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "CORYS",
			Website: "https://www.corys.com/",
			Careers: "https://www.corys.com/en/home/career/",
			About:   "https://www.corys.com/en/home/about-us/",
			Blog:    "https://www.corys.com/en/category/blog-en/engineering-blog/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                48613,
				IDs:               nil,
				Alias:             "corys",
				Name:              "CORYS",
				Followers:         "8K",
				Employees:         "201-500",
				AssociatedMembers: "541",
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
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-corys-EI_IE1128732.11,16.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/corys-Reviews-E1128732.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/corys-Jobs-E1128732.htm",
				Jobs:        "23",
				Reviews:     "22",
				Salaries:    "45",
				ReviewsRate: "3.7",
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
							Title:                "Software Developer — Power Simulation",
							ShortDescription:     "Rust or C++",
							SwitchingOpportunity: "Experience with Rust is strongly preferred",
							URL:                  "https://www.indeed.com/viewjob?jk=078b469307edba42",
							Date:                 mustDate("2025-03-25"),
							WithSalary:           true, // $85,000 - $120,000 per year
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
			ShortDescription: "French engineering company who provides real time high-fidelity dynamic process simulation, engineering, and training for process industries",
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "360Dialog",
			Website: "https://www.360dialog.com/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                2705257,
				IDs:               nil,
				Alias:             "360dialog",
				Name:              "360Dialog",
				Followers:         "24K",
				Employees:         "51-200",
				AssociatedMembers: "105",
				Verified:          false,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "360dialog",
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
			ShortDescription: "App for monitoring and enhancing WhatsApp performance in marketing campaigns",
			Ignore:           true, // Waiting for any vacancies
			SyncSources:      []domain.CompanySyncSource{domain.RustCompanies},
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Actyx",
			Website: "https://www.actyx.com/",
			Careers: "https://actyx.notion.site/Jobs-Actyx-2e33f9042fcc45bdaa0b369a1af4897d",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                10114352,
				IDs:               nil,
				Alias:             "actyx",
				Name:              "Actyx",
				Followers:         "2K",
				Employees:         "11-50",
				AssociatedMembers: "6",
				Verified:          false,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "actyx",
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
			ShortDescription: "",
			Ignore:           true, // Waiting for any vacancies
			SyncSources:      []domain.CompanySyncSource{domain.RustCompanies},
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Algorithmia",
			Website: "https://www.datarobot.com/",
			Careers: "https://www.datarobot.com/careers/open-positions/",
			About:   "https://www.datarobot.com/about-us/",
			Blog:    "https://www.datarobot.com/blog/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                89599097,
				IDs:               nil,
				Alias:             "algorithmiaio",
				Name:              "Algorithmia",
				Followers:         "165",
				Employees:         "51-200",
				AssociatedMembers: "3",
				Verified:          false,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "algorithmiaio",
				Verified: true,
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
			ShortDescription: "",
			Ignore:           true, // Waiting for any vacancies
			SyncSources:      []domain.CompanySyncSource{domain.RustCompanies},
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Amedia",
			Website: "https://www.amedia.no/",
			Careers: "https://www.amedia.no/alle-stillinger",
			About:   "https://www.amedia.no/english",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                2718150,
				IDs:               nil,
				Alias:             "amedia-as",
				Name:              "Amedia",
				Followers:         "10K",
				Employees:         "1K-5K",
				AssociatedMembers: "668",
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
			ShortDescription: "Norway newspaper publisher",
			Ignore:           true, // Waiting for any vacancies
			SyncSources:      []domain.CompanySyncSource{domain.RustCompanies},
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "ANIXE",
			Website: "https://www.anixe.pl/",
			Careers: "https://www.anixe.pl/jobs/",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                126161,
				IDs:               nil,
				Alias:             "anixe-polska-sp--z-o-o-",
				Name:              "ANIXE",
				Followers:         "4K",
				Employees:         "51-200",
				AssociatedMembers: "85",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "anixe",
				Verified: false,
			},
			BlindProfile: domain.BlindProfile{
				Alias: "",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "anixe",
				Employees: "121",
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
			ShortDescription: "Polish engineering company that specializes for the travel, tourism, and other industries",
			Ignore:           true, // Waiting for any vacancies
			SyncSources:      []domain.CompanySyncSource{domain.RustCompanies},
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "AppSignal",
			Website: "https://www.appsignal.com/",
			Careers: "https://www.appsignal.com/jobs",
			About:   "https://www.appsignal.com/about",
			Blog:    "https://blog.appsignal.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                2898136,
				IDs:               nil,
				Alias:             "appsignal",
				Name:              "AppSignal",
				Followers:         "866",
				Employees:         "11-50",
				AssociatedMembers: "18",
				Verified:          false,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "appsignal",
				Verified: true,
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
			ShortDescription: "Error tracker and performance monitor service for Ruby, Node, JavaScript and Elixir",
			Ignore:           true, // Waiting for any vacancies
			SyncSources:      []domain.CompanySyncSource{domain.RustCompanies},
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Atlassian",
			Website: "https://www.atlassian.com/",
			Careers: "https://www.atlassian.com/company/careers",
			About:   "https://www.atlassian.com/company",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                22688,
				IDs:               nil,
				Alias:             "atlassian",
				Name:              "Atlassian",
				Followers:         "2M",
				Employees:         "10K+",
				AssociatedMembers: "17,754",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "atlassian",
				Verified: true,
			},
			BlindProfile: domain.BlindProfile{
				Alias: "",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "atlassian",
				Employees: "6,000",
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
			ShortDescription: "Australian software company that makes software development and collaboration tools targeted at software developers",
			Ignore:           true, // Waiting for any vacancies
			SyncSources:      []domain.CompanySyncSource{domain.RustCompanies},
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Bencher",
			Website: "https://bencher.dev/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                18757175,
				IDs:               nil,
				Alias:             "bencher",
				Name:              "Bencher",
				Followers:         "66",
				Employees:         "0-1",
				AssociatedMembers: "9",
				Verified:          false,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "bencherdev",
				Verified: true,
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
			ShortDescription: "",
			Ignore:           true, // Waiting for any vacancies
			SyncSources:      []domain.CompanySyncSource{domain.RustCompanies},
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Bitfury",
			Website: "https://bitfury.com/",
			Careers: "",
			About:   "https://bitfury.com/about",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                5178019,
				IDs:               nil,
				Alias:             "bitfury",
				Name:              "Bitfury",
				Followers:         "16K",
				Employees:         "501-1K",
				AssociatedMembers: "89",
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
				Alias:     "the-bitfury-group",
				Employees: "240",
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
			ShortDescription: "Full-service blockchain technology company",
			Ignore:           true, // Crypto, waiting for any vacancies
			SyncSources:      []domain.CompanySyncSource{domain.RustCompanies},
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Braintree",
			Website: "https://www.braintreepayments.com/",
			Careers: "https://careers.pypl.com/",
			About:   "https://about.pypl.com/about-us/",
			Blog:    "https://braintree.github.io/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                1310825,
				IDs:               nil,
				Alias:             "braintree",
				Name:              "Braintree",
				Followers:         "46K",
				Employees:         "501-1K",
				AssociatedMembers: "249",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "braintree",
				Verified: false,
			},
			BlindProfile: domain.BlindProfile{
				Alias: "",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "braintree",
				Employees: "510",
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
			ShortDescription: "American company that primarily deals in mobile and web payment systems for e-commerce companies",
			Industries: []domain.Industry{
				domain.IndustryFinTech,
			},
			Ignore:      true, // Waiting for any vacancies
			SyncSources: []domain.CompanySyncSource{domain.RustCompanies},
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Brave",
			Website: "https://brave.com/",
			Careers: "https://brave.com/careers/",
			About:   "https://brave.com/about/",
			Blog:    "https://brave.com/blog/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                7929354,
				IDs:               nil,
				Alias:             "brave-software",
				Name:              "Brave",
				Followers:         "51K",
				Employees:         "51-200",
				AssociatedMembers: "272",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "brave",
				Verified: true,
			},
			BlindProfile: domain.BlindProfile{
				Alias: "",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "brave-software",
				Employees: "150",
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
			ShortDescription: "American privacy-focused browser",
			Ignore:           true, // Waiting for any vacancies
			SyncSources:      []domain.CompanySyncSource{domain.RustCompanies},
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Ceph",
			Website: "https://ceph.io/",
			Careers: "https://ceph.io/en/community/jobs/",
			About:   "https://ceph.io/en/foundation/about/",
			Blog:    "https://ceph.io/en/news/blog/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                2455652,
				IDs:               nil,
				Alias:             "ceph",
				Name:              "Ceph",
				Followers:         "8K",
				Employees:         "2-10",
				AssociatedMembers: "18",
				Verified:          false,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "ceph",
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
			ShortDescription: "Free and open-source software-defined storage platform",
			Ignore:           true, // Waiting for any vacancies
			SyncSources:      []domain.CompanySyncSource{domain.RustCompanies},
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Progress Chef",
			Website: "https://www.chef.io/",
			Careers: "https://www.chef.io/careers",
			About:   "https://www.progress.com/company",
			Blog:    "https://www.chef.io/blog",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                223176,
				IDs:               nil,
				Alias:             "chef-software",
				Name:              "Progress Chef",
				Followers:         "33K",
				Employees:         "201-500",
				AssociatedMembers: "162",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "chef",
				Verified: true,
			},
			BlindProfile: domain.BlindProfile{
				Alias: "",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "progress-chef",
				Employees: "160",
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
			ShortDescription: "Configuration management tool is used to streamline the task of configuring and maintaining a company's servers",
			Ignore:           true, // Waiting for any vacancies
			SyncSources:      []domain.CompanySyncSource{domain.RustCompanies},
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Clever Cloud",
			Website: "https://www.clever-cloud.com/",
			Careers: "",
			About:   "https://www.clever-cloud.com/about/",
			Blog:    "https://www.clever-cloud.com/blog/?cat=engineering#grid",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                1202148,
				IDs:               nil,
				Alias:             "clever-cloud",
				Name:              "Clever Cloud",
				Followers:         "6K",
				Employees:         "51-200",
				AssociatedMembers: "78",
				Verified:          false,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "clevercloud",
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
			ShortDescription: "IT Automation platform to manage all the ops work",
			Ignore:           true, // Waiting for any vacancies
			SyncSources:      []domain.CompanySyncSource{domain.RustCompanies},
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Craft AI",
			Website: "https://www.craft.ai/",
			Careers: "https://craft-ai.welcomekit.co/",
			About:   "https://www.craft.ai/en/qui-sommes-nous",
			Blog:    "https://www.craft.ai/en/blog",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                9460012,
				IDs:               nil,
				Alias:             "craft-ai",
				Name:              "Craft AI",
				Followers:         "4K",
				Employees:         "11-50",
				AssociatedMembers: "30",
				Verified:          false,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "craft-ai",
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
			ShortDescription: "",
			Ignore:           true, // Waiting for any vacancies
			SyncSources:      []domain.CompanySyncSource{domain.RustCompanies},
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Cryptape",
			Website: "https://cryptape.com/",
			Careers: "",
			About:   "",
			Blog:    "https://blog.cryptape.com/series/engineering",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                18153446,
				IDs:               nil,
				Alias:             "realcryptape",
				Name:              "Cryptape",
				Followers:         "735",
				Employees:         "51-200",
				AssociatedMembers: "23",
				Verified:          false,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "cryptape",
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
							Title:                "Software Engineer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/3964838010/",
							Date:                 mustDate("2024-06-29"),
							WithSalary:           true, // CN¥25K/month - CN¥50K/month
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
			ShortDescription: "Provider of blockchain technology for the cryptocurrency",
			Ignore:           true, // Crypto
			SyncSources:      []domain.CompanySyncSource{domain.RustCompanies},
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Deepgram",
			Website: "https://deepgram.com/",
			Careers: "https://deepgram.com/careers",
			About:   "https://deepgram.com/about",
			Blog:    "https://deepgram.com/learn/article",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                7602597,
				IDs:               nil,
				Alias:             "deepgram",
				Name:              "Deepgram",
				Followers:         "21K",
				Employees:         "51-200",
				AssociatedMembers: "175",
				Verified:          false,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "deepgram",
				Verified: false,
			},
			BlindProfile: domain.BlindProfile{
				Alias: "",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "deepgram",
				Employees: "160",
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
			ShortDescription: "AI company specializing in advanced speech recognition and transcription technology",
			Ignore:           true, // Waiting for any vacancies
			SyncSources:      []domain.CompanySyncSource{domain.RustCompanies},
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Imperva",
			Website: "https://www.imperva.com/",
			Careers: "https://www.imperva.com/company/careers/",
			About:   "https://www.imperva.com/company/about/",
			Blog:    "https://www.imperva.com/blog/?category=engineering",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                14312,
				IDs:               nil,
				Alias:             "imperva",
				Name:              "Imperva",
				Followers:         "91K",
				Employees:         "1K-5K",
				AssociatedMembers: "1,584",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "imperva",
				Verified: false,
			},
			BlindProfile: domain.BlindProfile{
				Alias: "",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "imperva",
				Employees: "1,000",
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
			ShortDescription: "American cyber security software and services company",
			Industries: []domain.Industry{
				domain.IndustryCyberSecurity,
			},
			Ignore:      true, // Waiting for any vacancies
			SyncSources: []domain.CompanySyncSource{domain.RustCompanies},
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Espressif Systems",
			Website: "https://www.espressif.com/",
			Careers: "https://www.espressif.com/en/careers",
			About:   "https://www.espressif.com/en/company/about-espressif",
			Blog:    "https://developer.espressif.com/blog/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                1211602,
				IDs:               nil,
				Alias:             "espressif-systems",
				Name:              "Espressif Systems",
				Followers:         "83K",
				Employees:         "501-1K",
				AssociatedMembers: "330",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "espressif",
				Verified: true,
			},
			BlindProfile: domain.BlindProfile{
				Alias: "",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "espressif-systems",
				Employees: "420",
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
			ShortDescription: "Chinese fabless semiconductor company",
			Ignore:           true, // Wennie, waiting for any vacancies
			SyncSources:      []domain.CompanySyncSource{domain.RustCompanies},
		},
	}
}
