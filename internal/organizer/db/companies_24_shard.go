package db

import "github.com/readytotouch/readytotouch/internal/domain"

func companies24Shard() []domain.CompanyProfile {
	return []domain.CompanyProfile{
		{
			ID:         0,  // system
			Type:       "", // system
			Name:       "FTAPI",
			BaseURL:    "https://www.ftapi.com/",
			CareersURL: "",
			AboutURL:   "",
			BlogURL:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                1086088,
				IDs:               nil,
				Alias:             "ftapi-software-gmbh",
				Name:              "FTAPI",
				Followers:         "7K",
				Employees:         "51-200",
				AssociatedMembers: "159",
				Verified:          true,
				Date:              mustDate("2026-02-17"),
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "",
				Verified: false,
			},
			BlindProfile: domain.BlindProfile{
				Alias: "",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
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
					GitHubRepositoryCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "(Senior) Backend Engineer (Go)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4316306633/",
							Location:             "Munich, Bavaria, Germany",
							CloudProviders:       []domain.CloudProvider{},
							Date:                 mustDate("2026-02-13"),
							WithSalary:           false,
							Remote:               false,
						},
					},
				},
				domain.Rust: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Zig: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Scala: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Elixir: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Clojure: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Haskell: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Erlang: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
			},
			ShortDescription: "Platform for secure data exchange",
			Industries: []domain.Industry{
				domain.IndustryCyberSecurity,
			},
		},
		{
			ID:         0,  // system
			Type:       "", // system
			Name:       "Shopfully",
			BaseURL:    "https://shopfully.com/",
			CareersURL: "",
			AboutURL:   "",
			BlogURL:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                2760711,
				IDs:               []int{1546328, 2760711},
				Alias:             "shopfully",
				Name:              "Shopfully",
				Followers:         "60K",
				Employees:         "201-500",
				AssociatedMembers: "545",
				Verified:          true,
				Date:              mustDate("2026-02-17"),
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "",
				Verified: false,
			},
			BlindProfile: domain.BlindProfile{
				Alias: "",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias: "",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Percona-LLC-EI_IE283779.11,22.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/index.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/ShopFully-Jobs-E1472579.htm",
				Jobs:        "", // Becomes empty in the new design version
				Reviews:     "", // Becomes empty in the new design version
				Salaries:    "", // Becomes empty in the new design version
				ReviewsRate: "",
				Verified:    false,
				Date:        mustDate("2026-02-17"),
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "",
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoryCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Senior Backend Software Engineer",
							ShortDescription:     "Proven experience with Go",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4372733783/",
							Location:             "Spain",
							CloudProviders:       []domain.CloudProvider{domain.AWS},
							Date:                 mustDate("2026-02-14"),
							WithSalary:           false,
							Remote:               true,
						},
					},
				},
				domain.Rust: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Zig: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Scala: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Elixir: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Clojure: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Haskell: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Erlang: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
			},
			ShortDescription: "",
			Industries:       []domain.Industry{
				// NOP
			},
		},
		{
			ID:         0,  // system
			Type:       "", // system
			Name:       "FXC Intelligence",
			BaseURL:    "https://www.fxcintel.com/",
			CareersURL: "https://careers.fxcintel.com/",
			AboutURL:   "https://www.fxcintel.com/about",
			BlogURL:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                748331,
				IDs:               nil,
				Alias:             "fxcintelligence",
				Name:              "FXC Intelligence",
				Followers:         "16K",
				Employees:         "51-200",
				AssociatedMembers: "84",
				Verified:          true,
				Date:              mustDate("2026-02-17"),
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "",
				Verified: false,
			},
			BlindProfile: domain.BlindProfile{
				Alias: "",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias: "",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Percona-LLC-EI_IE283779.11,22.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/index.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/FXC-Intelligence-Jobs-E988093.htm",
				Jobs:        "", // Becomes empty in the new design version
				Reviews:     "", // Becomes empty in the new design version
				Salaries:    "", // Becomes empty in the new design version
				ReviewsRate: "",
				Verified:    false,
				Date:        mustDate("2026-02-17"),
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "",
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoryCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Senior Software Engineer (Golang/TypeScript)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4373565827/",
							Location:             "European Economic Area",
							CloudProviders:       []domain.CloudProvider{},
							Date:                 mustDate("2026-02-16"),
							WithSalary:           false,
							Remote:               true,
						},
					},
				},
				domain.Rust: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Zig: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Scala: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Elixir: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Clojure: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Haskell: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Erlang: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
			},
			ShortDescription: "Financial data platform",
			Industries:       []domain.Industry{
				// NOP
			},
		},
		{
			ID:         0,  // system
			Type:       "", // system
			Name:       "Wobot AI",
			BaseURL:    "https://wobot.ai/",
			CareersURL: "https://wobot.ai/careers",
			AboutURL:   "https://wobot.ai/about-us",
			BlogURL:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                13654015,
				IDs:               nil,
				Alias:             "wobot-ai",
				Name:              "Wobot AI",
				Followers:         "12K",
				Employees:         "51-200",
				AssociatedMembers: "110",
				Verified:          true,
				Date:              mustDate("2026-02-20"),
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "",
				Verified: false,
			},
			BlindProfile: domain.BlindProfile{
				Alias: "",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
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
					GitHubRepositoryCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Golang Developer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4372926753/",
							Location:             "Greater Kolkata Area",
							CloudProviders:       []domain.CloudProvider{},
							Date:                 mustDate("2026-02-14"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Golang Developer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4377062327/",
							Location:             "Greater Kolkata Area",
							CloudProviders:       []domain.CloudProvider{},
							Date:                 mustDate("2026-02-25"),
							WithSalary:           false,
							Remote:               false,
						},
					},
				},
				domain.Rust: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Zig: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Scala: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Elixir: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Clojure: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Haskell: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Erlang: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
			},
			ShortDescription: "The product tracks workplace processes to address issues that negatively affect customer experiences",
			Industries:       []domain.Industry{
				// NOP
			},
		},
		{
			ID:         0,  // system
			Type:       "", // system
			Name:       "Gen",
			BaseURL:    "http://gendigital.com/",
			CareersURL: "",
			AboutURL:   "",
			BlogURL:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                86730975,
				IDs:               []int{28199, 34316, 86730975},
				Alias:             "gendigitalinc",
				Name:              "Gen",
				Followers:         "77K",
				Employees:         "5K-10K",
				AssociatedMembers: "3,441",
				Verified:          true,
				Date:              mustDate("2026-02-19"),
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "",
				Verified: false,
			},
			BlindProfile: domain.BlindProfile{
				Alias: "",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias: "",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Percona-LLC-EI_IE283779.11,22.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/index.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Gen-Jobs-E8043746.htm",
				Jobs:        "", // Becomes empty in the new design version
				Reviews:     "", // Becomes empty in the new design version
				Salaries:    "", // Becomes empty in the new design version
				ReviewsRate: "",
				Verified:    false,
				Date:        mustDate("2026-02-20"),
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "",
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Rust: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Zig: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Scala: {
					GitHubRepositoryCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Staff Engineer (Scala)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4374784932/",
							Location:             "New York, NY",
							CloudProviders:       []domain.CloudProvider{},
							Date:                 mustDate("2026-02-18"),
							WithSalary:           true, // $200k/yr - $225k/yr
							Remote:               false,
						},
					},
				},
				domain.Elixir: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Clojure: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Haskell: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Erlang: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
			},
			ShortDescription: "",
			Industries: []domain.Industry{
				domain.IndustryCyberSecurity,
			},
		},
		{
			ID:         0,  // system
			Type:       "", // system
			Name:       "Six Robotics",
			BaseURL:    "https://sixrobotics.com/",
			CareersURL: "https://sixrobotics.com/careers",
			AboutURL:   "https://sixrobotics.com/company",
			BlogURL:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                102401514,
				IDs:               nil,
				Alias:             "six-robotics",
				Name:              "Six Robotics",
				Followers:         "3K",
				Employees:         "51-200",
				AssociatedMembers: "42",
				Verified:          false,
				Date:              mustDate("2026-02-19"),
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "",
				Verified: false,
			},
			BlindProfile: domain.BlindProfile{
				Alias: "",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
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
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Rust: {
					GitHubRepositoryCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Senior Software Engineer (Rust)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4372675375/",
							Location:             "Oslo, Oslo, Norway",
							CloudProviders:       []domain.CloudProvider{},
							Date:                 mustDate("2026-02-19"),
							WithSalary:           false,
							Remote:               false,
						},
					},
				},
				domain.Zig: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Scala: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Elixir: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Clojure: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Haskell: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Erlang: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
			},
			ShortDescription: "Swarm autonomy as a force multiplier",
			Industries: []domain.Industry{
				domain.IndustryDefenseTech,
			},
		},
		{
			ID:         0,  // system
			Type:       "", // system
			Name:       "Input Output Group",
			BaseURL:    "https://www.iog.io/",
			CareersURL: "https://www.iog.io/careers",
			AboutURL:   "https://www.iog.io/about",
			BlogURL:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                6385405,
				IDs:               nil,
				Alias:             "input-output-group",
				Name:              "Input Output Group",
				Followers:         "87K",
				Employees:         "501-1K",
				AssociatedMembers: "425",
				Verified:          true,
				Date:              mustDate("2026-02-20"),
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "",
				Verified: false,
			},
			BlindProfile: domain.BlindProfile{
				Alias: "",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
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
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Rust: {
					GitHubRepositoryCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Software Engineer (Rust)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4361625819/",
							Location:             "United Kingdom",
							CloudProviders:       []domain.CloudProvider{},
							Date:                 mustDate("2026-02-18"),
							WithSalary:           false,
							Remote:               true,
						},
					},
				},
				domain.Zig: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Scala: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Elixir: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Clojure: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Haskell: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Erlang: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
			},
			ShortDescription: "Cardano blockchain platform",
			Industries:       []domain.Industry{
				// Blockchain
			},
			Ignore: true, // Blockchain
		},
		{
			ID:         0,  // system
			Type:       "", // system
			Name:       "Glovo",
			BaseURL:    "https://glovoapp.com/",
			CareersURL: "https://jobs.glovoapp.com/",
			AboutURL:   "https://about.glovoapp.com/",
			BlogURL:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                9262901,
				IDs:               nil,
				Alias:             "glovo-app",
				Name:              "Glovo",
				Followers:         "333K",
				Employees:         "1K-5K",
				AssociatedMembers: "14,207",
				Verified:          true,
				Date:              mustDate("2026-02-20"),
			},
			GitHubProfile: domain.GitHubProfile{
				Login:     "Glovo",
				Followers: "444",
				Verified:  true,
				Date:      mustDate("2026-02-20"),
			},
			BlindProfile: domain.BlindProfile{
				Alias: "",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias: "",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Percona-LLC-EI_IE283779.11,22.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/index.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Glovo-Jobs-E3424586.htm",
				Jobs:        "", // Becomes empty in the new design version
				Reviews:     "", // Becomes empty in the new design version
				Salaries:    "", // Becomes empty in the new design version
				ReviewsRate: "",
				Verified:    false,
				Date:        mustDate("2026-02-20"),
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "",
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoryCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Staff Software Engineer, Golang",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4374006073/",
							Location:             "Barcelona, Catalonia, Spain",
							CloudProviders:       []domain.CloudProvider{},
							Date:                 mustDate("2026-02-19"),
							WithSalary:           false,
							Remote:               false,
						},
					},
				},
				domain.Rust: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Zig: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Scala: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Elixir: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Clojure: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Haskell: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Erlang: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
			},
			ShortDescription: "Glovo is a multi-category app connecting users with businesses, and couriers, offering on-demand services from local restaurants, grocers and supermarkets, and high street retail stores",
			Industries: []domain.Industry{
				domain.IndustryLogisticsTech,
			},
		},
		{
			ID:         0,  // system
			Type:       "", // system
			Name:       "Paysafe",
			BaseURL:    "https://www.paysafe.com/",
			CareersURL: "https://www.paysafe.com/careers/",
			AboutURL:   "",
			BlogURL:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                132205,
				IDs:               []int{132205, 3492624, 10289663, 10666557},
				Alias:             "paysafe",
				Name:              "Paysafe",
				Followers:         "135K",
				Employees:         "1K-5K",
				AssociatedMembers: "3,508",
				Verified:          true,
				Date:              mustDate("2026-02-20"),
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "",
				Verified: false,
			},
			BlindProfile: domain.BlindProfile{
				Alias: "",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias: "",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Percona-LLC-EI_IE283779.11,22.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/index.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Paysafe-Jobs-E35498.htm",
				Jobs:        "", // Becomes empty in the new design version
				Reviews:     "", // Becomes empty in the new design version
				Salaries:    "", // Becomes empty in the new design version
				ReviewsRate: "",
				Verified:    false,
				Date:        mustDate("2026-02-20"),
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "",
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoryCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Senior Software Engineer (Go)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4366473415/",
							Location:             "Sofia, Sofia City, Bulgaria",
							CloudProviders:       []domain.CloudProvider{},
							Date:                 mustDate("2026-02-19"),
							WithSalary:           false,
							Remote:               false,
						},
					},
				},
				domain.Rust: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Zig: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Scala: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Elixir: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Clojure: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Haskell: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Erlang: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
			},
			ShortDescription: "Payment platform",
			Industries: []domain.Industry{
				domain.IndustryFinTech,
			},
		},
		{
			ID:         0,  // system
			Type:       "", // system
			Name:       "Bitkub",
			BaseURL:    "https://www.bitkub.com/",
			CareersURL: "",
			AboutURL:   "",
			BlogURL:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                13687085,
				IDs:               nil,
				Alias:             "bitkub",
				Name:              "Bitkub",
				Followers:         "44K",
				Employees:         "501-1K",
				AssociatedMembers: "702",
				Verified:          true,
				Date:              mustDate("2026-02-20"),
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "",
				Verified: false,
			},
			BlindProfile: domain.BlindProfile{
				Alias: "",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
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
					GitHubRepositoryCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Software Engineer Lead (Blockchain : Go)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4370630372/",
							Location:             "Bangkok, Bangkok City, Thailand",
							CloudProviders:       []domain.CloudProvider{domain.AWS, domain.GCP},
							Date:                 mustDate("2026-02-19"),
							WithSalary:           false,
							Remote:               false,
						},
					},
				},
				domain.Rust: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Zig: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Scala: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Elixir: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Clojure: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Haskell: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Erlang: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
			},
			ShortDescription: "Digital Asset Exchange",
			Industries:       []domain.Industry{
				// NOP
			},
			Ignore: true, // Blockchain
		},
		{
			ID:         0,  // system
			Type:       "", // system
			Name:       "Lookout",
			BaseURL:    "https://www.lookout.com/",
			CareersURL: "https://www.lookout.com/careers/join-our-team",
			AboutURL:   "https://www.lookout.com/company/about-lookout",
			BlogURL:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                824768,
				IDs:               []int{824768, 9317165},
				Alias:             "lookout",
				Name:              "Lookout",
				Followers:         "66K",
				Employees:         "201-500",
				AssociatedMembers: "520",
				Verified:          true,
				Date:              mustDate("2026-02-26"),
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "",
				Verified: false,
			},
			BlindProfile: domain.BlindProfile{
				Alias: "",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
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
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Rust: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Zig: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Scala: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Elixir: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Clojure: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Haskell: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Erlang: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
			},
			ShortDescription: "",
			Industries: []domain.Industry{
				domain.IndustryCyberSecurity,
			},
		},
		{
			ID:         0,  // system
			Type:       "", // system
			Name:       "MetalBear",
			BaseURL:    "https://metalbear.com/",
			CareersURL: "https://metalbear.com/careers/",
			AboutURL:   "",
			BlogURL:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                81275662,
				IDs:               nil,
				Alias:             "metalbearco",
				Name:              "MetalBear",
				Followers:         "6K",
				Employees:         "11-50",
				AssociatedMembers: "37",
				Verified:          true,
				Date:              mustDate("2026-02-26"),
			},
			GitHubProfile: domain.GitHubProfile{
				Login:     "metalbear-co",
				Followers: "110",
				Verified:  false,
				Date:      mustDate("2026-02-26"),
			},
			BlindProfile: domain.BlindProfile{
				Alias: "",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
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
					GitHubRepositoryCount: 6,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Rust: {
					GitHubRepositoryCount: 31,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Senior Backend Engineer (Rust)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4376718005/",
							Location:             "Warsaw, Mazowieckie, Poland",
							CloudProviders:       []domain.CloudProvider{},
							Date:                 mustDate("2026-02-25"),
							WithSalary:           false,
							Remote:               true,
						},
					},
				},
				domain.Zig: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Scala: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Elixir: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Clojure: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Haskell: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Erlang: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
			},
			ShortDescription: "mirrord introduces a new workflow where the entire team can run code against the staging environment concurrently, without interrupting each other, and from the very first stage of the development cycle",
			Industries: []domain.Industry{
				domain.IndustryDevOps,
			},
		},
		{
			ID:         0,  // system
			Type:       "", // system
			Name:       "mx51",
			BaseURL:    "https://mx51.com/",
			CareersURL: "https://mx51.com/work-with-us",
			AboutURL:   "https://mx51.com/our-story",
			BlogURL:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                31339633,
				IDs:               nil,
				Alias:             "mx51",
				Name:              "mx51",
				Followers:         "5K",
				Employees:         "51-200",
				AssociatedMembers: "66",
				Verified:          false,
				Date:              mustDate("2026-02-26"),
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "",
				Verified: false,
			},
			BlindProfile: domain.BlindProfile{
				Alias: "",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
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
					GitHubRepositoryCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Senior Backend Engineer (Golang)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4330441288/",
							Location:             "Vietnam",
							CloudProviders:       []domain.CloudProvider{domain.AWS},
							Date:                 mustDate("2026-02-25"),
							WithSalary:           false,
							Remote:               true,
						},
					},
				},
				domain.Rust: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Zig: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Scala: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Elixir: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Clojure: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Haskell: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Erlang: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
			},
			ShortDescription: "",
			Industries: []domain.Industry{
				domain.IndustryFinTech,
			},
		},
		{
			ID:         0,  // system
			Type:       "", // system
			Name:       "Bird",
			BaseURL:    "https://www.bird.co/",
			CareersURL: "https://www.bird.co/careers/",
			AboutURL:   "https://www.bird.co/about/",
			BlogURL:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                18359698,
				IDs:               nil,
				Alias:             "birdapp",
				Name:              "Bird",
				Followers:         "69K",
				Employees:         "201-500",
				AssociatedMembers: "942",
				Verified:          true,
				Date:              mustDate("2026-02-26"),
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "",
				Verified: false,
			},
			BlindProfile: domain.BlindProfile{
				Alias: "",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
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
					GitHubRepositoryCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Senior Backend Engineer (Golang)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4374829325/",
							Location:             "United States",
							CloudProviders:       []domain.CloudProvider{domain.AWS},
							Date:                 mustDate("2026-02-25"),
							WithSalary:           true, // $137k/yr - $185k/yr
							Remote:               true,
						},
					},
				},
				domain.Rust: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Zig: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Scala: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Elixir: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Clojure: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Haskell: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Erlang: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
			},
			ShortDescription: "Bird is on a mission to provide eco-friendly transportation for everyone",
			Industries: []domain.Industry{
				domain.IndustryLogisticsTech,
			},
		},
		{
			ID:         0,  // system
			Type:       "", // system
			Name:       "Equativ",
			BaseURL:    "https://www.equativ.com/",
			CareersURL: "https://www.equativ.com/careers",
			AboutURL:   "https://www.equativ.com/company",
			BlogURL:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                202220,
				IDs:               []int{202220, 3488199, 92785323},
				Alias:             "equativ",
				Name:              "Equativ",
				Followers:         "45K",
				Employees:         "501-1K",
				AssociatedMembers: "701",
				Verified:          true,
				Date:              mustDate("2026-02-26"),
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "",
				Verified: false,
			},
			BlindProfile: domain.BlindProfile{
				Alias: "",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
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
					GitHubRepositoryCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Backend Engineer (Golang)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4368197876/",
							Location:             "Berlin, Berlin, Germany",
							CloudProviders:       []domain.CloudProvider{},
							Date:                 mustDate("2026-02-25"),
							WithSalary:           false,
							Remote:               false,
						},
					},
				},
				domain.Rust: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Zig: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Scala: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Elixir: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Clojure: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Haskell: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Erlang: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
			},
			ShortDescription: "",
			Industries: []domain.Industry{
				domain.IndustryAdTech,
			},
		},
		{
			ID:         0,  // system
			Type:       "", // system
			Name:       "Checkbox",
			BaseURL:    "https://www.checkbox.ai/",
			CareersURL: "https://www.checkbox.ai/careers",
			AboutURL:   "",
			BlogURL:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                7943075,
				IDs:               nil,
				Alias:             "checkboxai",
				Name:              "Checkbox",
				Followers:         "12K",
				Employees:         "51-200",
				AssociatedMembers: "77",
				Verified:          true,
				Date:              mustDate("2026-02-26"),
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "",
				Verified: false,
			},
			BlindProfile: domain.BlindProfile{
				Alias: "",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
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
					GitHubRepositoryCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Staff Software Engineer (Golang)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4375718420/",
							Location:             "Greater Sydney Area",
							CloudProviders:       []domain.CloudProvider{},
							Date:                 mustDate("2026-02-24"),
							WithSalary:           false,
							Remote:               false,
						},
					},
				},
				domain.Rust: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Zig: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Scala: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Elixir: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Clojure: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Haskell: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Erlang: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
			},
			ShortDescription: "",
			Industries:       []domain.Industry{
				// NOP
			},
		},
		{
			ID:         0,  // system
			Type:       "", // system
			Name:       "Grocery TV",
			BaseURL:    "https://grocerytv.com/",
			CareersURL: "",
			AboutURL:   "",
			BlogURL:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                28870623,
				IDs:               nil,
				Alias:             "grocery-tv",
				Name:              "Grocery TV",
				Followers:         "4K",
				Employees:         "51-200",
				AssociatedMembers: "77",
				Verified:          true,
				Date:              mustDate("2026-02-26"),
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "",
				Verified: false,
			},
			BlindProfile: domain.BlindProfile{
				Alias: "",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
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
					GitHubRepositoryCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Senior Backend Engineer II – Golang",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4376157523/",
							Location:             "Austin, TX",
							CloudProviders:       []domain.CloudProvider{},
							Date:                 mustDate("2026-02-22"),
							WithSalary:           false,
							Remote:               false,
						},
					},
				},
				domain.Rust: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Zig: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Scala: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Elixir: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Clojure: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Haskell: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Erlang: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
			},
			ShortDescription: "Retail media platform",
			Industries: []domain.Industry{
				domain.IndustryAdTech,
			},
		},
		{
			ID:         0,  // system
			Type:       "", // system
			Name:       "Pismo",
			BaseURL:    "https://www.pismo.io/",
			CareersURL: "https://www.pismo.io/careers/",
			AboutURL:   "https://www.pismo.io/about-us/",
			BlogURL:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                3844919,
				IDs:               nil,
				Alias:             "pismo",
				Name:              "Pismo",
				Followers:         "63K",
				Employees:         "201-500",
				AssociatedMembers: "575",
				Verified:          true,
				Date:              mustDate("2026-02-26"),
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "",
				Verified: false,
			},
			BlindProfile: domain.BlindProfile{
				Alias: "",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
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
					GitHubRepositoryCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Senior Software Engineer – Golang",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4374303417/",
							Location:             "Warsaw, Mazowieckie, Poland",
							CloudProviders:       []domain.CloudProvider{},
							Date:                 mustDate("2026-02-20"),
							WithSalary:           false,
							Remote:               false,
						},
					},
				},
				domain.Rust: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Zig: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Scala: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Elixir: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Clojure: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Haskell: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Erlang: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
			},
			ShortDescription: "Platform for payments and banking",
			Industries: []domain.Industry{
				domain.IndustryFinTech,
			},
		},
		{
			ID:         0,  // system
			Type:       "", // system
			Name:       "Exinity",
			BaseURL:    "https://www.exinity.com/",
			CareersURL: "",
			AboutURL:   "",
			BlogURL:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                34900724,
				IDs:               nil,
				Alias:             "exinity",
				Name:              "Exinity",
				Followers:         "59K",
				Employees:         "501-1K",
				AssociatedMembers: "848",
				Verified:          true,
				Date:              mustDate("2026-02-26"),
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "",
				Verified: false,
			},
			BlindProfile: domain.BlindProfile{
				Alias: "",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
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
					GitHubRepositoryCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Senior Golang Developer",
							ShortDescription:     "Trading Systems",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4376577496/",
							Location:             "Limassol, Limassol, Cyprus",
							CloudProviders:       []domain.CloudProvider{},
							Date:                 mustDate("2026-02-24"),
							WithSalary:           false,
							Remote:               false,
						},
					},
				},
				domain.Rust: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Zig: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Scala: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Elixir: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Clojure: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Haskell: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Erlang: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
			},
			ShortDescription: "",
			Industries: []domain.Industry{
				domain.IndustryFinTech,
			},
		},
		{
			ID:         0,  // system
			Type:       "", // system
			Name:       "Graphcore",
			BaseURL:    "https://www.graphcore.ai/",
			CareersURL: "https://www.graphcore.ai/jobs",
			AboutURL:   "https://www.graphcore.ai/about",
			BlogURL:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                10812092,
				IDs:               nil,
				Alias:             "graphcore",
				Name:              "Graphcore",
				Followers:         "43K",
				Employees:         "201-500",
				AssociatedMembers: "555",
				Verified:          true,
				Date:              mustDate("2026-02-26"),
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "",
				Verified: false,
			},
			BlindProfile: domain.BlindProfile{
				Alias: "",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
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
					GitHubRepositoryCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Senior Kubernetes Software Engineer (Go)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4344851946/",
							Location:             "Gdańsk, Pomorskie, Poland",
							CloudProviders:       []domain.CloudProvider{},
							Date:                 mustDate("2026-02-21"),
							WithSalary:           false,
							Remote:               false,
						},
					},
				},
				domain.Rust: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Zig: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Scala: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Elixir: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Clojure: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Haskell: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Erlang: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
			},
			ShortDescription: "Semiconductor Manufacturing",
			Industries:       []domain.Industry{
				// NOP
			},
		},
	}
}
