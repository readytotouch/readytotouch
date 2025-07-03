package db

import "github.com/readytotouch/readytotouch/internal/domain"

func companiesOutsourceCompanies() []domain.CompanyProfile {
	return []domain.CompanyProfile{
		// OLD
		{
			ID:      0,                           // system
			Type:    domain.CompanyTypeOutsource, // system
			Name:    "Cognizant",
			Website: "https://www.cognizant.com/",
			Careers: "https://careers.cognizant.com/global-en/",
			About:   "https://www.cognizant.com/au/en/about-cognizant",
			Blog:    "https://www.cognizant.com/se/en/insights/blog/home-page-se",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                1680,
				Alias:             "cognizant",
				Name:              "Cognizant",
				Followers:         "8M",
				Employees:         "10K+",
				AssociatedMembers: "326,260",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "",
				Verified: false,
			},
			BlindProfile: domain.BlindProfile{
				Alias:       "Cognizant",
				Employees:   "10,000+",
				Salary:      "$50K ~ $220K a year",
				Reviews:     "612",
				ReviewsRate: "3.1",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "cognizant",
				Employees: "309,110",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Cognizant-Technology-Solutions-EI_IE8014.11,41.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Cognizant-Technology-Solutions-Reviews-E8014.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Cognizant-Technology-Solutions-Jobs-E8014.htm",
				Jobs:        "2.8K",
				Reviews:     "117K",
				Salaries:    "226K",
				ReviewsRate: "3.7",
				Verified:    true,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "Cognizant-Technology-Solutions",
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "https://www.youtube.com/@cognizant",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Senior Golang Developer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/3914839625/",
							Date:                 mustDate("2024-05-24"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Golang Developer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4120737001/",
							Date:                 mustDate("2025-01-24"), // mustDate("2025-01-11"),
							WithSalary:           true,                   // The annual salary for this position is between $95K to 115K depending on experience and other qualifications of the successful candidate
							Remote:               false,
						},
						{
							Title:                "Senior AWS developer with Golang",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4149258542/",
							Date:                 mustDate("2025-02-10"),
							WithSalary:           true, // $55.000 — $114.000 per year
							Remote:               false,
						},
						{
							Title:                "Senior Full Stack Engineer (Golang/AWS)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4144963593/",
							Date:                 mustDate("2025-02-14"),
							WithSalary:           true, // salary for this position is between $68,000 – $114,000
							Remote:               false,
						},
						{
							Title:                "Golang Developer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4181350285/",
							Date:                 mustDate("2025-03-14"),
							WithSalary:           false,
							Remote:               true,
						},
					},
				},
				domain.Rust: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Rust Developer (Remote)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4066607632/",
							Date:                 mustDate("2024-11-07"),
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
							Title:                "Data Engineer — AWS/Hive/Scala/Spark",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4124669157/",
							Date:                 mustDate("2025-01-15"),
							WithSalary:           true, // $68K — $108K per year
							Remote:               false,
						},
						{
							Title:                "Senior Java Developer with Scala/Kafka Experience",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4210424636/",
							Date:                 mustDate("2025-04-18"),
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
			Ignore: true, // Outsource
		},
		{
			ID:      0,                           // system
			Type:    domain.CompanyTypeOutsource, // system
			Name:    "Sparq",
			Website: "https://www.teamsparq.com/",
			Careers: "https://www.teamsparq.com/careers/",
			About:   "https://www.teamsparq.com/who-we-are/",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                90679425,
				Alias:             "teamsparq",
				Name:              "Sparq",
				Followers:         "13K",
				Employees:         "501-1K",
				AssociatedMembers: "790",
				Verified:          true,
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Senior Golang Developer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4049853199/",
							Date:                 mustDate("2025-01-14"), // mustDate("2024-12-03")
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
			Ignore: true,
		},
		{
			ID:      0,                           // system
			Type:    domain.CompanyTypeOutsource, // system
			Name:    "Coherent Solutions",
			Website: "https://www.coherentsolutions.com/",
			Careers: "https://www.coherentsolutions.com/careers",
			About:   "https://www.coherentsolutions.com/about-us",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                38745,
				Alias:             "coherent-solutions",
				Name:              "Coherent Solutions",
				Followers:         "13K",
				Employees:         "1K-5K",
				AssociatedMembers: "680",
				Verified:          true,
			},
			Ignore: true, // The deleted outsource company was added by mistake
		},
		{
			ID:      0,                           // system
			Type:    domain.CompanyTypeOutsource, // system
			Name:    "Antmicro",
			Website: "https://antmicro.com/",
			Careers: "https://careers.antmicro.com/",
			About:   "https://antmicro.com/about/company/",
			Blog:    "https://antmicro.com/blog/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                3267482,
				Alias:             "antmicro-ltd",
				Name:              "Antmicro",
				Followers:         "6K",
				Employees:         "51-200",
				AssociatedMembers: "83",
				Verified:          true,
			},
			ShortDescription: "Sweden software-driven tech company",
			Ignore:           true, // The deleted outsource company was added by mistake
		},
		{
			ID:      0,                           // system
			Type:    domain.CompanyTypeOutsource, // system
			Name:    "Arcanys",
			Website: "https://www.arcanys.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                1472152,
				IDs:               nil,
				Alias:             "arcanys",
				Name:              "Arcanys",
				Followers:         "27K",
				Employees:         "201-500",
				AssociatedMembers: "273",
				Verified:          true,
			},
			Ignore: true, // Outsourcing
		},
		{
			ID:      0,                           // system
			Type:    domain.CompanyTypeOutsource, // system
			Name:    "Skillspark",
			Website: "https://www.skillspark.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                80284986,
				IDs:               nil,
				Alias:             "skillsparkab",
				Name:              "Skillspark",
				Followers:         "68K",
				Employees:         "51-200",
				AssociatedMembers: "33",
				Verified:          false,
			},
			Ignore: true, // Outsourcing
		},
		{
			ID:      0,                           // system
			Type:    domain.CompanyTypeOutsource, // system
			Name:    "Qinshift",
			Website: "https://qinshift.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                92778277,
				IDs:               nil,
				Alias:             "qinshift",
				Name:              "Qinshift",
				Followers:         "52K",
				Employees:         "1K-5K",
				AssociatedMembers: "1,717",
				Verified:          true,
			},
			Ignore: true, // Outsourcing
		},
		{
			ID:      0,                           // system
			Type:    domain.CompanyTypeOutsource, // system
			Name:    "Brillio",
			Website: "https://www.brillio.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                5022712,
				IDs:               nil,
				Alias:             "brillio",
				Name:              "Brillio",
				Followers:         "475K",
				Employees:         "5K-10K",
				AssociatedMembers: "5,206",
				Verified:          true,
			},
			Ignore: true, // Outsourcing
		},
		{
			ID:      0,                           // system
			Type:    domain.CompanyTypeOutsource, // system
			Name:    "Akvelon",
			Website: "https://akvelon.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                309542,
				IDs:               nil,
				Alias:             "akvelon",
				Name:              "Akvelon, Inc.",
				Followers:         "20K",
				Employees:         "1K-5K",
				AssociatedMembers: "632",
				Verified:          false,
			},
			Ignore: true, // Outsourcing
		},
		{
			ID:      0,                           // system
			Type:    domain.CompanyTypeOutsource, // system
			Name:    "Multi Media",
			Website: "https://multimediallc.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                10496448,
				IDs:               nil,
				Alias:             "multimedia-llc",
				Name:              "Multi Media, LLC",
				Followers:         "33K",
				Employees:         "201-500",
				AssociatedMembers: "308",
				Verified:          false,
			},
			Ignore: true, // Outsourcing
		},
		{
			ID:      0,                           // system
			Type:    domain.CompanyTypeOutsource, // system
			Name:    "BairesDev",
			Website: "https://www.bairesdev.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                798671,
				IDs:               nil,
				Alias:             "bairesdev",
				Name:              "BairesDev",
				Followers:         "656K",
				Employees:         "1K-5K",
				AssociatedMembers: "3,011",
				Verified:          true,
			},
			Ignore: true, // Outsourcing
		},
		{
			ID:      0,                           // system
			Type:    domain.CompanyTypeOutsource, // system
			Name:    "Ubique Systems",
			Website: "https://www.ubique-systems.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                939675,
				IDs:               nil,
				Alias:             "ubique-systems",
				Name:              "Ubique Systems",
				Followers:         "229K",
				Employees:         "501-1K",
				AssociatedMembers: "307",
				Verified:          true,
			},
			Ignore: true, // Outsourcing
		},
		{
			ID:      0,                           // system
			Type:    domain.CompanyTypeOutsource, // system
			Name:    "Immunant",
			Website: "https://immunant.com/",
			Careers: "https://immunant.com/jobs/",
			About:   "",
			Blog:    "https://immunant.com/blog/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                27043749,
				IDs:               nil,
				Alias:             "immunant",
				Name:              "Immunant",
				Followers:         "167",
				Employees:         "2-10",
				AssociatedMembers: "4",
				Verified:          false,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "immunant",
				Verified: false,
			},
			ShortDescription: "Software consulting company",
			Ignore:           true, // Outsource
			SyncSources:      []domain.CompanySyncSource{domain.RustCompanies},
		},
		{
			ID:      0,                           // system
			Type:    domain.CompanyTypeOutsource, // system
			Name:    "Integer 32, LLC",
			Website: "https://www.integer32.com/",
			Careers: "",
			About:   "https://www.integer32.com/about/",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                10822533,
				IDs:               nil,
				Alias:             "integer-32-llc",
				Name:              "Integer 32, LLC",
				Followers:         "126",
				Employees:         "2-10",
				AssociatedMembers: "2",
				Verified:          false,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "integer32llc",
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
			Ignore:           true, // Outsource
			SyncSources:      []domain.CompanySyncSource{domain.RustCompanies},
		},
		{
			ID:      0,                           // system
			Type:    domain.CompanyTypeOutsource, // system
			Name:    "Red Iron Team",
			Website: "https://red-iron.eu/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                0,
				Alias:             "red-iron-team",
				Name:              "Red Iron Team",
				Followers:         "195",
				Employees:         "2-10",
				AssociatedMembers: "0",
				Verified:          false,
			},
			ShortDescription: "French software company",
			Ignore:           true, // Outsource, should be double-checked
			SyncSources:      []domain.CompanySyncSource{domain.RustCompanies},
		},
		{
			ID:      0,                           // system
			Type:    domain.CompanyTypeOutsource, // system
			Name:    "Rustunit",
			Website: "https://rustunit.com/",
			Careers: "",
			About:   "",
			Blog:    "https://rustunit.com/blog/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                100357404,
				Alias:             "rustunit",
				Name:              "Rustunit",
				Followers:         "288",
				Employees:         "2-10",
				AssociatedMembers: "1",
				Verified:          false,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "rustunit",
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
			Ignore:           true, // Outsource, should be double-checked
			SyncSources:      []domain.CompanySyncSource{domain.RustCompanies},
		},
		{
			ID:      0,                           // system
			Type:    domain.CompanyTypeOutsource, // system
			Name:    "49nord GmbH",
			Website: "https://49nord.de/",
			Careers: "",
			About:   "https://49nord.de/en/company",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                18586171,
				IDs:               nil,
				Alias:             "49nord",
				Name:              "49nord GmbH",
				Followers:         "18",
				Employees:         "2-10",
				AssociatedMembers: "1",
				Verified:          false,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "49nord",
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
			Ignore:           true, // Outsource, waiting for any vacancies
			SyncSources:      []domain.CompanySyncSource{domain.RustCompanies},
		},
		{
			ID:      0,                           // system
			Type:    domain.CompanyTypeOutsource, // system
			Name:    "corrode",
			Website: "https://corrode.dev/",
			Careers: "",
			About:   "https://corrode.dev/about/",
			Blog:    "https://corrode.dev/blog/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                82209241,
				IDs:               nil,
				Alias:             "corrode",
				Name:              "corrode",
				Followers:         "156",
				Employees:         "0-1",
				AssociatedMembers: "0",
				Verified:          false,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "corrode",
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
			ShortDescription: "Germans Rust consulting company",
			Ignore:           true, // Outsource, waiting for any vacancies
			SyncSources:      []domain.CompanySyncSource{domain.RustCompanies},
		},
		{
			ID:      0,                           // system
			Type:    domain.CompanyTypeOutsource, // system
			Name:    "Ferrous Systems GmbH",
			Website: "https://ferrous-systems.com/",
			Careers: "",
			About:   "",
			Blog:    "https://ferrous-systems.com/blog/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                11847446,
				IDs:               nil,
				Alias:             "ferrous-systems",
				Name:              "Ferrous Systems GmbH",
				Followers:         "3K",
				Employees:         "11-50",
				AssociatedMembers: "22",
				Verified:          false,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "ferrous-systems",
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
			ShortDescription: "Rust consulting company",
			Ignore:           true, // Outsource, waiting for any vacancies
			SyncSources:      []domain.CompanySyncSource{domain.RustCompanies},
		},
		{
			ID:      0,                           // system
			Type:    domain.CompanyTypeOutsource, // system
			Name:    "Photon",
			Website: "https://photon.com/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                165464,
				IDs:               nil,
				Alias:             "photon-interactive",
				Name:              "Photon",
				Followers:         "526K",
				Employees:         "5K-10K",
				AssociatedMembers: "5,404",
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
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4195243048/",
							Date:                 mustDate("2025-04-01"),
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
			Ignore:           true, // Outsource
		},
		{
			ID:      0,                           // system
			Type:    domain.CompanyTypeOutsource, // system
			Name:    "Vinova",
			Website: "https://vinova.sg/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                1614934,
				IDs:               nil,
				Alias:             "vinova-sg",
				Name:              "Vinova",
				Followers:         "4K",
				Employees:         "201-500",
				AssociatedMembers: "",
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
					Vacancies:               []domain.Vacancy{},
				},
				domain.Rust:    {},
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
			Ignore: true,
		},

		// NEW
		{
			Type:    domain.CompanyTypeOutsource,
			Name:    "Minsait",
			Website: "https://www.minsait.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                10963,
				IDs:               nil,
				Alias:             "minsait",
				Name:              "Minsait",
				Followers:         "599K",
				Employees:         "10K+",
				AssociatedMembers: "15,668",
				Verified:          true,
			},
			Ignore: true, // Outsource
		},
		{
			Type:    domain.CompanyTypeOutsource,
			Name:    "Virtusa",
			Website: "https://www.virtusa.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                5140,
				IDs:               nil,
				Alias:             "virtusa",
				Name:              "Virtusa",
				Followers:         "1M",
				Employees:         "10K+",
				AssociatedMembers: "17,503",
				Verified:          true,
			},
			Ignore: true, // Outsource
		},
		{
			Type:    domain.CompanyTypeOutsource,
			Name:    "Qubika",
			Website: "https://qubika.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                201026,
				IDs:               nil,
				Alias:             "qubika",
				Name:              "Qubika",
				Followers:         "58K",
				Employees:         "501-1K",
				AssociatedMembers: "726",
				Verified:          true,
			},
			Ignore: true, // Outsource
		},
		{
			Type:    domain.CompanyTypeOutsource,
			Name:    "Zartis",
			Website: "https://www.zartis.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                2274396,
				IDs:               nil,
				Alias:             "zartis",
				Name:              "Zartis",
				Followers:         "48K",
				Employees:         "201-500",
				AssociatedMembers: "216",
				Verified:          true,
			},
			Ignore: true, // Outsource
		},
		{
			Type:    domain.CompanyTypeOutsource,
			Name:    "Smart Consulting",
			Website: "https://www.smartconsulting.pt/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                9356133,
				IDs:               []int{9356133, 79396476},
				Alias:             "thesmartwaytodoit",
				Name:              "Smart Consulting",
				Followers:         "121K",
				Employees:         "201-500",
				AssociatedMembers: "322",
				Verified:          false,
			},
			Ignore: true, // Outsource
		},
		{
			Type:    domain.CompanyTypeOutsource,
			Name:    "Xoriant",
			Website: "https://www.xoriant.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                166996,
				IDs:               nil,
				Alias:             "xoriant",
				Name:              "Xoriant",
				Followers:         "464K",
				Employees:         "5K-10K",
				AssociatedMembers: "4,548",
				Verified:          true,
			},
			Ignore: true, // Outsource
		},
		{
			Type:    domain.CompanyTypeOutsource,
			Name:    "Trideca",
			Website: "https://www.trideca.com.au/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                2455945,
				IDs:               nil,
				Alias:             "tridecaptyltd",
				Name:              "Trideca",
				Followers:         "7K",
				Employees:         "51-200",
				AssociatedMembers: "97",
				Verified:          true,
			},
			Ignore: true, // Outsource
		},
		{
			Type:    domain.CompanyTypeOutsource,
			Name:    "Excelia",
			Website: "https://excelia.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                735796,
				IDs:               nil,
				Alias:             "excelia",
				Name:              "Excelia",
				Followers:         "109K",
				Employees:         "201-500",
				AssociatedMembers: "325",
				Verified:          false,
			},
			Ignore: true, // Outsource
		},
		{
			Type:    domain.CompanyTypeOutsource,
			Name:    "Xebia",
			Website: "https://xebia.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                16903,
				IDs:               []int{16903, 72738535},
				Alias:             "xebia",
				Name:              "Xebia",
				Followers:         "378K",
				Employees:         "5K-10K",
				AssociatedMembers: "5,399",
				Verified:          true,
			},
			Ignore: true, // Outsource
		},
		{
			Type:    domain.CompanyTypeOutsource,
			Name:    "Coltech",
			Website: "https://www.coltech.io/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                69762147,
				IDs:               nil,
				Alias:             "coltechglobal",
				Name:              "Coltech",
				Followers:         "57K",
				Employees:         "11-50",
				AssociatedMembers: "28",
				Verified:          false,
			},
			Ignore: true, // Outsource
		},
		{
			Type:    domain.CompanyTypeOutsource,
			Name:    "Akkodis",
			Website: "https://www.akkodis.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                79383535,
				IDs:               []int{224006, 320456, 79383535},
				Alias:             "akkodis",
				Name:              "Akkodis",
				Followers:         "3M",
				Employees:         "10K+",
				AssociatedMembers: "24,836",
				Verified:          true,
			},
			Ignore: true, // Outsource
		},
		{
			Type:    domain.CompanyTypeOutsource,
			Name:    "The Select Group",
			Website: "https://www.selectgroup.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                28526,
				IDs:               []int{28526, 3056209},
				Alias:             "the-select-group",
				Name:              "The Select Group",
				Followers:         "391K",
				Employees:         "1K-5K",
				AssociatedMembers: "937",
				Verified:          true,
			},
			Ignore: true, // Outsource
		},
		{
			Type:    domain.CompanyTypeOutsource,
			Name:    "Programmers.io",
			Website: "https://programmers.io/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                13714336,
				IDs:               nil,
				Alias:             "programmers-io",
				Name:              "Programmers.io",
				Followers:         "538K",
				Employees:         "1K-5K",
				AssociatedMembers: "1,155",
				Verified:          true,
			},
			Ignore: true, // Outsource
		},
		{
			Type:    domain.CompanyTypeOutsource,
			Name:    "ValueLabs",
			Website: "https://www.valuelabs.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                13489,
				IDs:               nil,
				Alias:             "valuelabs",
				Name:              "ValueLabs",
				Followers:         "1M",
				Employees:         "5K-10K",
				AssociatedMembers: "8,506",
				Verified:          true,
			},
			Ignore: true, // Outsource
		},
		{
			Type:    domain.CompanyTypeOutsource,
			Name:    "CDW",
			Website: "https://www.cdw.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                3334,
				IDs:               []int{3334, 10951, 3646541, 10599950, 13250389, 15264947, 15911608},
				Alias:             "cdw",
				Name:              "CDW",
				Followers:         "317K",
				Employees:         "10K+",
				AssociatedMembers: "17,976",
				Verified:          true,
			},
			Ignore: true, // Outsource
		},
		{
			Type:    domain.CompanyTypeOutsource,
			Name:    "Lumenalta",
			Website: "https://lumenalta.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                719642,
				IDs:               nil,
				Alias:             "lumenalta",
				Name:              "Lumenalta",
				Followers:         "126K",
				Employees:         "501-1K",
				AssociatedMembers: "560",
				Verified:          true,
			},
			Ignore: true, // Outsource
		},
		{
			Type:    domain.CompanyTypeOutsource,
			Name:    "Dualboot Partners",
			Website: "https://dualbootpartners.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                28138288,
				IDs:               nil,
				Alias:             "dualbootpartners",
				Name:              "Dualboot Partners",
				Followers:         "13K",
				Employees:         "201-500",
				AssociatedMembers: "279",
				Verified:          true,
			},
			Ignore: true, // Outsource
		},
		{
			Type:    domain.CompanyTypeOutsource,
			Name:    "UKG",
			Website: "https://www.ukg.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                67539665,
				IDs:               nil,
				Alias:             "ukg",
				Name:              "UKG",
				Followers:         "268K",
				Employees:         "10K+",
				AssociatedMembers: "15,144",
				Verified:          true,
			},
			Ignore: true, // Outsource
		},
		{
			Type:    domain.CompanyTypeOutsource,
			Name:    "Nagarro",
			Website: "https://www.nagarro.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                11902,
				IDs:               []int{11902, 59775},
				Alias:             "nagarro",
				Name:              "Nagarro",
				Followers:         "2M",
				Employees:         "10K+",
				AssociatedMembers: "20,911",
				Verified:          true,
			},
			Ignore: true, // Outsource
		},
		{
			Type:    domain.CompanyTypeOutsource,
			Name:    "E-Solutions",
			Website: "https://e-solutionsinc.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                750072,
				IDs:               nil,
				Alias:             "e-solutions-inc",
				Name:              "E-Solutions",
				Followers:         "336K",
				Employees:         "1K-5K",
				AssociatedMembers: "1,771",
				Verified:          false,
			},
			Ignore: true, // Outsource
		},
		{
			Type:    domain.CompanyTypeOutsource,
			Name:    "Talan",
			Website: "https://www.talan.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                28542,
				IDs:               []int{15432, 28542, 55444, 132835, 278066, 427360, 11202100, 17908204, 33298999},
				Alias:             "talan",
				Name:              "Talan",
				Followers:         "350K",
				Employees:         "5K-10K",
				AssociatedMembers: "4,432",
				Verified:          true,
			},
			Ignore: true, // Outsource
		},
		{
			Type:    domain.CompanyTypeOutsource,
			Name:    "Sage",
			Website: "https://www.sage.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                2802,
				IDs:               []int{2802, 599552},
				Alias:             "sage-software",
				Name:              "Sage",
				Followers:         "548K",
				Employees:         "10K+",
				AssociatedMembers: "14,427",
				Verified:          true,
			},
			Ignore: true, // Outsource
		},
		{
			Type:    domain.CompanyTypeOutsource,
			Name:    "Wizeline",
			Website: "https://www.wizeline.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                3543771,
				IDs:               nil,
				Alias:             "wizeline",
				Name:              "Wizeline",
				Followers:         "76K",
				Employees:         "1K-5K",
				AssociatedMembers: "1,544",
				Verified:          true,
			},
			Ignore: true, // Outsource
		},
		{
			Type:    domain.CompanyTypeOutsource,
			Name:    "Vertex Agility",
			Website: "https://www.vertexagility.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                694216,
				IDs:               nil,
				Alias:             "vertex-agility-ltd",
				Name:              "Vertex Agility",
				Followers:         "127K",
				Employees:         "51-200",
				AssociatedMembers: "254",
				Verified:          true,
			},
			Ignore: true, // Outsource
		},
		{
			Type:    domain.CompanyTypeOutsource,
			Name:    "EXL",
			Website: "https://www.exlservice.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                163743,
				IDs:               []int{28813, 163743},
				Alias:             "exl-service",
				Name:              "EXL",
				Followers:         "1M",
				Employees:         "10K+",
				AssociatedMembers: "38,166",
				Verified:          true,
			},
			Ignore: true, // Outsource
		},
		{
			Type:    domain.CompanyTypeOutsource,
			Name:    "Viveris",
			Website: "https://www.viveris.fr/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                35418,
				IDs:               nil,
				Alias:             "viveris",
				Name:              "Viveris",
				Followers:         "21K",
				Employees:         "1K-5K",
				AssociatedMembers: "866",
				Verified:          true,
			},
			Ignore: true, // Outsource
		},
		{
			Type:    domain.CompanyTypeOutsource,
			Name:    "Compunnel Inc.",
			Website: "https://www.compunnel.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                16690,
				IDs:               nil,
				Alias:             "compunnel-software-group",
				Name:              "Compunnel Inc.",
				Followers:         "718K",
				Employees:         "5K-10K",
				AssociatedMembers: "2,620",
				Verified:          false,
			},
			Ignore: true, // Outsource
		},
		{
			Type:    domain.CompanyTypeOutsource,
			Name:    "ERNI",
			Website: "https://www.betterask.erni/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                55018,
				IDs:               []int{55018, 1414690},
				Alias:             "erni",
				Name:              "ERNI",
				Followers:         "55K",
				Employees:         "501-1K",
				AssociatedMembers: "962",
				Verified:          true,
			},
			Ignore: true, // Outsource
		},
		{
			Type:    domain.CompanyTypeOutsource,
			Name:    "emagine",
			Website: "https://www.emagine.org/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                30899,
				IDs:               []int{30899, 18487093, 18499941, 83005407},
				Alias:             "emagine",
				Name:              "emagine",
				Followers:         "139K",
				Employees:         "1K-5K",
				AssociatedMembers: "1,906",
				Verified:          true,
			},
			Ignore: true, // Outsource
		},
		{
			Type:    domain.CompanyTypeOutsource,
			Name:    "Schuberg Philis",
			Website: "https://schubergphilis.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                16104,
				IDs:               nil,
				Alias:             "schuberg-philis",
				Name:              "Schuberg Philis",
				Followers:         "14K",
				Employees:         "201-500",
				AssociatedMembers: "471",
				Verified:          true,
			},
			Ignore: true, // Outsource
		},
		{
			Type:    domain.CompanyTypeOutsource,
			Name:    "RITS",
			Website: "https://rits.center/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                5182288,
				IDs:               nil,
				Alias:             "ritscenter",
				Name:              "RITS",
				Followers:         "13K",
				Employees:         "501-1K",
				AssociatedMembers: "236",
				Verified:          true,
			},
			Ignore: true, // Outsource
		},
		{
			Type:    domain.CompanyTypeOutsource,
			Name:    "Revature",
			Website: "https://www.revature.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                10510386,
				IDs:               nil,
				Alias:             "revature",
				Name:              "Revature",
				Followers:         "324K",
				Employees:         "1K-5K",
				AssociatedMembers: "1,521",
				Verified:          true,
			},
			Ignore: true, // Outsource
		},
		{
			Type:    domain.CompanyTypeOutsource,
			Name:    "TEKsystems",
			Website: "https://www.teksystems.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                2152,
				IDs:               []int{2152, 14785},
				Alias:             "teksystems",
				Name:              "TEKsystems",
				Followers:         "1M",
				Employees:         "10K+",
				AssociatedMembers: "25,945",
				Verified:          true,
			},
			Ignore: true, // Outsource
		},
		{
			Type:    domain.CompanyTypeOutsource,
			Name:    "Centraprise",
			Website: "https://centraprise.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                13345578,
				IDs:               nil,
				Alias:             "centraprise",
				Name:              "Centraprise",
				Followers:         "178K",
				Employees:         "501-1K",
				AssociatedMembers: "276",
				Verified:          false,
			},
			Ignore: true, // Outsource
		},
		{
			Type:    domain.CompanyTypeOutsource,
			Name:    "Tata Consultancy Services",
			Website: "https://www.tcs.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                1353,
				IDs:               nil,
				Alias:             "tata-consultancy-services",
				Name:              "Tata Consultancy Services",
				Followers:         "17M",
				Employees:         "10K+",
				AssociatedMembers: "674,364",
				Verified:          true,
			},
			Ignore: true, // Outsource
		},
		{
			Type:    domain.CompanyTypeOutsource,
			Name:    "Parser",
			Website: "https://parserdigital.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                11142678,
				IDs:               nil,
				Alias:             "parser",
				Name:              "Parser",
				Followers:         "32K",
				Employees:         "201-500",
				AssociatedMembers: "201",
				Verified:          true,
			},
			Ignore: true, // Outsource
		},
		{
			Type:    domain.CompanyTypeOutsource,
			Name:    "Aderen",
			Website: "https://aderen.es/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                5151361,
				IDs:               nil,
				Alias:             "aderen-consulting-sl",
				Name:              "Aderen",
				Followers:         "17K",
				Employees:         "51-200",
				AssociatedMembers: "20",
				Verified:          true,
			},
			Ignore: true, // Outsource
		},
		{
			Type:    domain.CompanyTypeOutsource,
			Name:    "Devsu",
			Website: "https://devsu.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                1102330,
				IDs:               nil,
				Alias:             "devsu",
				Name:              "Devsu",
				Followers:         "46K",
				Employees:         "201-500",
				AssociatedMembers: "231",
				Verified:          true,
			},
			Ignore: true, // Outsource
		},
	}
}
