package db

import "github.com/readytotouch/readytotouch/internal/domain"

func companiesOutsourceCompanies() []domain.CompanyProfile {
	return []domain.CompanyProfile{
		{
			Type:    domain.CompanyTypeOutsource,
			Name:    "Minsait",
			Website: "https://www.minsait.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                10963,
				IDs:               nil,
				Alias:             "minsait",
				Name:              "Minsait",
				Followers:         "599K",
				Employees:         "10K+",
				AssociatedMembers: "15,668",
				Verified:          true,
			},
			Ignore: true, // Outsource
		},
		{
			Type:    domain.CompanyTypeOutsource,
			Name:    "Virtusa",
			Website: "https://www.virtusa.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                5140,
				IDs:               nil,
				Alias:             "virtusa",
				Name:              "Virtusa",
				Followers:         "1M",
				Employees:         "10K+",
				AssociatedMembers: "17,503",
				Verified:          true,
			},
			Ignore: true, // Outsource
		},
		{
			Type:    domain.CompanyTypeOutsource,
			Name:    "Qubika",
			Website: "https://qubika.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                201026,
				IDs:               nil,
				Alias:             "qubika",
				Name:              "Qubika",
				Followers:         "58K",
				Employees:         "501-1K",
				AssociatedMembers: "726",
				Verified:          true,
			},
			Ignore: true, // Outsource
		},
		{
			Type:    domain.CompanyTypeOutsource,
			Name:    "Zartis",
			Website: "https://www.zartis.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                2274396,
				IDs:               nil,
				Alias:             "zartis",
				Name:              "Zartis",
				Followers:         "48K",
				Employees:         "201-500",
				AssociatedMembers: "216",
				Verified:          true,
			},
			Ignore: true, // Outsource
		},
		{
			Type:    domain.CompanyTypeOutsource,
			Name:    "Smart Consulting",
			Website: "https://www.smartconsulting.pt/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                9356133,
				IDs:               []int{9356133, 79396476},
				Alias:             "thesmartwaytodoit",
				Name:              "Smart Consulting",
				Followers:         "121K",
				Employees:         "201-500",
				AssociatedMembers: "322",
				Verified:          false,
			},
			Ignore: true, // Outsource
		},
	}
}
