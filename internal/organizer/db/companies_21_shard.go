package db

import "github.com/readytotouch/readytotouch/internal/domain"

func companies21Shard() []domain.CompanyProfile {
	return []domain.CompanyProfile{
		{
			ID:         0,  // system
			Type:       "", // system
			Name:       "CUJU",
			BaseURL:    "https://cuju.pro/",
			CareersURL: "",
			AboutURL:   "",
			BlogURL:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                93166111,
				IDs:               nil,
				Alias:             "cuju-app",
				Name:              "CUJU",
				Followers:         "981",
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
					GitHubRepositoryCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Senior Go Engineer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4294445930/",
							Location:             "Germany",
							Date:                 mustDate("2025-09-06"),
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
			ShortDescription: "An AI based technology which gives football talents the chance to improve their skills, compete against friends and the global CUJU community and get scouted by a record-proven network of international senior scouts and specialists",
			Industries:       []domain.Industry{
				// NOP
			},
		},
		{
			ID:         0,  // system
			Type:       "", // system
			Name:       "AutoStore™",
			BaseURL:    "https://www.autostoresystem.com/",
			CareersURL: "https://www.autostoresystem.com/careers",
			AboutURL:   "https://www.autostoresystem.com/company",
			BlogURL:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                22346788,
				IDs:               nil,
				Alias:             "autostoresystem",
				Name:              "AutoStore™",
				Followers:         "66K",
				Employees:         "1K-5K",
				AssociatedMembers: "856",
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
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Rust: {
					GitHubRepositoryCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Back-End Software Engineer – C++, Rust or C#",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4281833526/",
							Location:             "Vindafjord, Rogaland, Norway",
							Date:                 mustDate("2025-10-03", "2025-09-06"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Back-end Software Engineer – C++, Rust or C#",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4281836212/",
							Location:             "Stavanger, Rogaland, Norway",
							Date:                 mustDate("2025-11-14", "2025-10-03"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Back-end Software Engineer – C++, Rust or C#",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4281830966/",
							Location:             "Oslo, Oslo, Norway",
							Date:                 mustDate("2025-11-14", "2025-10-03"),
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
			ShortDescription: "",
			Industries:       []domain.Industry{
				// NOP
			},
		},
		{
			ID:         0,  // system
			Type:       "", // system
			Name:       "Supabase",
			BaseURL:    "https://supabase.com/",
			CareersURL: "https://supabase.com/careers",
			AboutURL:   "https://supabase.com/company",
			BlogURL:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                31546644,
				IDs:               nil,
				Alias:             "supabase",
				Name:              "Supabase",
				Followers:         "60K",
				Employees:         "51-200",
				AssociatedMembers: "178",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:     "supabase",
				Followers: "7.2k",
				Verified:  true,
			},
			BlindProfile: domain.BlindProfile{
				Alias: "",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "supabase",
				Employees: "126",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Supabase-EI_IE7639911.11,19.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Supabase-Reviews-E7639911.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Supabase-Jobs-E7639911.htm",
				Jobs:        "33",
				Reviews:     "5",
				Salaries:    "3",
				ReviewsRate: "4.8",
				Verified:    false,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "",
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoryCount: 16,
					Vacancies: []domain.Vacancy{
						{
							Title:                "CLI Engineer",
							SubTitle:             "",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							Location:             "",
							URL:                  "https://jobs.ashbyhq.com/supabase/b7b2f679-c6d8-4f06-8c9e-05a6667bccfa",
							Date:                 mustDate("2025-09-05"),
							WithSalary:           false,
							Remote:               true,
						},
					},
				},
				domain.Rust: {
					GitHubRepositoryCount: 11,
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
			ShortDescription: "Supabase is the Postgres development platform",
			Industries: []domain.Industry{
				domain.IndustryDevOps,
			},
		},
		{
			ID:         0,  // system
			Type:       "", // system
			Name:       "Gresham",
			BaseURL:    "https://www.greshamtech.com/",
			CareersURL: "https://www.greshamtech.com/careers",
			AboutURL:   "https://www.greshamtech.com/about/our-vision",
			BlogURL:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                22858,
				IDs:               []int{22858, 2548739},
				Alias:             "greshamtech",
				Name:              "Gresham",
				Followers:         "13K",
				Employees:         "201-500",
				AssociatedMembers: "356",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:     "gresham-computing",
				Followers: "14",
				Verified:  false,
			},
			BlindProfile: domain.BlindProfile{
				Alias: "",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias: "",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Gresham-Tech-EI_IE26058.11,23.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Gresham-Tech-Reviews-E26058.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Gresham-Tech-Jobs-E26058.htm",
				Jobs:        "",
				Reviews:     "62",
				Salaries:    "149",
				ReviewsRate: "3.4",
				Verified:    false,
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
					GitHubRepositoryCount: 3,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Software Engineer Clojure",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4294972674/",
							Location:             "City Of Bristol, England, United Kingdom",
							Date:                 mustDate("2025-09-03"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Clojure Backend Developer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4343575058/",
							Location:             "Bangalore Urban, Karnataka, India",
							Date:                 mustDate("2025-12-15"),
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
			ShortDescription: "Enterprise data automation solutions",
			Industries:       []domain.Industry{
				// NOP
			},
		},
		{
			ID:         0,  // system
			Type:       "", // system
			Name:       "Multiverse",
			BaseURL:    "https://www.multiverse.io/",
			CareersURL: "https://www.multiverse.io/careers",
			AboutURL:   "https://www.multiverse.io/about",
			BlogURL:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                10579146,
				IDs:               nil,
				Alias:             "joinmultiverse",
				Name:              "Multiverse",
				Followers:         "91K",
				Employees:         "501-1K",
				AssociatedMembers: "1,273",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:     "Multiverse-io",
				Followers: "24",
				Verified:  true,
			},
			BlindProfile: domain.BlindProfile{
				Alias: "",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias: "",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Multiverse-EI_IE1676004.11,21.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Multiverse-Reviews-E1676004.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Multiverse-Jobs-E1676004.htm",
				Jobs:        "34",
				Reviews:     "351",
				Salaries:    "1.1K",
				ReviewsRate: "3.1",
				Verified:    false,
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
					GitHubRepositoryCount: 19,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Senior Software Engineer",
							ShortDescription:     "Our stack includes Python, TypeScript, React, Elixir, AWS",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4285039642/",
							Location:             "London, England, United Kingdom",
							CloudProviders:       []domain.CloudProvider{domain.AWS},
							Date:                 mustDate("2025-09-10"),
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
			ShortDescription: "The upskilling platform for AI and tech adoption",
			Industries: []domain.Industry{
				domain.IndustryEdTech,
			},
		},
		{
			ID:         0,  // system
			Type:       "", // system
			Name:       "Blue River Technology",
			BaseURL:    "https://www.bluerivertechnology.com/",
			CareersURL: "",
			AboutURL:   "",
			BlogURL:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                2742193,
				IDs:               nil,
				Alias:             "bluerivertech",
				Name:              "Blue River Technology",
				Followers:         "23K",
				Employees:         "201-500",
				AssociatedMembers: "334",
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
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Blue-River-Technology-EI_IE1175124.11,32.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Blue-River-Technology-Reviews-E1175124.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Blue-River-Technology-Jobs-E1175124.htm",
				Jobs:        "7",
				Reviews:     "80",
				Salaries:    "160",
				ReviewsRate: "3.6",
				Verified:    false,
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
							Title:                "Embedded Software Engineer (C++, Rust)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4275796783/",
							Location:             "United States",
							Date:                 mustDate("2025-09-07"),
							WithSalary:           true, // $125k/yr - $218k/yr
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
			ShortDescription: "We empower our customers – farmers",
			Industries: []domain.Industry{
				domain.IndustryAgroTech,
			},
		},
		{
			ID:         0,  // system
			Type:       "", // system
			Name:       "Ultra Intelligence & Communications",
			BaseURL:    "https://www.ultra-ic.com/",
			CareersURL: "https://www.ultra-ic.com/careers/",
			AboutURL:   "https://www.ultra-ic.com/about-us/",
			BlogURL:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                69235707,
				IDs:               nil,
				Alias:             "ultra-ic",
				Name:              "Ultra Intelligence & Communications",
				Followers:         "14K",
				Employees:         "5K-10K",
				AssociatedMembers: "452",
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
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Rust: {
					GitHubRepositoryCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Embedded Rust Software Engineer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4283637769/",
							Location:             "Tampa, FL",
							Date:                 mustDate("2025-09-07"),
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
			ShortDescription: "Defense and Space Manufacturing",
			Industries: []domain.Industry{
				domain.IndustryDefenseTech,
			},
		},
		{
			ID:         0,  // system
			Type:       "", // system
			Name:       "Cognyte",
			BaseURL:    "https://www.cognyte.com/",
			CareersURL: "https://www.cognyte.com/careers/",
			AboutURL:   "https://www.cognyte.com/about/",
			BlogURL:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                3669,
				IDs:               nil,
				Alias:             "cognyte",
				Name:              "Cognyte",
				Followers:         "118K",
				Employees:         "1K-5K",
				AssociatedMembers: "1,506",
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
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Cognyte-EI_IE4430257.11,18.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Cognyte-Reviews-E4430257.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Cognyte-Jobs-E4430257.htm",
				Jobs:        "23",
				Reviews:     "318",
				Salaries:    "403",
				ReviewsRate: "4.1",
				Verified:    true,
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
							Title:                "Rust Software Engineer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4298710314/",
							Location:             "Limassol, Cyprus",
							Date:                 mustDate("2025-10-14", "2025-09-22", "2025-09-10"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Rust Software Engineer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4359463590/",
							Location:             "Limassol, Cyprus",
							Date:                 mustDate("2026-01-10"),
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
			ShortDescription: "Data processing and investigative analytics solutions",
			Industries:       []domain.Industry{
				// NOP
			},
		},
		{
			ID:         0,  // system
			Type:       "", // system
			Name:       "Catio",
			BaseURL:    "https://www.catio.tech/",
			CareersURL: "",
			AboutURL:   "https://www.catio.tech/about-us",
			BlogURL:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                92939924,
				IDs:               nil,
				Alias:             "catiotech",
				Name:              "Catio",
				Followers:         "1K",
				Employees:         "11-50",
				AssociatedMembers: "21",
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
					GitHubRepositoryCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Senior Backend Engineer, Go",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4297890217/",
							Location:             "United States",
							Date:                 mustDate("2025-09-10"),
							WithSalary:           true, // $160k/yr - $200k/yr
							Remote:               true,
						},
						{
							Title:                "Senior Backend Engineer, Go",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4300776433/",
							Location:             "United States",
							Date:                 mustDate("2025-09-16"),
							WithSalary:           true, // $160k/yr - $200k/yr
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
			ShortDescription: "Copilot for Tech Architecture – equipping companies with a data-driven and AI-powered platform to excel with evaluating, planning and evolving their tech stacks",
			Industries:       []domain.Industry{
				// NOP
			},
		},
		{
			ID:         0,  // system
			Type:       "", // system
			Name:       "CloudNuro",
			BaseURL:    "https://www.cloudnuro.ai/",
			CareersURL: "https://www.cloudnuro.ai/careers",
			AboutURL:   "https://www.cloudnuro.ai/about",
			BlogURL:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                79733634,
				IDs:               nil,
				Alias:             "cloudnuro",
				Name:              "CloudNuro",
				Followers:         "12K",
				Employees:         "11-50",
				AssociatedMembers: "39",
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
					GitHubRepositoryCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Backend Engineer (Go + Python)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4294797257/",
							Location:             "Bengaluru, Karnataka, India",
							Date:                 mustDate("2025-09-10"),
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
			ShortDescription: "Intelligent SaaS Management Platform",
			Industries:       []domain.Industry{
				// NOP
			},
		},
		{
			ID:         0,  // system
			Type:       "", // system
			Name:       "Bitexen",
			BaseURL:    "https://www.bitexen.com/",
			CareersURL: "",
			AboutURL:   "",
			BlogURL:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                28666238,
				IDs:               nil,
				Alias:             "bitexen",
				Name:              "Bitexen",
				Followers:         "39K",
				Employees:         "51-200",
				AssociatedMembers: "107",
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
					GitHubRepositoryCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Senior Backend Engineer (C# / Go)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4298360650/",
							Location:             "Istanbul, Türkiye",
							Date:                 mustDate("2025-09-10"),
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
				// Cryptocurrency
			},
			Ignore: true, // Cryptocurrency
		},
		{
			ID:         0,  // system
			Type:       "", // system
			Name:       "Hostaway",
			BaseURL:    "https://www.hostaway.com/",
			CareersURL: "https://careers.hostaway.com/",
			AboutURL:   "https://www.hostaway.com/about/",
			BlogURL:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                10388777,
				IDs:               nil,
				Alias:             "hostaway",
				Name:              "Hostaway",
				Followers:         "27K",
				Employees:         "201-500",
				AssociatedMembers: "291",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:     "Hostaway",
				Followers: "16",
				Verified:  false,
			},
			BlindProfile: domain.BlindProfile{
				Alias: "",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias: "",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Hostaway-EI_IE1358830.11,19.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Hostaway-Reviews-E1358830.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Hostaway-Jobs-E1358830.htm",
				Jobs:        "171",
				Reviews:     "72",
				Salaries:    "64",
				ReviewsRate: "4.2",
				Verified:    true,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "",
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoryCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Senior Backend Engineer, PHP/Go",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4298774762/",
							Location:             "Tottenham, England, United Kingdom",
							Date:                 mustDate("2025-10-03", "2025-09-10"),
							WithSalary:           false,
							Remote:               true,
						},
						{
							Title:                "Staff Backend Engineer, PHP + Go",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4336730403/",
							Location:             "European Union",
							Date:                 mustDate("2025-12-09", "2025-11-18"),
							WithSalary:           false,
							Remote:               true,
						},
						{
							Title:                "Senior Backend Engineer, PHP/Go",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4339377138/",
							Location:             "Tottenham, England, United Kingdom",
							Date:                 mustDate("2026-01-10", "2025-12-20", "2025-11-28"),
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
			ShortDescription: "Vacation rental management software",
			Industries: []domain.Industry{
				domain.IndustryPropTech,
			},
		},
		{
			ID:         0,  // system
			Type:       "", // system
			Name:       "Cirrus Logic",
			BaseURL:    "https://www.cirrus.com/",
			CareersURL: "https://www.cirrus.com/careers",
			AboutURL:   "https://www.cirrus.com/company",
			BlogURL:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                5617,
				IDs:               nil,
				Alias:             "cirrus-logic",
				Name:              "Cirrus Logic",
				Followers:         "55K",
				Employees:         "1K-5K",
				AssociatedMembers: "1,747",
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
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Cirrus-Logic-EI_IE1259.11,23.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Cirrus-Logic-Reviews-E1259.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Cirrus-Logic-Jobs-E1259.htm",
				Jobs:        "76",
				Reviews:     "419",
				Salaries:    "1.3K",
				ReviewsRate: "4.1",
				Verified:    false,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "",
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoryCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Full Stack Software Engineer – Go/Python ",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4298640048/",
							Location:             "Austin, TX",
							Date:                 mustDate("2025-09-10"),
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
			ShortDescription: "Low-Power Audio and High-Performance Mixed-Signal Processing",
			Industries:       []domain.Industry{
				// NOP
			},
		},
		{
			ID:         0,  // system
			Type:       "", // system
			Name:       "Surfshark",
			BaseURL:    "https://surfshark.com/",
			CareersURL: "",
			AboutURL:   "",
			BlogURL:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                69271673,
				IDs:               nil,
				Alias:             "surfshark",
				Name:              "Surfshark",
				Followers:         "22K",
				Employees:         "201-500",
				AssociatedMembers: "452",
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
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Surfshark-EI_IE6375899.11,20.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Surfshark-Reviews-E6375899.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Surfshark-Jobs-E6375899.htm",
				Jobs:        "14",
				Reviews:     "48",
				Salaries:    "68",
				ReviewsRate: "4.2",
				Verified:    false,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "",
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoryCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Lead Backend Software Engineer (Go)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4298799776/",
							Location:             "Vilnius, Vilniaus, Lithuania",
							Date:                 mustDate("2025-09-10"),
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
			ShortDescription: "VPN",
			Industries: []domain.Industry{
				domain.IndustryCyberSecurity,
			},
		},
		{
			ID:         0,  // system
			Type:       "", // system
			Name:       "Forte Group",
			BaseURL:    "https://fortegrp.com/",
			CareersURL: "https://fortegrp.com/careers",
			AboutURL:   "https://fortegrp.com/about-forte-group",
			BlogURL:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                1069110,
				IDs:               nil,
				Alias:             "fortegroup",
				Name:              "Forte Group",
				Followers:         "81K",
				Employees:         "501-1K",
				AssociatedMembers: "665",
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
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Forte-Group-EI_IE402276.11,22.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Forte-Group-Reviews-E402276.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Forte-Group-Jobs-E402276.htm",
				Jobs:        "",
				Reviews:     "167",
				Salaries:    "163",
				ReviewsRate: "4.0",
				Verified:    true,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "",
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoryCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Senior C#/Golang Engineer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4296467085/",
							Location:             "Poland",
							Date:                 mustDate("2025-09-06"),
							WithSalary:           false,
							Remote:               true,
						},
						{
							Title:                "Middle/Senior Golang Engineer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4306491024/",
							Location:             "Poland",
							Date:                 mustDate("2025-09-27"),
							WithSalary:           false,
							Remote:               true,
						},
						{
							Title:                "Senior C#/Golang Engineer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4310854508/",
							Location:             "Poland",
							Date:                 mustDate("2025-10-07"),
							WithSalary:           false,
							Remote:               true,
						},
						{
							Title:                "Golang Developer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4317345671/",
							Location:             "Ukraine",
							Date:                 mustDate("2025-10-21"),
							WithSalary:           false,
							Remote:               true,
						},
						{
							Title:                "Middle/Senior Golang Engineer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4332420427/",
							Location:             "Poland",
							Date:                 mustDate("2025-10-29"),
							WithSalary:           false,
							Remote:               true,
						},
						{
							Title:                "Full-Stack Golang Engineer (React)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4336181513/",
							Location:             "Ukraine",
							Date:                 mustDate("2025-11-15"),
							WithSalary:           false,
							Remote:               true,
						},
						{
							Title:                "Middle/Senior Golang Engineer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4339275791/",
							Location:             "Poland",
							Date:                 mustDate("2025-11-28"),
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
			ShortDescription: "Managed Engineering Solutions in Software Development and Data Engineering",
			Industries:       []domain.Industry{
				// NOP
			},
		},
		{
			ID:         0,  // system
			Type:       "", // system
			Name:       "Thunes",
			BaseURL:    "https://www.thunes.com/",
			CareersURL: "",
			AboutURL:   "https://www.thunes.com/about-us/",
			BlogURL:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                14552981,
				IDs:               nil,
				Alias:             "thunespayments",
				Name:              "Thunes",
				Followers:         "71K",
				Employees:         "201-500",
				AssociatedMembers: "481",
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
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Thunes-EI_IE2906161.11,17.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Thunes-Reviews-E2906161.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Thunes-Jobs-E2906161.htm",
				Jobs:        "26",
				Reviews:     "81",
				Salaries:    "132",
				ReviewsRate: "3.3",
				Verified:    false,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "",
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoryCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Golang Software Engineer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4156638094/",
							Location:             "Barcelona, Catalonia, Spain",
							CloudProviders:       []domain.CloudProvider{},
							Date:                 mustDate("2026-01-16", "2025-12-25", "2025-12-02", "2025-11-11", "2025-10-21", "2025-09-09"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Golang Software Engineer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4229362550/",
							Location:             "Paris, Île-de-France, France",
							Date:                 mustDate("2025-10-21", "2025-09-30"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Go Senior Software Engineer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4354767832/",
							Location:             "Barcelona, Catalonia, Spain",
							Date:                 mustDate("2026-01-13"),
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
			Name:       "Kayrros",
			BaseURL:    "https://www.kayrros.com/",
			CareersURL: "https://www.kayrros.com/career/",
			AboutURL:   "https://www.kayrros.com/about-about/",
			BlogURL:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                10603168,
				IDs:               nil,
				Alias:             "kayrros",
				Name:              "Kayrros",
				Followers:         "43K",
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
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Kayrros-EI_IE2362521.11,18.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Kayrros-Reviews-E2362521.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Kayrros-Jobs-E2362521.htm",
				Jobs:        "2",
				Reviews:     "62",
				Salaries:    "125",
				ReviewsRate: "4.1",
				Verified:    false,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "",
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoryCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Golang Software Engineer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4297160208/",
							Location:             "Paris, Île-de-France, France",
							Date:                 mustDate("2025-09-09"),
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
			ShortDescription: "We use AI and geoanalytics to turn raw satellite data into actionable insights on energy, supply chains, physical risks, nature",
			Industries:       []domain.Industry{
				// NOP
			},
		},
		{
			ID:         0,  // system
			Type:       "", // system
			Name:       "Fingerprint",
			BaseURL:    "https://fingerprint.com/",
			CareersURL: "https://fingerprint.com/careers/",
			AboutURL:   "https://fingerprint.com/about-us/",
			BlogURL:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                31432400,
				IDs:               nil,
				Alias:             "fingerprintjs",
				Name:              "Fingerprint",
				Followers:         "28K",
				Employees:         "51-200",
				AssociatedMembers: "178",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:     "fingerprintjs",
				Followers: "462",
				Verified:  true,
			},
			BlindProfile: domain.BlindProfile{
				Alias: "",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias: "",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Fingerprint-EI_IE4075078.11,22.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Fingerprint-Reviews-E4075078.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Fingerprint-Jobs-E4075078.htm",
				Jobs:        "",
				Reviews:     "24",
				Salaries:    "29",
				ReviewsRate: "4.8",
				Verified:    false,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "",
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoryCount: 1,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Backend Software Engineer – Golang",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4298766644/",
							Location:             "United States",
							Date:                 mustDate("2025-09-10"),
							WithSalary:           true, // $150k/yr - $200k/yr
							Remote:               true,
						},
						{
							Title:                "Senior Backend Software Engineer (Golang)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4316006409/",
							Location:             "Chicago, IL",
							Date:                 mustDate("2025-10-17"),
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
			ShortDescription: "Identify visitors you can trust",
			Industries: []domain.Industry{
				domain.IndustryCyberSecurity,
			},
		},
		{
			ID:         0,  // system
			Type:       "", // system
			Name:       "Perion",
			BaseURL:    "https://perion.com/",
			CareersURL: "https://perion.com/careers/",
			AboutURL:   "https://perion.com/about/",
			BlogURL:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                26531,
				IDs:               []int{26531, 167701, 15179046, 27238464},
				Alias:             "perion",
				Name:              "Perion",
				Followers:         "19K",
				Employees:         "501-1K",
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
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Perion-Network-EI_IE561834.11,25.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Perion-Network-Reviews-E561834.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Perion-Network-Jobs-E561834.htm",
				Jobs:        "",
				Reviews:     "162",
				Salaries:    "144",
				ReviewsRate: "4.3",
				Verified:    false,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "",
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoryCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Back-End Developer – Golang",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4295267899/",
							Location:             "Tel Aviv-Yafo, Tel Aviv District, Israel",
							Date:                 mustDate("2025-09-08"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Back-End Developer – Golang",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4348076228/",
							Location:             "Tel Aviv-Yafo, Tel Aviv District, Israel",
							Date:                 mustDate("2025-12-02"),
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
			ShortDescription: "Perion is helping agencies, brands and retailers get better results with their marketing investments by providing advanced technology across all major digital channels",
			Industries: []domain.Industry{
				domain.IndustryAdTech,
			},
		},
		{
			ID:         0,  // system
			Type:       "", // system
			Name:       "Upfluence",
			BaseURL:    "https://www.upfluence.com/",
			CareersURL: "https://www.upfluence.com/careers",
			AboutURL:   "https://www.upfluence.com/about",
			BlogURL:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                2956735,
				IDs:               nil,
				Alias:             "upfluence",
				Name:              "Upfluence",
				Followers:         "22K",
				Employees:         "51-200",
				AssociatedMembers: "121",
				Verified:          false,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:     "upfluence",
				Followers: "6",
				Verified:  false,
			},
			BlindProfile: domain.BlindProfile{
				Alias:       "Upfluence",
				Employees:   "",
				Salary:      "",
				Reviews:     "",
				ReviewsRate: "",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "upfluence",
				Employees: "150",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Upfluence-EI_IE803160.11,20.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Upfluence-Reviews-E803160.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Upfluence-Jobs-E803160.htm",
				Jobs:        "5",
				Reviews:     "50",
				Salaries:    "127",
				ReviewsRate: "3.1",
				Verified:    false,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "",
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoryCount: 30,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Golang Developer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4298255189/",
							Location:             "Lyon, Auvergne-Rhône-Alpes, France",
							Date:                 mustDate("2025-09-10"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Golang Developer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4302201156/",
							Location:             "Lyon, Auvergne-Rhône-Alpes, France",
							Date:                 mustDate("2025-09-18"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Golang Developer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4306630638/",
							Location:             "Lyon, Auvergne-Rhône-Alpes, France",
							Date:                 mustDate("2025-09-30"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Golang Developer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4315332594/",
							Location:             "Lyon, Auvergne-Rhône-Alpes, France",
							Date:                 mustDate("2025-10-15"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Golang Developer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4316613327/",
							Location:             "Lyon, Auvergne-Rhône-Alpes, France",
							Date:                 mustDate("2025-10-21"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Back-End Software Engineer (Golang)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4335675931/",
							Location:             "Lyon, Auvergne-Rhône-Alpes, France",
							Date:                 mustDate("2025-11-12"),
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
			ShortDescription: "Influencer marketing platform",
			Industries:       []domain.Industry{
				// NOP
			},
		},
		{
			ID:         0,  // system
			Type:       "", // system
			Name:       "Capco",
			BaseURL:    "https://www.capco.com/",
			CareersURL: "https://www.capco.com/Careers/Job-Search",
			AboutURL:   "https://www.capco.com/About-Us",
			BlogURL:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                5470,
				IDs:               nil,
				Alias:             "capco",
				Name:              "Capco",
				Followers:         "456K",
				Employees:         "5K-10K",
				AssociatedMembers: "7,701",
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
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Capco-EI_IE400565.11,16.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Capco-Reviews-E400565.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Capco-Jobs-E400565.htm",
				Jobs:        "572",
				Reviews:     "3.9K",
				Salaries:    "7.5K",
				ReviewsRate: "3.8",
				Verified:    false,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "",
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoryCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Golang developer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4290908334/",
							Location:             "Brno Metropolitan Area",
							Date:                 mustDate("2025-09-16"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Golang Developer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4305601440/",
							Location:             "Pune, Maharashtra, India",
							Date:                 mustDate("2025-10-17"),
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
					Vacancies: []domain.Vacancy{
						{
							Title:                "Middle/Senior Big Data Engineer (Scala)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4277567694/",
							Location:             "Poland",
							Date:                 mustDate("2025-09-13"),
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
					Vacancies: []domain.Vacancy{
						{
							Title:                "Haskell Developer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4279000425/",
							Location:             "Poland",
							Date:                 mustDate("2025-09-14"),
							WithSalary:           false,
							Remote:               true,
						},
					},
				},
				domain.Erlang: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
			},
			ShortDescription: "Driving digital transformation in the financial services industry",
			Industries: []domain.Industry{
				domain.IndustryFinTech,
			},
		},
		{
			ID:         0,  // system
			Type:       "", // system
			Name:       "FacilitaPay",
			BaseURL:    "https://www.facilitapay.com/",
			CareersURL: "",
			AboutURL:   "https://www.facilitapay.com/about-us",
			BlogURL:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                27212402,
				IDs:               nil,
				Alias:             "facilitapay",
				Name:              "FacilitaPay",
				Followers:         "11K",
				Employees:         "51-200",
				AssociatedMembers: "74",
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
							Title:                "Software Engineer Specialist – Elixir | Phoenix",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4298961123/",
							Location:             "Brazil",
							Date:                 mustDate("2025-09-10"),
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
			ShortDescription: "FacilitaPay transforms payments between developed and emerging markets",
			Industries: []domain.Industry{
				domain.IndustryFinTech,
			},
		},
		{
			ID:         0,  // system
			Type:       "", // system
			Name:       "McKesson",
			BaseURL:    "https://www.mckesson.com/",
			CareersURL: "https://careers.mckesson.com/",
			AboutURL:   "https://www.mckesson.com/about-us/",
			BlogURL:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                1900,
				IDs:               []int{1900, 3597039, 69724639},
				Alias:             "mckesson",
				Name:              "McKesson",
				Followers:         "545K",
				Employees:         "10K+",
				AssociatedMembers: "23,331",
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
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-McKesson-EI_IE434.11,19.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/McKesson-Reviews-E434.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/McKesson-Jobs-E434.htm",
				Jobs:        "721",
				Reviews:     "6.5K",
				Salaries:    "11K",
				ReviewsRate: "3.6",
				Verified:    true,
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
							Title:                "Software Engineer (JavaScript / Ruby / Elixir)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4299028566/",
							Location:             "Columbus, OH",
							Date:                 mustDate("2025-10-04", "2025-09-11"),
							WithSalary:           true, // $106.9k/yr - $178.1k/yr
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
			ShortDescription: "Healthcare company",
			Industries: []domain.Industry{
				domain.IndustryHealthTech,
			},
		},
		{
			ID:         0,  // system
			Type:       "", // system
			Name:       "Commerce",
			BaseURL:    "https://www.commerce.com/",
			CareersURL: "https://www.commerce.com/careers/",
			AboutURL:   "https://www.commerce.com/about/",
			BlogURL:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                108096657,
				IDs:               nil,
				Alias:             "poweredbycommerce",
				Name:              "Commerce",
				Followers:         "3K",
				Employees:         "1K-5K",
				AssociatedMembers: "1,731",
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
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Commerce-EI_IE10729914.11,19.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Commerce-Reviews-E10729914.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Commerce-Jobs-E10729914.htm",
				Jobs:        "32",
				Reviews:     "10",
				Salaries:    "11",
				ReviewsRate: "4.7",
				Verified:    true,
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
							URL:                  "https://www.linkedin.com/jobs/view/4286318398/",
							Location:             "United States",
							Date:                 mustDate("2025-09-12"),
							WithSalary:           true, // $112k/yr - $189k/yr
							Remote:               true,
						},
						{
							Title:                "Senior Software Engineer – Python/Go",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4337919671/",
							Location:             "Sydney, New South Wales, Australia",
							Date:                 mustDate("2025-12-21", "2025-11-07"),
							WithSalary:           false,
							Remote:               true,
						},
						{
							Title:                "Senior Software Engineer – PHP/Go",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4337889761/",
							Location:             "United States",
							Date:                 mustDate("2025-11-29"),
							WithSalary:           true, // $112k/yr - $190k/yr
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
					Vacancies: []domain.Vacancy{
						{
							Title:                "Software Engineer II (Scala)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4299256946/",
							Location:             "United Kingdom",
							Date:                 mustDate("2025-09-16"),
							WithSalary:           false,
							Remote:               true,
						},
						{
							Title:                "Senior Software Engineer (Scala)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4299264966/",
							Location:             "Ireland",
							Date:                 mustDate("2025-10-08"),
							WithSalary:           false,
							Remote:               true,
						},
						{
							Title:                "Lead Software Engineer (Scala)",
							ShortDescription:     "Storefront",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4355394004/",
							Location:             "Austin, TX",
							CloudProviders:       []domain.CloudProvider{},
							Date:                 mustDate("2026-01-24"),
							WithSalary:           true, // $158.3k/yr - $237.4k/yr
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
			ShortDescription: "We help brands unlock the full potential of their data, connect systems, and deliver seamless, personalized experiences across every channel",
			Industries:       []domain.Industry{
				// NOP
			},
			YahooFinanceURL:  "https://finance.yahoo.com/quote/CMRC/",
			GoogleFinanceURL: "https://www.google.com/finance/quote/CMRC:NASDAQ",
		},
		{
			ID:         0,  // system
			Type:       "", // system
			Name:       "Ciena",
			BaseURL:    "https://www.ciena.com/",
			CareersURL: "https://www.ciena.com/careers",
			AboutURL:   "https://www.ciena.com/about",
			BlogURL:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                3960,
				IDs:               []int{3960, 115098, 2171992},
				Alias:             "ciena",
				Name:              "Ciena",
				Followers:         "269K",
				Employees:         "5K-10K",
				AssociatedMembers: "10,128",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:     "ciena",
				Followers: "28",
				Verified:  false,
			},
			BlindProfile: domain.BlindProfile{
				Alias: "",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "ciena",
				Employees: "6,250",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Ciena-EI_IE6612.11,16.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Ciena-Reviews-E6612.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Ciena-Jobs-E6612.htm",
				Jobs:        "174",
				Reviews:     "2.5K",
				Salaries:    "3.8K",
				ReviewsRate: "4.1",
				Verified:    true,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "",
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoryCount: 8,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Senior Golang Developer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4300493896/",
							Location:             "Gurugram, Haryana, India",
							Date:                 mustDate("2025-11-20"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Senior Golang Developer",
							ShortDescription:     "Cloud & Kubernetes",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4327785528/",
							Location:             "Gurugram, Haryana, India",
							Date:                 mustDate("2025-12-16"),
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
					Vacancies: []domain.Vacancy{
						{
							Title:                "Senior Scala Developer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4301301751/",
							Location:             "Gurugram, Haryana, India",
							Date:                 mustDate("2025-10-08", "2025-09-16"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Senior Scala Developer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4301188938/",
							Location:             "Pune, Maharashtra, India",
							Date:                 mustDate("2025-10-31"),
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
			ShortDescription: "Data transmission and network management",
			Industries: []domain.Industry{
				domain.IndustryTelecom,
			},
			YahooFinanceURL:  "https://finance.yahoo.com/quote/CIEN/",
			GoogleFinanceURL: "https://www.google.com/finance/quote/CIEN:NYSE",
		},
		{
			ID:         0,  // system
			Type:       "", // system
			Name:       "Amber Group",
			BaseURL:    "https://www.ambergroup.io/",
			CareersURL: "https://www.ambergroup.io/careers",
			AboutURL:   "",
			BlogURL:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                13451447,
				IDs:               nil,
				Alias:             "amberbtc",
				Name:              "Amber Group",
				Followers:         "56K",
				Employees:         "201-500",
				AssociatedMembers: "588",
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
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Rust: {
					GitHubRepositoryCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Software Engineer (Go, Rust)",
							ShortDescription:     "Platform Trading",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4268498746/",
							Location:             "Hong Kong, Hong Kong SAR",
							Date:                 mustDate("2025-09-13"),
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
			ShortDescription: "",
			Industries:       []domain.Industry{
				// NOP
			},
			Ignore: true, // Cryptocurrency
		},
		{
			ID:         0,  // system
			Type:       "", // system
			Name:       "Parcelhero",
			BaseURL:    "https://www.parcelhero.com/",
			CareersURL: "",
			AboutURL:   "",
			BlogURL:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                2246954,
				IDs:               nil,
				Alias:             "parcelhero-com-limited",
				Name:              "Parcelhero",
				Followers:         "3K",
				Employees:         "51-200",
				AssociatedMembers: "69",
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
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Rust: {
					GitHubRepositoryCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Software Engineer Rust/Shopify",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4299491739/",
							Location:             "Philippines",
							Date:                 mustDate("2025-09-13"),
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
			ShortDescription: "Parcel delivery service",
			Industries: []domain.Industry{
				domain.IndustryLogisticsTech,
			},
		},
		{
			ID:         0,  // system
			Type:       "", // system
			Name:       "Admiral Money",
			BaseURL:    "https://money.admiral.com/",
			CareersURL: "",
			AboutURL:   "",
			BlogURL:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                80031909,
				IDs:               nil,
				Alias:             "admiral-money",
				Name:              "Admiral Money",
				Followers:         "4K",
				Employees:         "201-500",
				AssociatedMembers: "270",
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
					GitHubRepositoryCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Senior Software Engineer – Java or Golang",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4300700392/",
							Location:             "United Kingdom",
							Date:                 mustDate("2025-09-16"),
							WithSalary:           false,
							Remote:               true,
						},
						{
							Title:                "Senior Software Engineer – Golang or Java",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4317910513/",
							Location:             "United Kingdom",
							Date:                 mustDate("2025-12-05"),
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
			Name:       "MatchMove",
			BaseURL:    "https://matchmove.com/",
			CareersURL: "",
			AboutURL:   "https://matchmove.com/about-us",
			BlogURL:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                275523,
				IDs:               nil,
				Alias:             "matchmove",
				Name:              "MatchMove",
				Followers:         "16K",
				Employees:         "51-200",
				AssociatedMembers: "112",
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
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-MatchMovePay-EI_IE996396.11,23.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/MatchMovePay-Reviews-E996396.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/MatchMovePay-Jobs-E996396.htm",
				Jobs:        "",
				Reviews:     "141",
				Salaries:    "94",
				ReviewsRate: "4.0",
				Verified:    false,
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
							URL:                  "https://www.linkedin.com/jobs/view/4300991599/",
							Location:             "Chennai, Tamil Nadu, India",
							Date:                 mustDate("2025-09-16"),
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
			ShortDescription: "Banking-as-a-Service (BaaS)",
			Industries: []domain.Industry{
				domain.IndustryFinTech,
			},
		},
		{
			ID:         0,  // system
			Type:       "", // system
			Name:       "Arctic Wolf",
			BaseURL:    "https://arcticwolf.com/",
			CareersURL: "https://arcticwolf.com/company/careers/",
			AboutURL:   "https://arcticwolf.com/company/overview/",
			BlogURL:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                2760138,
				IDs:               []int{2760138, 36094474},
				Alias:             "arcticwolf",
				Name:              "Arctic Wolf",
				Followers:         "123K",
				Employees:         "1K-5K",
				AssociatedMembers: "3,224",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:     "rtkwlf",
				Followers: "374",
				Verified:  false,
			},
			BlindProfile: domain.BlindProfile{
				Alias: "",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias: "",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Arctic-Wolf-EI_IE1393965.11,22.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Arctic-Wolf-Reviews-E1393965.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Arctic-Wolf-Jobs-E1393965.htm",
				Jobs:        "1",
				Reviews:     "320",
				Salaries:    "684",
				ReviewsRate: "3.8",
				Verified:    false,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "",
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoryCount: 7,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Developer – Back-End (Golang)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4290540051/",
							Location:             "Bengaluru, Karnataka, India",
							Date:                 mustDate("2025-11-20", "2025-10-08", "2025-09-16"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Senior Developer (Golang)",
							ShortDescription:     "Managed Risk Platform",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4344030443/",
							Location:             "Bengaluru, Karnataka, India",
							Date:                 mustDate("2025-12-17"),
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
			ShortDescription: "We help organizations end cyber risk by providing security operations as a concierge service",
			Industries: []domain.Industry{
				domain.IndustryCyberSecurity,
			},
		},
		{
			ID:         0,  // system
			Type:       "", // system
			Name:       "DataNimbus",
			BaseURL:    "https://www.datanimbus.com/",
			CareersURL: "https://datanimbus.com/careers/",
			AboutURL:   "https://datanimbus.com/about-us/",
			BlogURL:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                31059088,
				IDs:               nil,
				Alias:             "datanimbusinc",
				Name:              "DataNimbus",
				Followers:         "41K",
				Employees:         "51-200",
				AssociatedMembers: "157",
				Verified:          false,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:     "datanimbus",
				Followers: "6",
				Verified:  false,
			},
			BlindProfile: domain.BlindProfile{
				Alias: "",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias: "",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-DataNimbus-EI_IE3265873.11,21.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/DataNimbus-Reviews-E3265873.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/DataNimbus-Jobs-E3265873.htm",
				Jobs:        "",
				Reviews:     "24",
				Salaries:    "39",
				ReviewsRate: "4.6",
				Verified:    false,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "",
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoryCount: 3,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Back-End Developer – Golang",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4299236251/",
							Location:             "Pune, Maharashtra, India",
							Date:                 mustDate("2025-09-16"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Golang Developer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4338408878/",
							Location:             "Chennai, Tamil Nadu, India",
							Date:                 mustDate("2025-11-22"),
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
			ShortDescription: "Cloud-native AI-powered intelligent automation platform",
			Industries:       []domain.Industry{
				// NOP
			},
		},
		{
			ID:         0,  // system
			Type:       "", // system
			Name:       "GE Vernova",
			BaseURL:    "https://www.gevernova.com/",
			CareersURL: "https://careers.gevernova.com/",
			AboutURL:   "https://www.gevernova.com/company/about",
			BlogURL:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                82315716,
				IDs:               []int{1021, 2444, 751723, 1957003, 2631616, 5165055, 16252943, 68318371, 82315716, 89995973, 90505467},
				Alias:             "gevernova",
				Name:              "GE Vernova",
				Followers:         "3M",
				Employees:         "10K+",
				AssociatedMembers: "59,845",
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
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-GE-Vernova-EI_IE740889.11,21.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/GE-Vernova-Reviews-E740889.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/GE-Vernova-Jobs-E740889.htm",
				Jobs:        "2.8K",
				Reviews:     "3.1K",
				Salaries:    "4.7K",
				ReviewsRate: "3.9",
				Verified:    false,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "",
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoryCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Senior Platform Engineer – Python/Go",
							ShortDescription:     "CI/CD Tooling",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4265261756/",
							Location:             "Bucharest, Bucharest, Romania",
							Date:                 mustDate("2025-09-15"),
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
			ShortDescription: "GE Vernova is a purpose-built energy technology company on a mission to electrify to thrive and decarbonize the world",
			Industries:       []domain.Industry{
				// NOP
			},
		},
		{
			ID:         0,  // system
			Type:       "", // system
			Name:       "VisionPlus",
			BaseURL:    "https://visionxpay.com/",
			CareersURL: "",
			AboutURL:   "https://visionxpay.com/about-us/",
			BlogURL:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                105445508,
				IDs:               nil,
				Alias:             "visionxplus",
				Name:              "VisionPlus",
				Followers:         "9K",
				Employees:         "11-50",
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
					GitHubRepositoryCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Senior Go Developer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4299119401/",
							Location:             "Bengaluru, Karnataka, India",
							Date:                 mustDate("2025-09-12"),
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
			Name:       "Teads",
			BaseURL:    "https://www.teads.com/",
			CareersURL: "https://www.teads.com/about-teads/",
			AboutURL:   "https://www.teads.com/teads-careers/",
			BlogURL:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                2691523,
				IDs:               []int{2691523, 70931947, 100988850},
				Alias:             "teads",
				Name:              "Teads",
				Followers:         "153K",
				Employees:         "1K-5K",
				AssociatedMembers: "1,825",
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
					GitHubRepositoryCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Senior Backend Engineer (Java/Scala/Go)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4316278689/",
							Location:             "Netanya, Center District, Israel",
							Date:                 mustDate("2025-11-15"),
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
					Vacancies: []domain.Vacancy{
						{
							Title:                "Senior Full Stack – Backend Oriented",
							ShortDescription:     "We use mostly Scala and TypeScript",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4293125191/",
							Location:             "Netanya, Center District, Israel",
							Date:                 mustDate("2025-09-10"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Senior Backend Engineer (Java/Scala/Go)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4316278689/",
							Location:             "Netanya, Center District, Israel",
							Date:                 mustDate("2025-11-15"),
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
			ShortDescription: "Omnichannel outcomes platform",
			Industries: []domain.Industry{
				domain.IndustryAdTech,
			},
			YahooFinanceURL:  "https://finance.yahoo.com/quote/TEAD/",
			GoogleFinanceURL: "https://www.google.com/finance/quote/TEAD:NASDAQ",
		},
		{
			ID:         0,  // system
			Type:       "", // system
			Name:       "CircuitHub",
			BaseURL:    "https://www.circuithub.com/",
			CareersURL: "https://www.circuithub.com/company/careers",
			AboutURL:   "https://www.circuithub.com/company/about-us",
			BlogURL:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                2928643,
				IDs:               nil,
				Alias:             "circuithub",
				Name:              "CircuitHub",
				Followers:         "3K",
				Employees:         "51-200",
				AssociatedMembers: "40",
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
					Vacancies: []domain.Vacancy{
						{
							Title:                "Haskell Engineer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4302201309/",
							Location:             "United Kingdom",
							Date:                 mustDate("2025-09-19"),
							WithSalary:           false,
							Remote:               true,
						},
					},
				},
				domain.Erlang: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
			},
			ShortDescription: "CircuitHub is on a mission to fix rapid electronics prototyping",
			Industries:       []domain.Industry{
				// NOP
			},
		},
		{
			ID:         0,  // system
			Type:       "", // system
			Name:       "edclub",
			BaseURL:    "https://www.edclub.com/",
			CareersURL: "",
			AboutURL:   "",
			BlogURL:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                5387783,
				IDs:               nil,
				Alias:             "edclub",
				Name:              "edclub",
				Followers:         "2K",
				Employees:         "11-50",
				AssociatedMembers: "44",
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
					GitHubRepositoryCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Lead Software Developer – Go & Python",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4303218078/",
							Location:             "Rockville, MD",
							Date:                 mustDate("2025-09-20"),
							WithSalary:           true, // starting at $130k/yr
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
			ShortDescription: "Learning platform",
			Industries: []domain.Industry{
				domain.IndustryEdTech,
			},
		},
		{
			ID:         0,  // system
			Type:       "", // system
			Name:       "AppZen",
			BaseURL:    "https://www.appzen.com/",
			CareersURL: "https://www.appzen.com/careers",
			AboutURL:   "https://www.appzen.com/about-us",
			BlogURL:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                2829210,
				IDs:               nil,
				Alias:             "appzen",
				Name:              "AppZen",
				Followers:         "29K",
				Employees:         "201-500",
				AssociatedMembers: "378",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:     "appzen-oss",
				Followers: "",
				Verified:  true,
			},
			BlindProfile: domain.BlindProfile{
				Alias: "",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias: "",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-AppZen-EI_IE1284188.11,17.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/AppZen-Reviews-E1284188.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/AppZen-Jobs-E1284188.htm",
				Jobs:        "39",
				Reviews:     "222",
				Salaries:    "284",
				ReviewsRate: "3.2",
				Verified:    false,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "",
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoryCount: 1,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Golang Developer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4302857179/",
							Location:             "Pune, Maharashtra, India",
							Date:                 mustDate("2025-09-20"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Golang Developer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4313348410/",
							Location:             "Pune, Maharashtra, India",
							Date:                 mustDate("2025-10-18"),
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
			ShortDescription: "AppZen delivers autonomous spend management for expenses, invoices, and cards",
			Industries: []domain.Industry{
				domain.IndustryFinTech,
			},
		},
		{
			ID:         0,  // system
			Type:       "", // system
			Name:       "Yonder",
			BaseURL:    "https://www.yonder.com/",
			CareersURL: "https://www.yonder.com/careers",
			AboutURL:   "",
			BlogURL:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                72211530,
				IDs:               nil,
				Alias:             "yondercard",
				Name:              "Yonder",
				Followers:         "13K",
				Employees:         "51-200",
				AssociatedMembers: "102",
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
							Title:                "Senior Product Engineer (Scala)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4304070735/",
							Location:             "London, England, United Kingdom",
							Date:                 mustDate("2025-09-23"),
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
				domain.IndustryFinTech,
			},
		},
		{
			ID:         0,  // system
			Type:       "", // system
			Name:       "Eskimi",
			BaseURL:    "https://www.eskimi.com/",
			CareersURL: "https://www.eskimi.com/about-us",
			AboutURL:   "https://careers.eskimi.com/",
			BlogURL:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                2493050,
				IDs:               nil,
				Alias:             "eskimi",
				Name:              "Eskimi",
				Followers:         "38K",
				Employees:         "201-500",
				AssociatedMembers: "238",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:     "eskimi",
				Followers: "2",
				Verified:  false,
			},
			BlindProfile: domain.BlindProfile{
				Alias: "",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias: "",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Eskimi-EI_IE1332996.11,17.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Eskimi-Reviews-E1332996.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Eskimi-Jobs-E1332996.htm",
				Jobs:        "4",
				Reviews:     "78",
				Salaries:    "65",
				ReviewsRate: "3.9",
				Verified:    false,
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
							Title:                "Senior Software Engineer (Java/Scala)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4304081084/",
							Location:             "Dhaka, Dhaka, Bangladesh",
							Date:                 mustDate("2025-09-23"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Lead Software Engineer (Java/Scala)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4320326842/",
							Location:             "Dhaka, Dhaka, Bangladesh",
							Date:                 mustDate("2025-11-03"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Junior Software Engineer (Java/Scala)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4327605030/",
							Location:             "Dhaka, Dhaka, Bangladesh",
							Date:                 mustDate("2025-12-15"),
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
			ShortDescription: "Creative and media tech platform",
			Industries: []domain.Industry{
				domain.IndustryAdTech,
			},
		},
		{
			ID:         0,  // system
			Type:       "", // system
			Name:       "Databento",
			BaseURL:    "https://databento.com/",
			CareersURL: "https://databento.com/careers",
			AboutURL:   "https://databento.com/about",
			BlogURL:    "https://databento.com/blog/engineering",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                35540938,
				IDs:               nil,
				Alias:             "databento",
				Name:              "Databento",
				Followers:         "17K",
				Employees:         "11-50",
				AssociatedMembers: "31",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:     "databento",
				Followers: "293",
				Verified:  true,
			},
			BlindProfile: domain.BlindProfile{
				Alias: "",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias: "",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Databento-EI_IE5650624.11,20.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Databento-Reviews-E5650624.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Databento-Jobs-E5650624.htm",
				Jobs:        "2",
				Reviews:     "2",
				Salaries:    "4",
				ReviewsRate: "5.0",
				Verified:    true,
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
					GitHubRepositoryCount: 3,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Software Engineer (C++/Rust)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4304492468/",
							Location:             "Chicago, IL",
							Date:                 mustDate("2025-09-23"),
							WithSalary:           false,
							Remote:               true,
						},
						{
							Title:                "Software Engineer (C++/Rust)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4347516514/",
							Location:             "Chicago, IL",
							Date:                 mustDate("2025-11-26"),
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
			ShortDescription: "Databento makes it simpler and faster to get market data",
			Industries: []domain.Industry{
				domain.IndustryFinTech,
			},
		},
		{
			ID:         0,  // system
			Type:       "", // system
			Name:       "Zattoo",
			BaseURL:    "https://zattoo.com/",
			CareersURL: "https://zattoo.com/company/jobs",
			AboutURL:   "https://zattoo.com/company",
			BlogURL:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                109704,
				IDs:               nil,
				Alias:             "zattoo",
				Name:              "Zattoo",
				Followers:         "8K",
				Employees:         "201-500",
				AssociatedMembers: "247",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:     "zattoo",
				Followers: "24",
				Verified:  true,
			},
			BlindProfile: domain.BlindProfile{
				Alias: "",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias: "",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Zattoo-EI_IE278956.11,17.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Zattoo-Reviews-E278956.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Zattoo-Jobs-E278956.htm",
				Jobs:        "15",
				Reviews:     "60",
				Salaries:    "137",
				ReviewsRate: "3.7",
				Verified:    false,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "",
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoryCount: 5,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Backend Systems Engineer (C++, Golang)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4302636216/",
							Location:             "Berlin, Berlin, Germany",
							Date:                 mustDate("2025-09-23"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Backend Systems Engineer (C++, Golang)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4347283787/",
							Location:             "Berlin, Berlin, Germany",
							Date:                 mustDate("2025-11-23"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Backend Systems Engineer (C++, Golang)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4344296359/",
							Location:             "Berlin, Berlin, Germany",
							Date:                 mustDate("2025-12-17"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Backend Systems Engineer (C++, Golang)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4353194797/",
							Location:             "Berlin, Berlin, Germany",
							CloudProviders:       []domain.CloudProvider{},
							Date:                 mustDate("2026-01-16"),
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
			ShortDescription: "TV streaming platform",
			Industries:       []domain.Industry{
				// NOP
			},
		},
		{
			ID:         0,  // system
			Type:       "", // system
			Name:       "Oxylabs.io",
			BaseURL:    "https://oxylabs.io/",
			CareersURL: "https://career.oxylabs.io/",
			AboutURL:   "",
			BlogURL:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                10349119,
				IDs:               nil,
				Alias:             "oxylabs-io",
				Name:              "Oxylabs.io",
				Followers:         "28K",
				Employees:         "201-500",
				AssociatedMembers: "401",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:     "oxylabs",
				Followers: "587",
				Verified:  true,
			},
			BlindProfile: domain.BlindProfile{
				Alias: "",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias: "",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Oxylabs-EI_IE3226058.11,18.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Oxylabs-Reviews-E3226058.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Oxylabs-Jobs-E3226058.htm",
				Jobs:        "",
				Reviews:     "55",
				Salaries:    "90",
				ReviewsRate: "4.2",
				Verified:    false,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "",
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoryCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "PHP Developer (Golang)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4304189615/",
							Location:             "Vilnius, Vilniaus, Lithuania",
							Date:                 mustDate("2025-09-23"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Golang Developer (Senior)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4344358769/",
							Location:             "Vilnius, Vilniaus, Lithuania",
							Date:                 mustDate("2025-12-18"),
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
			ShortDescription: "We offer web scraping solutions",
			Industries:       []domain.Industry{
				// NOP
			},
		},
		{
			ID:         0,  // system
			Type:       "", // system
			Name:       "Pentaleap",
			BaseURL:    "https://www.pentaleap.com/",
			CareersURL: "",
			AboutURL:   "https://www.pentaleap.com/about-us",
			BlogURL:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                104735661,
				IDs:               nil,
				Alias:             "pentaleaptech",
				Name:              "Pentaleap",
				Followers:         "2K",
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
					GitHubRepositoryCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Senior Backend Engineer (Go)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4302637338/",
							Location:             "Germany",
							Date:                 mustDate("2025-09-23"),
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
			ShortDescription: "Retail Media Platform",
			Industries:       []domain.Industry{
				// NOP
			},
		},
		{
			ID:         0,  // system
			Type:       "", // system
			Name:       "Doodle",
			BaseURL:    "https://doodle.com/",
			CareersURL: "https://careers.doodle.com/",
			AboutURL:   "",
			BlogURL:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                3199137,
				IDs:               nil,
				Alias:             "doodle-ag",
				Name:              "Doodle",
				Followers:         "6K",
				Employees:         "51-200",
				AssociatedMembers: "113",
				Verified:          false,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:     "doodlescheduling",
				Followers: "14",
				Verified:  false,
			},
			BlindProfile: domain.BlindProfile{
				Alias: "",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias: "",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Doodle-EI_IE1835186.11,17.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Doodle-Reviews-E1835186.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Doodle-Jobs-E1835186.htm",
				Jobs:        "7",
				Reviews:     "53",
				Salaries:    "97",
				ReviewsRate: "3.7",
				Verified:    false,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "",
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoryCount: 31,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Software Engineer, Backend (Go/Java, Reactive Systems)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4304071155/",
							Location:             "Warsaw, Mazowieckie, Poland",
							Date:                 mustDate("2025-09-23"),
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
			ShortDescription: "Schedule meetings and events",
			Industries:       []domain.Industry{
				// NOP
			},
		},
		{
			ID:         0,  // system
			Type:       "", // system
			Name:       "CollX",
			BaseURL:    "https://www.collx.app/",
			CareersURL: "https://www.collx.app/career/job-openings",
			AboutURL:   "",
			BlogURL:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                80130229,
				IDs:               nil,
				Alias:             "collx",
				Name:              "CollX",
				Followers:         "2K",
				Employees:         "11-50",
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
							Title:                "Senior Elixir Engineer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4306775477/",
							Location:             "Greater Philadelphia",
							Date:                 mustDate("2025-09-26"),
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
			ShortDescription: "Scan any trading card to get the value in seconds, and begin building your collection",
			Industries:       []domain.Industry{
				// NOP
			},
		},
		{
			ID:         0,  // system
			Type:       "", // system
			Name:       "Fudo Security",
			BaseURL:    "https://www.fudosecurity.com/",
			CareersURL: "https://www.fudosecurity.com/careers",
			AboutURL:   "https://www.fudosecurity.com/about-us",
			BlogURL:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                216468,
				IDs:               nil,
				Alias:             "fudosecurity",
				Name:              "Fudo Security",
				Followers:         "6K",
				Employees:         "51-200",
				AssociatedMembers: "83",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:     "fudosecurity",
				Followers: "5",
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
							URL:                  "https://www.linkedin.com/jobs/view/4305466070/",
							Location:             "Poland",
							Date:                 mustDate("2025-09-27"),
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
			ShortDescription: "Fudo Security delivers advanced Privileged Access Management (PAM)",
			Industries: []domain.Industry{
				domain.IndustryCyberSecurity,
			},
		},
		{
			ID:         0,  // system
			Type:       "", // system
			Name:       "Bayer",
			BaseURL:    "https://www.bayer.com/",
			CareersURL: "https://www.bayer.com/career",
			AboutURL:   "",
			BlogURL:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                1893,
				IDs:               []int{1893, 3422, 86467},
				Alias:             "bayer",
				Name:              "Bayer",
				Followers:         "6M",
				Employees:         "10K+",
				AssociatedMembers: "93,319",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:     "Bayer-Group",
				Followers: "2.4k",
				Verified:  true,
			},
			BlindProfile: domain.BlindProfile{
				Alias: "",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias: "",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Bayer-EI_IE4245.11,16.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Bayer-Reviews-E4245.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Bayer-Jobs-E4245.htm",
				Jobs:        "562",
				Reviews:     "12K",
				Salaries:    "20K",
				ReviewsRate: "3.9",
				Verified:    true,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "",
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoryCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Staff Data Engineer (Golang)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4292978839/",
							Location:             "United States",
							Date:                 mustDate("2025-12-12", "2025-11-20", "2025-09-27"),
							WithSalary:           true, // $119.1k/yr - $178.6k/yr
							Remote:               true,
						},
						{
							Title:                "Senior Staff Data Engineer (Golang/Python)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4300971531/",
							Location:             "Creve Coeur, MO",
							Date:                 mustDate("2025-10-08"),
							WithSalary:           true, // $136.9k/yr - $205.4k/yr
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
				domain.IndustryHealthTech,
			},
		},
		{
			ID:         0,  // system
			Type:       "", // system
			Name:       "Vereign AG",
			BaseURL:    "https://www.vereign.com/",
			CareersURL: "https://www.vereign.com/careers",
			AboutURL:   "https://www.vereign.com/about",
			BlogURL:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                18356285,
				IDs:               nil,
				Alias:             "vereign",
				Name:              "Vereign AG",
				Followers:         "913",
				Employees:         "11-50",
				AssociatedMembers: "17",
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
					GitHubRepositoryCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Software Engineer (Golang)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4305089239/",
							Location:             "Sofia, Sofia City, Bulgaria",
							Date:                 mustDate("2025-09-29"),
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
			ShortDescription: "Secure communication, data exchange and enterprise edge applications",
			Industries: []domain.Industry{
				domain.IndustryCyberSecurity,
			},
		},
		{
			ID:         0,  // system
			Type:       "", // system
			Name:       "KNOREX",
			BaseURL:    "https://www.knorex.com/",
			CareersURL: "https://www.knorex.com/careers/",
			AboutURL:   "https://www.knorex.com/about/",
			BlogURL:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                737596,
				IDs:               nil,
				Alias:             "knorex",
				Name:              "KNOREX",
				Followers:         "43K",
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
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Knorex-EI_IE653201.11,17.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Knorex-Reviews-E653201.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Knorex-Jobs-E653201.htm",
				Jobs:        "7",
				Reviews:     "115",
				Salaries:    "121",
				ReviewsRate: "3.8",
				Verified:    false,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "",
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoryCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Software Engineer (C++/ Golang)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4306348703/",
							Location:             "Ho Chi Minh City, Vietnam",
							Date:                 mustDate("2025-09-30"),
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
			ShortDescription: "Knorex is a technology company that provides programmatic advertising solutions",
			Industries: []domain.Industry{
				domain.IndustryAdTech,
			},
		},
		{
			ID:         0,  // system
			Type:       "", // system
			Name:       "JustRelate",
			BaseURL:    "https://www.justrelate.com/",
			CareersURL: "https://www.justrelate.com/careers",
			AboutURL:   "https://www.justrelate.com/about-us",
			BlogURL:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                206241,
				IDs:               nil,
				Alias:             "justrelate-group",
				Name:              "JustRelate",
				Followers:         "33K",
				Employees:         "51-200",
				AssociatedMembers: "131",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:     "infopark",
				Followers: "9",
				Verified:  true,
			},
			BlindProfile: domain.BlindProfile{
				Alias: "",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "justrelate-group",
				Employees: "150",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-JustRelate-EI_IE4712417.11,21.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/JustRelate-Reviews-E4712417.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/JustRelate-Jobs-E4712417.htm",
				Jobs:        "4",
				Reviews:     "6",
				Salaries:    "15",
				ReviewsRate: "4.4",
				Verified:    false,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "",
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoryCount: 2,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Senior Golang Engineer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4307389811/",
							Location:             "Germany",
							Date:                 mustDate("2025-10-01"),
							WithSalary:           false,
							Remote:               true,
						},
						{
							Title:                "Senior Go Developer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4319190761/",
							Location:             "Poland",
							Date:                 mustDate("2025-10-28"),
							WithSalary:           false,
							Remote:               true,
						},
						{
							Title:                "Senior Go Developer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4321081652/",
							Location:             "Poland",
							Date:                 mustDate("2025-11-06"),
							WithSalary:           false,
							Remote:               true,
						},
						{
							Title:                "Senior Go Developer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4323610162/",
							Location:             "Poland",
							Date:                 mustDate("2025-11-20"),
							WithSalary:           false,
							Remote:               true,
						},
						{
							Title:                "Senior Go Developer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4324932042/",
							Location:             "Poland",
							Date:                 mustDate("2025-11-29"),
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
			ShortDescription: "Customer engagement platforms for companies with complex offerings",
			Industries:       []domain.Industry{
				// NOP
			},
		},
		{
			ID:         0,  // system
			Type:       "", // system
			Name:       "SecureCo",
			BaseURL:    "https://www.secureco.com/",
			CareersURL: "https://www.secureco.com/careers/",
			AboutURL:   "https://www.secureco.com/about/",
			BlogURL:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                64894717,
				IDs:               nil,
				Alias:             "secureco-usa",
				Name:              "SecureCo",
				Followers:         "2K",
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
				domain.Go: {
					GitHubRepositoryCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Senior Software Engineer – Golang",
							ShortDescription:     "Distributed Systems ",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4307699883/",
							Location:             "United States",
							Date:                 mustDate("2025-10-02"),
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
			ShortDescription: "Data transport and API protection made secure and resilient via obfuscation, shielding data and network assets",
			Industries: []domain.Industry{
				domain.IndustryCyberSecurity,
			},
		},
		{
			ID:         0,  // system
			Type:       "", // system
			Name:       "DATATRONiQ",
			BaseURL:    "https://datatroniq.com/",
			CareersURL: "",
			AboutURL:   "",
			BlogURL:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                10120513,
				IDs:               nil,
				Alias:             "datatroniq",
				Name:              "DATATRONiQ",
				Followers:         "1K",
				Employees:         "11-50",
				AssociatedMembers: "9",
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
					GitHubRepositoryCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "(Senior) Developer Backend (Golang)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4305014286/",
							Location:             "Berlin, Berlin, Germany",
							Date:                 mustDate("2025-09-28"),
							WithSalary:           true, // €70k/yr - €80k/yr
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
			ShortDescription: "The Industrial Internet of Things (IIoT) with its connected always-on devices emits massive data streams and information",
			Industries:       []domain.Industry{
				// NOP
			},
		},
		{
			ID:         0,  // system
			Type:       "", // system
			Name:       "DEUNA",
			BaseURL:    "https://www.deuna.com/",
			CareersURL: "https://www.deuna.com/careers",
			AboutURL:   "",
			BlogURL:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                68605015,
				IDs:               nil,
				Alias:             "de-una",
				Name:              "DEUNA",
				Followers:         "22K",
				Employees:         "51-200",
				AssociatedMembers: "148",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:     "DUNA-E-Commmerce",
				Followers: "39",
				Verified:  false,
			},
			BlindProfile: domain.BlindProfile{
				Alias: "",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias:     "deuna",
				Employees: "126",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-DEUNA-EI_IE6942865.11,16.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/DEUNA-Reviews-E6942865.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/DEUNA-Jobs-E6942865.htm",
				Jobs:        "142",
				Reviews:     "30",
				Salaries:    "16",
				ReviewsRate: "3.4",
				Verified:    true,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "",
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoryCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Senior Golang Backend Developer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4307703206/",
							Location:             "Querétaro, Querétaro, Mexico",
							Date:                 mustDate("2025-09-29"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Golang Backend Developer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4312047910/",
							Location:             "Querétaro, Querétaro, Mexico",
							Date:                 mustDate("2025-10-08"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Senior Golang Backend Developer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4335497407/",
							Location:             "Bogota, D.C., Capital District, Colombia",
							Date:                 mustDate("2025-11-12"),
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
			ShortDescription: "The unified platform to simplify global payments",
			Industries: []domain.Industry{
				domain.IndustryFinTech,
			},
		},
		{
			ID:         0,  // system
			Type:       "", // system
			Name:       "AvantStay",
			BaseURL:    "https://www.avantstay.com/",
			CareersURL: "https://avantstay.workable.com/",
			AboutURL:   "https://www.avantstay.com/about",
			BlogURL:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                17999166,
				IDs:               nil,
				Alias:             "avantstay",
				Name:              "AvantStay",
				Followers:         "59K",
				Employees:         "501-1K",
				AssociatedMembers: "410",
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
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-AvantStay-EI_IE1723986.11,20.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/AvantStay-Reviews-E1723986.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/AvantStay-Jobs-E1723986.htm",
				Jobs:        "64",
				Reviews:     "336",
				Salaries:    "411",
				ReviewsRate: "3.7",
				Verified:    true,
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
							Title:                "Senior Backend Engineer (Scala)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4308983560/",
							Location:             "Spain",
							Date:                 mustDate("2025-10-02"),
							WithSalary:           true, // $75k/yr - $100k/yr
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
			ShortDescription: "Hospitality platform",
			Industries:       []domain.Industry{
				// NOP
			},
		},
		{
			ID:         0,  // system
			Type:       "", // system
			Name:       "Kaluza",
			BaseURL:    "https://www.kaluza.com/",
			CareersURL: "https://careers.kaluza.com/",
			AboutURL:   "https://www.kaluza.com/about",
			BlogURL:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                1895538,
				IDs:               []int{1895538, 18703412},
				Alias:             "kaluza",
				Name:              "Kaluza",
				Followers:         "23K",
				Employees:         "501-1K",
				AssociatedMembers: "633",
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
				Alias:     "kaluza",
				Employees: "270",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Kaluza-EI_IE3218974.11,17.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Kaluza-Reviews-E3218974.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Kaluza-Jobs-E3218974.htm",
				Jobs:        "17",
				Reviews:     "87",
				Salaries:    "404",
				ReviewsRate: "3.6",
				Verified:    true,
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
							Title:                "Scala Engineer II",
							ShortDescription:     "API Platform Team",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4307236421/",
							Location:             "Bristol, England, United Kingdom",
							Date:                 mustDate("2025-10-03"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Scala Engineer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://app.welcometothejungle.com/jobs/SyvliXfm",
							Location:             "Bristol, England, United Kingdom",
							Date:                 mustDate("2025-11-14"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Scala Engineer",
							ShortDescription:     "API Platform Team",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4354044556/",
							Location:             "Bristol, England, United Kingdom",
							CloudProviders:       []domain.CloudProvider{},
							Date:                 mustDate("2026-01-20"),
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
			ShortDescription: "Energy Intelligence Platform",
			Industries:       []domain.Industry{
				// NOP
			},
		},
		{
			ID:         0,  // system
			Type:       "", // system
			Name:       "Zego",
			BaseURL:    "https://www.zego.com/",
			CareersURL: "https://www.zego.com/careers/",
			AboutURL:   "https://www.zego.com/about-us/",
			BlogURL:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                10854241,
				IDs:               []int{10854241, 77603493},
				Alias:             "zegocover",
				Name:              "Zego",
				Followers:         "63K",
				Employees:         "201-500",
				AssociatedMembers: "321",
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
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Zego-Insurance-EI_IE1851459.11,25.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Zego-Insurance-Reviews-E1851459.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Zego-Insurance-Jobs-E1851459.htm",
				Jobs:        "23",
				Reviews:     "319",
				Salaries:    "550",
				ReviewsRate: "4.0",
				Verified:    true,
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
							Title:                "Senior Scala Engineer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4308245274/",
							Location:             "United Kingdom",
							Date:                 mustDate("2026-01-22", "2025-12-10", "2025-11-19", "2025-10-08"),
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
			ShortDescription: "Zego is a motor insurance provider",
			Industries:       []domain.Industry{
				// NOP
			},
		},
		{
			ID:         0,  // system
			Type:       "", // system
			Name:       "Lobster Robotics",
			BaseURL:    "https://www.lobster-robotics.com/",
			CareersURL: "https://www.lobster-robotics.com/careers",
			AboutURL:   "https://www.lobster-robotics.com/about",
			BlogURL:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                14021985,
				IDs:               nil,
				Alias:             "lobster-robotics",
				Name:              "Lobster Robotics",
				Followers:         "4K",
				Employees:         "11-50",
				AssociatedMembers: "15",
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
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Rust: {
					GitHubRepositoryCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Rust Software Engineer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4309608120/",
							Location:             "Delft, South Holland, Netherlands",
							Date:                 mustDate("2025-10-03"),
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
			ShortDescription: "We design, manufacture, and rent out fully integrated underwater drones",
			Industries: []domain.Industry{
				domain.IndustryDefenseTech,
			},
		},
		{
			ID:         0,  // system
			Type:       "", // system
			Name:       "Affinidi",
			BaseURL:    "https://www.affinidi.com/",
			CareersURL: "",
			AboutURL:   "",
			BlogURL:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                53236067,
				IDs:               nil,
				Alias:             "affinidi",
				Name:              "Affinidi",
				Followers:         "38K",
				Employees:         "201-500",
				AssociatedMembers: "150",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:     "affinidi",
				Followers: "51",
				Verified:  true,
			},
			BlindProfile: domain.BlindProfile{
				Alias: "",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias: "",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Affinidi-EI_IE4213335.11,19.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Affinidi-Reviews-E4213335.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Affinidi-Jobs-E4213335.htm",
				Jobs:        "6",
				Reviews:     "110",
				Salaries:    "156",
				ReviewsRate: "3.2",
				Verified:    true,
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
					GitHubRepositoryCount: 5,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Staff Backend Engineer with Rust",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4308288567/",
							Location:             "Berlin, Berlin, Germany",
							Date:                 mustDate("2025-10-03"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Staff Rust Developer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4321884203/",
							Location:             "Dublin, County Dublin, Ireland",
							Date:                 mustDate("2025-12-02"),
							WithSalary:           false,
							Remote:               true,
						},
						{
							Title:                "Senior Backend Engineer (Rust)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4328320716/",
							Location:             "Berlin, Berlin, Germany",
							Date:                 mustDate("2025-12-20"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Senior Backend Engineer (Rust)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4328310746/",
							Location:             "Berlin, Berlin, Germany",
							Date:                 mustDate("2026-01-11"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Staff Rust Developer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4321944130/",
							Location:             "Greater London, England, United Kingdom",
							CloudProviders:       []domain.CloudProvider{},
							Date:                 mustDate("2026-01-15"),
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
			ShortDescription: "Identity management",
			Industries: []domain.Industry{
				domain.IndustryCyberSecurity,
			},
		},
		{
			ID:         0,  // system
			Type:       "", // system
			Name:       "Tensordyne",
			BaseURL:    "https://www.tensordyne.ai/",
			CareersURL: "https://www.tensordyne.ai/careers",
			AboutURL:   "https://www.tensordyne.ai/about",
			BlogURL:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                18326385,
				IDs:               nil,
				Alias:             "tensordyne",
				Name:              "Tensordyne",
				Followers:         "9K",
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
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Rust: {
					GitHubRepositoryCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Software Engineer – Rust",
							ShortDescription:     "Cloud Services",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4311159672/",
							Location:             "Germany",
							Date:                 mustDate("2025-10-07"),
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
			ShortDescription: "Math, Chips, and Systems for AI Inference",
			Industries:       []domain.Industry{
				// NOP
			},
		},
		{
			ID:         0,  // system
			Type:       "", // system
			Name:       "Rust Foundation",
			BaseURL:    "http://foundation.rust-lang.org/",
			CareersURL: "",
			AboutURL:   "",
			BlogURL:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                57208706,
				IDs:               nil,
				Alias:             "rust-foundation",
				Name:              "Rust Foundation",
				Followers:         "18K",
				Employees:         "2-10",
				AssociatedMembers: "37",
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
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Rust: {
					GitHubRepositoryCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Rust Infrastructure Engineer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4311203106/",
							Location:             "United States",
							Date:                 mustDate("2025-10-08"),
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
			ShortDescription: "The nonprofit steward of the Rust programming language and its global community",
			Industries:       []domain.Industry{
				// NOP
			},
		},
		{
			ID:         0,  // system
			Type:       "", // system
			Name:       "Redox One",
			BaseURL:    "http://www.redoxone.com/",
			CareersURL: "https://redoxone.com/careers/",
			AboutURL:   "https://redoxone.com/company/",
			BlogURL:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                91554046,
				IDs:               nil,
				Alias:             "redox-one",
				Name:              "Redox One",
				Followers:         "4K",
				Employees:         "51-200",
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
				domain.Go: {
					GitHubRepositoryCount: 0,
					Vacancies:             []domain.Vacancy{},
				},
				domain.Rust: {
					GitHubRepositoryCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Rust Full-Stack Developer",
							ShortDescription:     "Energy Systems",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4310409066/",
							Location:             "Dortmund, North Rhine-Westphalia, Germany",
							Date:                 mustDate("2025-10-04"),
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
			ShortDescription: "Redox One manufactures iron-chromium flow batteries for the large-scale energy storage requirements of businesses, industry and electricity networks",
			Industries:       []domain.Industry{
				// NOP
			},
		},
		{
			ID:         0,  // system
			Type:       "", // system
			Name:       "Tazapay",
			BaseURL:    "https://tazapay.com/",
			CareersURL: "https://tazapay.com/careers",
			AboutURL:   "https://tazapay.com/about-us",
			BlogURL:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                67153211,
				IDs:               nil,
				Alias:             "tazapay",
				Name:              "Tazapay",
				Followers:         "24K",
				Employees:         "51-200",
				AssociatedMembers: "198",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:     "tazapay",
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
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Tazapay-EI_IE6209842.11,18.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Tazapay-Reviews-E6209842.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Tazapay-Jobs-E6209842.htm",
				Jobs:        "",
				Reviews:     "32",
				Salaries:    "43",
				ReviewsRate: "3.6",
				Verified:    false,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "",
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoryCount: 3,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Senior Software Engineer – Golang/Python/Ruby",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4310491901/",
							Location:             "Bengaluru, Karnataka, India",
							Date:                 mustDate("2025-10-05"),
							WithSalary:           false,
							Remote:               false,
						},
						{
							Title:                "Software Engineer – Golang",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4328541476/",
							Location:             "Chennai, Tamil Nadu, India",
							Date:                 mustDate("2025-12-21"),
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
			Name:       "PAVE",
			BaseURL:    "https://pave.ai/",
			CareersURL: "",
			AboutURL:   "https://pave.ai/company",
			BlogURL:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                11341570,
				IDs:               nil,
				Alias:             "discovery-loft-inc",
				Name:              "PAVE",
				Followers:         "4K",
				Employees:         "51-200",
				AssociatedMembers: "68",
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
					GitHubRepositoryCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Senior Golang Engineer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4309720049/",
							Location:             "Ho Chi Minh City, Vietnam",
							Date:                 mustDate("2025-10-07"),
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
			ShortDescription: "Our intelligent platform enables consumers and dealers to complete a guided vehicle inspection that generates a comprehensive condition report in only minutes",
			Industries:       []domain.Industry{
				// NOP
			},
		},
		{
			ID:         0,  // system
			Type:       "", // system
			Name:       "Roku",
			BaseURL:    "https://www.roku.com/",
			CareersURL: "",
			AboutURL:   "",
			BlogURL:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                50076,
				IDs:               []int{50076, 220799},
				Alias:             "roku",
				Name:              "Roku",
				Followers:         "458K",
				Employees:         "1K-5K",
				AssociatedMembers: "4,135",
				Verified:          true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:     "rokudev",
				Followers: "485",
				Verified:  false,
			},
			BlindProfile: domain.BlindProfile{
				Alias: "",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias: "",
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Roku-EI_IE26760.11,15.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Roku-Reviews-E26760.htm",
				JobsURL:     "https://www.glassdoor.com/Jobs/Roku-Jobs-E26760.htm",
				Jobs:        "185",
				Reviews:     "911",
				Salaries:    "2.2K",
				ReviewsRate: "3.6",
				Verified:    true,
			},
			IndeedProfile: domain.IndeedProfile{
				Alias: "",
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoryCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Senior Software Engineer (Golang / Kubernetes)",
							ShortDescription:     "Observability Platform",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4312301338/",
							Location:             "Aarhus, Central Denmark Region, Denmark",
							Date:                 mustDate("2025-12-13", "2025-11-20", "2025-10-08"),
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
			ShortDescription: "Platform for streaming television",
			Industries:       []domain.Industry{
				// NOP
			},
		},
	}
}
