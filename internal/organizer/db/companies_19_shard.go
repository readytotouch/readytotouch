package db

import "github.com/readytotouch/readytotouch/internal/domain"

func companies19Shard() []domain.CompanyProfile {
	return []domain.CompanyProfile{
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Valence Security",
			Website: "https://www.valencesecurity.com/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                76479154,
				IDs:               nil,
				Alias:             "valence-security",
				Name:              "Valence Security",
				Followers:         "7K",
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
							URL:                  "https://www.linkedin.com/jobs/view/4253472167/",
							Date:                 mustDate("2025-06-23"),
							WithSalary:           false,
							Remote:               false,
						},
					},
				},
				domain.Rust: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Zig: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Scala: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Elixir: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Clojure: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Haskell: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
			},
			ShortDescription: "Valence finds and fixes SaaS risks",
			Industries: []domain.Industry{
				domain.IndustryCyberSecurity,
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Claroty",
			Website: "https://www.claroty.com/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                12898104,
				IDs:               []int{12898104, 18340376},
				Alias:             "claroty",
				Name:              "Claroty",
				Followers:         "57K",
				Employees:         "501-1K",
				AssociatedMembers: "811",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:     "claroty",
				Followers: "121",
				Verified:  false,
			},
			BlindProfile: domain.BlindProfile{
				Alias: "",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "claroty",
				Employees: "351",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Claroty-EI_IE2534828.11,18.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Claroty-Reviews-E2534828.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Claroty-Jobs-E2534828.htm",
				Jobs:        "",
				Reviews:     "94",
				Salaries:    "138",
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
							Title:                "Full Stack Software Engineer",
							ShortDescription:     "5+ years of experience developing server applications in Python/Go — Must",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4255769021/",
							Date:                 mustDate("2025-06-30"),
							WithSalary:           false,
							Remote:               false,
						},
					},
				},
				domain.Rust: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Zig: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Scala: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Elixir: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Clojure: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Haskell: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
			},
			ShortDescription: "Cyber-physical systems protection",
			Industries: []domain.Industry{
				domain.IndustryCyberSecurity,
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Aleph Alpha",
			Website: "https://aleph-alpha.com/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                20143683,
				IDs:               nil,
				Alias:             "aleph-alpha",
				Name:              "Aleph Alpha",
				Followers:         "62K",
				Employees:         "201-500",
				AssociatedMembers: "327",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:     "Aleph-Alpha",
				Followers: "639",
				Verified:  false,
			},
			BlindProfile: domain.BlindProfile{
				Alias: "",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias: "",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Aleph-Alpha-EI_IE7593393.11,22.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Aleph-Alpha-Reviews-E7593393.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Aleph-Alpha-Jobs-E7593393.htm",
				Jobs:        "30",
				Reviews:     "9",
				Salaries:    "20",
				ReviewsRate: "3.2",
				Verified:    true,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "",
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Rust: {
					GitHubRepositoriesCount: 6,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Senior Software Engineer — Rust",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4253172194/",
							Date:                 mustDate("2025-07-10"), // mustDate("2025-06-23"),
							WithSalary:           false,
							Remote:               false,
						},
					},
				},
				domain.Zig: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Scala: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Elixir: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Clojure: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Haskell: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
			},
			ShortDescription: "Our full-stack solution, PhariaAI, enables enterprises and public institutions to build sovereign, secure, and trustworthy applications – without compromising control",
			Industries:       []domain.Industry{
				// NOP
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Perkbox",
			Website: "https://www.perkbox.com/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                9474278,
				IDs:               nil,
				Alias:             "perkbox",
				Name:              "Perkbox",
				Followers:         "44K",
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
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Perkbox-EI_IE1195707.11,18.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Perkbox-Reviews-E1195707.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Perkbox-Jobs-E1195707.htm",
				Jobs:        "4",
				Reviews:     "161",
				Salaries:    "336",
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
							Title:                "Go Developer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4246318989/",
							Date:                 mustDate("2025-06-16"),
							WithSalary:           false,
							Remote:               true,
						},
					},
				},
				domain.Rust: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Zig: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Scala: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Elixir: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Clojure: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Haskell: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
			},
			ShortDescription: "Rewards and benefits platform that helps companies to care for, connect with and celebrate their employees",
			Industries:       []domain.Industry{
				// NOP
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Underdog",
			Website: "https://underdogfantasy.com/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                49173095,
				IDs:               nil,
				Alias:             "underdogfantasy",
				Name:              "Underdog",
				Followers:         "59K",
				Employees:         "201-500",
				AssociatedMembers: "577",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:     "underdog-inc",
				Followers: "38",
				Verified:  true,
			},
			BlindProfile: domain.BlindProfile{
				Alias: "",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "underdog-fantasy",
				Employees: "351",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Underdog-Fantasy-EI_IE6372418.11,27.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Underdog-Fantasy-Reviews-E6372418.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Underdog-Fantasy-Jobs-E6372418.htm",
				Jobs:        "",
				Reviews:     "80",
				Salaries:    "110",
				ReviewsRate: "3.8",
				Verified:    true,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "",
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 1,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Software Engineer – Go",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4122421045/",
							Date:                 mustDate("2025-06-16"),
							WithSalary:           true,
							Remote:               true,
						},
					},
				},
				domain.Rust: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Zig: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Scala: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Elixir: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Clojure: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Haskell: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
			},
			ShortDescription: "Fastest-growing fantasy sports",
			Industries:       []domain.Industry{
				// NOP
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Glasskube",
			Website: "https://glasskube.dev/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                83055300,
				IDs:               nil,
				Alias:             "glasskube",
				Name:              "Glasskube",
				Followers:         "2K",
				Employees:         "2-10",
				AssociatedMembers: "7",
				Verified:          false,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:     "glasskube",
				Followers: "115",
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
					GitHubRepositoriesCount: 3,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Founding AI Engineer (Go, TypeScript, Kubernetes, Docker)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4249315524/",
							Date:                 mustDate("2025-06-16"),
							WithSalary:           true,
							Remote:               false,
						},
					},
				},
				domain.Rust: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Zig: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Scala: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Elixir: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Clojure: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Haskell: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
			},
			ShortDescription: "The Open Source control plane for self-managed, BYOC, and on-prem deployments",
			Industries: []domain.Industry{
				domain.IndustryDevOps,
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Litmus",
			Website: "https://litmus.io/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                5220045,
				IDs:               nil,
				Alias:             "litmus-automation",
				Name:              "Litmus",
				Followers:         "22K",
				Employees:         "51-200",
				AssociatedMembers: "118",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:     "litmusautomation",
				Followers: "23",
				Verified:  true,
			},
			BlindProfile: domain.BlindProfile{
				Alias: "",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "litmus-automation",
				Employees: "126",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Litmus-Automation-EI_IE1863601.11,28.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Litmus-Automation-Reviews-E1863601.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Litmus-Automation-Jobs-E1863601.htm",
				Jobs:        "9",
				Reviews:     "16",
				Salaries:    "20",
				ReviewsRate: "4.6",
				Verified:    true,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "",
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 4,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Senior Golang Developer",
							ShortDescription:     "Data Platform",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4250646664/",
							Date:                 mustDate("2025-06-23"),
							WithSalary:           true,
							Remote:               false,
						},
					},
				},
				domain.Rust: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Zig: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Scala: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Elixir: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Clojure: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Haskell: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
			},
			ShortDescription: "Data Platform that unifies device connectivity, data intelligence and data integration in a complete Industry 4.0 solution",
			Industries: []domain.Industry{
				domain.IndustryDevOps,
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "blockit",
			Website: "https://www.blockitnow.com/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                16188895,
				IDs:               nil,
				Alias:             "blockit",
				Name:              "blockit",
				Followers:         "2K",
				Employees:         "51-200",
				AssociatedMembers: "58",
				Verified:          false,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:     "blockit",
				Followers: "9",
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
					Vacancies:               []domain.Vacancy{},
				},
				domain.Rust: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Zig: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Scala: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Elixir: {
					GitHubRepositoriesCount: 5,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Senior Software Engineer, Elixir",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4263488829/",
							Date:                 mustDate("2025-07-07"),
							WithSalary:           true, // $141k/yr - $165k/yr
							Remote:               false,
						},
					},
				},
				domain.Clojure: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Haskell: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Erlang: {
					GitHubRepositoriesCount: 2,
					Vacancies:               []domain.Vacancy{},
				},
			},
			ShortDescription: "We eliminate barriers between providers and patients with our comprehensive platform featuring digital front door capabilities, patient self-scheduling, referral management, and care coordination tools",
			Industries: []domain.Industry{
				domain.IndustryMedTech,
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Experian",
			Website: "https://www.experianplc.com/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                2534,
				IDs:               []int{2534, 685433, 1133358, 73804409},
				Alias:             "experian",
				Name:              "Experian",
				Followers:         "1M",
				Employees:         "10K+",
				AssociatedMembers: "24,170",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:     "experianplc",
				Followers: "22",
				Verified:  false,
			},
			BlindProfile: domain.BlindProfile{
				Alias:       "Experian",
				Employees:   "10,000+",
				Salary:      "$50K ~ $240K a year",
				Reviews:     "108",
				ReviewsRate: "3.4",
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
				domain.Rust: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Zig: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Scala: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Senior Software Engineer – (Kafka, Java, Scala and Spark)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4263607684/",
							Date:                 mustDate("2025-07-07"),
							WithSalary:           false,
							Remote:               false,
						},
					},
				},
				domain.Elixir: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Clojure: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Haskell: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
			},
			ShortDescription: "Data and technology company",
			Industries:       []domain.Industry{
				// NOP
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Thomson Reuters",
			Website: "https://www.thomsonreuters.com/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                1400,
				IDs:               []int{1400, 14826, 2487151, 16197320, 67237271},
				Alias:             "thomson-reuters",
				Name:              "Thomson Reuters",
				Followers:         "2M",
				Employees:         "10K+",
				AssociatedMembers: "35,779",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:     "thomsonreuters",
				Followers: "58",
				Verified:  true,
			},
			BlindProfile: domain.BlindProfile{
				Alias:       "Thomson-Reuters",
				Employees:   "10,000+",
				Salary:      "$47K ~ $250K a year",
				Reviews:     "149",
				ReviewsRate: "3.7",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias: "",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Thomson-Reuters-EI_IE100303.11,26.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Thomson-Reuters-Reviews-E100303.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Thomson-Reuters-Jobs-E100303.htm",
				Jobs:        "1.3K",
				Reviews:     "18K",
				Salaries:    "24K",
				ReviewsRate: "4.1",
				Verified:    true,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "",
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Rust: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Zig: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Scala: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Senior Software Engineer – Java/Scala",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4235994956/",
							Date:                 mustDate("2025-07-07"),
							WithSalary:           false,
							Remote:               false,
						},
					},
				},
				domain.Elixir: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Clojure: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Haskell: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
			},
			ShortDescription: "Our products combine specialized software and insights to empower professionals with the data, intelligence, and solutions needed to make informed decisions, and to help institutions in their pursuit of justice, truth, and transparency",
			Industries:       []domain.Industry{
				// NOP
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Aiven",
			Website: "https://aiven.io/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                10294984,
				IDs:               nil,
				Alias:             "aiven",
				Name:              "Aiven",
				Followers:         "55K",
				Employees:         "201-500",
				AssociatedMembers: "439",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:     "aiven",
				Followers: "295",
				Verified:  false,
			},
			BlindProfile: domain.BlindProfile{
				Alias:       "Aiven",
				Employees:   "501 to 1,000",
				Salary:      "",
				Reviews:     "8",
				ReviewsRate: "3.1",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias: "",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Aiven-EI_IE2610934.11,16.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Aiven-Reviews-E2610934.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Aiven-Jobs-E2610934.htm",
				Jobs:        "1",
				Reviews:     "156",
				Salaries:    "341",
				ReviewsRate: "2.5",
				Verified:    true,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "",
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 18,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Senior Software Engineer (Golang)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4240055195/",
							Date:                 mustDate("2025-07-07"),
							WithSalary:           false,
							Remote:               false,
						},
					},
				},
				domain.Rust: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Zig: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Scala: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Elixir: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Clojure: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Haskell: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
			},
			ShortDescription: "AI-ready open source data platform",
			Industries: []domain.Industry{
				domain.IndustryDevOps,
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Manus AI",
			Website: "https://manus.im/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                106664974,
				IDs:               nil,
				Alias:             "manus-im",
				Name:              "Manus AI",
				Followers:         "48K",
				Employees:         "51-200",
				AssociatedMembers: "48",
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
							URL:                  "https://www.linkedin.com/jobs/view/4263265388/",
							Date:                 mustDate("2025-07-07"),
							WithSalary:           false,
							Remote:               false,
						},
					},
				},
				domain.Rust: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Zig: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Scala: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Elixir: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Clojure: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Haskell: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
			},
			ShortDescription: "General AI agent",
			Industries:       []domain.Industry{
				// NOP
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Stori",
			Website: "https://www.storicard.com/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                40716583,
				IDs:               nil,
				Alias:             "stori-card",
				Name:              "Stori",
				Followers:         "43K",
				Employees:         "501-1K",
				AssociatedMembers: "758",
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
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Stori-Card-EI_IE3357578.11,21.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Stori-Card-Reviews-E3357578.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Stori-Card-Jobs-E3357578.htm",
				Jobs:        "",
				Reviews:     "211",
				Salaries:    "249",
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
							Title:                "Lead Engineer (Golang)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4263453034/",
							Date:                 mustDate("2025-07-07"),
							WithSalary:           false,
							Remote:               false,
						},
					},
				},
				domain.Rust: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Zig: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Scala: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Elixir: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Clojure: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Haskell: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
			},
			ShortDescription: "Venture-backed financial technology company",
			Industries: []domain.Industry{
				domain.IndustryFinTech,
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Testsigma",
			Website: "https://testsigma.com/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                13359845,
				IDs:               nil,
				Alias:             "testsigma",
				Name:              "Testsigma",
				Followers:         "29K",
				Employees:         "51-200",
				AssociatedMembers: "236",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:     "testsigmahq",
				Followers: "105",
				Verified:  false,
			},
			BlindProfile: domain.BlindProfile{
				Alias: "",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias: "",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Testsigma-EI_IE2048845.11,20.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Testsigma-Reviews-E2048845.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Testsigma-Jobs-E2048845.htm",
				Jobs:        "",
				Reviews:     "76",
				Salaries:    "91",
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
							Title:                "Lead Engineer – Golang",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4262157113/",
							Date:                 mustDate("2025-07-07"),
							WithSalary:           false,
							Remote:               false,
						},
					},
				},
				domain.Rust: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Zig: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Scala: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Elixir: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Clojure: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Haskell: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
			},
			ShortDescription: "Agentic Test Automation Platform",
			Industries: []domain.Industry{
				domain.IndustryDevOps,
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Wolters Kluwer",
			Website: "https://www.wolterskluwer.com/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                2483,
				IDs:               []int{2481, 2483, 2487, 6013, 21924, 97171, 131097, 224387, 290256, 343128, 599146, 640280, 1100500, 1356211, 2034209, 2330375, 2380797, 2405141, 2478425, 2659281, 3513007},
				Alias:             "wolters-kluwer",
				Name:              "Wolters Kluwer",
				Followers:         "312K",
				Employees:         "10K+",
				AssociatedMembers: "23,230",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "",
				Verified: false,
			},
			BlindProfile: domain.BlindProfile{
				Alias:       "Wolters-Kluwer",
				Employees:   "10,000+",
				Salary:      "$34K ~ $240K a year",
				Reviews:     "65",
				ReviewsRate: "3.4",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias: "",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Wolters-Kluwer-EI_IE10497.11,25.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Wolters-Kluwer-Reviews-E10497.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Wolters-Kluwer-Jobs-E10497.htm",
				Jobs:        "372",
				Reviews:     "4.7K",
				Salaries:    "6.9K",
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
							Title:                "Golang Software Engineer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4262168890/",
							Date:                 mustDate("2025-07-07"),
							WithSalary:           false,
							Remote:               false,
						},
					},
				},
				domain.Rust: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Zig: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Scala: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Elixir: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Clojure: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Haskell: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
			},
			ShortDescription: "",
			Industries:       []domain.Industry{
				// NOP
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Engageware",
			Website: "https://engageware.com/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                61247,
				IDs:               nil,
				Alias:             "engageware",
				Name:              "Engageware",
				Followers:         "8K",
				Employees:         "201-500",
				AssociatedMembers: "122",
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
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Engageware-EI_IE6093728.11,21.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Engageware-Reviews-E6093728.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Engageware-Jobs-E6093728.htm",
				Jobs:        "3",
				Reviews:     "22",
				Salaries:    "39",
				ReviewsRate: "2.5",
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
							Title:                "Senior Backend Developer, Golang",
							ShortDescription:     "Aivo",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4243309060/",
							Date:                 mustDate("2025-07-07"),
							WithSalary:           false,
							Remote:               true,
						},
					},
				},
				domain.Rust: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Zig: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Scala: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Elixir: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Clojure: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Haskell: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
			},
			ShortDescription: "Customer engagement platform",
			Industries:       []domain.Industry{
				// NOP
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Mercuryo",
			Website: "https://mercuryo.io/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                12993966,
				IDs:               nil,
				Alias:             "mercuryo-io",
				Name:              "Mercuryo",
				Followers:         "24K",
				Employees:         "201-500",
				AssociatedMembers: "270",
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
							URL:                  "https://www.linkedin.com/jobs/view/4263480260/",
							Date:                 mustDate("2025-07-07"),
							WithSalary:           false,
							Remote:               false,
						},
					},
				},
				domain.Rust: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Zig: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Scala: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Elixir: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Clojure: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Haskell: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
			},
			ShortDescription: "Payments for web3 dApps",
			Industries:       []domain.Industry{
				// NOP
			},
			Ignore: true,
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Parallel Wireless",
			Website: "https://www.parallelwireless.com/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                2875687,
				IDs:               nil,
				Alias:             "parallel-wireless-inc",
				Name:              "Parallel Wireless",
				Followers:         "65K",
				Employees:         "501-1K",
				AssociatedMembers: "722",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "",
				Verified: false,
			},
			BlindProfile: domain.BlindProfile{
				Alias:       "Parallel-Wireless",
				Employees:   "501 to 1,000",
				Salary:      "",
				Reviews:     "10",
				ReviewsRate: "3.2",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias: "",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Parallel-Wireless-EI_IE1369255.11,28.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Parallel-Wireless-Reviews-E1369255.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Parallel-Wireless-Jobs-E1369255.htm",
				Jobs:        "53",
				Reviews:     "219",
				Salaries:    "328",
				ReviewsRate: "3.4",
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
							Title:                "Senior Cloud Applications Engineer (C/C++/Go)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4262117247/",
							Date:                 mustDate("2025-07-07"),
							WithSalary:           false,
							Remote:               false,
						},
					},
				},
				domain.Rust: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Zig: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Scala: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Elixir: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Clojure: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Haskell: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
			},
			ShortDescription: "We disrupt the ways wireless networks are built and operated",
			Industries: []domain.Industry{
				domain.IndustryTelecom,
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "MarkeTeam.ai",
			Website: "https://www.marketeam.ai/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                71279421,
				IDs:               nil,
				Alias:             "marketeam-ai",
				Name:              "MarkeTeam.ai",
				Followers:         "7K",
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
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Rust: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Zig: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Scala: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Elixir: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Elixir Developer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4260363542/",
							Date:                 mustDate("2025-07-08"),
							WithSalary:           false,
							Remote:               true,
						},
					},
				},
				domain.Clojure: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Haskell: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
			},
			ShortDescription: "Virtual Marketing Team",
			Industries: []domain.Industry{
				domain.IndustryMarTech,
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "White Hat Gaming",
			Website: "https://www.whitehatgaming.com/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                10775622,
				IDs:               nil,
				Alias:             "white-hat-gaming",
				Name:              "White Hat Gaming",
				Followers:         "14K",
				Employees:         "501-1K",
				AssociatedMembers: "243",
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
					Vacancies:               []domain.Vacancy{},
				},
				domain.Rust: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Zig: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Scala: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Scala Developer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4240642949/",
							Date:                 mustDate("2025-07-08"),
							WithSalary:           false,
							Remote:               true,
						},
					},
				},
				domain.Elixir: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Clojure: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Haskell: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
			},
			ShortDescription: "iGaming platform",
			Industries:       []domain.Industry{
				// Casino
			},
			Ignore: true, // Casino
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "WhiteFox Defense Technologies, Inc.",
			Website: "https://www.whitefoxdefense.com/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                17943090,
				IDs:               nil,
				Alias:             "whitefox-defense-technologies-inc.",
				Name:              "WhiteFox Defense Technologies, Inc.",
				Followers:         "5K",
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
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-WhiteFox-Defense-Technologies-EI_IE1831870.11,40.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/WhiteFox-Defense-Technologies-Reviews-E1831870.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/WhiteFox-Defense-Technologies-Jobs-E1831870.htm",
				Jobs:        "",
				Reviews:     "20",
				Salaries:    "26",
				ReviewsRate: "3.6",
				Verified:    false,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "",
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Rust: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Software Engineer II (Rust Focused)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4263131418/",
							Date:                 mustDate("2025-07-09"),
							WithSalary:           true, // $125k/yr - $135k/yr
							Remote:               true,
						},
					},
				},
				domain.Zig: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Scala: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Elixir: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Clojure: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Haskell: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
			},
			ShortDescription: "Counter-drone and airspace management industry",
			Industries: []domain.Industry{
				domain.IndustryDefenseTech,
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Mizzle",
			Website: "https://mizzle.io/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                103293432,
				IDs:               nil,
				Alias:             "mizzlecloud",
				Name:              "Mizzle",
				Followers:         "3K",
				Employees:         "11-50",
				AssociatedMembers: "32",
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
					Vacancies:               []domain.Vacancy{},
				},
				domain.Rust: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Rust developer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4262397768/",
							Date:                 mustDate("2025-07-08"),
							WithSalary:           false,
							Remote:               false,
						},
					},
				},
				domain.Zig: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Scala: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Elixir: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Clojure: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Haskell: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
			},
			ShortDescription: "Decentralized cloud platform",
			Industries: []domain.Industry{
				domain.IndustryDevOps,
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Nuitée",
			Website: "https://nuitee.com/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                18032726,
				IDs:               nil,
				Alias:             "nuitee",
				Name:              "Nuitée",
				Followers:         "56K",
				Employees:         "51-200",
				AssociatedMembers: "229",
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
							URL:                  "https://www.linkedin.com/jobs/view/4262885265/",
							Date:                 mustDate("2025-07-08"),
							WithSalary:           false,
							Remote:               true,
						},
					},
				},
				domain.Rust: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Zig: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Scala: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Elixir: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Clojure: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Haskell: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
			},
			ShortDescription: "Nuitée provides infrastructure APIs tailored for the travel sector, offering a user-friendly solution to transform hotel connectivity and distribution",
			Industries:       []domain.Industry{
				// NOP
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "NFON AG",
			Website: "https://www.nfon.com/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                934860,
				IDs:               []int{934860, 103760758, 104286793, 104291144, 104294333},
				Alias:             "nfon",
				Name:              "NFON AG",
				Followers:         "13K",
				Employees:         "501-1K",
				AssociatedMembers: "369",
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
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-NFON-EI_IE947578.11,15.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/NFON-Reviews-E947578.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/NFON-Jobs-E947578.htm",
				Jobs:        "12",
				Reviews:     "62",
				Salaries:    "111",
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
							Title:                "Senior Full-Stack Engineer Golang",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4260732659/",
							Date:                 mustDate("2025-07-08"),
							WithSalary:           false,
							Remote:               false,
						},
					},
				},
				domain.Rust: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Zig: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Scala: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Elixir: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Clojure: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Haskell: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
			},
			ShortDescription: "",
			Industries: []domain.Industry{
				domain.IndustryTelecom,
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Milestone Systems",
			Website: "https://www.milestonesys.com/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                9977,
				IDs:               []int{9977, 202961, 18081691},
				Alias:             "milestone-systems",
				Name:              "Milestone Systems",
				Followers:         "129K",
				Employees:         "1K-5K",
				AssociatedMembers: "1,846",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:     "milestonesys",
				Followers: "99",
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
					GitHubRepositoriesCount: 1,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Lead Golang Software Engineer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4264183942/",
							Date:                 mustDate("2025-07-08"),
							WithSalary:           true, // $175k/yr - $195k/yr
							Remote:               false,
						},
					},
				},
				domain.Rust: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Zig: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Scala: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Elixir: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Clojure: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Haskell: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
			},
			ShortDescription: "Video technology software",
			Industries:       []domain.Industry{
				// NOP
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Eltropy",
			Website: "https://eltropy.com/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                6464349,
				IDs:               nil,
				Alias:             "eltropy",
				Name:              "Eltropy",
				Followers:         "45K",
				Employees:         "51-200",
				AssociatedMembers: "249",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "",
				Verified: false,
			},
			BlindProfile: domain.BlindProfile{
				Alias:       "Eltropy",
				Employees:   "51 to 200",
				Salary:      "",
				Reviews:     "5",
				ReviewsRate: "3.2",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias: "",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Eltropy-EI_IE1320280.11,18.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Eltropy-Reviews-E1320280.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Eltropy-Jobs-E1320280.htm",
				Jobs:        "12",
				Reviews:     "84",
				Salaries:    "114",
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
							Title:                "Senior Software Engineer – Golang/Python",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4260746039/",
							Date:                 mustDate("2025-07-08"),
							WithSalary:           false,
							Remote:               false,
						},
					},
				},
				domain.Rust: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Zig: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Scala: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Elixir: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Clojure: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Haskell: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
			},
			ShortDescription: "",
			Industries: []domain.Industry{
				domain.IndustryFinTech,
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "UNEY",
			Website: "https://uney.com/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                101728390,
				IDs:               nil,
				Alias:             "weareuney",
				Name:              "UNEY",
				Followers:         "4K",
				Employees:         "51-200",
				AssociatedMembers: "49",
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
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-UNEY-EI_IE9833439.11,15.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/UNEY-Reviews-E9833439.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/UNEY-Jobs-E9833439.htm",
				Jobs:        "",
				Reviews:     "11",
				Salaries:    "9",
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
							Title:                "(Senior) Principal Backend Engineer (Go/Golang/Rust)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4263368347/",
							Date:                 mustDate("2025-07-09"),
							WithSalary:           false,
							Remote:               false,
						},
					},
				},
				domain.Rust: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "(Senior) Principal Backend Engineer (Go/Golang/Rust)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4263368347/",
							Date:                 mustDate("2025-07-09"),
							WithSalary:           false,
							Remote:               false,
						},
					},
				},
				domain.Zig: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Scala: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Elixir: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Clojure: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Haskell: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
			},
			ShortDescription: "Software solutions that prioritize privacy and security, ensuring our users' data is always protected",
			Industries:       []domain.Industry{
				// NOP
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "AIQ.com (Alpine IQ)",
			Website: "https://aiq.com/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                40940457,
				IDs:               nil,
				Alias:             "alpine-iq",
				Name:              "AIQ.com (Alpine IQ)",
				Followers:         "8K",
				Employees:         "51-200",
				AssociatedMembers: "94",
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
							Title:                "Software Engineer, Level 2 (Golang)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4264557725/",
							Date:                 mustDate("2025-07-09"),
							WithSalary:           true, // $120k/yr - $145k/yr
							Remote:               true,
						},
					},
				},
				domain.Rust: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Zig: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Scala: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Elixir: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Clojure: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Haskell: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
			},
			ShortDescription: "AIQ powers seamless customer loyalty and engagement",
			Industries:       []domain.Industry{
				// NOP
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Celonis",
			Website: "https://www.celonis.com/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                3118913,
				IDs:               []int{3118913, 3197504},
				Alias:             "celonis",
				Name:              "Celonis",
				Followers:         "284K",
				Employees:         "1K-5K",
				AssociatedMembers: "3,662",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:     "celonis",
				Followers: "248",
				Verified:  true,
			},
			BlindProfile: domain.BlindProfile{
				Alias:       "Celonis",
				Employees:   "1,001 to 5,000 Employees",
				Salary:      "$70K ~ $235K a year",
				Reviews:     "97",
				ReviewsRate: "3.3",
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
					GitHubRepositoriesCount: 1,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Golang – Staff Software Engineer",
							ShortDescription:     "Orchestration Engine",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4261148978/",
							Date:                 mustDate("2025-07-09"),
							WithSalary:           false,
							Remote:               false,
						},
					},
				},
				domain.Rust: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Zig: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Scala: {
					GitHubRepositoriesCount: 1,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Elixir: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Clojure: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Haskell: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
			},
			ShortDescription: "The Celonis Process Intelligence Platform takes the data from the systems you already use, and presents you with a living digital twin of your end-to-end processes",
			Industries:       []domain.Industry{
				// NOP
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Altea Healthcare",
			Website: "https://alteahc.com/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                89863875,
				IDs:               nil,
				Alias:             "altea-healthcare",
				Name:              "Altea Healthcare",
				Followers:         "42K",
				Employees:         "201-500",
				AssociatedMembers: "171",
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
							Title:                "Senior Full-Stack Developer (Golang, React, MongoDB)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4262881374/",
							Date:                 mustDate("2025-07-09"),
							WithSalary:           true, // $110k/yr - $160k/yr
							Remote:               false,
						},
					},
				},
				domain.Rust: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Zig: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Scala: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Elixir: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Clojure: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Haskell: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
			},
			ShortDescription: "ALTEA Healthcare offers a comprehensive range of services, including provider services, medical directorship, and advanced technology solutions tailored specifically for outpatient facilities and clinics",
			Industries: []domain.Industry{
				domain.IndustryMedTech,
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "SmartBear",
			Website: "https://smartbear.com/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                108422,
				IDs:               nil,
				Alias:             "smartbear",
				Name:              "SmartBear",
				Followers:         "50K",
				Employees:         "501-1K",
				AssociatedMembers: "1,025",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:     "smartbear",
				Followers: "257",
				Verified:  true,
			},
			BlindProfile: domain.BlindProfile{
				Alias:       "SmartBear-Software",
				Employees:   "501 to 1,000",
				Salary:      "",
				Reviews:     "8",
				ReviewsRate: "3.6",
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
					GitHubRepositoriesCount: 3,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Go Software Engineer",
							ShortDescription:     "BitBar",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4263312446/",
							Date:                 mustDate("2025-07-09"),
							WithSalary:           false,
							Remote:               false,
						},
					},
				},
				domain.Rust: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Zig: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Scala: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Elixir: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Clojure: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Haskell: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
			},
			ShortDescription: "",
			Industries:       []domain.Industry{
				// NOP
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Hello Health Group",
			Website: "https://hellohealthgroup.com/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                6448869,
				IDs:               []int{6448869, 7604161, 13348951, 14525874, 14575417, 14591689, 81843837, 82475638, 82616972},
				Alias:             "hello-health-group",
				Name:              "Hello Health Group",
				Followers:         "29K",
				Employees:         "201-500",
				AssociatedMembers: "298",
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
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Hello-Health-Group-EI_IE1617422.11,29.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Hello-Health-Group-Reviews-E1617422.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Hello-Health-Group-Jobs-E1617422.htm",
				Jobs:        "",
				Reviews:     "79",
				Salaries:    "58",
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
							Title:                "Senior Golang Developer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4264572720/",
							Date:                 mustDate("2025-07-09"),
							WithSalary:           false,
							Remote:               false,
						},
					},
				},
				domain.Rust: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Zig: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Scala: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Elixir: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Clojure: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Haskell: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
			},
			ShortDescription: "Hello Health is building Emerging Asia’s Digital Health Ecosystem",
			Industries: []domain.Industry{
				domain.IndustryMedTech,
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "RIIICO",
			Website: "https://riiico.com/",
			Careers: "",
			About:   "https://riiico.com/about/",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                69330782,
				IDs:               nil,
				Alias:             "riiico",
				Name:              "RIIICO",
				Followers:         "3K",
				Employees:         "11-50",
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
				domain.Rust: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Junior Systems Engineer",
							ShortDescription:     "Maintain and extend our Rust data export pipeline",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4261919706/",
							Date:                 mustDate("2025-07-10"),
							WithSalary:           false,
							Remote:               false,
						},
					},
				},
				domain.Zig: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Scala: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Elixir: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Clojure: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Haskell: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
			},
			ShortDescription: "Synchronisation of physical factories with their digital counterparts",
			Industries:       []domain.Industry{
				// NOP
			},
		},

		// Template
		//{
		//	ID:      0,  // system
		//	Type:    "", // system
		//	Name:    "",
		//	Website: "",
		//	Careers: "",
		//	About:   "",
		//	Blog:    "",
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
		//	BlindProfile: domain.BlindProfile{
		//		Alias: "",
		//	},
		//	LevelsFyiProfile: domain.LevelsFyiProfile{
		//		Alias: "",
		//	},
		//	GlassdoorProfile: domain.GlassdoorProfile{
		//		OverviewURL: "",
		//	},
		//	IndeedProfile: domain.IndeedProfile{
		//		Alias: "",
		//	},
		//	Languages: domain.Languages{
		//		domain.Go: {
		//			GitHubRepositoriesCount: 0,
		//			Vacancies:               []domain.Vacancy{},
		//		},
		//		domain.Rust: {
		//			GitHubRepositoriesCount: 0,
		//			Vacancies:               []domain.Vacancy{},
		//		},
		//		domain.Zig: {
		//			GitHubRepositoriesCount: 0,
		//			Vacancies:               []domain.Vacancy{},
		//		},
		//		domain.Scala: {
		//			GitHubRepositoriesCount: 0,
		//			Vacancies:               []domain.Vacancy{},
		//		},
		//		domain.Elixir: {
		//			GitHubRepositoriesCount: 0,
		//			Vacancies:               []domain.Vacancy{},
		//		},
		//		domain.Clojure: {
		//			GitHubRepositoriesCount: 0,
		//			Vacancies:               []domain.Vacancy{},
		//		},
		//		domain.Haskell: {
		//			GitHubRepositoriesCount: 0,
		//			Vacancies:               []domain.Vacancy{},
		//		},
		//	},
		//	ShortDescription: "",
		//	Industries:       []domain.Industry{
		//		// NOP
		//	},
		//},
	}
}
