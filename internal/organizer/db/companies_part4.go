package db

import (
	"github.com/readytotouch/readytotouch/internal/domain"
)

func companiesPart4() []domain.CompanyProfile {
	return []domain.CompanyProfile{
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "limango GmbH — A member of the otto group",
			Website: "https://limango.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:       9381585,
				Alias:    "limango-gmbh",
				Name:     "limango GmbH — A member of the otto group",
				Verified: true,
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Senior Go developer",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4096830999/",
							Date:             mustDate("2024-12-13"),
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
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Varonis",
			Website: "https://www.varonis.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:       25403,
				Alias:    "varonis",
				Name:     "Varonis",
				Verified: true,
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Backend Developer (Python, GO)",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4093559031/",
							Date:             mustDate("2024-12-10"),
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
			Industries: []domain.Industry{
				domain.IndustryCyberSecurity,
			},
			HasEmployeesFromCountries: []domain.Country{},
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "WorldTech IT, LLC",
			Website: "https://wtit.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:       1811897,
				Alias:    "worldtechit",
				Name:     "WorldTech IT, LLC",
				Verified: true,
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Golang Software Engineer",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4101282630/",
							Date:             mustDate("2024-12-17"),
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
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Applied Research Solutions",
			Website: "https://www.appliedres.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:       6017665,
				Alias:    "applied-research-solutions",
				Name:     "Applied Research Solutions",
				Verified: true,
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 11,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Senior Go Developer",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4101198898/",
							Date:             mustDate("2024-12-16"),
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
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "WEX",
			Website: "https://www.wexinc.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:       11637,
				Alias:    "wexinc",
				Name:     "WEX",
				Verified: true,
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Specialist Java/Golang Developer",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4048143482/",
							Date:             mustDate("2024-12-14"),
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
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Comcast",
			Website: "https://corporate.comcast.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:       1703,
				Alias:    "comcast",
				Name:     "Comcast",
				Verified: true,
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 48,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Senior Golang Engineer",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4081171024/",
							Date:             mustDate("2024-12-14"),
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
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Semrush",
			Website: "https://www.semrush.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:       2821922,
				Alias:    "semrush",
				Name:     "Semrush",
				Verified: true,
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 2,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Senior Golang Developer",
							ShortDescription: "Maroon Team",
							URL:              "https://www.linkedin.com/jobs/view/4063404057/",
							Date:             mustDate("2024-12-13"),
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
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Pi Square Technologies",
			Website: "https://www.pisquaretech.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:       10589520,
				Alias:    "pi-square-technology",
				Name:     "Pi Square Technologies",
				Verified: true,
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Senior Golang Developer",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4098692466/",
							Date:             mustDate("2024-12-13"),
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
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "DataOn (PT Indodev Niaga Internet)",
			Website: "https://dataon.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:       844561,
				Alias:    "dataoncorp",
				Name:     "DataOn (PT Indodev Niaga Internet)",
				Verified: true,
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Senior Golang Developer",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4099666536/",
							Date:             mustDate("2024-12-13"),
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
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Kiteworks",
			Website: "https://www.kiteworks.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:       33337,
				Alias:    "kiteworkscgcp",
				Name:     "Kiteworks",
				Verified: true,
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Golang Developer",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4096373782/",
							Date:             mustDate("2024-12-11"),
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
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Centripetal",
			Website: "https://www.centripetal.ai/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:       1321511,
				Alias:    "centripetal-ai",
				Name:     "Centripetal",
				Verified: true,
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Senior Golang Software Engineer",
							ShortDescription: "CIS Data Services (IE)",
							URL:              "https://www.linkedin.com/jobs/view/4094769985/",
							Date:             mustDate("2024-12-10"),
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
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Technology Innovation Institute",
			Website: "https://www.tii.ae/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:       28729816,
				Alias:    "tiiuae",
				Name:     "Technology Innovation Institute",
				Verified: true,
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 16,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Senior Golang Developer",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4091567553/",
							Date:             mustDate("2024-12-10"),
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
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "BAUHAUS Deutschland",
			Website: "https://bauhaus.info/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:       9417424,
				Alias:    "bauhaus-deutschland",
				Name:     "BAUHAUS Deutschland",
				Verified: true,
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Senior Golang Engineer",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4049064454/",
							Date:             mustDate("2024-12-10"),
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
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Sun* Vietnam",
			Website: "https://sun-asterisk.vn/en/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:       14571208,
				Alias:    "sunasterisk",
				Name:     "Sun* Vietnam",
				Verified: true,
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 4,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Golang Engineer",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4090103893/",
							Date:             mustDate("2024-12-10"),
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
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Miracle Software Systems, Inc",
			Website: "https://www.miraclesoft.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:       15388,
				Alias:    "miraclesoft",
				Name:     "Miracle Software Systems, Inc",
				Verified: true,
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Golang Developer",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4090375762/",
							Date:             mustDate("2024-12-10"),
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
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Collabera",
			Website: "https://www.collaberadigital.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:       24440,
				Alias:    "collabera",
				Name:     "Collabera",
				Verified: true,
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Java Developer",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4087446589/",
							Date:             mustDate("2024-12-03"),
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
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Sparq",
			Website: "https://www.teamsparq.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:       90679425,
				Alias:    "teamsparq",
				Name:     "Sparq",
				Verified: true,
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Senior Golang Developer",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4049853199/",
							Date:             mustDate("2024-12-03"),
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
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "SMG Swiss Marketplace Group",
			Website: "https://swissmarketplace.group/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:       76994292,
				Alias:    "smg-marketplace",
				Name:     "SMG Swiss Marketplace Group",
				Verified: true,
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Senior Golang Engineer",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4016921964/",
							Date:             mustDate("2024-12-03"),
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
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Hornetsecurity",
			Website: "https://www.hornetsecurity.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:       750470,
				Alias:    "hornetsecurity",
				Name:     "Hornetsecurity",
				Verified: true,
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Golang Software Engineer",
							ShortDescription: "Security Lab",
							URL:              "https://www.linkedin.com/jobs/view/4089330917/",
							Date:             mustDate("2024-12-03"),
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
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Gorilla Logic",
			Website: "https://www.gorillalogic.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:       68692,
				Alias:    "gorillalogic",
				Name:     "Gorilla Logic",
				Verified: true,
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 3,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Senior Go Engineer",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4085414974/",
							Date:             mustDate("2024-11-26"),
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
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "ZeptoLab",
			Website: "https://www.zeptolab.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:       1945804,
				Alias:    "zeptolab",
				Name:     "ZeptoLab",
				Verified: true,
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Senior Golang Developer",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4102533913/",
							Date:             mustDate("2024-12-17"),
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
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Getir",
			Website: "https://getir.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:       9866841,
				Alias:    "getir",
				Name:     "Getir",
				Verified: true,
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Go Engineer",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4050960562/",
							Date:             mustDate("2024-12-17"),
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
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Happening",
			Website: "https://www.happening.xyz/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:       86471246,
				Alias:    "happen-ing",
				Name:     "Happening",
				Verified: true,
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Senior Backend Engineer (Go)",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4102013036/",
							Date:             mustDate("2024-12-16"),
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
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Zeller",
			Website: "https://www.myzeller.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:       12687202,
				Alias:    "zeller",
				Name:     "Zeller",
				Verified: true,
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Senior Engineer (GO/Typescript + AWS)",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4047579348/",
							Date:             mustDate("2024-12-13"),
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
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Flip",
			Website: "https://flip.id/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:       13234824,
				Alias:    "flip.id",
				Name:     "Flip",
				Verified: true,
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Software Engineer (Backend-GO) — Risk Platform",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4096854844/",
							Date:             mustDate("2024-12-13"),
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
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "OpenX",
			Website: "https://www.openx.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:       42890,
				Alias:    "openx",
				Name:     "OpenX",
				Verified: true,
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 4,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Senior Software Engineer IV (Java/GO)",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4062276145/",
							Date:             mustDate("2024-12-12"),
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
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Lenskart.com",
			Website: "https://lenskart.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:       213339,
				Alias:    "lenskart-com",
				Name:     "Lenskart.com",
				Verified: true,
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Software Engineer, Golang",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4096376928/",
							Date:             mustDate("2024-12-11"),
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
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Lowe's Companies, Inc.",
			Website: "https://lowes.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:       4128,
				Alias:    "lowe's-home-improvement",
				Name:     "Lowe's Companies, Inc.",
				Verified: true,
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Senior Software Engineer",
							ShortDescription: "Mobile/Backend (Android, Go, Kotlin)",
							URL:              "https://www.linkedin.com/jobs/view/4093511115/",
							Date:             mustDate("2024-12-11"),
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
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Huntress",
			Website: "https://www.huntress.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:       10172550,
				Alias:    "huntress-labs",
				Name:     "Huntress",
				Verified: true,
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 5,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Senior Software Engineer",
							ShortDescription: "Windows SIEM Agent (Go)",
							URL:              "https://www.linkedin.com/jobs/view/4055142455/",
							Date:             mustDate("2024-12-10"),
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
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "1NCE",
			Website: "https://1nce.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:       18471615,
				Alias:    "1nce",
				Name:     "1NCE",
				Verified: true,
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 4,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Senior Software Engineer Go",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4088250312/",
							Date:             mustDate("2024-12-03"),
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
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Waystar",
			Website: "https://www.waystar.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:       11444741,
				Alias:    "waystar",
				Name:     "Waystar",
				Verified: true,
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Advanced Application Engineer (PHP / GO)",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4089563703/",
							Date:             mustDate("2024-12-03"),
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
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "efood",
			Website: "https://www.e-food.gr/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:       2410959,
				Alias:    "e-food-gr",
				Name:     "efood",
				Verified: true,
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Senior Software Engineer (Go)",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/3873641135/",
							Date:             mustDate("2024-12-03"),
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
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "ESET",
			Website: "https://www.eset.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:       28967,
				Alias:    "eset",
				Name:     "ESET",
				Verified: true,
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 2,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Software Engineer",
							ShortDescription: "Cloud Security Applications (.NET Core or GO)",
							URL:              "https://www.linkedin.com/jobs/view/4084785751/",
							Date:             mustDate("2024-12-03"),
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
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Praxent",
			Website: "https://praxent.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:       1365857,
				Alias:    "praxent",
				Name:     "Praxent",
				Verified: true,
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:            "LATAM Senior Backend Engineer (Python or GO)",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4085449585/",
							Date:             mustDate("2024-12-26"),
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
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Stacklok",
			Website: "https://stacklok.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:       94101540,
				Alias:    "stacklok",
				Name:     "Stacklok",
				Verified: true,
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 5,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Software Engineer II (Go & Python)",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4079391812/",
							Date:             mustDate("2024-11-20"),
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
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "PT Lion Super Indo",
			Website: "https://www.superindo.co.id/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:       1104467,
				Alias:    "superindo",
				Name:     "PT Lion Super Indo",
				Verified: true,
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Backend Developer (Golang)",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4102186636/",
							Date:             mustDate("2024-12-17"),
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
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Ubiquiti Inc.",
			Website: "https://www.ui.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:       109341,
				Alias:    "ubiquiti-",
				Name:     "Ubiquiti Inc.",
				Verified: true,
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 4,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Backend Developer (Golang)",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4100868953/",
							Date:             mustDate("2024-12-16"),
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
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "TextNow",
			Website: "https://textnow.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:       529693,
				Alias:    "enflick-inc-",
				Name:     "TextNow",
				Verified: true,
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 18,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Golang Backend Developer",
							ShortDescription: "Trust & Safety",
							URL:              "https://www.linkedin.com/jobs/view/4096314146/",
							Date:             mustDate("2024-12-11"),
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
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Altenar",
			Website: "https://altenar.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:       11063103,
				Alias:    "altenar",
				Name:     "Altenar",
				Verified: true,
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Senior Golang Software Developer",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4006862895/",
							Date:             mustDate("2024-12-10"),
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
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "VUS — The English Center",
			Website: "https://vus.edu.vn/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:       33233160,
				Alias:    "vustheenglishcenter",
				Name:     "VUS — The English Center",
				Verified: true,
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Senior Backend Developer (Golang)",
							ShortDescription: "3 years exp",
							URL:              "https://www.linkedin.com/jobs/view/4089423781/",
							Date:             mustDate("2024-12-10"),
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
		},
		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Exness",
			Website: "https://www.exness.uk/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:       3291356,
				Alias:    "exness",
				Name:     "Exness",
				Verified: true,
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Golang/Python Developer",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4088177442/",
							Date:             mustDate("2024-12-03"),
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
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Xsolla",
			Website: "",
			LinkedInProfile: domain.LinkedInProfile{
				ID:       875431,
				Alias:    "xsolla",
				Name:     "Xsolla",
				Verified: true,
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 3,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Golang Backend Developer",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4069628874/",
							Date:             mustDate("2024-12-03"),
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
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "emerchantpay",
			Website: "https://www.emerchantpay.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:       1432812,
				Alias:    "emerchantpay-ltd",
				Name:     "emerchantpay",
				Verified: true,
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Senior Developer with Golang",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4084138762/",
							Date:             mustDate("2024-12-03"),
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
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Power Factors",
			Website: "https://www.powerfactors.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:       4748453,
				Alias:    "power-factors-inc",
				Name:     "Power Factors",
				Verified: true,
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Senior Engineer, Back-End, Golang (Greece)",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4102953810/",
							Date:             mustDate("2024-12-17"),
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
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Airspace",
			Website: "https://www.airspace.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:       10504949,
				Alias:    "airspace-inc",
				Name:     "Airspace",
				Verified: true,
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Senior Software Engineer — Golang",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4101213697/",
							Date:             mustDate("2024-12-16"),
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
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Amartha",
			Website: "https://amartha.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:       3655032,
				Alias:    "amartha",
				Name:     "Amartha",
				Verified: true,
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Backend Engineer (Golang) — Senior/Principal Level",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4102510260/",
							Date:             mustDate("2024-12-16"),
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
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Creative Software",
			Website: "http://www.creativesoftware.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:       12688,
				Alias:    "creativesoftware",
				Name:     "Creative Software",
				Verified: true,
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Senior Software Engineer(Golang)",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4101950317/",
							Date:             mustDate("2024-12-16"),
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
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Cisco ThousandEyes",
			Website: "https://www.thousandeyes.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:       687352,
				Alias:    "thousandeyes",
				Name:     "Cisco ThousandEyes",
				Verified: true,
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 9,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Senior Site Reliability Engineer I (Python/Golang), Agent Ops",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4100798671/",
							Date:             mustDate("2024-12-14"),
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
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "U.S. Bank",
			Website: "https://usbank.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:       2532,
				Alias:    "us-bank",
				Name:     "U.S. Bank",
				Verified: true,
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Software Information Security Engineer (Golang or Python)",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4063303097/",
							Date:             mustDate("2024-12-12"),
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
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "NetApp",
			Website: "https://www.netapp.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:       2105,
				Alias:    "netapp",
				Name:     "NetApp",
				Verified: true,
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 23,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Senior Software Engineer — Golang",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4095534140/",
							Date:             mustDate("2024-12-12"),
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
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Trend Micro",
			Website: "https://www.trendmicro.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:       4312,
				Alias:    "trend-micro",
				Name:     "Trend Micro",
				Verified: true,
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 4,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Cloud Engineer (Golang/Python, Backend Focus)",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4063567336/",
							Date:             mustDate("2024-12-11"),
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
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Central Retail",
			Website: "https://www.centralretail.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:       26628492,
				Alias:    "centralretail",
				Name:     "Central Retail",
				Verified: true,
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Senior Software Engineer (Golang)",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4093163873/",
							Date:             mustDate("2024-12-10"),
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
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Kargo",
			Website: "https://www.kargo.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:       55830,
				Alias:    "kargo",
				Name:     "Kargo",
				Verified: true,
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Senior Software Engineer (Golang)",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4089015041/",
							Date:             mustDate("2024-12-10"),
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
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "Salesloft",
			Website: "https://www.salesloft.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:       2296178,
				Alias:    "salesloft",
				Name:     "Salesloft",
				Verified: true,
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 3,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Software Engineer (Golang)",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4072522353/",
							Date:             mustDate("2024-12-03"),
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
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "CJ MORE",
			Website: "https://www.cjexpress.co.th/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:       10081692,
				Alias:    "cj-express-group",
				Name:     "CJ MORE",
				Verified: true,
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Software Engineer (Golang)",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4088241281/",
							Date:             mustDate("2024-12-03"),
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
		},

		{
			ID:      0,  // system
			Type:    "", // system
			Name:    "H-E-B",
			Website: "https://www.heb.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:       164159,
				Alias:    "heb",
				Name:     "H-E-B",
				Verified: true,
			},
			Languages: domain.Languages{
				domain.Go: {
					GitHubRepositoriesCount: 0,
					Vacancies: []domain.Vacancy{
						{
							Title:            "Sr Software Engineer, Golang",
							ShortDescription: "",
							URL:              "https://www.linkedin.com/jobs/view/4000104862/",
							Date:             mustDate("2024-12-03"),
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
		},
		// Template short
		//{
		//	ID:      0,  // system
		//	Type:    "", // system
		//	Name:    "",
		//	Website: "",
		//	LinkedInProfile: domain.LinkedInProfile{
		//		ID:       0,
		//		Alias:    "",
		//		Name:     "",
		//		Verified: false,
		//	},
		//	Languages: domain.Languages{
		//		domain.Go: {
		//			GitHubRepositoriesCount: 0,
		//			Vacancies: []domain.Vacancy{
		//				{
		//					Title:            "",
		//					ShortDescription: "",
		//					URL:              "",
		//					Date:             mustDate(""),
		//				},
		//			},
		//		},
		//		domain.Rust:    {},
		//		domain.Zig:     {},
		//		domain.Scala:   {},
		//		domain.Elixir:  {},
		//		domain.Clojure: {},
		//		domain.Haskell: {},
		//	},
		//},
	}
}
