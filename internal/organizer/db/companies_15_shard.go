package db

import "github.com/readytotouch/readytotouch/internal/domain"

func companies15Shard() []domain.CompanyProfile {
	return []domain.CompanyProfile{
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Vidsy",
			Website: "https://www.vidsy.co/",
			Careers: "https://jobs.lever.co/vidsy",
			About:   "https://www.vidsy.co/about",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                5144149,
				Alias:             "vidsy",
				Name:              "Vidsy",
				Followers:         "7K",
				Employees:         "51-200",
				AssociatedMembers: "205",
				Verified:          false,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:     "vidsy",
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
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Vidsy-EI_IE1285045.11,16.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Vidsy-Reviews-E1285045.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Vidsy-Jobs-E1285045.htm",
				Jobs:        "10",
				Reviews:     "44",
				Salaries:    "84",
				ReviewsRate: "4.6",
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
							Title:                "Go Software Engineer",
							ShortDescription:     "2+ years experience in software engineering",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4194783969/",
							Date:                 mustDate("2025-03-27"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Go Software Engineer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4192382841/",
							Date:                 mustDate("2025-04-15"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Go Software Engineer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4233716036/",
							Date:                 mustDate("2025-06-21"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Go Software Engineer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4259604374/",
							Date:                 mustDate("2025-07-01"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Go Software Engineer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4268582045/",
							Location:             "England, United Kingdom",
							Date:                 mustDate("2025-09-04", "2025-08-16", "2025-07-30"),
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
			ShortDescription: "Creative technology engine powering mobile video ad creation to help global brands",
			Industries: []domain.Industry{
				domain.IndustryAdTech,
			},
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "RBC",
			Website: "https://www.rbc.com/",
			Careers: "https://jobs.rbc.com/",
			About:   "https://www.rbc.com/about-rbc.html",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                1808,
				Alias:             "rbc",
				Name:              "RBC",
				Followers:         "811K",
				Employees:         "10K+",
				AssociatedMembers: "94,723",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:     "rbc",
				Followers: "36",
				Verified:  true,
			},
			BlindProfile: domain.BlindProfile{
				Alias:       "RBC",
				Employees:   "10,000+",
				Salary:      "$44K ~ $250K a year",
				Reviews:     "229",
				ReviewsRate: "3.7",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "rbc",
				Employees: "88,840",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-RBC-EI_IE3358.11,14.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/RBC-Reviews-E3358.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/RBC-Jobs-E3358.htm",
				Jobs:        "2.1K",
				Reviews:     "18K",
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
							Title:                "Senior Software Developer (Go/Java/Python)",
							ShortDescription:     "3 years experience with Golang, Java or Python",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4193989908/",
							Date:                 mustDate("2025-03-29"),
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
			ShortDescription: "Canadian multinational financial services company",
			Industries: []domain.Industry{
				domain.IndustryFinTech,
			},
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Stake",
			Website: "https://hellostake.com/",
			Careers: "https://hellostake.com/careers",
			About:   "https://hellostake.com/about",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                17955503,
				Alias:             "hello-stake",
				Name:              "Stake",
				Followers:         "14K",
				Employees:         "51-200",
				AssociatedMembers: "145",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "stake-global",
				Verified: true,
			},
			BlindProfile: domain.BlindProfile{
				Alias: "",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias: "",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Stake-EI_IE2116271.11,16.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Stake-Reviews-E2116271.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Stake-Jobs-E2116271.htm",
				Jobs:        "8",
				Reviews:     "55",
				Salaries:    "79",
				ReviewsRate: "4.3",
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
							URL:                  "https://www.linkedin.com/jobs/view/4189008372/",
							Date:                 mustDate("2025-03-27"),
							WithSalary:           false,
							Remote:               true,
						},
						{
							Title:                "Golang Engineer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4247554683/",
							Date:                 mustDate("2025-06-10"),
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
			ShortDescription: "Australian financial services company",
			Industries: []domain.Industry{
				domain.IndustryFinTech,
			},
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "DriveNets",
			Website: "https://drivenets.com/",
			Careers: "https://drivenets.com/careers/",
			About:   "https://drivenets.com/company/",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                11282464,
				Alias:             "drivenets",
				Name:              "DriveNets",
				Followers:         "12K",
				Employees:         "201-500",
				AssociatedMembers: "372",
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
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-DriveNets-EI_IE2183997.11,20.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/DriveNets-Reviews-E2183997.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/DriveNets-Jobs-E2183997.htm",
				Jobs:        "29",
				Reviews:     "47",
				Salaries:    "83",
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
							Title:                "Software Engineer Team Leader — Go (Golang)",
							ShortDescription:     "6+ years of development experience",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4195065700/",
							Date:                 mustDate("2025-06-06", "2025-04-24", "2025-04-02"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Software Engineer Backend – Go (Golang)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4226202870/",
							Location:             "Raanana, Center District, Israel",
							Date:                 mustDate("2025-08-04"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Software Engineer Backend – Go (Golang)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4283027506/",
							Location:             "Raanana, Center District, Israel",
							Date:                 mustDate("2025-11-25", "2025-09-22", "2025-09-01"),
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
			ShortDescription: "Israel software company, vendor of a network infrastructure platform",
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "IFS",
			Website: "https://www.ifs.com/",
			Careers: "https://www.ifs.com/about/careers-at-ifs",
			About:   "https://www.ifs.com/about",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                164301,
				Alias:             "ifs",
				Name:              "IFS",
				Followers:         "380K",
				Employees:         "5K-10K",
				AssociatedMembers: "8,849",
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
				Alias:     "ifs",
				Employees: "4,500",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-IFS-EI_IE215666.11,14.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/IFS-Reviews-E215666.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/IFS-Jobs-E215666.htm",
				Jobs:        "87",
				Reviews:     "2K",
				Salaries:    "2.6K",
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
							Title:                "Senior Lead Software Engineer (C#, Go)",
							ShortDescription:     "R&D / Data Services",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4195036459/",
							Date:                 mustDate("2025-03-31"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Senior Lead Software Engineer (Go)",
							ShortDescription:     "R&D / Data Services",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4214546493/",
							Date:                 mustDate("2025-04-23"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Senior Software Engineer (Go)",
							ShortDescription:     "R&D / Data Services",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4226690264/",
							Date:                 mustDate("2025-05-10"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Lead Software Engineer (Golang)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4282921125/",
							Location:             "Colombo, Western Province, Sri Lanka",
							Date:                 mustDate("2025-08-11"),
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
			ShortDescription: "Swedish Industrial AI and enterprise software company",
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Lightspeed Systems",
			Website: "https://www.lightspeedsystems.com/",
			Careers: "https://www.lightspeedsystems.com/about/careers/",
			About:   "https://www.lightspeedsystems.com/about/",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                887484,
				Alias:             "lightspeed-systems",
				Name:              "Lightspeed Systems",
				Followers:         "6K",
				Employees:         "201-500",
				AssociatedMembers: "234",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:     "Lightspeed-Systems",
				Followers: "34",
				Verified:  false,
			},
			BlindProfile: domain.BlindProfile{
				Alias: "",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "lightspeed-systems",
				Employees: "240",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Lightspeed-Systems-EI_IE263489.11,29.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Lightspeed-Systems-Reviews-E263489.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Lightspeed-Systems-Jobs-E263489.htm",
				Jobs:        "14",
				Reviews:     "99",
				Salaries:    "203",
				ReviewsRate: "3.8",
				Verified:    false,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "",
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 10,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Software Engineer (Go / Full Stack)",
							ShortDescription:     "3+ years of experience developing large scale web applications using Go",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4174137347/",
							Date:                 mustDate("2025-05-07", "2025-04-18", "2025-03-27"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Software Engineer (Golang)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4338457303/",
							Location:             "Austin, TX",
							Date:                 mustDate("2025-11-25"),
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
			ShortDescription: "Platform for schools worldwide with web filtering, student safety monitoring, classroom management, device management, and analytics solutions",
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "KNIME",
			Website: "https://www.knime.com/",
			Careers: "https://www.knime.com/careers",
			About:   "https://www.knime.com/about",
			Blog:    "https://www.knime.com/blog",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                692207,
				Alias:             "knime",
				Name:              "KNIME",
				Followers:         "62K",
				Employees:         "201-500",
				AssociatedMembers: "255",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:     "knime",
				Followers: "220",
				Verified:  false,
			},
			BlindProfile: domain.BlindProfile{
				Alias: "",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "knime",
				Employees: "150",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-KNIME-EI_IE2901810.11,16.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/KNIME-Reviews-E2901810.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/KNIME-Jobs-E2901810.htm",
				Jobs:        "6",
				Reviews:     "48",
				Salaries:    "70",
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
							Title:                "Golang Backend Software Engineer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4132287368/",
							Date:                 mustDate("2025-04-03"),
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
			ShortDescription: "Free and open-source data analytics, reporting and integration platform",
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Twitch",
			Website: "https://www.twitch.tv/",
			Careers: "https://www.twitch.tv/jobs/",
			About:   "https://www.twitch.tv/about/",
			Blog:    "https://blog.twitch.tv/en/tags/engineering/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                2320329,
				Alias:             "twitch-tv",
				Name:              "Twitch",
				Followers:         "588K",
				Employees:         "1K-5K",
				AssociatedMembers: "16,077",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:     "twitchtv",
				Followers: "249",
				Verified:  true,
			},
			BlindProfile: domain.BlindProfile{
				Alias:       "twitch",
				Employees:   "1,001 to 5,000",
				Salary:      "$58K ~ $350K a year",
				Reviews:     "342",
				ReviewsRate: "3.8",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "twitch",
				Employees: "14,260",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Twitch-EI_IE639426.11,17.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Twitch-Reviews-E639426.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Twitch-Jobs-E639426.htm",
				Jobs:        "2",
				Reviews:     "968",
				Salaries:    "3.7K",
				ReviewsRate: "3.7",
				Verified:    true,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "",
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 6,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Software Engineer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://app.welcometothejungle.com/jobs/Go06Evd5",
							Date:                 mustDate("2025-04-03"),
							WithSalary:           true, // $99.5-200k per year
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
			ShortDescription: "American video live-streaming service",
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Assembled",
			Website: "https://www.assembled.com/",
			Careers: "https://www.assembled.com/careers-at-assembled",
			About:   "https://www.assembled.com/about",
			Blog:    "https://www.assembled.com/blog",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                26571599,
				Alias:             "assembledhq",
				Name:              "Assembled",
				Followers:         "8K",
				Employees:         "51-200",
				AssociatedMembers: "141",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:     "assembledhq",
				Followers: "3",
				Verified:  true,
			},
			BlindProfile: domain.BlindProfile{
				Alias: "",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "assembled",
				Employees: "126",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Assembled-EI_IE1379368.11,20.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Assembled-Reviews-E1379368.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Assembled-Jobs-E1379368.htm",
				Jobs:        "",
				Reviews:     "27",
				Salaries:    "50",
				ReviewsRate: "4.0",
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
							Title:                "Senior Software Engineer",
							ShortDescription:     "5+ years of experience in software engineering as an individual contributor in full-stack environments",
							SwitchingOpportunity: "",
							URL:                  "https://app.welcometothejungle.com/jobs/ogHfGiQ_",
							Date:                 mustDate("2025-04-03"),
							WithSalary:           true, // $125-220k per year + stock options
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
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Rundoo",
			Website: "https://getrundoo.com/",
			Careers: "https://jobs.ashbyhq.com/rundoo",
			About:   "https://www.getrundoo.com/about",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                71237489,
				Alias:             "rundoo",
				Name:              "Rundoo",
				Followers:         "2K",
				Employees:         "11-50",
				AssociatedMembers: "40",
				Verified:          false,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "rundoo",
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
							Title:                "Senior Backend Go Engineer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://app.welcometothejungle.com/jobs/XyNov7nK",
							Date:                 mustDate("2025-04-03"),
							WithSalary:           true, // $218k per year
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
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Zelis",
			Website: "https://www.zelis.com/",
			Careers: "https://www.zelis.com/company/careers/",
			About:   "https://www.zelis.com/company/",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                10797668,
				IDs:               []int{15879, 3544572, 10548546, 10797668},
				Alias:             "zelis",
				Name:              "Zelis",
				Followers:         "74K",
				Employees:         "1K-5K",
				AssociatedMembers: "2,555",
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
				Alias:     "zelis",
				Employees: "890",
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
							Title:                "Staff Engineer – Backend Systems (C# / Rust / AWS)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4210761460/",
							Date:                 mustDate("2025-06-26", "2025-06-05", "2025-05-14"),
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
			ShortDescription: "Platform that bridges the gaps and aligns interests across payers, providers, and healthcare consumers",
			Industries: []domain.Industry{
				domain.IndustryHealthTech,
			},
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Ford Motor Company",
			Website: "https://www.ford.com/",
			Careers: "https://www.careers.ford.com/",
			About:   "https://corporate.ford.com/",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                1483,
				Alias:             "ford-motor-company",
				Name:              "Ford Motor Company",
				Followers:         "4M",
				Employees:         "10K+",
				AssociatedMembers: "140,101",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:     "ford",
				Followers: "89",
				Verified:  false,
			},
			BlindProfile: domain.BlindProfile{
				Alias:       "Ford",
				Employees:   "10,000+ Employees",
				Salary:      "$64K ~ $200K a year",
				Reviews:     "468",
				ReviewsRate: "3.6",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "ford-motor",
				Employees: "187,520",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Ford-Motor-Company-EI_IE263.11,29.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Ford-Motor-Company-Reviews-E263.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Ford-Motor-Company-Jobs-E263.htm",
				Jobs:        "873",
				Reviews:     "17K",
				Salaries:    "29K",
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
							Title:                "Senior Rust Developer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4198124231/",
							Date:                 mustDate("2025-04-03"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Staff Rust Developer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4211759935/",
							Date:                 mustDate("2025-04-17"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Rust Software Developer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4300101816/",
							Location:             "Palo Alto, CA",
							Date:                 mustDate("2025-10-05", "2025-09-14"),
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
			ShortDescription: "American multinational automobile manufacturer",
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "GitGuardian",
			Website: "https://www.gitguardian.com/",
			Careers: "https://www.gitguardian.com/careers",
			About:   "https://www.gitguardian.com/about-us",
			Blog:    "https://blog.gitguardian.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                27139915,
				Alias:             "gitguardian",
				Name:              "GitGuardian",
				Followers:         "13K",
				Employees:         "51-200",
				AssociatedMembers: "162",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:     "GitGuardian",
				Followers: "567",
				Verified:  true,
			},
			BlindProfile: domain.BlindProfile{
				Alias: "",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "gitguardian",
				Employees: "75",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-GitGuardian-EI_IE3330952.11,22.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/GitGuardian-Reviews-E3330952.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/GitGuardian-Jobs-E3330952.htm",
				Jobs:        "12",
				Reviews:     "52",
				Salaries:    "56",
				ReviewsRate: "4.3",
				Verified:    true,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "",
			},
			Languages: domain.Languages{
				domain.Go: {},
				domain.Rust: {
					GitHubRepositoriesCount: 1,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Senior Software Engineer (Back-End, Rust)",
							ShortDescription:     "3+ years of experience working in a Rust environment, contributing to products where performance and portability are key",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4196474583/",
							Date:                 mustDate("2025-04-02"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Senior Software Engineer (Back-End, Rust)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4204807609/",
							Date:                 mustDate("2025-05-24", "2025-04-10"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Senior Software Engineer (Back-End, Rust)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4269801442/",
							Location:             "Paris, Île-de-France, France",
							Date:                 mustDate("2025-07-20"),
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
			ShortDescription: "Global cybersecurity startup",
		},

		// BigTech
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Dell Technologies",
			Website: "https://www.dell.com/",
			Careers: "https://jobs.dell.com/",
			About:   "https://www.dell.com/en-us/lp/dt/who-we-are",
			Blog:    "https://www.dell.com/community/en/topics/developer-blogs",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                15088102,
				IDs:               []int{9019, 13771, 472943, 15088102},
				Alias:             "delltechnologies",
				Name:              "Dell Technologies",
				Followers:         "5M",
				Employees:         "10K+",
				AssociatedMembers: "130,367",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:     "DELL",
				Followers: "676",
				Verified:  false,
			},
			BlindProfile: domain.BlindProfile{
				Alias: "",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "dell-technologies",
				Employees: "165,000",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Dell-Technologies-EI_IE1327.11,28.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Dell-Technologies-Reviews-E1327.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Dell-Technologies-Jobs-E1327.htm",
				Jobs:        "568",
				Reviews:     "41K",
				Salaries:    "65K",
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
							Title:                "Software Principal Engineer — Kubernetes & Go",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4199734022/",
							Date:                 mustDate("2025-06-09", "2025-04-27", "2025-04-04"),
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
			ShortDescription: "American BigTech company",
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "DailyPay",
			Website: "https://www.dailypay.com/",
			Careers: "https://www.dailypay.com/careers/",
			About:   "https://www.dailypay.com/about-us/",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                6654085,
				IDs:               nil,
				Alias:             "dailypay",
				Name:              "DailyPay",
				Followers:         "24K",
				Employees:         "501-1K",
				AssociatedMembers: "927",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:     "dailypay",
				Followers: "65",
				Verified:  true,
			},
			BlindProfile: domain.BlindProfile{
				Alias: "",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "dailypay",
				Employees: "210",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-DailyPay-EI_IE1197210.11,19.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/DailyPay-Reviews-E1197210.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/DailyPay-Jobs-E1197210.htm",
				Jobs:        "18",
				Reviews:     "312",
				Salaries:    "582",
				ReviewsRate: "3.6",
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
							Title:                "Senior Software Engineer (Go/AWS)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4200514129/",
							Date:                 mustDate("2025-04-04"),
							WithSalary:           true, // £55,000—£73,000 GBP
							Remote:               false,
						},
						{
							Title:                "Senior Software Engineer (Go/AWS)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4307313689/",
							Location:             "Minneapolis, MN",
							Date:                 mustDate("2025-10-01"),
							WithSalary:           true, // $129k/yr - $200k/yr
							Remote:               true,
						},
						{
							Title:                "Senior Golang Engineer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4342191086/",
							Location:             "United States",
							Date:                 mustDate("2025-12-05"),
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
			ShortDescription: "American technology company which provides payroll services",
			Industries: []domain.Industry{
				domain.IndustryFinTech,
			},
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Shopware",
			Website: "https://www.shopware.com/",
			Careers: "https://www.shopware.com/en/jobs/",
			About:   "https://www.shopware.com/en/company/",
			Blog:    "https://www.shopware.com/en/news/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                5300057,
				IDs:               nil,
				Alias:             "shopware-ag",
				Name:              "Shopware",
				Followers:         "22K",
				Employees:         "201-500",
				AssociatedMembers: "406",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:     "shopware",
				Followers: "513",
				Verified:  false,
			},
			BlindProfile: domain.BlindProfile{
				Alias: "",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias: "",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-shopware-EI_IE1977427.11,19.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/shopware-Reviews-E1977427.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/shopware-Jobs-E1977427.htm",
				Jobs:        "",
				Reviews:     "13",
				Salaries:    "53",
				ReviewsRate: "4.0",
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
							Title:                "Senior Full Stack Developer — PHP/ Symfony/ Vue.js/ Golang",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4200842385/",
							Date:                 mustDate("2025-04-04"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Senior Full Stack Developer — PHP/ Symfony/ Vue.js/ Golang",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4205329270/",
							Date:                 mustDate("2025-04-10"),
							WithSalary:           false,
							Remote:               true,
						},
						{
							Title:                "Principal Fullstack Engineer – PHP Symfony, Go and Vue.js",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4284091923/",
							Location:             "Schöppingen, North Rhine-Westphalia, Germany",
							Date:                 mustDate("2025-10-08"),
							WithSalary:           false,
							Remote:               true,
						},
						{
							Title:                "(Senior) Full Stack Developer – PHP Symfony, Go and Vue.js",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4336833344/",
							Location:             "Schöppingen, North Rhine-Westphalia, Germany",
							Date:                 mustDate("2025-11-19"),
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
			ShortDescription: "Written in PHP open source e-commerce software",
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "StarHub",
			Website: "https://www.starhub.com/",
			Careers: "https://careers.starhub.com/",
			About:   "https://corporate.starhub.com/about-us.html",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                166131,
				IDs:               nil,
				Alias:             "starhub",
				Name:              "StarHub",
				Followers:         "63K",
				Employees:         "1K-5K",
				AssociatedMembers: "2,727",
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
							Title:                "Application Developer (Golang)",
							ShortDescription:     "3+ years of professional experience with Go (Golang)",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4195690505/",
							Date:                 mustDate("2025-04-04"),
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
			ShortDescription: "Singaporean multinational telecommunications conglomerate",
		},

		// BigTech
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Rakuten",
			Website: "https://global.rakuten.com/",
			Careers: "https://global.rakuten.com/corp/careers/",
			About:   "https://global.rakuten.com/corp/about/",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                47584,
				IDs:               []int{47584, 3729600, 8961866, 14556045, 24772538, 75640232, 102070393},
				Alias:             "rakuten",
				Name:              "Rakuten",
				Followers:         "305K",
				Employees:         "10K+",
				AssociatedMembers: "10,339",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:     "rakutentech",
				Followers: "179",
				Verified:  false,
			},
			BlindProfile: domain.BlindProfile{
				Alias: "",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "rakuten",
				Employees: "340",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Rakuten-EI_IE40197.11,18.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Rakuten-Reviews-E40197.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Rakuten-Jobs-E40197.htm",
				Jobs:        "634",
				Reviews:     "4K",
				Salaries:    "5.9K",
				ReviewsRate: "3.8",
				Verified:    true,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "",
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 39,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Golang and C++ Developer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4248672779/",
							Location:             "Tokyo, Japan",
							Date:                 mustDate("2025-07-23", "2025-07-02"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "GCP/Go Software Engineer (Backend)",
							ShortDescription:     "Ad Platform",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4298156404/",
							Location:             "Tokyo, Japan",
							Date:                 mustDate("2025-11-12", "2025-10-22", "2025-09-10"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Golang and C++ Developer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4312083446/",
							Location:             "Tokyo, Japan",
							Date:                 mustDate("2025-10-08"),
							WithSalary:           false,
							Remote:               false,
						},
					},
				},
				domain.Rust: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Rust Developer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4199352625/",
							Date:                 mustDate("2025-06-08", "2025-04-25", "2025-04-04"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Rust developer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4231523156/",
							Location:             "Munich, Bavaria, Germany",
							Date:                 mustDate("2025-09-27", "2025-09-05", "2025-07-23", "2025-07-02", "2025-06-12"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Rust developer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4248673738/",
							Location:             "Munich, Bavaria, Germany",
							Date:                 mustDate("2025-10-22", "2025-10-01", "2025-09-10", "2025-08-15"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Rust developer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4299640810/",
							Location:             "Munich, Bavaria, Germany",
							Date:                 mustDate("2025-11-20", "2025-10-30"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Rust developer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4299650626/",
							Location:             "Munich, Bavaria, Germany",
							Date:                 mustDate("2025-11-20"),
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
			ShortDescription: "Japanese technology conglomerate",
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Vital Bio",
			Website: "https://vitalbio.com/",
			Careers: "https://vitalbio.com/careers",
			About:   "https://vitalbio.com/company",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                19033427,
				IDs:               nil,
				Alias:             "vitalbio",
				Name:              "Vital Bio",
				Followers:         "8K",
				Employees:         "51-200",
				AssociatedMembers: "160",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:     "VitalBio",
				Followers: "4",
				Verified:  true,
			},
			BlindProfile: domain.BlindProfile{
				Alias: "",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "vital-biosciences",
				Employees: "76",
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
					GitHubRepositoriesCount: 23,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Senior Rust Application Engineer",
							ShortDescription:     "Strong proficiency in Rust programming with experience developing production-grade applications",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4199869768/",
							Location:             "Oakville, ON",
							Date:                 mustDate("2025-04-03"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Senior Rust Application Engineer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4283951213/",
							Location:             "Oakville, ON",
							Date:                 mustDate("2025-11-04"),
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
			ShortDescription: "Healthcare technology company who provide designing a new generation of tools to empower patients in monitoring health and managing the disease",
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Stord",
			Website: "https://www.stord.com/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                10239521,
				IDs:               nil,
				Alias:             "stord",
				Name:              "Stord",
				Followers:         "24K",
				Employees:         "1K-5K",
				AssociatedMembers: "486",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:     "stordco",
				Followers: "30",
				Verified:  false,
			},
			BlindProfile: domain.BlindProfile{
				Alias: "",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias: "",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Stord-EI_IE1954757.11,16.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Stord-Reviews-E1954757.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Stord-Jobs-E1954757.htm",
				Jobs:        "",
				Reviews:     "290",
				Salaries:    "458",
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
					GitHubRepositoriesCount: 9,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Senior Software Engineer – Elixir",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4199652107/",
							Date:                 mustDate("2025-04-04"),
							WithSalary:           false,
							Remote:               true,
						},
						{
							Title:                "Senior Software Engineer – Elixir",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4295396868/",
							Location:             "Latin America",
							Date:                 mustDate("2025-09-05"),
							WithSalary:           false,
							Remote:               true,
						},
						{
							Title:                "Senior Software Engineer – Elixir",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4314089339/",
							Location:             "Latin America",
							Date:                 mustDate("2025-10-14"),
							WithSalary:           false,
							Remote:               true,
						},
						{
							Title:                "Senior Software Engineer – Elixir",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4321911327/",
							Location:             "Atlanta, GA",
							Date:                 mustDate("2025-11-11"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Senior Software Engineer – Elixir",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4347366946/",
							Location:             "United States",
							Date:                 mustDate("2025-11-26"),
							WithSalary:           false,
							Remote:               true,
						},
					},
				},
				domain.Clojure: {},
				domain.Haskell: {},
			},
			ShortDescription: "Commerce enablement platform that powers seamless checkout and delivery experiences for brands across all channels",
			Industries: []domain.Industry{
				domain.IndustryECommerce,
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Luminovo",
			Website: "https://luminovo.com/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                18362796,
				IDs:               nil,
				Alias:             "luminovo",
				Name:              "Luminovo",
				Followers:         "10K",
				Employees:         "51-200",
				AssociatedMembers: "69",
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
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Luminovo-EI_IE2039308.11,19.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Luminovo-Reviews-E2039308.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Luminovo-Jobs-E2039308.htm",
				Jobs:        "16",
				Reviews:     "55",
				Salaries:    "28",
				ReviewsRate: "4.9",
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
							Title:                "Software Engineer – Rust & React",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4200736486/",
							Date:                 mustDate("2025-04-04"),
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
			ShortDescription: "Electronics supply chain platform",
			Industries:       []domain.Industry{
				//
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Thought Machine",
			Website: "https://www.thoughtmachine.net/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                5265217,
				IDs:               nil,
				Alias:             "thought-machine",
				Name:              "Thought Machine",
				Followers:         "69K",
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
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Thought-Machine-EI_IE1637848.11,26.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Thought-Machine-Reviews-E1637848.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Thought-Machine-Jobs-E1637848.htm",
				Jobs:        "9",
				Reviews:     "233",
				Salaries:    "610",
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
							Title:                "Senior Software Engineer",
							ShortDescription:     "Experience with either Python or Golang",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4101150703/",
							Date:                 mustDate("2025-03-31"),
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
			ShortDescription: "Cloud-native banking",
			Industries: []domain.Industry{
				domain.IndustryFinTech,
			},
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Kantar",
			Website: "https://www.kantar.com/",
			Careers: "https://www.kantar.com/careers",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                83111,
				IDs:               []int{12722, 14524, 16297, 16308, 16947, 20991, 40733, 43790, 62257, 83111, 116924, 164738, 166509, 654511, 766695, 793257, 797397, 3607797},
				Alias:             "kantar",
				Name:              "Kantar",
				Followers:         "909K",
				Employees:         "10K+",
				AssociatedMembers: "28,968",
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
							Title:                "Platform Engineer/Golang Engineer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4144629838/",
							Date:                 mustDate("2025-05-04", "2025-04-11"),
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
			ShortDescription: "Marketing data and analytics company",
			Industries:       []domain.Industry{
				// NOP
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Tubia",
			Website: "https://tubia.com/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                101248841,
				IDs:               nil,
				Alias:             "tubia",
				Name:              "Tubia",
				Followers:         "1K",
				Employees:         "51-200",
				AssociatedMembers: "17",
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
							URL:                  "https://www.linkedin.com/jobs/view/4206003029/",
							Date:                 mustDate("2025-04-11"),
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
			ShortDescription: "Tubia is a platform that leverages a robust collection of over 19,000 games to enhance user engagement, attract new visitors, and increase profits for portals, websites, and media publishers",
			Industries: []domain.Industry{
				domain.IndustryEntertainment,
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Verkada",
			Website: "https://www.verkada.com/",
			Careers: "https://www.verkada.com/careers/",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                12699415,
				IDs:               nil,
				Alias:             "verkada",
				Name:              "Verkada",
				Followers:         "165K",
				Employees:         "1K-5K",
				AssociatedMembers: "2,421",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:     "verkada",
				Followers: "91",
				Verified:  true,
			},
			BlindProfile: domain.BlindProfile{
				Alias:       "Verkada",
				Employees:   "501 to 1,000",
				Salary:      "$65K ~ $250K a year",
				Reviews:     "192",
				ReviewsRate: "3.7",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "verkada",
				Employees: "1,500",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Verkada-EI_IE2153775.11,18.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Verkada-Reviews-E2153775.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Verkada-Jobs-E2153775.htm",
				Jobs:        "148",
				Reviews:     "993",
				Salaries:    "2K",
				ReviewsRate: "3.5",
				Verified:    true,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "Verkada",
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 13,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Go Software Engineer",
							ShortDescription:     "Develop and maintain Go firmware for embedded devices with focus on performance and security",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4205197888/",
							Date:                 mustDate("2025-04-10"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Go Software Engineer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4227134333/",
							Location:             "Krakowski, Małopolskie, Poland",
							Date:                 mustDate("2025-10-07", "2025-09-15", "2025-08-24", "2025-08-04", "2025-07-13", "2025-06-22", "2025-05-10"),
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
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Synthflow AI",
			Website: "https://synthflow.ai/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                98202322,
				IDs:               nil,
				Alias:             "synthflowai",
				Name:              "Synthflow AI",
				Followers:         "7K",
				Employees:         "11-50",
				AssociatedMembers: "43",
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
							Title:                "Senior Go Software Engineer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4205747084/",
							Date:                 mustDate("2025-04-10"),
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
			ShortDescription: "AI Agents",
			Industries:       []domain.Industry{
				// NOP
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "FYST",
			Website: "https://fyst.com/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                82992816,
				IDs:               nil,
				Alias:             "we-are-fyst",
				Name:              "FYST",
				Followers:         "28K",
				Employees:         "51-200",
				AssociatedMembers: "73",
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
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-FYST-EI_IE9261042.11,15.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/FYST-Reviews-E9261042.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/FYST-Jobs-E9261042.htm",
				Jobs:        "20",
				Reviews:     "4",
				Salaries:    "3",
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
							Title:                "Senior Go Developer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4201048109/",
							Date:                 mustDate("2025-04-05"),
							WithSalary:           false,
							Remote:               true,
						},
						{
							Title:                "Senior Go Developer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4271017879/",
							Location:             "Ukraine",
							Date:                 mustDate("2025-07-21"),
							WithSalary:           false,
							Remote:               true,
						},
						{
							Title:                "Senior Go Developer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4334616271/",
							Location:             "Lithuania",
							Date:                 mustDate("2025-11-06"),
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
			ShortDescription: "Payment solution",
			Industries: []domain.Industry{
				domain.IndustryFinTech,
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "letgo",
			Website: "https://www.letgo.com/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                10159884,
				IDs:               nil,
				Alias:             "letgo",
				Name:              "letgo",
				Followers:         "47K",
				Employees:         "201-500",
				AssociatedMembers: "306",
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
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-letgo-EI_IE1337214.11,16.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/letgo-Reviews-E1337214.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/letgo-Jobs-E1337214.htm",
				Jobs:        "",
				Reviews:     "109",
				Salaries:    "176",
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
							Title:                "Senior Go Engineer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4204378000/",
							Date:                 mustDate("2025-05-03", "2025-04-23", "2025-04-08"),
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
			ShortDescription: "Platform for buying and selling second-hand items",
			Industries: []domain.Industry{
				domain.IndustryECommerce,
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Daftcode",
			Website: "https://daftcode.pl/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                5055428,
				IDs:               nil,
				Alias:             "daftcode-sp-z-o-o-",
				Name:              "Daftcode",
				Followers:         "8K",
				Employees:         "51-200",
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
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-DaftCode-EI_IE1788865.11,19.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/DaftCode-Reviews-E1788865.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/DaftCode-Jobs-E1788865.htm",
				Jobs:        "",
				Reviews:     "15",
				Salaries:    "40",
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
							Title:                "Middle/Senior Golang Developer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4204597545/",
							Date:                 mustDate("2025-04-10"),
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
			ShortDescription: "Polish venture builder company",
			Industries:       []domain.Industry{
				// NOP
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Scorebuddy",
			Website: "https://www.scorebuddyqa.com/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                2587098,
				IDs:               nil,
				Alias:             "scorebuddy",
				Name:              "Scorebuddy",
				Followers:         "5K",
				Employees:         "51-200",
				AssociatedMembers: "48",
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
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Scorebuddy-EI_IE1098874.11,21.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Scorebuddy-Reviews-E1098874.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Scorebuddy-Jobs-E1098874.htm",
				Jobs:        "",
				Reviews:     "7",
				Salaries:    "11",
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
							Title:                "Senior Golang Developer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4205812062/",
							Date:                 mustDate("2025-04-11"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Senior Golang Developer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4251934373/",
							Location:             "Dublin, County Dublin, Ireland",
							Date:                 mustDate("2025-08-20", "2025-07-28", "2025-07-21", "2025-07-08"),
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
			ShortDescription: "Contact center quality assurance (QA) solution",
			Industries:       []domain.Industry{
				// NOP
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Criteo",
			Website: "https://www.criteo.com/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                37209,
				IDs:               []int{37209, 206709, 3035209, 5041701},
				Alias:             "criteo",
				Name:              "Criteo",
				Followers:         "235K",
				Employees:         "1K-5K",
				AssociatedMembers: "3,711",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:     "criteo",
				Followers: "104",
				Verified:  true,
			},
			BlindProfile: domain.BlindProfile{
				Alias: "",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "criteo",
				Employees: "3,500",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Criteo-EI_IE426672.11,17.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Criteo-Reviews-E426672.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Criteo-Jobs-E426672.htm",
				Jobs:        "182",
				Reviews:     "1.9K",
				Salaries:    "3.4K",
				ReviewsRate: "4.1",
				Verified:    true,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "Criteo",
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 29,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Software Development Engineer – Lua/Go",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4119745114/",
							Date:                 mustDate("2025-04-12"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Software Development Engineer – Lua/Go",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4119740775/",
							Date:                 mustDate("2025-04-14"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Junior Software Development Engineer – Lua/Go",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4211738171/",
							Date:                 mustDate("2025-04-17"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Senior Backend Software Engineer – Lua/Go",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4253644830/",
							Date:                 mustDate("2025-06-21"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Junior Back-End Developer – Lua/Go",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4278085466/",
							Location:             "Limassol Municipality, Limassol, Cyprus",
							Date:                 mustDate("2025-07-30"),
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
				domain.IndustryAdTech,
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "uDelta",
			Website: "https://udelta.io/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                101896582,
				IDs:               nil,
				Alias:             "udelta",
				Name:              "uDelta",
				Followers:         "7K",
				Employees:         "11-50",
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
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Senior Go Developer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4202973976/",
							Date:                 mustDate("2025-04-08"),
							WithSalary:           false,
							Remote:               true,
						},
						{
							Title:                "Go Developer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4214011730/",
							Date:                 mustDate("2025-04-21"),
							WithSalary:           false,
							Remote:               true,
						},
					},
				},
				domain.Rust: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Senior Software Engineer (Rust)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4206237641/",
							Date:                 mustDate("2025-04-11"),
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
			ShortDescription: "Platform for compensation management",
			Industries:       []domain.Industry{
				// NOP
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Monil",
			Website: "https://monil.co.uk/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                86305917,
				IDs:               nil,
				Alias:             "monil-virtual-fences",
				Name:              "Monil",
				Followers:         "1K",
				Employees:         "11-50",
				AssociatedMembers: "29",
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
							Title:                "Senior Backend Developer (Go)",
							ShortDescription:     "Powering the Future of Farming 🚜🐄",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4204790743/",
							Date:                 mustDate("2025-04-10"),
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
			ShortDescription: "We make virtual fences for easier livestock care, managed grazing, and effortless fencing",
			Industries: []domain.Industry{
				domain.IndustryAgroTech,
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Disney Entertainment",
			Website: "https://www.dgepress.com/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                1299,
				IDs:               nil,
				Alias:             "disney-entertainment",
				Name:              "Disney Entertainment",
				Followers:         "483K",
				Employees:         "5K-10K",
				AssociatedMembers: "4,916",
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
							Title:                "Software Engineer II – C++/Rust",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4207804796/",
							Date:                 mustDate("2025-04-12"),
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
							Title:                "Senior Software Engineer – Scala",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4212933068/",
							Date:                 mustDate("2025-04-18"),
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
			Industries: []domain.Industry{
				domain.IndustryEntertainment,
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Collins Aerospace",
			Website: "https://www.collinsaerospace.com/",
			Careers: "https://www.collinsaerospace.com/careers/engineering",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                11695727,
				IDs:               []int{3959, 2669459, 11695727},
				Alias:             "collins-aerospace",
				Name:              "Collins Aerospace",
				Followers:         "957K",
				Employees:         "10K+",
				AssociatedMembers: "53,149",
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
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Collins-Aerospace-EI_IE2375726.11,28.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Collins-Aerospace-Reviews-E2375726.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Collins-Aerospace-Jobs-E2375726.htm",
				Jobs:        "776",
				Reviews:     "5.3K",
				Salaries:    "11K",
				ReviewsRate: "3.8",
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
							Title:                "Rust Software Engineer II",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4208524792/",
							Date:                 mustDate("2025-04-12"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Rust Senior Principal Software Engineer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4208228319/",
							Date:                 mustDate("2025-05-04"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Rust Principal Software Engineer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4208524788/",
							Date:                 mustDate("2025-06-17", "2025-05-26"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Rust Senior Software Engineer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4208530387/",
							Date:                 mustDate("2025-06-14"),
							WithSalary:           true,
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
			Name:    "Prima",
			Website: "https://www.helloprima.com/",
			Careers: "https://www.helloprima.com/careers",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                9305779,
				IDs:               nil,
				Alias:             "prima-assicurazioni",
				Name:              "Prima",
				Followers:         "78K",
				Employees:         "1K-5K",
				AssociatedMembers: "1,680",
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
							Title:                "Rust Software Engineer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4204211704/",
							Location:             "London, England, United Kingdom",
							Date:                 mustDate("2025-12-02", "2025-11-10", "2025-09-06", "2025-06-13", "2025-05-22", "2025-05-01", "2025-04-09"),
							WithSalary:           false,
							Remote:               true,
						},
						{
							Title:                "Rust Software Engineer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4204214285/",
							Location:             "Madrid, Community of Madrid, Spain",
							Date:                 mustDate("2025-12-04", "2025-11-12", "2025-10-22", "2025-09-30", "2025-09-08", "2025-08-17", "2025-07-26", "2025-07-04", "2025-06-14"),
							WithSalary:           false,
							Remote:               true,
						},
						{
							Title:                "Rust Software Engineer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4204213541/",
							Location:             "Milan, Lombardy, Italy",
							Date:                 mustDate("2025-11-29", "2025-09-27", "2025-09-06"),
							WithSalary:           false,
							Remote:               true,
						},
					},
				},
				domain.Zig:   {},
				domain.Scala: {},
				domain.Elixir: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Elixir Software Engineer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4203425510/",
							Location:             "Milan, Lombardy, Italy",
							Date:                 mustDate("2025-07-23", "2025-07-02", "2025-04-08"),
							WithSalary:           false,
							Remote:               true,
						},
						{
							Title:                "Elixir Software Engineer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4203423628/",
							Date:                 mustDate("2025-05-22", "2025-04-29"),
							WithSalary:           false,
							Remote:               true,
						},
						{
							Title:                "Elixir Software Engineer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4203428166/",
							Location:             "London, England, United Kingdom",
							Date:                 mustDate("2025-11-29", "2025-11-09", "2025-10-19", "2025-09-27", "2025-09-06", "2025-08-15", "2025-07-24", "2025-07-03", "2025-06-13"),
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
				domain.IndustryInsurTech,
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Wasmer",
			Website: "https://wasmer.io/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                18946055,
				IDs:               nil,
				Alias:             "wasmerio",
				Name:              "Wasmer",
				Followers:         "1K",
				Employees:         "2-10",
				AssociatedMembers: "9",
				Verified:          false,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:     "wasmerio",
				Followers: "824",
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
				domain.Go: {},
				domain.Rust: {
					GitHubRepositoriesCount: 61,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Rust – Software Engineer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4204685114/",
							Date:                 mustDate("2025-04-21", "2025-04-11"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Compiler Engineer (Rust)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4222224974/",
							Date:                 mustDate("2025-05-07"),
							WithSalary:           false,
							Remote:               true,
						},
						{
							Title:                "Rust Engineer",
							ShortDescription:     "Distributed Systems",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4259391621/",
							Date:                 mustDate("2025-07-07"),
							WithSalary:           false,
							Remote:               true,
						},
						{
							Title:                "Rust Engineer, Distributed Systems",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4280216814/",
							Location:             "Spain",
							Date:                 mustDate("2025-08-09"),
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
				domain.IndustryCloudComputing,
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Takealot",
			Website: "https://www.takealot.com/",
			Careers: "https://www.takealot.com/careers",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                1167110,
				IDs:               nil,
				Alias:             "takealot",
				Name:              "takealot.com",
				Followers:         "240K",
				Employees:         "1K-5K",
				AssociatedMembers: "6,366",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:     "takealot",
				Followers: "59",
				Verified:  false,
			},
			BlindProfile: domain.BlindProfile{
				Alias: "",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias: "",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-takealot-com-EI_IE1002703.11,23.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/takealot-com-Reviews-E1002703.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/takealot-com-Jobs-E1002703.htm",
				Jobs:        "21",
				Reviews:     "386",
				Salaries:    "553",
				ReviewsRate: "3.4",
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
							Title:                "Software Engineer (Scala)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4207708792/",
							Date:                 mustDate("2025-04-12"),
							WithSalary:           false,
							Remote:               true,
						},
						{
							Title:                "Software Engineer (Scala)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4248316240/",
							Date:                 mustDate("2025-06-11"),
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
			Industries: []domain.Industry{
				domain.IndustryECommerce,
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Maropost",
			Website: "https://www.maropost.com/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                386756,
				IDs:               []int{386756, 390673},
				Alias:             "maropost",
				Name:              "Maropost",
				Followers:         "25K",
				Employees:         "201-500",
				AssociatedMembers: "356",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:     "marosolutions",
				Followers: "3",
				Verified:  true,
			},
			BlindProfile: domain.BlindProfile{
				Alias: "",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "maropost",
				Employees: "300",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Maropost-EI_IE994532.11,19.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Maropost-Reviews-E994532.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Maropost-Jobs-E994532.htm",
				Jobs:        "8",
				Reviews:     "348",
				Salaries:    "347",
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
							Title:                "Software Engineer (Golang)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4260914833/",
							Location:             "Sahibzada Ajit Singh Nagar, Punjab, India",
							Date:                 mustDate("2025-09-10", "2025-08-17", "2025-07-26", "2025-07-04"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Backend Engineer (Golang)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4309796286/",
							Location:             "Sahibzada Ajit Singh Nagar, Punjab, India",
							Date:                 mustDate("2025-12-09", "2025-11-26", "2025-10-07"),
							WithSalary:           false,
							Remote:               false,
						},
					},
				},
				domain.Rust: {},
				domain.Zig:  {},
				domain.Scala: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Software Engineer (Scala)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4207470731/",
							Date:                 mustDate("2025-07-02", "2025-06-24", "2025-06-18", "2025-04-12"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Software Engineer (Scala)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4265322755/",
							Location:             "Stockholm, Stockholm County, Sweden",
							Date:                 mustDate("2025-11-06", "2025-10-12", "2025-09-10", "2025-08-02", "2025-07-12"),
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
			Industries: []domain.Industry{
				domain.IndustryECommerce,
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Corsearch",
			Website: "https://corsearch.com/",
			Careers: "https://careers.corsearch.com/jobs",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                2593860,
				IDs:               []int{2345334, 2593860, 3252922, 13196073, 15225899},
				Alias:             "corsearchinc",
				Name:              "Corsearch",
				Followers:         "35K",
				Employees:         "1K-5K",
				AssociatedMembers: "1,642",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:     "corsearch",
				Followers: "37",
				Verified:  true,
			},
			BlindProfile: domain.BlindProfile{
				Alias: "",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias: "",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Corsearch-EI_IE2990497.11,20.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Corsearch-Reviews-E2990497.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Corsearch-Jobs-E2990497.htm",
				Jobs:        "36",
				Reviews:     "387",
				Salaries:    "480",
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
							Title:                "Software Engineer – Golang",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4260462701/",
							Location:             "Mumbai Metropolitan Region",
							Date:                 mustDate("2025-07-23", "2025-07-02"),
							WithSalary:           false,
							Remote:               false,
						},
					},
				},
				domain.Rust: {},
				domain.Zig:  {},
				domain.Scala: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Senior Scala Developer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4160578820/",
							Date:                 mustDate("2025-05-25", "2025-04-11"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Senior Scala Developer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4200184252/",
							Date:                 mustDate("2025-05-18", "2025-05-02", "2025-04-25"),
							WithSalary:           false,
							Remote:               false,
						},
					},
				},
				domain.Elixir:  {},
				domain.Clojure: {},
				domain.Haskell: {},
			},
			ShortDescription: "Intelligent Trademark & Brand Protection Solutions",
			Industries:       []domain.Industry{
				// NOP
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Peerspace",
			Website: "https://www.peerspace.com/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                3542206,
				IDs:               nil,
				Alias:             "peerspace",
				Name:              "Peerspace",
				Followers:         "10K",
				Employees:         "51-200",
				AssociatedMembers: "105",
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
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Peerspace-EI_IE1304957.11,20.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Peerspace-Reviews-E1304957.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Peerspace-Jobs-E1304957.htm",
				Jobs:        "5",
				Reviews:     "47",
				Salaries:    "57",
				ReviewsRate: "4.4",
				Verified:    false,
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
							Title:                "Senior Clojure Back-End Engineer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4204502807/",
							Date:                 mustDate("2025-04-28", "2025-04-23", "2025-04-15", "2025-04-12"),
							WithSalary:           false,
							Remote:               true,
						},
						{
							Title:                "Senior Full Stack Engineer, Clojure",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4324540820/",
							Location:             "United States",
							Date:                 mustDate("2025-12-11"),
							WithSalary:           true, // $170k/yr - $180k/yr
							Remote:               true,
						},
					},
				},
				domain.Haskell: {},
			},
			ShortDescription: "Marketplace for hourly venue rentals for meetings, productions, and events",
			Industries:       []domain.Industry{
				// NOP
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "InCharge Energy",
			Website: "https://inchargeus.com/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                19130345,
				IDs:               nil,
				Alias:             "incharge-energy",
				Name:              "InCharge Energy",
				Followers:         "16K",
				Employees:         "201-500",
				AssociatedMembers: "209",
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
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-InCharge-Energy-EI_IE3859105.11,26.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/InCharge-Energy-Reviews-E3859105.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/InCharge-Energy-Jobs-E3859105.htm",
				Jobs:        "10",
				Reviews:     "33",
				Salaries:    "47",
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
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Senior Elixir Engineer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4187189707/",
							Date:                 mustDate("2025-04-06"),
							WithSalary:           false,
							Remote:               true,
						},
						{
							Title:                "Senior Elixir Engineer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4208458029/",
							Date:                 mustDate("2025-04-14"),
							WithSalary:           false,
							Remote:               true,
						},
						{
							Title:                "Senior Elixir Engineer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4222875157/",
							Date:                 mustDate("2025-05-06"),
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
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "BEAT81",
			Website: "https://www.beat81.com/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                18363575,
				IDs:               nil,
				Alias:             "beat81",
				Name:              "BEAT81",
				Followers:         "6K",
				Employees:         "51-200",
				AssociatedMembers: "251",
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
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-BEAT81-EI_IE2855598.11,17.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/BEAT81-Reviews-E2855598.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/BEAT81-Jobs-E2855598.htm",
				Jobs:        "",
				Reviews:     "32",
				Salaries:    "59",
				ReviewsRate: "3.2",
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
							Title:                "Senior Backend Engineer – Elixir",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4204229878/",
							Date:                 mustDate("2025-04-09"),
							WithSalary:           false,
							Remote:               false,
						},
					},
				},
				domain.Clojure: {},
				domain.Haskell: {},
			},
			ShortDescription: "HIIT training",
			Industries:       []domain.Industry{
				// NOP
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Cars Commerce",
			Website: "https://www.carscommerce.inc/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                5364,
				IDs:               nil,
				Alias:             "cars-commerce",
				Name:              "Cars Commerce",
				Followers:         "42K",
				Employees:         "1K-5K",
				AssociatedMembers: "2,272",
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
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Cars-Commerce-EI_IE9365375.11,24.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Cars-Commerce-Reviews-E9365375.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Cars-Commerce-Jobs-E9365375.htm",
				Jobs:        "",
				Reviews:     "30",
				Salaries:    "62",
				ReviewsRate: "3.0",
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
							Title:                "Lead Senior Elixir Software Engineer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4204059007/",
							Date:                 mustDate("2025-05-09", "2025-04-16", "2025-04-11"),
							WithSalary:           false,
							Remote:               true,
						},
						{
							Title:                "Senior Software Engineer – Elixir",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4297469027/",
							Location:             "United States",
							Date:                 mustDate("2025-10-04", "2025-09-11"),
							WithSalary:           true, // $118.6k/yr - $148.2k/yr
							Remote:               true,
						},
						{
							Title:                "Software Engineer – Elixir",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4311912380/",
							Location:             "United States",
							Date:                 mustDate("2025-10-15"),
							WithSalary:           true, // $104.1k/yr - $130.2k/yr
							Remote:               true,
						},
					},
				},
				domain.Clojure: {},
				domain.Haskell: {},
			},
			ShortDescription: "The platform to simplify car buying and selling",
			Industries:       []domain.Industry{
				// NOP
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Loblaw Digital",
			Website: "https://loblawdigital.co/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                9450129,
				IDs:               nil,
				Alias:             "loblaw-digital",
				Name:              "Loblaw Digital",
				Followers:         "47K",
				Employees:         "501-1K",
				AssociatedMembers: "832",
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
							Title:                "Senior Developer (Elixir)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4205898506/",
							Date:                 mustDate("2025-04-24", "2025-04-11"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Senior Developer (Elixir)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4187442191/",
							Date:                 mustDate("2025-04-10"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Senior Developer (Elixir)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4228879444/",
							Date:                 mustDate("2025-05-13"),
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
				domain.IndustryECommerce,
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Claritev",
			Website: "https://www.claritev.com/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                6788,
				IDs:               nil,
				Alias:             "claritev",
				Name:              "Claritev",
				Followers:         "39K",
				Employees:         "1K-5K",
				AssociatedMembers: "2,541",
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
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Claritev-EI_IE16533.11,19.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Claritev-Reviews-E16533.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Claritev-Jobs-E16533.htm",
				Jobs:        "1",
				Reviews:     "508",
				Salaries:    "826",
				ReviewsRate: "3.7",
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
							Title:                "Senior Data Engineer (SQL, Scala)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4169971174/",
							Date:                 mustDate("2025-04-13"),
							WithSalary:           true,
							Remote:               true,
						},
					},
				},
				domain.Elixir:  {},
				domain.Clojure: {},
				domain.Haskell: {},
			},
			ShortDescription: "Claritev is an independent and public healthcare technology, data and insights company making healthcare more transparent, fair and affordable for all",
			Industries: []domain.Industry{
				domain.IndustryHealthTech,
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Duffel",
			Website: "https://duffel.com/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                27115109,
				IDs:               nil,
				Alias:             "duffelhq",
				Name:              "Duffel",
				Followers:         "8K",
				Employees:         "51-200",
				AssociatedMembers: "56",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:     "duffelhq",
				Followers: "34",
				Verified:  true,
			},
			BlindProfile: domain.BlindProfile{
				Alias: "",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "duffel",
				Employees: "30",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Duffel-EI_IE2327865.11,17.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Duffel-Reviews-E2327865.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Duffel-Jobs-E2327865.htm",
				Jobs:        "",
				Reviews:     "25",
				Salaries:    "61",
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
					GitHubRepositoriesCount: 13,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Software Engineer, Backend",
							ShortDescription:     "The tools used on the team include Elixir, Phoenix, Kubernetes and Google Cloud Platform",
							SwitchingOpportunity: "An interest in working with Elixir, this is what we mostly work with",
							URL:                  "https://www.linkedin.com/jobs/view/4204672355/",
							Date:                 mustDate("2025-04-10"),
							WithSalary:           false,
							Remote:               false,
						},
					},
				},
				domain.Clojure: {},
				domain.Haskell: {},
			},
			ShortDescription: "The complete toolkit for selling travel online",
			Industries:       []domain.Industry{
				// NOP
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "FOSSA",
			Website: "https://fossa.com/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                17957915,
				IDs:               nil,
				Alias:             "fossa",
				Name:              "FOSSA",
				Followers:         "7K",
				Employees:         "51-200",
				AssociatedMembers: "64",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:     "fossas",
				Followers: "135",
				Verified:  true,
			},
			BlindProfile: domain.BlindProfile{
				Alias: "",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias: "",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-FOSSA-EI_IE1901714.11,16.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/FOSSA-Reviews-E1901714.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/FOSSA-Jobs-E1901714.htm",
				Jobs:        "",
				Reviews:     "22",
				Salaries:    "29",
				ReviewsRate: "4.5",
				Verified:    false,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "",
			},
			Languages: domain.Languages{
				domain.Go: {},
				domain.Rust: {
					GitHubRepositoriesCount: 10,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Senior Software Engineer",
							ShortDescription:     "Work in a variety of languages including Rust, Go, TypeScript, and Haskell",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4190979349/",
							Date:                 mustDate("2025-03-30"),
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
			ShortDescription: "SBOM and software supply chain risk management platform",
			Industries:       []domain.Industry{
				// NOP
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "LSEG",
			Website: "https://www.lseg.com/",
			Careers: "https://www.lseg.com/careers",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                1314683,
				IDs:               []int{8388, 8916, 16619, 405895, 1314683, 2721642, 4986853, 5317799, 10426432, 14468558, 33186884},
				Alias:             "london-stock-exchange-group",
				Name:              "LSEG",
				Followers:         "478K",
				Employees:         "10K+",
				AssociatedMembers: "22,736",
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
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-LSEG-London-Stock-Exchange-Group-EI_IE11860.11,43.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/LSEG-London-Stock-Exchange-Group-Reviews-E11860.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/LSEG-London-Stock-Exchange-Group-Jobs-E11860.htm",
				Jobs:        "1K",
				Reviews:     "3.1K",
				Salaries:    "5.4K",
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
							Title:                "Senior Lead Scala Engineer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4144858808/",
							Date:                 mustDate("2025-05-03", "2025-04-12"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Lead Scala Engineer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4144858801/",
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
			ShortDescription: "LSEG (London Stock Exchange Group)",
			Industries: []domain.Industry{
				domain.IndustryFinTech,
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Getnet",
			Website: "https://www.getnetworld.com/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                76159171,
				IDs:               nil,
				Alias:             "getnetworld",
				Name:              "Getnet",
				Followers:         "372K",
				Employees:         "5K-10K",
				AssociatedMembers: "3,173",
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
							Title:                "Software Engineer (Scala/Spark)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4206647185/",
							Date:                 mustDate("2025-04-11"),
							WithSalary:           false,
							Remote:               true,
						},
					},
				},
				domain.Elixir:  {},
				domain.Clojure: {},
				domain.Haskell: {},
			},
			ShortDescription: "Customised payment solutions for merchants",
			Industries: []domain.Industry{
				domain.IndustryFinTech,
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Hunter.io",
			Website: "https://hunter.io/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                11347551,
				IDs:               nil,
				Alias:             "hunterio",
				Name:              "Hunter.io",
				Followers:         "10K",
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
							Title:                "Senior Back-End Engineer (Go)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4209630478/",
							Date:                 mustDate("2025-04-14"),
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
			ShortDescription: "Hunter is an all-in-one email outreach platform to identify relevant prospects, find their email addresses, and send cold email campaigns",
			Industries: []domain.Industry{
				domain.IndustryMarTech,
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "SawitPRO",
			Website: "https://www.sawitpro.com/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                89878086,
				IDs:               nil,
				Alias:             "sawitpro",
				Name:              "SawitPRO",
				Followers:         "39K",
				Employees:         "51-200",
				AssociatedMembers: "246",
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
							Title:                "Senior Backend Engineer (Golang)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4207981446/",
							Date:                 mustDate("2025-04-14"),
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
			ShortDescription: "Palm oil",
			Industries: []domain.Industry{
				domain.IndustryAgroTech,
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Palisade, Inc",
			Website: "https://www.palisade.co/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                80564954,
				IDs:               nil,
				Alias:             "palisade-inc",
				Name:              "Palisade, Inc",
				Followers:         "1K",
				Employees:         "2-10",
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
							Title:                "Blockchain Integration Engineer, Golang",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4208444390/",
							Date:                 mustDate("2025-04-14"),
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
			Ignore: true, // Blockchain
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "blp",
			Website: "https://blp.digital/",
			Careers: "https://blp.digital/careers",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                26561819,
				IDs:               nil,
				Alias:             "blp-digital",
				Name:              "blp",
				Followers:         "4K",
				Employees:         "51-200",
				AssociatedMembers: "69",
				Verified:          false,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:     "BLPDigital",
				Followers: "14",
				Verified:  false,
			},
			BlindProfile: domain.BlindProfile{
				Alias: "",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias: "",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-BLP-Digital-EI_IE5466795.11,22.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/BLP-Digital-Reviews-E5466795.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/BLP-Digital-Jobs-E5466795.htm",
				Jobs:        "7",
				Reviews:     "11",
				Salaries:    "10",
				ReviewsRate: "5.0",
				Verified:    false,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "",
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 2,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Senior Backend Developer – Golang",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4209809848/",
							Date:                 mustDate("2025-04-14"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Senior Backend Engineer – Go",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4264255330/",
							Location:             "Spain",
							Date:                 mustDate("2025-07-21"),
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
			ShortDescription: "ERP-Automation solution for Finance, Procurement, Logistics, Sales and many mor",
			Industries:       []domain.Industry{
				// NOP
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Wonderschool",
			Website: "https://www.wonderschool.com/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                18020048,
				IDs:               nil,
				Alias:             "wonderschool",
				Name:              "Wonderschool",
				Followers:         "13K",
				Employees:         "51-200",
				AssociatedMembers: "156",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:     "wonderschool",
				Followers: "7",
				Verified:  false,
			},
			BlindProfile: domain.BlindProfile{
				Alias: "",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias: "",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Wonderschool-EI_IE1645756.11,23.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Wonderschool-Reviews-E1645756.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Wonderschool-Jobs-E1645756.htm",
				Jobs:        "27",
				Reviews:     "61",
				Salaries:    "89",
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
					GitHubRepositoriesCount: 7,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Senior Software Engineer, Elixir",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4211013843/",
							Date:                 mustDate("2025-09-24", "2025-05-11", "2025-05-07", "2025-04-16"),
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
			Name:    "AppotaPay",
			Website: "https://appotapay.com/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                13358280,
				IDs:               nil,
				Alias:             "appotapay",
				Name:              "AppotaPay",
				Followers:         "1K",
				Employees:         "201-500",
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
							Title:                "Golang Developer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4208880078/",
							Date:                 mustDate("2025-04-15"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Golang Developer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4221961524/",
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
			ShortDescription: "Payment solution for businesses",
			Industries: []domain.Industry{
				domain.IndustryFinTech,
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Eightcap",
			Website: "https://www.eightcap.com/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                6582200,
				IDs:               nil,
				Alias:             "eightcap-aufx",
				Name:              "Eightcap",
				Followers:         "7K",
				Employees:         "501-1K",
				AssociatedMembers: "267",
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
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Eightcap-EI_IE3828636.11,19.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Eightcap-Reviews-E3828636.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Eightcap-Jobs-E3828636.htm",
				Jobs:        "21",
				Reviews:     "22",
				Salaries:    "45",
				ReviewsRate: "3.2",
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
							Title:                "Senior Golang Engineer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4171670883/",
							Date:                 mustDate("2025-05-07", "2025-04-15"),
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
				domain.IndustryFinTech,
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Saber",
			Website: "https://www.saber.app/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                101692321,
				IDs:               nil,
				Alias:             "saberapp",
				Name:              "Saber",
				Followers:         "393",
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
							Title:                "Senior Software Engineer (Go)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4209319499/",
							Date:                 mustDate("2025-04-15"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Senior Software Engineer (Go)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4212839400/",
							Date:                 mustDate("2025-05-07"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Senior Software Engineer (Go)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4266188512/",
							Location:             "Amsterdam, North Holland, Netherlands",
							Date:                 mustDate("2025-07-15"),
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
			ShortDescription: "Helping sellers double meeting rates",
			Industries:       []domain.Industry{
				// NOP
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Care.com",
			Website: "https://www.care.com/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                157258,
				IDs:               []int{157258, 2824888},
				Alias:             "caredotcom",
				Name:              "Care.com",
				Followers:         "93K",
				Employees:         "1K-5K",
				AssociatedMembers: "12,213",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:     "care-dot-com",
				Followers: "42",
				Verified:  false,
			},
			BlindProfile: domain.BlindProfile{
				Alias: "",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "carecom",
				Employees: "515",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Care-com-EI_IE248084.11,19.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Care-com-Reviews-E248084.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Care-com-Jobs-E248084.htm",
				Jobs:        "44",
				Reviews:     "442",
				Salaries:    "2K",
				ReviewsRate: "3.0",
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
							Title:                "Senior Software Engineer (Java / Golang)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4211003389/",
							Date:                 mustDate("2025-04-16"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Senior Software Engineer (Ruby / Golang)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4244807687/",
							Date:                 mustDate("2025-06-05"),
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
			Industries:       []domain.Industry{
				// NOP
			},
		},
		{
			ID:      0,                         // system
			Type:    domain.CompanyTypeStartup, // system
			Name:    "AccessOwl",
			Website: "https://www.accessowl.com/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                82764837,
				IDs:               nil,
				Alias:             "accessowl",
				Name:              "AccessOwl",
				Followers:         "1K",
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
				domain.Go:    {},
				domain.Rust:  {},
				domain.Zig:   {},
				domain.Scala: {},
				domain.Elixir: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Senior Software Engineer (TypeScript)",
							ShortDescription:     "",
							SwitchingOpportunity: "Elixir is a plus",
							URL:                  "https://www.linkedin.com/jobs/view/4117821802/",
							Date:                 mustDate("2025-01-17"),
							WithSalary:           false,
							Remote:               true,
						},
						{
							Title:                "Senior Software Engineer",
							ShortDescription:     "Elixir Focus",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4214745238/",
							Date:                 mustDate("2025-04-25"),
							WithSalary:           false,
							Remote:               true,
						},
						{
							Title:                "Senior Software Engineer",
							ShortDescription:     "Elixir Focus",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4262844732/",
							Date:                 mustDate("2025-07-08"),
							WithSalary:           true, // €70k/yr - €95k/yr
							Remote:               true,
						},
						{
							Title:                "Senior Software Engineer",
							ShortDescription:     "Elixir Focus",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4266510208/",
							Location:             "Hungary",
							Date:                 mustDate("2025-07-14"),
							WithSalary:           true, // €70k/yr - €95k/yr
							Remote:               true,
						},
						{
							Title:                "Senior Software Engineer",
							ShortDescription:     "Elixir Focus",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4289006848/",
							Location:             "Georgia",
							Date:                 mustDate("2025-08-22"),
							WithSalary:           true, // €70k/yr - €95k/yr
							Remote:               true,
						},
					},
				},
				domain.Clojure: {},
				domain.Haskell: {},
			},
			ShortDescription: "Manage SaaS apps, access, and spend in a single place",
			Industries:       []domain.Industry{
				// NOP
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Adobe",
			Website: "https://www.adobe.com/",
			Careers: "https://www.adobe.com/careers.html",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                1480,
				IDs:               []int{1480, 236175, 1053555, 5273045},
				Alias:             "adobe",
				Name:              "Adobe",
				Followers:         "5M",
				Employees:         "10K+",
				AssociatedMembers: "39,764",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:     "adobe",
				Followers: "2.5k",
				Verified:  true,
			},
			BlindProfile: domain.BlindProfile{
				Alias:       "Adobe",
				Employees:   "10,000+",
				Salary:      "$24K ~ $325K a year",
				Reviews:     "1,876",
				ReviewsRate: "4",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "adobe",
				Employees: "21,250",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Adobe-EI_IE1090.11,16.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Adobe-Reviews-E1090.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Adobe-Jobs-E1090.htm",
				Jobs:        "783",
				Reviews:     "11K",
				Salaries:    "23K",
				ReviewsRate: "4.2",
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
					GitHubRepositoriesCount: 5,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Staff Backend Engineer",
							ShortDescription:     "Build new features and services in Elixir",
							SwitchingOpportunity: "Experience in Elixir or a comparable functional programming language, or interest and a deep desire to learn",
							URL:                  "https://www.linkedin.com/jobs/view/4056647024/",
							Date:                 mustDate("2024-11-17"),
							WithSalary:           true,
							Remote:               true,
						},
						{
							Title:                "Fullstack Engineer",
							ShortDescription:     "Proficiency in Elixir, NodeJS and ReactJS is required",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4130665982/",
							Date:                 mustDate("2025-02-18"),
							WithSalary:           true,
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
			Name:    "Teller",
			Website: "https://teller.io/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                28841985,
				IDs:               nil,
				Alias:             "teller-inc",
				Name:              "Teller",
				Followers:         "1K",
				Employees:         "11-50",
				AssociatedMembers: "23",
				Verified:          false,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:     "tellerhq",
				Followers: "43",
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
				domain.Go:    {},
				domain.Rust:  {},
				domain.Zig:   {},
				domain.Scala: {},
				domain.Elixir: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Elixir Product Engineer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4140872032/",
							Date:                 mustDate("2025-02-17"),
							WithSalary:           false,
							Remote:               false,
						},
					},
				},
				domain.Clojure: {},
				domain.Haskell: {},
			},
			ShortDescription: "API for bank accounts",
			Industries: []domain.Industry{
				domain.IndustryFinTech,
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Mux",
			Website: "https://www.mux.com/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                7793078,
				IDs:               nil,
				Alias:             "mux",
				Name:              "Mux",
				Followers:         "8K",
				Employees:         "51-200",
				AssociatedMembers: "126",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:     "muxinc",
				Followers: "350",
				Verified:  true,
			},
			BlindProfile: domain.BlindProfile{
				Alias: "",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias: "",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Mux-EI_IE1500920.11,14.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Mux-Reviews-E1500920.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Mux-Jobs-E1500920.htm",
				Jobs:        "5",
				Reviews:     "31",
				Salaries:    "43",
				ReviewsRate: "4.3",
				Verified:    true,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "",
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 42,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Senior Backend Engineer – Platform Experience",
							ShortDescription:     "Develop user-facing, product-focused features in our Elixir and Go backend services",
							SwitchingOpportunity: "Elixir experience is a big plus, but Go is also great",
							URL:                  "https://www.linkedin.com/jobs/view/4152246608/",
							Date:                 mustDate("2025-02-17"),
							WithSalary:           true,
							Remote:               true,
						},
					},
				},
				domain.Rust:  {},
				domain.Zig:   {},
				domain.Scala: {},
				domain.Elixir: {
					GitHubRepositoriesCount: 7,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Senior Backend Engineer – Platform Experience",
							ShortDescription:     "Develop user-facing, product-focused features in our Elixir and Go backend services",
							SwitchingOpportunity: "Elixir experience is a big plus",
							URL:                  "https://www.linkedin.com/jobs/view/4152246608/",
							Date:                 mustDate("2025-02-17"),
							WithSalary:           true,
							Remote:               true,
						},
					},
				},
				domain.Clojure: {},
				domain.Haskell: {},
			},
			ShortDescription: "Mux is video infrastructure that makes it easy for developers to build video into their products, and do so quickly, reliably, and at global scale",
			Industries:       []domain.Industry{
				// NOP
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "CAVA",
			Website: "https://cava.com/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                1396725,
				IDs:               nil,
				Alias:             "cava-",
				Name:              "CAVA",
				Followers:         "223K",
				Employees:         "5K-10K",
				AssociatedMembers: "4,104",
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
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-CAVA-EI_IE877467.11,15.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/CAVA-Reviews-E877467.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/CAVA-Jobs-E877467.htm",
				Jobs:        "2K",
				Reviews:     "684",
				Salaries:    "1.1K",
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
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Senior Software Engineer",
							ShortDescription:     "We use Elixir and Phoenix",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4126340720/",
							Date:                 mustDate("2025-02-17"),
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
	}
}
