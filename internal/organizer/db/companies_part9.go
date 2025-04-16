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
				Alias:       "forcepoint",
				Employees:   "1,001 to 5,000 Employees",
				Salary:      "$75K ~ $230K a year",
				Reviews:     "25",
				ReviewsRate: "3.0",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "forcepoint",
				Employees: "2,250",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Forcepoint-EI_IE1262390.11,21.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Forcepoint-Reviews-E1262390.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Forcepoint-Jobs-E1262390.htm",
				Jobs:        "27",
				Reviews:     "1.3K",
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
							Title:                "Principal Engineer (Golang, Java, AWS & Data Engineering)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4157762453/",
							Date:                 mustDate("2025-03-06"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Senior Software Engineer – Golang",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4169998476/",
							Date:                 mustDate("2025-03-24"), // mustDate("2025-03-18"),
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
				Alias:       "Barracuda",
				Employees:   "1,001 to 5,000",
				Salary:      "$53K ~ $335K a year",
				Reviews:     "20",
				ReviewsRate: "3.6",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "barracuda-networks",
				Employees: "1,500",
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
						{
							Title:                "Software Engineer – Golang",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4192226484/",
							Date:                 mustDate("2025-03-24"),
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
				Login:     "teradata",
				Followers: "106",
				Verified:  false,
			},
			BlindProfile: domain.BlindProfile{
				Alias:       "teradata",
				Employees:   "10,000+",
				Salary:      "$50K ~ $252K a year",
				Reviews:     "122",
				ReviewsRate: "3.3",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "teradata",
				Employees: "10,760",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Teradata-EI_IE14638.11,19.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Teradata-Reviews-E14638.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Teradata-Jobs-E14638.htm",
				Jobs:        "47",
				Reviews:     "3.7K",
				Salaries:    "5.4K",
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
						{
							Title:                "Senior Golang Software Engineer",
							ShortDescription:     "2+ years experience in Golang",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4198065578/",
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
			ShortDescription: "Teramind Tracker is an employee monitoring and insider threat detection software that records screenshots, audio/video, keystrokes, and other activities",
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
				domain.Rust:  {},
				domain.Zig:   {},
				domain.Scala: {},
				domain.Elixir: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Elixir Developer",
							ShortDescription:     "Business Automation",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4153262881/",
							Date:                 mustDate("2025-04-01"),
							WithSalary:           false,
							Remote:               false,
						},
					},
				},
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
						{
							Title:                "Senior Full Stack Engineer — Climate Tech — Rust & TypeScript",
							ShortDescription:     "Experience using Rust in production or side projects with significant scope",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4200870310/",
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
						{
							Title:                "Senior Backend Engineer (Scala & Elasticsearch)",
							ShortDescription:     "5+ years experience as Software Engineer working with Scala",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4200235969/",
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
				Login:     "Viasat",
				Followers: "30",
				Verified:  true,
			},
			BlindProfile: domain.BlindProfile{
				Alias:       "Viasat",
				Employees:   "5,001 to 10,000",
				Salary:      "$43K ~ $874K a year",
				Reviews:     "140",
				ReviewsRate: "3.6",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "viasat",
				Employees: "6,000",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Viasat-EI_IE5500.11,17.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Viasat-Reviews-E5500.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Viasat-Jobs-E5500.htm",
				Jobs:        "419",
				Reviews:     "1.9K",
				Salaries:    "3.6K",
				ReviewsRate: "3.6",
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
				Followers:         "312K",
				Employees:         "10K+", // https://www.linkedin.com/search/results/people/?currentCompany=%5B%224298%22%5D
				AssociatedMembers: "15,000",
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
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-University-of-Virginia-EI_IE8336.11,33.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/University-of-Virginia-Reviews-E8336.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/University-of-Virginia-Jobs-E8336.htm",
				Jobs:        "579",
				Reviews:     "2.5K",
				Salaries:    "5.9K",
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
							Title:                "Senior Golang Developer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4177332766/",
							Date:                 mustDate("2025-03-13"), // mustDate("2025-03-07"),
							WithSalary:           true,                   // $110,000 - $130,000 per year
							Remote:               false,
						},
						{
							Title:                "Senior Golang Developer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4205310300/",
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
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Mr-Price-EI_IE224003.11,19.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Mr-Price-Reviews-E224003.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Mr-Price-Jobs-E224003.htm",
				Jobs:        "118",
				Reviews:     "382",
				Salaries:    "438",
				ReviewsRate: "3.6",
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
				Login:     "LANL",
				Followers: "487",
				Verified:  true,
			},
			BlindProfile: domain.BlindProfile{
				Alias: "",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "los-alamos-national-laboratory",
				Employees: "12,752",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Los-Alamos-National-Laboratory-EI_IE35234.11,41.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Los-Alamos-National-Laboratory-Reviews-E35234.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Los-Alamos-National-Laboratory-Jobs-E35234.htm",
				Jobs:        "345",
				Reviews:     "1.2K",
				Salaries:    "2.9K",
				ReviewsRate: "4.1",
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
							Title:                "DevOps Python/Rust Developer (Scientist 1/2)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4181750667/",
							Date:                 mustDate("2025-04-01"), // mustDate("2025-03-11"),
							WithSalary:           true,                   // Scientist 1 ($94,5-$154,600), Scientist 2 ($104,100-$172,200)
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
						{
							Title:                "Senior Backend Developer (Go)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4206443530/",
							Date:                 mustDate("2025-04-12"),
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
							Date:                 mustDate("2025-04-14"), // mustDate("2025-04-04"), // mustDate("2025-03-21"), // mustDate("2025-03-11"),
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
				Login:     "f5networks",
				Followers: "280",
				Verified:  false,
			},
			BlindProfile: domain.BlindProfile{
				Alias:       "F5-Networks",
				Employees:   "1,001 to 5,000",
				Salary:      "$71K ~ $265K a year",
				Reviews:     "292",
				ReviewsRate: "3.5",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "f5-networks",
				Employees: "6,500",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-F5-EI_IE9222.11,13.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/F5-Reviews-E9222.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/F5-Jobs-E9222.htm",
				Jobs:        "173",
				Reviews:     "2K",
				Salaries:    "3.9K",
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
							Title:                "Senior Rust Developer",
							ShortDescription:     "Proxy Solution",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4153485948/",
							Date:                 mustDate("2025-03-31"), // mustDate("2025-03-09"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Senior Rust Developer",
							ShortDescription:     "Proxy Solution",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4190616076/",
							Date:                 mustDate("2025-03-21"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Senior Rust Developer",
							ShortDescription:     "Proxy Solution",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4198709559/",
							Date:                 mustDate("2025-04-06"),
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
							Date:                 mustDate("2025-04-02"), // mustDate("2025-03-12"),
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
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-TD-EI_IE3767.11,13.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/TD-Reviews-E3767.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/TD-Jobs-E3767.htm",
				Jobs:        "1.3K",
				Reviews:     "23K",
				Salaries:    "35K",
				ReviewsRate: "3.9",
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
						{
							Title:                "API Engineer — Micro-Services",
							ShortDescription:     "Maintain existing codebases in languages like Go",
							SwitchingOpportunity: "Knowledge of modern programming languages (specifically Go is a plus)",
							URL:                  "https://www.indeed.com/viewjob?jk=8d7ce4fc383b9eac",
							Date:                 mustDate("2025-03-25"),
							WithSalary:           true, // $134,000 - $247,000 per year
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
						{
							Title:                "Senior DevOps Engineer (Golang)",
							ShortDescription:     "Strong proficiency in Golang and experience with writing production-grade software using Golang",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4154050315/",
							Date:                 mustDate("2025-04-02"),
							WithSalary:           false,
							Remote:               true,
						},
						{
							Title:                "Senior Software Engineer (Golang)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4138747697/",
							Date:                 mustDate("2025-04-05"),
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
						{
							Title:                "Sr. Software Engineer – Golang",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4071755601/",
							Date:                 mustDate("2025-03-19"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Software Engineer II – Edge (Golang Kubernetes)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4145763846/",
							Date:                 mustDate("2025-03-22"),
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
				Login:     "cisco",
				Followers: "986",
				Verified:  true,
			},
			BlindProfile: domain.BlindProfile{
				Alias:       "cisco",
				Employees:   "10,000+",
				Salary:      "$50K ~ $260K a year",
				Reviews:     "2,519",
				ReviewsRate: "3.6",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "cisco",
				Employees: "98,110",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Cisco-EI_IE1425.11,16.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Cisco-Reviews-E1425.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Cisco-Jobs-E1425.htm",
				Jobs:        "814",
				Reviews:     "37K",
				Salaries:    "111K",
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
							Title:                "Software Engineer – FullStack | Golang | Java",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4183841211/",
							Date:                 mustDate("2025-03-13"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Senior Software Engineer — Go",
							ShortDescription:     "3+ years of experience in Go",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4198239895/",
							Date:                 mustDate("2025-04-01"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Software Engineer – Golang Developer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4189691197/",
							Date:                 mustDate("2025-04-12"),
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
						{
							Title:                "Golang Senior Developer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4203313137/",
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
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Valsoft-EI_IE2564742.11,18.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Valsoft-Reviews-E2564742.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Valsoft-Jobs-E2564742.htm",
				Jobs:        "70",
				Reviews:     "167",
				Salaries:    "191",
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
				Alias:     "nielseniq",
				Employees: "14,430",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-NielsenIQ-EI_IE4796509.11,20.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/NielsenIQ-Reviews-E4796509.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/NielsenIQ-Jobs-E4796509.htm",
				Jobs:        "621",
				Reviews:     "4.4K",
				Salaries:    "7.7K",
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
						{
							Title:                "Java/Scala Developer | Spark & Hadoop",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4120994368/",
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
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-RingCentral-EI_IE197577.11,22.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/RingCentral-Reviews-E197577.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/RingCentral-Jobs-E197577.htm",
				Jobs:        "184",
				Reviews:     "2K",
				Salaries:    "2.7K",
				ReviewsRate: "3.4",
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
							Date:                 mustDate("2025-04-08"), // mustDate("2025-03-31"), // mustDate("2025-03-11"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Senior Systems Software Developer (Rust)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4151540281/",
							Date:                 mustDate("2025-04-11"),
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
							Title:                "Senior Engineer Golang",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4181915311/",
							Date:                 mustDate("2025-03-14"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Senior Engineer Golang",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4189553947/",
							Date:                 mustDate("2025-04-14"), // mustDate("2025-03-24"),
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
				Alias:       "Syniti",
				Employees:   "1,001 to 5,000",
				Salary:      "",
				Reviews:     "3",
				ReviewsRate: "4.3",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "syniti",
				Employees: "1,500",
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
				Alias:       "goldman-sachs",
				Employees:   "10,000+",
				Salary:      "$33K ~ $295K a year",
				Reviews:     "1,576",
				ReviewsRate: "3.4",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "goldman-sachs",
				Employees: "67,370",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Goldman-Sachs-EI_IE2800.11,24.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Goldman-Sachs-Reviews-E2800.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Goldman-Sachs-Jobs-E2800.htm",
				Jobs:        "2.3K",
				Reviews:     "22K",
				Salaries:    "52K",
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
							Title:                "Golang Software Engineer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4180220865/",
							Date:                 mustDate("2025-03-13"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Software Engineer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.indeed.com/viewjob?jk=7922a4920a8d87b5",
							Date:                 mustDate("2025-03-23"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Core Engineering, Golang Software Engineer, Analyst/ Associate",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4124674198/",
							Date:                 mustDate("2025-04-12"), // mustDate("2025-03-22"),
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
				domain.Rust: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Software Developer (Rust) — AI Tutor",
							ShortDescription:     "3 years of professional experience with Rust",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4182671096/",
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
				Login:     "HewlettPackard",
				Followers: "352",
				Verified:  false,
			},
			BlindProfile: domain.BlindProfile{
				Alias: "",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias: "",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Hewlett-Packard-Enterprise-HPE-EI_IE1093046.11,41.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Hewlett-Packard-Enterprise-HPE-Reviews-E1093046.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Hewlett-Packard-Enterprise-HPE-Jobs-E1093046.htm",
				Jobs:        "297",
				Reviews:     "20K",
				Salaries:    "32K",
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
				Name:              "Relai 🇨🇭",
				Followers:         "14K",
				Employees:         "11-50",
				AssociatedMembers: "64",
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
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-FIS-EI_IE313114.11,14.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/FIS-Reviews-E313114.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/FIS-Jobs-E313114.htm",
				Jobs:        "719",
				Reviews:     "14K",
				Salaries:    "21K",
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
				Login:     "nike-inc",
				Followers: "229",
				Verified:  false,
			},
			BlindProfile: domain.BlindProfile{
				Alias:       "Nike",
				Employees:   "10,000+",
				Salary:      "$65K ~ $280K a year",
				Reviews:     "484",
				ReviewsRate: "3.6",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "nike",
				Employees: "99,580",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-NIKE-EI_IE1699.11,15.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/NIKE-Reviews-E1699.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/NIKE-Jobs-E1699.htm",
				Jobs:        "1.3K",
				Reviews:     "16K",
				Salaries:    "26K",
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
							Title:                "Software Engineer II — Platform Engineering",
							ShortDescription:     "Experience developing Kubernetes controllers in Golang",
							SwitchingOpportunity: "",
							URL:                  "https://www.indeed.com/viewjob?jk=6ea60544f1c6f982",
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
							ShortDescription:     "Build and enhance automation tools and frameworks in Golang and Python to streamline operations",
							SwitchingOpportunity: "",
							URL:                  "https://www.indeed.com/viewjob?jk=7843b69900f3a0ab",
							Date:                 mustDate("2025-03-18"),
							WithSalary:           true, // $244,000 - $265,000 per year
							Remote:               true,
						},
						{
							Title:                "Staff Software Engineer – Golang Developer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4184984550/",
							Date:                 mustDate("2025-03-19"),
							WithSalary:           true, // $220,000-$226,000
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
			ShortDescription: "American multi-cloud networking software who provides connect, manage, and secure their connections between cloud providers",
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Firstbase",
			Website: "https://www.firstbase.io/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                12671153,
				IDs:               nil,
				Alias:             "firstbase-io",
				Name:              "Firstbase",
				Followers:         "14K",
				Employees:         "51-200",
				AssociatedMembers: "60",
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
							Title:                "Senior Full-Stack Engineer – Golang",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4183968735/",
							Date:                 mustDate("2025-03-18"),
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
			ShortDescription: "",
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "NatWest Group",
			Website: "https://natwestgroup.com/",
			Careers: "https://jobs.natwestgroup.com/",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                68360623,
				IDs:               []int{4777, 163231, 12643959, 29024025, 68360623},
				Alias:             "natwest-group",
				Name:              "NatWest Group",
				Followers:         "682K",
				Employees:         "10K+",
				AssociatedMembers: "41,426",
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
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-NatWest-Group-EI_IE10222.11,24.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/NatWest-Group-Reviews-E10222.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/NatWest-Group-Jobs-E10222.htm",
				Jobs:        "207",
				Reviews:     "12K",
				Salaries:    "19K",
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
							Title:                "Back-End Engineer – Golang",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4188231036/",
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
			ShortDescription: "",
			Industries: []domain.Industry{
				domain.IndustryFinTech,
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Commonwealth Bank",
			Website: "https://www.commbank.com.au/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                2848,
				IDs:               []int{2848, 2850, 2851, 270126},
				Alias:             "commonwealthbank",
				Name:              "Commonwealth Bank",
				Followers:         "600K",
				Employees:         "10K+",
				AssociatedMembers: "45,071",
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
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Commonwealth-Bank-of-Australia-EI_IE7922.11,41.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Commonwealth-Bank-of-Australia-Reviews-E7922.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Commonwealth-Bank-of-Australia-Jobs-E7922.htm",
				Jobs:        "350",
				Reviews:     "7.8K",
				Salaries:    "14K",
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
							Title:                "Senior Platform Engineer (AWS, Python, Golang)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4183953062/",
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
			ShortDescription: "",
			Industries: []domain.Industry{
				domain.IndustryFinTech,
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "IKEA",
			Website: "https://ikea.com/",
			Careers: "",
			About:   "https://about.ikea.com/",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                2743,
				IDs:               []int{2743, 9854},
				Alias:             "ikea",
				Name:              "IKEA",
				Followers:         "3M",
				Employees:         "10K+",
				AssociatedMembers: "88,900",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:     "IKEA",
				Followers: "40",
				Verified:  true,
			},
			BlindProfile: domain.BlindProfile{
				Alias:       "IKEA",
				Employees:   "10,000+",
				Salary:      "$41K ~ $185K a year",
				Reviews:     "45",
				ReviewsRate: "3.8",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "ikea",
				Employees: "87,550",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-IKEA-EI_IE3957.11,15.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/IKEA-Reviews-E3957.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/IKEA-Jobs-E3957.htm",
				Jobs:        "1.7K",
				Reviews:     "17K",
				Salaries:    "24K",
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
							Title:                "Junior Software Engineer — (Java/Golang) & GCP",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4185204926/",
							Date:                 mustDate("2025-03-17"),
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
			Name:    "nexos.ai",
			Website: "http://www.nexos.ai/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                105703851,
				IDs:               nil,
				Alias:             "nexos-ai",
				Name:              "nexos.ai",
				Followers:         "4K",
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
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Back-End Engineer | Middle-Senior | Go",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4188140675/",
							Date:                 mustDate("2025-03-18"),
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
			Name:    "GCash",
			Website: "http://wearegcash.com/",
			Careers: "http://wearegcash.com/careers",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                14693558,
				IDs:               nil,
				Alias:             "wearegcash",
				Name:              "GCash",
				Followers:         "213K",
				Employees:         "501-1K",
				AssociatedMembers: "3,024",
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
							URL:                  "https://www.linkedin.com/jobs/view/4183372178/",
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
			ShortDescription: "",
			Industries: []domain.Industry{
				domain.IndustryFinTech,
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "mailstep",
			Website: "http://www.mailstep.cz/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                10554555,
				IDs:               nil,
				Alias:             "mail-step",
				Name:              "mailstep",
				Followers:         "2K",
				Employees:         "201-500",
				AssociatedMembers: "96",
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
							Title:                "Software Engineer PHP/Go",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4187272758/",
							Date:                 mustDate("2025-03-17"),
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
			Name:    "Lexia Learning",
			Website: "https://www.lexialearning.com/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                47092,
				IDs:               nil,
				Alias:             "lexia-learning-systems",
				Name:              "Lexia Learning",
				Followers:         "73K",
				Employees:         "501-1K",
				AssociatedMembers: "877",
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
							Title:                "Senior Fullstack Software Engineer, Rust Focused",
							ShortDescription:     "Develop modular, intuitive and scalable components, services and APIs using Rust",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4183950416/",
							Date:                 mustDate("2025-03-18"),
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
				domain.IndustryEdTech,
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "CertiK",
			Website: "https://www.certik.com/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                11831043,
				IDs:               nil,
				Alias:             "certik",
				Name:              "CertiK",
				Followers:         "23K",
				Employees:         "201-500",
				AssociatedMembers: "184",
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
							Title:                "Blockchain Security Engineer – (Solidity / Rust / Golang – Senior Level)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4161545800/",
							Date:                 mustDate("2025-03-16"),
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
			ShortDescription: "",
			Ignore:           true, // Web3
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "HolidayCheck",
			Website: "http://holidaycheck.com/",
			Careers: "http://careers.holidaycheck.com/",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                289231,
				IDs:               nil,
				Alias:             "holidaycheck-ag",
				Name:              "HolidayCheck AG",
				Followers:         "8K",
				Employees:         "201-500",
				AssociatedMembers: "309",
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
							Title:                "Software Engineer – Scala",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4087274629/",
							Date:                 mustDate("2025-03-16"),
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
			Name:    "Dyad",
			Website: "https://dyad.net/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                67565681,
				IDs:               nil,
				Alias:             "dyad-ai",
				Name:              "Dyad",
				Followers:         "2K",
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
							URL:                  "https://www.linkedin.com/jobs/view/4188226908/",
							Date:                 mustDate("2025-03-18"),
							WithSalary:           false,
							Remote:               false,
						},
					},
				},
				domain.Clojure: {},
				domain.Haskell: {},
			},
			ShortDescription: "",
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Stylitics",
			Website: "http://www.stylitics.com/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                1616816,
				IDs:               nil,
				Alias:             "stylitics",
				Name:              "Stylitics",
				Followers:         "14K",
				Employees:         "51-200",
				AssociatedMembers: "177",
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
							Title:                "Clojure Backend Engineer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4181907066/",
							Date:                 mustDate("2025-03-15"),
							WithSalary:           false,
							Remote:               false,
						},
					},
				},
				domain.Haskell: {},
			},
			ShortDescription: "",
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "CERESTI HEALTH",
			Website: "http://ceresti.com/",
			Careers: "https://www.ceresti.com/careers",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                15403042,
				IDs:               nil,
				Alias:             "ceresti-health-inc.",
				Name:              "CERESTI HEALTH",
				Followers:         "1K",
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
				Alias: "Ceresti-Health",
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Senior Golang Backend Developer",
							ShortDescription:     "3+ years of experience with the Go",
							SwitchingOpportunity: "",
							URL:                  "https://www.indeed.com/viewjob?jk=ed0f373654d5bce5",
							Date:                 mustDate("2025-03-19"),
							WithSalary:           true, // From $135,000 per year
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
			ShortDescription: "A medical facility that helps clients with Alzheimer's",
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Amuse",
			Website: "https://www.amuse.io/",
			Careers: "https://careers.amuse.io/jobs",
			About:   "https://www.amuse.io/en/resources/about/",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                10607933,
				IDs:               nil,
				Alias:             "amuse.io",
				Name:              "Amuse",
				Followers:         "9K",
				Employees:         "51-200",
				AssociatedMembers: "279",
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
				Alias: "Amuse",
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Full Stack Developer",
							ShortDescription:     "You have experience using Go or a strongly typed OO language",
							SwitchingOpportunity: "",
							URL:                  "https://www.indeed.com/viewjob?jk=1aeebd1be05a4f6e",
							Date:                 mustDate("2025-03-19"),
							WithSalary:           true, // $90,000 - $115,000 per year
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
			ShortDescription: "Global distributed  music company",
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Lucid Motors",
			Website: "https://lucidmotors.com/",
			Careers: "https://lucidmotors.com/careers",
			About:   "https://lucidmotors.com/company",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                1684122,
				IDs:               nil,
				Alias:             "lucidmotors",
				Name:              "Lucid Motors",
				Followers:         "512K",
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
				Alias:     "lucid-motors",
				Employees: "2,250",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Lucid-Motors-EI_IE698305.11,23.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Lucid-Motors-Reviews-E698305.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Lucid-Motors-Jobs-E698305.htm",
				Jobs:        "434",
				Reviews:     "1.1K",
				Salaries:    "3.3K",
				ReviewsRate: "3.0",
				Verified:    true,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "Lucid-Motors",
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Senior Full Stack Developer",
							ShortDescription:     "Design and develop full-stack web applications, microservices, and REST APIs using languages such as Golang, Python, Java, or similar",
							SwitchingOpportunity: "",
							URL:                  "https://www.indeed.com/viewjob?jk=e35725dabddd890c",
							Date:                 mustDate("2025-03-19"),
							WithSalary:           true, // $154,000 - $211,750 per year
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
			ShortDescription: "American automotive and technology company that produces electric vehicles",
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "The ODP Corporation",
			Website: "https://www.theodpcorp.com/",
			Careers: "https://careers.theodpcorp.com/",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                75753067,
				IDs:               nil,
				Alias:             "the-odp-corporation",
				Name:              "The ODP Corporation",
				Followers:         "11K",
				Employees:         "10K+",
				AssociatedMembers: "27,156",
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
				Alias:     "the-odp-corporation",
				Employees: "120",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-The-ODP-Corporation-EI_IE9487978.11,30.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/The-ODP-Corporation-Reviews-E9487978.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/The-ODP-Corporation-Jobs-E9487978.htm",
				Jobs:        "7",
				Reviews:     "6",
				Salaries:    "17",
				ReviewsRate: "2.8",
				Verified:    true,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "The-Odp-Corporation",
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Senior Backend Developer",
							ShortDescription:     "Architect and implement solutions primarily using Golang",
							SwitchingOpportunity: "",
							URL:                  "https://www.indeed.com/viewjob?jk=164ee8448e688865",
							Date:                 mustDate("2025-03-19"),
							WithSalary:           true, // $75,500 - $117,950 per year
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
			ShortDescription: "American company who provides business services and supplies, products, and digital workplace technology solutions for small, medium, and enterprise businesses",
		},

		// BigTech
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Netflix",
			Website: "https://netflix.com/",
			Careers: "https://jobs.netflix.com/",
			About:   "https://about.netflix.com/",
			Blog:    "https://netflixtechblog.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                165158,
				IDs:               nil,
				Alias:             "netflix",
				Name:              "Netflix",
				Followers:         "11M",
				Employees:         "10K+",
				AssociatedMembers: "16,766",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:     "netflix",
				Followers: "7.9k",
				Verified:  false,
			},
			BlindProfile: domain.BlindProfile{
				Alias:       "Netflix",
				Employees:   "1,001 to 5,000",
				Salary:      "$87K ~ $827K a year",
				Reviews:     "576",
				ReviewsRate: "4.2",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "netflix",
				Employees: "14,600",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Netflix-EI_IE11891.11,18.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Netflix-Reviews-E11891.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Netflix-Jobs-E11891.htm",
				Jobs:        "466",
				Reviews:     "3.3K",
				Salaries:    "6.8K",
				ReviewsRate: "4.2",
				Verified:    true,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "Netflix",
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Software Engineer L5 — Linux Kernel Developer",
							ShortDescription:     "Proficiency in Golang, Python",
							SwitchingOpportunity: "",
							URL:                  "https://www.indeed.com/viewjob?jk=5a23d65371dc5435",
							Date:                 mustDate("2025-03-19"),
							WithSalary:           true, // $100,000 - $620,000 per year
							Remote:               true,
						},
						{
							Title:                "Software Engineer 6 — Games Experience Engineering",
							ShortDescription:     "Languages: JavaScript, Java, Python, Lua, Golang, Clojure",
							SwitchingOpportunity: "",
							URL:                  "https://www.indeed.com/viewjob?jk=ba7e24381e243687",
							Date:                 mustDate("2025-03-23"),
							WithSalary:           true, // $230,000 - $960,000 per year
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
			ShortDescription: "American online streaming service",
		},

		// BigTech
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Panasonic",
			Website: "https://holdings.panasonic/",
			Careers: "https://holdings.panasonic/global/corporate/careers.html",
			About:   "https://holdings.panasonic/global/corporate/about.html",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                3549587,
				IDs:               nil,
				Alias:             "panasonic",
				Name:              "Panasonic",
				Followers:         "302K",
				Employees:         "10K+",
				AssociatedMembers: "28,942",
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
				Alias:     "panasonic",
				Employees: "259,000",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Panasonic-EI_IE4279.11,20.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Panasonic-Reviews-E4279.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Panasonic-Jobs-E4279.htm",
				Jobs:        "1.2K",
				Reviews:     "5.1K",
				Salaries:    "8.8K",
				ReviewsRate: "3.7",
				Verified:    false,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "Panasonic-a3e6d439",
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Golang Network Developer",
							ShortDescription:     "Strong proven experience in Golang",
							SwitchingOpportunity: "",
							URL:                  "https://www.indeed.com/viewjob?jk=858278a26965db1b",
							Date:                 mustDate("2025-03-19"),
							WithSalary:           true, // $93,000 - $157,000 per year
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
			ShortDescription: "Japanese multinational company who manufactures batteries, automotive and aviation systems, industrial systems, and home repair and construction",
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "TRON DAO",
			Website: "https://tron.network/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                18289948,
				IDs:               []int{29722, 18289948},
				Alias:             "trondao",
				Name:              "TRON DAO",
				Followers:         "54K",
				Employees:         "201-500",
				AssociatedMembers: "375",
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
							Title:                "Blockchain Systems Engineer Golang",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4112783642/",
							Date:                 mustDate("2025-03-19"),
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
			Ignore:           true, // Blockchain | Crypto | Web3
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Gnosis",
			Website: "http://gnosis.io/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                18098361,
				IDs:               nil,
				Alias:             "gnosis-limited",
				Name:              "Gnosis",
				Followers:         "8K",
				Employees:         "51-200",
				AssociatedMembers: "162",
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
							URL:                  "https://www.linkedin.com/jobs/view/4186509026/",
							Date:                 mustDate("2025-03-19"),
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
			Ignore:           true, // Blockchain
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Netscribes",
			Website: "http://www.netscribes.com/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                25974,
				IDs:               nil,
				Alias:             "netscribes",
				Name:              "Netscribes",
				Followers:         "103K",
				Employees:         "1K-5K",
				AssociatedMembers: "1,401",
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
				Alias:     "netscribes",
				Employees: "2,000",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Netscribes-EI_IE218712.11,21.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Netscribes-Reviews-E218712.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Netscribes-Jobs-E218712.htm",
				Jobs:        "14",
				Reviews:     "861",
				Salaries:    "787",
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
							Title:                "Senior Software Engineer – Golang",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4184795681/",
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
			ShortDescription: "",
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Warner Bros. Discovery",
			Website: "https://wbd.com/",
			Careers: "https://careers.wbd.com/",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                73234923,
				IDs:               []int{2046, 2470, 4749, 9620, 167119, 200300, 1393227, 10538105, 26998697, 73234923},
				Alias:             "warner-bros-discovery",
				Name:              "Warner Bros. Discovery",
				Followers:         "811K",
				Employees:         "10K+",
				AssociatedMembers: "39,992",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "",
				Verified: false,
			},
			BlindProfile: domain.BlindProfile{
				Alias:       "Warner-Bros-Discovery",
				Employees:   "10,000+",
				Salary:      "",
				Reviews:     "61",
				ReviewsRate: "2.9",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "warner-bros-discovery",
				Employees: "24,250",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Warner-Bros-Discovery-EI_IE12777.11,32.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Warner-Bros-Discovery-Reviews-E12777.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Warner-Bros-Discovery-Jobs-E12777.htm",
				Jobs:        "3.4K",
				Reviews:     "4.8K",
				Salaries:    "9.3K",
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
							Title:                "Staff Software Engineer – Golang",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4187284222/",
							Date:                 mustDate("2025-03-17"),
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
			Name:    "Gamingtec",
			Website: "https://gamingtec.com/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                9276193,
				IDs:               nil,
				Alias:             "gamingtec",
				Name:              "Gamingtec",
				Followers:         "6K",
				Employees:         "51-200",
				AssociatedMembers: "145",
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
							URL:                  "https://www.linkedin.com/jobs/view/4187105723/",
							Date:                 mustDate("2025-03-19"),
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
			Ignore:           true, // Gambling
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "VOIS",
			Website: "http://www.vodafone.com/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                167822,
				IDs:               nil,
				Alias:             "vois",
				Name:              "VOIS",
				Followers:         "345K",
				Employees:         "10K+",
				AssociatedMembers: "18,326",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:     "vodafone",
				Followers: "75",
				Verified:  false,
			},
			BlindProfile: domain.BlindProfile{
				Alias:       "Vodafone",
				Employees:   "10,000+",
				Salary:      "$20K ~ $450K a year",
				Reviews:     "56",
				ReviewsRate: "3.8",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "vodafone",
				Employees: "165,610",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Vodafone-EI_IE5775.11,19.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Vodafone-Reviews-E5775.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Vodafone-Jobs-E5775.htm",
				Jobs:        "2.4K",
				Reviews:     "19K",
				Salaries:    "27K",
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
							Title:                "Manager Golang Developer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4180356164/",
							Date:                 mustDate("2025-03-19"),
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
			Name:    "Spydra",
			Website: "https://www.spydra.app/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                84968130,
				IDs:               nil,
				Alias:             "spydra",
				Name:              "Spydra",
				Followers:         "50K",
				Employees:         "11-50",
				AssociatedMembers: "36",
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
							URL:                  "https://www.linkedin.com/jobs/view/4184559478/",
							Date:                 mustDate("2025-03-19"),
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
			Ignore:           true, // Blockchain
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Veact GmbH",
			Website: "https://veact.net/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                1412127,
				IDs:               nil,
				Alias:             "veact",
				Name:              "Veact GmbH",
				Followers:         "2K",
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
							URL:                  "https://www.linkedin.com/jobs/view/4182628362/",
							Date:                 mustDate("2025-03-15"),
							WithSalary:           false,
							Remote:               false,
						},
					},
				},
				domain.Elixir:  {},
				domain.Clojure: {},
				domain.Haskell: {},
			},
			ShortDescription: "Veact is developing applications, processes and products to improve sales success in the car trade market using data-based marketing",
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Walmart Data Ventures",
			Website: "https://www.walmartdataventures.com/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                80427781,
				IDs:               nil,
				Alias:             "walmartdataventures",
				Name:              "Walmart Data Ventures",
				Followers:         "15K",
				Employees:         "201-500",
				AssociatedMembers: "244",
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
							Title:                "Senior, Data Engineer (Scala/Spark)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4189070192/",
							Date:                 mustDate("2025-03-19"),
							WithSalary:           true,
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
			Name:    "Funnel",
			Website: "https://funnel.io/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                9397630,
				IDs:               nil,
				Alias:             "funnel-io",
				Name:              "Funnel",
				Followers:         "54K",
				Employees:         "201-500",
				AssociatedMembers: "437",
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
							Title:                "Senior Data Platform Engineer – Rust",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4142024603/",
							Date:                 mustDate("2025-04-11"), // mustDate("2025-03-20"),
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
			ShortDescription: "Marketing intelligence platform",
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "OKX",
			Website: "http://www.okx.com/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                13323120,
				IDs:               nil,
				Alias:             "okxofficial",
				Name:              "OKX",
				Followers:         "462K",
				Employees:         "1K-5K",
				AssociatedMembers: "5,933",
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
							Title:                "Blockchain Engineer (Rust/Go)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4187130229/",
							Date:                 mustDate("2025-03-20"),
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
			Ignore:           true, // Crypto
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Lingopal",
			Website: "https://lingopal.com/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                98079583,
				IDs:               nil,
				Alias:             "lingopal-ai",
				Name:              "Lingopal",
				Followers:         "50K",
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
							URL:                  "https://www.linkedin.com/jobs/view/4188980146/",
							Date:                 mustDate("2025-03-20"),
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
			ShortDescription: "Translate LIVE speech into any language in real time",
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Scanbo",
			Website: "https://www.scanbo.com/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                13271827,
				IDs:               nil,
				Alias:             "scanbo",
				Name:              "Scanbo",
				Followers:         "2K",
				Employees:         "2-10",
				AssociatedMembers: "25",
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
							Title:                "Rust Blockchain Developer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4185486374/",
							Date:                 mustDate("2025-03-22"),
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
			ShortDescription: "SCANBO is a HealthCare platform dedicated to making medical diagnosis available at the point of care and at home",
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "SurrealDB",
			Website: "https://surrealdb.com/",
			Careers: "https://surrealdb.com/careers",
			About:   "https://medium.com/surrealdb/about",
			Blog:    "https://surrealdb.com/blog",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                71534355,
				IDs:               nil,
				Alias:             "surrealdb",
				Name:              "SurrealDB",
				Followers:         "7K",
				Employees:         "11-50",
				AssociatedMembers: "43",
				Verified:          false,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "surrealdb",
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
			OttaProfileSlug: "SurrealDB",
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 11,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Rust: {
					GitHubRepositoriesCount: 25,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Senior Software Engineer",
							ShortDescription:     "Proficient in programming languages including Rust, C, or Go",
							SwitchingOpportunity: "",
							URL:                  "https://rustjobs.dev/featured-jobs/SurrealDB-Senior-Software-Engineer-Query-Language-QL-TcCVMEInkwcjBpupa3Ys",
							Date:                 mustDate("2024-02-03"),
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
			ShortDescription: "Serverless cloud database",
			Industries: []domain.Industry{
				domain.IndustryDevOps,
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Clockwork Labs",
			Website: "https://clockworklabs.io/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                19088131,
				IDs:               nil,
				Alias:             "clockwork-labs",
				Name:              "Clockwork Labs",
				Followers:         "3K",
				Employees:         "2-10",
				AssociatedMembers: "31",
				Verified:          false,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "clockworklabs",
				Verified: false,
			},
			BlindProfile: domain.BlindProfile{
				Alias: "",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
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
					GitHubRepositoriesCount: 9,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Zig:     {},
				domain.Scala:   {},
				domain.Elixir:  {},
				domain.Clojure: {},
				domain.Haskell: {},
			},
			ShortDescription: "Computer Games",
		},
		{
			ID:      0,        // system
			Type:    "",       // system
			Name:    "Polars", // Polars BV
			Website: "https://pola.rs/",
			Careers: "",
			About:   "https://pola.rs/about-us/",
			Blog:    "https://pola.rs/posts/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                98439348,
				IDs:               nil,
				Alias:             "pola-rs",
				Name:              "Polars",
				Followers:         "20K",
				Employees:         "2-10",
				AssociatedMembers: "14",
				Verified:          false,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "pola-rs",
				Verified: false,
			},
			BlindProfile: domain.BlindProfile{
				Alias: "",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
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
					GitHubRepositoriesCount: 8,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Rust Backend Engineer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://hiring.pola.rs/o/rust-backend-engineer",
							Date:                 mustDate("2025-03-21"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Database Engineer",
							ShortDescription:     "Experience in Rust",
							SwitchingOpportunity: "",
							URL:                  "https://hiring.pola.rs/o/database-engineer",
							Date:                 mustDate("2025-03-21"),
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
			ShortDescription: "Lightning-fast DataFrame library for Rust and Python",
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Tsunami Tsolutions",
			Website: "https://tsunamitsolutions.com/",
			Careers: "https://tsunamitsolutions.com/index.php/careers/",
			About:   "https://tsunamitsolutions.com/index.php/who-we-are/",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                983003,
				IDs:               nil,
				Alias:             "tsunami-tsolutions-llc.",
				Name:              "Tsunami Tsolutions",
				Followers:         "3K",
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
				OverviewURL: "",
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "",
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{
						// Unclear if this is a Go position https://www.indeed.com/viewjob?jk=e5b01bd47b175b64
					},
				},
				domain.Rust:    {},
				domain.Zig:     {},
				domain.Scala:   {},
				domain.Elixir:  {},
				domain.Clojure: {},
				domain.Haskell: {},
			},
			ShortDescription: "American aviation & aerospace company that offers Engineering Support and Information Solutions",
			Ignore:           true, // Unclear about the tech stack
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Draper",
			Website: "https://www.draper.com/",
			Careers: "https://www.draper.com/careers",
			About:   "https://www.draper.com/about",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                163500,
				IDs:               nil,
				Alias:             "draper",
				Name:              "Draper",
				Followers:         "46K",
				Employees:         "1K-5K",
				AssociatedMembers: "2,475",
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
				Alias:     "draper",
				Employees: "1,750",
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
						// Unclear if this is a Go position https://www.indeed.com/viewjob?jk=f77b58384382b852
					},
				},
				domain.Rust:    {},
				domain.Zig:     {},
				domain.Scala:   {},
				domain.Elixir:  {},
				domain.Clojure: {},
				domain.Haskell: {},
			},
			ShortDescription: "American nonprofit engineering innovation company",
			Ignore:           true, // Unclear about the tech stack
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Mayo Clinic",
			Website: "https://www.mayoclinic.org/",
			Careers: "https://jobs.mayoclinic.org/",
			About:   "https://www.mayoclinic.org/about-mayo-clinic",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                4725,
				IDs:               nil,
				Alias:             "mayo-clinic",
				Name:              "Mayo Clinic",
				Followers:         "1M",
				Employees:         "10K+",
				AssociatedMembers: "48,372",
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
				Alias:     "mayo-clinic",
				Employees: "63,000",
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
						// Unclear if this is a Go position https://www.indeed.com/viewjob?jk=4c3c7f291e1c965f
					},
				},
				domain.Rust:    {},
				domain.Zig:     {},
				domain.Scala:   {},
				domain.Elixir:  {},
				domain.Clojure: {},
				domain.Haskell: {},
			},
			ShortDescription: "American nonprofit academic medical center focused on integrated health care, education, and research",
			Ignore:           true, // Unclear about the tech stack
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "University of Florida",
			Website: "http://ufl.edu/",
			Careers: "https://jobs.ufl.edu/",
			About:   "https://www.ufl.edu/about/",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                4282,
				IDs:               []int{4282, 7974417, 15101148, 15102044, 18869387},
				Alias:             "uflorida",
				Name:              "University of Florida",
				Followers:         "606K",
				Employees:         "",
				AssociatedMembers: "",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "smartsystemslab-uf",
				Verified: false,
			},
			BlindProfile: domain.BlindProfile{
				Alias: "",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "university-of-florida",
				Employees: "19,453",
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
						// Unclear if this is a Go position https://www.indeed.com/viewjob?jk=630a513e4b7e6d8e
					},
				},
				domain.Rust:    {},
				domain.Zig:     {},
				domain.Scala:   {},
				domain.Elixir:  {},
				domain.Clojure: {},
				domain.Haskell: {},
			},
			ShortDescription: "American public land-grant research university",
			Ignore:           true, // Unclear about the tech stack
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Storage Rentals of America",
			Website: "https://www.sroa.com/",
			Careers: "https://www.sroa.com/careers",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                33228180,
				IDs:               nil,
				Alias:             "storage-rentals-america",
				Name:              "Storage Rentals of America",
				Followers:         "5K",
				Employees:         "501-1K",
				AssociatedMembers: "200",
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
							Title:                "DevOps Engineer",
							ShortDescription:     "Experience in programming with TypeScript and Golang, with the ability to write SQL scripts for database management and operations",
							SwitchingOpportunity: "",
							URL:                  "https://www.indeed.com/viewjob?jk=7195850460dbc38c",
							Date:                 mustDate("2025-03-23"),
							WithSalary:           true, // $75,000 - $95,000 per year
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
			ShortDescription: "American self-storage company who provides wide range of rentable storage properties",
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Red Hat",
			Website: "https://www.redhat.com/",
			Careers: "https://www.redhat.com/en/jobs",
			About:   "https://www.redhat.com/en/about",
			Blog:    "https://www.redhat.com/en/blog",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                3545,
				IDs:               nil,
				Alias:             "red-hat",
				Name:              "Red Hat",
				Followers:         "1M",
				Employees:         "10K+",
				AssociatedMembers: "19,585",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:     "redhatofficial",
				Followers: "2.2k",
				Verified:  false,
			},
			BlindProfile: domain.BlindProfile{
				Alias:       "Red-Hat",
				Employees:   "10,000+",
				Salary:      "$37K ~ $250K a year",
				Reviews:     "443",
				ReviewsRate: "4.0",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "red-hat",
				Employees: "18,600",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Red-Hat-EI_IE8868.11,18.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Red-Hat-Reviews-E8868.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Red-Hat-Jobs-E8868.htm",
				Jobs:        "294",
				Reviews:     "5.3K",
				Salaries:    "10K",
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
							Title:                "Software Engineer",
							ShortDescription:     "2+ years of experience working in a Linux environment with Golang or Python",
							SwitchingOpportunity: "",
							URL:                  "https://www.indeed.com/viewjob?jk=974e99549d320db1",
							Date:                 mustDate("2025-03-23"),
							WithSalary:           true, // $94,550 - $151,170 per year
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
			ShortDescription: "American software company that provides open source software products to enterprises and is a subsidiary of IBM",
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "OCC",
			Website: "https://www.theocc.com/",
			Careers: "https://www.theocc.com/careers",
			About:   "https://www.theocc.com/company-information/what-is-occ",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                47420,
				IDs:               nil,
				Alias:             "occ",
				Name:              "OCC",
				Followers:         "31K",
				Employees:         "1K-5K",
				AssociatedMembers: "2,568",
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
				Alias:     "occ",
				Employees: "2,250",
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
							Title:                "Software Engineering: Java — Associate Principal",
							ShortDescription:     "Experience with app development in Golang",
							SwitchingOpportunity: "",
							URL:                  "https://www.indeed.com/viewjob?jk=0ce9ca11c038853d",
							Date:                 mustDate("2025-03-23"),
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
			ShortDescription: "American clearing house who specializes in equity derivatives clearing",
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Motive",
			Website: "https://gomotive.com/",
			Careers: "https://gomotive.com/company/careers/",
			About:   "",
			Blog:    "https://medium.com/motive-eng",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                3271606,
				IDs:               nil,
				Alias:             "motive-inc",
				Name:              "Motive",
				Followers:         "540K",
				Employees:         "1K-5K",
				AssociatedMembers: "4,492",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "GoMotive",
				Verified: false,
			},
			BlindProfile: domain.BlindProfile{
				Alias: "",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "motive",
				Employees: "4,024",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Motive-EI_IE1018424.11,17.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Motive-Reviews-E1018424.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Motive-Jobs-E1018424.htm",
				Jobs:        "102",
				Reviews:     "1.5K",
				Salaries:    "2K",
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
							Title:                "Staff System Software Engineer",
							ShortDescription:     "Proficiency in Golang, Kotlin, and C++",
							SwitchingOpportunity: "",
							URL:                  "https://www.indeed.com/viewjob?jk=8206a2447ac1e213",
							Date:                 mustDate("2025-03-23"),
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
			ShortDescription: "American technology company that creates software used by truck",
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Freddie Mac",
			Website: "https://www.freddiemac.com/",
			Careers: "https://careers.freddiemac.com/",
			About:   "https://www.freddiemac.com/about",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                3140,
				IDs:               nil,
				Alias:             "freddie-mac",
				Name:              "Freddie Mac",
				Followers:         "184K",
				Employees:         "5K-10K",
				AssociatedMembers: "11,147",
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
				Alias:     "freddie-mac",
				Employees: "9,250",
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
					Vacancies:               []domain.Vacancy{
						// Unclear if this is a Rust position https://www.indeed.com/viewjob?jk=1ddd9df71b91cb71
					},
				},
				domain.Zig:     {},
				domain.Scala:   {},
				domain.Elixir:  {},
				domain.Clojure: {},
				domain.Haskell: {},
			},
			ShortDescription: "American government-sponsored agency charged with keeping mortgage markets liquid",
			Ignore:           true, // Unclear about the tech stack
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Stoke Space",
			Website: "http://www.stokespace.com/",
			Careers: "https://www.stokespace.com/careers/",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                53254054,
				IDs:               nil,
				Alias:             "stoke-space",
				Name:              "Stoke Space",
				Followers:         "45K",
				Employees:         "201-500",
				AssociatedMembers: "215",
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
				Alias:     "stoke-space",
				Employees: "85",
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
							Title:                "Vehicle Software Rust Developer",
							ShortDescription:     "Experience in Rust",
							SwitchingOpportunity: "",
							URL:                  "https://www.indeed.com/viewjob?jk=e6442d48c3e89a90",
							Date:                 mustDate("2025-03-23"),
							WithSalary:           true, // $94,100 - $275,600 per year
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
			ShortDescription: "American space launch company who developing a fully and rapidly reusable space launch vehicle",
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "GreyNoise Intelligence",
			Website: "https://greynoise.io/",
			Careers: "https://www.greynoise.io/careers",
			About:   "https://www.greynoise.io/about",
			Blog:    "https://www.greynoise.io/blog",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                11204989,
				IDs:               nil,
				Alias:             "greynoise",
				Name:              "GreyNoise Intelligence",
				Followers:         "11K",
				Employees:         "11-50",
				AssociatedMembers: "62",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "greynoise-intelligence",
				Verified: false,
			},
			BlindProfile: domain.BlindProfile{
				Alias: "",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "greynoise-intelligence",
				Employees: "59",
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
							Title:                "Senior Software Engineer (Platform)",
							ShortDescription:     "5+ years of production software engineering experience",
							SwitchingOpportunity: "Working with Go (Golang) for backend development",
							URL:                  "https://www.indeed.com/viewjob?jk=40120128f063af25",
							Date:                 mustDate("2025-03-24"),
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
			ShortDescription: "American cybersecurity company",
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Universal Music Group",
			Website: "https://www.universalmusic.com/",
			Careers: "https://www.umusiccareers.com/",
			About:   "https://www.universalmusic.com/company/",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                3007,
				IDs:               nil,
				Alias:             "universalmusicgroup",
				Name:              "Universal Music Group",
				Followers:         "951K",
				Employees:         "5K-10K",
				AssociatedMembers: "12,018",
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
				Alias:     "universal-music-group",
				Employees: "8,250",
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
						// Unclear if this is a Go position https://www.indeed.com/viewjob?jk=b9f6faf52fb64156
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
			Ignore:           true, // Unclear about the tech stack
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Circles",
			Website: "https://circles.co/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                6449768,
				IDs:               nil,
				Alias:             "circles1",
				Name:              "Circles",
				Followers:         "69K",
				Employees:         "501-1K",
				AssociatedMembers: "1,324",
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
							Title:                "Staff Software Engineer (Golang / Node.js)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4187988746/",
							Date:                 mustDate("2025-03-21"),
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
			ShortDescription: "Circles is a technology company reimagining the telco industry with its SaaS platform, helping telco operators launch and operate successful digital brands",
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "PulsePoint",
			Website: "https://www.pulsepoint.com/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                2349382,
				IDs:               nil,
				Alias:             "pulsepoint",
				Name:              "PulsePoint",
				Followers:         "52K",
				Employees:         "51-200",
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
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Senior K8/Golang Engineer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4188677602/",
							Date:                 mustDate("2025-03-22"),
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
			ShortDescription: "PulsePoint is a technology company using data to accelerate healthcare marketing",
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Armada",
			Website: "https://www.armada.ai/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                97187862,
				IDs:               nil,
				Alias:             "armadaai",
				Name:              "Armada",
				Followers:         "13K",
				Employees:         "201-500",
				AssociatedMembers: "245",
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
							Title:                "Software Engineer (Golang)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4190620206/",
							Date:                 mustDate("2025-03-21"),
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
			ShortDescription: "Computing platform",
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Hubble.Build",
			Website: "https://hubble.build/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                13184745,
				IDs:               nil,
				Alias:             "hubble-pte.-ltd.",
				Name:              "Hubble.Build",
				Followers:         "34K",
				Employees:         "51-200",
				AssociatedMembers: "76",
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
							Title:                "Software Engineer (Golang / Ruby on Rails)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4190270866/",
							Date:                 mustDate("2025-03-21"),
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
			ShortDescription: "Hubble.Build is a construction management platform that seamlessly connects stakeholders across the entire value chain to build better, faster, safer, and more cost-effective projects",
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Skillz",
			Website: "https://www.skillz.com/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                3104354,
				IDs:               nil,
				Alias:             "skillz",
				Name:              "Skillz",
				Followers:         "20K",
				Employees:         "51-200",
				AssociatedMembers: "215",
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
							Title:                "Backend Engineer – Java / Golang",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4192054490/",
							Date:                 mustDate("2025-03-24"),
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
			Name:    "Unity",
			Website: "https://unity.com/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                212669,
				IDs:               []int{5676, 212669, 1388753, 15091249, 18301696},
				Alias:             "unity",
				Name:              "Unity",
				Followers:         "745K",
				Employees:         "5K-10K",
				AssociatedMembers: "6,927",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:     "unity-technologies",
				Followers: "10.8k",
				Verified:  false,
			},
			BlindProfile: domain.BlindProfile{
				Alias:       "unity%20technologies",
				Employees:   "1,001 to 5,000",
				Salary:      "$48K ~ $290K a year",
				Reviews:     "659",
				ReviewsRate: "3.2",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "unity",
				Employees: "5,440",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Unity-EI_IE455854.11,16.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Unity-Reviews-E455854.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Unity-Jobs-E455854.htm",
				Jobs:        "79",
				Reviews:     "1.8K",
				Salaries:    "4.3K",
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
							Title:                "Golang Backend Engineer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4191748355/",
							Date:                 mustDate("2025-04-14"), // mustDate("2025-03-23"),
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
			ShortDescription: "Platform for creating and operating real-time 3D content",
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "adjoe",
			Website: "https://adjoe.io/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                12601466,
				IDs:               nil,
				Alias:             "adjoe",
				Name:              "adjoe",
				Followers:         "11K",
				Employees:         "51-200",
				AssociatedMembers: "149",
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
							ShortDescription:     "Advertising Dashboard Team",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4190144087/",
							Date:                 mustDate("2025-03-24"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Senior Go Developer",
							ShortDescription:     "Advertising Dashboard Team",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4198154575/",
							Date:                 mustDate("2025-04-03"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Senior Go Backend Developer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4207721435/",
							Date:                 mustDate("2025-04-12"),
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
			Name:    "Docebo",
			Website: "https://www.docebo.com/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                210447,
				IDs:               nil,
				Alias:             "docebo",
				Name:              "Docebo",
				Followers:         "89K",
				Employees:         "501-1K",
				AssociatedMembers: "1,023",
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
							Title:                "Backend Developer – Go",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4192297683/",
							Date:                 mustDate("2025-03-24"),
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
			ShortDescription: "Learning platform",
			Industries: []domain.Industry{
				domain.IndustryEdTech,
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Sona",
			Website: "https://www.getsona.com/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                79810814,
				IDs:               nil,
				Alias:             "getsona",
				Name:              "Sona (getsona.com)",
				Followers:         "10K",
				Employees:         "51-200",
				AssociatedMembers: "137",
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
							Title:                "AI Software Engineer (Elixir, LLMs)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4175895182/",
							Date:                 mustDate("2025-03-22"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Full Stack Engineer (Elixir + LiveView)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4175892408/",
							Date:                 mustDate("2025-04-12"), // mustDate("2025-03-30"),
							WithSalary:           false,
							Remote:               true,
						},
					},
				},
				domain.Clojure: {},
				domain.Haskell: {},
			},
			ShortDescription: "Workforce management solution",
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "ActiveProspect",
			Website: "https://activeprospect.com/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                523272,
				IDs:               []int{523272, 10653877},
				Alias:             "activeprospect",
				Name:              "ActiveProspect",
				Followers:         "21K",
				Employees:         "51-200",
				AssociatedMembers: "150",
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
							Title:                "Senior Software Engineer (Elixir)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4190011323/",
							Date:                 mustDate("2025-03-20"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Senior Software Engineer (Elixir)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4201324957/",
							Date:                 mustDate("2025-04-04"),
							WithSalary:           false,
							Remote:               false,
						},
					},
				},
				domain.Clojure: {},
				domain.Haskell: {},
			},
			ShortDescription: "Consent-based marketing platform",
			Industries: []domain.Industry{
				domain.IndustryMarTech,
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Ledger",
			Website: "https://www.ledger.com/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                5389345,
				IDs:               nil,
				Alias:             "ledgerhq",
				Name:              "Ledger",
				Followers:         "79K",
				Employees:         "501-1K",
				AssociatedMembers: "710",
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
							Title:                "Senior Scala Backend Engineer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4192446500/",
							Date:                 mustDate("2025-03-24"),
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
			Ignore:           true, // Crypto
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "TileDB",
			Website: "https://tiledb.com/",
			Careers: "https://ats.rippling.com/tiledb-careers/jobs",
			About:   "https://tiledb.com/about/",
			Blog:    "https://tiledb.com/blog",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                11301003,
				IDs:               nil,
				Alias:             "tiledb-inc",
				Name:              "TileDB",
				Followers:         "13K",
				Employees:         "51-200",
				AssociatedMembers: "75",
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
							Title:                "Staff Software Engineer – Golang (required)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4201033819/",
							Date:                 mustDate("2025-04-04"),
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
			ShortDescription: "Embedded storage engine designed to support the storage and access of both dense and sparse multi-dimensional arrays",
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Redpanda Data",
			Website: "https://redpanda.com/",
			Careers: "https://www.redpanda.com/careers",
			About:   "https://www.redpanda.com/about-us",
			Blog:    "https://www.redpanda.com/blog",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                35528665,
				IDs:               nil,
				Alias:             "redpanda-data",
				Name:              "Redpanda Data",
				Followers:         "19K",
				Employees:         "51-200",
				AssociatedMembers: "170",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "redpanda-data",
				Verified: true,
			},
			BlindProfile: domain.BlindProfile{
				Alias: "",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "redpanda-data",
				Employees: "183",
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
							ShortDescription:     "Strong understanding and experience of Go",
							SwitchingOpportunity: "",
							URL:                  "https://www.indeed.com/viewjob?jk=5c8079196dde8330",
							Date:                 mustDate("2025-03-25"),
							WithSalary:           true, // $185,000 - $225,000 per year
							Remote:               true,
						},
					},
				},
				domain.Rust: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{
						// Waiting for Rust position
					},
				},
				domain.Zig:     {},
				domain.Scala:   {},
				domain.Elixir:  {},
				domain.Clojure: {},
				domain.Haskell: {},
			},
			ShortDescription: "Modern streaming data platform for developers",
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Quizlet",
			Website: "https://quizlet.com/",
			Careers: "https://quizlet.com/careers",
			About:   "https://quizlet.com/mission",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                807920,
				IDs:               nil,
				Alias:             "quizlet",
				Name:              "Quizlet",
				Followers:         "47K",
				Employees:         "201-500",
				AssociatedMembers: "572",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "quizlet",
				Verified: true,
			},
			BlindProfile: domain.BlindProfile{
				Alias: "",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "quizlet",
				Employees: "100",
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
						// Waiting for Go position
					},
				},
				domain.Rust:    {},
				domain.Zig:     {},
				domain.Scala:   {},
				domain.Elixir:  {},
				domain.Clojure: {},
				domain.Haskell: {},
			},
			ShortDescription: "American company that provides tools for studying and learning",
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Svix",
			Website: "https://www.svix.com/",
			Careers: "https://www.svix.com/careers/",
			About:   "https://www.svix.com/about/",
			Blog:    "https://www.svix.com/blog/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                71559487,
				IDs:               nil,
				Alias:             "svix",
				Name:              "Svix",
				Followers:         "2K",
				Employees:         "11-50",
				AssociatedMembers: "10",
				Verified:          false,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "svix",
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
				domain.Go: {},
				domain.Rust: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Rust Backend Engineer",
							ShortDescription:     "Experience building production services with Rust",
							SwitchingOpportunity: "",
							URL:                  "https://www.indeed.com/viewjob?jk=c5a456ba1962bc7e",
							Date:                 mustDate("2025-03-25"),
							WithSalary:           true, // $130,000 - $160,000 per year
							Remote:               true,
						},
						{
							Title:                "Tech lead",
							ShortDescription:     "Experience with Rust, Postgres, and Redis",
							SwitchingOpportunity: "",
							URL:                  "https://www.indeed.com/viewjob?jk=1f4022f006ac8456",
							Date:                 mustDate("2025-03-25"),
							WithSalary:           true, // $190,000 - $210,000 per year
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
			ShortDescription: "Webhooks-as-a-service platform that empowers companies of all sizes to send webhooks easily and reliably",
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Anduril Industries",
			Website: "https://www.anduril.com/",
			Careers: "https://www.anduril.com/careers/",
			About:   "https://www.anduril.com/mission/",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                18293159,
				IDs:               nil,
				Alias:             "andurilindustries",
				Name:              "Anduril Industries",
				Followers:         "256K",
				Employees:         "1K-5K",
				AssociatedMembers: "4,143",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "anduril",
				Verified: true,
			},
			BlindProfile: domain.BlindProfile{
				Alias: "",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "anduril-industries",
				Employees: "530",
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
							Title:                "Game Developer — Sensor Simulation",
							ShortDescription:     "Rust programming experience",
							SwitchingOpportunity: "",
							URL:                  "https://www.indeed.com/viewjob?jk=1645f9d3eb114feb",
							Date:                 mustDate("2025-03-25"),
							WithSalary:           true, // $138,000 - $207,000 per year
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
			ShortDescription: "American defense technology company that specializes in autonomous systems",
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "CORYS",
			Website: "https://www.corys.com/",
			Careers: "https://www.corys.com/en/home/career/",
			About:   "https://www.corys.com/en/home/about-us/",
			Blog:    "https://www.corys.com/en/category/blog-en/engineering-blog/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                48613,
				IDs:               nil,
				Alias:             "corys",
				Name:              "CORYS",
				Followers:         "8K",
				Employees:         "201-500",
				AssociatedMembers: "541",
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
							Title:                "Software Developer — Power Simulation",
							ShortDescription:     "Rust or C++",
							SwitchingOpportunity: "Experience with Rust is strongly preferred",
							URL:                  "https://www.indeed.com/viewjob?jk=078b469307edba42",
							Date:                 mustDate("2025-03-25"),
							WithSalary:           true, // $85,000 - $120,000 per year
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
			ShortDescription: "French engineering company who provides real time high-fidelity dynamic process simulation, engineering, and training for process industries",
		},
	}
}
