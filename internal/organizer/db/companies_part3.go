package db

import "github.com/readytotouch/readytotouch/internal/domain"

func companiesPart3() []domain.CompanyProfile {
	return []domain.CompanyProfile{
		// Pro5.ai  - AI-powered platform
		{
			ID:   0,  // system
			Type: "", // system
			Name: "Pro5.ai",
			URL:  "https://www.pro5.ai/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:       71633800,
				Alias:    "pro5-ai",
				Name:     "Pro5.ai",
				Verified: true,
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
					"https://www.linkedin.com/jobs/view/4057458776/",
				},
				domain.Rust:    nil,
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
			},
			ShortDescription:          "The leading AI-powered platform",
			Industries:                []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{},
		},

		// web application security and performance / Cloudflare
		{
			ID:   0,  // system
			Type: "", // system
			Name: "Cloudflare",
			URL:  "https://www.cloudflare.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:       407222,
				Alias:    "cloudflare",
				Name:     "Cloudflare",
				Verified: true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "Cloudflare",
				GoRepositoryCount: 58,
				Verified:          false,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "",
				ReviewsURL:  "",
				Verified:    false,
			},
			OttaProfileSlug:   "Cloudflare-1",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go: nil,
				domain.Rust: []string{
					"https://www.linkedin.com/jobs/view/4041774900/",
				},
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
			},
			ShortDescription: "Web application security and performance.",
			Industries:       []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
				domain.Czechia,
			},
		},

		// Smart Electric Vehicle charging stations
		{
			ID:   0,  // system
			Type: "", // system
			Name: "EverCharge ",
			URL:  "https://evercharge.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:       4817778,
				Alias:    "evercharge",
				Name:     "EverCharge ",
				Verified: true,
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
			ShortDescription: "Smart Electric Vehicle charging stations",
			Industries:       []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
			},
		},

		// VPN company - ExpressVPN
		{
			ID:   0,  // system
			Type: "", // system
			Name: "ExpressVPN ",
			URL:  "https://www.expressvpn.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:       14488856,
				Alias:    "expressvpn",
				Name:     "ExpressVPN",
				Verified: false,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "expressvpn",
				GoRepositoryCount: 3,
				Verified:          true,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "",
				ReviewsURL:  "",
				Verified:    false,
			},
			OttaProfileSlug:   "ExpressVPN",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go: []string{
					"https://www.linkedin.com/jobs/view/4058715430/,",
				},
				domain.Rust:    nil,
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
			},
			ShortDescription:          "VPN company",
			Industries:                []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{},
		},

		//GlobalLogic -  full-lifecycle product development services
		{
			ID:   0,  // system
			Type: "", // system
			Name: "GlobalLogic",
			URL:  "https://www.globallogic.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:       164008,
				Alias:    "globallogic",
				Name:     "GlobalLogic",
				Verified: true,
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
					"https://www.linkedin.com/jobs/view/4057277625/",
				},
				domain.Rust:    nil,
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
			},
			ShortDescription: "Full-lifecycle product development services",
			Industries:       []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{
				domain.Czechia,
				domain.Ukraine,
			},
		},

		// Advertising - Influ2
		{
			ID:   0,  // system
			Type: "", // system
			Name: "Influ2",
			URL:  "https://www.influ2.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:       0,
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
				OverviewURL: "",
				ReviewsURL:  "",
				Verified:    false,
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
			ShortDescription: "Person-based advertising engine",
			Industries:       []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
			},
		},

		// DoublU -  AI enabled process automation.
		{
			ID:   0,  // system
			Type: "", // system
			Name: "DoublU",
			URL:  "https://doublu.ai/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:       18898324,
				Alias:    "",
				Name:     "DoublU",
				Verified: false,
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
					"https://www.linkedin.com/jobs/view/4058240622/",
				},
				domain.Rust:    nil,
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
			},
			ShortDescription:          "AI enabled process automation.",
			Industries:                []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{},
		},

		// telecommunications
		{
			ID:   0,  // system
			Type: "", // system
			Name: "PLC Group",
			URL:  "https://www.plcgroup.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:       11759368,
				Alias:    "plcgroupinc",
				Name:     "PLC Group",
				Verified: false,
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
					"https://www.linkedin.com/jobs/view/4057742583/",
				},
				domain.Rust:    nil,
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
			},
			ShortDescription:          "Critical Facilities & Infrastructure Management solutions",
			Industries:                []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{},
		},

		// Capital One - banking
		{
			ID:   0,  // system
			Type: "", // system
			Name: "Capital One",
			URL:  "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:       1419,
				Alias:    "capital-one",
				Name:     "Capital One",
				Verified: true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "capitalone",
				GoRepositoryCount: 5,
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
					"https://www.linkedin.com/jobs/view/4058793283/",
				},
				domain.Rust:    nil,
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
			},
			ShortDescription: "American bank holding company",
			Industries:       []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{
				domain.Czechia,
				domain.Ukraine,
			},
		},

		//Dynatrace - analytics and automation
		{
			ID:   0,  // system
			Type: "", // system
			Name: "Dynatrace",
			URL:  "https://www.dynatrace.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:       125999,
				Alias:    "dynatrace",
				Name:     "Dynatrace",
				Verified: true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "dynatrace",
				GoRepositoryCount: 11,
				Verified:          true,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "",
				ReviewsURL:  "",
				Verified:    false,
			},
			OttaProfileSlug:   "Dynatrace",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go: []string{
					"https://www.linkedin.com/jobs/view/4057742583/",
				},
				domain.Rust:    nil,
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
			},
			ShortDescription: "Platform that delivers analytics and automation for unified observability and security",
			Industries:       []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{
				domain.Czechia,
				domain.Ukraine,
			},
		},

		// Tele2 - telecommunications
		{
			ID:   0,  // system
			Type: "", // system
			Name: "Tele2",
			URL:  "https://www.tele2.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:       2831,
				Alias:    "tele2",
				Name:     "Tele2",
				Verified: true,
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
					"https://www.linkedin.com/jobs/view/4058482144/",
				},
				domain.Rust:    nil,
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
			},
			ShortDescription: "Telecommunication services",
			Industries:       []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{
				domain.Czechia,
				domain.Ukraine,
			},
		},

		// emnify - telecommunications
		{
			ID:   0,  // system
			Type: "", // system
			Name: "emnify",
			URL:  "https://www.emnify.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:       0,
				Alias:    "emnify",
				Name:     "emnify",
				Verified: true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "emnify",
				GoRepositoryCount: 7,
				Verified:          true,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "",
				ReviewsURL:  "",
				Verified:    false,
			},
			OttaProfileSlug:   "emnify-1",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go: []string{
					"https://www.linkedin.com/jobs/view/4058102205/",
				},
				domain.Rust:    nil,
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
			},
			ShortDescription: "IoT connectivity",
			Industries:       []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
			},
		},

		// PayNearMe - payments technology
		{
			ID:   0,  // system
			Type: "", // system
			Name: "PayNearMe",
			URL:  "https://home.paynearme.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:       1137260,
				Alias:    "paynearme",
				Name:     "PayNearMe",
				Verified: true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "paynearme",
				GoRepositoryCount: 13,
				Verified:          false,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "",
				ReviewsURL:  "",
				Verified:    false,
			},
			OttaProfileSlug:   "PayNearMe",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go: []string{
					"https://www.linkedin.com/jobs/view/4058460891",
				},
				domain.Rust:    nil,
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
			},
			ShortDescription:          "Payments DevOps",
			Industries:                []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{},
		},

		// TikTok
		{
			ID:   0,  // system
			Type: "", // system
			Name: "ByteDance",
			URL:  "https://jobs.bytedance.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:       6575553,
				Alias:    "bytedance",
				Name:     "ByteDance",
				Verified: true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "bytedance",
				GoRepositoryCount: 33,
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
					"https://www.linkedin.com/jobs/view/4057761195/",
				},
				domain.Rust: []string{
					"https://www.linkedin.com/jobs/view/4059885479/",
				},
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
			},
			ShortDescription: "Chinese internet technology company",
			Industries:       []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{
				domain.Czechia,
				domain.Ukraine,
			},
		},

		// TikTok - short-form video
		{
			ID:   0,  // system
			Type: "", // system
			Name: "TikTok",
			URL:  "https://careers.tiktok.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:       33246798,
				Alias:    "tiktok",
				Name:     "TikTok",
				Verified: true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "tiktok",
				GoRepositoryCount: 15,
				Verified:          true,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "",
				ReviewsURL:  "",
				Verified:    false,
			},
			OttaProfileSlug:   "TikTok",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go: []string{
					"https://www.linkedin.com/jobs/view/4057303610/",
				},
				domain.Rust:    nil,
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
			},
			ShortDescription: "Social media platform",
			Industries:       []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{
				domain.Czechia,
				domain.Ukraine,
			},
		},

		// Moody's -  American business and financial services company
		{
			ID:   0,  // system
			Type: "", // system
			Name: "Moody's Corporation",
			URL:  "https://www.moodys.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:       165033,
				Alias:    "moodys-corporation",
				Name:     "Moody's Corporation",
				Verified: true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "moodysanalytics",
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
					"https://www.linkedin.com/jobs/view/4060347505/",
				},
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
			},
			ShortDescription: "American business and financial services company",
			Industries:       []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{
				domain.Czechia,
				domain.Ukraine,
			},
		},

		//Lytx - fleet and compliance management solutions
		{
			ID:   0,  // system
			Type: "", // system
			Name: "Lytx, Inc.",
			URL:  "https://www.lytx.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:       26167,
				Alias:    "lytxinc",
				Name:     "Lytx, Inc.",
				Verified: true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "lytx",
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
					"https://www.linkedin.com/jobs/view/3951803632/",
				},
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
			},
			ShortDescription: "Fleet and compliance management solutions",
			Industries:       []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{
				domain.Czechia,
				domain.Ukraine,
			},
		},

		// Agoda - booking platform
		{
			ID:   0,  // system
			Type: "", // system
			Name: "Agoda",
			URL:  "https://careersatagoda.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:       66719,
				Alias:    "agoda",
				Name:     "Agoda",
				Verified: true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "agoda-com",
				GoRepositoryCount: 10,
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
					"https://www.linkedin.com/jobs/view/4059885479/",
				},
				domain.Rust: []string{
					"https://www.linkedin.com/jobs/view/4059885479/",
				},
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
			},
			ShortDescription: "Online travel-booking platforms",
			Industries:       []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{
				domain.Czechia,
				domain.Ukraine,
			},
		},

		//
		{
			ID:   0,  // system
			Type: "", // system
			Name: "",
			URL:  "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:       0,
				Alias:    "",
				Name:     "",
				Verified: false,
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
				domain.Go:      nil,
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

		//
		{
			ID:   0,  // system
			Type: "", // system
			Name: "",
			URL:  "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:       0,
				Alias:    "",
				Name:     "",
				Verified: false,
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
				domain.Go:      nil,
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

		//
		{
			ID:   0,  // system
			Type: "", // system
			Name: "",
			URL:  "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:       0,
				Alias:    "",
				Name:     "",
				Verified: false,
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
				domain.Go:      nil,
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

		//
		{
			ID:   0,  // system
			Type: "", // system
			Name: "",
			URL:  "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:       0,
				Alias:    "",
				Name:     "",
				Verified: false,
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
				domain.Go:      nil,
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

		//
		{
			ID:   0,  // system
			Type: "", // system
			Name: "",
			URL:  "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:       0,
				Alias:    "",
				Name:     "",
				Verified: false,
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
				domain.Go:      nil,
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

		//
		{
			ID:   0,  // system
			Type: "", // system
			Name: "",
			URL:  "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:       0,
				Alias:    "",
				Name:     "",
				Verified: false,
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
				domain.Go:      nil,
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

		//
		{
			ID:   0,  // system
			Type: "", // system
			Name: "",
			URL:  "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:       0,
				Alias:    "",
				Name:     "",
				Verified: false,
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
				domain.Go:      nil,
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

		//
		{
			ID:   0,  // system
			Type: "", // system
			Name: "",
			URL:  "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:       0,
				Alias:    "",
				Name:     "",
				Verified: false,
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
				domain.Go:      nil,
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

		//
		{
			ID:   0,  // system
			Type: "", // system
			Name: "",
			URL:  "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:       0,
				Alias:    "",
				Name:     "",
				Verified: false,
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
				domain.Go:      nil,
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
	}
}
