package db

import "github.com/readytotouch/readytotouch/internal/domain"

func companies22Shard() []domain.CompanyProfile {
	return []domain.CompanyProfile{
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Centra",
			Website: "https://centra.com/",
			Careers: "https://centra.teamtailor.com/",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                10813841,
				IDs:               nil,
				Alias:             "centra-commerce",
				Name:              "Centra",
				Followers:         "11K",
				Employees:         "51-200",
				AssociatedMembers: "313",
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
							Title:                "Staff Engineer Go",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4310350066/",
							Location:             "Wrocław, Dolnośląskie, Poland",
							Date:                 mustDate("2025-10-08"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Go Staff Engineer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4326037685/",
							Location:             "Stockholm, Stockholm County, Sweden",
							Date:                 mustDate("2025-12-09"),
							WithSalary:           false,
							Remote:               false,
						},
					},
				},
				domain.Rust: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Zig: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Scala: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Elixir: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Clojure: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Haskell: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Erlang: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
			},
			ShortDescription: "The fashion brand commerce platform",
			Industries: []domain.Industry{
				domain.IndustryECommerce,
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Lifemote",
			Website: "https://www.lifemote.com/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                15143129,
				IDs:               nil,
				Alias:             "lifemote-networks",
				Name:              "Lifemote",
				Followers:         "5K",
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
							Title:                "Backend Engineer – Go Developer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4311254582/",
							Location:             "Türkiye",
							Date:                 mustDate("2025-10-08"),
							WithSalary:           false,
							Remote:               true,
						},
					},
				},
				domain.Rust: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Zig: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Scala: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Elixir: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Clojure: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Haskell: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Erlang: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
			},
			ShortDescription: "WiFi Analytics and Proactive Maintenance for Internet Service Providers to find and fix home network issues",
			Industries:       []domain.Industry{
				// NOP
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Chili Piper",
			Website: "https://www.chilipiper.com/",
			Careers: "https://www.chilipiper.com/careers",
			About:   "https://www.chilipiper.com/about",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                2687110,
				IDs:               nil,
				Alias:             "chilipiper",
				Name:              "Chili Piper",
				Followers:         "81K",
				Employees:         "51-200",
				AssociatedMembers: "256",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:     "Chili-Piper",
				Followers: "21",
				Verified:  true,
			},
			BlindProfile: domain.BlindProfile{
				Alias: "",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "chili-piper",
				Employees: "180",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Chili-Piper-EI_IE1835737.11,22.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Chili-Piper-Reviews-E1835737.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Chili-Piper-Jobs-E1835737.htm",
				Jobs:        "",
				Reviews:     "100",
				Salaries:    "151",
				ReviewsRate: "4.5",
				Verified:    true,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "Chili-Piper",
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 1,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Rust: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Zig: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Scala: {
					GitHubRepositoriesCount: 5,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Scala Engineer",
							SubTitle:             "",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							Location:             "",
							URL:                  "https://www.chilipiper.com/careers?ashby_jid=ab556557-83cf-467d-90fb-5119dabf146c#work",
							Date:                 mustDate("2025-10-07"),
							WithSalary:           true, // $87K – $138K
							Remote:               true,
						},
					},
				},
				domain.Elixir: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Clojure: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Haskell: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Erlang: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
			},
			ShortDescription: "Demand Conversion Platform",
			Industries:       []domain.Industry{
				// NOP
			},
			PinnedUntil: mustDate("2025-10-31"),
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "OpenVPN Inc.",
			Website: "https://openvpn.net/",
			Careers: "https://careers.openvpn.net/",
			About:   "https://openvpn.net/about/",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                15353748,
				IDs:               nil,
				Alias:             "openvpn",
				Name:              "OpenVPN Inc.",
				Followers:         "7K",
				Employees:         "51-200",
				AssociatedMembers: "183",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:     "openvpn",
				Followers: "1.5k",
				Verified:  false,
			},
			BlindProfile: domain.BlindProfile{
				Alias: "",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias: "",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-OpenVPN-Technologies-EI_IE1607550.11,31.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/OpenVPN-Technologies-Reviews-E1607550.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/OpenVPN-Technologies-Jobs-E1607550.htm",
				Jobs:        "30",
				Reviews:     "39",
				Salaries:    "20",
				ReviewsRate: "4.4",
				Verified:    false,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "",
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 4,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Senior Software Engineer (Golang)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4315588096/",
							Location:             "European Union",
							Date:                 mustDate("2025-10-22"),
							WithSalary:           false,
							Remote:               true,
						},
						{
							Title:                "Senior Golang Developer",
							ShortDescription:     "Desktop",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4321637674/",
							Location:             "Spain",
							Date:                 mustDate("2025-11-11"),
							WithSalary:           false,
							Remote:               true,
						},
						{
							Title:                "Senior Software Engineer (Golang)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4324838905/",
							Location:             "European Union",
							Date:                 mustDate("2025-11-29"),
							WithSalary:           false,
							Remote:               true,
						},
					},
				},
				domain.Rust: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Zig: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Scala: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Elixir: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Clojure: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Haskell: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Erlang: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
			},
			ShortDescription: "",
			Industries: []domain.Industry{
				domain.IndustryCyberSecurity,
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Wynd Labs",
			Website: "https://www.wyndlabs.ai/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                92612409,
				IDs:               nil,
				Alias:             "wyndlabs",
				Name:              "Wynd Labs",
				Followers:         "2K",
				Employees:         "11-50",
				AssociatedMembers: "8",
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
				domain.Rust: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Senior Rust Engineer (Backend)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4334700120/",
							Location:             "United States",
							Date:                 mustDate("2025-11-07"),
							WithSalary:           true, // $180k/yr - $300k/yr
							Remote:               true,
						},
					},
				},
				domain.Zig: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Scala: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Elixir: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Clojure: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Haskell: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Erlang: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
			},
			ShortDescription: "Wynd Labs enables large-scale access to public web data through a distributed proxy network serving research, analytics, and business intelligence",
			Industries:       []domain.Industry{
				// NOP
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Phantom",
			Website: "https://phantom.com/",
			Careers: "https://phantom.com/careers",
			About:   "https://phantom.com/about",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                72352899,
				IDs:               nil,
				Alias:             "phantomwallet",
				Name:              "Phantom",
				Followers:         "69K",
				Employees:         "51-200",
				AssociatedMembers: "311",
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
				domain.Rust: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Senior Rust Engineer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4291964082/",
							Location:             "United States",
							Date:                 mustDate("2025-11-01"),
							WithSalary:           true, // $200k/yr - $230k/yr
							Remote:               true,
						},
					},
				},
				domain.Zig: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Scala: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Elixir: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Clojure: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Haskell: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Erlang: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
			},
			ShortDescription: "Crypto wallet",
			Industries:       []domain.Industry{
				// NOP
			},
			Ignore: true, // Cryptocurrency
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "webAI",
			Website: "https://www.webai.com/",
			Careers: "https://www.webai.com/careers",
			About:   "https://www.webai.com/about",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                30756874,
				IDs:               nil,
				Alias:             "webai-official",
				Name:              "webAI",
				Followers:         "8K",
				Employees:         "51-200",
				AssociatedMembers: "118",
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
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-webAI-EI_IE7431732.11,16.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/webAI-Reviews-E7431732.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/webAI-Jobs-E7431732.htm",
				Jobs:        "",
				Reviews:     "12",
				Salaries:    "15",
				ReviewsRate: "4.2",
				Verified:    true,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "",
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Rust: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Network Software Engineer – Rust",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4333249443/",
							Location:             "Austin, Texas Metropolitan Area",
							Date:                 mustDate("2025-11-02"),
							WithSalary:           false,
							Remote:               false,
						},
					},
				},
				domain.Zig: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Scala: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Elixir: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Clojure: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Haskell: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Erlang: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
			},
			ShortDescription: "webAI is the sovereign end-to-end AI platform",
			Industries:       []domain.Industry{
				// NOP
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Perplexity",
			Website: "https://www.perplexity.ai/",
			Careers: "https://www.perplexity.ai/careers",
			About:   "https://www.perplexity.ai/about",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                88007673,
				IDs:               nil,
				Alias:             "perplexity-ai",
				Name:              "Perplexity",
				Followers:         "1M",
				Employees:         "201-500",
				AssociatedMembers: "1,749",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:     "perplexityai",
				Followers: "494",
				Verified:  true,
			},
			BlindProfile: domain.BlindProfile{
				Alias: "",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "perplexity-ai",
				Employees: "10",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Perplexity-AI-EI_IE8515634.11,24.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Perplexity-AI-Reviews-E8515634.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Perplexity-AI-Jobs-E8515634.htm",
				Jobs:        "",
				Reviews:     "41",
				Salaries:    "34",
				ReviewsRate: "4.6",
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
							Title:                "Search Golang Engineer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4322153707/",
							Location:             "Belgrade, Serbia",
							Date:                 mustDate("2025-12-05", "2025-11-13"),
							WithSalary:           false,
							Remote:               false,
						},
					},
				},
				domain.Rust: {
					GitHubRepositoriesCount: 2,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Search Rust Engineer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4320745357/",
							Location:             "Belgrade, Serbia",
							Date:                 mustDate("2025-11-05"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Search Rust Engineer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4322043806/",
							Location:             "Belgrade, Serbia",
							Date:                 mustDate("2025-11-13"),
							WithSalary:           false,
							Remote:               false,
						},
					},
				},
				domain.Zig: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Scala: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Elixir: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Clojure: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Haskell: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Erlang: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
			},
			ShortDescription: "Answer engine",
			Industries:       []domain.Industry{
				// NOP
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "LiveKit",
			Website: "https://livekit.io/",
			Careers: "https://livekit.io/careers",
			About:   "https://livekit.io/about",
			Blog:    "https://blog.livekit.io/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                74724401,
				IDs:               nil,
				Alias:             "livekitco",
				Name:              "LiveKit",
				Followers:         "6K",
				Employees:         "51-200",
				AssociatedMembers: "69",
				Verified:          false,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:     "livekit",
				Followers: "2.3k",
				Verified:  true,
			},
			BlindProfile: domain.BlindProfile{
				Alias: "",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "livekit",
				Employees: "",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-LiveKit-EI_IE10388730.11,18.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/LiveKit-Reviews-E10388730.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/LiveKit-Jobs-E10388730.htm",
				Jobs:        "17",
				Reviews:     "1",
				Salaries:    "1",
				ReviewsRate: "5.0",
				Verified:    false,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "",
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 32,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Rust: {
					GitHubRepositoriesCount: 2,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Staff Rust SDK Engineer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4334381558/",
							Location:             "United States",
							Date:                 mustDate("2025-11-06"),
							WithSalary:           true, // $120k/yr - $250k/yr
							Remote:               true,
						},
						{
							Title:                "Staff Rust SDK Engineer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4339420067/",
							Location:             "Canada",
							Date:                 mustDate("2025-11-29"),
							WithSalary:           true, // $120k/yr - $250k/yr
							Remote:               true,
						},
					},
				},
				domain.Zig: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Scala: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Elixir: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Clojure: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Haskell: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Erlang: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
			},
			ShortDescription: "LiveKit offers open source frameworks and a cloud platform for building voice, video, and physical AI agents",
			Industries:       []domain.Industry{
				// NOP
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Watergenics",
			Website: "https://www.watergenics.tech/",
			Careers: "https://www.watergenics.tech/career",
			About:   "https://www.watergenics.tech/about-us",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                20534349,
				IDs:               nil,
				Alias:             "watergenics",
				Name:              "Watergenics",
				Followers:         "3K",
				Employees:         "11-50",
				AssociatedMembers: "20",
				Verified:          false,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:     "watergenics",
				Followers: "7",
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
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Rust: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Rust Software Engineer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4315680232/",
							Location:             "Berlin, Berlin, Germany",
							Date:                 mustDate("2025-11-06"),
							WithSalary:           false,
							Remote:               false,
						},
					},
				},
				domain.Zig: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Scala: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Elixir: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Clojure: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Haskell: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Erlang: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
			},
			ShortDescription: "Watergenics delivers solutions in real-time water quality measurement and analysis through a patented sensor technology and data analytics platform",
			Industries:       []domain.Industry{
				// NOP
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Verse",
			Website: "https://verse.inc/",
			Careers: "https://verse.inc/careers/",
			About:   "https://verse.inc/about-verse/",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                89737076,
				IDs:               nil,
				Alias:             "versedotinc",
				Name:              "Verse",
				Followers:         "5K",
				Employees:         "11-50",
				AssociatedMembers: "55",
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
				domain.Rust: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Senior Software Engineer, Rust",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4334348860/",
							Location:             "San Francisco, CA",
							Date:                 mustDate("2025-11-06"),
							WithSalary:           true, // $160k/yr - $194k/yr
							Remote:               false,
						},
					},
				},
				domain.Zig: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Scala: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Elixir: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Clojure: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Haskell: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Erlang: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
			},
			ShortDescription: "Verse delivers energy cost intelligence for the enterprise",
			Industries:       []domain.Industry{
				// NOP
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "lowRISC CIC",
			Website: "https://lowrisc.org/",
			Careers: "https://lowrisc.org/jobs/",
			About:   "https://lowrisc.org/about-us/",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                25176997,
				IDs:               nil,
				Alias:             "lowrisc",
				Name:              "lowRISC CIC",
				Followers:         "6K",
				Employees:         "11-50",
				AssociatedMembers: "27",
				Verified:          false,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:     "lowrisc",
				Followers: "719",
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
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Rust: {
					GitHubRepositoriesCount: 8,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Senior/Principal Rust Engineer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4334630561/",
							Location:             "Cambridge, England, United Kingdom",
							Date:                 mustDate("2025-11-07"),
							WithSalary:           false,
							Remote:               false,
						},
					},
				},
				domain.Zig: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Scala: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Elixir: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Clojure: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Haskell: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Erlang: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
			},
			ShortDescription: "",
			Industries:       []domain.Industry{
				// NOP
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Polygon.io",
			Website: "https://massive.com/",
			Careers: "https://massive.com/careers",
			About:   "https://massive.com/about",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                11871081,
				IDs:               nil,
				Alias:             "polygon-io",
				Name:              "Polygon.io",
				Followers:         "3K",
				Employees:         "51-200",
				AssociatedMembers: "50",
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
							Title:                "Go Engineer",
							ShortDescription:     "Market Data",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4333383984/",
							Location:             "United States",
							Date:                 mustDate("2025-11-01"),
							WithSalary:           false,
							Remote:               true,
						},
					},
				},
				domain.Rust: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Zig: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Scala: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Elixir: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Clojure: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Haskell: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Erlang: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
			},
			ShortDescription: "Massive empowers participation in the financial markets by providing fair access to market data through a developer-focused platform",
			Industries:       []domain.Industry{
				// NOP
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "ReveniQ",
			Website: "https://reveniq.com/",
			Careers: "",
			About:   "https://reveniq.com/about-us/",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                93815400,
				IDs:               nil,
				Alias:             "reveniq",
				Name:              "ReveniQ",
				Followers:         "1K",
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
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Full Stack Rails/Golang Developer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4334548798/",
							Location:             "Ireland",
							Date:                 mustDate("2025-11-06"),
							WithSalary:           false,
							Remote:               true,
						},
					},
				},
				domain.Rust: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Zig: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Scala: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Elixir: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Clojure: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Haskell: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Erlang: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
			},
			ShortDescription: "Helping B2B SaaS software owners scale their recurring revenue through advanced analytics",
			Industries:       []domain.Industry{
				// NOP
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "The Weather Company",
			Website: "https://www.weather.com/",
			Careers: "https://www.weathercompany.com/careers/",
			About:   "https://www.weathercompany.com/about-us/",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                3634137,
				IDs:               nil,
				Alias:             "weather-company",
				Name:              "The Weather Company",
				Followers:         "29K",
				Employees:         "1K-5K",
				AssociatedMembers: "713",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:     "TheWeatherCompany",
				Followers: "86",
				Verified:  true,
			},
			BlindProfile: domain.BlindProfile{
				Alias: "",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias: "",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-The-Weather-Company-EI_IE9457786.11,30.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/The-Weather-Company-Reviews-E9457786.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/The-Weather-Company-Jobs-E9457786.htm",
				Jobs:        "",
				Reviews:     "27",
				Salaries:    "63",
				ReviewsRate: "3.0",
				Verified:    true,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "",
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 14,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Full-Stack Software Engineer (Golang)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4297895819/",
							Location:             "Atlanta, GA",
							Date:                 mustDate("2025-11-29", "2025-11-08"),
							WithSalary:           false,
							Remote:               false,
						},
					},
				},
				domain.Rust: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Zig: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Scala: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Elixir: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Clojure: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Haskell: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Erlang: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
			},
			ShortDescription: "The Weather Company provides insights and solutions that harness the power of weather in a scalable, privacy-forward way",
			Industries:       []domain.Industry{
				// NOP
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Bibit.id",
			Website: "https://bibit.id/",
			Careers: "https://careers.stockbit.com/jobs",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                14476124,
				IDs:               nil,
				Alias:             "bibit",
				Name:              "Bibit.id",
				Followers:         "235K",
				Employees:         "501-1K",
				AssociatedMembers: "407",
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
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Bibit-id-EI_IE3341391.11,19.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Bibit-id-Reviews-E3341391.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Bibit-id-Jobs-E3341391.htm",
				Jobs:        "",
				Reviews:     "87",
				Salaries:    "178",
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
							Title:                "Lead Backend Engineer (Golang)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4321468352/",
							Location:             "Desa Penengahan, Lampung, Indonesia",
							Date:                 mustDate("2025-11-11"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Lead Backend Engineer (Golang)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4324814184/",
							Location:             "Jakarta, Jakarta, Indonesia",
							Date:                 mustDate("2025-11-28"),
							WithSalary:           false,
							Remote:               false,
						},
					},
				},
				domain.Rust: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Zig: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Scala: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Elixir: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Clojure: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Haskell: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Erlang: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
			},
			ShortDescription: "",
			Industries: []domain.Industry{
				domain.IndustryFinTech,
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Perceptyx",
			Website: "https://www.perceptyx.com/",
			Careers: "https://go.perceptyx.com/careers",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                2150008,
				IDs:               nil,
				Alias:             "perceptyx",
				Name:              "Perceptyx",
				Followers:         "43K",
				Employees:         "201-500",
				AssociatedMembers: "327",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:     "Perceptyx",
				Followers: "2",
				Verified:  true,
			},
			BlindProfile: domain.BlindProfile{
				Alias: "",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "perceptyx",
				Employees: "360",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Perceptyx-EI_IE1397735.11,20.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Perceptyx-Reviews-E1397735.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Perceptyx-Jobs-E1397735.htm",
				Jobs:        "",
				Reviews:     "159",
				Salaries:    "275",
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
							Title:                "Software Engineer – Golang",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4335474716/",
							Location:             "Montevideo, Montevideo, Uruguay",
							Date:                 mustDate("2025-11-12"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Software Engineer – Golang",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4336253707/",
							Location:             "Montevideo, Montevideo, Uruguay",
							Date:                 mustDate("2025-11-15"),
							WithSalary:           false,
							Remote:               false,
						},
					},
				},
				domain.Rust: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Zig: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Scala: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Elixir: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Clojure: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Haskell: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Erlang: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
			},
			ShortDescription: "Platform that helps enterprises turn employee voice into business impact",
			Industries:       []domain.Industry{
				// NOP
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Threecolts",
			Website: "https://www.threecolts.com/",
			Careers: "https://www.threecolts.com/careers",
			About:   "https://www.threecolts.com/about-us",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                71701124,
				IDs:               []int{15872062, 68321800, 71701124},
				Alias:             "threecolts",
				Name:              "Threecolts",
				Followers:         "13K",
				Employees:         "501-1K",
				AssociatedMembers: "312",
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
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Threecolts-EI_IE5755608.11,21.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Threecolts-Reviews-E5755608.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Threecolts-Jobs-E5755608.htm",
				Jobs:        "8",
				Reviews:     "68",
				Salaries:    "54",
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
							Title:                "Principal Software Engineer (Golang)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4335774601/",
							Location:             "Dublin, County Dublin, Ireland",
							Date:                 mustDate("2025-11-12"),
							WithSalary:           true, // $140k/yr
							Remote:               true,
						},
					},
				},
				domain.Rust: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Zig: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Scala: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Elixir: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Clojure: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Haskell: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Erlang: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
			},
			ShortDescription: "The super suite for sellers, retail vendors, and multichannel businesses",
			Industries: []domain.Industry{
				domain.IndustryECommerce,
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Spendbase",
			Website: "https://www.spendbase.co/",
			Careers: "https://spendbase.peopleforce.io/careers",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                88402031,
				IDs:               nil,
				Alias:             "spendbase",
				Name:              "Spendbase",
				Followers:         "4K",
				Employees:         "51-200",
				AssociatedMembers: "113",
				Verified:          false,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:     "Spendbase",
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
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Spendbase-EI_IE9355753.11,20.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Spendbase-Reviews-E9355753.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Spendbase-Jobs-E9355753.htm",
				Jobs:        "",
				Reviews:     "28",
				Salaries:    "6",
				ReviewsRate: "4.5",
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
							Title:                "Senior Golang Developer",
							ShortDescription:     "B2B Fintech",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4335653359/",
							Location:             "Kyiv City, Ukraine",
							Date:                 mustDate("2025-11-12"),
							WithSalary:           false,
							Remote:               false,
						},
					},
				},
				domain.Rust: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Zig: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Scala: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Elixir: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Clojure: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Haskell: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Erlang: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
			},
			ShortDescription: "The Spendbase platform lets you visualise your software costs and real-time usage so that it becomes obvious for you how to optimise your tech stack",
			Industries: []domain.Industry{
				domain.IndustryFinTech,
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Aonic",
			Website: "https://aonic.com/",
			Careers: "https://careers.aonic.com/",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                7803337,
				IDs:               nil,
				Alias:             "aonicgroup",
				Name:              "Aonic",
				Followers:         "11K",
				Employees:         "201-500",
				AssociatedMembers: "183",
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
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Aonic-EI_IE2977044.11,16.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Aonic-Reviews-E2977044.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Aonic-Jobs-E2977044.htm",
				Jobs:        "48",
				Reviews:     "57",
				Salaries:    "65",
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
							Title:                "Full Stack Developer (Nodejs / Nestjs / Golang)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4322200964/",
							Location:             "Subang Jaya, Selangor, Malaysia",
							Date:                 mustDate("2025-11-12"),
							WithSalary:           false,
							Remote:               false,
						},
					},
				},
				domain.Rust: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Zig: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Scala: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Elixir: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Clojure: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Haskell: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Erlang: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
			},
			ShortDescription: "Drone solutions provider",
			Industries:       []domain.Industry{
				// Aerospace
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Bud",
			Website: "https://www.thisisbud.com/",
			Careers: "https://www.thisisbud.com/company/careers",
			About:   "https://www.thisisbud.com/company/about",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                10051647,
				IDs:               nil,
				Alias:             "bud-financial",
				Name:              "Bud",
				Followers:         "20K",
				Employees:         "51-200",
				AssociatedMembers: "195",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:     "thisisbud",
				Followers: "13",
				Verified:  true,
			},
			BlindProfile: domain.BlindProfile{
				Alias: "",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias: "",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Bud-EI_IE1907066.11,14.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Bud-Reviews-E1907066.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Bud-Jobs-E1907066.htm",
				Jobs:        "35",
				Reviews:     "54",
				Salaries:    "103",
				ReviewsRate: "4.1",
				Verified:    false,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "",
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 6,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Senior Software Engineer Go",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4335644563/",
							Location:             "United Kingdom",
							Date:                 mustDate("2025-11-12"),
							WithSalary:           false,
							Remote:               true,
						},
						{
							Title:                "Senior Software Engineer Go",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4343250515/",
							Location:             "United Kingdom",
							Date:                 mustDate("2025-12-10"),
							WithSalary:           false,
							Remote:               true,
						},
					},
				},
				domain.Rust: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Zig: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Scala: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Elixir: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Clojure: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Haskell: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Erlang: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
			},
			ShortDescription: "",
			Industries: []domain.Industry{
				domain.IndustryFinTech,
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Zixflow",
			Website: "https://zixflow.com/",
			Careers: "https://zixflow.com/careers",
			About:   "https://zixflow.com/about-us",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                72446145,
				IDs:               nil,
				Alias:             "zixflow",
				Name:              "Zixflow",
				Followers:         "9K",
				Employees:         "11-50",
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
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Go Backend Engineer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4322232069/",
							Location:             "India",
							Date:                 mustDate("2025-11-12"),
							WithSalary:           false,
							Remote:               true,
						},
					},
				},
				domain.Rust: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Zig: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Scala: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Elixir: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Clojure: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Haskell: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Erlang: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
			},
			ShortDescription: "CRM and Messaging Across WhatsApp, RCS, SMS and Email",
			Industries:       []domain.Industry{
				// NOP
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Jabil",
			Website: "https://www.jabil.com/",
			Careers: "https://careers.jabil.com/",
			About:   "https://jabil.com/about-us.html",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                4852,
				IDs:               []int{4852, 10512},
				Alias:             "jabil",
				Name:              "Jabil",
				Followers:         "622K",
				Employees:         "10K+",
				AssociatedMembers: "58,012",
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
				Alias:     "jabil",
				Employees: "42,080",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Jabil-EI_IE2330.11,16.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Jabil-Reviews-E2330.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Jabil-Jobs-E2330.htm",
				Jobs:        "1.1K",
				Reviews:     "5.1K",
				Salaries:    "7.9K",
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
							Title:                "Go Engineer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4322062869/",
							Location:             "Budapest, Budapest, Hungary",
							Date:                 mustDate("2025-11-13"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Senior Backend Engineer (Go)",
							ShortDescription:     "Distributed Systems",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4325638466/",
							Location:             "European Union",
							Date:                 mustDate("2025-12-05"),
							WithSalary:           false,
							Remote:               true,
						},
					},
				},
				domain.Rust: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Zig: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Scala: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Elixir: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Clojure: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Haskell: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Erlang: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
			},
			ShortDescription: "",
			Industries:       []domain.Industry{
				// NOP
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Severalnines",
			Website: "https://severalnines.com/",
			Careers: "https://severalnines.com/career/",
			About:   "https://severalnines.com/about-us/",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                1685836,
				IDs:               nil,
				Alias:             "severalnines",
				Name:              "Severalnines",
				Followers:         "6K",
				Employees:         "11-50",
				AssociatedMembers: "37",
				Verified:          false,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:     "severalnines",
				Followers: "22",
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
				domain.Go: {
					GitHubRepositoriesCount: 10,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Senior Go Developer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4335435441/",
							Location:             "Zagreb, Zagreb, Croatia",
							Date:                 mustDate("2025-11-12"),
							WithSalary:           false,
							Remote:               true,
						},
					},
				},
				domain.Rust: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Zig: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Scala: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Elixir: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Clojure: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Haskell: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Erlang: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
			},
			ShortDescription: "DBaaS",
			Industries: []domain.Industry{
				domain.IndustryDevOps,
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Plannatech",
			Website: "https://www.plannatech.com/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                65515903,
				IDs:               nil,
				Alias:             "plannatech",
				Name:              "Plannatech",
				Followers:         "7K",
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
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Senior Golang Developer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4335797154/",
							Location:             "San Jose, Costa Rica",
							Date:                 mustDate("2025-11-13"),
							WithSalary:           false,
							Remote:               false,
						},
					},
				},
				domain.Rust: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Zig: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Scala: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Elixir: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Clojure: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Haskell: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Erlang: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
			},
			ShortDescription: "B2B platform",
			Industries:       []domain.Industry{
				// NOP
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Profound",
			Website: "https://www.tryprofound.com/",
			Careers: "https://www.tryprofound.com/careers",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                104065246,
				IDs:               nil,
				Alias:             "tryprofound",
				Name:              "Profound",
				Followers:         "20K",
				Employees:         "51-200",
				AssociatedMembers: "110",
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
					Vacancies:               []domain.Vacancy{},
				},
				domain.Rust: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Senior Software Engineer, Rust",
							ShortDescription:     "Workflows Team",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4321715155/",
							Location:             "New York City Metropolitan Area",
							Date:                 mustDate("2025-11-10"),
							WithSalary:           true, // $140k/yr - $230k/yr
							Remote:               false,
						},
						{
							Title:                "Senior System Engineer, Rust",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4323603760/",
							Location:             "New York City Metropolitan Area",
							Date:                 mustDate("2025-11-20"),
							WithSalary:           true, // $140k/yr - $230k/yr
							Remote:               false,
						},
						{
							Title:                "Senior System Engineer, Rust",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4324691416/",
							Location:             "New York City Metropolitan Area",
							Date:                 mustDate("2025-11-27"),
							WithSalary:           true, // $140k/yr - $230k/yr
							Remote:               false,
						},
					},
				},
				domain.Zig: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Scala: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Elixir: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Clojure: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Haskell: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Erlang: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
			},
			ShortDescription: "Marketing intelligence platform",
			Industries: []domain.Industry{
				domain.IndustryMarTech,
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Shopify",
			Website: "https://www.shopify.com/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                784652,
				IDs:               []int{784652, 10462386},
				Alias:             "shopify",
				Name:              "Shopify",
				Followers:         "1M",
				Employees:         "10K+",
				AssociatedMembers: "24,513",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:     "shopify",
				Followers: "7.3k",
				Verified:  true,
			},
			BlindProfile: domain.BlindProfile{
				Alias:       "Shopify",
				Employees:   "",
				Salary:      "",
				Reviews:     "1,224",
				ReviewsRate: "3.4",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "shopify",
				Employees: "7,000",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Shopify-EI_IE675933.11,18.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Shopify-Reviews-E675933.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Shopify-Jobs-E675933.htm",
				Jobs:        "53",
				Reviews:     "4K",
				Salaries:    "9.2K",
				ReviewsRate: "3.3",
				Verified:    false,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "Shopify",
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 123,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Rust: {
					GitHubRepositoriesCount: 18,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Software Engineer – Rust",
							ShortDescription:     "Monorepo Systems",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4315099000/",
							Location:             "NAMER",
							Date:                 mustDate("2025-11-29", "2025-11-12"),
							WithSalary:           false,
							Remote:               true,
						},
					},
				},
				domain.Zig: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Scala: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Elixir: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Clojure: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Haskell: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Erlang: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
			},
			ShortDescription: "",
			Industries: []domain.Industry{
				domain.IndustryECommerce,
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Exein",
			Website: "https://www.exein.io/",
			Careers: "https://www.exein.io/careers",
			About:   "https://www.exein.io/about",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                18793584,
				IDs:               nil,
				Alias:             "exein",
				Name:              "Exein",
				Followers:         "18K",
				Employees:         "11-50",
				AssociatedMembers: "42",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:     "exein-io",
				Followers: "116",
				Verified:  true,
			},
			BlindProfile: domain.BlindProfile{
				Alias: "",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias: "",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Exein-EI_IE2895315.11,16.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Exein-Reviews-E2895315.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Exein-Jobs-E2895315.htm",
				Jobs:        "",
				Reviews:     "22",
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
					Vacancies:               []domain.Vacancy{},
				},
				domain.Rust: {
					GitHubRepositoriesCount: 4,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Senior Software Engineer (Rust)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4321874435/",
							Location:             "Rome, Latium, Italy",
							Date:                 mustDate("2025-12-10", "2025-11-12"),
							WithSalary:           false,
							Remote:               false,
						},
					},
				},
				domain.Zig: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Scala: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Elixir: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Clojure: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Haskell: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Erlang: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
			},
			ShortDescription: "Embedded cybersecurity",
			Industries: []domain.Industry{
				domain.IndustryCyberSecurity,
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Cybereason",
			Website: "https://www.cybereason.com/",
			Careers: "https://www.cybereason.com/company",
			About:   "https://www.cybereason.com/company/careers",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                3173675,
				IDs:               nil,
				Alias:             "cybereason",
				Name:              "Cybereason",
				Followers:         "101K",
				Employees:         "1K-5K",
				AssociatedMembers: "556",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:    "",
				Verified: false,
			},
			BlindProfile: domain.BlindProfile{
				Alias:       "Cybereason",
				Employees:   "",
				Salary:      "",
				Reviews:     "11",
				ReviewsRate: "3",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "cybereason",
				Employees: "1,500",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Cybereason-EI_IE1077181.11,21.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Cybereason-Reviews-E1077181.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Cybereason-Jobs-E1077181.htm",
				Jobs:        "1",
				Reviews:     "486",
				Salaries:    "592",
				ReviewsRate: "2.7",
				Verified:    false,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "Cybereason",
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Rust: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Rust Engineer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4321915810/",
							Location:             "Tokyo, Tokyo, Japan",
							Date:                 mustDate("2025-12-10", "2025-12-01", "2025-11-24", "2025-11-20", "2025-11-12"),
							WithSalary:           false,
							Remote:               false,
						},
					},
				},
				domain.Zig: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Scala: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Elixir: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Clojure: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Haskell: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Erlang: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
			},
			ShortDescription: "",
			Industries: []domain.Industry{
				domain.IndustryCyberSecurity,
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Axonius",
			Website: "https://www.axonius.com/",
			Careers: "https://www.axonius.com/company/careers",
			About:   "https://www.axonius.com/company",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                18226665,
				IDs:               nil,
				Alias:             "axonius",
				Name:              "Axonius",
				Followers:         "49K",
				Employees:         "501-1K",
				AssociatedMembers: "830",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:     "Axonius",
				Followers: "92",
				Verified:  true,
			},
			BlindProfile: domain.BlindProfile{
				Alias: "",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "axonius",
				Employees: "430",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Axonius-EI_IE2193939.11,18.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Axonius-Reviews-E2193939.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Axonius-Jobs-E2193939.htm",
				Jobs:        "11",
				Reviews:     "228",
				Salaries:    "335",
				ReviewsRate: "4.1",
				Verified:    true,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "",
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 3,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Rust: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Senior Software Engineer (Rust)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4315928124/",
							Location:             "Tel Aviv-Yafo, Tel Aviv District, Israel",
							Date:                 mustDate("2025-11-12"),
							WithSalary:           false,
							Remote:               false,
						},
					},
				},
				domain.Zig: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Scala: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Elixir: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Clojure: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Haskell: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Erlang: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
			},
			ShortDescription: "",
			Industries:       []domain.Industry{
				// NOP
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "HubSpot",
			Website: "https://www.hubspot.com/",
			Careers: "https://www.hubspot.com/careers",
			About:   "https://www.hubspot.com/our-story",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                68529,
				IDs:               nil,
				Alias:             "hubspot",
				Name:              "HubSpot",
				Followers:         "2M",
				Employees:         "5K-10K",
				AssociatedMembers: "10,938",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:     "hubspot",
				Followers: "666",
				Verified:  false,
			},
			BlindProfile: domain.BlindProfile{
				Alias:       "HubSpot",
				Employees:   "",
				Salary:      "",
				Reviews:     "664",
				ReviewsRate: "3.7",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "hubspot",
				Employees: "3,500",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-HubSpot-EI_IE227605.11,18.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/HubSpot-Reviews-E227605.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/HubSpot-Jobs-E227605.htm",
				Jobs:        "106",
				Reviews:     "4.4K",
				Salaries:    "9.5K",
				ReviewsRate: "3.5",
				Verified:    true,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "Hubspot",
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 20,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Rust: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Zig: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Scala: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Elixir: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Clojure: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Haskell: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Senior Software Engineer, Data Integration (Haskell)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4308484530/",
							Location:             "Belgium",
							Date:                 mustDate("2025-11-12"),
							WithSalary:           false,
							Remote:               true,
						},
						{
							Title:                "Senior Software Engineer, Data Integration (Haskell)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4306246217/",
							Location:             "Ghent, Flemish Region, Belgium",
							Date:                 mustDate("2025-11-29"),
							WithSalary:           false,
							Remote:               false,
						},
					},
				},
				domain.Erlang: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
			},
			ShortDescription: "CRM platform that provides software and support to help businesses grow better",
			Industries:       []domain.Industry{
				// NOP
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "loveholidays",
			Website: "https://www.loveholidays.com/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                2519454,
				IDs:               nil,
				Alias:             "loveholidays",
				Name:              "loveholidays",
				Followers:         "20K",
				Employees:         "201-500",
				AssociatedMembers: "451",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:     "loveholidays",
				Followers: "42",
				Verified:  true,
			},
			BlindProfile: domain.BlindProfile{
				Alias: "",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias: "",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-loveholidays-EI_IE1138047.11,23.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/loveholidays-Reviews-E1138047.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/loveholidays-Jobs-E1138047.htm",
				Jobs:        "",
				Reviews:     "185",
				Salaries:    "280",
				ReviewsRate: "3.7",
				Verified:    true,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "",
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 17,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Senior Software Engineer",
							ShortDescription:     "Expertise in some of our main programming languages - TypeScript, Java, Golang, Rust, Python",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4280933839/",
							Location:             "London, England, United Kingdom",
							Date:                 mustDate("2025-11-07"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Principal Software Engineer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://app.welcometothejungle.com/jobs/K4rT3xXa",
							Date:                 mustDate("2025-11-14"),
							WithSalary:           false,
							Remote:               false,
						},
					},
				},
				domain.Rust: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Senior Software Engineer",
							ShortDescription:     "Expertise in some of our main programming languages - TypeScript, Java, Golang, Rust, Python",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4280933839/",
							Location:             "London, England, United Kingdom",
							Date:                 mustDate("2025-11-07"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Principal Software Engineer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://app.welcometothejungle.com/jobs/K4rT3xXa",
							Date:                 mustDate("2025-11-14"),
							WithSalary:           false,
							Remote:               false,
						},
					},
				},
				domain.Zig: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Scala: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Elixir: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Clojure: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Haskell: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Erlang: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
			},
			ShortDescription: "",
			Industries: []domain.Industry{
				domain.IndustryTravelTech,
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Nominal",
			Website: "https://nominal.io/",
			Careers: "https://nominal.io/careers",
			About:   "https://nominal.io/about-us",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                92924343,
				IDs:               nil,
				Alias:             "nominal-inc",
				Name:              "Nominal",
				Followers:         "11K",
				Employees:         "51-200",
				AssociatedMembers: "124",
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
				Alias:     "nominal",
				Employees: "",
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
				domain.Rust: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Software Engineer",
							ShortDescription:     "Rust, Bevy, WebGPU",
							SwitchingOpportunity: "",
							URL:                  "https://app.welcometothejungle.com/jobs/GBw_kEIb",
							Date:                 mustDate("2025-11-14"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Software Engineer",
							ShortDescription:     "Rust, Bevy, WebGPU",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4348528670/",
							Location:             "New York, NY",
							Date:                 mustDate("2025-12-06"),
							WithSalary:           true, // $130k/yr - $230k/yr
							Remote:               false,
						},
					},
				},
				domain.Zig: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Scala: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Elixir: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Clojure: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Haskell: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Erlang: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
			},
			ShortDescription: "Nominal builds the essential software stack that enables hardware teams to test and iterate as rapidly as software teams",
			Industries:       []domain.Industry{
				// NOP
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Upsun (formerly Platform.sh)",
			Website: "https://upsun.com/",
			Careers: "https://upsun.com/jobs/",
			About:   "https://upsun.com/about-us/",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                10304431,
				IDs:               nil,
				Alias:             "upsundotcom",
				Name:              "Upsun (formerly Platform.sh)",
				Followers:         "19K",
				Employees:         "201-500",
				AssociatedMembers: "287",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:     "upsun",
				Followers: "14",
				Verified:  true,
			},
			BlindProfile: domain.BlindProfile{
				Alias: "",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias: "",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Upsun-EI_IE1963948.11,16.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Upsun-Reviews-E1963948.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Upsun-Jobs-E1963948.htm",
				Jobs:        "8",
				Reviews:     "122",
				Salaries:    "238",
				ReviewsRate: "3.5",
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
							ShortDescription:     "Hands-on experience building RESTful applications in Go",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4315845952/",
							Location:             "Spain",
							Date:                 mustDate("2025-11-13"),
							WithSalary:           false,
							Remote:               true,
						},
						{
							Title:                "Software Engineer",
							ShortDescription:     "3+ years of development experience in Go",
							SwitchingOpportunity: "",
							URL:                  "https://app.welcometothejungle.com/jobs/a89BBmyB",
							Date:                 mustDate("2025-11-14"),
							WithSalary:           false,
							Remote:               false,
						},
					},
				},
				domain.Rust: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Zig: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Scala: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Elixir: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Clojure: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Haskell: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Erlang: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
			},
			ShortDescription: "Cloud Platform",
			Industries: []domain.Industry{
				domain.IndustryCloudComputing,
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "MotherDuck",
			Website: "https://motherduck.com/",
			Careers: "https://motherduck.com/careers/",
			About:   "https://motherduck.com/about-us/",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                81922043,
				IDs:               nil,
				Alias:             "motherduck",
				Name:              "MotherDuck",
				Followers:         "26K",
				Employees:         "11-50",
				AssociatedMembers: "104",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:     "motherduckdb",
				Followers: "140",
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
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Backend Software Engineer",
							ShortDescription:     "Hands-on experience with distributed systems, preferably using Go, C++, or similar languages",
							SwitchingOpportunity: "",
							URL:                  "https://app.welcometothejungle.com/jobs/ZE2xDHam",
							Date:                 mustDate("2025-11-14"),
							WithSalary:           false,
							Remote:               false,
						},
					},
				},
				domain.Rust: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Zig: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Scala: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Elixir: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Clojure: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Haskell: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Erlang: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
			},
			ShortDescription: "Data warehouse scaling in the cloud to make big data feel small using the speed, efficiency and ergonomics of DuckDB",
			Industries: []domain.Industry{
				domain.IndustryDevOps,
				domain.IndustryCloudComputing,
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "honeycomb.io",
			Website: "https://www.honeycomb.io/",
			Careers: "https://www.honeycomb.io/careers",
			About:   "https://www.honeycomb.io/about",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                12957659,
				IDs:               nil,
				Alias:             "honeycomb.io",
				Name:              "honeycomb.io",
				Followers:         "24K",
				Employees:         "201-500",
				AssociatedMembers: "299",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:     "honeycombio",
				Followers: "222",
				Verified:  true,
			},
			BlindProfile: domain.BlindProfile{
				Alias:       "Honeycombio",
				Employees:   "",
				Salary:      "",
				Reviews:     "7",
				ReviewsRate: "3.7",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "honeycomb",
				Employees: "150",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Honeycomb-EI_IE3970329.11,20.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Honeycomb-Reviews-E3970329.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Honeycomb-Jobs-E3970329.htm",
				Jobs:        "23",
				Reviews:     "36",
				Salaries:    "90",
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
							Title:                "Senior Software Engineer II",
							ShortDescription:     "Go",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4315779307/",
							Location:             "Canada",
							Date:                 mustDate("2025-10-31"),
							WithSalary:           false,
							Remote:               true,
						},
						{
							Title:                "Senior Software Engineer",
							ShortDescription:     "Go",
							SwitchingOpportunity: "",
							URL:                  "https://app.welcometothejungle.com/jobs/pWV5kzG8",
							Location:             "Canada",
							Date:                 mustDate("2025-11-14"),
							WithSalary:           false,
							Remote:               false,
						},
					},
				},
				domain.Rust: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Zig: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Scala: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Elixir: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Clojure: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Haskell: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Erlang: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
			},
			ShortDescription: "Observability platform",
			Industries: []domain.Industry{
				domain.IndustryDevOps,
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Meilisearch",
			Website: "https://www.meilisearch.com/",
			Careers: "https://www.meilisearch.com/careers",
			About:   "https://www.meilisearch.com/blog",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                13016868,
				IDs:               nil,
				Alias:             "meilisearch",
				Name:              "Meilisearch",
				Followers:         "4K",
				Employees:         "11-50",
				AssociatedMembers: "27",
				Verified:          false,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:     "meilisearch",
				Followers: "1.2k",
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
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Rust: {
					GitHubRepositoriesCount: 46,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Senior Software Engineer, Meilisearch",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://app.welcometothejungle.com/jobs/O2Ko2bd8",
							Location:             "Paris, Île-de-France, France",
							Date:                 mustDate("2025-11-14"),
							WithSalary:           false,
							Remote:               false,
						},
					},
				},
				domain.Zig: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Scala: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Elixir: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Clojure: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Haskell: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Erlang: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
			},
			ShortDescription: "Search solution",
			Industries: []domain.Industry{
				domain.IndustryDevOps,
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Cloudera",
			Website: "https://www.cloudera.com/",
			Careers: "https://www.cloudera.com/careers",
			About:   "https://www.cloudera.com/about",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                229433,
				IDs:               nil,
				Alias:             "cloudera",
				Name:              "Cloudera",
				Followers:         "286K",
				Employees:         "1K-5K",
				AssociatedMembers: "3,293",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:     "CLOUDERA",
				Followers: "309",
				Verified:  false,
			},
			BlindProfile: domain.BlindProfile{
				Alias:       "Cloudera",
				Employees:   "",
				Salary:      "",
				Reviews:     "455",
				ReviewsRate: "3.4",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "cloudera",
				Employees: "3,500",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Cloudera-EI_IE360671.11,19.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Cloudera-Reviews-E360671.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Cloudera-Jobs-E360671.htm",
				Jobs:        "80",
				Reviews:     "1.4K",
				Salaries:    "2.3K",
				ReviewsRate: "4.0",
				Verified:    true,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "Cloudera",
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Staff Software Engineer (Go)",
							ShortDescription:     "Infrastructure Tooling",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4316427285/",
							Location:             "Prague, Czechia",
							Date:                 mustDate("2025-12-02"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Senior Software Engineer (Go)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4326565099/",
							Location:             "Budapest Metropolitan Area",
							Date:                 mustDate("2025-12-10"),
							WithSalary:           false,
							Remote:               true,
						},
					},
				},
				domain.Rust: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Zig: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Scala: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Staff Software Engineer (Java/Scala)",
							ShortDescription:     "Iceberg",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4317466938/",
							Location:             "Budapest Metropolitan Area",
							Date:                 mustDate("2025-12-06", "2025-11-15"),
							WithSalary:           false,
							Remote:               true,
						},
					},
				},
				domain.Elixir: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Clojure: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Haskell: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Erlang: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
			},
			ShortDescription: "Data and AI platform",
			Industries: []domain.Industry{
				domain.IndustryDevOps,
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Twilio",
			Website: "https://www.twilio.com/",
			Careers: "https://www.twilio.com/company/jobs",
			About:   "https://www.twilio.com/company",
			Blog:    "https://www.twilio.com/blog/developers",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                400528,
				IDs:               []int{400528, 2425698},
				Alias:             "twilio-inc-",
				Name:              "Twilio",
				Followers:         "427K",
				Employees:         "5K-10K",
				AssociatedMembers: "6,590",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:     "twilio",
				Followers: "862",
				Verified:  true,
			},
			BlindProfile: domain.BlindProfile{
				Alias:       "Twilio",
				Employees:   "",
				Salary:      "",
				Reviews:     "1,228",
				ReviewsRate: "3.5",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "twilio",
				Employees: "7,140",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Twilio-EI_IE410790.11,17.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Twilio-Reviews-E410790.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Twilio-Jobs-E410790.htm",
				Jobs:        "57",
				Reviews:     "2.2K",
				Salaries:    "5.8K",
				ReviewsRate: "3.9",
				Verified:    false,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "Twilio",
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Rust: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Software Engineer (Java and Rust)",
							ShortDescription:     "Video Signalling",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4336250122/",
							Location:             "Ireland",
							Date:                 mustDate("2025-12-06", "2025-11-15"),
							WithSalary:           false,
							Remote:               true,
						},
					},
				},
				domain.Zig: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Scala: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Elixir: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Clojure: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Haskell: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Erlang: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
			},
			ShortDescription: "Customer Engagement Platform",
			YahooFinanceURL:  "https://finance.yahoo.com/quote/TWLO/",
			GoogleFinanceURL: "https://www.google.com/finance/quote/TWLO:NYSE",
			YCombinatorURL:   "",
			Industries: []domain.Industry{
				domain.IndustryDevOps,
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Albatross AI",
			Website: "https://usealbatross.ai/",
			Careers: "https://apply.workable.com/usealbatross/",
			About:   "https://usealbatross.ai/about-us",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                104747602,
				IDs:               nil,
				Alias:             "albatross-ai",
				Name:              "Albatross AI",
				Followers:         "2K",
				Employees:         "11-50",
				AssociatedMembers: "20",
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
					Vacancies:               []domain.Vacancy{},
				},
				domain.Rust: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Rust Software Engineer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4336878749/",
							Location:             "France",
							Date:                 mustDate("2025-11-19"),
							WithSalary:           false,
							Remote:               true,
						},
					},
				},
				domain.Zig: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Scala: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Elixir: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Clojure: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Haskell: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Erlang: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
			},
			ShortDescription: "AI platform for real-time discovery",
			Industries:       []domain.Industry{
				// NOP
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "DISQO",
			Website: "https://www.disqo.com/",
			Careers: "https://www.disqo.com/careers/",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                12178400,
				IDs:               nil,
				Alias:             "disqo",
				Name:              "DISQO",
				Followers:         "35K",
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
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-DISQO-EI_IE2266003.11,16.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/DISQO-Reviews-E2266003.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/DISQO-Jobs-E2266003.htm",
				Jobs:        "6",
				Reviews:     "136",
				Salaries:    "241",
				ReviewsRate: "4.1",
				Verified:    false,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "",
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Rust: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Senior Staff Software Engineer – Rust",
							ShortDescription:     "Analytics Platform",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4324341674/",
							Location:             "Los Angeles, CA",
							Date:                 mustDate("2025-12-09", "2025-11-25"),
							WithSalary:           true, // $200k/yr - $240k/yr
							Remote:               false,
						},
					},
				},
				domain.Zig: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Scala: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Elixir: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Clojure: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Haskell: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Erlang: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
			},
			ShortDescription: "Platform for measuring how advertising works",
			Industries: []domain.Industry{
				domain.IndustryAdTech,
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "SpotOn",
			Website: "https://www.spoton.com/",
			Careers: "https://www.spoton.com/careers/",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                1206743,
				IDs:               nil,
				Alias:             "spoton",
				Name:              "SpotOn",
				Followers:         "151K",
				Employees:         "1K-5K",
				AssociatedMembers: "2,603",
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
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-SpotOn-EI_IE504721.11,17.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/SpotOn-Reviews-E504721.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/SpotOn-Jobs-E504721.htm",
				Jobs:        "108",
				Reviews:     "1K",
				Salaries:    "1.4K",
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
							Title:                "Senior Software Engineer II (Golang)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4324396276/",
							Location:             "Chicago, IL",
							Date:                 mustDate("2025-11-26"),
							WithSalary:           true, // $145k/yr - $190k/yr
							Remote:               false,
						},
						{
							Title:                "Senior Software Engineer II (Golang)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4326711280/",
							Location:             "Detroit, MI",
							Date:                 mustDate("2025-12-10"),
							WithSalary:           true, // $145k/yr - $190k/yr
							Remote:               false,
						},
					},
				},
				domain.Rust: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Zig: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Scala: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Elixir: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Clojure: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Haskell: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Erlang: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
			},
			ShortDescription: "SpotOn helps restaurants grow profits, streamline operations, and own their customer experience",
			Industries:       []domain.Industry{
				// NOP
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Nebul",
			Website: "https://nebul.com/",
			Careers: "",
			About:   "https://nebul.com/about/",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                98627236,
				IDs:               nil,
				Alias:             "nebul",
				Name:              "Nebul",
				Followers:         "5K",
				Employees:         "11-50",
				AssociatedMembers: "38",
				Verified:          false,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:     "NEBUL-AI",
				Followers: "4",
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
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Backend Engineer (Golang + Python)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4336898660/",
							Location:             "Leiden, South Holland, Netherlands",
							Date:                 mustDate("2025-11-26"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Backend Engineer (Golang + Python)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4342869486/",
							Location:             "Leiden, South Holland, Netherlands",
							Date:                 mustDate("2025-12-10"),
							WithSalary:           false,
							Remote:               false,
						},
					},
				},
				domain.Rust: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Zig: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Scala: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Elixir: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Clojure: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Haskell: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Erlang: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
			},
			ShortDescription: "European Sovereign Cloud",
			Industries: []domain.Industry{
				domain.IndustryCloudComputing,
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Botify",
			Website: "https://www.botify.com/",
			Careers: "https://www.botify.com/careers",
			About:   "https://www.botify.com/about",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                2793570,
				IDs:               nil,
				Alias:             "botify",
				Name:              "Botify",
				Followers:         "32K",
				Employees:         "201-500",
				AssociatedMembers: "274",
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
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Botify-EI_IE1466856.11,17.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Botify-Reviews-E1466856.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Botify-Jobs-E1466856.htm",
				Jobs:        "14",
				Reviews:     "118",
				Salaries:    "235",
				ReviewsRate: "3.5",
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
							URL:                  "https://www.linkedin.com/jobs/view/4323860548/",
							Location:             "Paris, Île-de-France, France",
							Date:                 mustDate("2025-11-22"),
							WithSalary:           false,
							Remote:               false,
						},
					},
				},
				domain.Rust: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Zig: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Scala: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Elixir: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Clojure: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Haskell: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Erlang: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
			},
			ShortDescription: "Brands maximize visibility wherever consumers, bots, and AI agents search, protecting and capturing revenue across all search platforms",
			Industries:       []domain.Industry{
				// NOP
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Vara",
			Website: "https://www.vara.ai/",
			Careers: "https://www.vara.ai/careers",
			About:   "https://www.vara.ai/about-us",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                30122341,
				IDs:               nil,
				Alias:             "vara-ai",
				Name:              "Vara",
				Followers:         "11K",
				Employees:         "11-50",
				AssociatedMembers: "54",
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
					Vacancies:               []domain.Vacancy{},
				},
				domain.Rust: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Zig: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Scala: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Elixir: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Clojure: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Clojure(Script) Engineer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4338701346/",
							Location:             "Berlin, Germany",
							Date:                 mustDate("2025-11-25"),
							WithSalary:           false,
							Remote:               false,
						},
					},
				},
				domain.Haskell: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Erlang: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
			},
			ShortDescription: "Making data-driven breast cancer screening accessible",
			Industries: []domain.Industry{
				domain.IndustryMedTech,
				domain.IndustryHealthTech,
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "LoopMe",
			Website: "https://loopme.com/",
			Careers: "https://loopme.com/careers/",
			About:   "https://loopme.com/company/",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                2746023,
				IDs:               nil,
				Alias:             "loopme",
				Name:              "LoopMe",
				Followers:         "69K",
				Employees:         "201-500",
				AssociatedMembers: "461",
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
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-LoopMe-EI_IE911164.11,17.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/LoopMe-Reviews-E911164.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/LoopMe-Jobs-E911164.htm",
				Jobs:        "23",
				Reviews:     "133",
				Salaries:    "186",
				ReviewsRate: "3.5",
				Verified:    true,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "",
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Rust: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Zig: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Scala: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Senior Software Engineer (Scala/Java)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4320575418/",
							Location:             "Cracow, Małopolskie, Poland",
							Date:                 mustDate("2025-11-27"),
							WithSalary:           false,
							Remote:               false,
						},
					},
				},
				domain.Elixir: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Clojure: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Haskell: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Erlang: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
			},
			ShortDescription: "",
			Industries: []domain.Industry{
				domain.IndustryAdTech,
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "smartclip",
			Website: "https://smartclip.tv/",
			Careers: "https://smartclip.tv/jobs/",
			About:   "https://smartclip.tv/about/",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                386206,
				IDs:               nil,
				Alias:             "smartclip",
				Name:              "smartclip",
				Followers:         "11K",
				Employees:         "51-200",
				AssociatedMembers: "135",
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
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-smartclip-EI_IE765378.11,20.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/smartclip-Reviews-E765378.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/smartclip-Jobs-E765378.htm",
				Jobs:        "9",
				Reviews:     "17",
				Salaries:    "30",
				ReviewsRate: "4.3",
				Verified:    true,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "",
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Rust: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Zig: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Scala: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Software Engineer – Scala",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4339118450/",
							Location:             "Germany",
							Date:                 mustDate("2025-11-27"),
							WithSalary:           false,
							Remote:               true,
						},
					},
				},
				domain.Elixir: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Clojure: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Haskell: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Erlang: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
			},
			ShortDescription: "",
			Industries: []domain.Industry{
				domain.IndustryAdTech,
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Jump - Advisor AI",
			Website: "https://jump.ai/",
			Careers: "https://jump.ai/careers",
			About:   "https://jump.ai/about",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                99834287,
				IDs:               nil,
				Alias:             "jump-app-ai",
				Name:              "Jump - Advisor AI",
				Followers:         "9K",
				Employees:         "51-200",
				AssociatedMembers: "149",
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
					Vacancies:               []domain.Vacancy{},
				},
				domain.Rust: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Zig: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Scala: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Elixir: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Senior Elixir Engineer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4338746366/",
							Location:             "United States",
							Date:                 mustDate("2025-11-25"),
							WithSalary:           true, // $150k/yr - $260k/yr
							Remote:               true,
						},
					},
				},
				domain.Clojure: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Haskell: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Erlang: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
			},
			ShortDescription: "Jump uses AI to help financial managers automatically take notes, stay compliant, update their CRM, and serve clients",
			Industries: []domain.Industry{
				domain.IndustryFinTech,
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Fresh Clinics",
			Website: "https://www.freshclinics.com/",
			Careers: "https://www.freshclinics.com/careers",
			About:   "https://www.freshclinics.com/about-us",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                67767012,
				IDs:               nil,
				Alias:             "fresh-clinics",
				Name:              "Fresh Clinics",
				Followers:         "6K",
				Employees:         "51-200",
				AssociatedMembers: "112",
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
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Fresh-Clinics-EI_IE6821496.11,24.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Fresh-Clinics-Reviews-E6821496.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Fresh-Clinics-Jobs-E6821496.htm",
				Jobs:        "",
				Reviews:     "1",
				Salaries:    "4",
				ReviewsRate: "5.0",
				Verified:    true,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "",
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Rust: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Zig: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Scala: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Elixir: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Senior Software Engineer (Elixir)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4324393869/",
							Location:             "Surry Hills, New South Wales, Australia",
							Date:                 mustDate("2025-11-26"),
							WithSalary:           false,
							Remote:               true,
						},
					},
				},
				domain.Clojure: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Haskell: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Erlang: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
			},
			ShortDescription: "Fresh Clinics is a medical technology company on a mission to elevate care across the aesthetics and proactive healthcare space",
			Industries: []domain.Industry{
				domain.IndustryMedTech,
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Modern Milkman",
			Website: "https://themodernmilkman.co.uk/",
			Careers: "https://careers.themodernmilkman.co.uk/",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                18846815,
				IDs:               nil,
				Alias:             "the-modern-milkman",
				Name:              "Modern Milkman",
				Followers:         "14K",
				Employees:         "51-200",
				AssociatedMembers: "141",
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
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Modern-Milkman-EI_IE4066814.11,25.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Modern-Milkman-Reviews-E4066814.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Modern-Milkman-Jobs-E4066814.htm",
				Jobs:        "3",
				Reviews:     "83",
				Salaries:    "161",
				ReviewsRate: "3.0",
				Verified:    true,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "",
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Rust: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Senior Backend Engineer – Rust",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4340102154/",
							Location:             "London, England, United Kingdom",
							Date:                 mustDate("2025-11-29"),
							WithSalary:           true, // £70k/yr - £100k/yr
							Remote:               false,
						},
					},
				},
				domain.Zig: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Scala: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Elixir: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Clojure: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Haskell: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Erlang: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
			},
			ShortDescription: "At Modern Milkman, we deliver high-quality, farm-fresh essentials to your door",
			Industries:       []domain.Industry{
				// NOP
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Cynet Security",
			Website: "https://www.cynet.com/",
			Careers: "https://www.cynet.com/careers/",
			About:   "https://www.cynet.com/about-us/",
			Blog:    "https://www.cynet.com/blog/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                9363621,
				IDs:               nil,
				Alias:             "cynet-security",
				Name:              "Cynet Security",
				Followers:         "34K",
				Employees:         "201-500",
				AssociatedMembers: "317",
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
							URL:                  "https://www.linkedin.com/jobs/view/4347788624/",
							Location:             "Tel Aviv-Yafo, Tel Aviv District, Israel",
							Date:                 mustDate("2025-11-28"),
							WithSalary:           false,
							Remote:               false,
						},
					},
				},
				domain.Rust: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Zig: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Scala: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Elixir: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Clojure: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Haskell: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Erlang: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
			},
			ShortDescription: "All-in-One Cybersecurity Platform",
			Industries: []domain.Industry{
				domain.IndustryCyberSecurity,
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Ankorstore",
			Website: "https://www.ankorstore.com/",
			Careers: "https://company.ankorstore.com/",
			About:   "https://company.ankorstore.com/about-us",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                46903379,
				IDs:               nil,
				Alias:             "ankorstore",
				Name:              "Ankorstore",
				Followers:         "33K",
				Employees:         "51-200",
				AssociatedMembers: "216",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:     "ankorstore",
				Followers: "50",
				Verified:  true,
			},
			BlindProfile: domain.BlindProfile{
				Alias: "",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias: "",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Ankorstore-EI_IE4090028.11,21.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Ankorstore-Reviews-E4090028.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Ankorstore-Jobs-E4090028.htm",
				Jobs:        "10",
				Reviews:     "239",
				Salaries:    "529",
				ReviewsRate: "4.0",
				Verified:    true,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "",
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 14,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Staff Engineer – Golang / PHP",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4339345535/",
							Location:             "Paris, Île-de-France, France",
							Date:                 mustDate("2025-11-28"),
							WithSalary:           false,
							Remote:               false,
						},
					},
				},
				domain.Rust: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Zig: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Scala: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Elixir: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Clojure: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Haskell: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Erlang: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
			},
			ShortDescription: "B2B wholesale business partner connecting independent retailers and brands",
			Industries:       []domain.Industry{
				// NOP
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "PeakTew",
			Website: "https://www.peaktew.com/",
			Careers: "",
			About:   "https://www.peaktew.com/about-us",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                101096804,
				IDs:               nil,
				Alias:             "peaktew",
				Name:              "PeakTew",
				Followers:         "24K",
				Employees:         "11-50",
				AssociatedMembers: "22",
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
							Title:                "Golang Backend Developer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4325122416/",
							Location:             "Canada",
							Date:                 mustDate("2025-11-29"),
							WithSalary:           false,
							Remote:               true,
						},
					},
				},
				domain.Rust: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Zig: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Scala: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Elixir: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Clojure: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Haskell: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Erlang: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
			},
			ShortDescription: "Social network",
			Industries: []domain.Industry{
				domain.IndustrySocialMedia,
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "CloudTalk",
			Website: "https://www.cloudtalk.io/",
			Careers: "https://www.cloudtalk.io/careers/",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                17944841,
				IDs:               nil,
				Alias:             "cloudtalk.io",
				Name:              "CloudTalk",
				Followers:         "42K",
				Employees:         "51-200",
				AssociatedMembers: "184",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:     "CloudTalk-io",
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
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Cloudtalk-EI_IE509172.11,20.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Cloudtalk-Reviews-E509172.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Cloudtalk-Jobs-E509172.htm",
				Jobs:        "18",
				Reviews:     "178",
				Salaries:    "154",
				ReviewsRate: "4.3",
				Verified:    true,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "",
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 2,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Senior Software Engineer (Go)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4347912294/",
							Location:             "Prague, Prague, Czechia",
							Date:                 mustDate("2025-11-29"),
							WithSalary:           false,
							Remote:               true,
						},
					},
				},
				domain.Rust: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Zig: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Scala: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Elixir: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Clojure: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Haskell: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Erlang: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
			},
			ShortDescription: "Calling software",
			Industries:       []domain.Industry{
				// NOP
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Back Market",
			Website: "https://www.backmarket.com/",
			Careers: "",
			About:   "https://www.backmarket.com/about-us",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                4793443,
				IDs:               nil,
				Alias:             "back-market",
				Name:              "Back Market",
				Followers:         "92K",
				Employees:         "501-1K",
				AssociatedMembers: "978",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:     "backmarket-oss",
				Followers: "26",
				Verified:  true,
			},
			BlindProfile: domain.BlindProfile{
				Alias: "",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias: "",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Back-Market-EI_IE1780208.11,22.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Back-Market-Reviews-E1780208.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Back-Market-Jobs-E1780208.htm",
				Jobs:        "50",
				Reviews:     "338",
				Salaries:    "667",
				ReviewsRate: "4.1",
				Verified:    true,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "",
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 13,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Staff Backend Engineer – Go",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4339801157/",
							Location:             "Barcelona, Catalonia, Spain",
							Date:                 mustDate("2025-11-29"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Staff Backend Engineer – Go",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4318543641/",
							Location:             "Paris, Île-de-France, France",
							Date:                 mustDate("2025-12-06"),
							WithSalary:           false,
							Remote:               false,
						},
					},
				},
				domain.Rust: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Zig: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Scala: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Elixir: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Clojure: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Haskell: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Erlang: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
			},
			ShortDescription: "Marketplace for refurbished electronics",
			Industries: []domain.Industry{
				domain.IndustryECommerce,
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Zesty",
			Website: "https://zesty.co/",
			Careers: "https://zesty.co/careers/",
			About:   "https://zesty.co/about-us/",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                19106475,
				IDs:               nil,
				Alias:             "zestyco",
				Name:              "Zesty",
				Followers:         "21K",
				Employees:         "51-200",
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
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Zesty-EI_IE5284749.11,16.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Zesty-Reviews-E5284749.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Zesty-Jobs-E5284749.htm",
				Jobs:        "",
				Reviews:     "91",
				Salaries:    "85",
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
							Title:                "Senior Backend Developer – Go",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4339955174/",
							Location:             "Tel Aviv-Yafo, Tel Aviv District, Israel",
							Date:                 mustDate("2025-11-26"),
							WithSalary:           false,
							Remote:               false,
						},
					},
				},
				domain.Rust: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Zig: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Scala: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Elixir: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Clojure: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Haskell: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Erlang: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
			},
			ShortDescription: "A Kubernetes optimization platform powered by multi-layer automation, built to ensure peak performance and cost efficiency across resource and financial layers",
			Industries: []domain.Industry{
				domain.IndustryDevOps,
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Loyal Guru",
			Website: "https://www.loyal.guru/",
			Careers: "https://www.loyal.guru/about-us/",
			About:   "https://jobs.loyal.guru/",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                5222128,
				IDs:               nil,
				Alias:             "loyal-guru",
				Name:              "Loyal Guru",
				Followers:         "9K",
				Employees:         "51-200",
				AssociatedMembers: "65",
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
							Title:                "Backend Developer Go",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4347642829/",
							Location:             "Spain",
							Date:                 mustDate("2025-11-27"),
							WithSalary:           false,
							Remote:               true,
						},
					},
				},
				domain.Rust: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Zig: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Scala: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Elixir: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Clojure: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Haskell: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Erlang: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
			},
			ShortDescription: "Create & scale personalized loyalty programs with gamification, promotions, rewards, and budget control",
			Industries: []domain.Industry{
				domain.IndustryMarTech,
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Vyro",
			Website: "https://vyro.ai/",
			Careers: "https://vyro.ai/careers",
			About:   "https://vyro.ai/about",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                64501680,
				IDs:               nil,
				Alias:             "vyro-ai",
				Name:              "Vyro",
				Followers:         "77K",
				Employees:         "51-200",
				AssociatedMembers: "165",
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
							Title:                "Go (Golang) Developer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4339012184/",
							Location:             "Bandar Banang Jaya, Johore, Malaysia",
							Date:                 mustDate("2025-11-26"),
							WithSalary:           false,
							Remote:               true,
						},
					},
				},
				domain.Rust: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Zig: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Scala: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Elixir: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Clojure: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Haskell: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Erlang: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
			},
			ShortDescription: "Content creation",
			Industries: []domain.Industry{
				domain.IndustrySocialMedia,
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Foxglove",
			Website: "https://foxglove.dev/",
			Careers: "https://foxglove.dev/careers",
			About:   "https://foxglove.dev/about",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                72660627,
				IDs:               nil,
				Alias:             "foxglovedev",
				Name:              "Foxglove",
				Followers:         "17K",
				Employees:         "51-200",
				AssociatedMembers: "62",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:     "foxglove",
				Followers: "558",
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
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Rust: {
					GitHubRepositoriesCount: 3,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Senior Rust Engineer, Frontend",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4285040837/",
							Location:             "San Francisco Bay Area",
							Date:                 mustDate("2025-11-25"),
							WithSalary:           true, // $165k/yr - $200k/yr
							Remote:               false,
						},
						{
							Title:                "Senior Software Engineer, Rust",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4285041821/",
							Location:             "New South Wales, Australia",
							Date:                 mustDate("2025-12-06"),
							WithSalary:           true, // $180k/yr - $220k/yr
							Remote:               true,
						},
					},
				},
				domain.Zig: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Scala: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Elixir: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Clojure: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Haskell: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Erlang: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
			},
			ShortDescription: "Visualize, debug, and manage multimodal data in one purpose-built platform for robotics and embodied AI development",
			Industries: []domain.Industry{
				domain.IndustryDevOps,
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Sky",
			Website: "https://www.sky.com/",
			Careers: "https://careers.sky.com/",
			About:   "https://careers.sky.com/about",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                3318,
				IDs:               []int{3318, 10983400},
				Alias:             "sky",
				Name:              "Sky",
				Followers:         "567K",
				Employees:         "10K+",
				AssociatedMembers: "35,523",
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
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Sky-EI_IE3903.11,14.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Sky-Reviews-E3903.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Sky-Jobs-E3903.htm",
				Jobs:        "190",
				Reviews:     "7K",
				Salaries:    "14K",
				ReviewsRate: "3.5",
				Verified:    false,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "",
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Rust: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Zig: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Scala: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Scala Developer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4341524010/",
							Location:             "Isleworth, England, United Kingdom",
							Date:                 mustDate("2025-12-02"),
							WithSalary:           false,
							Remote:               false,
						},
					},
				},
				domain.Elixir: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Clojure: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Haskell: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Erlang: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
			},
			ShortDescription: "TV",
			Industries: []domain.Industry{
				domain.IndustryEntertainment,
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Reducate",
			Website: "https://www.reducate.com/",
			Careers: "https://www.reducate.com/careers/",
			About:   "https://www.reducate.com/about-us/",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                85115348,
				IDs:               nil,
				Alias:             "reducate",
				Name:              "Reducate",
				Followers:         "2K",
				Employees:         "51-200",
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
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Software Engineer (Python, Go, Node)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4325399277/",
							Location:             "The Randstad, Netherlands",
							Date:                 mustDate("2025-12-03"),
							WithSalary:           false,
							Remote:               false,
						},
					},
				},
				domain.Rust: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Zig: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Scala: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Elixir: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Clojure: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Haskell: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Erlang: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
			},
			ShortDescription: "",
			Industries: []domain.Industry{
				domain.IndustryEdTech,
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "myo",
			Website: "https://myo.de/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                10141276,
				IDs:               nil,
				Alias:             "myo-healthcare",
				Name:              "myo",
				Followers:         "2K",
				Employees:         "11-50",
				AssociatedMembers: "39",
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
					Vacancies:               []domain.Vacancy{},
				},
				domain.Rust: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Zig: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Scala: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Elixir: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Senior Fullstack Engineer (Elixir/Phoenix)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4348395855/",
							Location:             "Germany",
							Date:                 mustDate("2025-12-05"),
							WithSalary:           false,
							Remote:               true,
						},
					},
				},
				domain.Clojure: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Haskell: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Erlang: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
			},
			ShortDescription: "Communication service for the care sector",
			Industries: []domain.Industry{
				domain.IndustryHealthTech,
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Ripple",
			Website: "https://ripple.com/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                3249443,
				IDs:               nil,
				Alias:             "ripple",
				Name:              "Ripple",
				Followers:         "424K",
				Employees:         "501-1K",
				AssociatedMembers: "1,459",
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
					Vacancies:               []domain.Vacancy{},
				},
				domain.Rust: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Senior Software Engineer, C++ / Rust",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4318356923/",
							Location:             "New York, NY",
							Date:                 mustDate("2025-12-08"),
							WithSalary:           true, // $152k/yr - $190k/yr
							Remote:               false,
						},
					},
				},
				domain.Zig: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Scala: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Elixir: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Clojure: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Haskell: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Erlang: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
			},
			ShortDescription: "",
			Industries:       []domain.Industry{
				// NOP
			},
			Ignore: true, // Cryptocurrency
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Innoflight",
			Website: "https://innoflight.com/",
			Careers: "https://innoflight.com/careers/",
			About:   "https://innoflight.com/about-us/",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                3028857,
				IDs:               nil,
				Alias:             "innoflight-llc",
				Name:              "Innoflight",
				Followers:         "14K",
				Employees:         "51-200",
				AssociatedMembers: "162",
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
					Vacancies:               []domain.Vacancy{},
				},
				domain.Rust: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Embedded Software Engineer IV – Rust",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4343181290/",
							Location:             "San Diego, CA",
							Date:                 mustDate("2025-12-11"),
							WithSalary:           true, // $150k/yr - $170k/yr
							Remote:               false,
						},
					},
				},
				domain.Zig: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Scala: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Elixir: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Clojure: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Haskell: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Erlang: {
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
			},
			ShortDescription: "",
			Industries: []domain.Industry{
				domain.IndustryCyberSecurity,
			},
		},
	}
}
