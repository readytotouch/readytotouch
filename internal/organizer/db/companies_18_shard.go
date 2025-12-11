package db

import "github.com/readytotouch/readytotouch/internal/domain"

func companies18Shard() []domain.CompanyProfile {
	return []domain.CompanyProfile{
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Trainline",
			Website: "https://www.thetrainline.com/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                96955,
				IDs:               nil,
				Alias:             "trainline",
				Name:              "Trainline",
				Followers:         "51K",
				Employees:         "1K-5K",
				AssociatedMembers: "1,105",
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
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Trainline-EI_IE249203.11,20.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Trainline-Reviews-E249203.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Trainline-Jobs-E249203.htm",
				Jobs:        "19",
				Reviews:     "443",
				Salaries:    "1K",
				ReviewsRate: "3.8",
				Verified:    true,
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
							Title:                "Data Engineer – Scala",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4245930632/",
							Location:             "London, England, United Kingdom",
							Date:                 mustDate("2025-07-18", "2025-06-28", "2025-06-07"),
							WithSalary:           false,
							Remote:               false,
						},
					},
				},
				domain.Elixir:  {},
				domain.Clojure: {},
				domain.Haskell: {},
			},
			ShortDescription: "Rail and coach travel platform",
			Industries:       []domain.Industry{
				// NOP
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Samba.ai",
			Website: "https://www.samba.ai/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                19090027,
				IDs:               nil,
				Alias:             "samba-ai",
				Name:              "Samba.ai",
				Followers:         "991",
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
				domain.Go:   {},
				domain.Rust: {},
				domain.Zig:  {},
				domain.Scala: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Functional developer (Scala)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4244484254/",
							Date:                 mustDate("2025-06-05"),
							WithSalary:           false,
							Remote:               false,
						},
					},
				},
				domain.Elixir:  {},
				domain.Clojure: {},
				domain.Haskell: {},
			},
			ShortDescription: "Our initial focus of providing data analysis for businesses culminated in the creation of a predictive model that would help businesses meet customer needs",
			Industries:       []domain.Industry{
				// NOP
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Nordiska",
			Website: "https://www.nordiska.se/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                18323957,
				IDs:               nil,
				Alias:             "bankaktiebolaget-nordiska",
				Name:              "Nordiska",
				Followers:         "2K",
				Employees:         "51-200",
				AssociatedMembers: "103",
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
							URL:                  "https://www.linkedin.com/jobs/view/4244496625/",
							Date:                 mustDate("2025-06-05"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Scala Developer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4260902083/",
							Date:                 mustDate("2025-07-04"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Scala Developer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4304080418/",
							Location:             "Stockholm, Stockholm County, Sweden",
							Date:                 mustDate("2025-09-23"),
							WithSalary:           false,
							Remote:               false,
						},
					},
				},
				domain.Elixir:  {},
				domain.Clojure: {},
				domain.Haskell: {},
			},
			ShortDescription: "Our goal is to be the leader in embedded finance in Europe",
			Industries: []domain.Industry{
				domain.IndustryFinTech,
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Florida Blue",
			Website: "http://www.floridablue.com/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                239879,
				IDs:               nil,
				Alias:             "florida-blue",
				Name:              "Florida Blue",
				Followers:         "66K",
				Employees:         "5K-10K",
				AssociatedMembers: "8,396",
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
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Florida-Blue-EI_IE17308.11,23.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Florida-Blue-Reviews-E17308.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Florida-Blue-Jobs-E17308.htm",
				Jobs:        "",
				Reviews:     "1.5K",
				Salaries:    "2.4K",
				ReviewsRate: "3.7",
				Verified:    false,
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
							Title:                "Senior Spark/Scala Developer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4244952523/",
							Date:                 mustDate("2025-06-09", "2025-06-06"),
							WithSalary:           true,
							Remote:               true,
						},
						{
							Title:                "Senior Spark/Scala Developer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4255337607/",
							Date:                 mustDate("2025-06-24"),
							WithSalary:           false,
							Remote:               false,
						},
					},
				},
				domain.Elixir:  {},
				domain.Clojure: {},
				domain.Haskell: {},
			},
			ShortDescription: "Florida Blue is a subsidiary of a not-for-profit health solutions company dedicated to serving all Floridians in the pursuit of health",
			Industries: []domain.Industry{
				domain.IndustryInsurTech,
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "GreyOrange",
			Website: "https://www.greyorange.com/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                3195463,
				IDs:               nil,
				Alias:             "gogreyorange",
				Name:              "GreyOrange",
				Followers:         "205K",
				Employees:         "501-1K",
				AssociatedMembers: "971",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:     "greyorange",
				Followers: "80",
				Verified:  true,
			},
			BlindProfile: domain.BlindProfile{
				Alias: "",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias: "",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-GreyOrange-EI_IE1034918.11,21.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/GreyOrange-Reviews-E1034918.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/GreyOrange-Jobs-E1034918.htm",
				Jobs:        "12",
				Reviews:     "611",
				Salaries:    "666",
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
							Title:                "Staff Software Engineer",
							ShortDescription:     "Preferably Java/Golang",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4210733413/",
							Date:                 mustDate("2025-06-08"),
							WithSalary:           false,
							Remote:               false,
						},
					},
				},
				domain.Rust:  {},
				domain.Zig:   {},
				domain.Scala: {},
				domain.Elixir: {
					GitHubRepositoriesCount: 1,
				},
				domain.Clojure: {},
				domain.Haskell: {},
				domain.Erlang: {
					GitHubRepositoriesCount: 29,
				},
			},
			ShortDescription: "Warehouse orchestration and store inventory management software",
			Industries:       []domain.Industry{
				// NOP
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "HealthEquity",
			Website: "https://www.healthequity.com/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                26087,
				IDs:               nil,
				Alias:             "healthequity",
				Name:              "HealthEquity",
				Followers:         "43K",
				Employees:         "1K-5K",
				AssociatedMembers: "2,299",
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
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-HealthEquity-Inc-EI_IE199470.11,27.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/HealthEquity-Inc-Reviews-E199470.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/HealthEquity-Inc-Jobs-E199470.htm",
				Jobs:        "24",
				Reviews:     "916",
				Salaries:    "1.2K",
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
							Title:                "Golang Software Engineer",
							ShortDescription:     "Cybersecurity",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4229352263/",
							Date:                 mustDate("2025-06-05"),
							WithSalary:           true,
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
			ShortDescription: "HealthEquity is a administrator of Health Savings Accounts (HSAs) and other consumer-directed benefits—FSA, HRA, COBRA, and Commuter",
			Industries: []domain.Industry{
				domain.IndustryFinTech,
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "foodpanda",
			Website: "https://www.foodpanda.com/",
			Careers: "https://careers.foodpanda.com/",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                2569398,
				IDs:               nil,
				Alias:             "foodpanda",
				Name:              "foodpanda",
				Followers:         "472K",
				Employees:         "5K-10K",
				AssociatedMembers: "15,612",
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
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-foodpanda-EI_IE709546.11,20.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/foodpanda-Reviews-E709546.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/foodpanda-Jobs-E709546.htm",
				Jobs:        "42",
				Reviews:     "2.7K",
				Salaries:    "3.1K",
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
							Title:                "Backend Engineer (Golang)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4242927462/",
							Date:                 mustDate("2025-06-06"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Senior Backend Engineer (Golang)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4282454513/",
							Location:             "Singapore, Singapore",
							Date:                 mustDate("2025-08-11"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Senior Backend Software Engineer (Golang)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4289190767/",
							Location:             "Singapore, Singapore",
							Date:                 mustDate("2025-11-15", "2025-10-03", "2025-09-12", "2025-08-21"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Backend Software Engineer (Golang)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4314019421/",
							Location:             "Singapore, Singapore",
							Date:                 mustDate("2025-11-26"),
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
			ShortDescription: "Delivery platform",
			Industries:       []domain.Industry{
				// NOP
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Opendoor",
			Website: "https://www.opendoor.com/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                9398436,
				IDs:               nil,
				Alias:             "opendoor-com",
				Name:              "Opendoor",
				Followers:         "164K",
				Employees:         "1K-5K",
				AssociatedMembers: "1,549",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:     "opendoor-labs",
				Followers: "73",
				Verified:  false,
			},
			BlindProfile: domain.BlindProfile{
				Alias:       "OpenDoor",
				Employees:   "501 to 1,000",
				Salary:      "$77K ~ $300K a year",
				Reviews:     "208",
				ReviewsRate: "3.2",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "opendoor",
				Employees: "1,500",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Opendoor-EI_IE1021515.11,19.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Opendoor-Reviews-E1021515.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Opendoor-Jobs-E1021515.htm",
				Jobs:        "27",
				Reviews:     "984",
				Salaries:    "1.7K",
				ReviewsRate: "3.3",
				Verified:    true,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "Opendoor-2",
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 40,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Software Engineer – Fullstack",
							ShortDescription:     "TypesSript, React, Golang, SQL",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4215964620/",
							Date:                 mustDate("2025-06-07"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Software Engineer – Fullstack (Go, SQL, TypeScript)",
							ShortDescription:     "Pricing",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4250001632/",
							Location:             "United States",
							Date:                 mustDate("2025-07-26", "2025-07-04"),
							WithSalary:           true, // $143.4k/yr - $196.9k/yr
							Remote:               true,
						},
						{
							Title:                "Software Engineer – Fullstack (Go, SQL, TypeScript)",
							ShortDescription:     "Pricing",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4308121465/",
							Location:             "Seattle, WA",
							Date:                 mustDate("2025-12-02", "2025-11-12"),
							WithSalary:           true, // $143.4k/yr - $196.9k/yr
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
				domain.IndustryPropTech,
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "SuperAnnotate",
			Website: "https://www.superannotate.com/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                18999422,
				IDs:               nil,
				Alias:             "superannotate",
				Name:              "SuperAnnotate",
				Followers:         "30K",
				Employees:         "51-200",
				AssociatedMembers: "307",
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
				Alias:     "superannotate",
				Employees: "150",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-SuperAnnotate-AI-EI_IE5649190.11,27.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/SuperAnnotate-AI-Reviews-E5649190.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/SuperAnnotate-AI-Jobs-E5649190.htm",
				Jobs:        "",
				Reviews:     "56",
				Salaries:    "60",
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
							Title:                "Senior Rust Developer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4145729615/",
							Date:                 mustDate("2025-02-10"),
							WithSalary:           true,
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
			ShortDescription: "Platform for building, fine-tuning, iterating, and managing your AI models faster with the highest-quality training data",
			Industries:       []domain.Industry{
				// NOP
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Snowplow",
			Website: "https://snowplow.io/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                3177331,
				IDs:               nil,
				Alias:             "snowplow",
				Name:              "Snowplow",
				Followers:         "15K",
				Employees:         "51-200",
				AssociatedMembers: "162",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:     "snowplow",
				Followers: "239",
				Verified:  true,
			},
			BlindProfile: domain.BlindProfile{
				Alias: "",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias: "",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Snowplow-EI_IE1482594.11,19.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Snowplow-Reviews-E1482594.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Snowplow-Jobs-E1482594.htm",
				Jobs:        "5",
				Reviews:     "63",
				Salaries:    "125",
				ReviewsRate: "3.4",
				Verified:    true,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "",
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 8,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Senior Software Engineer",
							ShortDescription:     "We are looking for a Senior Software Engineer with experience in Go development to join our Data Processing team",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4189936616/",
							Date:                 mustDate("2025-04-10"),
							WithSalary:           false,
							Remote:               true,
						},
					},
				},
				domain.Rust: {
					GitHubRepositoriesCount: 0,
				},
				domain.Zig: {
					GitHubRepositoriesCount: 0,
				},
				domain.Scala: {
					GitHubRepositoriesCount: 17,
				},
				domain.Elixir: {
					GitHubRepositoriesCount: 0,
				},
				domain.Clojure: {
					GitHubRepositoriesCount: 0,
				},
				domain.Haskell: {
					GitHubRepositoriesCount: 0,
				},
			},
			ShortDescription: "Customer data infrastructure for AI, enabling every organization to transform raw behavioral data into governed, high-fidelity fuel for AI-powered applications—including advanced analytics, real-time personalization engines, and AI agent",
			Industries: []domain.Industry{
				domain.IndustryDevOps,
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Monumental",
			Website: "https://www.monumental.co/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                74482238,
				IDs:               nil,
				Alias:             "monumentalco",
				Name:              "Monumental",
				Followers:         "6K",
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
				},
				domain.Rust: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Software Engineer, Full Stack",
							ShortDescription:     "Writing Rust and TypeScript code that controls the robot in real-time",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4245525591/",
							Date:                 mustDate("2025-06-06"),
							WithSalary:           false,
							Remote:               false,
						},
					},
				},
				domain.Zig: {
					GitHubRepositoriesCount: 0,
				},
				domain.Scala: {
					GitHubRepositoriesCount: 0,
				},
				domain.Elixir: {
					GitHubRepositoriesCount: 0,
				},
				domain.Clojure: {
					GitHubRepositoriesCount: 0,
				},
				domain.Haskell: {
					GitHubRepositoriesCount: 0,
				},
			},
			ShortDescription: "Automating on-site construction with robotics and software",
			Industries:       []domain.Industry{
				// NOP
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "SIDES",
			Website: "https://www.get-sides.de/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                10433986,
				IDs:               nil,
				Alias:             "sides-dach",
				Name:              "SIDES",
				Followers:         "3K",
				Employees:         "51-200",
				AssociatedMembers: "89",
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
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-SIDES-EI_IE4904216.11,16.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/SIDES-Reviews-E4904216.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/SIDES-Jobs-E4904216.htm",
				Jobs:        "12",
				Reviews:     "13",
				Salaries:    "34",
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
							Title:                "Golang Developer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4226876078/",
							Date:                 mustDate("2025-05-20"),
							WithSalary:           false,
							Remote:               false,
						},
					},
				},
				domain.Rust: {
					// NOP
				},
				domain.Zig: {
					// NOP
				},
				domain.Scala: {
					// NOP
				},
				domain.Elixir: {
					// NOP
				},
				domain.Clojure: {
					// NOP
				},
				domain.Haskell: {
					// NOP
				},
			},
			ShortDescription: "We have developed a cloud-based software package that is precisely tailored to the requirements of gastronomic delivery services",
			Industries:       []domain.Industry{
				// NOP
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Insider.",
			Website: "https://useinsider.com/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                2560882,
				IDs:               []int{2560882, 41621219},
				Alias:             "useinsider",
				Name:              "Insider.",
				Followers:         "127K",
				Employees:         "1K-5K",
				AssociatedMembers: "1,425",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:     "useinsider",
				Followers: "209",
				Verified:  false,
			},
			BlindProfile: domain.BlindProfile{
				Alias: "",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias: "",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Insider-EI_IE1673833.11,18.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Insider-Reviews-E1673833.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Insider-Jobs-E1673833.htm",
				Jobs:        "",
				Reviews:     "1.5K",
				Salaries:    "1K",
				ReviewsRate: "4.4",
				Verified:    true,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "",
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 3,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Senior Software Engineer (Golang)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4230962545/",
							Date:                 mustDate("2025-05-20"),
							WithSalary:           false,
							Remote:               true,
						},
						{
							Title:                "Senior Software Engineer (Golang)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4276101442/",
							Location:             "Istanbul, Istanbul, Türkiye",
							Date:                 mustDate("2025-07-27"),
							WithSalary:           false,
							Remote:               true,
						},
						{
							Title:                "Senior Software Engineer – Golang",
							ShortDescription:     "Email Team",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4278688430/",
							Location:             "Istanbul, Istanbul, Türkiye",
							Date:                 mustDate("2025-08-02"),
							WithSalary:           false,
							Remote:               true,
						},
						{
							Title:                "Senior Software Engineer – Golang",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4309049827/",
							Location:             "Istanbul, Istanbul, Türkiye",
							Date:                 mustDate("2025-10-02"),
							WithSalary:           false,
							Remote:               true,
						},
						{
							Title:                "Software Engineer (Golang)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4336992593/",
							Location:             "Hamburg, Hamburg, Germany",
							Date:                 mustDate("2025-11-19"),
							WithSalary:           false,
							Remote:               true,
						},
					},
				},
				domain.Rust: {
					// NOP
				},
				domain.Zig: {
					// NOP
				},
				domain.Scala: {
					// NOP
				},
				domain.Elixir: {
					// NOP
				},
				domain.Clojure: {
					// NOP
				},
				domain.Haskell: {
					// NOP
				},
			},
			ShortDescription: "Platform for individualized, cross-channel experiences—enables enterprise marketers to connect customer data across channels and systems, predict their future behavior with an AI intent engine and individualize customer experiences",
			Industries:       []domain.Industry{
				// NOP
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Lucanet",
			Website: "https://www.lucanet.com/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                1691101,
				IDs:               nil,
				Alias:             "lucanet-ag",
				Name:              "Lucanet",
				Followers:         "49K",
				Employees:         "501-1K",
				AssociatedMembers: "848",
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
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Lucanet-EI_IE413584.11,18.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Lucanet-Reviews-E413584.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Lucanet-Jobs-E413584.htm",
				Jobs:        "27",
				Reviews:     "80",
				Salaries:    "221",
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
							Title:                "Senior Go Backend Engineer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4245965954/",
							Location:             "Romania",
							Date:                 mustDate("2025-09-23", "2025-09-02", "2025-06-07"),
							WithSalary:           false,
							Remote:               true,
						},
					},
				},
				domain.Rust: {
					// NOP
				},
				domain.Zig: {
					// NOP
				},
				domain.Scala: {
					// NOP
				},
				domain.Elixir: {
					// NOP
				},
				domain.Clojure: {
					// NOP
				},
				domain.Haskell: {
					// NOP
				},
			},
			ShortDescription: "Corporate Performance Management",
			Industries:       []domain.Industry{
				// NOP
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Uniphore",
			Website: "https://www.uniphore.com/",
			Careers: "https://www.uniphore.com/careers/",
			About:   "https://www.uniphore.com/about-us/",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                1331150,
				IDs:               nil,
				Alias:             "uniphore",
				Name:              "Uniphore",
				Followers:         "70K",
				Employees:         "501-1K",
				AssociatedMembers: "861",
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
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Uniphore-Software-Systems-EI_IE924913.11,36.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Uniphore-Software-Systems-Reviews-E924913.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Uniphore-Software-Systems-Jobs-E924913.htm",
				Jobs:        "39",
				Reviews:     "431",
				Salaries:    "486",
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
							Title:                "Senior Software Engineer (Java/Go+Looker)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4185030868/",
							Location:             "Bengaluru, Karnataka, India",
							Date:                 mustDate("2025-09-04", "2025-07-23", "2025-06-09"),
							WithSalary:           false,
							Remote:               false,
						},
					},
				},
				domain.Rust: {
					// NOP
				},
				domain.Zig: {
					// NOP
				},
				domain.Scala: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Senior Scala Engineer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4247302848/",
							Date:                 mustDate("2025-06-12"),
							WithSalary:           true,
							Remote:               false,
						},
					},
				},
				domain.Elixir: {
					// NOP
				},
				domain.Clojure: {
					// NOP
				},
				domain.Haskell: {
					// NOP
				},
			},
			ShortDescription: "Our sovereign, composable, and secure AI platform connects enterprise data, fine-tunes AI models and deploys agentic AI across the enterprise",
			Industries:       []domain.Industry{
				// NOP
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "The Home Depot",
			Website: "https://homedepot.com/",
			Careers: "https://careers.homedepot.com/",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                1534,
				IDs:               []int{1534, 164209},
				Alias:             "the-home-depot",
				Name:              "The Home Depot",
				Followers:         "961K",
				Employees:         "10K+",
				AssociatedMembers: "121,537",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:     "homedepot",
				Followers: "56",
				Verified:  false,
			},
			BlindProfile: domain.BlindProfile{
				Alias:       "The-Home-Depot",
				Employees:   "10,000+",
				Salary:      "$60K ~ $250K a year",
				Reviews:     "368",
				ReviewsRate: "3.7",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "the-home-depot",
				Employees: "500,000",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-The-Home-Depot-EI_IE655.11,25.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/The-Home-Depot-Reviews-E655.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/The-Home-Depot-Jobs-E655.htm",
				Jobs:        "18K",
				Reviews:     "56K",
				Salaries:    "107K",
				ReviewsRate: "3.8",
				Verified:    true,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "",
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 28,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Backend Software Engineer – Go",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4243288093/",
							Date:                 mustDate("2025-06-09"),
							WithSalary:           false,
							Remote:               true,
						},
						{
							Title:                "Software Engineer II – Golang",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4280054568/",
							Location:             "Atlanta, GA",
							Date:                 mustDate("2025-08-06"),
							WithSalary:           true, // $90k/yr - $170k/yr
							Remote:               true,
						},
					},
				},
				domain.Rust: {
					// NOP
				},
				domain.Zig: {
					// NOP
				},
				domain.Scala: {
					// NOP
				},
				domain.Elixir: {
					// NOP
				},
				domain.Clojure: {
					// NOP
				},
				domain.Haskell: {
					// NOP
				},
			},
			ShortDescription: "Home improvement specialty retailer",
			Industries:       []domain.Industry{
				// NOP
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Takt",
			Website: "https://www.takt.io/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                74506126,
				IDs:               nil,
				Alias:             "takt-io",
				Name:              "Takt",
				Followers:         "2K",
				Employees:         "11-50",
				AssociatedMembers: "17",
				Verified:          false,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:     "takt-corp",
				Followers: "5",
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
					GitHubRepositoriesCount: 6,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Golang Engineer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4247471344/",
							Date:                 mustDate("2025-06-09"),
							WithSalary:           false,
							Remote:               false,
						},
					},
				},
				domain.Rust: {
					// NOP
				},
				domain.Zig: {
					// NOP
				},
				domain.Scala: {
					// NOP
				},
				domain.Elixir: {
					// NOP
				},
				domain.Clojure: {
					// NOP
				},
				domain.Haskell: {
					// NOP
				},
			},
			ShortDescription: "Labor Management Systems",
			Industries:       []domain.Industry{
				// NOP
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Travelport",
			Website: "https://www.travelport.com/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                2124,
				IDs:               nil,
				Alias:             "travelport",
				Name:              "Travelport",
				Followers:         "219K",
				Employees:         "1K-5K",
				AssociatedMembers: "3,769",
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
				Alias:     "travelport",
				Employees: "4,250",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Travelport-EI_IE41925.11,21.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Travelport-Reviews-E41925.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Travelport-Jobs-E41925.htm",
				Jobs:        "18",
				Reviews:     "1.2K",
				Salaries:    "1.8K",
				ReviewsRate: "3.3",
				Verified:    true,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "",
			},
			Languages: domain.Languages{
				domain.Go: {
					// NOP
				},
				domain.Rust: {
					// NOP
				},
				domain.Zig: {
					// NOP
				},
				domain.Scala: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Software Development Engineer Scala",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4247551737/",
							Date:                 mustDate("2025-06-10"),
							WithSalary:           false,
							Remote:               false,
						},
					},
				},
				domain.Elixir: {
					// NOP
				},
				domain.Clojure: {
					// NOP
				},
				domain.Haskell: {
					// NOP
				},
			},
			ShortDescription: "Travel retail platform",
			Industries:       []domain.Industry{
				// NOP
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Marketer.com",
			Website: "https://www.marketer.com/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                91317726,
				IDs:               nil,
				Alias:             "marketer-com",
				Name:              "Marketer.com",
				Followers:         "811",
				Employees:         "11-50",
				AssociatedMembers: "32",
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
				},
				domain.Rust: {
					// NOP
				},
				domain.Zig: {
					// NOP
				},
				domain.Scala: {
					// NOP
				},
				domain.Elixir: {
					// NOP
				},
				domain.Clojure: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Senior Clojure Engineer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4249821311/",
							Date:                 mustDate("2025-06-13"),
							WithSalary:           false,
							Remote:               true,
						},
					},
				},
				domain.Haskell: {
					// NOP
				},
			},
			ShortDescription: "App for fully automated Shopify sales",
			Industries: []domain.Industry{
				domain.IndustryMarTech,
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "WPP Media",
			Website: "https://www.wppmedia.com/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                106681730,
				IDs:               []int{3764, 4703, 5935, 13602, 22070, 242315, 359382, 796114, 1519820, 1915522, 2527962, 2550970, 10528967, 11266049, 26058981, 72211530, 80676023, 91693237, 106681730},
				Alias:             "wpp-media",
				Name:              "WPP Media",
				Followers:         "765K",
				Employees:         "10K+",
				AssociatedMembers: "25,858",
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
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-WPP-Media-EI_IE318409.11,20.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/WPP-Media-Reviews-E318409.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/WPP-Media-Jobs-E318409.htm",
				Jobs:        "1",
				Reviews:     "3.5K",
				Salaries:    "6.4K",
				ReviewsRate: "3.6",
				Verified:    false,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "",
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
				},
				domain.Rust: {
					// NOP
				},
				domain.Zig: {
					// NOP
				},
				domain.Scala: {
					// NOP
				},
				domain.Elixir: {
					// NOP
				},
				domain.Clojure: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Senior Clojure Developer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4248410642/",
							Location:             "London, England, United Kingdom",
							Date:                 mustDate("2025-08-05", "2025-07-26", "2025-07-04", "2025-06-14"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Senior Clojure Developer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4296287334/",
							Location:             "London, England, United Kingdom",
							Date:                 mustDate("2025-09-20", "2025-09-10"),
							WithSalary:           false,
							Remote:               false,
						},
					},
				},
				domain.Haskell: {
					// NOP
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
			Name:    "Lonely Planet",
			Website: "https://www.lonelyplanet.com/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                164764,
				IDs:               nil,
				Alias:             "lonely-planet",
				Name:              "Lonely Planet",
				Followers:         "115K",
				Employees:         "51-200",
				AssociatedMembers: "783",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:     "lonelyplanet",
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
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Lonely-Planet-EI_IE5746.11,24.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Lonely-Planet-Reviews-E5746.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Lonely-Planet-Jobs-E5746.htm",
				Jobs:        "",
				Reviews:     "149",
				Salaries:    "219",
				ReviewsRate: "3.7",
				Verified:    false,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "",
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
				},
				domain.Rust: {
					// NOP
				},
				domain.Zig: {
					// NOP
				},
				domain.Scala: {
					// NOP
				},
				domain.Elixir: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Software Engineer – Elixir",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4248263842/",
							Date:                 mustDate("2025-06-11"),
							WithSalary:           true,
							Remote:               true,
						},
						{
							Title:                "Elixir Software Engineer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4269516629/",
							Location:             "United States",
							Date:                 mustDate("2025-07-18"),
							WithSalary:           true, // $80k/yr - $120k/yr
							Remote:               true,
						},
					},
				},
				domain.Clojure: {
					// NOP
				},
				domain.Haskell: {
					// NOP
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
			Name:    "Event Inc Group",
			Website: "https://www.eventinc.de/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                3646433,
				IDs:               nil,
				Alias:             "event-inc",
				Name:              "Event Inc Group",
				Followers:         "2K",
				Employees:         "51-200",
				AssociatedMembers: "77",
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
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Event-EI_IE1343788.11,16.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Event-Reviews-E1343788.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Event-Jobs-E1343788.htm",
				Jobs:        "20",
				Reviews:     "20",
				Salaries:    "19",
				ReviewsRate: "4.5",
				Verified:    false,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "",
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
				},
				domain.Rust: {
					// NOP
				},
				domain.Zig: {
					// NOP
				},
				domain.Scala: {
					// NOP
				},
				domain.Elixir: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Software Engineer – Elixir",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4247993545/",
							Date:                 mustDate("2025-06-14"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Senior DevOps Engineer – Elixir",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4294469302/",
							Location:             "Hamburg, Hamburg, Germany",
							Date:                 mustDate("2025-09-06"),
							WithSalary:           false,
							Remote:               false,
						},
					},
				},
				domain.Clojure: {
					// NOP
				},
				domain.Haskell: {
					// NOP
				},
			},
			ShortDescription: "Online marketplace for conference & event spaces",
			Industries:       []domain.Industry{
				// NOP
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Box",
			Website: "https://www.box.com/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                54178,
				IDs:               nil,
				Alias:             "box",
				Name:              "Box",
				Followers:         "179K",
				Employees:         "1K-5K",
				AssociatedMembers: "3,980",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:     "box",
				Followers: "167",
				Verified:  true,
			},
			BlindProfile: domain.BlindProfile{
				Alias:       "Box",
				Employees:   "1,001 to 5,000",
				Salary:      "$73K ~ $333K a year",
				Reviews:     "414",
				ReviewsRate: "3.9",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "box",
				Employees: "3,500",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Box-EI_IE254092.11,14.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Box-Reviews-E254092.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Box-Jobs-E254092.htm",
				Jobs:        "206",
				Reviews:     "1.3K",
				Salaries:    "3.3K",
				ReviewsRate: "4.3",
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
							Title:                "Software Engineer III, SRE (Python/Go)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4255716681/",
							Date:                 mustDate("2025-06-24"),
							WithSalary:           false,
							Remote:               false,
						},
					},
				},
				domain.Rust: {
					// NOP
				},
				domain.Zig: {
					// NOP
				},
				domain.Scala: {
					GitHubRepositoriesCount: 1,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Senior Scala Engineer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4215392486/",
							Date:                 mustDate("2025-06-12"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Senior Scala Engineer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4252590022/",
							Date:                 mustDate("2025-07-09", "2025-06-18"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Middle Scala Engineer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4252582963/",
							Location:             "Warsaw, Mazowieckie, Poland",
							Date:                 mustDate("2025-07-30"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Senior Scala Engineer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4277500586/",
							Location:             "Warsaw, Mazowieckie, Poland",
							Date:                 mustDate("2025-08-22"),
							WithSalary:           false,
							Remote:               false,
						},
					},
				},
				domain.Elixir: {
					// NOP
				},
				domain.Clojure: {
					// NOP
				},
				domain.Haskell: {
					// NOP
				},
			},
			ShortDescription: "Intelligent Content Cloud, a single platform that enables organizations to fuel collaboration, manage the entire content lifecycle, secure critical content, and transform business workflows with enterprise AI",
			Industries:       []domain.Industry{
				// NOP
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Fortanix",
			Website: "https://www.fortanix.com/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                12417189,
				IDs:               nil,
				Alias:             "fortanix",
				Name:              "Fortanix",
				Followers:         "145K",
				Employees:         "201-500",
				AssociatedMembers: "228",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:     "fortanix",
				Followers: "45",
				Verified:  false,
			},
			BlindProfile: domain.BlindProfile{
				Alias: "",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "fortanix",
				Employees: "150",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Fortanix-EI_IE1977675.11,19.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Fortanix-Reviews-E1977675.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Fortanix-Jobs-E1977675.htm",
				Jobs:        "13",
				Reviews:     "155",
				Salaries:    "226",
				ReviewsRate: "2.8",
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
							Title:                "Staff Software Engineer – Rust, C/C++, Golang",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4250628040/",
							Date:                 mustDate("2025-06-14"),
							WithSalary:           true,
							Remote:               false,
						},
					},
				},
				domain.Rust: {
					GitHubRepositoriesCount: 50,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Staff Software Engineer – Rust, C/C++, Golang",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4250628040/",
							Date:                 mustDate("2025-06-14"),
							WithSalary:           true,
							Remote:               false,
						},
						{
							Title:                "Software Engineer III – Rust",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4317329718/",
							Location:             "Bengaluru, Karnataka, India",
							Date:                 mustDate("2025-10-21"),
							WithSalary:           false,
							Remote:               false,
						},
					},
				},
				domain.Zig: {
					// NOP
				},
				domain.Scala: {
					// NOP
				},
				domain.Elixir: {
					// NOP
				},
				domain.Clojure: {
					// NOP
				},
				domain.Haskell: {
					// NOP
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
			Name:    "Take-Two Interactive",
			Website: "https://www.take2games.com/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                10205,
				IDs:               []int{10205, 5037214},
				Alias:             "take-2-interactive-software-inc-",
				Name:              "Take-Two Interactive",
				Followers:         "91K",
				Employees:         "10K+",
				AssociatedMembers: "1,146",
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
				Alias:     "take-two-interactive-software",
				Employees: "760",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Take-Two-EI_IE7413.11,19.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Take-Two-Reviews-E7413.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Take-Two-Jobs-E7413.htm",
				Jobs:        "122",
				Reviews:     "195",
				Salaries:    "519",
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
							Title:                "Senior Go Engineer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4203042089/",
							Date:                 mustDate("2025-06-12"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Senior Go Engineer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4203036962/",
							Date:                 mustDate("2025-07-04"),
							WithSalary:           false,
							Remote:               false,
						},
					},
				},
				domain.Rust: {
					// NOP
				},
				domain.Zig: {
					// NOP
				},
				domain.Scala: {
					// NOP
				},
				domain.Elixir: {
					// NOP
				},
				domain.Clojure: {
					// NOP
				},
				domain.Haskell: {
					// NOP
				},
			},
			ShortDescription: "Developer, publisher, and marketer of interactive entertainment for consumers",
			Industries: []domain.Industry{
				domain.IndustryGameDev,
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Bullish",
			Website: "https://bullish.com/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                75109083,
				IDs:               nil,
				Alias:             "bebullish",
				Name:              "Bullish",
				Followers:         "22K",
				Employees:         "201-500",
				AssociatedMembers: "303",
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
							Title:                "Senior/Lead Golang Backend Engineer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4247818350/",
							Date:                 mustDate("2025-06-11"),
							WithSalary:           false,
							Remote:               false,
						},
					},
				},
				domain.Rust: {
					// NOP
				},
				domain.Zig: {
					// NOP
				},
				domain.Scala: {
					// NOP
				},
				domain.Elixir: {
					// NOP
				},
				domain.Clojure: {
					// NOP
				},
				domain.Haskell: {
					// NOP
				},
			},
			ShortDescription: "",
			Industries:       []domain.Industry{
				// NOP
			},
			Ignore: true, // Crypto
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Pizza Hut",
			Website: "https://www.pizzahut.com/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                3406,
				IDs:               nil,
				Alias:             "pizza-hut",
				Name:              "Pizza Hut",
				Followers:         "346K",
				Employees:         "10K+",
				AssociatedMembers: "82,738",
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
				Alias:     "pizza-hut",
				Employees: "54,201",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Pizza-Hut-EI_IE10090.11,20.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Pizza-Hut-Reviews-E10090.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Pizza-Hut-Jobs-E10090.htm",
				Jobs:        "23K",
				Reviews:     "15K",
				Salaries:    "16K",
				ReviewsRate: "3.5",
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
							Title:                "Staff Software Engineer – Full Stack Engineer (Golang)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4219890382/",
							Date:                 mustDate("2025-06-14"),
							WithSalary:           true,
							Remote:               false,
						},
					},
				},
				domain.Rust: {
					// NOP
				},
				domain.Zig: {
					// NOP
				},
				domain.Scala: {
					// NOP
				},
				domain.Elixir: {
					// NOP
				},
				domain.Clojure: {
					// NOP
				},
				domain.Haskell: {
					// NOP
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
			Name:    "Further",
			Website: "https://www.gofurther.com/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                238000,
				IDs:               nil,
				Alias:             "further-worldwide",
				Name:              "Further",
				Followers:         "6K",
				Employees:         "201-500",
				AssociatedMembers: "168",
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
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Further-EI_IE1316197.11,18.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Further-Reviews-E1316197.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Further-Jobs-E1316197.htm",
				Jobs:        "",
				Reviews:     "72",
				Salaries:    "183",
				ReviewsRate: "3.1",
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
							Title:                "Full Stack Golang + Typescript Application Developer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4219327413/",
							Date:                 mustDate("2025-06-10"),
							WithSalary:           false,
							Remote:               false,
						},
					},
				},
				domain.Rust: {
					// NOP
				},
				domain.Zig: {
					// NOP
				},
				domain.Scala: {
					// NOP
				},
				domain.Elixir: {
					// NOP
				},
				domain.Clojure: {
					// NOP
				},
				domain.Haskell: {
					// NOP
				},
			},
			ShortDescription: "Further is a data, cloud, and AI company whose sole focus is helping you turn raw data into the right decisions",
			Industries:       []domain.Industry{
				// NOP
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Titan OS",
			Website: "https://www.titanos.tv/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                91587657,
				IDs:               nil,
				Alias:             "titanos",
				Name:              "Titan OS",
				Followers:         "10K",
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
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Golang Developer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4249271541/",
							Date:                 mustDate("2025-06-12"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Golang Developer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4260858136/",
							Date:                 mustDate("2025-07-02"),
							WithSalary:           false,
							Remote:               false,
						},
					},
				},
				domain.Rust: {
					// NOP
				},
				domain.Zig: {
					// NOP
				},
				domain.Scala: {
					// NOP
				},
				domain.Elixir: {
					// NOP
				},
				domain.Clojure: {
					// NOP
				},
				domain.Haskell: {
					// NOP
				},
			},
			ShortDescription: "The company specialises in developing software and solutions to unleash the full potential of Connected TV",
			Industries:       []domain.Industry{
				// NOP
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Akka",
			Website: "https://akka.io/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                7789814,
				IDs:               nil,
				Alias:             "akka-io",
				Name:              "Akka",
				Followers:         "16K",
				Employees:         "51-200",
				AssociatedMembers: "93",
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
							Title:                "Senior Software Engineer Agentic systems, Go, Scala, Kubernetes",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4252856747/",
							Date:                 mustDate("2025-06-18"),
							WithSalary:           false,
							Remote:               true,
						},
					},
				},
				domain.Rust: {
					// NOP
				},
				domain.Zig: {
					// NOP
				},
				domain.Scala: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Senior Software Engineer Agentic systems, Go, Scala, Kubernetes",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4252856747/",
							Date:                 mustDate("2025-06-18"),
							WithSalary:           false,
							Remote:               true,
						},
					},
				},
				domain.Elixir: {
					// NOP
				},
				domain.Clojure: {
					// NOP
				},
				domain.Haskell: {
					// NOP
				},
			},
			ShortDescription: "Akka is a platform to build and run data, API, and agentic AI services",
			Industries: []domain.Industry{
				domain.IndustryDevOps,
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "LivePerson",
			Website: "https://www.liveperson.com/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                164748,
				IDs:               nil,
				Alias:             "liveperson",
				Name:              "LivePerson",
				Followers:         "333K",
				Employees:         "501-1K",
				AssociatedMembers: "1,407",
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
				Alias:     "liveperson",
				Employees: "2,000",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-LivePerson-EI_IE11463.11,21.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/LivePerson-Reviews-E11463.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/LivePerson-Jobs-E11463.htm",
				Jobs:        "25",
				Reviews:     "944",
				Salaries:    "1.7K",
				ReviewsRate: "2.7",
				Verified:    false,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "",
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
				},
				domain.Rust: {
					// NOP
				},
				domain.Zig: {
					// NOP
				},
				domain.Scala: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Backend Engineer (Java/Scala)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4250367265/",
							Date:                 mustDate("2025-06-18"),
							WithSalary:           false,
							Remote:               true,
						},
					},
				},
				domain.Elixir: {
					// NOP
				},
				domain.Clojure: {
					// NOP
				},
				domain.Haskell: {
					// NOP
				},
			},
			ShortDescription: "Conversational Cloud platform",
			Industries:       []domain.Industry{
				// NOP
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "bunny.net",
			Website: "https://bunny.net/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                10484043,
				IDs:               nil,
				Alias:             "bunnynet",
				Name:              "bunny.net",
				Followers:         "6K",
				Employees:         "51-200",
				AssociatedMembers: "86",
				Verified:          false,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:     "BunnyWay",
				Followers: "76",
				Verified:  false,
			},
			BlindProfile: domain.BlindProfile{
				Alias: "",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias: "",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-bunny-net-EI_IE5045467.11,20.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/bunny-net-Reviews-E5045467.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/bunny-net-Jobs-E5045467.htm",
				Jobs:        "",
				Reviews:     "15",
				Salaries:    "15",
				ReviewsRate: "4.5",
				Verified:    false,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "",
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 3,
				},
				domain.Rust: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Senior Software Engineer (Rust)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4248785636/",
							Date:                 mustDate("2025-06-14"),
							WithSalary:           false,
							Remote:               true,
						},
						{
							Title:                "Senior Software Engineer (Rust)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4304078444/",
							Location:             "Ljubljana, Ljubljana, Slovenia",
							Date:                 mustDate("2025-09-23"),
							WithSalary:           false,
							Remote:               true,
						},
					},
				},
				domain.Zig: {
					// NOP
				},
				domain.Scala: {
					// NOP
				},
				domain.Elixir: {
					// NOP
				},
				domain.Clojure: {
					// NOP
				},
				domain.Haskell: {
					// NOP
				},
			},
			ShortDescription: "bunny.net is the content delivery platform",
			Industries: []domain.Industry{
				domain.IndustryDevOps,
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Lynk",
			Website: "https://lynk.world/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                11024220,
				IDs:               nil,
				Alias:             "lynk-global-inc",
				Name:              "Lynk",
				Followers:         "10K",
				Employees:         "11-50",
				AssociatedMembers: "70",
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
				},
				domain.Rust: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Senior Satellite Flight Software Engineer (Rust)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4253168651/",
							Date:                 mustDate("2025-06-18"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Full Stack Developer (Rust)",
							ShortDescription:     "CI/CD & Test Environments",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4274985946/",
							Location:             "Headquarters, KY",
							Date:                 mustDate("2025-07-23"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Senior Satellite Flight Software Engineer (Rust)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4281438944/",
							Location:             "Chantilly, VA",
							Date:                 mustDate("2025-09-02"),
							WithSalary:           false,
							Remote:               false,
						},
					},
				},
				domain.Zig: {
					// NOP
				},
				domain.Scala: {
					// NOP
				},
				domain.Elixir: {
					// NOP
				},
				domain.Clojure: {
					// NOP
				},
				domain.Haskell: {
					// NOP
				},
			},
			ShortDescription: "Lynk allows commercial subscribers to send and receive text messages to and from space",
			Industries:       []domain.Industry{
				// NOP
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "audibene",
			Website: "https://www.audibene.de/",
			Careers: "https://karriere.audibene.de/",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                3248758,
				IDs:               nil,
				Alias:             "audibene",
				Name:              "audibene",
				Followers:         "26K",
				Employees:         "1K-5K",
				AssociatedMembers: "456",
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
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-audibene-EI_IE1155737.11,19.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/audibene-Reviews-E1155737.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/audibene-Jobs-E1155737.htm",
				Jobs:        "30",
				Reviews:     "215",
				Salaries:    "282",
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
					Vacancies: []domain.Vacancy{
						{
							Title:                "Senior Rust Developer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4253509850/",
							Date:                 mustDate("2025-07-12", "2025-06-20"),
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
			ShortDescription: "Hearing care company",
			Industries: []domain.Industry{
				domain.IndustryHealthTech,
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Toast",
			Website: "https://pos.toasttab.com/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                5102948,
				IDs:               nil,
				Alias:             "toast-inc",
				Name:              "Toast",
				Followers:         "238K",
				Employees:         "1K-5K",
				AssociatedMembers: "6,765",
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
							Title:                "Staff Software Engineer, Infrastructure Engineer (Golang)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4244205425/",
							Date:                 mustDate("2025-06-24"),
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
			ShortDescription: "Restaurant platform",
			Industries:       []domain.Industry{
				// NOP
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Rumble",
			Website: "https://rumble.com/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                3496158,
				IDs:               nil,
				Alias:             "rumble-com",
				Name:              "Rumble",
				Followers:         "20K",
				Employees:         "11-50",
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
							URL:                  "https://www.linkedin.com/jobs/view/4246502912/",
							Location:             "Toronto, ON",
							Date:                 mustDate("2025-08-09", "2025-07-11", "2025-07-04", "2025-06-24"),
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
			ShortDescription: "Video platform that is creating the rails and independent infrastructure designed to be immune to cancel culture",
			Industries:       []domain.Industry{
				// NOP
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "SOAX",
			Website: "https://soax.com/",
			Careers: "https://soax.com/careers",
			About:   "https://soax.com/about",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                34924582,
				IDs:               nil,
				Alias:             "soax-network",
				Name:              "SOAX",
				Followers:         "6K",
				Employees:         "51-200",
				AssociatedMembers: "54",
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
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-SOAX-COM-EI_IE4888929.11,19.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/SOAX-COM-Reviews-E4888929.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/SOAX-COM-Jobs-E4888929.htm",
				Jobs:        "",
				Reviews:     "18",
				Salaries:    "25",
				ReviewsRate: "4.6",
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
							ShortDescription:     "Proxy",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4253254160/",
							Date:                 mustDate("2025-06-20"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Backend Engineer (Go & Python)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4260878493/",
							Date:                 mustDate("2025-07-02"),
							WithSalary:           false,
							Remote:               true,
						},
						{
							Title:                "Backend Engineer (Go & Python)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4260861412/",
							Date:                 mustDate("2025-07-02"),
							WithSalary:           false,
							Remote:               true,
						},
						{
							Title:                "Golang Software Engineer",
							ShortDescription:     "Proxy",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4266341426/",
							Location:             "Spain",
							Date:                 mustDate("2025-07-15"),
							WithSalary:           false,
							Remote:               true,
						},
						{
							Title:                "Backend Engineer (Go & Python)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4275330408/",
							Location:             "Spain",
							Date:                 mustDate("2025-07-24"),
							WithSalary:           false,
							Remote:               true,
						},
						{
							Title:                "Golang Software Engineer",
							ShortDescription:     "Proxy",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4318273888/",
							Location:             "Cyprus",
							Date:                 mustDate("2025-10-22"),
							WithSalary:           false,
							Remote:               true,
						},
						{
							Title:                "Golang Software Engineer (Proxy)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4335403626/",
							Location:             "Serbia",
							Date:                 mustDate("2025-11-12"),
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
			ShortDescription: "Data extraction platform",
			Industries:       []domain.Industry{
				// NOP
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Artificial Labs",
			Website: "https://artificial.io/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                5044394,
				IDs:               nil,
				Alias:             "artificiallabs",
				Name:              "Artificial Labs",
				Followers:         "6K",
				Employees:         "51-200",
				AssociatedMembers: "66",
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
							Title:                "Haskell Engineer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4258314586/",
							Date:                 mustDate("2025-07-01"),
							WithSalary:           false,
							Remote:               true,
						},
					},
				},
			},
			ShortDescription: "",
			Industries: []domain.Industry{
				domain.IndustryInsurTech,
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Quorum Software",
			Website: "https://www.quorumsoftware.com/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                17518,
				IDs:               []int{17518, 44347, 143840, 1616533, 2365867, 5588778},
				Alias:             "quorum-software",
				Name:              "Quorum Software",
				Followers:         "29K",
				Employees:         "1K-5K",
				AssociatedMembers: "1,680",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "",
				Verified: false,
			},
			BlindProfile: domain.BlindProfile{
				Alias:       "Quorum-Software",
				Employees:   "1,001 to 5,000",
				Salary:      "",
				Reviews:     "5",
				ReviewsRate: "3.6",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias: "",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Quorum-Software-EI_IE16174.11,26.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Quorum-Software-Reviews-E16174.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Quorum-Software-Jobs-E16174.htm",
				Jobs:        "121",
				Reviews:     "455",
				Salaries:    "911",
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
							Title:                "Lead Engineer – Node.js or Elixir",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4089011061/",
							Date:                 mustDate("2025-06-28"),
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
			ShortDescription: "Provider of energy software",
			Industries:       []domain.Industry{
				// NOP
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "State Street",
			Website: "https://www.statestreet.com/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                1767,
				IDs:               []int{1767, 1089754},
				Alias:             "state-street",
				Name:              "State Street",
				Followers:         "696K",
				Employees:         "10K+",
				AssociatedMembers: "44,571",
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
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-State-Street-EI_IE1911.11,23.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/State-Street-Reviews-E1911.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/State-Street-Jobs-E1911.htm",
				Jobs:        "693",
				Reviews:     "11K",
				Salaries:    "19K",
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
					Vacancies: []domain.Vacancy{
						{
							Title:                "Java Scala Developer — Senior Manager",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4259646441/",
							Date:                 mustDate("2025-07-01"),
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
			Name:    "RELX",
			Website: "https://www.relx.com/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                3022,
				IDs:               nil,
				Alias:             "relx-group",
				Name:              "RELX",
				Followers:         "93K",
				Employees:         "1K-5K",
				AssociatedMembers: "2,761",
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
							Title:                "Rust Senior Software Engineer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4256117557/",
							Date:                 mustDate("2025-06-26"),
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
			ShortDescription: "Provider of information-based analytics and decision tools for professional and business customers",
			Industries:       []domain.Industry{
				// NOP
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Castor",
			Website: "https://www.castoredc.com/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                3075405,
				IDs:               nil,
				Alias:             "castoredc",
				Name:              "Castor",
				Followers:         "23K",
				Employees:         "201-500",
				AssociatedMembers: "275",
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
							Title:                "Senior Software Engineer (Rust)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4256124146/",
							Date:                 mustDate("2025-06-26"),
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
			ShortDescription: "Modular clinical trial platform",
			Industries: []domain.Industry{
				domain.IndustryMedTech,
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Symbolica AI",
			Website: "https://www.symbolica.ai/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                82669556,
				IDs:               nil,
				Alias:             "symbolica-ai",
				Name:              "Symbolica AI",
				Followers:         "3K",
				Employees:         "11-50",
				AssociatedMembers: "18",
				Verified:          false,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:     "symbolica-ai",
				Followers: "201",
				Verified:  true,
			},
			BlindProfile: domain.BlindProfile{
				Alias: "",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias: "",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Symbolica-AI-EI_IE10326705.11,23.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Symbolica-AI-Reviews-E10326705.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Symbolica-AI-Jobs-E10326705.htm",
				Jobs:        "",
				Reviews:     "14",
				Salaries:    "13",
				ReviewsRate: "4.4",
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
					GitHubRepositoriesCount: 1,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Senior Software Engineer (Rust)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4255852068/",
							Date:                 mustDate("2025-06-28"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Senior Software Engineer (Rust)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4265460789/",
							Location:             "San Francisco, CA",
							Date:                 mustDate("2025-11-25", "2025-11-03", "2025-07-15"),
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
			ShortDescription: "AI research lab",
			Industries:       []domain.Industry{
				// NOP
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Fern",
			Website: "https://www.buildwithfern.com/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                88002225,
				IDs:               nil,
				Alias:             "buildwithfern",
				Name:              "Fern",
				Followers:         "5K",
				Employees:         "11-50",
				AssociatedMembers: "34",
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
							Title:                "Software Engineer, Go",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4261072865/",
							Date:                 mustDate("2025-07-03"),
							WithSalary:           true,
							Remote:               true,
						},
					},
				},
				domain.Rust: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Software Engineer, Rust",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4258500752/",
							Date:                 mustDate("2025-06-28"),
							WithSalary:           true,
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
			ShortDescription: "Generate SDKs in multiple languages and interactive API documentation",
			Industries: []domain.Industry{
				domain.IndustryDevOps,
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Bullet",
			Website: "https://www.bullet.xyz/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                83484451,
				IDs:               nil,
				Alias:             "bulletxyz",
				Name:              "Bullet",
				Followers:         "2K",
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
					Vacancies:               []domain.Vacancy{},
				},
				domain.Rust: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Senior Rust Developer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4256368070/",
							Date:                 mustDate("2025-06-27"),
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
			ShortDescription: "Solana's trading layer",
			Industries:       []domain.Industry{
				// NOP
			},
			Ignore: true, // Cryptocurrency
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Famedly GmbH",
			Website: "https://famedly.com/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                12976028,
				IDs:               nil,
				Alias:             "famedly",
				Name:              "Famedly GmbH",
				Followers:         "2K",
				Employees:         "11-50",
				AssociatedMembers: "36",
				Verified:          false,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:     "famedly",
				Followers: "58",
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
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Rust: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Python / Rust Developer",
							ShortDescription:     "Matrix Synapse Homeserver",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4258947366/",
							Date:                 mustDate("2025-06-29"),
							WithSalary:           true,
							Remote:               true,
						},
						{
							Title:                "Rust Developer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4331456289/",
							Location:             "Germany",
							Date:                 mustDate("2025-10-30"),
							WithSalary:           true, // €50k/yr - €70k/yr
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
			ShortDescription: "Medical collaboration platform",
			Industries: []domain.Industry{
				domain.IndustryMedTech,
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Chapa",
			Website: "https://chapa.co/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                75151479,
				IDs:               nil,
				Alias:             "chapa-financial-technologies",
				Name:              "Chapa",
				Followers:         "25K",
				Employees:         "11-50",
				AssociatedMembers: "61",
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
							Title:                "Backend Developer (Rust)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4258996962/",
							Date:                 mustDate("2025-06-29"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Backend Developer (Rust)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4305719609/",
							Location:             "Addis Ababa, Addis Ababa, Ethiopia",
							Date:                 mustDate("2025-09-27"),
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
			ShortDescription: "Payment gateway",
			Industries: []domain.Industry{
				domain.IndustryFinTech,
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "WebBeds",
			Website: "https://www.webbeds.com/",
			Careers: "https://www.webbeds.com/careers/",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                13598746,
				IDs:               []int{11694121, 13598746, 37510949},
				Alias:             "webbeds",
				Name:              "WebBeds",
				Followers:         "43K",
				Employees:         "1K-5K",
				AssociatedMembers: "1,407",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:     "WebBeds",
				Followers: "49",
				Verified:  true,
			},
			BlindProfile: domain.BlindProfile{
				Alias: "",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias: "",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-WebBeds-EI_IE3219218.11,18.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/WebBeds-Reviews-E3219218.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/WebBeds-Jobs-E3219218.htm",
				Jobs:        "104",
				Reviews:     "139",
				Salaries:    "244",
				ReviewsRate: "3.7",
				Verified:    false,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "",
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 3,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Golang Engineer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4256478997/",
							Date:                 mustDate("2025-07-01"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Senior Backend Developer (.NET / Go / PHP)",
							ShortDescription:     "High-Scale Systems | Cloud & Microservices",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4268595844/",
							Location:             "London, England, United Kingdom",
							Date:                 mustDate("2025-07-23"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Golang Engineer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4348622020/",
							Location:             "Palma, Balearic Islands, Spain",
							Date:                 mustDate("2025-12-06"),
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
			ShortDescription: "Marketplace for the travel trade",
			Industries:       []domain.Industry{
				// NOP
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Chemelex",
			Website: "https://chemelex.com/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                106314506,
				IDs:               nil,
				Alias:             "chemelex",
				Name:              "Chemelex",
				Followers:         "2K",
				Employees:         "5K-10K",
				AssociatedMembers: "319",
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
							Title:                "Embedded Rust/C++ Firmware Developer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4173967092/",
							Location:             "Edmonton, AB",
							Date:                 mustDate("2025-08-23", "2025-07-02"),
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
			ShortDescription: "Electric thermal and sensing solutions",
			Industries:       []domain.Industry{
				// NOP
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "WP Engine",
			Website: "https://wpengine.com/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                1556180,
				IDs:               nil,
				Alias:             "wpengine",
				Name:              "WP Engine",
				Followers:         "39K",
				Employees:         "1K-5K",
				AssociatedMembers: "1,123",
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
							Title:                "Senior Software Engineer, Platform Networking (Python/Go, Kubernetes)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4261099779/",
							Date:                 mustDate("2025-07-03"),
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
			ShortDescription: "",
			Industries:       []domain.Industry{
				// NOP
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "BNP Paribas",
			Website: "https://group.bnpparibas/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                1508,
				IDs:               []int{1508, 1519, 10070, 98774, 166278, 251838, 256009, 281207, 282760, 631981, 963211, 2642837, 3625182, 3627928, 3683364, 3880216, 11092068, 15245937, 18735883, 30590982},
				Alias:             "bnp-paribas",
				Name:              "BNP Paribas",
				Followers:         "2M",
				Employees:         "10K+",
				AssociatedMembers: "161,221",
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
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-BNP-Paribas-EI_IE10342.11,22.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/BNP-Paribas-Reviews-E10342.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/BNP-Paribas-Jobs-E10342.htm",
				Jobs:        "3K",
				Reviews:     "19K",
				Salaries:    "30K",
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
							Title:                "Go Full-Stack Developer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4261003515/",
							Location:             "Tamil Nadu, India",
							Date:                 mustDate("2025-07-24", "2025-07-02"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Senior Backend Developer (Golang/Python)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4318534106/",
							Location:             "Warsaw, Mazowieckie, Poland",
							Date:                 mustDate("2025-10-22"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Senior Backend Developer (Golang/Python)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4339050924/",
							Location:             "Warsaw, Mazowieckie, Poland",
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
					Vacancies: []domain.Vacancy{
						{
							Title:                "Senior Spark/Scala Developer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4313221564/",
							Location:             "Greater Chennai Area",
							Date:                 mustDate("2025-10-11"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Senior Spark/Scala Developer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4335431028/",
							Location:             "Chennai, Tamil Nadu, India",
							Date:                 mustDate("2025-12-02"),
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
			ShortDescription: "Bank",
			Industries: []domain.Industry{
				domain.IndustryFinTech,
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Distribusion Technologies",
			Website: "https://www.distribusion.com/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                2747271,
				IDs:               nil,
				Alias:             "distribusion",
				Name:              "Distribusion Technologies",
				Followers:         "14K",
				Employees:         "201-500",
				AssociatedMembers: "303",
				Verified:          false,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:     "distribusion",
				Followers: "24",
				Verified:  false,
			},
			BlindProfile: domain.BlindProfile{
				Alias: "",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias: "",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Distribusion-EI_IE2207226.11,23.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Distribusion-Reviews-E2207226.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Distribusion-Jobs-E2207226.htm",
				Jobs:        "50",
				Reviews:     "35",
				Salaries:    "64",
				ReviewsRate: "3.9",
				Verified:    true,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "",
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 1,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Rust: {
					GitHubRepositoriesCount: 1,
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
							Title:                "Senior Elixir Software Engineer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4261717421/",
							Location:             "Berlin, Berlin, Germany",
							Date:                 mustDate("2025-09-06", "2025-08-16", "2025-07-25", "2025-07-03"),
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
			ShortDescription: "Dround transport marketplace",
			Industries:       []domain.Industry{
				// NOP
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Thales",
			Website: "https://www.thalesgroup.com/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                1951,
				IDs:               []int{1951, 239078, 1003745},
				Alias:             "thales",
				Name:              "Thales",
				Followers:         "2M",
				Employees:         "10K+",
				AssociatedMembers: "66,621",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:     "thalesgroup",
				Followers: "414",
				Verified:  false,
			},
			BlindProfile: domain.BlindProfile{
				Alias: "",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "thales",
				Employees: "81,000",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Thales-EI_IE10358.11,17.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Thales-Reviews-E10358.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Thales-Jobs-E10358.htm",
				Jobs:        "2.6K",
				Reviews:     "9K",
				Salaries:    "16K",
				ReviewsRate: "3.9",
				Verified:    true,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "",
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 37,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Software Engineer – III (C++/Golang Developer)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4161416034/",
							Location:             "Noida, Uttar Pradesh, India",
							Date:                 mustDate("2025-07-23", "2025-07-01"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Senior Software Engineer – Golang, C/C++",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4299786439/",
							Location:             "Bengaluru, Karnataka, India",
							Date:                 mustDate("2025-09-13"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Senior Software Engineer – Golang, C/C++",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4310599223/",
							Location:             "Bengaluru, Karnataka, India",
							Date:                 mustDate("2025-10-05"),
							WithSalary:           false,
							Remote:               false,
						},
					},
				},
				domain.Rust: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Senior Software Engineer (Rust, JavaScript)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4283913630/",
							Location:             "Bengaluru, Karnataka, India",
							Date:                 mustDate("2025-09-04"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Senior Software Engineer (Rust, JavaScript)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4334557414/",
							Location:             "Bengaluru, Karnataka, India",
							Date:                 mustDate("2025-11-29", "2025-11-06"),
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
			ShortDescription: "",
			Industries: []domain.Industry{
				domain.IndustryCyberSecurity,
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Kissht",
			Website: "https://kissht.com/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                13216565,
				IDs:               []int{6471115, 13216565},
				Alias:             "kissht",
				Name:              "Kissht",
				Followers:         "39K",
				Employees:         "501-1K",
				AssociatedMembers: "1,617",
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
				Alias:     "kissht-finance",
				Employees: "1,000",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Kissht-EI_IE2068851.11,17.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Kissht-Reviews-E2068851.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Kissht-Jobs-E2068851.htm",
				Jobs:        "11",
				Reviews:     "177",
				Salaries:    "179",
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
							Title:                "Senior Golang Software Engineer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4260223469/",
							Date:                 mustDate("2025-07-01"),
							WithSalary:           false,
							Remote:               true,
						},
						{
							Title:                "Senior Golang Software Engineer (SDE 3)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4279998684/",
							Location:             "India",
							Date:                 mustDate("2025-08-06"),
							WithSalary:           false,
							Remote:               true,
						},
						{
							Title:                "Senior Golang Software Engineer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4285509703/",
							Location:             "India",
							Date:                 mustDate("2025-08-19"),
							WithSalary:           false,
							Remote:               true,
						},
						{
							Title:                "Senior Golang Software Engineer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4333235706/",
							Location:             "India",
							Date:                 mustDate("2025-11-01"),
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
			ShortDescription: "India’s credit platform",
			Industries: []domain.Industry{
				domain.IndustryFinTech,
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Motadata",
			Website: "https://www.motadata.com/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                13187469,
				IDs:               nil,
				Alias:             "motadata",
				Name:              "Motadata",
				Followers:         "21K",
				Employees:         "201-500",
				AssociatedMembers: "221",
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
				Alias:     "motadata",
				Employees: "180",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Motadata-EI_IE9939867.11,19.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Motadata-Reviews-E9939867.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Motadata-Jobs-E9939867.htm",
				Jobs:        "",
				Reviews:     "15",
				Salaries:    "17",
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
							Title:                "Full Stack Engineer – Golang",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4262651044/",
							Date:                 mustDate("2025-07-05"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Full Stack Developer – Golang/React.js",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4299703305/",
							Location:             "Ahmedabad, Gujarat, India",
							Date:                 mustDate("2025-09-13"),
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
			ShortDescription: "IT monitoring and management software",
			Industries: []domain.Industry{
				domain.IndustryDevOps,
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "CloudBees",
			Website: "https://www.cloudbees.com/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                1189836,
				IDs:               nil,
				Alias:             "cloudbees",
				Name:              "CloudBees",
				Followers:         "66K",
				Employees:         "501-1K",
				AssociatedMembers: "594",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:     "cloudbees",
				Followers: "112",
				Verified:  true,
			},
			BlindProfile: domain.BlindProfile{
				Alias:       "CloudBees",
				Employees:   "501 to 1,000",
				Salary:      "",
				Reviews:     "11",
				ReviewsRate: "3.6",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "cloudbees",
				Employees: "510",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-CloudBees-EI_IE1106279.11,20.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/CloudBees-Reviews-E1106279.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/CloudBees-Jobs-E1106279.htm",
				Jobs:        "76",
				Reviews:     "227",
				Salaries:    "357",
				ReviewsRate: "4.0",
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
							Title:                "Senior Software Engineer (Golang)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4208216235/",
							Date:                 mustDate("2025-07-06"),
							WithSalary:           false,
							Remote:               true,
						},
						{
							Title:                "Senior Software Engineer – Golang",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4315319289/",
							Location:             "Bengaluru, Karnataka, India",
							Date:                 mustDate("2025-11-28", "2025-10-16"),
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
			ShortDescription: "DevOps solution",
			Industries: []domain.Industry{
				domain.IndustryDevOps,
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "SecPod",
			Website: "https://www.secpod.com/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                216212,
				IDs:               nil,
				Alias:             "secpod-technologies",
				Name:              "SecPod",
				Followers:         "57K",
				Employees:         "51-200",
				AssociatedMembers: "147",
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
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-SecPod-Technologies-EI_IE929477.11,30.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/SecPod-Technologies-Reviews-E929477.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/SecPod-Technologies-Jobs-E929477.htm",
				Jobs:        "",
				Reviews:     "36",
				Salaries:    "36",
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
							Title:                "Full Stack Developer (Python/Golang + React)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4260697255/",
							Date:                 mustDate("2025-07-04"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Full Stack Developer (Python/Golang + React)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4271809120/",
							Location:             "Warsaw, Mazowieckie, Poland",
							Date:                 mustDate("2025-07-24"),
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
			ShortDescription: "CyberSecurity company committed to preventing cyberattacks through proactive security",
			Industries: []domain.Industry{
				domain.IndustryCyberSecurity,
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Cielara AI",
			Website: "https://cielara.ai/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                104159883,
				IDs:               nil,
				Alias:             "cielara",
				Name:              "Cielara AI",
				Followers:         "1K",
				Employees:         "2-10",
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
					Vacancies: []domain.Vacancy{
						{
							Title:                "Backend Developer (Golang)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4261529674/",
							Date:                 mustDate("2025-07-06"),
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
			ShortDescription: "Cloud computing with a real-time digital twin and an AI reasoning layer",
			Industries: []domain.Industry{
				domain.IndustryDevOps,
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "SolarWinds",
			Website: "https://www.solarwinds.com/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                166039,
				IDs:               []int{129912, 166039, 262758, 550008, 2299482},
				Alias:             "solarwinds",
				Name:              "SolarWinds",
				Followers:         "240K",
				Employees:         "1K-5K",
				AssociatedMembers: "2,803",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:     "solarwinds",
				Followers: "166",
				Verified:  true,
			},
			BlindProfile: domain.BlindProfile{
				Alias:       "SolarWinds",
				Employees:   "1,001 to 5,000",
				Salary:      "$50K ~ $180K a year",
				Reviews:     "26",
				ReviewsRate: "3.3",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "solarwinds",
				Employees: "3,500",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-SolarWinds-EI_IE100286.11,21.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/SolarWinds-Reviews-E100286.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/SolarWinds-Jobs-E100286.htm",
				Jobs:        "89",
				Reviews:     "1.5K",
				Salaries:    "2.3K",
				ReviewsRate: "3.6",
				Verified:    true,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "Solarwinds",
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 24,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Senior Software Engineer – Go Developer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4255188786/",
							Date:                 mustDate("2025-07-03"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Senior Software Engineer – Go Developer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4279380810/",
							Location:             "Cracow, Małopolskie, Poland",
							Date:                 mustDate("2025-08-04"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Senior Software Engineer with Go",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4287963124/",
							Location:             "Brno, South Moravia, Czechia",
							Date:                 mustDate("2025-08-19"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Senior Software Engineer with Go",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4306313219/",
							Location:             "Brno, South Moravia, Czechia",
							Date:                 mustDate("2025-10-06", "2025-09-29"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Senior Software Engineer with Go",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4334029045/",
							Location:             "Brno, South Moravia, Czechia",
							Date:                 mustDate("2025-11-16"),
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
				domain.IndustryDevOps,
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Blizzard Entertainment",
			Website: "https://blizzard.com/",
			Careers: "https://careers.blizzard.com/",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                5329,
				IDs:               nil,
				Alias:             "blizzard-entertainment",
				Name:              "Blizzard Entertainment",
				Followers:         "1M",
				Employees:         "1K-5K",
				AssociatedMembers: "5,798",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:     "blizzard",
				Followers: "413",
				Verified:  false,
			},
			BlindProfile: domain.BlindProfile{
				Alias:       "Activision-Blizzard",
				Employees:   "5,001 to 10,000",
				Salary:      "$49K ~ $315K a year",
				Reviews:     "179",
				ReviewsRate: "3.3",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "blizzard-entertainment",
				Employees: "5,500",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Blizzard-Entertainment-EI_IE24858.11,33.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Blizzard-Entertainment-Reviews-E24858.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Blizzard-Entertainment-Jobs-E24858.htm",
				Jobs:        "85",
				Reviews:     "1.6K",
				Salaries:    "3.7K",
				ReviewsRate: "3.4",
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
							Title:                "Senior Go Software Engineer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4204059285/",
							Location:             "Santa Monica, CA",
							Date:                 mustDate("2025-08-19", "2025-07-27", "2025-07-27", "2025-07-06"),
							WithSalary:           true, // $101k/yr - $186.8k/yr
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
			ShortDescription: "Developer and publisher of entertainment software",
			Industries: []domain.Industry{
				domain.IndustryGameDev,
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "NinjaOne",
			Website: "https://www.ninjaone.com/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                6436301,
				IDs:               nil,
				Alias:             "ninjaone",
				Name:              "NinjaOne",
				Followers:         "47K",
				Employees:         "1K-5K",
				AssociatedMembers: "1,596",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:     "NinjaRMM",
				Followers: "72",
				Verified:  false,
			},
			BlindProfile: domain.BlindProfile{
				Alias: "",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "ninjaone",
				Employees: "330",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-NinjaOne-EI_IE1595050.11,19.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/NinjaOne-Reviews-E1595050.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/NinjaOne-Jobs-E1595050.htm",
				Jobs:        "77",
				Reviews:     "331",
				Salaries:    "478",
				ReviewsRate: "4.0",
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
							Title:                "Senior Software Engineer, Go",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4207030100/",
							Date:                 mustDate("2025-07-06"),
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
			ShortDescription: "Endpoint management platform",
			Industries: []domain.Industry{
				domain.IndustryDevOps,
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Aviva",
			Website: "https://www.aviva.com/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                2922,
				IDs:               []int{2922, 228054, 717083, 958252, 1170046, 1389172, 2504608},
				Alias:             "aviva-plc",
				Name:              "Aviva",
				Followers:         "309K",
				Employees:         "10K+",
				AssociatedMembers: "24,985",
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
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Aviva-EI_IE311641.11,16.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Aviva-Reviews-E311641.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Aviva-Jobs-E311641.htm",
				Jobs:        "228",
				Reviews:     "2.9K",
				Salaries:    "5.4K",
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
							Title:                "Go .NET Developer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4258636824/",
							Date:                 mustDate("2025-07-02"),
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
			ShortDescription: "Insurance, Wealth & Retirement business",
			Industries: []domain.Industry{
				domain.IndustryInsurTech,
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Sectigo",
			Website: "https://www.sectigo.com/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                35431914,
				IDs:               nil,
				Alias:             "sectigo",
				Name:              "Sectigo",
				Followers:         "65K",
				Employees:         "501-1K",
				AssociatedMembers: "539",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:     "sectigo",
				Followers: "61",
				Verified:  true,
			},
			BlindProfile: domain.BlindProfile{
				Alias: "",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "sectigo",
				Employees: "351",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Sectigo-EI_IE2544320.11,18.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Sectigo-Reviews-E2544320.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Sectigo-Jobs-E2544320.htm",
				Jobs:        "140",
				Reviews:     "86",
				Salaries:    "144",
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
							Title:                "Golang Developer with DevOps",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4248510432/",
							Date:                 mustDate("2025-07-05"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Golang & Python Software Developer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4247335602/",
							Location:             "Ottawa, ON",
							Date:                 mustDate("2025-07-25"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Golang & Python Software Developer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4282001020/",
							Location:             "Ottawa, ON",
							Date:                 mustDate("2025-08-10"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Middle Golang and Java Developer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4289593401/",
							Location:             "Ottawa, ON",
							Date:                 mustDate("2025-10-05", "2025-08-23"),
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
			ShortDescription: "Certificate Lifecycle Management",
			Industries:       []domain.Industry{
				// NOP
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Tabeo",
			Website: "https://tabeo.co.uk/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                18303026,
				IDs:               nil,
				Alias:             "tabeo",
				Name:              "Tabeo",
				Followers:         "8K",
				Employees:         "11-50",
				AssociatedMembers: "44",
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
							Title:                "Backend Engineer (Go)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4251716865/",
							Date:                 mustDate("2025-06-23"),
							WithSalary:           false,
							Remote:               true,
						},
						{
							Title:                "Senior Backend Engineer (Go)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4300901459/",
							Location:             "London, England, United Kingdom",
							Date:                 mustDate("2025-09-16"),
							WithSalary:           false,
							Remote:               true,
						},
						{
							Title:                "Senior Backend Engineer (Go)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4304112042/",
							Location:             "United Kingdom",
							Date:                 mustDate("2025-09-22"),
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
			ShortDescription: "Tabeo’s mission is to simplify payments",
			Industries: []domain.Industry{
				domain.IndustryFinTech,
			},
		},
	}
}
