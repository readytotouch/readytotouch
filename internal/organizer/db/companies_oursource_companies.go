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
		{
			Type:    domain.CompanyTypeOutsource,
			Name:    "Xoriant",
			Website: "https://www.xoriant.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                166996,
				IDs:               nil,
				Alias:             "xoriant",
				Name:              "Xoriant",
				Followers:         "464K",
				Employees:         "5K-10K",
				AssociatedMembers: "4,548",
				Verified:          true,
			},
			Ignore: true, // Outsource
		},
		{
			Type:    domain.CompanyTypeOutsource,
			Name:    "Trideca",
			Website: "https://www.trideca.com.au/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                2455945,
				IDs:               nil,
				Alias:             "tridecaptyltd",
				Name:              "Trideca",
				Followers:         "7K",
				Employees:         "51-200",
				AssociatedMembers: "97",
				Verified:          true,
			},
			Ignore: true, // Outsource
		},
		{
			Type:    domain.CompanyTypeOutsource,
			Name:    "Excelia",
			Website: "https://excelia.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                735796,
				IDs:               nil,
				Alias:             "excelia",
				Name:              "Excelia",
				Followers:         "109K",
				Employees:         "201-500",
				AssociatedMembers: "325",
				Verified:          false,
			},
			Ignore: true, // Outsource
		},
		{
			Type:    domain.CompanyTypeOutsource,
			Name:    "Xebia",
			Website: "https://xebia.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                16903,
				IDs:               []int{16903, 72738535},
				Alias:             "xebia",
				Name:              "Xebia",
				Followers:         "378K",
				Employees:         "5K-10K",
				AssociatedMembers: "5,399",
				Verified:          true,
			},
			Ignore: true, // Outsource
		},
		{
			Type:    domain.CompanyTypeOutsource,
			Name:    "Coltech",
			Website: "https://www.coltech.io/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                69762147,
				IDs:               nil,
				Alias:             "coltechglobal",
				Name:              "Coltech",
				Followers:         "57K",
				Employees:         "11-50",
				AssociatedMembers: "28",
				Verified:          false,
			},
			Ignore: true, // Outsource
		},
		{
			Type:    domain.CompanyTypeOutsource,
			Name:    "Akkodis",
			Website: "https://www.akkodis.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                79383535,
				IDs:               []int{224006, 320456, 79383535},
				Alias:             "akkodis",
				Name:              "Akkodis",
				Followers:         "3M",
				Employees:         "10K+",
				AssociatedMembers: "24,836",
				Verified:          true,
			},
			Ignore: true, // Outsource
		},
	}
}
