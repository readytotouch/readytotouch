package db

import "github.com/readytotouch/readytotouch/internal/domain"

func companiesPart2() []domain.CompanyProfile {
	return []domain.CompanyProfile{
		{
			ID:   0,  // system
			Type: "", // system
			Name: "Trusting Social",
			URL:  "https://www.trustingsocial.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    3506386,
				Alias: "trustingsocial",
				Name:  "Trusting Social",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "",
				GoRepositoryCount: 0,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-TrustingSocial-EI_IE741535.11,25.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/TrustingSocial-Reviews-E741535.htm",
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go: []string{
					"https://www.linkedin.com/jobs/view/4042510486/",
				},
				domain.Rust:    nil,
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
			},
			ShortDescription: "AI fintech company revolutionizing credit scoring using big data technology",
			Industries: []domain.Industry{
				domain.IndustryFinTech,
			},
			HasEmployeesFromCountries: []domain.Country{},
		},
		{
			ID:   0,  // system
			Type: "", // system
			Name: "McAfee",
			URL:  "https://www.mcafee.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    2336,
				Alias: "mcafee",
				Name:  "McAfee",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "mcafee",
				GoRepositoryCount: 0,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-McAfee-EI_IE2244.11,17.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/McAfee-Reviews-E2244.htm",
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go: []string{
					"https://www.linkedin.com/jobs/view/4041686516/",
				},
				domain.Rust:    nil,
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
			},
			ShortDescription: "Antivirus software company",
			Industries: []domain.Industry{
				domain.IndustryCyberSecurity,
			},
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
				domain.Czechia,
			},
		},
		{
			ID:   0,                         // system
			Type: domain.CompanyTypeStartup, // system
			Name: "Hadrian",
			URL:  "https://www.hadrian.co/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    71668100,
				Alias: "hadrianspace",
				Name:  "Hadrian",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "",
				GoRepositoryCount: 0,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Hadrian-EI_IE5364733.11,18.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Hadrian-Reviews-E5364733.htm",
			},
			OttaProfileSlug:   "Hadrian-2",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go: []string{
					"https://app.welcometothejungle.com/jobs/63inUEkF",
				},
				domain.Rust:    nil,
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
			},
			ShortDescription:          "Building autonomous factories for aerospace components",
			Industries:                []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{},
		},
		{
			ID:   0,  // system
			Type: "", // system
			Name: "Scaleway",
			URL:  "https://www.scaleway.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    9427185,
				Alias: "scaleway",
				Name:  "Scaleway",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "scaleway",
				GoRepositoryCount: 23,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Scaleway-EI_IE1861876.11,19.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Scaleway-Reviews-E1861876.htm",
			},
			OttaProfileSlug:   "Scaleway",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go: []string{
					"https://app.welcometothejungle.com/jobs/dFlIR1ow",
				},
				domain.Rust:    nil,
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
			},
			ShortDescription:          "Cloud provider",
			Industries:                []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{},
		},
		{
			ID:   0,  // system
			Type: "", // system
			Name: "Demandbase",
			URL:  "https://www.demandbase.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    89759,
				Alias: "demandbase",
				Name:  "Demandbase",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "Demandbase",
				GoRepositoryCount: 1,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Demandbase-EI_IE271272.11,21.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Demandbase-Reviews-E271272.htm",
			},
			OttaProfileSlug:   "Demandbase",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go:   nil,
				domain.Rust: nil,
				domain.Zig:  nil,
				domain.Scala: []string{
					"https://app.welcometothejungle.com/jobs/sKrHcCx0",
				},
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
			},
			ShortDescription:          "Go-To-Market engagement & marketing software",
			Industries:                []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{},
		},
		{
			ID:   0,  // system
			Type: "", // system
			Name: "Overwolf",
			URL:  "https://www.overwolf.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    519584,
				Alias: "overwolf.com",
				Name:  "Overwolf",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "overwolf",
				GoRepositoryCount: 0,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Overwolf-EI_IE1963582.11,19.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Overwolf-Reviews-E1963582.htm",
			},
			OttaProfileSlug:   "Overwolf",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go: []string{
					"https://app.welcometothejungle.com/jobs/duQ9FWwK",
				},
				domain.Rust:    nil,
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
			},
			ShortDescription: "Gaming app & mod development framework",
			Industries: []domain.Industry{
				domain.IndustryGameDev,
			},
			HasEmployeesFromCountries: []domain.Country{},
		},
		{
			ID:   0,  // system
			Type: "", // system
			Name: "Bitfount",
			URL:  "https://www.bitfount.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    69183486,
				Alias: "bitfount",
				Name:  "Bitfount",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "bitfount",
				GoRepositoryCount: 0, // Rust 1
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "",
				ReviewsURL:  "",
			},
			OttaProfileSlug:   "Bitfount",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go: nil,
				domain.Rust: []string{
					"https://www.linkedin.com/jobs/view/4044732400/",
					"https://app.welcometothejungle.com/jobs/AGxChIIS",
				},
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
			},
			ShortDescription:          "Federated AI & distributed data science platform",
			Industries:                []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{},
		},
		{
			ID:   0,                         // system
			Type: domain.CompanyTypeStartup, // system
			Name: "Hinge",
			URL:  "https://hinge.co/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    2603651,
				Alias: "hinge-app",
				Name:  "Hinge",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "Hinge",
				GoRepositoryCount: 4,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Hinge-EI_IE871901.11,16.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Hinge-Reviews-E871901.htm",
			},
			OttaProfileSlug:   "Hinge",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go: []string{
					"https://www.linkedin.com/jobs/view/4043430809/",
					"https://app.welcometothejungle.com/jobs/V73cck97",
				},
				domain.Rust:    nil,
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
			},
			ShortDescription:          "Dating app for meaningful relationships",
			Industries:                []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{},
		},
		{
			ID:   0,                         // system
			Type: domain.CompanyTypeStartup, // system
			Name: "Firework",
			URL:  "https://firework.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    18700473,
				Alias: "fireworkhq",
				Name:  "Firework",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "",
				GoRepositoryCount: 0,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Firework-EI_IE2300706.11,19.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Firework-Reviews-E2300706.htm",
			},
			OttaProfileSlug:   "Firework",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go:    nil,
				domain.Rust:  nil,
				domain.Zig:   nil,
				domain.Scala: nil,
				domain.Elixir: []string{
					"https://app.welcometothejungle.com/jobs/srNX_4cO",
				},
				domain.Clojure: nil,
				domain.Haskell: nil,
			},
			ShortDescription: "Video-based eCommerce platform",
			Industries: []domain.Industry{
				domain.IndustryECommerce,
			},
			HasEmployeesFromCountries: []domain.Country{},
		},
		{
			ID:   0,                         // system
			Type: domain.CompanyTypeStartup, // system
			Name: "Breakroom",
			URL:  "https://www.breakroom.cc/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    18458137,
				Alias: "breakroom",
				Name:  "Breakroom",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "breakroom",
				GoRepositoryCount: 0, // Elixir 13
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Breakroom-EI_IE8002334.11,20.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Breakroom-Reviews-E8002334.htm",
			},
			OttaProfileSlug:   "Breakroom",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go:    nil,
				domain.Rust:  nil,
				domain.Zig:   nil,
				domain.Scala: nil,
				domain.Elixir: []string{
					"https://www.linkedin.com/jobs/view/4052987616/",
					"https://app.welcometothejungle.com/jobs/p_OpKzJ8",
				},
				domain.Clojure: nil,
				domain.Haskell: nil,
			},
			ShortDescription:          "The people-powered job comparison site",
			Industries:                []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{},
		},
		{
			ID:   0,                         // system
			Type: domain.CompanyTypeStartup, // system
			Name: "Ansa",
			URL:  "https://www.ansa.dev/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    91158109,
				Alias: "getansa",
				Name:  "Ansa",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "",
				GoRepositoryCount: 0,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "",
				ReviewsURL:  "",
			},
			OttaProfileSlug:   "Ansa",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go: []string{
					"https://app.welcometothejungle.com/jobs/EjVpa6gl",
				},
				domain.Rust:    nil,
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
			},
			ShortDescription:          "Branded customer wallet infrastructure",
			Industries:                []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{},
		},
		{
			ID:   0,  // system
			Type: "", // system
			Name: "1GLOBAL",
			URL:  "https://www.1global.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    39711,
				Alias: "1global",
				Name:  "1GLOBAL",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "",
				GoRepositoryCount: 0,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-1GLOBAL-EI_IE9805631.11,18.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/1GLOBAL-Reviews-E9805631.htm",
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go: []string{
					"https://www.linkedin.com/jobs/view/4048972208/",
				},
				domain.Rust:    nil,
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
			},
			ShortDescription: "eSIM",
			Industries:       []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
				domain.Czechia,
			},
		},
		{
			ID:   0,  // system
			Type: "", // system
			Name: "Eco-Movement",
			URL:  "https://eco-movement.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    15223273,
				Alias: "eco-movement-charge-point-data",
				Name:  "Eco-Movement",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "",
				GoRepositoryCount: 0,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Eco-Movement-EI_IE5353422.11,23.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Eco-Movement-Reviews-E5353422.htm",
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go: []string{
					"https://www.linkedin.com/jobs/view/4046511967/",
				},
				domain.Rust:    nil,
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
			},
			ShortDescription:          "Charging station data platform",
			Industries:                []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{},
		},
		{
			ID:   0,  // system
			Type: "", // system
			Name: "Network Optix",
			URL:  "https://www.networkoptix.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    1896041,
				Alias: "network-optix",
				Name:  "Network Optix",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "networkoptix",
				GoRepositoryCount: 0,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Network-Optix-EI_IE1103945.11,24.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Network-Optix-Reviews-E1103945.htm",
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go: []string{
					"https://www.linkedin.com/jobs/view/4028418131/",
				},
				domain.Rust:    nil,
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
			},
			ShortDescription: "Convert video data into practical insights for businesses and industries of all kinds",
			Industries:       []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
			},
		},
		{
			ID:   0,  // system
			Type: "", // system
			Name: "Cinemo",
			URL:  "https://www.cinemo.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    564661,
				Alias: "cinemo",
				Name:  "Cinemo",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "",
				GoRepositoryCount: 0,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Cinemo-EI_IE2189711.11,17.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Cinemo-Reviews-E2189711.htm",
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go: []string{
					"https://www.linkedin.com/jobs/view/4050001754/",
				},
				domain.Rust:    nil,
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
			},
			ShortDescription: "Make every screen an opportunity",
			Industries:       []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
				domain.Czechia,
			},
		},
		{
			ID:   0,  // system
			Type: "", // system
			Name: "Readdle",
			URL:  "https://readdle.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    629551,
				Alias: "readdle",
				Name:  "Readdle",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "readdle",
				GoRepositoryCount: 5,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Readdle-EI_IE613771.11,18.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Readdle-Reviews-E613771.htm",
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go: []string{
					"https://www.linkedin.com/jobs/view/3971228819/",
				},
				domain.Rust:    nil,
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
			},
			ShortDescription: "Creating apps",
			Industries:       []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
			},
		},
		{
			ID:   0,  // system
			Type: "", // system
			Name: "Ipsos",
			URL:  "https://www.ipsos.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    4318,
				Alias: "ipsos",
				Name:  "Ipsos",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "",
				GoRepositoryCount: 0,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Ipsos-EI_IE13063.11,16.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Ipsos-Reviews-E13063.htm",
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go: []string{
					"https://www.linkedin.com/jobs/view/4053349624/",
				},
				domain.Rust:    nil,
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
			},
			ShortDescription: "Market research company",
			Industries:       []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
				domain.Czechia,
			},
		},
		{
			ID:   0,  // system
			Type: "", // system
			Name: "Bedrock Streaming",
			URL:  "https://bedrockstreaming.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    64853943,
				Alias: "bedrock-streaming",
				Name:  "Bedrock Streaming",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "bedrockstreaming",
				GoRepositoryCount: 7,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Bedrock-Streaming-EI_IE5955934.11,28.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Bedrock-Streaming-Reviews-E5955934.htm",
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go: []string{
					"https://www.linkedin.com/jobs/view/4049210489/",
				},
				domain.Rust:    nil,
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
			},
			ShortDescription:          "Bedrock creates and operates full-scope streaming platforms",
			Industries:                []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{},
		},
		{
			ID:   0,  // system
			Type: "", // system
			Name: "Whalebone",
			URL:  "https://www.whalebone.io/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    10961729,
				Alias: "whaleboneio",
				Name:  "Whalebone",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "whalebone",
				GoRepositoryCount: 5,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Whalebone-EI_IE4889910.11,20.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Whalebone-Reviews-E4889910.htm",
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go: []string{
					"https://www.linkedin.com/jobs/view/4007694502/",
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
				domain.Czechia,
			},
		},
		{
			ID:   0,  // system
			Type: "", // system
			Name: "OneSignal",
			URL:  "https://onesignal.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    6424376,
				Alias: "onesignal",
				Name:  "OneSignal",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "OneSignal",
				GoRepositoryCount: 5, // Rust 32
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-OneSignal-EI_IE1952551.11,20.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/OneSignal-Reviews-E1952551.htm",
			},
			OttaProfileSlug:   "OneSignal",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go: []string{
					"https://app.welcometothejungle.com/jobs/wxzE8h5B",
				},
				domain.Rust: []string{
					"https://app.welcometothejungle.com/jobs/wxzE8h5B",
				},
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
			},
			ShortDescription:          "Customer messaging and engagement platform",
			Industries:                []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{},
		},
		{
			ID:   0,                         // system
			Type: domain.CompanyTypeStartup, // system
			Name: "Metabase",
			URL:  "https://www.metabase.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    6460313,
				Alias: "metabase",
				Name:  "Metabase",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "metabase",
				GoRepositoryCount: 0, // Clojure 30
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Metabase-CA-EI_IE6095700.11,22.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Metabase-CA-Reviews-E6095700.htm",
			},
			OttaProfileSlug:   "Metabase",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go:     nil,
				domain.Rust:   nil,
				domain.Zig:    nil,
				domain.Scala:  nil,
				domain.Elixir: nil,
				domain.Clojure: []string{
					"https://app.welcometothejungle.com/jobs/K21IaTJ2",
				},
				domain.Haskell: nil,
			},
			ShortDescription: "Open-source Business Intelligence tool",
			Industries:       []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
			},
		},
		{
			ID:   0,  // system
			Type: "", // system
			Name: "Anjuna Security",
			URL:  "https://www.anjuna.io/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    18438300,
				Alias: "anjuna-security",
				Name:  "Anjuna Security",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "anjuna-security",
				GoRepositoryCount: 2,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Anjuna-Security-EI_IE3150191.11,26.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Anjuna-Security-Reviews-E3150191.htm",
			},
			OttaProfileSlug:   "Anjuna",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go: []string{
					"https://www.linkedin.com/jobs/view/4052028718/",
					"https://app.welcometothejungle.com/jobs/yxGdxMhb",
				},
				domain.Rust: []string{
					"https://www.linkedin.com/jobs/view/4052028718/",
					"https://app.welcometothejungle.com/jobs/yxGdxMhb",
				},
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
			},
			ShortDescription: "Public-cloud network security",
			Industries: []domain.Industry{
				domain.IndustryCyberSecurity,
			},
			HasEmployeesFromCountries: []domain.Country{},
		},
		{
			ID:   0,  // system
			Type: "", // system
			Name: "Aqua Security",
			URL:  "https://www.aquasec.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    10034420,
				Alias: "aquasecteam",
				Name:  "Aqua Security",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "aquasecurity",
				GoRepositoryCount: 101,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Aqua-Security-Software-EI_IE1785939.11,33.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Aqua-Security-Software-Reviews-E1785939.htm",
			},
			OttaProfileSlug:   "Aqua-Security",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go: []string{
					"https://www.linkedin.com/jobs/view/4045761623/",
				},
				domain.Rust:    nil,
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
			},
			ShortDescription: "Cloud native security platform",
			Industries: []domain.Industry{
				domain.IndustryCyberSecurity,
			},
			HasEmployeesFromCountries: []domain.Country{},
		},
		{
			ID:   0,  // system
			Type: "", // system
			Name: "Verve",
			URL:  "https://verve.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    108605,
				Alias: "verve-ad-solutions",
				Name:  "Verve",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "",
				GoRepositoryCount: 0,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Verve-EI_IE4432646.11,16.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Verve-Reviews-E4432646.htm",
			},
			OttaProfileSlug:   "Verve",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go: []string{
					"https://www.linkedin.com/jobs/view/4015809270/",
				},
				domain.Rust:    nil,
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
			},
			ShortDescription: "Real time omnichannel ad platform",
			Industries: []domain.Industry{
				domain.IndustryAdTech,
			},
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
				domain.Czechia,
			},
		},
		{
			ID:   0,  // system
			Type: "", // system
			Name: "Wargaming",
			URL:  "https://wargaming.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    127309,
				Alias: "wargaming-net",
				Name:  "Wargaming",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "wgnet",
				GoRepositoryCount: 3,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Wargaming-EI_IE381713.11,20.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Wargaming-Reviews-E381713.htm",
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go: []string{
					"https://www.linkedin.com/jobs/view/4048943270/",
				},
				domain.Rust:    nil,
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
			},
			ShortDescription: "Game developer and publisher",
			Industries: []domain.Industry{
				domain.IndustryGameDev,
			},
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
				domain.Czechia,
			},
		},

		// Template
		//{
		//	ID:   0,  // system
		//	Type: "", // system
		//	Name: "",
		//	URL:  "",
		//	LinkedInProfile: domain.LinkedInProfile{
		//		ID:    0,
		//		Alias: "",
		//		Name:  "",
		//	},
		//	GitHubProfile: domain.GitHubProfile{
		//		Login:             "",
		//		GoRepositoryCount: 0,
		//	},
		//	GlassdoorProfile: domain.GlassdoorProfile{
		//		OverviewURL: "",
		//		ReviewsURL:  "",
		//	},
		//	OttaProfileSlug:   "",
		//	YouTubeChannelURL: "",
		//	GoMainLanguage:    false,
		//	Vacancies: domain.Vacancies{
		//		domain.Go:      []string{},
		//		domain.Rust:    nil,
		//		domain.Zig:     nil,
		//		domain.Scala:   nil,
		//		domain.Elixir:  nil,
		//		domain.Clojure: nil,
		//		domain.Haskell: nil,
		//	},
		//	ShortDescription:          "",
		//	Industries:                []domain.Industry{},
		//	HasEmployeesFromCountries: []domain.Country{},
		//},
	}
}
