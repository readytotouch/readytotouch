package db

import "github.com/readytotouch/readytotouch/internal/domain"

func companies21Shard() []domain.CompanyProfile {
	return []domain.CompanyProfile{
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "CUJU",
			Website: "https://cuju.pro/",
			Careers: "",
			About:   "",
			Blog:    "",
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
					GitHubRepositoriesCount: 0,
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
			ShortDescription: "An AI based technology which gives football talents the chance to improve their skills, compete against friends and the global CUJU community and get scouted by a record-proven network of international senior scouts and specialists",
			Industries:       []domain.Industry{
				// NOP
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "AutoStore™",
			Website: "https://www.autostoresystem.com/",
			Careers: "https://www.autostoresystem.com/careers",
			About:   "https://www.autostoresystem.com/company",
			Blog:    "",
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
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Rust: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Back-End Software Engineer – C++, Rust or C#",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4281833526/",
							Location:             "Vindafjord, Rogaland, Norway",
							Date:                 mustDate("2025-09-06"),
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
			Name:    "Supabase",
			Website: "https://supabase.com/",
			Careers: "https://supabase.com/careers",
			About:   "https://supabase.com/company",
			Blog:    "",
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
					GitHubRepositoriesCount: 16,
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
					GitHubRepositoriesCount: 11,
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
			ShortDescription: "Supabase is the Postgres development platform",
			Industries: []domain.Industry{
				domain.IndustryDevOps,
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Gresham",
			Website: "https://www.greshamtech.com/",
			Careers: "https://www.greshamtech.com/careers",
			About:   "https://www.greshamtech.com/about/our-vision",
			Blog:    "",
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
							Title:                "Software Engineer Clojure",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4294972674/",
							Location:             "City Of Bristol, England, United Kingdom",
							Date:                 mustDate("2025-09-03"),
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
			ShortDescription: "Enterprise data automation solutions",
			Industries:       []domain.Industry{
				// NOP
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Multiverse",
			Website: "https://www.multiverse.io/",
			Careers: "https://www.multiverse.io/careers",
			About:   "https://www.multiverse.io/about",
			Blog:    "",
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
					GitHubRepositoriesCount: 19,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Senior Software Engineer",
							ShortDescription:     "Our stack includes Python, TypeScript, React, Elixir, AWS",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4285039642/",
							Location:             "London, England, United Kingdom",
							Date:                 mustDate("2025-09-10"),
							WithSalary:           false,
							Remote:               false,
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
			ShortDescription: "The upskilling platform for AI and tech adoption",
			Industries: []domain.Industry{
				domain.IndustryEdTech,
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Blue River Technology",
			Website: "https://www.bluerivertechnology.com/",
			Careers: "",
			About:   "",
			Blog:    "",
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
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Rust: {
					GitHubRepositoriesCount: 0,
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
			ShortDescription: "We empower our customers – farmers",
			Industries: []domain.Industry{
				domain.IndustryAgroTech,
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Ultra Intelligence & Communications",
			Website: "https://www.ultra-ic.com/",
			Careers: "https://www.ultra-ic.com/careers/",
			About:   "https://www.ultra-ic.com/about-us/",
			Blog:    "",
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
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
				domain.Rust: {
					GitHubRepositoriesCount: 0,
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
			ShortDescription: "Defense and Space Manufacturing",
			Industries: []domain.Industry{
				domain.IndustryDefenseTech,
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Cognyte",
			Website: "https://www.cognyte.com/",
			Careers: "https://www.cognyte.com/careers/",
			About:   "https://www.cognyte.com/about/",
			Blog:    "",
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
				Jobs:        "27",
				Reviews:     "314",
				Salaries:    "397",
				ReviewsRate: "4.2",
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
							Title:                "Rust Software Engineer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4298710314/",
							Location:             "Limassol, Cyprus",
							Date:                 mustDate("2025-09-10"),
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
			ShortDescription: "Data processing and investigative analytics solutions",
			Industries:       []domain.Industry{
				// NOP
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Catio",
			Website: "https://www.catio.tech/",
			Careers: "",
			About:   "https://www.catio.tech/about-us",
			Blog:    "",
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
					GitHubRepositoriesCount: 0,
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
			ShortDescription: "Copilot for Tech Architecture – equipping companies with a data-driven and AI-powered platform to excel with evaluating, planning and evolving their tech stacks",
			Industries:       []domain.Industry{
				// NOP
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "CloudNuro",
			Website: "https://www.cloudnuro.ai/",
			Careers: "https://www.cloudnuro.ai/careers",
			About:   "https://www.cloudnuro.ai/about",
			Blog:    "",
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
					GitHubRepositoriesCount: 0,
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
			ShortDescription: "Intelligent SaaS Management Platform",
			Industries:       []domain.Industry{
				// NOP
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Bitexen",
			Website: "https://www.bitexen.com/",
			Careers: "",
			About:   "",
			Blog:    "",
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
					GitHubRepositoriesCount: 0,
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
				// Cryptocurrency
			},
			Ignore: true, // Cryptocurrency
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Hostaway",
			Website: "https://www.hostaway.com/",
			Careers: "https://careers.hostaway.com/",
			About:   "https://www.hostaway.com/about/",
			Blog:    "",
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
				Jobs:        "136",
				Reviews:     "70",
				Salaries:    "63",
				ReviewsRate: "4.4",
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
							Title:                "Senior Backend Engineer, PHP/Go",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4298774762/",
							Location:             "Tottenham, England, United Kingdom",
							Date:                 mustDate("2025-09-10"),
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
			ShortDescription: "Vacation rental management software",
			Industries: []domain.Industry{
				domain.IndustryPropTech,
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Cirrus Logic",
			Website: "https://www.cirrus.com/",
			Careers: "https://www.cirrus.com/careers",
			About:   "https://www.cirrus.com/company",
			Blog:    "",
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
					GitHubRepositoriesCount: 0,
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
			ShortDescription: "Low-Power Audio and High-Performance Mixed-Signal Processing",
			Industries:       []domain.Industry{
				// NOP
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Surfshark",
			Website: "https://surfshark.com/",
			Careers: "",
			About:   "",
			Blog:    "",
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
					GitHubRepositoriesCount: 0,
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
			ShortDescription: "VPN",
			Industries: []domain.Industry{
				domain.IndustryCyberSecurity,
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Forte Group",
			Website: "https://fortegrp.com/",
			Careers: "https://fortegrp.com/careers",
			About:   "https://fortegrp.com/about-forte-group",
			Blog:    "",
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
				Reviews:     "163",
				Salaries:    "159",
				ReviewsRate: "4.1",
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
							Title:                "Senior C#/Golang Engineer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4296467085/",
							Location:             "Poland",
							Date:                 mustDate("2025-09-06"),
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
			ShortDescription: "Managed Engineering Solutions in Software Development and Data Engineering",
			Industries:       []domain.Industry{
				// NOP
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Thunes",
			Website: "https://www.thunes.com/",
			Careers: "",
			About:   "https://www.thunes.com/about-us/",
			Blog:    "",
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
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:                "Golang Software Engineer",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4156638094/",
							Location:             "Barcelona, Catalonia, Spain",
							Date:                 mustDate("2025-09-09"),
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
			Name:    "Kayrros",
			Website: "https://www.kayrros.com/",
			Careers: "https://www.kayrros.com/career/",
			About:   "https://www.kayrros.com/about-about/",
			Blog:    "",
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
					GitHubRepositoriesCount: 0,
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
			ShortDescription: "We use AI and geoanalytics to turn raw satellite data into actionable insights on energy, supply chains, physical risks, nature",
			Industries:       []domain.Industry{
				// NOP
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Fingerprint",
			Website: "https://fingerprint.com/",
			Careers: "https://fingerprint.com/careers/",
			About:   "https://fingerprint.com/about-us/",
			Blog:    "",
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
					GitHubRepositoriesCount: 1,
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
			ShortDescription: "Identify visitors you can trust",
			Industries: []domain.Industry{
				domain.IndustryCyberSecurity,
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Perion",
			Website: "https://perion.com/",
			Careers: "https://perion.com/careers/",
			About:   "https://perion.com/about/",
			Blog:    "",
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
					GitHubRepositoriesCount: 0,
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
			ShortDescription: "Perion is helping agencies, brands and retailers get better results with their marketing investments by providing advanced technology across all major digital channels",
			Industries: []domain.Industry{
				domain.IndustryAdTech,
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Upfluence",
			Website: "https://www.upfluence.com/",
			Careers: "https://www.upfluence.com/careers",
			About:   "https://www.upfluence.com/about",
			Blog:    "",
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
				Alias: "",
			},
			LevelsFyiProfile: domain.LevelsFyiProfile{
				Alias: "",
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
					GitHubRepositoriesCount: 30,
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
			ShortDescription: "Influencer marketing platform",
			Industries:       []domain.Industry{
				// NOP
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Capco",
			Website: "https://www.capco.com/",
			Careers: "https://www.capco.com/Careers/Job-Search",
			About:   "https://www.capco.com/About-Us",
			Blog:    "",
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
					Vacancies:               []domain.Vacancy{},
				},
				domain.Haskell: {
					GitHubRepositoriesCount: 0,
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
					GitHubRepositoriesCount: 0,
					Vacancies:               []domain.Vacancy{},
				},
			},
			ShortDescription: "Driving digital transformation in the financial services industry",
			Industries: []domain.Industry{
				domain.IndustryFinTech,
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "FacilitaPay",
			Website: "https://www.facilitapay.com/",
			Careers: "",
			About:   "https://www.facilitapay.com/about-us",
			Blog:    "",
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
			ShortDescription: "FacilitaPay transforms payments between developed and emerging markets",
			Industries: []domain.Industry{
				domain.IndustryFinTech,
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "McKesson",
			Website: "https://www.mckesson.com/",
			Careers: "https://careers.mckesson.com/",
			About:   "https://www.mckesson.com/about-us/",
			Blog:    "",
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
							Title:                "Software Engineer (JavaScript / Ruby / Elixir)",
							ShortDescription:     "",
							SwitchingOpportunity: "",
							URL:                  "https://www.linkedin.com/jobs/view/4299028566/",
							Location:             "Columbus, OH",
							Date:                 mustDate("2025-09-11"),
							WithSalary:           true, // $106.9k/yr - $178.1k/yr
							Remote:               false,
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
			ShortDescription: "Healthcare company",
			Industries: []domain.Industry{
				domain.IndustryHealthTech,
			},
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "",
			Website: "",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                0,
				Alias:             "",
				Name:              "",
				Followers:         "",
				Employees:         "",
				AssociatedMembers: "",
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
			Name:    "",
			Website: "",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                0,
				Alias:             "",
				Name:              "",
				Followers:         "",
				Employees:         "",
				AssociatedMembers: "",
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
			Name:    "",
			Website: "",
			Careers: "",
			About:   "",
			Blog:    "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                0,
				Alias:             "",
				Name:              "",
				Followers:         "",
				Employees:         "",
				AssociatedMembers: "",
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

		// Template
		//{
		//	ID:      0,  // system
		//	Type:    "", // system
		//	Name:    "",
		//	Website: "",
		//	Careers: "",
		//	About:   "",
		//	Blog:    "",
		//	LinkedInProfile: domain.LinkedInProfile{
		//		ID:                0,
		//		Alias:             "",
		//		Name:              "",
		//		Followers:         "",
		//		Employees:         "",
		//		AssociatedMembers: "",
		//		Verified:          false,
		//	},
		//	GitHubProfile: domain.GitHubProfile{
		//		Login:    "",
		//		Verified: false,
		//	},
		//	BlindProfile: domain.BlindProfile{
		//		Alias: "",
		//	},
		//	LevelsFyiProfile: domain.LevelsFyiProfile{
		//		Alias: "",
		//	},
		//	GlassdoorProfile: domain.GlassdoorProfile{
		//		OverviewURL: "",
		//	},
		//	IndeedProfile: domain.IndeedProfile{
		//		Alias: "",
		//	},
		//	Languages: domain.Languages{
		//		domain.Go: {
		//			GitHubRepositoriesCount: 0,
		//			Vacancies:               []domain.Vacancy{},
		//		},
		//		domain.Rust: {
		//			GitHubRepositoriesCount: 0,
		//			Vacancies:               []domain.Vacancy{},
		//		},
		//		domain.Zig: {
		//			GitHubRepositoriesCount: 0,
		//			Vacancies:               []domain.Vacancy{},
		//		},
		//		domain.Scala: {
		//			GitHubRepositoriesCount: 0,
		//			Vacancies:               []domain.Vacancy{},
		//		},
		//		domain.Elixir: {
		//			GitHubRepositoriesCount: 0,
		//			Vacancies:               []domain.Vacancy{},
		//		},
		//		domain.Clojure: {
		//			GitHubRepositoriesCount: 0,
		//			Vacancies:               []domain.Vacancy{},
		//		},
		//		domain.Haskell: {
		//			GitHubRepositoriesCount: 0,
		//			Vacancies:               []domain.Vacancy{},
		//		},
		//		domain.Erlang: {
		//			GitHubRepositoriesCount: 0,
		//			Vacancies:               []domain.Vacancy{},
		//		},
		//	},
		//	ShortDescription: "",
		//	Industries:       []domain.Industry{
		//		// NOP
		//	},
		//},
	}
}
