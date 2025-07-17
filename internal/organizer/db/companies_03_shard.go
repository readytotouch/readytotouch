package db

import (
	"github.com/readytotouch/readytotouch/internal/domain"
)

func companies03Shard() []domain.CompanyProfile {
	return []domain.CompanyProfile{
		// Some | Flix
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Flix",
			Website: "https://www.flixbus.com/",
			Careers: "https://flix.careers/",
			About:   "https://corporate.flix.com/about-flix/",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                2726149,
				Alias:             "flixbus",
				Name:              "Flix",
				Followers:         "97K",
				Employees:         "1K-5K",
				AssociatedMembers: "2,246",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "flix-tech",
				Verified: false,
			},
			BlindProfile: domain.BlindProfile{
				Alias:       "Flixbus",
				Employees:   "1,001-5,000",
				Salary:      "",
				Reviews:     "4",
				ReviewsRate: "3.5",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "flix",
				Employees: "2,250",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Flix-EI_IE976463.11,15.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Flix-Reviews-E976463.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Flix-Jobs-E976463.htm",
				Jobs:        "200",
				Reviews:     "589",
				Salaries:    "1.1K",
				ReviewsRate: "4.0",
				Verified:    true,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "Flix-Se-2",
			},
			OttaProfileSlug:   "FlixBus-1",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 17,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Middle Golang Engineer — Content Platform team",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4101217107/",
							Date:                 mustDate("2024-12-16"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Middle Back-End (Go) Engineer  — Content Platform team",
							ShortDescription:     "Backend: Go, ArangoDB, MySQL, Confluent Kafka, Snowflake, GraphQL + different AWS services",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4091237831/",
							Date:                 mustDate("2025-01-03"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Middle Golang Engineer — Content Platform team",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4093658841/",
							Date:                 mustDate("2025-01-20", "2025-01-16", "2025-01-11"),
							WithSalary:           false,
							Remote:               false, // Up to 60 days of working from (m)anywhere
						},
						{
							Title:                "Golang Engineer — Content Platform team",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4143253999/",
							Date:                 mustDate("2025-02-28", "2025-02-06"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Middle Back-End (Go) Engineer — Content Platform team",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4138356213/",
							Date:                 mustDate("2025-03-07", "2025-02-21"),
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

		// Some | Press Ganey
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Press Ganey",
			Website: "https://www.pressganey.com/",
			Careers: "https://www.pressganey.com/company/careers/",
			About:   "https://www.pressganey.com/company/",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                18007,
				Alias:             "press-ganey-associates",
				Name:              "Press Ganey",
				Followers:         "90K",
				Employees:         "1K-5K",
				AssociatedMembers: "2,070",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "",
				Verified: false,
			},
			BlindProfile: domain.BlindProfile{
				Alias:       "Press-Ganey",
				Employees:   "1,001 to 5,000",
				Salary:      "",
				Reviews:     "10",
				ReviewsRate: "2.9",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "press-ganey",
				Employees: "2,500",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Press-Ganey-EI_IE35100.11,22.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Press-Ganey-Reviews-E35100.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Press-Ganey-Jobs-E35100.htm",
				Jobs:        "57",
				Reviews:     "784",
				Salaries:    "1.2K",
				ReviewsRate: "3.2",
				Verified:    true,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "Press-Ganey-4",
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "https://www.youtube.com/@PressGaney",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Senior Golang Developer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/3898197086/",
							Date:                 mustDate("2024-05-24"),
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
			Industries:       []domain.Industry{},
		},

		// Some | Atmail
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Atmail",
			Website: "https://www.atmail.com/",
			Careers: "",
			About:   "https://www.atmail.com/company/",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                918978,
				Alias:             "atmail",
				Name:              "Atmail",
				Followers:         "5K",
				Employees:         "51-200",
				AssociatedMembers: "32",
				Verified:          false,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "",
				Verified: false,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "",
				ReviewsURL:  "",
				JobsURL:     "",
				Jobs:        "",
				Reviews:     "",
				Salaries:    "",
				ReviewsRate: "",
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
							Title:                "Golang Developer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/3843643222/",
							Date:                 mustDate("2024-04-25"),
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
			Industries:       []domain.Industry{},
		},

		// Some | Dusty Robotics
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Dusty Robotics",
			Website: "https://www.dustyrobotics.com/",
			Careers: "https://www.dustyrobotics.com/careers",
			About:   "https://www.dustyrobotics.com/about",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                33298433,
				Alias:             "dusty-robotics",
				Name:              "Dusty Robotics",
				Followers:         "20K",
				Employees:         "51-200",
				AssociatedMembers: "88",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "",
				Verified: false,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Dusty-Robotics-EI_IE3518259.11,25.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Dusty-Robotics-Reviews-E3518259.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Dusty-Robotics-Jobs-E3518259.htm",
				Jobs:        "",
				Reviews:     "25",
				Salaries:    "34",
				ReviewsRate: "4.7",
				Verified:    false,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "",
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "https://www.youtube.com/@dustyrobotics",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Senior Golang Developer fullstack/backend",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/3804188613/",
							Date:                 mustDate("2024-04-25"),
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
			Industries:       []domain.Industry{},
		},

		// Some | Cimri
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Cimri",
			Website: "https://www.cimri.com/",
			Careers: "",
			About:   "https://www.cimri.com/cimri-hakkinda",
			Blog:    "https://engineering.cimri.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                411498,
				Alias:             "cimri",
				Name:              "Cimri",
				Followers:         "32K",
				Employees:         "51-200",
				AssociatedMembers: "178",
				Verified:          false,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "",
				Verified: false,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Cimri-EI_IE2401296.11,16.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Cimri-Reviews-E2401296.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Cimri-Jobs-E2401296.htm",
				Jobs:        "",
				Reviews:     "2",
				Salaries:    "20",
				ReviewsRate: "4.0",
				Verified:    false,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "",
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "https://www.youtube.com/@cimri",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Golang Backend Developer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/3863565726/",
							Date:                 mustDate("2024-04-25"),
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
			Industries:       []domain.Industry{},
		},

		// Some | Quadcode
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Quadcode",
			Website: "https://quadcode.com/",
			Careers: "https://jobs.quadcode.com/jobs",
			About:   "https://quadcode.com/about",
			Blog:    "https://quadcode.com/blog",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                42345997,
				Alias:             "quadcodecareer",
				Name:              "Quadcode",
				Followers:         "38K",
				Employees:         "501-1K",
				AssociatedMembers: "422",
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
				Alias:     "quadcode",
				Employees: "450",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Quadcode-EI_IE3293771.11,19.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Quadcode-Reviews-E3293771.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Quadcode-Jobs-E3293771.htm",
				Jobs:        "60",
				Reviews:     "52",
				Salaries:    "107",
				ReviewsRate: "3.3",
				Verified:    true,
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
							Title:                "Senior Golang Developer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/3907613945/",
							Date:                 mustDate("2024-04-25"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Senior Golang Developer",
							ShortDescription:     "Trading Exchange",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4242631355/",
							Date:                 mustDate("2025-06-27", "2025-06-06"),
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
			ShortDescription: "Fintech company developing software for the trading and investment industry",
			Industries: []domain.Industry{
				domain.IndustryFinTech,
			},
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
			},
		},

		// Some | FinCompare
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "FinCompare",
			Website: "https://fincompare.de/",
			Careers: "https://fincompare.de/jobs",
			About:   "https://fincompare.de/ueber-fincompare",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                10976436,
				Alias:             "fincompare",
				Name:              "FinCompare",
				Followers:         "5K",
				Employees:         "51-200",
				AssociatedMembers: "64",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "fincompare",
				Verified: true,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-FinCompare-EI_IE1644778.11,21.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/FinCompare-Reviews-E1644778.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/FinCompare-Jobs-E1644778.htm",
				Jobs:        "4",
				Reviews:     "45",
				Salaries:    "51",
				ReviewsRate: "4.3",
				Verified:    true,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "",
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 1,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Senior Golang Engineer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/3846776554/",
							Date:                 mustDate("2024-04-25"),
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
			Industries:       []domain.Industry{},
		},

		// Some | Mellifera
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Mellifera",
			Website: "https://mellifera.team/",
			Careers: "https://mellifera.team/careers/",
			About:   "https://mellifera.team/about-us/",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                98533069,
				IDs:               nil,
				Alias:             "mellifera-operations-limited",
				Name:              "Mellifera",
				Followers:         "8K",
				Employees:         "51-200",
				AssociatedMembers: "88",
				Verified:          false,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "",
				Verified: false,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "",
				ReviewsURL:  "",
				JobsURL:     "",
				Jobs:        "",
				Reviews:     "",
				Salaries:    "",
				ReviewsRate: "",
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
							Title:                "Go (Golang) Developer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/3912533167/",
							Date:                 mustDate("2024-04-25"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Senior Go Engineer",
							ShortDescription:     "Payment Gateway Solutions",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4182021240/",
							Date:                 mustDate("2025-03-11"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Go Engineer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4184612851/",
							Date:                 mustDate("2025-03-14"),
							WithSalary:           true, // 4500-6000 EUR gross per month
							Remote:               false,
						},
						{
							Title:                "Senior Go Engineer",
							ShortDescription:     "Payment Gateway Solutions",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4218758725/",
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
			ShortDescription: "",
			Industries:       []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
			},
		},

		// Some | BNP Paribas — Securities Services
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "BNP Paribas — Securities Services",
			Website: "https://securities.cib.bnpparibas/",
			Careers: "https://securities.cib.bnpparibas/who-we-are/your-career-with-us/",
			About:   "https://securities.cib.bnpparibas/",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                3625182,
				Alias:             "bnpparibassecuritiesservices",
				Name:              "BNP Paribas - Securities Services",
				Followers:         "137K",
				Employees:         "5K-10K",
				AssociatedMembers: "6,620",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "",
				Verified: false,
			},
			BlindProfile: domain.BlindProfile{
				Alias:       "BNP-Paribas",
				Employees:   "10,000+",
				Salary:      "$45K ~ $236K a year",
				Reviews:     "43",
				ReviewsRate: "2.9",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "bnp-paribas-securities",
				Employees: "17,000",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-BNP-Paribas-EI_IE10342.11,22.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/BNP-Paribas-Reviews-E10342.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/BNP-Paribas-Jobs-E10342.htm",
				Jobs:        "3.1K",
				Reviews:     "18K",
				Salaries:    "28K",
				ReviewsRate: "3.8",
				Verified:    true,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "",
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "https://www.youtube.com/@labanquedunmondequichange",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Tech Lead Golang Developer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/3909846279/",
							Date:                 mustDate("2024-06-25"),
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
			Industries:       []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
			},
		},

		// Some | Apifonica
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Apifonica",
			Website: "https://www.apifonica.com/",
			Careers: "https://www.apifonica.com/en/industries/hr-and-recruitment/",
			About:   "https://www.apifonica.com/en/company/about/",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                10943475,
				Alias:             "apifonica",
				Name:              "Apifonica",
				Followers:         "15K",
				Employees:         "51-200",
				AssociatedMembers: "45",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "apifonica",
				Verified: false,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Apifonica-EI_IE1805118.11,20.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Apifonica-Reviews-E1805118.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Apifonica-Jobs-E1805118.htm",
				Jobs:        "",
				Reviews:     "8",
				Salaries:    "10",
				ReviewsRate: "4.3",
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
					GitHubRepositoriesCount: 4,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Lead Golang Developer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/3909698518/",
							Date:                 mustDate("2024-04-25"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Senior PHP/Go Developer",
							ShortDescription:     "3+ years of engineering and development experience (PHP/Golang)",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4122032404/",
							Date:                 mustDate("2025-01-10"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Senior PHP/Go Developer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4142622628/",
							Date:                 mustDate("2025-02-04"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Lead PHP/Go Developer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4142623623/",
							Date:                 mustDate("2025-02-04"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Senior PHP/Go Developer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4161464959/",
							Date:                 mustDate("2025-02-21"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Lead PHP/Go Developer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4161466631/",
							Date:                 mustDate("2025-02-21"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Senior Golang Developer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4245963867/",
							Date:                 mustDate("2025-06-07"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Lead Golang Developer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4255188444/",
							Date:                 mustDate("2025-06-24"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Senior/Lead Golang Developer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4262447536/",
							Date:                 mustDate("2025-07-04"),
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
			Industries:       []domain.Industry{},
		},

		// Some | Cybus
		{
			ID:      0,                         // system
			Type:    domain.CompanyTypeStartup, // system
			Name:    "Cybus",
			Website: "https://www.cybus.io/",
			Careers: "https://www.cybus.io/en/jobs/career/",
			About:   "https://www.cybus.io/en/company/about-cybus/",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                7798667,
				Alias:             "cybus",
				Name:              "Cybus",
				Followers:         "3K",
				Employees:         "51-200",
				AssociatedMembers: "62",
				Verified:          false,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "cybusio",
				Verified: false,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Cybus-EI_IE2928520.11,16.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Cybus-Reviews-E2928520.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Cybus-Jobs-E2928520.htm",
				Jobs:        "3",
				Reviews:     "5",
				Salaries:    "11",
				ReviewsRate: "4.8",
				Verified:    false,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "",
			},
			OttaProfileSlug:   "Cybus",
			YouTubeChannelURL: "https://www.youtube.com/@cybus_io",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 1,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Senior Backend (Node.js) Engineer",
							ShortDescription:     "Work with our application core technologies, which is Node.js and Go and modern broker technologies",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4054430926/",
							Date:                 mustDate("2024-10-25"),
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
			ShortDescription:          "Data architecture for the industrial IoT world",
			Industries:                []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{},
		},

		// Some | Flink
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Flink",
			Website: "https://www.goflink.com/",
			Careers: "https://careers.smartrecruiters.com/Flink3/joinus",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                71241902,
				Alias:             "goflink",
				Name:              "Flink",
				Followers:         "52K",
				Employees:         "5K-10K",
				AssociatedMembers: "2,478",
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
				Alias:     "flink",
				Employees: "12,000",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Flink-EI_IE4921496.11,16.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Flink-Reviews-E4921496.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Flink-Jobs-E4921496.htm",
				Jobs:        "370",
				Reviews:     "724",
				Salaries:    "1.2K",
				ReviewsRate: "2.9",
				Verified:    false,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "",
			},
			OttaProfileSlug:   "Flink",
			YouTubeChannelURL: "https://www.youtube.com/@flink7309",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Software Engineer",
							ShortDescription:     "2 years of commercial experience with Golang",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/3906306472/",
							Date:                 mustDate("2024-04-25"),
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
			Industries:       []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{
				domain.Czechia,
			},
		},

		// Some | Greenbone
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Greenbone",
			Website: "https://www.greenbone.net/",
			Careers: "https://www.greenbone.net/careers/",
			About:   "https://www.greenbone.net/about-greenbone/",
			Blog:    "https://community.greenbone.net/blog/category/blog/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                697428,
				Alias:             "greenbone-ag",
				PreviousAliases:   []string{"greenbone-networks-gmbh"},
				Name:              "Greenbone AG",
				Followers:         "7K",
				Employees:         "51-200",
				AssociatedMembers: "104",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:     "greenbone",
				Followers: "1k",
				Verified:  true,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Greenbone-Networks-EI_IE379229.11,29.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Greenbone-Networks-Reviews-E379229.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Greenbone-Networks-Jobs-E379229.htm",
				Jobs:        "",
				Reviews:     "8",
				Salaries:    "12",
				ReviewsRate: "4.1",
				Verified:    false,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "",
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "https://www.youtube.com/@greenbone",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 9,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Golang Software Engineer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/3892843872/",
							Date:                 mustDate("2024-04-25"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Senior Golang Software Engineer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4227038724/",
							Date:                 mustDate("2025-05-10"),
							WithSalary:           false,
							Remote:               true,
						},
						{
							Title:                "Senior Golang Software Engineer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4249613510/",
							Date:                 mustDate("2025-06-13"),
							WithSalary:           false,
							Remote:               true,
						},
						{
							Title:                "Senior Golang Software Engineer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4264589613/",
							Location:             "Germany",
							Date:                 mustDate("2025-07-16", "2025-07-09"),
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
			Industries:       []domain.Industry{},
		},

		// Some | OLX
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "OLX",
			Website: "https://www.olxgroup.com/",
			Careers: "https://careers.olxgroup.com/",
			About:   "https://www.olxgroup.com/about-us/",
			Blog:    "https://tech.olx.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                167557,
				Alias:             "olx-group",
				Name:              "OLX",
				Followers:         "213K",
				Employees:         "1K-5K",
				AssociatedMembers: "4,306",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "",
				Verified: false,
			},
			BlindProfile: domain.BlindProfile{
				Alias:       "OLX",
				Employees:   "1,001 to 5,000",
				Salary:      "",
				Reviews:     "11",
				ReviewsRate: "3.7",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "olx-group",
				Employees: "87",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-OLX-Group-EI_IE517166.11,20.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/OLX-Group-Reviews-E517166.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/OLX-Group-Jobs-E517166.htm",
				Jobs:        "",
				Reviews:     "1.2K",
				Salaries:    "2K",
				ReviewsRate: "4.0",
				Verified:    true,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "Olx-Group",
			},
			OttaProfileSlug:   "OLX-Group",
			YouTubeChannelURL: "https://www.youtube.com/@OLXGroup",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Middle/Senior Software Engineer (Golang)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4055838812/",
							Date:                 mustDate("2024-12-22"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Senior Software Engineer (Golang)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4133389022/",
							Date:                 mustDate("2025-02-15", "2025-01-24"),
							WithSalary:           false,
							Remote:               true,
						},
						{
							Title:                "Senior Backend Engineer (Golang)",
							ShortDescription:     "Real Estate",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4142492503/",
							Date:                 mustDate("2025-03-19", "2025-02-25", "2025-02-03"),
							WithSalary:           false,
							Remote:               true,
						},
						{
							Title:                "Senior Software Engineer (Golang, Java, PHP)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4137532355/",
							Date:                 mustDate("2025-02-19"),
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
			ShortDescription: "Marketplace ecosystems",
			Industries: []domain.Industry{
				domain.IndustryECommerce,
			},
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
				domain.Czechia,
			},
		},

		// Some | Snyk
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Snyk",
			Website: "https://snyk.io/",
			Careers: "https://snyk.io/careers/",
			About:   "https://snyk.io/about/",
			Blog:    "https://snyk.io/blog/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                10043614,
				Alias:             "snyk",
				Name:              "Snyk",
				Followers:         "92K",
				Employees:         "1K-5K",
				AssociatedMembers: "1,290",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "snyk",
				Verified: true,
			},
			BlindProfile: domain.BlindProfile{
				Alias:       "Snyk",
				Employees:   "1,001 to 5,000 Employees",
				Salary:      "",
				Reviews:     "83",
				ReviewsRate: "3.6",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "snyk",
				Employees: "930",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Snyk-EI_IE2094989.11,15.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Snyk-Reviews-E2094989.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Snyk-Jobs-E2094989.htm",
				Jobs:        "40",
				Reviews:     "299",
				Salaries:    "655",
				ReviewsRate: "3.9",
				Verified:    true,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "Snyk",
			},
			OttaProfileSlug:   "Snyk",
			YouTubeChannelURL: "https://www.youtube.com/@Snyksec",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 36,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Software Engineer (Go)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4006143681/",
							Date:                 mustDate("2025-06-09", "2025-04-27", "2025-04-05", "2025-01-08", "2024-11-14"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Senior Software Engineer — Policies Team (TypeScript, Go)",
							ShortDescription:     "7+ years experience in software engineering",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4024339260/",
							Date:                 mustDate("2025-02-28", "2025-01-14"),
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
			Industries:                []domain.Industry{},
			HasEmployeesFromCountries: nil,
		},

		// Some | Sinch
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Sinch",
			Website: "https://www.sinch.com/",
			Careers: "https://sinch.com/careers/",
			About:   "https://sinch.com/about-us/",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                3726743,
				Alias:             "sinch",
				Name:              "Sinch",
				Followers:         "398K",
				Employees:         "1K-5K",
				AssociatedMembers: "4,169",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "sinch",
				Verified: false,
			},
			BlindProfile: domain.BlindProfile{
				Alias:       "Sinch",
				Employees:   "1,001 to 5,000",
				Salary:      "",
				Reviews:     "9",
				ReviewsRate: "3.0",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "sinch",
				Employees: "1,500",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Sinch-EI_IE788805.11,16.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Sinch-Reviews-E788805.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Sinch-Jobs-E788805.htm",
				Jobs:        "118",
				Reviews:     "547",
				Salaries:    "1K",
				ReviewsRate: "3.7",
				Verified:    true,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "Sinch",
			},
			OttaProfileSlug:   "sinch",
			YouTubeChannelURL: "https://www.youtube.com/@WeAreSinch",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Senior Software Engineer — Golang",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/3903797937/",
							Date:                 mustDate("2024-04-25"),
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
			Industries:       []domain.Industry{},
		},

		// Some | FOX Tech
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "FOX Tech",
			Website: "https://tech.fox.com/",
			Careers: "https://www.foxcareers.com/",
			About:   "https://www.foxcorporation.com/about-us/",
			Blog:    "https://medium.com/fox-tech",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                14850572,
				Alias:             "foxtechteam",
				Name:              "FOX Tech",
				Followers:         "11K",
				Employees:         "5K-10K",
				AssociatedMembers: "111",
				Verified:          false,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "",
				Verified: false,
			},
			BlindProfile: domain.BlindProfile{
				Alias:       "Fox",
				Employees:   "10,000+",
				Salary:      "$66K ~ $318K a year",
				Reviews:     "67",
				ReviewsRate: "3.6",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "fox",
				Employees: "120",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-FOX-Broadcasting-EI_IE13272.11,27.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/FOX-Broadcasting-Reviews-E13272.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/FOX-Broadcasting-Jobs-E13272.htm",
				Jobs:        "173",
				Reviews:     "313",
				Salaries:    "552",
				ReviewsRate: "3.8",
				Verified:    false,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "Fox-Corporation",
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Senior Software Engineer, Golang",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/3903338524/",
							Date:                 mustDate("2024-08-26"),
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
			Industries:       []domain.Industry{},
		},

		// Some | SailPoint
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "SailPoint",
			Website: "https://www.sailpoint.com/",
			Careers: "https://www.sailpoint.com/company/careers",
			About:   "https://www.sailpoint.com/why-us/about-us",
			Blog:    "https://developer.sailpoint.com/blog",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                47456,
				Alias:             "sailpoint-technologies",
				Name:              "SailPoint",
				Followers:         "125K",
				Employees:         "1K-5K",
				AssociatedMembers: "2,830",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:     "sailpoint-oss",
				Followers: "273",
				Verified:  true,
			},
			BlindProfile: domain.BlindProfile{
				Alias:       "SailPoint-Technologies",
				Employees:   "1,001 to 5,000",
				Salary:      "$60K ~ $250K a year",
				Reviews:     "98",
				ReviewsRate: "3.8",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "sailpoint",
				Employees: "2,000",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-SailPoint-Technologies-EI_IE449696.11,33.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/SailPoint-Technologies-Reviews-E449696.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/SailPoint-Technologies-Jobs-E449696.htm",
				Jobs:        "94",
				Reviews:     "939",
				Salaries:    "1.3K",
				ReviewsRate: "4.6",
				Verified:    true,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "Sailpoint-Technologies-2",
			},
			OttaProfileSlug:   "Sailpoint",
			YouTubeChannelURL: "https://www.youtube.com/@SailPointTechnologies",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 3,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Senior Golang Engineer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4030943277/",
							Date:                 mustDate("2024-10-25"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Senior Golang Software Engineer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4107495969/",
							Date:                 mustDate("2024-12-22"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Senior Golang Software Engineer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4077458993/",
							Date:                 mustDate("2025-01-23"),
							WithSalary:           false,
							Remote:               true,
						},
						{
							Title:                "Senior Golang Software Engineer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4149693865/",
							Date:                 mustDate("2025-06-27", "2025-06-05", "2025-05-15", "2025-04-24", "2025-03-27", "2025-03-06", "2025-02-11"),
							WithSalary:           true, // $110,250 - $157,500 - $204,750 (min-mid-max)
							Remote:               true,
						},
						{
							Title:                "Senior Golang Engineer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4148940672/",
							Date:                 mustDate("2025-03-06", "2025-02-13"),
							WithSalary:           false,
							Remote:               true,
						},
						{
							Title:                "Software Engineer (Java, Golang)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4151653458/",
							Date:                 mustDate("2025-02-14"),
							WithSalary:           true, // salary range $74,060 - $105,800 - $137,540
							Remote:               true,
						},
						{
							Title:                "Senior Software Engineer — Golang",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4137755120/",
							Date:                 mustDate("2025-02-19"),
							WithSalary:           false,
							Remote:               true,
						},
						{
							Title:                "Golang Software Engineer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4158963804/",
							Date:                 mustDate("2025-02-19"),
							WithSalary:           true, // salary range $69,860 - $99,800 - $129,740
							Remote:               true,
						},
						{
							Title:                "Senior Golang Software Engineer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4213006840/",
							Location:             "Guadalajara, Jalisco, Mexico",
							Date:                 mustDate("2025-07-15", "2025-06-24", "2025-05-10", "2025-04-19"),
							WithSalary:           false,
							Remote:               true,
						},
						{
							Title:                "Senior Golang Engineer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4219688456/",
							Date:                 mustDate("2025-06-24"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Senior Golang Engineer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4260870925/",
							Date:                 mustDate("2025-07-02"),
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
			Industries:       []domain.Industry{},
		},

		// Some | Proofpoint
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Proofpoint",
			Website: "https://www.proofpoint.com/",
			Careers: "https://www.proofpoint.com/us/company/careers",
			About:   "https://www.proofpoint.com/us/company/about",
			Blog:    "https://www.proofpoint.com/us/blog/engineering-insights",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                11681,
				Alias:             "proofpoint",
				Name:              "Proofpoint",
				Followers:         "154K",
				Employees:         "1K-5K",
				AssociatedMembers: "4,698",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "proofpoint",
				Verified: false,
			},
			BlindProfile: domain.BlindProfile{
				Alias:       "Proofpoint",
				Employees:   "1,001 to 5,000",
				Salary:      "$62K ~ $324K a year",
				Reviews:     "162",
				ReviewsRate: "3.5",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "proofpoint",
				Employees: "3,750",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Proofpoint-EI_IE39140.11,21.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Proofpoint-Reviews-E39140.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Proofpoint-Jobs-E39140.htm",
				Jobs:        "140",
				Reviews:     "1.1K",
				Salaries:    "2.1K",
				ReviewsRate: "3.6",
				Verified:    true,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "",
			},
			OttaProfileSlug:   "Proofpoint",
			YouTubeChannelURL: "https://www.youtube.com/@proofpoint",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 13,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Senior Software Engineer (Golang)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/3797062001/",
							Date:                 mustDate("2024-06-25"),
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
			Industries:       []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
				domain.Czechia,
			},
		},

		// Some | Asset Reality
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Asset Reality",
			Website: "https://www.assetreality.com/",
			Careers: "",
			About:   "https://www.assetreality.com/about",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                42805677,
				Alias:             "asset-reality",
				Name:              "Asset Reality",
				Followers:         "3K",
				Employees:         "11-50",
				AssociatedMembers: "50",
				Verified:          false,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "",
				Verified: false,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "",
				ReviewsURL:  "",
				JobsURL:     "",
				Jobs:        "",
				Reviews:     "",
				Salaries:    "",
				ReviewsRate: "",
				Verified:    false,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "",
			},
			OttaProfileSlug:   "Asset-Reality",
			YouTubeChannelURL: "https://www.youtube.com/@assetreality",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						/*
							{
								Title:                "",
								ShortDescription:     "",
								SwitchingOpportunity: "",
								URL:                  "",
								Date:                 mustDate(""),
								WithSalary:           false,
								Remote:               false,
							},
						*/
						{
							Title:                "Senior Backend Engineer — Go",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4165243684/",
							Date:                 mustDate("2025-02-25"),
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
			Industries:       []domain.Industry{},
		},

		// Some | Limango
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Limango",
			Website: "https://www.limango.pl/",
			Careers: "https://praca.limango.pl/oferty-pracy/oferta-product-owner/",
			About:   "https://www.limango.pl/info/limango",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                2966982,
				Alias:             "limango-sp-z-o-o-",
				Name:              "Limango",
				Followers:         "2K",
				Employees:         "501-1K",
				AssociatedMembers: "179",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "",
				Verified: false,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Limango-Polska-EI_IE2884426.11,25.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Limango-Polska-Reviews-E2884426.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Limango-Polska-Jobs-E2884426.htm",
				Jobs:        "",
				Reviews:     "8",
				Salaries:    "28",
				ReviewsRate: "3.5",
				Verified:    false,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "",
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "https://www.youtube.com/@limangoPolska",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{
						/*
							{
								Title:            "",
								ShortDescription: "",
								URL:              "https://justjoin.it/offers/limango-polska-mid-senior-backend-developer-golang",
								Date:             mustDate(""),
							},
						*/
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
		},

		// Some | RxBenefits
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "RxBenefits",
			Website: "https://www.rxbenefits.com/",
			Careers: "https://www.rxbenefits.com/about-us/culture/",
			About:   "https://www.rxbenefits.com/about-us/about-rxbenefits/",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                2873210,
				Alias:             "rxbenefits-inc-",
				Name:              "RxBenefits, Inc.",
				Followers:         "38K",
				Employees:         "1K-5K",
				AssociatedMembers: "879",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "",
				Verified: false,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-RxBenefits-EI_IE1175839.11,21.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/RxBenefits-Reviews-E1175839.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/RxBenefits-Jobs-E1175839.htm",
				Jobs:        "103",
				Reviews:     "147",
				Salaries:    "247",
				ReviewsRate: "3.6",
				Verified:    true,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "Rxbenefits,-Inc.-1",
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Senior Golang Software Engineer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/3910221910/",
							Date:                 mustDate("2024-04-25"),
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
			Industries:       []domain.Industry{},
		},

		// Some | SmithRx
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "SmithRx",
			Website: "https://www.smithrx.com/",
			Careers: "https://www.smithrx.com/careers",
			About:   "https://www.smithrx.com/our-story",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                10886362,
				Alias:             "smithrx",
				Name:              "SmithRx",
				Followers:         "11K",
				Employees:         "501-1K",
				AssociatedMembers: "409",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "",
				Verified: false,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-SmithRx-EI_IE1901555.11,18.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/SmithRx-Reviews-E1901555.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/SmithRx-Jobs-E1901555.htm",
				Jobs:        "11",
				Reviews:     "66",
				Salaries:    "112",
				ReviewsRate: "3.6",
				Verified:    true,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "",
			},
			OttaProfileSlug:   "SmithRx",
			YouTubeChannelURL: "https://www.youtube.com/@SmithRxPBM",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{
						/*
							{
								Title:                "",
								ShortDescription:     "",
								SwitchingOpportunity: "",
								URL:                  "",
								Date:                 mustDate(""),
								WithSalary:           false,
								Remote:               false,
							},
						*/
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
		},

		// Some | k-ID
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "k-ID",
			Website: "https://www.k-id.com/",
			Careers: "https://www.k-id.com/careers",
			About:   "https://www.k-id.com/about",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                45973117,
				Alias:             "k-id",
				Name:              "k-ID",
				Followers:         "5K",
				Employees:         "11-50",
				AssociatedMembers: "41",
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
				Alias:     "k-id",
				Employees: "42",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "",
				ReviewsURL:  "",
				JobsURL:     "",
				Jobs:        "",
				Reviews:     "",
				Salaries:    "",
				ReviewsRate: "",
				Verified:    false,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "",
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "https://www.youtube.com/@k-IDofficial",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Senior Golang Engineer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/3912660665/",
							Date:                 mustDate("2024-04-25"),
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
			Industries:       []domain.Industry{},
		},

		// Some | CAFU
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "CAFU",
			Website: "https://www.cafu.com/",
			Careers: "https://www.cafu.com/careers",
			About:   "https://www.cafu.com/about",
			Blog:    "https://cafu.engineering/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                13892019,
				Alias:             "mycafu",
				Name:              "CAFU",
				Followers:         "293K",
				Employees:         "51-200",
				AssociatedMembers: "1,090",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "",
				Verified: false,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-CAFU-EI_IE3713615.11,15.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/CAFU-Reviews-E3713615.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/CAFU-Jobs-E3713615.htm",
				Jobs:        "",
				Reviews:     "93",
				Salaries:    "118",
				ReviewsRate: "3.0",
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
					Vacancies:               []domain.Vacancy{
						/*
							{
								Title:                "",
								ShortDescription:     "",
								SwitchingOpportunity: "",
								URL:                  "",
								Date:                 mustDate(""),
								WithSalary:           false,
								Remote:               false,
							},
						*/
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
		},

		// Some | Rollee
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Rollee",
			Website: "https://www.getrollee.com/",
			Careers: "https://www.getrollee.com/company/careers",
			About:   "https://www.getrollee.com/company/mission",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                76353840,
				Alias:             "rollee",
				Name:              "Rollee",
				Followers:         "3K",
				Employees:         "11-50",
				AssociatedMembers: "26",
				Verified:          false,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "",
				Verified: false,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "",
				ReviewsURL:  "",
				JobsURL:     "",
				Jobs:        "",
				Reviews:     "",
				Salaries:    "",
				ReviewsRate: "",
				Verified:    false,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "",
			},
			OttaProfileSlug:   "Rollee",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{
						/*
							{
								Title:            "",
								ShortDescription: "",
								URL:              "https://app.welcometothejungle.com/jobs/4ajPfsAU",
								Date:             mustDate(""),
							},
						*/
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
		},

		// Some | Net2Phone
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Net2Phone",
			Website: "https://www.net2phone.com/",
			Careers: "https://www.net2phone.com/careers",
			About:   "https://www.net2phone.com/about",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                2965,
				Alias:             "net2phone",
				Name:              "Net2Phone",
				Followers:         "10K",
				Employees:         "1K-5K",
				AssociatedMembers: "259",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "net2phone",
				Verified: true,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Net2Phone-EI_IE9342.11,20.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Net2Phone-Reviews-E9342.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Net2Phone-Jobs-E9342.htm",
				Jobs:        "5",
				Reviews:     "62",
				Salaries:    "104",
				ReviewsRate: "3.6",
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
							Title:                "Senior Golang Developer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/3914839595/",
							Date:                 mustDate("2024-04-25"),
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
			Industries:       []domain.Industry{},
		},

		// Some | Ola Chat
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Ola Chat",
			Website: "https://olachat.sg/",
			Careers: "",
			About:   "https://olachat.sg/about",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                68332088,
				Alias:             "ola-chat",
				Name:              "Ola Chat",
				Followers:         "47K",
				Employees:         "1K-5K",
				AssociatedMembers: "243",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "olachat",
				Verified: false,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "", // NOP
				ReviewsURL:  "", // NOP
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "",
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 8,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Senior Golang Engineer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/3912130901/",
							Date:                 mustDate("2024-05-24"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Back-End Developer — Golang",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4105518392/",
							Date:                 mustDate("2024-12-19"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Software Engineer (Go)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4132246194/",
							Date:                 mustDate("2025-01-22"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Software Engineer (Go)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4181042382/",
							Date:                 mustDate("2025-03-10"),
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
			Industries:       []domain.Industry{},
		},

		// Some | Veracity Software Inc
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Veracity Software Inc",
			Website: "https://veracity-us.com/",
			Careers: "https://veracity-us.com//career",
			About:   "https://veracity-us.com//about",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                11137552,
				Alias:             "veracitysoftwareinc",
				Name:              "Veracity Software Inc",
				Followers:         "46K",
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
				Alias:     "veracity-software",
				Employees: "30",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Veracity-Software-EI_IE1357198.11,28.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Veracity-Software-Reviews-E1357198.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Veracity-Software-Jobs-E1357198.htm",
				Jobs:        "11",
				Reviews:     "78",
				Salaries:    "115",
				ReviewsRate: "4.1",
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
							Title:                "Full Stack Java / Golang Developer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/3818383367/",
							Date:                 mustDate("2024-02-26"),
							WithSalary:           false,
							Remote:               false,
						},
					},
				},
				domain.Rust: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Rust, AI Platform Engineer, Director",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4252441556/",
							Date:                 mustDate("2025-06-22"),
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
			Industries:       []domain.Industry{},
		},

		// Some | Treecard
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Treecard",
			Website: "https://www.treecard.org/",
			Careers: "https://www.treecard.org/careers",
			About:   "https://www.treecard.org/the-mission",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                68821773,
				Alias:             "treecardapp",
				Name:              "Treecard",
				Followers:         "16K",
				Employees:         "11-50",
				AssociatedMembers: "19",
				Verified:          false,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "TreeCard",
				Verified: false,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Treecard-EI_IE5675051.11,19.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Treecard-Reviews-E5675051.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Treecard-Jobs-E5675051.htm",
				Jobs:        "",
				Reviews:     "14",
				Salaries:    "15",
				ReviewsRate: "4.7",
				Verified:    false,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "",
			},
			OttaProfileSlug:   "Treecard",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 2,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Senior Product Engineer (Backend — Golang)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/3909975571/",
							Date:                 mustDate("2024-04-25"),
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
			Industries:       []domain.Industry{},
		},

		// Some | Openprovider
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Openprovider",
			Website: "https://www.openprovider.com/",
			Careers: "https://www.openprovider.com/company/careers",
			About:   "https://www.openprovider.com/company/about-us",
			Blog:    "https://www.openprovider.com/blog",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                550698,
				Alias:             "openprovider",
				Name:              "Openprovider",
				Followers:         "6K",
				Employees:         "51-200",
				AssociatedMembers: "73",
				Verified:          false,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "openprovider",
				Verified: true,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Openprovider-EI_IE1267124.11,23.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Openprovider-Reviews-E1267124.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Openprovider-Jobs-E1267124.htm",
				Jobs:        "17",
				Reviews:     "30",
				Salaries:    "28",
				ReviewsRate: "4.7",
				Verified:    true,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "",
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "https://www.youtube.com/@openprovider453",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 10,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Backend Software Engineer (Golang)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/3914419419/",
							Date:                 mustDate("2024-04-25"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Full Stack Developer (ReactJS & Golang)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4166998360/",
							Date:                 mustDate("2025-02-24"),
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
			Industries:       []domain.Industry{},
		},

		// Some | fiskaly
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "fiskaly",
			Website: "https://www.fiskaly.com/",
			Careers: "https://www.fiskaly.com/jobs",
			About:   "https://www.fiskaly.com/about",
			Blog:    "https://developer.fiskaly.com/blog",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                18929063,
				Alias:             "fiskaly",
				Name:              "fiskaly",
				Followers:         "8K",
				Employees:         "51-200",
				AssociatedMembers: "86",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:     "fiskaly",
				Followers: "36",
				Verified:  true,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-fiskaly-EI_IE3059515.11,18.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/fiskaly-Reviews-E3059515.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/fiskaly-Jobs-E3059515.htm",
				Jobs:        "8",
				Reviews:     "4",
				Salaries:    "25",
				ReviewsRate: "4.3",
				Verified:    false,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "",
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "https://www.youtube.com/@fiskaly",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 5,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Senior Backend Developer Golang",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/3914242491/", // 5 years working with Golang or other OOP languages
							Date:                 mustDate("2024-05-24"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Senior Backend Developer Golang",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4055768396/", // 5 years working with Golang or other OOP languages
							Date:                 mustDate("2024-12-19"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Senior Backend Developer Golang",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4157742979/",
							Date:                 mustDate("2025-02-18"),
							WithSalary:           true, // €56.000 - €70.000 per year
							Remote:               false,
						},
						{
							Title:                "Senior Backend Engineer (Golang)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4260892737/",
							Date:                 mustDate("2025-07-02"),
							WithSalary:           false,
							Remote:               true,
						},
						{
							Title:                "Senior Backend Developer Golang",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4261154544/",
							Location:             "Vienna, Vienna, Austria",
							Date:                 mustDate("2025-07-09"),
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
			ShortDescription: "fiskaly builds innovative tools to make receipts simple",
			Industries: []domain.Industry{
				domain.IndustryFinTech,
			},
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
			},
		},

		// Some | Vay
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Vay",
			Website: "https://vay.io/",
			Careers: "https://vay.io/careers/",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                12584218,
				Alias:             "vaytechnology",
				Name:              "Vay",
				Followers:         "15K",
				Employees:         "51-200",
				AssociatedMembers: "162",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "Reemote",
				Verified: false,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "",
				ReviewsURL:  "",
				Verified:    false,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "",
			},
			OttaProfileSlug:   "Vay",
			YouTubeChannelURL: "https://www.youtube.com/@vay_io",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 2,
					Vacancies:               []domain.Vacancy{
						/*
							{
								Title:                "",
								ShortDescription:     "",
								SwitchingOpportunity: "",
								URL:                  "",
								Date:                 mustDate(""),
								WithSalary:           false,
								Remote:               false,
							},
						*/
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
		},

		// Some | Voltus
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Voltus",
			Website: "https://www.voltus.co/",
			Careers: "https://www.voltus.co/careers",
			About:   "https://www.voltus.co/about",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                10866628,
				Alias:             "voltus-inc.",
				Name:              "Voltus",
				Followers:         "14K",
				Employees:         "201-500",
				AssociatedMembers: "256",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "voltusdev",
				Verified: false,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Voltus-EI_IE2090197.11,17.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Voltus-Reviews-E2090197.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Voltus-Jobs-E2090197.htm",
				Jobs:        "",
				Reviews:     "93",
				Salaries:    "178",
				ReviewsRate: "3.7",
				Verified:    false,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "",
			},
			OttaProfileSlug:   "Voltus",
			YouTubeChannelURL: "https://www.youtube.com/@Voltusinc",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 1,
					Vacancies:               []domain.Vacancy{
						/*
							{
								Title:            "",
								ShortDescription: "",
								URL:              "https://app.welcometothejungle.com/jobs/pB_hGf0W",
								Date:             mustDate(""),
							},
						*/
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
		},

		// Some | Stonebranch
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Stonebranch",
			Website: "https://www.stonebranch.com/",
			Careers: "https://www.stonebranch.com/careers",
			About:   "https://www.stonebranch.com/about-us",
			Blog:    "https://www.stonebranch.com/blog",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                71261,
				Alias:             "stonebranch",
				Name:              "Stonebranch",
				Followers:         "8K",
				Employees:         "51-200",
				AssociatedMembers: "161",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "stonebranch-marketplace",
				Verified: false,
			},
			BlindProfile: domain.BlindProfile{
				Alias: "",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "stonebranch",
				Employees: "150",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Stonebranch-EI_IE408996.11,22.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Stonebranch-Reviews-E408996.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Stonebranch-Jobs-E408996.htm",
				Jobs:        "17",
				Reviews:     "30",
				Salaries:    "46",
				ReviewsRate: "3.9",
				Verified:    true,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "Stonebranch",
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "https://www.youtube.com/@Stonebranch",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Go Developer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/3914103960/",
							Date:                 mustDate("2024-05-24"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Go Developer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4106867221/",
							Date:                 mustDate("2024-12-24"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Go Developer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4201851657/",
							Date:                 mustDate("2025-04-08"),
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
			Industries:       []domain.Industry{},
		},

		// Some | Rapid7
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Rapid7",
			Website: "https://www.rapid7.com/",
			Careers: "https://careers.rapid7.com/",
			About:   "https://www.rapid7.com/about/",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                39624,
				Alias:             "rapid7",
				Name:              "Rapid7",
				Followers:         "183K",
				Employees:         "1K-5K",
				AssociatedMembers: "3,070",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "rapid7",
				Verified: true,
			},
			BlindProfile: domain.BlindProfile{
				Alias:       "Rapid8",
				Employees:   "1,001 to 5,000",
				Salary:      "$75K ~ $350K a year",
				Reviews:     "69",
				ReviewsRate: "3.8",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "rapid7",
				Employees: "2,500",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Rapid7-EI_IE243542.11,17.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Rapid7-Reviews-E243542.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Rapid7-Jobs-E243542.htm",
				Jobs:        "193",
				Reviews:     "893",
				Salaries:    "2.1K",
				ReviewsRate: "3.7",
				Verified:    true,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "Rapid7",
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "https://www.youtube.com/@GoRapid7",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 13,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Senior Software Engineer — Golang",
							ShortDescription:     "Open-source platform — Velociraptor",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/3915341496/",
							Date:                 mustDate("2024-09-25"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Senior Software Engineer — Golang",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4159762093/",
							Date:                 mustDate("2025-04-05", "2025-03-14", "2025-02-20"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Senior Software Engineer – Golang",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4181500800/",
							Date:                 mustDate("2025-04-27"),
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
			Industries:       []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
				domain.Czechia,
			},
		},

		// Some | Toggle AI
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Toggle AI",
			Website: "https://toggle.ai/",
			Careers: "https://toggle.ai/careers",
			About:   "https://toggle.ai/#aboutUs",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                28508827,
				Alias:             "toggle-ai",
				Name:              "Toggle AI",
				Followers:         "4K",
				Employees:         "11-50",
				AssociatedMembers: "47",
				Verified:          false,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "",
				Verified: false,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Toggle-EI_IE3924898.11,17.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Toggle-Reviews-E3924898.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Toggle-Jobs-E3924898.htm",
				Jobs:        "",
				Reviews:     "14",
				Salaries:    "11",
				ReviewsRate: "4.8",
				Verified:    false,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "",
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "https://www.youtube.com/@ToggleAI-Investing",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Backend Software Engineer — Golang",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/3915194291/",
							Date:                 mustDate("2024-05-24"),
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
			Industries:       []domain.Industry{},
		},

		// Some | hearX Group
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "hearX Group",
			Website: "https://www.hearxgroup.com/",
			Careers: "https://hearxgroup.simplify.hr/",
			About:   "https://hearxgroup.com/about-us",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                18002825,
				Alias:             "hearx-group",
				Name:              "hearX Group",
				Followers:         "11K",
				Employees:         "51-200",
				AssociatedMembers: "160",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "hearxgroup",
				Verified: false,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-hearX-Group-EI_IE5800566.11,22.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/hearX-Group-Reviews-E5800566.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/hearX-Group-Jobs-E5800566.htm",
				Jobs:        "",
				Reviews:     "4",
				Salaries:    "11",
				ReviewsRate: "4.5",
				Verified:    false,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "",
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "https://www.youtube.com/@hearxgroupptyltd8061",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{
						/*
							{
								Title:                "",
								ShortDescription:     "",
								SwitchingOpportunity: "",
								URL:                  "",
								Date:                 mustDate(""),
								WithSalary:           false,
								Remote:               false,
							},
						*/
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
		},

		// Some | MarkiTech.AI
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "MarkiTech.AI",
			Website: "https://markitech.ca/",
			Careers: "https://markitech.ca/jobs/",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                9294422,
				Alias:             "markitech-ai",
				Name:              "MarkiTech.AI",
				Followers:         "3K",
				Employees:         "11-50",
				AssociatedMembers: "32",
				Verified:          false,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "",
				Verified: false,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Markitech-EI_IE4190913.11,20.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Markitech-Reviews-E4190913.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Markitech-Jobs-E4190913.htm",
				Jobs:        "1",
				Reviews:     "33",
				Salaries:    "33",
				ReviewsRate: "4.1",
				Verified:    false,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "",
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "https://www.youtube.com/@markitech-digitaltransform6173",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{
						/*
							{
								Title:                "",
								ShortDescription:     "",
								SwitchingOpportunity: "",
								URL:                  "",
								Date:                 mustDate(""),
								WithSalary:           false,
								Remote:               false,
							},
						*/
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
		},

		// Some | Lantronix
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Lantronix",
			Website: "https://www.lantronix.com/",
			Careers: "https://www.lantronix.com/about-us/careers/",
			About:   "https://www.lantronix.com/about-us/",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                12612,
				Alias:             "lantronix",
				Name:              "Lantronix",
				Followers:         "10K",
				Employees:         "201-500",
				AssociatedMembers: "443",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "Lantronix",
				Verified: false,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Lantronix-EI_IE5498.11,20.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Lantronix-Reviews-E5498.htm",
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "",
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "https://www.youtube.com/@LantronixInc",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{
						/*
							{
								Title:                "",
								ShortDescription:     "",
								SwitchingOpportunity: "",
								URL:                  "",
								Date:                 mustDate(""),
								WithSalary:           false,
								Remote:               false,
							},
						*/
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
		},

		// Some | INFOLOB
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "INFOLOB",
			Website: "https://www.infolob.com/",
			Careers: "https://www.infolob.com/careers/",
			About:   "https://www.infolob.com/about/",
			Blog:    "https://www.infolob.com/blog/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                404211,
				Alias:             "infolob-global",
				Name:              "INFOLOB",
				Followers:         "48K",
				Employees:         "201-500",
				AssociatedMembers: "468",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "",
				Verified: false,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Infolob-Solutions-EI_IE423113.11,28.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Infolob-Solutions-Reviews-E423113.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Infolob-Solutions-Jobs-E423113.htm",
				Jobs:        "6",
				Reviews:     "120",
				Salaries:    "524",
				ReviewsRate: "4.3",
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
							Title:                "Senior Python AWS Developer / Golang Developer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/3914415181/",
							Date:                 mustDate("2024-05-24"),
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
			Industries:       []domain.Industry{},
		},

		// Some | Argela Technologies
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Argela Technologies",
			Website: "https://www.argela.com.tr/",
			Careers: "https://www.argela.com.tr/en/career",
			About:   "https://www.argela.com.tr/en/about-us",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                26805,
				Alias:             "argela-technologies",
				Name:              "Argela Technologies",
				Followers:         "21K",
				Employees:         "51-200",
				AssociatedMembers: "190",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "",
				Verified: false,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Argela-EI_IE389528.11,17.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Argela-Reviews-E389528.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Argela-Jobs-E389528.htm",
				Jobs:        "",
				Reviews:     "57",
				Salaries:    "57",
				ReviewsRate: "4.0",
				Verified:    false,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "",
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "https://www.youtube.com/@argelatechnologies",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Software Engineer — C/C++, Golang",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/3914290466/",
							Date:                 mustDate("2024-04-25"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Software Engineer — C/C++, Golang",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4151349412/",
							Date:                 mustDate("2025-02-13"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Software Engineer — C/C++, Golang",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4184638545/",
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
			Industries:       []domain.Industry{},
		},

		// Some | Top Doctors
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Top Doctors",
			Website: "https://topdoctors.com/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                3035481,
				Alias:             "top-doctors-europe",
				Name:              "Top Doctors",
				Followers:         "35K",
				Employees:         "201-500",
				AssociatedMembers: "452",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "",
				Verified: false,
			},
			BlindProfile: domain.BlindProfile{
				Alias: "",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Top-Doctors-EI_IE1712721.11,22.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Top-Doctors-Reviews-E1712721.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Top-Doctors-Jobs-E1712721.htm",
				Jobs:        "27",
				Reviews:     "95",
				Salaries:    "145",
				ReviewsRate: "2.7",
				Verified:    false,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "Top-Doctors",
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "https://www.youtube.com/@TopDoctorsUK",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Developer Backend Golang",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4077191404/",
							Date:                 mustDate("2024-11-19"),
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
			Industries:                []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{},
		},

		// Some | Recurly
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Recurly",
			Website: "https://recurly.com/",
			Careers: "https://recurly.com/careers/",
			About:   "https://recurly.com/company/",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                810383,
				Alias:             "recurly",
				Name:              "Recurly",
				Followers:         "17K",
				Employees:         "201-500",
				AssociatedMembers: "300",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "recurly",
				Verified: true,
			},
			BlindProfile: domain.BlindProfile{
				Alias: "",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "recurly",
				Employees: "270",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Recurly-EI_IE692611.11,18.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Recurly-Reviews-E692611.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Recurly-Jobs-E692611.htm",
				Jobs:        "25",
				Reviews:     "178",
				Salaries:    "319",
				ReviewsRate: "3.7",
				Verified:    true,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "Recurly",
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "https://www.youtube.com/@Recurly",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 4,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Senior Software Engineer, Features (Golang)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/3912991626/",
							Date:                 mustDate("2024-05-24"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Principal Golang Engineer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4238083331/",
							Date:                 mustDate("2025-06-06"),
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
			Industries:       []domain.Industry{},
		},

		// Some | Cynet Systems
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Cynet Systems",
			Website: "https://www.cynetsystems.com/",
			Careers: "https://www.cynetsystems.com/jobs/",
			About:   "https://www.cynetsystems.com/company/",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                5003556,
				Alias:             "cynet-systems",
				Name:              "Cynet Systems",
				Followers:         "72K",
				Employees:         "1K-5K",
				AssociatedMembers: "682",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "",
				Verified: false,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Cynet-Systems-EI_IE654628.11,24.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Cynet-Systems-Reviews-E654628.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Cynet-Systems-Jobs-E654628.htm",
				Jobs:        "5",
				Reviews:     "436",
				Salaries:    "343",
				ReviewsRate: "4.2",
				Verified:    true,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "Cynet-Systems",
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "https://www.youtube.com/@cynetsystems2026",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Senior Golang Developer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/3912977786/",
							Date:                 mustDate("2024-05-24"),
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
			Industries:       []domain.Industry{},
		},

		// Some | Odyssey Information Services
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Odyssey Information Services",
			Website: "https://www.odysseyis.com/",
			Careers: "",
			About:   "https://www.odysseyis.com/about",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                56445,
				Alias:             "odyssey-information-services",
				Name:              "Odyssey Information Services",
				Followers:         "117K",
				Employees:         "201-500",
				AssociatedMembers: "164",
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
				Alias:     "odyssey-information-services",
				Employees: "210",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Odyssey-Information-Services-EI_IE558201.11,39.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Odyssey-Information-Services-Reviews-E558201.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Odyssey-Information-Services-Jobs-E558201.htm",
				Jobs:        "18",
				Reviews:     "47",
				Salaries:    "43",
				ReviewsRate: "4.6",
				Verified:    true,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "Odyssey-Information-Services",
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "https://www.youtube.com/@odysseyinformationservices1641",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Golang Cloud Security Developer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/3905563294/",
							Date:                 mustDate("2024-05-24"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Golang Developer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4152455199/",
							Date:                 mustDate("2025-03-04", "2025-03-01", "2025-02-26", "2025-02-19", "2025-02-18", "2025-02-17", "2025-02-13"),
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
			Industries:       []domain.Industry{},
		},

		// Some | Infomatics Corp
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Infomatics Corp",
			Website: "https://infomaticscorp.com/",
			Careers: "https://infomaticscorp.com/careers/",
			About:   "https://infomaticscorp.com/about-us/",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                821065,
				Alias:             "infomatics-corp",
				Name:              "Infomatics Corp",
				Followers:         "36K",
				Employees:         "51-200",
				AssociatedMembers: "99",
				Verified:          false,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "",
				Verified: false,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Infomatics-EI_IE925978.11,21.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Infomatics-Reviews-E925978.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Infomatics-Jobs-E925978.htm",
				Jobs:        "",
				Reviews:     "54",
				Salaries:    "253",
				ReviewsRate: "4.5",
				Verified:    false,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "Infomatics-Corp",
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Golang Developer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/3912956713/",
							Date:                 mustDate("2024-05-24"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Golang Developer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4124982450/",
							Date:                 mustDate("2025-01-15"),
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
			Industries:       []domain.Industry{},
		},

		// Some | Mindera
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Mindera",
			Website: "https://mindera.com/",
			Careers: "https://mindera.com/careers",
			About:   "https://mindera.com/about",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                80044268,
				Alias:             "mindera-world",
				Name:              "Mindera",
				Followers:         "96K",
				Employees:         "1K-5K",
				AssociatedMembers: "952",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "Mindera",
				Verified: true,
			},
			BlindProfile: domain.BlindProfile{
				Alias:       "Mindera",
				Employees:   "501 to 1,000",
				Salary:      "",
				Reviews:     "6",
				ReviewsRate: "4.5",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "mindera",
				Employees: "760",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Mindera-EI_IE1139926.11,18.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Mindera-Reviews-E1139926.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Mindera-Jobs-E1139926.htm",
				Jobs:        "53",
				Reviews:     "265",
				Salaries:    "516",
				ReviewsRate: "4.3",
				Verified:    true,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "Mindera",
			},
			OttaProfileSlug:   "Mindera",
			YouTubeChannelURL: "https://www.youtube.com/@MinderaSoftware",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Back-End Golang Senior Developer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/3915092213/",
							Date:                 mustDate("2024-05-24"),
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
			Industries:       []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{
				domain.Czechia,
			},
		},

		// Some | Sytac
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Sytac",
			Website: "https://sytac.io/",
			Careers: "https://sytac.io/careers/",
			About:   "https://sytac.io/aboutus/",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                626751,
				Alias:             "sytac",
				Name:              "Sytac",
				Followers:         "7K",
				Employees:         "51-200",
				AssociatedMembers: "102",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "sytac",
				Verified: false,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Sytac-EI_IE1255983.11,16.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Sytac-Reviews-E1255983.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Sytac-Jobs-E1255983.htm",
				Jobs:        "",
				Reviews:     "55",
				Salaries:    "92",
				ReviewsRate: "4.0",
				Verified:    false,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "",
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "https://www.youtube.com/@sytac",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{
						/*
							{
								Title:                "",
								ShortDescription:     "",
								SwitchingOpportunity: "",
								URL:                  "",
								Date:                 mustDate(""),
								WithSalary:           false,
								Remote:               false,
							},
						*/
					},
				},
				domain.Rust: {},
				domain.Zig:  {},
				domain.Scala: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Senior Scala Engineer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4027504713/",
							Date:                 mustDate("2025-06-17"),
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
			Industries:       []domain.Industry{},
		},

		// Some | Qumulus Cloud Platform
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Qumulus Cloud Platform",
			Website: "https://www.qumulus.io/",
			Careers: "",
			About:   "https://www.qumulus.io/about",
			Blog:    "https://www.qumulus.io/blog-pages/blogs",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                81905634,
				Alias:             "qumuluscloudplatform",
				Name:              "Qumulus Cloud Platform",
				Followers:         "1K",
				Employees:         "2-10",
				AssociatedMembers: "13",
				Verified:          false,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "QumulusTechnology",
				Verified: false,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "", // NOP
				ReviewsURL:  "", // NOP
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "",
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 3,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Lead Golang Engineer — Kubernetes",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/3914358799/",
							Date:                 mustDate("2024-05-24"),
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
			Industries:       []domain.Industry{},
		},

		// Some | Saxon AI
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Saxon AI",
			Website: "https://saxon.ai/",
			Careers: "https://saxon.ai/career/",
			About:   "https://saxon.ai/about-us/",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                224935,
				Alias:             "saxonai",
				Name:              "Saxon AI",
				Followers:         "135K",
				Employees:         "201-500",
				AssociatedMembers: "374",
				Verified:          false,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "",
				Verified: false,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "", // NOP
				ReviewsURL:  "", // NOP
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "",
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "https://www.youtube.com/@saxonai",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{
						/*
							{
								Title:                "",
								ShortDescription:     "",
								SwitchingOpportunity: "",
								URL:                  "",
								Date:                 mustDate(""),
								WithSalary:           false,
								Remote:               false,
							},
						*/
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
		},

		// Some | Dyninno Group
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Dyninno Group",
			Website: "https://dyninno.com/",
			Careers: "https://dyninno.com/en/careers/",
			About:   "https://dyninno.com/en/about/",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                9456141,
				Alias:             "dyninno-group",
				Name:              "Dyninno Group",
				Followers:         "60K",
				Employees:         "1K-5K",
				AssociatedMembers: "1,713",
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
				Alias:     "dyninno",
				Employees: "5,400",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-DYNINNO-Group-EI_IE2562842.11,24.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/DYNINNO-Group-Reviews-E2562842.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/DYNINNO-Group-Jobs-E2562842.htm",
				Jobs:        "11",
				Reviews:     "285",
				Salaries:    "316",
				ReviewsRate: "4.1",
				Verified:    true,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "Dyninno-1",
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "https://www.youtube.com/@dyninnogroup2702",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Senior Golang Developer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/3914135515/",
							Date:                 mustDate("2024-05-24"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Senior Golang Developer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4171323905/",
							Date:                 mustDate("2025-03-01"),
							WithSalary:           true, // salary range from 3.800 to 4.600 EUR per month
							Remote:               false,
						},
						{
							Title:                "Senior Golang Developer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4249876193/",
							Date:                 mustDate("2025-06-13"),
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
			ShortDescription: "We create and promote our own products in travel, finance, and entertech",
			Industries:       []domain.Industry{},
		},

		// Some | Holland & Barrett
		{
			ID:      0,
			Name:    "Holland & Barrett",
			Website: "https://www.hollandandbarrett.com/",
			Careers: "https://www.hollandandbarrettjobs.com/",
			About:   "https://www.hollandandbarrett.com/info/who-we-are/",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                815488,
				IDs:               nil,
				Alias:             "815488",
				Name:              "Holland & Barrett",
				Followers:         "68K",
				Employees:         "5K-10K",
				AssociatedMembers: "4,282",
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
				Alias:     "holland-and-barrett",
				Employees: "7,000",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Holland-and-Barrett-EI_IE36174.11,30.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Holland-and-Barrett-Reviews-E36174.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Holland-and-Barrett-Jobs-E36174.htm",
				Jobs:        "182",
				Reviews:     "2.1K",
				Salaries:    "2.7K",
				ReviewsRate: "3.4",
				Verified:    true,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "",
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "https://www.youtube.com/@hollandandbarrett",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Principal Engineer (Golang Engineer)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/3911059005/",
							Date:                 mustDate("2024-05-24"),
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
		},

		// Some | Group Avows
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Group Avows",
			Website: "https://avowstech.com/",
			Careers: "https://avowstech.com/career/",
			About:   "https://avowstech.com/about-us/",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                3651016,
				Alias:             "group-avows",
				Name:              "Group Avows",
				Followers:         "31K",
				Employees:         "501-1K",
				AssociatedMembers: "267",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "",
				Verified: false,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-AVOWS-Technologies-EI_IE870406.11,29.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/AVOWS-Technologies-Reviews-E870406.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/AVOWS-Technologies-Jobs-E870406.htm",
				Jobs:        "18",
				Reviews:     "77",
				Salaries:    "71",
				ReviewsRate: "3.5",
				Verified:    false,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "",
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "https://www.youtube.com/@avowsgroupofficial9907",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Full Stack Developer (Golang)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/3911007016/",
							Date:                 mustDate("2024-05-24"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Golang Developer",
							ShortDescription:     "3 years experience as Golang Developer",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4123654018/",
							Date:                 mustDate("2025-01-14"),
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
			Industries:       []domain.Industry{},
		},

		// Some | Nuro
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Nuro",
			Website: "https://www.nuro.ai/",
			Careers: "https://www.nuro.ai/careers",
			About:   "https://www.nuro.ai/company",
			Blog:    "https://medium.com/nuro",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                12957486,
				Alias:             "nuro-inc.",
				Name:              "Nuro",
				Followers:         "86K",
				Employees:         "501-1K",
				AssociatedMembers: "930",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "",
				Verified: false,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Nuro-EI_IE1550693.11,15.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Nuro-Reviews-E1550693.htm",
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "",
			},
			OttaProfileSlug:   "Nuro",
			YouTubeChannelURL: "https://www.youtube.com/@NuroTeam",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{
						/*
							{
								Title:            "",
								ShortDescription: "",
								URL:              "https://app.welcometothejungle.com/jobs/P5a_50Xb",
								Date:             mustDate(""),
							},
						*/
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
		},

		// Some | CloudWalk
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "CloudWalk",
			Website: "https://cloudwalk.io/",
			Careers: "https://www.cloudwalk.io/jobs",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                3523168,
				Alias:             "cloudwalk-inc",
				Name:              "CloudWalk, Inc.",
				Followers:         "40K",
				Employees:         "501-1K",
				AssociatedMembers: "601",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "cloudwalk",
				Verified: false,
			},
			BlindProfile: domain.BlindProfile{
				Alias: "",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "cloudwalk",
				Employees: "180",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-CloudWalk-Inc-EI_IE2722308.11,24.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/CloudWalk-Inc-Reviews-E2722308.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/CloudWalk-Inc-Jobs-E2722308.htm",
				Jobs:        "",
				Reviews:     "282",
				Salaries:    "490",
				ReviewsRate: "4.3",
				Verified:    true,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "",
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "https://www.youtube.com/@cloudwalk_shorts",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 6,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Senior Software Engineer — Golang",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/3917078597/",
							Date:                 mustDate("2024-08-26"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Senior Software Engineer — Golang",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4088770341/",
							Date:                 mustDate("2024-12-25"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Senior Software Engineer — Golang",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4121841670/",
							Date:                 mustDate("2025-06-05", "2025-05-14", "2025-04-23", "2025-03-10", "2025-02-15", "2025-01-14"),
							WithSalary:           false,
							Remote:               true,
						},
						{
							Title:                "Senior AI-Driven Engineer — Golang",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4149153063/",
							Date:                 mustDate("2025-06-21", "2025-03-04"),
							WithSalary:           false,
							Remote:               true,
						},
					},
				},
				domain.Rust: {
					GitHubRepositoriesCount: 6,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Senior AI-Driven Engineer — Rust",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4129248449/",
							Date:                 mustDate("2025-04-20"), //  mustDate("2025-03-06", "2025-02-12", "2025-01-20"),
							WithSalary:           false,
							Remote:               true,
						},
						{
							Title:                "Senior AI-Driven Engineer – Rust",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4180652646/",
							Date:                 mustDate("2025-06-08", "2025-04-25", "2025-04-04", "2025-03-13"),
							WithSalary:           false,
							Remote:               true,
						},
						{
							Title:                "Senior AI-Driven Engineer – Rust",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4152784417/",
							Date:                 mustDate("2025-06-05", "2025-05-15", "2025-04-23", "2025-04-01"),
							WithSalary:           false,
							Remote:               true,
						},
						{
							Title:                "AI-Driven Engineer – Rust (PWD/PCD)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4255818917/",
							Date:                 mustDate("2025-06-28"),
							WithSalary:           false,
							Remote:               true,
						},
						{
							Title:                "AI-Driven Engineer – Rust",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4253975338/",
							Location:             "São Paulo, São Paulo, Brazil",
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
			ShortDescription: "",
			Industries:       []domain.Industry{},
		},

		// Some | Transition Technologies PSC
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Transition Technologies PSC",
			Website: "https://ttpsc.com/",
			Careers: "https://kariera.ttpsc.com/en/job-offers/",
			About:   "https://ttpsc.com/en/who-we-are/",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                17880075,
				Alias:             "transition-technologies-psc",
				Name:              "Transition Technologies PSC",
				Followers:         "6K",
				Employees:         "501-1K",
				AssociatedMembers: "622",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "",
				Verified: false,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Transition-Technologies-PSC-EI_IE1875542.11,38.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Transition-Technologies-PSC-Reviews-E1875542.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Transition-Technologies-PSC-Jobs-E1875542.htm",
				Jobs:        "",
				Reviews:     "47",
				Salaries:    "83",
				ReviewsRate: "3.6",
				Verified:    false,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "",
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "https://www.youtube.com/@TransitionTechnologiesPSC",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Golang Developer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/3886949034/",
							Date:                 mustDate("2024-05-24"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Golang Developer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4167039797/",
							Date:                 mustDate("2025-02-24"),
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
			Industries:       []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
			},
		},

		// Some | Kroger
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Kroger",
			Website: "https://www.kroger.com/",
			Careers: "https://www.krogerfamilycareers.com/en/sites/CX_2001",
			About:   "https://www.thekrogerco.com/about-kroger/",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                4914,
				Alias:             "kroger",
				Name:              "Kroger",
				Followers:         "454K",
				Employees:         "10K+",
				AssociatedMembers: "135,606",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:     "Kroger-Technology",
				Followers: "14",
				Verified:  false,
			},
			BlindProfile: domain.BlindProfile{
				Alias:       "Kroger",
				Employees:   "10,000+",
				Salary:      "$23K ~ $200K a year",
				Reviews:     "72",
				ReviewsRate: "3.7",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "kroger",
				Employees: "465,000",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Kroger-EI_IE386.11,17.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Kroger-Reviews-E386.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Kroger-Jobs-E386.htm",
				Jobs:        "9.9K",
				Reviews:     "24K",
				Salaries:    "34K",
				ReviewsRate: "3.2",
				Verified:    false,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "Kroger-54356a75",
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "https://www.youtube.com/@KrogerCo",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Advanced Software Engineer (Java/Golang)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/3917025739/",
							Date:                 mustDate("2024-05-24"),
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
			Industries:       []domain.Industry{},
		},

		// Some | Precisely
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Precisely",
			Website: "https://www.precisely.com/",
			Careers: "https://www.precisely.com/careers-and-culture",
			About:   "https://www.precisely.com/about-us",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                64863146,
				Alias:             "preciselydata",
				Name:              "Precisely",
				Followers:         "149K",
				Employees:         "1K-5K",
				AssociatedMembers: "2,958",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "PreciselyData",
				Verified: false,
			},
			BlindProfile: domain.BlindProfile{
				Alias: "",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "precisely",
				Employees: "2,000",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Precisely-EI_IE3372755.11,20.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Precisely-Reviews-E3372755.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Precisely-Jobs-E3372755.htm",
				Jobs:        "24",
				Reviews:     "434",
				Salaries:    "789",
				ReviewsRate: "4.1",
				Verified:    true,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "Precisely-1",
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "https://www.youtube.com/@PreciselyData",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 1,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Principal Software Engineer (Golang)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/3917000992/",
							Date:                 mustDate("2024-06-25"),
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
			Industries:       []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
				domain.Czechia,
			},
		},

		// Some | R Systems
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "R Systems",
			Website: "https://www.rsystems.com/",
			Careers: "https://www.rsystems.com/apac/career/",
			About:   "https://www.rsystems.com/about-us/",
			Blog:    "https://eu.rsystems.com/category/blog/tech/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                165636,
				Alias:             "r-systems",
				Name:              "R Systems",
				Followers:         "146K",
				Employees:         "1K-5K",
				AssociatedMembers: "5,830",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "",
				Verified: false,
			},
			BlindProfile: domain.BlindProfile{
				Alias:       "R-Systems",
				Employees:   "1,001 to 5,000",
				Salary:      "",
				Reviews:     "5",
				ReviewsRate: "4.2",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "r-systems",
				Employees: "6,000",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-R-Systems-EI_IE32444.11,20.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/R-Systems-Reviews-E32444.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/R-Systems-Jobs-E32444.htm",
				Jobs:        "12",
				Reviews:     "1.5K",
				Salaries:    "2K",
				ReviewsRate: "3.7",
				Verified:    true,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "R-Systems-1",
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "https://www.youtube.com/@RSystems_inc",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Golang Developer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/3916555239/",
							Date:                 mustDate("2024-05-24"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Golang developer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4245442819/",
							Date:                 mustDate("2025-07-02"),
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
			Industries:       []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
			},
		},

		// Some | Hays
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Hays",
			Website: "https://www.haysplc.com/",
			Careers: "https://www.haysplc.com/joinhays",
			About:   "https://www.haysplc.com/about-us",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                3486,
				Alias:             "hays",
				Name:              "Hays",
				Followers:         "8M",
				Employees:         "5K-10K",
				AssociatedMembers: "26,240",
				Verified:          true,
			},
			Ignore: true, // The deleted recruiting agency was added by mistake
		},

		// Some | Consort Group
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Consort Group",
			Website: "https://consort-group.com/",
			Careers: "https://consort-group.com/en/join-us/our-jobs/",
			About:   "https://consort-group.com/en/consort-group-2/",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                46088,
				Alias:             "consortgroup",
				Name:              "Consort Group",
				Followers:         "65K",
				Employees:         "1K-5K",
				AssociatedMembers: "1,666",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "",
				Verified: false,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Groupe-Consort-EI_IE915503.11,25.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Groupe-Consort-Reviews-E915503.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Groupe-Consort-Jobs-E915503.htm",
				Jobs:        "",
				Reviews:     "237",
				Salaries:    "469",
				ReviewsRate: "3.6",
				Verified:    false,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "Consort-Group-5",
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "https://www.youtube.com/@consortgroup",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Golang Developer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/3913831485/",
							Date:                 mustDate("2024-05-24"),
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
			Industries:       []domain.Industry{},
		},

		// Some | Ascendion
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Ascendion",
			Website: "https://ascendion.com/",
			Careers: "https://ascendion.com/careers/",
			About:   "https://ascendion.com/who-we-are/",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                86694680,
				Alias:             "ascendion",
				Name:              "Ascendion",
				Followers:         "1M",
				Employees:         "5K-10K",
				AssociatedMembers: "3,039",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "",
				Verified: false,
			},
			BlindProfile: domain.BlindProfile{
				Alias:       "Ascendion",
				Employees:   "1,001 to 5,000",
				Salary:      "",
				Reviews:     "1",
				ReviewsRate: "4.0",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "ascendion",
				Employees: "150",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Ascendion-EI_IE7774544.11,20.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Ascendion-Reviews-E7774544.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Ascendion-Jobs-E7774544.htm",
				Jobs:        "67",
				Reviews:     "549",
				Salaries:    "682",
				ReviewsRate: "4.1",
				Verified:    true,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "Ascendion",
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "https://www.youtube.com/@ascendioninc",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Senior Golang Developer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4248770904/",
							Date:                 mustDate("2025-06-14"),
							WithSalary:           true,
							Remote:               true,
						},
					},
				},
				domain.Rust: {},
				domain.Zig:  {},
				domain.Scala: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Scala Developer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4117260547/",
							Date:                 mustDate("2025-01-27"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Scala Developer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4221272747/",
							Date:                 mustDate("2025-05-02"),
							WithSalary:           true,
							Remote:               false,
						},
						{
							Title:                "Big Data Developer (Scala/Spark + React)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4225312611/",
							Date:                 mustDate("2025-05-07"),
							WithSalary:           false,
							Remote:               true,
						},
						{
							Title:                "Scala Engineer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4267351389/",
							Location:             "New Jersey, United States",
							Date:                 mustDate("2025-07-15"),
							WithSalary:           false,
							Remote:               true,
						},
					},
				},
				domain.Elixir:  {},
				domain.Clojure: {},
				domain.Haskell: {},
			},
			ShortDescription: "",
			Industries:       []domain.Industry{},
		},

		// Some | Checkout.com
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Checkout.com",
			Website: "https://www.checkout.com/",
			Careers: "https://www.checkout.com/careers",
			About:   "https://www.checkout.com/mission",
			Blog:    "https://medium.com/checkout-com-techblog",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                3110635,
				Alias:             "checkout",
				Name:              "Checkout.com",
				Followers:         "205K",
				Employees:         "1K-5K",
				AssociatedMembers: "1,956",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "checkout",
				Verified: false,
			},
			BlindProfile: domain.BlindProfile{
				Alias:       "Checkoutcom",
				Employees:   "1,001 to 5,000",
				Salary:      "",
				Reviews:     "27",
				ReviewsRate: "2.4",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "checkoutcom",
				Employees: "1,500",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Checkout-com-EI_IE837487.11,23.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Checkout-com-Reviews-E837487.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Checkout-com-Jobs-E837487.htm",
				Jobs:        "10",
				Reviews:     "837",
				Salaries:    "1.8K",
				ReviewsRate: "4.0",
				Verified:    true,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "Checkout.com",
			},
			OttaProfileSlug:   "Checkout-com",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 1,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Senior Software Engineer (Golang) — Payment Performance",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/3920739722/",
							Date:                 mustDate("2024-06-25"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Software Engineer (Golang)",
							ShortDescription:     "ProcessOut — Payments",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4125751681/",
							Date:                 mustDate("2025-02-04", "2025-01-14"),
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
			Industries:       []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
				domain.Czechia,
			},
		},
	}
}
