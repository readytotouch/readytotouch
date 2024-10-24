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

		//
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
			ShortDescription: "",
			Industries:       []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
			},
		},

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
