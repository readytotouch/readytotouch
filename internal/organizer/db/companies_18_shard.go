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
							Date:                 mustDate("2025-06-07"),
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
							Date:                 mustDate("2025-06-09"), // mustDate("2025-06-06"),
							WithSalary:           true,
							Remote:               true,
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
				Alias: "",
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
				Alias: "",
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
							Date:                 mustDate("2025-06-07"),
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
			Careers: "",
			About:   "",
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
				Alias: "",
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
							Title:                "Senior Clojure Developer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4248410642/",
							Date:                 mustDate("2025-06-14"),
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
				Alias: "",
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
							Date:                 mustDate("2025-06-18"),
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
							Title:                "Senior Software Engineer (Rust)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4248785636/",
							Date:                 mustDate("2025-06-14"),
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
							URL:                  "https://www.linkedin.com/jobs/view/4253509850/",
							Date:                 mustDate("2025-06-20"),
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
