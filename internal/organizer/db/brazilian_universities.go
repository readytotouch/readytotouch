package db

import "github.com/readytotouch/readytotouch/internal/domain"

func BrazilianUniversities() []domain.University {
	return []domain.University{
		{
			LinkedInProfile: domain.LinkedInProfile{
				ID:       239895,
				Alias:    "uspoficial",
				Name:     "USP — Universidade de São Paulo",
				Verified: true,
			},
		},
		{
			LinkedInProfile: domain.LinkedInProfile{
				ID:       986104,
				Alias:    "universidade-estadual-de-campinas",
				Name:     "Universidade Estadual de Campinas",
				Verified: true,
			},
		},
		{
			LinkedInProfile: domain.LinkedInProfile{
				ID:       15171,
				Alias:    "ufmg",
				Name:     "Universidade Federal de Minas Gerais",
				Verified: true,
			},
		},
		{
			LinkedInProfile: domain.LinkedInProfile{
				ID:       69715404,
				Alias:    "ufrjoficial",
				Name:     "Federal University of Rio de Janeiro",
				Verified: true,
			},
		},
		{
			LinkedInProfile: domain.LinkedInProfile{
				ID:       10866,
				Alias:    "ufrgs",
				Name:     "Federal University of Rio Grande do Sul",
				Verified: true,
			},
		},
		{
			LinkedInProfile: domain.LinkedInProfile{
				ID:       28514,
				Alias:    "ufpe",
				Name:     "Universidade Federal de Pernambuco",
				Verified: true,
			},
		},
		{
			LinkedInProfile: domain.LinkedInProfile{
				ID:       17959,
				Alias:    "ufsc",
				Name:     "Universidade Federal de Santa Catarina",
				Verified: true,
			},
		},
		{
			LinkedInProfile: domain.LinkedInProfile{
				ID:       312647,
				Alias:    "universidade-de-bras-lia",
				Name:     "Universidade de Brasília",
				Verified: true,
			},
		},
		{
			LinkedInProfile: domain.LinkedInProfile{
				ID:       1379596,
				Alias:    "unesp---universidade-estadual-paulista-j-lio-de-mesquita-filho-",
				Name:     "Universidade Estadual Paulista Júlio de Mesquita Filho",
				Verified: true,
			},
		},
		{
			LinkedInProfile: domain.LinkedInProfile{
				ID:       38307,
				Alias:    "ufpr",
				Name:     "Universidade Federal do Paraná",
				Verified: true,
			},
		},
		{
			LinkedInProfile: domain.LinkedInProfile{
				ID:       760298,
				Alias:    "ufcinforma",
				Name:     "Federal University of Ceara",
				Verified: true,
			},
		},

		//{
		//	LinkedInProfile: domain.LinkedInProfile{
		//		ID:       0,
		//		Alias:    "",
		//		Name:     "",
		//		Verified: false,
		//	},
		//},
	}
}
