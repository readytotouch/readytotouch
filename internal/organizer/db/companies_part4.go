package db

import (
	"github.com/readytotouch/readytotouch/internal/domain"
)

func companiesPart4() []domain.CompanyProfile {
	return []domain.CompanyProfile{
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "limango GmbH — A member of the otto group",
			Website: "https://limango.com/",
			Careers: "https://joinus.limango.com/",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                9381585,
				Alias:             "limango-gmbh",
				Name:              "limango GmbH - A member of the otto group",
				Followers:         "4K",
				Employees:         "201-500",
				AssociatedMembers: "195",
				Verified:          true,
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Senior Go developer",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4096830999/",
							Date:             mustDate("2024-12-13"),
						},
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

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Varonis",
			Website: "https://www.varonis.com/",
			Careers: "https://www.varonis.com/careers",
			About:   "https://www.varonis.com/company",
			Blog:    "https://www.varonis.com/blog",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                25403,
				Alias:             "varonis",
				Name:              "Varonis",
				Followers:         "111K",
				Employees:         "1K-5K",
				AssociatedMembers: "2,433",
				Verified:          true,
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "varonis",
				Employees: "2,250",
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Backend Developer (Python, Go)",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4093559031/",
							Date:             mustDate("2024-12-10"),
						},
						{
							Title:                "Senior Go Engineer",
							ShortDescription:     "At least 5 years experience as a Software Engineer in Go",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4081442058/",
							Date:                 mustDate("2025-01-25"), // mustDate("2025-01-03"),
							WithSalary:           false,
						},
					},
				},
				domain.Rust:    {},
				domain.Zig:     {},
				domain.Scala:   {},
				domain.Elixir:  {},
				domain.Clojure: {},
				domain.Haskell: {},
			},
			Industries: []domain.Industry{
				domain.IndustryCyberSecurity,
			},
			HasEmployeesFromCountries: []domain.Country{},
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "WorldTech IT, LLC",
			Website: "https://wtit.com/",
			Careers: "https://wtit.com/worldtech-it-jobs/",
			About:   "https://wtit.com/about-us/",
			Blog:    "https://wtit.com/blog/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                1811897,
				Alias:             "worldtechit",
				Name:              "WorldTech IT, LLC",
				Followers:         "18K",
				Employees:         "11-50",
				AssociatedMembers: "46",
				Verified:          true,
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Golang Software Engineer",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4101282630/",
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
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Applied Research Solutions",
			Website: "https://www.appliedres.com/",
			Careers: "https://www.appliedres.com/now-hiring",
			About:   "https://www.appliedres.com/meet-the-team",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                6017665,
				Alias:             "applied-research-solutions",
				Name:              "Applied Research Solutions",
				Followers:         "6K",
				Employees:         "501-1K",
				AssociatedMembers: "572",
				Verified:          true,
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "applied-systems",
				Employees: "2,000",
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 11,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Senior Go Developer",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4101198898/",
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
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "WEX",
			Website: "https://www.wexinc.com/",
			Careers: "https://careers.wexinc.com/us/en",
			About:   "https://www.wexinc.com/about/",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                11637,
				Alias:             "wexinc",
				Name:              "WEX",
				Followers:         "92K",
				Employees:         "5K-10K",
				AssociatedMembers: "4,872",
				Verified:          true,
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "wex",
				Employees: "3,500",
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Java/Golang Developer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4048143482/",
							Date:                 mustDate("2025-01-27"), // mustDate("2025-01-10"), // mustDate("2024-12-14"),
							WithSalary:           false,
							Remote:               true,
						},
						{
							Title:                "Golang Developer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4085853366/",
							Date:                 mustDate("2025-02-11"), // mustDate("2025-01-17"),
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
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Comcast",
			Website: "https://corporate.comcast.com/",
			Careers: "https://corporate.comcast.com/careers",
			About:   "https://corporate.comcast.com/company",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                1703,
				Alias:             "comcast",
				Name:              "Comcast",
				Followers:         "648K",
				Employees:         "10K+",
				AssociatedMembers: "59,631",
				Verified:          true,
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "comcast",
				Employees: "164,000",
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 48,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Senior Golang Engineer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4081171024/",
							Date:                 mustDate("2025-02-01"), // mustDate("2024-12-14"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Senior Golang Engineer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4128303283/",
							Date:                 mustDate("2025-01-18"),
							WithSalary:           true, // $132.800 - $208.300 per year
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

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Semrush",
			Website: "https://www.semrush.com/",
			Careers: "https://careers.semrush.com/",
			About:   "https://www.semrush.com/company/",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                2821922,
				Alias:             "semrush",
				Name:              "Semrush",
				Followers:         "502K",
				Employees:         "1K-5K",
				AssociatedMembers: "2,091",
				Verified:          true,
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "semrush",
				Employees: "990",
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 2,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Senior Golang Developer",
							ShortDescription:     "Maroon Team",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4063404057/",
							Date:                 mustDate("2025-02-01"), // mustDate("2024-12-13"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Golang Developer",
							ShortDescription:     "Go as a primary programming language",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4077464276/",
							Date:                 mustDate("2025-01-03"),
							WithSalary:           false,
							Remote:               true,
						},
						{
							Title:                "Golang Developer",
							ShortDescription:     "Enterprise Solutions Unit",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4132135610/",
							Date:                 mustDate("2025-01-23"),
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
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Pi Square Technologies",
			Website: "https://www.pisquaretech.com/",
			Careers: "https://www.pisquaretech.com/careers",
			About:   "https://www.pisquaretech.com/whoweare",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                10589520,
				Alias:             "pi-square-technology",
				Name:              "Pi Square Technologies",
				Followers:         "45K",
				Employees:         "201-500",
				AssociatedMembers: "226",
				Verified:          true,
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Senior Golang Developer",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4098692466/",
							Date:             mustDate("2024-12-13"),
						},
						{
							Title:            "Python Developer",
							ShortDescription: "with Pen-testing & Golang experience",
							URL:              "https://www.linkedin.com/jobs/view/4132052658/",
							Date:             mustDate("2025-01-23"),
						},
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

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "DataOn (PT Indodev Niaga Internet)",
			Website: "https://dataon.com/",
			Careers: "https://dataon.com/careers/",
			About:   "https://dataon.com/about-us/",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                844561,
				Alias:             "dataoncorp",
				Name:              "DataOn (PT Indodev Niaga Internet)",
				Followers:         "6K",
				Employees:         "201-500",
				AssociatedMembers: "468",
				Verified:          true,
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Senior Golang Developer",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4099666536/",
							Date:             mustDate("2024-12-13"),
						},
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

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Kiteworks",
			Website: "https://www.kiteworks.com/",
			Careers: "https://www.kiteworks.com/company/careers/",
			About:   "https://www.kiteworks.com/company/",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                33337,
				Alias:             "kiteworkscgcp",
				Name:              "Kiteworks",
				Followers:         "25K",
				Employees:         "201-500",
				AssociatedMembers: "263",
				Verified:          true,
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "kiteworks",
				Employees: "351",
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Golang Developer",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4096373782/",
							Date:             mustDate("2024-12-11"),
						},
						{
							Title:                "Golang Developer / Backend Developer (Golang)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4123639315/",
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
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Centripetal",
			Website: "https://www.centripetal.ai/",
			Careers: "https://www.centripetal.ai/careers/",
			About:   "https://www.centripetal.ai/about/",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                1321511,
				Alias:             "centripetal-ai",
				Name:              "Centripetal",
				Followers:         "4K",
				Employees:         "51-200",
				AssociatedMembers: "102",
				Verified:          true,
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Senior Golang Software Engineer",
							ShortDescription: "CIS Data Services (IE)",
							URL:              "https://www.linkedin.com/jobs/view/4094769985/",
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
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Technology Innovation Institute",
			Website: "https://www.tii.ae/",
			Careers: "https://www.tii.ae/careers",
			About:   "https://www.tii.ae/about-us",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                28729816,
				Alias:             "tiiuae",
				Name:              "Technology Innovation Institute",
				Followers:         "162K",
				Employees:         "1K-5K",
				AssociatedMembers: "1,122",
				Verified:          true,
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "technology-innovation-institute",
				Employees: "690",
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 16,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Senior Golang Developer",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4091567553/",
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
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "BAUHAUS Deutschland",
			Website: "https://bauhaus.info/",
			Careers: "https://jobs.bauhaus.info/",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                9417424,
				Alias:             "bauhaus-deutschland",
				Name:              "BAUHAUS Deutschland",
				Followers:         "15K",
				Employees:         "10K+",
				AssociatedMembers: "1,633",
				Verified:          true,
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Senior Golang Engineer",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4049064454/",
							Date:             mustDate("2025-01-26"), // mustDate("2024-12-10"),
						},
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

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Sun* Vietnam",
			Website: "https://sun-asterisk.vn/en/",
			Careers: "https://sun-asterisk.vn/en/jobs-en/",
			About:   "https://sun-asterisk.vn/en/about-us/",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                14571208,
				Alias:             "sunasterisk",
				Name:              "Sun* Vietnam",
				Followers:         "11K",
				Employees:         "1K-5K",
				AssociatedMembers: "843",
				Verified:          true,
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 4,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Golang Engineer",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4090103893/",
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
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Miracle Software Systems, Inc",
			Website: "https://www.miraclesoft.com/",
			Careers: "https://careers.miraclesoft.com/",
			About:   "https://www.miraclesoft.com/company/",
			Blog:    "https://blog.miraclesoft.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                15388,
				Alias:             "miraclesoft",
				Name:              "Miracle Software Systems, Inc",
				Followers:         "115K",
				Employees:         "1K-5K",
				AssociatedMembers: "2,581",
				Verified:          true,
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "miracle-software-systems",
				Employees: "2,500",
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Golang Developer",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4090375762/",
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
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Collabera",
			Website: "https://www.collaberadigital.com/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                24440,
				Alias:             "collabera",
				Name:              "Collabera",
				Followers:         "1M",
				Employees:         "10K+",
				AssociatedMembers: "6,048",
				Verified:          true,
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "collabera",
				Employees: "6,250",
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Java Developer",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4087446589/",
							Date:             mustDate("2024-12-03"),
						},
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
							Title:            "Senior Golang Developer",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4049853199/",
							Date:             mustDate("2025-01-14"), // mustDate("2024-12-03")
							WithSalary:       false,
							Remote:           true,
						},
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
			ID:      0,  // system
			Type:    "", // system
			Name:    "SMG Swiss Marketplace Group",
			Website: "https://swissmarketplace.group/",
			Careers: "https://swissmarketplace.group/career/",
			About:   "https://swissmarketplace.group/about/",
			Blog:    "https://swissmarketplace.group/tech-blog/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                76994292,
				IDs:               []int{14415749, 76994292, 103188619},
				Alias:             "smg-marketplace",
				Name:              "SMG Swiss Marketplace Group",
				Followers:         "24K",
				Employees:         "501-1K",
				AssociatedMembers: "636",
				Verified:          true,
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Senior Golang Engineer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4016921964/",
							Date:                 mustDate("2025-01-16"), // mustDate("2024-12-03"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Senior Golang Engineer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4145034904/",
							Date:                 mustDate("2025-02-06"),
							WithSalary:           false,
							Remote:               true,
						},
						{
							Title:                "Senior Golang Engineer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4149648184/",
							Date:                 mustDate("2025-02-11"),
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
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Hornetsecurity",
			Website: "https://www.hornetsecurity.com/",
			Careers: "https://www.hornetsecurity.com/career/",
			About:   "https://www.hornetsecurity.com/company/",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                750470,
				IDs:               nil,
				Alias:             "hornetsecurity",
				Name:              "Hornetsecurity",
				Followers:         "20K",
				Employees:         "501-1K",
				AssociatedMembers: "414",
				Verified:          true,
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Golang Software Engineer",
							ShortDescription:     "Security Lab",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4089330917/",
							Date:                 mustDate("2024-12-03"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Golang Software Engineer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4149649448/",
							Date:                 mustDate("2025-02-11"),
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

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Gorilla Logic",
			Website: "https://www.gorillalogic.com/",
			Careers: "https://gorillalogic.com/careers",
			About:   "https://www.gorillalogic.com/who-we-are",
			Blog:    "https://www.gorillalogic.com/blog-and-resources?contentType=Blog",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                68692,
				Alias:             "gorillalogic",
				Name:              "Gorilla Logic",
				Followers:         "41K",
				Employees:         "501-1K",
				AssociatedMembers: "524",
				Verified:          true,
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 3,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Senior Go Engineer",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4085414974/",
							Date:             mustDate("2024-11-26"),
						},
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

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "ZeptoLab",
			Website: "https://www.zeptolab.com/",
			Careers: "https://careers.zeptolab.com/",
			About:   "https://www.zeptolab.com/about",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                1945804,
				Alias:             "zeptolab",
				Name:              "ZeptoLab",
				Followers:         "32K",
				Employees:         "51-200",
				AssociatedMembers: "143",
				Verified:          true,
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Senior Golang Developer",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4102533913/",
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
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Getir",
			Website: "https://getir.com/",
			Careers: "https://career.getir.com/",
			About:   "https://getir.com/en/about/",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                9866841,
				Alias:             "getir",
				Name:              "Getir",
				Followers:         "415K",
				Employees:         "1K-5K",
				AssociatedMembers: "6,610",
				Verified:          true,
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "getir",
				Employees: "27,520",
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Go Engineer",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4050960562/",
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
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Happening",
			Website: "https://www.happening.xyz/",
			Careers: "https://www.happening.xyz/careers",
			About:   "https://www.happening.xyz/about-us",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                86471246,
				Alias:             "happen-ing",
				Name:              "Happening",
				Followers:         "7K",
				Employees:         "501-1K",
				AssociatedMembers: "514",
				Verified:          true,
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Senior Backend Engineer (Go)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4102013036/",
							Date:                 mustDate("2025-02-07"), // mustDate("2025-02-05"), // mustDate("2025-01-20"), // mustDate("2024-12-16"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Senior Golang Engineer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4130508027/",
							Date:                 mustDate("2025-02-05"), // mustDate("2025-01-20"),
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
							Title:            "Backend Engineer (Elixir)",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/3884244005/",
							Date:             mustDate("2025-01-20"), // mustDate("2024-11-29"),
						},
						{
							Title:                "Senior Backend Engineer (Elixir)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4103278708/",
							Date:                 mustDate("2025-01-20"), // mustDate("2025-01-17"),
							WithSalary:           false,
							Remote:               false,
						},
					},
				},
				domain.Clojure: {},
				domain.Haskell: {},
			},
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Zeller",
			Website: "https://www.myzeller.com/",
			Careers: "https://www.myzeller.com/careers",
			About:   "https://www.myzeller.com/about",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                12687202,
				Alias:             "zeller",
				Name:              "Zeller",
				Followers:         "16K",
				Employees:         "201-500",
				AssociatedMembers: "242",
				Verified:          true,
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Senior Engineer (Go/Typescript + AWS)",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4047579348/",
							Date:             mustDate("2024-12-13"),
						},
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

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Flip",
			Website: "https://flip.id/",
			Careers: "https://career.flip.id/jobs",
			About:   "https://flip.id/en/tentang-flip",
			Blog:    "https://tech.flip.id/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                13234824,
				Alias:             "flip.id",
				Name:              "Flip",
				Followers:         "222K",
				Employees:         "201-500",
				AssociatedMembers: "610",
				Verified:          true,
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Software Engineer (Backend-Go) — Risk Platform",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4096854844/",
							Date:             mustDate("2024-12-13"),
						},
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

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "OpenX",
			Website: "https://www.openx.com/",
			Careers: "https://www.openx.com/careers/",
			About:   "https://www.openx.com/company/",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                42890,
				Alias:             "openx",
				Name:              "OpenX",
				Followers:         "30K",
				Employees:         "201-500",
				AssociatedMembers: "440",
				Verified:          true,
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "openx",
				Employees: "210",
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 4,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Senior Software Engineer IV (Java/Go)",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4062276145/",
							Date:             mustDate("2025-02-04"), // mustDate("2024-12-12"),
						},
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
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Lenskart.com",
			Website: "https://lenskart.com/",
			Careers: "https://hiring.lenskart.com/",
			About:   "https://www.lenskart.com/about-us.html",
			Blog:    "https://lenskart.medium.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                213339,
				Alias:             "lenskart-com",
				Name:              "Lenskart.com",
				Followers:         "397K",
				Employees:         "10K+",
				AssociatedMembers: "10,127",
				Verified:          true,
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "lenskart",
				Employees: "500",
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Software Engineer, Golang",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4096376928/",
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
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Lowe's Companies, Inc.",
			Website: "https://lowes.com/",
			Careers: "https://talent.lowes.com/us/en",
			About:   "https://www.lowes.com/l/about",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				// LinkedIn alias that is hard to encode
				ID:                4128,
				Alias:             "4128", // "lowe's-home-improvement",
				Name:              "Lowe's Companies, Inc.",
				Followers:         "756K",
				Employees:         "10K+",
				AssociatedMembers: "118,402",
				Verified:          true,
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "lowes",
				Employees: "118,840",
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Senior Software Engineer",
							ShortDescription:     "Mobile/Backend (Android, Go, Kotlin)",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4093511115/",
							Date:                 mustDate("2025-01-14"), // mustDate("2024-12-11"),
							WithSalary:           true,                   // $92,300.00 — $175,400.00
							Remote:               false,
						},
						{
							Title:                "Senior Software Engineer",
							ShortDescription:     "Mobile/Backend (Android, Go, Kotlin)",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4144331556/",
							Date:                 mustDate("2025-02-05"),
							WithSalary:           true, // $95.100 — $180,700 per year
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
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Huntress",
			Website: "https://www.huntress.com/",
			Careers: "https://www.huntress.com/company/careers",
			About:   "https://www.huntress.com/company/our-company",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                10172550,
				Alias:             "huntress-labs",
				Name:              "Huntress",
				Followers:         "71K",
				Employees:         "201-500",
				AssociatedMembers: "542",
				Verified:          true,
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "huntress",
				Employees: "150",
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 5,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Senior Software Engineer",
							ShortDescription: "Windows SIEM Agent (Go)",
							URL:              "https://www.linkedin.com/jobs/view/4055142455/",
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
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "1NCE",
			Website: "https://1nce.com/",
			Careers: "https://1nce.com/en-eu/about/careers",
			About:   "https://1nce.com/en-eu/about/",
			Blog:    "https://1nce.com/en-us/resources/news/blog",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                18471615,
				Alias:             "1nce",
				Name:              "1NCE",
				Followers:         "10K",
				Employees:         "201-500",
				AssociatedMembers: "289",
				Verified:          true,
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 4,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Senior Software Engineer Go",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4088250312/",
							Date:             mustDate("2024-12-03"),
						},
						{
							Title:                "Software Engineer Go (Middle/Senior)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4125950329/",
							Date:                 mustDate("2025-01-16"),
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

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Waystar",
			Website: "https://www.waystar.com/",
			Careers: "https://careers.waystar.com/",
			About:   "https://www.waystar.com/about-us/",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                11444741,
				Alias:             "waystar",
				Name:              "Waystar",
				Followers:         "21K",
				Employees:         "1K-5K",
				AssociatedMembers: "1,457",
				Verified:          true,
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "waystar",
				Employees: "1,040",
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Advanced Application Engineer (PHP / Go)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4089563703/",
							Date:                 mustDate("2024-12-03"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Application Engineer (PHP, Golang)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/3871186016/",
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
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "efood",
			Website: "https://www.e-food.gr/",
			Careers: "https://www.e-food.gr/blog/efood-is-seeking-on-foot-delivery-team-members/",
			About:   "https://www.e-food.gr/page/who",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                2410959,
				Alias:             "e-food-gr",
				Name:              "efood",
				Followers:         "78K",
				Employees:         "1K-5K",
				AssociatedMembers: "1,190",
				Verified:          true,
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Senior Software Engineer (Go)",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/3873641135/",
							Date:             mustDate("2024-12-03"),
						},
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

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "ESET",
			Website: "https://www.eset.com/",
			Careers: "https://www.eset.com/int/about/careers/",
			About:   "https://www.eset.com/us/about/",
			Blog:    "https://www.eset.com/blog/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                28967,
				Alias:             "eset",
				Name:              "ESET",
				Followers:         "77K",
				Employees:         "1K-5K",
				AssociatedMembers: "1,838",
				Verified:          true,
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "eset",
				Employees: "1,750",
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 2,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Software Engineer",
							ShortDescription: "Cloud Security Applications (.NET Core or Go)",
							URL:              "https://www.linkedin.com/jobs/view/4084785751/",
							Date:             mustDate("2024-12-03"),
						},
						{
							Title:                "Software Engineer — Cloud Security Applications (.NET Core or Go)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4122199318/",
							Date:                 mustDate("2025-01-14"),
							WithSalary:           true, // from 2 600 EUR
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

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Praxent",
			Website: "https://praxent.com/",
			Careers: "https://praxent.com/careers",
			About:   "https://praxent.com/overview",
			Blog:    "https://praxent.com/blog",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                1365857,
				Alias:             "praxent",
				Name:              "Praxent",
				Followers:         "5K",
				Employees:         "51-200",
				AssociatedMembers: "116",
				Verified:          true,
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "praxent",
				Employees: "70",
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:            "LATAM Senior Backend Engineer (Python or Go)",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4085449585/",
							Date:             mustDate("2024-11-27"),
						},
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

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Stacklok",
			Website: "https://stacklok.com/",
			Careers: "https://stacklok.com/careers",
			About:   "https://stacklok.com/about",
			Blog:    "https://stacklok.com/blog",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                94101540,
				Alias:             "stacklok",
				Name:              "Stacklok",
				Followers:         "2K",
				Employees:         "11-50",
				AssociatedMembers: "39",
				Verified:          true,
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 5,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Software Engineer II (Go & Python)",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4079391812/",
							Date:             mustDate("2024-11-20"),
						},
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

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "PT Lion Super Indo",
			Website: "https://www.superindo.co.id/",
			Careers: "https://superindo.jobseeker.software/",
			About:   "https://www.superindo.co.id/korporasi-keberlanjutan/corporate/about_us",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                1104467,
				Alias:             "superindo",
				Name:              "PT Lion Super Indo",
				Followers:         "630K",
				Employees:         "10K+",
				AssociatedMembers: "13,961",
				Verified:          true,
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Backend Developer (Golang)",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4102186636/",
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
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Ubiquiti Inc.",
			Website: "https://www.ui.com/",
			Careers: "https://careers.ui.com/",
			About:   "https://www.ui.com/introduction",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                109341,
				Alias:             "ubiquiti-",
				Name:              "Ubiquiti Inc.",
				Followers:         "103K",
				Employees:         "1K-5K",
				AssociatedMembers: "1,320",
				Verified:          true,
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "ubiquiti",
				Employees: "930",
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 4,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Backend Developer (Golang)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4100868953/",
							Date:                 mustDate("2025-01-17"), // mustDate("2024-12-16"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Cloud Software Developer (Go)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4144403457/",
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
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "TextNow",
			Website: "https://textnow.com/",
			Careers: "https://careers.textnow.com/careers",
			About:   "https://careers.textnow.com/",
			Blog:    "https://medium.com/textnowengineering",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                529693,
				Alias:             "enflick-inc-",
				Name:              "TextNow",
				Followers:         "21K",
				Employees:         "51-200",
				AssociatedMembers: "249",
				Verified:          true,
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "textnow",
				Employees: "180",
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 18,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Golang Backend Developer",
							ShortDescription:     "Trust & Safety",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4096314146/",
							Date:                 mustDate("2025-01-27"), // mustDate("2024-12-11"),
							WithSalary:           false,
							Remote:               true,
						},
						{
							Title:                "Golang Backend Developer",
							ShortDescription:     "Maintain and refactor backend systems to improve our existing detections in languages like Golang, PHP, and Python",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4096309405/",
							Date:                 mustDate("2025-01-02"),
							WithSalary:           false,
							Remote:               true,
						},
						{
							Title:                "Senior Backend Software Developer (Golang)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4124814543/",
							Date:                 mustDate("2025-01-15"),
							WithSalary:           false,
							Remote:               true,
						},
						{
							Title:                "Golang Backend Developer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4075845150/",
							Date:                 mustDate("2025-02-05"),
							WithSalary:           true, // $96,840 — $132,420 per year
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
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Altenar",
			Website: "https://altenar.com/",
			Careers: "https://altenar.com/career/",
			About:   "https://altenar.com/about/",
			Blog:    "https://altenar.com/blog/technology/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                11063103,
				Alias:             "altenar",
				Name:              "Altenar",
				Followers:         "16K",
				Employees:         "501-1K",
				AssociatedMembers: "299",
				Verified:          true,
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Senior Golang Software Developer",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4006862895/",
							Date:             mustDate("2024-12-10"),
						},
						{
							Title:                "Senior Golang Software Developer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4127518517/",
							Date:                 mustDate("2025-01-16"),
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
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "VUS — The English Center",
			Website: "https://vus.edu.vn/",
			Careers: "",
			About:   "https://vus.edu.vn/gioi-thieu-vus",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                33233160,
				IDs:               nil,
				Alias:             "vustheenglishcenter",
				Name:              "VUS - The English Center",
				Followers:         "14K",
				Employees:         "1K-5K",
				AssociatedMembers: "1,992",
				Verified:          true,
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Senior Backend Developer (Golang)",
							ShortDescription:     "3 years exp",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4089423781/",
							Date:                 mustDate("2024-12-10"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Senior Backend Developer — Golang",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4148839012/",
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
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Exness",
			Website: "https://www.exness.uk/",
			Careers: "https://exness-careers.com/",
			About:   "https://www.exness.uk/about-us/",
			Blog:    "https://medium.com/exness-blog",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                3291356,
				Alias:             "exness",
				Name:              "Exness",
				Followers:         "92K",
				Employees:         "1K-5K",
				AssociatedMembers: "4,664",
				Verified:          true,
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "exness",
				Employees: "3,000",
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Golang/Python Developer",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4088177442/",
							Date:             mustDate("2024-12-03"),
						},
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

		{
			ID:   0,  // system
			Type: "", // system
			Name: "Xsolla",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                875431,
				Alias:             "xsolla",
				Name:              "Xsolla",
				Followers:         "32K",
				Employees:         "501-1K",
				AssociatedMembers: "944",
				Verified:          true,
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 3,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Golang Backend Developer",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4069628874/",
							Date:             mustDate("2025-01-15"), // mustDate("2024-12-03"),
						},
					},
				},
				domain.Rust:    {},
				domain.Zig:     {},
				domain.Scala:   {},
				domain.Elixir:  {},
				domain.Clojure: {},
				domain.Haskell: {},
			},
			Ignore: true, // https://gamedev.dou.ua/forums/topic/50138/
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "emerchantpay",
			Website: "https://www.emerchantpay.com/",
			Careers: "https://www.emerchantpay.com/careers/",
			About:   "https://www.emerchantpay.com/about-us/",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                1432812,
				Alias:             "emerchantpay-ltd",
				Name:              "emerchantpay",
				Followers:         "32K",
				Employees:         "201-500",
				AssociatedMembers: "408",
				Verified:          true,
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Senior Developer with Golang",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4084138762/",
							Date:             mustDate("2024-12-03"),
						},
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

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Power Factors",
			Website: "https://www.powerfactors.com/",
			Careers: "https://www.powerfactors.com/careers",
			About:   "https://www.powerfactors.com/about",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                4748453,
				Alias:             "power-factors-inc",
				Name:              "Power Factors",
				Followers:         "15K",
				Employees:         "501-1K",
				AssociatedMembers: "586",
				Verified:          true,
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "power-factors",
				Employees: "751",
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Senior Engineer, Back-End, Golang (Greece)",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4102953810/",
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
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Airspace",
			Website: "https://www.airspace.com/",
			Careers: "https://www.airspace.com/careers",
			About:   "https://www.airspace.com/company/who-we-are",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                10504949,
				Alias:             "airspace-inc",
				Name:              "Airspace",
				Followers:         "9K",
				Employees:         "201-500",
				AssociatedMembers: "360",
				Verified:          true,
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "airspace",
				Employees: "351",
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Senior Software Engineer — Golang",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4101213697/",
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
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Amartha",
			Website: "https://amartha.com/",
			Careers: "https://amartha.com/en/career/",
			About:   "https://amartha.com/en/about-us/",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                3655032,
				Alias:             "amartha",
				Name:              "Amartha",
				Followers:         "123K",
				Employees:         "5K-10K",
				AssociatedMembers: "2,197",
				Verified:          true,
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Backend Engineer (Golang) — Senior/Principal Level",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4102510260/",
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
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Creative Software",
			Website: "http://www.creativesoftware.com/",
			Careers: "https://www.creativesoftware.com/careers",
			About:   "https://www.creativesoftware.com/aboutus",
			Blog:    "https://www.creativesoftware.com/blog-categories/tech-insights",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                12688,
				Alias:             "creativesoftware",
				Name:              "Creative Software",
				Followers:         "61K",
				Employees:         "201-500",
				AssociatedMembers: "531",
				Verified:          true,
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Senior Software Engineer(Golang)",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4101950317/",
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
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Cisco ThousandEyes",
			Website: "https://www.thousandeyes.com/",
			Careers: "https://www.thousandeyes.com/careers/",
			About:   "https://www.thousandeyes.com/about/",
			Blog:    "https://www.thousandeyes.com/blog/?cat=thousandeyes-engineering",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                687352,
				Alias:             "thousandeyes",
				Name:              "Cisco ThousandEyes",
				Followers:         "62K",
				Employees:         "1K-5K",
				AssociatedMembers: "959",
				Verified:          true,
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "thousandeyes",
				Employees: "840",
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 9,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Senior Site Reliability Engineer I (Python/Golang)",
							ShortDescription:     "Agent Ops",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4100798671/",
							Date:                 mustDate("2025-02-04"), // mustDate("2024-12-14"),
							WithSalary:           true,                   // $146.600 — $203.100 per year
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

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "U.S. Bank",
			Website: "https://usbank.com/",
			Careers: "https://careers.usbank.com/",
			About:   "https://www.usbank.com/about-us-bank.html",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                2532,
				Alias:             "us-bank",
				Name:              "U.S. Bank",
				Followers:         "499K",
				Employees:         "10K+",
				AssociatedMembers: "75,830",
				Verified:          true,
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "us-bank",
				Employees: "64,510",
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Software Information Security Engineer (Golang or Python)",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4063303097/",
							Date:             mustDate("2024-12-12"),
						},
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

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "NetApp",
			Website: "https://www.netapp.com/",
			Careers: "https://careers.netapp.com/",
			About:   "https://www.netapp.com/company/",
			Blog:    "https://community.netapp.com/t5/Blogs/ct-p/netapp-blogs",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                2105,
				Alias:             "netapp",
				Name:              "NetApp",
				Followers:         "709K",
				Employees:         "10K+",
				AssociatedMembers: "12,935",
				Verified:          true,
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "netapp",
				Employees: "11,000",
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 23,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Senior Software Engineer — Golang",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4095534140/",
							Date:                 mustDate("2024-12-12"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Senior Software Engineer — Golang",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4094529149/",
							Date:                 mustDate("2025-01-16"),
							WithSalary:           true, // $161.100 - $226.500 per year
							Remote:               false,
						},
						{
							Title:                "Software Engineer — Golang",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4132416977/",
							Date:                 mustDate("2025-01-23"),
							WithSalary:           true, // $138.780 — $195.030 per year
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

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Trend Micro",
			Website: "https://www.trendmicro.com/",
			Careers: "https://www.trendmicro.com/en_us/about/careers.html",
			About:   "https://www.trendmicro.com/en_us/about.html",
			Blog:    "https://www.trendmicro.com/en_us/research.html",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                4312,
				Alias:             "trend-micro",
				Name:              "Trend Micro",
				Followers:         "269K",
				Employees:         "5K-10K",
				AssociatedMembers: "7,565",
				Verified:          true,
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "trend-micro",
				Employees: "7,000",
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 4,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Cloud Engineer (Golang/Python, Backend Focus)",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4063567336/",
							Date:             mustDate("2024-12-11"), // mustDate("2024-12-11"),
						},
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

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Central Retail",
			Website: "https://www.centralretail.com/",
			Careers: "https://www.centralretail.com/en/careers/life-at-central-retail",
			About:   "https://www.centralretail.com/en/who-we-are/our-purpose",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                26628492,
				Alias:             "centralretail",
				Name:              "Central Retail",
				Followers:         "67K",
				Employees:         "10K+",
				AssociatedMembers: "1,066",
				Verified:          true,
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Senior Software Engineer (Golang)",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4093163873/",
							Date:             mustDate("2024-12-10"),
						},
						{
							Title:                "Software Engineer (Golang)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4133877020/",
							Date:                 mustDate("2025-01-24"),
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

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Kargo",
			Website: "https://www.kargo.com/",
			Careers: "https://www.kargo.com/careers",
			About:   "https://www.kargo.com/about",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                55830,
				Alias:             "kargo",
				Name:              "Kargo",
				Followers:         "37K",
				Employees:         "201-500",
				AssociatedMembers: "980",
				Verified:          true,
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "kargo",
				Employees: "67",
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Senior Software Engineer (Golang)",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4089015041/",
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
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Salesloft",
			Website: "https://www.salesloft.com/",
			Careers: "https://www.salesloft.com/company/careers",
			About:   "https://www.salesloft.com/company",
			Blog:    "https://www.salesloft.com/resources/blog",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                2296178,
				Alias:             "salesloft",
				Name:              "Salesloft",
				Followers:         "110K",
				Employees:         "501-1K",
				AssociatedMembers: "1,253",
				Verified:          true,
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "salesloft",
				Employees: "740",
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 3,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Software Engineer (Golang)",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4072522353/",
							Date:             mustDate("2024-12-03"),
						},
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

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "CJ MORE",
			Website: "https://www.cjexpress.co.th/",
			Careers: "https://www.cjexpress.co.th/career",
			About:   "https://www.cjexpress.co.th/about",
			Blog:    "https://medium.com/cj-express-tech-tildi",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                10081692,
				Alias:             "cj-express-group",
				Name:              "CJ MORE",
				Followers:         "35K",
				Employees:         "501-1K",
				AssociatedMembers: "669",
				Verified:          true,
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Software Engineer (Golang)",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4088241281/",
							Date:             mustDate("2024-12-03"),
						},
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

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "H-E-B",
			Website: "https://www.heb.com/",
			Careers: "https://careers.heb.com/",
			About:   "https://careers.heb.com/about-h-e-b/",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                164159,
				Alias:             "heb",
				Name:              "H-E-B",
				Followers:         "273K",
				Employees:         "10K+",
				AssociatedMembers: "42,800",
				Verified:          true,
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "h-e-b",
				Employees: "23,870",
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Senior Software Engineer, Golang",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4000104862/",
							Date:             mustDate("2024-12-03"),
						},
					},
				},
				domain.Rust: {},
				domain.Zig:  {},
				domain.Scala: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Senior Software Engineer — Scala/Java",
							ShortDescription:     "Order Management System",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4113015625/",
							Date:                 mustDate("2025-01-24"),
							WithSalary:           false,
							Remote:               false,
						},
					},
				},
				domain.Elixir:  {},
				domain.Clojure: {},
				domain.Haskell: {},
			},
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
