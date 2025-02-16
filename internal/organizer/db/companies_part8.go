package db

import (
	"github.com/readytotouch/readytotouch/internal/domain"
)

func companiesPart8() []domain.CompanyProfile {
	return []domain.CompanyProfile{

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
			ShortDescription: "Software development company",
			Ignore:           true, // The deleted outsource company was added by mistake
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "ATG (Auction Technology Group)",
			Website: "https://www.auctiontechnologygroup.com/",
			Careers: "https://www.auctiontechnologygroup.com/careers",
			About:   "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                565927,
				Alias:             "auctiontechnologygroup",
				Name:              "ATG (Auction Technology Group)",
				Followers:         "17K",
				Employees:         "201-500",
				AssociatedMembers: "509",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "",
				Verified: false,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Auction-Technology-Group-EI_IE2071353.11,35.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Auction-Technology-Group-Reviews-E2071353.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Auction-Technology-Group-Jobs-E2071353.htm",
				Jobs:        "",
				Reviews:     "83",
				Salaries:    "103",
				ReviewsRate: "3.5",
				Verified:    true,
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Senior Golang Developer",
							ShortDescription:     "3+ years experience with Golang",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4117172166/",
							Date:                 mustDate("2025-01-08"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Senior Golang Developer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4148683119/",
							Date:                 mustDate("2025-02-10"),
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
			ShortDescription: "Online platform for auctions",
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "gridX",
			Website: "https://www.gridx.ai/",
			Careers: "https://www.gridx.ai/careers",
			About:   "https://www.gridx.ai/about",
			Blog:    "https://www.gridx.ai/blog-category/tech",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                10599040,
				Alias:             "gridx",
				Name:              "gridX",
				Followers:         "21K",
				Employees:         "51-200",
				AssociatedMembers: "199",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "grid-x",
				Verified: true,
			},
			BlindProfile: domain.BlindProfile{},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "gridx",
				Employees: "126",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-gridX-EI_IE2866338.11,16.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/gridX-Reviews-E2866338.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/gridX-Jobs-E2866338.htm",
				Jobs:        "6",
				Reviews:     "17",
				Salaries:    "39",
				ReviewsRate: "4.1",
				Verified:    false,
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Golang Developer",
							ShortDescription:     "Golang for at least 2+ years and preferably have an understanding of IoT that you can apply",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4118507935/",
							Date:                 mustDate("2025-01-07"),
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
			ShortDescription: "Developments in the energy industry",
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Evolve",
			Website: "https://evolve.com/",
			Careers: "https://evolve.com/careers",
			About:   "https://evolve.com/our-people",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                2114845,
				Alias:             "evolve-vacation-rental",
				Name:              "Evolve",
				Followers:         "24K",
				Employees:         "501-1K",
				AssociatedMembers: "1,529",
				Verified:          true,
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
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Senior Software Engineer Golang",
							ShortDescription:     "5+ years of backend development experience building applications in Go, NodeJS, and integrations with 3rd Party tools, APIs, and other services",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4119005072/",
							Date:                 mustDate("2025-01-07"),
							WithSalary:           true, // salary range is $120,000 to $185,000
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
			ShortDescription: "Technology-driven vacation rental company",
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "StrangeBee",
			Website: "https://strangebee.com/",
			Careers: "https://strangebee.com/careers-at-strangebee/",
			About:   "https://strangebee.com/about-strangebee/",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                18993016,
				Alias:             "strangebee",
				Name:              "StrangeBee",
				Followers:         "6K",
				Employees:         "51-200",
				AssociatedMembers: "43",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "StrangeBeeCorp",
				Verified: true,
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
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Golang Engineer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4121276014/",
							Date:                 mustDate("2025-01-09"),
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
			ShortDescription: "CyberSecurity software company",
			Industries: []domain.Industry{
				domain.IndustryCyberSecurity,
			},
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "thatgamecompany",
			Website: "https://thatgamecompany.com/",
			Careers: "https://thatgamecompany.com/careers/",
			About:   "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                1637083,
				Alias:             "thatgamecompany",
				Name:              "thatgamecompany",
				Followers:         "24K",
				Employees:         "51-200",
				AssociatedMembers: "197",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "thatgameco",
				Verified: false,
			},
			BlindProfile: domain.BlindProfile{},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "thatgamecompany",
				Employees: "150",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-thatgame-EI_IE2130792.11,19.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/thatgame-Reviews-E2130792.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/thatgame-Jobs-E2130792.htm",
				Jobs:        "28",
				Reviews:     "8",
				Salaries:    "17",
				ReviewsRate: "4.4",
				Verified:    false,
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Senior Software Engineer — Golang",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4120264764/",
							Date:                 mustDate("2025-01-09"),
							WithSalary:           true, // Salary range is $95,000 USD to $177,000 USD
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
			ShortDescription: "American video game development company",
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "BetterMe",
			Website: "https://betterme.world/",
			Careers: "https://betterme.world/career",
			About:   "https://betterme.world/about",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                18550246,
				Alias:             "betterme-company",
				Name:              "BetterMe",
				Followers:         "12K",
				Employees:         "201-500",
				AssociatedMembers: "482",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "betterme-dev",
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
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Backend Engineer (Go + PHP)",
							ShortDescription:     "Design and implement backend features from conception through to deployment, ensuring they meet strategic business objectives",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4121587727/",
							Date:                 mustDate("2025-01-10"),
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
			ShortDescription: "Digital fitness company",
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "SPE Solution",
			Website: "https://www.spesolution.com/",
			Careers: "https://www.spesolution.com/career",
			About:   "https://www.spesolution.com/company",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                30205768,
				Alias:             "spe-solution",
				Name:              "SPE Solution",
				Followers:         "7K",
				Employees:         "201-500",
				AssociatedMembers: "190",
				Verified:          true,
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
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Backend Developer — Golang",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4120414743/",
							Date:                 mustDate("2025-01-10"), // mustDate("2025-01-10"),
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
			ShortDescription: "Indonesian FinTech solution company",
			Industries: []domain.Industry{
				domain.IndustryFinTech,
			},
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Quantum Systems",
			Website: "https://quantum-systems.com/",
			Careers: "https://career.quantum-systems.com/",
			About:   "https://quantum-systems.com/us/about-us/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                25007323,
				Alias:             "quantum-systems-gmbh",
				Name:              "Quantum Systems",
				Followers:         "37K",
				Employees:         "201-500",
				AssociatedMembers: "284",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "",
				Verified: false,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Quantum-Systems-EI_IE3879838.11,26.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Quantum-Systems-Reviews-E3879838.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Quantum-Systems-Jobs-E3879838.htm",
				Jobs:        "27",
				Reviews:     "24",
				Salaries:    "27",
				ReviewsRate: "4.7",
				Verified:    true,
			},
			Languages: domain.Languages{
				domain.Go: {},
				domain.Rust: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Rust Developer",
							ShortDescription:     "5+ years of professional experience in backend development with a strong focus on real-time systems",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4119548956/",
							Date:                 mustDate("2025-01-16"), // mustDate("2025-01-10"),
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
			ShortDescription: "German technology company",
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "General Dynamics Mission Systems",
			Website: "https://gdmissionsystems.com/",
			Careers: "https://gdmissionsystems.com/careers",
			About:   "https://gdmissionsystems.com/about-us",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                1908,
				Alias:             "gdms",
				Name:              "General Dynamics Mission Systems",
				Followers:         "165K",
				Employees:         "10K",
				AssociatedMembers: "8,822",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "",
				Verified: false,
			},
			BlindProfile: domain.BlindProfile{},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "general-dynamics-mission-systems",
				Employees: "8,450",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-General-Dynamics-Mission-Systems-EI_IE941740.11,43.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/General-Dynamics-Mission-Systems-Reviews-E941740.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/General-Dynamics-Mission-Systems-Jobs-E941740.htm",
				Jobs:        "410",
				Reviews:     "2K",
				Salaries:    "3.8K",
				ReviewsRate: "3.9",
				Verified:    true,
			},
			Languages: domain.Languages{
				domain.Go: {},
				domain.Rust: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Rust Embedded Software Engineer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4118966168/",
							Date:                 mustDate("2025-02-10"), // mustDate("2025-01-09"),
							WithSalary:           true,                   // Starting at $108K/yr
							Remote:               false,
						},
						{
							Title:                "Rust Embedded Software Engineer",
							ShortDescription:     "Clearable",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4148907792/",
							Date:                 mustDate("2025-02-13"),
							WithSalary:           true, // salary range $95.100 - $256.000 per year
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
			ShortDescription: "Integrates secure communication and information systems and technology",
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
			ShortDescription: "Sweden software-driven tech company",
			Ignore:           true, // The deleted outsource company was added by mistake
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Mondu",
			Website: "https://www.mondu.ai/",
			Careers: "https://www.mondu.ai/careers/",
			About:   "https://www.mondu.ai/about/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                76198623,
				Alias:             "mondu-ai",
				Name:              "Mondu",
				Followers:         "12K",
				Employees:         "51-200",
				AssociatedMembers: "111",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "mondu-ai",
				Verified: true,
			},
			BlindProfile: domain.BlindProfile{},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "mondu",
				Employees: "145",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Mondu-EI_IE6725727.11,16.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Mondu-Reviews-E6725727.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Mondu-Jobs-E6725727.htm",
				Jobs:        "3",
				Reviews:     "16",
				Salaries:    "51",
				ReviewsRate: "3.7",
				Verified:    false,
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Senior Golang Engineer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4121324436/",
							Date:                 mustDate("2025-01-13"),
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
			ShortDescription: "European Fintech offering innovative payment solutions for B2B commerce",
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "株式会社マネーフォワード — Money Forward",
			Website: "https://corp.moneyforward.com/en/",
			Careers: "https://recruit.moneyforward.com/en/",
			About:   "https://corp.moneyforward.com/en/aboutus/",
			Blog:    "https://mfi.engineering/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                5775746,
				Alias:             "money-forward",
				Name:              "株式会社マネーフォワード — Money Forward",
				Followers:         "9K",
				Employees:         "1K-5K",
				AssociatedMembers: "893",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "moneyforward",
				Verified: true,
			},
			BlindProfile: domain.BlindProfile{},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "money-forward",
				Employees: "1,250",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Money-Forward-EI_IE3841421.11,24.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Money-Forward-Reviews-E3841421.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Money-Forward-Jobs-E3841421.htm",
				Jobs:        "481",
				Reviews:     "63",
				Salaries:    "100",
				ReviewsRate: "3.9",
				Verified:    false,
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Senior Backend Engineer (Go)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4126118648/",
							Date:                 mustDate("2025-01-15"),
							WithSalary:           true, // 659,000 JPY — 〜834,000 JPY / month
						},
					},
				},
				domain.Rust:    {},
				domain.Zig:     {},
				domain.Scala:   {},
				domain.Elixir:  {},
				domain.Clojure: {},
				domain.Haskell: {},
			},
			ShortDescription: "FinTech company",
			Industries: []domain.Industry{
				domain.IndustryFinTech,
			},
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "PrizePicks",
			Website: "https://prizepicks.com/",
			Careers: "https://www.prizepicks.com/careers",
			About:   "https://www.prizepicks.com/resources/what-is-prizepicks-how-to-play-promo-code",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                11477578,
				Alias:             "prize-picks",
				Name:              "PrizePicks",
				Followers:         "21K",
				Employees:         "501-1K",
				AssociatedMembers: "543",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "",
				Verified: false,
			},
			BlindProfile: domain.BlindProfile{},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "prizepicks",
				Employees: "60",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-PrizePicks-EI_IE2582151.11,21.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/PrizePicks-Reviews-E2582151.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/PrizePicks-Jobs-E2582151.htm",
				Jobs:        "16",
				Reviews:     "90",
				Salaries:    "98",
				ReviewsRate: "4.2",
				Verified:    true,
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Software Engineer III (Go)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4124190818/",
							Date:                 mustDate("2025-02-12"), // mustDate("2025-01-15"),
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
			ShortDescription: "A daily fantasy sports operator",
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "CHEP",
			Website: "https://www.chep.com/",
			Careers: "https://www.chep.com/careers/your-career-chep",
			About:   "https://www.chep.com/about-us/about-us",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                5085,
				Alias:             "chep",
				Name:              "CHEP",
				Followers:         "398K",
				Employees:         "10K+",
				AssociatedMembers: "8,831",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "",
				Verified: false,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-CHEP-EI_IE22580.11,15.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/CHEP-Reviews-E22580.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/CHEP-Jobs-E22580.htm",
				Jobs:        "56",
				Reviews:     "1.1K",
				Salaries:    "1.7K",
				ReviewsRate: "4.0",
				Verified:    true,
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Full Stack Developer (Golang and Svelte)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4123557217/",
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
			ShortDescription: "Company that specializes in pooling and managing pallets, crates, and containers for efficient and sustainable supply chain operations",
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Optimizely",
			Website: "https://www.optimizely.com/",
			Careers: "https://careers.optimizely.com/",
			About:   "https://www.optimizely.com/company/",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                1189697,
				Alias:             "optimizely",
				Name:              "Optimizely",
				Followers:         "140K",
				Employees:         "1K-5K",
				AssociatedMembers: "1,528",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "optimizely",
				Verified: false,
			},
			BlindProfile: domain.BlindProfile{},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "optimizely",
				Employees: "840",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Optimizely-EI_IE498757.11,21.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Optimizely-Reviews-E498757.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Optimizely-Jobs-E498757.htm",
				Jobs:        "27",
				Reviews:     "546",
				Salaries:    "913",
				ReviewsRate: "4.1",
				Verified:    true,
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Senior Software Engineer (Golang)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4120974608/",
							Date:                 mustDate("2025-02-05"), // mustDate("2025-01-15"),
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
			ShortDescription: "American Digital Experience Platform",
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "OPSWAT",
			Website: "https://www.opswat.com/",
			Careers: "https://www.opswat.com/careers",
			About:   "https://www.opswat.com/company/about",
			Blog:    "https://www.opswat.com/blog",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                105936,
				Alias:             "opswat",
				Name:              "OPSWAT",
				Followers:         "106K",
				Employees:         "501-1K",
				AssociatedMembers: "952",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "OPSWAT",
				Verified: false,
			},
			BlindProfile: domain.BlindProfile{},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "opswat",
				Employees: "490",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-OPSWAT-EI_IE266447.11,17.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/OPSWAT-Reviews-E266447.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/OPSWAT-Jobs-E266447.htm",
				Jobs:        "79",
				Reviews:     "251",
				Salaries:    "473",
				ReviewsRate: "3.4",
				Verified:    false,
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Software Engineer (Golang)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4124852341/",
							Date:                 mustDate("2025-01-15"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Backend Developer (Golang)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4130863951/",
							Date:                 mustDate("2025-02-11"), // mustDate("2025-01-21"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Software Engineer (Golang)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4152664916/",
							Date:                 mustDate("2025-02-14"),
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
			ShortDescription: "American cybersecurity company",
			Industries: []domain.Industry{
				domain.IndustryCyberSecurity,
			},
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "VideoAmp",
			Website: "https://videoamp.com/",
			Careers: "https://videoamp.com/jobs/",
			About:   "https://videoamp.com/our-company/",
			Blog:    "https://videoamp.com/blog/category/engineering/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                3704165,
				Alias:             "videoamp",
				Name:              "VideoAmp",
				Followers:         "19K",
				Employees:         "201-500",
				AssociatedMembers: "611",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "",
				Verified: false,
			},
			BlindProfile: domain.BlindProfile{},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "videoamp",
				Employees: "420",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-VideoAmp-EI_IE1143644.11,19.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/VideoAmp-Reviews-E1143644.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/VideoAmp-Jobs-E1143644.htm",
				Jobs:        "9",
				Reviews:     "154",
				Salaries:    "343",
				ReviewsRate: "3.0",
				Verified:    true,
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Senior Full Stack Engineer — HTMX / Golang",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4122141025/",
							Date:                 mustDate("2025-02-04"), // mustDate("2025-01-14"),
							WithSalary:           true,                   // This position has a minimum salary of $150,000- $170,000
							Remote:               true,
						},
						{
							Title:                "Senior Full Stack Engineer — HTMX / Golang",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4150889365/",
							Date:                 mustDate("2025-02-12"),
							WithSalary:           true, // This position has a minimum salary of $150,000- $170,000
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
			ShortDescription: "Media measurement company revolutionizing advertising",
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Broadcom",
			Website: "https://www.broadcom.com/",
			Careers: "https://www.broadcom.com/company/careers",
			About:   "https://www.broadcom.com/company/about-us",
			Blog:    "https://www.broadcom.com/blog",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                3072,
				Alias:             "broadcom",
				Name:              "Broadcom",
				Followers:         "498K",
				Employees:         "10K+",
				AssociatedMembers: "60,347",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "broadcom",
				Verified: true,
			},
			BlindProfile: domain.BlindProfile{},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "broadcom",
				Employees: "37,990",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Broadcom-EI_IE6926.11,19.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Broadcom-Reviews-E6926.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Broadcom-Jobs-E6926.htm",
				Jobs:        "272",
				Reviews:     "6.2K",
				Salaries:    "14K",
				ReviewsRate: "3.2",
				Verified:    false,
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Senior Golang Developer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4127239750/",
							Date:                 mustDate("2025-01-17"),
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
			ShortDescription: "American global technology company",
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Toyota North America",
			Website: "https://www.toyota.com/",
			Careers: "https://www.toyota.com/careers",
			About:   "https://www.toyota.com/brand/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                2107,
				Alias:             "toyota-north-america",
				Name:              "Toyota North America",
				Followers:         "435K",
				Employees:         "10K+",
				AssociatedMembers: "19,741",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "",
				Verified: false,
			},
			BlindProfile: domain.BlindProfile{},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "toyota-usa",
				Employees: "30,860",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Toyota-North-America-EI_IE3544.11,31.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Toyota-North-America-Reviews-E3544.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Toyota-North-America-Jobs-E3544.htm",
				Jobs:        "211",
				Reviews:     "4.3K",
				Salaries:    "6.8K",
				ReviewsRate: "4.0",
				Verified:    true,
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Golang Developer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4125175696/",
							Date:                 mustDate("2025-01-16"),
							WithSalary:           true, // $60.00 - $77.00/ hr
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
			ShortDescription: "Manufacturing companies",
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "ABOUT YOU",
			Website: "https://aboutyou.de",
			Careers: "https://corporate.aboutyou.de/en/career",
			About:   "https://corporate.aboutyou.de/en/about-us",
			Blog:    "https://medium.com/about-developer-blog",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                5214302,
				Alias:             "about-you",
				Name:              "ABOUT YOU",
				Followers:         "102K",
				Employees:         "1K-5K",
				AssociatedMembers: "1,529",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "",
				Verified: false,
			},
			BlindProfile: domain.BlindProfile{},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "about-you",
				Employees: "690",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-About-You-EI_IE1111018.11,20.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/About-You-Reviews-E1111018.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/About-You-Jobs-E1111018.htm",
				Jobs:        "57",
				Reviews:     "294",
				Salaries:    "691",
				ReviewsRate: "3.5",
				Verified:    true,
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Senior Golang Developer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4124350189/",
							Date:                 mustDate("2025-01-16"),
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
			ShortDescription: "German fashion online retailer",
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Magicmotorsport",
			Website: "https://www.magicmotorsport.com/",
			Careers: "https://www.magicmotorsport.com/about-us/careers/",
			About:   "https://www.magicmotorsport.com/about-us/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                10541609,
				Alias:             "magicmotorsport",
				Name:              "Magicmotorsport",
				Followers:         "6K",
				Employees:         "51-200",
				AssociatedMembers: "128",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "",
				Verified: false,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-MAGICMOTORSPORT-EI_IE6079983.11,26.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/MAGICMOTORSPORT-Reviews-E6079983.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/MAGICMOTORSPORT-Jobs-E6079983.htm",
				Jobs:        "",
				Reviews:     "7",
				Salaries:    "8",
				ReviewsRate: "4.7",
				Verified:    false,
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
							Title:                "Ruby/Elixir Back-End Developer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4119784307/",
							Date:                 mustDate("2025-02-07"), // mustDate("2025-01-10"),
							WithSalary:           false,
							Remote:               false,
						},
					},
				},
				domain.Clojure: {},
				domain.Haskell: {},
			},
			ShortDescription: "Development and production of equipment for chip tuning and diagnostics of electronic control units (ECU) of cars",
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "STATION F",
			Website: "https://stationf.co/",
			Careers: "https://jobs.stationf.co/",
			About:   "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                10776313,
				Alias:             "stationf",
				Name:              "STATION F",
				Followers:         "171K",
				Employees:         "11-50",
				AssociatedMembers: "266",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "stationf",
				Verified: false,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Station-F-EI_IE2991167.11,20.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Station-F-Reviews-E2991167.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Station-F-Jobs-E2991167.htm",
				Jobs:        "",
				Reviews:     "19",
				Salaries:    "41",
				ReviewsRate: "4.4",
				Verified:    false,
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
							Title:                "Senior Full Stack Engineer (Elixir/React)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4124020794/",
							Date:                 mustDate("2025-01-14"),
							WithSalary:           true, // 55000€ — 69000€
							Remote:               false,
						},
					},
				},
				domain.Clojure: {},
				domain.Haskell: {},
			},
			ShortDescription: "French startup campus",
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Nordea",
			Website: "https://www.nordea.com/",
			Careers: "https://www.nordea.com/en/careers",
			About:   "https://www.nordea.com/en/about-us",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                2007,
				Alias:             "nordea",
				Name:              "Nordea",
				Followers:         "283K",
				Employees:         "10K+",
				AssociatedMembers: "29,892",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "",
				Verified: false,
			},
			BlindProfile: domain.BlindProfile{},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "nordea",
				Employees: "27,000",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Nordea-Bank-EI_IE10381.11,22.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Nordea-Bank-Reviews-E10381.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Nordea-Bank-Jobs-E10381.htm",
				Jobs:        "52",
				Reviews:     "1.6K",
				Salaries:    "2.8K",
				ReviewsRate: "4.0",
				Verified:    false,
			},
			Languages: domain.Languages{
				domain.Go:   {},
				domain.Rust: {},
				domain.Zig:  {},
				domain.Scala: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Senior IT Developer — Scala/Spark/Hadoop/Big Data",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4056511122/",
							Date:                 mustDate("2025-01-17"),
							WithSalary:           false,
							Remote:               false,
						},
					},
				},
				domain.Elixir:  {},
				domain.Clojure: {},
				domain.Haskell: {},
			},
			ShortDescription: "Nordic financial services group",
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Deutsche Börse",
			Website: "https://www.deutsche-boerse.com/dbg-en/",
			Careers: "https://careers.deutsche-boerse.com/",
			About:   "https://www.deutsche-boerse.com/dbg-en/about-us",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                163430,
				Alias:             "deutscheboerse",
				Name:              "Deutsche Börse",
				Followers:         "97K",
				Employees:         "10K+",
				AssociatedMembers: "3,934",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "",
				Verified: false,
			},
			BlindProfile: domain.BlindProfile{},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "deutsche-borse",
				Employees: "10,200",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Deutsche-B%C3%B6rse-Group-EI_IE11878.11,31.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Deutsche-B%C3%B6rse-Group-Reviews-E11878.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Deutsche-B%C3%B6rse-Group-Jobs-E11878.htm",
				Jobs:        "320",
				Reviews:     "775",
				Salaries:    "1.3K",
				ReviewsRate: "4.1",
				Verified:    true,
			},
			Languages: domain.Languages{
				domain.Go:   {},
				domain.Rust: {},
				domain.Zig:  {},
				domain.Scala: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Big Data Developer / Engineer (Scala/Java)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4124323501/",
							Date:                 mustDate("2025-01-15"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Big Data Developer (Scala/Spark)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4088735395/",
							Date:                 mustDate("2025-01-30"),
							WithSalary:           false,
							Remote:               false,
						},
					},
				},
				domain.Elixir:  {},
				domain.Clojure: {},
				domain.Haskell: {},
			},
			ShortDescription: "German stock exchange",
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Hootsuite",
			Website: "https://www.hootsuite.com/",
			Careers: "https://careers.hootsuite.com/",
			About:   "https://www.hootsuite.com/about",
			Blog:    "https://medium.com/hootsuite-engineering",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                288540,
				Alias:             "hootsuite",
				Name:              "Hootsuite",
				Followers:         "511K",
				Employees:         "1K-5K",
				AssociatedMembers: "1,668",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "",
				Verified: false,
			},
			BlindProfile: domain.BlindProfile{},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "hootsuite",
				Employees: "1,500",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Hootsuite-EI_IE617216.11,20.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Hootsuite-Reviews-E617216.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Hootsuite-Jobs-E617216.htm",
				Jobs:        "36",
				Reviews:     "823",
				Salaries:    "1.5K",
				ReviewsRate: "3.9",
				Verified:    false,
			},
			Languages: domain.Languages{
				domain.Go:   {},
				domain.Rust: {},
				domain.Zig:  {},
				domain.Scala: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Senior Software Developer, Backend (Scala)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4123118327/",
							Date:                 mustDate("2025-02-07"), // mustDate("2025-01-15"),
							WithSalary:           true,                   // $98,400—$137,800 CAD
							Remote:               false,
						},
					},
				},
				domain.Elixir:  {},
				domain.Clojure: {},
				domain.Haskell: {},
			},
			ShortDescription: "Social media management platform",
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "HSBC",
			Website: "https://www.hsbc.com/",
			Careers: "https://www.hsbc.com/careers",
			About:   "https://www.hsbc.com/who-we-are",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                1241,
				Alias:             "hsbc",
				Name:              "HSBC",
				Followers:         "4M",
				Employees:         "10K+",
				AssociatedMembers: "180,941",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "hsbc",
				Verified: true,
			},
			BlindProfile: domain.BlindProfile{},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "hsbc",
				Employees: "205,340",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-HSBC-EI_IE3482.11,15.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/HSBC-Reviews-E3482.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/HSBC-Jobs-E3482.htm",
				Jobs:        "1.8K",
				Reviews:     "33K",
				Salaries:    "50K",
				ReviewsRate: "3.9",
				Verified:    true,
			},
			Languages: domain.Languages{
				domain.Go:   {},
				domain.Rust: {},
				domain.Zig:  {},
				domain.Scala: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Senior React and Scala/Java Developer (Investment Banking Tech)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4122037400/",
							Date:                 mustDate("2025-01-10"),
							WithSalary:           false,
							Remote:               false,
						},
					},
				},
				domain.Elixir:  {},
				domain.Clojure: {},
				domain.Haskell: {},
			},
			ShortDescription: "British FinTech company",
			Industries: []domain.Industry{
				domain.IndustryFinTech,
			},
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "AffiniPay",
			Website: "https://www.affinipay.com/",
			Careers: "https://www.affinipay.com/careers/",
			About:   "https://www.affinipay.com/company/about-us/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                1276962,
				Alias:             "affinipay",
				Name:              "AffiniPay",
				Followers:         "12K",
				Employees:         "201-500",
				AssociatedMembers: "651",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "affinipay",
				Verified: false,
			},
			BlindProfile: domain.BlindProfile{},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "affinipay",
				Employees: "210",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-AffiniPay-EI_IE528570.11,20.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/AffiniPay-Reviews-E528570.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/AffiniPay-Jobs-E528570.htm",
				Jobs:        "",
				Reviews:     "117",
				Salaries:    "255",
				ReviewsRate: "3.5",
				Verified:    true,
			},
			Languages: domain.Languages{
				domain.Go:   {},
				domain.Rust: {},
				domain.Zig:  {},
				domain.Scala: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Contractor — Senior Back End Developer (Scala)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4120259237/",
							Date:                 mustDate("2025-01-10"),
							WithSalary:           false,
							Remote:               true,
						},
					},
				},
				domain.Elixir:  {},
				domain.Clojure: {},
				domain.Haskell: {},
			},
			ShortDescription: "American FinTech company",
			Industries: []domain.Industry{
				domain.IndustryFinTech,
			},
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Vivid Money",
			Website: "https://vivid.money/",
			Careers: "https://careers.vivid.money/",
			About:   "https://vivid.money/about-us/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                42385216,
				Alias:             "vividmoney",
				Name:              "Vivid Money",
				Followers:         "42K",
				Employees:         "201-500",
				AssociatedMembers: "234",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "vivid-money",
				Verified: false,
			},
			BlindProfile: domain.BlindProfile{},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "vivid-money",
				Employees: "200",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Vivid-Germany-EI_IE3416158.11,24.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Vivid-Germany-Reviews-E3416158.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Vivid-Germany-Jobs-E3416158.htm",
				Jobs:        "",
				Reviews:     "74",
				Salaries:    "135",
				ReviewsRate: "3.3",
				Verified:    true,
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Backend Developer (Golang)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4127601895/",
							Date:                 mustDate("2025-01-17"),
							WithSalary:           false,
							Remote:               false, // Hybrid or Remote
						},
					},
				},
				domain.Rust:    {},
				domain.Zig:     {},
				domain.Scala:   {},
				domain.Elixir:  {},
				domain.Clojure: {},
				domain.Haskell: {},
			},
			ShortDescription: "Mobile banking",
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "IDEMIA",
			Website: "https://www.idemia.com/",
			Careers: "https://www.idemia.com/careers-idemia",
			About:   "https://www.idemia.com/making-world-safer-place",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                3488,
				Alias:             "idemiagroup",
				Name:              "IDEMIA",
				Followers:         "560K",
				Employees:         "10K+",
				AssociatedMembers: "13,534",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "idemia",
				Verified: false,
			},
			BlindProfile: domain.BlindProfile{},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "idemia",
				Employees: "15,000",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-IDEMIA-EI_IE1835341.11,17.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/IDEMIA-Reviews-E1835341.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/IDEMIA-Jobs-E1835341.htm",
				Jobs:        "123",
				Reviews:     "2.1K",
				Salaries:    "2.8K",
				ReviewsRate: "3.9",
				Verified:    true,
			},
			Languages: domain.Languages{
				domain.Go:   {},
				domain.Rust: {},
				domain.Zig:  {},
				domain.Scala: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Senior Scala Developer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4129086458/",
							Date:                 mustDate("2025-01-19"),
							WithSalary:           false,
							Remote:               false,
						},
					},
				},
				domain.Elixir:  {},
				domain.Clojure: {},
				domain.Haskell: {},
			},
			ShortDescription: "Digital identity solutions",
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "NinjaTrader",
			Website: "https://ninjatrader.com/",
			Careers: "https://ninjatrader.com/careers/",
			About:   "",
			Blog:    "https://ninjacoding.net/ninjatrader/blog",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                2689269,
				Alias:             "ninjatrader-group-llc",
				Name:              "NinjaTrader",
				Followers:         "8K",
				Employees:         "201-500",
				AssociatedMembers: "359",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "",
				Verified: false,
			},
			BlindProfile: domain.BlindProfile{},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "ninjatrader",
				Employees: "180",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-NinjaTrader-EI_IE244011.11,22.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/NinjaTrader-Reviews-E244011.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/NinjaTrader-Jobs-E244011.htm",
				Jobs:        "21",
				Reviews:     "65",
				Salaries:    "157",
				ReviewsRate: "4.3",
				Verified:    true,
			},
			Languages: domain.Languages{
				domain.Go:   {},
				domain.Rust: {},
				domain.Zig:  {},
				domain.Scala: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Staff Software Engineer — Scala",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4124923617/",
							Date:                 mustDate("2025-01-15"),
							WithSalary:           true, // $175K - $205K per year
							Remote:               false,
						},
					},
				},
				domain.Elixir:  {},
				domain.Clojure: {},
				domain.Haskell: {},
			},
			ShortDescription: "Online trading platform",
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Tiko",
			Website: "https://tiko.org/",
			Careers: "https://tiko.org/category/job-listings/",
			About:   "https://tiko.org/about-us/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                5374111,
				Alias:             "tikoafrica",
				Name:              "Tiko",
				Followers:         "15K",
				Employees:         "51-200",
				AssociatedMembers: "314",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "Triggerise",
				Verified: true,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Tiko-EI_IE1300210.11,15.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Tiko-Reviews-E1300210.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Tiko-Jobs-E1300210.htm",
				Jobs:        "",
				Reviews:     "52",
				Salaries:    "44",
				ReviewsRate: "4.0",
				Verified:    true,
			},
			Languages: domain.Languages{
				domain.Go:   {},
				domain.Rust: {},
				domain.Zig:  {},
				domain.Scala: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Scala Backend Engineer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4125716142/",
							Date:                 mustDate("2025-01-14"),
							WithSalary:           false,
							Remote:               true,
						},
					},
				},
				domain.Elixir:  {},
				domain.Clojure: {},
				domain.Haskell: {},
			},
			ShortDescription: "HealthTech",
			Industries: []domain.Industry{
				domain.IndustryHealthTech,
			},
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Topgolf",
			Website: "https://topgolf.com/",
			Careers: "https://careers.topgolf.com/",
			About:   "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                1339209,
				Alias:             "topgolf",
				Name:              "Topgolf",
				Followers:         "126K",
				Employees:         "10K+",
				AssociatedMembers: "11,145",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "topgolf",
				Verified: true,
			},
			BlindProfile: domain.BlindProfile{},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "topgolf",
				Employees: "8,250",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Topgolf-EI_IE510279.11,18.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Topgolf-Reviews-E510279.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Topgolf-Jobs-E510279.htm",
				Jobs:        "576",
				Reviews:     "4.6K",
				Salaries:    "5.3K",
				ReviewsRate: "3.4",
				Verified:    true,
			},
			Languages: domain.Languages{
				domain.Go:   {},
				domain.Rust: {},
				domain.Zig:  {},
				domain.Scala: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Senior Scala Engineer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4121847142/",
							Date:                 mustDate("2025-02-14"), // mustDate("2025-01-13"),
							WithSalary:           false,
							Remote:               false,
						},
					},
				},
				domain.Elixir:  {},
				domain.Clojure: {},
				domain.Haskell: {},
			},
			ShortDescription: "American multinational sports entertainment company",
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Hapag-Lloyd AG",
			Website: "https://www.hapag-lloyd.com/",
			Careers: "https://www.hapag-lloyd.com/en/company/career/overview.html",
			About:   "https://www.hapag-lloyd.com/en/company.html",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                133790,
				Alias:             "hapag-lloyd-ag",
				Name:              "Hapag-Lloyd AG",
				Followers:         "735K",
				Employees:         "10K+",
				AssociatedMembers: "13,990",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "hapag-lloyd",
				Verified: false,
			},
			BlindProfile: domain.BlindProfile{},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "hapag-lloyd",
				Employees: "13,117",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Hapag-Lloyd-EI_IE10450.11,22.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Hapag-Lloyd-Reviews-E10450.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Hapag-Lloyd-Jobs-E10450.htm",
				Jobs:        "1",
				Reviews:     "999",
				Salaries:    "1.5K",
				ReviewsRate: "3.8",
				Verified:    false,
			},
			Languages: domain.Languages{
				domain.Go:   {},
				domain.Rust: {},
				domain.Zig:  {},
				domain.Scala: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Software Engineer Interim (Scala/Java/Kotlin)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4121334590/",
							Date:                 mustDate("2025-01-13"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Software Engineer Interim (Scala/Java/Kotlin)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4149473644/",
							Date:                 mustDate("2025-02-13"),
							WithSalary:           false,
							Remote:               false,
						},
					},
				},
				domain.Elixir:  {},
				domain.Clojure: {},
				domain.Haskell: {},
			},
			ShortDescription: "German international shipping and container transportation company",
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "CV-Library",
			Website: "https://www.cv-library.co.uk/",
			Careers: "https://www.cv-library.co.uk/work-for-us",
			About:   "https://www.cv-library.co.uk/about",
			Blog:    "https://tech-blog.cv-library.co.uk/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                298012,
				Alias:             "cv-library",
				Name:              "CV-Library",
				Followers:         "136K",
				Employees:         "201-500",
				AssociatedMembers: "843",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "cv-library",
				Verified: false,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-CV-Library-EI_IE1076756.11,21.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/CV-Library-Reviews-E1076756.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/CV-Library-Jobs-E1076756.htm",
				Jobs:        "13",
				Reviews:     "143",
				Salaries:    "220",
				ReviewsRate: "3.2",
				Verified:    false,
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Senior Go Developer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4131178111/",
							Date:                 mustDate("2025-01-21"),
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
			ShortDescription: "UK job board",
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "r2p Group",
			Website: "https://www.r2p.com/",
			Careers: "https://www.r2p.com/career/",
			About:   "https://www.r2p.com/about-r2p/company-info/",
			Blog:    "https://r2p.dev/blog/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                11242428,
				Alias:             "r2p-group",
				Name:              "r2p Group",
				Followers:         "2K",
				Employees:         "201-500",
				AssociatedMembers: "180",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "",
				Verified: false,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-R2P-EI_IE3027715.11,14.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/R2P-Reviews-E3027715.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/R2P-Jobs-E3027715.htm",
				Jobs:        "",
				Reviews:     "13",
				Salaries:    "27",
				ReviewsRate: "3.6",
				Verified:    false,
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Embedded Rust Developer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4129472012/",
							Date:                 mustDate("2025-01-21"),
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
			ShortDescription: "Global provider of intelligent technology systems for public transport",
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "nVent",
			Website: "https://www.nvent.com/",
			Careers: "https://www.nvent.com/about-nvent/overview",
			About:   "https://www.nvent.com/careers",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                11418434,
				Alias:             "nvent",
				Name:              "nVent",
				Followers:         "109K",
				Employees:         "10K+",
				AssociatedMembers: "2,972",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "",
				Verified: false,
			},
			BlindProfile: domain.BlindProfile{},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "nvent",
				Employees: "9,900",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-nVent-EI_IE1166699.11,16.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/nVent-Reviews-E1166699.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/nVent-Jobs-E1166699.htm",
				Jobs:        "233",
				Reviews:     "464",
				Salaries:    "815",
				ReviewsRate: "3.7",
				Verified:    true,
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Embedded Rust/C++ Firmware Developer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4075185834/",
							Date:                 mustDate("2025-01-20"),
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
			ShortDescription: "nVent is a leader in enclosures, electric heat-tracing solutions, complete heat-management systems and electrical and fastening solutions",
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Endress+Hauser Group",
			Website: "https://www.endress.com",
			Careers: "https://careers.endress.com/",
			About:   "https://www.endress.com/en/endress-hauser-group",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                11708,
				Alias:             "endress-hauser-group",
				Name:              "Endress+Hauser Group",
				Followers:         "308K",
				Employees:         "10K+",
				AssociatedMembers: "9,669",
				Verified:          true,
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
			Ignore:           true, // Was added by mistake because of the name "Rust Automation & Controls" https://www.linkedin.com/jobs/view/4131415073/
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "PDF Solutions",
			Website: "https://www.pdf.com/",
			Careers: "https://www.pdf.com/company/careers/",
			About:   "https://www.pdf.com/company/overview/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                12658,
				Alias:             "pdf-solutions",
				Name:              "PDF Solutions",
				Followers:         "14K",
				Employees:         "201-500",
				AssociatedMembers: "775",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "",
				Verified: false,
			},
			BlindProfile: domain.BlindProfile{},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "pdf-solutions",
				Employees: "420",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-PDF-Solutions-EI_IE12373.11,24.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/PDF-Solutions-Reviews-E12373.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/PDF-Solutions-Jobs-E12373.htm",
				Jobs:        "11",
				Reviews:     "80",
				Salaries:    "186",
				ReviewsRate: "3.8",
				Verified:    false,
			},
			Languages: domain.Languages{
				domain.Go:   {},
				domain.Rust: {},
				domain.Zig:  {},
				domain.Scala: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Senior Scala/Spark Software Developer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4041068965/",
							Date:                 mustDate("2025-01-23"),
							WithSalary:           true, // CAD $100,000.00 — CAD $130,000.00 per year
							Remote:               true,
						},
						{
							Title:                "Senior Scala Software Developer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4147607601/",
							Date:                 mustDate("2025-02-11"),
							WithSalary:           true, // salary range CAD $100,000 - CAD $150,000
							Remote:               true,
						},
					},
				},
				domain.Elixir:  {},
				domain.Clojure: {},
				domain.Haskell: {},
			},
			// NASDAQ: PDFS
			ShortDescription: "PDF Solutions offers an end-to-end analytics platform that empowers engineers and data scientists across the semiconductor ecosystem to rapidly improve the yield, quality, and profitability of their products",
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Quantexa",
			Website: "https://www.quantexa.com/",
			Careers: "https://www.quantexa.com/careers/",
			About:   "https://www.quantexa.com/about/",
			Blog:    "https://www.quantexa.com/blog/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                10607739,
				Alias:             "quantexa",
				Name:              "Quantexa",
				Followers:         "80K",
				Employees:         "501-1K",
				AssociatedMembers: "794",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "",
				Verified: false,
			},
			BlindProfile: domain.BlindProfile{},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "quantexa",
				Employees: "450",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Quantexa-EI_IE1337455.11,19.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Quantexa-Reviews-E1337455.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Quantexa-Jobs-E1337455.htm",
				Jobs:        "40",
				Reviews:     "153",
				Salaries:    "410",
				ReviewsRate: "4.5",
				Verified:    true,
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
							URL:                  "https://www.linkedin.com/jobs/view/4131108929/",
							Date:                 mustDate("2025-01-22"),
							WithSalary:           false,
							Remote:               false,
						},
					},
				},
				domain.Elixir:  {},
				domain.Clojure: {},
				domain.Haskell: {},
			},
			ShortDescription: "Data and analytics software company",
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "PagoNxt (a Santander company)",
			Website: "https://www.pagonxt.com/",
			Careers: "https://www.pagonxt.com/join-us",
			About:   "https://www.pagonxt.com/who-we-are",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                71200791,
				Alias:             "pagonxt",
				Name:              "PagoNxt (a Santander company)",
				Followers:         "106K",
				Employees:         "1K-5K",
				AssociatedMembers: "1,071",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "pagonxt",
				Verified: false,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-PagoNxt-EI_IE5603235.11,18.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/PagoNxt-Reviews-E5603235.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/PagoNxt-Jobs-E5603235.htm",
				Jobs:        "57",
				Reviews:     "253",
				Salaries:    "471",
				ReviewsRate: "3.4",
				Verified:    true,
			},
			Languages: domain.Languages{
				domain.Go:   {},
				domain.Rust: {},
				domain.Zig:  {},
				domain.Scala: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Software Engineer (Java/Scala)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4132333213/",
							Date:                 mustDate("2025-01-23"),
							WithSalary:           false,
							Remote:               true,
						},
					},
				},
				domain.Elixir:  {},
				domain.Clojure: {},
				domain.Haskell: {},
			},
			ShortDescription: "Global payment solution provider",
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "QAD",
			Website: "https://www.qad.com/",
			Careers: "https://www.qad.com/about/careers",
			About:   "https://www.qad.com/about",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                5099,
				Alias:             "qad",
				Name:              "QAD",
				Followers:         "160K",
				Employees:         "1K-5K",
				AssociatedMembers: "1,707",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "",
				Verified: false,
			},
			BlindProfile: domain.BlindProfile{},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "qad",
				Employees: "2,250",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-QAD-EI_IE6446.11,14.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/QAD-Reviews-E6446.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/QAD-Jobs-E6446.htm",
				Jobs:        "24",
				Reviews:     "627",
				Salaries:    "919",
				ReviewsRate: "4.2",
				Verified:    true,
			},
			Languages: domain.Languages{
				domain.Go:   {},
				domain.Rust: {},
				domain.Zig:  {},
				domain.Scala: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Senior Scala Engineer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4128235975/",
							Date:                 mustDate("2025-02-08"), // mustDate("2025-01-17"),
							WithSalary:           false,
							Remote:               true,
						},
						{
							Title:                "Senior Golang Software Engineer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4139628979/",
							Date:                 mustDate("2025-01-31"),
							WithSalary:           false,
							Remote:               true,
						},
					},
				},
				domain.Elixir:  {},
				domain.Clojure: {},
				domain.Haskell: {},
			},
			ShortDescription: "A provider of next-generation manufacturing and supply chain solutions in the cloud",
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "New Work SE",
			Website: "https://new-work.se/",
			Careers: "https://new-work.se/en/career",
			About:   "https://new-work.se/en/about-new-work-se",
			Blog:    "https://tech.new-work.se/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                28513290,
				Alias:             "new-work-se",
				Name:              "New Work SE",
				Followers:         "25K",
				Employees:         "1K-5K",
				AssociatedMembers: "559",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "new-work",
				Verified: false,
			},
			BlindProfile: domain.BlindProfile{},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "new-work-se",
				Employees: "550",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-New-Work-SE-EI_IE100328.11,22.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/New-Work-SE-Reviews-E100328.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/New-Work-SE-Jobs-E100328.htm",
				Jobs:        "2",
				Reviews:     "342",
				Salaries:    "911",
				ReviewsRate: "3.8",
				Verified:    false,
			},
			Languages: domain.Languages{
				domain.Go:   {},
				domain.Rust: {},
				domain.Zig:  {},
				domain.Scala: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Data Engineer — Scala",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4125983022/",
							Date:                 mustDate("2025-01-16"),
							WithSalary:           true, // 42.000-58.000 EUR
							Remote:               false,
						},
					},
				},
				domain.Elixir:  {},
				domain.Clojure: {},
				domain.Haskell: {},
			},
			ShortDescription: "Germany-based operator of a social network for business professionals",
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "PT. Sarana Abadi Makmur Bersama",
			Website: "https://samb.co.id/",
			Careers: "https://samb.co.id/pages/career",
			About:   "https://samb.co.id/pages/profile",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                66337745,
				Alias:             "sambindonesia",
				Name:              "PT. Sarana Abadi Makmur Bersama",
				Followers:         "3K",
				Employees:         "501-1K",
				AssociatedMembers: "196",
				Verified:          true,
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
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Senior Golang Developer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4134786158/",
							Date:                 mustDate("2025-01-25"),
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
			ShortDescription: "Distribution services, logistics and transportation solutions",
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Walmart",
			Website: "https://walmart.com/",
			Careers: "https://careers.walmart.com/",
			About:   "https://corporate.walmart.com/about",
			Blog:    "https://tech.walmart.com/content/walmart-global-tech/en_us/blog/post.html",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                2646,
				Alias:             "walmart",
				Name:              "Walmart",
				Followers:         "4M",
				Employees:         "10K+",
				AssociatedMembers: "498,289",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "",
				Verified: false,
			},
			BlindProfile: domain.BlindProfile{},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "walmart",
				Employees: "",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Walmart-EI_IE715.11,18.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Walmart-Reviews-E715.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Walmart-Jobs-E715.htm",
				Jobs:        "33K",
				Reviews:     "140K",
				Salaries:    "218K",
				ReviewsRate: "3.4",
				Verified:    true,
			},
			Languages: domain.Languages{
				domain.Go:   {},
				domain.Rust: {},
				domain.Zig:  {},
				domain.Scala: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Data Engineer III (Scala/Spark)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4133956837/",
							Date:                 mustDate("2025-01-25"),
							WithSalary:           true, // $117.000 $234.000 per year
							Remote:               false,
						},
					},
				},
				domain.Elixir:  {},
				domain.Clojure: {},
				domain.Haskell: {},
			},
			ShortDescription: "American multinational retail corporation",
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Magnite",
			Website: "https://www.magnite.com/",
			Careers: "https://www.magnite.com/careers/",
			About:   "https://www.magnite.com/about-us/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                67332446,
				Alias:             "magnite",
				Name:              "Magnite",
				Followers:         "28K",
				Employees:         "501-1K",
				AssociatedMembers: "891",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "MagniteEngineering",
				Verified: false,
			},
			BlindProfile: domain.BlindProfile{},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "magnite",
				Employees: "510",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Magnite-EI_IE4576454.11,18.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Magnite-Reviews-E4576454.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Magnite-Jobs-E4576454.htm",
				Jobs:        "32",
				Reviews:     "113",
				Salaries:    "376",
				ReviewsRate: "3.7",
				Verified:    true,
			},
			Languages: domain.Languages{
				domain.Go:   {},
				domain.Rust: {},
				domain.Zig:  {},
				domain.Scala: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Engineer II, Reporting and Insights, Scala/Spark",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4116388763/",
							Date:                 mustDate("2025-01-24"),
							WithSalary:           true, // $95.000 — $135.000 per year
							Remote:               false,
						},
					},
				},
				domain.Elixir:  {},
				domain.Clojure: {},
				domain.Haskell: {},
			},
			ShortDescription: "American online advertising technology firm",
			Industries: []domain.Industry{
				domain.IndustryAdTech,
			},
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Prezi",
			Website: "https://prezi.com/",
			Careers: "https://prezi.com/jobs/",
			About:   "https://prezi.com/about/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                216295,
				Alias:             "prezi",
				Name:              "Prezi",
				Followers:         "65K",
				Employees:         "201-500",
				AssociatedMembers: "258",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "prezi",
				Verified: false,
			},
			BlindProfile: domain.BlindProfile{},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "prezi",
				Employees: "420",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Prezi-EI_IE453105.11,16.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Prezi-Reviews-E453105.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Prezi-Jobs-E453105.htm",
				Jobs:        "12",
				Reviews:     "180",
				Salaries:    "295",
				ReviewsRate: "3.8",
				Verified:    false,
			},
			Languages: domain.Languages{
				domain.Go:   {},
				domain.Rust: {},
				domain.Zig:  {},
				domain.Scala: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Senior Scala Developer (AI Projects)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4135250887/",
							Date:                 mustDate("2025-01-27"),
							WithSalary:           false,
							Remote:               true,
						},
					},
				},
				domain.Elixir:  {},
				domain.Clojure: {},
				domain.Haskell: {},
			},
			ShortDescription: "Cloud-based presentation software that allows users to create dynamic and interactive presentations.",
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Yuno",
			Website: "https://www.y.uno/",
			Careers: "https://www.y.uno/careers",
			About:   "https://www.y.uno/about",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                79990231,
				Alias:             "yunopay",
				Name:              "Yuno",
				Followers:         "22K",
				Employees:         "51-200",
				AssociatedMembers: "548",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "yuno-payments",
				Verified: false,
			},
			BlindProfile: domain.BlindProfile{},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "yuno",
				Employees: "420",
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
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Backend Go Developer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4139299697/",
							Date:                 mustDate("2025-02-08"), // mustDate("2025-02-05"),
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
			ShortDescription: "Fintech company",
			Industries: []domain.Industry{
				domain.IndustryFinTech,
			},
		},

		{
			ID:      0,                         // system
			Type:    domain.CompanyTypeStartup, // system
			Name:    "Qustodio (part of the Qoria family)",
			Website: "https://www.qustodio.com/",
			Careers: "https://www.qustodio.com/careers/",
			About:   "https://www.qustodio.com/company/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                5229163,
				Alias:             "qustodio",
				Name:              "Qustodio (part of the Qoria family)",
				Followers:         "4K",
				Employees:         "51-200",
				AssociatedMembers: "89",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "qustodio",
				Verified: false,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Qustodio-EI_IE1604075.11,19.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Qustodio-Reviews-E1604075.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Qustodio-Jobs-E1604075.htm",
				Jobs:        "",
				Reviews:     "84",
				Salaries:    "73",
				ReviewsRate: "4.5",
				Verified:    true,
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Senior Software Engineer (Go & Python)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4144099112/",
							Date:                 mustDate("2025-02-05"),
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
			ShortDescription: "Is a parental control app for computers, tablets, and smartphones",
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Pexip",
			Website: "https://www.pexip.com/",
			Careers: "https://www.pexip.com/careers",
			About:   "https://www.pexip.com/about",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                2914650,
				Alias:             "pexip",
				Name:              "Pexip",
				Followers:         "25K",
				Employees:         "201-500",
				AssociatedMembers: "313",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "pexip",
				Verified: true,
			},
			BlindProfile: domain.BlindProfile{},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "pexip",
				Employees: "600",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Pexip-EI_IE2575465.11,16.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Pexip-Reviews-E2575465.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Pexip-Jobs-E2575465.htm",
				Jobs:        "11",
				Reviews:     "79",
				Salaries:    "111",
				ReviewsRate: "3.7",
				Verified:    true,
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Software Engineer, Go/Golang",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4127825676/",
							Date:                 mustDate("2025-02-04"),
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
			ShortDescription: "Is a video communications platform for organizations",
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Vinted",
			Website: "https://www.vinted.com/",
			Careers: "https://careers.vinted.com/",
			About:   "https://www.vinted.com/about",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                2909062,
				Alias:             "vinted",
				Name:              "Vinted",
				Followers:         "120K",
				Employees:         "1K-5K",
				AssociatedMembers: "2,845",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "vinted",
				Verified: true,
			},
			BlindProfile: domain.BlindProfile{},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "vinted",
				Employees: "1,500",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Vinted-EI_IE1028123.11,17.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Vinted-Reviews-E1028123.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Vinted-Jobs-E1028123.htm",
				Jobs:        "12",
				Reviews:     "214",
				Salaries:    "446",
				ReviewsRate: "4.0",
				Verified:    true,
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Middle/Senior Go Backend Engineer",
							ShortDescription:     "Transactional Experience",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4139709944/",
							Date:                 mustDate("2025-02-04"),
							WithSalary:           true, // The salary range for this position is €3,975 — €6,725 per month
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
			ShortDescription: "Marketplace ecosystems",
			Industries: []domain.Industry{
				domain.IndustryECommerce,
			},
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "BigCommerce",
			Website: "https://www.bigcommerce.com/",
			Careers: "https://careers.bigcommerce.com/",
			About:   "https://www.bigcommerce.com/company/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                103315,
				Alias:             "bigcommerce",
				Name:              "BigCommerce",
				Followers:         "144K",
				Employees:         "1K-5K",
				AssociatedMembers: "1,268",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "bigcommerce",
				Verified: true,
			},
			BlindProfile: domain.BlindProfile{},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "bigcommerce",
				Employees: "1,000",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-BigCommerce-EI_IE446630.11,22.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/BigCommerce-Reviews-E446630.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/BigCommerce-Jobs-E446630.htm",
				Jobs:        "28",
				Reviews:     "481",
				Salaries:    "1.2K",
				ReviewsRate: "3.7",
				Verified:    true,
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Senior Software Engineer — Python/Go",
							ShortDescription:     "B2B",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4139297081/",
							Date:                 mustDate("2025-01-31"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Senior Software Engineer — Python/Go",
							ShortDescription:     "B2B",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4148659905/",
							Date:                 mustDate("2025-02-10"),
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
			ShortDescription: "Is a eCommerce platform provides software to businesses that helps them set up and manage online and mobile stores",
			Industries: []domain.Industry{
				domain.IndustryECommerce,
			},
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Grab",
			Website: "https://www.grab.com/",
			Careers: "https://www.grab.careers/",
			About:   "https://www.grab.careers/about-us/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                5382086,
				Alias:             "grabapp",
				Name:              "Grab",
				Followers:         "852K",
				Employees:         "5K-10K",
				AssociatedMembers: "48,818",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "grab",
				Verified: true,
			},
			BlindProfile: domain.BlindProfile{},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "grab",
				Employees: "36,080",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Grab-EI_IE958580.11,15.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Grab-Reviews-E958580.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Grab-Jobs-E958580.htm",
				Jobs:        "10",
				Reviews:     "4.8K",
				Salaries:    "5.9K",
				ReviewsRate: "3.9",
				Verified:    true,
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Senior Software Engineer, Backend (Golang)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4141112561/",
							Date:                 mustDate("2025-02-03"),
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
			ShortDescription: "Developer of app for ride-hailing, food delivery, and digital payment services on mobile devices",
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "A1 Bulgaria",
			Website: "https://www.a1.bg/bg",
			Careers: "https://jobs.a1.com/bg/",
			About:   "https://www.a1.bg/a1-overview-en",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                18367062,
				IDs:               nil,
				Alias:             "a1bulgaria",
				Name:              "A1 Bulgaria",
				Followers:         "27K",
				Employees:         "1K-5K",
				AssociatedMembers: "2,073",
				Verified:          true,
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
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Go Developer",
							ShortDescription:     "Business Process Automation",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4145776890/",
							Date:                 mustDate("2025-02-07"),
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
			ShortDescription: "Is a telecommunications company",
			Ignore:           true, // The company description should be in English
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Qlik",
			Website: "https://www.qlik.com/",
			Careers: "https://www.qlik.com/us/company/careers",
			About:   "https://www.qlik.com/us/why-qlik-is-different",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                10162,
				IDs:               []int{10162, 65485, 18212129},
				Alias:             "qlik",
				Name:              "Qlik",
				Followers:         "257K",
				Employees:         "1K-5K",
				AssociatedMembers: "4,099",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "qlik-oss",
				Verified: false,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Qlik-EI_IE354111.11,15.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Qlik-Reviews-E354111.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Qlik-Jobs-E354111.htm",
				Jobs:        "156",
				Reviews:     "1.2K",
				Salaries:    "1.6K",
				ReviewsRate: "3.9",
				Verified:    true,
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Junior Backend Developer (Java/Go)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4148284507/",
							Date:                 mustDate("2025-02-10"),
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
			ShortDescription: "Company specializing in software development for data integration, analytics, and artificial intelligence.",
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Ryanair",
			Website: "https://www.ryanair.com/",
			Careers: "https://careers.ryanair.com/",
			About:   "https://corporate.ryanair.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                16584,
				IDs:               nil,
				Alias:             "ryanair",
				Name:              "Ryanair - Europe's Favourite Airline",
				Followers:         "812K",
				Employees:         "10K+",
				AssociatedMembers: "19,464",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "",
				Verified: false,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Ryanair-EI_IE6965.11,18.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Ryanair-Reviews-E6965.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Ryanair-Jobs-E6965.htm",
				Jobs:        "54",
				Reviews:     "2.7K",
				Salaries:    "3.8K",
				ReviewsRate: "3.3",
				Verified:    false,
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Graduate Cloud Software Developer (Java/Go)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4145691848/",
							Date:                 mustDate("2025-02-06"),
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
			ShortDescription: "Europe's  largest low cost — airline company",
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Jamf",
			Website: "https://www.jamf.com/",
			Careers: "https://www.jamf.com/about/careers/",
			About:   "https://www.jamf.com/about/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                717074,
				IDs:               nil,
				Alias:             "jamf-software",
				Name:              "Jamf",
				Followers:         "100K",
				Employees:         "1K-5K",
				AssociatedMembers: "2,637",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "",
				Verified: false,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Jamf-EI_IE636547.11,15.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Jamf-Reviews-E636547.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Jamf-Jobs-E636547.htm",
				Jobs:        "39",
				Reviews:     "525",
				Salaries:    "1.2K",
				ReviewsRate: "3.6",
				Verified:    true,
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Software Engineer II (Java/Go)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4144388320/",
							Date:                 mustDate("2025-02-06"),
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
			ShortDescription: "Software company specializing in the management and security of Apple devices",
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "LegalZoom",
			Website: "https://www.legalzoom.com/",
			Careers: "https://www.legalzoom.com/careers",
			About:   "https://www.legalzoom.com/about-us",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                164705,
				Alias:             "legalzoom",
				Name:              "LegalZoom",
				Followers:         "71K",
				Employees:         "501-1K",
				AssociatedMembers: "1,365",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "legalzoom",
				Verified: false,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-LegalZoom-EI_IE270834.11,20.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/LegalZoom-Reviews-E270834.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/LegalZoom-Jobs-E270834.htm",
				Jobs:        "16",
				Reviews:     "835",
				Salaries:    "1.2K",
				ReviewsRate: "3.5",
				Verified:    true,
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Principal Backend Engineer (Golang)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4148714692/",
							Date:                 mustDate("2025-02-12"),
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
			ShortDescription: "Is an American online legal technology and services company",
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "NCR Atleos",
			Website: "https://www.ncratleos.com/",
			Careers: "https://www.ncratleos.com/careers",
			About:   "https://www.ncratleos.com/about-us",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                96148266,
				Alias:             "ncratleos",
				Name:              "NCR Atleos",
				Followers:         "72K",
				Employees:         "10K+",
				AssociatedMembers: "6,905",
				Verified:          true,
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
			Languages: domain.Languages{
				domain.Go: {},
				domain.Rust: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Senior Software Engineer (C++ / Rust)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4148124360/",
							Date:                 mustDate("2025-02-12"),
							WithSalary:           true, // Annual Incentive Base Range 48,000 - 60,000 - 72,000 GBP
							Remote:               false,
						},
						{
							Title:                "Software Engineer (C++ / Rust)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4148121704/",
							Date:                 mustDate("2025-02-12"),
							WithSalary:           true, // Annual Incentive Base Range 36,800 - 46,000 - 55,200 GBP
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
			ShortDescription: "American FinTech company",
			Industries: []domain.Industry{
				domain.IndustryFinTech,
			},
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "KnowBe4",
			Website: "https://www.knowbe4.com/",
			Careers: "https://www.knowbe4.com/careers",
			About:   "https://www.knowbe4.com/about-us",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                2225282,
				Alias:             "knowbe4",
				Name:              "KnowBe4",
				Followers:         "292K",
				Employees:         "1K-5K",
				AssociatedMembers: "2,055",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "knowbe4",
				Verified: false,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-KnowBe4-EI_IE969384.11,18.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/KnowBe4-Reviews-E969384.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/KnowBe4-Jobs-E969384.htm",
				Jobs:        "52",
				Reviews:     "1.4K",
				Salaries:    "1.7K",
				ReviewsRate: "3.6",
				Verified:    true,
			},
			Languages: domain.Languages{
				domain.Go: {},
				domain.Rust: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Senior Software Engineer (Python, C#, Rust)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4150365284/",
							Date:                 mustDate("2025-02-11"),
							WithSalary:           true, // for this position ranges from $130,000-$150,000
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
			ShortDescription: "Cyber security awareness training platform",
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Seldon",
			Website: "https://www.seldon.io/",
			Careers: "https://www.seldon.io/careers",
			About:   "https://www.seldon.io/about",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                2028462,
				Alias:             "seldon",
				Name:              "Seldon",
				Followers:         "11K",
				Employees:         "51-200",
				AssociatedMembers: "99",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "SeldonIO",
				Verified: true,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Seldon-EI_IE2956572.11,17.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Seldon-Reviews-E2956572.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Seldon-Jobs-E2956572.htm",
				Jobs:        "",
				Reviews:     "26",
				Salaries:    "37",
				ReviewsRate: "3.6",
				Verified:    true,
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Software Engineer (Go)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4149751377/",
							Date:                 mustDate("2025-02-13"),
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
			ShortDescription: "British open-source technology company",
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Rockstar Games",
			Website: "https://www.rockstargames.com/",
			Careers: "https://www.rockstargames.com/careers",
			About:   "https://www.rockstargames.com/corpinfo",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                6684,
				Alias:             "rockstar-games",
				Name:              "Rockstar Games",
				Followers:         "694K",
				Employees:         "1K-5K",
				AssociatedMembers: "6,036",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "RockstarGames",
				Verified: true,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Rockstar-Games-EI_IE20887.11,25.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Rockstar-Games-Reviews-E20887.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Rockstar-Games-Jobs-E20887.htm",
				Jobs:        "77",
				Reviews:     "651",
				Salaries:    "1.5K",
				ReviewsRate: "3.8",
				Verified:    false,
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Software Engineer (Go)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4152181957/",
							Date:                 mustDate("2025-02-13"),
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
			ShortDescription: "American game developer and publisher company",
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Kiwi.com",
			Website: "http://kiwi.com/",
			Careers: "https://jobs.kiwi.com/",
			About:   "https://jobs.kiwi.com/about-us/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                3010943,
				Alias:             "kiwi.com",
				Name:              "Kiwi.com",
				Followers:         "27K",
				Employees:         "1K-5K",
				AssociatedMembers: "983",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "kiwicom",
				Verified: false,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Kiwi-com-EI_IE1162174.11,19.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Kiwi-com-Reviews-E1162174.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Kiwi-com-Jobs-E1162174.htm",
				Jobs:        "1",
				Reviews:     "412",
				Salaries:    "615",
				ReviewsRate: "3.6",
				Verified:    true,
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Software Engineer (Golang)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4139671820/",
							Date:                 mustDate("2025-02-14"),
							WithSalary:           true, // Salary range starting from 2.640 EUR per month gross depending on relevant experience and skills
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
			ShortDescription: "Czech online travel agency",
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Tenable",
			Website: "https://www.tenable.com/",
			Careers: "https://www.tenable.com/careers",
			About:   "https://www.tenable.com/about-tenable/about-us",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                25452,
				Alias:             "tenableinc",
				Name:              "Tenable",
				Followers:         "170K",
				Employees:         "1K-5K",
				AssociatedMembers: "2,176",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "tenable",
				Verified: true,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Tenable-EI_IE17494.11,18.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Tenable-Reviews-E17494.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Tenable-Jobs-E17494.htm",
				Jobs:        "11",
				Reviews:     "620",
				Salaries:    "970",
				ReviewsRate: "3.9",
				Verified:    true,
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Staff Software Engineer — Golang",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4076914112/",
							Date:                 mustDate("2025-02-13"),
							WithSalary:           true, // $161,500—$215,500 per year
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
			Name:    "Ramp",
			Website: "https://ramp.com/",
			Careers: "https://ramp.com/careers",
			About:   "https://ramp.com/about-us",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                1406226,
				Alias:             "ramp",
				Name:              "Ramp",
				Followers:         "168K",
				Employees:         "501-1K",
				AssociatedMembers: "1,454",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "",
				Verified: false,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Ramp-EI_IE4211228.11,15.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Ramp-Reviews-E4211228.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Ramp-Jobs-E4211228.htm",
				Jobs:        "",
				Reviews:     "116",
				Salaries:    "234",
				ReviewsRate: "4.2",
				Verified:    true,
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
							Title:                "Senior Software Engineer | Elixir",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4137064133/",
							Date:                 mustDate("2025-01-30"),
							WithSalary:           true, // $188.000 - $258.000 per year
							Remote:               true,
						},
					},
				},
				domain.Clojure: {},
				domain.Haskell: {},
			},
			ShortDescription: "American FinTech company",
			Industries: []domain.Industry{
				domain.IndustryFinTech,
			},
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Hatch",
			Website: "https://www.usehatchapp.com/",
			Careers: "https://www.usehatchapp.com/careers",
			About:   "https://www.usehatchapp.com/why-hatch",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                11433680,
				Alias:             "usehatchapp",
				Name:              "Hatch",
				Followers:         "10K",
				Employees:         "11-50",
				AssociatedMembers: "89",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "Hatch1fy",
				Verified: false,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Hatch-EI_IE2914673.11,16.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Hatch-Reviews-E2914673.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Hatch-Jobs-E2914673.htm",
				Jobs:        "",
				Reviews:     "51",
				Salaries:    "138",
				ReviewsRate: "3.9",
				Verified:    false,
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
							Title:                "Senior Elixir Engineer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4130659582/",
							Date:                 mustDate("2025-01-23"),
							WithSalary:           true, // $140.000 - $200.000 per year
							Remote:               true,
						},
					},
				},
				domain.Clojure: {},
				domain.Haskell: {},
			},
			ShortDescription: "Communication AI platform for businesses",
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
		//					Remote:               false,
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
