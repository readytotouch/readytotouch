package db

import (
	"github.com/readytotouch/readytotouch/internal/domain"
)

func companiesPart11() []domain.CompanyProfile {
	return []domain.CompanyProfile{

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Fathom.io",
			Website: "https://fathom.io/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                18699507,
				IDs:               nil,
				Alias:             "fathom-io",
				Name:              "Fathom.io",
				Followers:         "8K",
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
				domain.Go: {},
				domain.Rust: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Rust Software Engineer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4038633965/",
							Date:                 mustDate("2024-10-01"),
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
			ShortDescription: "Enterprise AI-Driven Platform For Informed Decision-Making and Fast Execution",
		},
		{
			ID:      0,                           // system
			Type:    domain.CompanyTypeOutsource, // system
			Name:    "Photon",
			Website: "https://photon.com/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                165464,
				IDs:               nil,
				Alias:             "photon-interactive",
				Name:              "Photon",
				Followers:         "526K",
				Employees:         "5K-10K",
				AssociatedMembers: "5,404",
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
							Title:                "Rust Developer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4195243048/",
							Date:                 mustDate("2025-04-01"),
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
			Ignore:           true, // Outsource
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Dragonfly",
			Website: "https://www.dragonfly.xyz/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                18845152,
				IDs:               nil,
				Alias:             "dragonfly-capital-partners",
				Name:              "Dragonfly",
				Followers:         "37K",
				Employees:         "11-50",
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
				domain.Go: {},
				domain.Rust: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Rust Software Engineer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4181606878/",
							Date:                 mustDate("2025-04-01"),
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
			ShortDescription: "Crypto venture fund",
			Ignore:           true, // Crypto
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Veeva Systems",
			Website: "https://www.veeva.com/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                402048,
				IDs:               []int{141456, 402048},
				Alias:             "veeva-systems",
				Name:              "Veeva Systems",
				Followers:         "236K",
				Employees:         "5K-10K",
				AssociatedMembers: "8,910",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:     "veeva",
				Followers: "71",
				Verified:  false,
			},
			BlindProfile: domain.BlindProfile{
				Alias:       "Veeva",
				Employees:   "1,001 to 5,000",
				Salary:      "$65K ~ $283K a year",
				Reviews:     "247",
				ReviewsRate: "3.5",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "veeva-systems",
				Employees: "4,500",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Veeva-Systems-EI_IE459351.11,24.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Veeva-Systems-Reviews-E459351.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Veeva-Systems-Jobs-E459351.htm",
				Jobs:        "812",
				Reviews:     "1.6K",
				Salaries:    "3.6K",
				ReviewsRate: "3.6",
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
							Title:                "Principal Software Engineer – Rust",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4144651944/",
							Date:                 mustDate("2025-04-15"), // mustDate("2025-04-08"), // mustDate("2025-03-31"),
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
			ShortDescription: "Software for the global life sciences industry",
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Rivian and Volkswagen Group Technologies",
			Website: "https://rivianvw.tech/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                105032955,
				IDs:               nil,
				Alias:             "rivian-and-vw-group-technologies",
				Name:              "Rivian and Volkswagen Group Technologies",
				Followers:         "9K",
				Employees:         "501-1K",
				AssociatedMembers: "640",
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
							Title:                "Connectivity Software Engineer – Rust",
							ShortDescription:     "Design and implementation of embedded applications in Rust",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4094024474/",
							Date:                 mustDate("2025-04-11"), // mustDate("2025-03-31"),
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
			Industries: []domain.Industry{
				domain.IndustryAutomotive,
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "PULT",
			Website: "https://www.pult.com/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                20488600,
				IDs:               nil,
				Alias:             "joinpult",
				Name:              "PULT",
				Followers:         "1K",
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
				domain.Go: {},
				domain.Rust: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Full-Stack Developer TypeScript / Rust",
							ShortDescription:     "Hands-on experience with the Rust",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4194668550/",
							Date:                 mustDate("2025-03-30"),
							WithSalary:           false,
							Remote:               true,
						},
						{
							Title:                "Founding Full Stack Engineer TypeScript / Rust",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4197407019/",
							Date:                 mustDate("2025-04-01"),
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
			ShortDescription: "Enabling modern workplaces to go hybrid",
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Dria",
			Website: "https://www.firstbatch.xyz/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                80364972,
				IDs:               nil,
				Alias:             "dria-ai",
				Name:              "Dria",
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
				domain.Go: {},
				domain.Rust: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Rust Backend Engineer",
							ShortDescription:     "Design and develop high-performance backend services using Rust",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4192746950/",
							Date:                 mustDate("2025-03-28"),
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
			ShortDescription: "Democratizing access to synthetic data for AI/ML models",
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Anaplan",
			Website: "https://www.anaplan.com/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                658814,
				IDs:               nil,
				Alias:             "anaplan",
				Name:              "Anaplan",
				Followers:         "309K",
				Employees:         "1K-5K",
				AssociatedMembers: "2,462",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "",
				Verified: false,
			},
			BlindProfile: domain.BlindProfile{
				Alias:       "anaplan",
				Employees:   "1,001 to 5,000",
				Salary:      "$52K ~ $248K a year",
				Reviews:     "199",
				ReviewsRate: "2.9",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "anaplan",
				Employees: "2,500",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Anaplan-EI_IE695685.11,18.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Anaplan-Reviews-E695685.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Anaplan-Jobs-E695685.htm",
				Jobs:        "81",
				Reviews:     "1.1K",
				Salaries:    "1.8K",
				ReviewsRate: "3.3",
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
							Title:                "Software Engineer (Rust)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4192707647/",
							Date:                 mustDate("2025-03-28"),
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
			ShortDescription: "Scenario planning and analysis platform designed to optimize decision-making",
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "St. Jude Children's Research Hospital",
			Website: "https://www.stjude.org/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                15223686,
				IDs:               nil,
				Alias:             "st-jude-childrens-research-hospital",
				Name:              "St. Jude Children's Research Hospital",
				Followers:         "90K",
				Employees:         "5K-10K",
				AssociatedMembers: "4,430",
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
							Title:                "Senior Staff / Principal Software Engineer, Rust Genomics",
							ShortDescription:     "This position requires a deep understanding of the Rust programming language",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4081324333/",
							Date:                 mustDate("2025-03-28"),
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
			ShortDescription: "Pediatric cancer and other life-threatening diseases of childhood",
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Strong Compute",
			Website: "https://strongcompute.com/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                76229717,
				IDs:               nil,
				Alias:             "strongcompute",
				Name:              "Strong Compute",
				Followers:         "1K",
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
				domain.Go:    {},
				domain.Rust:  {},
				domain.Zig:   {},
				domain.Scala: {},
				domain.Elixir: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Elixir Engineer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4198256844/",
							Date:                 mustDate("2025-04-01"),
							WithSalary:           false,
							Remote:               false,
						},
					},
				},
				domain.Clojure: {},
				domain.Haskell: {},
			},
			ShortDescription: "Fast compute for ML use cases",
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Yolo Group",
			Website: "https://yolo.com/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                18303220,
				IDs:               nil,
				Alias:             "yolo-group",
				Name:              "Yolo Group",
				Followers:         "23K",
				Employees:         "501-1K",
				AssociatedMembers: "813",
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
							Title:                "Senior Software Engineer (Elixir)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4194883244/",
							Date:                 mustDate("2025-04-01"),
							WithSalary:           false,
							Remote:               false,
						},
					},
				},
				domain.Clojure: {},
				domain.Haskell: {},
			},
			ShortDescription: "",
			Ignore:           true, // Blockchain
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Hard Rock Digital",
			Website: "https://www.hardrockdigital.com/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                70993771,
				IDs:               nil,
				Alias:             "hardrockdigital",
				Name:              "Hard Rock Digital",
				Followers:         "64K",
				Employees:         "501-1K",
				AssociatedMembers: "781",
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
							Title:                "Senior Elixir Engineer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4193232651/",
							Date:                 mustDate("2025-03-31"),
							WithSalary:           false,
							Remote:               true,
						},
					},
				},
				domain.Clojure: {},
				domain.Haskell: {},
			},
			ShortDescription: "Online sports betting",
			Ignore:           true, // Betting
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Tubi",
			Website: "http://tubitv.com/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                3650994,
				IDs:               nil,
				Alias:             "tubi-tv",
				Name:              "Tubi",
				Followers:         "113K",
				Employees:         "501-1K",
				AssociatedMembers: "788",
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
							Title:                "Staff Software Engineer (Elixir)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4195366164/",
							Date:                 mustDate("2025-03-28"),
							WithSalary:           false,
							Remote:               false,
						},
					},
				},
				domain.Clojure: {},
				domain.Haskell: {},
			},
			ShortDescription: "Free TV and movie streaming service",
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Samsung Food",
			Website: "https://samsungfood.com/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                2650412,
				IDs:               nil,
				Alias:             "samsungfoodofficial",
				Name:              "Samsung Food",
				Followers:         "4K",
				Employees:         "51-200",
				AssociatedMembers: "77",
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
							Title:                "Backend Developer (Scala)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4149148628/",
							Date:                 mustDate("2025-03-31"),
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
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Ace & Tate",
			Website: "https://www.aceandtate.com/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                3085253,
				IDs:               nil,
				Alias:             "3085253", // "ace-&-tate",
				PreviousAliases:   []string{"ace-&-tate"},
				Name:              "Ace & Tate",
				Followers:         "41K",
				Employees:         "501-1K",
				AssociatedMembers: "472",
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
							Title:                "Backend Engineer (Scala)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4195967500/",
							Date:                 mustDate("2025-03-29"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Backend Engineer (Scala)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4198318379/",
							Date:                 mustDate("2025-04-03"),
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
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "GR8 Tech",
			Website: "https://gr8.tech/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                90353573,
				IDs:               nil,
				Alias:             "gr8-tech",
				Name:              "GR8 Tech",
				Followers:         "22K",
				Employees:         "501-1K",
				AssociatedMembers: "652",
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
							Title:                "Senior Scala Developer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4198081792/",
							Date:                 mustDate("2025-04-02"),
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
			Ignore:           true, // Gambling
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Axpo Group",
			Website: "https://www.axpo.com/",
			Careers: "https://www.axpo.com/career",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                79459,
				IDs:               []int{68129, 79459, 4867117, 6026057, 10039680, 10053807, 11857624, 16197419, 23309232, 24783919, 68529775, 68558805, 68558888, 68808354, 69500006, 69509458, 104087292, 104684084, 106328455},
				Alias:             "axpo-group",
				Name:              "Axpo Group",
				Followers:         "71K",
				Employees:         "5K-10K",
				AssociatedMembers: "4,502",
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
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Axpo-EI_IE755157.11,15.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Axpo-Reviews-E755157.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Axpo-Jobs-E755157.htm",
				Jobs:        "194",
				Reviews:     "171",
				Salaries:    "356",
				ReviewsRate: "4.1",
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
							Title:                "Scala Developer / Data Engineer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4198014569/",
							Date:                 mustDate("2025-04-01"),
							WithSalary:           false,
							Remote:               false,
						},
					},
				},
				domain.Elixir:  {},
				domain.Clojure: {},
				domain.Haskell: {},
			},
			ShortDescription: "Axpo is driven by a single purpose – to enable a sustainable future through innovative energy solutions",
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "SwissBorg",
			Website: "https://swissborg.com/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                11040884,
				IDs:               nil,
				Alias:             "swissborg",
				Name:              "SwissBorg",
				Followers:         "39K",
				Employees:         "201-500",
				AssociatedMembers: "241",
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
							Title:                "Scala Engineer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4197856511/",
							Date:                 mustDate("2025-04-02"),
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
			Ignore:           true, // Crypto
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
