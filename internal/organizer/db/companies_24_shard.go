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
							Date:                 mustDate("2026-03-07", "2026-02-13"),
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
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-ShopFully-EI_IE1472579.11,20.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/ShopFully-Reviews-E1472579.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/ShopFully-Jobs-E1472579.htm",
				Jobs:        "", // Becomes empty in the new design version
				Reviews:     "", // Becomes empty in the new design version
				Salaries:    "", // Becomes empty in the new design version
				ReviewsRate: "",
				Verified:    false,
				Date:        mustDate("2026-02-26"),
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
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-FXC-Intelligence-EI_IE988093.11,27.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/FXC-Intelligence-Reviews-E988093.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/FXC-Intelligence-Jobs-E988093.htm",
				Jobs:        "", // Becomes empty in the new design version
				Reviews:     "", // Becomes empty in the new design version
				Salaries:    "", // Becomes empty in the new design version
				ReviewsRate: "",
				Verified:    false,
				Date:        mustDate("2026-02-26"),
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
				Followers:         "79K",
				Employees:         "5K-10K",
				AssociatedMembers: "3,466",
				Verified:          true,
				Date:              mustDate("2026-03-14"),
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "",
				Verified: false,
			},
			BlindProfile: domain.BlindProfile{
				Alias: "",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias: "",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Gen-EI_IE8043746.11,14.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Gen-Reviews-E8043746.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Gen-Jobs-E8043746.htm",
				Jobs:        "", // Becomes empty in the new design version
				Reviews:     "", // Becomes empty in the new design version
				Salaries:    "", // Becomes empty in the new design version
				ReviewsRate: "",
				Verified:    false,
				Date:        mustDate("2026-02-26"),
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
							Date:                 mustDate("2026-03-13", "2026-02-18"),
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
				Followers:         "341K",
				Employees:         "1K-5K",
				AssociatedMembers: "14,285",
				Verified:          true,
				Date:              mustDate("2026-03-14"),
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
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Glovo-EI_IE3424586.11,16.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Glovo-Reviews-E3424586.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Glovo-Jobs-E3424586.htm",
				Jobs:        "", // Becomes empty in the new design version
				Reviews:     "", // Becomes empty in the new design version
				Salaries:    "", // Becomes empty in the new design version
				ReviewsRate: "",
				Verified:    false,
				Date:        mustDate("2026-02-26"),
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
							Date:                 mustDate("2026-03-13", "2026-02-19"),
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
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Paysafe-EI_IE35498.11,18.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Paysafe-Reviews-E35498.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Paysafe-Jobs-E35498.htm",
				Jobs:        "", // Becomes empty in the new design version
				Reviews:     "", // Becomes empty in the new design version
				Salaries:    "", // Becomes empty in the new design version
				ReviewsRate: "",
				Verified:    false,
				Date:        mustDate("2026-02-26"),
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
				AssociatedMembers: "522",
				Verified:          true,
				Date:              mustDate("2026-03-02"),
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "",
				Verified: false,
			},
			BlindProfile: domain.BlindProfile{
				Alias: "",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias: "",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Lookout-EI_IE318794.11,18.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Lookout-Reviews-E318794.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Lookout-Jobs-E318794.htm",
				Jobs:        "", // Becomes empty in the new design version
				Reviews:     "", // Becomes empty in the new design version
				Salaries:    "", // Becomes empty in the new design version
				ReviewsRate: "",
				Verified:    false,
				Date:        mustDate("2026-03-03"),
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
							Title:                "Staff Backend Software Engineer (Java/Scala)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4377020542/",
							Location:             "Canada",
							CloudProviders:       []domain.CloudProvider{},
							Date:                 mustDate("2026-02-28"),
							WithSalary:           false,
							Remote:               true,
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
				Login:     "mx51",
				Followers: "19",
				Verified:  true,
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
					GitHubRepositoryCount: 9,
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
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Bird-EI_IE2143140.11,15.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Bird-Reviews-E2143140.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Bird-Jobs-E2143140.htm",
				Jobs:        "", // Becomes empty in the new design version
				Reviews:     "", // Becomes empty in the new design version
				Salaries:    "", // Becomes empty in the new design version
				ReviewsRate: "",
				Verified:    false,
				Date:        mustDate("2026-02-26"),
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
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Equativ-EI_IE423392.11,18.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Equativ-Reviews-E423392.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Equativ-Jobs-E423392.htm",
				Jobs:        "", // Becomes empty in the new design version
				Reviews:     "", // Becomes empty in the new design version
				Salaries:    "", // Becomes empty in the new design version
				ReviewsRate: "",
				Verified:    false,
				Date:        mustDate("2026-02-26"),
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
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Pismo-EI_IE2744250.11,16.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Pismo-Reviews-E2744250.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Pismo-Jobs-E2744250.htm",
				Jobs:        "", // Becomes empty in the new design version
				Reviews:     "", // Becomes empty in the new design version
				Salaries:    "", // Becomes empty in the new design version
				ReviewsRate: "",
				Verified:    false,
				Date:        mustDate("2026-02-26"),
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
						{
							Title:                "Senior Software Engineer – Golang",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4377241502/",
							Location:             "Warsaw, Mazowieckie, Poland",
							CloudProviders:       []domain.CloudProvider{},
							Date:                 mustDate("2026-02-27"),
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
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Exinity-EI_IE4254120.11,18.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Exinity-Reviews-E4254120.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Exinity-Jobs-E4254120.htm",
				Jobs:        "", // Becomes empty in the new design version
				Reviews:     "", // Becomes empty in the new design version
				Salaries:    "", // Becomes empty in the new design version
				ReviewsRate: "",
				Verified:    false,
				Date:        mustDate("2026-02-26"),
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
				AssociatedMembers: "578",
				Verified:          true,
				Date:              mustDate("2026-03-14"),
			},
			GitHubProfile: domain.GitHubProfile{
				Login:     "graphcore",
				Followers: "323",
				Verified:  true,
				Date:      mustDate("2026-03-14"),
			},
			BlindProfile: domain.BlindProfile{
				Alias: "",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias: "",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Graphcore-EI_IE2500374.11,20.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Graphcore-Reviews-E2500374.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Graphcore-Jobs-E2500374.htm",
				Jobs:        "", // Becomes empty in the new design version
				Reviews:     "", // Becomes empty in the new design version
				Salaries:    "", // Becomes empty in the new design version
				ReviewsRate: "",
				Verified:    false,
				Date:        mustDate("2026-02-26"),
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
							Date:                 mustDate("2026-03-13", "2026-02-21"),
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
		{
			ID:         0,  // system
			Type:       "", // system
			Name:       "FeedMe",
			BaseURL:    "https://feedme.ai/",
			CareersURL: "https://feedme.ai/company/career",
			AboutURL:   "https://feedme.ai/company/about",
			BlogURL:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                68880880,
				IDs:               nil,
				Alias:             "feedme-pos-sdn-bhd",
				Name:              "FeedMe",
				Followers:         "8K",
				Employees:         "51-200",
				AssociatedMembers: "64",
				Verified:          true,
				Date:              mustDate("2026-02-26"),
			},
			GitHubProfile: domain.GitHubProfile{
				Login:     "feedmepos",
				Followers: "22",
				Verified:  true,
				Date:      mustDate("2026-02-26"),
			},
			BlindProfile: domain.BlindProfile{
				Alias: "",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias: "",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Feedme-POS-EI_IE4709290.11,21.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Feedme-POS-Reviews-E4709290.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Feedme-POS-Jobs-E4709290.htm",
				Jobs:        "", // Becomes empty in the new design version
				Reviews:     "", // Becomes empty in the new design version
				Salaries:    "", // Becomes empty in the new design version
				ReviewsRate: "",
				Verified:    false,
				Date:        mustDate("2026-02-26"),
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "",
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoryCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Senior Software Engineer",
							ShortDescription:     "Willingness and enthusiasm to learn and write Go (Golang)",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4374571187/",
							Location:             "Johor Bahru, Johore, Malaysia",
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
			ShortDescription: "We manage to help F&B owners sell online, manage inventory, run a busy kitchen, book appointments, engage loyal customers, cut their labour costs, and improve their operation hassle",
			Industries:       []domain.Industry{
				// NOP
			},
		},
		{
			ID:         0,  // system
			Type:       "", // system
			Name:       "Wavestore",
			BaseURL:    "https://www.wavestore.com/",
			CareersURL: "",
			AboutURL:   "https://www.wavestore.com/about-us/our-company",
			BlogURL:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                919408,
				IDs:               nil,
				Alias:             "wavestore-global",
				Name:              "Wavestore",
				Followers:         "5K",
				Employees:         "11-50",
				AssociatedMembers: "46",
				Verified:          false,
				Date:              mustDate("2026-03-02"),
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "",
				Verified: false,
			},
			BlindProfile: domain.BlindProfile{
				Alias: "",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
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
							Title:                "Rust Developer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4378869197/",
							Location:             "Ottawa, ON",
							CloudProviders:       []domain.CloudProvider{domain.AWS, domain.GCP, domain.Azure}, //
							Date:                 mustDate("2026-03-02"),
							WithSalary:           true, // ca$55k/yr - ca$80k/yr
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
			ShortDescription: "Video Management Software solutions",
			Industries:       []domain.Industry{
				// NOP
			},
		},
		{
			ID:         0,  // system
			Type:       "", // system
			Name:       "Narvar",
			BaseURL:    "https://corp.narvar.com/",
			CareersURL: "https://corp.narvar.com/careers",
			AboutURL:   "https://corp.narvar.com/about",
			BlogURL:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                3064987,
				IDs:               nil,
				Alias:             "narvar",
				Name:              "Narvar",
				Followers:         "31K",
				Employees:         "201-500",
				AssociatedMembers: "376",
				Verified:          true,
				Date:              mustDate("2026-03-02"),
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "",
				Verified: false,
			},
			BlindProfile: domain.BlindProfile{
				Alias: "",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias: "",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Narvar-EI_IE1268435.11,17.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Narvar-Reviews-E1268435.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Narvar-Jobs-E1268435.htm",
				Jobs:        "", // Becomes empty in the new design version
				Reviews:     "", // Becomes empty in the new design version
				Salaries:    "", // Becomes empty in the new design version
				ReviewsRate: "",
				Verified:    true,
				Date:        mustDate("2026-03-11"),
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
							Title:                "Senior Rust Software Developer II",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4374292185/",
							Location:             "Canada",
							CloudProviders:       []domain.CloudProvider{},
							Date:                 mustDate("2026-03-03", "2026-02-26"),
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
			ShortDescription: "Platform for personalization",
			Industries:       []domain.Industry{
				// NOP
			},
		},
		{
			ID:         0,  // system
			Type:       "", // system
			Name:       "Balyasny Asset Management L.P.",
			BaseURL:    "https://www.bamfunds.com/",
			CareersURL: "https://www.bamfunds.com/careers",
			AboutURL:   "https://www.bamfunds.com/about-us",
			BlogURL:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                1383805,
				IDs:               nil,
				Alias:             "balyasny-asset-management-l.p.",
				Name:              "Balyasny Asset Management L.P.",
				Followers:         "159K",
				Employees:         "1K-5K",
				AssociatedMembers: "2,408",
				Verified:          true,
				Date:              mustDate("2026-03-03"),
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "",
				Verified: false,
			},
			BlindProfile: domain.BlindProfile{
				Alias: "",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
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
							Title:                "Senior Rust Engineer",
							ShortDescription:     "Trade Processing Technology",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4377688601/",
							Location:             "Warsaw Metropolitan Area",
							CloudProviders:       []domain.CloudProvider{},
							Date:                 mustDate("2026-02-28"),
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
			ShortDescription: "Investment firm",
			Industries: []domain.Industry{
				domain.IndustryFinTech,
			},
		},
		{
			ID:         0,  // system
			Type:       "", // system
			Name:       "Flight Science",
			BaseURL:    "https://www.flightscience.ai/",
			CareersURL: "",
			AboutURL:   "https://www.flightscience.ai/company",
			BlogURL:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                102052784,
				IDs:               nil,
				Alias:             "flight-science",
				Name:              "Flight Science",
				Followers:         "629",
				Employees:         "11-50",
				AssociatedMembers: "10",
				Verified:          false,
				Date:              mustDate("2026-03-03"),
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "",
				Verified: false,
			},
			BlindProfile: domain.BlindProfile{
				Alias: "",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
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
							Title:                "Rust Backend Engineer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4376901695/",
							Location:             "West Hollywood, CA",
							CloudProviders:       []domain.CloudProvider{domain.AWS},
							Date:                 mustDate("2026-02-27"),
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
			ShortDescription: "Platform for airlines",
			Industries:       []domain.Industry{
				// NOP
			},
		},
		{
			ID:         0,  // system
			Type:       "", // system
			Name:       "Swiss Re",
			BaseURL:    "https://www.swissre.com/",
			CareersURL: "https://www.swissre.com/careers.html",
			AboutURL:   "https://www.swissre.com/about-us.html",
			BlogURL:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                3845,
				IDs:               []int{3845, 15073937},
				Alias:             "swiss-re",
				Name:              "Swiss Re",
				Followers:         "585K",
				Employees:         "10K+",
				AssociatedMembers: "15,169",
				Verified:          true,
				Date:              mustDate("2026-03-03"),
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "",
				Verified: false,
			},
			BlindProfile: domain.BlindProfile{
				Alias: "",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias: "",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Swiss-Re-EI_IE4716.11,19.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Swiss-Re-Reviews-E4716.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Swiss-Re-Jobs-E4716.htm",
				Jobs:        "", // Becomes empty in the new design version
				Reviews:     "", // Becomes empty in the new design version
				Salaries:    "", // Becomes empty in the new design version
				ReviewsRate: "",
				Verified:    false,
				Date:        mustDate("2026-03-03"),
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "",
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoryCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Senior Backend Engineer – Golang",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4370033475/",
							Location:             "Madrid, Community of Madrid, Spain",
							CloudProviders:       []domain.CloudProvider{},
							Date:                 mustDate("2026-03-02"),
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
			ShortDescription: "Provider of reinsurance, insurance and other insurance-based forms of risk transfer",
			Industries: []domain.Industry{
				domain.IndustryInsurTech,
			},
		},
		{
			ID:         0,  // system
			Type:       "", // system
			Name:       "Tide",
			BaseURL:    "https://www.tide.co/",
			CareersURL: "https://www.tide.co/careers/",
			AboutURL:   "https://www.tide.co/about/",
			BlogURL:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                10877408,
				IDs:               nil,
				Alias:             "tide-banking",
				Name:              "Tide",
				Followers:         "610K",
				Employees:         "1K-5K",
				AssociatedMembers: "3,149",
				Verified:          true,
				Date:              mustDate("2026-03-03"),
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "",
				Verified: false,
			},
			BlindProfile: domain.BlindProfile{
				Alias: "",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias: "",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Tide-EI_IE2058934.11,15.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Tide-Reviews-E2058934.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Tide-Jobs-E2058934.htm",
				Jobs:        "", // Becomes empty in the new design version
				Reviews:     "", // Becomes empty in the new design version
				Salaries:    "", // Becomes empty in the new design version
				ReviewsRate: "",
				Verified:    false,
				Date:        mustDate("2026-03-03"),
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "",
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoryCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Senior Golang Engineer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4358249813/",
							Location:             "Belgrade, Serbia",
							CloudProviders:       []domain.CloudProvider{},
							Date:                 mustDate("2026-02-28"),
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
			ShortDescription: "SME banking & business management",
			Industries: []domain.Industry{
				domain.IndustryFinTech,
			},
		},
		{
			ID:         0,  // system
			Type:       "", // system
			Name:       "Apex Fintech Solutions",
			BaseURL:    "https://apexfintechsolutions.com/",
			CareersURL: "https://careers.apexfintechsolutions.com/",
			AboutURL:   "https://apexfintechsolutions.com/about",
			BlogURL:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                3008681,
				IDs:               []int{3008681, 68595869},
				Alias:             "apex-fintech",
				Name:              "Apex Fintech Solutions",
				Followers:         "25K",
				Employees:         "501-1K",
				AssociatedMembers: "1,042",
				Verified:          true,
				Date:              mustDate("2026-03-03"),
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "",
				Verified: false,
			},
			BlindProfile: domain.BlindProfile{
				Alias: "",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
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
							Title:                "Senior Software Engineer (Golang)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4363561357/",
							Location:             "Austin, TX",
							CloudProviders:       []domain.CloudProvider{},
							Date:                 mustDate("2026-02-28"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Staff Software Engineer (Golang)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4371462871/",
							Location:             "Austin, TX",
							CloudProviders:       []domain.CloudProvider{},
							Date:                 mustDate("2026-03-06"),
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
			Name:       "HDFC Bank",
			BaseURL:    "https://www.hdfc.bank.in/",
			CareersURL: "https://www.hdfc.bank.in/careers",
			AboutURL:   "https://www.hdfc.bank.in/about-us",
			BlogURL:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                164151,
				IDs:               nil,
				Alias:             "hdfc-bank",
				Name:              "HDFC Bank",
				Followers:         "4M",
				Employees:         "10K+",
				AssociatedMembers: "211,044",
				Verified:          true,
				Date:              mustDate("2026-03-03"),
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "",
				Verified: false,
			},
			BlindProfile: domain.BlindProfile{
				Alias: "",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
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
							Title:                "Senior Backend Engineer- Golang",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4375512699/",
							Location:             "Bengaluru, Karnataka, India",
							CloudProviders:       []domain.CloudProvider{},
							Date:                 mustDate("2026-02-27"),
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
			ShortDescription: "India's largest private sector bank",
			Industries: []domain.Industry{
				domain.IndustryFinTech,
			},
		},
		{
			ID:         0,                         // system
			Type:       domain.CompanyTypeStartup, // system
			Name:       "Trio",
			BaseURL:    "https://www.trio.so/",
			CareersURL: "",
			AboutURL:   "https://www.trio.so/about-us",
			BlogURL:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                92430741,
				IDs:               nil,
				Alias:             "trio-so",
				Name:              "Trio",
				Followers:         "9K",
				Employees:         "11-50",
				AssociatedMembers: "51",
				Verified:          true,
				Date:              mustDate("2026-03-03"),
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "",
				Verified: false,
			},
			BlindProfile: domain.BlindProfile{
				Alias: "",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
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
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4377798858/",
							Location:             "Dhaka, Dhaka, Bangladesh",
							CloudProviders:       []domain.CloudProvider{},
							Date:                 mustDate("2026-03-03"),
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
			ShortDescription: "MDM (Mobile Device Management) and EDR (Endpoint Detection & Response) solutions",
			Industries: []domain.Industry{
				domain.IndustryCyberSecurity,
			},
		},
		{
			ID:         0,  // system
			Type:       "", // system
			Name:       "EarnIn",
			BaseURL:    "https://www.earnin.com/",
			CareersURL: "https://www.earnin.com/careers",
			AboutURL:   "https://www.earnin.com/about",
			BlogURL:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                2837417,
				IDs:               nil,
				Alias:             "earnin",
				Name:              "EarnIn",
				Followers:         "199K",
				Employees:         "201-500",
				AssociatedMembers: "586",
				Verified:          true,
				Date:              mustDate("2026-03-03"),
			},
			GitHubProfile: domain.GitHubProfile{
				Login:     "swissre",
				Followers: "30",
				Verified:  true,
				Date:      mustDate("2026-03-03"),
			},
			BlindProfile: domain.BlindProfile{
				Alias: "",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias: "",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-EarnIn-EI_IE962632.11,17.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/EarnIn-Reviews-E962632.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/EarnIn-Jobs-E962632.htm",
				Jobs:        "", // Becomes empty in the new design version
				Reviews:     "", // Becomes empty in the new design version
				Salaries:    "", // Becomes empty in the new design version
				ReviewsRate: "",
				Verified:    false,
				Date:        mustDate("2026-03-03"),
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "",
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoryCount: 1,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Senior Software Engineer (Go)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4369220734/",
							Location:             "Mexico",
							CloudProviders:       []domain.CloudProvider{},
							Date:                 mustDate("2026-02-27"),
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
			Name:       "happyhotel",
			BaseURL:    "https://www.happyhotel.io/",
			CareersURL: "https://www.happyhotel.io/en/career",
			AboutURL:   "",
			BlogURL:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                18842021,
				IDs:               nil,
				Alias:             "happyhotel",
				Name:              "happyhotel",
				Followers:         "3K",
				Employees:         "11-50",
				AssociatedMembers: "39",
				Verified:          true,
				Date:              mustDate("2026-03-03"),
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "",
				Verified: false,
			},
			BlindProfile: domain.BlindProfile{
				Alias: "",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
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
							Title:                "Senior Software Engineer (Go)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4378343310/",
							Location:             "Offenburg, Baden-Württemberg, Germany",
							CloudProviders:       []domain.CloudProvider{domain.AWS},
							Date:                 mustDate("2026-02-27"),
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
			ShortDescription: "The revenue management software for the hotel industry",
			Industries:       []domain.Industry{},
		},
		{
			ID:         0,  // system
			Type:       "", // system
			Name:       "Griffin",
			BaseURL:    "https://griffin.com/",
			CareersURL: "https://griffin.com/careers",
			AboutURL:   "https://griffin.com/about",
			BlogURL:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                11087261,
				IDs:               nil,
				Alias:             "griffin-bank",
				Name:              "Griffin",
				Followers:         "15K",
				Employees:         "51-200",
				AssociatedMembers: "144",
				Verified:          true,
				Date:              mustDate("2026-03-03"),
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "",
				Verified: false,
			},
			BlindProfile: domain.BlindProfile{
				Alias: "",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
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
					Vacancies: []domain.Vacancy{
						{
							Title:                "Software Engineer",
							ShortDescription:     "Our backend stack is Clojure, FoundationDB, Bazel, Kubernetes and AWS",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4378321804/",
							Location:             "United Kingdom",
							CloudProviders:       []domain.CloudProvider{domain.AWS},
							Date:                 mustDate("2026-02-27"),
							WithSalary:           false,
							Remote:               true,
						},
					},
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
			ShortDescription: "Embed bank accounts, payments, and all kinds of financial products quickly via our simple API",
			Industries: []domain.Industry{
				domain.IndustryFinTech,
			},
		},
		{
			ID:         0,  // system
			Type:       "", // system
			Name:       "Glia",
			BaseURL:    "https://www.glia.com/",
			CareersURL: "https://www.glia.com/careers",
			AboutURL:   "https://www.glia.com/about",
			BlogURL:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                2687116,
				IDs:               nil,
				Alias:             "gliainc",
				Name:              "Glia",
				Followers:         "34K",
				Employees:         "201-500",
				AssociatedMembers: "449",
				Verified:          true,
				Date:              mustDate("2026-03-03"),
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "",
				Verified: false,
			},
			BlindProfile: domain.BlindProfile{
				Alias: "",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias: "",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Glia-EI_IE1041718.11,15.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Glia-Reviews-E1041718.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Glia-Jobs-E1041718.htm",
				Jobs:        "", // Becomes empty in the new design version
				Reviews:     "", // Becomes empty in the new design version
				Salaries:    "", // Becomes empty in the new design version
				ReviewsRate: "",
				Verified:    false,
				Date:        mustDate("2026-03-03"),
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
					Vacancies: []domain.Vacancy{
						{
							Title:                "Software Engineer",
							ShortDescription:     "Backend: Elixir, Go and Ruby",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4375555017/",
							Location:             "Canada",
							CloudProviders:       []domain.CloudProvider{},
							Date:                 mustDate("2026-02-27"),
							WithSalary:           false,
							Remote:               true,
						},
					},
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
			ShortDescription: "Platform for Intelligent Banking Interactions",
			Industries: []domain.Industry{
				domain.IndustryFinTech,
			},
		},
		{
			ID:         0,  // system
			Type:       "", // system
			Name:       "OTR Solutions",
			BaseURL:    "https://otrsolutions.com/",
			CareersURL: "https://otrsolutions.com/careers",
			AboutURL:   "https://otrsolutions.com/about-us",
			BlogURL:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                3488196,
				IDs:               nil,
				Alias:             "otrsolutions",
				Name:              "OTR Solutions",
				Followers:         "48K",
				Employees:         "201-500",
				AssociatedMembers: "497",
				Verified:          true,
				Date:              mustDate("2026-03-08"),
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "",
				Verified: false,
			},
			BlindProfile: domain.BlindProfile{
				Alias: "",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias: "",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-OTR-Solutions-EI_IE788595.11,24.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/OTR-Solutions-Reviews-E788595.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/OTR-Solutions-Jobs-E788595.htm",
				Jobs:        "", // Becomes empty in the new design version
				Reviews:     "", // Becomes empty in the new design version
				Salaries:    "", // Becomes empty in the new design version
				ReviewsRate: "",
				Verified:    true,
				Date:        mustDate("2026-03-11"),
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
					Vacancies: []domain.Vacancy{
						{
							Title:                "Senior Clojure Backend Engineer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4379896178/",
							Location:             "Roswell, GA",
							CloudProviders:       []domain.CloudProvider{domain.AWS},
							Date:                 mustDate("2026-03-06"),
							WithSalary:           false,
							Remote:               false,
						},
					},
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
			Name:       "Techex",
			BaseURL:    "https://www.techex.tv/",
			CareersURL: "https://www.techex.tv/pages/careers",
			AboutURL:   "",
			BlogURL:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                65589,
				IDs:               nil,
				Alias:             "techex",
				Name:              "Techex",
				Followers:         "5K",
				Employees:         "51-200",
				AssociatedMembers: "128",
				Verified:          true,
				Date:              mustDate("2026-03-08"),
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "",
				Verified: false,
			},
			BlindProfile: domain.BlindProfile{
				Alias: "",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
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
							Title:                "Senior Rust Developer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4381005660/",
							Location:             "Bracknell, England, United Kingdom",
							CloudProviders:       []domain.CloudProvider{},
							Date:                 mustDate("2026-03-07"),
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
			ShortDescription: "Video over IP",
			Industries:       []domain.Industry{
				// NOP
			},
		},
		{
			ID:         0,  // system
			Type:       "", // system
			Name:       "simplesurance",
			BaseURL:    "https://www.simplesurance.com/",
			CareersURL: "https://www.simplesurance.com/careers/",
			AboutURL:   "https://www.simplesurance.com/who-we-are/",
			BlogURL:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                2867578,
				IDs:               nil,
				Alias:             "simplesurance",
				Name:              "simplesurance",
				Followers:         "7K",
				Employees:         "51-200",
				AssociatedMembers: "122",
				Verified:          true,
				Date:              mustDate("2026-03-08"),
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "",
				Verified: false,
			},
			BlindProfile: domain.BlindProfile{
				Alias: "",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
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
							Title:                "Senior Golang Software Engineer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4380966201/",
							Location:             "Lisbon, Lisbon, Portugal",
							CloudProviders:       []domain.CloudProvider{},
							Date:                 mustDate("2026-03-05"),
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
				domain.IndustryInsurTech,
			},
		},
		{
			ID:         0,  // system
			Type:       "", // system
			Name:       "Spacelift",
			BaseURL:    "https://spacelift.io/",
			CareersURL: "https://careers.spacelift.io/",
			AboutURL:   "https://spacelift.io/about-us",
			BlogURL:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                34594817,
				IDs:               nil,
				Alias:             "spacelift-io",
				Name:              "Spacelift",
				Followers:         "18K",
				Employees:         "51-200",
				AssociatedMembers: "143",
				Verified:          true,
				Date:              mustDate("2026-03-08"),
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "",
				Verified: false,
			},
			BlindProfile: domain.BlindProfile{
				Alias: "",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias: "",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Spacelift-EI_IE3644715.11,20.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Spacelift-Reviews-E3644715.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Spacelift-Jobs-E3644715.htm",
				Jobs:        "", // Becomes empty in the new design version
				Reviews:     "", // Becomes empty in the new design version
				Salaries:    "", // Becomes empty in the new design version
				ReviewsRate: "",
				Verified:    false,
				Date:        mustDate("2026-03-11"),
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "",
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoryCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Product Software Engineer (Go)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4379853305/",
							Location:             "European Economic Area",
							CloudProviders:       []domain.CloudProvider{},
							Date:                 mustDate("2026-03-06"),
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
			ShortDescription: "Spacelift is an infrastructure orchestration platform",
			Industries: []domain.Industry{
				domain.IndustryDevOps,
			},
		},
		{
			ID:         0,  // system
			Type:       "", // system
			Name:       "Raidiam",
			BaseURL:    "https://www.raidiam.com/",
			CareersURL: "https://www.raidiam.com/company/careers",
			AboutURL:   "https://www.raidiam.com/company/about-us",
			BlogURL:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                10985183,
				IDs:               nil,
				Alias:             "raidiam",
				Name:              "Raidiam",
				Followers:         "5K",
				Employees:         "51-200",
				AssociatedMembers: "69",
				Verified:          true,
				Date:              mustDate("2026-03-08"),
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "",
				Verified: false,
			},
			BlindProfile: domain.BlindProfile{
				Alias: "",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
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
							Title:                "Backend Engineer (Go)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4380769640/",
							Location:             "São Paulo, São Paulo, Brazil",
							CloudProviders:       []domain.CloudProvider{},
							Date:                 mustDate("2026-03-05"),
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
			Name:       "Skool",
			BaseURL:    "https://www.skool.com/",
			CareersURL: "https://www.skool.com/careers",
			AboutURL:   "",
			BlogURL:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                34674684,
				IDs:               nil,
				Alias:             "skool-com",
				Name:              "Skool",
				Followers:         "13K",
				Employees:         "11-50",
				AssociatedMembers: "431",
				Verified:          true,
				Date:              mustDate("2026-03-08"),
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "",
				Verified: false,
			},
			BlindProfile: domain.BlindProfile{
				Alias: "",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
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
							Title:                "Principal Software Engineer, Backend (Go)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4362002600/",
							Location:             "United States",
							CloudProviders:       []domain.CloudProvider{},
							Date:                 mustDate("2026-03-04"),
							WithSalary:           true, // $300k/yr - $450k/yr + bonus, stock options
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
			ShortDescription: "Skool is a community platform. You can discover communities or create your own.",
			Industries: []domain.Industry{
				domain.IndustrySocialMedia,
			},
		},
		{
			ID:         0,  // system
			Type:       "", // system
			Name:       "Sprout Social, Inc.",
			BaseURL:    "https://sproutsocial.com/",
			CareersURL: "https://sproutsocial.com/careers/",
			AboutURL:   "https://sproutsocial.com/about/",
			BlogURL:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                1175268,
				IDs:               nil,
				Alias:             "sprout-social-inc-",
				Name:              "Sprout Social, Inc.",
				Followers:         "169K",
				Employees:         "1K-5K",
				AssociatedMembers: "1,793",
				Verified:          true,
				Date:              mustDate("2026-03-10"),
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "",
				Verified: false,
			},
			BlindProfile: domain.BlindProfile{
				Alias: "",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
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
					Vacancies: []domain.Vacancy{
						{
							Title:                "Senior Software Engineer – Back-end (Scala)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4381069460/",
							Location:             "Poland",
							CloudProviders:       []domain.CloudProvider{},
							Date:                 mustDate("2026-03-07"),
							WithSalary:           false,
							Remote:               true,
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
			ShortDescription: "Social media management and analytics software",
			Industries: []domain.Industry{
				domain.IndustryMarTech,
			},
		},
		{
			ID:         0,  // system
			Type:       "", // system
			Name:       "Torchlight®",
			BaseURL:    "https://www.torchlight.ai/",
			CareersURL: "https://www.torchlight.ai/careers/",
			AboutURL:   "https://www.torchlight.ai/about/",
			BlogURL:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                79459437,
				IDs:               nil,
				Alias:             "torchlightai",
				Name:              "Torchlight®",
				Followers:         "20K",
				Employees:         "11-50",
				AssociatedMembers: "47",
				Verified:          true,
				Date:              mustDate("2026-03-11"),
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "",
				Verified: false,
			},
			BlindProfile: domain.BlindProfile{
				Alias: "",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
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
							Title:                "Senior Software Engineer (Golang)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4372132912/",
							Location:             "United States",
							CloudProviders:       []domain.CloudProvider{},
							Date:                 mustDate("2026-03-10"),
							WithSalary:           true, // $175k/yr - $195k/yr
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
				domain.IndustryDefenseTech,
			},
		},
		{
			ID:         0,  // system
			Type:       "", // system
			Name:       "Xperi Inc.",
			BaseURL:    "https://xperi.com/",
			CareersURL: "https://xperi.com/careers/",
			AboutURL:   "",
			BlogURL:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                17962490,
				IDs:               nil,
				Alias:             "xperi-inc",
				Name:              "Xperi Inc.",
				Followers:         "35K",
				Employees:         "1K-5K",
				AssociatedMembers: "1,581",
				Verified:          true,
				Date:              mustDate("2026-03-11"),
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "",
				Verified: false,
			},
			BlindProfile: domain.BlindProfile{
				Alias: "",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias: "",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Xperi-EI_IE1563777.11,16.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Xperi-Reviews-E1563777.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Xperi-Jobs-E1563777.htm",
				Jobs:        "", // Becomes empty in the new design version
				Reviews:     "", // Becomes empty in the new design version
				Salaries:    "", // Becomes empty in the new design version
				ReviewsRate: "",
				Verified:    true,
				Date:        mustDate("2026-03-14"),
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "",
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoryCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Senior Golang Engineer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4382297495/",
							Location:             "Warsaw, Mazowieckie, Poland",
							CloudProviders:       []domain.CloudProvider{},
							Date:                 mustDate("2026-03-10"),
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
			Name:       "Planitar Inc.",
			BaseURL:    "https://goiguide.com/",
			CareersURL: "https://goiguide.com/careers",
			AboutURL:   "https://goiguide.com/about",
			BlogURL:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                2990083,
				IDs:               nil,
				Alias:             "planitar",
				Name:              "Planitar Inc.",
				Followers:         "4K",
				Employees:         "51-200",
				AssociatedMembers: "163",
				Verified:          true,
				Date:              mustDate("2026-03-11"),
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "",
				Verified: false,
			},
			BlindProfile: domain.BlindProfile{
				Alias: "",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
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
							Title:                "Full Stack Engineer (Golang)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4334974804/",
							Location:             "Toronto, ON",
							CloudProviders:       []domain.CloudProvider{},
							Date:                 mustDate("2026-03-10"),
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
			ShortDescription: "Planitar Inc. is the maker of iGUIDE®, a proprietary camera and software platform for documenting properties and delivering immersive 3D virtual tours and detailed property information, including accurate floor plans and reliable measurements.",
			Industries:       []domain.Industry{
				// NOP
			},
		},
		{
			ID:         0,  // system
			Type:       "", // system
			Name:       "Cenareo",
			BaseURL:    "https://www.cenareo.com/",
			CareersURL: "",
			AboutURL:   "",
			BlogURL:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                3169814,
				IDs:               nil,
				Alias:             "cenareo",
				Name:              "Cenareo",
				Followers:         "4K",
				Employees:         "51-200",
				AssociatedMembers: "62",
				Verified:          false,
				Date:              mustDate("2026-03-11"),
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "",
				Verified: false,
			},
			BlindProfile: domain.BlindProfile{
				Alias: "",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
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
			ShortDescription: "Cenareo is a digital signage platform for managing smart video content and screen fleets",
			Industries: []domain.Industry{
				domain.IndustryMarTech,
			},
		},
		{
			ID:         0,  // system
			Type:       "", // system
			Name:       "Alluvium Health",
			BaseURL:    "https://www.alluviumhealth.com/",
			CareersURL: "",
			AboutURL:   "",
			BlogURL:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                16188895,
				IDs:               nil,
				Alias:             "alluviumhealth",
				Name:              "Alluvium Health",
				Followers:         "2K",
				Employees:         "51-200",
				AssociatedMembers: "61",
				Verified:          false,
				Date:              mustDate("2026-03-14"),
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "",
				Verified: false,
			},
			BlindProfile: domain.BlindProfile{
				Alias: "",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
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
					Vacancies: []domain.Vacancy{
						{
							Title:                "Elixir Software Engineer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4384326649/",
							Location:             "United States",
							CloudProviders:       []domain.CloudProvider{},
							Date:                 mustDate("2026-03-12"),
							WithSalary:           true, // $107k/yr - $142k/yr
							Remote:               true,
						},
					},
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
			ShortDescription: "Capacity performance platform — unify data, predict demand, and optimize patient access at scale",
			Industries: []domain.Industry{
				domain.IndustryMedTech,
			},
		},
		{
			ID:         0,  // system
			Type:       "", // system
			Name:       "HG Insights",
			BaseURL:    "https://hginsights.com/",
			CareersURL: "https://hginsights.com/hg-insights-careers/",
			AboutURL:   "https://hginsights.com/about/",
			BlogURL:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                1311279,
				IDs:               nil,
				Alias:             "hg-insights",
				Name:              "HG Insights",
				Followers:         "22K",
				Employees:         "51-200",
				AssociatedMembers: "955",
				Verified:          true,
				Date:              mustDate("2026-03-14"),
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "",
				Verified: false,
			},
			BlindProfile: domain.BlindProfile{
				Alias: "",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias: "",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-HG-Insights-EI_IE1150802.11,22.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/HG-Insights-Reviews-E1150802.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/HG-Insights-Jobs-E1150802.htm",
				Jobs:        "", // Becomes empty in the new design version
				Reviews:     "", // Becomes empty in the new design version
				Salaries:    "", // Becomes empty in the new design version
				ReviewsRate: "",
				Verified:    false,
				Date:        mustDate("2026-03-14"),
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
					Vacancies: []domain.Vacancy{
						{
							Title:                "Senior Software Engineer – (Elixir/TypeScript)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4383230327/",
							Location:             "Pune District, Maharashtra, India",
							CloudProviders:       []domain.CloudProvider{},
							Date:                 mustDate("2026-03-12"),
							WithSalary:           false,
							Remote:               false,
						},
					},
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
			ShortDescription: "Go-to-market Technology Intelligence",
			Industries:       []domain.Industry{
				// NOP
			},
		},
		{
			ID:         0,  // system
			Type:       "", // system
			Name:       "TrueLayer",
			BaseURL:    "https://truelayer.com/",
			CareersURL: "https://truelayer.com/careers/",
			AboutURL:   "https://truelayer.com/about/",
			BlogURL:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                17954388,
				IDs:               nil,
				Alias:             "truelayer",
				Name:              "TrueLayer",
				Followers:         "63K",
				Employees:         "201-500",
				AssociatedMembers: "276",
				Verified:          true,
				Date:              mustDate("2026-03-14"),
			},
			GitHubProfile: domain.GitHubProfile{
				Login:     "truelayer",
				Followers: "92",
				Verified:  true,
				Date:      mustDate("2026-03-14"),
			},
			BlindProfile: domain.BlindProfile{
				Alias: "",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias: "",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-TrueLayer-EI_IE2862409.11,20.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/TrueLayer-Reviews-E2862409.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/TrueLayer-Jobs-E2862409.htm",
				Jobs:        "", // Becomes empty in the new design version
				Reviews:     "", // Becomes empty in the new design version
				Salaries:    "", // Becomes empty in the new design version
				ReviewsRate: "",
				Verified:    false,
				Date:        mustDate("2026-03-14"),
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
							URL:                  "https://www.linkedin.com/jobs/view/4312155633/",
							Location:             "Italy",
							CloudProviders:       []domain.CloudProvider{},
							Date:                 mustDate("2026-03-17", "2026-03-13"),
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
			ShortDescription: "Open banking payments network",
			Industries: []domain.Industry{
				domain.IndustryFinTech,
			},
		},
		{
			ID:         0,  // system
			Type:       "", // system
			Name:       "Toku",
			BaseURL:    "https://toku.co/",
			CareersURL: "https://toku.co/careers/",
			AboutURL:   "https://toku.co/about-us/",
			BlogURL:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                13650928,
				IDs:               []int{10983916, 13650928},
				Alias:             "toku-global",
				Name:              "Toku",
				Followers:         "8K",
				Employees:         "51-200",
				AssociatedMembers: "122",
				Verified:          true,
				Date:              mustDate("2026-03-14"),
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "",
				Verified: false,
			},
			BlindProfile: domain.BlindProfile{
				Alias: "",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
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
							Title:                "Backend Software Engineer – Go and AWS",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4382656503/",
							Location:             "Singapore",
							CloudProviders:       []domain.CloudProvider{domain.AWS},
							Date:                 mustDate("2026-03-13"),
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
			ShortDescription: "Bespoke cloud communications solutions for enhanced CX in APAC",
			Industries:       []domain.Industry{
				// NOP
			},
		},
		{
			ID:         0,  // system
			Type:       "", // system
			Name:       "Qvantel",
			BaseURL:    "https://www.qvantel.com/",
			CareersURL: "https://www.qvantel.com/careers",
			AboutURL:   "https://www.qvantel.com/company",
			BlogURL:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                65215,
				IDs:               nil,
				Alias:             "qvantel",
				Name:              "Qvantel",
				Followers:         "30K",
				Employees:         "501-1K",
				AssociatedMembers: "687",
				Verified:          true,
				Date:              mustDate("2026-03-18"),
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "",
				Verified: false,
			},
			BlindProfile: domain.BlindProfile{
				Alias: "",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
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
					Vacancies: []domain.Vacancy{
						{
							Title:                "Scala Developer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4384019554/",
							Location:             "Hyderabad, Telangana, India",
							CloudProviders:       []domain.CloudProvider{},
							Date:                 mustDate("2026-03-17"),
							WithSalary:           false,
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
				domain.IndustryTelecom,
			},
		},
		{
			ID:         0,  // system
			Type:       "", // system
			Name:       "BitMart",
			BaseURL:    "https://www.bitmart.com/",
			CareersURL: "",
			AboutURL:   "",
			BlogURL:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                11398876,
				IDs:               nil,
				Alias:             "bitmart",
				Name:              "BitMart",
				Followers:         "61K",
				Employees:         "501-1K",
				AssociatedMembers: "530",
				Verified:          true,
				Date:              mustDate("2026-03-18"),
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "",
				Verified: false,
			},
			BlindProfile: domain.BlindProfile{
				Alias: "",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
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
							Title:                "Blockchain Developer (Rust)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4380097700/",
							Location:             "APAC",
							CloudProviders:       []domain.CloudProvider{},
							Date:                 mustDate("2026-03-17"),
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
			ShortDescription: "Digital asset exchange platform",
			Industries: []domain.Industry{
				domain.IndustryCryptoCurrency,
			},
		},
	}
}
