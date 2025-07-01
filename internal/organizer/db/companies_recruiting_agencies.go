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
		{
			Name:    "Spectrum IT Recruitment",
			Website: "https://www.spectrumit.co.uk/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                206206,
				IDs:               nil,
				Alias:             "spectrum-it-recruitment",
				Name:              "Spectrum IT Recruitment",
				Followers:         "149K",
				Employees:         "11-50",
				AssociatedMembers: "41",
				Verified:          false,
			},
			Ignore: true, // Recruiting agency
		},
		{
			Name:    "RJC Group",
			Website: "https://www.rjcgroup.co.uk/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                10333877,
				IDs:               nil,
				Alias:             "rjc-group-ltd",
				Name:              "RJC Group",
				Followers:         "41K",
				Employees:         "11-50",
				AssociatedMembers: "29",
				Verified:          false,
			},
			Ignore: true, // Recruiting agency
		},
		{
			Name:    "Identify Solutions",
			Website: "https://identifysolutions.co.uk/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                19023795,
				IDs:               nil,
				Alias:             "identify-solutions-ltd",
				Name:              "Identify Solutions",
				Followers:         "44K",
				Employees:         "11-50",
				AssociatedMembers: "8",
				Verified:          false,
			},
			Ignore: true, // Recruiting agency
		},
		{
			Name:    "IDR, Inc.",
			Website: "https://www.idr-inc.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                259589,
				IDs:               nil,
				Alias:             "idrinc",
				Name:              "IDR, Inc.",
				Followers:         "527K",
				Employees:         "501-1K",
				AssociatedMembers: "514",
				Verified:          true,
			},
			Ignore: true, // Recruiting agency
		},
		{
			Name:    "Phaxis",
			Website: "https://phaxis.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                106083,
				IDs:               nil,
				Alias:             "phaxis-consulting",
				Name:              "Phaxis",
				Followers:         "273K",
				Employees:         "51-200",
				AssociatedMembers: "201",
				Verified:          true,
			},
			Ignore: true, // Recruiting agency
		},
		{
			Name:    "Liberty Personnel Services, Inc.",
			Website: "https://libertyjobs.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                40956,
				IDs:               nil,
				Alias:             "liberty-personnel-services-inc",
				Name:              "Liberty Personnel Services, Inc.",
				Followers:         "289K",
				Employees:         "51-200",
				AssociatedMembers: "71",
				Verified:          true,
			},
			Ignore: true, // Recruiting agency
		},
		{
			Name:    "Robert Half",
			Website: "https://www.roberthalf.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                1681,
				IDs:               nil,
				Alias:             "robert-half-international",
				Name:              "Robert Half",
				Followers:         "4M",
				Employees:         "10K+",
				AssociatedMembers: "27,978",
				Verified:          true,
			},
			Ignore: true, // Recruiting agency
		},
		{
			Name:    "Cpl",
			Website: "https://www.cpl.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                163237,
				IDs:               []int{163237, 333317, 486582, 830729, 868838, 1134057, 14831142, 15251866, 64752574, 69269271, 87085983, 94292973},
				Alias:             "cpl",
				Name:              "Cpl",
				Followers:         "382K",
				Employees:         "10K+",
				AssociatedMembers: "4,654",
				Verified:          true,
			},
			Ignore: true, // Recruiting agency
		},
		{
			Name:    "HireArt",
			Website: "https://www.hireart.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                2506549,
				IDs:               nil,
				Alias:             "hireart",
				Name:              "HireArt",
				Followers:         "20K",
				Employees:         "51-200",
				AssociatedMembers: "286",
				Verified:          true,
			},
			Ignore: true, // https://www.linkedin.com/jobs/view/4211617362/ Rust robotaxi
		},
		{
			Name:    "Solomon Page",
			Website: "https://www.solomonpage.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                14352,
				IDs:               nil,
				Alias:             "solomon-page-group",
				Name:              "Solomon Page",
				Followers:         "641K",
				Employees:         "1K-5K",
				AssociatedMembers: "760",
				Verified:          true,
			},
			Ignore: true, // Recruiting agency
		},
		{
			Name:    "Seer",
			Website: "https://www.seerrecruit.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                26560101,
				IDs:               nil,
				Alias:             "seerrecruit",
				Name:              "Seer",
				Followers:         "37K",
				Employees:         "11-50",
				AssociatedMembers: "24",
				Verified:          true,
			},
			Ignore: true, // Recruiting agency
		},
		{
			Name:    "Robert Walters",
			Website: "https://www.robertwalters.co.uk/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                165757,
				IDs:               nil,
				Alias:             "robert-walters",
				Name:              "Robert Walters",
				Followers:         "2M",
				Employees:         "1K-5K",
				AssociatedMembers: "",
				Verified:          true,
			},
			Ignore: true, // Recruiting agency
		},
		{
			Name:    "Client Server",
			Website: "https://www.client-server.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                99339,
				IDs:               nil,
				Alias:             "client-server",
				Name:              "Client Server",
				Followers:         "159K",
				Employees:         "51-200",
				AssociatedMembers: "114",
				Verified:          true,
			},
			Ignore: true, // Recruiting agency
		},
		{
			Name:    "RED Global",
			Website: "https://www.redglobal.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                15014,
				IDs:               nil,
				Alias:             "red-commerce",
				Name:              "RED Global",
				Followers:         "566K",
				Employees:         "201-500",
				AssociatedMembers: "669",
				Verified:          true,
			},
			Ignore: true, // Recruiting agency
		},
		{
			Name:    "Signify Technology",
			Website: "https://www.signifytechnology.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                15087386,
				IDs:               nil,
				Alias:             "signify-technology",
				Name:              "Signify Technology",
				Followers:         "302K",
				Employees:         "11-50",
				AssociatedMembers: "61",
				Verified:          true,
			},
			Ignore: true, // Recruiting agency
		},
		{
			Name:    "Finna Group",
			Website: "https://www.finnagroup.io/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                67969725,
				IDs:               nil,
				Alias:             "finna-group",
				Name:              "Finna Group",
				Followers:         "3K",
				Employees:         "2-10",
				AssociatedMembers: "6",
				Verified:          true,
			},
			Ignore: true, // Recruiting agency
		},
		{
			Name:    "Lensa",
			Website: "https://lensa.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                5192530,
				IDs:               nil,
				Alias:             "lensa",
				Name:              "Lensa",
				Followers:         "34K",
				Employees:         "51-200",
				AssociatedMembers: "135",
				Verified:          true,
			},
			Ignore: true, // Recruiting agency
		},
		{
			Name:    "W3Global",
			Website: "https://www.w3global.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                1286891,
				IDs:               nil,
				Alias:             "w3global",
				Name:              "W3Global",
				Followers:         "106K",
				Employees:         "501-1K",
				AssociatedMembers: "1,010",
				Verified:          true,
			},
			Ignore: true, // Recruiting agency
		},
		{
			Name:    "Uplers",
			Website: "https://www.uplers.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                51692521,
				IDs:               nil,
				Alias:             "weareuplers",
				Name:              "Uplers",
				Followers:         "1M",
				Employees:         "1K-5K",
				AssociatedMembers: "1,068",
				Verified:          true,
			},
			Ignore: true, // Freelance Platforms
		},
		{
			Name:    "Coopers Group AG",
			Website: "https://www.coopers.ch/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                2282474,
				IDs:               nil,
				Alias:             "coopers-group-gmbh",
				Name:              "Coopers Group AG",
				Followers:         "61K",
				Employees:         "51-200",
				AssociatedMembers: "71",
				Verified:          false,
			},
			Ignore: true, // Recruiting agency
		},
		{
			Name:    "UMATR",
			Website: "https://www.umatr.io/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                70926319,
				IDs:               nil,
				Alias:             "umatr",
				Name:              "UMATR",
				Followers:         "110K",
				Employees:         "11-50",
				AssociatedMembers: "18",
				Verified:          false,
			},
			Ignore: true, // Recruiting agency
		},
		{
			Name:    "Edison Smart®",
			Website: "https://edisonsmart.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                76357488,
				IDs:               nil,
				Alias:             "edisonsmart",
				Name:              "Edison Smart®",
				Followers:         "135K",
				Employees:         "11-50",
				AssociatedMembers: "47",
				Verified:          false,
			},
			Ignore: true, // Recruiting agency
		},
		{
			Name:    "Ambition",
			Website: "https://www.ambition.com.au/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                3842697,
				IDs:               nil,
				Alias:             "ambitionrecruitment",
				Name:              "Ambition",
				Followers:         "240K",
				Employees:         "201-500",
				AssociatedMembers: "655",
				Verified:          true,
			},
			Ignore: true, // Recruiting agency
		},
		{
			Name:    "Michael Page",
			Website: "https://www.michaelpage.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                3476,
				IDs:               nil,
				Alias:             "michael-page",
				Name:              "Michael Page",
				Followers:         "4M",
				Employees:         "5K-10K",
				AssociatedMembers: "11,527",
				Verified:          true,
			},
			Ignore: true, // Recruiting agency
		},
		{
			Name:    "Morgan McKinley",
			Website: "https://www.morganmckinley.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                16175,
				IDs:               nil,
				Alias:             "morgan-mckinley",
				Name:              "Morgan McKinley",
				Followers:         "826K",
				Employees:         "1K-5K",
				AssociatedMembers: "1,359",
				Verified:          true,
			},
			Ignore: true, // Recruiting agency
		},
		{
			Name:    "Spectrum Search",
			Website: "https://www.spectrum-search.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                12581649,
				IDs:               nil,
				Alias:             "spectrumsearch",
				Name:              "Spectrum Search",
				Followers:         "132K",
				Employees:         "11-50",
				AssociatedMembers: "11",
				Verified:          false,
			},
			Ignore: true, // Recruiting agency
		},
		{
			Name:    "LHH",
			Website: "https://www.lhh.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                5235,
				IDs:               nil,
				Alias:             "lhhworldwide",
				Name:              "LHH",
				Followers:         "3M",
				Employees:         "10K+",
				AssociatedMembers: "18,119",
				Verified:          true,
			},
			Ignore: true, // Recruiting agency
		},
		{
			Name:    "Swisslinx",
			Website: "https://www.swisslinx.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                50444,
				IDs:               nil,
				Alias:             "swisslinx-ag",
				Name:              "Swisslinx",
				Followers:         "68K",
				Employees:         "51-200",
				AssociatedMembers: "60",
				Verified:          false,
			},
			Ignore: true, // Recruiting agency
		},
		{
			Name:    "Randstad Switzerland",
			Website: "https://www.randstad.ch/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                2127592,
				IDs:               nil,
				Alias:             "randstad-switzerland",
				Name:              "Randstad Switzerland",
				Followers:         "106K",
				Employees:         "201-500",
				AssociatedMembers: "285",
				Verified:          true,
			},
			Ignore: true, // Recruiting agency
		},
		{
			Name:    "Durlston Partners",
			Website: "https://durlstonpartners.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                1113019,
				IDs:               nil,
				Alias:             "durlston-partners",
				Name:              "Durlston Partners",
				Followers:         "140K",
				Employees:         "11-50",
				AssociatedMembers: "27",
				Verified:          true,
			},
			Ignore: true, // Recruiting agency
		},
		{
			Name:    "Anson McCade",
			Website: "https://www.ansonmccade.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                31392,
				IDs:               nil,
				Alias:             "anson-mccade",
				Name:              "Anson McCade",
				Followers:         "276K",
				Employees:         "11-50",
				AssociatedMembers: "125",
				Verified:          true,
			},
			Ignore: true, // Recruiting agency
		},
		{
			Name:    "Soni",
			Website: "https://www.soniresources.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                10599076,
				IDs:               nil,
				Alias:             "soni-resources-group",
				Name:              "Soni",
				Followers:         "101K",
				Employees:         "201-500",
				AssociatedMembers: "165",
				Verified:          true,
			},
			Ignore: true, // Recruiting agency
		},
		{
			Name:    "Oliver James",
			Website: "https://www.oliverjames.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                212201,
				IDs:               nil,
				Alias:             "oliver-james-",
				Name:              "Oliver James",
				Followers:         "133K",
				Employees:         "501-1K",
				AssociatedMembers: "875",
				Verified:          true,
			},
			Ignore: true, // Recruiting agency
		},
		{
			Name:    "Concero",
			Website: "https://concero.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                2699082,
				IDs:               nil,
				Alias:             "get-concero",
				Name:              "Concero",
				Followers:         "118K",
				Employees:         "51-200",
				AssociatedMembers: "85",
				Verified:          true,
			},
			Ignore: true, // Recruiting agency
		},
		{
			Name:    "OFS",
			Website: "https://weareofs.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                103150231,
				IDs:               nil,
				Alias:             "orbis-group-financial-services",
				Name:              "OFS",
				Followers:         "17K",
				Employees:         "11-50",
				AssociatedMembers: "11",
				Verified:          false,
			},
			Ignore: true, // Recruiting agency
		},
		{
			Name:    "Oliver Bernard",
			Website: "https://www.oliverbernard.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                798377,
				IDs:               nil,
				Alias:             "oliverbernard",
				Name:              "Oliver Bernard",
				Followers:         "432K",
				Employees:         "11-50",
				AssociatedMembers: "33",
				Verified:          true,
			},
			Ignore: true, // Recruiting agency
		},
		{
			Name:    "Talent Groups",
			Website: "https://www.talentgroups.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                51701268,
				IDs:               nil,
				Alias:             "talentgroups-connections",
				Name:              "Talent Groups",
				Followers:         "432K",
				Employees:         "501-1K",
				AssociatedMembers: "514",
				Verified:          false,
			},
			Ignore: true, // Recruiting agency
		},
		{
			Name:    "Workana",
			Website: "https://www.workana.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                2524324,
				IDs:               nil,
				Alias:             "workana",
				Name:              "Workana",
				Followers:         "128K",
				Employees:         "51-200",
				AssociatedMembers: "4,883",
				Verified:          true,
			},
			Ignore: true, // Recruiting agency
		},
		{
			Name:    "Red Oak Technologies",
			Website: "https://www.redoaktech.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                47475,
				IDs:               nil,
				Alias:             "red-oak-technologies",
				Name:              "Red Oak Technologies",
				Followers:         "129K",
				Employees:         "501-1K",
				AssociatedMembers: "220",
				Verified:          true,
			},
			Ignore: true, // Recruiting agency
		},
		{
			Name:    "iXceed Solutions",
			Website: "https://ixceed-solutions.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                1626442,
				IDs:               nil,
				Alias:             "ixceed-solutions",
				Name:              "iXceed Solutions",
				Followers:         "127K",
				Employees:         "201-500",
				AssociatedMembers: "121",
				Verified:          true,
			},
			Ignore: true, // Recruiting agency
		},
		{
			Name:    "Vallum Associates",
			Website: "https://vallumassociates.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                10348811,
				IDs:               nil,
				Alias:             "vallum-associates-limited",
				Name:              "Vallum Associates",
				Followers:         "196K",
				Employees:         "11-50",
				AssociatedMembers: "63",
				Verified:          true,
			},
			Ignore: true, // Recruiting agency
		},
		{
			Name:    "Russell Tobin",
			Website: "https://russelltobin.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                827183,
				IDs:               nil,
				Alias:             "827183", // "russell-tobin-&-associates-llc",
				Name:              "Russell Tobin",
				Followers:         "1M",
				Employees:         "201-500",
				AssociatedMembers: "816",
				Verified:          true,
			},
			Ignore: true, // Recruiting agency
		},
		{
			Name:    "Gazelle Global",
			Website: "https://www.gazellegc.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                9312644,
				IDs:               nil,
				Alias:             "gazelle-global-consulting",
				Name:              "Gazelle Global",
				Followers:         "210K",
				Employees:         "51-200",
				AssociatedMembers: "28",
				Verified:          true,
			},
			Ignore: true, // Recruiting agency
		},
		{
			Name:    "Harvey Nash",
			Website: "https://www.harveynash.co.uk/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                164131,
				IDs:               []int{26199, 28173, 164131, 2110112, 72217697},
				Alias:             "harvey-nash",
				Name:              "Harvey Nash",
				Followers:         "701K",
				Employees:         "5K-10K",
				AssociatedMembers: "1,302",
				Verified:          true,
			},
			Ignore: true, // Recruiting agency
		},
		{
			Name:    "Ocho",
			Website: "https://www.ochopeople.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:                89286373,
				IDs:               nil,
				Alias:             "ocho-people",
				Name:              "Ocho",
				Followers:         "36K",
				Employees:         "2-10",
				AssociatedMembers: "16",
				Verified:          false,
			},
			Ignore: true, // Recruiting agency
		},
	}
}
