package db

import "github.com/readytotouch/readytotouch/internal/domain"

func companiesRecruitingAgencies() []domain.CompanyProfile {
	return []domain.CompanyProfile{
		{
			Name:    "La Fosse",
			Website: "https://www.lafosse.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                103069,
				IDs:               nil,
				Alias:             "la-fosse",
				Name:              "La Fosse",
				Followers:         "336K",
				Employees:         "201-500",
				AssociatedMembers: "626",
				Verified:          true,
			},
			Ignore: true, // Recruiting agency
		},
		{
			Name:    "Zeektek",
			Website: "https://www.zeektek.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                15205156,
				IDs:               nil,
				Alias:             "zeektek",
				Name:              "Zeektek",
				Followers:         "184K",
				Employees:         "11-50",
				AssociatedMembers: "55",
				Verified:          false,
			},
			Ignore: true, // Recruiting agency
		},
		{
			Name:    "Propel",
			Website: "https://www.propel-together.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                41226,
				IDs:               nil,
				Alias:             "propel-together",
				Name:              "Propel",
				Followers:         "286K",
				Employees:         "11-50",
				AssociatedMembers: "24",
				Verified:          true,
			},
			Ignore: true, // Recruiting agency
		},
		{
			Name:    "Undelucram.ro",
			Website: "https://www.undelucram.ro/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                18046146,
				IDs:               nil,
				Alias:             "undelucram.ro",
				Name:              "Undelucram.ro",
				Followers:         "89K",
				Employees:         "11-50",
				AssociatedMembers: "20",
				Verified:          true,
			},
			Ignore: true, // Recruiting agency
		},
		{
			Name:    "Adecco",
			Website: "https://www.adecco.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                1104359,
				IDs:               []int{598768, 1104359},
				Alias:             "adecco",
				Name:              "Adecco",
				Followers:         "9M",
				Employees:         "10K+",
				AssociatedMembers: "161,507",
				Verified:          true,
			},
			Ignore: true, // Recruiting agency
		},
		{
			Name:    "Fruition Group",
			Website: "https://www.fruitiongroup.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                1403805,
				IDs:               nil,
				Alias:             "fruition-it",
				Name:              "Fruition Group",
				Followers:         "151K",
				Employees:         "11-50",
				AssociatedMembers: "66",
				Verified:          false,
			},
			Ignore: true, // Recruiting agency
		},
		{
			Name:    "micro1",
			Website: "https://www.micro1.ai/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                81490090,
				IDs:               nil,
				Alias:             "micro1",
				Name:              "micro1",
				Followers:         "352K",
				Employees:         "51-200",
				AssociatedMembers: "791",
				Verified:          false,
			},
			Ignore: true, // Recruiting agency
		},
		{
			Name:    "PrincePerelson and Associates",
			Website: "https://perelson.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                70603,
				IDs:               nil,
				Alias:             "prince-perelson-and-associates",
				Name:              "PrincePerelson and Associates",
				Followers:         "60K",
				Employees:         "11-50",
				AssociatedMembers: "61",
				Verified:          false,
			},
			Ignore: true, // Recruiting agency
		},
		{
			Name:    "Fynity",
			Website: "https://www.fynitytalent.io/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                963360,
				IDs:               nil,
				Alias:             "fynity-png",
				Name:              "Fynity",
				Followers:         "75K",
				Employees:         "11-50",
				AssociatedMembers: "20",
				Verified:          false,
			},
			Ignore: true, // Recruiting agency
		},
		{
			Name:    "TieTalent",
			Website: "https://tietalent.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                11010661,
				IDs:               nil,
				Alias:             "tietalent",
				Name:              "TieTalent",
				Followers:         "49K",
				Employees:         "11-50",
				AssociatedMembers: "20",
				Verified:          true,
			},
			Ignore: true, // Recruiting agency
		},
		{
			Name:    "Akraya, Inc.",
			Website: "https://www.akraya.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                41771,
				IDs:               nil,
				Alias:             "akraya-inc",
				Name:              "Akraya, Inc.",
				Followers:         "173K",
				Employees:         "201-500",
				AssociatedMembers: "325",
				Verified:          true,
			},
			Ignore: true, // Recruiting agency
		},
		{
			Name:    "Burns Sheehan",
			Website: "https://www.burnssheehan.co.uk/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                103683,
				IDs:               nil,
				Alias:             "jcl-burns-sheehan",
				Name:              "Burns Sheehan",
				Followers:         "233K",
				Employees:         "11-50",
				AssociatedMembers: "46",
				Verified:          false,
			},
			Ignore: true, // Recruiting agency
		},
		{
			Name:    "Radley James",
			Website: "https://radleyjames.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                1116353,
				IDs:               nil,
				Alias:             "radley-james",
				Name:              "Radley James",
				Followers:         "308K",
				Employees:         "51-200",
				AssociatedMembers: "38",
				Verified:          false,
			},
			Ignore: true, // Recruiting agency
		},
	}
}
