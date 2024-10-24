package db

import "github.com/readytotouch/readytotouch/internal/domain"

func companiesPart3() []domain.CompanyProfile {
	return []domain.CompanyProfile{
		//  monitoring platform for cloud applications / Datadog
		{
			ID:   0,  // system
			Type: "", // system
			Name: "Datadog ",
			URL:  "https://www.datadoghq.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:       1066442,
				Alias:    "datadog",
				Name:     "Datadog",
				Verified: true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "datadog",
				GoRepositoryCount: 85,
				Verified:          true,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.es/Resumen/Trabajar-en-Datadog-EI_IE762009.12,19.htm",
				ReviewsURL:  "",
				Verified:    true,
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go: nil,
				domain.Rust: []string{
					"https://www.linkedin.com/jobs/view/3967138677/",
				},
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
			},
			ShortDescription: "observability platform that supports every phase of software development",
			Industries:       []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
				domain.Czechia,
			},
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
