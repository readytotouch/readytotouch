package db

import (
	"github.com/readytotouch/readytotouch/internal/domain"
)

func companiesPart12() []domain.CompanyProfile {
	return []domain.CompanyProfile{

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Starship Technologies",
			Website: "https://www.starship.xyz/",
			Careers: "https://www.starship.xyz/careers/",
			About:   "https://www.starship.xyz/about/",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                6646555,
				Alias:             "starshiptechnologies",
				Name:              "Starship Technologies",
				Followers:         "36K",
				Employees:         "201-500",
				AssociatedMembers: "511",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "starship-technologies",
				Verified: false,
			},
			BlindProfile: domain.BlindProfile{
				Alias: "",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "starship-technologies",
				Employees: "420",
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
							Title:                "Senior Software Engineer",
							ShortDescription:     "Developing software and robot infrastructure services. Go and Kubernetes in the cloud. Rust on the robot",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4174644239/",
							Date:                 mustDate("2025-03-05"),
							WithSalary:           false,
							Remote:               false,
						},
					},
				},
				domain.Rust: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Computer Vision Lead",
							ShortDescription:     "Proficiency in C/C++/Rust, OpenCV, and Python",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4174437544/",
							Date:                 mustDate("2025-03-02"),
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
			ShortDescription: "Estonian company developing autonomous delivery vehicles",
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Datalust",
			Website: "https://datalust.co/",
			Careers: "https://datalust.co/careers/senior-software-engineer-1022",
			About:   "",
			Blog:    "https://blog.datalust.co/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                7947624,
				Alias:             "datalust",
				Name:              "Datalust",
				Followers:         "184",
				Employees:         "2-10",
				AssociatedMembers: "3",
				Verified:          false,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "datalust",
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
				domain.Go:      {},
				domain.Rust:    {},
				domain.Zig:     {},
				domain.Scala:   {},
				domain.Elixir:  {},
				domain.Clojure: {},
				domain.Haskell: {},
			},
			ShortDescription: "",
			Ignore:           true,
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Stan.",
			Website: "https://www.stan.com.au/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                4858643,
				Alias:             "stan-entertainment",
				Name:              "Stan.",
				Followers:         "51K",
				Employees:         "201-500",
				AssociatedMembers: "318",
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
							Title:                "Go Developer",
							ShortDescription:     "1 year experience with Go",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4195296291/",
							Date:                 mustDate("2025-04-02"),
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
			ShortDescription: "Australian streaming platform",
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Skyro",
			Website: "https://www.skyro.ph/",
			Careers: "",
			About:   "https://www.skyro.ph/about-us",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                91642572,
				Alias:             "skyroph",
				Name:              "Skyro",
				Followers:         "7K",
				Employees:         "1K-5K",
				AssociatedMembers: "591",
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
							Title:                "Middle Golang Developer",
							ShortDescription:     "3 years of professional experience in Golang development",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4193235343/",
							Date:                 mustDate("2025-04-01"),
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
			ShortDescription: "Philippines FinTech company",
			Industries: []domain.Industry{
				domain.IndustryFinTech,
			},
		},

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
				Login:    "vidsy",
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
							Title:                "Go Software Engineer",
							ShortDescription:     "2+ years experience in software engineering",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4194783969/",
							Date:                 mustDate("2025-03-27"),
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
				Login:    "rbc",
				Verified: true,
			},
			BlindProfile: domain.BlindProfile{
				Alias: "",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "rbc",
				Employees: "88,840",
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
							Title:                "Go Developer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4189008372/",
							Date:                 mustDate("2025-03-27"),
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
							Title:                "Software Engineer Team Leader — Go (Golang)",
							ShortDescription:     "6+ years of development experience",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4195065700/",
							Date:                 mustDate("2025-04-02"),
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
							Title:                "Senior Lead Software Engineer (C#, Go) — R&D / Data Services",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4195036459/",
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
				Login:    "Lightspeed-Systems",
				Verified: false,
			},
			BlindProfile: domain.BlindProfile{
				Alias: "",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "lightspeed-systems",
				Employees: "240",
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
							Title:                "Software Engineer (Go / Full Stack)",
							ShortDescription:     "3+ years of experience developing large scale web applications using Go",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4174137347/",
							Date:                 mustDate("2025-03-27"),
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
				Login:    "knime",
				Verified: false,
			},
			BlindProfile: domain.BlindProfile{
				Alias: "",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "knime",
				Employees: "150",
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
							URL:                  "https://www.linkedin.com/jobs/view/4132287368/",
							Date:                 mustDate(""),
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
			Name:    "Thema",
			Website: "https://www.thema.ai/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                85651783,
				Alias:             "thema-ai",
				Name:              "Thema",
				Followers:         "585",
				Employees:         "2-10",
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
				domain.Go:      {},
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
				Login:    "twitchtv",
				Verified: true,
			},
			BlindProfile: domain.BlindProfile{
				Alias: "",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "twitch",
				Employees: "14,260",
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
				Login:    "assembledhq",
				Verified: true,
			},
			BlindProfile: domain.BlindProfile{
				Alias: "",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "assembled",
				Employees: "126",
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
			Name:    "Snowplow",
			Website: "https://snowplow.io/",
			Careers: "https://snowplow.io/careers",
			About:   "https://snowplow.io/about-us",
			Blog:    "https://snowplow.io/blog",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                3177331,
				Alias:             "snowplow",
				Name:              "Snowplow",
				Followers:         "15K",
				Employees:         "51-200",
				AssociatedMembers: "161",
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
				Alias:     "snowplow",
				Employees: "150",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "",
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "",
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
				Alias:             "zelis",
				Name:              "Zelis",
				Followers:         "70K",
				Employees:         "1K-5K",
				AssociatedMembers: "2,525",
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
				domain.Go:      {},
				domain.Rust:    {},
				domain.Zig:     {},
				domain.Scala:   {},
				domain.Elixir:  {},
				domain.Clojure: {},
				domain.Haskell: {},
			},
			ShortDescription: "Healthcare technology company",
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "SuperAnnotate",
			Website: "https://www.superannotate.com/",
			Careers: "https://www.superannotate.com/careers",
			About:   "https://www.superannotate.com/company",
			Blog:    "https://www.superannotate.com/blog",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                18999422,
				Alias:             "superannotate",
				Name:              "SuperAnnotate",
				Followers:         "27K",
				Employees:         "51-200",
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
				Alias:     "superannotate",
				Employees: "150",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "",
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "",
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
				Login:    "ford",
				Verified: false,
			},
			BlindProfile: domain.BlindProfile{
				Alias: "",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "ford-motor",
				Employees: "187,520",
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
							Title:                "Senior Rust Developer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4198124231/",
							Date:                 mustDate("2025-04-03"),
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
				Login:    "GitGuardian",
				Verified: true,
			},
			BlindProfile: domain.BlindProfile{
				Alias: "",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "gitguardian",
				Employees: "75",
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
							Title:                "Senior Software Engineer (BackEnd, Rust)",
							ShortDescription:     "3+ years of experience working in a Rust environment, contributing to products where performance and portability are key",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4196474583/",
							Date:                 mustDate("2025-04-02"),
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
