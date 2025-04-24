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
						{
							Title:                "Go Software Engineer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4192382841/",
							Date:                 mustDate("2025-04-15"),
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
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Software Engineer (Go / Full Stack)",
							ShortDescription:     "3+ years of experience developing large scale web applications using Go",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4174137347/",
							Date:                 mustDate("2025-04-18"), // mustDate("2025-03-27"),
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
				Login:    "twitchtv",
				Verified: true,
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
				Jobs:        "775",
				Reviews:     "16K",
				Salaries:    "28K",
				ReviewsRate: "4.0",
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
					GitHubRepositoriesCount: 0,
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
							Date:                 mustDate("2025-04-10"),
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
				Login:    "dailypay",
				Verified: true,
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
				Jobs:        "",
				Reviews:     "295",
				Salaries:    "549",
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
							Title:                "Senior Software Engineer (Go/AWS)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4200514129/",
							Date:                 mustDate("2025-04-04"),
							WithSalary:           true, // £55,000—£73,000 GBP
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
				Login:    "shopware",
				Verified: false,
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
				Reviews:     "12",
				Salaries:    "47",
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
							Title:                "Rust Developer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4199352625/",
							Date:                 mustDate("2025-04-04"),
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
				Login:    "VitalBio",
				Verified: true,
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
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Senior Rust Application Engineer",
							ShortDescription:     "Strong proficiency in Rust programming with experience developing production-grade applications",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4199869768/",
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
					GitHubRepositoriesCount: 0,
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
			ID:      0,                           // system
			Type:    domain.CompanyTypeOutsource, // system
			Name:    "Vinova",
			Website: "https://vinova.sg/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                1614934,
				IDs:               nil,
				Alias:             "vinova-sg",
				Name:              "Vinova",
				Followers:         "4K",
				Employees:         "201-500",
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
			Ignore: true,
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Kantar",
			Website: "https://www.kantar.com/",
			Careers: "",
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
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Verkada-EI_IE2153775.11,18.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Verkada-Reviews-E2153775.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Verkada-Jobs-E2153775.htm",
				Jobs:        "136",
				Reviews:     "902",
				Salaries:    "1.8K",
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
							Title:                "Go Software Engineer",
							ShortDescription:     "Develop and maintain Go firmware for embedded devices with focus on performance and security",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4205197888/",
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
							Date:                 mustDate("2025-04-08"),
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
							URL:                  "https://www.linkedin.com/jobs/view/4205812062/",
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
				Alias: "",
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
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
			Careers: "",
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
			Careers: "",
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
							Date:                 mustDate("2025-04-09"),
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
							Date:                 mustDate("2025-04-08"),
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
							Title:                "Rust – Software Engineer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4204685114/",
							Date:                 mustDate("2025-04-21"), // mustDate("2025-04-11"),
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
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Maropost-EI_IE994532.11,19.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Maropost-Reviews-E994532.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Maropost-Jobs-E994532.htm",
				Jobs:        "23",
				Reviews:     "321",
				Salaries:    "320",
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
							Title:                "Software Engineer (Scala)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4207470731/",
							Date:                 mustDate("2025-04-12"),
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
			Careers: "",
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
				domain.Go:   {},
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
							Date:                 mustDate("2025-04-11"),
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
							Date:                 mustDate("2025-04-23"), // mustDate("2025-04-15"), // mustDate("2025-04-12"),
							WithSalary:           false,
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
				Reviews:     "20",
				Salaries:    "42",
				ReviewsRate: "3.6",
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
							Date:                 mustDate("2025-04-16"), // mustDate("2025-04-11"),
							WithSalary:           false,
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
							Date:                 mustDate("2025-04-11"),
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
					GitHubRepositoriesCount: 0,
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
					GitHubRepositoriesCount: 0,
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
			Name:    "LSEG (London Stock Exchange Group)",
			Website: "https://www.lseg.com/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                1314683,
				IDs:               []int{8388, 8916, 16619, 405895, 1314683, 2721642, 4986853, 5317799, 10426432, 14468558, 33186884},
				Alias:             "london-stock-exchange-group",
				Name:              "LSEG (London Stock Exchange Group)",
				Followers:         "456K",
				Employees:         "10K+",
				AssociatedMembers: "22,632",
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
							Date:                 mustDate("2025-04-12"),
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
							Title:                "Senior Back-end Engineer (Go)",
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
			Careers: "",
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
							Title:                "Senior Backend Developer – Golang",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4209809848/",
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
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Senior Software Engineer, Elixir",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4211013843/",
							Date:                 mustDate("2025-04-16"),
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
							URL:                  "https://www.linkedin.com/jobs/view/4171670883/",
							Date:                 mustDate("2025-04-15"),
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
				Alias: "",
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
			Careers: "",
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
				Alias: "",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias: "",
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
					GitHubRepositoriesCount: 0,
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
					GitHubRepositoriesCount: 0,
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
				domain.Go:    {},
				domain.Rust:  {},
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
			Careers: "",
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
				Alias: "",
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
							Date:                 mustDate("2024-04-17"), // mustDate("2025-04-17"),
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
							Title:                "Senior Software Engineer – Elixir",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/3860724318/",
							Date:                 mustDate("2024-04-17"), // mustDate("2025-04-17"),
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
				domain.Go:    {},
				domain.Rust:  {},
				domain.Zig:   {},
				domain.Scala: {},
				domain.Elixir: {
					GitHubRepositoriesCount: 0,
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
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Elixir Software Developer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4128427094/",
							Date:                 mustDate("2025-02-18"),
							WithSalary:           true,
							Remote:               true,
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
					GitHubRepositoriesCount: 0,
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
				Alias: "",
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
				Alias: "",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias: "",
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
					},
				},
				domain.Rust:    {},
				domain.Zig:     {},
				domain.Scala:   {},
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
							ShortDescription:     "3+ years of experience with Go",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4092820547/",
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
			Careers: "",
			About:   "",
			Blog:    "",
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
				Alias: "",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Muzz-EI_IE3256123.11,15.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Muzz-Reviews-E3256123.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Muzz-Jobs-E3256123.htm",
				Jobs:        "13",
				Reviews:     "44",
				Salaries:    "70",
				ReviewsRate: "3.3",
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
							Title:                "Go Engineer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4207580495/",
							Date:                 mustDate("2025-04-13"),
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
							Title:                "Senior Full Stack Engineer (Golang & React.js)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4208614176/",
							Date:                 mustDate("2025-04-17"),
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
							Title:                "Backend Go Developer – Kubernetes infrastructure",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4209740652/",
							Date:                 mustDate("2025-04-17"),
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
			Careers: "",
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
				Alias: "",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias: "",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Rockwell-Automation-EI_IE567.11,30.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Rockwell-Automation-Reviews-E567.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Rockwell-Automation-Jobs-E567.htm",
				Jobs:        "726",
				Reviews:     "4.3K",
				Salaries:    "7.5K",
				ReviewsRate: "3.7",
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
							Title:                "Senior Engineer, Software (Rust)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4212980639/",
							Date:                 mustDate("2025-04-18"),
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
				Alias: "",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias: "",
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
				Alias: "",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias: "",
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
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Omniful-EI_IE7591023.11,18.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Omniful-Reviews-E7591023.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Omniful-Jobs-E7591023.htm",
				Jobs:        "",
				Reviews:     "8",
				Salaries:    "9",
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
							Title:                "Golang Developer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4212688691/",
							Date:                 mustDate("2025-04-21"),
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
		//	Industries:       []domain.Industry{
		//		// NOP
		//	},
		//},
	}
}
