package db

import "github.com/readytotouch/readytotouch/internal/domain"

func companies14Shard() []domain.CompanyProfile {
	return []domain.CompanyProfile{
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "EVO",
			Website: "https://evo.company/",
			Careers: "https://evo.company/vacancies/",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                1267199,
				IDs:               nil,
				Alias:             "evo.company",
				Name:              "EVO",
				Followers:         "11K",
				Employees:         "501-1K",
				AssociatedMembers: "789",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "evo-company",
				Verified: false,
			},
			BlindProfile: domain.BlindProfile{
				Alias: "",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
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
			ShortDescription: "Ukrainian developer of marketplace product",
			Ignore:           true, // Skip for now as it's Ukrainian only, expected English, will be added later
			SyncSources:      []domain.CompanySyncSource{domain.RustCompanies},
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Faraday",
			Website: "https://faraday.ai/",
			Careers: "",
			About:   "https://faraday.ai/about",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                3011900,
				IDs:               nil,
				Alias:             "faradayai",
				Name:              "Faraday",
				Followers:         "3K",
				Employees:         "11-50",
				AssociatedMembers: "48",
				Verified:          false,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "faradayio",
				Verified: false,
			},
			BlindProfile: domain.BlindProfile{
				Alias: "",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "faraday",
				Employees: "126",
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
			Ignore:           true, // Waiting for any vacancies
			SyncSources:      []domain.CompanySyncSource{domain.RustCompanies},
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Fedi",
			Website: "https://www.fedi.xyz/",
			Careers: "https://bitcoinerjobs.com/company/fedi",
			About:   "https://www.fedi.xyz/aboutus",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                92585320,
				IDs:               nil,
				Alias:             "fedi",
				Name:              "Fedi",
				Followers:         "3K",
				Employees:         "11-50",
				AssociatedMembers: "29",
				Verified:          false,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "fedibtc",
				Verified: false,
			},
			BlindProfile: domain.BlindProfile{
				Alias: "",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
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
			Ignore:           true, // Crypto, waiting for any vacancies
			SyncSources:      []domain.CompanySyncSource{domain.RustCompanies},
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "FifthTry",
			Website: "https://www.fifthtry.com/",
			Careers: "",
			About:   "https://www.fifthtry.com/about/",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                69613898,
				IDs:               nil,
				Alias:             "fifthtry",
				Name:              "FifthTry",
				Followers:         "667",
				Employees:         "2-10",
				AssociatedMembers: "5",
				Verified:          false,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "fifthtry",
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
				domain.Go:      {},
				domain.Rust:    {},
				domain.Zig:     {},
				domain.Scala:   {},
				domain.Elixir:  {},
				domain.Clojure: {},
				domain.Haskell: {},
			},
			ShortDescription: "Cloud based platform providing automated project documentation",
			Ignore:           true, // Waiting for any vacancies
			SyncSources:      []domain.CompanySyncSource{domain.RustCompanies},
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Fire and Emergency NZ",
			Website: "https://www.fireandemergency.nz/",
			Careers: "",
			About:   "https://www.fireandemergency.nz/about-us/",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                18135341,
				IDs:               nil,
				Alias:             "fireandemergencynz",
				Name:              "Fire and Emergency NZ",
				Followers:         "13K",
				Employees:         "10K+",
				AssociatedMembers: "1,417",
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
			ShortDescription: "New Zealand integrated fire and emergency services organisation",
			Ignore:           true, // Waiting for any vacancies
			SyncSources:      []domain.CompanySyncSource{domain.RustCompanies},
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Firo Solutions",
			Website: "https://firosolutions.com/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                13022579,
				IDs:               nil,
				Alias:             "firo-solutions",
				Name:              "Firo Solutions",
				Followers:         "24",
				Employees:         "2-10",
				AssociatedMembers: "2",
				Verified:          false,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "firosolutions",
				Verified: false,
			},
			BlindProfile: domain.BlindProfile{
				Alias: "",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
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
			Ignore:           true, // Should be double-checked, waiting for any vacancies
			SyncSources:      []domain.CompanySyncSource{domain.RustCompanies},
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Fly.io",
			Website: "https://fly.io/",
			Careers: "https://fly.io/jobs/",
			About:   "https://fly.io/about/",
			Blog:    "https://fly.io/blog/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                33182031,
				IDs:               nil,
				Alias:             "flydotio",
				Name:              "Fly.io",
				Followers:         "6K",
				Employees:         "51-200",
				AssociatedMembers: "61",
				Verified:          false,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "superfly",
				Verified: true,
			},
			BlindProfile: domain.BlindProfile{
				Alias: "",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "flyio",
				Employees: "20",
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
			ShortDescription: "Hosting",
			Ignore:           true, // Waiting for any vacancies
			SyncSources:      []domain.CompanySyncSource{domain.RustCompanies},
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Gremlin",
			Website: "https://www.gremlin.com/",
			Careers: "https://www.gremlin.com/team",
			About:   "https://www.gremlin.com/about",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                17919967,
				Alias:             "gremlin-inc.", // ADDED: 781
				Name:              "Gremlin",
				Followers:         "12K",
				Employees:         "51-200",
				AssociatedMembers: "64",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "gremlin",
				Verified: false,
			},
			BlindProfile: domain.BlindProfile{
				Alias: "",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "gremlin",
				Employees: "150",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Gremlin-EI_IE2056439.11,18.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Gremlin-Reviews-E2056439.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Gremlin-Jobs-E2056439.htm",
				Jobs:        "",
				Reviews:     "24",
				Salaries:    "39",
				ReviewsRate: "3.4",
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
							Title:                "Backend Software Engineer",
							ShortDescription:     "Experience in Rust & Systems Level Programming",
							SwitchingOpportunity: "",
							URL:                  "https://app.welcometothejungle.com/jobs/NU9vNkVF",
							Date:                 mustDate("2025-03-29"),
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
			ShortDescription: "Reliability Management Platform",
			SyncSources:      []domain.CompanySyncSource{domain.RustCompanies},
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Hove",
			Website: "https://hove.com/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                269166,
				Alias:             "hovedata",
				Name:              "Hove",
				Followers:         "9K",
				Employees:         "51-200",
				AssociatedMembers: "100",
				Verified:          false,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "hove-io",
				Verified: false,
			},
			BlindProfile: domain.BlindProfile{
				Alias: "",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
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
			Ignore:           true, // Waiting for any vacancies
			SyncSources:      []domain.CompanySyncSource{domain.RustCompanies},
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Huawei",
			Website: "https://www.huawei.com/",
			Careers: "https://career.huawei.com/",
			About:   "https://www.huawei.com/en/corporate-information",
			Blog:    "https://medium.com/huawei-developers",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                3014,
				Alias:             "huawei",
				Name:              "Huawei",
				Followers:         "5M",
				Employees:         "10K+",
				AssociatedMembers: "140,063",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "huawei",
				Verified: false,
			},
			BlindProfile: domain.BlindProfile{
				Alias: "",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "huawei",
				Employees: "158,910",
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
			ShortDescription: "Chinese multinational corporation and technology company",
			Ignore:           true, // Winnie
			SyncSources:      []domain.CompanySyncSource{domain.RustCompanies},
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Hugging Face",
			Website: "https://huggingface.co/",
			Careers: "https://apply.workable.com/huggingface/",
			About:   "https://huggingface.co/huggingface",
			Blog:    "https://huggingface.co/blog",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                11193683,
				Alias:             "huggingface",
				Name:              "Hugging Face",
				Followers:         "922K",
				Employees:         "51-200",
				AssociatedMembers: "513",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "huggingface",
				Verified: true,
			},
			BlindProfile: domain.BlindProfile{
				Alias: "",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "hugging-face",
				Employees: "60",
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
			ShortDescription: "American machine learning and data science platform",
			Ignore:           true, // Waiting for any vacancies
			SyncSources:      []domain.CompanySyncSource{domain.RustCompanies},
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "JUSPAY",
			Website: "https://juspay.io/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                3590874,
				IDs:               nil,
				Alias:             "juspay-technologies",
				Name:              "JUSPAY",
				Followers:         "205K",
				Employees:         "1K-5K",
				AssociatedMembers: "1,629",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "juspay",
				Verified: false,
			},
			BlindProfile: domain.BlindProfile{
				Alias: "",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "juspay",
				Employees: "450",
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
			ShortDescription: "Payments Platform",
			Ignore:           true, // Waiting for any vacancies
			SyncSources:      []domain.CompanySyncSource{domain.RustCompanies},
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "IAMBOT Inc.",
			Website: "https://iambot.ai/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                18069717,
				Alias:             "iambot",
				Name:              "IAMBOT Inc.",
				Followers:         "109",
				Employees:         "11-50",
				AssociatedMembers: "0",
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
			Ignore:           true, // Should be double-checked, waiting for any vacancies
			SyncSources:      []domain.CompanySyncSource{domain.RustCompanies},
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Lechev.space",
			Website: "https://lechev.space/",
			Careers: "",
			About:   "https://lechev.space/about-us/",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                82071520,
				IDs:               nil,
				Alias:             "lechev-space",
				Name:              "Lechev.space",
				Followers:         "143",
				Employees:         "0-1",
				AssociatedMembers: "0",
				Verified:          false,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "LechevSpace",
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
				domain.Go:      {},
				domain.Rust:    {},
				domain.Zig:     {},
				domain.Scala:   {},
				domain.Elixir:  {},
				domain.Clojure: {},
				domain.Haskell: {},
			},
			ShortDescription: "",
			Ignore:           true, // Should be double-checked, waiting for any vacancies
			SyncSources:      []domain.CompanySyncSource{domain.RustCompanies},
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Diem Association",
			Website: "https://www.diem.com/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                40670075,
				Alias:             "diemassociation",
				Name:              "Diem Association",
				Followers:         "10K",
				Employees:         "11-50",
				AssociatedMembers: "4",
				Verified:          false,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "diem",
				Verified: false,
			},
			BlindProfile: domain.BlindProfile{
				Alias: "",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
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
			Ignore:           true, // Blockchain
			SyncSources:      []domain.CompanySyncSource{domain.RustCompanies},
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Materialize",
			Website: "https://materialize.com/",
			Careers: "https://materialize.com/careers/",
			About:   "https://materialize.com/about/",
			Blog:    "https://materialize.com/blog/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                14059168,
				Alias:             "materializeinc", // ADDED: 782
				Name:              "Materialize",
				Followers:         "7K",
				Employees:         "51-200",
				AssociatedMembers: "59",
				Verified:          false,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "materializeinc",
				Verified: true,
			},
			BlindProfile: domain.BlindProfile{
				Alias: "",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "materialize",
				Employees: "30",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Materialize-EI_IE2898477.11,22.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Materialize-Reviews-E2898477.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Materialize-Jobs-E2898477.htm",
				Jobs:        "",
				Reviews:     "8",
				Salaries:    "17",
				ReviewsRate: "4.2",
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
							Title:                "Senior Full Stack Engineer — Cloud",
							ShortDescription:     "Familiarity with Python or Rust",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4177804070/",
							Date:                 mustDate("2025-03-14"),
							WithSalary:           true, // $160,000 - 215,000/year + equity
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
			ShortDescription: "Real-time data integration platform",
			SyncSources:      []domain.CompanySyncSource{domain.RustCompanies},
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "npm, Inc.",
			Website: "https://www.npmjs.com/",
			Careers: "",
			About:   "https://www.npmjs.com/about",
			Blog:    "https://github.blog/tag/npm/",
			LinkedInProfile: domain.LinkedInProfile{
				// npm, Inc. was acquired by GitHub. To see what’s new, visit GitHub.
				ID:                3585636,
				Alias:             "npm-inc-",
				Name:              "npm, Inc.",
				Followers:         "11K",
				Employees:         "11-50",
				AssociatedMembers: "15",
				Verified:          false,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "npm",
				Verified: true,
			},
			BlindProfile: domain.BlindProfile{
				Alias: "",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "npm",
				Employees: "31",
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
			ShortDescription: "Package manager for the JavaScript programming language",
			Ignore:           true, // Waiting for any vacancies
			SyncSources:      []domain.CompanySyncSource{domain.RustCompanies},
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "OVHcloud",
			Website: "https://www.ovhcloud.com/",
			Careers: "https://careers.ovhcloud.com/",
			About:   "https://www.ovhcloud.com/en/about-us/",
			Blog:    "https://blog.ovhcloud.com/category/engineering/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                1883877,
				Alias:             "ovhgroup",
				Name:              "OVHcloud",
				Followers:         "240K",
				Employees:         "1K-5K",
				AssociatedMembers: "3,056",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "ovh",
				Verified: true,
			},
			BlindProfile: domain.BlindProfile{
				Alias: "",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "ovhcloud",
				Employees: "2,766",
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
							Title:                "Software Engineer", // Rust
							ShortDescription:     "Connaissance d'un langage de programmation (C, Rust, C++)",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4185781244/",
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
			ShortDescription: "French cloud computing company which offers VPS, dedicated servers, and other web services",
			Ignore:           true, // Skip for now as it's French only, expected English, will be added later
			SyncSources:      []domain.CompanySyncSource{domain.RustCompanies},
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Parity Technologies",
			Website: "https://www.parity.io/",
			Careers: "https://www.parity.io/careers",
			About:   "https://www.parity.io/about-us",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                10211733,
				Alias:             "paritytech",
				Name:              "Parity Technologies",
				Followers:         "18K",
				Employees:         "51-200",
				AssociatedMembers: "176",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "paritytech",
				Verified: true,
			},
			BlindProfile: domain.BlindProfile{
				Alias: "",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "parity-technologies",
				Employees: "150",
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
							Title:                "Mobile Engineer",
							ShortDescription:     "Rust (for deeper blockchain integrations)",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4193420901/",
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
			ShortDescription: "",
			Ignore:           true, // Blockchain
			SyncSources:      []domain.CompanySyncSource{domain.RustCompanies},
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Polysync Technologies",
			Website: "https://www.polysync.io/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                4996073,
				Alias:             "polysync",
				Name:              "Polysync Technologies",
				Followers:         "1K",
				Employees:         "11-50",
				AssociatedMembers: "0",
				Verified:          false,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "PolySync",
				Verified: false,
			},
			BlindProfile: domain.BlindProfile{
				Alias: "",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
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
			Ignore:           true, // Should be double-checked, waiting for any vacancies
			SyncSources:      []domain.CompanySyncSource{domain.RustCompanies},
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Qumulo",
			Website: "https://qumulo.com/",
			Careers: "https://qumulo.com/careers/",
			About:   "https://qumulo.com/company/about-qumulo/",
			Blog:    "https://qumulo.com/blog/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                2533593,
				Alias:             "qumulo", // ADDED: 783
				Name:              "Qumulo",
				Followers:         "35K",
				Employees:         "201-500",
				AssociatedMembers: "502",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "qumulo",
				Verified: true,
			},
			BlindProfile: domain.BlindProfile{
				Alias: "",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "qumulo",
				Employees: "480",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Qumulo-EI_IE678884.11,17.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Qumulo-Reviews-E678884.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Qumulo-Jobs-E678884.htm",
				Jobs:        "8",
				Reviews:     "188",
				Salaries:    "291",
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
							Title:                "Software Engineer — Core Data Services",
							ShortDescription:     "Contribute to feature development and testing using languages like C and Rust",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/3895536895/",
							Date:                 mustDate("2025-03-28"),
							WithSalary:           true, // $140.000 - $190.000 per year
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
			ShortDescription: "American data storage company",
			SyncSources:      []domain.CompanySyncSource{domain.RustCompanies},
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Red Sift",
			Website: "https://redsift.com/",
			Careers: "https://careers.redsift.com/",
			About:   "https://redsift.com/about-us",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                10001774,
				Alias:             "red-sift",
				Name:              "Red Sift",
				Followers:         "10K",
				Employees:         "51-200",
				AssociatedMembers: "113",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "redsift",
				Verified: true,
			},
			BlindProfile: domain.BlindProfile{
				Alias: "",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "red-sift",
				Employees: "90",
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
			ShortDescription: "Global cybersecurity company",
			Industries: []domain.Industry{
				domain.IndustryCyberSecurity,
			},
			Ignore:      true, // Waiting for any vacancies
			SyncSources: []domain.CompanySyncSource{domain.RustCompanies},
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Routific",
			Website: "https://www.routific.com/",
			Careers: "",
			About:   "https://www.routific.com/about-us",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                3821017,
				Alias:             "routific",
				Name:              "Routific",
				Followers:         "3K",
				Employees:         "11-50",
				AssociatedMembers: "31",
				Verified:          false,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "routific",
				Verified: false,
			},
			BlindProfile: domain.BlindProfile{
				Alias: "",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
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
			ShortDescription: "Delivery management software with an intelligent route optimization feature",
			Ignore:           true, // Waiting for any vacancies
			SyncSources:      []domain.CompanySyncSource{domain.RustCompanies},
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Sandstorm.io",
			Website: "https://sandstorm.io/",
			Careers: "",
			About:   "https://sandstorm.io/about",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                7957250,
				Alias:             "sandstorm.io",
				Name:              "Sandstorm.io",
				Followers:         "97",
				Employees:         "2-10",
				AssociatedMembers: "0",
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
			Ignore:           true, // Should be double-checked
			SyncSources:      []domain.CompanySyncSource{domain.RustCompanies},
		},

		{
			// https://www.shiftleft.io/ redirects to https://qwiet.ai/
			ID:      0,  // system
			Type:    "", // system
			Name:    "Qwiet AI",
			Website: "https://qwiet.ai/",
			Careers: "https://qwiet.ai/careers/",
			About:   "https://qwiet.ai/about-us/",
			Blog:    "https://qwiet.ai/blogs/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                89946885,
				Alias:             "qwiet",
				Name:              "Qwiet AI",
				Followers:         "8K",
				Employees:         "51-200",
				AssociatedMembers: "46",
				Verified:          false,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "ShiftLeftSecurity",
				Verified: false,
			},
			BlindProfile: domain.BlindProfile{
				Alias: "",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "qwiet-ai",
				Employees: "",
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
						// Waiting for the Rust vacancy https://qwiet.ai/careers/?gh_jid=6787364002
					},
				},
				domain.Zig:     {},
				domain.Scala:   {},
				domain.Elixir:  {},
				domain.Clojure: {},
				domain.Haskell: {},
			},
			ShortDescription: "",
			Ignore:           true, // Waiting for any vacancies
			SyncSources:      []domain.CompanySyncSource{domain.RustCompanies},
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "slowtec GmbH",
			Website: "https://slowtec.de/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                11779846,
				Alias:             "slowtec",
				Name:              "slowtec GmbH",
				Followers:         "99",
				Employees:         "2-10",
				AssociatedMembers: "3",
				Verified:          false,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "slowtec",
				Verified: false,
			},
			BlindProfile: domain.BlindProfile{
				Alias: "",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
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
			Ignore:           true, // Waiting for any vacancies
			SyncSources:      []domain.CompanySyncSource{domain.RustCompanies},
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "SmartThings",
			Website: "https://www.samsung.com/us/smartthings/",
			Careers: "https://smartthings.pinpointhq.com/",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                2686346,
				Alias:             "smartthings",
				Name:              "SmartThings",
				Followers:         "19K",
				Employees:         "201-500",
				AssociatedMembers: "277",
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
				Alias:     "smartthings",
				Employees: "240",
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
			ShortDescription: "American home automation company",
			Ignore:           true, // Waiting for any vacancies
			SyncSources:      []domain.CompanySyncSource{domain.RustCompanies},
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Snips",
			Website: "https://snips.ai/",
			Careers: "",
			About:   "",
			Blog:    "https://medium.com/snips-ai/tagged/engineering",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                5024356,
				Alias:             "snips",
				Name:              "Snips",
				Followers:         "13K",
				Employees:         "51-200",
				AssociatedMembers: "55",
				Verified:          false,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "snipsco",
				Verified: false,
			},
			BlindProfile: domain.BlindProfile{
				Alias: "",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
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
			ShortDescription: "Artificial intelligence voice platform",
			Ignore:           true, // Waiting for any vacancies
			SyncSources:      []domain.CompanySyncSource{domain.RustCompanies},
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Solana Foundation",
			Website: "https://solana.com/",
			Careers: "https://jobs.solana.com/jobs",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                72057622,
				Alias:             "solana-foundation",
				Name:              "Solana Foundation",
				Followers:         "80K",
				Employees:         "51-200",
				AssociatedMembers: "356",
				Verified:          false,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "", // also Solana Foundation with solana.org information
				Verified: false,
			},
			BlindProfile: domain.BlindProfile{
				Alias: "",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
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
							Title:                "Data Engineer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://jobs.solana.com/companies/solana-foundation-2/jobs/48071693-data-engineer",
							Date:                 mustDate("2025-03-29"),
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
			SyncSources:      []domain.CompanySyncSource{domain.RustCompanies},
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Stackable",
			Website: "https://stackable.tech/",
			Careers: "https://jobs.stackable.de/",
			About:   "https://stackable.tech/en/about-us/",
			Blog:    "https://stackable.tech/en/blog/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                71663246,
				Alias:             "stackabletech",
				Name:              "Stackable",
				Followers:         "2K",
				Employees:         "11-50",
				AssociatedMembers: "19",
				Verified:          false,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "stackabletech",
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
				domain.Go:      {},
				domain.Rust:    {},
				domain.Zig:     {},
				domain.Scala:   {},
				domain.Elixir:  {},
				domain.Clojure: {},
				domain.Haskell: {},
			},
			ShortDescription: "German-made open-source data platform",
			Ignore:           true, // Waiting for any vacancies
			SyncSources:      []domain.CompanySyncSource{domain.RustCompanies},
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Star Lab, a Wind River Company",
			Website: "https://www.starlab.io/",
			Careers: "https://www.starlab.io/careers",
			About:   "https://www.starlab.io/about",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                3654418,
				Alias:             "star-lab",
				Name:              "Star Lab, a Wind River Company",
				Followers:         "3K",
				Employees:         "11-50",
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
				domain.Go:      {},
				domain.Rust:    {},
				domain.Zig:     {},
				domain.Scala:   {},
				domain.Elixir:  {},
				domain.Clojure: {},
				domain.Haskell: {},
			},
			ShortDescription: "Company who development and productization of embedded security technologies for defense and commercial applications",
			Ignore:           true, // Waiting for any vacancies
			SyncSources:      []domain.CompanySyncSource{domain.RustCompanies},
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Starry, Inc.",
			Website: "https://starry.com/",
			Careers: "https://starry.com/careers",
			About:   "https://starry.com/about",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                10435968,
				Alias:             "starryinternet",
				Name:              "Starry, Inc.",
				Followers:         "18K",
				Employees:         "201-500",
				AssociatedMembers: "259",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "starryinternet",
				Verified: true,
			},
			BlindProfile: domain.BlindProfile{
				Alias: "",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "starry",
				Employees: "520",
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
			ShortDescription: "American internet provider company",
			Ignore:           true, // Waiting for any vacancies
			SyncSources:      []domain.CompanySyncSource{domain.RustCompanies},
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "System76",
			Website: "https://system76.com/",
			Careers: "https://system76.com/careers/",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                2274140,
				Alias:             "system76",
				Name:              "System76",
				Followers:         "16K",
				Employees:         "51-200",
				AssociatedMembers: "43",
				Verified:          false,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "system76",
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
				domain.Go:      {},
				domain.Rust:    {},
				domain.Zig:     {},
				domain.Scala:   {},
				domain.Elixir:  {},
				domain.Clojure: {},
				domain.Haskell: {},
			},
			ShortDescription: "American computer manufacturer company",
			Ignore:           true, // Waiting for any vacancies
			SyncSources:      []domain.CompanySyncSource{domain.RustCompanies},
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Tailcall",
			Website: "https://tailcall.run/",
			Careers: "",
			About:   "",
			Blog:    "https://tailcall.run/blog/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                87392304,
				Alias:             "tailcall",
				Name:              "Tailcall",
				Followers:         "2K",
				Employees:         "2-10",
				AssociatedMembers: "7",
				Verified:          false,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "tailcallhq",
				Verified: false,
			},
			BlindProfile: domain.BlindProfile{
				Alias: "",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
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
			ShortDescription: "Open-source platform for building GraphQL APIs",
			Ignore:           true, // Waiting for any vacancies
			SyncSources:      []domain.CompanySyncSource{domain.RustCompanies},
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "tCell",
			Website: "https://www.tcell.io/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                4841079,
				Alias:             "tcell-io-inc",
				Name:              "tCell",
				Followers:         "380",
				Employees:         "11-50",
				AssociatedMembers: "2",
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
			Ignore:           true, // Waiting for any vacancies
			SyncSources:      []domain.CompanySyncSource{domain.RustCompanies},
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "TenX Technologies",
			Website: "https://tenx.tech/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                3241555,
				Alias:             "tenx-technologies",
				Name:              "TenX Technologies",
				Followers:         "29",
				Employees:         "2-10",
				AssociatedMembers: "5",
				Verified:          false,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "tenx-tech",
				Verified: false,
			},
			BlindProfile: domain.BlindProfile{
				Alias: "",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
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
			Ignore:           true, // Waiting for any vacancies
			SyncSources:      []domain.CompanySyncSource{domain.RustCompanies},
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "ThreatX",
			Website: "https://www.threatx.com/",
			Careers: "https://www.a10networks.com/company/careers/", // redirect from main site
			About:   "https://www.a10networks.com/company/leadership/",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                4811828,
				Alias:             "threatx",
				Name:              "ThreatX",
				Followers:         "9K",
				Employees:         "11-50",
				AssociatedMembers: "27",
				Verified:          false,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "ThreatX",
				Verified: false,
			},
			BlindProfile: domain.BlindProfile{
				Alias: "",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "threatx",
				Employees: "62",
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
			ShortDescription: "API and web application protection platform",
			Ignore:           true, // Waiting for any vacancies
			SyncSources:      []domain.CompanySyncSource{domain.RustCompanies},
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "TreeScale",
			Website: "https://treescale.com/",
			Careers: "",
			About:   "https://treescale.com/about",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                93230551,
				Alias:             "treescale",
				Name:              "TreeScale",
				Followers:         "67",
				Employees:         "2-10",
				AssociatedMembers: "1",
				Verified:          false,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "treescale",
				Verified: false,
			},
			BlindProfile: domain.BlindProfile{
				Alias: "",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
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
			Ignore:           true, // Waiting for any vacancies
			SyncSources:      []domain.CompanySyncSource{domain.RustCompanies},
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Tweede golf",
			Website: "https://tweedegolf.nl/",
			Careers: "https://tweedegolf.nl/en/vacancies",
			About:   "https://tweedegolf.nl/en/about",
			Blog:    "https://tweedegolf.nl/en/blog",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                704244,
				Alias:             "tweede-golf-software-engineering",
				Name:              "Tweede golf",
				Followers:         "2K",
				Employees:         "11-50",
				AssociatedMembers: "15",
				Verified:          false,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "tweedegolf",
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
				domain.Go:      {},
				domain.Rust:    {},
				domain.Zig:     {},
				domain.Scala:   {},
				domain.Elixir:  {},
				domain.Clojure: {},
				domain.Haskell: {},
			},
			ShortDescription: "Rust engineering consultancy company",
			Ignore:           true, // Waiting for any vacancies
			SyncSources:      []domain.CompanySyncSource{domain.RustCompanies},
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "VersionEye",
			Website: "https://www.versioneye.com/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                2883860,
				Alias:             "versioneye",
				Name:              "VersionEye",
				Followers:         "41",
				Employees:         "2-10",
				AssociatedMembers: "0",
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
			ShortDescription: "Cross-platform search engine and crowdsourcing app for open source libraries",
			Ignore:           true, // Waiting for any vacancies
			SyncSources:      []domain.CompanySyncSource{domain.RustCompanies},
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Wildfish Limited",
			Website: "https://www.wildfish.com/",
			Careers: "",
			About:   "https://www.wildfish.com/about/",
			Blog:    "https://www.wildfish.com/blog/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                1518933,
				Alias:             "wildfish",
				Name:              "Wildfish Limited",
				Followers:         "316",
				Employees:         "2-10",
				AssociatedMembers: "12",
				Verified:          false,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "wildfish",
				Verified: false,
			},
			BlindProfile: domain.BlindProfile{
				Alias: "",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
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
			ShortDescription: "Web Development and Consultancy company",
			Ignore:           true, // Waiting for any vacancies
			SyncSources:      []domain.CompanySyncSource{domain.RustCompanies},
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "YomuraFiber",
			Website: "https://yomurafiber.com/",
			Careers: "",
			About:   "https://yomurafiber.com/about-yomura-fiber/",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                13013886,
				Alias:             "yomurafiber",
				Name:              "YomuraFiber",
				Followers:         "35",
				Employees:         "201-500",
				AssociatedMembers: "1",
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
			ShortDescription: "American internet provider company",
			Ignore:           true, // Waiting for any vacancies
			SyncSources:      []domain.CompanySyncSource{domain.RustCompanies},
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Zama",
			Website: "https://www.zama.ai/",
			Careers: "https://jobs.zama.ai/",
			About:   "https://www.zama.ai/about",
			Blog:    "https://www.zama.ai/category/engineering",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                34914422,
				Alias:             "zama-ai", // ADDED: 784
				Name:              "Zama",
				Followers:         "8K",
				Employees:         "51-200",
				AssociatedMembers: "157",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "zama-ai",
				Verified: true,
			},
			BlindProfile: domain.BlindProfile{
				Alias: "",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "zama",
				Employees: "90",
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
							URL:                  "https://jobs.zama.ai/companies/zama/jobs/rust-engineer_paris",
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
			ShortDescription: "Open source cryptography company",
			SyncSources:      []domain.CompanySyncSource{domain.RustCompanies},
		},
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
				Alias: "Veeva-Systems",
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
							Date:                 mustDate("2025-04-22", "2025-04-15", "2025-04-08", "2025-03-31"),
							WithSalary:           true,
							Remote:               true,
						},
						{
							Title:                "Principal Software Engineer – Rust",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4280891809/",
							Location:             "United States",
							Date:                 mustDate("2025-12-02", "2025-10-11", "2025-08-29"),
							WithSalary:           true, // $150k/yr - $300k/yr
							Remote:               true,
						},
						{
							Title:                "Senior Software Engineer – Rust",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4280890834/",
							Location:             "United States",
							Date:                 mustDate("2025-12-11", "2025-11-11", "2025-11-01"),
							WithSalary:           true, // $110k/yr - $270k/yr
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
							Date:                 mustDate("2025-04-11", "2025-03-31"),
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
						{
							Title:                "Full-Stack Developer TypeScript / Rust",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4204006233/",
							Date:                 mustDate("2025-04-12"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Full-Stack Developer TypeScript / Rust",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4236387997/",
							Date:                 mustDate("2025-05-30"),
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
				Jobs:        "163",
				Reviews:     "1.1K",
				Salaries:    "1.9K",
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
							Date:                 mustDate("2025-04-18", "2025-03-28"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Senior Software Engineer (Java, Rust, k8s)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4207668934/",
							Date:                 mustDate("2025-05-30", "2025-05-07"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Software Engineer – Rust",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4224710915/",
							Date:                 mustDate("2025-05-10"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Senior Software Engineer – Rust",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4287142210/",
							Location:             "Manchester, England, United Kingdom",
							Date:                 mustDate("2025-09-14"),
							WithSalary:           false,
							Remote:               true,
						},
						{
							Title:                "Software Engineer – Rust",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4309596226/",
							Location:             "York, England, United Kingdom",
							Date:                 mustDate("2025-11-20", "2025-10-30", "2025-10-09"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Software Engineer – Rust",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4287143144/",
							Location:             "Manchester, England, United Kingdom",
							Date:                 mustDate("2025-10-06"),
							WithSalary:           false,
							Remote:               true,
						},
						{
							Title:                "Software Engineer – Rust",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4311988455/",
							Location:             "Manchester, England, United Kingdom",
							Date:                 mustDate("2025-11-26", "2025-11-06"),
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
			Website: "https://tubitv.com/",
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
				Alias:     "tubi",
				Employees: "420",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Tubi-EI_IE1264931.11,15.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Tubi-Reviews-E1264931.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Tubi-Jobs-E1264931.htm",
				Jobs:        "11",
				Reviews:     "109",
				Salaries:    "281",
				ReviewsRate: "3.6",
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
							Title:                "Staff Software Engineer (Rust)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4240340574/",
							Location:             "San Francisco Bay Area",
							Date:                 mustDate("2025-10-29", "2025-10-08", "2025-09-15", "2025-08-24", "2025-08-02", "2025-07-13", "2025-06-21", "2025-05-30"),
							WithSalary:           true, // $181k/yr - $259k/yr
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
							Title:                "Staff Software Engineer (Elixir)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4195366164/",
							Date:                 mustDate("2025-04-18", "2025-03-28"),
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
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Samsung-Food-EI_IE6824700.11,23.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Samsung-Food-Reviews-E6824700.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Samsung-Food-Jobs-E6824700.htm",
				Jobs:        "",
				Reviews:     "32",
				Salaries:    "7",
				ReviewsRate: "4.7",
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
							Title:                "Backend Developer (Scala)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4149148628/",
							Date:                 mustDate("2025-03-31"),
							WithSalary:           false,
							Remote:               true,
						},
						{
							Title:                "Scala Developer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4211841294/",
							Date:                 mustDate("2025-04-18"),
							WithSalary:           false,
							Remote:               true,
						},
						{
							Title:                "Backend Developer (Scala)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4223337902/",
							Date:                 mustDate("2025-05-07"),
							WithSalary:           false,
							Remote:               true,
						},
						{
							Title:                "Middle Scala Developer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4228131232/",
							Date:                 mustDate("2025-05-15"),
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
			Ignore:           true, // https://dou.ua/forums/topic/56908/
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
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Ace-and-Tate-EI_IE1736109.11,23.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Ace-and-Tate-Reviews-E1736109.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Ace-and-Tate-Jobs-E1736109.htm",
				Jobs:        "60",
				Reviews:     "139",
				Salaries:    "258",
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
						{
							Title:                "Backend Engineer (Scala)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4232131818/",
							Date:                 mustDate("2025-06-08"),
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
							Date:                 mustDate("2025-06-07", "2025-05-15", "2025-04-01"),
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
				domain.Go: {},
				domain.Rust: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Rust Engineer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4261402273/",
							Date:                 mustDate("2025-07-01"),
							WithSalary:           false,
							Remote:               true,
						},
					},
				},
				domain.Zig: {},
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
						{
							Title:                "Go Developer",
							ShortDescription:     "(Streaming Media)",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4223259085/",
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
	}
}
