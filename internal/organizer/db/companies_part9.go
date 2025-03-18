package db

import (
	"github.com/readytotouch/readytotouch/internal/domain"
)

func companiesPart9() []domain.CompanyProfile {
	return []domain.CompanyProfile{

		{
			ID:      0,                           // system
			Type:    domain.CompanyTypeOutsource, // system
			Name:    "Brillio",
			Website: "https://www.brillio.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                5022712,
				IDs:               nil,
				Alias:             "brillio",
				Name:              "Brillio",
				Followers:         "475K",
				Employees:         "5K-10K",
				AssociatedMembers: "5,206",
				Verified:          true,
			},
			Ignore: true, // Outsourcing
		},
		{
			ID:      0,                           // system
			Type:    domain.CompanyTypeOutsource, // system
			Name:    "Akvelon",
			Website: "https://akvelon.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                309542,
				IDs:               nil,
				Alias:             "akvelon",
				Name:              "Akvelon, Inc.",
				Followers:         "20K",
				Employees:         "1K-5K",
				AssociatedMembers: "632",
				Verified:          false,
			},
			Ignore: true, // Outsourcing
		},
		{
			ID:      0,                           // system
			Type:    domain.CompanyTypeOutsource, // system
			Name:    "Multi Media",
			Website: "https://multimediallc.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                10496448,
				IDs:               nil,
				Alias:             "multimedia-llc",
				Name:              "Multi Media, LLC",
				Followers:         "33K",
				Employees:         "201-500",
				AssociatedMembers: "308",
				Verified:          false,
			},
			Ignore: true, // Outsourcing
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "DNSFilter",
			Website: "https://www.dnsfilter.com/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                10288375,
				IDs:               nil,
				Alias:             "dnsfilter",
				Name:              "DNSFilter",
				Followers:         "23K",
				Employees:         "51-200",
				AssociatedMembers: "171",
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
							Title:                "Senior Software Engineer – Golang",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4174530724/",
							Date:                 mustDate("2025-03-06"),
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
			ShortDescription: "DNS security that identifies threats",
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Forcepoint",
			Website: "https://www.forcepoint.com/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                7584467,
				IDs:               nil,
				Alias:             "forcepoint",
				Name:              "Forcepoint",
				Followers:         "166K",
				Employees:         "1K-5K",
				AssociatedMembers: "1,813",
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
							Title:                "Principal Engineer (Golang, Java, AWS & Data Engineering)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4157762453/",
							Date:                 mustDate("2025-03-06"),
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
		},
		{
			ID:      0,                           // system
			Type:    domain.CompanyTypeOutsource, // system
			Name:    "BairesDev",
			Website: "https://www.bairesdev.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                798671,
				IDs:               nil,
				Alias:             "bairesdev",
				Name:              "BairesDev",
				Followers:         "656K",
				Employees:         "1K-5K",
				AssociatedMembers: "3,011",
				Verified:          true,
			},
			Ignore: true, // Outsourcing
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "NU",
			Website: "https://www.thenu.com/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                76545362,
				IDs:               nil,
				Alias:             "thenu",
				Name:              "NU",
				Followers:         "5K",
				Employees:         "11-50",
				AssociatedMembers: "98",
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
							Title:                "Golang Back-End Engineer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4174841291/",
							Date:                 mustDate("2025-03-05"),
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
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Barracuda",
			Website: "http://www.barracuda.com/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                13957,
				IDs:               nil,
				Alias:             "barracuda-networks",
				Name:              "Barracuda",
				Followers:         "69K",
				Employees:         "1K-5K",
				AssociatedMembers: "2,148",
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
							Title:                "Software Engineer (React, Java Spring, Golang)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4173786381/",
							Date:                 mustDate("2025-03-05"),
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
				domain.IndustryCyberSecurity,
			},
		},
		{
			ID:      0,                         // system
			Type:    domain.CompanyTypeStartup, // system
			Name:    "FlatPeak",
			Website: "https://flatpeak.com/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                82675445,
				IDs:               nil,
				Alias:             "flatpeak",
				Name:              "FlatPeak",
				Followers:         "1K",
				Employees:         "11-50",
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
							URL:                  "https://www.linkedin.com/jobs/view/4172858065/",
							Date:                 mustDate("2025-03-04"),
							WithSalary:           true, // £65,000-£80,000 + bonus and stock options
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
			ShortDescription: "Energy tariff data and optimisation platform",
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Teradata",
			Website: "https://www.teradata.com/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                1466,
				IDs:               []int{1466, 14996},
				Alias:             "teradata",
				Name:              "Teradata",
				Followers:         "404K",
				Employees:         "10K+",
				AssociatedMembers: "10,263",
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
							Title:                "Cloud Engineer (Golang/Java/Python – RestAPI)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4149196862/",
							Date:                 mustDate("2025-03-04"),
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
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Teramind",
			Website: "https://www.teramind.co/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                5090184,
				IDs:               nil,
				Alias:             "teramindco",
				Name:              "Teramind",
				Followers:         "60K",
				Employees:         "51-200",
				AssociatedMembers: "161",
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
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4172639728/",
							Date:                 mustDate("2025-03-04"),
							WithSalary:           false,
							Remote:               true,
						},
						{
							Title:                "Senior Golang Software Engineer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4172643376/",
							Date:                 mustDate("2025-03-04"),
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
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Lotus's",
			Website: "http://www.lotuss.com/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                3830488,
				IDs:               nil,
				Alias:             "lotusretails",
				Name:              "Lotus's",
				Followers:         "70K",
				Employees:         "10K+",
				AssociatedMembers: "2,932",
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
							Title:                "Senior Software Engineer – Backend (Golang)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4117492485/",
							Date:                 mustDate("2025-03-04"),
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
			ShortDescription: "Lotus’s is an omni-channel retailer in Thailand",
		},
		{
			ID:      0,                           // system
			Type:    domain.CompanyTypeOutsource, // system
			Name:    "Ubique Systems",
			Website: "https://www.ubique-systems.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                939675,
				IDs:               nil,
				Alias:             "ubique-systems",
				Name:              "Ubique Systems",
				Followers:         "229K",
				Employees:         "501-1K",
				AssociatedMembers: "307",
				Verified:          true,
			},
			Ignore: true, // Outsourcing
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Aescape",
			Website: "http://aescape.com/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                11234163,
				IDs:               nil,
				Alias:             "aescape",
				Name:              "Aescape",
				Followers:         "13K",
				Employees:         "51-200",
				AssociatedMembers: "159",
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
							Title:                "Senior/Staff Backend Engineer, Golang",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4172541606/",
							Date:                 mustDate("2025-03-04"),
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
			ShortDescription: "Building intelligent massage therapy to help people live better, longer",
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "SMC",
			Website: "https://www.smcusa.com/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                6122,
				IDs:               nil,
				Alias:             "smc",
				Name:              "SMC",
				Followers:         "156K",
				Employees:         "10K+",
				AssociatedMembers: "8,749",
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
							Title:                "Lead Software Engineer – Golang",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4172507663/",
							Date:                 mustDate("2025-03-04"),
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
		},
		{
			ID:      0,                           // system
			Type:    domain.CompanyTypeOutsource, // system
			Name:    "Arcanys",
			Website: "http://www.arcanys.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                1472152,
				IDs:               nil,
				Alias:             "arcanys",
				Name:              "Arcanys",
				Followers:         "27K",
				Employees:         "201-500",
				AssociatedMembers: "273",
				Verified:          true,
			},
			Ignore: true, // Outsourcing
		},
		{
			ID:      0,                           // system
			Type:    domain.CompanyTypeOutsource, // system
			Name:    "Skillspark",
			Website: "http://www.skillspark.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                80284986,
				IDs:               nil,
				Alias:             "skillsparkab",
				Name:              "Skillspark",
				Followers:         "68K",
				Employees:         "51-200",
				AssociatedMembers: "33",
				Verified:          false,
			},
			Ignore: true, // Outsourcing
		},
		{
			ID:      0,                           // system
			Type:    domain.CompanyTypeOutsource, // system
			Name:    "Qinshift",
			Website: "https://qinshift.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                92778277,
				IDs:               nil,
				Alias:             "qinshift",
				Name:              "Qinshift",
				Followers:         "52K",
				Employees:         "1K-5K",
				AssociatedMembers: "1,717",
				Verified:          true,
			},
			Ignore: true, // Outsourcing
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Crypto",
			Website: "https://www.crypto.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                13322447,
				IDs:               nil,
				Alias:             "cryptocom",
				Name:              "Crypto.com",
				Followers:         "590K",
				Employees:         "1K-5K",
				AssociatedMembers: "6,502",
				Verified:          true,
			},
			Ignore: true, // Crypto
		},
		{
			ID:      0,                         // system
			Type:    domain.CompanyTypeStartup, // system
			Name:    "Depixen",
			Website: "http://www.depixen.com/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                19062020,
				IDs:               nil,
				Alias:             "depixen-sw",
				Name:              "Depixen",
				Followers:         "8K",
				Employees:         "11-50",
				AssociatedMembers: "33",
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
							Title:                "Senior Backend Developer | Senior Go Developer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4174831245/",
							Date:                 mustDate("2025-03-05"),
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
			ShortDescription: "AI-Powered Autonomous Product Management Systems, Ontology-Based Digital Transformation, RDF-Based Semantic Applications",
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Climatiq",
			Website: "https://www.climatiq.io/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                72159281,
				IDs:               nil,
				Alias:             "climatiq",
				Name:              "Climatiq",
				Followers:         "8K",
				Employees:         "11-50",
				AssociatedMembers: "30",
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
							Title:                "Senior Full Stack Engineer – Climate Tech – Rust & TypeScript",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4175571523/",
							Date:                 mustDate("2025-03-07"),
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
			ShortDescription: "The scope 1, 2, and 3 carbon calculation engine",
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "commercetools",
			Website: "https://www.commercetools.com/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                126050,
				IDs:               nil,
				Alias:             "commercetools",
				Name:              "commercetools",
				Followers:         "41K",
				Employees:         "501-1K",
				AssociatedMembers: "696",
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
							Title:                "Back-End Engineer – Scala",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4175367766/",
							Date:                 mustDate("2025-03-07"),
							WithSalary:           false,
							Remote:               true,
						},
						{
							Title:                "Senior Scala Engineer",
							ShortDescription:     "Typelevel & Elasticsearch",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4172530640/",
							Date:                 mustDate("2025-03-11"),
							WithSalary:           false,
							Remote:               false,
						},
					},
				},
				domain.Elixir:  {},
				domain.Clojure: {},
				domain.Haskell: {},
			},
			ShortDescription: "Composable commerce",
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Net Health",
			Website: "https://www.nethealth.com/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                120271,
				IDs:               []int{90636, 120271},
				Alias:             "net-health-systems-inc-",
				Name:              "Net Health",
				Followers:         "24K",
				Employees:         "501-1K",
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
				domain.Go:    {},
				domain.Rust:  {},
				domain.Zig:   {},
				domain.Scala: {},
				domain.Elixir: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Senior Elixir Developer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4148903233/",
							Date:                 mustDate("2025-03-12"), // mustDate("2025-03-06"),
							WithSalary:           true,                   // $104.000 - $130.000 per year
							Remote:               false,
						},
					},
				},
				domain.Clojure: {},
				domain.Haskell: {},
			},
			ShortDescription: "Specialized software and analytics that serves the continuum of restorative care, from hospital to home",
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Dscout",
			Website: "https://dscout.com/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                2438748,
				IDs:               nil,
				Alias:             "dscout",
				Name:              "Dscout",
				Followers:         "25K",
				Employees:         "201-500",
				AssociatedMembers: "255",
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
							Title:                "Lead Software Engineer – Elixir",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4149944830/",
							Date:                 mustDate("2025-02-14"),
							WithSalary:           false,
							Remote:               true,
						},
					},
				},
				domain.Clojure: {},
				domain.Haskell: {},
			},
			ShortDescription: "Experience Research Platform",
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Viasat",
			Website: "http://www.viasat.com/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                5770,
				IDs:               nil,
				Alias:             "viasat",
				Name:              "Viasat",
				Followers:         "166K",
				Employees:         "5K-10K",
				AssociatedMembers: "6,382",
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
							Title:                "Senior Software Engineer (Java, Clojure)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4166326665/",
							Date:                 mustDate("2025-02-28"),
							WithSalary:           false,
							Remote:               false,
						},
					},
				},
				domain.Haskell: {},
			},
			ShortDescription: "Telecommunications",
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Pitch",
			Website: "http://pitch.com/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                27222819,
				IDs:               nil,
				Alias:             "pitchhq",
				Name:              "Pitch",
				Followers:         "30K",
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
							Title:                "Senior Backend Clojure Engineer (m/f/d)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4128083196/",
							Date:                 mustDate("2025-02-07"),
							WithSalary:           false,
							Remote:               true,
						},
					},
				},
				domain.Haskell: {},
			},
			ShortDescription: "The complete pitching platform",
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Doccla",
			Website: "http://www.doccla.com/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                33432702,
				IDs:               nil,
				Alias:             "doccla",
				Name:              "Doccla",
				Followers:         "11K",
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
							Title:                "Senior Software Engineer – Clojure",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4068248164/",
							Date:                 mustDate("2024-11-07"),
							WithSalary:           false,
							Remote:               false,
						},
					},
				},
				domain.Haskell: {},
			},
			ShortDescription: "Supporting healthcare organisations deliver virtual ward services proven to reduce costs while improving outcomes",
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "University of Virginia",
			Website: "https://www.virginia.edu/",
			Careers: "https://jobs.virginia.edu/",
			About:   "https://www.virginia.edu/aboutuva/",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                4298,
				IDs:               nil,
				Alias:             "university-of-virginia",
				Name:              "University of Virginia",
				Followers:         "311K",
				Employees:         "",
				AssociatedMembers: "",
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
				Alias:     "university-of-virginia",
				Employees: "8,397",
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
							URL:                  "https://www.linkedin.com/jobs/view/4177332766/",
							Date:                 mustDate("2025-03-13"), // mustDate("2025-03-07"),
							WithSalary:           true,                   // $110,000 - $130,000 per year
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
			ShortDescription: "American public institution of higher learning",
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Phrase",
			Website: "https://phrase.com/",
			Careers: "https://phrase.com/careers/",
			About:   "https://phrase.com/about/",
			Blog:    "https://phrase.com/blog/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                1344091,
				IDs:               nil,
				Alias:             "phraseplatform",
				Name:              "Phrase",
				Followers:         "36K",
				Employees:         "201-500",
				AssociatedMembers: "374",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "phrase",
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
				domain.Go:    {},
				domain.Rust:  {},
				domain.Zig:   {},
				domain.Scala: {},
				domain.Elixir: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Senior Full Stack (Elixir) Engineer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4176875904/",
							Date:                 mustDate("2025-03-09"),
							WithSalary:           false,
							Remote:               true,
						},
					},
				},
				domain.Clojure: {},
				domain.Haskell: {},
			},
			ShortDescription: "Czech AI-led translation technology",
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Mr Price Group",
			Website: "https://mrpricegroup.com/",
			Careers: "https://mrpcareers.com/our-careers/",
			About:   "https://mrpricegroup.com/about-us/",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                47016,
				IDs:               []int{47016, 3110091},
				Alias:             "mr-price-group",
				Name:              "Mr Price Group",
				Followers:         "622K",
				Employees:         "10K+",
				AssociatedMembers: "16,870",
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
							Title:                "Lead Clojure Developer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4173567999/",
							Date:                 mustDate("2025-03-05"),
							WithSalary:           false,
							Remote:               true,
						},
					},
				},
				domain.Haskell: {},
			},
			ShortDescription: "SA omni-channel, fashion value retailer",
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "PlentyONE",
			Website: "https://www.plentyone.com/",
			Careers: "https://www.plentyone.com/careers",
			About:   "https://www.plentyone.com/about",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                5079248,
				IDs:               nil,
				Alias:             "plentyone-ecommerce",
				Name:              "PlentyONE",
				Followers:         "5K",
				Employees:         "201-500",
				AssociatedMembers: "237",
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
							URL:                  "https://www.linkedin.com/jobs/view/4177522085/",
							Date:                 mustDate("2025-03-09"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Golang Developer",
							ShortDescription:     "Plentychannel",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4178761414/",
							Date:                 mustDate("2025-03-13"),
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
			ShortDescription: "E-Commerce ERP system for online retailers, brands and manufacturers",
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Future",
			Website: "https://www.future.co/",
			Careers: "https://www.future.co/careers",
			About:   "https://www.future.co/about",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                33275761,
				IDs:               nil,
				Alias:             "future-research-inc",
				Name:              "Future",
				Followers:         "13K",
				Employees:         "201-500",
				AssociatedMembers: "654",
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
							Title:                "Backend Engineer (Golang)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4179354991/",
							Date:                 mustDate("2025-03-10"),
							WithSalary:           true, // Salary range: $160,000 - 215,000 / year + equity
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
			ShortDescription: "Online fitness training platform",
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Los Alamos National Laboratory",
			Website: "https://www.lanl.gov/",
			Careers: "https://www.lanl.gov/careers",
			About:   "https://www.lanl.gov/about",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                5327,
				IDs:               nil,
				Alias:             "los-alamos-national-laboratory",
				Name:              "Los Alamos National Laboratory",
				Followers:         "157K",
				Employees:         "10K+",
				AssociatedMembers: "12,716",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "LANL",
				Verified: true,
			},
			BlindProfile: domain.BlindProfile{
				Alias: "",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "los-alamos-national-laboratory",
				Employees: "12,752",
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
							Title:                "DevOps Python/Rust Developer (Scientist 1/2)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4181750667/",
							Date:                 mustDate("2025-03-11"),
							WithSalary:           true, // Scientist 1 ($94,5-$154,600), Scientist 2 ($104,100-$172,200)
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
			ShortDescription: "American laboratory's who provide scientific and engineering support to national security programs",
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Volkswagen Digital Solutions",
			Website: "https://www.vwds.pt/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                19035234,
				IDs:               nil,
				Alias:             "volkswagen-digital-solutions",
				Name:              "Volkswagen Digital Solutions",
				Followers:         "47K",
				Employees:         "501-1K",
				AssociatedMembers: "694",
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
							URL:                  "https://www.linkedin.com/jobs/view/4157350713/",
							Date:                 mustDate("2025-03-11"),
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
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Binance",
			Website: "https://www.binance.com/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                13336409,
				IDs:               nil,
				Alias:             "binance",
				Name:              "Binance",
				Followers:         "824K",
				Employees:         "5K-10K",
				AssociatedMembers: "10,428",
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
							URL:                  "https://www.linkedin.com/jobs/view/4179397668/",
							Date:                 mustDate("2025-03-11"),
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
			ShortDescription: "Digital Asset Exchange",
			Ignore:           true, // Cryptocurrency
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Crossover",
			Website: "https://www.crossover.com/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                9308035,
				IDs:               nil,
				Alias:             "crossover",
				Name:              "Crossover",
				Followers:         "8M",
				Employees:         "5K-10K",
				AssociatedMembers: "2,742",
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
				domain.Go:      {},
				domain.Rust:    {},
				domain.Zig:     {},
				domain.Scala:   {},
				domain.Elixir:  {},
				domain.Clojure: {},
				domain.Haskell: {},
			},
			ShortDescription: "Recruitment company",
			Ignore:           true,
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Optable",
			Website: "http://optable.co/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                64779622,
				IDs:               nil,
				Alias:             "optableco",
				Name:              "Optable",
				Followers:         "3K",
				Employees:         "11-50",
				AssociatedMembers: "57",
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
							Title:                "Experienced Backend Developer (Golang)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4178647999/",
							Date:                 mustDate("2025-03-11"),
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
			ShortDescription: "Optable is an identity management and data collaboration platform",
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "ScaleOps",
			Website: "https://scaleops.com/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                79471518,
				IDs:               nil,
				Alias:             "scaleops-sh",
				Name:              "ScaleOps - Cloud-Native Optimization",
				Followers:         "9K",
				Employees:         "51-200",
				AssociatedMembers: "62",
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
							Title:                "Senior Backend Engineer Go",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4181796179/",
							Date:                 mustDate("2025-03-11"),
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
			ShortDescription: "Effortlessly reduce cloud costs by up to 80% by automatically optimizing and scaling Kubernetes resources in runtime",
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "bet365",
			Website: "https://www.bet365careers.com/en",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                40163,
				IDs:               nil,
				Alias:             "bet365",
				Name:              "bet365",
				Followers:         "237K",
				Employees:         "5K-10K",
				AssociatedMembers: "6,619",
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
							Title:                "Software Engineer (Go), Sports Platform",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4179415502/",
							Date:                 mustDate("2025-03-11"),
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
			Ignore:           true, // Gambling
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Zenport",
			Website: "https://zenport.io/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                13725003,
				IDs:               nil,
				Alias:             "zenport",
				Name:              "Zenport Inc.",
				Followers:         "2K",
				Employees:         "11-50",
				AssociatedMembers: "10",
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
							Title:                "Senior Backend Developer (Go)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4181727305/",
							Date:                 mustDate("2025-03-12"),
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
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "FirstBatch",
			Website: "https://www.firstbatch.xyz/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                80364972,
				IDs:               nil,
				Alias:             "firstbatchai",
				Name:              "FirstBatch",
				Followers:         "6K",
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
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4178518636/",
							Date:                 mustDate("2025-03-11"),
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
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Syndica",
			Website: "https://www.syndica.io/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                76168224,
				IDs:               nil,
				Alias:             "syndica-io",
				Name:              "Syndica",
				Followers:         "2K",
				Employees:         "11-50",
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
							Title:                "Senior Rust Software Engineer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4179432496/",
							Date:                 mustDate("2025-03-11"),
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
			Ignore:           true, // Web3
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Nelly Solutions",
			Website: "http://www.getnelly.de/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                74522753,
				IDs:               nil,
				Alias:             "nelly-solutions",
				Name:              "Nelly Solutions",
				Followers:         "6K",
				Employees:         "51-200",
				AssociatedMembers: "132",
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
							Title:                "(Senior) Software Engineer, Rust",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4125742838/",
							Date:                 mustDate("2025-03-11"),
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
			ShortDescription: "Nelly is a healthcare payments company",
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Polygon Labs",
			Website: "https://polygon.technology/",
			Careers: "https://polygon.technology/careers/",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                13449964,
				IDs:               nil,
				Alias:             "polygonlabs",
				Name:              "Polygon Labs",
				Followers:         "191K",
				Employees:         "201-500",
				AssociatedMembers: "447",
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
							Title:                "Senior Rust Engineer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4102719913/",
							Date:                 mustDate("2025-03-11"),
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
			Ignore:           true, // Blockchain
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Skyworks Solutions",
			Website: "http://www.skyworksinc.com/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                165998,
				IDs:               nil,
				Alias:             "skyworks-solutions-inc",
				Name:              "Skyworks Solutions, Inc.",
				Followers:         "82K",
				Employees:         "5K-10K",
				AssociatedMembers: "5,202",
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
							URL:                  "https://www.linkedin.com/jobs/view/4027089913/",
							Date:                 mustDate("2025-03-11"),
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
			ShortDescription: "American semiconductor company",
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "F5",
			Website: "https://www.f5.com/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                4841,
				IDs:               nil,
				Alias:             "f5",
				Name:              "F5",
				Followers:         "358K",
				Employees:         "5K-10K",
				AssociatedMembers: "6,141",
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
							Title:                "Senior Rust Developer",
							ShortDescription:     "Proxy Solution",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4153485948/",
							Date:                 mustDate("2025-03-09"),
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
			ShortDescription: "F5 is a multi-cloud application services and security company",
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Megaport",
			Website: "https://www.megaport.com/",
			Careers: "https://www.megaport.com/careers/",
			About:   "https://www.megaport.com/about-megaport/",
			Blog:    "https://www.megaport.com/blog/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                3225864,
				IDs:               nil,
				Alias:             "megaport",
				Name:              "Megaport",
				Followers:         "32K",
				Employees:         "201-500",
				AssociatedMembers: "377",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "megaport",
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
							Title:                "Software Engineer (Go)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4181511391/",
							Date:                 mustDate("2025-03-13"),
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
			ShortDescription: "Australian AWS Direct Connect eco-system",
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "FairMoney",
			Website: "https://fairmoney.io/",
			Careers: "",
			About:   "https://fairmoney.io/about/",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                11707115,
				IDs:               nil,
				Alias:             "fairmoney",
				Name:              "FairMoney",
				Followers:         "142K",
				Employees:         "501-1K",
				AssociatedMembers: "1,029",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "fairmoney",
				Verified: false,
			},
			BlindProfile: domain.BlindProfile{
				Alias: "",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "fairmoney",
				Employees: "570",
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
							Title:                "Software Engineer — Backend (Golang)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4178759491/",
							Date:                 mustDate("2025-03-12"),
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
			ShortDescription: "Nigerian mobile banking for private and business borrowers",
			Industries: []domain.Industry{
				domain.IndustryFinTech,
			},
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Zalopay",
			Website: "https://zalopay.vn/",
			Careers: "",
			About:   "",
			Blog:    "https://engineering.zalopay.vn/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                31297705,
				IDs:               nil,
				Alias:             "zalopay",
				Name:              "Zalopay",
				Followers:         "44K",
				Employees:         "501-1K",
				AssociatedMembers: "406",
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
							Title:                "Senior Software Engineer — Golang",
							ShortDescription:     "Merchant Platform",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4180086053/",
							Date:                 mustDate("2025-03-12"),
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
			ShortDescription: "Vietnamese mobile wallet and payment platform associated with the Zalo messaging app",
			Industries: []domain.Industry{
				domain.IndustryFinTech,
			},
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "TD",
			Website: "https://www.td.com/",
			Careers: "https://careers.td.com/",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                2775,
				IDs:               nil,
				Alias:             "td",
				Name:              "TD",
				Followers:         "900K",
				Employees:         "10K+",
				AssociatedMembers: "97,091",
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
				Alias:     "td-bank",
				Employees: "26.000",
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
							Title:                "Software Engineer II (Scala/Spark)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4182972537/",
							Date:                 mustDate("2025-03-12"),
							WithSalary:           true, // $76,800 - $115,200 CAD per year
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
			ShortDescription: "Canadian bank who offers a full range of financial products and services",
			Industries: []domain.Industry{
				domain.IndustryFinTech,
			},
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "AI Rudder",
			Website: "https://airudder.com/",
			Careers: "https://airudder.com/careers/",
			About:   "https://airudder.com/about-us/",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                47635489,
				IDs:               nil,
				Alias:             "airudder",
				Name:              "AI Rudder",
				Followers:         "16K",
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
							Title:                "Software Engineer (Go/Golang)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4184051342/",
							Date:                 mustDate("2025-03-13"),
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
			ShortDescription: "AI startup who provides AI-powered voice solutions to improve B2C communications",
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Bitpanda",
			Website: "https://www.bitpanda.com/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                18164565,
				IDs:               nil,
				Alias:             "bitpanda",
				Name:              "Bitpanda",
				Followers:         "68K",
				Employees:         "501-1K",
				AssociatedMembers: "744",
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
							Title:                "Senior Software Engineer, Golang (Digital Asset Infrastructure)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4168215806/",
							Date:                 mustDate("2025-03-13"),
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
			Ignore:           true, // Crypto
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Guidewire Software",
			Website: "http://www.guidewire.com/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                9624,
				IDs:               nil,
				Alias:             "guidewire-software",
				Name:              "Guidewire Software",
				Followers:         "261K",
				Employees:         "1K-5K",
				AssociatedMembers: "3,597",
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
							Title:                "Software Engineer III (Golang)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4184033707/",
							Date:                 mustDate("2025-03-13"),
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
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Makro PRO",
			Website: "https://www.makro.pro/",
			Careers: "https://www.makrodigitalcareer.com/",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                76238800,
				IDs:               nil,
				Alias:             "makropro",
				Name:              "Makro PRO",
				Followers:         "67K",
				Employees:         "51-200",
				AssociatedMembers: "131",
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
							URL:                  "https://www.linkedin.com/jobs/view/4138748551/",
							Date:                 mustDate("2025-03-13"),
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
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Infoblox",
			Website: "http://www.infoblox.com/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                8697,
				IDs:               nil,
				Alias:             "infoblox",
				Name:              "Infoblox",
				Followers:         "172K",
				Employees:         "1K-5K",
				AssociatedMembers: "2,608",
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
							Title:                "Senior Software Engineer (Golang)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4183852152/",
							Date:                 mustDate("2025-03-13"),
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
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "NCR Voyix",
			Website: "https://www.ncrvoyix.com/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                96157938,
				IDs:               nil,
				Alias:             "ncrvoyix",
				Name:              "NCR Voyix",
				Followers:         "112K",
				Employees:         "10K+",
				AssociatedMembers: "6,668",
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
							Title:                "Software Engineer II (Golang – Kubernetes)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4137578941/",
							Date:                 mustDate("2025-03-13"),
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
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Cisco",
			Website: "http://www.cisco.com/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                1063,
				IDs:               []int{1063, 3805, 11434, 92950, 687352, 2792574, 3517498, 72566046},
				Alias:             "cisco",
				Name:              "Cisco",
				Followers:         "7M",
				Employees:         "10K+",
				AssociatedMembers: "95,120",
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
							Title:                "Software Engineer – FullStack | Golang | Java",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4183841211/",
							Date:                 mustDate("2025-03-13"),
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
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Priority Software",
			Website: "https://www.priority-software.com/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                89771,
				IDs:               nil,
				Alias:             "prioritysoftware",
				Name:              "Priority Software",
				Followers:         "19K",
				Employees:         "201-500",
				AssociatedMembers: "405",
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
							Title:                "Golang Senior Developer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4178758249/",
							Date:                 mustDate("2025-03-12"),
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
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Aspire Software",
			Website: "https://www.aspiresoftware.com/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                52129049,
				IDs:               nil,
				Alias:             "aspire-software",
				Name:              "Aspire Software",
				Followers:         "36K",
				Employees:         "1K-5K",
				AssociatedMembers: "248",
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
							Title:                "Senior Software Developer – Golang",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4182723217/",
							Date:                 mustDate("2025-03-12"),
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
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Valsoft Corporation",
			Website: "http://www.valsoftcorp.com/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                10396649,
				IDs:               nil,
				Alias:             "valsoft-corporation",
				Name:              "Valsoft Corporation",
				Followers:         "60K",
				Employees:         "1K-5K",
				AssociatedMembers: "336",
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
							Title:                "Senior Software Developer – Golang",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4182769280/",
							Date:                 mustDate("2025-03-12"),
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
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Flutter Entertainment",
			Website: "http://www.flutter.com/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                14072241,
				IDs:               nil,
				Alias:             "flutter-entertainment",
				Name:              "Flutter Entertainment",
				Followers:         "64K",
				Employees:         "10K+",
				AssociatedMembers: "7,847",
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
							Title:                "Software Engineer (Scala)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4158759621/",
							Date:                 mustDate("2025-03-13"),
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
			Ignore:           true, // Gambling
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "NielsenIQ",
			Website: "https://nielseniq.com/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                69088863,
				IDs:               []int{1718, 419609, 1364767, 3126409, 18360338, 69088863},
				Alias:             "nielseniq",
				Name:              "NielsenIQ",
				Followers:         "1M",
				Employees:         "10K+",
				AssociatedMembers: "25,585",
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
							Title:                "Data Engineer (Python, Scala, Azure)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4181286159/",
							Date:                 mustDate("2025-03-14"),
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
			Name:    "Adform",
			Website: "https://adform.com/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                72240,
				IDs:               nil,
				Alias:             "adform",
				Name:              "Adform",
				Followers:         "42K",
				Employees:         "501-1K",
				AssociatedMembers: "858",
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
							Title:                "Software Engineer, Java/Scala",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/3885026236/",
							Date:                 mustDate("2025-03-13"),
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
			Name:    "Avathon",
			Website: "http://www.avathon.com/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                5155679,
				IDs:               nil,
				Alias:             "avathonai",
				Name:              "Avathon",
				Followers:         "45K",
				Employees:         "201-500",
				AssociatedMembers: "302",
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
							Title:                "Software Engineer (Scala Back-End)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4053922033/",
							Date:                 mustDate("2025-03-11"),
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
			Name:    "Kraken Digital Asset Exchange",
			Website: "https://kraken.com/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                3787845,
				IDs:               []int{3787845, 69222723},
				Alias:             "krakenfx",
				Name:              "Kraken Digital Asset Exchange",
				Followers:         "208K",
				Employees:         "1K-5K",
				AssociatedMembers: "851",
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
							Title:                "Senior Software Engineer – Rust",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4137684962/",
							Date:                 mustDate("2025-03-14"),
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
			ShortDescription: "",
			Ignore:           true, // Crypto
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Cryptio",
			Website: "https://cryptio.co/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                11423070,
				IDs:               nil,
				Alias:             "cryptio",
				Name:              "Cryptio",
				Followers:         "8K",
				Employees:         "11-50",
				AssociatedMembers: "97",
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
							Title:                "Senior Rust Engineer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4183189232/",
							Date:                 mustDate("2025-03-13"),
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
			Ignore:           true, // Crypto
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "RingCentral",
			Website: "https://www.ringcentral.com/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                60868,
				IDs:               nil,
				Alias:             "ringcentral",
				Name:              "RingCentral",
				Followers:         "254K",
				Employees:         "5K-10K",
				AssociatedMembers: "6,328",
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
							Title:                "Rust developer",
							ShortDescription:     "Core Media AI",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4168118534/",
							Date:                 mustDate("2025-03-13"),
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
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "QNX",
			Website: "https://blackberry.qnx.com/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                9889,
				IDs:               nil,
				Alias:             "qnx",
				Name:              "QNX",
				Followers:         "26K",
				Employees:         "201-500",
				AssociatedMembers: "704",
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
							Title:                "Rust Systems Software Developer II",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4151541268/",
							Date:                 mustDate("2025-03-11"),
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
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Solaris",
			Website: "https://www.solarisgroup.com/",
			Careers: "https://www.solarisgroup.com/careers/",
			About:   "https://www.solarisgroup.com/about/",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                10533980,
				IDs:               nil,
				Alias:             "solariscompany",
				Name:              "Solaris SE",
				Followers:         "55K",
				Employees:         "201-500",
				AssociatedMembers: "683",
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
							Title:                "Senior Engineer Golang (f/m/d) at Solaris",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4181915311/",
							Date:                 mustDate("2025-03-14"),
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
			Name:    "smava",
			Website: "https://www.smava.de/",
			Careers: "https://www.smava.de/jobs",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                643119,
				IDs:               nil,
				Alias:             "smava",
				Name:              "smava",
				Followers:         "7K",
				Employees:         "501-1K",
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
							Title:                "Senior Software Engineer – Golang & NodeJS",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4184824435/",
							Date:                 mustDate("2025-03-14"),
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
			Name:    "CoinGate",
			Website: "https://coingate.com/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                10586817,
				IDs:               nil,
				Alias:             "coingate",
				Name:              "CoinGate",
				Followers:         "9K",
				Employees:         "51-200",
				AssociatedMembers: "81",
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
							URL:                  "https://www.linkedin.com/jobs/view/4184624827/",
							Date:                 mustDate("2025-03-14"),
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
			Ignore:           true, // Crypto
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Kiln",
			Website: "https://www.kiln.fi/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                28150174,
				IDs:               nil,
				Alias:             "kiln-fi",
				Name:              "Kiln",
				Followers:         "8K",
				Employees:         "51-200",
				AssociatedMembers: "124",
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
							Title:                "Senior Backend Engineer – Golang",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4184623805/",
							Date:                 mustDate("2025-03-14"),
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
			Ignore:           true, // Crypto
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Syniti",
			Website: "http://www.syniti.com/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                19112642,
				IDs:               []int{2724939, 10417630, 19112642, 69646738},
				Alias:             "synitidata",
				Name:              "Syniti",
				Followers:         "67K",
				Employees:         "1K-5K",
				AssociatedMembers: "1,452",
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
							Title:                "Software Engineer ll (Golang)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4181840665/",
							Date:                 mustDate("2025-03-14"),
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
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Ringover",
			Website: "https://www.ringover.com/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                79585140,
				IDs:               []int{9180058, 79583448, 79585140},
				Alias:             "ringover",
				Name:              "Ringover North America",
				Followers:         "9K",
				Employees:         "201-500",
				AssociatedMembers: "386",
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
							Title:                "Software Engineer Golang (Senior)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4184269668/",
							Date:                 mustDate("2025-03-14"),
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
			Ignore:           true, // Waiting for the English version of the vacancy description
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Goldman Sach",
			Website: "http://www.goldmansachs.com/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                1382,
				IDs:               nil,
				Alias:             "goldman-sachs",
				Name:              "Goldman Sachs",
				Followers:         "5M",
				Employees:         "10K+",
				AssociatedMembers: "60,944",
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
							Title:                "Golang Software Engineer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4180220865/",
							Date:                 mustDate("2025-03-13"),
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
			Name:    "Chợ Tốt",
			Website: "https://chotot.com/",
			Careers: "https://careers.chotot.com/",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                3780320,
				IDs:               []int{3780320, 78359837},
				Alias:             "cho-tot",
				Name:              "Chợ Tốt",
				Followers:         "28K",
				Employees:         "201-500",
				AssociatedMembers: "232",
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
							Title:                "Backend Engineer (Java, Golang, Python)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4178756795/",
							Date:                 mustDate("2025-03-12"),
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
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Simulmedia",
			Website: "https://www.simulmedia.com/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                237095,
				IDs:               nil,
				Alias:             "simulmedia",
				Name:              "Simulmedia",
				Followers:         "13K",
				Employees:         "51-200",
				AssociatedMembers: "89",
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
							URL:                  "https://www.linkedin.com/jobs/view/4182751037/",
							Date:                 mustDate("2025-03-12"),
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
			ShortDescription: "Simulmedia is a cross-channel TV advertising platform",
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Mindrift",
			Website: "https://mindrift.ai/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                101576014,
				IDs:               nil,
				Alias:             "mindrift-ai",
				Name:              "Mindrift",
				Followers:         "380K",
				Employees:         "501-1K",
				AssociatedMembers: "1,135",
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
							Title:                "Software Developer (Golang)",
							ShortDescription:     "AI Tutor",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4182668629/",
							Date:                 mustDate("2025-03-14"),
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
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Hewlett Packard Enterprise",
			Website: "http://hpe.com/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                1025,
				IDs:               []int{1025, 162533},
				Alias:             "hewlett-packard-enterprise",
				Name:              "Hewlett Packard Enterprise",
				Followers:         "4M",
				Employees:         "10K+",
				AssociatedMembers: "74,605",
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
							Title:                "Cloud Developer (Golang)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4184377636/",
							Date:                 mustDate("2025-03-14"),
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
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "IONOS",
			Website: "https://www.ionos.com/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                10298,
				IDs:               nil,
				Alias:             "ionos",
				Name:              "IONOS",
				Followers:         "63K",
				Employees:         "1K-5K",
				AssociatedMembers: "2,599",
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
							Title:                "GO/Golang Developer (DNS)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4182960038/",
							Date:                 mustDate("2025-03-12"),
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
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Relai",
			Website: "https://relai.app/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                19078807,
				IDs:               nil,
				Alias:             "relai-app",
				Name:              "Relai",
				Followers:         "14K",
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
					Vacancies: []domain.Vacancy{
						{
							Title:                "Senior Go Developer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4157760436/",
							Date:                 mustDate("2025-03-14"),
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
			Ignore:           true, // Crypto
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Velmie",
			Website: "https://www.velmie.com/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                18778012,
				IDs:               nil,
				Alias:             "velmie",
				Name:              "Velmie",
				Followers:         "2K",
				Employees:         "51-200",
				AssociatedMembers: "47",
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
							URL:                  "https://www.linkedin.com/jobs/view/4184379493/",
							Date:                 mustDate("2025-03-14"),
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
			Industries: []domain.Industry{
				domain.IndustryFinTech,
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Clinia",
			Website: "https://clinia.com/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                11160528,
				IDs:               nil,
				Alias:             "clinia",
				Name:              "Clinia",
				Followers:         "3K",
				Employees:         "11-50",
				AssociatedMembers: "41",
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
							URL:                  "https://www.linkedin.com/jobs/view/4183894600/",
							Date:                 mustDate("2025-03-14"),
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
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Keyrock",
			Website: "https://keyrock.com/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                11259345,
				IDs:               nil,
				Alias:             "keyrock",
				Name:              "Keyrock",
				Followers:         "20K",
				Employees:         "51-200",
				AssociatedMembers: "166",
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
							Title:                "Senior Software Engineer – High Frequency Trading (C++ OR Rust)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4180667729/",
							Date:                 mustDate("2025-03-14"),
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
			Ignore:           true, // Crypto
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Lazer Technologies",
			Website: "https://www.lazertechnologies.com/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                37569333,
				IDs:               nil,
				Alias:             "lazer-technologies",
				Name:              "Lazer Technologies",
				Followers:         "8K",
				Employees:         "51-200",
				AssociatedMembers: "125",
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
							Title:                "Senior Full Stack Rust Engineer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4184889475/",
							Date:                 mustDate("2025-03-15"),
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
			Ignore:           true, // Crypto
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Cloudwick",
			Website: "https://cloudwick.com/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                2043823,
				IDs:               nil,
				Alias:             "cloudwick",
				Name:              "Cloudwick",
				Followers:         "13K",
				Employees:         "201-500",
				AssociatedMembers: "143",
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
							Title:                "Data Engineer (Health Domain)",
							ShortDescription:     "Design, develop, and maintain ETL processes and data pipelines with Scala/PySpark",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4182758670/",
							Date:                 mustDate("2025-03-13"),
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
			Name:    "FIS",
			Website: "http://www.fisglobal.com/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                3128,
				IDs:               []int{3128, 804638, 76916563},
				Alias:             "fis",
				Name:              "FIS",
				Followers:         "756K",
				Employees:         "10K+",
				AssociatedMembers: "46,542",
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
							Title:                "Lead Engineer – Development (BigData, Scala)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4172072369/",
							Date:                 mustDate("2025-03-12"),
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
			Name:    "Nike",
			Website: "https://nike.com/",
			Careers: "https://careers.nike.com/",
			About:   "https://about.nike.com/",
			Blog:    "https://engineering.nike.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                2029,
				IDs:               []int{2029, 10818},
				Alias:             "nike",
				Name:              "Nike",
				Followers:         "6M",
				Employees:         "10K+",
				AssociatedMembers: "81,633",
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
				Alias:     "nike",
				Employees: "99,580",
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
							Title:                "Software Engineer II — Platform Engineering",
							ShortDescription:     "3+ years of professional software development experience",
							SwitchingOpportunity: "",
							URL:                  "https://www.indeed.com/viewjob?jk=6ea60544f1c6f982&",
							Date:                 mustDate("2025-03-18"),
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
			ShortDescription: "American athletic footwear and apparel corporation",
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "CVS Health",
			Website: "https://www.cvshealth.com/",
			Careers: "https://jobs.cvshealth.com/",
			About:   "https://www.cvshealth.com/about.html",
			Blog:    "https://medium.com/cvs-health-tech-blog",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                4680,
				IDs:               []int{2807, 4680, 13059, 11098302},
				Alias:             "cvshealth",
				Name:              "CVS Health",
				Followers:         "972K",
				Employees:         "10K+",
				AssociatedMembers: "123,272",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "cvs-health",
				Verified: false,
			},
			BlindProfile: domain.BlindProfile{
				Alias: "",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "cvs-health",
				Employees: "300,000",
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
							Title:                "Software Development Engineer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.indeed.com/viewjob?jk=b5208bb10cbb651b&",
							Date:                 mustDate("2025-03-18"),
							WithSalary:           true, // $72,100 - $144,200 per year
							Remote:               false,
						},
						{
							Title:                "Software Engineer III",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.indeed.com/viewjob?jk=eff1f9f3ed90066d&",
							Date:                 mustDate("2025-03-18"),
							WithSalary:           true, // $92,700 - $185,400 per year
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
			ShortDescription: "American healthcare/retail  corporation",
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Airlines Reporting Corporation (ARC)",
			Website: "https://www2.arccorp.com/",
			Careers: "https://www2.arccorp.com/about-us/careers/?utm_source=Global_Navigation",
			About:   "https://www2.arccorp.com/about-us/our-story/?utm_source=Global_Navigation",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                6509,
				IDs:               nil,
				Alias:             "arc",
				Name:              "Airlines Reporting Corporation (ARC)",
				Followers:         "19K",
				Employees:         "201-500",
				AssociatedMembers: "3,250",
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
				Alias:     "airlines-reporting",
				Employees: "2,750",
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
							Title:                "Manager, Software Engineering",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.indeed.com/viewjob?jk=ff3ad58cfe2f8295&",
							Date:                 mustDate("2025-03-18"),
							WithSalary:           true, // $123,000 - $204,300 per year
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
			ShortDescription: "American intelligence behind air travel, connecting the industry ecosystem and powering commercial decisions for airlines and their partners",
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Epsilon",
			Website: "https://www.epsilon.com/",
			Careers: "https://www.epsilon.com/emea/careers",
			About:   "https://www.epsilon.com/emea/about-us/",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                3776,
				IDs:               []int{3776, 3648054},
				Alias:             "epsilon",
				Name:              "Epsilon",
				Followers:         "464K",
				Employees:         "5K-10K",
				AssociatedMembers: "11,167",
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
				Alias:     "Epsilon",
				Employees: "7,500",
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
							Title:                "Lead, System and Platform Operations",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.indeed.com/viewjob?jk=788c43c55093c166&",
							Date:                 mustDate("2025-03-18"),
							WithSalary:           true, // $115,000 - $130,000 per year
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
			ShortDescription: "American global data, technology and services company",
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Travelers",
			Website: "https://www.travelers.com/",
			Careers: "https://www.travelers.com/careers",
			About:   "https://www.travelers.com/about-travelers",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                3088,
				IDs:               nil,
				Alias:             "travelers",
				Name:              "Travelers",
				Followers:         "426K",
				Employees:         "10K+",
				AssociatedMembers: "37,299",
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
				Alias:     "travelers",
				Employees: "31,000",
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
							Title:                "Data Engineer I (AWS/Snowflake/Databricks/Abinitio)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.indeed.com/viewjob?jk=322ff7ac1003f649&",
							Date:                 mustDate("2025-03-18"),
							WithSalary:           true, // $106,300 - $175,400 per year
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
			ShortDescription: "American insurance and financial company",
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Aviatrix",
			Website: "https://aviatrix.com/",
			Careers: "https://aviatrix.com/careers/",
			About:   "https://aviatrix.com/about-us/",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                6440846,
				IDs:               nil,
				Alias:             "aviatrix-systems",
				Name:              "Aviatrix",
				Followers:         "38K",
				Employees:         "201-500",
				AssociatedMembers: "495",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "aviatrixsystems",
				Verified: false,
			},
			BlindProfile: domain.BlindProfile{
				Alias: "",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "aviatrix",
				Employees: "10",
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
							Title:                "Principal Architect/Technical Lead Manager",
							ShortDescription:     "Site Reliability Engineering",
							SwitchingOpportunity: "",
							URL:                  "https://www.indeed.com/viewjob?jk=7843b69900f3a0ab&",
							Date:                 mustDate("2025-03-18"),
							WithSalary:           true, // $244,000 - $265,000 per year
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
			ShortDescription: "American multi-cloud networking software who provides connect, manage, and secure their connections between cloud providers",
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
