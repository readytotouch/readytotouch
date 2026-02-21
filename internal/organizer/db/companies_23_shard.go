package db

import "github.com/readytotouch/readytotouch/internal/domain"

func companies23Shard() []domain.CompanyProfile {
	return []domain.CompanyProfile{
		{
			ID:         0,  // system
			Type:       "", // system
			Name:       "Alma Security",
			BaseURL:    "https://alma.security/",
			CareersURL: "https://alma.security/Career",
			AboutURL:   "https://alma.security/About-us",
			BlogURL:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                100717639,
				IDs:               nil,
				Alias:             "alma-security",
				Name:              "Alma Security",
				Followers:         "3K",
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
					GitHubRepositoryCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Backend Developer, Golang",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4350210925/",
							Location:             "Tel Aviv District, Israel",
							Date:                 mustDate("2025-12-15"),
							WithSalary:           false,
							Remote:               false,
						},
					},
				},
				domain.Rust: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Zig: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Scala: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Elixir: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Clojure: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Haskell: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Erlang: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
			},
			ShortDescription: "",
			Industries: []domain.Industry{
				domain.IndustryCyberSecurity,
			},
		},
		{
			ID:         0,  // system
			Type:       "", // system
			Name:       "ScorePlay",
			BaseURL:    "https://www.scoreplay.io/",
			CareersURL: "https://www.scoreplay.io/careers",
			AboutURL:   "https://www.scoreplay.io/company",
			BlogURL:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                13009905,
				IDs:               nil,
				Alias:             "scoreplay",
				Name:              "ScorePlay",
				Followers:         "12K",
				Employees:         "51-200",
				AssociatedMembers: "61",
				Verified:          false,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:     "ScorePlay-Inc",
				Followers: "12",
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
					GitHubRepositoryCount: 1,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Senior Back-end Engineer (GO)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4343636595/",
							Location:             "Paris, Île-de-France, France",
							Date:                 mustDate("2025-12-15"),
							WithSalary:           false,
							Remote:               false,
						},
					},
				},
				domain.Rust: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Zig: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Scala: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Elixir: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Clojure: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Haskell: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Erlang: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
			},
			ShortDescription: "ScorePlay builds the media backbone for sports organization",
			Industries: []domain.Industry{
				domain.IndustryEntertainment,
			},
		},
		{
			ID:         0,  // system
			Type:       "", // system
			Name:       "Open Cosmos",
			BaseURL:    "https://www.open-cosmos.com/",
			CareersURL: "https://www.open-cosmos.com/careers",
			AboutURL:   "https://www.open-cosmos.com/about",
			BlogURL:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                10352999,
				IDs:               nil,
				Alias:             "opencosmos",
				Name:              "Open Cosmos",
				Followers:         "30K",
				Employees:         "51-200",
				AssociatedMembers: "193",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:     "opencosmos",
				Followers: "12",
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
					GitHubRepositoryCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Junior Backend Engineer",
							ShortDescription:     "",
							SwitchingOpportunity: "Experience or strong interest in backend software development (Go, Python, or similar)",
							URL:                  "https://app.welcometothejungle.com/jobs/_EoB7Jm4",
							Date:                 mustDate("2025-12-16"),
							WithSalary:           false,
							Remote:               false,
						},
					},
				},
				domain.Rust: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Zig: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Scala: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Elixir: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Clojure: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Haskell: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Erlang: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
			},
			ShortDescription: "Open Cosmos provides simple and affordable satellite missions",
			Industries:       []domain.Industry{
				// NOP
			},
		},
		{
			ID:         0,  // system
			Type:       "", // system
			Name:       "Axiom",
			BaseURL:    "https://axiom.co/",
			CareersURL: "",
			AboutURL:   "https://axiom.co/company",
			BlogURL:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                18582643,
				IDs:               nil,
				Alias:             "axiomhq",
				Name:              "Axiom",
				Followers:         "6K",
				Employees:         "11-50",
				AssociatedMembers: "59",
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
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Rust: {
					GitHubRepositoryCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Software Engineer",
							ShortDescription:     "2+ years experience with the Rust programming language",
							SwitchingOpportunity: "",
							URL:                  "https://app.welcometothejungle.com/jobs/sLh3GI7U",
							Date:                 mustDate("2025-12-16"),
							WithSalary:           false,
							Remote:               false,
						},
					},
				},
				domain.Zig: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Scala: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Elixir: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Clojure: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Haskell: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Erlang: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
			},
			ShortDescription: "Axiom has developed its own cloud native data storage solution that its says is easier to integrate and scale than its competitors",
			Industries: []domain.Industry{
				domain.IndustryDevOps,
			},
		},
		{
			ID:         0,  // system
			Type:       "", // system
			Name:       "Staffbase",
			BaseURL:    "https://staffbase.com/",
			CareersURL: "https://staffbase.com/jobs/",
			AboutURL:   "https://staffbase.com/about-us",
			BlogURL:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                9223407,
				IDs:               []int{2575394, 9223407, 10008057},
				Alias:             "staffbase",
				Name:              "Staffbase",
				Followers:         "38K",
				Employees:         "501-1K",
				AssociatedMembers: "945",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:     "staffbase",
				Followers: "125",
				Verified:  true,
			},
			BlindProfile: domain.BlindProfile{
				Alias: "",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias: "",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Percona-LLC-EI_IE283779.11,22.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/index.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Staffbase-Jobs-E1980315.htm",
				Jobs:        "", // Becomes empty in the new design version
				Reviews:     "", // Becomes empty in the new design version
				Salaries:    "", // Becomes empty in the new design version
				ReviewsRate: "4.3",
				Verified:    true,
				Date:        mustDate("2026-02-04"),
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "",
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoryCount: 16,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Backend Engineer – Golang",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4343666538/",
							Location:             "Leipzig, Saxony, Germany",
							Date:                 mustDate("2025-12-17"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Backend Engineer – Golang",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4344830063/",
							Location:             "Berlin, Berlin, Germany",
							Date:                 mustDate("2025-12-25"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Staff Engineer – Golang",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4365212068/",
							Location:             "Berlin, Berlin, Germany",
							CloudProviders:       []domain.CloudProvider{},
							Date:                 mustDate("2026-01-30", "2026-01-23"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Backend Engineer – Golang",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4369346699/",
							Location:             "Berlin, Berlin, Germany",
							CloudProviders:       []domain.CloudProvider{},
							Date:                 mustDate("2026-02-10"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Senior Backend Engineer – Golang",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4373861093/",
							Location:             "Dresden, Saxony, Germany",
							CloudProviders:       []domain.CloudProvider{},
							Date:                 mustDate("2026-02-19"),
							WithSalary:           false,
							Remote:               false,
						},
					},
				},
				domain.Rust: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Zig: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Scala: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Elixir: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Clojure: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Haskell: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Erlang: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
			},
			ShortDescription: "Employee Experience Platform",
			Industries:       []domain.Industry{
				// HR
			},
		},
		{
			ID:         0,  // system
			Type:       "", // system
			Name:       "CubiCasa",
			BaseURL:    "https://www.cubi.casa/",
			CareersURL: "",
			AboutURL:   "",
			BlogURL:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                3274772,
				IDs:               nil,
				Alias:             "cubicasa",
				Name:              "CubiCasa",
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
				domain.Go: {
					GitHubRepositoryCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "BackEnd Developer (Golang)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4327912288/",
							Location:             "Ho Chi Minh City, Vietnam",
							Date:                 mustDate("2025-12-17"),
							WithSalary:           false,
							Remote:               false,
						},
					},
				},
				domain.Rust: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Zig: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Scala: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Elixir: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Clojure: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Haskell: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Erlang: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
			},
			ShortDescription: "Interior property data mapping through smartphone scanning",
			Industries: []domain.Industry{
				domain.IndustryPropTech,
			},
		},
		{
			ID:         0,  // system
			Type:       "", // system
			Name:       "AxxonSoft",
			BaseURL:    "https://www.axxonsoft.com/",
			CareersURL: "https://www.axxonsoft.com/about-us/careers",
			AboutURL:   "https://www.axxonsoft.com/about-us/company",
			BlogURL:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                433282,
				IDs:               nil,
				Alias:             "axxonsoft",
				Name:              "AxxonSoft",
				Followers:         "17K",
				Employees:         "501-1K",
				AssociatedMembers: "216",
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
					GitHubRepositoryCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Junior Golang Developer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4344302127/",
							Location:             "Islamabad, Islāmābād, Pakistan",
							Date:                 mustDate("2025-12-17"),
							WithSalary:           false,
							Remote:               false,
						},
					},
				},
				domain.Rust: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Zig: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Scala: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Elixir: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Clojure: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Haskell: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Erlang: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
			},
			ShortDescription: "",
			Industries:       []domain.Industry{
				// NOP
			},
		},
		{
			ID:         0,  // system
			Type:       "", // system
			Name:       "CareerOS",
			BaseURL:    "https://www.careeros.com/",
			CareersURL: "https://www.careeros.com/about-us",
			AboutURL:   "",
			BlogURL:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                88681295,
				IDs:               nil,
				Alias:             "careeros",
				Name:              "CareerOS",
				Followers:         "5K",
				Employees:         "11-50",
				AssociatedMembers: "39",
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
					GitHubRepositoryCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Full Stack Engineer (Go/React)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4343858204/",
							Location:             "Barcelona, Catalonia, Spain",
							Date:                 mustDate("2025-12-17"),
							WithSalary:           false,
							Remote:               false,
						},
					},
				},
				domain.Rust: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Zig: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Scala: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Elixir: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Clojure: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Haskell: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Erlang: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
			},
			ShortDescription: "CareerOS helps students and professionals advance their careers by providing tools to connect with key decision-makers, strategically network, and organize every aspect of their job search in one platform",
			Industries:       []domain.Industry{
				// NOP
			},
		},
		{
			ID:         0,  // system
			Type:       "", // system
			Name:       "Tele2",
			BaseURL:    "https://www.tele2.com/",
			CareersURL: "https://www.tele2.com/join-us/",
			AboutURL:   "https://www.tele2.com/about/",
			BlogURL:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                2831,
				IDs:               []int{2831, 15196159},
				Alias:             "tele2",
				Name:              "Tele2",
				Followers:         "70K",
				Employees:         "1K-5K",
				AssociatedMembers: "5,280",
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
					GitHubRepositoryCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Software Engineer (Go)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4344015925/",
							Location:             "Stockholm, Stockholm County, Sweden",
							Date:                 mustDate("2025-12-17"),
							WithSalary:           false,
							Remote:               false,
						},
					},
				},
				domain.Rust: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Zig: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Scala: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Elixir: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Clojure: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Haskell: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Erlang: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
			},
			ShortDescription: "",
			Industries: []domain.Industry{
				domain.IndustryTelecom,
			},
		},
		{
			ID:         0,  // system
			Type:       "", // system
			Name:       "Tipalti",
			BaseURL:    "https://tipalti.com/",
			CareersURL: "https://tipalti.com/company/careers/",
			AboutURL:   "https://tipalti.com/company/",
			BlogURL:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                1770643,
				IDs:               nil,
				Alias:             "tipalti",
				Name:              "Tipalti",
				Followers:         "62K",
				Employees:         "1K-5K",
				AssociatedMembers: "1,234",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:     "tipalti",
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
					GitHubRepositoryCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Senior Backend Engineer (Go)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4338655342/",
							Location:             "Tel Aviv District, Israel",
							Date:                 mustDate("2025-12-17"),
							WithSalary:           false,
							Remote:               false,
						},
					},
				},
				domain.Rust: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Zig: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Scala: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Elixir: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Clojure: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Haskell: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Erlang: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
			},
			ShortDescription: "",
			Industries: []domain.Industry{
				domain.IndustryFinTech,
			},
		},
		{
			ID:         0,  // system
			Type:       "", // system
			Name:       "Unbabel",
			BaseURL:    "https://unbabel.com/",
			CareersURL: "https://careers.unbabel.com/",
			AboutURL:   "https://unbabel.com/about/",
			BlogURL:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                3327165,
				IDs:               nil,
				Alias:             "unbabel",
				Name:              "Unbabel",
				Followers:         "42K",
				Employees:         "201-500",
				AssociatedMembers: "716",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:     "Unbabel",
				Followers: "",
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
					GitHubRepositoryCount: 3,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Software Engineer (Python or Go)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4314026533/",
							Location:             "Lisbon, Portugal",
							Date:                 mustDate("2025-12-17"),
							WithSalary:           false,
							Remote:               false,
						},
					},
				},
				domain.Rust: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Zig: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Scala: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Elixir: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Clojure: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Haskell: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Erlang: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
			},
			ShortDescription: "",
			Industries:       []domain.Industry{
				// NOP
			},
		},
		{
			ID:         0,  // system
			Type:       "", // system
			Name:       "CARIAD",
			BaseURL:    "https://cariad.technology/",
			CareersURL: "https://cariad.technology/de/en/careers.html",
			AboutURL:   "https://cariad.technology/de/en/company.html",
			BlogURL:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                70099120,
				IDs:               nil,
				Alias:             "cariad-tech",
				Name:              "CARIAD",
				Followers:         "117K",
				Employees:         "5K-10K",
				AssociatedMembers: "4,677",
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
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Rust: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Zig: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Scala: {
					GitHubRepositoryCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "(Senior) Cloud Scala Engineer",
							ShortDescription:     "AI Platform",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4344599600/",
							Location:             "Ingolstadt, Bavaria, Germany",
							Date:                 mustDate("2025-12-19"),
							WithSalary:           false,
							Remote:               false,
						},
					},
				},
				domain.Elixir: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Clojure: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Haskell: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Erlang: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
			},
			ShortDescription: "CARIAD is the automotive software company of the Volkswagen Group",
			Industries: []domain.Industry{
				domain.IndustryAutomotive,
			},
		},
		{
			ID:         0,  // system
			Type:       "", // system
			Name:       "OpenAI",
			BaseURL:    "https://openai.com/",
			CareersURL: "https://openai.com/careers/",
			AboutURL:   "https://openai.com/about/",
			BlogURL:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                11130470,
				IDs:               nil,
				Alias:             "openai",
				Name:              "OpenAI",
				Followers:         "10M",
				Employees:         "201-500",
				AssociatedMembers: "6,824",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:     "OPENAI",
				Followers: "110k",
				Verified:  true,
			},
			BlindProfile: domain.BlindProfile{
				Alias: "",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias: "",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Percona-LLC-EI_IE283779.11,22.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/index.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/OpenAI-Jobs-E2210885.htm",
				Jobs:        "", // Becomes empty in the new design version
				Reviews:     "", // Becomes empty in the new design version
				Salaries:    "", // Becomes empty in the new design version
				ReviewsRate: "4.4",
				Verified:    false,
				Date:        mustDate("2026-02-04"),
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "",
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoryCount: 7,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Software Engineer (Python, Golang)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4346138871/",
							Location:             "United States",
							Date:                 mustDate("2025-12-25"),
							WithSalary:           true, // $255k/yr - $325k/yr
							Remote:               true,
							PinnedUntil:          mustDate("2026-01-14"),
						},
					},
					PinnedUntil: mustDate("2026-01-14"),
				},
				domain.Rust: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Zig: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Scala: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Elixir: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Clojure: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Haskell: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Erlang: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
			},
			ShortDescription: "",
			Industries:       []domain.Industry{
				// NOP
			},
		},
		{
			ID:         0,  // system
			Type:       "", // system
			Name:       "Cargill",
			BaseURL:    "https://www.cargill.com/",
			CareersURL: "https://careers.cargill.com/",
			AboutURL:   "https://www.cargill.com/about",
			BlogURL:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                2678,
				IDs:               []int{2678, 41345, 74321, 120416},
				Alias:             "cargill",
				Name:              "Cargill",
				Followers:         "3M",
				Employees:         "10K+",
				AssociatedMembers: "68,836",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:     "cargill",
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
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Percona-LLC-EI_IE283779.11,22.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/index.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Cargill-Jobs-E2739.htm",
				Jobs:        "", // Becomes empty in the new design version
				Reviews:     "", // Becomes empty in the new design version
				Salaries:    "", // Becomes empty in the new design version
				ReviewsRate: "4.0",
				Verified:    true,
				Date:        mustDate("2026-02-11"),
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "",
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoryCount: 8,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Software Engineer (Golang)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4346184199/",
							Location:             "Atlanta, GA",
							Date:                 mustDate("2026-02-06", "2026-01-16", "2025-12-25"),
							WithSalary:           false,
							Remote:               false,
						},
					},
				},
				domain.Rust: {
					GitHubRepositoryCount: 10,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Zig: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Scala: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Elixir: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Clojure: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Haskell: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Erlang: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
			},
			ShortDescription: "Food and Beverage Manufacturing",
			Industries:       []domain.Industry{
				// NOP
			},
		},
		{
			ID:         0,  // system
			Type:       "", // system
			Name:       "Sardine",
			BaseURL:    "https://www.sardine.ai/",
			CareersURL: "https://www.sardine.ai/careers",
			AboutURL:   "https://www.sardine.ai/about",
			BlogURL:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                76451674,
				IDs:               nil,
				Alias:             "sardineai",
				Name:              "Sardine",
				Followers:         "43K",
				Employees:         "201-500",
				AssociatedMembers: "292",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:     "sardine-ai",
				Followers: "33",
				Verified:  true,
			},
			BlindProfile: domain.BlindProfile{
				Alias: "",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias: "",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Percona-LLC-EI_IE283779.11,22.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/index.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Sardine-Jobs-E7044956.htm",
				Jobs:        "", // Becomes empty in the new design version
				Reviews:     "", // Becomes empty in the new design version
				Salaries:    "", // Becomes empty in the new design version
				ReviewsRate: "4.5",
				Verified:    true,
				Date:        mustDate("2026-02-15"),
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "",
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoryCount: 8,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Go Backend Engineer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4345959274/",
							Location:             "European Union",
							Date:                 mustDate("2025-12-24"),
							WithSalary:           false,
							Remote:               true,
						},
						{
							Title:                "Software Engineer – Backend (Go)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4372184776/",
							Location:             "Canada",
							CloudProviders:       []domain.CloudProvider{},
							Date:                 mustDate("2026-02-14"),
							WithSalary:           false,
							Remote:               true,
						},
					},
				},
				domain.Rust: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Zig: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Scala: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Elixir: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Clojure: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Haskell: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Erlang: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
			},
			ShortDescription: "AI Risk Platform for fraud, credit, and compliance",
			Industries: []domain.Industry{
				domain.IndustryCyberSecurity,
			},
		},
		{
			ID:         0,  // system
			Type:       "", // system
			Name:       "AppDirect",
			BaseURL:    "https://www.appdirect.com/",
			CareersURL: "",
			AboutURL:   "https://www.appdirect.com/about",
			BlogURL:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                1418447,
				IDs:               nil,
				Alias:             "appdirect",
				Name:              "AppDirect",
				Followers:         "144K",
				Employees:         "501-1K",
				AssociatedMembers: "1,109",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:     "appdirect",
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
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Percona-LLC-EI_IE283779.11,22.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/index.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/AppDirect-Jobs-E614387.htm",
				Jobs:        "", // Becomes empty in the new design version
				Reviews:     "", // Becomes empty in the new design version
				Salaries:    "", // Becomes empty in the new design version
				ReviewsRate: "3.7",
				Verified:    true,
				Date:        mustDate("2026-02-04"),
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "",
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoryCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Senior Full Stack Developer (Go/React)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4328874966/",
							Location:             "Greater Buenos Aires",
							Date:                 mustDate("2025-12-24"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Software Development Engineer (Golang)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4358139738/",
							Location:             "Greater Buenos Aires",
							Date:                 mustDate("2026-01-30", "2026-01-08"),
							WithSalary:           false,
							Remote:               false,
						},
					},
				},
				domain.Rust: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Zig: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Scala: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Elixir: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Clojure: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Haskell: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Erlang: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
			},
			ShortDescription: "B2B subscription commerce platform",
			Industries:       []domain.Industry{
				// NOP
			},
		},
		{
			ID:         0,  // system
			Type:       "", // system
			Name:       "Wix",
			BaseURL:    "https://www.wix.com/",
			CareersURL: "https://careers.wix.com/",
			AboutURL:   "https://www.wix.com/about/us",
			BlogURL:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                97250,
				IDs:               nil,
				Alias:             "wix-com",
				Name:              "Wix",
				Followers:         "365K",
				Employees:         "1K-5K",
				AssociatedMembers: "7,632",
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
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Rust: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Zig: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Scala: {
					GitHubRepositoryCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Server Engineer (Java/Scala)",
							ShortDescription:     "eCommerce Platform",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4345823191/",
							Location:             "Dnipro, Dnipropetrovsk, Ukraine",
							Date:                 mustDate("2025-12-23"),
							WithSalary:           false,
							Remote:               false,
						},
					},
				},
				domain.Elixir: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Clojure: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Haskell: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Erlang: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
			},
			ShortDescription: "Wix is the comprehensive platform",
			Industries:       []domain.Industry{
				// NOP
			},
		},
		{
			ID:         0,  // system
			Type:       "", // system
			Name:       "ZorgDomein",
			BaseURL:    "https://zorgdomein.com/",
			CareersURL: "",
			AboutURL:   "",
			BlogURL:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                2688962,
				IDs:               nil,
				Alias:             "zorgdomein",
				Name:              "ZorgDomein",
				Followers:         "12K",
				Employees:         "51-200",
				AssociatedMembers: "230",
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
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Rust: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Zig: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Scala: {
					GitHubRepositoryCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Software Engineer (Scala)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4328476948/",
							Location:             "Breukelen, Utrecht, Netherlands",
							Date:                 mustDate("2025-12-23"),
							WithSalary:           false,
							Remote:               false,
						},
					},
				},
				domain.Elixir: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Clojure: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Haskell: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Erlang: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
			},
			ShortDescription: "We are leveraging technology to make more room for people work in healthcare",
			Industries: []domain.Industry{
				domain.IndustryHealthTech,
			},
		},
		{
			ID:         0,  // system
			Type:       "", // system
			Name:       "Apple",
			BaseURL:    "https://www.apple.com/",
			CareersURL: "https://www.apple.com/careers/",
			AboutURL:   "",
			BlogURL:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                162479,
				IDs:               nil,
				Alias:             "apple",
				Name:              "Apple",
				Followers:         "18M",
				Employees:         "10K+",
				AssociatedMembers: "169,279",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:     "APPLE",
				Followers: "34.8k",
				Verified:  true,
			},
			BlindProfile: domain.BlindProfile{
				Alias: "",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias: "",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Percona-LLC-EI_IE283779.11,22.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/index.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Apple-Jobs-E1138.htm",
				Jobs:        "", // Becomes empty in the new design version
				Reviews:     "", // Becomes empty in the new design version
				Salaries:    "", // Becomes empty in the new design version
				ReviewsRate: "4.2",
				Verified:    true,
				Date:        mustDate("2026-02-04"),
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "",
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoryCount: 6,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Distributed Systems Software Engineer (Golang)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4359295380/",
							Location:             "San Diego, CA",
							CloudProviders:       []domain.CloudProvider{},
							Date:                 mustDate("2026-01-16"),
							WithSalary:           true, // $171.6k/yr - $302.2k/yr
							Remote:               false,
						},
						{
							Title:                "Senior Software Engineer – Golang",
							ShortDescription:     "Cloud Infrastructure",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4321352292/",
							Location:             "San Francisco, CA",
							CloudProviders:       []domain.CloudProvider{},
							Date:                 mustDate("2026-01-24"),
							WithSalary:           true, // $171.6k/yr - $302.2k/yr
							Remote:               false,
						},
						{
							Title:                "Senior Software Engineer – Golang",
							ShortDescription:     "Cloud Infrastructure",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4370779514/",
							Location:             "Austin, TX",
							CloudProviders:       []domain.CloudProvider{},
							Date:                 mustDate("2026-02-12"),
							WithSalary:           true, // $171.6k/yr - $302.2k/yr
							Remote:               false,
						},
					},
				},
				domain.Rust: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Zig: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Scala: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Elixir: {
					GitHubRepositoryCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Senior Elixir Software Engineer",
							ShortDescription:     "Environment and Supply Chain Innovation",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4359225502/",
							Location:             "Sunnyvale, CA",
							CloudProviders:       []domain.CloudProvider{},
							Date:                 mustDate("2026-01-16"),
							WithSalary:           true, // $181.1k/yr - $318.4k/yr
							Remote:               false,
						},
					},
				},
				domain.Clojure: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Haskell: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Erlang: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
			},
			ShortDescription: "",
			Industries:       []domain.Industry{
				// NOP
			},
		},
		{
			ID:         0,  // system
			Type:       "", // system
			Name:       "Dragonfly",
			BaseURL:    "https://www.dragonflydb.io/",
			CareersURL: "https://www.dragonflydb.io/careers",
			AboutURL:   "https://www.dragonflydb.io/about",
			BlogURL:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                82089379,
				IDs:               nil,
				Alias:             "dragonflydb",
				Name:              "Dragonfly",
				Followers:         "8K",
				Employees:         "11-50",
				AssociatedMembers: "33",
				Verified:          false,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:     "dragonflydb",
				Followers: "414",
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
					GitHubRepositoryCount: 6,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Senior Control Plane Backend Engineer (Golang)",
							SubTitle:             "",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.dragonflydb.io/careers/senior-control-plane-backend-engineer-golang",
							Location:             "Israel",
							Date:                 mustDate("2025-08-10"),
							WithSalary:           false,
							Remote:               false,
						},
					},
				},
				domain.Rust: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Zig: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Scala: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Elixir: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Clojure: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Haskell: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Erlang: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
			},
			ShortDescription: "Key-value database",
			Industries: []domain.Industry{
				domain.IndustryDevOps,
			},
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
			},
		},
		{
			ID:         0,  // system
			Type:       "", // system
			Name:       "BRP",
			BaseURL:    "https://www.brp.com/",
			CareersURL: "https://www.brp.com/careers.html",
			AboutURL:   "https://www.brp.com/our-company.html",
			BlogURL:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                10580,
				IDs:               nil,
				Alias:             "brp",
				Name:              "BRP",
				Followers:         "185K",
				Employees:         "10K+",
				AssociatedMembers: "8,829",
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
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Percona-LLC-EI_IE283779.11,22.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/index.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/BRP-Jobs-E314918.htm",
				Jobs:        "", // Becomes empty in the new design version
				Reviews:     "", // Becomes empty in the new design version
				Salaries:    "", // Becomes empty in the new design version
				ReviewsRate: "3.7",
				Verified:    true,
				Date:        mustDate("2026-02-11"),
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "",
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Rust: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Zig: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Scala: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Elixir: {
					GitHubRepositoryCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Backend Elixir Developer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4352623829/",
							Location:             "Greater Campinas",
							CloudProviders:       []domain.CloudProvider{},
							Date:                 mustDate("2026-02-03", "2026-01-14"),
							WithSalary:           false,
							Remote:               false,
						},
					},
				},
				domain.Clojure: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Haskell: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Erlang: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
			},
			ShortDescription: "",
			Industries:       []domain.Industry{
				// NOP
			},
		},
		{
			ID:         0,  // system
			Type:       "", // system
			Name:       "Commvault",
			BaseURL:    "https://www.commvault.com/",
			CareersURL: "https://www.commvault.com/careers",
			AboutURL:   "https://www.commvault.com/about-us",
			BlogURL:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                163166,
				IDs:               []int{163166, 3771318},
				Alias:             "commvault",
				Name:              "Commvault",
				Followers:         "185K",
				Employees:         "1K-5K",
				AssociatedMembers: "3,482",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:     "commvault",
				Followers: "69",
				Verified:  false,
				Date:      mustDate("2026-02-15"),
			},
			BlindProfile: domain.BlindProfile{
				Alias: "",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias: "",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Percona-LLC-EI_IE283779.11,22.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/index.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Commvault-Jobs-E16184.htm",
				Jobs:        "", // Becomes empty in the new design version
				Reviews:     "", // Becomes empty in the new design version
				Salaries:    "", // Becomes empty in the new design version
				ReviewsRate: "3.8",
				Verified:    true,
				Date:        mustDate("2026-02-11"),
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "",
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoryCount: 2,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Rust: {
					GitHubRepositoryCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Principal Engineer (Rust/C/C++)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4362002359/",
							Location:             "Israel",
							CloudProviders:       []domain.CloudProvider{},
							Date:                 mustDate("2026-02-12", "2026-01-20"),
							WithSalary:           false,
							Remote:               false,
						},
					},
				},
				domain.Zig: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Scala: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Elixir: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Clojure: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Haskell: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Erlang: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
			},
			ShortDescription: "Cybersecurity and data protection software company",
			YahooFinanceURL:  "https://finance.yahoo.com/quote/CVLT/",
			GoogleFinanceURL: "https://www.google.com/finance/quote/CVLT:NASDAQ",
			Industries: []domain.Industry{
				domain.IndustryCyberSecurity,
			},
		},
		{
			ID:         0,  // system
			Type:       "", // system
			Name:       "Elastio",
			BaseURL:    "https://elastio.com/",
			CareersURL: "",
			AboutURL:   "",
			BlogURL:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                65520227,
				IDs:               nil,
				Alias:             "elastio",
				Name:              "Elastio",
				Followers:         "3K",
				Employees:         "51-200",
				AssociatedMembers: "42",
				Verified:          false,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:     "elastio",
				Followers: "31",
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
					GitHubRepositoryCount: 1,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Senior Go Engineer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4361981306/",
							Location:             "Greater Toronto Area, Canada",
							CloudProviders:       []domain.CloudProvider{},
							Date:                 mustDate("2026-01-20"),
							WithSalary:           false,
							Remote:               false,
						},
					},
				},
				domain.Rust: {
					GitHubRepositoryCount: 20,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Zig: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Scala: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Elixir: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Clojure: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Haskell: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Erlang: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
			},
			ShortDescription: "Elastio proves ransomware recovery and ROI by validating data and backups, cutting downtime, removing ransom leverage, and making recovery risk measurable and insurable",
			Industries: []domain.Industry{
				domain.IndustryCyberSecurity,
			},
		},
		{
			ID:         0,  // system
			Type:       "", // system
			Name:       "Tyba",
			BaseURL:    "https://www.tyba.ai/",
			CareersURL: "https://www.tyba.ai/careers/",
			AboutURL:   "https://www.tyba.ai/about/",
			BlogURL:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                74758359,
				IDs:               nil,
				Alias:             "tybaenergy",
				Name:              "Tyba",
				Followers:         "4K",
				Employees:         "11-50",
				AssociatedMembers: "37",
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
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Rust: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Zig: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Scala: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Elixir: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Clojure: {
					GitHubRepositoryCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Clojure Software Engineer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4365415737/",
							Location:             "United States",
							CloudProviders:       []domain.CloudProvider{},
							Date:                 mustDate("2026-01-23"),
							WithSalary:           true, // $160k/yr - $200k/yr
							Remote:               true,
						},
					},
				},
				domain.Haskell: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Erlang: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
			},
			ShortDescription: "Enabling the clean energy transition through profitable, scalable energy storage operations",
			Industries:       []domain.Industry{
				// NOP
			},
		},
		{
			ID:         0,  // system
			Type:       "", // system
			Name:       "Fractile",
			BaseURL:    "https://www.fractile.ai/",
			CareersURL: "",
			AboutURL:   "",
			BlogURL:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                75705060,
				IDs:               nil,
				Alias:             "fractile",
				Name:              "Fractile",
				Followers:         "8K",
				Employees:         "51-200",
				AssociatedMembers: "68",
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
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Rust: {
					GitHubRepositoryCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Senior Rust Software Engineer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4365068054/",
							Location:             "Bristol, England, United Kingdom",
							CloudProviders:       []domain.CloudProvider{},
							Date:                 mustDate("2026-02-12", "2026-01-22"),
							WithSalary:           false,
							Remote:               false,
						},
					},
				},
				domain.Zig: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Scala: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Elixir: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Clojure: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Haskell: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Erlang: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
			},
			ShortDescription: "Fractile is building chips that remove every bottleneck to running large AI models",
			Industries:       []domain.Industry{
				// NOP
			},
		},
		{
			ID:         0,  // system
			Type:       "", // system
			Name:       "SensitHaptics",
			BaseURL:    "https://sensithaptics.com/",
			CareersURL: "",
			AboutURL:   "",
			BlogURL:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                20487540,
				IDs:               nil,
				Alias:             "sensit-haptics",
				Name:              "SensitHaptics",
				Followers:         "2K",
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
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Rust: {
					GitHubRepositoryCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "(Senior) Software Engineer (Rust)",
							ShortDescription:     "Real-Time Systems",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4353968645/",
							Location:             "Potsdam, Brandenburg, Germany",
							CloudProviders:       []domain.CloudProvider{},
							Date:                 mustDate("2026-01-22"),
							WithSalary:           false,
							Remote:               false,
						},
					},
				},
				domain.Zig: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Scala: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Elixir: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Clojure: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Haskell: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Erlang: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
			},
			ShortDescription: "SensitHaptics is the global leader for advanced haptics technologies for sim racing, sim gaming and entertainment",
			Industries: []domain.Industry{
				domain.IndustryEntertainment,
			},
		},
		{
			ID:         0,  // system
			Type:       "", // system
			Name:       "Temporal Technologies",
			BaseURL:    "https://temporal.io/",
			CareersURL: "",
			AboutURL:   "",
			BlogURL:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                67081245,
				IDs:               nil,
				Alias:             "temporal-technologies",
				Name:              "Temporal Technologies",
				Followers:         "24K",
				Employees:         "201-500",
				AssociatedMembers: "399",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:     "temporalio",
				Followers: "2.7k",
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
					GitHubRepositoryCount: 54,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Senior Staff Software Engineer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://app.welcometothejungle.com/jobs/4rzK4S2Q",
							Date:                 mustDate("2026-01-27"),
							WithSalary:           false,
							Remote:               false,
						},
					},
				},
				domain.Rust: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Zig: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Scala: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Elixir: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Clojure: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Haskell: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Erlang: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
			},
			ShortDescription: "Open source durable execution system",
			Industries: []domain.Industry{
				domain.IndustryDevOps,
			},
		},
		{
			ID:         0,  // system
			Type:       "", // system
			Name:       "Synctera",
			BaseURL:    "https://www.synctera.com/",
			CareersURL: "https://www.synctera.com/careers",
			AboutURL:   "https://www.synctera.com/about",
			BlogURL:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                68275317,
				IDs:               nil,
				Alias:             "synctera",
				Name:              "Synctera",
				Followers:         "84K",
				Employees:         "51-200",
				AssociatedMembers: "100",
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
					GitHubRepositoryCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Senior Software Engineer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://app.welcometothejungle.com/jobs/Aj9OCtLp",
							Date:                 mustDate("2026-01-27"),
							WithSalary:           false,
							Remote:               false,
						},
					},
				},
				domain.Rust: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Zig: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Scala: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Elixir: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Clojure: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Haskell: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Erlang: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
			},
			ShortDescription: "Synctera’s platform gives companies of all sizes the technology infrastructure, sponsor bank connection, and compliance framework they need to launch FinTech or embedded banking products",
			Industries: []domain.Industry{
				domain.IndustryFinTech,
			},
		},
		{
			ID:         0,  // system
			Type:       "", // system
			Name:       "Finary",
			BaseURL:    "https://finary.com/",
			CareersURL: "",
			AboutURL:   "",
			BlogURL:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                65846455,
				IDs:               nil,
				Alias:             "finaryhq",
				Name:              "Finary",
				Followers:         "43K",
				Employees:         "51-200",
				AssociatedMembers: "159",
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
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Percona-LLC-EI_IE283779.11,22.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/index.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Finary-Jobs-E6643956.htm",
				Jobs:        "", // Becomes empty in the new design version
				Reviews:     "", // Becomes empty in the new design version
				Salaries:    "", // Becomes empty in the new design version
				ReviewsRate: "4.7",
				Verified:    true,
				Date:        mustDate("2026-02-04"),
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "",
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Rust: {
					GitHubRepositoryCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Senior Software Engineer (Rust)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4365057961/",
							Location:             "Paris, Île-de-France, France",
							CloudProviders:       []domain.CloudProvider{},
							Date:                 mustDate("2026-01-22"),
							WithSalary:           false,
							Remote:               true,
						},
					},
				},
				domain.Zig: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Scala: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Elixir: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Clojure: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Haskell: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Erlang: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
			},
			ShortDescription: "",
			Industries: []domain.Industry{
				domain.IndustryFinTech,
			},
		},
		{
			ID:         0,  // system
			Type:       "", // system
			Name:       "Lyft",
			BaseURL:    "https://www.lyft.com/",
			CareersURL: "https://www.lyft.com/careers",
			AboutURL:   "",
			BlogURL:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                2620735,
				IDs:               nil,
				Alias:             "lyft",
				Name:              "Lyft",
				Followers:         "396K",
				Employees:         "5K-10K",
				AssociatedMembers: "27,444",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:     "lyft",
				Followers: "704",
				Verified:  true,
			},
			BlindProfile: domain.BlindProfile{
				Alias: "",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias: "",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Percona-LLC-EI_IE283779.11,22.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/index.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Lyft-Jobs-E700614.htm",
				Jobs:        "", // Becomes empty in the new design version
				Reviews:     "", // Becomes empty in the new design version
				Salaries:    "", // Becomes empty in the new design version
				ReviewsRate: "3.8",
				Verified:    false,
				Date:        mustDate("2026-02-04"),
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "",
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoryCount: 58,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Software Engineer",
							ShortDescription:     "Experience with Python, Go",
							SwitchingOpportunity: "",
							URL:                  "https://app.welcometothejungle.com/jobs/8uY1vXia",
							Date:                 mustDate("2026-01-27"),
							WithSalary:           false,
							Remote:               false,
						},
					},
				},
				domain.Rust: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Zig: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Scala: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Elixir: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Clojure: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Haskell: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Erlang: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
			},
			ShortDescription: "",
			Industries: []domain.Industry{
				domain.IndustryLogisticsTech,
			},
		},
		{
			ID:         0,  // system
			Type:       "", // system
			Name:       "Pulley",
			BaseURL:    "https://pulley.com/",
			CareersURL: "https://pulley.com/careers",
			AboutURL:   "",
			BlogURL:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                42343737,
				IDs:               nil,
				Alias:             "pulley-cap-table",
				Name:              "Pulley",
				Followers:         "22K",
				Employees:         "51-200",
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
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Percona-LLC-EI_IE283779.11,22.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/index.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Pulley-Jobs-E7426761.htm",
				Jobs:        "", // Becomes empty in the new design version
				Reviews:     "", // Becomes empty in the new design version
				Salaries:    "", // Becomes empty in the new design version
				ReviewsRate: "3.7",
				Verified:    true,
				Date:        mustDate("2026-02-11"),
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "",
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoryCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Backend Engineer",
							ShortDescription:     "Proficiency in Go, with experience building web services",
							SwitchingOpportunity: "",
							URL:                  "https://app.welcometothejungle.com/jobs/zdAIjgww",
							Date:                 mustDate("2026-01-27"),
							WithSalary:           false,
							Remote:               false,
						},
					},
				},
				domain.Rust: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Zig: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Scala: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Elixir: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Clojure: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Haskell: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Erlang: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
			},
			ShortDescription: "Equity management for high growth startups",
			Industries:       []domain.Industry{
				// NOP
			},
		},
		{
			ID:         0,  // system
			Type:       "", // system
			Name:       "TigerGraph",
			BaseURL:    "https://www.tigergraph.com/",
			CareersURL: "",
			AboutURL:   "",
			BlogURL:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                3693966,
				IDs:               nil,
				Alias:             "tigergraph",
				Name:              "TigerGraph",
				Followers:         "52K",
				Employees:         "201-500",
				AssociatedMembers: "130",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:     "tigergraph",
				Followers: "152",
				Verified:  false,
			},
			BlindProfile: domain.BlindProfile{
				Alias: "",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias: "",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Percona-LLC-EI_IE283779.11,22.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/index.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/TigerGraph-Jobs-E1145722.htm",
				Jobs:        "", // Becomes empty in the new design version
				Reviews:     "", // Becomes empty in the new design version
				Salaries:    "", // Becomes empty in the new design version
				ReviewsRate: "3.1",
				Verified:    true,
				Date:        mustDate("2026-02-04"),
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "",
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoryCount: 5,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Software Engineer",
							ShortDescription:     "",
							SwitchingOpportunity: "Solid programming fundamentals; experienced with C++, Go, or any other major programming language",
							URL:                  "https://app.welcometothejungle.com/jobs/bxA9L-FI",
							Date:                 mustDate("2026-01-27"),
							WithSalary:           false,
							Remote:               false,
						},
					},
				},
				domain.Rust: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Zig: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Scala: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Elixir: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Clojure: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Haskell: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Erlang: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
			},
			ShortDescription: "Graph database",
			Industries: []domain.Industry{
				domain.IndustryDevOps,
			},
		},
		{
			ID:         0,  // system
			Type:       "", // system
			Name:       "Carta Healthcare",
			BaseURL:    "http://carta.healthcare/",
			CareersURL: "https://www.carta.healthcare/careers/",
			AboutURL:   "https://www.carta.healthcare/our-story/",
			BlogURL:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                18706683,
				IDs:               nil,
				Alias:             "cartahealthcare",
				Name:              "Carta Healthcare",
				Followers:         "19K",
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
				domain.Go: {
					GitHubRepositoryCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Senior Software Engineer",
							ShortDescription:     "",
							SwitchingOpportunity: "Experience or interest in Go",
							URL:                  "https://app.welcometothejungle.com/jobs/1LWG87a7",
							Date:                 mustDate("2026-01-27"),
							WithSalary:           false,
							Remote:               false,
						},
					},
				},
				domain.Rust: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Zig: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Scala: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Elixir: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Clojure: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Haskell: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Erlang: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
			},
			ShortDescription: "Provider of clinical data management solutions",
			Industries: []domain.Industry{
				domain.IndustryMedTech,
			},
		},
		{
			ID:         0,  // system
			Type:       "", // system
			Name:       "League",
			BaseURL:    "https://league.com/",
			CareersURL: "https://league.com/careers/",
			AboutURL:   "https://league.com/about/",
			BlogURL:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                5401900,
				IDs:               nil,
				Alias:             "league-inc-",
				Name:              "League",
				Followers:         "53K",
				Employees:         "501-1K",
				AssociatedMembers: "530",
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
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Percona-LLC-EI_IE283779.11,22.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/index.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/League-Jobs-E1165876.htm",
				Jobs:        "", // Becomes empty in the new design version
				Reviews:     "", // Becomes empty in the new design version
				Salaries:    "", // Becomes empty in the new design version
				ReviewsRate: "3.8",
				Verified:    true,
				Date:        mustDate("2026-02-11"),
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "",
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoryCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Senior Backend Software Engineer",
							ShortDescription:     "Working primarily with Go and MongoDB",
							SwitchingOpportunity: "",
							URL:                  "https://app.welcometothejungle.com/jobs/FEj04ps5",
							Date:                 mustDate("2026-01-27"),
							WithSalary:           false,
							Remote:               false,
						},
					},
				},
				domain.Rust: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Zig: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Scala: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Elixir: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Clojure: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Haskell: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Erlang: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
			},
			ShortDescription: "Healthcare consumer experience platform",
			Industries: []domain.Industry{
				domain.IndustryMedTech,
			},
		},
		{
			ID:         0,  // system
			Type:       "", // system
			Name:       "iManage",
			BaseURL:    "https://imanage.com/",
			CareersURL: "https://imanage.com/about/the-company/careers/",
			AboutURL:   "https://imanage.com/about/the-company/",
			BlogURL:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                4526,
				IDs:               nil,
				Alias:             "imanage",
				Name:              "iManage",
				Followers:         "72K",
				Employees:         "1K-5K",
				AssociatedMembers: "1,245",
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
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Percona-LLC-EI_IE283779.11,22.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/index.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/iManage-Jobs-E1311409.htm",
				Jobs:        "", // Becomes empty in the new design version
				Reviews:     "", // Becomes empty in the new design version
				Salaries:    "", // Becomes empty in the new design version
				ReviewsRate: "4.4",
				Verified:    true,
				Date:        mustDate("2026-02-04"),
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "",
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Rust: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Zig: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Scala: {
					GitHubRepositoryCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Principal AI Software Engineer (Java, Scala)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4365966011/",
							Location:             "London, England, United Kingdom",
							CloudProviders:       []domain.CloudProvider{},
							Date:                 mustDate("2026-01-28"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "AI Software Engineer (Java, Scala)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4321918504/",
							Location:             "London, England, United Kingdom",
							CloudProviders:       []domain.CloudProvider{},
							Date:                 mustDate("2026-02-06"),
							WithSalary:           false,
							Remote:               false,
						},
					},
				},
				domain.Elixir: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Clojure: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Haskell: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Erlang: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
			},
			ShortDescription: "iManage is the company dedicated to Making Knowledge Work",
			Industries:       []domain.Industry{
				// NOP
			},
		},
		{
			ID:         0,  // system
			Type:       "", // system
			Name:       "SentiOne",
			BaseURL:    "https://sentione.com/",
			CareersURL: "",
			AboutURL:   "https://sentione.com/about",
			BlogURL:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                2320009,
				IDs:               nil,
				Alias:             "sentione",
				Name:              "SentiOne",
				Followers:         "6K",
				Employees:         "51-200",
				AssociatedMembers: "65",
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
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Rust: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Zig: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Scala: {
					GitHubRepositoryCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Scala Developer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4366245449/",
							Location:             "Gdańsk, Pomorskie, Poland",
							CloudProviders:       []domain.CloudProvider{domain.AWS},
							Date:                 mustDate("2026-01-28"),
							WithSalary:           false,
							Remote:               false,
						},
					},
				},
				domain.Elixir: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Clojure: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Haskell: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Erlang: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
			},
			ShortDescription: "SentiOne is an Omnichannel Conversational AI Platform",
			Industries:       []domain.Industry{
				// NOP
			},
		},
		{
			ID:         0,  // system
			Type:       "", // system
			Name:       "Pulsar",
			BaseURL:    "https://pulsar.com/",
			CareersURL: "https://pulsar.com/careers/",
			AboutURL:   "https://pulsar.com/about-us/",
			BlogURL:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                10045951,
				IDs:               nil,
				Alias:             "pulsar-crypto",
				Name:              "Pulsar",
				Followers:         "6K",
				Employees:         "51-200",
				AssociatedMembers: "79",
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
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Rust: {
					GitHubRepositoryCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Senior Rust Developer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4356492665/",
							Location:             "London Area, United Kingdom",
							CloudProviders:       []domain.CloudProvider{},
							Date:                 mustDate("2026-01-30"),
							WithSalary:           false,
							Remote:               false,
						},
					},
				},
				domain.Zig: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Scala: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Elixir: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Clojure: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Haskell: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Erlang: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
			},
			ShortDescription: "Pulsar is a proprietary algorithmic trading firm, specialized in high frequency trading and market making in the crypto market",
			Industries: []domain.Industry{
				domain.IndustryCryptoCurrency,
			},
		},
		{
			ID:         0,  // system
			Type:       "", // system
			Name:       "Superbet",
			BaseURL:    "https://www.super.xyz/",
			CareersURL: "https://www.super.xyz/careers/",
			AboutURL:   "",
			BlogURL:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                1162677,
				IDs:               nil,
				Alias:             "superbet-international",
				Name:              "Superbet",
				Followers:         "66K",
				Employees:         "1K-5K",
				AssociatedMembers: "1,722",
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
					GitHubRepositoryCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Senior Backend Engineer (Go)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4330753912/",
							Location:             "Zagreb, Croatia",
							CloudProviders:       []domain.CloudProvider{domain.AWS},
							Date:                 mustDate("2026-01-29"),
							WithSalary:           false,
							Remote:               false,
						},
					},
				},
				domain.Rust: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Zig: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Scala: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Elixir: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Clojure: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Haskell: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Erlang: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
			},
			ShortDescription: "",
			Industries:       []domain.Industry{
				// Betting
			},
			Ignore: true,
		},
		{
			ID:         0,  // system
			Type:       "", // system
			Name:       "Paribu",
			BaseURL:    "https://www.paribu.com/",
			CareersURL: "",
			AboutURL:   "",
			BlogURL:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                17978220,
				IDs:               nil,
				Alias:             "paribucom",
				Name:              "Paribu",
				Followers:         "50K",
				Employees:         "201-500",
				AssociatedMembers: "160",
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
					GitHubRepositoryCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Senior Backend Engineer (Go)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4339870078/",
							Location:             "Istanbul, Türkiye",
							CloudProviders:       []domain.CloudProvider{domain.AWS},
							Date:                 mustDate("2026-01-28"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Senior Backend Engineer (Go)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4363635733/",
							Location:             "Istanbul, Türkiye",
							CloudProviders:       []domain.CloudProvider{domain.AWS},
							Date:                 mustDate("2026-02-13"),
							WithSalary:           false,
							Remote:               false,
						},
					},
				},
				domain.Rust: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Zig: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Scala: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Elixir: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Clojure: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Haskell: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Erlang: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
			},
			ShortDescription: "Türkiye’s crypto asset ecosystem",
			Industries: []domain.Industry{
				domain.IndustryCryptoCurrency,
			},
		},
		{
			ID:         0,  // system
			Type:       "", // system
			Name:       "BitGo",
			BaseURL:    "https://www.bitgo.com/",
			CareersURL: "https://www.bitgo.com/company/careers/",
			AboutURL:   "https://www.bitgo.com/company/about-us/",
			BlogURL:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                3597797,
				IDs:               []int{3597797, 18745410},
				Alias:             "bitgo",
				Name:              "BitGo",
				Followers:         "56K",
				Employees:         "201-500",
				AssociatedMembers: "148",
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
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Percona-LLC-EI_IE283779.11,22.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/index.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/BitGo-Jobs-E1912848.htm",
				Jobs:        "", // Becomes empty in the new design version
				Reviews:     "", // Becomes empty in the new design version
				Salaries:    "", // Becomes empty in the new design version
				ReviewsRate: "4.5",
				Verified:    false,
				Date:        mustDate("2026-02-11"),
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "",
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoryCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Software Engineer",
							ShortDescription:     "Experience with server-side languages like Golang",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4365876616/",
							Location:             "New York, United States",
							CloudProviders:       []domain.CloudProvider{},
							Date:                 mustDate("2026-01-27"),
							WithSalary:           true, // $160k/yr - $200k/yr
							Remote:               false,
						},
					},
				},
				domain.Rust: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Zig: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Scala: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Elixir: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Clojure: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Haskell: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Erlang: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
			},
			ShortDescription: "Infrastructure provider of digital asset solutions",
			Industries: []domain.Industry{
				domain.IndustryCryptoCurrency,
			},
		},
		{
			ID:         0,  // system
			Type:       "", // system
			Name:       "Flybits",
			BaseURL:    "https://flybits.com/",
			CareersURL: "https://flybits.com/careers/",
			AboutURL:   "https://flybits.com/about-us/",
			BlogURL:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                2630845,
				IDs:               nil,
				Alias:             "flybits",
				Name:              "Flybits",
				Followers:         "28K",
				Employees:         "51-200",
				AssociatedMembers: "66",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:     "flybits",
				Followers: "15",
				Verified:  true,
			},
			BlindProfile: domain.BlindProfile{
				Alias: "",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias: "",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Percona-LLC-EI_IE283779.11,22.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/index.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Flybits-Jobs-E1910414.htm",
				Jobs:        "", // Becomes empty in the new design version
				Reviews:     "", // Becomes empty in the new design version
				Salaries:    "", // Becomes empty in the new design version
				ReviewsRate: "3.6",
				Verified:    true,
				Date:        mustDate("2026-02-04"),
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "",
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoryCount: 2,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Senior Backend Engineer",
							SubTitle:             "",
							ShortDescription:     "",
							SwitchingOpportunity: "Expertise in Go or at least one other major systems language with a strong desire to work in Go",
							URL:                  "https://ats.rippling.com/flybits/jobs/8785b045-6ce7-475e-8d55-3e65f634a6d7",
							Location:             "Toronto, ON", // Toronto, Canada
							CloudProviders:       []domain.CloudProvider{domain.AWS, domain.GCP, domain.Azure},
							Date:                 mustDate("2026-02-03"),
							WithSalary:           true, // $145,000-$195,000
							Remote:               false,
						},
					},
				},
				domain.Rust: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Zig: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Scala: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Elixir: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Clojure: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Haskell: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Erlang: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
			},
			ShortDescription: "Contextual engagement platform for the financial industry",
			Industries: []domain.Industry{
				domain.IndustryFinTech,
			},
		},
		{
			ID:         0,  // system
			Type:       "", // system
			Name:       "WorldQuant",
			BaseURL:    "https://www.worldquant.com/",
			CareersURL: "https://www.worldquant.com/careers/",
			AboutURL:   "",
			BlogURL:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                276383,
				IDs:               nil,
				Alias:             "worldquant",
				Name:              "WorldQuant",
				Followers:         "149K",
				Employees:         "501-1K",
				AssociatedMembers: "2,518",
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
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Percona-LLC-EI_IE283779.11,22.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/index.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/WorldQuant-Jobs-E309841.htm",
				Jobs:        "", // Becomes empty in the new design version
				Reviews:     "", // Becomes empty in the new design version
				Salaries:    "", // Becomes empty in the new design version
				ReviewsRate: "4.2",
				Verified:    true,
				Date:        mustDate("2026-02-04"),
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "",
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoryCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Python/Go Software Engineer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4367337482/",
							Location:             "Budapest, Budapest, Hungary",
							CloudProviders:       []domain.CloudProvider{},
							Date:                 mustDate("2026-01-30"),
							WithSalary:           false,
							Remote:               false,
						},
					},
				},
				domain.Rust: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Zig: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Scala: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Elixir: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Clojure: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Haskell: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Erlang: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
			},
			ShortDescription: "Quantitative asset management",
			Industries: []domain.Industry{
				domain.IndustryFinTech,
			},
		},
		{
			ID:         0,  // system
			Type:       "", // system
			Name:       "Delinea",
			BaseURL:    "https://delinea.com/",
			CareersURL: "",
			AboutURL:   "https://delinea.com/about",
			BlogURL:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                24070,
				IDs:               nil,
				Alias:             "delinea",
				Name:              "Delinea",
				Followers:         "42K",
				Employees:         "1K-5K",
				AssociatedMembers: "1,233",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:     "DelineaXPM",
				Followers: "158",
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
					GitHubRepositoryCount: 10,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Go Engineer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4328183145/",
							Location:             "Mexico City, Mexico",
							CloudProviders:       []domain.CloudProvider{},
							Date:                 mustDate("2026-02-01"),
							WithSalary:           false,
							Remote:               false,
						},
					},
				},
				domain.Rust: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Zig: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Scala: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Elixir: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Clojure: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Haskell: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Erlang: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
			},
			ShortDescription: "Securing identities through centralized authorization",
			Industries: []domain.Industry{
				domain.IndustryCyberSecurity,
			},
		},
		{
			ID:         0,  // system
			Type:       "", // system
			Name:       "Monks",
			BaseURL:    "https://www.monks.com/",
			CareersURL: "https://www.monks.com/careers",
			AboutURL:   "https://www.monks.com/what-we-do",
			BlogURL:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                46926,
				IDs:               nil,
				Alias:             "monks",
				Name:              "Monks",
				Followers:         "543K",
				Employees:         "5K-10K",
				AssociatedMembers: "6,521",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:     "mediamonks",
				Followers: "61",
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
					GitHubRepositoryCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Staff Engineer – Ruby, Golang, Python",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4367672197/",
							Location:             "Colombia",
							CloudProviders:       []domain.CloudProvider{domain.AWS},
							Date:                 mustDate("2026-02-01"),
							WithSalary:           false,
							Remote:               true,
						},
					},
				},
				domain.Rust: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Zig: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Scala: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Elixir: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Clojure: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Haskell: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Erlang: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
			},
			ShortDescription: "",
			Industries: []domain.Industry{
				domain.IndustryAdTech,
			},
		},
		{
			ID:         0,  // system
			Type:       "", // system
			Name:       "MotorInc",
			BaseURL:    "https://motorinc.com/",
			CareersURL: "",
			AboutURL:   "",
			BlogURL:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                87400465,
				IDs:               nil,
				Alias:             "motorinc",
				Name:              "MotorInc",
				Followers:         "14K",
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
				domain.Go: {
					GitHubRepositoryCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Backend Developer (Golang / Elixir / Python)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4367553323/",
							Location:             "Navi Mumbai, Maharashtra, India",
							CloudProviders:       []domain.CloudProvider{},
							Date:                 mustDate("2026-01-31"),
							WithSalary:           false,
							Remote:               false,
						},
					},
				},
				domain.Rust: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Zig: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Scala: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Elixir: {
					GitHubRepositoryCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Backend Developer (Golang / Elixir / Python)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4367553323/",
							Location:             "Navi Mumbai, Maharashtra, India",
							CloudProviders:       []domain.CloudProvider{},
							Date:                 mustDate("2026-01-31"),
							WithSalary:           false,
							Remote:               false,
						},
					},
				},
				domain.Clojure: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Haskell: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Erlang: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
			},
			ShortDescription: "MotorInc helps people buy vehicles with confidence",
			Industries:       []domain.Industry{
				// NOP
			},
		},
		{
			ID:         0,  // system
			Type:       "", // system
			Name:       "ng-voice",
			BaseURL:    "https://www.ng-voice.com/",
			CareersURL: "https://www.ng-voice.com/careers",
			AboutURL:   "https://www.ng-voice.com/about-us",
			BlogURL:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                3215598,
				IDs:               nil,
				Alias:             "ngvoice",
				Name:              "ng-voice",
				Followers:         "6K",
				Employees:         "51-200",
				AssociatedMembers: "115",
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
					GitHubRepositoryCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Software Engineer – Go",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4368231130/",
							Location:             "United States",
							CloudProviders:       []domain.CloudProvider{},
							Date:                 mustDate("2026-02-03"),
							WithSalary:           false,
							Remote:               true,
						},
					},
				},
				domain.Rust: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Zig: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Scala: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Elixir: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Clojure: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Haskell: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Erlang: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
			},
			ShortDescription: "IMS (IP Multimedia Subsystem) Solution",
			Industries: []domain.Industry{
				domain.IndustryTelecom,
			},
		},
		{
			ID:         0,  // system
			Type:       "", // system
			Name:       "Soundtrack",
			BaseURL:    "https://www.soundtrack.io/",
			CareersURL: "https://careers.soundtrack.io/",
			AboutURL:   "",
			BlogURL:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                3296806,
				IDs:               nil,
				Alias:             "soundtrackio",
				Name:              "Soundtrack",
				Followers:         "11K",
				Employees:         "51-200",
				AssociatedMembers: "261",
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
					GitHubRepositoryCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Backend Engineer (Go)",
							ShortDescription:     "Business systems",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4344990032/",
							Location:             "Stockholm, Stockholm County, Sweden",
							CloudProviders:       []domain.CloudProvider{},
							Date:                 mustDate("2026-02-05"),
							WithSalary:           false,
							Remote:               false,
						},
					},
				},
				domain.Rust: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Zig: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Scala: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Elixir: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Clojure: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Haskell: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Erlang: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
			},
			ShortDescription: "Music streaming service",
			Industries: []domain.Industry{
				domain.IndustryEntertainment,
			},
		},
		{
			ID:         0,  // system
			Type:       "", // system
			Name:       "Kraken",
			BaseURL:    "https://kraken.tech/",
			CareersURL: "https://kraken.tech/working-at-kraken",
			AboutURL:   "https://kraken.tech/about",
			BlogURL:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                68175052,
				IDs:               nil,
				Alias:             "krakentech",
				Name:              "Kraken",
				Followers:         "132K",
				Employees:         "1K-5K",
				AssociatedMembers: "1,984",
				Verified:          true,
				Date:              mustDate("2026-02-11"),
			},
			GitHubProfile: domain.GitHubProfile{
				Login:     "kraken-tech",
				Followers: "129",
				Verified:  true,
				Date:      mustDate("2026-02-11"),
			},
			BlindProfile: domain.BlindProfile{
				Alias: "",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias: "",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Percona-LLC-EI_IE283779.11,22.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/index.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Kraken-Jobs-E8860582.htm",
				Jobs:        "", // Becomes empty in the new design version
				Reviews:     "", // Becomes empty in the new design version
				Salaries:    "", // Becomes empty in the new design version
				ReviewsRate: "4.5",
				Verified:    true,
				Date:        mustDate("2026-02-11"),
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "",
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Rust: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Zig: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Scala: {
					GitHubRepositoryCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Senior Software Engineer – Scala",
							ShortDescription:     "Market Gateway",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4370934479/",
							Location:             "Melbourne, Victoria, Australia",
							CloudProviders:       []domain.CloudProvider{},
							Date:                 mustDate("2026-02-10"),
							WithSalary:           false,
							Remote:               true,
						},
					},
				},
				domain.Elixir: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Clojure: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Haskell: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Erlang: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
			},
			ShortDescription: "Operating system for energy utilities",
			Industries:       []domain.Industry{
				// NOP
			},
		},
		{
			ID:         0,  // system
			Type:       "", // system
			Name:       "ServiceNow",
			BaseURL:    "https://www.servicenow.com/",
			CareersURL: "https://careers.servicenow.com/",
			AboutURL:   "https://www.servicenow.com/company.html",
			BlogURL:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                29352,
				IDs:               []int{29352, 2978982, 7595344, 82334548},
				Alias:             "servicenow",
				Name:              "ServiceNow",
				Followers:         "1M",
				Employees:         "10K+",
				AssociatedMembers: "32,462",
				Verified:          true,
				Date:              mustDate("2026-02-11"),
			},
			GitHubProfile: domain.GitHubProfile{
				Login:     "servicenow",
				Followers: "885",
				Verified:  false,
				Date:      mustDate("2026-02-11"),
			},
			BlindProfile: domain.BlindProfile{
				Alias: "",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias: "",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Percona-LLC-EI_IE283779.11,22.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/index.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/ServiceNow-Jobs-E403326.htm",
				Jobs:        "", // Becomes empty in the new design version
				Reviews:     "", // Becomes empty in the new design version
				Salaries:    "", // Becomes empty in the new design version
				ReviewsRate: "4.2",
				Verified:    true,
				Date:        mustDate("2026-02-11"),
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "",
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoryCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Staff Software Engineer – Python & Golang",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4364881730/",
							Location:             "Hyderabad, Telangana, India",
							CloudProviders:       []domain.CloudProvider{},
							Date:                 mustDate("2026-02-10"),
							WithSalary:           false,
							Remote:               false,
						},
					},
				},
				domain.Rust: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Zig: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Scala: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Elixir: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Clojure: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Haskell: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Erlang: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
			},
			ShortDescription: "ServiceNow is a platform-as-a-service, that is designed to support IT service management and help desk functionality with automated workflows",
			Industries:       []domain.Industry{
				// NOP
			},
		},
		{
			ID:         0,  // system
			Type:       "", // system
			Name:       "Allianz Partners",
			BaseURL:    "https://www.allianz-partners.com/",
			CareersURL: "",
			AboutURL:   "",
			BlogURL:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                2255755,
				IDs:               []int{36034, 1088545, 2255755, 2894024, 3026041, 3276710, 3524677, 3778336, 10120220, 11207062, 19068744, 27126591, 27788082},
				Alias:             "allianz-partners",
				Name:              "Allianz Partners",
				Followers:         "247K",
				Employees:         "10K+",
				AssociatedMembers: "12,342",
				Verified:          true,
				Date:              mustDate("2026-02-11"),
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
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Percona-LLC-EI_IE283779.11,22.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/index.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Allianz-Partners-Jobs-E497552.htm",
				Jobs:        "", // Becomes empty in the new design version
				Reviews:     "", // Becomes empty in the new design version
				Salaries:    "", // Becomes empty in the new design version
				ReviewsRate: "3.7",
				Verified:    true,
				Date:        mustDate("2026-02-11"),
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "",
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoryCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Golang Software Engineer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4369786187/",
							Location:             "Lisbon, Lisbon, Portugal",
							CloudProviders:       []domain.CloudProvider{},
							Date:                 mustDate("2026-02-11"),
							WithSalary:           false,
							Remote:               false,
						},
					},
				},
				domain.Rust: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Zig: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Scala: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Elixir: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Clojure: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Haskell: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Erlang: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
			},
			ShortDescription: "Insurance and assistance company",
			Industries: []domain.Industry{
				domain.IndustryInsurTech,
			},
		},
		{
			ID:         0,  // system
			Type:       "", // system
			Name:       "adWMG",
			BaseURL:    "https://adwmg.com/",
			CareersURL: "https://adwmg.com/career/",
			AboutURL:   "",
			BlogURL:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                33203389,
				IDs:               nil,
				Alias:             "wmg-international",
				Name:              "adWMG",
				Followers:         "6K",
				Employees:         "51-200",
				AssociatedMembers: "47",
				Verified:          true,
				Date:              mustDate("2026-02-11"),
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
					GitHubRepositoryCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Middle Golang Developer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4363064533/",
							Location:             "Time, Rogaland, Norway",
							CloudProviders:       []domain.CloudProvider{},
							Date:                 mustDate("2026-02-10"),
							WithSalary:           false,
							Remote:               true,
						},
					},
				},
				domain.Rust: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Zig: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Scala: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Elixir: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Clojure: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Haskell: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Erlang: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
			},
			ShortDescription: "AdTech company specializing in programmatic game promotion and monetization",
			Industries: []domain.Industry{
				domain.IndustryAdTech,
			},
		},
		{
			ID:         0,  // system
			Type:       "", // system
			Name:       "Array",
			BaseURL:    "https://array.com/",
			CareersURL: "https://array.com/company/careers",
			AboutURL:   "https://array.com/company/about-us",
			BlogURL:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                68706605,
				IDs:               nil,
				Alias:             "array-io",
				Name:              "Array",
				Followers:         "19K",
				Employees:         "201-500",
				AssociatedMembers: "295",
				Verified:          true,
				Date:              mustDate("2026-02-11"),
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
					GitHubRepositoryCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Senior Go Software Engineer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4361873587/",
							Location:             "United States",
							CloudProviders:       []domain.CloudProvider{},
							Date:                 mustDate("2026-02-10"),
							WithSalary:           true, // $80/hr - $100/hr
							Remote:               true,
						},
					},
				},
				domain.Rust: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Zig: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Scala: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Elixir: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Clojure: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Haskell: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Erlang: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
			},
			ShortDescription: "Array is a financial platform that helps digital brands, financial institutions, and fintechs get compelling consumer products to market faster",
			Industries: []domain.Industry{
				domain.IndustryFinTech,
			},
		},
		{
			ID:         0,  // system
			Type:       "", // system
			Name:       "Emma – The Sleep Company",
			BaseURL:    "https://www.emma-sleep.com/",
			CareersURL: "",
			AboutURL:   "",
			BlogURL:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                26225188,
				IDs:               nil,
				Alias:             "lifeatemma",
				Name:              "Emma – The Sleep Company",
				Followers:         "78K",
				Employees:         "1K-5K",
				AssociatedMembers: "1,251",
				Verified:          true,
				Date:              mustDate("2026-02-11"),
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
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Percona-LLC-EI_IE283779.11,22.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/index.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Emma-Jobs-E2114470.htm",
				Jobs:        "", // Becomes empty in the new design version
				Reviews:     "", // Becomes empty in the new design version
				Salaries:    "", // Becomes empty in the new design version
				ReviewsRate: "3.5",
				Verified:    true,
				Date:        mustDate("2026-02-11"),
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "",
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoryCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "(Senior) Backend Software Engineer (Go / TypeScript)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4364275353/",
							Location:             "Frankfurt, Hesse, Germany",
							CloudProviders:       []domain.CloudProvider{domain.AWS},
							Date:                 mustDate("2026-02-10"),
							WithSalary:           false,
							Remote:               true,
						},
					},
				},
				domain.Rust: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Zig: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Scala: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Elixir: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Clojure: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Haskell: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Erlang: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
			},
			ShortDescription: "",
			Industries:       []domain.Industry{
				// NOP
			},
		},
		{
			ID:         0,  // system
			Type:       "", // system
			Name:       "RAPIDFORT",
			BaseURL:    "https://www.rapidfort.com/",
			CareersURL: "https://www.rapidfort.com/company/careers",
			AboutURL:   "https://www.rapidfort.com/company/about-us",
			BlogURL:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                74337267,
				IDs:               nil,
				Alias:             "rapidfort",
				Name:              "RAPIDFORT",
				Followers:         "24K",
				Employees:         "51-200",
				AssociatedMembers: "90",
				Verified:          true,
				Date:              mustDate("2026-02-11"),
			},
			GitHubProfile: domain.GitHubProfile{
				Login:     "rapidfort",
				Followers: "43",
				Verified:  false,
				Date:      mustDate("2026-02-11"),
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
					GitHubRepositoryCount: 1,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Microservice Backend Engineer (Python/Go)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4370711609/",
							Location:             "United States",
							CloudProviders:       []domain.CloudProvider{domain.AWS, domain.GCP, domain.Azure},
							Date:                 mustDate("2026-02-11"),
							WithSalary:           true, // $120k/yr - $200k/yr
							Remote:               true,
						},
					},
				},
				domain.Rust: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Zig: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Scala: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Elixir: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Clojure: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Haskell: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Erlang: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
			},
			ShortDescription: "Software supply chain security company that provides a platform designed to automatically secure container applications and accelerate compliance processes",
			Industries: []domain.Industry{
				domain.IndustryCyberSecurity,
			},
		},
		{
			ID:         0,  // system
			Type:       "", // system
			Name:       "Nosto",
			BaseURL:    "https://www.nosto.com/",
			CareersURL: "https://www.nosto.com/company/careers/",
			AboutURL:   "https://www.nosto.com/company/about/",
			BlogURL:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                2621568,
				IDs:               nil,
				Alias:             "nosto",
				Name:              "Nosto",
				Followers:         "14K",
				Employees:         "201-500",
				AssociatedMembers: "171",
				Verified:          true,
				Date:              mustDate("2026-02-14"),
			},
			GitHubProfile: domain.GitHubProfile{
				Login:     "nosto",
				Followers: "14",
				Verified:  false,
				Date:      mustDate("2026-02-15"),
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
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Rust: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Zig: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Scala: {
					GitHubRepositoryCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Senior Search Engineer, Scala",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4364423727/",
							Location:             "Helsinki, Uusimaa, Finland",
							CloudProviders:       []domain.CloudProvider{domain.AWS},
							Date:                 mustDate("2026-02-13"),
							WithSalary:           false,
							Remote:               false,
						},
					},
				},
				domain.Elixir: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Clojure: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Haskell: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Erlang: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
			},
			ShortDescription: "Commerce Experience Platform",
			Industries: []domain.Industry{
				domain.IndustryECommerce,
			},
		},
		{
			ID:         0,  // system
			Type:       "", // system
			Name:       "GridPoint",
			BaseURL:    "https://www.gridpoint.com/",
			CareersURL: "https://www.gridpoint.com/company/careers/",
			AboutURL:   "https://www.gridpoint.com/company/",
			BlogURL:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                362584,
				IDs:               nil,
				Alias:             "gridpoint",
				Name:              "GridPoint",
				Followers:         "8K",
				Employees:         "201-500",
				AssociatedMembers: "186",
				Verified:          true,
				Date:              mustDate("2026-02-14"),
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
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Rust: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Zig: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Scala: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Elixir: {
					GitHubRepositoryCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Senior Software Engineer – React Native, Elixir",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4370739065/",
							Location:             "United States",
							CloudProviders:       []domain.CloudProvider{},
							Date:                 mustDate("2026-02-11"),
							WithSalary:           false,
							Remote:               true,
						},
					},
				},
				domain.Clojure: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Haskell: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Erlang: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
			},
			ShortDescription: "Energy intelligence technology that helps commercial buildings reduce energy consumption",
			Industries:       []domain.Industry{
				// NOP
			},
		},
		{
			ID:         0,  // system
			Type:       "", // system
			Name:       "VIAVI Solutions",
			BaseURL:    "https://www.viavisolutions.com/",
			CareersURL: "https://www.viavisolutions.com/corporate/careers",
			AboutURL:   "https://www.viavisolutions.com/corporate/about-us",
			BlogURL:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                3218,
				IDs:               []int{3218, 100688, 787185},
				Alias:             "viavi-solutions",
				Name:              "VIAVI Solutions",
				Followers:         "107K",
				Employees:         "1K-5K",
				AssociatedMembers: "4,984",
				Verified:          true,
				Date:              mustDate("2026-02-15"),
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
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Rust: {
					GitHubRepositoryCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Software Development Engineer II (Rust)",
							ShortDescription:     "Enterprise AI Platform",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4370703961/",
							Location:             "Hopkins, MN",
							CloudProviders:       []domain.CloudProvider{},
							Date:                 mustDate("2026-02-11"),
							WithSalary:           false,
							Remote:               false,
						},
					},
				},
				domain.Zig: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Scala: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Elixir: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Clojure: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Haskell: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Erlang: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
			},
			ShortDescription: "VIAVI Solutions develops network enablement and optical security performance solutions",
			Industries: []domain.Industry{
				domain.IndustryTelecom,
			},
		},
		{
			ID:         0,  // system
			Type:       "", // system
			Name:       "Machinify",
			BaseURL:    "https://www.machinify.com/",
			CareersURL: "https://www.machinify.com/careers",
			AboutURL:   "https://www.machinify.com/about",
			BlogURL:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                7596960,
				IDs:               nil,
				Alias:             "machinify",
				Name:              "Machinify",
				Followers:         "14K",
				Employees:         "1K-5K",
				AssociatedMembers: "723",
				Verified:          true,
				Date:              mustDate("2026-02-15"),
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
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Percona-LLC-EI_IE283779.11,22.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/index.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Machinify-Jobs-E3102945.htm",
				Jobs:        "", // Becomes empty in the new design version
				Reviews:     "", // Becomes empty in the new design version
				Salaries:    "", // Becomes empty in the new design version
				ReviewsRate: "2.3",
				Verified:    true,
				Date:        mustDate("2026-02-15"),
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "",
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Rust: {
					GitHubRepositoryCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Rust Developer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4364140207/",
							Location:             "United States",
							CloudProviders:       []domain.CloudProvider{},
							Date:                 mustDate("2026-02-13"),
							WithSalary:           true, // $215k/yr - $245k/yr
							Remote:               true,
						},
					},
				},
				domain.Zig: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Scala: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Elixir: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Clojure: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Haskell: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Erlang: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
			},
			ShortDescription: "",
			Industries:       []domain.Industry{
				// NOP
			},
		},
		{
			ID:         0,  // system
			Type:       "", // system
			Name:       "Synamedia",
			BaseURL:    "https://www.synamedia.com/",
			CareersURL: "https://www.synamedia.com/careers/",
			AboutURL:   "https://www.synamedia.com/about-us/",
			BlogURL:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                33194934,
				IDs:               nil,
				Alias:             "synamedia",
				Name:              "Synamedia",
				Followers:         "79K",
				Employees:         "1K-5K",
				AssociatedMembers: "1,463",
				Verified:          true,
				Date:              mustDate("2026-02-15"),
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
					GitHubRepositoryCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Senior Software Engineer – Golang",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4371419072/",
							Location:             "Bengaluru, Karnataka, India",
							CloudProviders:       []domain.CloudProvider{},
							Date:                 mustDate("2026-02-12"),
							WithSalary:           false,
							Remote:               false,
						},
					},
				},
				domain.Rust: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Zig: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Scala: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Elixir: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Clojure: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Haskell: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Erlang: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
			},
			ShortDescription: "",
			Industries:       []domain.Industry{
				// NOP
			},
		},
		{
			ID:         0,  // system
			Type:       "", // system
			Name:       "Ocrolus",
			BaseURL:    "https://www.ocrolus.com/",
			CareersURL: "https://www.ocrolus.com/careers/",
			AboutURL:   "https://www.ocrolus.com/company/",
			BlogURL:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                3651177,
				IDs:               nil,
				Alias:             "ocrolus",
				Name:              "Ocrolus",
				Followers:         "32K",
				Employees:         "201-500",
				AssociatedMembers: "918",
				Verified:          true,
				Date:              mustDate("2026-02-15"),
			},
			GitHubProfile: domain.GitHubProfile{
				Login:     "Ocrolus",
				Followers: "9",
				Verified:  true,
				Date:      mustDate("2026-02-15"),
			},
			BlindProfile: domain.BlindProfile{
				Alias: "",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias: "",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Percona-LLC-EI_IE283779.11,22.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/index.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Ocrolus-Jobs-E2857732.htm",
				Jobs:        "", // Becomes empty in the new design version
				Reviews:     "", // Becomes empty in the new design version
				Salaries:    "", // Becomes empty in the new design version
				ReviewsRate: "3.7",
				Verified:    true,
				Date:        mustDate("2026-02-15"),
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "",
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoryCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Senior Software Engineer (Golang)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4372171272/",
							Location:             "New York, NY",
							CloudProviders:       []domain.CloudProvider{},
							Date:                 mustDate("2026-02-14"),
							WithSalary:           false,
							Remote:               true,
						},
					},
				},
				domain.Rust: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Zig: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Scala: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Elixir: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Clojure: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Haskell: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Erlang: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
			},
			ShortDescription: "Ocrolus is an automation and analytics platform for financial decision-making",
			Industries: []domain.Industry{
				domain.IndustryFinTech,
			},
		},
		{
			ID:         0,  // system
			Type:       "", // system
			Name:       "MirrorWeb",
			BaseURL:    "https://www.mirrorweb.com/",
			CareersURL: "https://www.mirrorweb.com/careers/",
			AboutURL:   "https://www.mirrorweb.com/why-mirrorweb/about",
			BlogURL:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                3035248,
				IDs:               nil,
				Alias:             "mirrorweb",
				Name:              "MirrorWeb",
				Followers:         "6K",
				Employees:         "51-200",
				AssociatedMembers: "77",
				Verified:          true,
				Date:              mustDate("2026-02-15"),
			},
			GitHubProfile: domain.GitHubProfile{
				Login:     "mirrorweb",
				Followers: "7",
				Verified:  false,
				Date:      mustDate("2026-02-15"),
			},
			BlindProfile: domain.BlindProfile{
				Alias: "",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias: "",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Percona-LLC-EI_IE283779.11,22.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/index.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/MirrorWeb-Jobs-E2347924.htm",
				Jobs:        "", // Becomes empty in the new design version
				Reviews:     "", // Becomes empty in the new design version
				Salaries:    "", // Becomes empty in the new design version
				ReviewsRate: "3.6",
				Verified:    false,
				Date:        mustDate("2026-02-15"),
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "",
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoryCount: 6,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Senior Software Engineer (Go)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4371181676/",
							Location:             "Manchester, England, United Kingdom",
							CloudProviders:       []domain.CloudProvider{domain.AWS},
							Date:                 mustDate("2026-02-12"),
							WithSalary:           false,
							Remote:               false,
						},
					},
				},
				domain.Rust: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Zig: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Scala: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Elixir: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Clojure: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Haskell: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Erlang: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
			},
			ShortDescription: "Through the MirrorWeb Insight platform we capture, archive and monitor web, SMS, email, instant messaging and social media channels, ensuring digital communication regulations are met and content remains compliant.",
			Industries:       []domain.Industry{
				// NOP
			},
		},
		{
			ID:         0,  // system
			Type:       "", // system
			Name:       "Wells Fargo",
			BaseURL:    "https://www.wellsfargo.com/",
			CareersURL: "https://www.wellsfargo.com/about/careers/",
			AboutURL:   "https://www.wellsfargo.com/about/",
			BlogURL:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                1235,
				IDs:               []int{1235, 1409},
				Alias:             "wellsfargo",
				Name:              "Wells Fargo",
				Followers:         "3M",
				Employees:         "10K+",
				AssociatedMembers: "224,205",
				Verified:          true,
				Date:              mustDate("2026-02-15"),
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
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Percona-LLC-EI_IE283779.11,22.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/index.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Wells-Fargo-Jobs-E8876.htm",
				Jobs:        "", // Becomes empty in the new design version
				Reviews:     "", // Becomes empty in the new design version
				Salaries:    "", // Becomes empty in the new design version
				ReviewsRate: "3.6",
				Verified:    true,
				Date:        mustDate("2026-02-15"),
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "",
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoryCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Lead Protocol Engineer (Go / Cryptography)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4372242028/",
							Location:             "Concord, CA",
							CloudProviders:       []domain.CloudProvider{},
							Date:                 mustDate("2026-02-14"),
							WithSalary:           false,
							Remote:               false,
						},
					},
				},
				domain.Rust: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Zig: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Scala: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Elixir: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Clojure: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Haskell: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Erlang: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
			},
			ShortDescription: "Community-based financial services company",
			Industries: []domain.Industry{
				domain.IndustryFinTech,
			},
		},
		{
			ID:         0,  // system
			Type:       "", // system
			Name:       "Trade Republic",
			BaseURL:    "https://traderepublic.com/",
			CareersURL: "",
			AboutURL:   "",
			BlogURL:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                13014668,
				IDs:               nil,
				Alias:             "trade-republic",
				Name:              "Trade Republic",
				Followers:         "147K",
				Employees:         "501-1K",
				AssociatedMembers: "1,332",
				Verified:          true,
				Date:              mustDate("2026-02-15"),
			},
			GitHubProfile: domain.GitHubProfile{
				Login:     "traderepublic",
				Followers: "198",
				Verified:  true,
				Date:      mustDate("2026-02-15"),
			},
			BlindProfile: domain.BlindProfile{
				Alias: "",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias: "",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Percona-LLC-EI_IE283779.11,22.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/index.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Trade-Republic-Jobs-E3554739.htm",
				Jobs:        "", // Becomes empty in the new design version
				Reviews:     "", // Becomes empty in the new design version
				Salaries:    "", // Becomes empty in the new design version
				ReviewsRate: "2.7",
				Verified:    true,
				Date:        mustDate("2026-02-15"),
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "",
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoryCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "(Senior) Platform Engineer (Go)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4372008252/",
							Location:             "Berlin, Germany",
							CloudProviders:       []domain.CloudProvider{},
							Date:                 mustDate("2026-02-14"),
							WithSalary:           false,
							Remote:               false,
						},
					},
				},
				domain.Rust: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Zig: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Scala: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Elixir: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Clojure: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Haskell: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Erlang: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
			},
			ShortDescription: "",
			Industries: []domain.Industry{
				domain.IndustryFinTech,
			},
		},
		{
			ID:         0,  // system
			Type:       "", // system
			Name:       "Arkeus",
			BaseURL:    "https://arkeus.com/",
			CareersURL: "https://arkeus.com/careers",
			AboutURL:   "https://arkeus.com/who-we-are",
			BlogURL:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                42907079,
				IDs:               nil,
				Alias:             "arkeus-sensors",
				Name:              "Arkeus",
				Followers:         "7K",
				Employees:         "11-50",
				AssociatedMembers: "57",
				Verified:          false,
				Date:              mustDate("2026-02-17"),
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
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Rust: {
					GitHubRepositoryCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Senior Software Engineer (Rust)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4364708193/",
							Location:             "Port Melbourne, Victoria, Australia",
							CloudProviders:       []domain.CloudProvider{},
							Date:                 mustDate("2026-02-17"),
							WithSalary:           false,
							Remote:               false,
						},
					},
				},
				domain.Zig: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Scala: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Elixir: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Clojure: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Haskell: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Erlang: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
			},
			ShortDescription: "Arkeus wide area autonomous optical search and surveillance systems",
			Industries:       []domain.Industry{
				// NOP
			},
		},
	}
}
