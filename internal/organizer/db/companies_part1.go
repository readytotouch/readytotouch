package db

import "github.com/readytotouch/readytotouch/internal/domain"

func companiesPart1() []domain.CompanyProfile {
	return []domain.CompanyProfile{
		// Favorites
		// Favorites | ReadyToTouch
		{
			ID:   0,  // system
			Type: "", // system
			Name: "ReadyToTouch",
			URL:  "https://readytotouch.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    97909464,
				IDs:   nil,
				Alias: "readytotouch",
				Name:  "ReadyToTouch",
			},
			GitHubProfile:     domain.GitHubProfile{},
			GlassdoorProfile:  domain.GlassdoorProfile{},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Vacancies:         domain.Vacancies{},
			ShortDescription:  "Service for simplifying job search",
			Industries:        []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
			},
		},

		// Favorites | DocHQ
		{
			ID:   0,  // system
			Type: "", // system
			Name: "DocHQ",
			URL:  "https://dochq.co.uk/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    14136494,
				Alias: "dochq",
				Name:  "DocHQ",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "dochq",
				GoRepositoryCount: 9,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "",
				ReviewsURL:  "",
				Verified:    false,
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "",
			GoMainLanguage:    true,
			Vacancies: domain.Vacancies{
				domain.Go:      []string{},
				domain.Rust:    nil,
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
			},
			ShortDescription: "Health Tech company",
			Industries: []domain.Industry{
				domain.IndustryHealthTech,
			},
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
			},
		},

		// BigTech
		// BigTech | Google
		{
			ID:   0,  // system
			Type: "", // system
			Name: "Google",
			URL:  "https://www.google.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    1441,
				Alias: "google",
				Name:  "Google",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "google",
				GoRepositoryCount: 157,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Google-EI_IE9079.11,17.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Google-Reviews-E9079.htm",
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go:      []string{},
				domain.Rust:    nil,
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
			},
			ShortDescription: "American multinational technology company",
			Industries:       []domain.Industry{
				// Too many industries
			},
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
				domain.Czechia,
			},
		},

		// BigTech | Mozilla
		{
			ID:   0,  // system
			Type: "", // system
			Name: "Mozilla",
			URL:  "https://www.mozilla.org/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    13948,
				Alias: "mozilla-corporation",
				Name:  "Mozilla",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "mozilla",
				GoRepositoryCount: 24,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Mozilla-EI_IE19129.11,18.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Mozilla-Reviews-E19129.htm",
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go: []string{
					"https://app.otta.com/jobs/x1nUGqRw",
				},
				domain.Rust: []string{
					"https://app.otta.com/jobs/x1nUGqRw",
				},
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
			},
			ShortDescription: "Organization dedicated to making the web better",
			Industries:       []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{
				domain.Czechia,
			},
		},

		// BigTech | Discord
		{
			ID:   0,  // system
			Type: "", // system
			Name: "Discord",
			URL:  "https://discord.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    3765675,
				Alias: "discord",
				Name:  "Discord",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "discord",
				GoRepositoryCount: 10, // Rust 35
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Discord-EI_IE910317.11,18.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Discord-Reviews-E910317.htm",
			},
			OttaProfileSlug:   "Discord",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go: nil,
				domain.Rust: []string{
					"https://app.welcometothejungle.com/jobs/q5H-48wM",
				},
				domain.Zig:   nil,
				domain.Scala: nil,
				domain.Elixir: []string{
					"https://app.welcometothejungle.com/jobs/HxEBR9jp",
				},
				domain.Clojure: nil,
				domain.Haskell: nil,
			},
			ShortDescription:          "Chat",
			Industries:                []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{},
		},

		// BigTech | Figma
		{
			ID:   0,  // system
			Type: "", // system
			Name: "Figma",
			URL:  "https://www.figma.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    3650502,
				Alias: "figma",
				Name:  "Figma",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "figma",
				GoRepositoryCount: 14, // Rust 4
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Figma-EI_IE1537286.11,16.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Figma-Reviews-E1537286.htm",
			},
			OttaProfileSlug:   "Figma",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go: []string{
					"https://app.welcometothejungle.com/jobs/YXcMzOW8",
				},
				domain.Rust: []string{
					"https://app.welcometothejungle.com/jobs/YXcMzOW8",
				},
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
			},
			ShortDescription: "Design platform for product teams",
			Industries:       []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
				domain.Czechia,
			},
		},

		// BigTech | Microsoft
		{
			ID:   0,  // system
			Type: "", // system
			Name: "Microsoft",
			URL:  "https://www.microsoft.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    1035,
				Alias: "microsoft",
				Name:  "Microsoft",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "microsoft",
				GoRepositoryCount: 78,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Microsoft-EI_IE1651.11,20.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Microsoft-Reviews-E1651.htm",
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "https://www.youtube.com/@Microsoft",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go: []string{
					"https://www.linkedin.com/jobs/view/3908309859/",
					"https://www.linkedin.com/jobs/view/3905989512/",
				},
				domain.Rust: []string{
					"https://www.linkedin.com/jobs/view/4038788830/",
				},
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
			},
			ShortDescription: "American multinational technology company",
			Industries:       []domain.Industry{
				// Too many industries
			},
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
				domain.Czechia,
			},
		},

		// BigTech | Amazon
		{
			ID:   0,  // system
			Type: "", // system
			Name: "Amazon",
			URL:  "https://www.aboutamazon.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID: 0,
				IDs: []int{
					1586,    // Amazon
					2382910, // Amazon Web Services (AWS)
				},
				Alias: "amazon",
				Name:  "Amazon",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "aws", // "amzn",
				GoRepositoryCount: 71,    // 2,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Amazon-EI_IE6036.11,17.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Amazon-Reviews-E6036.htm",
			},
			OttaProfileSlug:   "Amazon",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go: []string{
					"https://app.otta.com/jobs/G0JaYF4Z",
				},
				domain.Rust: []string{
					"https://app.otta.com/jobs/G0JaYF4Z",
				},
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
			},
			ShortDescription: "",
			Industries:       []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
				domain.Czechia,
			},
		},

		// BigTech | IBM
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "IBM",
			Website: "https://www.ibm.com/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:       1009,
				Alias:    "ibm",
				Name:     "IBM",
				Verified: true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "IBM",
				GoRepositoryCount: 199,
				Verified:          true,
			},
			BlindProfile:     domain.BlindProfile{},
			LevelsFyiProfile: domain.LevelsFyiProfile{},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-IBM-EI_IE354.11,14.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/IBM-Reviews-E354.htm",
				Verified:    true,
			},
			IndeedProfile:     domain.IndeedProfile{},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "https://www.youtube.com/@IBMTechnology",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Software Developer (Back-End) — Go/Python",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4036994106/",
							Date:             mustDate("2024-10-25"),
						},
					},
				},
				domain.Rust: {},
				domain.Zig:  {},
				domain.Scala: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Lead Scala Software Engineer",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/3998480636/",
							Date:             mustDate("2024-11-08"),
						},
						{
							Title:            "Scala Developer",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4063746170/",
							Date:             mustDate("2024-11-02"),
						},
					},
				},
				domain.Elixir:  {},
				domain.Clojure: {},
				domain.Haskell: {},
			},
			ShortDescription: "American multinational technology company",
			DealroomURL:      "",
			CrunchbaseURL:    "",
			PitchbookURL:     "",
			YahooFinanceURL:  "",
			GoogleFinanceURL: "",
			YCombinatorURL:   "",
			Industries:       []domain.Industry{
				// Too many industries
			},
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
				domain.Czechia,
			},
		},

		// BigTech | SAP
		{
			ID:   0,  // system
			Type: "", // system
			Name: "SAP",
			URL:  "https://www.sap.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:       1115,
				IDs:      []int{1115, 2573558, 2818, 166185, 5822},
				Alias:    "sap",
				Name:     "SAP",
				Verified: true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "SAP",
				GoRepositoryCount: 53,
				Verified:          true,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/sap",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/SAP-Reviews-E10471.htm",
				Verified:    true,
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go: []string{
					"https://www.linkedin.com/jobs/view/3943874981/", // strong understanding of Golang
				},
				domain.Rust: []string{
					"https://www.linkedin.com/jobs/view/4059368134/",
				},
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
			},
			ShortDescription: "Enterprise application and business AI company",
			Industries:       []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
				domain.Czechia,
			},
		},

		// BigTech | Oracle
		{
			ID:   0,  // system
			Type: "", // system
			Name: "Oracle",
			URL:  "https://www.oracle.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:       1028,
				Alias:    "oracle",
				Name:     "Oracle",
				Verified: true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "oracle",
				GoRepositoryCount: 31,
				Verified:          false,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Oracle-EI_IE1737.11,17.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Oracle-Reviews-E1737.htm",
				Verified:    true,
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go:   nil,
				domain.Rust: nil,
				domain.Zig:  nil,
				domain.Scala: []string{
					"https://www.linkedin.com/jobs/view/4034901984/",
				},
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
			},
			ShortDescription: "American multinational technology company",
			Industries:       []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
				domain.Czechia,
			},
		},

		// Social
		// Social | GitHub
		{
			ID:   0,  // system
			Type: "", // system
			Name: "GitHub",
			URL:  "https://github.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    1418841,
				Alias: "github",
				Name:  "GitHub",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "github",
				GoRepositoryCount: 18,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-GitHub-EI_IE671945.11,17.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/GitHub-Reviews-E671945.htm",
			},
			OttaProfileSlug:   "Github",
			YouTubeChannelURL: "https://www.youtube.com/@GitHub",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go: []string{
					"https://www.linkedin.com/jobs/view/3914880703/",
				},
				domain.Rust:    nil,
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
			},
			ShortDescription: "Developer platform",
			Industries: []domain.Industry{
				domain.IndustryDevOps,
			},
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
				domain.Czechia,
			},
		},

		// Social | GitLab
		{
			ID:   0,  // system
			Type: "", // system
			Name: "GitLab",
			URL:  "https://gitlab.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:       5101804,
				Alias:    "gitlab-com",
				Name:     "GitLab",
				Verified: true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "gitlabhq",
				GoRepositoryCount: 3,
				Verified:          false,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-GitLab-EI_IE1296544.11,17.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/GitLab-Reviews-E1296544.htm",
				Verified:    true,
			},
			OttaProfileSlug:   "GitLab-1",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go: []string{
					"https://www.linkedin.com/jobs/view/4018780627/", // Significant experience with Go
				},
				domain.Rust:    nil,
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
			},
			ShortDescription: "DevOps platform",
			Industries: []domain.Industry{
				domain.IndustryDevOps,
			},
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
				domain.Czechia,
			},
		},

		// Social | LinkedIn
		{
			ID:   0,  // system
			Type: "", // system
			Name: "LinkedIn",
			URL:  "https://www.linkedin.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    1337,
				Alias: "linkedin",
				Name:  "LinkedIn",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "linkedin",
				GoRepositoryCount: 8,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-LinkedIn-EI_IE34865.11,19.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/LinkedIn-Reviews-E34865.htm",
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go: []string{
					"https://www.linkedin.com/jobs/view/4023583351/",
				},
				domain.Rust: []string{
					"https://www.linkedin.com/jobs/view/4023579503/",
					"https://www.linkedin.com/jobs/view/4023583351/",
					"https://www.linkedin.com/jobs/view/3981396826/",
				},
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
			},
			ShortDescription: "Business and employment-focused social media platform",
			Industries: []domain.Industry{
				domain.IndustrySocialMedia,
			},
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
				domain.Czechia,
			},
		},

		// Social | Reddit
		{
			ID:   0,  // system
			Type: "", // system
			Name: "Reddit",
			URL:  "https://www.reddit.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    150573,
				Alias: "reddit-com",
				Name:  "Reddit, Inc.",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "reddit",
				GoRepositoryCount: 20,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Reddit-EI_IE796358.11,17.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Reddit-Reviews-E796358.htm",
			},
			OttaProfileSlug:   "Reddit-1",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go: []string{
					"https://app.otta.com/jobs/l-8sLbJL",
				},
				domain.Rust:    nil,
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
			},
			ShortDescription: "Network of communities where people can dive into their interests, hobbies and passions",
			Industries: []domain.Industry{
				domain.IndustrySocialMedia,
			},
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
			},
		},

		// Social | Medium
		{
			ID:   0,  // system
			Type: "", // system
			Name: "Medium",
			URL:  "https://medium.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    3039001,
				Alias: "medium-com",
				Name:  "Medium",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "Medium",
				GoRepositoryCount: 20,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Medium-EI_IE784883.11,17.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Medium-Reviews-E784883.htm",
			},
			OttaProfileSlug:   "Medium",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go: []string{
					"https://www.linkedin.com/jobs/view/3910516815/",
				},
				domain.Rust:    nil,
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
			},
			ShortDescription: "Home for human stories and ideas",
			Industries: []domain.Industry{
				domain.IndustrySocialMedia,
			},
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
				domain.Czechia,
			},
		},

		// Social | Pinterest
		{
			ID:   0,  // system
			Type: "", // system
			Name: "Pinterest",
			URL:  "https://www.pinterest.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    1124131,
				Alias: "pinterest",
				Name:  "Pinterest",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "pinterest",
				GoRepositoryCount: 3,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Pinterest-EI_IE503467.11,20.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Pinterest-Reviews-E503467.htm",
			},
			OttaProfileSlug:   "Pinterest",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go: []string{
					"https://app.otta.com/jobs/JwXeQ-gm",
				},
				domain.Rust:    nil,
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
			},
			ShortDescription: "Visual discovery engine for finding ideas",
			Industries: []domain.Industry{
				domain.IndustrySocialMedia,
			},
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
				domain.Czechia,
			},
		},

		// Social | Snap
		{
			ID:   0,  // system
			Type: "", // system
			Name: "Snap",
			URL:  "https://snap.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    15191764,
				Alias: "snap-inc-co",
				Name:  "Snap Inc.",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "Snapchat",
				GoRepositoryCount: 1,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Snap-EI_IE671946.11,15.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Snap-Reviews-E671946.htm",
			},
			OttaProfileSlug:   "Snap",
			YouTubeChannelURL: "https://www.youtube.com/@snapinc",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go: []string{
					"https://app.otta.com/jobs/hr1D1_FZ",
				},
				domain.Rust:    nil,
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
			},
			ShortDescription: "American multimedia instant messaging app",
			Industries: []domain.Industry{
				domain.IndustrySocialMedia,
			},
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
				domain.Czechia,
			},
		},

		// Social | BeReal.
		{
			ID:   0,
			Type: "",
			Name: "BeReal.",
			URL:  "https://bereal.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    34731272,
				Alias: "bereal-app",
				Name:  "BeReal.",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "BeReal-App",
				GoRepositoryCount: 0,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-BeReal-EI_IE7468524.11,17.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/BeReal-Reviews-E7468524.htm",
			},
			OttaProfileSlug:   "BeReal",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go:      []string{},
				domain.Rust:    nil,
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
			},
			ShortDescription: "Simplest photo sharing app",
			Industries: []domain.Industry{
				domain.IndustrySocialMedia,
			},
			HasEmployeesFromCountries: []domain.Country{},
		},

		// Favorites
		// Favorites | VictoriaMetrics
		{
			ID:   0,  // system
			Type: "", // system
			Name: "VictoriaMetrics",
			URL:  "https://victoriametrics.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    30169914,
				Alias: "victoriametrics",
				Name:  "VictoriaMetrics",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "VictoriaMetrics",
				GoRepositoryCount: 10,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "",
				ReviewsURL:  "",
				Verified:    false,
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "",
			GoMainLanguage:    true,
			Vacancies: domain.Vacancies{
				domain.Go:      []string{},
				domain.Rust:    nil,
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
			},
			ShortDescription: "Time series database and easy-to-use high performance monitoring solutions",
			Industries: []domain.Industry{
				domain.IndustryDevOps,
			},
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
			},
		},

		// Favorites | Grammarly
		{
			ID:   0,  // system
			Type: "", // system
			Name: "Grammarly",
			URL:  "https://www.grammarly.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    1253671,
				Alias: "grammarly",
				Name:  "Grammarly",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "grammarly",
				GoRepositoryCount: 0,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Grammarly-EI_IE636873.11,20.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Grammarly-Reviews-E636873.htm",
			},
			OttaProfileSlug:   "Grammarly",
			YouTubeChannelURL: "https://www.youtube.com/@grammarly",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go: []string{
					"https://app.otta.com/jobs/jonw8YN_",
				},
				domain.Rust:    nil,
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
			},
			ShortDescription: "Typing assistant",
			Industries:       []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
				domain.Czechia,
			},
		},

		// Favorites | Influ2
		{
			ID:   0,                         // system
			Type: domain.CompanyTypeStartup, // system
			Name: "Influ2",
			URL:  "https://www.influ2.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:       18119989,
				Alias:    "influ2",
				Name:     "Influ2",
				Verified: true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "",
				GoRepositoryCount: 0,
				Verified:          false,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Influ2-EI_IE4066744.11,17.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Influ2-Reviews-E4066744.htm",
				Verified:    true,
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go: []string{
					"https://www.linkedin.com/jobs/view/4058445939/",
				},
				domain.Rust:    nil,
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
			},
			ShortDescription: "Person-based advertising platform",
			Industries: []domain.Industry{
				domain.IndustryAdTech,
			},
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
			},
		},

		// Tech
		// Tech | Netlify
		{
			ID:   0,  // system
			Type: "", // system
			Name: "Netlify",
			URL:  "https://www.netlify.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    6392431,
				Alias: "netlify",
				Name:  "Netlify",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "netlify",
				GoRepositoryCount: 35,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Netlify-EI_IE1426251.11,18.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Netlify-Reviews-E1426251.htm",
			},
			OttaProfileSlug:   "Netlify",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go: []string{
					"https://app.otta.com/jobs/qdIzdfOu",
				},
				domain.Rust:    nil,
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
			},
			ShortDescription: "Static hosting service",
			Industries: []domain.Industry{
				domain.IndustryDevOps,
			},
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
			},
		},

		// Tech | Vercel
		{
			ID:   0,  // system
			Type: "", // system
			Name: "Vercel",
			URL:  "https://vercel.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    16181286,
				Alias: "vercel",
				Name:  "Vercel",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "vercel",
				GoRepositoryCount: 6,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Vercel-EI_IE6510369.11,17.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Vercel-Reviews-E6510369.htm",
			},
			OttaProfileSlug:   "Vercel",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go: []string{
					"https://www.linkedin.com/jobs/view/4048326475/",
				},
				domain.Rust:    nil,
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
			},
			ShortDescription: "Static hosting service",
			Industries: []domain.Industry{
				domain.IndustryDevOps,
			},
			HasEmployeesFromCountries: []domain.Country{},
		},

		// Tech | Fastly
		{
			ID:   0,  // system
			Type: "", // system
			Name: "Fastly",
			URL:  "https://www.fastly.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    2602522,
				Alias: "fastly",
				Name:  "Fastly",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "fastly",
				GoRepositoryCount: 35,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Fastly-EI_IE814479.11,17.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Fastly-Reviews-E814479.htm",
			},
			OttaProfileSlug:   "Fastly",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go: []string{
					"https://app.welcometothejungle.com/jobs/mZtAqCtB",
				},
				domain.Rust: []string{
					"https://app.welcometothejungle.com/jobs/mZtAqCtB",
				},
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
			},
			ShortDescription: "Edge cloud platform",
			Industries: []domain.Industry{
				domain.IndustryDevOps,
			},
			HasEmployeesFromCountries: []domain.Country{},
		},

		// Tech | Dropbox
		{
			ID:   0,  // system
			Type: "", // system
			Name: "Dropbox",
			URL:  "https://www.dropbox.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    167251,
				Alias: "dropbox",
				Name:  "Dropbox",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "dropbox",
				GoRepositoryCount: 22, // Rust 16
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Dropbox-EI_IE415350.11,18.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Dropbox-Reviews-E415350.htm",
			},
			OttaProfileSlug:   "Dropbox",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go: []string{
					"https://app.welcometothejungle.com/jobs/WKQvjdTx",
				},
				domain.Rust: []string{
					"https://app.welcometothejungle.com/jobs/WKQvjdTx",
				},
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
			},
			ShortDescription: "File hosting service",
			Industries: []domain.Industry{
				domain.IndustryDevOps,
			},
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
			},
		},

		// Tech | Docker
		{
			ID:   0,  // system
			Type: "", // system
			Name: "Docker",
			URL:  "https://www.docker.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    1301808,
				Alias: "docker",
				Name:  "Docker, Inc",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "docker",
				GoRepositoryCount: 28,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Docker-EI_IE1089506.11,17.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Docker-Reviews-E1089506.htm",
			},
			OttaProfileSlug:   "Docker-1",
			YouTubeChannelURL: "",
			GoMainLanguage:    true,
			Vacancies: domain.Vacancies{
				domain.Go:      []string{},
				domain.Rust:    nil,
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
			},
			ShortDescription: "Multi-platform application for developers",
			Industries: []domain.Industry{
				domain.IndustryDevOps,
			},
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
			},
		},

		// Tech | Grafana Labs
		{
			ID:   0,  // system
			Type: "", // system
			Name: "Grafana Labs",
			URL:  "https://grafana.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    11062162,
				Alias: "grafana-labs",
				Name:  "Grafana Labs",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "grafana",
				GoRepositoryCount: 233,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Grafana-Labs-EI_IE2300269.11,23.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Grafana-Labs-Reviews-E2300269.htm",
			},
			OttaProfileSlug:   "Grafana-Labs",
			YouTubeChannelURL: "https://www.youtube.com/channel/UCYCwgQAMm9sTJv0rgwQLCxw",
			GoMainLanguage:    true,
			Vacancies: domain.Vacancies{
				domain.Go:      []string{},
				domain.Rust:    nil,
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
			},
			ShortDescription: "Open source monitoring and observability platform",
			Industries: []domain.Industry{
				domain.IndustryDevOps,
			},
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
			},
		},

		// Tech | HashiCorp
		{
			ID:   0,  // system
			Type: "", // system
			Name: "HashiCorp",
			URL:  "https://www.hashicorp.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    2830763,
				Alias: "hashicorp",
				Name:  "HashiCorp",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "hashicorp",
				GoRepositoryCount: 296,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-HashiCorp-EI_IE1359860.11,20.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/HashiCorp-Reviews-E1359860.htm",
			},
			OttaProfileSlug:   "HashiCorp",
			YouTubeChannelURL: "",
			GoMainLanguage:    true,
			Vacancies: domain.Vacancies{
				domain.Go:      []string{},
				domain.Rust:    nil,
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
			},
			ShortDescription: "Software company that provides modular DevOps infrastructure provisioning and management products",
			Industries: []domain.Industry{
				domain.IndustryDevOps,
			},
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
				domain.Czechia,
			},
		},

		// Tech | CrowdStrike
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "CrowdStrike",
			Website: "https://www.crowdstrike.com/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:       2497653,
				Alias:    "crowdstrike",
				Name:     "CrowdStrike",
				Verified: true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "CrowdStrike",
				GoRepositoryCount: 20,
			},
			BlindProfile:     domain.BlindProfile{},
			LevelsFyiProfile: domain.LevelsFyiProfile{},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-CrowdStrike-EI_IE795976.11,22.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/CrowdStrike-Reviews-E795976.htm",
			},
			IndeedProfile:     domain.IndeedProfile{},
			OttaProfileSlug:   "CrowdStrike-1",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Golang Engineer III, Identity Protection (Remote)",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/3996818515/",
							Date:             mustDate("2024-11-03"),
						},
					},
				},
				domain.Rust: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Software Engineer — Rust (Remote)",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4040870810/",
							Date:             mustDate("2024-11-01"),
						},
					},
				},
				domain.Zig:     {},
				domain.Scala:   {},
				domain.Elixir:  {},
				domain.Clojure: {},
				domain.Haskell: {},
			},
			ShortDescription: "Cybersecurity technology company",
			DealroomURL:      "",
			CrunchbaseURL:    "",
			PitchbookURL:     "",
			YahooFinanceURL:  "",
			GoogleFinanceURL: "",
			YCombinatorURL:   "",
			Industries:       []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{
				domain.Czechia,
			},
		},

		// Tech | Cockroach Labs
		{
			ID:   0,  // system
			Type: "", // system
			Name: "Cockroach Labs",
			URL:  "https://www.cockroachlabs.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    9309408,
				Alias: "cockroach-labs",
				Name:  "Cockroach Labs",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "cockroachdb",
				GoRepositoryCount: 92,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Cockroach-Labs-EI_IE1168502.11,25.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Cockroach-Labs-Reviews-E1168502.htm",
			},
			OttaProfileSlug:   "Cockroach-Labs",
			YouTubeChannelURL: "https://www.youtube.com/@cockroachdb",
			GoMainLanguage:    true,
			Vacancies: domain.Vacancies{
				domain.Go: []string{
					"https://app.otta.com/jobs/uPMrLAMV",
				},
				domain.Rust:    nil,
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
			},
			ShortDescription: "Distributed database with standard SQL for cloud applications",
			Industries: []domain.Industry{
				domain.IndustryDevOps,
			},
		},

		// Tech | Timescale
		{
			ID:   0,  // system
			Type: "", // system
			Name: "Timescale",
			URL:  "https://www.timescale.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    11065434,
				Alias: "timescaledb",
				Name:  "Timescale",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "timescale",
				GoRepositoryCount: 11,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Timescale-EI_IE2214356.11,20.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Timescale-Reviews-E2214356.htm",
			},
			OttaProfileSlug:   "Timescale",
			YouTubeChannelURL: "https://www.youtube.com/@TimescaleDB",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go: []string{
					"https://app.otta.com/jobs/qZT84OgV",
				},
				domain.Rust:    nil,
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
			},
			ShortDescription: "Relational open-source database with PostgreSQL",
			Industries: []domain.Industry{
				domain.IndustryDevOps,
			},
		},

		// Tech | ScyllaDB
		{
			ID:   0,  // system
			Type: "", // system
			Name: "ScyllaDB",
			URL:  "https://www.scylladb.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    10201068,
				Alias: "scylladb",
				Name:  "ScyllaDB",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "scylladb",
				GoRepositoryCount: 18,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-ScyllaDB-EI_IE1622223.11,19.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/ScyllaDB-Reviews-E1622223.htm",
			},
			OttaProfileSlug:   "ScyllaDB",
			YouTubeChannelURL: "https://www.youtube.com/@ScyllaDB",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go: []string{
					"https://app.otta.com/jobs/Xj1J-uAC",
				},
				domain.Rust:    nil,
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
			},
			ShortDescription: "ScyllaDB is the distributed database",
			Industries: []domain.Industry{
				domain.IndustryDevOps,
			},
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
				domain.Czechia,
			},
		},

		// Tech | Percona
		{
			ID:   0,  // system
			Type: "", // system
			Name: "Percona",
			URL:  "https://www.percona.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    421929,
				Alias: "percona",
				Name:  "Percona",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "percona",
				GoRepositoryCount: 46,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Percona-EI_IE283779.11,18.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Percona-Reviews-E283779.htm",
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go: []string{
					"https://app.otta.com/jobs/5n3UU6ZU",
				},
				domain.Rust:    nil,
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
			},
			ShortDescription: "Open source database software",
			Industries: []domain.Industry{
				domain.IndustryDevOps,
			},
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
			},
		},

		// Tech | Elastic
		{
			ID:   0,  // system
			Type: "", // system
			Name: "Elastic",
			URL:  "https://www.elastic.co/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    814025,
				Alias: "elastic-co",
				Name:  "Elastic",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "elastic",
				GoRepositoryCount: 124,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Elastic-EI_IE751551.11,18.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Elastic-Reviews-E751551.htm",
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go: []string{
					"https://www.linkedin.com/jobs/view/3999755835/",
				},
				domain.Rust:    nil,
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
			},
			ShortDescription: "Search engine",
			Industries: []domain.Industry{
				domain.IndustryDevOps,
			},
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
				domain.Czechia,
			},
		},

		// Tech | MongoDB
		{
			ID:   0,  // system
			Type: "", // system
			Name: "MongoDB",
			URL:  "https://www.mongodb.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    783611,
				Alias: "mongodbinc",
				Name:  "MongoDB",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "mongodb",
				GoRepositoryCount: 33,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-MongoDB-EI_IE433703.11,18.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/MongoDB-Reviews-E433703.htm",
			},
			OttaProfileSlug:   "MongoDB",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go: []string{
					"https://app.otta.com/jobs/jWhHaFVW",
				},
				domain.Rust:    nil,
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
			},
			ShortDescription: "",
			Industries: []domain.Industry{
				domain.IndustryDevOps,
			},
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
			},
		},

		// Tech | FerretDB
		{
			ID:   0,  // system
			Type: "", // system
			Name: "FerretDB",
			URL:  "https://www.ferretdb.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    80672744,
				Alias: "ferretdb",
				Name:  "FerretDB",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "FerretDB",
				GoRepositoryCount: 5,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "",
				ReviewsURL:  "",
				Verified:    false,
			},
			OttaProfileSlug:   "FerretDB",
			YouTubeChannelURL: "",
			GoMainLanguage:    true,
			Vacancies: domain.Vacancies{
				domain.Go:      []string{},
				domain.Rust:    nil,
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
			},
			ShortDescription: "MongoDB compatible database",
			Industries: []domain.Industry{
				domain.IndustryDevOps,
			},
		},

		// Tech | Redis
		{
			ID:   0,  // system
			Type: "", // system
			Name: "Redis",
			URL:  "https://redis.io/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    2014725,
				Alias: "redisinc",
				Name:  "Redis",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "redis",
				GoRepositoryCount: 2,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Redis-EI_IE928722.11,16.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Redis-Reviews-E928722.htm",
			},
			OttaProfileSlug:   "Redis",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go: []string{
					"https://www.linkedin.com/jobs/view/4021697979/",
				},
				domain.Rust: []string{
					"https://www.linkedin.com/jobs/view/3963199017/",
				},
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
			},
			ShortDescription: "Key-value database",
			Industries: []domain.Industry{
				domain.IndustryDevOps,
			},
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
			},
		},

		// Tech | DigitalOcean
		{
			ID:   0,  // system
			Type: "", // system
			Name: "DigitalOcean",
			URL:  "https://www.digitalocean.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    2601253,
				Alias: "digitalocean",
				Name:  "DigitalOcean",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "digitalocean",
				GoRepositoryCount: 116,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-DigitalOcean-EI_IE823482.11,23.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/DigitalOcean-Reviews-E823482.htm",
			},
			OttaProfileSlug:   "DigitalOcean",
			YouTubeChannelURL: "https://www.youtube.com/@DigitalOcean",
			GoMainLanguage:    true,
			Vacancies: domain.Vacancies{
				domain.Go: []string{
					"https://www.linkedin.com/jobs/view/4049491677/",
					"https://app.otta.com/jobs/bwHc4GXE",
				},
				domain.Rust:    nil,
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
			},
			ShortDescription: "Cloud service provider",
			Industries: []domain.Industry{
				domain.IndustryDevOps,
			},
		},

		// Tech | Canonical
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Canonical",
			Website: "https://canonical.com/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:       234280,
				Alias:    "canonical",
				Name:     "Canonical",
				Verified: true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "canonical",
				GoRepositoryCount: 83,
				Verified:          true,
			},
			BlindProfile:     domain.BlindProfile{},
			LevelsFyiProfile: domain.LevelsFyiProfile{},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Canonical-EI_IE230560.11,20.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Canonical-Reviews-E230560.htm",
				Verified:    true,
			},
			IndeedProfile:     domain.IndeedProfile{},
			OttaProfileSlug:   "canonical",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Rust: []string{
					"https://www.linkedin.com/jobs/view/3210061536/",
				},
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Golang Engineer",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/3736604544/",
							Date:             mustDate("2024-11-01"),
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
			ShortDescription: "Ubuntu publisher",
			DealroomURL:      "",
			CrunchbaseURL:    "",
			PitchbookURL:     "",
			YahooFinanceURL:  "",
			GoogleFinanceURL: "",
			YCombinatorURL:   "",
			Industries: []domain.Industry{
				domain.IndustryDevOps,
			},
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
				domain.Czechia,
			},
		},

		// Tech | SUSE
		{
			ID:   0,  // system
			Type: "", // system
			Name: "SUSE",
			URL:  "https://www.suse.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:       1455,
				Alias:    "suse",
				Name:     "SUSE",
				Verified: true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "suse",
				GoRepositoryCount: 69,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-SUSE-EI_IE466462.11,15.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/SUSE-Reviews-E466462.htm",
				Verified:    true,
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go: []string{
					"https://www.linkedin.com/jobs/view/4034618699/",
				},
				domain.Rust:    nil,
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
			},
			ShortDescription: "Secure enterprise open source solutions",
			Industries:       []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{
				domain.Czechia,
			},
		},

		// Tech | Kong
		{
			ID:   0,  // system
			Type: "", // system
			Name: "Kong",
			URL:  "https://konghq.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    278819,
				Alias: "konghq",
				Name:  "Kong Inc.",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "Kong",
				GoRepositoryCount: 51,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Kong-EI_IE801963.11,15.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Kong-Reviews-E801963.htm",
			},
			OttaProfileSlug:   "Kong",
			YouTubeChannelURL: "https://www.youtube.com/@KongInc",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go: []string{
					"https://app.otta.com/jobs/7-RRNkSk",
				},
				domain.Rust:    nil,
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
			},
			ShortDescription: "Cloud-native, platform-agnostic, scalable API Gateway",
			Industries: []domain.Industry{
				domain.IndustryDevOps,
			},
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
			},
		},

		// Tech | Exasol
		{
			ID:   0,  // system
			Type: "", // system
			Name: "Exasol",
			URL:  "https://www.exasol.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    1741694,
				Alias: "exasol-ag",
				Name:  "Exasol",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "exasol",
				GoRepositoryCount: 7,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Exasol-EI_IE968677.11,17.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Exasol-Reviews-E968677.htm",
			},
			OttaProfileSlug:   "Exasol",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go: []string{
					"https://app.otta.com/jobs/yfqng_As",
				},
				domain.Rust:    nil,
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
			},
			ShortDescription: "Analytics database",
			Industries: []domain.Industry{
				domain.IndustryDevOps,
			},
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
				domain.Czechia,
			},
		},

		// Tech | Palantir
		{
			ID:   0,  // system
			Type: "", // system
			Name: "Palantir",
			URL:  "https://www.palantir.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    20708,
				Alias: "palantir-technologies",
				Name:  "Palantir Technologies",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "palantir",
				GoRepositoryCount: 57,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Palantir-Technologies-EI_IE236375.11,32.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Palantir-Technologies-Reviews-E236375.htm",
			},
			OttaProfileSlug:   "Palantir",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go:      []string{},
				domain.Rust:    nil,
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
			},
			ShortDescription: "Platform for big data analytics",
			Industries: []domain.Industry{
				domain.IndustryDevOps,
			},
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
				domain.Czechia,
			},
		},

		// Tech | Buf
		// Template
		{
			ID:   0,  // system
			Type: "", // system
			Name: "Buf",
			URL:  "https://buf.build/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    68684262,
				Alias: "bufbuild",
				Name:  "Buf",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "bufbuild",
				GoRepositoryCount: 25,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "",
				ReviewsURL:  "",
				Verified:    false,
			},
			OttaProfileSlug:   "Buf",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go: []string{
					"https://app.otta.com/jobs/9KLKULQr",
				},
				domain.Rust:    nil,
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
			},
			ShortDescription: "Protobuf developer platform",
			Industries: []domain.Industry{
				domain.IndustryDevOps,
			},
		},

		// FinTech
		// FinTech | Stripe
		{
			ID:   0,  // system
			Type: "", // system
			Name: "Stripe",
			URL:  "https://stripe.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    2135371,
				Alias: "stripe",
				Name:  "Stripe",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "stripe",
				GoRepositoryCount: 12,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Stripe-EI_IE671932.11,17.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Stripe-Reviews-E671932.htm",
			},
			OttaProfileSlug:   "Stripe",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go: []string{
					"https://www.linkedin.com/jobs/view/4032599929/",
					"https://app.welcometothejungle.com/jobs/pnIwTyGO",
				},
				domain.Rust:    nil,
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
			},
			ShortDescription: "A financial infrastructure platform for businesses",
			Industries: []domain.Industry{
				domain.IndustryFinTech,
			},
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
				domain.Czechia,
			},
		},

		// FinTech | Wise
		{
			ID:   0,  // system
			Type: "", // system
			Name: "Wise",
			URL:  "https://wise.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    1769571,
				Alias: "wiseaccount",
				Name:  "Wise",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "transferwise",
				GoRepositoryCount: 35,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Wise-EI_IE637715.11,15.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Wise-Reviews-E637715.htm",
			},
			OttaProfileSlug:   "Wise",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go: []string{
					"https://app.otta.com/jobs/UVRcHQhe",
				},
				domain.Rust:    nil,
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
			},
			ShortDescription: "FinTech company",
			Industries: []domain.Industry{
				domain.IndustryFinTech,
			},
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
				domain.Czechia,
			},
		},

		// FinTech | American Express
		{
			ID:   0,  // system
			Type: "", // system
			Name: "American Express",
			URL:  "https://www.americanexpress.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    1277,
				Alias: "american-express",
				Name:  "American Express",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "americanexpress",
				GoRepositoryCount: 4,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-American-Express-EI_IE35.11,27.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/American-Express-Reviews-E35.htm",
			},
			OttaProfileSlug:   "American-Express",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go: []string{
					"https://www.linkedin.com/jobs/view/3926321626/",
				},
				domain.Rust:    nil,
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
			},
			ShortDescription: "Bank holding",
			Industries: []domain.Industry{
				domain.IndustryFinTech,
			},
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
				domain.Czechia,
			},
		},

		// FinTech | Mastercard
		{
			ID:   0,  // system
			Type: "", // system
			Name: "Mastercard",
			URL:  "https://www.mastercard.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    3015,
				Alias: "mastercard",
				Name:  "Mastercard",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "Mastercard",
				GoRepositoryCount: 3,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Mastercard-EI_IE3677.11,21.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Mastercard-Reviews-E3677.htm",
			},
			OttaProfileSlug:   "Mastercard",
			YouTubeChannelURL: "https://www.youtube.com/@MasterCard",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go: []string{
					"https://www.linkedin.com/jobs/view/3842471708/",
					"https://www.linkedin.com/jobs/view/3928338037/",
				},
				domain.Rust:    nil,
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
			},
			ShortDescription: "Payment card services corporation",
			Industries: []domain.Industry{
				domain.IndustryFinTech,
			},
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
				domain.Czechia,
			},
		},

		// FinTech | Morgan Stanley
		{
			ID:   0,  // system
			Type: "", // system
			Name: "Morgan Stanley",
			URL:  "https://www.morganstanley.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID: 0,
				IDs: []int{
					497017,  // Morgan Stanley
					1292145, // Graystone Consulting from Morgan Stanley
				},
				Alias: "morgan-stanley",
				Name:  "Morgan Stanley",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "MorganStanley",
				GoRepositoryCount: 0,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Morgan-Stanley-EI_IE2282.11,25.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Morgan-Stanley-Reviews-E2282.htm",
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go: []string{
					"https://www.linkedin.com/jobs/view/3952881485/",
				},
				domain.Rust:    nil,
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
			},
			ShortDescription: "",
			Industries: []domain.Industry{
				domain.IndustryFinTech,
			},
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
			},
		},

		// FinTech | Monzo
		{
			ID:   0,  // system
			Type: "", // system
			Name: "Monzo",
			URL:  "https://monzo.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    9471107,
				Alias: "monzo-bank",
				Name:  "Monzo Bank",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "monzo",
				GoRepositoryCount: 76,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Monzo-Bank-EI_IE1557148.11,21.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Monzo-Bank-Reviews-E1557148.htm",
			},
			OttaProfileSlug:   "Monzo",
			YouTubeChannelURL: "",
			GoMainLanguage:    true,
			Vacancies: domain.Vacancies{
				domain.Go: []string{
					"https://app.otta.com/jobs/ZwnXtENr",
				},
				domain.Rust:    nil,
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
			},
			ShortDescription: "Mobile banking",
			Industries: []domain.Industry{
				domain.IndustryFinTech,
			},
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
				domain.Czechia,
			},
		},

		// FinTech | Cynergy Bank
		{
			ID:   0,  // system
			Type: "", // system
			Name: "Cynergy Bank",
			URL:  "https://www.cynergybank.co.uk/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    18921842,
				Alias: "cynergy-bank",
				Name:  "Cynergy Bank",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "", // NOP
				GoRepositoryCount: 0,  // NOP
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Cynergy-Bank-EI_IE769090.11,23.htm",
				ReviewsURL:  "https://www.glassdoor.co.uk/Reviews/Cynergy-Bank-Reviews-E769090.htm",
			},
			OttaProfileSlug:   "", // NOP
			YouTubeChannelURL: "https://www.youtube.com/@cynergybank",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go: []string{
					"https://www.linkedin.com/jobs/view/3917097524/",
					"https://www.linkedin.com/jobs/view/3868356139/",
				},
				domain.Rust:    nil,
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
			},
			ShortDescription: "",
			Industries: []domain.Industry{
				domain.IndustryFinTech,
			},
		},

		// FinTech | Atom bank
		{
			ID:   0,  // system
			Type: "", // system
			Name: "Atom bank",
			URL:  "https://www.atombank.co.uk/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    5197064,
				Alias: "atom-bank",
				Name:  "Atom bank",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "atombank",
				GoRepositoryCount: 0,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Atom-Bank-EI_IE1354887.11,20.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Atom-Bank-Reviews-E1354887.htm",
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go: []string{
					"https://www.linkedin.com/jobs/view/3921514553/",
				},
				domain.Rust:    nil,
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
			},
			ShortDescription: "",
			Industries: []domain.Industry{
				domain.IndustryFinTech,
			},
		},

		// FinTech | Citi
		{
			ID:   0,  // system
			Type: "", // system
			Name: "Citi",
			URL:  "https://www.citi.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    11448,
				Alias: "citi",
				Name:  "Citi",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "Citi",
				GoRepositoryCount: 1,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/citi",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Citi-Reviews-E8843.htm",
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "https://www.youtube.com/@Citi",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go: []string{
					"https://www.linkedin.com/jobs/view/3912849197/",
				},
				domain.Rust:    nil,
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
			},
			ShortDescription: "",
			Industries: []domain.Industry{
				domain.IndustryFinTech,
			},
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
				domain.Czechia,
			},
		},

		// Internet
		// Internet Bitly
		{
			ID:   0,  // system
			Type: "", // system
			Name: "Bitly",
			URL:  "https://bitly.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    552285,
				Alias: "bitly",
				Name:  "Bitly",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "bitly",
				GoRepositoryCount: 11,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Bitly-EI_IE800313.11,16.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Bitly-Reviews-E800313.htm",
			},
			OttaProfileSlug:   "Bitly",
			YouTubeChannelURL: "",
			GoMainLanguage:    true, // https://bitly.com/blog/why-we-write-everything-in-go/
			Vacancies: domain.Vacancies{
				domain.Go:      []string{},
				domain.Rust:    nil,
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
			},
			ShortDescription: "URL shortening service and a link management platform",
			Industries: []domain.Industry{
				domain.IndustryMarTech,
			},
		},

		// Internet Cloudflare
		{
			ID:   0,  // system
			Type: "", // system
			Name: "Cloudflare",
			URL:  "https://www.cloudflare.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    407222,
				Alias: "cloudflare",
				Name:  "Cloudflare",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "cloudflare",
				GoRepositoryCount: 98,
				Verified:          false,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Cloudflare-EI_IE430862.11,21.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Cloudflare-Reviews-E430862.htm",
				Verified:    true,
			},
			OttaProfileSlug:   "Cloudflare-1",
			YouTubeChannelURL: "https://www.youtube.com/@cloudflare",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go: []string{
					"https://www.linkedin.com/jobs/view/4038005635/",
					"https://app.otta.com/jobs/RWRkVWNJ",
					"https://app.otta.com/jobs/B_RaZ6l5",
					"https://app.otta.com/jobs/e4CavNqx",
				},
				domain.Rust: []string{
					"https://www.linkedin.com/jobs/view/4041774900/",
					"https://www.linkedin.com/jobs/view/4038005635/",
					"https://startup.jobs/software-engineer-rust-cloudflare-3768119",
				},
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
			},
			ShortDescription: "Web application security and performance",
			Industries: []domain.Industry{
				domain.IndustryDevOps,
				domain.IndustryCyberSecurity,
			},
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
				domain.Czechia,
			},
		},

		// Internet Namecheap
		{
			ID:   0,  // system
			Type: "", // system
			Name: "Namecheap",
			URL:  "https://www.namecheap.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    486932,
				Alias: "namecheap-inc",
				Name:  "Namecheap, Inc",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "namecheap",
				GoRepositoryCount: 3,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Namecheap-EI_IE994113.11,20.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Namecheap-Reviews-E994113.htm",
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "https://www.youtube.com/@namecheap",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go: []string{
					"https://www.linkedin.com/jobs/view/3902607066/",
				},
				domain.Rust:    nil,
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
			},
			ShortDescription: "Domain and web-hosting company",
			Industries:       []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
			},
		},

		// Internet GoDaddy
		{
			ID:   0,  // system
			Type: "", // system
			Name: "GoDaddy",
			URL:  "https://www.godaddy.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    7846,
				Alias: "godaddy",
				Name:  "GoDaddy",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "godaddy",
				GoRepositoryCount: 6,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-GoDaddy-EI_IE35337.11,18.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/GoDaddy-Reviews-E35337.htm",
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go: []string{
					"https://www.linkedin.com/jobs/view/4011177772/",
				},
				domain.Rust:    nil,
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
			},
			ShortDescription: "Domain and web-hosting company",
			Industries:       []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
			},
		},

		// Security
		// Security 1Password
		{
			ID:   0,  // system
			Type: "", // system
			Name: "1Password",
			URL:  "https://1password.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    18648301,
				Alias: "1password",
				Name:  "1Password",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "1Password",
				GoRepositoryCount: 19,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-1Password-EI_IE2984143.11,20.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/1Password-Reviews-E2984143.htm",
			},
			OttaProfileSlug:   "1Password",
			YouTubeChannelURL: "https://www.youtube.com/@1PasswordVideos",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go: []string{
					"https://www.linkedin.com/jobs/view/3905310871/",
				},
				domain.Rust:    nil,
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
			},
			ShortDescription: "Service for users to store various passwords",
			Industries: []domain.Industry{
				domain.IndustryCyberSecurity,
			},
		},

		// Security Okta
		{
			ID:   0,  // system
			Type: "", // system
			Name: "Okta",
			URL:  "https://www.okta.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    926041,
				Alias: "okta-inc-",
				Name:  "Okta",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "okta",
				GoRepositoryCount: 5,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Okta-EI_IE444756.11,15.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Okta-Reviews-E444756.htm",
			},
			OttaProfileSlug:   "Okta",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go: []string{
					"https://app.otta.com/jobs/FtbF4gpB",
					"https://app.otta.com/jobs/QcMdOhR_",
				},
				domain.Rust:    nil,
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
			},
			ShortDescription: "Identity and access management company",
			Industries:       []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
				domain.Czechia,
			},
		},

		// Security Nord Security
		{
			ID:   0,  // system
			Type: "", // system
			Name: "Nord Security",
			URL:  "https://nordsecurity.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    64277258,
				Alias: "nordsecurity",
				Name:  "Nord Security",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "NordSecurity",
				GoRepositoryCount: 7, // Rust 22 repositories
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Nord-Security-EI_IE4015819.11,24.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Nord-Security-Reviews-E4015819.htm",
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "https://www.youtube.com/@nordsecurity",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go: []string{
					"https://www.linkedin.com/jobs/view/3890883832/",
					"https://www.linkedin.com/jobs/view/3892388207/",
				},
				domain.Rust:    nil,
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
			},
			ShortDescription: "",
			Industries: []domain.Industry{
				domain.IndustryCyberSecurity,
			},
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
				domain.Czechia,
			},
		},

		// Security Proton
		{
			ID:   0,  // system
			Type: "", // system
			Name: "Proton",
			URL:  "https://proton.me/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    5241679,
				Alias: "protonprivacy",
				Name:  "Proton",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "ProtonMail",
				GoRepositoryCount: 36,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Proton-privacy-by-default-EI_IE1405328.11,36.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Proton-privacy-by-default-Reviews-E1405328.htm",
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go: nil,
				domain.Rust: []string{
					"https://www.linkedin.com/jobs/view/3955798795/",
					"https://www.linkedin.com/jobs/view/4012882343/",
				},
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
			},
			ShortDescription: "",
			Industries:       []domain.Industry{},
		},

		// Security Fortinet
		{
			ID:   0,  // system
			Type: "", // system
			Name: "Fortinet",
			URL:  "https://www.fortinet.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    6460,
				Alias: "fortinet",
				Name:  "Fortinet",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "fortinet",
				GoRepositoryCount: 1,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Fortinet-EI_IE23128.11,19.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Fortinet-Reviews-E23128.htm",
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "https://www.youtube.com/@fortinet",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go: []string{
					"https://www.linkedin.com/jobs/view/3907970778/",
				},
				domain.Rust:    nil,
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
			},
			ShortDescription: "",
			Industries: []domain.Industry{
				domain.IndustryCyberSecurity,
			},
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
				domain.Czechia,
			},
		},

		// Security SentinelOne
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "SentinelOne",
			Website: "https://www.sentinelone.com/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:       2886771,
				Alias:    "sentinelone",
				Name:     "SentinelOne",
				Verified: true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "Sentinel-One",
				GoRepositoryCount: 2,
			},
			BlindProfile:     domain.BlindProfile{},
			LevelsFyiProfile: domain.LevelsFyiProfile{},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-SentinelOne-EI_IE1361978.11,22.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/SentinelOne-Reviews-E1361978.htm",
			},
			IndeedProfile:     domain.IndeedProfile{},
			OttaProfileSlug:   "SentinelOne",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go: []string{
					"https://app.otta.com/jobs/RN0fzEWC",
				},
				domain.Rust:    nil,
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
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
			DealroomURL:      "",
			CrunchbaseURL:    "",
			PitchbookURL:     "",
			YahooFinanceURL:  "",
			GoogleFinanceURL: "",
			YCombinatorURL:   "",
			Industries: []domain.Industry{
				domain.IndustryCyberSecurity,
			},
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
				domain.Czechia,
			},
		},

		// Famous
		// Famous Uber
		{
			ID:   0,  // system
			Type: "", // system
			Name: "Uber",
			URL:  "https://www.uber.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    1815218,
				Alias: "uber-com",
				Name:  "Uber",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "uber",
				GoRepositoryCount: 30,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Uber-EI_IE575263.11,15.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Uber-Reviews-E575263.htm",
			},
			OttaProfileSlug:   "Uber",
			YouTubeChannelURL: "",
			GoMainLanguage:    true,
			Vacancies: domain.Vacancies{
				domain.Go:      []string{},
				domain.Rust:    nil,
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
			},
			ShortDescription: "Transportation company",
			Industries:       []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
				domain.Czechia,
			},
		},

		// Famous Siemens
		{
			ID:   0,  // system
			Type: "", // system
			Name: "Siemens",
			URL:  "https://www.siemens.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    0,
				IDs:   []int{1043, 157241},
				Alias: "siemens",
				Name:  "Siemens",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "siemens",
				GoRepositoryCount: 19,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Siemens-EI_IE3510.11,18.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Siemens-Reviews-E3510.htm",
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "https://www.youtube.com/@SiemensKnowledgeHub",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go: []string{
					"https://www.linkedin.com/jobs/view/3901324430/",
				},
				domain.Rust: []string{
					"https://www.linkedin.com/jobs/view/4018413094/",
				},
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
			},
			ShortDescription: "",
			Industries:       []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
				domain.Czechia,
			},
		},

		// Famous Ericsson
		{
			ID:   0,  // system
			Type: "", // system
			Name: "Ericsson",
			URL:  "https://www.ericsson.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    1060,
				Alias: "ericsson",
				Name:  "Ericsson",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "Ericsson",
				GoRepositoryCount: 4,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Ericsson-Worldwide-EI_IE3472.11,29.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Ericsson-Worldwide-Reviews-E3472.htm",
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go: []string{
					"https://www.linkedin.com/jobs/view/4038208151/",
				},
				domain.Rust:    nil,
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
			},
			ShortDescription: "",
			Industries:       []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
			},
		},

		// Famous SoundCloud
		{
			ID:   0,  // system
			Type: "", // system
			Name: "SoundCloud",
			URL:  "https://soundcloud.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    200200,
				Alias: "soundcloud",
				Name:  "SoundCloud",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "soundcloud",
				GoRepositoryCount: 5,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-SoundCloud-EI_IE407066.11,21.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/SoundCloud-Reviews-E407066.htm",
			},
			OttaProfileSlug:   "SoundCloud",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go: []string{
					"https://app.otta.com/jobs/Q8WY5QwP",
				},
				domain.Rust:    nil,
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
			},
			ShortDescription: "Audio streaming service",
			Industries:       []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
				domain.Czechia,
			},
		},

		// Famous Spotify
		{
			ID:   0,  // system
			Type: "", // system
			Name: "Spotify",
			URL:  "https://www.lifeatspotify.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    207470,
				Alias: "spotify",
				Name:  "Spotify",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "spotify",
				GoRepositoryCount: 12,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/spotify",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Spotify-Reviews-E408251.htm",
			},
			OttaProfileSlug:   "Spotify",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go: []string{
					"https://app.welcometothejungle.com/jobs/BAoL6B1_",
				},
				domain.Rust:    nil,
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
			},
			ShortDescription: "Music streaming service",
			Industries: []domain.Industry{
				domain.IndustryEntertainment,
			},
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
				domain.Czechia,
			},
		},

		// Famous DoorDash
		{
			ID:   0,  // system
			Type: "", // system
			Name: "DoorDash",
			URL:  "https://doordash.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    3205573,
				Alias: "doordash",
				Name:  "DoorDash",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "doordash",
				GoRepositoryCount: 7,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-DoorDash-EI_IE813073.11,19.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/DoorDash-Reviews-E813073.htm",
			},
			OttaProfileSlug:   "DoorDash",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go:      []string{},
				domain.Rust:    nil,
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
			},
			ShortDescription: "Food delivery",
			Industries:       []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
				domain.Czechia,
			},
		},

		{
			ID:   0,  // system
			Type: "", // system
			Name: "Just Eat Takeaway.com",
			URL:  "https://careers.justeattakeaway.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID: 0,
				IDs: []int{
					103531,  // Just Eat Takeaway.com
					2708300, //  SkipTheDishes
				},
				Alias: "just-eat-takeaway-com",
				Name:  "Just Eat Takeaway.com",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "justeattakeaway",
				GoRepositoryCount: 0,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Just-Eat-Takeaway-com-EI_IE490124.11,32.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Just-Eat-Takeaway-com-Reviews-E490124.htm",
			},
			OttaProfileSlug:   "Just-Eat",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go: []string{
					"https://www.linkedin.com/jobs/view/4039505901/",
				},
				domain.Rust:    nil,
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
			},
			ShortDescription: "Food delivery",
			Industries:       []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
			},
		},

		// Famous Sixt
		{
			ID:   0,
			Type: "",
			Name: "Sixt",
			URL:  "https://www.sixt.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    17120,
				Alias: "sixt",
				Name:  "SIXT",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "Sixt",
				GoRepositoryCount: 5,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Sixt-EI_IE10875.11,15.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Sixt-Reviews-E10875.htm",
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Vacancies:         domain.Vacancies{},
			ShortDescription:  "Car Rental service",
			Industries:        []domain.Industry{},
		},

		// Famous Motorola Solutions
		{
			ID:   0,  // system
			Type: "", // system
			Name: "Motorola Solutions",
			URL:  "https://www.motorolasolutions.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    1771432,
				Alias: "motorolasolutions",
				Name:  "Motorola Solutions",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "",
				GoRepositoryCount: 0,
				Verified:          false,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Motorola-Solutions-EI_IE427189.11,29.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Motorola-Solutions-Reviews-E427189.htm",
			},
			OttaProfileSlug:   "Motorola-Solutions",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go: []string{
					"https://app.otta.com/jobs/L8SSyDNH",
				},
				domain.Rust:    nil,
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
			},
			ShortDescription: "American multinational telecommunications company",
			Industries:       []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
				domain.Czechia,
			},
		},

		// Famous Samsung
		{
			ID:   0,  // system
			Type: "", // system
			Name: "Samsung",
			URL:  "https://www.samsung.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    0,
				IDs:   []int{1753, 13749561, 3270132, 77752302, 3812597, 3220515, 10633911, 1447520, 3238801, 2711168, 10261221, 87464967, 5089912, 9500165, 10218505, 895705, 78467539, 10332849, 10660176, 79815984, 3290134, 76958044, 85881048, 27127559, 9278177, 37470202, 22316801, 81590082, 11229641, 15213487, 68478415, 5552815, 14472582},
				Alias: "samsung-electronics",
				Name:  "Samsung Electronics",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "Samsung",
				GoRepositoryCount: 0,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Samsung-Electronics-EI_IE3363.11,30.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Samsung-Electronics-Reviews-E3363.htm",
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "https://www.youtube.com/@Samsung",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go: []string{
					"https://www.linkedin.com/jobs/view/3919334242/",
					"https://www.linkedin.com/jobs/view/3882780776/",
					"https://www.linkedin.com/jobs/view/3864617797/",
				},
				domain.Rust:    nil,
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
			},
			ShortDescription: "Electronics corporation",
			Industries:       []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
				domain.Czechia,
			},
		},

		// Famous Salesforge
		{
			ID:   0,  // system
			Type: "", // system
			Name: "Salesforge",
			URL:  "https://www.salesforge.ai/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    91706116,
				Alias: "salesforgeai",
				Name:  "Salesforge",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "",
				GoRepositoryCount: 0,
				Verified:          false,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "",
				ReviewsURL:  "",
				Verified:    false,
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "https://www.youtube.com/@salesforge",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go: []string{
					"https://www.linkedin.com/jobs/view/3912551615/",
				},
				domain.Rust:    nil,
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
			},
			ShortDescription: "All-inclusive sales app",
			Industries:       []domain.Industry{},
		},

		// Some
		// Some | Careem
		{
			ID:   0,  // system
			Type: "", // system
			Name: "Careem",
			URL:  "https://www.careem.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    2852511,
				Alias: "careem",
				Name:  "Careem",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "careem",
				GoRepositoryCount: 3,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Careem-EI_IE1438731.11,17.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Careem-Reviews-E1438731.htm",
			},
			OttaProfileSlug:   "Careem",
			YouTubeChannelURL: "",
			GoMainLanguage:    true,
			Vacancies: domain.Vacancies{
				domain.Go:      []string{},
				domain.Rust:    nil,
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
			},
			ShortDescription: "SuperApp",
			Industries:       []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
				domain.Czechia,
			},
		},

		// Some | Dailymotion
		{
			ID:   0,  // system
			Type: "", // system
			Name: "Dailymotion",
			URL:  "https://www.dailymotion.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    24411,
				Alias: "dailymotion",
				Name:  "Dailymotion",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "dailymotion",
				GoRepositoryCount: 18,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Dailymotion-EI_IE372880.11,22.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Dailymotion-Reviews-E372880.htm",
			},
			OttaProfileSlug:   "Dailymotion",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go:      []string{},
				domain.Rust:    nil,
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
			},
			ShortDescription: "Video hosting and sharing platforms",
			Industries:       []domain.Industry{},
		},

		// Some | Stream
		{
			ID:   0,  // system
			Type: "", // system
			Name: "Stream",
			URL:  "https://getstream.io/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    5338728,
				Alias: "getstream",
				Name:  "Stream",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "GetStream",
				GoRepositoryCount: 32,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Stream-CO-EI_IE1703813.11,20.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Stream-CO-Reviews-E1703813.htm",
			},
			OttaProfileSlug:   "Stream",
			YouTubeChannelURL: "https://www.youtube.com/@streamdevelopers",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go: []string{
					"https://www.linkedin.com/jobs/view/3989947743/",
					"https://www.linkedin.com/jobs/view/3873637399/",
				},
				domain.Rust:    nil,
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
			},
			ShortDescription: "Scalable and fast APIs for building social networks and apps",
			Industries:       []domain.Industry{},
		},

		// Some | Workato
		{
			ID:   0,  // system
			Type: "", // system
			Name: "Workato",
			URL:  "https://www.workato.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    3675685,
				Alias: "workato",
				Name:  "Workato",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "workato",
				GoRepositoryCount: 4,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Workato-EI_IE1333040.11,18.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Workato-Reviews-E1333040.htm",
			},
			OttaProfileSlug:   "Workato",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go:      []string{},
				domain.Rust:    nil,
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
			},
			ShortDescription: "",
			Industries:       []domain.Industry{},
		},

		// Some | Form3
		{
			ID:   0,  // system
			Type: "", // system
			Name: "Form3",
			URL:  "https://www.form3.tech/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    15156804,
				Alias: "form3-financial-cloud",
				Name:  "Form3",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "form3tech-oss",
				GoRepositoryCount: 43,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Form3-EI_IE2008415.11,16.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Form3-Reviews-E2008415.htm",
			},
			OttaProfileSlug:   "Form3",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go: []string{
					"https://relocate.me/the-united-kingdom/london/form3/senior-software-engineer-go-7646",
				},
				domain.Rust:    nil,
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
			},
			ShortDescription: "Financial technology company focused on money transfers",
			Industries: []domain.Industry{
				domain.IndustryFinTech,
			},
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
				domain.Czechia,
			},
		},

		// Some | Assertive Yield
		{
			ID:   0,  // system
			Type: "", // system
			Name: "Assertive Yield",
			URL:  "https://www.assertiveyield.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    76806664,
				Alias: "assertive-yield",
				Name:  "Assertive Yield B.V.",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "Assertive-Yield",
				GoRepositoryCount: 1,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "",
				ReviewsURL:  "",
				Verified:    false,
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "",
			GoMainLanguage:    true,
			Vacancies: domain.Vacancies{
				domain.Go:      []string{},
				domain.Rust:    nil,
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
			},
			ShortDescription: "Marketing services company",
			Industries: []domain.Industry{
				domain.IndustryAdTech,
				domain.IndustryMarTech,
			},
		},

		// Some | Splunk
		{
			ID:   0,  // system
			Type: "", // system
			Name: "Splunk",
			URL:  "https://www.splunk.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    20226,
				Alias: "splunk",
				Name:  "Splunk",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "splunk",
				GoRepositoryCount: 43,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Splunk-EI_IE117313.11,17.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Splunk-Reviews-E117313.htm",
			},
			OttaProfileSlug:   "Splunk",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go: []string{
					"https://justjoin.it/offers/splunk-senior-software-engineer-backend-go-olkusz",
				},
				domain.Rust:    nil,
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
			},
			ShortDescription: "Big Data platform",
			Industries: []domain.Industry{
				domain.IndustryBigData,
			},
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
				domain.Czechia,
			},
		},

		// Some | 90POE
		{
			ID:   0,  // system
			Type: "", // system
			Name: "90POE",
			URL:  "https://www.90poe.io/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    18466590,
				Alias: "90poe",
				Name:  "Ninety Percent of Everything (90POE)",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "90poe",
				GoRepositoryCount: 28,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-90-POE-EI_IE3898368.11,17.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/90-POE-Reviews-E3898368.htm",
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "",
			GoMainLanguage:    true,
			Vacancies: domain.Vacancies{
				domain.Go:      []string{},
				domain.Rust:    nil,
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
			},
			ShortDescription: "Digital platform for maritime transporting",
			Industries:       []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
				domain.Czechia,
			},
		},

		// Some | HelloFresh
		{
			ID:   0,  // system
			Type: "", // system
			Name: "HelloFresh",
			URL:  "https://www.hellofresh.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    2454643,
				Alias: "hellofresh",
				Name:  "HelloFresh",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "hellofresh",
				GoRepositoryCount: 9,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-HelloFresh-EI_IE998728.11,21.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/HelloFresh-Reviews-E998728.htm",
			},
			OttaProfileSlug:   "HelloFresh",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go:      []string{},
				domain.Rust:    nil,
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
			},
			ShortDescription: "",
			Industries:       []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
				domain.Czechia,
			},
		},

		// Some | AUTODOC
		{
			ID:   0,  // system
			Type: "", // system
			Name: "AUTODOC",
			URL:  "https://autodoc.group/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    7703911,
				Alias: "autodoc",
				Name:  "AUTODOC",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "", // unknown
				GoRepositoryCount: 0,  // unknown
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-AUTODOC-EI_IE2179604.11,18.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/AUTODOC-Reviews-E2179604.htm",
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go:      []string{},
				domain.Rust:    nil,
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
			},
			ShortDescription: "",
			Industries:       []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
			},
		},

		// Some | Gymondo
		{
			ID:   0,  // system
			Type: "", // system
			Name: "Gymondo",
			URL:  "https://www.gymondo.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    5233814,
				Alias: "gymondo-gmbh",
				Name:  "Gymondo",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "Gymondo-git",
				GoRepositoryCount: 3,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Gymondo-EI_IE1344198.11,18.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Gymondo-Reviews-E1344198.htm",
			},
			OttaProfileSlug:   "Gymondo",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go:      []string{},
				domain.Rust:    nil,
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
			},
			ShortDescription: "",
			Industries:       []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
			},
		},

		// Some | Delivery Hero
		{
			ID:   0,  // system
			Type: "", // system
			Name: "Delivery Hero",
			URL:  "https://www.deliveryhero.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    2393200,
				Alias: "delivery-hero-se",
				Name:  "Delivery Hero",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "deliveryhero",
				GoRepositoryCount: 11,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Delivery-Hero-EI_IE504556.11,24.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Delivery-Hero-Reviews-E504556.htm",
			},
			OttaProfileSlug:   "Delivery-Hero",
			YouTubeChannelURL: "https://www.youtube.com/@DeliveryHeroDH",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go: []string{
					"https://www.linkedin.com/jobs/view/3908507037/",
				},
				domain.Rust:    nil,
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
			},
			ShortDescription: "",
			Industries:       []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
				domain.Czechia,
			},
		},

		// Some | Weaviate
		{
			ID:   0,  // system
			Type: "", // system
			Name: "Weaviate",
			URL:  "https://weaviate.io/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    11702022,
				Alias: "weaviate-io",
				Name:  "Weaviate",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "weaviate",
				GoRepositoryCount: 13,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Weaviate-EI_IE7479527.11,19.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Weaviate-Reviews-E7479527.htm",
			},
			OttaProfileSlug:   "Weaviate",
			YouTubeChannelURL: "",
			GoMainLanguage:    true,
			Vacancies: domain.Vacancies{
				domain.Go:      []string{},
				domain.Rust:    nil,
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
			},
			ShortDescription: "",
			Industries:       []domain.Industry{},
		},

		// Some | Fubo
		{
			ID:   0,  // system
			Type: "", // system
			Name: "Fubo",
			URL:  "https://www.fubo.tv/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    5316737,
				Alias: "fubotv",
				Name:  "Fubo",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "fubotv",
				GoRepositoryCount: 6,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-FuboTV-EI_IE1006878.11,17.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/FuboTV-Reviews-E1006878.htm",
			},
			OttaProfileSlug:   "fubo",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go:      []string{},
				domain.Rust:    nil,
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
			},
			ShortDescription: "",
			Industries:       []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
			},
		},

		// Some | Yassir
		{
			ID:   0,  // system
			Type: "", // system
			Name: "Yassir",
			URL:  "https://yassir.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    19069709,
				Alias: "yassir",
				Name:  "Yassir",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "YAtechnologies",
				GoRepositoryCount: 0, // 0
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-YASSIR-EI_IE2601333.11,17.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/YASSIR-Reviews-E2601333.htm",
			},
			OttaProfileSlug:   "Yassir",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go:      []string{},
				domain.Rust:    nil,
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
			},
			ShortDescription: "",
			Industries:       []domain.Industry{},
		},

		// Some | Vio.com
		{
			ID:   0,  // system
			Type: "", // system
			Name: "Vio.com",
			URL:  "https://www.vio.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    1192098,
				Alias: "viodotcom",
				Name:  "Vio.com",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "viodotcom",
				GoRepositoryCount: 2,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Vio-com-EI_IE940798.11,18.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Vio-com-Reviews-E940798.htm",
			},
			OttaProfileSlug:   "Vio-com",
			YouTubeChannelURL: "",
			GoMainLanguage:    false, // https://www.linkedin.com/jobs/view/3819771736/
			Vacancies: domain.Vacancies{
				domain.Go:      []string{},
				domain.Rust:    nil,
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
			},
			ShortDescription: "",
			Industries:       []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
				domain.Czechia,
			},
		},

		// Some | Vodeno
		{
			ID:   0,  // system
			Type: "", // system
			Name: "Vodeno",
			URL:  "https://vodeno.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    19016391,
				Alias: "vodeno",
				Name:  "Vodeno",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "vodeno",
				GoRepositoryCount: 0,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Vodeno-EI_IE2877672.11,17.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Vodeno-Reviews-E2877672.htm",
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go: []string{
					"https://justjoin.it/offers/vodeno-go-developer-kielce-358668",
				},
				domain.Rust:    nil,
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
			},
			ShortDescription: "",
			Industries:       []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{
				domain.Czechia,
			},
		},

		// Some | Utility Warehouse
		{
			ID:   0,  // system
			Type: "", // system
			Name: "Utility Warehouse",
			URL:  "https://uw.co.uk/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    457903,
				Alias: "utilitywarehouse",
				Name:  "Utility Warehouse",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "utilitywarehouse",
				GoRepositoryCount: 85,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Utility-Warehouse-EI_IE512935.11,28.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Utility-Warehouse-Reviews-E512935.htm",
			},
			OttaProfileSlug:   "Utility-Warehouse",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go:      []string{},
				domain.Rust:    nil,
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
			},
			ShortDescription: "",
			Industries:       []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{
				domain.Czechia,
			},
		},

		// Some | Codenotary
		{
			ID:   0,  // system
			Type: "", // system
			Name: "Codenotary",
			URL:  "https://codenotary.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    35523736,
				Alias: "codenotary",
				Name:  "Codenotary",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "codenotary",
				GoRepositoryCount: 23,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-CodeNotary-EI_IE3677292.11,21.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/CodeNotary-Reviews-E3677292.htm",
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "",
			GoMainLanguage:    true,
			Vacancies: domain.Vacancies{
				domain.Go:      []string{},
				domain.Rust:    nil,
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
			},
			ShortDescription: "",
			Industries:       []domain.Industry{},
		},

		// Some | Audigent
		{
			ID:   0,
			Type: "",
			Name: "Audigent",
			URL:  "https://audigent.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    10642467,
				Alias: "audigent",
				Name:  "Audigent",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "AuDigent",
				GoRepositoryCount: 2,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Audigent-EI_IE5815298.11,19.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Audigent-Reviews-E5815298.htm",
			},
			OttaProfileSlug:   "Audigent",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Vacancies:         domain.Vacancies{},
			ShortDescription:  "",
			Industries:        []domain.Industry{},
		},

		// Some | runZero
		{
			ID:   0,  // system
			Type: "", // system
			Name: "runZero",
			URL:  "https://www.runzero.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    33274038,
				Alias: "runzero",
				Name:  "runZero",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "runZeroInc",
				GoRepositoryCount: 7,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-runZero-EI_IE7717209.11,18.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/runZero-Reviews-E7717209.htm",
			},
			OttaProfileSlug:   "runZero",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go:      []string{},
				domain.Rust:    nil,
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
			},
			ShortDescription: "",
			Industries:       []domain.Industry{},
		},

		// Some | Tyk
		{
			ID:   0,  // system
			Type: "", // system
			Name: "Tyk",
			URL:  "https://tyk.io/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    10775050,
				Alias: "tyk",
				Name:  "Tyk",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "TykTechnologies",
				GoRepositoryCount: 79,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Tyk-EI_IE2321465.11,14.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Tyk-Reviews-E2321465.htm",
			},
			OttaProfileSlug:   "Tyk",
			YouTubeChannelURL: "",
			GoMainLanguage:    true,
			Vacancies: domain.Vacancies{
				domain.Go: []string{
					"https://app.welcometothejungle.com/jobs/CwkAD14C",
				},
				domain.Rust:    nil,
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
			},
			ShortDescription: "Open-source API gateway & API management platform",
			Industries: []domain.Industry{
				domain.IndustryDevOps,
			},
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
			},
		},

		// Some | OpenTag
		{
			ID:   0,  // system
			Type: "", // system
			Name: "OpenTag",
			URL:  "https://theopentag.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    20565935,
				Alias: "theopentag",
				Name:  "OpenTag",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "OpenTagOS",
				GoRepositoryCount: 3,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-OpenTag-EI_IE3310869.11,18.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/OpenTag-Reviews-E3310869.htm",
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go:      []string{},
				domain.Rust:    nil,
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
			},
			ShortDescription: "",
			Industries:       []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
			},
		},

		// Some | Oxla
		{
			ID:   0,  // system
			Type: "", // system
			Name: "Oxla",
			URL:  "https://oxla.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    79378182,
				Alias: "oxla",
				Name:  "Oxla",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "", // NOP
				GoRepositoryCount: 0,  // NOP
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "",
				ReviewsURL:  "",
				Verified:    false,
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go: []string{
					"https://justjoin.it/offers/oxla-golang-developer-gdansk-362959",
				},
				domain.Rust:    nil,
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
			},
			ShortDescription: "",
			Industries:       []domain.Industry{},
		},

		// Some | Lightspeed
		{
			ID:   0,  // system
			Type: "", // system
			Name: "Lightspeed",
			URL:  "https://www.lightspeedhq.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:       1557218,
				Alias:    "lightspeedcommerce",
				Name:     "Lightspeed Commerce",
				Verified: true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "lightspeed",
				GoRepositoryCount: 0,
				Verified:          true,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Lightspeed-EI_IE648762.11,21.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Lightspeed-Reviews-E648762.htm",
				Verified:    true,
			},
			OttaProfileSlug:   "Lightspeed",
			YouTubeChannelURL: "https://www.youtube.com/channel/UCqOEKwLpolZBcj4LfU3R0Fg",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go: []string{
					"https://www.linkedin.com/jobs/view/4040512029/", // Develop and maintain our web applications using Golang and React
					"https://app.otta.com/jobs/RnXM1YTv",
				},
				domain.Rust:    nil,
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
			},
			ShortDescription: "",
			Industries:       []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
				domain.Czechia,
			},
		},

		// Some | Squarespace
		{
			ID:   0,  // system
			Type: "", // system
			Name: "Squarespace",
			URL:  "https://www.squarespace.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:       265314,
				Alias:    "squarespace",
				Name:     "Squarespace",
				Verified: true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "squarespace",
				GoRepositoryCount: 2,
				Verified:          true,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Squarespace-EI_IE466343.11,22.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Squarespace-Reviews-E466343.htm",
				Verified:    true,
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "https://www.youtube.com/squarespace",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go: []string{
					"https://www.linkedin.com/jobs/view/4022342302/",
					"https://app.otta.com/jobs/qAnITcbo",
				},
				domain.Rust:    nil,
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
			},
			ShortDescription: "",
			Industries:       []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
			},
		},

		// Some | Curve
		{
			ID:   0,  // system
			Type: "", // system
			Name: "Curve",
			URL:  "https://curve.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    10023464,
				Alias: "curve-ltd",
				Name:  "Curve",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "", // NOP
				GoRepositoryCount: 0,  // NOP
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Curve-EI_IE1739754.11,16.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Curve-Reviews-E1739754.htm",
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go: []string{
					"https://www.linkedin.com/jobs/view/3872933701/",
				},
				domain.Rust:    nil,
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
			},
			ShortDescription:          "",
			Industries:                []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{
				//
			},
		},

		// Some | Tradevest
		{
			ID:   0,  // system
			Type: "", // system
			Name: "Tradevest",
			URL:  "https://www.tradevest.ai/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    92827682,
				Alias: "tradevestgmbh",
				Name:  "Tradevest",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "", // NOP
				GoRepositoryCount: 0,  // NOP
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "",
				ReviewsURL:  "",
				Verified:    false,
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go: []string{
					"https://justjoin.it/offers/tv-development-gmbh-senior-backend-developer",
				},
				domain.Rust:    nil,
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
			},
			ShortDescription: "",
			Industries:       []domain.Industry{},
		},

		// Some | Woolsocks
		{
			ID:   0,  // system
			Type: "", // system
			Name: "Woolsocks",
			URL:  "https://woolsocks.eu/",
			LinkedInProfile: domain.LinkedInProfile{

				ID:    79728837,
				Alias: "woolsocks",
				Name:  "Woolsocks",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "amsterdam-platform-creation",
				GoRepositoryCount: 3,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "",
				ReviewsURL:  "",
				Verified:    false,
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go: []string{
					"https://www.linkedin.com/jobs/view/3869628047/",
				},
				domain.Rust:    nil,
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
			},
			ShortDescription: "",
			Industries:       []domain.Industry{},
		},

		// Some | Applied Systems Canada
		{
			ID:   0,  // system
			Type: "", // system
			Name: "Applied Systems Canada",
			URL:  "https://www.appliedsystems.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    908801,
				Alias: "applied-systems-canada",
				Name:  "Applied Systems Canada",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "",
				GoRepositoryCount: 0,
				Verified:          false,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Applied-Systems-EI_IE8534.11,26.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Applied-Systems-Reviews-E8534.htm",
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go: []string{
					"https://www.linkedin.com/jobs/view/3879152828/",
				},
				domain.Rust:    nil,
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
			},
			ShortDescription: "",
			Industries:       []domain.Industry{},
		},

		// Some | Autodesk
		{
			ID:   0,  // system
			Type: "", // system
			Name: "Autodesk",
			URL:  "https://www.autodesk.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    1879,
				Alias: "autodesk",
				Name:  "Autodesk",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "Autodesk",
				GoRepositoryCount: 5,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Autodesk-EI_IE1155.11,19.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Autodesk-Reviews-E1155.htm",
			},
			OttaProfileSlug:   "Autodesk",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go: []string{
					"https://www.linkedin.com/jobs/view/3843045776/",
				},
				domain.Rust:    nil,
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
			},
			ShortDescription: "",
			Industries:       []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
				domain.Czechia,
			},
		},

		// Some | Vonage
		{
			ID:   0,  // system
			Type: "", // system
			Name: "Vonage",
			URL:  "https://www.vonage.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    0,
				IDs:   []int{5028, 1345545, 66428, 76778, 2102},
				Alias: "vonage",
				Name:  "Vonage",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "Vonage",
				GoRepositoryCount: 4,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Vonage-EI_IE23019.11,17.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Vonage-Reviews-E23019.htm",
			},
			OttaProfileSlug:   "Vonage",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go: []string{
					"https://www.linkedin.com/jobs/view/3818119706/",
					"https://app.otta.com/jobs/0bz2Xy75",
				},
				domain.Rust:    nil,
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
			},
			ShortDescription: "",
			Industries:       []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
			},
		},

		// Some | OpenWeb
		{
			ID:   0,  // system
			Type: "", // system
			Name: "OpenWeb",
			URL:  "https://www.openweb.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    2506872,
				Alias: "openwebhq",
				Name:  "OpenWeb",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "",
				GoRepositoryCount: 0,
				Verified:          false,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-OpenWeb-EI_IE1675932.11,18.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/OpenWeb-Reviews-E1675932.htm",
			},
			OttaProfileSlug:   "OpenWeb",
			YouTubeChannelURL: "https://www.youtube.com/@openwebhq",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go: []string{
					"https://app.otta.com/jobs/x-Xm2wSF",
				},
				domain.Rust:    nil,
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
			},
			ShortDescription: "",
			Industries:       []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
			},
		},

		// Some | Arenko
		{
			ID:   0,  // system
			Type: "", // system
			Name: "Arenko",
			URL:  "https://arenko.group/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    10288973,
				Alias: "arenko-cleantech",
				Name:  "Arenko",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "arenko-group",
				GoRepositoryCount: 2,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Arenko-Group-EI_IE4554199.11,23.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Arenko-Group-Reviews-E4554199.htm",
			},
			OttaProfileSlug:   "Arenko-Group",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go: []string{
					"https://www.linkedin.com/jobs/view/4037978333/",
					"https://www.linkedin.com/jobs/view/3889419131/",
					"https://app.otta.com/jobs/tvHMOMCd",
					"https://app.otta.com/jobs/hKvg-QNr",
				},
				domain.Rust:    nil,
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
			},
			ShortDescription:          "Technology provider enabling the clean energy transition",
			Industries:                []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{},
		},

		// Some | Xata
		{
			ID:   0,  // system
			Type: "", // system
			Name: "Xata",
			URL:  "https://xata.io/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    69560619,
				Alias: "xataio",
				Name:  "Xata.io",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "xataio",
				GoRepositoryCount: 3,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "",
				ReviewsURL:  "",
				Verified:    false,
			},
			OttaProfileSlug:   "Xata",
			YouTubeChannelURL: "https://www.youtube.com/@xataio",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go: []string{
					"https://www.linkedin.com/jobs/view/3887215661/",
				},
				domain.Rust:    nil,
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
			},
			ShortDescription: "",
			Industries:       []domain.Industry{},
		},

		// Some | Dojo
		{
			ID:   0,  // system
			Type: "", // system
			Name: "Dojo",
			URL:  "https://dojo.careers/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:       42391390,
				Alias:    "dojo-tech",
				Name:     "Dojo",
				Verified: true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "",
				GoRepositoryCount: 0,
				Verified:          false,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Dojo-EI_IE687810.11,15.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Dojo-Reviews-E687810.htm",
				Verified:    true,
			},
			OttaProfileSlug:   "Dojo",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go: []string{
					"https://www.linkedin.com/jobs/view/3926951138/",
					"https://www.linkedin.com/jobs/view/3887771635/",
					"https://app.otta.com/jobs/kdaq4_rK",
				},
				domain.Rust:    nil,
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
			},
			ShortDescription: "",
			Industries:       []domain.Industry{},
		},

		// Some | Unnax
		{
			ID:   0,  // system
			Type: "", // system
			Name: "Unnax",
			URL:  "https://www.unnax.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    10295665,
				Alias: "unnax-emi",
				Name:  "Unnax",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "unnax",
				GoRepositoryCount: 1,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Unnax-EI_IE2108310.11,16.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Unnax-Reviews-E2108310.htm",
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "https://www.youtube.com/@unnax-emi",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go: []string{
					"https://www.linkedin.com/jobs/view/3890431726/",
				},
				domain.Rust:    nil,
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
			},
			ShortDescription: "",
			Industries:       []domain.Industry{},
		},

		// Some | AB Tasty
		{
			ID:   0,  // system
			Type: "", // system
			Name: "AB Tasty",
			URL:  "https://www.abtasty.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    1885711,
				Alias: "ab-tasty",
				Name:  "AB Tasty",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "",
				GoRepositoryCount: 0,
				Verified:          false,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-AB-Tasty-EI_IE1309242.11,19.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/AB-Tasty-Reviews-E1309242.htm",
			},
			OttaProfileSlug:   "AB-Tasty",
			YouTubeChannelURL: "https://www.youtube.com/@abtasty",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go: []string{
					"https://www.linkedin.com/jobs/view/3872953963/",
					"https://app.otta.com/jobs/eYKE7Ta_",
				},
				domain.Rust:    nil,
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
			},
			ShortDescription: "",
			Industries:       []domain.Industry{},
		},

		// Some | Firebolt
		{
			ID:   0,  // system
			Type: "", // system
			Name: "Firebolt",
			URL:  "https://www.firebolt.io/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    40719957,
				Alias: "firebolt",
				Name:  "Firebolt",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "firebolt-db",
				GoRepositoryCount: 2,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Firebolt-EI_IE5001853.11,19.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Firebolt-Reviews-E5001853.htm",
			},
			OttaProfileSlug:   "firebolt",
			YouTubeChannelURL: "https://www.youtube.com/@FireboltHQ",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go: []string{
					"https://www.linkedin.com/jobs/view/3855486951/",
				},
				domain.Rust:    nil,
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
			},
			ShortDescription: "",
			Industries:       []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
			},
		},

		// Some | Nine
		{
			ID:   0,  // system
			Type: "", // system
			Name: "Nine",
			URL:  "https://www.nineforbrands.com.au/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    14933,
				Alias: "nine-entertainment-co.",
				Name:  "Nine",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "",
				GoRepositoryCount: 0,
				Verified:          false,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Nine-Entertainment-EI_IE229827.11,29.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Nine-Entertainment-Reviews-E229827.htm",
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go: []string{
					"https://www.linkedin.com/jobs/view/3890491488/",
				},
				domain.Rust:    nil,
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
			},
			ShortDescription: "",
			Industries:       []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{
				domain.Czechia,
			},
		},

		// Some | Isovalent
		{
			ID:   0,  // system
			Type: "", // system
			Name: "Isovalent",
			URL:  "https://isovalent.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    34714411,
				Alias: "isovalent",
				Name:  "Isovalent",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "isovalent",
				GoRepositoryCount: 10,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Isovalent-EI_IE3180689.11,20.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Isovalent-Reviews-E3180689.htm",
			},
			OttaProfileSlug:   "Isovalent",
			YouTubeChannelURL: "https://www.youtube.com/@isovalent",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go: []string{
					"https://www.linkedin.com/jobs/view/3893426061/",
				},
				domain.Rust:    nil,
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
			},
			ShortDescription: "",
			Industries:       []domain.Industry{},
		},

		// Some | ABC Fitness
		{
			ID:   0,  // system
			Type: "", // system
			Name: "ABC Fitness",
			URL:  "https://abcfitness.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    110372,
				Alias: "abc-fitness",
				Name:  "ABC Fitness",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "",
				GoRepositoryCount: 0,
				Verified:          false,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-ABC-Fitness-EI_IE28305.11,22.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/ABC-Fitness-Reviews-E28305.htm",
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "https://www.youtube.com/@ABC-Fitness",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go: []string{
					"https://www.linkedin.com/jobs/view/3890939506/",
				},
				domain.Rust:    nil,
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
			},
			ShortDescription: "",
			Industries:       []domain.Industry{},
		},

		// Some | Device42
		{
			ID:   0,  // system
			Type: "", // system
			Name: "Device42",
			URL:  "https://www.device42.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    2345405,
				Alias: "device42",
				Name:  "Device42",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "device42",
				GoRepositoryCount: 0,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Device42-EI_IE1705087.11,19.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Device42-Reviews-E1705087.htm",
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "https://www.youtube.com/@Device42",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go: []string{
					"https://www.linkedin.com/jobs/view/3887639999/",
				},
				domain.Rust:    nil,
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
			},
			ShortDescription: "",
			Industries:       []domain.Industry{},
		},

		// Some | Acronis
		{
			ID:   0,  // system
			Type: "", // system
			Name: "Acronis",
			URL:  "https://www.acronis.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    13179,
				Alias: "acronis",
				Name:  "Acronis",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "acronis",
				GoRepositoryCount: 3,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Acronis-EI_IE152824.11,18.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Acronis-Reviews-E152824.htm",
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "https://www.youtube.com/@Acronis",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go: []string{
					"https://www.linkedin.com/jobs/view/3888974596/",
				},
				domain.Rust:    nil,
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
			},
			ShortDescription: "",
			Industries:       []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
			},
		},

		// Some | Gcore
		{
			ID:   0,  // system
			Type: "", // system
			Name: "Gcore",
			URL:  "https://gcore.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    10351246,
				Alias: "g-core",
				Name:  "Gcore",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "Gcore",
				GoRepositoryCount: 11,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Gcore-EI_IE2156026.11,16.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Gcore-Reviews-E2156026.htm",
			},
			OttaProfileSlug:   "Gcore",
			YouTubeChannelURL: "https://www.youtube.com/@GCoreOfficial",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go: []string{
					"https://www.linkedin.com/jobs/view/3888070191/",
				},
				domain.Rust:    nil,
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
			},
			ShortDescription: "",
			Industries:       []domain.Industry{},
		},

		// Some | Zep AI
		{
			ID:   0,  // system
			Type: "", // system
			Name: "Zep AI",
			URL:  "https://www.getzep.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    97181599,
				Alias: "zep-ai",
				Name:  "Zep AI",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "",
				GoRepositoryCount: 0,
				Verified:          false,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "",
				ReviewsURL:  "",
				Verified:    false,
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go: []string{
					"https://www.linkedin.com/jobs/view/3885860612/",
				},
				domain.Rust:    nil,
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
			},
			ShortDescription: "",
			Industries:       []domain.Industry{},
		},

		// Some | Gelato
		{
			ID:   0,  // system
			Type: "", // system
			Name: "Gelato",
			URL:  "https://www.gelato.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    5037871,
				Alias: "gelato",
				Name:  "Gelato",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "gelatoas",
				GoRepositoryCount: 0,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Gelato-EI_IE1297272.11,17.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Gelato-Reviews-E1297272.htm",
			},
			OttaProfileSlug:   "Gelato",
			YouTubeChannelURL: "https://www.youtube.com/@GelatoConnects",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go: []string{
					"https://www.linkedin.com/jobs/view/3866739991/",
				},
				domain.Rust:    nil,
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
			},
			ShortDescription: "",
			Industries:       []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{
				domain.Czechia,
			},
		},

		// Some | SumUp
		{
			ID:   0,  // system
			Type: "", // system
			Name: "SumUp",
			URL:  "https://www.sumup.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:       2619512,
				Alias:    "sumup",
				Name:     "SumUp",
				Verified: true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "sumup",
				GoRepositoryCount: 11,
				Verified:          true,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-SumUp-EI_IE673829.11,16.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/SumUp-Reviews-E673829.htm",
				Verified:    true,
			},
			OttaProfileSlug:   "sumup",
			YouTubeChannelURL: "https://www.youtube.com/@SumUpGlobal",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go: []string{
					"https://www.linkedin.com/jobs/view/3848461191/",
				},
				domain.Rust:  nil,
				domain.Zig:   nil,
				domain.Scala: nil,
				domain.Elixir: []string{
					"https://www.linkedin.com/jobs/view/4014335109/", // Senior Backend Engineer (Elixir & Go)
				},
				domain.Clojure: nil,
				domain.Haskell: nil,
			},
			ShortDescription: "Payments company",
			Industries: []domain.Industry{
				domain.IndustryFinTech,
			},
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
				domain.Czechia,
			},
		},

		// Some | Level Home
		{
			ID:   0,  // system
			Type: "", // system
			Name: "Level Home",
			URL:  "https://level.co/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    28779237,
				Alias: "levelhome",
				Name:  "Level Home Inc.",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "LevelHome",
				GoRepositoryCount: 0,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Level-Home-EI_IE3556695.11,21.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Level-Home-Reviews-E3556695.htm",
			},
			OttaProfileSlug:   "Level-Home",
			YouTubeChannelURL: "https://www.youtube.com/@LevelHome",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go: []string{
					"https://www.linkedin.com/jobs/view/3897014183/",
				},
				domain.Rust:    nil,
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
			},
			ShortDescription: "",
			Industries:       []domain.Industry{},
		},

		// Some | SonicWall
		{
			ID:   0,  // system
			Type: "", // system
			Name: "SonicWall",
			URL:  "https://www.sonicwall.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    4926,
				Alias: "sonicwall",
				Name:  "SonicWall",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "sonicwall",
				GoRepositoryCount: 0,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-SonicWall-EI_IE9777.11,20.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/SonicWall-Reviews-E9777.htm",
			},
			OttaProfileSlug:   "SonicWall",
			YouTubeChannelURL: "https://www.youtube.com/@SonicWallInc",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go: []string{
					"https://www.linkedin.com/jobs/view/3875719837/",
				},
				domain.Rust:    nil,
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
			},
			ShortDescription: "",
			Industries:       []domain.Industry{},
		},

		// Some | Pindrop
		{
			ID:   0,  // system
			Type: "", // system
			Name: "Pindrop",
			URL:  "https://www.pindrop.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    2326557,
				Alias: "pindrop",
				Name:  "Pindrop",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "", // NOP
				GoRepositoryCount: 0,  // NOP
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Pindrop-EI_IE709157.11,18.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Pindrop-Reviews-E709157.htm",
			},
			OttaProfileSlug:   "Pindrop",
			YouTubeChannelURL: "https://www.youtube.com/@Pindropsecurity",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go: []string{
					"https://www.linkedin.com/jobs/view/3901803458/",
				},
				domain.Rust:    nil,
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
			},
			ShortDescription: "",
			Industries:       []domain.Industry{},
		},

		// Some | Seedtag
		{
			ID:   0,  // system
			Type: "", // system
			Name: "Seedtag",
			URL:  "https://www.seedtag.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    5171806,
				Alias: "seedtag",
				Name:  "Seedtag",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "seedtag",
				GoRepositoryCount: 0,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Seedtag-EI_IE1421858.11,18.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Seedtag-Reviews-E1421858.htm",
			},
			OttaProfileSlug:   "Seedtag",
			YouTubeChannelURL: "https://www.youtube.com/@seedtag.advertising",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go: []string{
					"https://www.linkedin.com/jobs/view/3853123918/",
				},
				domain.Rust:    nil,
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
			},
			ShortDescription: "",
			Industries:       []domain.Industry{},
		},

		// Some | Flix
		{
			ID:   0,  // system
			Type: "", // system
			Name: "Flix",
			URL:  "https://www.flixbus.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    2726149,
				Alias: "flixbus",
				Name:  "Flix",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "flix-tech",
				GoRepositoryCount: 17,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Flix-EI_IE976463.11,15.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Flix-Reviews-E976463.htm",
			},
			OttaProfileSlug:   "FlixBus-1",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go: []string{
					"https://www.linkedin.com/jobs/view/3901538380/",
				},
				domain.Rust:    nil,
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
			},
			ShortDescription: "",
			Industries:       []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
				domain.Czechia,
			},
		},

		// Some | Press Ganey
		{
			ID:   0,  // system
			Type: "", // system
			Name: "Press Ganey",
			URL:  "https://www.pressganey.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    18007,
				Alias: "press-ganey-associates",
				Name:  "Press Ganey",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "",
				GoRepositoryCount: 0,
				Verified:          false,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Press-Ganey-EI_IE35100.11,22.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Press-Ganey-Reviews-E35100.htm",
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "https://www.youtube.com/@PressGaney",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go: []string{
					"https://www.linkedin.com/jobs/view/3898197086/",
				},
				domain.Rust:    nil,
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
			},
			ShortDescription: "",
			Industries:       []domain.Industry{},
		},

		// Some | Atmail
		{
			ID:   0,  // system
			Type: "", // system
			Name: "Atmail",
			URL:  "https://www.atmail.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    918978,
				Alias: "atmail",
				Name:  "Atmail",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "", // NOP
				GoRepositoryCount: 0,  // NOP
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "", // NOP
				ReviewsURL:  "", // NOP
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go: []string{
					"https://www.linkedin.com/jobs/view/3843643222/",
				},
				domain.Rust:    nil,
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
			},
			ShortDescription: "",
			Industries:       []domain.Industry{},
		},

		// Some | Dusty Robotics
		{
			ID:   0,  // system
			Type: "", // system
			Name: "Dusty Robotics",
			URL:  "https://www.dustyrobotics.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    33298433,
				Alias: "dusty-robotics",
				Name:  "Dusty Robotics",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "", // NOP
				GoRepositoryCount: 0,  // NOP
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Dusty-Robotics-EI_IE3518259.11,25.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Dusty-Robotics-Reviews-E3518259.htm",
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "https://www.youtube.com/@dustyrobotics",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go: []string{
					"https://www.linkedin.com/jobs/view/3804188613/",
					"https://www.linkedin.com/jobs/view/3832850412/",
				},
				domain.Rust:    nil,
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
			},
			ShortDescription: "",
			Industries:       []domain.Industry{},
		},

		// Some | Cimri
		{
			ID:   0,  // system
			Type: "", // system
			Name: "Cimri",
			URL:  "https://www.cimri.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    411498,
				Alias: "cimri",
				Name:  "Cimri",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "",
				GoRepositoryCount: 0,
				Verified:          false,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Cimri-EI_IE2401296.11,16.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Cimri-Reviews-E2401296.htm",
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "https://www.youtube.com/@cimri",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go: []string{
					"https://www.linkedin.com/jobs/view/3863565726/",
				},
				domain.Rust:    nil,
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
			},
			ShortDescription: "",
			Industries:       []domain.Industry{},
		},

		// Some | Quadcode
		{
			ID:   0,  // system
			Type: "", // system
			Name: "Quadcode",
			URL:  "https://quadcode.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    42345997,
				Alias: "quadcodecareer",
				Name:  "Quadcode",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "",
				GoRepositoryCount: 0,
				Verified:          false,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Quadcode-EI_IE3293771.11,19.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Quadcode-Reviews-E3293771.htm",
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go: []string{
					"https://www.linkedin.com/jobs/view/3907613945/",
				},
				domain.Rust:    nil,
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
			},
			ShortDescription: "Fintech company developing software for the trading and investment industry",
			Industries: []domain.Industry{
				domain.IndustryFinTech,
			},
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
			},
		},

		// Some | FinCompare
		{
			ID:   0,  // system
			Type: "", // system
			Name: "FinCompare",
			URL:  "https://fincompare.de/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    10976436,
				Alias: "fincompare",
				Name:  "FinCompare",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "fincompare",
				GoRepositoryCount: 1,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-FinCompare-EI_IE1644778.11,21.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/FinCompare-Reviews-E1644778.htm",
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go: []string{
					"https://www.linkedin.com/jobs/view/3846776554/",
				},
				domain.Rust:    nil,
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
			},
			ShortDescription: "",
			Industries:       []domain.Industry{},
		},

		// Some | Mellifera
		{
			ID:   0,  // system
			Type: "", // system
			Name: "Mellifera",
			URL:  "https://mellifera.team/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    98533069,
				Alias: "mellifera-operations-limited",
				Name:  "Mellifera Operations Limited",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "",
				GoRepositoryCount: 0,
				Verified:          false,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "",
				ReviewsURL:  "",
				Verified:    false,
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go: []string{
					"https://www.linkedin.com/jobs/view/3912533167/",
					"https://www.linkedin.com/jobs/view/3907553857/",
				},
				domain.Rust:    nil,
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
			},
			ShortDescription: "",
			Industries:       []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
			},
		},

		// Some | BNP Paribas - Securities Services
		{
			ID:   0,  // system
			Type: "", // system
			Name: "BNP Paribas - Securities Services",
			URL:  "https://securities.cib.bnpparibas/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    3625182,
				Alias: "bnpparibassecuritiesservices",
				Name:  "BNP Paribas - Securities Services",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "",
				GoRepositoryCount: 0,
				Verified:          false,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-BNP-Paribas-EI_IE10342.11,22.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/BNP-Paribas-Reviews-E10342.htm",
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "https://www.youtube.com/@labanquedunmondequichange",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go: []string{
					"https://www.linkedin.com/jobs/view/3909846279/",
				},
				domain.Rust:    nil,
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
			},
			ShortDescription: "",
			Industries:       []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
			},
		},

		// Some | Apifonica
		{
			ID:   0,  // system
			Type: "", // system
			Name: "Apifonica",
			URL:  "https://www.apifonica.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    10943475,
				Alias: "apifonica",
				Name:  "Apifonica",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "apifonica",
				GoRepositoryCount: 4,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Apifonica-EI_IE1805118.11,20.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Apifonica-Reviews-E1805118.htm",
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go: []string{
					"https://www.linkedin.com/jobs/view/3909698518/",
				},
				domain.Rust:    nil,
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
			},
			ShortDescription: "",
			Industries:       []domain.Industry{},
		},

		// Some | Cybus
		{
			ID:   0,                         // system
			Type: domain.CompanyTypeStartup, // system
			Name: "Cybus",
			URL:  "https://www.cybus.io/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:       7798667,
				Alias:    "cybus",
				Name:     "Cybus",
				Verified: false,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "cybusio",
				GoRepositoryCount: 1,
				Verified:          false,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Cybus-EI_IE2928520.11,16.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Cybus-Reviews-E2928520.htm",
				Verified:    false,
			},
			OttaProfileSlug:   "Cybus",
			YouTubeChannelURL: "https://www.youtube.com/@cybus_io",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go: []string{
					"https://www.linkedin.com/jobs/view/4054430926/", // Highly proficient in GO and proficient in Node.js for backend development
					"https://www.linkedin.com/jobs/view/3905530455/",
				},
				domain.Rust:    nil,
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
			},
			ShortDescription:          "Data architecture for the industrial IoT world",
			Industries:                []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{},
		},

		// Some | Flink
		{
			ID:   0,  // system
			Type: "", // system
			Name: "Flink",
			URL:  "https://www.goflink.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    71241902,
				Alias: "goflink",
				Name:  "Flink",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "",
				GoRepositoryCount: 0,
				Verified:          false,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Flink-EI_IE4921496.11,16.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Flink-Reviews-E4921496.htm",
			},
			OttaProfileSlug:   "Flink",
			YouTubeChannelURL: "https://www.youtube.com/@flink7309",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go: []string{
					"https://www.linkedin.com/jobs/view/3906306472/",
				},
				domain.Rust:    nil,
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
			},
			ShortDescription: "",
			Industries:       []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{
				domain.Czechia,
			},
		},

		// Some | Greenbone
		{
			ID:   0,  // system
			Type: "", // system
			Name: "Greenbone",
			URL:  "https://www.greenbone.net/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    697428,
				Alias: "greenbone-networks-gmbh",
				Name:  "Greenbone AG",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "greenbone",
				GoRepositoryCount: 7,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Greenbone-Networks-EI_IE379229.11,29.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Greenbone-Networks-Reviews-E379229.htm",
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "https://www.youtube.com/@greenbone",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go: []string{
					"https://www.linkedin.com/jobs/view/3892843872/",
				},
				domain.Rust:    nil,
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
			},
			ShortDescription: "",
			Industries:       []domain.Industry{},
		},

		// Some | OLX
		{
			ID:   0,  // system
			Type: "", // system
			Name: "OLX",
			URL:  "https://www.olxgroup.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:       167557,
				Alias:    "olx-group",
				Name:     "OLX",
				Verified: true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "",
				GoRepositoryCount: 0,
				Verified:          false,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-OLX-Group-EI_IE517166.11,20.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/OLX-Group-Reviews-E517166.htm",
				Verified:    true,
			},
			OttaProfileSlug:   "OLX-Group",
			YouTubeChannelURL: "https://www.youtube.com/@OLXGroup",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go: []string{
					"https://www.linkedin.com/jobs/view/4055838812/", // Our systems are built with Golang and they run on AWS
					"https://www.linkedin.com/jobs/view/3904643356/",
				},
				domain.Rust:    nil,
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
			},
			ShortDescription: "Marketplace ecosystems",
			Industries: []domain.Industry{
				domain.IndustryECommerce,
			},
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
				domain.Czechia,
			},
		},

		// Some | Snyk
		{
			ID:   0,  // system
			Type: "", // system
			Name: "Snyk",
			URL:  "https://snyk.io/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    10043614,
				Alias: "snyk",
				Name:  "Snyk",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "snyk",
				GoRepositoryCount: 22,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Snyk-EI_IE2094989.11,15.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Snyk-Reviews-E2094989.htm",
			},
			OttaProfileSlug:   "Snyk",
			YouTubeChannelURL: "https://www.youtube.com/@Snyksec",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go: []string{
					"https://www.linkedin.com/jobs/view/3903469567/",
					"https://app.otta.com/jobs/AOeyI9AM",
					"https://app.otta.com/jobs/HX7eW7pM",
				},
				domain.Rust:    nil,
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
			},
			ShortDescription: "",
			Industries:       []domain.Industry{},
		},

		// Some | Sinch
		{
			ID:   0,  // system
			Type: "", // system
			Name: "Sinch",
			URL:  "https://www.sinch.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    3726743,
				Alias: "sinch",
				Name:  "Sinch",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "sinch",
				GoRepositoryCount: 0,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Sinch-EI_IE788805.11,16.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Sinch-Reviews-E788805.htm",
			},
			OttaProfileSlug:   "sinch",
			YouTubeChannelURL: "https://www.youtube.com/@WeAreSinch",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go: []string{
					"https://www.linkedin.com/jobs/view/3903797937/",
				},
				domain.Rust:    nil,
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
			},
			ShortDescription: "",
			Industries:       []domain.Industry{},
		},

		// Some | FOX Tech
		{
			ID:   0,  // system
			Type: "", // system
			Name: "FOX Tech",
			URL:  "https://tech.fox.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    14850572,
				Alias: "foxtechteam",
				Name:  "FOX Tech",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "",
				GoRepositoryCount: 0,
				Verified:          false,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-FOX-Broadcasting-EI_IE13272.11,27.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/FOX-Broadcasting-Reviews-E13272.htm",
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go: []string{
					"https://www.linkedin.com/jobs/view/3903338524/",
				},
				domain.Rust:    nil,
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
			},
			ShortDescription: "",
			Industries:       []domain.Industry{},
		},

		// Some | SailPoint
		{
			ID:   0,  // system
			Type: "", // system
			Name: "SailPoint",
			URL:  "https://www.sailpoint.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    47456,
				Alias: "sailpoint-technologies",
				Name:  "SailPoint",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "sailpoint-oss",
				GoRepositoryCount: 2,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-SailPoint-Technologies-EI_IE449696.11,33.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/SailPoint-Technologies-Reviews-E449696.htm",
			},
			OttaProfileSlug:   "Sailpoint",
			YouTubeChannelURL: "https://www.youtube.com/@SailPointTechnologies",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go: []string{
					"https://www.linkedin.com/jobs/view/4030943277/",
					"https://www.linkedin.com/jobs/view/3903153316/",
				},
				domain.Rust:    nil,
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
			},
			ShortDescription: "",
			Industries:       []domain.Industry{},
		},

		// Some | Proofpoint
		{
			ID:   0,  // system
			Type: "", // system
			Name: "Proofpoint",
			URL:  "https://www.proofpoint.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    11681,
				Alias: "proofpoint",
				Name:  "Proofpoint",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "proofpoint",
				GoRepositoryCount: 0,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Proofpoint-EI_IE39140.11,21.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Proofpoint-Reviews-E39140.htm",
			},
			OttaProfileSlug:   "Proofpoint",
			YouTubeChannelURL: "https://www.youtube.com/@proofpoint",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go: []string{
					"https://www.linkedin.com/jobs/view/3797062001/",
				},
				domain.Rust:    nil,
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
			},
			ShortDescription: "",
			Industries:       []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
				domain.Czechia,
			},
		},

		// Some | Asset Reality
		{
			ID:   0,  // system
			Type: "", // system
			Name: "Asset Reality",
			URL:  "https://www.assetreality.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    42805677,
				Alias: "asset-reality",
				Name:  "Asset Reality",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "",
				GoRepositoryCount: 0,
				Verified:          false,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "",
				ReviewsURL:  "",
				Verified:    false,
			},
			OttaProfileSlug:   "Asset-Reality",
			YouTubeChannelURL: "https://www.youtube.com/@assetreality",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go: []string{
					"https://www.linkedin.com/jobs/view/3897435235/",
				},
				domain.Rust:    nil,
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
			},
			ShortDescription: "",
			Industries:       []domain.Industry{},
		},

		// Some | Limango
		{
			ID:   0,  // system
			Type: "", // system
			Name: "Limango",
			URL:  "https://www.limango.pl/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    2966982,
				Alias: "limango-sp-z-o-o-",
				Name:  "Limango",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "",
				GoRepositoryCount: 0,
				Verified:          false,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Limango-Polska-EI_IE2884426.11,25.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Limango-Polska-Reviews-E2884426.htm",
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "https://www.youtube.com/@limangoPolska",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go: []string{
					"https://justjoin.it/offers/limango-polska-mid-senior-backend-developer-golang",
				},
				domain.Rust:    nil,
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
			},
			ShortDescription: "",
			Industries:       []domain.Industry{},
		},

		// Some | RxBenefits
		{
			ID:   0,  // system
			Type: "", // system
			Name: "RxBenefits",
			URL:  "https://www.rxbenefits.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    2873210,
				Alias: "rxbenefits-inc-",
				Name:  "RxBenefits, Inc.",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "", // NOP
				GoRepositoryCount: 0,  // NOP
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-RxBenefits-EI_IE1175839.11,21.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/RxBenefits-Reviews-E1175839.htm",
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go: []string{
					"https://www.linkedin.com/jobs/view/3910221910/",
				},
				domain.Rust:    nil,
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
			},
			ShortDescription: "",
			Industries:       []domain.Industry{},
		},

		// Some | SmithRx
		{
			ID:   0,  // system
			Type: "", // system
			Name: "SmithRx",
			URL:  "https://www.smithrx.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    10886362,
				Alias: "smithrx",
				Name:  "SmithRx",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "", // NOP
				GoRepositoryCount: 0,  // NOP
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-SmithRx-EI_IE1901555.11,18.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/SmithRx-Reviews-E1901555.htm",
			},
			OttaProfileSlug:   "SmithRx",
			YouTubeChannelURL: "https://www.youtube.com/@SmithRxPBM",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go: []string{
					"https://www.linkedin.com/jobs/view/3910420916/",
				},
				domain.Rust:    nil,
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
			},
			ShortDescription: "",
			Industries:       []domain.Industry{},
		},

		// Some | k-ID
		{
			ID:   0,  // system
			Type: "", // system
			Name: "k-ID",
			URL:  "https://www.k-id.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    45973117,
				Alias: "k-id",
				Name:  "k-ID",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "", // NOP
				GoRepositoryCount: 0,  // NOP
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "",
				ReviewsURL:  "",
				Verified:    false,
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "https://www.youtube.com/@k-IDofficial",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go: []string{
					"https://www.linkedin.com/jobs/view/3912660665/",
				},
				domain.Rust:    nil,
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
			},
			ShortDescription: "",
			Industries:       []domain.Industry{},
		},

		// Some | CAFU
		{
			ID:   0,  // system
			Type: "", // system
			Name: "CAFU",
			URL:  "https://www.cafu.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    13892019,
				Alias: "mycafu",
				Name:  "CAFU",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "", // NOP
				GoRepositoryCount: 0,  // NOP
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-CAFU-EI_IE3713615.11,15.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/CAFU-Reviews-E3713615.htm",
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go: []string{
					"https://www.linkedin.com/jobs/view/3910059830/",
				},
				domain.Rust:    nil,
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
			},
			ShortDescription: "",
			Industries:       []domain.Industry{},
		},

		// Some | Rollee
		{
			ID:   0,  // system
			Type: "", // system
			Name: "Rollee",
			URL:  "https://www.getrollee.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    76353840,
				Alias: "rollee",
				Name:  "Rollee",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "", // NOP
				GoRepositoryCount: 0,  // NOP
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "", // NOP
				ReviewsURL:  "", // NOP
			},
			OttaProfileSlug:   "Rollee",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go: []string{
					"https://www.linkedin.com/jobs/view/3908431251/",
					"https://app.otta.com/jobs/4ajPfsAU",
				},
				domain.Rust:    nil,
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
			},
			ShortDescription: "",
			Industries:       []domain.Industry{},
		},

		// Some | Net2Phone
		{
			ID:   0,  // system
			Type: "", // system
			Name: "Net2Phone",
			URL:  "https://www.net2phone.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    2965,
				Alias: "net2phone",
				Name:  "Net2Phone",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "net2phone",
				GoRepositoryCount: 0,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Net2Phone-EI_IE9342.11,20.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Net2Phone-Reviews-E9342.htm",
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go: []string{
					"https://www.linkedin.com/jobs/view/3914839595/",
				},
				domain.Rust:    nil,
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
			},
			ShortDescription: "",
			Industries:       []domain.Industry{},
		},

		// Some | Ola Chat
		{
			ID:   0,  // system
			Type: "", // system
			Name: "Ola Chat",
			URL:  "https://olachat.sg/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    68332088,
				Alias: "ola-chat",
				Name:  "Ola Chat",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "olachat",
				GoRepositoryCount: 8,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "", // NOP
				ReviewsURL:  "", // NOP
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go: []string{
					"https://www.linkedin.com/jobs/view/3912130901/",
				},
				domain.Rust:    nil,
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
			},
			ShortDescription: "",
			Industries:       []domain.Industry{},
		},

		// Some | Veracity Software Inc
		{
			ID:   0,  // system
			Type: "", // system
			Name: "Veracity Software Inc",
			URL:  "https://veracity-us.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    11137552,
				Alias: "veracitysoftwareinc",
				Name:  "Veracity Software Inc",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "", // NOP
				GoRepositoryCount: 0,  // NOP
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Veracity-Software-EI_IE1357198.11,28.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Veracity-Software-Reviews-E1357198.htm",
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go: []string{
					"https://www.linkedin.com/jobs/view/3913559527/",
				},
				domain.Rust:    nil,
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
			},
			ShortDescription: "",
			Industries:       []domain.Industry{},
		},

		// Some | Treecard
		{
			ID:   0,  // system
			Type: "", // system
			Name: "Treecard",
			URL:  "https://www.treecard.org/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    68821773,
				Alias: "treecardapp",
				Name:  "Treecard",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "TreeCard",
				GoRepositoryCount: 2,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Treecard-EI_IE5675051.11,19.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Treecard-Reviews-E5675051.htm",
			},
			OttaProfileSlug:   "Treecard",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go: []string{
					"https://www.linkedin.com/jobs/view/3909975571/",
				},
				domain.Rust:    nil,
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
			},
			ShortDescription: "",
			Industries:       []domain.Industry{},
		},

		// Some | Openprovider
		{
			ID:   0,  // system
			Type: "", // system
			Name: "Openprovider",
			URL:  "https://www.openprovider.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    550698,
				Alias: "openprovider",
				Name:  "Openprovider",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "openprovider",
				GoRepositoryCount: 10,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Openprovider-EI_IE1267124.11,23.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Openprovider-Reviews-E1267124.htm",
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "https://www.youtube.com/@openprovider453",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go: []string{
					"https://www.linkedin.com/jobs/view/3914419419/",
				},
				domain.Rust:    nil,
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
			},
			ShortDescription: "",
			Industries:       []domain.Industry{},
		},

		// Some | fiskaly
		{
			ID:   0,  // system
			Type: "", // system
			Name: "fiskaly",
			URL:  "https://www.fiskaly.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:       18929063,
				Alias:    "fiskaly",
				Name:     "fiskaly",
				Verified: true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "fiskaly",
				GoRepositoryCount: 5,
				Verified:          true,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-fiskaly-EI_IE3059515.11,18.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/fiskaly-Reviews-E3059515.htm",
				Verified:    true,
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "https://www.youtube.com/@fiskaly",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go: []string{
					"https://www.linkedin.com/jobs/view/3914242491/", // 5 years working with Golang or other OOP languages
				},
				domain.Rust:    nil,
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
			},
			ShortDescription: "fiskaly builds innovative tools to make receipts simple",
			Industries: []domain.Industry{
				domain.IndustryFinTech,
			},
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
			},
		},

		// Some | Vay
		{
			ID:   0,  // system
			Type: "", // system
			Name: "Vay",
			URL:  "https://vay.io/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    12584218,
				Alias: "vaytechnology",
				Name:  "Vay",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "Reemote",
				GoRepositoryCount: 2,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "",
				ReviewsURL:  "",
				Verified:    false,
			},
			OttaProfileSlug:   "Vay",
			YouTubeChannelURL: "https://www.youtube.com/@vay_io",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go: []string{
					"https://www.linkedin.com/jobs/view/3912106948/",
				},
				domain.Rust:    nil,
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
			},
			ShortDescription: "",
			Industries:       []domain.Industry{},
		},

		// Some | Voltus
		{
			ID:   0,  // system
			Type: "", // system
			Name: "Voltus",
			URL:  "https://www.voltus.co/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    10866628,
				Alias: "voltus-inc.",
				Name:  "Voltus",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "voltusdev",
				GoRepositoryCount: 2,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Voltus-EI_IE2090197.11,17.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Voltus-Reviews-E2090197.htm",
			},
			OttaProfileSlug:   "Voltus",
			YouTubeChannelURL: "https://www.youtube.com/@Voltusinc",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go: []string{
					"https://app.otta.com/jobs/pB_hGf0W",
				},
				domain.Rust:    nil,
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
			},
			ShortDescription: "",
			Industries:       []domain.Industry{},
		},

		// Some | Stonebranch
		{
			ID:   0,  // system
			Type: "", // system
			Name: "Stonebranch",
			URL:  "https://www.stonebranch.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    71261,
				Alias: "stonebranch",
				Name:  "Stonebranch",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "stonebranch-marketplace",
				GoRepositoryCount: 0,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Stonebranch-EI_IE408996.11,22.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Stonebranch-Reviews-E408996.htm",
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "https://www.youtube.com/@Stonebranch",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go: []string{
					"https://www.linkedin.com/jobs/view/3914103960/",
				},
				domain.Rust:    nil,
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
			},
			ShortDescription: "",
			Industries:       []domain.Industry{},
		},

		// Some | Rapid7
		{
			ID:   0,  // system
			Type: "", // system
			Name: "Rapid7",
			URL:  "https://www.rapid7.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    39624,
				Alias: "rapid7",
				Name:  "Rapid7",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "rapid7",
				GoRepositoryCount: 7,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Rapid7-EI_IE243542.11,17.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Rapid7-Reviews-E243542.htm",
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "https://www.youtube.com/@GoRapid7",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go: []string{
					"https://www.linkedin.com/jobs/view/3915341496/",
				},
				domain.Rust:    nil,
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
			},
			ShortDescription: "",
			Industries:       []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
				domain.Czechia,
			},
		},

		// Some | Toggle AI
		{
			ID:   0,  // system
			Type: "", // system
			Name: "Toggle AI",
			URL:  "https://toggle.ai/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    28508827,
				Alias: "toggle-ai",
				Name:  "Toggle AI",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "",
				GoRepositoryCount: 0,
				Verified:          false,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Toggle-EI_IE3924898.11,17.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Toggle-Reviews-E3924898.htm",
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "https://www.youtube.com/@ToggleAI-Investing",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go: []string{
					"https://www.linkedin.com/jobs/view/3915194291/",
				},
				domain.Rust:    nil,
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
			},
			ShortDescription: "",
			Industries:       []domain.Industry{},
		},

		// Some | hearX Group
		{
			ID:   0,  // system
			Type: "", // system
			Name: "hearX Group",
			URL:  "https://www.hearxgroup.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    18002825,
				Alias: "hearx-group",
				Name:  "hearX Group",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "hearxgroup",
				GoRepositoryCount: 0,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-hearX-Group-EI_IE5800566.11,22.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/hearX-Group-Reviews-E5800566.htm",
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "https://www.youtube.com/@hearxgroupptyltd8061",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go: []string{
					"https://www.linkedin.com/jobs/view/3913102407/",
				},
				domain.Rust:    nil,
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
			},
			ShortDescription: "",
			Industries:       []domain.Industry{},
		},

		// Some | MarkiTech.AI
		{
			ID:   0,  // system
			Type: "", // system
			Name: "MarkiTech.AI",
			URL:  "https://markitech.ca/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    9294422,
				Alias: "markitech-ai",
				Name:  "MarkiTech.AI",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "", // NOP
				GoRepositoryCount: 0,  // NOP
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Markitech-EI_IE4190913.11,20.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Markitech-Reviews-E4190913.htm",
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "https://www.youtube.com/@markitech-digitaltransform6173",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go: []string{
					"https://www.linkedin.com/jobs/view/3912880597/",
				},
				domain.Rust:    nil,
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
			},
			ShortDescription: "",
			Industries:       []domain.Industry{},
		},

		// Some | Lantronix
		{
			ID:   0,  // system
			Type: "", // system
			Name: "Lantronix",
			URL:  "https://www.lantronix.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    12612,
				Alias: "lantronix",
				Name:  "Lantronix",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "Lantronix",
				GoRepositoryCount: 0,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Lantronix-EI_IE5498.11,20.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Lantronix-Reviews-E5498.htm",
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "https://www.youtube.com/@LantronixInc",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go: []string{
					"https://www.linkedin.com/jobs/view/3915346698/",
				},
				domain.Rust:    nil,
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
			},
			ShortDescription: "",
			Industries:       []domain.Industry{},
		},

		// Some | INFOLOB
		{
			ID:   0,  // system
			Type: "", // system
			Name: "INFOLOB",
			URL:  "https://www.infolob.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    404211,
				Alias: "infolob-global",
				Name:  "INFOLOB",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "", // NOP
				GoRepositoryCount: 0,  // NOP
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Infolob-Solutions-EI_IE423113.11,28.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Infolob-Solutions-Reviews-E423113.htm",
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go: []string{
					"https://www.linkedin.com/jobs/view/3914415181/",
				},
				domain.Rust:    nil,
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
			},
			ShortDescription: "",
			Industries:       []domain.Industry{},
		},

		// Some | Argela Technologies
		{
			ID:   0,  // system
			Type: "", // system
			Name: "Argela Technologies",
			URL:  "https://www.argela.com.tr/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    26805,
				Alias: "argela-technologies",
				Name:  "Argela Technologies",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "", // NOP
				GoRepositoryCount: 0,  // NOP
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Argela-EI_IE389528.11,17.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Argela-Reviews-E389528.htm",
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "https://www.youtube.com/@argelatechnologies",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go: []string{
					"https://www.linkedin.com/jobs/view/3914290466/",
				},
				domain.Rust:    nil,
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
			},
			ShortDescription: "",
			Industries:       []domain.Industry{},
		},

		// Some | Top Doctors
		{
			ID:   0,  // system
			Type: "", // system
			Name: "Top Doctors",
			URL:  "https://topdoctors.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    3035481,
				Alias: "top-doctors-europe",
				Name:  "Top Doctors",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "", // NOP
				GoRepositoryCount: 0,  // NOP
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Top-Doctors-EI_IE1712721.11,22.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Top-Doctors-Reviews-E1712721.htm",
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "https://www.youtube.com/@TopDoctorsUK",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go: []string{
					"https://www.linkedin.com/jobs/view/3912190311/",
				},
				domain.Rust:    nil,
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
			},
			ShortDescription: "",
			Industries:       []domain.Industry{},
		},

		// Some | Recurly
		{
			ID:   0,  // system
			Type: "", // system
			Name: "Recurly",
			URL:  "https://recurly.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    810383,
				Alias: "recurly",
				Name:  "Recurly",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "recurly",
				GoRepositoryCount: 2,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Recurly-EI_IE692611.11,18.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Recurly-Reviews-E692611.htm",
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "https://www.youtube.com/@Recurly",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go: []string{
					"https://www.linkedin.com/jobs/view/3912991626/",
				},
				domain.Rust:    nil,
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
			},
			ShortDescription: "",
			Industries:       []domain.Industry{},
		},

		// Some | Cynet Systems
		{
			ID:   0,  // system
			Type: "", // system
			Name: "Cynet Systems",
			URL:  "https://www.cynetsystems.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    5003556,
				Alias: "cynet-systems",
				Name:  "Cynet Systems",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "", // NOP
				GoRepositoryCount: 0,  // NOP
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Cynet-Systems-EI_IE654628.11,24.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Cynet-Systems-Reviews-E654628.htm",
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "https://www.youtube.com/@cynetsystems2026",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go: []string{
					"https://www.linkedin.com/jobs/view/3912977786/",
				},
				domain.Rust:    nil,
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
			},
			ShortDescription: "",
			Industries:       []domain.Industry{},
		},

		// Some | Odyssey Information Services
		{
			ID:   0,  // system
			Type: "", // system
			Name: "Odyssey Information Services",
			URL:  "https://www.odysseyis.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    56445,
				Alias: "odyssey-information-services",
				Name:  "Odyssey Information Services",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "", // NOP
				GoRepositoryCount: 0,  // NOP
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Odyssey-Information-Services-EI_IE558201.11,39.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Odyssey-Information-Services-Reviews-E558201.htm",
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "https://www.youtube.com/@odysseyinformationservices1641",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go: []string{
					"https://www.linkedin.com/jobs/view/3905563294/",
				},
				domain.Rust:    nil,
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
			},
			ShortDescription: "",
			Industries:       []domain.Industry{},
		},

		// Some | Infomatics Corp
		{
			ID:   0,  // system
			Type: "", // system
			Name: "Infomatics Corp",
			URL:  "https://infomaticscorp.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    821065,
				Alias: "infomatics-corp",
				Name:  "Infomatics Corp",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "", // NOP
				GoRepositoryCount: 0,  // NOP
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Infomatics-EI_IE925978.11,21.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Infomatics-Reviews-E925978.htm",
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go: []string{
					"https://www.linkedin.com/jobs/view/3912956713/",
				},
				domain.Rust:    nil,
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
			},
			ShortDescription: "",
			Industries:       []domain.Industry{},
		},

		// Some | Mindera
		{
			ID:   0,  // system
			Type: "", // system
			Name: "Mindera",
			URL:  "https://mindera.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    80044268,
				Alias: "mindera-world",
				Name:  "Mindera",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "Mindera",
				GoRepositoryCount: 3,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Mindera-EI_IE1139926.11,18.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Mindera-Reviews-E1139926.htm",
			},
			OttaProfileSlug:   "Mindera",
			YouTubeChannelURL: "https://www.youtube.com/@MinderaSoftware",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go: []string{
					"https://www.linkedin.com/jobs/view/3915092213/",
				},
				domain.Rust:    nil,
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
			},
			ShortDescription: "",
			Industries:       []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{
				domain.Czechia,
			},
		},

		// Some | Sytac
		{
			ID:   0,  // system
			Type: "", // system
			Name: "Sytac",
			URL:  "https://sytac.io/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    626751,
				Alias: "sytac",
				Name:  "Sytac",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "sytac",
				GoRepositoryCount: 0,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Sytac-EI_IE1255983.11,16.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Sytac-Reviews-E1255983.htm",
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "https://www.youtube.com/@sytac",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go: []string{
					"https://www.linkedin.com/jobs/view/3912222258/",
				},
				domain.Rust:    nil,
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
			},
			ShortDescription: "",
			Industries:       []domain.Industry{},
		},

		// Some | Qumulus Cloud Platform
		{
			ID:   0,  // system
			Type: "", // system
			Name: "Qumulus Cloud Platform",
			URL:  "https://www.qumulus.io/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    81905634,
				Alias: "qumuluscloudplatform",
				Name:  "Qumulus Cloud Platform",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "QumulusTechnology",
				GoRepositoryCount: 2,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "", // NOP
				ReviewsURL:  "", // NOP
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go: []string{
					"https://www.linkedin.com/jobs/view/3914358799/",
				},
				domain.Rust:    nil,
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
			},
			ShortDescription: "",
			Industries:       []domain.Industry{},
		},

		// Some | Saxon AI
		{
			ID:   0,  // system
			Type: "", // system
			Name: "Saxon AI",
			URL:  "https://saxon.ai/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    224935,
				Alias: "saxonai",
				Name:  "Saxon AI",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "", // NOP
				GoRepositoryCount: 0,  // NOP
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "", // NOP
				ReviewsURL:  "", // NOP
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "https://www.youtube.com/@saxonai",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go: []string{
					"https://www.linkedin.com/jobs/view/3914185384/",
				},
				domain.Rust:    nil,
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
			},
			ShortDescription: "",
			Industries:       []domain.Industry{},
		},

		// Some | Dyninno Group
		{
			ID:   0,  // system
			Type: "", // system
			Name: "Dyninno Group",
			URL:  "https://dyninno.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    9456141,
				Alias: "dyninno-group",
				Name:  "Dyninno Group",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "", // NOP
				GoRepositoryCount: 0,  // NOP
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-DYNINNO-Group-EI_IE2562842.11,24.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/DYNINNO-Group-Reviews-E2562842.htm",
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "https://www.youtube.com/@dyninnogroup2702",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go: []string{
					"https://www.linkedin.com/jobs/view/3914135515/",
				},
				domain.Rust:    nil,
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
			},
			ShortDescription: "",
			Industries:       []domain.Industry{},
		},

		// TODO solve alias later "holland-&-barrett"
		// Some | Holland & Barrett
		//{
		//	ID: 0,
		//	Name: "Holland & Barrett",
		//	URL:  "https://www.hollandandbarrett.com/",
		//	LinkedInProfile: domain.LinkedInProfile{
		//		ID:    815488,
		//		Alias: "holland-&-barrett",
		//		Name:  "Holland & Barrett",
		//	},
		//	GitHubProfile: domain.GitHubProfile{
		//		Login:             "", // NOP
		//		GoRepositoryCount: 0,  // NOP
		//	},
		//	GlassdoorProfile: domain.GlassdoorProfile{
		//		OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Holland-and-Barrett-EI_IE36174.11,30.htm",
		//		ReviewsURL:  "https://www.glassdoor.com/Reviews/Holland-and-Barrett-Reviews-E36174.htm",
		//	},
		//	OttaProfileSlug:   "",
		//	YouTubeChannelURL: "https://www.youtube.com/@hollandandbarrett",
		//	GoMainLanguage:    false,
		//	Vacancies: domain.Vacancies{
		//		domain.Go: []string{
		//			"https://www.linkedin.com/jobs/view/3911059005/",
		//		},
		//		domain.Rust:    nil,
		//		domain.Zig:     nil,
		//		domain.Scala:   nil,
		//		domain.Elixir:  nil,
		//		domain.Clojure: nil,
		//		domain.Haskell: nil,
		//	},
		//},

		// Some | Group Avows
		{
			ID:   0,  // system
			Type: "", // system
			Name: "Group Avows",
			URL:  "https://avowstech.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    3651016,
				Alias: "group-avows",
				Name:  "Group Avows",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "", // NOP
				GoRepositoryCount: 0,  // NOP
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-AVOWS-Technologies-EI_IE870406.11,29.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/AVOWS-Technologies-Reviews-E870406.htm",
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "https://www.youtube.com/@avowsgroupofficial9907",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go: []string{
					"https://www.linkedin.com/jobs/view/3911007016/",
				},
				domain.Rust:    nil,
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
			},
			ShortDescription: "",
			Industries:       []domain.Industry{},
		},

		// Some | Cognizant
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Cognizant",
			Website: "https://www.cognizant.com/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:       1680,
				Alias:    "cognizant",
				Name:     "Cognizant",
				Verified: true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "", // NOP
				GoRepositoryCount: 0,  // NOP
			},
			BlindProfile:     domain.BlindProfile{},
			LevelsFyiProfile: domain.LevelsFyiProfile{},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Cognizant-Technology-Solutions-EI_IE8014.11,41.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Cognizant-Technology-Solutions-Reviews-E8014.htm",
			},
			IndeedProfile:     domain.IndeedProfile{},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "https://www.youtube.com/@cognizant",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go: []string{
					"https://www.linkedin.com/jobs/view/3914839625/",
				},
				domain.Rust:    nil,
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
			},
			Languages: domain.Languages{
				domain.Go: {},
				domain.Rust: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Rust Developer (Remote)",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4066607632/",
							Date:             mustDate("2024-11-07"),
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
			DealroomURL:      "",
			CrunchbaseURL:    "",
			PitchbookURL:     "",
			YahooFinanceURL:  "",
			GoogleFinanceURL: "",
			YCombinatorURL:   "",
			Industries:       []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
				domain.Czechia,
			},
		},

		// Some | Nuro
		{
			ID:   0,  // system
			Type: "", // system
			Name: "Nuro",
			URL:  "https://www.nuro.ai/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    12957486,
				Alias: "nuro-inc.",
				Name:  "Nuro",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "", // NOP
				GoRepositoryCount: 0,  // NOP
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Nuro-EI_IE1550693.11,15.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Nuro-Reviews-E1550693.htm",
			},
			OttaProfileSlug:   "Nuro",
			YouTubeChannelURL: "https://www.youtube.com/@NuroTeam",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go: []string{
					"https://app.otta.com/jobs/P5a_50Xb",
				},
				domain.Rust:    nil,
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
			},
			ShortDescription: "",
			Industries:       []domain.Industry{},
		},

		// Some | CloudWalk
		{
			ID:   0,  // system
			Type: "", // system
			Name: "CloudWalk",
			URL:  "https://cloudwalk.io/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    3523168,
				Alias: "cloudwalk-inc",
				Name:  "CloudWalk, Inc.",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "cloudwalk",
				GoRepositoryCount: 6,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-CloudWalk-Inc-EI_IE2722308.11,24.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/CloudWalk-Inc-Reviews-E2722308.htm",
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "https://www.youtube.com/@cloudwalk_shorts",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go: []string{
					"https://www.linkedin.com/jobs/view/3917078597/",
				},
				domain.Rust:    nil,
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
			},
			ShortDescription: "",
			Industries:       []domain.Industry{},
		},

		// Some | Transition Technologies PSC
		{
			ID:   0,  // system
			Type: "", // system
			Name: "Transition Technologies PSC",
			URL:  "https://ttpsc.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    17880075,
				Alias: "transition-technologies-psc",
				Name:  "Transition Technologies PSC",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "", // NOP
				GoRepositoryCount: 0,  // NOP
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Transition-Technologies-PSC-EI_IE1875542.11,38.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Transition-Technologies-PSC-Reviews-E1875542.htm",
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "https://www.youtube.com/@TransitionTechnologiesPSC",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go: []string{
					"https://www.linkedin.com/jobs/view/3886949034/",
				},
				domain.Rust:    nil,
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
			},
			ShortDescription: "",
			Industries:       []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
			},
		},

		// Some | Kroger
		{
			ID:   0,  // system
			Type: "", // system
			Name: "Kroger",
			URL:  "https://www.kroger.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    4914,
				Alias: "kroger",
				Name:  "Kroger",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "Kroger-Technology",
				GoRepositoryCount: 0,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Kroger-EI_IE386.11,17.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Kroger-Reviews-E386.htm",
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "https://www.youtube.com/@KrogerCo",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go: []string{
					"https://www.linkedin.com/jobs/view/3917025739/",
				},
				domain.Rust:    nil,
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
			},
			ShortDescription: "",
			Industries:       []domain.Industry{},
		},

		// Some | Precisely
		{
			ID:   0,  // system
			Type: "", // system
			Name: "Precisely",
			URL:  "https://www.precisely.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    64863146,
				Alias: "preciselydata",
				Name:  "Precisely",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "PreciselyData",
				GoRepositoryCount: 1,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Precisely-EI_IE3372755.11,20.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Precisely-Reviews-E3372755.htm",
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "https://www.youtube.com/@PreciselyData",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go: []string{
					"https://www.linkedin.com/jobs/view/3917000992/",
				},
				domain.Rust:    nil,
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
			},
			ShortDescription: "",
			Industries:       []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
				domain.Czechia,
			},
		},

		// Some | R Systems
		{
			ID:   0,  // system
			Type: "", // system
			Name: "R Systems",
			URL:  "https://www.rsystems.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    165636,
				Alias: "r-systems",
				Name:  "R Systems",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "", // NOP
				GoRepositoryCount: 0,  // NOP
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-R-Systems-EI_IE32444.11,20.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/R-Systems-Reviews-E32444.htm",
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "https://www.youtube.com/@RSystems_inc",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go: []string{
					"https://www.linkedin.com/jobs/view/3916555239/",
				},
				domain.Rust:    nil,
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
			},
			ShortDescription: "",
			Industries:       []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
			},
		},

		// Some | Hays
		{
			ID:   0,  // system
			Type: "", // system
			Name: "Hays",
			URL:  "https://www.haysplc.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    3486,
				Alias: "hays",
				Name:  "Hays",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "", // NOP
				GoRepositoryCount: 0,  // NOP
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Hays-EI_IE10166.11,15.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Hays-Reviews-E10166.htm",
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "https://www.youtube.com/@HaysTV",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go: []string{
					"https://www.linkedin.com/jobs/view/3913846166/",
				},
				domain.Rust:    nil,
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
			},
			ShortDescription:          "",
			Industries:                []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{
				//
			},
		},

		// Some | Consort Group
		{
			ID:   0,  // system
			Type: "", // system
			Name: "Consort Group",
			URL:  "https://consort-group.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    46088,
				Alias: "consortgroup",
				Name:  "Consort Group",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "", // NOP
				GoRepositoryCount: 0,  // NOP
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Groupe-Consort-EI_IE915503.11,25.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Groupe-Consort-Reviews-E915503.htm",
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "https://www.youtube.com/@consortgroup",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go: []string{
					"https://www.linkedin.com/jobs/view/3913831485/",
				},
				domain.Rust:    nil,
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
			},
			ShortDescription: "",
			Industries:       []domain.Industry{},
		},

		// Some | Ascendion
		{
			ID:   0,  // system
			Type: "", // system
			Name: "Ascendion",
			URL:  "https://ascendion.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    86694680,
				Alias: "ascendion",
				Name:  "Ascendion",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "", // NOP
				GoRepositoryCount: 0,  // NOP
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Ascendion-EI_IE7774544.11,20.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Ascendion-Reviews-E7774544.htm",
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "https://www.youtube.com/@ascendioninc",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go: []string{
					"https://www.linkedin.com/jobs/view/3916325061/",
				},
				domain.Rust:    nil,
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
			},
			ShortDescription: "",
			Industries:       []domain.Industry{},
		},

		// Some | Checkout.com
		{
			ID:   0,  // system
			Type: "", // system
			Name: "Checkout.com",
			URL:  "https://www.checkout.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    3110635,
				Alias: "checkout",
				Name:  "Checkout.com",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "checkout",
				GoRepositoryCount: 1,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Checkout-com-EI_IE837487.11,23.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Checkout-com-Reviews-E837487.htm",
			},
			OttaProfileSlug:   "Checkout-com",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go: []string{
					"https://www.linkedin.com/jobs/view/3920739722/",
				},
				domain.Rust:    nil,
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
			},
			ShortDescription: "",
			Industries:       []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
				domain.Czechia,
			},
		},

		// Some | Unlimit
		{
			ID:   0,  // system
			Type: "", // system
			Name: "Unlimit",
			URL:  "https://www.unlimit.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    56459436,
				Alias: "unlimit-com",
				Name:  "Unlimit",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "",
				GoRepositoryCount: 0,
				Verified:          false,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "",
				ReviewsURL:  "",
				Verified:    false,
			},
			OttaProfileSlug:   "Unlimit",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go: []string{
					"https://app.otta.com/jobs/P4ept_aQ",
				},
				domain.Rust:    nil,
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
			},
			ShortDescription: "",
			Industries:       []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
				domain.Czechia,
			},
		},

		// Some | Chime
		{
			ID:   0,  // system
			Type: "", // system
			Name: "Chime",
			URL:  "https://www.chime.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    3738695,
				Alias: "chime-card",
				Name:  "Chime",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "",
				GoRepositoryCount: 0,
				Verified:          false,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Chime-EI_IE1493686.11,16.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Chime-Reviews-E1493686.htm",
			},
			OttaProfileSlug:   "Chime",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go: []string{
					"https://app.otta.com/jobs/nKhnfPzD",
				},
				domain.Rust:    nil,
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
			},
			ShortDescription: "",
			Industries:       []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
				domain.Czechia,
			},
		},

		// Some | Collective Minds Radiology
		{
			ID:   0,  // system
			Type: "", // system
			Name: "Collective Minds Radiology",
			URL:  "https://cmrad.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    11564593,
				Alias: "cmrad",
				Name:  "Collective Minds Radiology",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "",
				GoRepositoryCount: 0,
				Verified:          false,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Collective-Minds-Radiology-EI_IE4686955.11,37.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Collective-Minds-Radiology-Reviews-E4686955.htm",
			},
			OttaProfileSlug:   "Collective-Minds-Radiology",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go:      []string{},
				domain.Rust:    nil,
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
			},
			ShortDescription: "",
			Industries:       []domain.Industry{},
		},

		// Some | Cruise
		{
			ID:   0,  // system
			Type: "", // system
			Name: "Cruise",
			URL:  "https://www.getcruise.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    6577635,
				Alias: "getcruise",
				Name:  "Cruise",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "cruise-automation",
				GoRepositoryCount: 6,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Cruise-EI_IE977351.11,17.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Cruise-Reviews-E977351.htm",
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go: []string{
					"https://app.otta.com/jobs/xpGrXDnd",
				},
				domain.Rust:    nil,
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
			},
			ShortDescription: "",
			Industries:       []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
			},
		},

		// Some | Compass
		{
			ID:   0,  // system
			Type: "", // system
			Name: "Compass",
			URL:  "https://www.compass.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    2866215,
				Alias: "compassinc",
				Name:  "Compass",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "UrbanCompass",
				GoRepositoryCount: 10,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Compass-EI_IE719025.11,18.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Compass-Reviews-E719025.htm",
			},
			OttaProfileSlug:   "Compass-1",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go: []string{
					"https://app.otta.com/jobs/EK9Av13p",
				},
				domain.Rust:    nil,
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
			},
			ShortDescription: "",
			Industries:       []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
			},
		},

		// Some | Mercury
		{
			ID:   0,  // system
			Type: "", // system
			Name: "Mercury",
			URL:  "https://mercury.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    19107985,
				Alias: "mercuryhq",
				Name:  "Mercury",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "",
				GoRepositoryCount: 0,
				Verified:          false,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Mercury-EI_IE3583070.11,18.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Mercury-Reviews-E3583070.htm",
			},
			OttaProfileSlug:   "Mercury",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go:      nil,
				domain.Rust:    nil,
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: []string{
					"https://app.welcometothejungle.com/jobs/UzSGGV5V",
					"https://app.otta.com/jobs/8o2_A2QN",
				},
			},
			ShortDescription: "FinTech company",
			Industries: []domain.Industry{
				domain.IndustryFinTech,
			},
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
				domain.Czechia,
			},
		},

		{
			ID:   0,  // system
			Type: "", // system
			Name: "Tailscale",
			URL:  "https://tailscale.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    35653234,
				Alias: "tailscale",
				Name:  "Tailscale",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "tailscale",
				GoRepositoryCount: 68,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "",
				ReviewsURL:  "",
				Verified:    false,
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go: []string{
					"https://app.otta.com/jobs/N2SJQvgM",
				},
				domain.Rust:    nil,
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
			},
			ShortDescription: "VPN",
			Industries: []domain.Industry{
				domain.IndustryCyberSecurity,
			},
		},
		{
			ID:   0,  // system
			Type: "", // system
			Name: "Customer.io",
			URL:  "https://customer.io/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    2492471,
				Alias: "customer-io",
				Name:  "Customer.io",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "customerio",
				GoRepositoryCount: 21,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Customer-io-EI_IE1308885.11,22.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Customer-io-Reviews-E1308885.htm",
			},
			OttaProfileSlug:   "Customer-io",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go: []string{
					"https://app.otta.com/jobs/QaTOxLFC",
				},
				domain.Rust:    nil,
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
			},
			ShortDescription:          "",
			Industries:                []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{},
		},
		{
			ID:   0,  // system
			Type: "", // system
			Name: "Tabby",
			URL:  "https://tabby.ai/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    26615638,
				Alias: "tabbypay",
				Name:  "Tabby",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "",
				GoRepositoryCount: 0,
				Verified:          false,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-tabby-EI_IE6075206.11,16.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/tabby-Reviews-E6075206.htm",
			},
			OttaProfileSlug:   "Tabby",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go: []string{
					"https://www.linkedin.com/jobs/view/4012619449/",
					"https://www.linkedin.com/jobs/view/4012621025/",
					"https://app.otta.com/jobs/lZl0yCwo",
				},
				domain.Rust:    nil,
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
			},
			ShortDescription: "",
			Industries: []domain.Industry{
				domain.IndustryFinTech,
			},
			HasEmployeesFromCountries: []domain.Country{},
		},
		{
			ID:   0,  // system
			Type: "", // system
			Name: "Sonatus",
			URL:  "https://www.sonatus.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    18780890,
				Alias: "sonatus",
				Name:  "Sonatus",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "",
				GoRepositoryCount: 0,
				Verified:          false,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Sonatus-EI_IE3258616.11,18.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Sonatus-Reviews-E3258616.htm",
			},
			OttaProfileSlug:   "Sonatus",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go: []string{
					"https://app.otta.com/jobs/0g6xbMY_",
				},
				domain.Rust:    nil,
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
			},
			ShortDescription:          "",
			Industries:                []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{},
		},
		{
			ID:   0,  // system
			Type: "", // system
			Name: "KOHO",
			URL:  "https://www.koho.ca/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    9443598,
				Alias: "koho",
				Name:  "KOHO",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "kohofinancial",
				GoRepositoryCount: 4,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-KOHO-EI_IE2155372.11,15.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/KOHO-Reviews-E2155372.htm",
			},
			OttaProfileSlug:   "KOHO-Financial",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go: []string{
					"https://www.linkedin.com/jobs/view/4032999298/",
					"https://app.otta.com/jobs/7pHaeAso",
				},
				domain.Rust:    nil,
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
			},
			ShortDescription: "",
			Industries: []domain.Industry{
				domain.IndustryFinTech,
			},
			HasEmployeesFromCountries: []domain.Country{},
		},
		{
			ID:   0,  // system
			Type: "", // system
			Name: "Nicolab",
			URL:  "https://www.nicolab.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    10873366,
				Alias: "nico-lab",
				Name:  "Nicolab",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "",
				GoRepositoryCount: 0,
				Verified:          false,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "",
				ReviewsURL:  "",
				Verified:    false,
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go: []string{
					"https://app.otta.com/jobs/h0EkrVdy",
				},
				domain.Rust:    nil,
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
			},
			ShortDescription: "",
			Industries: []domain.Industry{
				domain.IndustryHealthTech,
				domain.IndustryMedTech,
			},
			HasEmployeesFromCountries: []domain.Country{},
		},
		{
			ID:   0,  // system
			Type: "", // system
			Name: "Operant AI",
			URL:  "https://www.operant.ai/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    67522377,
				Alias: "operantai",
				Name:  "Operant AI",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "OperantAI",
				GoRepositoryCount: 1,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "",
				ReviewsURL:  "",
				Verified:    false,
			},
			OttaProfileSlug:   "Operant-AI",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go: []string{
					"https://www.linkedin.com/jobs/view/4000951216/",
					"https://app.otta.com/jobs/wily0aG5",
				},
				domain.Rust:    nil,
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
			},
			ShortDescription: "",
			Industries: []domain.Industry{
				domain.IndustryCyberSecurity,
			},
			HasEmployeesFromCountries: []domain.Country{},
		},
		{
			ID:   0,  // system
			Type: "", // system
			Name: "CrowdRiff",
			URL:  "https://crowdriff.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    2781710,
				Alias: "crowd-riff",
				Name:  "CrowdRiff",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "crowdriff",
				GoRepositoryCount: 10,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-CrowdRiff-EI_IE1643945.11,20.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/CrowdRiff-Reviews-E1643945.htm",
			},
			OttaProfileSlug:   "CrowdRiff",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go: []string{
					"https://www.linkedin.com/jobs/view/4022030075/",
					"https://app.otta.com/jobs/0yGgNSdM",
				},
				domain.Rust:    nil,
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
			},
			ShortDescription: "",
			Industries: []domain.Industry{
				domain.IndustryMarTech,
			},
			HasEmployeesFromCountries: []domain.Country{},
		},
		{
			ID:   0,  // system
			Type: "", // system
			Name: "Swiss Post",
			URL:  "https://www.post.ch/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    338556,
				Alias: "swiss-post",
				Name:  "Swiss Post",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "swisspost",
				GoRepositoryCount: 5,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Die-Schweizerische-Post-EI_IE12870.11,34.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Die-Schweizerische-Post-Reviews-E12870.htm",
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go: []string{
					"https://www.linkedin.com/jobs/view/4040049773/",
				},
				domain.Rust:    nil,
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
			},
			ShortDescription: "",
			Industries:       []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Rialtic",
			Website: "https://www.rialtic.io/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:       68821761,
				Alias:    "rialtic-io",
				Name:     "Rialtic",
				Verified: true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "Rialtic",
				GoRepositoryCount: 0,
			},
			BlindProfile:     domain.BlindProfile{},
			LevelsFyiProfile: domain.LevelsFyiProfile{},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Rialtic-EI_IE4497416.11,18.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Rialtic-Reviews-E4497416.htm",
			},
			IndeedProfile:     domain.IndeedProfile{},
			OttaProfileSlug:   "Rialtic",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Senior Golang Engineer",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4062048996/",
							Date:             mustDate("2024-11-01"),
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
			ShortDescription: "Administrative and financial technology for healthcare",
			DealroomURL:      "",
			CrunchbaseURL:    "",
			PitchbookURL:     "",
			YahooFinanceURL:  "",
			GoogleFinanceURL: "",
			YCombinatorURL:   "",
			Industries: []domain.Industry{
				domain.IndustryHealthTech,
			},
			HasEmployeesFromCountries: []domain.Country{},
		},
		{
			ID:   0,  // system
			Type: "", // system
			Name: "Prisma",
			URL:  "https://www.prisma.io/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    10614939,
				Alias: "prisma-io",
				Name:  "Prisma",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "prisma",
				GoRepositoryCount: 3,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Prisma-Data-EI_IE2431237.11,22.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Prisma-Data-Reviews-E2431237.htm",
			},
			OttaProfileSlug:   "Prisma",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go: []string{
					"https://app.otta.com/jobs/k6ASPhV7",
				},
				domain.Rust: []string{
					"https://app.otta.com/jobs/k6ASPhV7",
				},
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
			},
			ShortDescription: "",
			Industries: []domain.Industry{
				domain.IndustryDevOps,
			},
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
			},
		},
		{
			ID:   0,  // system
			Type: "", // system
			Name: "Zipline",
			URL:  "https://www.flyzipline.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    7602863,
				Alias: "flyzipline",
				Name:  "Zipline",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "",
				GoRepositoryCount: 0,
				Verified:          false,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Zipline-EI_IE1394276.11,18.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Zipline-Reviews-E1394276.htm",
			},
			OttaProfileSlug:   "Zipline",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go: []string{
					"https://app.otta.com/jobs/DwReS2t8",
				},
				domain.Rust:    nil,
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
			},
			ShortDescription:          "",
			Industries:                []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{},
		},

		{
			ID:   0,                         // system
			Type: domain.CompanyTypeStartup, // system
			Name: "VisionAI",
			URL:  "https://www.visionai.co/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    18858131,
				Alias: "wearevisionai",
				Name:  "VisionAI",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "",
				GoRepositoryCount: 0,
				Verified:          false,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-VisionAI-EI_IE9666263.11,19.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/VisionAI-Reviews-E9666263.htm",
			},
			OttaProfileSlug:   "VisionAI",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go: nil,
				domain.Rust: []string{
					"https://www.linkedin.com/jobs/view/4020879940/",
				},
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
			},
			ShortDescription:          "",
			Industries:                []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{},
		},
		{
			ID:   0,  // system
			Type: "", // system
			Name: "Vector Atomic",
			URL:  "https://www.vectoratomic.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    18630145,
				Alias: "vectoratomic",
				Name:  "Vector Atomic",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "",
				GoRepositoryCount: 0,
				Verified:          false,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Vector-Atomic-EI_IE3378320.11,24.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Vector-Atomic-Reviews-E3378320.htm",
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go: nil,
				domain.Rust: []string{
					"https://www.linkedin.com/jobs/view/4023783229/",
				},
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
			},
			ShortDescription:          "",
			Industries:                []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{},
		},
		{
			ID:   0,                         // system
			Type: domain.CompanyTypeStartup, // system
			Name: "Freeform",
			URL:  "https://freeform.co/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    19207744,
				Alias: "freeformfuture",
				Name:  "Freeform",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "",
				GoRepositoryCount: 0,
				Verified:          false,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Freeform-Future-EI_IE8374898.11,26.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Freeform-Future-Reviews-E8374898.htm",
			},
			OttaProfileSlug:   "Freeform",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go: nil,
				domain.Rust: []string{
					"https://www.linkedin.com/jobs/view/4023717146/",
				},
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
			},
			ShortDescription:          "Metal 3D printing factories",
			Industries:                []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{},
		},
		{
			ID:   0,  // system
			Type: "", // system
			Name: "Daedalean AI",
			URL:  "https://daedalean.ai/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    10999325,
				Alias: "daedalean",
				Name:  "Daedalean AI",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "daedaleanai",
				GoRepositoryCount: 14,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Daedalean-EI_IE3150803.11,20.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Daedalean-Reviews-E3150803.htm",
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go: nil,
				domain.Rust: []string{
					"https://www.linkedin.com/jobs/view/4029289382/",
				},
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
			},
			ShortDescription: "Autonomous flight control for the aircraft",
			Industries:       []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
			},
		},
		{
			ID:   0,  // system
			Type: "", // system
			Name: "Statista",
			URL:  "https://www.statista.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:       0,
				IDs:      []int{1042846, 42875927},
				Alias:    "statista",
				Name:     "Statista",
				Verified: true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "",
				GoRepositoryCount: 0,
				Verified:          false,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Statista-EI_IE800158.11,19.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Statista-Reviews-E800158.htm",
				Verified:    true,
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go: nil,
				domain.Rust: []string{
					"https://www.linkedin.com/jobs/view/4031949216/",
				},
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
			},
			ShortDescription: "Business-data platform",
			Industries:       []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
				domain.Czechia,
			},
		},
		{
			ID:   0,  // system
			Type: "", // system
			Name: "DroneSense",
			URL:  "https://www.dronesense.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    10128253,
				Alias: "dronesense",
				Name:  "DroneSense",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "",
				GoRepositoryCount: 0,
				Verified:          false,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-DroneSense-EI_IE1830147.11,21.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/DroneSense-Reviews-E1830147.htm",
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go: nil,
				domain.Rust: []string{
					"https://www.linkedin.com/jobs/view/4035121961/",
				},
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
			},
			ShortDescription:          "Drone software platform",
			Industries:                []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{},
		},
		{
			ID:   0,  // system
			Type: "", // system
			Name: "Zurich Instruments",
			URL:  "https://www.zhinst.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    2289023,
				Alias: "zurich-instruments-ag",
				Name:  "Zurich Instruments",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "zhinst",
				GoRepositoryCount: 0,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Zurich-Instruments-AG-EI_IE3109985.11,32.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Zurich-Instruments-AG-Reviews-E3109985.htm",
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go: nil,
				domain.Rust: []string{
					"https://www.linkedin.com/jobs/view/4033619128/",
				},
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
			},
			ShortDescription: "Instrumentation for Quantum Computing and Periodic Signal Measurement",
			Industries:       []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
				domain.Czechia,
			},
		},
		{
			ID:   0,  // system
			Type: "", // system
			Name: "Palo Alto Networks",
			URL:  "https://www.paloaltonetworks.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    30086,
				Alias: "palo-alto-networks",
				Name:  "Palo Alto Networks",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "paloaltonetworks",
				GoRepositoryCount: 46,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Palo-Alto-Networks-EI_IE115142.11,29.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Palo-Alto-Networks-Reviews-E115142.htm",
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go: []string{
					"https://www.linkedin.com/jobs/view/4033692599/",
				},
				domain.Rust: []string{
					"https://www.linkedin.com/jobs/view/4035556879/",
				},
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
			},
			ShortDescription: "",
			Industries: []domain.Industry{
				domain.IndustryCyberSecurity,
			},
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
				domain.Czechia,
			},
		},
		{
			ID:   0,  // system
			Type: "", // system
			Name: "SciTec",
			URL:  "https://scitec.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:       22302976,
				Alias:    "scitecinc",
				Name:     "SciTec",
				Verified: true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "",
				GoRepositoryCount: 0,
				Verified:          false,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-SciTec-EI_IE1000832.11,17.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/SciTec-Reviews-E1000832.htm",
				Verified:    true,
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go: nil,
				domain.Rust: []string{
					"https://www.linkedin.com/jobs/view/4018720987/",
				},
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
			},
			ShortDescription: "",
			Industries: []domain.Industry{
				domain.IndustryCyberSecurity,
			},
			HasEmployeesFromCountries: []domain.Country{},
		},
		{
			ID:   0,  // system
			Type: "", // system
			Name: "Stackx.me",
			URL:  "https://stackx.me/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    71717406,
				Alias: "stack-x-me",
				Name:  "Stackx.me",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "",
				GoRepositoryCount: 0,
				Verified:          false,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "",
				ReviewsURL:  "",
				Verified:    false,
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go: nil,
				domain.Rust: []string{
					"https://www.linkedin.com/jobs/view/4034993832/",
				},
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
			},
			ShortDescription: "Social investment app",
			Industries: []domain.Industry{
				domain.IndustryFinTech,
			},
			HasEmployeesFromCountries: []domain.Country{},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Walt Disney",
			Website: "https://www.disneycareers.com/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:       1292,
				Alias:    "the-walt-disney-company",
				Name:     "The Walt Disney Company",
				Verified: true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "disney",
				GoRepositoryCount: 1,
			},
			BlindProfile:     domain.BlindProfile{},
			LevelsFyiProfile: domain.LevelsFyiProfile{},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Walt-Disney-Company-EI_IE717.11,30.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Walt-Disney-Company-Reviews-E717.htm",
			},
			IndeedProfile:     domain.IndeedProfile{},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go: {},
				domain.Rust: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Senior Software Engineer (Rust Developer)",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4071554607/",
							Date:             mustDate("2024-11-07"),
						},
					},
				},
				domain.Zig: {},
				domain.Scala: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Senior Software Engineer (Scala)",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4063482617/",
							Date:             mustDate("2024-11-01"),
						},
					},
				},
				domain.Elixir:  {},
				domain.Clojure: {},
				domain.Haskell: {},
			},
			ShortDescription: "",
			DealroomURL:      "",
			CrunchbaseURL:    "",
			PitchbookURL:     "",
			YahooFinanceURL:  "",
			GoogleFinanceURL: "",
			YCombinatorURL:   "",
			Industries: []domain.Industry{
				domain.IndustryEntertainment,
			},
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
				domain.Czechia,
			},
		},
		{
			ID:   0,  // system
			Type: "", // system
			Name: "Materialise",
			URL:  "https://www.materialise.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    4463,
				Alias: "materialise",
				Name:  "Materialise",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "",
				GoRepositoryCount: 0,
				Verified:          false,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Materialise-EI_IE223927.11,22.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Materialise-Reviews-E223927.htm",
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go: nil,
				domain.Rust: []string{
					"https://www.linkedin.com/jobs/view/3948655227/",
				},
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
			},
			ShortDescription: "",
			Industries:       []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
				domain.Czechia,
			},
		},
		{
			ID:   0,                         // system
			Type: domain.CompanyTypeStartup, // system
			Name: "InfluxData",
			URL:  "https://www.influxdata.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    5159145,
				Alias: "influxdb",
				Name:  "InfluxData",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "influxdata",
				GoRepositoryCount: 72,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-InfluxData-EI_IE1402855.11,21.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/InfluxData-Reviews-E1402855.htm",
			},
			OttaProfileSlug:   "InfluxData",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go: nil,
				domain.Rust: []string{
					"https://www.linkedin.com/jobs/view/3978294129/",
				},
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
			},
			ShortDescription: "Time series database",
			Industries: []domain.Industry{
				domain.IndustryDevOps,
			},
			HasEmployeesFromCountries: []domain.Country{},
		},
		{
			ID:   0,  // system
			Type: "", // system
			Name: "amo",
			URL:  "https://amo.co/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    92897504,
				Alias: "amoamoamo",
				Name:  "amo",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "",
				GoRepositoryCount: 0,
				Verified:          false,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "",
				ReviewsURL:  "",
				Verified:    false,
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go: nil,
				domain.Rust: []string{
					"https://www.linkedin.com/jobs/view/4038779132/",
				},
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
			},
			ShortDescription:          "Social Networking Platforms",
			Industries:                []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{},
		},
		{
			ID:   0,  // system
			Type: "", // system
			Name: "Sony Interactive Entertainment",
			URL:  "http://www.sonyinteractive.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    57993393,
				Alias: "sony-interactive-entertainment-llc",
				Name:  "Sony Interactive Entertainment",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "",
				GoRepositoryCount: 0,
				Verified:          false,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Sony-Interactive-Entertainment-EI_IE5580180.11,41.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Sony-Interactive-Entertainment-Reviews-E5580180.htm",
			},
			OttaProfileSlug:   "PlayStation",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go: nil,
				domain.Rust: []string{
					"https://www.linkedin.com/jobs/view/4022259692/",
				},
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
			},
			ShortDescription: "PlayStation",
			Industries: []domain.Industry{
				domain.IndustryEntertainment,
			},
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
				domain.Czechia,
			},
		},
		{
			ID:   0,  // system
			Type: "", // system
			Name: "Datadog",
			URL:  "https://www.datadoghq.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    1066442,
				Alias: "datadog",
				Name:  "Datadog",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "datadog",
				GoRepositoryCount: 259,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Datadog-EI_IE762009.11,18.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Datadog-Reviews-E762009.htm",
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go: []string{
					"https://www.linkedin.com/jobs/view/4030885124/",
				},
				domain.Rust: []string{
					"https://www.linkedin.com/jobs/view/3967138677/",
				},
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
			},
			ShortDescription: "Datadog provides cloud-scale monitoring and security for metrics, traces and logs in one unified platform",
			Industries: []domain.Industry{
				domain.IndustryDevOps,
			},
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
				domain.Czechia,
			},
		},
		{
			ID:   0,               // system
			Type: "",              // system
			Name: "Genius Sports", // "Second Spectrum" somehow related to "Genius Sports"
			URL:  "https://www.geniussports.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    22208562,
				Alias: "geniussports",
				Name:  "Genius Sports",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "",
				GoRepositoryCount: 0,
				Verified:          false,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Genius-Sports-EI_IE769838.11,24.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Genius-Sports-Reviews-E769838.htm",
			},
			OttaProfileSlug:   "Genius-Sports",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go: nil,
				domain.Rust: []string{
					"https://www.linkedin.com/jobs/view/4039692138/",
				},
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
			},
			ShortDescription: "",
			Industries:       []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
				domain.Czechia,
			},
		},
		{
			ID:   0,  // system
			Type: "", // system
			Name: "Windmill",
			URL:  "https://www.windmill.dev/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    82268299,
				Alias: "windmill-dev",
				Name:  "Windmill",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "windmill-labs",
				GoRepositoryCount: 1,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "",
				ReviewsURL:  "",
				Verified:    false,
			},
			OttaProfileSlug:   "Windmill",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go: nil,
				domain.Rust: []string{
					"https://www.linkedin.com/jobs/view/3992496369/",
				},
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
			},
			ShortDescription: "OSS self-hostable developer platform for APIs, background jobs, workflows and UIs",
			Industries: []domain.Industry{
				domain.IndustryDevOps,
			},
			HasEmployeesFromCountries: []domain.Country{},
		},
		{
			ID:   0,                         // system
			Type: domain.CompanyTypeStartup, // system
			Name: "Fetcherr",
			URL:  "https://fetcherr.io",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    31454791,
				Alias: "fetcherr-ltd",
				Name:  "Fetcherr",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "",
				GoRepositoryCount: 0,
				Verified:          false,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Fetcherr-EI_IE7854340.11,19.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Fetcherr-Reviews-E7854340.htm",
			},
			OttaProfileSlug:   "Fetcherr",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go: nil,
				domain.Rust: []string{
					"https://www.linkedin.com/jobs/view/4043996376/",
				},
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
			},
			ShortDescription:          "AI-driven solutions for the airline industry",
			Industries:                []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{},
		},
		{
			ID:   0,  // system
			Type: "", // system
			Name: "Flok Health",
			URL:  "https://flok.health/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    84475363,
				Alias: "flok-health",
				Name:  "Flok Health",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "",
				GoRepositoryCount: 0,
				Verified:          false,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "",
				ReviewsURL:  "",
				Verified:    false,
			},
			OttaProfileSlug:   "Flok-Health",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go: nil,
				domain.Rust: []string{
					"https://www.linkedin.com/jobs/view/4047041896/",
				},
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
			},
			ShortDescription: "Digital-native healthcare provider",
			Industries: []domain.Industry{
				domain.IndustryHealthTech,
			},
			HasEmployeesFromCountries: []domain.Country{},
		},
		{
			ID:   0,  // system
			Type: "", // system
			Name: "Rhebo",
			URL:  "https://rhebo.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    5312885,
				Alias: "rhebo",
				Name:  "Rhebo",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "",
				GoRepositoryCount: 0,
				Verified:          false,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "",
				ReviewsURL:  "",
				Verified:    false,
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go: nil,
				domain.Rust: []string{
					"https://www.linkedin.com/jobs/view/4050077229/",
				},
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
			},
			ShortDescription: "",
			Industries: []domain.Industry{
				domain.IndustryCyberSecurity,
			},
			HasEmployeesFromCountries: []domain.Country{},
		},
		{
			ID:   0,  // system
			Type: "", // system
			Name: "Bjak",
			URL:  "https://bjak.my/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    14607755,
				Alias: "bjak",
				Name:  "Bjak",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "",
				GoRepositoryCount: 0,
				Verified:          false,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Bjak-EI_IE3055055.11,15.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Bjak-Reviews-E3055055.htm",
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go: nil,
				domain.Rust: []string{
					"https://www.linkedin.com/jobs/view/4046735639/",
				},
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
			},
			ShortDescription:          "",
			Industries:                []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{},
		},
		{
			ID:   0,  // system
			Type: "", // system
			Name: "EverCharge",
			URL:  "https://evercharge.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    4817778,
				Alias: "evercharge",
				Name:  "EverCharge",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "",
				GoRepositoryCount: 0,
				Verified:          false,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "",
				ReviewsURL:  "",
				Verified:    false,
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go: nil,
				domain.Rust: []string{
					"https://www.linkedin.com/jobs/view/4040225807/",
				},
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
			},
			ShortDescription: "Electric vehicle (EV) charging",
			Industries:       []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
			},
		},
		{
			ID:   0,  // system
			Type: "", // system
			Name: "Ebury",
			URL:  "https://ebury.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    963919,
				Alias: "eburyfintech",
				Name:  "Ebury",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "",
				GoRepositoryCount: 0,
				Verified:          false,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Ebury-Partners-EI_IE823195.11,25.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Ebury-Partners-Reviews-E823195.htm",
			},
			OttaProfileSlug:   "Ebury",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go: []string{
					"https://www.linkedin.com/jobs/view/3990418832/",
					"https://www.linkedin.com/jobs/view/3990418832/",
				},
				domain.Rust:    nil,
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
			},
			ShortDescription: "Ebury is a hyper-growth FinTech firm",
			Industries: []domain.Industry{
				domain.IndustryFinTech,
			},
			HasEmployeesFromCountries: []domain.Country{
				domain.Czechia,
			},
		},
		{
			ID:   0,  // system
			Type: "", // system
			Name: "Talon.One",
			URL:  "https://talon.one",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    10393857,
				Alias: "talon.one",
				Name:  "Talon.One",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "talon-one",
				GoRepositoryCount: 18,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Talon-One-EI_IE2176357.11,20.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Talon-One-Reviews-E2176357.htm",
			},
			OttaProfileSlug:   "Talon-One",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go: []string{
					"https://www.linkedin.com/jobs/view/3945724263/",
				},
				domain.Rust:    nil,
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
			},
			ShortDescription: "Loyalty and promotion engine for enterprises",
			Industries:       []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
			},
		},
		{
			ID:   0,                         // system
			Type: domain.CompanyTypeStartup, // system
			Name: "NewHomesMate",
			URL:  "https://newhomesmate.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    86332598,
				Alias: "newhomesmate",
				Name:  "NewHomesMate",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "",
				GoRepositoryCount: 0,
				Verified:          false,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-NewHomesMate-EI_IE8164563.11,23.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/NewHomesMate-Reviews-E8164563.htm",
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go: []string{
					"https://jobs.dou.ua/companies/propertymate/vacancies/264309/",
				},
				domain.Rust:    nil,
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
			},
			ShortDescription: "Platform to search, compare, and purchase newly constructed homes",
			Industries: []domain.Industry{
				domain.IndustryPropTech,
			},
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
			},
		},
		{
			ID:   0,                         // system
			Type: domain.CompanyTypeStartup, // system
			Name: "Density",
			URL:  "https://www.density.io/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    7589398,
				Alias: "density-inc-",
				Name:  "Density",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "",
				GoRepositoryCount: 0,
				Verified:          false,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Density-EI_IE1627818.11,18.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Density-Reviews-E1627818.htm",
			},
			OttaProfileSlug:   "Density",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go: []string{
					"https://app.welcometothejungle.com/jobs/t7vxWHyI",
				},
				domain.Rust:    nil,
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
			},
			ShortDescription:          "Analytics platform for measuring and improving workplaces",
			Industries:                []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{},
		},
		{
			ID:   0,  // system
			Type: "", // system
			Name: "Surfly",
			URL:  "https://www.surfly.com",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    1943587,
				Alias: "surfly",
				Name:  "Surfly",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "surfly",
				GoRepositoryCount: 3,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Surfly-EI_IE1265203.11,17.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Surfly-Reviews-E1265203.htm",
			},
			OttaProfileSlug:   "Surfly",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go: []string{
					"https://www.linkedin.com/jobs/view/3931538747/",
					"https://app.welcometothejungle.com/jobs/M6OnKOSN",
				},
				domain.Rust:    nil,
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
			},
			ShortDescription:          "Co-browsing middleware provider",
			Industries:                []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{},
		},
		{
			ID:   0,  // system
			Type: "", // system
			Name: "Cross River",
			URL:  "https://www.crossriver.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:       967395,
				Alias:    "cross-river-bank",
				Name:     "Cross River",
				Verified: true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "",
				GoRepositoryCount: 0,
				Verified:          false,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Cross-River-Bank-EI_IE1177112.11,27.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Cross-River-Bank-Reviews-E1177112.htm",
				Verified:    true,
			},
			OttaProfileSlug:   "Cross-River",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go: []string{
					"https://www.linkedin.com/jobs/view/4015069840/",
					"https://app.welcometothejungle.com/jobs/gGYEvKLd",
				},
				domain.Rust:    nil,
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
			},
			ShortDescription: "Bank",
			Industries: []domain.Industry{
				domain.IndustryFinTech,
			},
			HasEmployeesFromCountries: []domain.Country{},
		},
		{
			ID:   0,                         // system
			Type: domain.CompanyTypeStartup, // system
			Name: "Abnormal Security",
			URL:  "https://abnormalsecurity.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    18586257,
				Alias: "abnormalsecurity",
				Name:  "Abnormal Security",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "abnormal-security",
				GoRepositoryCount: 0,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Abnormal-Security-EI_IE3146005.11,28.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Abnormal-Security-Reviews-E3146005.htm",
			},
			OttaProfileSlug:   "Abnormal-Security",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go: []string{
					"https://app.welcometothejungle.com/jobs/pRJX4yTR",
				},
				domain.Rust:    nil,
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
			},
			ShortDescription: "Keeps your email protected",
			Industries: []domain.Industry{
				domain.IndustryCyberSecurity,
			},
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
			},
		},
		{
			ID:   0,                         // system
			Type: domain.CompanyTypeStartup, // system
			Name: "Scope3",
			URL:  "https://scope3.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    80943372,
				Alias: "scope3data",
				Name:  "Scope3",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "scope3data",
				GoRepositoryCount: 0,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "",
				ReviewsURL:  "",
				Verified:    false,
			},
			OttaProfileSlug:   "Scope3",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go: []string{
					"https://app.welcometothejungle.com/jobs/LhvHDkyc",
				},
				domain.Rust:    nil,
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
			},
			ShortDescription:          "Sustainability platform for media & advertising",
			Industries:                []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{},
		},
		{
			ID:   0,                         // system
			Type: domain.CompanyTypeStartup, // system
			Name: "Apollo GraphQL",
			URL:  "https://www.apollographql.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    28602935,
				Alias: "apollo-graphql",
				Name:  "Apollo GraphQL",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "apollographql",
				GoRepositoryCount: 1,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Apollo-GraphQL-EI_IE893757.11,25.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Apollo-GraphQL-Reviews-E893757.htm",
			},
			OttaProfileSlug:   "Apollo",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go: nil,
				domain.Rust: []string{
					"https://www.linkedin.com/jobs/view/4043213870/",
				},
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
			},
			ShortDescription: "Deliver GraphQL at any scale",
			Industries: []domain.Industry{
				domain.IndustryDevOps,
			},
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
			},
		},
		{
			ID:   0,  // system
			Type: "", // system
			Name: "Paessler",
			URL:  "https://www.paessler.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    409588,
				Alias: "paessler-gmbh",
				Name:  "Paessler",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "PaesslerAG",
				GoRepositoryCount: 2,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Paessler-GmbH-EI_IE1430882.11,24.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Paessler-GmbH-Reviews-E1430882.htm",
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go: nil,
				domain.Rust: []string{
					"https://www.linkedin.com/jobs/view/4039210719/",
				},
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
			},
			ShortDescription: "",
			Industries: []domain.Industry{
				domain.IndustryDevOps,
			},
			HasEmployeesFromCountries: []domain.Country{},
		},
		{
			ID:   0,                         // system
			Type: domain.CompanyTypeStartup, // system
			Name: "Hasura",
			URL:  "https://hasura.io/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    13211008,
				Alias: "hasura",
				Name:  "Hasura",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "hasura",
				GoRepositoryCount: 33,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Hasura-EI_IE1451757.11,17.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Hasura-Reviews-E1451757.htm",
			},
			OttaProfileSlug:   "Hasura",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go: []string{
					"https://www.linkedin.com/jobs/view/4026175934/",
				},
				domain.Rust: []string{
					"https://www.linkedin.com/jobs/view/4036409356/",
				},
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
			},
			ShortDescription: "API platform, built for data delivery",
			Industries: []domain.Industry{
				domain.IndustryDevOps,
			},
			HasEmployeesFromCountries: []domain.Country{},
		},
		{
			ID:   0,                         // system
			Type: domain.CompanyTypeStartup, // system
			Name: "Aira",
			URL:  "https://company.airahome.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    93205887,
				Alias: "airahome",
				Name:  "Aira",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "",
				GoRepositoryCount: 0,
				Verified:          false,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Aira-EI_IE8976549.11,15.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Aira-Reviews-E8976549.htm",
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go: nil,
				domain.Rust: []string{
					"https://www.linkedin.com/jobs/view/4039345477/",
				},
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
			},
			ShortDescription:          "Clean energy-tech accessible and affordable for millions of homes",
			Industries:                []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "DDN",
			Website: "https://www.ddn.com/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:       21665,
				Alias:    "ddn", // before ddn-storage
				Name:     "DDN",
				Verified: true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "DDNStorage",
				GoRepositoryCount: 3,
			},
			BlindProfile:     domain.BlindProfile{},
			LevelsFyiProfile: domain.LevelsFyiProfile{},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-DDN-Storage-EI_IE262933.11,22.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/DDN-Storage-Reviews-E262933.htm",
			},
			IndeedProfile:     domain.IndeedProfile{},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go: {},
				domain.Rust: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Senior Development Engineer — RUST",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4064738212/",
							Date:             mustDate("2024-11-02"),
						},
					},
				},
				domain.Zig:     {},
				domain.Scala:   {},
				domain.Elixir:  {},
				domain.Clojure: {},
				domain.Haskell: {},
			},
			ShortDescription:          "Data Intelligence Platform",
			DealroomURL:               "",
			CrunchbaseURL:             "",
			PitchbookURL:              "",
			YahooFinanceURL:           "",
			GoogleFinanceURL:          "",
			YCombinatorURL:            "",
			Industries:                []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{},
		},
		{
			ID:   0,  // system
			Type: "", // system
			Name: "Starling Bank",
			URL:  "https://www.starlingbank.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    9180896,
				Alias: "starlingbank",
				Name:  "Starling Bank",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "starlingbank",
				GoRepositoryCount: 0,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Starling-Bank-EI_IE1337967.11,24.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Starling-Bank-Reviews-E1337967.htm",
			},
			OttaProfileSlug:   "Starling-Bank",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go: []string{
					"https://app.welcometothejungle.com/jobs/lfnyRzC0",
				},
				domain.Rust:    nil,
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
			},
			ShortDescription: "Mobile-first bank",
			Industries: []domain.Industry{
				domain.IndustryFinTech,
			},
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
			},
		},
		{
			ID:   0,  // system
			Type: "", // system
			Name: "Spice AI",
			URL:  "https://spice.ai/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    74148478,
				Alias: "spice-ai",
				Name:  "Spice AI",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "spiceai",
				GoRepositoryCount: 6,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "",
				ReviewsURL:  "",
				Verified:    false,
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go: []string{
					"https://app.welcometothejungle.com/jobs/WlVadVlu",
					"https://app.welcometothejungle.com/jobs/emdLdWlp",
				},
				domain.Rust: []string{
					"https://app.welcometothejungle.com/jobs/WlVadVlu",
					"https://app.welcometothejungle.com/jobs/emdLdWlp",
				},
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
			},
			ShortDescription:          "Blocks for data and time-series AI applications",
			Industries:                []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{},
		},
		{
			ID:   0,  // system
			Type: "", // system
			Name: "Algolia",
			URL:  "https://www.algolia.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    2728700,
				Alias: "algolia",
				Name:  "Algolia",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "algolia",
				GoRepositoryCount: 14,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Algolia-EI_IE998983.11,18.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Algolia-Reviews-E998983.htm",
			},
			OttaProfileSlug:   "Algolia",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go: []string{
					"https://app.welcometothejungle.com/jobs/tsBqRgc_",
					"https://www.linkedin.com/jobs/view/3984452580/",
				},
				domain.Rust:    nil,
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
			},
			ShortDescription: "End-to-end AI Search solution",
			Industries:       []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{
				domain.Czechia,
			},
		},
		{
			ID:   0,                         // system
			Type: domain.CompanyTypeStartup, // system
			Name: "Observe",
			URL:  "https://www.observeinc.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    18796376,
				Alias: "observe-inc",
				Name:  "Observe",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "observeinc",
				GoRepositoryCount: 21,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Observe-CA-EI_IE4567528.11,21.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Observe-CA-Reviews-E4567528.htm",
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go: []string{
					"https://app.welcometothejungle.com/jobs/bjD5ZhEh",
				},
				domain.Rust:    nil,
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
			},
			ShortDescription: "SaaS Observability means fewer incidents",
			Industries: []domain.Industry{
				domain.IndustryDevOps,
			},
			HasEmployeesFromCountries: []domain.Country{},
		},
		{
			ID:   0,  // system
			Type: "", // system
			Name: "Chalk",
			URL:  "https://chalk.ai/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    78829224,
				Alias: "chalkai",
				Name:  "Chalk",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "",
				GoRepositoryCount: 0,
				Verified:          false,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "",
				ReviewsURL:  "",
				Verified:    false,
			},
			OttaProfileSlug:   "Chalk",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go: nil,
				domain.Rust: []string{
					"https://app.welcometothejungle.com/jobs/2SMONxFZ",
				},
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
			},
			ShortDescription:          "Data platform for machine learning",
			Industries:                []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{},
		},
		{
			ID:   0,  // system
			Type: "", // system
			Name: "Enova",
			URL:  "https://www.enova.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    670584,
				Alias: "enova-international",
				Name:  "Enova",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "enova",
				GoRepositoryCount: 7,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Enova-EI_IE298072.11,16.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Enova-Reviews-E298072.htm",
			},
			OttaProfileSlug:   "Enova",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go: []string{
					"https://app.welcometothejungle.com/jobs/6TIHCD3L",
				},
				domain.Rust:    nil,
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
			},
			ShortDescription: "Financial credit provider",
			Industries: []domain.Industry{
				domain.IndustryFinTech,
			},
			HasEmployeesFromCountries: []domain.Country{},
		},
		{
			ID:   0,  // system
			Type: "", // system
			Name: "JustWatch",
			URL:  "https://media.justwatch.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    9216347,
				Alias: "justwatch",
				Name:  "JustWatch",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "justwatch",
				GoRepositoryCount: 7,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-JustWatch-EI_IE1348307.11,20.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/JustWatch-Reviews-E1348307.htm",
			},
			OttaProfileSlug:   "JustWatch-1",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go: []string{
					"https://app.welcometothejungle.com/jobs/2BshPA1c",
				},
				domain.Rust:    nil,
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
			},
			ShortDescription: "Cross platform streaming guide",
			Industries:       []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
			},
		},
		{
			ID:   0,                         // system
			Type: domain.CompanyTypeStartup, // system
			Name: "Sourcegraph",
			URL:  "https://sourcegraph.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    4803356,
				Alias: "sourcegraph",
				Name:  "Sourcegraph",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "sourcegraph",
				GoRepositoryCount: 190,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Sourcegraph-EI_IE1356770.11,22.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Sourcegraph-Reviews-E1356770.htm",
			},
			OttaProfileSlug:   "Sourcegraph",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go: []string{
					"https://app.welcometothejungle.com/jobs/wTAsZImo",
				},
				domain.Rust:    nil,
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
			},
			ShortDescription:          "Universal code search and code intelligence for developers",
			Industries:                []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{},
		},
		{
			ID:   0,  // system
			Type: "", // system
			Name: "Plaid",
			URL:  "https://plaid.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    2684737,
				Alias: "plaid-",
				Name:  "Plaid",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "plaid",
				GoRepositoryCount: 9,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Plaid-EI_IE1156368.11,16.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Plaid-Reviews-E1156368.htm",
			},
			OttaProfileSlug:   "Plaid",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go: []string{
					"https://app.welcometothejungle.com/jobs/7_pckxdd",
				},
				domain.Rust:    nil,
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
			},
			ShortDescription: "Infrastructure and APIs for FinTech",
			Industries: []domain.Industry{
				domain.IndustryFinTech,
			},
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
			},
		},
		{
			ID:   0,  // system
			Type: "", // system
			Name: "Coalition",
			URL:  "https://www.coalitioninc.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    11193112,
				Alias: "coalitioninc",
				Name:  "Coalition",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "",
				GoRepositoryCount: 0,
				Verified:          false,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Coalition-EI_IE1717118.11,20.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Coalition-Reviews-E1717118.htm",
			},
			OttaProfileSlug:   "Coalition",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go: []string{
					"https://app.welcometothejungle.com/jobs/7o-foHYQ",
				},
				domain.Rust:    nil,
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
			},
			ShortDescription:          "Cyber insurance and security for businesses",
			Industries:                []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{},
		},
		{
			ID:   0,  // system
			Type: "", // system
			Name: "Mechanical Orchard",
			URL:  "https://www.mechanical-orchard.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    89798793,
				Alias: "mechanical-orchard",
				Name:  "Mechanical Orchard",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "mechanical-orchard",
				GoRepositoryCount: 0, // Elixir 9
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "",
				ReviewsURL:  "",
				Verified:    false,
			},
			OttaProfileSlug:   "Mechanical-Orchard",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go:    nil,
				domain.Rust:  nil,
				domain.Zig:   nil,
				domain.Scala: nil,
				domain.Elixir: []string{
					"https://app.welcometothejungle.com/jobs/zLruxuAk",
					"https://app.welcometothejungle.com/jobs/h2lbAL5a",
				},
				domain.Clojure: nil,
				domain.Haskell: nil,
			},
			ShortDescription:          "",
			Industries:                []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{},
		},
		{
			ID:   0,  // system
			Type: "", // system
			Name: "Moveworks",
			URL:  "https://www.moveworks.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    18311685,
				Alias: "moveworksai",
				Name:  "Moveworks",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "moveworks",
				GoRepositoryCount: 0,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Moveworks-EI_IE1730936.11,20.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Moveworks-Reviews-E1730936.htm",
			},
			OttaProfileSlug:   "Moveworks",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go: []string{
					"https://www.linkedin.com/jobs/view/3923093175/",
					"https://app.welcometothejungle.com/jobs/X1IQMOCK",
				},
				domain.Rust:    nil,
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
			},
			ShortDescription:          "Enterprise-built conversational AI platform",
			Industries:                []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{},
		},
		{
			ID:   0,                         // system
			Type: domain.CompanyTypeStartup, // system
			Name: "Emitwise",
			URL:  "https://emitwise.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    20511470,
				Alias: "emitwise",
				Name:  "Emitwise",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "emitwise",
				GoRepositoryCount: 0,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Emitwise-EI_IE3915102.11,19.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Emitwise-Reviews-E3915102.htm",
			},
			OttaProfileSlug:   "Emitwise",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go: []string{
					"https://app.welcometothejungle.com/jobs/w8X1TqCu",
				},
				domain.Rust:    nil,
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
			},
			ShortDescription:          "Carbon accounting platform",
			Industries:                []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{},
		},
		{
			ID:   0,  // system
			Type: "", // system
			Name: "Upvest",
			URL:  "https://upvest.co/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    18388245,
				Alias: "upvest",
				Name:  "Upvest",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "upvestco",
				GoRepositoryCount: 3,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Upvest-EI_IE2948356.11,17.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Upvest-Reviews-E2948356.htm",
			},
			OttaProfileSlug:   "Upvest",
			YouTubeChannelURL: "",
			GoMainLanguage:    true, // https://app.welcometothejungle.com/jobs/JTNWvpsj Work with cutting-edge technologies (Go is the primary language) without a legacy codebase
			Vacancies: domain.Vacancies{
				domain.Go: []string{
					"https://app.welcometothejungle.com/jobs/JTNWvpsj",
				},
				domain.Rust:    nil,
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
			},
			ShortDescription: "The Investment API",
			Industries: []domain.Industry{
				domain.IndustryFinTech,
			},
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
			},
		},
		{
			ID:   0,  // system
			Type: "", // system
			Name: "CoreWeave",
			URL:  "https://www.coreweave.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    36121341,
				Alias: "coreweave",
				Name:  "CoreWeave",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "coreweave",
				GoRepositoryCount: 44,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-CoreWeave-EI_IE5711823.11,20.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/CoreWeave-Reviews-E5711823.htm",
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go: []string{
					"https://www.linkedin.com/jobs/view/4029510008/",
					"https://app.welcometothejungle.com/jobs/-02MRxuv",
				},
				domain.Rust:    nil,
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
			},
			ShortDescription: "Specialized cloud provider",
			Industries: []domain.Industry{
				domain.IndustryDevOps,
			},
			HasEmployeesFromCountries: []domain.Country{},
		},
		{
			ID:   0,  // system
			Type: "", // system
			Name: "Neo4j",
			URL:  "https://www.neo4j.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    828370,
				Alias: "neo4j",
				Name:  "Neo4j",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "neo4j",
				GoRepositoryCount: 3,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Neo4j-EI_IE665767.11,16.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Neo4j-Reviews-E665767.htm",
			},
			OttaProfileSlug:   "Neo4j",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go: []string{
					"https://app.welcometothejungle.com/jobs/xOagXoE8",
				},
				domain.Rust:    nil,
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
			},
			ShortDescription: "Graph-based data analysis",
			Industries: []domain.Industry{
				domain.IndustryDevOps,
			},
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
				domain.Czechia,
			},
		},
		{
			ID:   0,  // system
			Type: "", // system
			Name: "Jack Henry",
			URL:  "https://www.jackhenry.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:       6970,
				Alias:    "jack-henry",
				Name:     "Jack Henry",
				Verified: true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "jkhy",
				GoRepositoryCount: 0,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Jack-Henry-and-Associates-EI_IE1543.11,36.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Jack-Henry-and-Associates-Reviews-E1543.htm",
				Verified:    true, // openCompany ???
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go: []string{
					"https://www.linkedin.com/jobs/view/4057290510/", // Professional experience programming in Golang
					"https://www.linkedin.com/jobs/view/4044914356/",
				},
				domain.Rust:    nil,
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
			},
			ShortDescription: "Financial technology company that strengthens connections between people and their financial institutions",
			Industries: []domain.Industry{
				domain.IndustryFinTech,
			},
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
				domain.Czechia,
			},
		},
		{
			ID:   0,  // system
			Type: "", // system
			Name: "GoCardless",
			URL:  "https://gocardless.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    2808012,
				Alias: "gocardless",
				Name:  "GoCardless",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "gocardless",
				GoRepositoryCount: 24,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-GoCardless-EI_IE1001543.11,21.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/GoCardless-Reviews-E1001543.htm",
			},
			OttaProfileSlug:   "GoCardless",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go: []string{
					"https://app.welcometothejungle.com/jobs/xxfBLoSU",
				},
				domain.Rust:    nil,
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
			},
			ShortDescription: "",
			Industries: []domain.Industry{
				domain.IndustryFinTech,
			},
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
				domain.Czechia,
			},
		},
		{
			ID:   0,  // system
			Type: "", // system
			Name: "Zalando",
			URL:  "https://jobs.zalando.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:       698916,
				Alias:    "zalando",
				Name:     "Zalando",
				Verified: true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "zalando",
				GoRepositoryCount: 5,
				Verified:          false,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Zalando-EI_IE613421.11,18.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Zalando-Reviews-E613421.htm",
				Verified:    true,
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go: []string{
					"https://www.linkedin.com/jobs/view/3959001476/",
				},
				domain.Rust: nil,
				domain.Zig:  nil,
				domain.Scala: []string{
					"https://www.linkedin.com/jobs/view/4055507332/",
				},
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
			},
			ShortDescription: "Online fashion platform",
			Industries: []domain.Industry{
				domain.IndustryECommerce,
			},
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
				domain.Czechia,
			},
		},
		{
			ID:   0,  // system
			Type: "", // system
			Name: "FreeWheel",
			URL:  "https://www.freewheel.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    458871,
				Alias: "freewheel",
				Name:  "FreeWheel",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "freewheel",
				GoRepositoryCount: 0,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-FreeWheel-EI_IE313920.11,20.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/FreeWheel-Reviews-E313920.htm",
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go: []string{
					"https://www.linkedin.com/jobs/view/4008715409/",
				},
				domain.Rust:    nil,
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
			},
			ShortDescription: "Platform makes TV and video advertising work",
			Industries: []domain.Industry{
				domain.IndustryAdTech,
			},
			HasEmployeesFromCountries: []domain.Country{},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "GeoComply",
			Website: "https://www.geocomply.com/",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:       2457172,
				Alias:    "geocomply",
				Name:     "GeoComply",
				Verified: true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "GeoComply",
				GoRepositoryCount: 0,
			},
			BlindProfile:     domain.BlindProfile{},
			LevelsFyiProfile: domain.LevelsFyiProfile{},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-GeoComply-EI_IE2333177.11,20.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/GeoComply-Reviews-E2333177.htm",
			},
			IndeedProfile:     domain.IndeedProfile{},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Senior Software Engineer (Golang)",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4026401818/",
							Date:             mustDate("2024-11-01"),
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
			ShortDescription: "Fraud prevention and cybersecurity solutions",
			DealroomURL:      "",
			CrunchbaseURL:    "",
			PitchbookURL:     "",
			YahooFinanceURL:  "",
			GoogleFinanceURL: "",
			YCombinatorURL:   "",
			Industries: []domain.Industry{
				domain.IndustryCyberSecurity,
			},
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
			},
		},
		{
			ID:   0,  // system
			Type: "", // system
			Name: "Electronic Arts",
			URL:  "https://www.ea.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    1449,
				Alias: "electronic-arts",
				Name:  "Electronic Arts",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "electronicarts",
				GoRepositoryCount: 3,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Electronic-Arts-EI_IE1628.11,26.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Electronic-Arts-Reviews-E1628.htm",
			},
			OttaProfileSlug:   "Electronic-Arts",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go: []string{
					"https://www.linkedin.com/jobs/view/4043032750/",
				},
				domain.Rust:    nil,
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
			},
			ShortDescription: "",
			Industries: []domain.Industry{
				domain.IndustryEntertainment,
			},
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
				domain.Czechia,
			},
		},
	}
}
