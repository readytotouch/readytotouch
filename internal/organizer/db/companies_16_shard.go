package db

import "github.com/readytotouch/readytotouch/internal/domain"

func companies16Shard() []domain.CompanyProfile {
	return []domain.CompanyProfile{
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Cabify",
			Website: "https://cabify.com/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                2394728,
				IDs:               nil,
				Alias:             "cabify",
				Name:              "Cabify",
				Followers:         "175K",
				Employees:         "1K-5K",
				AssociatedMembers: "3,945",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:     "cabify",
				Followers: "46",
				Verified:  false,
			},
			BlindProfile: domain.BlindProfile{
				Alias: "",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "cabify",
				Employees: "2,500",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Cabify-EI_IE975202.11,17.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Cabify-Reviews-E975202.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Cabify-Jobs-E975202.htm",
				Jobs:        "78",
				Reviews:     "531",
				Salaries:    "975",
				ReviewsRate: "3.8",
				Verified:    true,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "",
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 38,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Backend Software Engineer",
							ShortDescription:     "",
							SwitchingOpportunity: "We are looking for engineers with experience working with any backend technology, you will be working with Go, Elixir or Ruby for your day-to-day work",
							URL:                  "https://www.linkedin.com/jobs/view/4068277723/",
							Date:                 mustDate("2024-11-17"),
							WithSalary:           false,
							Remote:               true,
						},
					},
				},
				domain.Rust:  {},
				domain.Zig:   {},
				domain.Scala: {},
				domain.Elixir: {
					GitHubRepositoriesCount: 6,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Backend Software Engineer",
							ShortDescription:     "",
							SwitchingOpportunity: "We are looking for engineers with experience working with any backend technology, you will be working with Go, Elixir or Ruby for your day-to-day work",
							URL:                  "https://www.linkedin.com/jobs/view/4068277723/",
							Date:                 mustDate("2024-11-17"),
							WithSalary:           false,
							Remote:               true,
						},
					},
				},
				domain.Clojure: {},
				domain.Haskell: {},
			},
			ShortDescription: "Transform cities into better living spaces and make them more sustainable, accessible, and humane",
			Industries:       []domain.Industry{
				// NOP
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "CoverMyMeds",
			Website: "https://www.covermymeds.health/",
			Careers: "",
			About:   "https://experience.covermymeds.com/what-we-do",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                1741261,
				IDs:               nil,
				Alias:             "covermymeds",
				Name:              "CoverMyMeds",
				Followers:         "55K",
				Employees:         "1K-5K",
				AssociatedMembers: "30",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:     "covermymeds",
				Followers: "20",
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
				domain.Go:    {},
				domain.Rust:  {},
				domain.Zig:   {},
				domain.Scala: {},
				domain.Elixir: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{
						// https://www.linkedin.com/jobs/view/software-developer-elixir-phoenix-at-covermymeds-4102099804
					},
				},
				domain.Clojure: {},
				domain.Haskell: {},
			},
			ShortDescription: "",
			Industries:       []domain.Industry{
				// NOP
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Together AI",
			Website: "https://www.together.ai/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                89816302,
				IDs:               nil,
				Alias:             "togethercomputer",
				Name:              "Together AI",
				Followers:         "53K",
				Employees:         "51-200",
				AssociatedMembers: "193",
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
				domain.Go: {},
				domain.Rust: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Rust Systems Engineer",
							ShortDescription:     "Inference",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4214931469/",
							Location:             "San Francisco, CA",
							Date:                 mustDate("2025-09-22", "2025-09-01", "2025-08-11", "2025-07-20", "2025-06-28", "2025-06-05", "2025-05-15", "2025-04-24"),
							WithSalary:           true, // $160k/yr - $240k/yr
							Remote:               false,
						},
					},
				},
				domain.Zig:   {},
				domain.Scala: {},
				domain.Elixir: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Elixir Backend Engineer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4177515860/",
							Date:                 mustDate("2025-03-17"),
							WithSalary:           false,
							Remote:               false,
						},
					},
				},
				domain.Clojure: {},
				domain.Haskell: {},
			},
			ShortDescription: "Research-driven artificial intelligence company",
			Industries:       []domain.Industry{
				// NOP
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Telnyx",
			Website: "https://telnyx.com/",
			Careers: "https://telnyx.com/careers",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                3349412,
				IDs:               nil,
				Alias:             "telnyx",
				Name:              "Telnyx",
				Followers:         "20K",
				Employees:         "201-500",
				AssociatedMembers: "309",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:     "team-telnyx",
				Followers: "103",
				Verified:  true,
			},
			BlindProfile: domain.BlindProfile{
				Alias:       "Telnyx",
				Employees:   "51 to 200",
				Salary:      "",
				Reviews:     "7",
				ReviewsRate: "3.1",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias: "",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Telnyx-EI_IE841349.11,17.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Telnyx-Reviews-E841349.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Telnyx-Jobs-E841349.htm",
				Jobs:        "2",
				Reviews:     "122",
				Salaries:    "229",
				ReviewsRate: "3.7",
				Verified:    true,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "",
			},
			Languages: domain.Languages{
				domain.Go:    {},
				domain.Rust:  {},
				domain.Zig:   {},
				domain.Scala: {},
				domain.Elixir: {
					GitHubRepositoriesCount: 32,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Senior Software Engineer, Elixir",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4211721655/",
							Date:                 mustDate("2025-04-16"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Senior Software Engineer, Elixir",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4226344237/",
							Date:                 mustDate("2025-05-31", "2025-05-09"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Senior Software Engineer, Elixir",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4233635680/",
							Date:                 mustDate("2025-05-21"),
							WithSalary:           false,
							Remote:               true,
						},
						{
							Title:                "Software Engineer, Elixir",
							ShortDescription:     "Wireless",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4250479645/",
							Date:                 mustDate("2025-06-14"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Software Engineer, Elixir",
							ShortDescription:     "Wireless",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4261024126/",
							Location:             "Warsaw, Mazowieckie, Poland",
							Date:                 mustDate("2025-07-24", "2025-07-02"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Software Engineer, Elixir",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4286234802/",
							Location:             "Cracow, Małopolskie, Poland",
							Date:                 mustDate("2025-10-18", "2025-09-26", "2025-09-06", "2025-08-16"),
							WithSalary:           false,
							Remote:               true,
						},
						{
							Title:                "Senior Software Engineer, Elixir",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4301511792/",
							Location:             "São Paulo, São Paulo, Brazil",
							Date:                 mustDate("2025-09-16"),
							WithSalary:           false,
							Remote:               true,
						},
						{
							Title:                "Senior Software Engineer, Elixir",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4304194138/",
							Location:             "Sydney, New South Wales, Australia",
							Date:                 mustDate("2025-10-15"),
							WithSalary:           false,
							Remote:               true,
						},
						{
							Title:                "Senior Elixir Engineer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4343031500/",
							Location:             "Dublin, County Dublin, Ireland",
							Date:                 mustDate("2025-12-10"),
							WithSalary:           false,
							Remote:               false,
						},
					},
				},
				domain.Clojure: {},
				domain.Haskell: {},
			},
			ShortDescription: "Build real-time conversational AI with global voice",
			Industries:       []domain.Industry{
				// NOP
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Euna Solutions",
			Website: "https://eunasolutions.com/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                60200222,
				IDs:               []int{9604627, 60200222},
				Alias:             "eunasolutions",
				Name:              "Euna Solutions",
				Followers:         "7K",
				Employees:         "501-1K",
				AssociatedMembers: "558",
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
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Euna-Solutions-EI_IE8940329.11,25.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Euna-Solutions-Reviews-E8940329.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Euna-Solutions-Jobs-E8940329.htm",
				Jobs:        "29",
				Reviews:     "33",
				Salaries:    "93",
				ReviewsRate: "3.5",
				Verified:    false,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "",
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
							Title:                "Software Developer (Elixir)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4186581023/",
							Date:                 mustDate("2025-04-11"),
							WithSalary:           false,
							Remote:               false,
						},
					},
				},
				domain.Clojure: {},
				domain.Haskell: {},
			},
			ShortDescription: "SaaS solutions for the public sector",
			Industries:       []domain.Industry{
				// NOP
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Driftrock",
			Website: "https://www.driftrock.com/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                3087205,
				IDs:               nil,
				Alias:             "driftrock",
				Name:              "Driftrock",
				Followers:         "5K",
				Employees:         "11-50",
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
				OverviewURL: "",
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "",
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
							Title:                "Senior Software Engineer",
							ShortDescription:     "Mainly built in Elixir",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4188235715/",
							Date:                 mustDate("2025-03-20"),
							WithSalary:           false,
							Remote:               true,
						},
					},
				},
				domain.Clojure: {},
				domain.Haskell: {},
			},
			ShortDescription: "Automotive Marketing Platform",
			Industries: []domain.Industry{
				domain.IndustryMarTech,
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "GreyScout",
			Website: "https://greyscout.com/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                65338573,
				IDs:               nil,
				Alias:             "greyscout",
				Name:              "GreyScout",
				Followers:         "2K",
				Employees:         "11-50",
				AssociatedMembers: "21",
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
				domain.Go:    {},
				domain.Rust:  {},
				domain.Zig:   {},
				domain.Scala: {},
				domain.Elixir: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Elixir Developer",
							ShortDescription:     "Our backend engineering team is composed of Elixir engineers, who have a strong Ruby background",
							SwitchingOpportunity: "At a minimum you should have a strong interest in Elixir",
							URL:                  "https://www.linkedin.com/jobs/view/3758901108/",
							Date:                 mustDate("2024-04-17", "2025-04-17"),
							WithSalary:           false,
							Remote:               false,
						},
					},
				},
				domain.Clojure: {},
				domain.Haskell: {},
			},
			ShortDescription: "Platform that lets you quickly detect, authenticate, and take enforcement action against grey market sellers",
			Industries:       []domain.Industry{
				// NOP
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Padlet",
			Website: "https://padlet.com/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                2644238,
				IDs:               nil,
				Alias:             "padlet",
				Name:              "Padlet",
				Followers:         "7K",
				Employees:         "51-200",
				AssociatedMembers: "66",
				Verified:          false,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:     "padlet",
				Followers: "17",
				Verified:  false,
			},
			BlindProfile: domain.BlindProfile{
				Alias: "",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias: "",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Padlet-EI_IE2379906.11,17.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Padlet-Reviews-E2379906.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Padlet-Jobs-E2379906.htm",
				Jobs:        "17",
				Reviews:     "17",
				Salaries:    "38",
				ReviewsRate: "4.1",
				Verified:    true,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "",
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
							Title:                "Senior Software Engineer – Elixir",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/3860724318/",
							Date:                 mustDate("2024-04-17", "2025-04-17"),
							WithSalary:           false,
							Remote:               false,
						},
					},
				},
				domain.Clojure: {},
				domain.Haskell: {},
			},
			ShortDescription: "",
			Industries: []domain.Industry{
				domain.IndustryEdTech,
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "proSapient",
			Website: "https://www.prosapient.com/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                11165963,
				IDs:               nil,
				Alias:             "prosapient",
				Name:              "proSapient",
				Followers:         "48K",
				Employees:         "201-500",
				AssociatedMembers: "437",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:     "prosapient",
				Followers: "1",
				Verified:  false,
			},
			BlindProfile: domain.BlindProfile{
				Alias: "",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias: "",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-ProSapient-EI_IE2288318.11,21.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/ProSapient-Reviews-E2288318.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/ProSapient-Jobs-E2288318.htm",
				Jobs:        "12",
				Reviews:     "216",
				Salaries:    "521",
				ReviewsRate: "3.4",
				Verified:    true,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "",
			},
			Languages: domain.Languages{
				domain.Go:    {},
				domain.Rust:  {},
				domain.Zig:   {},
				domain.Scala: {},
				domain.Elixir: {
					GitHubRepositoriesCount: 23,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Elixir Developer",
							ShortDescription:     "Develop and maintain scalable, efficient, and high-quality Elixir web applications",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4178411182/",
							Date:                 mustDate("2025-02-17"),
							WithSalary:           false,
							Remote:               true,
						},
						{
							Title:                "Elixir Developer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4321865679/",
							Location:             "Greater London, England, United Kingdom",
							Date:                 mustDate("2025-11-12"),
							WithSalary:           false,
							Remote:               false,
						},
					},
				},
				domain.Clojure: {},
				domain.Haskell: {},
			},
			ShortDescription: "proSapient is a expert network and primary research platform that enables professionals to share insights with clients through short calls, online surveys and long-term engagements",
			Industries:       []domain.Industry{
				// NOP
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Angel Studios",
			Website: "https://www.angel.com/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                72350313,
				IDs:               nil,
				Alias:             "angel-studios-page",
				Name:              "Angel Studios",
				Followers:         "26K",
				Employees:         "51-200",
				AssociatedMembers: "276",
				Verified:          false,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:     "Angel-Studios",
				Followers: "17",
				Verified:  false,
			},
			BlindProfile: domain.BlindProfile{
				Alias: "",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias: "",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Angel-Studios-EI_IE5687780.11,24.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Angel-Studios-Reviews-E5687780.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Angel-Studios-Jobs-E5687780.htm",
				Jobs:        "",
				Reviews:     "53",
				Salaries:    "94",
				ReviewsRate: "4.1",
				Verified:    true,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "",
			},
			Languages: domain.Languages{
				domain.Go:    {},
				domain.Rust:  {},
				domain.Zig:   {},
				domain.Scala: {},
				domain.Elixir: {
					GitHubRepositoriesCount: 10,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Senior Backend Engineer (Elixir)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4075796424/",
							Date:                 mustDate("2024-11-17"),
							WithSalary:           true,
							Remote:               false,
						},
					},
				},
				domain.Clojure: {},
				domain.Haskell: {},
			},
			ShortDescription: "We are a film studio platform that helps creators come together with viewers to create high-quality TV and film without answering to Hollywood",
			Industries:       []domain.Industry{
				// NOP
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Sleeper",
			Website: "https://sleeper.com/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                3742895,
				IDs:               nil,
				Alias:             "sleeperhq",
				Name:              "Sleeper",
				Followers:         "14K",
				Employees:         "51-200",
				AssociatedMembers: "187",
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
				domain.Go:    {},
				domain.Rust:  {},
				domain.Zig:   {},
				domain.Scala: {},
				domain.Elixir: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Backend Engineer",
							ShortDescription:     "Our backend consists primarily of Elixir and ScyllaDB",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4130988748/",
							Date:                 mustDate("2025-02-17"),
							WithSalary:           true,
							Remote:               true,
						},
					},
				},
				domain.Clojure: {},
				domain.Haskell: {},
			},
			ShortDescription: "Digital playground for sports fans and their friends to hang out together",
			Industries:       []domain.Industry{
				// NOP
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Cadmus",
			Website: "https://cadmus.io/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                7941136,
				IDs:               nil,
				Alias:             "cadmus-io",
				Name:              "Cadmus",
				Followers:         "18K",
				Employees:         "51-200",
				AssociatedMembers: "69",
				Verified:          false,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:     "Vericus",
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
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Cadmus-EI_IE4368640.11,17.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Cadmus-Reviews-E4368640.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Cadmus-Jobs-E4368640.htm",
				Jobs:        "45",
				Reviews:     "10",
				Salaries:    "41",
				ReviewsRate: "2.7",
				Verified:    false,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "",
			},
			Languages: domain.Languages{
				domain.Go: {},
				domain.Rust: {
					GitHubRepositoriesCount: 3,
				},
				domain.Zig:   {},
				domain.Scala: {},
				domain.Elixir: {
					GitHubRepositoriesCount: 5,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Senior Software Engineer",
							ShortDescription:     "BE: Elixir, FE: React",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4108509122/",
							Date:                 mustDate("2024-12-17"),
							WithSalary:           false,
							Remote:               true,
						},
					},
				},
				domain.Clojure: {},
				domain.Haskell: {},
			},
			ShortDescription: "",
			Industries: []domain.Industry{
				domain.IndustryEdTech,
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Simpli.fi",
			Website: "https://simpli.fi/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                801567,
				IDs:               nil,
				Alias:             "simpli.fi",
				Name:              "Simpli.fi",
				Followers:         "18K",
				Employees:         "501-1K",
				AssociatedMembers: "665",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:     "simplifi",
				Followers: "28",
				Verified:  true,
			},
			BlindProfile: domain.BlindProfile{
				Alias: "",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "simplifi",
				Employees: "518",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Simpli-fi-EI_IE755206.11,20.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Simpli-fi-Reviews-E755206.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Simpli-fi-Jobs-E755206.htm",
				Jobs:        "3",
				Reviews:     "156",
				Salaries:    "459",
				ReviewsRate: "3.5",
				Verified:    true,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "",
			},
			Languages: domain.Languages{
				domain.Go:    {},
				domain.Rust:  {},
				domain.Zig:   {},
				domain.Scala: {},
				domain.Elixir: {
					GitHubRepositoriesCount: 8,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Elixir Software Developer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4128427094/",
							Date:                 mustDate("2025-02-18"),
							WithSalary:           true, // $140k/yr - $160k/yr
							Remote:               true,
						},
						{
							Title:                "Elixir Software Developer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4265373102/",
							Location:             "Fort Worth, TX",
							Date:                 mustDate("2025-07-12"),
							WithSalary:           true, // $140k/yr - $160k/yr
							Remote:               false,
						},
						{
							Title:                "Elixir Software Developer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4296340972/",
							Location:             "Fort Worth, TX",
							Date:                 mustDate("2025-09-06"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Elixir Software Developer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4301948734/",
							Location:             "Fort Worth, TX",
							Date:                 mustDate("2025-09-23", "2025-09-18"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Senior Elixir Software Developer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4315785503/",
							Location:             "Fort Worth, TX",
							Date:                 mustDate("2025-10-19"),
							WithSalary:           false,
							Remote:               false,
						},
					},
				},
				domain.Clojure: {},
				domain.Haskell: {},
			},
			ShortDescription: "Advertising success platform for digital marketers",
			Industries: []domain.Industry{
				domain.IndustryAdTech,
				domain.IndustryMarTech,
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "DraftKings Inc.",
			Website: "https://www.draftkings.com/",
			Careers: "https://careers.draftkings.com/",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                2507852,
				IDs:               []int{2038354, 2507852},
				Alias:             "draftkings-inc-",
				Name:              "DraftKings Inc.",
				Followers:         "138K",
				Employees:         "1K-5K",
				AssociatedMembers: "5,038",
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
				domain.Go:    {},
				domain.Rust:  {},
				domain.Zig:   {},
				domain.Scala: {},
				domain.Elixir: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Senior Software Engineer, Elixir",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4156575025/",
							Date:                 mustDate("2025-03-18"),
							WithSalary:           true,
							Remote:               true,
						},
					},
				},
				domain.Clojure: {},
				domain.Haskell: {},
			},
			ShortDescription: "",
			Industries:       []domain.Industry{
				// NOP
			},
			Ignore: true, // Casino
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Instinct Science",
			Website: "https://www.instinct.vet/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                16187981,
				IDs:               []int{3293342, 16187981},
				Alias:             "instinct-science",
				Name:              "Instinct Science",
				Followers:         "7K",
				Employees:         "51-200",
				AssociatedMembers: "177",
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
				domain.Go:    {},
				domain.Rust:  {},
				domain.Zig:   {},
				domain.Scala: {},
				domain.Elixir: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Engineering Manager",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4168905529/",
							Date:                 mustDate("2025-03-18"),
							WithSalary:           true,
							Remote:               true,
						},
					},
				},
				domain.Clojure: {},
				domain.Haskell: {},
			},
			ShortDescription: "Veterinary software system",
			Industries: []domain.Industry{
				domain.IndustryMedTech,
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "PENN Interactiv",
			Website: "https://www.pennentertainment.com/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                16182384,
				IDs:               nil,
				Alias:             "penn-interactive-pi",
				Name:              "PENN Interactive",
				Followers:         "24K",
				Employees:         "201-500",
				AssociatedMembers: "354",
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
				domain.Go:    {},
				domain.Rust:  {},
				domain.Zig:   {},
				domain.Scala: {},
				domain.Elixir: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Engineering Manager",
							ShortDescription:     "",
							SwitchingOpportunity: "Nice to have: Elixir",
							URL:                  "https://www.linkedin.com/jobs/view/4144534826/",
							Date:                 mustDate("2025-02-18"),
							WithSalary:           false,
							Remote:               false,
						},
					},
				},
				domain.Clojure: {},
				domain.Haskell: {},
			},
			ShortDescription: "",
			Industries:       []domain.Industry{
				// NOP
			},
			Ignore: true, // Gambling Facilities and Casinos
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Peek",
			Website: "https://www.peek.com/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                2468911,
				IDs:               nil,
				Alias:             "peek-com",
				Name:              "Peek",
				Followers:         "18K",
				Employees:         "51-200",
				AssociatedMembers: "299",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:     "peek-travel",
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
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Peek-EI_IE926518.11,15.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Peek-Reviews-E926518.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Peek-Jobs-E926518.htm",
				Jobs:        "13",
				Reviews:     "352",
				Salaries:    "364",
				ReviewsRate: "4.0",
				Verified:    true,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "",
			},
			Languages: domain.Languages{
				domain.Go:    {},
				domain.Rust:  {},
				domain.Zig:   {},
				domain.Scala: {},
				domain.Elixir: {
					GitHubRepositoriesCount: 33,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Full Stack Engineer",
							ShortDescription:     "We have a modern tech stack, built with Elixir, Ruby, Ember.js, and Swift",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4148678553/",
							Date:                 mustDate("2025-02-18"),
							WithSalary:           false,
							Remote:               true,
						},
					},
				},
				domain.Clojure: {},
				domain.Haskell: {},
			},
			ShortDescription: "Marketplace for consumers to book fun things to do like wine tours, watersports, skydiving, art classes, and more",
			Industries:       []domain.Industry{
				// NOP
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "PFF",
			Website: "https://www.pff.com/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                2992703,
				IDs:               nil,
				Alias:             "pro-football-focus",
				Name:              "PFF",
				Followers:         "17K",
				Employees:         "201-500",
				AssociatedMembers: "743",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:     "pro-football-focus",
				Followers: "39",
				Verified:  false,
			},
			BlindProfile: domain.BlindProfile{
				Alias: "",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias: "",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Pro-Football-Focus-EI_IE1628531.11,29.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Pro-Football-Focus-Reviews-E1628531.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Pro-Football-Focus-Jobs-E1628531.htm",
				Jobs:        "3",
				Reviews:     "102",
				Salaries:    "84",
				ReviewsRate: "4.1",
				Verified:    false,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "",
			},
			Languages: domain.Languages{
				domain.Go:    {},
				domain.Rust:  {},
				domain.Zig:   {},
				domain.Scala: {},
				domain.Elixir: {
					GitHubRepositoriesCount: 15,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Head of Engineering",
							ShortDescription:     "",
							SwitchingOpportunity: "Working knowledge of Elixir, Phoenix, Broadway, Oban is a plus",
							URL:                  "https://www.linkedin.com/jobs/view/4183960567/",
							Date:                 mustDate("2025-03-18"),
							WithSalary:           false,
							Remote:               true,
						},
					},
				},
				domain.Clojure: {},
				domain.Haskell: {},
			},
			ShortDescription: "PFF is transforming the sports performance analytics experience end-to-end, building products that help NFL & NCAA and soccer front offices, coaching staffs, media partners and digital subscribers evaluate player performance at an unmatched level",
			Industries:       []domain.Industry{
				// NOP
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "NIOV LABS",
			Website: "https://www.niovlabs.com/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                104105558,
				IDs:               nil,
				Alias:             "niov-labs",
				Name:              "NIOV LABS",
				Followers:         "187",
				Employees:         "2-10",
				AssociatedMembers: "5",
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
				domain.Go:    {},
				domain.Rust:  {},
				domain.Zig:   {},
				domain.Scala: {},
				domain.Elixir: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Elixir Developer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4152641983/",
							Date:                 mustDate("2025-02-18"),
							WithSalary:           false,
							Remote:               true,
						},
					},
				},
				domain.Clojure: {},
				domain.Haskell: {},
			},
			ShortDescription: "Company focused on transforming how businesses and users interact with data in the digital age",
			Industries:       []domain.Industry{
				// NOP
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Remote",
			Website: "https://remote.com/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                17955831,
				IDs:               nil,
				Alias:             "remote.com",
				Name:              "Remote",
				Followers:         "737K",
				Employees:         "1K-5K",
				AssociatedMembers: "9,165",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:     "remoteoss",
				Followers: "88",
				Verified:  false,
			},
			BlindProfile: domain.BlindProfile{
				Alias:       "Remote",
				Employees:   "201 to 500",
				Salary:      "",
				Reviews:     "21",
				ReviewsRate: "3.0",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "remote",
				Employees: "2,000",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Remote-EI_IE3871683.11,17.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Remote-Reviews-E3871683.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Remote-Jobs-E3871683.htm",
				Jobs:        "5",
				Reviews:     "611",
				Salaries:    "1.3K",
				ReviewsRate: "3.6",
				Verified:    true,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "",
			},
			Languages: domain.Languages{
				domain.Go:    {},
				domain.Rust:  {},
				domain.Zig:   {},
				domain.Scala: {},
				domain.Elixir: {
					GitHubRepositoriesCount: 17,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Backend Engineer – Elixir",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4005825587/",
							Date:                 mustDate("2024-08-18"),
							WithSalary:           false,
							Remote:               true,
						},
						{
							Title:                "Senior Backend Engineer (Elixir)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4220171393/",
							Date:                 mustDate("2025-05-01"),
							WithSalary:           false,
							Remote:               true,
						},
						{
							Title:                "Senior Backend Engineer (Elixir)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4228998540/",
							Date:                 mustDate("2025-05-13"),
							WithSalary:           false,
							Remote:               true,
						},
						{
							Title:                "Senior Backend Engineer (Elixir)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4304693089/",
							Location:             "Portugal",
							Date:                 mustDate("2025-09-29", "2025-09-23"),
							WithSalary:           false,
							Remote:               true,
						},
						{
							Title:                "Senior Backend Engineer (Elixir)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4311735011/",
							Location:             "Monaco, Monaco, Monaco",
							Date:                 mustDate("2025-10-08"),
							WithSalary:           false,
							Remote:               true,
						},
						{
							Title:                "Senior Backend Engineer (Elixir)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4311716888/",
							Location:             "Liechtenstein",
							Date:                 mustDate("2025-10-16", "2025-10-15"),
							WithSalary:           false,
							Remote:               true,
						},
						{
							Title:                "Senior Backend Engineer (Elixir)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4333101476/",
							Location:             "Lithuania",
							Date:                 mustDate("2025-11-06"),
							WithSalary:           false,
							Remote:               true,
						},
						{
							Title:                "Senior Backend Engineer (Elixir)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4333817343/",
							Location:             "Ciudad Bolivia, Barinas State, Venezuela",
							Date:                 mustDate("2025-11-12"),
							WithSalary:           false,
							Remote:               true,
						},
						{
							Title:                "Senior Backend Engineer (Elixir)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4336651004/",
							Location:             "Brazil",
							Date:                 mustDate("2025-11-20"),
							WithSalary:           false,
							Remote:               true,
						},
					},
				},
				domain.Clojure: {},
				domain.Haskell: {},
			},
			ShortDescription: "HR Platform that helps companies hire, manage, and pay their entire team",
			Industries:       []domain.Industry{
				// NOP
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Fresha",
			Website: "https://www.fresha.com/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                13230763,
				IDs:               nil,
				Alias:             "heyfresha",
				Name:              "Fresha",
				Followers:         "50K",
				Employees:         "201-500",
				AssociatedMembers: "532",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:     "surgeventures",
				Followers: "61",
				Verified:  false,
			},
			BlindProfile: domain.BlindProfile{
				Alias: "",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "fresha",
				Employees: "330",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Fresha-EI_IE3506248.11,17.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Fresha-Reviews-E3506248.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Fresha-Jobs-E3506248.htm",
				Jobs:        "5",
				Reviews:     "220",
				Salaries:    "384",
				ReviewsRate: "3.5",
				Verified:    false,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "",
			},
			Languages: domain.Languages{
				domain.Go:    {},
				domain.Rust:  {},
				domain.Zig:   {},
				domain.Scala: {},
				domain.Elixir: {
					GitHubRepositoriesCount: 53,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Senior Backend Engineer",
							ShortDescription:     "Strong understanding of languages such as Elixir, Ruby, GraphQL or TypeScript",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/3933499435/",
							Date:                 mustDate("2025-04-14"),
							WithSalary:           false,
							Remote:               false,
						},
					},
				},
				domain.Clojure: {},
				domain.Haskell: {},
			},
			ShortDescription: "Marketplace platform for the beauty, wellness, and self-care industry",
			Industries: []domain.Industry{
				domain.IndustryECommerce,
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "The RealReal",
			Website: "https://www.therealreal.com/",
			Careers: "https://careers.therealreal.com/",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                2378851,
				IDs:               nil,
				Alias:             "the-realreal",
				Name:              "The RealReal",
				Followers:         "186K",
				Employees:         "1K-5K",
				AssociatedMembers: "2,464",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:     "therealreal",
				Followers: "30",
				Verified:  true,
			},
			BlindProfile: domain.BlindProfile{
				Alias: "",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias: "",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-The-RealReal-EI_IE729017.11,23.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/The-RealReal-Reviews-E729017.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/The-RealReal-Jobs-E729017.htm",
				Jobs:        "113",
				Reviews:     "1.6K",
				Salaries:    "2.5K",
				ReviewsRate: "2.8",
				Verified:    true,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "",
			},
			Languages: domain.Languages{
				domain.Go:    {},
				domain.Rust:  {},
				domain.Zig:   {},
				domain.Scala: {},
				domain.Elixir: {
					GitHubRepositoriesCount: 21,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Staff Software Engineer",
							ShortDescription:     "",
							SwitchingOpportunity: "Experience in Ruby or Elixir",
							URL:                  "https://www.linkedin.com/jobs/view/4091569026/",
							Date:                 mustDate("2025-04-04"),
							WithSalary:           true,
							Remote:               true,
						},
					},
				},
				domain.Clojure: {},
				domain.Haskell: {},
			},
			ShortDescription: "Marketplace for authenticated, consigned luxury goods",
			Industries:       []domain.Industry{
				// NOP
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "BlackRock",
			Website: "https://www.blackrock.com/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                4764,
				IDs:               nil,
				Alias:             "blackrock",
				Name:              "BlackRock",
				Followers:         "2M",
				Employees:         "10K+",
				AssociatedMembers: "29,706",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:     "blackrock",
				Followers: "166",
				Verified:  false,
			},
			BlindProfile: domain.BlindProfile{
				Alias:       "BlackRock",
				Employees:   "10,000+",
				Salary:      "$71K ~ $250K a year",
				Reviews:     "430",
				ReviewsRate: "3.6",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "blackrock",
				Employees: "22,090",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-BlackRock-EI_IE9331.11,20.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/BlackRock-Reviews-E9331.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/BlackRock-Jobs-E9331.htm",
				Jobs:        "11",
				Reviews:     "7K",
				Salaries:    "16K",
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
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4193174525/",
							Date:                 mustDate("2025-04-16"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Golang Software Engineer, Vice President",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4194420551/",
							Date:                 mustDate("2025-05-09"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Golang Software Engineer, Associate",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4219948079/",
							Date:                 mustDate("2025-05-25"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Golang Software Engineer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4236185509/",
							Date:                 mustDate("2025-06-14"),
							WithSalary:           false,
							Remote:               false,
						},
					},
				},
				domain.Rust: {
					GitHubRepositoriesCount: 1,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Rust, AI Platform Engineer, Director",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4206654255/",
							Date:                 mustDate("2025-09-14", "2025-08-23", "2025-08-01", "2025-07-10", "2025-06-19", "2025-05-07"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Rust AI Engineer, Vice President",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4184944616/",
							Date:                 mustDate("2025-05-22"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Rust, AI Platform Engineer, Vice President",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4244046826/",
							Location:             "New York, NY",
							Date:                 mustDate("2025-09-21", "2025-08-30", "2025-07-18", "2025-06-26", "2025-06-04"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Rust, AI Platform Engineer, Director",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4280988230/",
							Location:             "New York, NY",
							Date:                 mustDate("2025-11-22", "2025-11-01", "2025-10-11", "2025-09-20", "2025-08-06"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Rust, AI Platform Engineer, Vice President",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4331174966/",
							Location:             "New York, NY",
							Date:                 mustDate("2025-12-07", "2025-10-25"),
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
							Title:                "Data Engineer – Python and Scala, Associate",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4261039927/",
							Date:                 mustDate("2025-07-02"),
							WithSalary:           false,
							Remote:               false,
						},
					},
				},
				domain.Elixir:  {},
				domain.Clojure: {},
				domain.Haskell: {},
			},
			ShortDescription: "Asset manager",
			Industries: []domain.Industry{
				domain.IndustryFinTech,
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Diagrid",
			Website: "https://www.diagrid.io/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                78416785,
				IDs:               nil,
				Alias:             "diagrid-inc",
				Name:              "Diagrid",
				Followers:         "3K",
				Employees:         "11-50",
				AssociatedMembers: "35",
				Verified:          false,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:     "diagrid-labs", // diagridio
				Followers: "53",
				Verified:  false,
			},
			BlindProfile: domain.BlindProfile{
				Alias: "",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias: "",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Diagrid-EI_IE8577874.11,18.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Diagrid-Reviews-E8577874.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Diagrid-Jobs-E8577874.htm",
				Jobs:        "",
				Reviews:     "9",
				Salaries:    "3",
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
							ShortDescription:     "3+ years of experience with Go",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4092820547/",
							Date:                 mustDate("2025-05-15", "2025-04-17"),
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
			ShortDescription: "Making distributed systems easier for developers",
			Industries: []domain.Industry{
				domain.IndustryDevOps,
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Foodji",
			Website: "https://www.foodji.com/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                18006920,
				IDs:               nil,
				Alias:             "foodji-marketplace-gmbh",
				Name:              "Foodji",
				Followers:         "7K",
				Employees:         "51-200",
				AssociatedMembers: "104",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:     "foodji",
				Followers: "7",
				Verified:  false,
			},
			BlindProfile: domain.BlindProfile{
				Alias: "",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "foodji",
				Employees: "60",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-foodji-marketplace-EI_IE4919054.11,29.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/foodji-marketplace-Reviews-E4919054.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/foodji-marketplace-Jobs-E4919054.htm",
				Jobs:        "9",
				Reviews:     "1",
				Salaries:    "12",
				ReviewsRate: "5.0",
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
							URL:                  "https://www.linkedin.com/jobs/view/4208639737/",
							Date:                 mustDate("2025-04-17"),
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
			Industries:       []domain.Industry{
				// NOP
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "inulti",
			Website: "https://inulti.com/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                18099279,
				IDs:               nil,
				Alias:             "inulti",
				Name:              "inulti",
				Followers:         "9K",
				Employees:         "11-50",
				AssociatedMembers: "45",
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
							ShortDescription:     "Backend development with hands-on expertise in PHP (Laravel/Symfony)",
							SwitchingOpportunity: "Ideally some Golang (or willingness to learn)",
							URL:                  "https://www.linkedin.com/jobs/view/4212795004/",
							Date:                 mustDate("2025-04-18"),
							WithSalary:           true, // Competitive range starting at €5000–5700 gross, with room to talk based on your expertise
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
			ShortDescription: "Insurance industry with innovative digital marketing solutions",
			Industries: []domain.Industry{
				domain.IndustryMarTech,
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "B. Braun Group",
			Website: "https://www.bbraun.com/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                6773,
				IDs:               []int{6773, 6776, 76074448},
				Alias:             "bbraun-group",
				Name:              "B. Braun Group",
				Followers:         "532K",
				Employees:         "10K+",
				AssociatedMembers: "23,031",
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
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-B-Braun-EI_IE6877.11,18.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/B-Braun-Reviews-E6877.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/B-Braun-Jobs-E6877.htm",
				Jobs:        "954",
				Reviews:     "1.7K",
				Salaries:    "2.6K",
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
							Title:                "Go Developer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4206846116/",
							Date:                 mustDate("2025-04-11"),
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
			ShortDescription: "Provider and manufacturer of healthcare solutions",
			Industries: []domain.Industry{
				domain.IndustryHealthTech,
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Muzz",
			Website: "https://muzz.com/",
			Careers: "https://muzz.com/careers/",
			About:   "",
			Blog:    "https://engineering.muzz.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                11089336,
				IDs:               nil,
				Alias:             "muzz",
				Name:              "Muzz",
				Followers:         "17K",
				Employees:         "51-200",
				AssociatedMembers: "161",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:     "muzzapp",
				Followers: "18",
				Verified:  true,
			},
			BlindProfile: domain.BlindProfile{
				Alias: "",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "muzz",
				Employees: "30",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Muzz-EI_IE3256123.11,15.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Muzz-Reviews-E3256123.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Muzz-Jobs-E3256123.htm",
				Jobs:        "19",
				Reviews:     "48",
				Salaries:    "79",
				ReviewsRate: "3.5",
				Verified:    false,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "",
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 5,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Go Engineer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4207580495/",
							Date:                 mustDate("2025-04-13"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Go Engineer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4214257562/",
							Date:                 mustDate("2025-04-26"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Go Engineer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4221277098/",
							Date:                 mustDate("2025-05-03"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Go Engineer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4236557748/",
							Date:                 mustDate("2025-05-28"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Go Engineer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4254615496/",
							Date:                 mustDate("2025-06-23"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Go Engineer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4261828971/",
							Date:                 mustDate("2025-07-07"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Go Engineer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4271802793/",
							Location:             "London Area, United Kingdom",
							Date:                 mustDate("2025-07-24"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Go Engineer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4278462191/",
							Location:             "London Area, United Kingdom",
							Date:                 mustDate("2025-08-04"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Go Engineer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4287563146/",
							Location:             "London Area, United Kingdom",
							Date:                 mustDate("2025-08-19"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Go Engineer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4294229505/",
							Location:             "London Area, United Kingdom",
							Date:                 mustDate("2025-09-02"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Go Engineer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4301978409/",
							Location:             "London Area, United Kingdom",
							Date:                 mustDate("2025-09-18"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Go Engineer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4307662741/",
							Location:             "London Area, United Kingdom",
							Date:                 mustDate("2025-10-01"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Go Engineer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4316093784/",
							Location:             "London Area, United Kingdom",
							Date:                 mustDate("2025-10-18"),
							WithSalary:           true, // £60k/yr - £80k/yr
							Remote:               false,
						},
						{
							Title:                "Go Engineer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4319259257/",
							Location:             "London Area, United Kingdom",
							Date:                 mustDate("2025-10-29"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Go Engineer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4321257129/",
							Location:             "London Area, United Kingdom",
							Date:                 mustDate("2025-11-08"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Go Engineer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4323337330/",
							Location:             "London Area, United Kingdom",
							Date:                 mustDate("2025-11-20"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Go Engineer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4326459110/",
							Location:             "London Area, United Kingdom",
							Date:                 mustDate("2025-12-10"),
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
			ShortDescription: "Muslim marriage app",
			Industries:       []domain.Industry{
				// Dating Services
			},
		},
		{
			ID:      0,                         // system
			Type:    domain.CompanyTypeStartup, // system
			Name:    "Oolio",
			Website: "https://www.oolio.com/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                82589090,
				IDs:               nil,
				Alias:             "oolio-au",
				Name:              "Oolio",
				Followers:         "8K",
				Employees:         "11-50",
				AssociatedMembers: "118",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:     "oolio-group",
				Followers: "16",
				Verified:  false,
			},
			BlindProfile: domain.BlindProfile{
				Alias: "",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias: "",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Oolio-EI_IE8473361.11,16.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Oolio-Reviews-E8473361.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Oolio-Jobs-E8473361.htm",
				Jobs:        "",
				Reviews:     "9",
				Salaries:    "24",
				ReviewsRate: "3.0",
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
							Title:                "Senior Full Stack Engineer (Golang & React.js)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4208614176/",
							Date:                 mustDate("2025-04-17"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Senior Full Stack Engineer (Golang & React.js)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4224184412/",
							Date:                 mustDate("2025-05-07"),
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
			ShortDescription: "Australian-based start-up on a mission to empower the hospitality industry",
			Industries: []domain.Industry{
				domain.IndustryMedTech,
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "QuickNode",
			Website: "https://www.quicknode.com/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                35476643,
				IDs:               nil,
				Alias:             "quicknode",
				Name:              "QuickNode",
				Followers:         "7K",
				Employees:         "51-200",
				AssociatedMembers: "117",
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
							Title:                "Senior Software Engineer (Rust/Golang)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4212330488/",
							Date:                 mustDate("2025-04-17"),
							WithSalary:           false,
							Remote:               true,
						},
					},
				},
				domain.Rust: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Senior Software Engineer (Rust/Golang)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4212330488/",
							Date:                 mustDate("2025-04-17"),
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
			Industries:       []domain.Industry{
				// NOP
			},
			Ignore: true, // Blockchain
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Submer",
			Website: "https://submer.com/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                10389939,
				IDs:               nil,
				Alias:             "submer",
				Name:              "Submer",
				Followers:         "22K",
				Employees:         "51-200",
				AssociatedMembers: "127",
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
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Submer-EI_IE3525623.11,17.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Submer-Reviews-E3525623.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Submer-Jobs-E3525623.htm",
				Jobs:        "19",
				Reviews:     "12",
				Salaries:    "18",
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
							Title:                "Backend Go Developer",
							ShortDescription:     "Kubernetes infrastructure",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4209740652/",
							Date:                 mustDate("2025-04-17"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Backend Go Developer",
							ShortDescription:     "Kubernetes infrastructure",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4214839336/",
							Date:                 mustDate("2025-04-24"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Backend Go Developer",
							ShortDescription:     "Kubernetes infrastructure",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4214476626/",
							Date:                 mustDate("2025-05-28"),
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
			ShortDescription: "Cooling and automation for data & energy-intense environments",
			Industries:       []domain.Industry{
				// NOP
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Prospect 33",
			Website: "https://prospect33.com/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                200651,
				IDs:               nil,
				Alias:             "prospect-33",
				Name:              "Prospect 33",
				Followers:         "23K",
				Employees:         "51-200",
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
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Prospect-33-EI_IE360832.11,22.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Prospect-33-Reviews-E360832.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Prospect-33-Jobs-E360832.htm",
				Jobs:        "3",
				Reviews:     "114",
				Salaries:    "80",
				ReviewsRate: "4.6",
				Verified:    false,
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
							Title:                "Rust AI/ML Engineer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4211750274/",
							Date:                 mustDate("2025-04-18"),
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
			Industries: []domain.Industry{
				domain.IndustryFinTech,
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Rockwell Automation",
			Website: "https://www.rockwellautomation.com/",
			Careers: "https://www.rockwellautomation.com/careers.html",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                2498,
				IDs:               []int{2498, 20931, 33920, 770755, 8422791, 10829528, 10831502, 10970690},
				Alias:             "rockwell-automation",
				Name:              "Rockwell Automation",
				Followers:         "1M",
				Employees:         "10K+",
				AssociatedMembers: "20,705",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:     "RockwellAutomation",
				Followers: "94",
				Verified:  true,
			},
			BlindProfile: domain.BlindProfile{
				Alias:       "Rockwell-Automation",
				Employees:   "10,000+",
				Salary:      "$56K ~ $160K a year",
				Reviews:     "50",
				ReviewsRate: "3.4",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "rockwell-automation",
				Employees: "22,270",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Rockwell-Automation-EI_IE567.11,30.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Rockwell-Automation-Reviews-E567.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Rockwell-Automation-Jobs-E567.htm",
				Jobs:        "422",
				Reviews:     "4.5K",
				Salaries:    "7.8K",
				ReviewsRate: "3.7",
				Verified:    true,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "Rockwell-Automation",
			},
			Languages: domain.Languages{
				domain.Go: {},
				domain.Rust: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Senior Software Engineer – Rust",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4212980639/",
							Location:             "Prague, Prague, Czechia",
							Date:                 mustDate("2025-04-18"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Senior Software Engineer – Rust",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4279624000/",
							Location:             "Prague, Prague, Czechia",
							Date:                 mustDate("2025-10-30", "2025-10-08", "2025-09-16", "2025-08-26", "2025-08-04"),
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
			Industries:       []domain.Industry{
				// NOP
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Wise Home",
			Website: "https://wisehome.dk/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                24789239,
				IDs:               nil,
				Alias:             "wisehome",
				Name:              "Wise Home",
				Followers:         "1K",
				Employees:         "11-50",
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
				OverviewURL: "",
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "",
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
							Title:                "Back-End Developer",
							ShortDescription:     "Our backend is written in Elixir",
							SwitchingOpportunity: "We expect no prior knowledge about Elixir",
							URL:                  "https://www.linkedin.com/jobs/view/3992422571/",
							Date:                 mustDate("2024-08-20"),
							WithSalary:           false,
							Remote:               false,
						},
					},
				},
				domain.Clojure: {},
				domain.Haskell: {},
			},
			ShortDescription: "Software for property administrators",
			Industries: []domain.Industry{
				domain.IndustryPropTech,
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Weedmaps",
			Website: "https://weedmaps.com/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                3676172,
				IDs:               nil,
				Alias:             "weedmaps",
				Name:              "Weedmaps",
				Followers:         "76K",
				Employees:         "501-1K",
				AssociatedMembers: "665",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:     "ghostgroup",
				Followers: "51",
				Verified:  true,
			},
			BlindProfile: domain.BlindProfile{
				Alias:       "Weedmaps",
				Employees:   "201 to 500",
				Salary:      "$95K ~ $200K a year",
				Reviews:     "83",
				ReviewsRate: "2.9",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "weedmaps",
				Employees: "700",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Weedmaps-EI_IE1013617.11,19.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Weedmaps-Reviews-E1013617.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Weedmaps-Jobs-E1013617.htm",
				Jobs:        "4",
				Reviews:     "308",
				Salaries:    "597",
				ReviewsRate: "2.3",
				Verified:    false,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "",
			},
			Languages: domain.Languages{
				domain.Go:    {},
				domain.Rust:  {},
				domain.Zig:   {},
				domain.Scala: {},
				domain.Elixir: {
					GitHubRepositoriesCount: 16,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Staff Software Engineer",
							ShortDescription:     "Experience with Node.js, Elixir, and Ruby",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4205786538/",
							Date:                 mustDate("2025-04-13"),
							WithSalary:           true, // $202,846.00 - $228,071.00 to per year
							Remote:               false,
						},
					},
				},
				domain.Clojure: {},
				domain.Haskell: {},
			},
			ShortDescription: "Weedmaps is the technology platform powering the cannabis industry",
			YahooFinanceURL:  "https://finance.yahoo.com/quote/MAPS/",
			GoogleFinanceURL: "https://www.google.com/finance/quote/MAPS:NASDAQ",
			Industries:       []domain.Industry{
				// NOP
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Versus Systems",
			Website: "https://www.versussystems.com/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                3587255,
				IDs:               nil,
				Alias:             "versus-gaming-network",
				Name:              "Versus Systems",
				Followers:         "1K",
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
				domain.Go:    {},
				domain.Rust:  {},
				domain.Zig:   {},
				domain.Scala: {},
				domain.Elixir: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Senior Software Engineer, Elixir",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4124935788/",
							Date:                 mustDate("2025-01-20"),
							WithSalary:           false,
							Remote:               false,
						},
					},
				},
				domain.Clojure: {},
				domain.Haskell: {},
			},
			ShortDescription: "",
			Industries: []domain.Industry{
				domain.IndustryAdTech,
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Truely Travel eSIM",
			Website: "https://truely.com/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                10844540,
				IDs:               nil,
				Alias:             "truely",
				Name:              "Truely Travel eSIM",
				Followers:         "3K",
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
				domain.Go:    {},
				domain.Rust:  {},
				domain.Zig:   {},
				domain.Scala: {},
				domain.Elixir: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Elixir Developer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4152691095/",
							Date:                 mustDate("2025-02-20"),
							WithSalary:           false,
							Remote:               true,
						},
					},
				},
				domain.Clojure: {},
				domain.Haskell: {},
			},
			ShortDescription: "Travel eSIM",
			Industries:       []domain.Industry{
				// NOP
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Spreedly",
			Website: "https://www.spreedly.com/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                1869926,
				IDs:               nil,
				Alias:             "spreedly",
				Name:              "Spreedly",
				Followers:         "10K",
				Employees:         "51-200",
				AssociatedMembers: "137",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:     "spreedly",
				Followers: "34",
				Verified:  false,
			},
			BlindProfile: domain.BlindProfile{
				Alias:       "Spreedly",
				Employees:   "51 to 200",
				Salary:      "",
				Reviews:     "5",
				ReviewsRate: "4.4",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "spreedly",
				Employees: "150",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Spreedly-EI_IE1362343.11,19.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Spreedly-Reviews-E1362343.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Spreedly-Jobs-E1362343.htm",
				Jobs:        "",
				Reviews:     "56",
				Salaries:    "114",
				ReviewsRate: "3.6",
				Verified:    true,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "",
			},
			Languages: domain.Languages{
				domain.Go:    {},
				domain.Rust:  {},
				domain.Zig:   {},
				domain.Scala: {},
				domain.Elixir: {
					GitHubRepositoriesCount: 7,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Software Engineering Manager",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4064317166/",
							Date:                 mustDate("2024-12-20"),
							WithSalary:           false,
							Remote:               true,
						},
					},
				},
				domain.Clojure: {},
				domain.Haskell: {},
			},
			ShortDescription: "Build payment systems by connecting to any payment service",
			Industries: []domain.Industry{
				domain.IndustryFinTech,
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "River",
			Website: "https://river.com/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                18932591,
				IDs:               nil,
				Alias:             "riverfinancial",
				Name:              "River",
				Followers:         "10K",
				Employees:         "51-200",
				AssociatedMembers: "101",
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
				domain.Go:    {},
				domain.Rust:  {},
				domain.Zig:   {},
				domain.Scala: {},
				domain.Elixir: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Senior Software Engineer",
							ShortDescription:     "Our web stack consists primarily of Elixir (Phoenix and LiveView) with Postgres for data storage",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4183674310/",
							Date:                 mustDate("2025-04-13"),
							WithSalary:           false,
							Remote:               true,
						},
					},
				},
				domain.Clojure: {},
				domain.Haskell: {},
			},
			ShortDescription: "",
			Industries:       []domain.Industry{
				// NOP
			},
			Ignore: true, // Cryptocurrency
		},
		{
			ID:      0,                         // system
			Type:    domain.CompanyTypeStartup, // system
			Name:    "Tab",
			Website: "https://tab.technology/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                101631732,
				IDs:               nil,
				Alias:             "tabtechnology",
				Name:              "Tab",
				Followers:         "274",
				Employees:         "2-10",
				AssociatedMembers: "6",
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
							URL:                  "https://www.linkedin.com/jobs/view/4208320140/",
							Date:                 mustDate("2025-04-20"),
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
			ShortDescription: "Tab is the API-first company that enables smart receipts to arrive straight to your mobile bank app, linked to the correct transaction",
			Industries: []domain.Industry{
				domain.IndustryFinTech,
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Rakuten Symphony",
			Website: "https://symphony.rakuten.com/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                47672239,
				IDs:               []int{2675449, 5398295, 47672239},
				Alias:             "rakuten-symphony",
				Name:              "Rakuten Symphony",
				Followers:         "200K",
				Employees:         "1K-5K",
				AssociatedMembers: "2,957",
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
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Rakuten-Symphony-EI_IE5928296.11,27.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Rakuten-Symphony-Reviews-E5928296.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Rakuten-Symphony-Jobs-E5928296.htm",
				Jobs:        "",
				Reviews:     "458",
				Salaries:    "624",
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
							Title:                "Technical Lead (Python/Golang/Rust)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4211532061/",
							Date:                 mustDate("2025-04-19"),
							WithSalary:           false,
							Remote:               false,
						},
					},
				},
				domain.Rust: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Technical Lead (Python/Golang/Rust)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4211532061/",
							Date:                 mustDate("2025-04-19"),
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
			Industries:       []domain.Industry{
				// NOP
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "airfocus",
			Website: "https://airfocus.com/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                10393984,
				IDs:               nil,
				Alias:             "airfocus",
				Name:              "airfocus by Lucid",
				Followers:         "9K",
				Employees:         "51-200",
				AssociatedMembers: "44",
				Verified:          false,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:     "airfocusio",
				Followers: "6",
				Verified:  true,
			},
			BlindProfile: domain.BlindProfile{
				Alias: "",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias: "",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-airfocus-EI_IE3247690.11,19.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/airfocus-Reviews-E3247690.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/airfocus-Jobs-E3247690.htm",
				Jobs:        "4",
				Reviews:     "40",
				Salaries:    "34",
				ReviewsRate: "4.8",
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
							Title:                "Senior Scala Engineer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4210201874/",
							Date:                 mustDate("2025-04-21"),
							WithSalary:           false,
							Remote:               false,
						},
					},
				},
				domain.Elixir:  {},
				domain.Clojure: {},
				domain.Haskell: {},
			},
			ShortDescription: "Product management platform",
			Industries:       []domain.Industry{
				// NOP
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Paradigm",
			Website: "https://www.paradigm.xyz/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                18995890,
				IDs:               nil,
				Alias:             "paradigm-xyz",
				Name:              "Paradigm",
				Followers:         "23K",
				Employees:         "51-200",
				AssociatedMembers: "116",
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
				domain.Go: {},
				domain.Rust: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Rust Engineer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4213274043/",
							Date:                 mustDate("2025-04-20"),
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
			ShortDescription: "Crypto investment firm",
			Industries:       []domain.Industry{
				// NOP
			},
			Ignore: true, // Cryptocurrency
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    ".txt",
			Website: "https://dottxt.co/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                100259483,
				IDs:               nil,
				Alias:             "dottxt",
				Name:              ".txt",
				Followers:         "5K",
				Employees:         "2-10",
				AssociatedMembers: "24",
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
				domain.Go: {},
				domain.Rust: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Staff Rust Developer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4212846266/",
							Date:                 mustDate("2025-04-22"),
							WithSalary:           false,
							Remote:               true,
						},
						{
							Title:                "Staff Rust Developer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4289500880/",
							Location:             "United States",
							Date:                 mustDate("2025-08-22"),
							WithSalary:           true, // $170k/yr - $200k/yr
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
			ShortDescription: "Making AI speak computers",
			Industries:       []domain.Industry{
				// NOP
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Omniful",
			Website: "https://www.omniful.ai/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                68537379,
				IDs:               nil,
				Alias:             "omniful",
				Name:              "Omniful",
				Followers:         "9K",
				Employees:         "51-200",
				AssociatedMembers: "94",
				Verified:          false,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:     "omniful",
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
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Omniful-EI_IE7591023.11,18.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Omniful-Reviews-E7591023.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Omniful-Jobs-E7591023.htm",
				Jobs:        "",
				Reviews:     "30",
				Salaries:    "26",
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
							Title:                "Golang Developer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4212688691/",
							Date:                 mustDate("2025-04-21"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Senior Golang Developer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4264324374/",
							Location:             "Gurugram, Haryana, India",
							Date:                 mustDate("2025-07-10"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Senior Golang Developer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4299778413/",
							Location:             "Gurugram, Haryana, India",
							Date:                 mustDate("2025-09-13"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Senior Golang Developer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4308965697/",
							Location:             "Gurugram, Haryana, India",
							Date:                 mustDate("2025-10-04"),
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
				domain.IndustryECommerce,
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Alterante, Inc.",
			Website: "https://www.alterante.com/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                3167880,
				IDs:               nil,
				Alias:             "alterante",
				Name:              "Alterante, Inc.",
				Followers:         "762",
				Employees:         "2-10",
				AssociatedMembers: "3",
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
				domain.Go:     {},
				domain.Rust:   {},
				domain.Zig:    {},
				domain.Scala:  {},
				domain.Elixir: {},
				domain.Clojure: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Senior Software Engineer (Java + Clojure)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4212268324/",
							Date:                 mustDate("2025-04-20"),
							WithSalary:           false,
							Remote:               false,
						},
					},
				},
				domain.Haskell: {},
			},
			ShortDescription: "",
			Industries:       []domain.Industry{
				// NOP
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Railsr",
			Website: "https://www.railsr.com/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                10632671,
				IDs:               nil,
				Alias:             "railsrglobal",
				Name:              "Railsr",
				Followers:         "29K",
				Employees:         "201-500",
				AssociatedMembers: "167",
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
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Railsr-EI_IE2978897.11,17.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Railsr-Reviews-E2978897.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Railsr-Jobs-E2978897.htm",
				Jobs:        "7",
				Reviews:     "74",
				Salaries:    "198",
				ReviewsRate: "3.4",
				Verified:    true,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "",
			},
			Languages: domain.Languages{
				domain.Go:     {},
				domain.Rust:   {},
				domain.Zig:    {},
				domain.Scala:  {},
				domain.Elixir: {},
				domain.Clojure: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Senior Clojure Engineer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4215346791/",
							Date:                 mustDate("2025-04-23"),
							WithSalary:           false,
							Remote:               true,
						},
						{
							Title:                "Senior Clojure Engineer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4215347539/",
							Date:                 mustDate("2025-04-23"),
							WithSalary:           false,
							Remote:               true,
						},
					},
				},
				domain.Haskell: {},
			},
			ShortDescription: "Embedded Finance Experiences Platform",
			Industries: []domain.Industry{
				domain.IndustryFinTech,
			},
		},
		{
			ID:      0,                         // system
			Type:    domain.CompanyTypeStartup, // system
			Name:    "Metric",
			Website: "https://www.metric.tech/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                91317726,
				IDs:               nil,
				Alias:             "metric-app",
				Name:              "Metric: AI Powered Marketing",
				Followers:         "615",
				Employees:         "11-50",
				AssociatedMembers: "26",
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
				domain.Go:     {},
				domain.Rust:   {},
				domain.Zig:    {},
				domain.Scala:  {},
				domain.Elixir: {},
				domain.Clojure: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Senior Clojure Engineer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4214482668/",
							Date:                 mustDate("2025-04-23"),
							WithSalary:           false,
							Remote:               true,
						},
					},
				},
				domain.Haskell: {},
			},
			ShortDescription: "AI-powered marketing agency",
			Industries: []domain.Industry{
				domain.IndustryMarTech,
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Prima Group",
			Website: "https://www.primagroup.org/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                18787171,
				IDs:               nil,
				Alias:             "prima-group-housing",
				Name:              "Prima Group",
				Followers:         "2K",
				Employees:         "51-200",
				AssociatedMembers: "72",
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
				domain.Go:    {},
				domain.Rust:  {},
				domain.Zig:   {},
				domain.Scala: {},
				domain.Elixir: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Elixir Software Engineer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4215017248/",
							Date:                 mustDate("2025-04-23"),
							WithSalary:           false,
							Remote:               false,
						},
					},
				},
				domain.Clojure: {},
				domain.Haskell: {},
			},
			ShortDescription: "",
			Industries:       []domain.Industry{
				// NOP
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Parachute Health",
			Website: "https://www.parachutehealth.com/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                3293664,
				IDs:               nil,
				Alias:             "parachute-health",
				Name:              "Parachute Health",
				Followers:         "11K",
				Employees:         "201-500",
				AssociatedMembers: "233",
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
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Parachute-Health-EI_IE2022769.11,27.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Parachute-Health-Reviews-E2022769.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Parachute-Health-Jobs-E2022769.htm",
				Jobs:        "",
				Reviews:     "30",
				Salaries:    "73",
				ReviewsRate: "4.1",
				Verified:    true,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "",
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
							Title:                "Senior / Staff Software Engineer (Elixir)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4215023789/",
							Date:                 mustDate("2025-04-23"),
							WithSalary:           false,
							Remote:               true,
						},
					},
				},
				domain.Clojure: {},
				domain.Haskell: {},
			},
			ShortDescription: "The Parachute Platform empowers healthcare providers with delightfully simple ePrescribing for medical equipment (DME), supplies, and services and powers suppliers with digital transformation tools - to streamline workflows, increase clinician satisfaction and improve patient outcomes",
			Industries: []domain.Industry{
				domain.IndustryHealthTech,
				domain.IndustryMedTech,
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Allocator One",
			Website: "https://www.allocator.one/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                98395156,
				IDs:               nil,
				Alias:             "allocator-one",
				Name:              "Allocator One",
				Followers:         "3K",
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
				domain.Go:    {},
				domain.Rust:  {},
				domain.Zig:   {},
				domain.Scala: {},
				domain.Elixir: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Senior Full Stack Developer | Elixir, Phoenix, LiveView",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4213116359/",
							Date:                 mustDate("2025-04-24"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Senior Full Stack Developer | Elixir, Phoenix, LiveView",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4234115706/",
							Date:                 mustDate("2025-05-25"),
							WithSalary:           true,
							Remote:               false,
						},
						{
							Title:                "Senior Full Stack Developer | Elixir, Phoenix, LiveView",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4303171328/",
							Location:             "Philippines",
							Date:                 mustDate("2025-09-21"),
							WithSalary:           true, // $30k/yr - $45k/yr
							Remote:               true,
						},
						{
							Title:                "Senior Full Stack Developer | Elixir, Phoenix, LiveView",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4314927536/",
							Location:             "Philippines",
							Date:                 mustDate("2025-10-22"),
							WithSalary:           true, // $30k/yr - $45k/yr
							Remote:               true,
						},
						{
							Title:                "Senior Full Stack Developer | Elixir, Phoenix, LiveView",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4341961080/",
							Location:             "Philippines",
							Date:                 mustDate("2025-11-20"),
							WithSalary:           true, // $30k/yr - $45k/yr
							Remote:               true,
						},
					},
				},
				domain.Clojure: {},
				domain.Haskell: {},
			},
			ShortDescription: "Platform for first-time venture capital funds",
			Industries: []domain.Industry{
				domain.IndustryFinTech,
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Simacan",
			Website: "https://simacan.com/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                3045752,
				IDs:               nil,
				Alias:             "simacanbv",
				Name:              "Simacan",
				Followers:         "4K",
				Employees:         "51-200",
				AssociatedMembers: "46",
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
							Title:                "Scala Backend Developer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4194786604/",
							Date:                 mustDate("2025-04-22"),
							WithSalary:           false,
							Remote:               false,
						},
					},
				},
				domain.Elixir:  {},
				domain.Clojure: {},
				domain.Haskell: {},
			},
			ShortDescription: "End-to-end supply chain platform",
			Industries: []domain.Industry{
				domain.IndustryLogisticsTech,
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "impact.com",
			Website: "https://impact.com/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                608678,
				IDs:               nil,
				Alias:             "impactdotcom",
				Name:              "impact.com",
				Followers:         "81K",
				Employees:         "1K-5K",
				AssociatedMembers: "3,073",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:     "ImpactInc",
				Followers: "53",
				Verified:  true,
			},
			BlindProfile: domain.BlindProfile{
				Alias: "",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "impactcom",
				Employees: "1,090",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-impact-com-EI_IE319986.11,21.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/impact-com-Reviews-E319986.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/impact-com-Jobs-E319986.htm",
				Jobs:        "69",
				Reviews:     "354",
				Salaries:    "886",
				ReviewsRate: "4.2",
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
					GitHubRepositoriesCount: 2,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Software Engineer I (Java, Scala)",
							ShortDescription:     "Data Platform",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4209595207/",
							Date:                 mustDate("2025-06-01", "2025-05-10", "2025-04-19"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Software Engineer I (Java, Scala)",
							ShortDescription:     "Data Platform",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4264910255/",
							Location:             "Columbus, OH",
							Date:                 mustDate("2025-07-15"),
							WithSalary:           true, // $95k/yr - $105k/yr
							Remote:               false,
						},
					},
				},
				domain.Elixir:  {},
				domain.Clojure: {},
				domain.Haskell: {},
			},
			ShortDescription: "Partnership management platform",
			Industries:       []domain.Industry{
				// NOP
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "SymphonyAI",
			Website: "https://www.symphonyai.com/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                18254848,
				IDs:               []int{51715, 10327626, 16153604, 18254848},
				Alias:             "symphonyai",
				Name:              "SymphonyAI",
				Followers:         "154K",
				Employees:         "1K-5K",
				AssociatedMembers: "2,065",
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
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-SymphonyAI-EI_IE2376728.11,21.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/SymphonyAI-Reviews-E2376728.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/SymphonyAI-Jobs-E2376728.htm",
				Jobs:        "73",
				Reviews:     "602",
				Salaries:    "842",
				ReviewsRate: "3.5",
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
							Title:                "Scala Engineer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4214511153/",
							Date:                 mustDate("2025-06-05", "2025-04-22"),
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
			Industries:       []domain.Industry{
				// NOP
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Volka Games",
			Website: "https://volkagames.com/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                18589641,
				IDs:               nil,
				Alias:             "volka",
				Name:              "Volka Games",
				Followers:         "8K",
				Employees:         "51-200",
				AssociatedMembers: "66",
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
							Title:                "Senior Frontend Developer (Rust)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4211302509/",
							Date:                 mustDate("2025-05-21", "2025-04-22"),
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
			ShortDescription: "Mobile games developer and publisher",
			Industries: []domain.Industry{
				domain.IndustryGameDev,
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Squarepoint",
			Website: "https://www.squarepoint-capital.com/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                9218046,
				IDs:               nil,
				Alias:             "squarepoint-capital",
				Name:              "Squarepoint",
				Followers:         "57K",
				Employees:         "1K-5K",
				AssociatedMembers: "1,624",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "",
				Verified: false,
			},
			BlindProfile: domain.BlindProfile{
				Alias:       "Squarepoint-Capital",
				Employees:   "501 to 1,000",
				Salary:      "$85K ~ $202K a year",
				Reviews:     "56",
				ReviewsRate: "3.8",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias: "",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Squarepoint-Capital-EI_IE1442647.11,30.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Squarepoint-Capital-Reviews-E1442647.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Squarepoint-Capital-Jobs-E1442647.htm",
				Jobs:        "15",
				Reviews:     "251",
				Salaries:    "548",
				ReviewsRate: "3.9",
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
							Title:                "Software Engineer – Rust Developer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4213316323/",
							Date:                 mustDate("2025-04-22"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Software Developer (Rust)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4231900511/",
							Date:                 mustDate("2025-06-18"),
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
			ShortDescription: "Investment management firm",
			Industries: []domain.Industry{
				domain.IndustryFinTech,
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Flow Traders",
			Website: "https://www.flowtraders.com/",
			Careers: "https://www.flowtraders.com/careers/",
			About:   "https://www.flowtraders.com/about/",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                63592,
				IDs:               nil,
				Alias:             "flow-traders",
				Name:              "Flow Traders",
				Followers:         "87K",
				Employees:         "501-1K",
				AssociatedMembers: "827",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "",
				Verified: false,
			},
			BlindProfile: domain.BlindProfile{
				Alias:       "Flow-Traders",
				Employees:   "201 to 500",
				Salary:      "",
				Reviews:     "22",
				ReviewsRate: "3.3",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "flow-traders",
				Employees: "554",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Flow-Traders-EI_IE258344.11,23.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Flow-Traders-Reviews-E258344.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Flow-Traders-Jobs-E258344.htm",
				Jobs:        "38",
				Reviews:     "230",
				Salaries:    "486",
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
							Title:                "Senior Go Engineer",
							ShortDescription:     "Digital Assets",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4318022178/",
							Location:             "Amsterdam, North Holland, Netherlands",
							Date:                 mustDate("2025-10-28"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Senior Go Engineer",
							ShortDescription:     "Digital Assets",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4324821055/",
							Location:             "Amsterdam, North Holland, Netherlands",
							Date:                 mustDate("2025-11-26"),
							WithSalary:           false,
							Remote:               false,
						},
					},
				},
				domain.Rust: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Software Engineer – Rust",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4214447561/",
							Date:                 mustDate("2025-06-05", "2025-04-23"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Software Engineer – Rust",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4287336668/",
							Location:             "Hong Kong, Hong Kong SAR",
							Date:                 mustDate("2025-11-19", "2025-10-29", "2025-08-22"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Software Engineer – Rust",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4292760005/",
							Location:             "Amsterdam, North Holland, Netherlands",
							Date:                 mustDate("2025-11-20", "2025-10-30", "2025-10-03", "2025-09-13", "2025-09-02"),
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
			ShortDescription: "Trading firm",
			Industries: []domain.Industry{
				domain.IndustryFinTech,
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "World",
			Website: "https://world.org/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                76499266,
				IDs:               nil,
				Alias:             "worldofficial",
				Name:              "World",
				Followers:         "36K",
				Employees:         "51-200",
				AssociatedMembers: "473",
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
				domain.Go: {},
				domain.Rust: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Rust Backend Engineer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4199014747/",
							Date:                 mustDate("2025-04-23"),
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
			Industries:       []domain.Industry{
				// NOP
			},
			Ignore: true, // Blockchain
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Black Duck",
			Website: "https://www.blackduck.com/",
			Careers: "https://www.blackduck.com/company/careers.html",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                18957,
				IDs:               nil,
				Alias:             "black-duck-software",
				Name:              "Black Duck",
				Followers:         "23K",
				Employees:         "5K-10K",
				AssociatedMembers: "1,284",
				Verified:          false,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:     "blackducksoftware",
				Followers: "173",
				Verified:  false,
			},
			BlindProfile: domain.BlindProfile{
				Alias: "",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias: "",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Black-Duck-Software-EI_IE10393181.11,30.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Black-Duck-Software-Reviews-E10393181.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Black-Duck-Software-Jobs-E10393181.htm",
				Jobs:        "",
				Reviews:     "17",
				Salaries:    "55",
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
							Title:                "Senior Staff Software Engineer – Java, Golang & K8s",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4282310276/",
							Location:             "Bengaluru, Karnataka, India",
							Date:                 mustDate("2025-08-09"),
							WithSalary:           false,
							Remote:               false,
						},
					},
				},
				domain.Rust: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Senior C++/Rust Software Engineer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4214495115/",
							Date:                 mustDate("2025-04-23"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Senior C++/Rust Software Engineer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4262898655/",
							Date:                 mustDate("2025-07-08"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Senior C++/Rust Software Engineer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4271585493/",
							Location:             "Belfast, Northern Ireland, United Kingdom",
							Date:                 mustDate("2025-07-24"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "C++/Rust Software Engineer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4297179843/",
							Location:             "United Kingdom",
							Date:                 mustDate("2025-09-09"),
							WithSalary:           false,
							Remote:               true,
						},
						{
							Title:                "C++/Rust Software Engineer, Staff",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4315723560/",
							Location:             "Calgary, AB",
							Date:                 mustDate("2025-11-29", "2025-11-09", "2025-10-19"),
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
			ShortDescription: "Application security solutions",
			Industries: []domain.Industry{
				domain.IndustryCyberSecurity,
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Stems Labs",
			Website: "https://www.stems.art/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                84136875,
				IDs:               nil,
				Alias:             "stemsmusic",
				Name:              "Stems Labs",
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
							Title:                "Software Engineer (Backend – Golang)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4215016329/",
							Date:                 mustDate("2025-04-24"),
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
			ShortDescription: "Stems is a social music app that lets anyone remix and share music",
			Industries: []domain.Industry{
				domain.IndustryEntertainment,
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Deutsche Telekom",
			Website: "https://www.telekom.com/",
			Careers: "https://www.telekom.com/careers",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                1593,
				IDs:               []int{1592, 1593, 1596, 6068, 7492, 207359, 238752, 241577, 244342, 260502, 303294, 717403, 999183, 1580644, 1626602, 2317339, 2540191, 2543729, 3027822, 3341259, 3542077, 3678770, 3797745, 4864118, 5096726, 8039514, 10801559, 10809618, 11764786, 15220855, 16232139, 17961408, 18589501, 19225562, 35889975, 68901287, 69211820, 79530291, 80028435, 86401257, 87388970, 88656958, 93643937, 94131659, 101832630},
				Alias:             "telekom",
				Name:              "Deutsche Telekom",
				Followers:         "269K",
				Employees:         "10K+",
				AssociatedMembers: "67,279",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:     "telekom",
				Followers: "252",
				Verified:  true,
			},
			BlindProfile: domain.BlindProfile{
				Alias: "",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "deutsche-telekom",
				Employees: "215,680",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Deutsche-Telekom-EI_IE4092.11,27.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Deutsche-Telekom-Reviews-E4092.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Deutsche-Telekom-Jobs-E4092.htm",
				Jobs:        "1.8K",
				Reviews:     "2.6K",
				Salaries:    "5.2K",
				ReviewsRate: "4.0",
				Verified:    true,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "",
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 23,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Software Engineer (Golang)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4204905697/",
							Date:                 mustDate("2025-04-24"),
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
				domain.IndustryTelecom,
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "ParshipMeet Group",
			Website: "https://www.parshipmeet.com/",
			Careers: "https://www.parshipmeet.com/careers/",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                68485243,
				IDs:               []int{16002, 64690, 1404480, 3234901, 68485243},
				Alias:             "parshipmeet",
				Name:              "ParshipMeet Group",
				Followers:         "2K",
				Employees:         "501-1K",
				AssociatedMembers: "687",
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
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Parship-Group-EI_IE469564.11,24.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Parship-Group-Reviews-E469564.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Parship-Group-Jobs-E469564.htm",
				Jobs:        "7",
				Reviews:     "16",
				Salaries:    "34",
				ReviewsRate: "4.2",
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
							Title:                "Golang Engineer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4215387718/",
							Date:                 mustDate("2025-04-24"),
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
			ShortDescription: "Dating",
			Industries:       []domain.Industry{
				// NOP
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Fast Track",
			Website: "https://www.fasttrack.ai/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                10631279,
				IDs:               nil,
				Alias:             "fast-track-crm",
				Name:              "Fast Track",
				Followers:         "10K",
				Employees:         "51-200",
				AssociatedMembers: "296",
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
							Title:                "Senior Back-End Developer (Golang)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4083032052/",
							Date:                 mustDate("2025-04-24"),
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
			Industries:       []domain.Industry{
				// NOP
			},
			Ignore: true, // iGaming
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Hark",
			Website: "https://harksys.com/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                10931594,
				IDs:               nil,
				Alias:             "harksys",
				Name:              "Hark",
				Followers:         "3K",
				Employees:         "11-50",
				AssociatedMembers: "42",
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
							Title:                "Senior Go Engineer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4216368076/",
							Date:                 mustDate("2025-04-24"),
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
			ShortDescription: "Industrial IoT and Energy Management Platform",
			Industries:       []domain.Industry{
				// NOP
			},
		},
	}
}
