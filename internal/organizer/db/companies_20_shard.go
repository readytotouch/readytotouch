package db

import "github.com/readytotouch/readytotouch/internal/domain"

func companies20Shard() []domain.CompanyProfile {
	return []domain.CompanyProfile{
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Buenro",
			Website: "https://buenro.com/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                77603395,
				IDs:               nil,
				Alias:             "buenro",
				Name:              "Buenro",
				Followers:         "979",
				Employees:         "11-50",
				AssociatedMembers: "13",
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
							Title:                "Senior Data Engineer (Golang)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4278310297/",
							Location:             "Estonia",
							Date:                 mustDate("2025-07-30"),
							WithSalary:           false,
							Remote:               true,
						},
						{
							Title:                "Senior Golang Engineer",
							ShortDescription:     "Data",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4278079889/",
							Location:             "European Union",
							Date:                 mustDate("2025-09-16", "2025-09-02"),
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
				domain.Erlang: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
			},
			ShortDescription: "Social platform",
			Industries:       []domain.Industry{
				// NOP
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Komerce",
			Website: "https://komerce.id/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                43188359,
				IDs:               nil,
				Alias:             "komerceid",
				Name:              "Komerce",
				Followers:         "5K",
				Employees:         "51-200",
				AssociatedMembers: "198",
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
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Komerce-EI_IE6591742.11,18.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Komerce-Reviews-E6591742.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Komerce-Jobs-E6591742.htm",
				Jobs:        "",
				Reviews:     "7",
				Salaries:    "20",
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
							Title:                "Backend Developer (Golang)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4273562235/",
							Location:             "Purbalingga, Central Java, Indonesia",
							Date:                 mustDate("2025-07-30"),
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
				domain.Erlang: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
			},
			ShortDescription: "",
			Industries: []domain.Industry{
				domain.IndustryECommerce,
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Mayflower",
			Website: "https://mayflower.work/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                88008730,
				IDs:               nil,
				Alias:             "mayflower-tech",
				Name:              "Mayflower",
				Followers:         "15K",
				Employees:         "201-500",
				AssociatedMembers: "358",
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
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Mayflower-EI_IE8288295.11,20.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Mayflower-Reviews-E8288295.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Mayflower-Jobs-E8288295.htm",
				Jobs:        "",
				Reviews:     "3",
				Salaries:    "9",
				ReviewsRate: "3.8",
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
							Title:                "Go/PHP Backend Developer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4274301081/",
							Location:             "Limassol, Limassol, Cyprus",
							Date:                 mustDate("2025-07-30"),
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
				domain.Erlang: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
			},
			ShortDescription: "FunTech company that alters the entertainment industry",
			Industries: []domain.Industry{
				domain.IndustryEntertainment,
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Tontine Trust",
			Website: "https://tontine.com/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                11177312,
				IDs:               nil,
				Alias:             "tontinetrust",
				Name:              "Tontine Trust",
				Followers:         "2K",
				Employees:         "11-50",
				AssociatedMembers: "27",
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
					Vacancies:               []domain.Vacancy{},
				},
				domain.Clojure: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Haskell: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Haskell Developer Intern",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4279708121/",
							Location:             "France",
							Date:                 mustDate("2025-08-04"),
							WithSalary:           false,
							Remote:               true,
						},
						{
							Title:                "Haskell / Scala Engineer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4300913262/",
							Location:             "Serbia",
							Date:                 mustDate("2025-09-15"),
							WithSalary:           false,
							Remote:               true,
						},
					},
				},
				domain.Erlang: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
			},
			ShortDescription: "A Fintech Trust Company offering Trust Funds, Pensions & a TontineIRA™ that all pay you a monthly income for life",
			Industries: []domain.Industry{
				domain.IndustryFinTech,
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Wooga",
			Website: "https://www.wooga.com/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                454295,
				IDs:               nil,
				Alias:             "wooga",
				Name:              "Wooga",
				Followers:         "31K",
				Employees:         "201-500",
				AssociatedMembers: "380",
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
				Alias:     "wooga",
				Employees: "420",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Wooga-EI_IE421202.11,16.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Wooga-Reviews-E421202.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Wooga-Jobs-E421202.htm",
				Jobs:        "7",
				Reviews:     "229",
				Salaries:    "366",
				ReviewsRate: "3.9",
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
							Title:                "Backend Engineer – Elixir",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4278634861/",
							Location:             "Berlin, Germany",
							Date:                 mustDate("2025-08-01"),
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
				domain.Erlang: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
			},
			ShortDescription: "Story-driven casual games",
			Industries: []domain.Industry{
				domain.IndustryGameDev,
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "BoodleBox",
			Website: "https://boodlebox.ai/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                11033108,
				IDs:               nil,
				Alias:             "boodleboxai",
				Name:              "BoodleBox",
				Followers:         "4K",
				Employees:         "11-50",
				AssociatedMembers: "58",
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
					Vacancies: []domain.Vacancy{
						{
							Title:                "Backend Scala Engineer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4277593613/",
							Location:             "United States",
							Date:                 mustDate("2025-08-01"),
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
				domain.Erlang: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
			},
			ShortDescription: "BoodleBox is a collaborative GenAI platform for higher education that enables professors and students to work together responsibly with top AI tools like GPT-4o, Claude, Gemini, and Perplexity and offers features like efficient class preparation, custom AI bots, and academic assessment",
			Industries: []domain.Industry{
				domain.IndustryEdTech,
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "NN",
			Website: "https://www.nn-group.com/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                2230096,
				IDs:               nil,
				Alias:             "nn",
				Name:              "NN",
				Followers:         "121K",
				Employees:         "1K-5K",
				AssociatedMembers: "10,086",
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
							Title:                "Rust Developer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4275926654/",
							Location:             "Budapest, Budapest, Hungary",
							Date:                 mustDate("2025-07-31"),
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
				domain.Erlang: {
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
			Name:    "JustPark",
			Website: "https://www.justpark.com/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                563069,
				IDs:               []int{563069, 5312876},
				Alias:             "justpark",
				Name:              "JustPark",
				Followers:         "30K",
				Employees:         "51-200",
				AssociatedMembers: "229",
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
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-JustPark-EI_IE954923.11,19.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/JustPark-Reviews-E954923.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/JustPark-Jobs-E954923.htm",
				Jobs:        "3",
				Reviews:     "89",
				Salaries:    "123",
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
							Title:                "Senior Golang Engineer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4278940094/",
							Location:             "Dallas, TX",
							Date:                 mustDate("2025-08-01"),
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
				domain.Erlang: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
			},
			ShortDescription: "Parking Management Software",
			Industries:       []domain.Industry{
				// NOP
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Deloitte",
			Website: "https://www.deloitte.com/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                1038,
				IDs:               []int{1038, 116005, 634399, 945079, 1399138, 1729712, 2449847},
				Alias:             "deloitte",
				Name:              "Deloitte",
				Followers:         "19M",
				Employees:         "10K+",
				AssociatedMembers: "490,278",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "",
				Verified: false,
			},
			BlindProfile: domain.BlindProfile{
				Alias:       "Deloitte",
				Employees:   "10,000+",
				Salary:      "$45K ~ $261K a year",
				Reviews:     "1,674",
				ReviewsRate: "3.6",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "deloitte",
				Employees: "335,000",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Deloitte-EI_IE2763.11,19.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Deloitte-Reviews-E2763.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Deloitte-Jobs-E2763.htm",
				Jobs:        "24K",
				Reviews:     "142K",
				Salaries:    "395K",
				ReviewsRate: "3.8",
				Verified:    true,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "Deloitte",
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Golang developer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4277397320/",
							Location:             "Sofia, Sofia City, Bulgaria",
							Date:                 mustDate("2025-10-08", "2025-09-18", "2025-08-26", "2025-08-04"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Senior Golang Developer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4331739119/",
							Location:             "Budapest, Budapest, Hungary",
							Date:                 mustDate("2025-11-01"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Senior Java & Golang Software Engineer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4337640170/",
							Location:             "Toronto, ON",
							Date:                 mustDate("2025-11-20"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Senior Golang Developer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4348328585/",
							Location:             "Budapest, Budapest, Hungary",
							Date:                 mustDate("2025-12-05"),
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
				domain.Erlang: {
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
			Name:    "1touch.io",
			Website: "https://1touch.io/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                18546264,
				IDs:               nil,
				Alias:             "1touch-io",
				Name:              "1touch.io",
				Followers:         "7K",
				Employees:         "51-200",
				AssociatedMembers: "68",
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
							URL:                  "https://www.linkedin.com/jobs/view/4279716505/",
							Location:             "Poland",
							Date:                 mustDate("2025-08-04"),
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
				domain.Erlang: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
			},
			ShortDescription: "1touch.io Inventa delivers a fully autonomous, holistic, and up-to-date view of an organization’s sensitive data usage",
			Industries: []domain.Industry{
				domain.IndustryCyberSecurity,
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "italki",
			Website: "https://www.italki.com/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                167445,
				IDs:               nil,
				Alias:             "italki",
				Name:              "italki",
				Followers:         "118K",
				Employees:         "51-200",
				AssociatedMembers: "4,134",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:     "italki",
				Followers: "8",
				Verified:  true,
			},
			BlindProfile: domain.BlindProfile{
				Alias: "",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias: "",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-italki-EI_IE348674.11,17.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/italki-Reviews-E348674.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/italki-Jobs-E348674.htm",
				Jobs:        "",
				Reviews:     "645",
				Salaries:    "435",
				ReviewsRate: "4.4",
				Verified:    false,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "",
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 1,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Back-End Developer (Go)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4278782092/",
							Location:             "Community of Madrid, Spain",
							Date:                 mustDate("2025-08-02"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Back-End Developer (Go)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4312070903/",
							Location:             "Community of Madrid, Spain",
							Date:                 mustDate("2025-10-08"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Back-End Developer (Go)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4338774285/",
							Location:             "Community of Madrid, Spain",
							Date:                 mustDate("2025-11-25"),
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
				domain.Erlang: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
			},
			ShortDescription: "Language learning marketplace that brings students and teachers together for 1-on-1 and Group-class online language lessons",
			Industries: []domain.Industry{
				domain.IndustryEdTech,
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Revvity",
			Website: "https://www.revvity.com/",
			Careers: "https://jobs.revvity.com/",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                90499149,
				IDs:               []int{90499149, 91410483},
				Alias:             "revvity",
				Name:              "Revvity",
				Followers:         "217K",
				Employees:         "10K+",
				AssociatedMembers: "3,374",
				Verified:          false,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:     "Revvity",
				Followers: "175",
				Verified:  false,
			},
			BlindProfile: domain.BlindProfile{
				Alias: "",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias: "",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Revvity-EI_IE8682798.11,18.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Revvity-Reviews-E8682798.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Revvity-Jobs-E8682798.htm",
				Jobs:        "149",
				Reviews:     "208",
				Salaries:    "462",
				ReviewsRate: "3.9",
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
					GitHubRepositoriesCount: 4,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Senior Software Engineer (Scala, Spark)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4254962827/",
							Location:             "Boston, MA",
							Date:                 mustDate("2025-10-08", "2025-09-16", "2025-08-06"),
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
				domain.Erlang: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
			},
			ShortDescription: "Revvity provides health science solutions, technologies, expertise and services that deliver complete workflows from discovery to development, and diagnosis to cure",
			Industries:       []domain.Industry{
				// NOP
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Seagate Technology",
			Website: "https://www.seagate.com/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                2868,
				IDs:               nil,
				Alias:             "seagate-technology",
				Name:              "Seagate Technology",
				Followers:         "280K",
				Employees:         "10K+",
				AssociatedMembers: "16,389",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:     "seagate",
				Followers: "240",
				Verified:  true,
			},
			BlindProfile: domain.BlindProfile{
				Alias:       "Seagate-Technology",
				Employees:   "10,000+",
				Salary:      "$16K ~ $315K a year",
				Reviews:     "54",
				ReviewsRate: "3.5",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "seagate",
				Employees: "16,000",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Seagate-Technology-EI_IE1864.11,29.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Seagate-Technology-Reviews-E1864.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Seagate-Technology-Jobs-E1864.htm",
				Jobs:        "79",
				Reviews:     "3.3K",
				Salaries:    "5.4K",
				ReviewsRate: "3.7",
				Verified:    false,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "",
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 15,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Cloud Software Engineer (Golang)",
							ShortDescription:     "EC2, AWS S3",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4280834608/",
							Location:             "Poland",
							Date:                 mustDate("2025-08-06"),
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
				domain.Erlang: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
			},
			ShortDescription: "Computer Hardware Manufacturing",
			Industries:       []domain.Industry{
				// NOP
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Batuta",
			Website: "https://batuta.com/",
			Careers: "",
			About:   "https://batuta.com/about/",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                18752548,
				IDs:               nil,
				Alias:             "batutaespm",
				Name:              "Batuta",
				Followers:         "16K",
				Employees:         "51-200",
				AssociatedMembers: "102",
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
							Title:                "Software Developer Golang",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4280995720/",
							Location:             "Argentina",
							Date:                 mustDate("2025-08-07"),
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
				domain.Erlang: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
			},
			ShortDescription: "By intelligently orchestrating every endpoint, the platform eliminates complexity and fragmentation, turning noise into actionable insights",
			Industries: []domain.Industry{
				domain.IndustryCyberSecurity,
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Nozomi Networks",
			Website: "https://www.nozominetworks.com/",
			Careers: "https://www.nozominetworks.com/company/careers",
			About:   "https://www.nozominetworks.com/company",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                5093151,
				IDs:               nil,
				Alias:             "nozomi-networks-sa",
				Name:              "Nozomi Networks",
				Followers:         "86K",
				Employees:         "201-500",
				AssociatedMembers: "365",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:     "NozomiNetworks",
				Followers: "54",
				Verified:  false,
			},
			BlindProfile: domain.BlindProfile{
				Alias: "",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "nozomi-networks",
				Employees: "351",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Nozomi-Networks-EI_IE2409639.11,26.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Nozomi-Networks-Reviews-E2409639.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Nozomi-Networks-Jobs-E2409639.htm",
				Jobs:        "31",
				Reviews:     "83",
				Salaries:    "58",
				ReviewsRate: "4.2",
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
							Title:                "Software Engineer (Python/Go)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4280758824/",
							Location:             "Switzerland",
							Date:                 mustDate("2025-08-06"),
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
				domain.Erlang: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
			},
			ShortDescription: "We protect the world’s critical infrastructure from cyber threats",
			Industries: []domain.Industry{
				domain.IndustryCyberSecurity,
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Nasuni",
			Website: "https://www.nasuni.com/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                580624,
				IDs:               nil,
				Alias:             "nasuni",
				Name:              "Nasuni",
				Followers:         "19K",
				Employees:         "501-1K",
				AssociatedMembers: "563",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "",
				Verified: false,
			},
			BlindProfile: domain.BlindProfile{
				Alias:       "Nasuni",
				Employees:   "201 to 500",
				Salary:      "",
				Reviews:     "2",
				ReviewsRate: "2.5",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias: "",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Nasuni-EI_IE417222.11,17.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Nasuni-Reviews-E417222.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Nasuni-Jobs-E417222.htm",
				Jobs:        "31",
				Reviews:     "118",
				Salaries:    "267",
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
							Title:                "Principal C++/Go/Linux Engineer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4266776966/",
							Location:             "Boston, MA",
							Date:                 mustDate("2025-09-20", "2025-08-06"),
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
				domain.Erlang: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
			},
			ShortDescription: "Nasuni is a unified file data platform for enterprises facing an explosion of unstructured data",
			Industries: []domain.Industry{
				domain.IndustryCloudComputing,
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Simplisales",
			Website: "https://simplisales.com/",
			Careers: "",
			About:   "https://simplisales.com/about-us",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                84357358,
				IDs:               nil,
				Alias:             "simplisales",
				Name:              "Simplisales",
				Followers:         "602",
				Employees:         "11-50",
				AssociatedMembers: "12",
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
							URL:                  "https://www.linkedin.com/jobs/view/4278268730/",
							Location:             "Istanbul, Türkiye",
							Date:                 mustDate("2025-08-06"),
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
				domain.Erlang: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
			},
			ShortDescription: "B2B commerce platform tailored for wholesalers, distributors, and manufacturers",
			Industries: []domain.Industry{
				domain.IndustryECommerce,
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "BioCatch",
			Website: "https://www.biocatch.com/",
			Careers: "https://www.biocatch.com/company",
			About:   "https://www.biocatch.com/cybersecurity-careers",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                2689103,
				IDs:               nil,
				Alias:             "biocatch",
				Name:              "BioCatch",
				Followers:         "33K",
				Employees:         "201-500",
				AssociatedMembers: "413",
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
						// Waiting for the next job to be posted
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
				domain.Erlang: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
			},
			ShortDescription: "BioCatch prevents financial crime by recognizing patterns in human behavior",
			Industries: []domain.Industry{
				domain.IndustryCyberSecurity,
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Infact",
			Website: "https://infact.io/",
			Careers: "https://infact.io/about",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                86375533,
				IDs:               nil,
				Alias:             "infactsystems",
				Name:              "Infact",
				Followers:         "3K",
				Employees:         "11-50",
				AssociatedMembers: "22",
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
							Title:                "Full stack engineer (React & Golang)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4281394183/",
							Location:             "London, England, United Kingdom",
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
				domain.Erlang: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
			},
			ShortDescription: "Our mission is to increase responsible lending through fair, transparent, and up to date credit reporting. We provide effective consumer data for decisioning, through a single API.",
			Industries: []domain.Industry{
				domain.IndustryFinTech,
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Planet",
			Website: "https://www.weareplanet.com/",
			Careers: "https://www.weareplanet.com/careers",
			About:   "https://www.weareplanet.com/about",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                114374,
				IDs:               []int{114374, 2240839},
				Alias:             "planet",
				Name:              "Planet",
				Followers:         "86K",
				Employees:         "1K-5K",
				AssociatedMembers: "2,784",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:     "weareplanet",
				Followers: "160",
				Verified:  false,
			},
			BlindProfile: domain.BlindProfile{
				Alias: "",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias: "",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Planet-EI_IE756506.11,17.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Planet-Reviews-E756506.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Planet-Jobs-E756506.htm",
				Jobs:        "83",
				Reviews:     "343",
				Salaries:    "671",
				ReviewsRate: "3.0",
				Verified:    false,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "",
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 1,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Backend Engineer – Golang",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4240429040/",
							Location:             "Porto, Porto, Portugal",
							Date:                 mustDate("2025-09-21", "2025-08-09"),
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
				domain.Erlang: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
			},
			ShortDescription: "Planet is a global provider of integrated technology and payments solutions for retail and hospitality customers",
			Industries: []domain.Industry{
				domain.IndustryFinTech,
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Swapcard",
			Website: "https://www.swapcard.com/",
			Careers: "https://www.swapcard.com/careers",
			About:   "https://www.swapcard.com/about-swapcard",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                3177729,
				IDs:               nil,
				Alias:             "swapcard",
				Name:              "Swapcard",
				Followers:         "19K",
				Employees:         "51-200",
				AssociatedMembers: "180",
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
							Title:                "Golang Backend Software Engineer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4280234585/",
							Location:             "Valletta, Valletta, Malta",
							Date:                 mustDate("2025-08-09"),
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
				domain.Erlang: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
			},
			ShortDescription: "AI-powered event engagement platform for associations, trade shows, and conference organizers",
			Industries:       []domain.Industry{
				// NOP
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Eachlabs",
			Website: "https://www.eachlabs.ai/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                101532219,
				IDs:               nil,
				Alias:             "eachlabs",
				Name:              "Eachlabs",
				Followers:         "2K",
				Employees:         "2-10",
				AssociatedMembers: "10",
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
							URL:                  "https://www.linkedin.com/jobs/view/4210762801/",
							Location:             "Türkiye",
							Date:                 mustDate("2025-08-10"),
							WithSalary:           false,
							Remote:               true,
						},
						{
							Title:                "Senior Software Engineer (Golang)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4293947042/",
							Location:             "Greater Istanbul",
							Date:                 mustDate("2025-09-04"),
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
				domain.Erlang: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
			},
			ShortDescription: "Turn your AI models into secure, fast, and infinitely scalable inference endpoints",
			Industries: []domain.Industry{
				domain.IndustryDevOps,
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "SandboxAQ",
			Website: "https://www.sandboxaq.com/",
			Careers: "https://www.sandboxaq.com/careers",
			About:   "https://www.sandboxaq.com/company/about-us",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                79089195,
				IDs:               nil,
				Alias:             "sandboxaq",
				Name:              "SandboxAQ",
				Followers:         "43K",
				Employees:         "51-200",
				AssociatedMembers: "330",
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
				Alias:     "sandboxaq",
				Employees: "126",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-SandboxAQ-EI_IE7881847.11,20.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/SandboxAQ-Reviews-E7881847.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/SandboxAQ-Jobs-E7881847.htm",
				Jobs:        "",
				Reviews:     "15",
				Salaries:    "35",
				ReviewsRate: "4.5",
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
							Title:                "Senior Software Engineer, Golang",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4185662443/",
							Location:             "United States",
							Date:                 mustDate("2025-08-11"),
							WithSalary:           true, // $154k/yr - $216k/yr
							Remote:               true,
						},
						{
							Title:                "Senior Software Engineer, Golang",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4295340495/",
							Location:             "United States",
							Date:                 mustDate("2025-09-04"),
							WithSalary:           true, // $154k/yr - $216k/yr
							Remote:               true,
						},
						{
							Title:                "Senior Software Engineer, Golang",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4311340543/",
							Location:             "United States",
							Date:                 mustDate("2025-10-08"),
							WithSalary:           true, // $154k/yr - $216k/yr
							Remote:               true,
						},
					},
				},
				domain.Rust: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Staff Rust Software Engineer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4319834867/",
							Location:             "Canada",
							Date:                 mustDate("2025-10-31"),
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
				domain.Erlang: {
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
			Name:    "Roslin",
			Website: "https://www.roslin.us/",
			Careers: "https://www.roslin.us/careers",
			About:   "https://www.roslin.us/about-us",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                79301238,
				IDs:               nil,
				Alias:             "roslinhq",
				Name:              "Roslin",
				Followers:         "1K",
				Employees:         "201-500",
				AssociatedMembers: "113",
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
							URL:                  "https://www.linkedin.com/jobs/view/4280646692/",
							Location:             "Warsaw, Mazowieckie, Poland",
							Date:                 mustDate("2025-08-11"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Senior Golang Engineer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4291727451/",
							Location:             "Iaşi, Iaşi, Romania",
							Date:                 mustDate("2025-09-01"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Senior Golang Engineer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4297420276/",
							Location:             "Bucharest, Romania",
							Date:                 mustDate("2025-09-12"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Middle Golang Engineer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4313386241/",
							Location:             "Chişinău, Chișinău, Moldova",
							Date:                 mustDate("2025-10-18"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Senior Golang Engineer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4348197360/",
							Location:             "Warsaw, Mazowieckie, Poland",
							Date:                 mustDate("2025-12-03"),
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
				domain.Erlang: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
			},
			ShortDescription: "Management software and call center services for the transportation industry",
			Industries:       []domain.Industry{
				// NOP
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Lenovo",
			Website: "https://www.lenovo.com/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                3653,
				IDs:               nil,
				Alias:             "lenovo",
				Name:              "Lenovo",
				Followers:         "1M",
				Employees:         "10K+",
				AssociatedMembers: "45,133",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:     "lenovo",
				Followers: "160",
				Verified:  false,
			},
			BlindProfile: domain.BlindProfile{
				Alias:       "Lenovo",
				Employees:   "10,000+",
				Salary:      "$44K ~ $250K a year",
				Reviews:     "93",
				ReviewsRate: "3.6",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "lenovo",
				Employees: "72,000",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Lenovo-EI_IE8034.11,17.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Lenovo-Reviews-E8034.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Lenovo-Jobs-E8034.htm",
				Jobs:        "",
				Reviews:     "5K",
				Salaries:    "7.5K",
				ReviewsRate: "3.9",
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
							Title:                "Golang Developer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4271553021/",
							Location:             "Petaling Jaya, Selangor, Malaysia",
							Date:                 mustDate("2025-08-15"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Golang Developer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4288753154/",
							Location:             "Petaling Jaya, Selangor, Malaysia",
							Date:                 mustDate("2025-08-22"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Golang Developer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4303653528/",
							Location:             "Petaling Jaya, Selangor, Malaysia",
							Date:                 mustDate("2025-09-22"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Golang Developer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4326115280/",
							Location:             "Singapore, Singapore",
							Date:                 mustDate("2025-12-08"),
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
				domain.Erlang: {
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
			Name:    "EasyPost",
			Website: "https://www.easypost.com/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                3085688,
				IDs:               nil,
				Alias:             "easypost",
				Name:              "EasyPost",
				Followers:         "21K",
				Employees:         "201-500",
				AssociatedMembers: "383",
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
				Alias:     "easypost",
				Employees: "210",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-EasyPost-EI_IE1300009.11,19.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/EasyPost-Reviews-E1300009.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/EasyPost-Jobs-E1300009.htm",
				Jobs:        "3",
				Reviews:     "238",
				Salaries:    "340",
				ReviewsRate: "3.0",
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
							Title:                "Senior Software Engineer (Golang)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4288462487/",
							Location:             "United States",
							Date:                 mustDate("2025-08-20"),
							WithSalary:           true, // $150k/yr - $200k/yr
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
				domain.Erlang: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
			},
			ShortDescription: "EasyPost is a Shipping API that solves complex logistics problems for eCommerce businesses",
			Industries: []domain.Industry{
				domain.IndustryDevOps,
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "N-able",
			Website: "https://www.n-able.com/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                76322470,
				IDs:               nil,
				Alias:             "n-able",
				Name:              "N-able",
				Followers:         "57K",
				Employees:         "1K-5K",
				AssociatedMembers: "2,094",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:     "n-able",
				Followers: "70",
				Verified:  false,
			},
			BlindProfile: domain.BlindProfile{
				Alias: "",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "n-able-technologies",
				Employees: "1,454",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-N-able-EI_IE5067371.11,17.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/N-able-Reviews-E5067371.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/N-able-Jobs-E5067371.htm",
				Jobs:        "69",
				Reviews:     "359",
				Salaries:    "605",
				ReviewsRate: "3.9",
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
							Title:                "Senior Software Engineer (Python, Golang)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4272812123/",
							Location:             "Ottawa, ON",
							Date:                 mustDate("2025-09-10", "2025-08-21"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Software Engineer (Golang)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4347824357/",
							Location:             "Edinburgh, Scotland, United Kingdom",
							Date:                 mustDate("2025-11-29"),
							WithSalary:           false,
							Remote:               false,
						},
					},
				},
				domain.Rust: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Staff Software Engineer (Rust)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4312455087/",
							Location:             "Warsaw, Mazowieckie, Poland",
							Date:                 mustDate("2025-11-28", "2025-11-06", "2025-10-17"),
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
				domain.Erlang: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
			},
			ShortDescription: "N-able’s mission is to protect businesses against evolving cyberthreats with a unified cyber resiliency platform to manage, secure, and recover",
			Industries: []domain.Industry{
				domain.IndustryCyberSecurity,
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Forest",
			Website: "https://www.forest.me/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                36086360,
				IDs:               nil,
				Alias:             "forestbikeapp",
				Name:              "Forest",
				Followers:         "7K",
				Employees:         "51-200",
				AssociatedMembers: "334",
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
							Title:                "Back-End Engineer (Golang)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4288290559/",
							Location:             "Santiago, Santiago Metropolitan Region, Chile",
							Date:                 mustDate("2025-08-22"),
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
				domain.Erlang: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
			},
			ShortDescription: "London's most affordable and sustainable eBike app",
			Industries:       []domain.Industry{
				// NOP
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Postman",
			Website: "https://www.postman.com/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                3795851,
				IDs:               nil,
				Alias:             "postman-platform",
				Name:              "Postman",
				Followers:         "526K",
				Employees:         "501-1K",
				AssociatedMembers: "3,129",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:     "postmanlabs",
				Followers: "3.4k",
				Verified:  false,
			},
			BlindProfile: domain.BlindProfile{
				Alias:       "Postman",
				Employees:   "201 to 500",
				Salary:      "",
				Reviews:     "188",
				ReviewsRate: "3.4",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "postman",
				Employees: "800",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Postman-EI_IE1926052.11,18.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Postman-Reviews-E1926052.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Postman-Jobs-E1926052.htm",
				Jobs:        "4",
				Reviews:     "245",
				Salaries:    "510",
				ReviewsRate: "3.5",
				Verified:    false,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "",
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 13,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Backend Software Engineer, API Client (Node.js and/or Golang)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4278739224/",
							Location:             "San Francisco, CA",
							Date:                 mustDate("2025-09-14", "2025-08-22"),
							WithSalary:           true, // $180k/yr - $240k/yr
							Remote:               false,
						},
						{
							Title:                "Backend and System Engineer (Node.js or Golang)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4058765320/",
							Location:             "New York, NY",
							Date:                 mustDate("2025-11-26"),
							WithSalary:           true, // $150k/yr - $245k/yr
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
				domain.Erlang: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
			},
			ShortDescription: "API platform",
			Industries: []domain.Industry{
				domain.IndustryDevOps,
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Havas",
			Website: "https://www.havas.com/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                157305,
				IDs:               []int{17319, 42152, 47476, 64196, 100086, 150292, 157304, 157305, 164133, 164304, 203609, 211535, 269842, 309870, 393791, 453550, 536111, 643295, 650442, 957406, 975440, 1046784, 1076618, 1195336, 1342316, 1522649, 1569087, 1639124, 1870266, 1941122, 2233228, 2285252, 2628806, 2770689, 2778033, 2892571, 3116226, 3124684, 3136819, 3497715, 3533001, 3774123, 4839659, 5010069, 6608955, 7931229, 10449384, 10806387, 11104454, 11262319, 13234016, 13268754, 18476963, 18666924, 24767234, 27028517, 27148810, 28167444, 38110178, 42145892, 46243929, 70251746, 71646516, 71729933, 72310129, 73200768, 73853527, 76631708, 77285270, 77763425, 77960624, 79516873, 79922611, 80112741, 82330593, 85936711, 86106758, 99893734, 99899910},
				Alias:             "havas",
				Name:              "Havas",
				Followers:         "255K",
				Employees:         "10K+",
				AssociatedMembers: "15,575",
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
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Havas-Worldwide-EI_IE14455.11,26.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Havas-Worldwide-Reviews-E14455.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Havas-Worldwide-Jobs-E14455.htm",
				Jobs:        "1",
				Reviews:     "1.3K",
				Salaries:    "2.3K",
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
							Title:                "Senior Golang Developer",
							ShortDescription:     "Azure, DevOps",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4285626267/",
							Location:             "Ho Chi Minh City, Vietnam",
							Date:                 mustDate("2025-10-25", "2025-08-22"),
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
				domain.Erlang: {
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
			Name:    "Tietoevry",
			Website: "https://www.tietoevry.com/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                40669199,
				IDs:               nil,
				Alias:             "tietoevry",
				Name:              "Tietoevry",
				Followers:         "462K",
				Employees:         "10K+",
				AssociatedMembers: "14,643",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "",
				Verified: false,
			},
			BlindProfile: domain.BlindProfile{
				Alias:       "TietoEVRY",
				Employees:   "10,000+",
				Salary:      "",
				Reviews:     "8",
				ReviewsRate: "3.5",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "tietoevry",
				Employees: "24,000",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Tietoevry-EI_IE315536.11,20.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Tietoevry-Reviews-E315536.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Tietoevry-Jobs-E315536.htm",
				Jobs:        "317",
				Reviews:     "4.3K",
				Salaries:    "4.6K",
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
							Title:                "Senior Golang Developer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4289016402/",
							Location:             "Sofia, Sofia City, Bulgaria",
							Date:                 mustDate("2025-08-22"),
							WithSalary:           false,
							Remote:               true,
						},
						{
							Title:                "Golang Developer",
							ShortDescription:     "Telecom",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4294098298/",
							Location:             "Bengaluru, Karnataka, India",
							Date:                 mustDate("2025-11-06"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Backend Developer (Go/Python)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4312121952/",
							Location:             "Hyderabad, Telangana, India",
							Date:                 mustDate("2025-11-22"),
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
					Vacancies: []domain.Vacancy{
						{
							Title:                "Fullstack Java/Scala developer",
							ShortDescription:     "Banking",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4314013932/",
							Location:             "Trondheim, Trøndelag, Norway",
							Date:                 mustDate("2025-11-05"),
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
				domain.Erlang: {
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
			Name:    "Torq",
			Website: "https://torq.io/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                64511369,
				IDs:               nil,
				Alias:             "torqio",
				Name:              "Torq",
				Followers:         "30K",
				Employees:         "201-500",
				AssociatedMembers: "332",
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
				Alias:     "torq",
				Employees: "90",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Torq-EI_IE7415271.11,15.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Torq-Reviews-E7415271.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Torq-Jobs-E7415271.htm",
				Jobs:        "",
				Reviews:     "17",
				Salaries:    "25",
				ReviewsRate: "4.7",
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
							Title:                "Senior Backend Engineer (Python + Go)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4289053211/",
							Location:             "Tel Aviv-Yafo, Tel Aviv District, Israel",
							Date:                 mustDate("2025-08-22"),
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
				domain.Erlang: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
			},
			ShortDescription: "Autonomous Security Operations",
			Industries: []domain.Industry{
				domain.IndustryCyberSecurity,
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Unanet",
			Website: "https://unanet.com/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                71832,
				IDs:               nil,
				Alias:             "unanet-technologies",
				Name:              "Unanet",
				Followers:         "13K",
				Employees:         "201-500",
				AssociatedMembers: "464",
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
				Alias:     "unanet",
				Employees: "330",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Unanet-EI_IE355390.11,17.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Unanet-Reviews-E355390.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Unanet-Jobs-E355390.htm",
				Jobs:        "",
				Reviews:     "128",
				Salaries:    "240",
				ReviewsRate: "3.9",
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
							Title:                "Staff Software Engineer (Vue.js/Go/TypeScript)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4289364017/",
							Location:             "United States",
							Date:                 mustDate("2025-09-10", "2025-09-02", "2025-08-22"),
							WithSalary:           true, // $152k/yr - $170k/yr + bonus, stock options
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
				domain.Erlang: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
			},
			ShortDescription: "Provider of ERP and CRM solutions purpose-built for Government Contractors, AEC & Prof Services",
			Industries:       []domain.Industry{
				// NOP
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Entrust",
			Website: "https://www.entrust.com/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                8011,
				IDs:               nil,
				Alias:             "entrust",
				Name:              "Entrust",
				Followers:         "146K",
				Employees:         "1K-5K",
				AssociatedMembers: "3,772",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:     "EntrustCorporation",
				Followers: "63",
				Verified:  false,
			},
			BlindProfile: domain.BlindProfile{
				Alias:       "Entrust",
				Employees:   "1,001 to 5,000",
				Salary:      "",
				Reviews:     "11",
				ReviewsRate: "3.5",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "entrust",
				Employees: "3,000",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Entrust-EI_IE1104613.11,18.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Entrust-Reviews-E1104613.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Entrust-Jobs-E1104613.htm",
				Jobs:        "59",
				Reviews:     "483",
				Salaries:    "938",
				ReviewsRate: "3.6",
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
							Title:                "Software Developer II – C/Python/Go",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4289131757/",
							Location:             "Sunrise, FL",
							Date:                 mustDate("2025-08-22"),
							WithSalary:           true, // $100.8k/yr - $147.8k/yr
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
				domain.Erlang: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
			},
			ShortDescription: "Identity-centric security solutions, providing an integrated platform of scalable, AI-enabled security offerings",
			Industries: []domain.Industry{
				domain.IndustryCyberSecurity,
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Yolcu360",
			Website: "https://yolcu360.com/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                7934593,
				IDs:               nil,
				Alias:             "yolcu-360",
				Name:              "Yolcu360",
				Followers:         "19K",
				Employees:         "51-200",
				AssociatedMembers: "140",
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
							Title:                "Backend Developer (Go)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4289245232/",
							Location:             "Istanbul, Türkiye",
							Date:                 mustDate("2025-08-22"),
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
				domain.Erlang: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
			},
			ShortDescription: "Travel tech company operating a global marketplace for car rentals and flight bookings",
			Industries: []domain.Industry{
				domain.IndustryTravelTech,
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Leaseweb",
			Website: "https://www.leaseweb.com/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                242620,
				IDs:               nil,
				Alias:             "leaseweb",
				Name:              "Leaseweb",
				Followers:         "25K",
				Employees:         "501-1K",
				AssociatedMembers: "630",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:     "Leaseweb",
				Followers: "33",
				Verified:  false,
			},
			BlindProfile: domain.BlindProfile{
				Alias: "",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "leaseweb",
				Employees: "510",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Leaseweb-EI_IE697046.11,19.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Leaseweb-Reviews-E697046.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Leaseweb-Jobs-E697046.htm",
				Jobs:        "18",
				Reviews:     "213",
				Salaries:    "329",
				ReviewsRate: "3.4",
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
							Title:                "Senior Golang Developer",
							ShortDescription:     "Developer Platform Team",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4233711269/",
							Location:             "European Union",
							Date:                 mustDate("2025-08-22"),
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
				domain.Erlang: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
			},
			ShortDescription: "Cloud services and Infrastructure-as-a-Service (IaaS) provider",
			Industries: []domain.Industry{
				domain.IndustryCloudComputing,
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Netcore Cloud",
			Website: "https://netcorecloud.com/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                31757,
				IDs:               nil,
				Alias:             "netcore-cloud",
				Name:              "Netcore Cloud",
				Followers:         "168K",
				Employees:         "1K-5K",
				AssociatedMembers: "1,266",
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
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Netcore-Cloud-EI_IE519713.11,24.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Netcore-Cloud-Reviews-E519713.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Netcore-Cloud-Jobs-E519713.htm",
				Jobs:        "",
				Reviews:     "638",
				Salaries:    "566",
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
							Title:                "Senior Go Developer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4289477990/",
							Location:             "Mumbai, Maharashtra, India",
							Date:                 mustDate("2025-08-22"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Senior Go Developer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4347018010/",
							Location:             "Mumbai Metropolitan Region",
							Date:                 mustDate("2025-11-22"),
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
				domain.Erlang: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
			},
			ShortDescription: "Comprehensive Customer Engagement & Experience platform to build meaningful connections between consumers and brands",
			Industries:       []domain.Industry{
				// NOP
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "FluidCloud",
			Website: "https://www.fluidcloud.com/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                106725587,
				IDs:               nil,
				Alias:             "fluidcloud",
				Name:              "FluidCloud",
				Followers:         "2K",
				Employees:         "11-50",
				AssociatedMembers: "15",
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
							URL:                  "https://www.linkedin.com/jobs/view/4287515381/",
							Location:             "Pleasanton, CA",
							Date:                 mustDate("2025-08-19"),
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
				domain.Erlang: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
			},
			ShortDescription: "Multi-cloud infrastructure",
			Industries: []domain.Industry{
				domain.IndustryCloudComputing,
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "iFood",
			Website: "https://www.ifood.com.br/",
			Careers: "https://carreiras.ifood.com.br/",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                247645,
				IDs:               nil,
				Alias:             "ifood-",
				Name:              "iFood",
				Followers:         "2M",
				Employees:         "5K-10K",
				AssociatedMembers: "21,806",
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
							Title:                "Senior Golang Software Engineer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4096799427/",
							Location:             "Brazil",
							Date:                 mustDate("2025-10-01", "2025-09-10", "2025-08-20"),
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
				domain.Erlang: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
			},
			ShortDescription: "FoodTech company",
			Industries:       []domain.Industry{
				// NOP
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "ampliFI Loyalty Solutions",
			Website: "https://www.amplifiloyalty.com/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                33516651,
				IDs:               nil,
				Alias:             "amplifiloyalty",
				Name:              "ampliFI Loyalty Solutions",
				Followers:         "2K",
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
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-AmpliFi-Loyalty-Solutions-EI_IE4449023.11,36.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/AmpliFi-Loyalty-Solutions-Reviews-E4449023.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/AmpliFi-Loyalty-Solutions-Jobs-E4449023.htm",
				Jobs:        "5",
				Reviews:     "11",
				Salaries:    "34",
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
					Vacancies:               []domain.Vacancy{},
				},
				domain.Elixir: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Clojure: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Clojure / Java Developer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4287526043/",
							Location:             "Naperville, IL",
							Date:                 mustDate("2025-08-19"),
							WithSalary:           true, // $110k/yr - $120k/yr
							Remote:               false,
						},
						{
							Title:                "Clojure / Java Developer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4320054232/",
							Location:             "Naperville, IL",
							Date:                 mustDate("2025-10-31"),
							WithSalary:           true, // $110k/yr - $120k/yr
							Remote:               false,
						},
					},
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
			ShortDescription: "ampliFI Loyalty Solutions is a loyalty and engagement program provider specializing in data-driven, custom rewards programs for financial institutions nationwide",
			Industries:       []domain.Industry{
				// NOP
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Rublon",
			Website: "https://rublon.com/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                2772205,
				IDs:               nil,
				Alias:             "rublon",
				Name:              "Rublon",
				Followers:         "2K",
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
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Rust: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Rust/C++ Software Engineer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4289132808/",
							Location:             "Zielona Góra, Lubuskie, Poland",
							Date:                 mustDate("2025-08-21"),
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
				domain.Erlang: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
			},
			ShortDescription: "Protect access to your networks, servers and applications using workforce multi-factor authentication ",
			Industries: []domain.Industry{
				domain.IndustryCyberSecurity,
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Bigabid",
			Website: "https://www.bigabid.com/",
			Careers: "https://www.bigabid.com/careers/",
			About:   "https://www.bigabid.com/about-us/",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                10033874,
				IDs:               nil,
				Alias:             "bigabid",
				Name:              "Bigabid",
				Followers:         "5K",
				Employees:         "51-200",
				AssociatedMembers: "106",
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
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-BigaBid-EI_IE3055132.11,18.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/BigaBid-Reviews-E3055132.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/BigaBid-Jobs-E3055132.htm",
				Jobs:        "16",
				Reviews:     "30",
				Salaries:    "33",
				ReviewsRate: "3.7",
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
							Title:                "Senior Big Data Engineer",
							ShortDescription:     "6+ years of proven experience as a Data/Backend Engineer with strong coding skills in Scala, Java, or Python",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4290544541/",
							Location:             "Tel Aviv-Yafo, Tel Aviv District, Israel",
							Date:                 mustDate("2025-08-24"),
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
				domain.Erlang: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
			},
			ShortDescription: "Programmatic User Acquisition and Retargeting DSP for Mobile Apps",
			Industries: []domain.Industry{
				domain.IndustryAdTech,
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Permutive",
			Website: "https://permutive.com/",
			Careers: "https://permutive.com/careers",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                9314455,
				IDs:               nil,
				Alias:             "permutive",
				Name:              "Permutive",
				Followers:         "10K",
				Employees:         "51-200",
				AssociatedMembers: "108",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:     "permutive",
				Followers: "23",
				Verified:  true,
			},
			BlindProfile: domain.BlindProfile{
				Alias: "",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias: "",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Permutive-EI_IE2116966.11,20.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Permutive-Reviews-E2116966.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Permutive-Jobs-E2116966.htm",
				Jobs:        "",
				Reviews:     "60",
				Salaries:    "164",
				ReviewsRate: "3.7",
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
					GitHubRepositoriesCount: 3,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Scala Engineer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4288896427/",
							Location:             "London Area, United Kingdom",
							Date:                 mustDate("2025-08-26"),
							WithSalary:           true, // £70k/yr - £80k/yr
							Remote:               false,
						},
						{
							Title:                "Senior Scala Engineer (Tech Lead)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4304163804/",
							Location:             "London, England, United Kingdom",
							Date:                 mustDate("2025-09-22"),
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
				domain.Erlang: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
			},
			ShortDescription: "Permutive puts publishers in control of their revenue by enabling them to reach and monetize 100% of their audience while respecting consumer privacy",
			Industries: []domain.Industry{
				domain.IndustryAdTech,
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Innovatrics",
			Website: "https://www.innovatrics.com/",
			Careers: "https://careers.innovatrics.com/",
			About:   "https://www.innovatrics.com/about-us/",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                1016818,
				IDs:               nil,
				Alias:             "innovatrics",
				Name:              "Innovatrics",
				Followers:         "10K",
				Employees:         "51-200",
				AssociatedMembers: "204",
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
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Innovatrics-EI_IE1862436.11,22.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Innovatrics-Reviews-E1862436.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Innovatrics-Jobs-E1862436.htm",
				Jobs:        "2",
				Reviews:     "22",
				Salaries:    "47",
				ReviewsRate: "3.4",
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
							Title:                "C++/Rust Software Engineer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4291261200/",
							Location:             "Brno, South Moravia, Czechia",
							Date:                 mustDate("2025-08-26"),
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
				domain.Erlang: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
			},
			ShortDescription: "Innovatrics is an independent EU-based provider of trusted biometric solutions for governments and enterprises",
			Industries:       []domain.Industry{
				// NOP
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Spire",
			Website: "https://spire.com/",
			Careers: "https://spire.com/careers/job-openings/",
			About:   "https://spire.com/about-us/",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                3506863,
				IDs:               nil,
				Alias:             "spireglobal",
				Name:              "Spire",
				Followers:         "53K",
				Employees:         "201-500",
				AssociatedMembers: "494",
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
				Alias:     "spire-global",
				Employees: "411",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Spire-Global-EI_IE794910.11,23.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Spire-Global-Reviews-E794910.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Spire-Global-Jobs-E794910.htm",
				Jobs:        "128",
				Reviews:     "142",
				Salaries:    "286",
				ReviewsRate: "2.8",
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
							Title:                "Spacecraft Embedded Software Engineer (Rust)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4291616116/",
							Location:             "Cambridge, ON",
							Date:                 mustDate("2025-10-08", "2025-09-18", "2025-08-26"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Senior Spacecraft Embedded Software Engineer (Rust)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4296629231/",
							Location:             "Boulder, CO",
							Date:                 mustDate("2025-09-06"),
							WithSalary:           true, // $130.5k/yr - $171k/yr
							Remote:               false,
						},
						{
							Title:                "Senior Software Engineer (Rust)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4297667474/",
							Location:             "Luxembourg, Luxembourg, Luxembourg",
							Date:                 mustDate("2025-10-01", "2025-09-09"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Senior Software Engineer (Rust)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4297662626/",
							Location:             "Boulder, CO",
							Date:                 mustDate("2025-10-22"),
							WithSalary:           true, // $130.5k/yr - $171k/yr
							Remote:               false,
						},
						{
							Title:                "Spacecraft Embedded Software Engineer (Rust)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4291616085/",
							Location:             "Boulder, CO",
							Date:                 mustDate("2025-11-20", "2025-10-30"),
							WithSalary:           true, // $90k/yr - $130.5k/yr
							Remote:               false,
						},
						{
							Title:                "Senior Software Engineer (Rust)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4297659747/",
							Location:             "Harwell, England, United Kingdom",
							Date:                 mustDate("2025-11-12"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Senior Software Engineer (Rust)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4297666414/",
							Location:             "Munich, Bavaria, Germany",
							Date:                 mustDate("2025-12-05"),
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
				domain.Erlang: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
			},
			ShortDescription: "A provider of space-based data, analytics, and space services",
			Industries:       []domain.Industry{
				// NOP
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "LG Ad Solutions",
			Website: "https://lgads.tv/",
			Careers: "https://lgads.tv/careers/",
			About:   "https://lgads.tv/company/",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                3100031,
				IDs:               nil,
				Alias:             "lgads",
				Name:              "LG Ad Solutions",
				Followers:         "21K",
				Employees:         "201-500",
				AssociatedMembers: "464",
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
							Title:                "Backend Software Engineer (C++, Golang, Node.js)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4289995599/",
							Location:             "Denver, CO",
							Date:                 mustDate("2025-08-23"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Senior Software Engineer II, Backend (Golang)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4294812095/",
							Location:             "Bengaluru, Karnataka, India",
							Date:                 mustDate("2025-09-04"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Senior Software Engineer II, Backend (Golang/Python)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4324641696/",
							Location:             "Bengaluru, Karnataka, India",
							Date:                 mustDate("2025-11-26"),
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
				domain.Erlang: {
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
			Name:    "Brevo",
			Website: "https://www.brevo.com/",
			Careers: "https://www.brevo.com/careers/",
			About:   "https://www.brevo.com/company/about/",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                2346309,
				IDs:               nil,
				Alias:             "brevo",
				Name:              "Brevo",
				Followers:         "70K",
				Employees:         "501-1K",
				AssociatedMembers: "832",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:     "getbrevo",
				Followers: "41",
				Verified:  true,
			},
			BlindProfile: domain.BlindProfile{
				Alias: "",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias: "",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Brevo-EI_IE809418.11,16.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Brevo-Reviews-E809418.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Brevo-Jobs-E809418.htm",
				Jobs:        "23",
				Reviews:     "321",
				Salaries:    "564",
				ReviewsRate: "3.9",
				Verified:    false,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "",
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 1,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Golang Developer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4289604251/",
							Location:             "Noida, Uttar Pradesh, India",
							Date:                 mustDate("2025-08-26"),
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
				domain.Erlang: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
			},
			ShortDescription: "CRM",
			Industries:       []domain.Industry{
				// NOP
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Hudl",
			Website: "https://www.hudl.com/",
			Careers: "https://www.hudl.com/jobs",
			About:   "https://www.hudl.com/about",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                2464624,
				IDs:               nil,
				Alias:             "hudl",
				Name:              "Hudl",
				Followers:         "94K",
				Employees:         "1K-5K",
				AssociatedMembers: "3,836",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:     "hudl",
				Followers: "150",
				Verified:  true,
			},
			BlindProfile: domain.BlindProfile{
				Alias: "",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias: "",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Hudl-EI_IE826705.11,15.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Hudl-Reviews-E826705.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Hudl-Jobs-E826705.htm",
				Jobs:        "13",
				Reviews:     "1.2K",
				Salaries:    "1.9K",
				ReviewsRate: "4.0",
				Verified:    false,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "",
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 14,
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
					Vacancies: []domain.Vacancy{
						{
							Title:                "Senior Engineer (Clojure)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4294245074/",
							Location:             "London, England, United Kingdom",
							Date:                 mustDate("2025-09-02"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Senior Engineer (Clojure)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4326520542/",
							Location:             "London, England, United Kingdom",
							Date:                 mustDate("2025-12-10"),
							WithSalary:           false,
							Remote:               false,
						},
					},
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
			ShortDescription: "Performance analysis for sports",
			Industries:       []domain.Industry{
				// NOP
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Recorded Future",
			Website: "https://www.recordedfuture.com/",
			Careers: "https://www.recordedfuture.com/work-with-us/jobs",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                678036,
				IDs:               nil,
				Alias:             "recorded-future",
				Name:              "Recorded Future",
				Followers:         "77K",
				Employees:         "501-1K",
				AssociatedMembers: "1,079",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:     "recordedfuture",
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
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Recorded-Future-EI_IE1322290.11,26.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Recorded-Future-Reviews-E1322290.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Recorded-Future-Jobs-E1322290.htm",
				Jobs:        "70",
				Reviews:     "233",
				Salaries:    "454",
				ReviewsRate: "3.7",
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
					Vacancies:               []domain.Vacancy{},
				},
				domain.Elixir: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Backend Software Engineer (Elixir)",
							ShortDescription:     "Attack Surface Intelligence",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4293000776/",
							Location:             "Boston, VA",
							Date:                 mustDate("2025-09-21", "2025-08-30"),
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
				domain.Erlang: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
			},
			ShortDescription: "Threat intelligence company",
			Industries: []domain.Industry{
				domain.IndustryCyberSecurity,
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Opensignal",
			Website: "https://www.opensignal.com/",
			Careers: "https://www.opensignal.com/about/opensignal-careers",
			About:   "https://www.opensignal.com/about-us",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                1954386,
				IDs:               nil,
				Alias:             "opensignal",
				Name:              "Opensignal",
				Followers:         "37K",
				Employees:         "201-500",
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
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Opensignal-EI_IE723974.11,21.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Opensignal-Reviews-E723974.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Opensignal-Jobs-E723974.htm",
				Jobs:        "8",
				Reviews:     "44",
				Salaries:    "113",
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
							Title:                "Senior Scala Data Engineer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4258545583/",
							Location:             "Warsaw, Mazowieckie, Poland",
							Date:                 mustDate("2025-10-12", "2025-09-21", "2025-08-30"),
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
				domain.Erlang: {
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
			Name:    "Football Radar",
			Website: "https://www.footballradar.com/",
			Careers: "https://www.footballradar.com/careers/",
			About:   "https://www.footballradar.com/about/",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                1697361,
				IDs:               nil,
				Alias:             "football-radar-ltd",
				Name:              "Football Radar",
				Followers:         "12K",
				Employees:         "51-200",
				AssociatedMembers: "109",
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
							Title:                "Senior Scala Engineer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4293313610/",
							Location:             "London Area, United Kingdom",
							Date:                 mustDate("2025-08-31"),
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
				domain.Erlang: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
			},
			ShortDescription: "Provider of football betting advice",
			Industries:       []domain.Industry{
				// Betting
			},
			Ignore: true, // Betting
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Yoti",
			Website: "https://www.yoti.com/",
			Careers: "https://www.yoti.com/jobs/",
			About:   "https://www.yoti.com/about/",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                9428931,
				IDs:               nil,
				Alias:             "yoti-digital-identity",
				Name:              "Yoti",
				Followers:         "38K",
				Employees:         "201-500",
				AssociatedMembers: "427",
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
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Yoti-EI_IE1339394.11,15.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Yoti-Reviews-E1339394.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Yoti-Jobs-E1339394.htm",
				Jobs:        "",
				Reviews:     "103",
				Salaries:    "234",
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
							Title:                "Golang Backend Engineer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4294137065/",
							Location:             "London, England, United Kingdom",
							Date:                 mustDate("2025-09-02"),
							WithSalary:           true, // £55k/yr - £60k/yr
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
				domain.Erlang: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
			},
			ShortDescription: "Privacy-focused identity and age verification solutions",
			Industries:       []domain.Industry{
				// NOP
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "SoLo Funds",
			Website: "https://solofunds.com/",
			Careers: "https://solofunds.com/careers/",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                9424483,
				IDs:               nil,
				Alias:             "solo-funds",
				Name:              "SoLo Funds",
				Followers:         "5K",
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
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-SoLo-Funds-EI_IE4677633.11,21.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/SoLo-Funds-Reviews-E4677633.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/SoLo-Funds-Jobs-E4677633.htm",
				Jobs:        "",
				Reviews:     "20",
				Salaries:    "27",
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
							Title:                "Senior Software Engineer (Golang)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4294145998/",
							Location:             "Brazil",
							Date:                 mustDate("2025-09-03"),
							WithSalary:           true, // $3,000/month - $7,000/month
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
				domain.Erlang: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
			},
			ShortDescription: "SoLo Funds is a community finance platform",
			Industries: []domain.Industry{
				domain.IndustryFinTech,
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Avey",
			Website: "http://www.avey.ai/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                18320282,
				IDs:               nil,
				Alias:             "avey-ai",
				Name:              "Avey",
				Followers:         "18K",
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
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Backend Engineer (Golang)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4292729748/",
							Location:             "Doha, Qatar",
							Date:                 mustDate("2025-09-03"),
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
				domain.Erlang: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
			},
			ShortDescription: "Avey specializes in researching, designing, and developing novel AI models that empower healthcare providers & patients",
			Industries: []domain.Industry{
				domain.IndustryHealthTech,
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "EVANA AG",
			Website: "https://evana.ai/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                10247434,
				IDs:               nil,
				Alias:             "evana",
				Name:              "EVANA AG",
				Followers:         "3K",
				Employees:         "51-200",
				AssociatedMembers: "22",
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
							URL:                  "https://www.linkedin.com/jobs/view/4294341517/",
							Location:             "Germany",
							Date:                 mustDate("2025-09-03"),
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
				domain.Erlang: {
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
			Name:    "Airspace Link, Inc.",
			Website: "https://airspacelink.com/",
			Careers: "https://airspacelink.com/careers",
			About:   "https://airspacelink.com/about",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                28592338,
				IDs:               nil,
				Alias:             "airspacelink",
				Name:              "Airspace Link, Inc.",
				Followers:         "6K",
				Employees:         "51-200",
				AssociatedMembers: "52",
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
							Title:                "Golang Engineer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4294093341/",
							Location:             "Detroit Metropolitan Area",
							Date:                 mustDate("2025-09-03"),
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
				domain.Erlang: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
			},
			ShortDescription: "Drone Operations Management System",
			Industries:       []domain.Industry{
				// NOP
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Lumel",
			Website: "https://lumel.com/",
			Careers: "https://lumel.com/careers/",
			About:   "https://lumel.com/about-us/",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                75039399,
				IDs:               []int{14621722, 20414266, 27034875, 71563074, 75039399},
				Alias:             "lumeltech",
				Name:              "Lumel",
				Followers:         "12K",
				Employees:         "201-500",
				AssociatedMembers: "371",
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
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Lumel-Technologies-EI_IE5932478.11,29.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Lumel-Technologies-Reviews-E5932478.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Lumel-Technologies-Jobs-E5932478.htm",
				Jobs:        "",
				Reviews:     "45",
				Salaries:    "74",
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
							Title:                "Senior/Lead Backend Developer – Golang/Node.js",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4293417775/",
							Location:             "Chennai, Tamil Nadu, India",
							Date:                 mustDate("2025-09-01"),
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
				domain.Erlang: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
			},
			ShortDescription: "Enterprise Performance Management",
			Industries:       []domain.Industry{
				// NOP
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "RoboMarkets",
			Website: "https://www.robomarkets.com.cy/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                68275262,
				IDs:               nil,
				Alias:             "robomarkets",
				Name:              "RoboMarkets",
				Followers:         "3K",
				Employees:         "51-200",
				AssociatedMembers: "128",
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
							Title:                "Golang developer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4295629301/",
							Location:             "Limassol Municipality, Limassol, Cyprus",
							Date:                 mustDate("2025-09-04"),
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
				domain.Erlang: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
			},
			ShortDescription: "European brokerage company",
			Industries: []domain.Industry{
				domain.IndustryFinTech,
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "bTaskee",
			Website: "https://www.btaskee.com/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                10682858,
				IDs:               nil,
				Alias:             "btaskee",
				Name:              "bTaskee",
				Followers:         "5K",
				Employees:         "201-500",
				AssociatedMembers: "149",
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
							Title:                "Golang Developer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4295317964/",
							Location:             "Ho Chi Minh City, Vietnam",
							Date:                 mustDate("2025-09-04"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Golang Developer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4301174724/",
							Location:             "Ho Chi Minh City, Vietnam",
							Date:                 mustDate("2025-09-16"),
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
				domain.Erlang: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
			},
			ShortDescription: "bTaskee provides on-demand household cleaning and maid services by connecting customers with local experts through a mobile platform",
			Industries:       []domain.Industry{
				// NOP
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Cruit",
			Website: "https://www.joincruit.com/",
			Careers: "",
			About:   "https://www.joincruit.com/about",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                81771726,
				IDs:               nil,
				Alias:             "joincruit",
				Name:              "Cruit",
				Followers:         "1K",
				Employees:         "11-50",
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
							Title:                "Full-Stack Developer Golang",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4295918906/",
							Location:             "Manila, National Capital Region, Philippines",
							Date:                 mustDate("2025-09-04"),
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
				domain.Erlang: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
			},
			ShortDescription: "Cruit is an employee referral software designed to craft highly engaging programs that optimize referral recruitment, boost visibility and increase employee retention",
			Industries:       []domain.Industry{
				// NOP
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Blackpoint Cyber",
			Website: "https://blackpointcyber.com/",
			Careers: "https://blackpointcyber.com/company/careers/",
			About:   "https://blackpointcyber.com/company/",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                9381059,
				IDs:               nil,
				Alias:             "blackpoint",
				Name:              "Blackpoint Cyber",
				Followers:         "18K",
				Employees:         "51-200",
				AssociatedMembers: "191",
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
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Blackpoint-Cyber-EI_IE2940227.11,27.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Blackpoint-Cyber-Reviews-E2940227.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Blackpoint-Cyber-Jobs-E2940227.htm",
				Jobs:        "11",
				Reviews:     "56",
				Salaries:    "83",
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
							Title:                "Software Engineer Go",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4294835755/",
							Location:             "Canada",
							Date:                 mustDate("2025-09-04"),
							WithSalary:           true, // $150k/yr - $165k/yr
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
				domain.Erlang: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
			},
			ShortDescription: "",
			Industries: []domain.Industry{
				domain.IndustryCyberSecurity,
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "RapidAI",
			Website: "https://www.rapidai.com/",
			Careers: "https://www.rapidai.com/careers",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                37506734,
				IDs:               nil,
				Alias:             "rapid-ai",
				Name:              "RapidAI",
				Followers:         "21K",
				Employees:         "51-200",
				AssociatedMembers: "256",
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
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-RapidAI-EI_IE3241437.11,18.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/RapidAI-Reviews-E3241437.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/RapidAI-Jobs-E3241437.htm",
				Jobs:        "10",
				Reviews:     "51",
				Salaries:    "102",
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
							Title:                "Senior Software Engineer – Golang",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4296963567/",
							Location:             "Bengaluru, Karnataka, India",
							Date:                 mustDate("2025-09-06"),
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
				domain.Erlang: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
			},
			ShortDescription: "",
			Industries: []domain.Industry{
				domain.IndustryHealthTech,
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Schmooze",
			Website: "https://schmooze.in/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                72037888,
				IDs:               nil,
				Alias:             "schmooze-dating",
				Name:              "Schmooze",
				Followers:         "24K",
				Employees:         "11-50",
				AssociatedMembers: "68",
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
							Title:                "Software Engineer – Backend (Python/Go)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4296026137/",
							Location:             "Bengaluru, Karnataka, India",
							Date:                 mustDate("2025-09-06"),
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
				domain.Erlang: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
			},
			ShortDescription: "The first ever meme based dating app",
			Industries:       []domain.Industry{
				// NOP
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "OpenText",
			Website: "https://www.opentext.com/",
			Careers: "https://careers.opentext.com/",
			About:   "https://www.opentext.com/about",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                2709,
				IDs:               []int{2709, 9242, 21240},
				Alias:             "opentext",
				Name:              "OpenText",
				Followers:         "497K",
				Employees:         "10K+",
				AssociatedMembers: "22,746",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:     "opentext",
				Followers: "95",
				Verified:  true,
			},
			BlindProfile: domain.BlindProfile{
				Alias: "",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias: "",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-OpenText-EI_IE5694.11,19.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/OpenText-Reviews-E5694.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/OpenText-Jobs-E5694.htm",
				Jobs:        "515",
				Reviews:     "6K",
				Salaries:    "9.6K",
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
							Title:                "C++/Go Software Engineer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4245104258/",
							Location:             "Cluj-Napoca, Cluj, Romania",
							Date:                 mustDate("2025-11-19", "2025-10-29", "2025-10-08", "2025-09-16", "2025-09-06"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Golang Senior Engineer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4297794298/",
							Location:             "Bengaluru, Karnataka, India",
							Date:                 mustDate("2025-09-16"),
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
				domain.Erlang: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
			},
			ShortDescription: "",
			Industries: []domain.Industry{
				domain.IndustryCloudComputing,
			},
		},
	}
}
