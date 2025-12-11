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
				Jobs:        "27",
				Reviews:     "12",
				Salaries:    "24",
				ReviewsRate: "3.3",
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
							Date:                 mustDate("2025-07-10", "2025-06-23"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Senior Software Engineer — Rust",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4298779685/",
							Location:             "Berlin, Berlin, Germany",
							Date:                 mustDate("2025-10-08", "2025-09-10"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Senior Software Engineer — Rust",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4298772901/",
							Location:             "Heidelberg, Baden-Württemberg, Germany",
							Date:                 mustDate("2025-10-08"),
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
						{
							Title:                "Go Developer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4298390655/",
							Location:             "United Kingdom",
							Date:                 mustDate("2025-09-10"),
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
						{
							Title:                "Senior Backend Engineer – Go",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4342904107/",
							Location:             "United States",
							Date:                 mustDate("2025-12-10"),
							WithSalary:           true, // $160k/yr - $240k/yr
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
							Location:             "San Jose, CA",
							Date:                 mustDate("2025-12-08"), // mustDate("2025-11-17", "2025-10-27", "2025-10-06", "2025-09-14", "2025-08-23", "2025-08-02", "2025-07-10", "2025-06-23"),
							WithSalary:           true,                   // $150k/yr - $175k/yr
							Remote:               false,
						},
						{
							Title:                "Senior Software Engineer – Golang",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4289985968/",
							Location:             "Toronto, ON",
							Date:                 mustDate("2025-08-31", "2025-08-23"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Senior Software Engineer – Golang",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4306332686/",
							Location:             "Toronto, ON",
							Date:                 mustDate("2025-10-21"),
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
			Website: "https://www.experian.com/",
			Careers: "https://www.experian.com/careers/",
			About:   "https://www.experian.com/corporate/about-experian",
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
				Alias: "Experian",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Experian-EI_IE42406.11,19.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Experian-Reviews-E42406.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Experian-Jobs-E42406.htm",
				Jobs:        "754",
				Reviews:     "7.7K",
				Salaries:    "11K",
				ReviewsRate: "4.1",
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
						{
							Title:                "Senior Software Engineer – (Big Data, Scala)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4274816427/",
							Location:             "Hyderabad, Telangana, India",
							Date:                 mustDate("2025-09-06", "2025-07-23"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Software Engineer – Go",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4315346445/",
							Location:             "London, England, United Kingdom",
							Date:                 mustDate("2025-11-29", "2025-11-07"),
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
				Alias:     "thomson-reuters",
				Employees: "24,000",
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
							Location:             "Bengaluru, Karnataka, India",
							Date:                 mustDate("2025-07-28", "2025-07-07"),
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
				Alias:     "aiven",
				Employees: "180",
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
							Location:             "Helsinki, Uusimaa, Finland",
							Date:                 mustDate("2025-09-10", "2025-08-19", "2025-07-28", "2025-07-07"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Senior Software Engineer (Golang)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4265470899/",
							Location:             "Berlin, Berlin, Germany",
							Date:                 mustDate("2025-09-20", "2025-08-06", "2025-07-15"),
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
				Reviews:     "232",
				Salaries:    "270",
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
							Title:                "Lead Engineer (Golang)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4263453034/",
							Date:                 mustDate("2025-07-07"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Software Engineer – Go",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4306865654/",
							Location:             "Mexico City, Mexico",
							Date:                 mustDate("2025-09-28"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Software Engineer – Go",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4312048947/",
							Location:             "Mérida, Yucatán, Mexico",
							Date:                 mustDate("2025-10-08"),
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
						{
							Title:                "Lead Engineer – Golang",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4276796401/",
							Location:             "Bengaluru, Karnataka, India",
							Date:                 mustDate("2025-07-28"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Lead Engineer – Golang",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4303072316/",
							Location:             "Bengaluru, Karnataka, India",
							Date:                 mustDate("2025-09-23"),
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
				Alias:     "wolters-kluwer",
				Employees: "19,000",
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
						{
							Title:                "(Senior) Golang Software Engineer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4289035251/",
							Location:             "Leuven, Flemish Region, Belgium",
							Date:                 mustDate("2025-09-13", "2025-08-22"),
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
							Location:             "Argentina",
							Date:                 mustDate("2025-07-30", "2025-07-07"),
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
				Alias:     "parallel-wireless",
				Employees: "750",
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
							Location:             "Kfar Saba, Center District, Israel",
							Date:                 mustDate("2025-07-30", "2025-07-07"),
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
						{
							Title:                "Senior Elixir Developer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4272135150/",
							Location:             "Warsaw, Mazowieckie, Poland",
							Date:                 mustDate("2025-07-30"),
							WithSalary:           false,
							Remote:               true,
						},
						{
							Title:                "Elixir Developer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4341941098/",
							Location:             "Tel Aviv-Yafo, Tel Aviv District, Israel",
							Date:                 mustDate("2025-11-20"),
							WithSalary:           false,
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
			Careers: "https://www.whitefoxdefense.com/careers",
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
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Nuitee-Travel-EI_IE5318150.11,24.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Nuitee-Travel-Reviews-E5318150.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Nuitee-Travel-Jobs-E5318150.htm",
				Jobs:        "",
				Reviews:     "57",
				Salaries:    "65",
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
							Title:                "Senior Golang Engineer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4262885265/",
							Date:                 mustDate("2025-07-08"),
							WithSalary:           false,
							Remote:               true,
						},
						{
							Title:                "Senior Golang Engineer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4287085768/",
							Location:             "United Kingdom",
							Date:                 mustDate("2025-08-17"),
							WithSalary:           false,
							Remote:               true,
						},
						{
							Title:                "Senior Golang Engineer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4289486465/",
							Location:             "United Kingdom",
							Date:                 mustDate("2025-08-22"),
							WithSalary:           false,
							Remote:               true,
						},
						{
							Title:                "Senior Golang Engineer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4295362915/",
							Location:             "United Kingdom",
							Date:                 mustDate("2025-09-04"),
							WithSalary:           false,
							Remote:               true,
						},
						{
							Title:                "Senior Golang Engineer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4299882918/",
							Location:             "United Kingdom",
							Date:                 mustDate("2025-09-14"),
							WithSalary:           false,
							Remote:               true,
						},
						{
							Title:                "Senior Golang Engineer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4304024501/",
							Location:             "United Kingdom",
							Date:                 mustDate("2025-09-23"),
							WithSalary:           false,
							Remote:               true,
						},
						{
							Title:                "Senior Golang Engineer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4307008431/",
							Location:             "United Kingdom",
							Date:                 mustDate("2025-10-01"),
							WithSalary:           false,
							Remote:               true,
						},
						{
							Title:                "Senior Backend (Golang) Engineer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4326630654/",
							Location:             "Barcelona, Catalonia, Spain",
							Date:                 mustDate("2025-12-10"),
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
				Jobs:        "5",
				Reviews:     "65",
				Salaries:    "118",
				ReviewsRate: "4.1",
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
						{
							Title:                "Senior Full-Stack Engineer (Go / TypeScript)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4309201003/",
							Location:             "Lisbon, Lisbon, Portugal",
							Date:                 mustDate("2025-10-07"),
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
				Reviews:     "14",
				Salaries:    "13",
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
						{
							Title:                "Principal Backend Engineer – Rust",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4309255202/",
							Location:             "Dubai, United Arab Emirates",
							Date:                 mustDate("2025-10-08"),
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
				Alias:     "celonis",
				Employees: "2,000",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Celonis-EI_IE1307503.11,18.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Celonis-Reviews-E1307503.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Celonis-Jobs-E1307503.htm",
				Jobs:        "193",
				Reviews:     "1.1K",
				Salaries:    "2K",
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
							Title:                "Golang – Staff Software Engineer",
							ShortDescription:     "Orchestration Engine",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4261148978/",
							Location:             "Madrid, Community of Madrid, Spain",
							Date:                 mustDate("2025-07-09"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Golang – Software Engineer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4270882682/",
							Location:             "Madrid, Community of Madrid, Spain",
							Date:                 mustDate("2025-07-27"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Senior Golang Backend Engineer",
							ShortDescription:     "Orchestration Engine",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4296750297/",
							Location:             "Madrid, Community of Madrid, Spain",
							Date:                 mustDate("2025-10-01"),
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
			Careers: "https://smartbear.com/company/careers/",
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
				Alias:     "smartbear-software",
				Employees: "780",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-SmartBear-EI_IE410916.11,20.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/SmartBear-Reviews-E410916.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/SmartBear-Jobs-E410916.htm",
				Jobs:        "29",
				Reviews:     "386",
				Salaries:    "664",
				ReviewsRate: "3.6",
				Verified:    false,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "Smartbear",
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
							Location:             "Wrocław, Dolnośląskie, Poland",
							Date:                 mustDate("2025-07-30", "2025-07-09"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Senior Go Software Engineer",
							ShortDescription:     "BitBar",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4260683559/",
							Location:             "Wrocław, Dolnośląskie, Poland",
							Date:                 mustDate("2025-07-25"),
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
			Careers: "https://hellohealthgroup.com/careers/",
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
				Reviews:     "81",
				Salaries:    "60",
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
						{
							Title:                "Senior Backend Golang Developer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4265110120/",
							Location:             "Ho Chi Minh City, Vietnam",
							Date:                 mustDate("2025-07-12"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Backend Developer (Golang)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4310096314/",
							Location:             "Ho Chi Minh City, Vietnam",
							Date:                 mustDate("2025-10-08"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Backend Developer (Golang)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4317175505/",
							Location:             "Ho Chi Minh City, Vietnam",
							Date:                 mustDate("2025-10-22"),
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
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Boozt",
			Website: "https://www.boozt.com/",
			Careers: "https://careers.booztgroup.com/",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                478478,
				IDs:               nil,
				Alias:             "boozt-fashion",
				Name:              "Boozt",
				Followers:         "27K",
				Employees:         "1K-5K",
				AssociatedMembers: "771",
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
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Boozt-EI_IE2391881.11,16.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Boozt-Reviews-E2391881.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Boozt-Jobs-E2391881.htm",
				Jobs:        "11",
				Reviews:     "164",
				Salaries:    "240",
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
							Title:                "Senior Software Engineer (Golang/PHP)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4222848827/",
							Location:             "Malmo, Skåne County, Sweden",
							Date:                 mustDate("2025-08-23", "2025-07-10"),
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
			ShortDescription: "Nordic department store selling curated Fashion, Kids, Sport, Home and Beauty online",
			Industries: []domain.Industry{
				domain.IndustryECommerce,
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "FrankieOne",
			Website: "https://frankieone.com/",
			Careers: "https://careers.frankieone.com/",
			About:   "https://frankieone.com/frankieone-about-us",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                35577295,
				IDs:               nil,
				Alias:             "frankieone",
				Name:              "FrankieOne",
				Followers:         "8K",
				Employees:         "51-200",
				AssociatedMembers: "90",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:     "FrankieFinancial",
				Followers: "5",
				Verified:  true,
			},
			BlindProfile: domain.BlindProfile{
				Alias: "",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias: "",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-FrankieOne-EI_IE6239581.11,21.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/FrankieOne-Reviews-E6239581.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/FrankieOne-Jobs-E6239581.htm",
				Jobs:        "8",
				Reviews:     "53",
				Salaries:    "60",
				ReviewsRate: "3.4",
				Verified:    true,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "",
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 9,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Senior Backend Engineer (Golang)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4264350447/",
							Location:             "Melbourne, Victoria, Australia",
							Date:                 mustDate("2025-07-10"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Senior Backend Engineer (Golang)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4276108751/",
							Location:             "Melbourne, Victoria, Australia",
							Date:                 mustDate("2025-08-26", "2025-07-28"),
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
			ShortDescription: "Identity verification and fraud detection",
			Industries: []domain.Industry{
				domain.IndustryCyberSecurity,
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Caliber Smart",
			Website: "https://caliberreps.com/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                18856651,
				IDs:               nil,
				Alias:             "calibersmart",
				Name:              "Caliber Smart",
				Followers:         "1K",
				Employees:         "501-1K",
				AssociatedMembers: "129",
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
							Title:                "Java developer with Golang",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4262553994/",
							Location:             "Los Angeles, CA",
							Date:                 mustDate("2025-07-10"),
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
			ShortDescription: "Caliber specializes in direct to consumer sales",
			Industries:       []domain.Industry{
				// NOP
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Archie",
			Website: "https://archieapp.co/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                11107334,
				IDs:               nil,
				Alias:             "archieapp",
				Name:              "Archie",
				Followers:         "4K",
				Employees:         "51-200",
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
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Senior Go Developer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4261955207/",
							Location:             "Montreal, QC",
							Date:                 mustDate("2025-07-10"),
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
			ShortDescription: "Archie provides 1000+ flexible & hybrid workspaces around the world with an all-in-one platform to manage flexible offices from an intuitive and elegant platform",
			Industries:       []domain.Industry{
				// NOP
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "OXIO",
			Website: "https://oxio.com/",
			Careers: "https://oxio.com/join-the-team/",
			About:   "https://oxio.com/company/",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                28141767,
				IDs:               nil,
				Alias:             "oxiocorp",
				Name:              "OXIO",
				Followers:         "5K",
				Employees:         "51-200",
				AssociatedMembers: "108",
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
				Alias:     "oxio",
				Employees: "126",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-OXIO-EI_IE2395820.11,15.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/OXIO-Reviews-E2395820.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/OXIO-Jobs-E2395820.htm",
				Jobs:        "5",
				Reviews:     "19",
				Salaries:    "29",
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
							Title:                "Go Developer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4262532716/",
							Location:             "New York City Metropolitan Area",
							Date:                 mustDate("2025-08-02", "2025-07-10"),
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
			ShortDescription: "Telecom-as-a-Service (TaaS) platform",
			Industries:       []domain.Industry{
				// NOP
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Bellroy",
			Website: "https://bellroy.com/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                926134,
				IDs:               nil,
				Alias:             "bellroy",
				Name:              "Bellroy",
				Followers:         "15K",
				Employees:         "51-200",
				AssociatedMembers: "141",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:     "bellroy",
				Followers: "32",
				Verified:  false,
			},
			BlindProfile: domain.BlindProfile{
				Alias: "",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias: "",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Bellroy-EI_IE811984.11,18.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Bellroy-Reviews-E811984.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Bellroy-Jobs-E811984.htm",
				Jobs:        "15",
				Reviews:     "26",
				Salaries:    "44",
				ReviewsRate: "4.6",
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
					GitHubRepositoriesCount: 24,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Haskell Developer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4263012229/",
							Location:             "Collingwood, Victoria, Australia",
							Date:                 mustDate("2025-11-06", "2025-08-05", "2025-07-12"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Haskell Developer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4278484490/",
							Location:             "Australia and New Zealand",
							Date:                 mustDate("2025-08-04"),
							WithSalary:           false,
							Remote:               true,
						},
					},
				},
			},
			ShortDescription: "Bellroy is an Australian accessories brand making carry goods",
			Industries: []domain.Industry{
				domain.IndustryECommerce,
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Azion",
			Website: "https://www.azion.com/",
			Careers: "https://www.azion.com/en/careers/",
			About:   "https://www.azion.com/en/about-us/",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                2137446,
				IDs:               nil,
				Alias:             "aziontech",
				Name:              "Azion",
				Followers:         "40K",
				Employees:         "201-500",
				AssociatedMembers: "201",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:     "aziontech",
				Followers: "120",
				Verified:  false,
			},
			BlindProfile: domain.BlindProfile{
				Alias: "",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "azion",
				Employees: "351",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Azion-Technologies-EI_IE1525794.11,29.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Azion-Technologies-Reviews-E1525794.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Azion-Technologies-Jobs-E1525794.htm",
				Jobs:        "19",
				Reviews:     "157",
				Salaries:    "283",
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
							Title:                "Senior Software Engineer (Go)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4266554424/",
							Location:             "Greater Porto Alegre",
							Date:                 mustDate("2025-07-15"),
							WithSalary:           false,
							Remote:               false,
						},
					},
				},
				domain.Rust: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Software Engineer (Rust)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4266549864/",
							Location:             "Greater Porto Alegre",
							Date:                 mustDate("2025-07-14"),
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
			ShortDescription: "Azion is computing platform that simplifies how you build and run applications anywhere",
			Industries: []domain.Industry{
				domain.IndustryCloudComputing,
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Riverside Research",
			Website: "https://www.riversideresearch.org/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                1945102,
				IDs:               nil,
				Alias:             "riverside-research",
				Name:              "Riverside Research",
				Followers:         "11K",
				Employees:         "501-1K",
				AssociatedMembers: "753",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:     "riversideresearch",
				Followers: "7",
				Verified:  false,
			},
			BlindProfile: domain.BlindProfile{
				Alias: "",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "riverside-research",
				Employees: "630",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Riverside-Research-EI_IE193384.11,29.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Riverside-Research-Reviews-E193384.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Riverside-Research-Jobs-E193384.htm",
				Jobs:        "31",
				Reviews:     "135",
				Salaries:    "257",
				ReviewsRate: "4.0",
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
					Vacancies: []domain.Vacancy{
						{
							Title:                "Software Engineer – Low-level Systems (C, C++, Rust)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4249280373/",
							Location:             "Lexington, MA",
							Date:                 mustDate("2025-07-15"),
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
			ShortDescription: "Founded in 1967, Riverside Research is a nonprofit organization chartered to advance scientific research for the benefit of the US government and in the public interest",
			Industries: []domain.Industry{
				domain.IndustryCyberSecurity,
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Ecosia",
			Website: "https://www.ecosia.org/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                1114590,
				IDs:               nil,
				Alias:             "ecosia",
				Name:              "Ecosia",
				Followers:         "154K",
				Employees:         "51-200",
				AssociatedMembers: "139",
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
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Ecosia-EI_IE1627398.11,17.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Ecosia-Reviews-E1627398.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Ecosia-Jobs-E1627398.htm",
				Jobs:        "",
				Reviews:     "33",
				Salaries:    "55",
				ReviewsRate: "4.4",
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
							Title:                "Software Engineer – Backend",
							ShortDescription:     "Node.js, Golang, and Python services using REST and gRPC",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4253805508/",
							Location:             "Berlin, Berlin, Germany",
							Date:                 mustDate("2025-06-24"),
							WithSalary:           true, // €64k/yr - €85k/yr
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
			ShortDescription: "Ecosia is the search engine that plants trees. 100% of our profits go to climate action.",
			Industries:       []domain.Industry{
				// NOP
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Breeze",
			Website: "https://breeze.social/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                26560342,
				IDs:               nil,
				Alias:             "breezesocial",
				Name:              "Breeze",
				Followers:         "5K",
				Employees:         "11-50",
				AssociatedMembers: "59",
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
							URL:                  "https://www.linkedin.com/jobs/view/4264494875/",
							Location:             "Rotterdam, South Holland, Netherlands",
							Date:                 mustDate("2025-07-15"),
							WithSalary:           true, // €4,500/month - €6,000/month
							Remote:               false,
						},
						{
							Title:                "Backend Engineer (Golang)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4281160160/",
							Location:             "Rotterdam, South Holland, Netherlands",
							Date:                 mustDate("2025-08-08"),
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
			ShortDescription: "The app for offline dates",
			Industries:       []domain.Industry{
				// NOP
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Winnow",
			Website: "https://www.winnowsolutions.com/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                5366082,
				IDs:               nil,
				Alias:             "winnow-ltd-",
				Name:              "Winnow",
				Followers:         "44K",
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
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Winnow-Solutions-EI_IE1857734.11,27.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Winnow-Solutions-Reviews-E1857734.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Winnow-Solutions-Jobs-E1857734.htm",
				Jobs:        "",
				Reviews:     "53",
				Salaries:    "85",
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
							Title:                "Golang (Polyglot) Software Developer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4266723296/",
							Location:             "London, England, United Kingdom",
							Date:                 mustDate("2025-07-15"),
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
			ShortDescription: "Winnow develops artificial intelligence tools to help chefs run more profitable and sustainable kitchens",
			Industries:       []domain.Industry{
				// NOP
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "gwidotcom",
			Website: "https://www.gwi.com/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                2515651,
				IDs:               nil,
				Alias:             "gwidotcom",
				Name:              "GWI",
				Followers:         "64K",
				Employees:         "501-1K",
				AssociatedMembers: "740",
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
				Alias:     "gwi",
				Employees: "760",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-GWI-EI_IE1333381.11,14.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/GWI-Reviews-E1333381.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/GWI-Jobs-E1333381.htm",
				Jobs:        "21",
				Reviews:     "236",
				Salaries:    "589",
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
							Title:                "Senior Backend Engineer (Go)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4266237220/",
							Location:             "Athens, Attiki, Greece",
							Date:                 mustDate("2025-08-05", "2025-07-15"),
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
			ShortDescription: "Target audience company",
			Industries:       []domain.Industry{
				// NOP
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Pepperstone",
			Website: "https://pepperstone.com/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                1537446,
				IDs:               nil,
				Alias:             "pepperstone",
				Name:              "Pepperstone",
				Followers:         "53K",
				Employees:         "201-500",
				AssociatedMembers: "784",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:     "pepperstone",
				Followers: "53",
				Verified:  true,
			},
			BlindProfile: domain.BlindProfile{
				Alias: "",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias: "",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Pepperstone-Group-EI_IE1971578.11,28.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Pepperstone-Group-Reviews-E1971578.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Pepperstone-Group-Jobs-E1971578.htm",
				Jobs:        "15",
				Reviews:     "114",
				Salaries:    "155",
				ReviewsRate: "3.6",
				Verified:    true,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "",
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 2,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Senior Engineer — Go",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4255303317/",
							Location:             "Limassol, Limassol, Cyprus",
							Date:                 mustDate("2025-07-15"),
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
			ShortDescription: "Pepperstone is an online Forex and CFD Broker",
			Industries: []domain.Industry{
				domain.IndustryFinTech,
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "FICO",
			Website: "https://www.fico.com/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                3398,
				IDs:               nil,
				Alias:             "fico",
				Name:              "FICO",
				Followers:         "487K",
				Employees:         "1K-5K",
				AssociatedMembers: "3,770",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "",
				Verified: false,
			},
			BlindProfile: domain.BlindProfile{
				Alias:       "FICO",
				Employees:   "1,001 to 5,000",
				Salary:      "$57K ~ $215K a year",
				Reviews:     "46",
				ReviewsRate: "3.5",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "fico",
				Employees: "4,000",
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
							Title:                "Full-Stack / Front-End Engineer — Angular JS | TypeScript | Go",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4255383137/",
							Location:             "United States",
							Date:                 mustDate("2025-08-06", "2025-07-15"),
							WithSalary:           true, // $105k/yr - $165k/yr
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
			ShortDescription: "Analytics company",
			Industries:       []domain.Industry{
				// NOP
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "allpay Limited",
			Website: "https://www.allpay.net/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                72115,
				IDs:               nil,
				Alias:             "allpay-limited",
				Name:              "allpay Limited",
				Followers:         "8K",
				Employees:         "201-500",
				AssociatedMembers: "279",
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
				Alias:     "allpay-limited",
				Employees: "210",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-allpay-EI_IE213649.11,17.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/allpay-Reviews-E213649.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/allpay-Jobs-E213649.htm",
				Jobs:        "11",
				Reviews:     "95",
				Salaries:    "212",
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
							Title:                "Go Backend developer",
							ShortDescription:     "Flutter experience as a plus",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4266332868/",
							Location:             "Hereford, England, United Kingdom",
							Date:                 mustDate("2025-07-15"),
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
			Name:    "Adyen",
			Website: "https://www.adyen.com/",
			Careers: "https://careers.adyen.com/",
			About:   "https://www.adyen.com/about",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                279565,
				IDs:               nil,
				Alias:             "adyen",
				Name:              "Adyen",
				Followers:         "320K",
				Employees:         "1K-5K",
				AssociatedMembers: "5,048",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:     "adyen",
				Followers: "390",
				Verified:  true,
			},
			BlindProfile: domain.BlindProfile{
				Alias: "",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "adyen",
				Employees: "2,500",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "",
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "",
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 2,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Infrastructure Developer (Go)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4266555529/",
							Location:             "Amsterdam, North Holland, Netherlands",
							Date:                 mustDate("2025-07-15"),
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
			ShortDescription: "Financial technology platform",
			Industries: []domain.Industry{
				domain.IndustryFinTech,
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Immutable",
			Website: "https://www.immutable.com/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                18693504,
				IDs:               nil,
				Alias:             "immutable-1",
				Name:              "Immutable",
				Followers:         "54K",
				Employees:         "201-500",
				AssociatedMembers: "231",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:     "immutable",
				Followers: "205",
				Verified:  true,
			},
			BlindProfile: domain.BlindProfile{
				Alias: "",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "immutable",
				Employees: "150",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Immutable-EI_IE3253189.11,20.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Immutable-Reviews-E3253189.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Immutable-Jobs-E3253189.htm",
				Jobs:        "10",
				Reviews:     "169",
				Salaries:    "252",
				ReviewsRate: "4.2",
				Verified:    true,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "",
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 5,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Senior Software Engineer (Golang, TypeScript)",
							ShortDescription:     "Wallets",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4267047562/",
							Location:             "Sydney, New South Wales, Australia",
							Date:                 mustDate("2025-10-04", "2025-08-21", "2025-07-31", "2025-07-16"),
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
			ShortDescription: "Platform for building games on Ethereum",
			Industries: []domain.Industry{
				domain.IndustryGameDev,
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Rincon Research Corporation",
			Website: "https://www.rincon.com/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                70870,
				IDs:               nil,
				Alias:             "rincon-research-corporation",
				Name:              "Rincon Research Corporation",
				Followers:         "4K",
				Employees:         "201-500",
				AssociatedMembers: "273",
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
							Title:                "Advanced Software Engineer – Golang",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4244085841/",
							Location:             "Tucson, AZ",
							Date:                 mustDate("2025-07-16"),
							WithSalary:           true, // $135k/yr - $165k/yr
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
			ShortDescription: "At Rincon Research Corporation, our core business is to design, build, test, and field digital signal processing (DSP) products and services for the United States Defense",
			Industries:       []domain.Industry{
				// NOP
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Azberry BV",
			Website: "https://azberry.com/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                100755618,
				IDs:               nil,
				Alias:             "azberry-bv",
				Name:              "Azberry BV",
				Followers:         "2K",
				Employees:         "51-200",
				AssociatedMembers: "11",
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
							Title:                "Middle Backend Engineer (Golang)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4266412740/",
							Location:             "Türkiye",
							Date:                 mustDate("2025-07-16"),
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
			ShortDescription: "",
			Industries: []domain.Industry{
				domain.IndustryAdTech,
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Octopus Money",
			Website: "https://octopusmoney.com/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                11186488,
				IDs:               nil,
				Alias:             "octopusmoney",
				Name:              "Octopus Money",
				Followers:         "10K",
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
				Alias:     "octopus-money",
				Employees: "151",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Octopus-Money-EI_IE1948876.11,24.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Octopus-Money-Reviews-E1948876.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Octopus-Money-Jobs-E1948876.htm",
				Jobs:        "6",
				Reviews:     "13",
				Salaries:    "53",
				ReviewsRate: "4.7",
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
							Title:                "Senior Backend Engineer (Golang)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4267971361/",
							Location:             "London, England, United Kingdom",
							Date:                 mustDate("2025-07-16"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Senior Backend Engineer",
							ShortDescription:     "Golang",
							SwitchingOpportunity: "",
							URL:                  "https://app.welcometothejungle.com/jobs/FBaRkiuG",
							Date:                 mustDate("2025-11-14"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Senior Backend Engineer (Golang)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4316935978/",
							Location:             "London, England, United Kingdom",
							Date:                 mustDate("2025-11-20"),
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
			ShortDescription: "At Octopus Money, we’re on a mission to make money advice accessible to all",
			Industries: []domain.Industry{
				domain.IndustryFinTech,
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Swan",
			Website: "https://www.swan.io/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                19116430,
				IDs:               nil,
				Alias:             "swan-embedded-banking",
				Name:              "Swan",
				Followers:         "19K",
				Employees:         "201-500",
				AssociatedMembers: "290",
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
							Title:                "Senior Scala Software Engineer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4268438536/",
							Location:             "Paris, Île-de-France, France",
							Date:                 mustDate("2025-07-18"),
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
			ShortDescription: "Embedded banking",
			Industries: []domain.Industry{
				domain.IndustryFinTech,
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Julius Baer",
			Website: "https://www.juliusbaer.com/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                7660,
				IDs:               nil,
				Alias:             "julius-baer",
				Name:              "Julius Baer",
				Followers:         "200K",
				Employees:         "5K-10K",
				AssociatedMembers: "7,707",
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
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Julius-Baer-EI_IE12799.11,22.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Julius-Baer-Reviews-E12799.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Julius-Baer-Jobs-E12799.htm",
				Jobs:        "106",
				Reviews:     "1K",
				Salaries:    "1.4K",
				ReviewsRate: "4.0",
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
							Title:                "Senior Scala Developer",
							ShortDescription:     "Markets Quant Team",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4214640651/",
							Location:             "Zurich, Zurich, Switzerland",
							Date:                 mustDate("2025-09-02", "2025-07-19"),
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
			ShortDescription: "",
			Industries: []domain.Industry{
				domain.IndustryFinTech,
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Capital Rx",
			Website: "https://www.cap-rx.com/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                10008489,
				IDs:               nil,
				Alias:             "cap-rx",
				Name:              "Capital Rx",
				Followers:         "111K",
				Employees:         "501-1K",
				AssociatedMembers: "800",
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
							Title:                "Senior Software Developer (Rust)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4268126029/",
							Location:             "Denver, CO",
							Date:                 mustDate("2025-07-18"),
							WithSalary:           true, // $150k/yr - $180k/yr
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
			ShortDescription: "Capital Rx is an enterprise healthcare technology company advancing our nation’s electronic healthcare infrastructure to improve drug price visibility and patient outcomes",
			Industries: []domain.Industry{
				domain.IndustryMedTech,
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "P2P.org",
			Website: "https://www.p2p.org/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                37427350,
				IDs:               nil,
				Alias:             "p2p-org",
				Name:              "P2P.org",
				Followers:         "7K",
				Employees:         "201-500",
				AssociatedMembers: "181",
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
							Title:                "Senior Blockchain Developer – Rust",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4270628559/",
							Location:             "Warsaw, Mazowieckie, Poland",
							Date:                 mustDate("2025-07-21"),
							WithSalary:           false,
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
			ShortDescription: "Staking-as-a-Business",
			Industries:       []domain.Industry{
				// NOP
			},
			Ignore: true, // Blockchain
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Papa",
			Website: "https://www.papa.com/",
			Careers: "https://www.papa.com/careers",
			About:   "https://www.papa.com/about",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                10879786,
				IDs:               nil,
				Alias:             "papainc",
				Name:              "Papa",
				Followers:         "39K",
				Employees:         "501-1K",
				AssociatedMembers: "976",
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
				Alias: "Papa-8",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Papa-EI_IE2525288.11,15.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Papa-Reviews-E2525288.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Papa-Jobs-E2525288.htm",
				Jobs:        "17",
				Reviews:     "196",
				Salaries:    "231",
				ReviewsRate: "3.5",
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
					Vacancies:               []domain.Vacancy{},
				},
				domain.Elixir: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Senior Back-End Software Engineer (Elixir)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4271482521/",
							Location:             "Miami, FL",
							Date:                 mustDate("2025-07-23"),
							WithSalary:           false,
							Remote:               true,
						},
						{
							Title:                "Senior Back-End Software Engineer (Elixir)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4270777580/",
							Location:             "United States",
							Date:                 mustDate("2025-08-17"),
							WithSalary:           false,
							Remote:               true,
						},
						{
							Title:                "Senior Back-End Software Engineer (Elixir)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4315504772/",
							Location:             "United States",
							Date:                 mustDate("2025-11-29", "2025-11-08", "2025-10-17"),
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
			ShortDescription: "Papa is an end-to-end human care network supporting the social needs of older adults, other underserved populations, and families",
			Industries: []domain.Industry{
				domain.IndustryMedTech,
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Bitwarden",
			Website: "https://bitwarden.com/",
			Careers: "https://bitwarden.com/careers/",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                24776937,
				IDs:               nil,
				Alias:             "bitwarden1",
				Name:              "Bitwarden",
				Followers:         "28K",
				Employees:         "201-500",
				AssociatedMembers: "208",
				Verified:          false,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:     "bitwarden",
				Followers: "7.9k",
				Verified:  true,
			},
			BlindProfile: domain.BlindProfile{
				Alias: "",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "bitwarden",
				Employees: "126",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Bitwarden-EI_IE4337610.11,20.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Bitwarden-Reviews-E4337610.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Bitwarden-Jobs-E4337610.htm",
				Jobs:        "9",
				Reviews:     "41",
				Salaries:    "47",
				ReviewsRate: "3.7",
				Verified:    true,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "",
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 4,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Rust: {
					GitHubRepositoriesCount: 6,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Senior Rust Software Engineer",
							ShortDescription:     "Autofill",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4271123395/",
							Location:             "Santa Barbara, CA",
							Date:                 mustDate("2025-07-23"),
							WithSalary:           true, // $140k/yr - $200k/yr
							Remote:               true,
						},
						{
							Title:                "Senior Rust Software Engineer",
							ShortDescription:     "Autofill",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4279522160/",
							Location:             "Santa Barbara, CA",
							Date:                 mustDate("2025-08-04"),
							WithSalary:           true, // $140k/yr - $200k/yr
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
			ShortDescription: "Open source password management solutions",
			Industries: []domain.Industry{
				domain.IndustryCyberSecurity,
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "KYP",
			Website: "https://kyp.io/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                76490514,
				IDs:               nil,
				Alias:             "kyp",
				Name:              "KYP",
				Followers:         "1K",
				Employees:         "2-10",
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
							Title:                "Back-End Engineer Golang",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4267262281/",
							Location:             "United Kingdom",
							Date:                 mustDate("2025-07-24"),
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
			ShortDescription: "UK based Real-Time Third Party Enterprise Risk Intelligence Platform",
			Industries: []domain.Industry{
				domain.IndustryFinTech,
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Virtuozzo",
			Website: "https://www.virtuozzo.com/",
			Careers: "",
			About:   "https://www.virtuozzo.com/company/about/",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                10358578,
				IDs:               []int{1068325, 2426564, 2868387, 10358578},
				Alias:             "virtuozzo",
				Name:              "Virtuozzo",
				Followers:         "17K",
				Employees:         "201-500",
				AssociatedMembers: "252",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:     "virtuozzo",
				Followers: "28",
				Verified:  false,
			},
			BlindProfile: domain.BlindProfile{
				Alias: "",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias: "",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Virtuozzo-EI_IE1099190.11,20.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Virtuozzo-Reviews-E1099190.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Virtuozzo-Jobs-E1099190.htm",
				Jobs:        "22",
				Reviews:     "45",
				Salaries:    "49",
				ReviewsRate: "4.1",
				Verified:    true,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "",
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 8,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Rust: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Lead Rust Developer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4271713155/",
							Location:             "Barcelona, Catalonia, Spain",
							Date:                 mustDate("2025-07-28"),
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
			ShortDescription: "Cloud Platform",
			Industries: []domain.Industry{
				domain.IndustryCloudComputing,
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "NVIDIA",
			Website: "https://www.nvidia.com/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                3608,
				IDs:               nil,
				Alias:             "nvidia",
				Name:              "NVIDIA",
				Followers:         "4M",
				Employees:         "10K+",
				AssociatedMembers: "40,840",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:     "nvidia",
				Followers: "17.9k",
				Verified:  true,
			},
			BlindProfile: domain.BlindProfile{
				Alias:       "NVIDIA",
				Employees:   "10,000+",
				Salary:      "$29K ~ $317K a year",
				Reviews:     "1,054",
				ReviewsRate: "4.3",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "nvidia",
				Employees: "22,810",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-NVIDIA-EI_IE7633.11,17.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/NVIDIA-Reviews-E7633.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/NVIDIA-Jobs-E7633.htm",
				Jobs:        "1.5K",
				Reviews:     "6.5K",
				Salaries:    "65K",
				ReviewsRate: "4.5",
				Verified:    true,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "Nvidia",
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 44,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Senior Systems Software Engineer, Golang",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4275852563/",
							Location:             "United States",
							Date:                 mustDate("2025-07-26"),
							WithSalary:           true, // $148k/yr - $287.5k/yr
							Remote:               true,
						},
						{
							Title:                "Senior Systems Software Engineer, Golang",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4275848868/",
							Location:             "Redmond, WA",
							Date:                 mustDate("2025-08-17"),
							WithSalary:           true, // $148k/yr - $287.5k/yr
							Remote:               false,
						},
						{
							Title:                "Senior Distributed Software Engineer, Golang",
							ShortDescription:     "DGX Cloud",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4282100495/",
							Location:             "United States",
							Date:                 mustDate("2025-09-20"),
							WithSalary:           true, // $148k/yr - $287.5k/yr
							Remote:               true,
						},
						{
							Title:                "Senior Software Cloud Engineer – Golang",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4316039489/",
							Location:             "Santa Clara, CA",
							Date:                 mustDate("2025-11-29", "2025-11-08", "2025-10-18"),
							WithSalary:           true, // $168k/yr - $322k/yr
							Remote:               false,
						},
						{
							Title:                "Senior Software Engineer – Kubernetes and Golang",
							ShortDescription:     "DGX Cloud",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4342275619/",
							Location:             "Santa Clara, CA",
							Date:                 mustDate("2025-12-06"),
							WithSalary:           true, // $184k/yr - $356.5k/yr
							Remote:               false,
						},
					},
				},
				domain.Rust: {
					GitHubRepositoriesCount: 4,
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
			Name:    "Klarna",
			Website: "https://www.klarna.com/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                748731,
				IDs:               nil,
				Alias:             "klarna",
				Name:              "Klarna",
				Followers:         "370K",
				Employees:         "1K-5K",
				AssociatedMembers: "4,718",
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
				Alias:     "klarna",
				Employees: "4,000",
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
				domain.Erlang: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Senior Erlang Engineer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4257029469/",
							Location:             "Warsaw, Mazowieckie, Poland",
							Date:                 mustDate("2025-07-24"),
							WithSalary:           false,
							Remote:               false,
						},
					},
				},
			},
			ShortDescription: "Company that provides online financial services",
			Industries: []domain.Industry{
				domain.IndustryFinTech,
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "xAI",
			Website: "https://x.ai/",
			Careers: "https://x.ai/careers",
			About:   "https://x.ai/company",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                96151950,
				IDs:               nil,
				Alias:             "xai",
				Name:              "xAI",
				Followers:         "84K",
				Employees:         "11-50",
				AssociatedMembers: "3,179",
				Verified:          false,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:     "xai-org",
				Followers: "6.3k",
				Verified:  false,
			},
			BlindProfile: domain.BlindProfile{
				Alias: "",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias: "",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-xAI-EI_IE10404667.11,14.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/xAI-Reviews-E10404667.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/xAI-Jobs-E10404667.htm",
				Jobs:        "",
				Reviews:     "31",
				Salaries:    "62",
				ReviewsRate: "3.9",
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
							Title:                "Rust/C++ Backend Engineer",
							ShortDescription:     "Enterprise Agent",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4278324635/",
							Location:             "London, England, United Kingdom",
							Date:                 mustDate("2025-10-04", "2025-09-13", "2025-07-30"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Rust/C++ Backend Engineer",
							ShortDescription:     "Enterprise Agent",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4262573291/",
							Location:             "Palo Alto, CA",
							Date:                 mustDate("2025-12-09", "2025-11-17", "2025-10-27", "2025-10-05", "2025-09-13", "2025-08-22"),
							WithSalary:           true, // $180k/yr - $440k/yr
							Remote:               false,
						},
						{
							Title:                "Systems Engineer – Rust/C++",
							ShortDescription:     "Remote Sandbox Service",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4306851124/",
							Location:             "Palo Alto, CA",
							Date:                 mustDate("2025-11-29", "2025-11-09", "2025-10-19", "2025-09-28"),
							WithSalary:           true, // $180k/yr - $440k/yr
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
				domain.Erlang: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
			},
			ShortDescription: "Company working on building artificial intelligence",
			Industries:       []domain.Industry{
				// NOP
			},
		},
	}
}
